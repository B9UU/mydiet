-- Insert dummy users
INSERT INTO users (name, email, password_hash) VALUES
('Alice Smith', 'alice@example.com', 'hashed_password_1'),
('Bob Johnson', 'bob@example.com', 'hashed_password_2');

INSERT INTO foods (name, serving_size, serving_unit, calories, protein, fat, carbs, fiber, sugar, sodium) VALUES
('Banana', 118, 'gram', 105, 1.3, 0.4, 27, 3.1, 14, 1),
('Chicken Breast', 120, 'gram', 198, 37, 4.5, 0, 0, 0, 74),
('White Bread Slice', 25, 'slice', 67, 2, 1, 12, 1, 1.4, 127),
('Apple', 182, 'gram', 95, 0.5, 0.3, 25, 4.4, 19, 2),
('Boiled Egg', 50, 'piece', 78, 6.3, 5.3, 0.6, 0, 0.6, 62),
('Brown Rice', 195, 'gram', 216, 5, 1.8, 45, 3.5, 1, 10),
('Broccoli', 91, 'gram', 31, 2.5, 0.3, 6, 2.4, 1.5, 30),
('Almonds', 28, 'gram', 164, 6, 14, 6, 3.5, 1.2, 0),
('Greek Yogurt (Plain)', 150, 'gram', 100, 17, 0.7, 6, 0, 6, 50),
('Oatmeal (Cooked)', 234, 'gram', 154, 6, 3.2, 27, 4, 1.1, 2),
('Carrots', 61, 'gram', 25, 0.6, 0.1, 6, 1.7, 2.9, 42),
('Peanut Butter', 32, 'gram', 188, 8, 16, 6, 2, 3, 147),
('Avocado', 150, 'gram', 240, 3, 22, 12, 10, 1, 10),
('Salmon (Cooked)', 100, 'gram', 206, 22, 13, 0, 0, 0, 50),
('Potato (Baked)', 173, 'gram', 161, 4.3, 0.2, 37, 3.8, 1.7, 17),
('Lentils (Cooked)', 198, 'gram', 230, 18, 0.8, 40, 16, 2, 4),
('Spinach', 30, 'gram', 7, 0.9, 0.1, 1, 0.7, 0.1, 24),
('Tuna (canned in water)', 85, 'gram', 100, 22, 1, 0, 0, 0, 300),
('Quinoa (Cooked)', 185, 'gram', 222, 8, 3.6, 39, 5, 1, 13),
('Milk (Whole)', 244, 'gram', 149, 8, 8, 12, 0, 12, 105),
('Cheddar Cheese', 28, 'gram', 113, 7, 9, 1, 0, 0.5, 174),
('Olive Oil', 15, 'gram', 119, 0, 14, 0, 0, 0, 0),
('Tomato', 123, 'gram', 22, 1.1, 0.2, 4.8, 1.5, 3.2, 6),
('Cucumber', 104, 'gram', 16, 0.7, 0.1, 3.8, 0.5, 1.7, 2),
('Bell Pepper (Red)', 92, 'gram', 24, 0.9, 0.2, 6, 2, 4, 3),
('Hummus', 30, 'gram', 82, 2.4, 6, 6, 2, 0.1, 148),
('Whole Wheat Bread', 32, 'slice', 81, 4, 1.1, 14, 2, 1.5, 146),
('Black Beans (Cooked)', 172, 'gram', 227, 15, 0.9, 41, 15, 0, 1),
('Strawberries', 152, 'gram', 49, 1, 0.5, 12, 3, 7, 1),
('Blueberries', 148, 'gram', 84, 1.1, 0.5, 21, 3.6, 15, 1),
('Cottage Cheese (2%)', 113, 'gram', 92, 12, 2.5, 5, 0, 4, 360),
('Tofu (Firm)', 81, 'gram', 145, 16, 9, 3, 2, 0.5, 7),
('Shrimp (Cooked)', 85, 'gram', 84, 20, 0.3, 0, 0, 0, 111),
('Sweet Potato (Baked)', 180, 'gram', 162, 3.6, 0.3, 37, 6, 11, 99),
('Ground Beef (90% lean, cooked)', 85, 'gram', 184, 22, 10, 0, 0, 0, 76),
('Corn on the Cob', 143, 'gram', 127, 4.7, 1.9, 29, 3.9, 8.9, 21);

-- Insert dummy food logs
INSERT INTO food_logs (user_id, food_id, quantity, unit, timestamp) VALUES
(1, 1, 150, 'gram', '2025-06-30 08:00:00'),
(1, 2, 1, 'serving', '2025-06-30 12:30:00'),
(2, 3, 3, 'slice', '2025-06-30 07:30:00');
