package repository

import (
	"database/sql"
	"fmt"
	"time"

	"hungerycat-backend.com/main/services/models"
)

var DB *sql.DB

// Customer GET, POST, PUT And DELETE methods
func PostCustomer(name, email, password, phone_number, customer_id, profile_image, address, location string, created_date time.Time) (uint, error) {
	var id uint

	currentTime := time.Now()

	// Insert into customer table
	err := DB.QueryRow(
		"INSERT INTO customer(name, email, password, phone_number, customer_id, profile_image, address, location, created_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		name, email, password, phone_number, customer_id, profile_image, address, location, currentTime).Scan(&id)

	if err != nil {
		return 0, err
	}

	fmt.Println("Post Customer Successfully")

	return id, nil
}

func GetCustomer() ([]models.Customer, error) {
	// implement get all customer logic here
	query := "SELECT id, name, email, password, phone_number, customer_id, profile_image, address, location, created_date FROM customer"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []models.Customer

	for rows.Next() {
		var customer models.Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Password, &customer.PhoneNumber, &customer.CustomerID, &customer.ProfileImage, &customer.Address, &customer.Location, &customer.CreatedDate); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println("Get Customer Successfully")

	return customers, nil
}

func PutCustomer(id uint, name, email, password, phone_number, customer_id, profile_image, address, location string) error {

	// implement update logic here
	_, err := DB.Exec(
		"UPDATE customer SET id=$1, name=$2, email=$3, password=$4, phone_number=$5, profile_image=$6, address=$7, location=$8 WHERE customer_id=$9",
		id, name, email, password, phone_number, profile_image, address, location, customer_id)

	fmt.Println("Update Customer Successfully")

	return err
}

func DeleteCustomer(customer_id string) error {

	// implement delete logic here
	_, err := DB.Exec("DELETE FROM customer WHERE customer_id=$1", customer_id)

	fmt.Println("Delete Customer Successfully")

	return err
}

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

func PutAdmin(id uint, username, email, password, phone_number, admin_id, profile_image string, last_signin time.Time) error {

	// implement update logic here
	_, err := DB.Exec(
		"UPDATE admin SET id=$1, username=$2, email=$3, password=$4, phone_number=$5, profile_image=$6, last_signin=$7 WHERE admin_id=$8",
		id, username, email, password, phone_number, profile_image, time.Now(), admin_id)

	fmt.Println("Update Admin Successfully")

	return err
}

func DeleteAdmin(admin_id string) error {

	// implement delete logic here
	_, err := DB.Exec("DELETE FROM admin WHERE admin_id=$1", admin_id)

	fmt.Println("Delete Admin Successfully")

	return err
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

// Food GET ,POST, PUT and DELETE
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

func PutFood(id uint, name, description, category, product_id, image, hotel_name, hotel_id string, price int, stock bool, updated_date time.Time) error {

	// implement update logic here
	_, err := DB.Exec(
		"UPDATE food SET id=$1, name=$2, description=$3, category=$4, price=$5, stock=$6, image=$7, hotel_name=$8, hotel_id=$9, updated_date=$10 WHERE product_id=$11",
		id, name, description, category, price, stock, image, hotel_name, hotel_id, time.Now(), product_id)

	fmt.Println("Update Food Successfully")

	return err
}

func DeleteFood(product_id string) error {

	// implement delete logic here
	_, err := DB.Exec("DELETE FROM food WHERE product_id=$1", product_id)

	fmt.Println("Delete Food Successfully")

	return err
}

// Restaurant GET, POST, PUT and DELETE
func PostRestaurant(hotel_id, hotel_name, description, address, location, phone_number, email, website, menu, profile_image, open_time, close_time string, ratings int, created_date, updated_date time.Time) (uint, error) {
	var id uint

	currentTime := time.Now()

	// Insert into Restaurant table
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

	// implement get all Restaurant logic here
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

func PutRestaurant(id uint, hotel_id, hotel_name, description, address, location, phone_number, email, website, menu, profile_image, open_time, close_time string, ratings int, updated_date time.Time) error {

	// implement update logic here
	_, err := DB.Exec(
		"UPDATE restaurant SET id=$1, hotel_name=$2, description=$3, address=$4, location=$5, phone_number=$6, email=$7, website=$8, menu=$9, profile_image=$10, open_time=$11, close_time=$12, ratings=$13, updated_date=$14 WHERE hotel_id=$15",
		id, hotel_name, description, address, location, phone_number, email, website, menu, profile_image, open_time, close_time, ratings, time.Now(), hotel_id)

	fmt.Println("Update Restaurant Successfully")

	return err
}

func DeleteRestaurant(hotel_id string) error {

	// implement delete logic here
	_, err := DB.Exec("DELETE FROM restaurant WHERE hotel_id=$1", hotel_id)

	fmt.Println("Delete Restaurant Successfully")

	return err
}

// Order GET, POST, PUT and DELETE
func PostOrder(order_id, customer_id, product_id, hotel_id, order_address, order_location, order_status string, quantity int, order_time, created_date, updated_date time.Time) (uint, error) {
	var id uint

	currentTime := time.Now()

	// Insert into Order table
	err := DB.QueryRow(
		"INSERT INTO orders(order_id, customer_id, product_id, quantity, hotel_id, order_address, order_location, order_time, order_status, created_date, updated_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8 ,$9, $10, $11) RETURNING id",
		order_id, customer_id, product_id, quantity, hotel_id, order_address, order_location, currentTime, order_status, currentTime, currentTime).Scan(&id)

	if err != nil {
		return 0, err
	}

	fmt.Println("Post Order Successfully")

	return id, nil
}

func GetOrder() ([]models.Order, error) {
	// implement get all orders logic here
	query := "SELECT id, order_id, customer_id, product_id, quantity, hotel_id, order_address, order_location, order_time, order_status, created_date, updated_date FROM orders"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order

	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.OrderID, &order.CustomerID, &order.ProductId, &order.Quantity, &order.HotelId, &order.OrderAddress, &order.OrderLocation, &order.OrderTime, &order.OrderStatus, &order.CreatedDate, &order.UpdatedDate); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println("Get Order Successfully")

	return orders, nil
}

func PutOrder(id uint, order_id, customer_id, product_id, hotel_id, order_address, order_location, order_status string, quantity int, order_time, updated_date time.Time) error {

	// implement update logic here
	_, err := DB.Exec(
		"UPDATE orders SET id=$1, customer_id=$2, product_id=$3, hotel_id=$4, order_address=$5, order_location=$6, order_status=$7, quantity=$8, order_time=$9, updated_date=$10 WHERE order_id=$11",
		id, customer_id, product_id, hotel_id, order_address, order_location, order_status, quantity, time.Now(), time.Now(), order_id)

	fmt.Println("Update Order Successfully")

	return err
}

func DeleteOrder(order_id string) error {

	// implement delete logic here
	_, err := DB.Exec("DELETE FROM orders WHERE order_id=$1", order_id)

	fmt.Println("Delete Order Successfully")

	return err
}
