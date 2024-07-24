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
	food.Offer, _ = strconv.Atoi(r.FormValue("offer"))

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

	// Parse form data to get the file and food details
	err := r.ParseMultipartForm(10 << 20) // 10MB max file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Extract food details from request body
	var restaurant models.Restaurant
	restaurant.HotelId = r.FormValue("hotel_id")
	restaurant.HotelName = r.FormValue("hotel_name")
	restaurant.Description = r.FormValue("description")
	restaurant.Address = r.FormValue("address")
	restaurant.Location = r.FormValue("location")
	restaurant.PhoneNumber = r.FormValue("phone_number")
	restaurant.Email = r.FormValue("email")
	restaurant.Website = r.FormValue("website")
	restaurant.Menu = r.FormValue("menu")
	restaurant.OpenTime = r.FormValue("open_time")
	restaurant.CloseTime = r.FormValue("close_time")
	restaurant.Ratings, _ = strconv.Atoi(r.FormValue("ratings"))

	// Process uploaded image
	file, _, err := r.FormFile("profile_image")
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
	imageKey := fmt.Sprintf("ResturantImage/%d.jpg", time.Now().Unix()) // Adjust key as needed
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

	id, err := repository.PostRestaurant(restaurant.HotelId, restaurant.HotelName, restaurant.Description, restaurant.Address, restaurant.Location, restaurant.PhoneNumber, restaurant.Email, restaurant.Website, restaurant.Menu, imageURL, restaurant.OpenTime, restaurant.CloseTime, restaurant.Ratings, restaurant.CreatedDate, restaurant.UpdatedDate)
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

// Category Handler
func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		PostCategoryHandler(w, r)
	} else if r.Method == http.MethodGet {
		GetCategoryHandler(w, r)
	} else if r.Method == http.MethodPut {
		PutCategoryHandler(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteCategoryHandler(w, r)
	} else {
		http.Error(w, "Invalid Method", http.StatusBadRequest)
	}
}

func PostCategoryHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20) // 10MB max file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	var category models.Category
	category.Title = r.FormValue("title")
	category.CategoryID = r.FormValue("category_id")

	// Process uploaded image
	file, _, err := r.FormFile("icon")
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
	imageKey := fmt.Sprintf("Icons/%d.jpg", time.Now().Unix()) // Adjust key as needed
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

	id, err := repository.PostCategory(category.Title, category.CategoryID, imageURL, category.CreatedDate, category.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	category.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

func GetCategoryHandler(w http.ResponseWriter, r *http.Request) {
	category, err := repository.GetCategory()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(category)
}

func PutCategoryHandler(w http.ResponseWriter, r *http.Request) {
	// Extract category_id from query parameters
	queryParams := r.URL.Query()
	CategoryID := queryParams.Get("category_id")
	if CategoryID == "" {
		http.Error(w, "Missing category_id query parameter", http.StatusBadRequest)
		return
	}

	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the category_id from the query parameter
	category.CategoryID = CategoryID

	// Set the updated date to the current time
	category.UpdatedDate = time.Now()

	// Update the Category in the repository
	err := repository.PutCategory(category.ID, category.Title, category.CategoryID, category.Icon, category.UpdatedDate)
	if err != nil {
		if err.Error() == "Category not found" {
			http.Error(w, "Category not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update Category: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	CategoryID := r.URL.Query().Get("category_id")
	if CategoryID == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err := repository.DeleteCategory(CategoryID)
	if err != nil {
		if err.Error() == "Category not found" {
			http.Error(w, "Category boy not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to delete Category: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// Banner Handler
func BannerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		PostBannerHandler(w, r)
	} else if r.Method == http.MethodGet {
		GetBannerHandler(w, r)
	} else if r.Method == http.MethodPut {
		PutBannerHandler(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteBannerHandler(w, r)
	} else {
		http.Error(w, "Invalid Method", http.StatusBadRequest)
	}
}

func PostBannerHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB max file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	var banner models.Banner
	banner.Title = r.FormValue("title")
	banner.BannerID = r.FormValue("banner_id")
	banner.Subtitle = r.FormValue("subtitle")
	banner.Discount = r.FormValue("discount")

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
	imageKey := fmt.Sprintf("BannerImages/%d.jpg", time.Now().Unix()) // Adjust key as needed
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

	id, err := repository.PostBanner(banner.Title, banner.BannerID, imageURL, banner.Subtitle, banner.Discount, banner.CreatedDate, banner.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	banner.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(banner)
}

func GetBannerHandler(w http.ResponseWriter, r *http.Request) {
	banner, err := repository.GetBanner()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(banner)
}

func PutBannerHandler(w http.ResponseWriter, r *http.Request) {
	// Extract category_id from query parameters
	queryParams := r.URL.Query()
	BannerID := queryParams.Get("banner_id")
	if BannerID == "" {
		http.Error(w, "Missing category_id query parameter", http.StatusBadRequest)
		return
	}

	var banner models.Banner
	if err := json.NewDecoder(r.Body).Decode(&banner); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the banner_id from the query parameter
	banner.BannerID = BannerID

	// Set the updated date to the current time
	banner.UpdatedDate = time.Now()

	// Update the Banner in the repository
	err := repository.PutBanner(banner.ID, banner.Title, banner.BannerID, banner.Image, banner.Subtitle, banner.Discount, banner.UpdatedDate)
	if err != nil {
		if err.Error() == "Banner not found" {
			http.Error(w, "Banner not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update Banner: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(banner); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteBannerHandler(w http.ResponseWriter, r *http.Request) {
	BannerID := r.URL.Query().Get("banner_id")
	if BannerID == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err := repository.DeleteBanner(BannerID)
	if err != nil {
		if err.Error() == "Banner not found" {
			http.Error(w, "Banner not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to delete Banner: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// Offer Handler
func OfferHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		PostOfferHandler(w, r)
	} else if r.Method == http.MethodGet {
		GetOfferHandler(w, r)
	} else if r.Method == http.MethodPut {
		PutOfferHandler(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteOfferHandler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PostOfferHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB max file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	var offer models.Offer
	offer.Title = r.FormValue("title")
	offer.Subtitle = r.FormValue("subtitle")
	offer.OfferID = r.FormValue("offer_id")
	offer.Offer, _ = strconv.Atoi(r.FormValue("offer"))

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
	imageKey := fmt.Sprintf("BannerImages/%d.jpg", time.Now().Unix()) // Adjust key as needed
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

	id, err := repository.PostOffer(offer.Title, offer.Subtitle, offer.OfferID, imageURL, offer.Offer, offer.CreatedDate, offer.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	offer.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(offer)
}

func GetOfferHandler(w http.ResponseWriter, r *http.Request) {
	offer, err := repository.GetOffer()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(offer)
}

func PutOfferHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	OfferID := queryParams.Get("offer_id")
	if OfferID == "" {
		http.Error(w, "Missing offer_id query parameter", http.StatusBadRequest)
		return
	}

	var offer models.Offer
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the offer_id from the query parameter
	offer.OfferID = OfferID

	// Set the updated date to the current time
	offer.UpdatedDate = time.Now()

	// Update the Offer in the repository
	err := repository.PutOffer(offer.ID, offer.Title, offer.Subtitle, offer.OfferID, offer.Image, offer.Offer, offer.UpdatedDate)
	if err != nil {
		if err.Error() == "offer not found" {
			http.Error(w, "Offer not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update offer: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(offer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteOfferHandler(w http.ResponseWriter, r *http.Request) {
	OfferID := r.URL.Query().Get("offer_id")
	if OfferID == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err := repository.DeleteOffer(OfferID)
	if err != nil {
		if err.Error() == "Offer not found" {
			http.Error(w, "Offer not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to delete Offer: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// Address Handler
func AddressHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		PostAddressHandler(w, r)
	} else if r.Method == http.MethodGet {
		GetAddressHandler(w, r)
	} else if r.Method == http.MethodPut {
		PutAddressHandler(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteAddressHandler(w, r)
	} else {
		http.Error(w, "Invalid Method", http.StatusBadRequest)
	}
}

func PostAddressHandler(w http.ResponseWriter, r *http.Request) {
	var address models.Details
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := repository.PostAddress(address.Address, address.BuildingType, address.CustomerID, address.CreatedDate, address.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	address.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(address)
}

func GetAddressHandler(w http.ResponseWriter, r *http.Request) {
	address, err := repository.GetAddress()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(address)
}

func PutAddressHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var address models.Details
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the offer_id from the query parameter
	address.ID = uint(id)

	// Set the updated date to the current time
	address.UpdatedDate = time.Now()

	// Update the Offer in the repository
	err = repository.PutAddress(address.ID, address.Address, address.BuildingType, address.CustomerID, address.UpdatedDate)
	if err != nil {
		if err.Error() == "Address not found" {
			http.Error(w, "Address not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update address: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(address); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteAddressHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = repository.DeleteAddress(uint(id))
	if err != nil {
		if err.Error() == "Address not found" {
			http.Error(w, "Address not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to delete Address: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
