# MyDiet

MyDiet is a terminal-based application for tracking your daily food intake. It provides a simple and efficient way to log your meals and monitor your nutritional information.

## Features

*   **Food Logging:** Easily log your meals (breakfast, lunch, dinner, and snacks).
*   **Nutritional Information:** View detailed nutritional information for each food item.
*   **Search:** Quickly find food items from the database.
*   **Simple Interface:** A clean and intuitive terminal-based user interface.

## Technologies Used

*   **Go:** The application is written in the Go programming language.
*   **Bubble Tea:** A Go library for building terminal-based user interfaces.
*   **SQLite:** A lightweight, serverless database for storing food and nutrition data.

## Getting Started

### Prerequisites

*   Go 1.23.0 or higher
*   A C compiler (for the SQLite driver)

### Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/b9uu/mydiet.git
    cd mydiet
    ```

2.  **Build the application:**

    ```bash
    go build -o mydiet ./cmd/main.go
    ```

3.  **Run the application:**

    ```bash
    ./mydiet
    ```

## Usage

When you run the application, you will be presented with a view of your daily food log. You can use the following keys to navigate and interact with the application:

*   **Ctrl+C:** Quit the application.
*   **Tab:** Switch between different sections of the interface.
*   **Enter:** Select an item or confirm an action.

## Database

The application uses a SQLite database named `nutrition.db` to store food and nutrition data. The database schema is managed using SQL migration files located in the `migrations` directory.
