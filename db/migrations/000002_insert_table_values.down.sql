-- delete purchases
DELETE FROM purchases
WHERE user_id = 1
  AND (restaurant_id IN (2, 3));

-- delete users
DELETE FROM users
WHERE id = 1 AND name = 'Edith Johnson';

-- delete menu_items
DELETE FROM menu_items
WHERE (restaurant_id = 1 AND dish_name IN ('Postum cereal coffee', 'Coffee Cocktail (Port Wine', 'Sweet Virginia Pickles'))
   OR (restaurant_id = 2 AND dish_name = 'Olives')
   OR (restaurant_id = 3 AND dish_name = 'Roast Young Turkey');

-- delete restaurants
DELETE FROM restaurants
WHERE id IN (1, 2, 3);
