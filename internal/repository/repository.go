package repository

import (
	"database/sql"
	"errors"
	"task/model"
	"time"
)

type UserRepository interface {
	QueryRestaurants(userID int) ([]model.Restaurant, error)
	PurchaseTX(userID, menuItemID int) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) QueryRestaurants(userID int) ([]model.Restaurant, error) {
	rows, err := r.db.Query(`
		SELECT DISTINCT r.id, r.name, r.cash_balance
		FROM restaurants r
		JOIN purchases p ON r.id = p.restaurant_id
		WHERE p.user_id = $1`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var restaurants []model.Restaurant
	for rows.Next() {
		var r model.Restaurant
		if err := rows.Scan(&r.ID, &r.Name, &r.CashBalance); err != nil {
			return nil, err
		}
		restaurants = append(restaurants, r)
	}

	return restaurants, nil
}

func (r *userRepo) PurchaseTX(userID, menuItemID int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// Get price and restaurant_id
	var price float32
	var restaurantID int
	err = tx.QueryRow(`
		SELECT price, restaurant_id
		FROM menu_items
		WHERE id = $1
	`, menuItemID).Scan(&price, &restaurantID)

	if err != nil {
		return err
	}

	// Get user balance
	var userBalance float32
	err = tx.QueryRow(`
		SELECT cash_balance
		FROM users
		WHERE id = $1
	`, userID).Scan(&userBalance)

	if err != nil {
		return err
	}

	// Check balance
	if userBalance < price {
		return errors.New("repo: insufficient balance")
	}

	// Deduct from user
	if _, err = tx.Exec(`
		UPDATE users
		SET cash_balance = cash_balance - $1
		WHERE id = $2
	`, price, userID); err != nil {
		return err
	}

	// Add to restaurant
	if _, err = tx.Exec(`
		UPDATE restaurants
		SET cash_balance = cash_balance + $1
		WHERE id = $2
	`, price, restaurantID); err != nil {
		return err
	}

	// Log the purchase
	if _, err = tx.Exec(`
		INSERT INTO purchases (user_id, restaurant_id, menu_item_id, amount, purchased_at)
		VALUES ($1, $2, $3, $4, $5)
	`, userID, restaurantID, menuItemID, price, time.Now()); err != nil {
		return err
	}

	// Commit
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
