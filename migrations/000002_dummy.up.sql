-- Insert dummy users
INSERT INTO users (name, email, password_hash) VALUES
('Alice Smith', 'alice@example.com', 'hashed_password_1'),
('Bob Johnson', 'bob@example.com', 'hashed_password_2');

-- Insert into foods
INSERT INTO foods (name, calories, protein, fat, carbs, fiber, sugar, sodium) VALUES
('Apple', 52, 0.3, 0.2, 14, 2.4, 10, 1),
('Banana', 89, 1.1, 0.3, 23, 2.6, 12, 1),
('Chicken Breast', 165, 31, 3.6, 0, 0, 0, 74),
('White Rice', 130, 2.7, 0.3, 28, 0.4, 0.1, 1),
('Broccoli', 55, 3.7, 0.6, 11, 3.8, 2.2, 33),
('Almonds', 579, 21, 50, 22, 12, 4.4, 1),
('Whole Milk', 61, 3.2, 3.3, 5, 0, 5, 44),
('Egg', 155, 13, 11, 1.1, 0, 1.1, 124),
('Salmon', 208, 20, 13, 0, 0, 0, 59),
('Carrot', 41, 0.9, 0.2, 10, 2.8, 4.7, 69);

-- Insert into food_units (assuming the food_ids from 1 to 10 for the inserted foods)
INSERT INTO food_units (food_id, unit, size_in_grams) VALUES
(1, 'piece', 182),   -- Apple average medium size
(1, 'gram', 1),
(2, 'piece', 118),   -- Banana average medium size
(2, 'gram', 1),
(3, 'gram', 1),
(3, 'piece', 120),   -- average chicken breast weight
(4, 'cup', 158),     -- cooked white rice per cup
(4, 'gram', 1),
(5, 'cup chopped', 91), -- broccoli chopped
(5, 'gram', 1),
(6, 'gram', 1),
(6, 'oz', 28.35),
(7, 'ml', 244),      -- 1 cup whole milk
(7, 'gram', 1),
(8, 'piece', 50),    -- 1 large egg approx 50g
(9, 'gram', 1),
(9, 'oz', 28.35),
(10, 'piece', 61),   -- 1 medium carrot approx 61g
(10, 'gram', 1);

-- Insert dummy food logs
INSERT INTO food_logs (
    user_id, food_id,
    food_unit_id, quantity,
    meal, timestamp) VALUES
(1, 1, 2, 100, 'snacks', '2025-07-03 15:30:00'),      -- 100 grams apple as snacks
(2, 2, 4, 150, 'breakfast', '2025-07-03 08:15:00'),   -- 150 grams banana at breakfast
(1, 3, 5, 180, 'lunch', '2025-07-03 13:00:00'),       -- 180 grams chicken breast for lunch
(2, 4, 7, 2, 'dinner', '2025-07-03 19:45:00'),        -- 2 cups cooked white rice at dinner
(1, 5, 9, 1, 'dinner', '2025-07-03 20:00:00'),        -- 1 cup chopped broccoli at dinner
(2, 1, 1, 1, 'breakfast', '2025-07-04 07:50:00'),     -- 1 piece apple at breakfast
(1, 2, 3, 2, 'snacks', '2025-07-04 16:00:00'),        -- 2 pieces banana as snacks
(2, 3, 6, 1, 'lunch', '2025-07-04 12:30:00'),         -- 1 chicken breast at lunch
(1, 4, 8, 250, 'dinner', '2025-07-04 18:30:00'),      -- 250 grams white rice at dinner
(2, 5, 10, 80, 'snacks', '2025-07-04 16:45:00');       -- 80 grams broccoli as snacks
