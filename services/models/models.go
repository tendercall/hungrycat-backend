package models

import "time"

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
	Stock       bool      `json:"stock"`
	Image       string    `json:"image"`
	HotelName   string    `json:"hotel_name"`
	HotelId     string    `json:"hotel_id"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
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
