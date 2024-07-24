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

func GetCustomerHandlerById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idstr := r.URL.Query().Get("id")
	if idstr == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Retrieve customer from repository by ID
	customer, err := repository.GetCustomerByID(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no customer found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with customer details in JSON format
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

func GetAdminHandlerById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idstr := r.URL.Query().Get("id")
	if idstr == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Retrieve admin from repository by ID
	admin, err := repository.GetAdminByID(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no admin found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with admin details in JSON format
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

func GetFoodHandlerById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idstr := r.URL.Query().Get("id")
	if idstr == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Retrieve food from repository by ID
	food, err := repository.GetFoodByID(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no food found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with food details in JSON format
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
	restaurant.Ratings, err = strconv.ParseFloat("ratings", 64)

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

func GetRestuarantHandlerById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idstr := r.URL.Query().Get("id")
	if idstr == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Retrieve restaurant from repository by ID
	restaurant, err := repository.GetRestaurantByID(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no restaurant found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with restaurant details in JSON format
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

func GetOrdeHandlerById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idstr := r.URL.Query().Get("id")
	if idstr == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Retrieve order from repository by ID
	order, err := repository.GetOrderByID(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no order found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with order details in JSON format
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
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idstr := r.URL.Query().Get("id")
	if idstr == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Retrieve test from repository by ID
	test, err := repository.TestGetById(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no test found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
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

func GetDeliveryBoyByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idstr := r.URL.Query().Get("id")
	if idstr == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Retrieve customer from repository by ID
	delivery, err := repository.GetDeliveryBoyByID(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no delivery found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with delivery details in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(delivery)
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

func GetCategoryByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idstr := r.URL.Query().Get("id")
	if idstr == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Retrieve category from repository by ID
	category, err := repository.GetCategoryByID(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no category found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with category details in JSON format
	w.Header().Set("Content-Type", "application/json")
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

func GetBannerByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idstr := r.URL.Query().Get("id")
	if idstr == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Retrieve banner from repository by ID
	Banner, err := repository.GetBannerByID(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no banner found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with banner details in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Banner)
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

func GetOfferByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idstr := r.URL.Query().Get("id")
	if idstr == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Retrieve offer from repository by ID
	offer, err := repository.GetOfferByID(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no offer found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with offer details in JSON format
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

func GetAddressByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Retrieve address from repository by ID
	address, err := repository.GetAddressByID(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no offer found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with address details in JSON format
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

// Ratings Handler
func RatingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		PostRatingsHandler(w, r)
	} else if r.Method == http.MethodGet {
		GetRatingsHandler(w, r)
	} else if r.Method == http.MethodPut {
		PutRatingHandler(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteRatingHandler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PostRatingsHandler(w http.ResponseWriter, r *http.Request) {
	var ratings models.Rating

	if err := json.NewDecoder(r.Body).Decode(&ratings); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	id, err := repository.PostRating(ratings.Title, ratings.CustomerID, ratings.ProductId, ratings.Rating, ratings.CreatedDate, ratings.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ratings.ID = id

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(ratings)
}

func GetRatingsHandler(w http.ResponseWriter, r *http.Request) {
	ratings, err := repository.GetRating()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(ratings)
}

func GetRatingByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Retrieve rating from repository by ID
	rating, err := repository.GetRatingByID(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no offer found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with rating details in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rating)
}

func PutRatingHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var rating models.Rating
	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the offer_id from the query parameter
	rating.ID = uint(id)

	// Set the updated date to the current time
	rating.UpdatedDate = time.Now()

	// Update the Offer in the repository
	err = repository.PutRating(rating.ID, rating.Title, rating.CustomerID, rating.ProductId, rating.Rating, rating.UpdatedDate)
	if err != nil {
		if err.Error() == "Rating not found" {
			http.Error(w, "Rating not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update rating: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(rating); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteRatingHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = repository.DeleteRating(uint(id))
	if err != nil {
		if err.Error() == "Rating not found" {
			http.Error(w, "Rating not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update rating: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// Cart Handler
func CartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		PostCartHandler(w, r)
	} else if r.Method == http.MethodGet {
		GetCartHandler(w, r)
	} else if r.Method == http.MethodPut {
		PutCartHandler(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteCartHandler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PostCartHandler(w http.ResponseWriter, r *http.Request) {
	var cart models.Cart

	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	id, err := repository.PostCart(cart.ProductID, cart.CustomerID, cart.Discount, cart.Quantity, cart.OrderDate, cart.CreatedDate, cart.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cart.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func GetCartHandler(w http.ResponseWriter, r *http.Request) {
	cart, err := repository.GetCart()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func GetCartByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Retrieve cart from repository by ID
	cart, err := repository.GetCartByID(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no offer found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with cart details in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func PutCartHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var cart models.Cart
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the offer_id from the query parameter
	cart.ID = uint(id)

	// Set the updated date to the current time
	cart.UpdatedDate = time.Now()

	// Update the Offer in the repository
	err = repository.PutCart(cart.ID, cart.ProductID, cart.CustomerID, cart.Discount, cart.Quantity, cart.UpdatedDate)
	if err != nil {
		if err.Error() == "Cart not found" {
			http.Error(w, "Cart not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update Cart: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cart); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteCartHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = repository.DeleteCart(uint(id))
	if err != nil {
		if err.Error() == "Cart not found" {
			http.Error(w, "Cart not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update Cart: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// Cart Handler
func CancelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		PostCancelHandler(w, r)
	} else if r.Method == http.MethodGet {
		GetCancelHandler(w, r)
	} else if r.Method == http.MethodPut {
		PutCancelHandler(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteCancelHandler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PostCancelHandler(w http.ResponseWriter, r *http.Request) {
	var cancel models.Cancel

	if err := json.NewDecoder(r.Body).Decode(&cancel); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	id, err := repository.PostCancel(cancel.ProductID, cancel.CancelReason, cancel.OtherReason, cancel.CustomerID, cancel.CancelledDate, cancel.CreatedDate, cancel.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cancel.ID = id

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(cancel)
}

func GetCancelHandler(w http.ResponseWriter, r *http.Request) {
	cancel, err := repository.GetCancel()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(cancel)
}

func GetCancelByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Retrieve cancel from repository by ID
	cancel, err := repository.GetCancelByID(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no offer found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with cancel details in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cancel)
}

func PutCancelHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var cancel models.Cancel
	if err := json.NewDecoder(r.Body).Decode(&cancel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the id from the query parameter
	cancel.ID = uint(id)

	// Set the updated date to the current time
	cancel.UpdatedDate = time.Now()

	// Update the Offer in the repository
	err = repository.PutCancel(cancel.ID, cancel.ProductID, cancel.CancelReason, cancel.OtherReason, cancel.CustomerID, cancel.UpdatedDate)
	if err != nil {
		if err.Error() == "Cancelled order not found" {
			http.Error(w, "Cancelled order not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update Cancelled order: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cancel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteCancelHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = repository.DeleteCancel(uint(id))
	if err != nil {
		if err.Error() == "Cancelled order not found" {
			http.Error(w, "Cancelled order not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update Cancelled order: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// Chat Handler
func ChatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		PostChatHandler(w, r)
	} else if r.Method == http.MethodGet {
		GetChatHandler(w, r)
	} else if r.Method == http.MethodPut {
		PutChatHandler(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteChatHandler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PostChatHandler(w http.ResponseWriter, r *http.Request) {
	var chat models.Chat

	if err := json.NewDecoder(r.Body).Decode(&chat); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	id, err := repository.PostChat(chat.Message, chat.CustomerID, chat.DbID, chat.ProductID, chat.OrderID, chat.HotelId, chat.IsActive, chat.CreatedDate, chat.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	chat.ID = id

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(chat)
}

func GetChatHandler(w http.ResponseWriter, r *http.Request) {
	chat, err := repository.GetChat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(chat)
}

func GetChatByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Retrieve chat from repository by ID
	chat, err := repository.GetChatByID(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no offer found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with chat details in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chat)
}

func PutChatHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var chat models.Chat
	if err := json.NewDecoder(r.Body).Decode(&chat); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the id from the query parameter
	chat.ID = uint(id)

	// Set the updated date to the current time
	chat.UpdatedDate = time.Now()

	// Update the Offer in the repository
	err = repository.PutChat(chat.ID, chat.Message, chat.CustomerID, chat.DbID, chat.ProductID, chat.OrderID, chat.HotelId, chat.IsActive, chat.UpdatedDate)
	if err != nil {
		if err.Error() == "Chat not found" {
			http.Error(w, "Chat not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update Chat: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(chat); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteChatHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = repository.DeleteChat(uint(id))
	if err != nil {
		if err.Error() == "Chat not found" {
			http.Error(w, "Chat not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update Chat: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// Payment method
func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		PostPaymentHandler(w, r)
	} else if r.Method == http.MethodGet {
		GetPaymentHandler(w, r)
	} else if r.Method == http.MethodPut {
		PutPaymentHandler(w, r)
	} else if r.Method == http.MethodDelete {
		DeletePaymentHandler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PostPaymentHandler(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	id, err := repository.PostPayment(payment.OrderID, payment.ProductID, payment.HotelId, payment.CustomerID, payment.DbID, payment.PaymentMethod, payment.Amount, payment.CreatedDate, payment.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payment.ID = id

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(payment)
}

func GetPaymentHandler(w http.ResponseWriter, r *http.Request) {
	payment, err := repository.GetPayment()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(payment)
}

func GetPaymentByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Retrieve payment from repository by ID
	payment, err := repository.GetPaymentByID(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no offer found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with payment details in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}

func PutPaymentHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the id from the query parameter
	payment.ID = uint(id)

	// Set the updated date to the current time
	payment.UpdatedDate = time.Now()

	// Update the Offer in the repository
	err = repository.PutPayment(payment.ID, payment.OrderID, payment.ProductID, payment.HotelId, payment.CustomerID, payment.DbID, payment.PaymentMethod, payment.Amount, payment.UpdatedDate)
	if err != nil {
		if err.Error() == "Payment not found" {
			http.Error(w, "Payment not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update payment: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(payment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeletePaymentHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = repository.DeletePayment(uint(id))
	if err != nil {
		if err.Error() == "Payment not found" {
			http.Error(w, "Payment not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update payment: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// Support Handler
func SupportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		PostSupportHandler(w, r)
	} else if r.Method == http.MethodGet {
		GetSupportHandler(w, r)
	} else if r.Method == http.MethodPut {
		PutSupportHandler(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteSupportHandler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PostSupportHandler(w http.ResponseWriter, r *http.Request) {
	var support models.Support

	if err := json.NewDecoder(r.Body).Decode(&support); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	id, err := repository.PostSupport(support.Message, support.CustomerID, support.CreatedDate, support.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	support.ID = id

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(support)
}

func GetSupportHandler(w http.ResponseWriter, r *http.Request) {
	support, err := repository.GetSupport()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(support)
}

func GetSupportByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Retrieve support from repository by ID
	support, err := repository.GetPaymentByID(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no offer found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with support details in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(support)
}

func PutSupportHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var support models.Support
	if err := json.NewDecoder(r.Body).Decode(&support); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the id from the query parameter
	support.ID = uint(id)

	// Set the updated date to the current time
	support.UpdatedDate = time.Now()

	// Update the Offer in the repository
	err = repository.PutSupport(support.ID, support.Message, support.CustomerID, support.UpdatedDate)
	if err != nil {
		if err.Error() == "Support not found" {
			http.Error(w, "Support not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update Support: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(support); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteSupportHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = repository.DeleteSupport(uint(id))
	if err != nil {
		if err.Error() == "Support not found" {
			http.Error(w, "Support not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update support: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// Log Handler
func LogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		PostLogHandler(w, r)
	} else if r.Method == http.MethodGet {
		GetLogHandler(w, r)
	} else if r.Method == http.MethodPut {
		PutLogHandler(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteLogHandler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PostLogHandler(w http.ResponseWriter, r *http.Request) {
	var log models.Logs

	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	id, err := repository.PostLog(log.Functions, log.LogMessage, log.CustomerID, log.DeviceID, log.CreatedDate, log.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.ID = id

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(log)
}

func GetLogHandler(w http.ResponseWriter, r *http.Request) {
	log, err := repository.GetSupport()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(log)
}

func GetLogByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL query parameter
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Retrieve support from repository by ID
	log, err := repository.GetLogByID(uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("no log found with id %d", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with support details in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(log)
}

func PutLogHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var log models.Logs
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the id from the query parameter
	log.ID = uint(id)

	// Set the updated date to the current time
	log.UpdatedDate = time.Now()

	// Update the Offer in the repository
	err = repository.PutLog(log.ID, log.Functions, log.LogMessage, log.CustomerID, log.DeviceID, log.UpdatedDate)
	if err != nil {
		if err.Error() == "Log not found" {
			http.Error(w, "Log not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update log: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(log); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteLogHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = repository.DeleteLog(uint(id))
	if err != nil {
		if err.Error() == "Log not found" {
			http.Error(w, "Log not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("Failed to update Log: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
