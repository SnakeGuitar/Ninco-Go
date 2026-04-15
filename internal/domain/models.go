package domain

import (
	"time"
)

// Account represents a user account for authentication
type Account struct {
	ID        int       `json:"id" db:"account_id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"` // Never export password in JSON
	Role      Role      `json:"role" db:"role"`
	State     State     `json:"state" db:"state"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Store represents a physical branch/store
type Store struct {
	ID        int       `json:"id" db:"store_id"`
	Name      string    `json:"name" db:"name"`
	Address   string    `json:"address" db:"address"`
	Phone     string    `json:"phone" db:"phone"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Employee represents a person working in a store
type Employee struct {
	ID        int       `json:"id" db:"employee_id"`
	Email     string    `json:"email" db:"email"` // Acts as FK to Account
	Name      string    `json:"name" db:"name"`
	LastName  string    `json:"last_name" db:"last_name"`
	StoreID   *int      `json:"store_id,omitempty" db:"store_id"` // Pointer because it can be NULL in SQL
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Navigation property (Optional, depending on your ORM/Query builder)
	Account *Account `json:"account,omitempty" db:"-"`
}

// Access represents login/logout logs for employees
type Access struct {
	ID         int       `json:"id" db:"access_id"`
	EmployeeID int       `json:"employee_id" db:"employee_id"`
	Action     Action    `json:"action" db:"action"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// Product represents an item for sale
type Product struct {
	ID          int       `json:"id" db:"product_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Brand       string    `json:"brand" db:"brand"`
	Price       float64   `json:"price" db:"price"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// Stock represents the quantity of a product in a specific store (Composite PK)
type Stock struct {
	ProductID int       `json:"product_id" db:"product_id"`
	StoreID   int       `json:"store_id" db:"store_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Navigation property
	Product *Product `json:"product,omitempty" db:"-"`
}

// Invoice represents a sale transaction header
type Invoice struct {
	ID         int       `json:"id" db:"invoice_id"`
	StoreID    int       `json:"store_id" db:"store_id"`
	NameClient string    `json:"name_client" db:"name_client"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`

	// Navigation property (Replaces the non-existent Detail struct)
	Sales []Sale `json:"sales,omitempty" db:"-"`
}

// Sale represents the line items / transactional details of an Invoice
type Sale struct {
	ID         int       `json:"sale_id" db:"sale_id"`
	EmployeeID int       `json:"employee_id" db:"employee_id"`
	InvoiceID  int       `json:"invoice_id" db:"invoice_id"`
	ProductID  int       `json:"product_id" db:"product_id"`
	StoreID    int       `json:"store_id" db:"store_id"`
	Amount     int       `json:"amount" db:"amount"`
	Price      float64   `json:"price" db:"price"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// PendingRegistration represents a user waiting for validation
type PendingRegistration struct {
	ID        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Pin       string    `json:"pin" db:"pin"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Password  string    `json:"-" db:"password"`
	Role      Role      `json:"role" db:"role"`
}

// Session represents an active employee session
type Session struct {
	TokenID    string    `json:"token_id" db:"token_id"` // PK is a string (VARCHAR) in SQL
	EmployeeID int       `json:"employee_id" db:"employee_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	ExpiresAt  time.Time `json:"expires_at" db:"expires_at"`
}
