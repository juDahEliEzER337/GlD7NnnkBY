// 代码生成时间: 2025-08-18 09:39:19
package main

import (
    "fmt"
    "gorm.io/driver/sqlite" // Assuming using SQLite for the database
    "gorm.io/gorm"
)

// CartItem represents an item in the shopping cart
type CartItem struct {
    gorm.Model
    ProductID uint
    Quantity  int
}

// Cart represents a shopping cart
type Cart struct {
    gorm.Model
    Items       []CartItem
}

// Product represents a product that can be added to the cart
type Product struct {
    gorm.Model
    Name  string
    Price float64
}

// DatabaseClient represents a connection to the database
type DatabaseClient struct {
    db *gorm.DB
}

// NewDatabaseClient creates a new database client with a connection to the SQLite database
func NewDatabaseClient() (*DatabaseClient, error) {
    db, err := gorm.Open(sqlite.Open("shoppingCart.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    // Migrate the schema
    db.AutoMigrate(&Cart{}, &CartItem{}, &Product{})
    return &DatabaseClient{db: db}, nil
}

// AddItemToCart adds an item to the cart
func (dc *DatabaseClient) AddItemToCart(cartID uint, productID uint, quantity int) error {
    var cart Cart
    if err := dc.db.First(&cart, cartID).Error; err != nil {
        return err
    }
    // Check if the item already exists in the cart
    for _, item := range cart.Items {
        if item.ProductID == productID {
            item.Quantity += quantity
            if err := dc.db.Save(&item).Error; err != nil {
                return err
            }
            return nil
        }
    }
    // Add a new item to the cart
    newItem := CartItem{ProductID: productID, Quantity: quantity}
    if err := dc.db.Create(&newItem).Error; err != nil {
        return err
    }
    // Add the new item to the cart
    cart.Items = append(cart.Items, newItem)
    if err := dc.db.Model(&cart).Association("Items").Append(&newItem).Error; err != nil {
        return err
    }
    return nil
}

// RemoveItemFromCart removes an item from the cart
func (dc *DatabaseClient) RemoveItemFromCart(cartID uint, productID uint) error {
    var cart Cart
    if err := dc.db.First(&cart, cartID).Error; err != nil {
        return err
    }
    for i, item := range cart.Items {
        if item.ProductID == productID {
            if err := dc.db.Delete(&item).Error; err != nil {
                return err
            }
            cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("product with id %d not found in cart", productID)
}

// GetCart returns the cart with its items
func (dc *DatabaseClient) GetCart(cartID uint) (*Cart, error)
func (dc *DatabaseClient) GetCart(cartID uint) (*Cart, error) {
    var cart Cart
    if err := dc.db.Preload("Items").Preload("Items.Product").First(&cart, cartID).Error; err != nil {
        return nil, err
    }
    return &cart, nil
}

func main() {
    dbClient, err := NewDatabaseClient()
    if err != nil {
        panic("failed to connect to database: " + err.Error())
    }
    // Example usage
    // Add items to cart
    if err := dbClient.AddItemToCart(1, 1, 2); err != nil {
        panic(err)
    }
    // Get cart
    cart, err := dbClient.GetCart(1)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Cart: %+v
", cart)
    // Remove item from cart
    if err := dbClient.RemoveItemFromCart(1, 1); err != nil {
        panic(err)
    }
}
