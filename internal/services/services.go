package services

import (
	"errors"
	"fmt"
	"task/internal/db"
	"time"
)

type Restaurant struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	CashBalance float32 `json:"cash_balance"`
}

func GetUserRestaurants(userID int) ([]Restaurant, error) {
	rows, err := db.DB.Query(`
	SELECT DISTINCT r.id, r.name, r.cash_balance
	FROM restaurants r
	JOIN purchases p ON r.id = p.restaurant_id
	WHERE p.user_id = $1`, userID)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	var restaurants []Restaurant

	for rows.Next() {
		var r Restaurant
		if err := rows.Scan(&r.ID, &r.Name, &r.CashBalance); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		restaurants = append(restaurants, r)
	}

	return restaurants, nil
}

func PurchaseMenuItem(userID, menuItemID int) error {
	commit := false
	tx, err := db.DB.Begin()

	if err != nil {
		fmt.Println("services: Error intializing transaction")
		return err
	}

	defer func() {
		if !commit {
			tx.Rollback()
		}
	}()

	var price float32
	var restaurantID int

	err = tx.QueryRow("SELECT price, restaurant_id FROM menu_items WHERE id=$1", menuItemID).Scan(&price, &restaurantID)

	if err != nil {
		fmt.Println("services: Error quering database")
		return err
	}

	var userBalance float32

	err = tx.QueryRow("SELECT cash_balance FROM users WHERE id=$1", userID).Scan(&userBalance)

	if err != nil {
		fmt.Println("services: Error quering database")
		return err
	}

	if userBalance < price {
		return errors.New("insufficient balance")
	}

	_, err = tx.Exec("UPDATE users SET cash_balance = cash_balance - $1 WHERE id=$2", price, userID)

	if err != nil {
		fmt.Println("services: Error quering database")
		return err
	}

	_, err = tx.Exec("UPDATE restaurants SET cash_balance = cash_balance + $1 WHERE id=$2", price, restaurantID)

	if err != nil {
		fmt.Println("services: Error quering database")
		return err
	}

	_, err = tx.Exec(`
        INSERT INTO purchases (user_id, restaurant_id, menu_item_id, amount, purchased_at)
        VALUES ($1, $2, $3, $4, $5)
    `, userID, restaurantID, menuItemID, price, time.Now())

	if err != nil {
		fmt.Println("services: Error quering database")
		return err
	}

	err = tx.Commit()

	if err != nil {
		fmt.Println("services: Error quering database")
		return err
	} else {
		commit = true
		return nil
	}
}
