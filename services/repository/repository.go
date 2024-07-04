package repository

import (
	"database/sql"
	"fmt"
	"time"

	"hungerycat-backend.com/main/services/models"
)

var DB *sql.DB

// Admin GET, POST, PUT and DELETE methods
func PostAdmin(username, email, password, phone_number, admin_id, profile_image string, created_at, last_signin time.Time) (uint, error) {
	var id uint

	currentTime := time.Now()

	// Insert into admin table
	err := DB.QueryRow(
		"INSERT INTO admin(username, email, password, phone_number, admin_id, profile_image, created_at, last_signin) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		username, email, password, phone_number, admin_id, profile_image, currentTime, currentTime).Scan(&id)

	if err != nil {
		return 0, err
	}

	fmt.Println("Post Admin Successfully")

	return id, nil
}

func GetAdmin() ([]models.Admin, error) {
	// implement get all orders logic here
	query := "SELECT id, username, email, password, phone_number, admin_id, profile_image, created_at, last_signin FROM admin"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var admins []models.Admin

	for rows.Next() {
		var admin models.Admin
		if err := rows.Scan(&admin.ID, &admin.Username, &admin.Email, &admin.Password, &admin.PhoneNumber, &admin.AdminId, &admin.ProfileImage, &admin.CreatedAt, &admin.LastSingIn); err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println("Get Admin Successfully")

	return admins, nil
}

func CheckEmailAndPassword(email, password string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM admin WHERE email = $1 AND password = $2);"
	err := DB.QueryRow(query, email, password).Scan(&exists)
	if err != nil {
		return false, err
	}

	fmt.Println("Checking email and password")

	return exists, nil
}

// Food GET And POST
func PostFood(name, description, category, product_id, image, hotel_name, hotel_id string, price int, stock bool, created_date, updated_date time.Time) (uint, error) {
	var id uint

	currentTime := time.Now()

	// Insert into Food table
	err := DB.QueryRow(
		"INSERT INTO food(name, description, category, product_id, price, stock, image, hotel_name, hotel_id, created_date, updated_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8 ,$9, $10, $11) RETURNING id",
		name, description, category, product_id, price, stock, image, hotel_name, hotel_id, currentTime, currentTime).Scan(&id)

	if err != nil {
		return 0, err
	}

	fmt.Println("Post food Successfully")

	return id, nil
}

func GetFood() ([]models.Food, error) {
	// implement get all orders logic here
	query := "SELECT id, name, description, category, product_id, price, stock, image, hotel_name, hotel_id, created_date, updated_date FROM food"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foods []models.Food

	for rows.Next() {
		var food models.Food
		if err := rows.Scan(&food.ID, &food.Name, &food.Description, &food.Category, &food.ProductId, &food.Price, &food.Stock, &food.Image, &food.HotelName, &food.HotelId, &food.CreatedDate, &food.UpdatedDate); err != nil {
			return nil, err
		}
		foods = append(foods, food)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println("Get Food Successfully")

	return foods, nil
}

// Restaurant GET And POST
func PostRestaurant(hotel_id, hotel_name, description, address, location, phone_number, email, website, menu, profile_image, open_time, close_time string, ratings int, created_date, updated_date time.Time) (uint, error) {
	var id uint

	currentTime := time.Now()

	// Insert into Food table
	err := DB.QueryRow(
		"INSERT INTO restaurant(hotel_id, hotel_name, description, address, location, phone_number, email, website, menu, profile_image, open_time, close_time, ratings, created_date, updated_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8 ,$9, $10, $11, $12, $13, $14, $15) RETURNING id",
		hotel_id, hotel_name, description, address, location, phone_number, email, website, menu, profile_image, open_time, close_time, ratings, currentTime, currentTime).Scan(&id)

	if err != nil {
		return 0, err
	}

	fmt.Println("Post Restaurant Successfully")

	return id, nil
}

func GetRestaurant() ([]models.Restaurant, error) {
	// implement get all orders logic here
	query := "SELECT id, hotel_id, hotel_name, description, address, location, phone_number, email, website, menu, profile_image, open_time, close_time, ratings, created_date, updated_date FROM restaurant"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var restaurants []models.Restaurant

	for rows.Next() {
		var restaurant models.Restaurant
		if err := rows.Scan(&restaurant.ID, &restaurant.HotelId, &restaurant.HotelName, &restaurant.Description, &restaurant.Address, &restaurant.Location, &restaurant.PhoneNumber, &restaurant.Email, &restaurant.Website, &restaurant.Menu, &restaurant.ProfileImage, &restaurant.OpenTime, &restaurant.CloseTime, &restaurant.Ratings, &restaurant.CreatedDate, &restaurant.UpdatedDate); err != nil {
			return nil, err
		}
		restaurants = append(restaurants, restaurant)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println("Get Restaurants Successfully")

	return restaurants, nil
}
