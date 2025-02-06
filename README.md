# Finance Control API (Open Source)

Financial control project developed in Go, using the Gin framework and PostgreSQL as the database.  
This project is under development.

## Authors

- **Andressa Silva** - [AuroraDark](https://github.com/auroradark)

## Current Features

- **REST API** with endpoints for:
  - **Transactions**
    - List transactions with date filters (using the `transaction_date` column)  
      `GET /api/transactions?start=YYYY-MM-DD&end=YYYY-MM-DD`
    - Create new transactions  
      `POST /api/transactions`
    - Update transactions  
      `PUT /api/transactions/:id`
    - Delete transactions  
      `DELETE /api/transactions/:id`
    - Support for two types of transactions: `income` and `expense`
  - **Categories**
    - List categories  
      `GET /api/categories`
    - Get details of a category  
      `GET /api/categories/:id`
    - Create new categories  
      `POST /api/categories`
    - Update categories  
      `PUT /api/categories/:id`
    - Delete categories  
      `DELETE /api/categories/:id`
  - **Users and Authentication**
    - User CRUD  
      `GET /api/users`, `GET /api/users/:id`, `POST /api/users`, `PUT /api/users/:id`, `DELETE /api/users/:id`
    - Login  
      `POST /api/login`
    - Logout  
      `POST /api/logout`
    - Session management (retrieve the logged-in user via `GET /api/session`)
  - **Investments**
    - List investments  
      `GET /api/investments`
    - Get details of an investment  
      `GET /api/investments/:id`
    - Create new investments  
      `POST /api/investments`
    - Update investments  
      `PUT /api/investments/:id`
    - Delete investments  
      `DELETE /api/investments/:id`
  - **Investment Movements**
    - List movements  
      `GET /api/investment_movements`
    - Get details of a movement  
      `GET /api/investment_movements/:id`
    - Create new movements  
      `POST /api/investment_movements`
    - Update movements  
      `PUT /api/investment_movements/:id`
    - Delete movements  
      `DELETE /api/investment_movements/:id`
- **Data Validation** using Gin's binding tags.
- **PostgreSQL Integration** via GORM.
- **Code Organization** using the Repository Pattern to decouple data access layers.
- **Session Management** via cookies using `gin-contrib/sessions`.
- **Automatic Documentation** generated with Swagger (Swag).

## Prerequisites

- [Go](https://golang.org/dl/) 1.16 or higher
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/finance-control.git
   cd finance-control


2. Install the modules with:

`go mod download`

3. In your PostgreSQL, create a database called finance_db (or adjust the name in the configuration file).

3. Rename the file config.json.example to config.json and adjust your connection values.

4. Run the project using:

`go run main.go`