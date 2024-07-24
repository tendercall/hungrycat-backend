package repository

import (
	"database/sql"
	"fmt"
	"log"
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

func GetCustomerByID(id uint) (*models.Customer, error) {
	var customer models.Customer

	err := DB.QueryRow("SELECT id, name, email, password, phone_number, customer_id, profile_image, address, location, created_date FROM customer WHERE id = $1", id).Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Password, &customer.PhoneNumber, &customer.CustomerID, &customer.ProfileImage, &customer.Address, &customer.Location, &customer.CreatedDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no customer found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}

	return &customer, nil
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

func GetAdminByID(id uint) (*models.Admin, error) {
	var admin models.Admin

	err := DB.QueryRow("SELECT id, username, email, password, phone_number, admin_id, profile_image, created_at, last_signin FROM admin WHERE id = $1", id).Scan(&admin.ID, &admin.Username, &admin.Email, &admin.Password, &admin.PhoneNumber, &admin.AdminId, &admin.ProfileImage, &admin.CreatedAt, &admin.LastSingIn)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no admin found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}

	return &admin, nil
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
func PostFood(name, description, category, productID, image, hotelName, hotelID string, price, stock, offer int, createdDate, updatedDate time.Time) (uint, error) {
	var id uint
	currentTime := time.Now()

	err := DB.QueryRow(`
		INSERT INTO food(name, description, category, product_id, price, stock, image, hotel_name, hotel_id, created_date, updated_date, offer)VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`, name, description, category, productID, price, stock, image, hotelName, hotelID, currentTime, currentTime, offer).Scan(&id)

	if err != nil {
		return 0, err
	}

	fmt.Println("Successfully inserted food into database")

	return id, nil
}

func GetFood() ([]models.Food, error) {
	// implement get all orders logic here
	query := "SELECT id, name, description, category, product_id, price, stock, image, hotel_name, hotel_id, created_date, updated_date, offer FROM food"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foods []models.Food

	for rows.Next() {
		var food models.Food
		if err := rows.Scan(&food.ID, &food.Name, &food.Description, &food.Category, &food.ProductId, &food.Price, &food.Stock, &food.Image, &food.HotelName, &food.HotelId, &food.CreatedDate, &food.UpdatedDate, &food.Offer); err != nil {
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

func GetFoodByID(id uint) (*models.Food, error) {
	var food models.Food

	err := DB.QueryRow("SELECT id, name, description, category, product_id, price, stock, image, hotel_name, hotel_id, created_date, updated_date, offer FROM food WHERE id = $1", id).Scan(&food.ID, &food.Name, &food.Description, &food.Category, &food.ProductId, &food.Price, &food.Stock, &food.Image, &food.HotelName, &food.HotelId, &food.CreatedDate, &food.UpdatedDate, &food.Offer)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no food found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}

	return &food, nil
}

func PutFood(id uint, name, description, category, product_id, image, hotel_name, hotel_id string, price, stock, offer int, updated_date time.Time) error {

	// implement update logic here
	_, err := DB.Exec(
		"UPDATE food SET id=$1, name=$2, description=$3, category=$4, price=$5, stock=$6, offer=$7, image=$8, hotel_name=$9, hotel_id=$10, updated_date=$11 WHERE product_id=$12",
		id, name, description, category, price, stock, image, hotel_name, hotel_id, time.Now(), offer, product_id)

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
func PostRestaurant(hotel_id, hotel_name, description, address, location, phone_number, email, website, menu, profile_image, open_time, close_time string, ratings float64, created_date, updated_date time.Time) (uint, error) {
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

func GetRestaurantByID(id uint) (*models.Restaurant, error) {
	var restaurant models.Restaurant

	err := DB.QueryRow("SELECT id, hotel_id, hotel_name, description, address, location, phone_number, email, website, menu, profile_image, open_time, close_time, ratings, created_date, updated_date FROM restaurant WHERE id = $1", id).Scan(&restaurant.ID, &restaurant.HotelId, &restaurant.HotelName, &restaurant.Description, &restaurant.Address, &restaurant.Location, &restaurant.PhoneNumber, &restaurant.Email, &restaurant.Website, &restaurant.Menu, &restaurant.ProfileImage, &restaurant.OpenTime, &restaurant.CloseTime, &restaurant.Ratings, &restaurant.CreatedDate, &restaurant.UpdatedDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no restaurant found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}

	return &restaurant, nil
}

func PutRestaurant(id uint, hotel_id, hotel_name, description, address, location, phone_number, email, website, menu, profile_image, open_time, close_time string, ratings float64, updated_date time.Time) error {

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

func GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order

	err := DB.QueryRow("SELECT id, order_id, customer_id, product_id, quantity, hotel_id, order_address, order_location, order_time, order_status, created_date, updated_date FROM orders WHERE id = $1", id).Scan(&order.ID, &order.OrderID, &order.CustomerID, &order.ProductId, &order.Quantity, &order.HotelId, &order.OrderAddress, &order.OrderLocation, &order.OrderTime, &order.OrderStatus, &order.CreatedDate, &order.UpdatedDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no order found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}

	return &order, nil
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

// test POST and GET method
func TestPost(email, password string) (uint, error) {
	// implement post logic here
	var id uint

	err := DB.QueryRow(
		"INSERT INTO test(email, password) VALUES ($1 , $2) RETURNING id",
		email, password).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func TestGet() ([]models.Test, error) {
	// implement get logic here
	rows, err := DB.Query("SELECT id,email,password FROM test")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tests []models.Test
	for rows.Next() {
		var test models.Test
		err = rows.Scan(&test.ID, &test.Email, &test.Password)
		if err != nil {
			return nil, err
		}

		tests = append(tests, test)
	}
	return tests, nil
}

func TestGetById(id uint) (*models.Test, error) {
	// implement get logic here
	var test models.Test

	err := DB.QueryRow("SELECT id,email,password FROM test WHERE id = $1", id).Scan(&test.ID, &test.Email, &test.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no test found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}
	return &test, nil

}

// Delivery POST, GET, PUT and DELETE
func PostDelivery(name, phone_number, db_id, location, latitude, longitude string, total_payment, total_orders int, created_date, updated_date time.Time) (uint, error) {
	// implement post logic here
	var id uint

	currentTime := time.Now()

	// insert into delivery table
	err := DB.QueryRow("INSERT INTO delivery_boy (name, phone_number, db_id, location, latitude, longitude, total_payment, total_orders, created_date, updated_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id", name, phone_number, db_id, location, latitude, longitude, total_payment, total_orders, currentTime, currentTime).Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println("Post successfull")

	return id, nil
}

func GetDelivery() ([]models.DeliveryBoy, error) {
	// implement get logic here
	rows, err := DB.Query("SELECT id, name, phone_number, db_id, location, latitude, longitude, total_payment, total_orders, created_date, updated_date FROM delivery_boy")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var deliveryBoys []models.DeliveryBoy
	for rows.Next() {
		var deliveryBoy models.DeliveryBoy
		err := rows.Scan(&deliveryBoy.ID, &deliveryBoy.Name, &deliveryBoy.PhoneNumber, &deliveryBoy.DbID, &deliveryBoy.Location, &deliveryBoy.Latitude, &deliveryBoy.Longitude, &deliveryBoy.TotalPayment, &deliveryBoy.TotalOrder, &deliveryBoy.CreatedDate, &deliveryBoy.TotalOrder)
		if err != nil {
			log.Println("Error", err)
		}

		deliveryBoys = append(deliveryBoys, deliveryBoy)
	}

	fmt.Println("Get successfull")

	return deliveryBoys, nil
}

func GetDeliveryBoyByID(id uint) (*models.DeliveryBoy, error) {
	var deliveryBoy models.DeliveryBoy

	err := DB.QueryRow("SELECT id, name, phone_number, db_id, location, latitude, longitude, total_payment, total_orders, created_date, updated_date FROM delivery_boy WHERE id = $1", id).Scan(&deliveryBoy.ID, &deliveryBoy.Name, &deliveryBoy.PhoneNumber, &deliveryBoy.DbID, &deliveryBoy.Location, &deliveryBoy.Latitude, &deliveryBoy.Longitude, &deliveryBoy.TotalPayment, &deliveryBoy.TotalOrder, &deliveryBoy.CreatedDate, &deliveryBoy.TotalOrder)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no delivery boy found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}

	return &deliveryBoy, nil
}

func PutDelivery(id uint, name, phone_number, db_id, location, latitude, longitude string, total_payment, total_orders int, updated_date time.Time) error {
	// implement put logic here
	result, err := DB.Exec("UPDATE delivery_boy SET id=$1, name=$2, phone_number=$3, location=$4, latitude=$5, longitude=$6, total_payment=$7, total_orders=$8, updated_date=$9 WHERE db_id=$10", id, name, phone_number, location, latitude, longitude, total_payment, total_orders, time.Now(), db_id)
	if err != nil {
		return fmt.Errorf("failed to query delivery boy: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("delivery boy not found")
	}

	fmt.Println("Update successfull")

	return nil
}

func DeleteDelivery(db_id string) error {
	// implement delete logic here
	result, err := DB.Exec("DELETE FROM delivery_boy WHERE db_id=$1", db_id)
	if err != nil {
		return fmt.Errorf("failed to query delivery boy: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("delivery boy not found")
	}

	fmt.Println("Delete successfull")

	return nil
}

// Category GET, POST , PUT and DELETE
func PostCategory(title, category_id, icon string, created_date, updated_date time.Time) (uint, error) {
	// implement post logic here
	var id uint

	currentTime := time.Now()

	err := DB.QueryRow("INSERT INTO category (title, category_id, icon, created_date, updated_date) VALUES ($1, $2, $3, $4, $5) RETURNING id", title, category_id, icon, currentTime, currentTime).Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println("Post Successful")

	return id, nil
}

func GetCategory() ([]models.Category, error) {
	// implement get logic here
	rows, err := DB.Query("SELECT id, title, category_id, icon, created_date, updated_date FROM category")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Title, &category.CategoryID, &category.Icon, &category.CreatedDate, &category.UpdatedDate)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	fmt.Println("Get Successful")

	return categories, nil
}

func GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category

	err := DB.QueryRow("SELECT id, title, category_id, icon, created_date, updated_date FROM category WHERE id = $1", id).Scan(&category.ID, &category.Title, &category.CategoryID, &category.Icon, &category.CreatedDate, &category.UpdatedDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no category found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}

	return &category, nil
}

func PutCategory(id uint, title, category_id, icon string, updated_date time.Time) error {
	// implement put logic here
	result, err := DB.Exec("UPDATE category SET id=$1, title=$2, icon=$3, updated_date=$4 WHERE category_id=$5", id, title, icon, time.Now(), category_id)

	if err != nil {
		return fmt.Errorf("failed to query category: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("category not found")
	}

	fmt.Println("Update successfull")

	return nil
}

func DeleteCategory(category_id string) error {
	// implement delete logic here
	result, err := DB.Exec("DELETE FROM category WHERE category_id=$1", category_id)

	if err != nil {
		return fmt.Errorf("failed to query category: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("category boy not found")
	}

	fmt.Println("Delete successfull")

	return nil
}

// Banner POST, GET, PUT and DELETE
func PostBanner(title, banner_id, image, subtitle, discount string, created_date, updated_date time.Time) (uint, error) {
	// implement post logic here
	var id uint

	currentTime := time.Now()

	err := DB.QueryRow("INSERT INTO banner (title, banner_id, image, created_date, updated_date, subtitle, discount) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", title, banner_id, image, currentTime, currentTime, subtitle, discount).Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println("Post Successful")

	return id, nil
}

func GetBanner() ([]models.Banner, error) {
	// implement get logic here
	rows, err := DB.Query("SELECT id, title, banner_id, image, created_date, updated_date, subtitle, discount FROM banner")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banners []models.Banner
	for rows.Next() {
		var banner models.Banner
		err := rows.Scan(&banner.ID, &banner.Title, &banner.BannerID, &banner.Image, &banner.CreatedDate, &banner.UpdatedDate, &banner.Subtitle, &banner.Discount)
		if err != nil {
			return nil, err
		}
		banners = append(banners, banner)
	}

	fmt.Println("Get Successful")

	return banners, nil
}

func GetBannerByID(id uint) (*models.Banner, error) {
	var banner models.Banner

	err := DB.QueryRow("SELECT id, title, banner_id, image, created_date, updated_date, subtitle, discount FROM banner WHERE id = $1", id).Scan(&banner.ID, &banner.Title, &banner.BannerID, &banner.Image, &banner.CreatedDate, &banner.UpdatedDate, &banner.Subtitle, &banner.Discount)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no banner found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}

	return &banner, nil
}

func PutBanner(id uint, title, banner_id, image, subtitle, discount string, updated_date time.Time) error {
	// implement put logic here
	result, err := DB.Exec("UPDATE banner SET id=$1, title=$2, image=$3, updated_date=$4, subtitle=$5, discount=$6 WHERE banner_id=$7", id, title, image, time.Now(), subtitle, discount, banner_id)
	if err != nil {
		return fmt.Errorf("failed to query Banner: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("banner not found")
	}

	fmt.Println("Update successfull")

	return nil
}

func DeleteBanner(banner_id string) error {
	// implement delete logic here
	result, err := DB.Exec("DELETE FROM banner WHERE banner_id=$1", banner_id)

	if err != nil {
		return fmt.Errorf("failed to query category: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("banner not found")
	}

	fmt.Println("Delete successfull")

	return nil
}

// Offer GET, POST, PUT and DELETE
func PostOffer(title, subtitle, offer_id, image string, offer int, created_date, updated_date time.Time) (uint, error) {
	// implement post logic here
	var id uint

	currentTime := time.Now()

	err := DB.QueryRow("INSERT INTO offer (title, subtitle, image, offer, offer_id, created_date, updated_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", title, subtitle, image, offer, offer_id, currentTime, currentTime).Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println("Post Succesful")

	return id, nil
}

func GetOffer() ([]models.Offer, error) {
	// implement get logic here
	rows, err := DB.Query("SELECT id, title, subtitle, image,offer, offer_id, created_date, updated_date FROM offer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offers []models.Offer
	for rows.Next() {
		var offer models.Offer
		err := rows.Scan(&offer.ID, &offer.Title, &offer.Subtitle, &offer.Image, &offer.Offer, &offer.OfferID, &offer.CreatedDate, &offer.UpdatedDate)
		if err != nil {
			return nil, err
		}

		offers = append(offers, offer)
	}

	fmt.Println("Get Successful")

	return offers, nil
}

func GetOfferByID(id uint) (*models.Offer, error) {
	var offer models.Offer

	err := DB.QueryRow("SELECT id, title, subtitle, image, offer, offer_id, created_date, updated_date FROM offer WHERE id = $1", id).Scan(&offer.ID, &offer.Title, &offer.Subtitle, &offer.Image, &offer.Offer, &offer.OfferID, &offer.CreatedDate, &offer.UpdatedDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no offer found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}

	return &offer, nil
}

func PutOffer(id uint, title, subtitle, offer_id, image string, offer int, updated_date time.Time) error {
	// Correct the placeholders in the SQL query
	result, err := DB.Exec("UPDATE offer SET id=$1, title=$2, subtitle=$3, image=$4, offer=$5, updated_date=$6 WHERE offer_id=$7", id, title, subtitle, image, offer, updated_date, offer_id)
	if err != nil {
		return fmt.Errorf("failed to query Offer: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("offer not found")
	}

	fmt.Println("Update successful")

	return nil
}

func DeleteOffer(offer_id string) error {
	// implement delete logic here
	result, err := DB.Exec("DELETE FROM offer WHERE offer_id=$1", offer_id)
	if err != nil {
		return fmt.Errorf("failed to query Offer: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("offer not found")
	}

	fmt.Println("Delete successfull")

	return nil
}

// Address POST, GET, PUT and DELETE
func PostAddress(address, building_type, customer_id string, created_date, updated_date time.Time) (uint, error) {
	// implement post logic here
	var id uint

	currentTime := time.Now()

	err := DB.QueryRow("INSERT INTO address (address, building_type, customer_id, created_date, updated_date) VALUES ($1, $2, $3, $4, $5) RETURNING id", address, building_type, customer_id, currentTime, currentTime).Scan(&id)

	if err != nil {
		return 0, err
	}

	fmt.Println("Post Successful")

	return id, nil
}

func GetAddress() ([]models.Details, error) {
	// implement get logic here
	rows, err := DB.Query("SELECT id, address, building_type, customer_id, created_date, updated_date FROM address")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []models.Details
	for rows.Next() {
		var address models.Details
		err := rows.Scan(&address.ID, &address.Address, &address.BuildingType, &address.CustomerID, &address.CreatedDate, &address.UpdatedDate)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	fmt.Println("Get Successful")

	return addresses, nil
}

func GetAddressByID(id uint) (*models.Details, error) {
	var address models.Details

	err := DB.QueryRow("SELECT id, address, building_type, customer_id, created_date, updated_date FROM address WHERE id = $1", id).Scan(&address.ID, &address.Address, &address.BuildingType, &address.CustomerID, &address.CreatedDate, &address.UpdatedDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no address found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}

	return &address, nil
}

func PutAddress(id uint, address, building_type, customer_id string, updated_date time.Time) error {
	// implement put logic here
	result, err := DB.Exec("UPDATE address SET address=$1, building_type=$2, customer_id=$3, updated_date=$4 WHERE id=$5", address, building_type, customer_id, updated_date, id)
	if err != nil {
		return fmt.Errorf("failed to query Address: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("address not found")
	}

	fmt.Println("Update successfull")

	return nil
}

func DeleteAddress(id uint) error {
	// implement delete logic here
	result, err := DB.Exec("DELETE FROM address WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("failed to query Address: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("address not found")
	}

	fmt.Println("Delete successfull")

	return nil
}

// Rating POST, GET, PUT and DELETE
func PostRating(title, customer_id, product_id string, rating float64, created_date, updated_date time.Time) (uint, error) {
	// implement post logic here
	var id uint

	currentTime := time.Now()

	err := DB.QueryRow("INSERT INTO rating (title, rating, customer_id, product_id, created_date, updated_date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", title, rating, customer_id, product_id, currentTime, currentTime).Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println("Post Success")

	return id, nil
}

func GetRating() ([]models.Rating, error) {
	// implement get logic here
	rows, err := DB.Query("SELECT id, title, rating, customer_id, product_id, created_date, updated_date FROM rating")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ratings []models.Rating
	for rows.Next() {
		var rating models.Rating
		err = rows.Scan(&rating.ID, &rating.Title, &rating.Rating, &rating.CustomerID, &rating.ProductId, &rating.CreatedDate, &rating.UpdatedDate)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, rating)
	}

	fmt.Println("Get Success")

	return ratings, nil
}

func GetRatingByID(id uint) (*models.Rating, error) {
	var rating models.Rating

	err := DB.QueryRow("SELECT id, title, rating, customer_id, product_id, created_date, updated_date FROM rating WHERE id = $1", id).Scan(&rating.ID, &rating.Title, &rating.Rating, &rating.CustomerID, &rating.ProductId, &rating.CreatedDate, &rating.UpdatedDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no rating found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}

	return &rating, nil
}

func PutRating(id uint, title, customer_id, product_id string, rating float64, updated_date time.Time) error {
	// implement put logic here
	result, err := DB.Exec("UPDATE rating SET title=$1, rating=$2, customer_id=$3, product_id=$4, updated_date=$5 WHERE id=$6", title, rating, customer_id, product_id, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to query Rating: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("rating not found")
	}

	fmt.Println("Update successfull")

	return nil
}

func DeleteRating(id uint) error {
	// implement delete logic here
	result, err := DB.Exec("DELETE FROM rating WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("failed to query Rating: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("rating not found")
	}

	fmt.Println("Delete successfull")

	return nil
}

// Cart POST, GET, PUT and DELETE
func PostCart(product_id, customer_id, discount string, quantity int, order_date, created_date, updated_date time.Time) (uint, error) {
	// implement post logic here
	var id uint

	currentTime := time.Now()

	err := DB.QueryRow("INSERT INTO cart (product_id, quantity, customer_id, discount, order_date, created_date, updated_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", product_id, quantity, customer_id, discount, currentTime, currentTime, currentTime).Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println("Post Successful")

	return id, nil
}

func GetCart() ([]models.Cart, error) {
	// implement get logic here
	rows, err := DB.Query("SELECT id, product_id, quantity, customer_id, discount, order_date, created_date, updated_date FROM cart")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []models.Cart
	for rows.Next() {
		var cart models.Cart
		err = rows.Scan(&cart.ID, &cart.ProductID, &cart.Quantity, &cart.CustomerID, &cart.Discount, &cart.OrderDate, &cart.CreatedDate, &cart.UpdatedDate)
		if err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}

	fmt.Println("Get Success")

	return carts, nil
}

func GetCartByID(id uint) (*models.Cart, error) {
	var cart models.Cart

	err := DB.QueryRow("SELECT id, product_id, quantity, customer_id, discount, order_date, created_date, updated_date FROM cart WHERE id = $1", id).Scan(&cart.ID, &cart.ProductID, &cart.Quantity, &cart.CustomerID, &cart.Discount, &cart.OrderDate, &cart.CreatedDate, &cart.UpdatedDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no cart found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}

	return &cart, nil
}

func PutCart(id uint, product_id, customer_id, discount string, quantity int, updated_date time.Time) error {
	// implement put logic here
	result, err := DB.Exec("UPDATE cart SET product_id=$1, quantity=$2, customer_id=$3, discount=$4, updated_date=$5 WHERE id=$6", product_id, quantity, customer_id, discount, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to query cart: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("cart not found")
	}

	fmt.Println("Update successfull")

	return nil
}

func DeleteCart(id uint) error {
	// implement delete logic here
	result, err := DB.Exec("DELETE FROM cart WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("failed to query cart: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("cart not found")
	}

	fmt.Println("Delete successfull")

	return nil
}

// Cancel POST, GET, PUT and DELETE
func PostCancel(product_id, cancel_reason, other_reason, customer_id string, cancelled_date, created_date, updated_date time.Time) (uint, error) {
	// implement post logic here
	var id uint

	currentTime := time.Now()

	err := DB.QueryRow("INSERT INTO cancel_orders (product_id, cancel_reason, other_reason, customer_id, cancelled_date, created_date, updated_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", product_id, cancel_reason, other_reason, customer_id, currentTime, currentTime, currentTime).Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println("Post Successful")

	return id, nil
}

func GetCancel() ([]models.Cancel, error) {
	// implement get logic here
	rows, err := DB.Query("SELECT id, product_id, cancel_reason, other_reason, customer_id, cancelled_date, created_date, updated_date FROM cancel_orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cancels []models.Cancel
	for rows.Next() {
		var cancel models.Cancel
		err = rows.Scan(&cancel.ID, &cancel.ProductID, &cancel.CancelReason, &cancel.OtherReason, &cancel.CustomerID, &cancel.CancelledDate, &cancel.CreatedDate, &cancel.UpdatedDate)
		if err != nil {
			return nil, err
		}
		cancels = append(cancels, cancel)
	}

	fmt.Println("Get successful")

	return cancels, nil
}

func GetCancelByID(id uint) (*models.Cancel, error) {
	var cancel models.Cancel

	err := DB.QueryRow("SELECT id, product_id, cancel_reason, other_reason, customer_id, cancelled_date, created_date, updated_date FROM cancel_orders WHERE id = $1", id).Scan(&cancel.ID, &cancel.ProductID, &cancel.CancelReason, &cancel.OtherReason, &cancel.CustomerID, &cancel.CancelledDate, &cancel.CreatedDate, &cancel.UpdatedDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no cancelled order found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}

	return &cancel, nil
}

func PutCancel(id uint, product_id, cancel_reason, other_reason, customer_id string, updated_date time.Time) error {
	// implement put logic here
	result, err := DB.Exec("UPDATE cancel_orders SET product_id=$1, cancel_reason=$2, other_reason=$3, customer_id=$4, updated_date=$5 WHERE id=$6", product_id, cancel_reason, other_reason, customer_id, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to query cancelled order: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("cancelled order not found")
	}

	fmt.Println("Update successfull")

	return nil
}

func DeleteCancel(id uint) error {
	// implement delete logic here
	result, err := DB.Exec("DELETE FROM cancel_orders WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("failed to query cancelled order: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("cancelled order not found")
	}

	fmt.Println("Delete successfull")

	return nil
}

// Chat POST, GET, PUT and DELETE
func PostChat(message, customer_id, db_id, product_id, order_id, hotel_id, is_active string, created_date, updated_date time.Time) (uint, error) {
	// implement post logic here
	var id uint

	currentTime := time.Now()

	err := DB.QueryRow("INSERT INTO chat (message, customer_id, db_id, product_id, order_id, hotel_id, is_active, created_date, updated_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id", message, customer_id, db_id, product_id, order_id, hotel_id, is_active, currentTime, currentTime).Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println("Post Successful")

	return id, nil
}

func GetChat() ([]models.Chat, error) {
	// implement get logic here
	rows, err := DB.Query("SELECT id, message, customer_id, db_id, product_id, order_id, hotel_id, is_active, created_date, updated_date FROM chat")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []models.Chat
	for rows.Next() {
		var chat models.Chat
		err = rows.Scan(&chat.ID, &chat.Message, &chat.CustomerID, &chat.DbID, &chat.ProductID, &chat.OrderID, &chat.HotelId, &chat.IsActive, &chat.CreatedDate, &chat.UpdatedDate)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}

	fmt.Println("Get Successful")

	return chats, nil
}

func GetChatByID(id uint) (*models.Chat, error) {
	var chat models.Chat

	err := DB.QueryRow("SELECT id, message, customer_id, db_id, product_id, order_id, hotel_id, is_active, created_date, updated_date FROM chat WHERE id = $1", id).Scan(&chat.ID, &chat.Message, &chat.CustomerID, &chat.DbID, &chat.ProductID, &chat.OrderID, &chat.HotelId, &chat.IsActive, &chat.CreatedDate, &chat.UpdatedDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no chat found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}

	return &chat, nil
}

func PutChat(id uint, message, customer_id, db_id, product_id, order_id, hotel_id, is_active string, updated_date time.Time) error {
	// implement put logic here
	result, err := DB.Exec("UPDATE chat SET message=$1, customer_id=$2, db_id=$3, product_id=$4, order_id=$5, hotel_id=$6, is_active=$7, updated_date=$8 WHERE id=$9", message, customer_id, db_id, product_id, order_id, hotel_id, is_active, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to query chat: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("chat not found")
	}

	fmt.Println("Update successfull")

	return nil
}

func DeleteChat(id uint) error {
	// implement delete logic here
	result, err := DB.Exec("DELETE FROM chat WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("failed to query chat: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("chat not found")
	}

	fmt.Println("Delete successfull")

	return nil
}

// Payment POST, GET, PUT and DELETE
func PostPayment(order_id, product_id, hotel_id, customer_id, db_id, payment_method string, amount int, created_date, updated_date time.Time) (uint, error) {
	// implement post logic here
	var id uint

	currentTime := time.Now()

	err := DB.QueryRow("INSERT INTO payment (order_id, product_id, hotel_id, customer_id, db_id, amount, payment_method, created_date, updated_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id", order_id, product_id, hotel_id, customer_id, db_id, amount, payment_method, currentTime, currentTime).Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println("Post Successful")

	return id, nil
}

func GetPayment() ([]models.Payment, error) {
	// implement get logic here
	rows, err := DB.Query("SELECT id, order_id, product_id, hotel_id, customer_id, db_id, amount, payment_method, created_date, updated_date FROM payment")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []models.Payment
	for rows.Next() {
		var payment models.Payment
		err = rows.Scan(&payment.ID, &payment.OrderID, &payment.ProductID, &payment.HotelId, &payment.CustomerID, &payment.DbID, &payment.Amount, &payment.PaymentMethod, &payment.CreatedDate, &payment.UpdatedDate)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	fmt.Println("Get Successful")

	return payments, nil
}

func GetPaymentByID(id uint) (*models.Payment, error) {
	var payment models.Payment

	err := DB.QueryRow("SELECT id, order_id, product_id, hotel_id, customer_id, db_id, amount, payment_method, created_date, updated_date FROM payment WHERE id = $1", id).Scan(&payment.ID, &payment.OrderID, &payment.ProductID, &payment.HotelId, &payment.CustomerID, &payment.DbID, &payment.Amount, &payment.PaymentMethod, &payment.CreatedDate, &payment.UpdatedDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no payment found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}

	return &payment, nil
}

func PutPayment(id uint, order_id, product_id, hotel_id, customer_id, db_id, payment_method string, amount int, updated_date time.Time) error {
	// implement put logic here
	result, err := DB.Exec("UPDATE payment SET order_id=$1, product_id=$2, hotel_id=$3, customer_id=$4, db_id=$5, payment_method=$6, amount=$7, updated_date=$8 WHERE id=$9", order_id, product_id, hotel_id, customer_id, db_id, payment_method, amount, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to query payment: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("payment not found")
	}

	fmt.Println("Update successfull")

	return nil
}

func DeletePayment(id uint) error {
	// implement delete logic here
	result, err := DB.Exec("DELETE FROM payment WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("failed to query payment: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("payment not found")
	}

	fmt.Println("Delete successfull")

	return nil
}

// Support POST, GET, PUT and DELETE
func PostSupport(message, customer_id string, created_date, updated_date time.Time) (uint, error) {
	// implement post logic here
	var id uint

	currentTime := time.Now()

	err := DB.QueryRow("INSERT INTO customer_support (message, customer_id, created_date, updated_date) VALUES ($1, $2, $3, $4) RETURNING id", message, customer_id, currentTime, currentTime).Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println("Post Successful")

	return id, nil
}

func GetSupport() ([]models.Support, error) {
	// implement get logic here
	rows, err := DB.Query("SELECT id, message, customer_id, created_date, updated_date FROM customer_support")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var supports []models.Support
	for rows.Next() {
		var support models.Support
		err = rows.Scan(&support.ID, &support.Message, &support.CustomerID, &support.CreatedDate, &support.UpdatedDate)
		if err != nil {
			return nil, err
		}
		supports = append(supports, support)
	}

	fmt.Println("Get Successful")

	return supports, nil
}

func GetSupportByID(id uint) (*models.Support, error) {
	var support models.Support

	err := DB.QueryRow("SELECT id, message, customer_id, created_date, updated_date FROM customer_support WHERE id = $1", id).Scan(&support.ID, &support.Message, &support.CustomerID, &support.CreatedDate, &support.UpdatedDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no support found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}

	return &support, nil
}

func PutSupport(id uint, message, customer_id string, updated_date time.Time) error {
	// implement put logic here
	result, err := DB.Exec("UPDATE customer_support SET message=$1, customer_id=$2, updated_date=$3 WHERE id=$4", message, customer_id, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to query support: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("support not found")
	}

	fmt.Println("Update successfull")

	return nil
}

func DeleteSupport(id uint) error {
	// implement delete logic here
	result, err := DB.Exec("DELETE FROM customer_support WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("failed to query support: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("support not found")
	}

	fmt.Println("Delete successfull")

	return nil
}

// Log POST, GET, PUT and DELETE
func PostLog(function, log_message, customer_id, device_id string, created_date, updated_date time.Time) (uint, error) {
	// implement post logic here
	var id uint

	currentTime := time.Now()

	err := DB.QueryRow("INSERT INTO logs (function, log_message, customer_id, device_id, created_date, updated_date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", function, log_message, customer_id, device_id, currentTime, currentTime).Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println("Post Successful")

	return id, nil
}

func GetLog() ([]models.Logs, error) {
	// implement get logic here
	rows, err := DB.Query("SELECT id, function, log_message, customer_id, device_id, created_date, updated_date FROM logs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Logs []models.Logs
	for rows.Next() {
		var Log models.Logs
		err = rows.Scan(&Log.ID, &Log.LogMessage, &Log.CustomerID, &Log.DeviceID, &Log.CreatedDate, &Log.UpdatedDate)
		if err != nil {
			return nil, err
		}
		Logs = append(Logs, Log)
	}

	fmt.Println("Get Successful")

	return Logs, nil
}

func GetLogByID(id uint) (*models.Logs, error) {
	var Log models.Logs

	err := DB.QueryRow("SELECT id, function, log_message, customer_id, device_id, created_date, updated_date FROM logs WHERE id = $1", id).Scan(&Log.ID, &Log.LogMessage, &Log.CustomerID, &Log.DeviceID, &Log.CreatedDate, &Log.UpdatedDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no Log found with id %d", id)
		}
		log.Println("Error", err)
		return nil, err
	}

	return &Log, nil
}

func PutLog(id uint, function, log_message, customer_id, device_id string, updated_date time.Time) error {
	// implement put logic here
	result, err := DB.Exec("UPDATE logs SET function=$1, log_message=$2, customer_id=$3, device_id=$4, updated_date=$5 WHERE id=$6", function, log_message, customer_id, device_id, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to query log: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("log not found")
	}

	fmt.Println("Update successfull")

	return nil
}

func DeleteLog(id uint) error {
	// implement delete logic here
	result, err := DB.Exec("DELETE FROM logs WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("failed to query log: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("log not found")
	}

	fmt.Println("Delete successfull")

	return nil
}
