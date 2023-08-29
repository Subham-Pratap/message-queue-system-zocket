// models.go
package models

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID                 uint `gorm:"primaryKey"`
	UserID             uint
	ProductName        string
	ProductDescription string   `gorm:"type:text"`
	ProductImages      []string `gorm:"type:json"`
	ProductPrice       float64
}

// package models

// import "time"

// type Product struct {
// 	ProductID               int       `json:"product_id"`
// 	ProductName             string    `json:"product_name"`
// 	ProductDescription      string    `json:"product_description"`
// 	ProductImages           []string  `json:"product_images"`
// 	ProductPrice            float64   `json:"product_price"`
// 	CompressedProductImages []string  `json:"compressed_product_images"`
// 	CreatedAt               time.Time `json:"created_at"`
// 	UpdatedAt               time.Time `json:"updated_at"`
// }

// package models

// import "time"

// type User struct {
// 	UserID    int       `json:"user_id"`
// 	Name      string    `json:"name"`
// 	Mobile    string    `json:"mobile"`
// 	Latitude  float64   `json:"latitude"`
// 	Longitude float64   `json:"longitude"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }
