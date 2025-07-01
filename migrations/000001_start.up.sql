-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL
);

-- Create foods table
CREATE TABLE IF NOT EXISTS foods (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    calories REAL NOT NULL,
    protein REAL NOT NULL,
    fat REAL NOT NULL,
    carbs REAL NOT NULL,
    fiber REAL NOT NULL,
    sugar REAL NOT NULL,
    sodium REAL NOT NULL
);

CREATE TABLE IF NOT EXISTS food_units (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    food_id INTEGER NOT NULL,
    unit TEXT NOT NULL,       -- e.g., "grams", "piece", "oz"
    size_in_grams REAL NOT NULL,  -- how much one unit weighs in grams
    FOREIGN KEY (food_id) REFERENCES foods(id)
);

-- Create food_logs table
CREATE TABLE IF NOT EXISTS food_logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    food_id INTEGER NOT NULL,
    food_unit_id INTEGER NOT NULL,   -- Reference to food_units table
    meal TEXT NOT NULL,
    quantity REAL NOT NULL,           -- How many units consumed
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (food_id) REFERENCES foods(id),
    FOREIGN KEY (food_unit_id) REFERENCES food_units(id)
);
