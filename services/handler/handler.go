package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nfnt/resize"
	"hungerycat-backend.com/main/services/models"
	"hungerycat-backend.com/main/services/repository"
)

// Customer Handler
func CustomerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetCustomerHandler(w, r)
	} else if r.Method == http.MethodPost {
		PostCustomerHandler(w, r)
	} else if r.Method == http.MethodPut {
		PutCustomerHandler(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteCustomerHandler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PostCustomerHandler(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	var customer models.Customer

	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := repository.PostCustomer(customer.Name, customer.Email, customer.Password, customer.PhoneNumber, customer.CustomerID, customer.ProfileImage, customer.Address, customer.Location, customer.CreatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	customer.ID = id

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	fmt.Printf("Function executed in %v\n", executionTime)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func GetCustomerHandler(w http.ResponseWriter, r *http.Request) {

	customer, err := repository.GetCustomer()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func PutCustomerHandler(w http.ResponseWriter, r *http.Request) {

	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := repository.PutCustomer(customer.ID, customer.Name, customer.Email, customer.Password, customer.PhoneNumber, customer.CustomerID, customer.ProfileImage, customer.Address, customer.Location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func DeleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	customer_id := r.URL.Query().Get("customer_id")
	if customer_id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err := repository.DeleteCustomer(string(customer_id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Admin Handler
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetAdminHandler(w, r)
	} else if r.Method == http.MethodPost {
		PostAdminHandler(w, r)
	} else if r.Method == http.MethodPut {
		PutAdminHandler(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteAdminHandler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PostAdminHandler(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	var admin models.Admin

	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := repository.PostAdmin(admin.Username, admin.Email, admin.Password, admin.PhoneNumber, admin.AdminId, admin.ProfileImage, admin.CreatedAt, admin.LastSingIn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	admin.ID = id

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	fmt.Printf("Function executed in %v\n", executionTime)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(admin)
}

func GetAdminHandler(w http.ResponseWriter, r *http.Request) {

	admin, err := repository.GetAdmin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(admin)
}

func PutAdminHandler(w http.ResponseWriter, r *http.Request) {

	var Admin models.Admin
	if err := json.NewDecoder(r.Body).Decode(&Admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the updated date to the current time
	Admin.LastSingIn = time.Now()

	err := repository.PutAdmin(Admin.ID, Admin.Username, Admin.Email, Admin.Password, Admin.PhoneNumber, Admin.AdminId, Admin.ProfileImage, Admin.LastSingIn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Admin)
}

func DeleteAdminHandler(w http.ResponseWriter, r *http.Request) {
	admin_id := r.URL.Query().Get("admin_id")
	if admin_id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err := repository.DeleteAdmin(string(admin_id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func CheckEmailAndPasswordHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var admin models.Admin
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exists, err := repository.CheckEmailAndPassword(admin.Email, admin.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if exists {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Admin exists")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Admin not found")
	}
}

// Food Handler
func FoodHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetFoodHandler(w, r)
	} else if r.Method == http.MethodPost {
		PostFoodHandler(w, r)
	} else if r.Method == http.MethodPut {
		PutFoodHandler(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteFoodHandler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PostFoodHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	// Parse form data to get the file and food details
	err := r.ParseMultipartForm(10 << 20) // 10MB max file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Extract food details from request body
	var food models.Food
	food.Name = r.FormValue("name")
	food.Description = r.FormValue("description")
	food.Category = r.FormValue("category")
	food.ProductId = r.FormValue("product_id")
	food.Price, _ = strconv.Atoi(r.FormValue("price"))
	food.Stock, _ = strconv.Atoi(r.FormValue("stock"))
	food.HotelName = r.FormValue("hotel_name")
	food.HotelId = r.FormValue("hotel_id")

	// Process uploaded image
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error uploading file", http.StatusBadRequest)
		fmt.Println("Error uploading file:", err)
		return
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file content", http.StatusInternalServerError)
		fmt.Println("Error reading file content:", err)
		return
	}

	// Resize image if it exceeds 3MB
	if len(fileBytes) > 3*1024*1024 {
		img, _, err := image.Decode(bytes.NewReader(fileBytes))
		if err != nil {
			http.Error(w, "Error decoding image", http.StatusInternalServerError)
			fmt.Println("Error decoding image:", err)
			return
		}

		newImage := resize.Resize(800, 0, img, resize.Lanczos3)
		var buf bytes.Buffer
		err = jpeg.Encode(&buf, newImage, nil)
		if err != nil {
			http.Error(w, "Error encoding compressed image", http.StatusInternalServerError)
			fmt.Println("Error encoding compressed image:", err)
			return
		}
		fileBytes = buf.Bytes()
	}

	// Upload image to AWS S3
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Replace with your AWS region
		Credentials: credentials.NewStaticCredentials(
			"AKIAYS2NVN4MBSHP33FF",                     // Replace with your AWS access key ID
			"aILySGhiQAB7SaFnqozcRZe1MhZ0zNODLof2Alr4", // Replace with your AWS secret access key
			""), // Optional token, leave blank if not using
	})
	if err != nil {
		log.Printf("Failed to create AWS session: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	svc := s3.New(sess)
	imageKey := fmt.Sprintf("FoodImage/%d.jpg", time.Now().Unix()) // Adjust key as needed
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("tendercall-db"), // Replace with your S3 bucket name
		Key:    aws.String(imageKey),
		Body:   bytes.NewReader(fileBytes),
	})
	if err != nil {
		log.Printf("Failed to upload image to S3: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Construct imageURL assuming it's from your S3 bucket
	imageURL := fmt.Sprintf("https://tendercall-db.s3.amazonaws.com/%s", imageKey)

	// Save food details to database
	id, err := repository.PostFood(food.Name, food.Description, food.Category, food.ProductId, imageURL, food.HotelName, food.HotelId, food.Price, food.Stock, food.Offer, time.Now(), time.Now())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Update the food object with retrieved data
	food.ID = id

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Function executed in %v\n", executionTime)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(food)
}

func GetFoodHandler(w http.ResponseWriter, r *http.Request) {

	food, err := repository.GetFood()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(food)
}

func PutFoodHandler(w http.ResponseWriter, r *http.Request) {

	var food models.Food
	if err := json.NewDecoder(r.Body).Decode(&food); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the updated date to the current time
	food.UpdatedDate = time.Now()

	err := repository.PutFood(food.ID, food.Name, food.Description, food.Category, food.ProductId, food.Image, food.HotelName, food.HotelId, food.Price, food.Stock, food.Offer, food.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(food)
}

func DeleteFoodHandler(w http.ResponseWriter, r *http.Request) {
	product_id := r.URL.Query().Get("product_id")
	if product_id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err := repository.DeleteFood(string(product_id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Restaurant Handler
func RestaurantHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetRestaurantHandler(w, r)
	} else if r.Method == http.MethodPost {
		PostRestaurantHandler(w, r)
	} else if r.Method == http.MethodPut {
		PutRestaurantHandler(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteRestaurantHandler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PostRestaurantHandler(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	var restaurant models.Restaurant

	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := repository.PostRestaurant(restaurant.HotelId, restaurant.HotelName, restaurant.Description, restaurant.Address, restaurant.Location, restaurant.PhoneNumber, restaurant.Email, restaurant.Website, restaurant.Menu, restaurant.ProfileImage, restaurant.OpenTime, restaurant.CloseTime, restaurant.Ratings, restaurant.CreatedDate, restaurant.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	restaurant.ID = id

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	fmt.Printf("Function executed in %v\n", executionTime)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurant)
}

func GetRestaurantHandler(w http.ResponseWriter, r *http.Request) {

	restaurant, err := repository.GetRestaurant()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurant)
}

func PutRestaurantHandler(w http.ResponseWriter, r *http.Request) {

	var restaurant models.Restaurant
	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the updated date to the current time
	restaurant.UpdatedDate = time.Now()

	err := repository.PutRestaurant(restaurant.ID, restaurant.HotelId, restaurant.HotelName, restaurant.Description, restaurant.Address, restaurant.Location, restaurant.PhoneNumber, restaurant.Email, restaurant.Website, restaurant.Menu, restaurant.ProfileImage, restaurant.OpenTime, restaurant.CloseTime, restaurant.Ratings, restaurant.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurant)
}

func DeleteRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	hotel_id := r.URL.Query().Get("hotel_id")
	if hotel_id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err := repository.DeleteRestaurant(string(hotel_id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Order Handler
func OrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetOrderHandler(w, r)
	} else if r.Method == http.MethodPost {
		PostOrderHandler(w, r)
	} else if r.Method == http.MethodPut {
		PutOrderHandler(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteOrderHandler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PostOrderHandler(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	var order models.Order

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := repository.PostOrder(order.OrderID, order.CustomerID, order.ProductId, order.HotelId, order.OrderAddress, order.OrderLocation, order.OrderStatus, order.Quantity, order.OrderTime, order.CreatedDate, order.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	order.ID = id

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	fmt.Printf("Function executed in %v\n", executionTime)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {

	order, err := repository.GetOrder()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func PutOrderHandler(w http.ResponseWriter, r *http.Request) {

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the updated date to the current time
	order.UpdatedDate = time.Now()
	order.OrderTime = time.Now()

	err := repository.PutOrder(order.ID, order.OrderID, order.CustomerID, order.ProductId, order.HotelId, order.OrderAddress, order.OrderLocation, order.OrderStatus, order.Quantity, order.OrderTime, order.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	order_id := r.URL.Query().Get("order_id")
	if order_id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err := repository.DeleteOrder(string(order_id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Test Handler
func TestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		TestGetHandler(w, r)
	} else if r.Method == http.MethodPost {
		TestPostHandler(w, r)
	} else {
		http.Error(w, "Invalid Method", http.StatusBadRequest)
	}
}

func TestPostHandler(w http.ResponseWriter, r *http.Request) {
	var test models.Test
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := repository.TestPost(test.Email, test.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	test.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(test)
}

func TestGetHandler(w http.ResponseWriter, r *http.Request) {
	test, err := repository.TestGet()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(test)
}

func TestGetByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid Method", http.StatusBadRequest)
	}

	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	test, err := repository.TestGetById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(test)
}

// Delivery Handler
func DeliveryBoyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		PostBeliveryBoyHandler(w, r)
	} else if r.Method == http.MethodGet {
		GetDeliveryBoyHandler(w, r)
	} else if r.Method == http.MethodPut {
		PutDeliveryBoyHandler(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteDeliveryBoyHandler(w, r)
	} else {
		http.Error(w, "Invalid Method", http.StatusBadRequest)
	}
}

func PostBeliveryBoyHandler(w http.ResponseWriter, r *http.Request) {
	var deliveryBoy models.DeliveryBoy
	if err := json.NewDecoder(r.Body).Decode(&deliveryBoy); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := repository.PostDelivery(deliveryBoy.Name, deliveryBoy.PhoneNumber, deliveryBoy.DbID, deliveryBoy.Location, deliveryBoy.Latitude, deliveryBoy.Longitude, deliveryBoy.TotalPayment, deliveryBoy.TotalOrder, deliveryBoy.CreatedDate, deliveryBoy.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	deliveryBoy.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deliveryBoy)
}

func GetDeliveryBoyHandler(w http.ResponseWriter, r *http.Request) {
	deliveryBoy, err := repository.GetDelivery()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(deliveryBoy)
}

func PutDeliveryBoyHandler(w http.ResponseWriter, r *http.Request) {
	// Extract db_id from query parameters
	queryParams := r.URL.Query()
	dbID := queryParams.Get("db_id")
	if dbID == "" {
		http.Error(w, "Missing db_id query parameter", http.StatusBadRequest)
		return
	}

	var deliveryBoy models.DeliveryBoy
	if err := json.NewDecoder(r.Body).Decode(&deliveryBoy); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the db_id from the query parameter
	deliveryBoy.DbID = dbID

	// Set the updated date to the current time
	deliveryBoy.UpdatedDate = time.Now()

	// Update the delivery boy in the repository
	err := repository.PutDelivery(deliveryBoy.ID, deliveryBoy.Name, deliveryBoy.PhoneNumber, deliveryBoy.DbID, deliveryBoy.Location, deliveryBoy.Latitude, deliveryBoy.Longitude, deliveryBoy.TotalPayment, deliveryBoy.TotalOrder, deliveryBoy.UpdatedDate)
	if err != nil {
		if err.Error() == "delivery boy not found" {
			http.Error(w, "Delivery boy not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update delivery boy: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(deliveryBoy); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteDeliveryBoyHandler(w http.ResponseWriter, r *http.Request) {
	dbID := r.URL.Query().Get("db_id")
	if dbID == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err := repository.DeleteDelivery(dbID)
	if err != nil {
		if err.Error() == "delivery boy not found" {
			http.Error(w, "Delivery boy not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to delete delivery boy: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
