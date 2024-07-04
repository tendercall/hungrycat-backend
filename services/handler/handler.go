package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"hungerycat-backend.com/main/services/models"
	"hungerycat-backend.com/main/services/repository"
)

//Adin Handler

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetAdminHandler(w, r)
	} else if r.Method == http.MethodPost {
		PostAdminHandler(w, r)
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

// Restaurant Handler
func RestaurantHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetRestaurantHandler(w, r)
	} else if r.Method == http.MethodPost {
		PostRestaurantHandler(w, r)
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
