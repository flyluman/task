-- insert into restaurant
INSERT INTO restaurants (id, name, cash_balance) VALUES
(1, 'Ulu Ocean Grill and Sushi Lounge', 4483.84),
(2, 'Roma Ristorante', 1000),
(3, 'Naked City Pizza', 1200);

-- insert into menu_items
INSERT INTO menu_items (restaurant_id, dish_name, price) VALUES
(1, 'Postum cereal coffee', 13.88),
(1, 'Coffee Cocktail (Port Wine', 12.45),
(1, 'Sweet Virginia Pickles', 10.15),
(2, 'Olives', 13.18),
(3, 'Roast Young Turkey', 12.81);

-- insert into user
INSERT INTO users (id, name, cash_balance)
VALUES (1, 'Edith Johnson', 700.7);

-- insert into purchases
INSERT INTO purchases (user_id, restaurant_id, menu_item_id, amount)
VALUES
(1, 2, (SELECT id FROM menu_items WHERE restaurant_id=2 AND dish_name='Olives'), 13.18),
(1, 3, (SELECT id FROM menu_items WHERE restaurant_id=3 AND dish_name='Roast Young Turkey'), 12.81);
