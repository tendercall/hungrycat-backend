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

	//Test router
	http.Handle("/test", middleware.AuthMiddleware(http.HandlerFunc(handler.TestHandler)))
	http.Handle("/test/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.TestGetByIdHandler)))

	//Delivery router
	http.Handle("/delivery", middleware.AuthMiddleware(http.HandlerFunc(handler.DeliveryBoyHandler)))

	//Categoory router
	http.Handle("/category", middleware.AuthMiddleware(http.HandlerFunc(handler.CategoryHandler)))

	//Banner router
	http.Handle("/banner", middleware.AuthMiddleware(http.HandlerFunc(handler.BannerHandler)))

	//Offer router
	http.Handle("/offer", middleware.AuthMiddleware(http.HandlerFunc(handler.OfferHandler)))

	//Address router
	http.Handle("/address", middleware.AuthMiddleware(http.HandlerFunc(handler.AddressHandler)))

	http.ListenAndServe(":8080", nil)
}
