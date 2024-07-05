package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"hungerycat-backend.com/main/services/models"
	"hungerycat-backend.com/main/services/repository"
)

//Adin Handler

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
	var food models.Food

	if err := json.NewDecoder(r.Body).Decode(&food); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := repository.PostFood(food.Name, food.Description, food.Category, food.ProductId, food.Image, food.HotelName, food.HotelId, food.Price, food.Stock, food.CreatedDate, food.UpdatedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	food.ID = id

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

	err := repository.PutFood(food.ID, food.Name, food.Description, food.Category, food.ProductId, food.Image, food.HotelName, food.HotelId, food.Price, food.Stock, food.UpdatedDate)
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
