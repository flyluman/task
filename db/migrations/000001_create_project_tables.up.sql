-- create users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    cash_balance NUMERIC NOT NULL
);

-- create restaurants table
CREATE TABLE IF NOT EXISTS restaurants (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    cash_balance NUMERIC NOT NULL
);

-- create menu_items table
CREATE TABLE IF NOT EXISTS menu_items (
    id SERIAL PRIMARY KEY,
    restaurant_id INT REFERENCES restaurants(id) ON DELETE CASCADE,
    dish_name TEXT NOT NULL,
    price NUMERIC NOT NULL
);

-- create purchases table
CREATE TABLE IF NOT EXISTS purchases (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    restaurant_id INT REFERENCES restaurants(id) ON DELETE CASCADE,
    menu_item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
    amount NUMERIC NOT NULL,
    purchased_at TIMESTAMP DEFAULT NOW()
);
