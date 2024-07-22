package models

import "time"

type Customer struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	PhoneNumber  string    `json:"phone_number"`
	CustomerID   string    `json:"customer_id"`
	ProfileImage string    `json:"profile_image"`
	Address      string    `json:"address"`
	Location     string    `json:"location"`
	CreatedDate  time.Time `json:"created_date"`
}

type Admin struct {
	ID           uint      `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	PhoneNumber  string    `json:"phone_number"`
	AdminId      string    `json:"admin_id"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
	LastSingIn   time.Time `json:"last_singin"`
}

type Food struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	ProductId   string    `json:"product_id"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Image       string    `json:"image"`
	HotelName   string    `json:"hotel_name"`
	HotelId     string    `json:"hotel_id"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	Offer       int       `json:"offer"`
}

type Restaurant struct {
	ID           uint      `json:"id"`
	HotelId      string    `json:"hotel_id"`
	HotelName    string    `json:"Hotel_name"`
	Description  string    `json:"description"`
	Address      string    `json:"address"`
	Location     string    `json:"location"`
	PhoneNumber  string    `json:"phone_number"`
	Email        string    `json:"email"`
	Website      string    `json:"website"`
	Menu         string    `json:"menu"`
	ProfileImage string    `json:"profile_image"`
	OpenTime     string    `json:"open_time"`
	CloseTime    string    `json:"close_time"`
	Ratings      int       `json:"ratings"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedDate  time.Time `json:"updated_date"`
}

type Order struct {
	ID            uint      `json:"id"`
	OrderID       string    `json:"order_id"`
	CustomerID    string    `json:"customer_id"`
	ProductId     string    `json:"product_id"`
	Quantity      int       `json:"quantity"`
	HotelId       string    `json:"hotel_id"`
	OrderAddress  string    `json:"order_address"`
	OrderLocation string    `json:"order_location"`
	OrderTime     time.Time `json:"order_time"`
	OrderStatus   string    `json:"order_status"`
	CreatedDate   time.Time `json:"created_date"`
	UpdatedDate   time.Time `json:"updated_date"`
}

type Test struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DeliveryBoy struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	PhoneNumber  string    `json:"phone_number"`
	DbID         string    `json:"db_id"`
	Location     string    `json:"location"`
	Latitude     string    `json:"latitude"`
	Longitude    string    `json:"longitude"`
	TotalPayment int       `json:"total_payment"`
	TotalOrder   int       `json:"total_orders"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedDate  time.Time `json:"updated_date"`
}

type Category struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	CategoryID  string    `json:"category_id"`
	Icon        string    `json:"icon"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

type Banner struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	BannerID    string    `json:"banner_id"`
	Image       string    `json:"image"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

type Offer struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Offer       int       `json:"offer"`
	OfferID     string    `json:"offer_id"`
	Image       string    `json:"image"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

type Details struct {
	ID           uint      `json:"id"`
	Address      string    `json:"address"`
	BuildingType string    `json:"building_type"`
	CustomerID   string    `json:"customer_id"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedDate  time.Time `json:"updated_date"`
}
