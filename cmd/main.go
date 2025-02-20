package main

import (
	"net/http"

	"hungerycat-backend.com/main/database"
	"hungerycat-backend.com/main/middleware"
	"hungerycat-backend.com/main/services/handler"
)

func main() {
	database.Initdb()

	//Customer router
	http.Handle("/customer", middleware.AuthMiddleware(http.HandlerFunc(handler.CustomerHandler)))

	//Admin router
	http.Handle("/admin", middleware.AuthMiddleware(http.HandlerFunc(handler.AdminHandler)))

	//sign in router
	http.Handle("/signin", middleware.AuthMiddleware(http.HandlerFunc(handler.CheckEmailAndPasswordHandler)))

	//Food router
	http.Handle("/food", middleware.AuthMiddleware(http.HandlerFunc(handler.FoodHandler)))

	//Restaurant router
	http.Handle("/restaurant", middleware.AuthMiddleware(http.HandlerFunc(handler.RestaurantHandler)))

	//Order router
	http.Handle("/order", middleware.AuthMiddleware(http.HandlerFunc(handler.OrderHandler)))

	http.ListenAndServe(":8080", nil)
}
