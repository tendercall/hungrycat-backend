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
	http.Handle("/customers", middleware.AuthMiddleware(http.HandlerFunc(handler.GetCustomerHandlerById)))

	//Admin router
	http.Handle("/admin", middleware.AuthMiddleware(http.HandlerFunc(handler.AdminHandler)))
	http.Handle("/admins", middleware.AuthMiddleware(http.HandlerFunc(handler.GetAdminHandlerById)))

	//sign in router
	http.Handle("/signin", middleware.AuthMiddleware(http.HandlerFunc(handler.CheckEmailAndPasswordHandler)))

	//Food router
	http.Handle("/food", middleware.AuthMiddleware(http.HandlerFunc(handler.FoodHandler)))
	http.Handle("/foods", middleware.AuthMiddleware(http.HandlerFunc(handler.GetFoodHandlerById)))

	//Restaurant router
	http.Handle("/restaurant", middleware.AuthMiddleware(http.HandlerFunc(handler.RestaurantHandler)))
	http.Handle("/restaurants", middleware.AuthMiddleware(http.HandlerFunc(handler.GetRestuarantHandlerById)))

	//Order router
	http.Handle("/order", middleware.AuthMiddleware(http.HandlerFunc(handler.OrderHandler)))
	http.Handle("/orders", middleware.AuthMiddleware(http.HandlerFunc(handler.GetOrdeHandlerById)))

	//Test router
	http.Handle("/test", middleware.AuthMiddleware(http.HandlerFunc(handler.TestHandler)))
	http.Handle("/tests", middleware.AuthMiddleware(http.HandlerFunc(handler.TestGetByIdHandler)))

	//Delivery router
	http.Handle("/delivery", middleware.AuthMiddleware(http.HandlerFunc(handler.DeliveryBoyHandler)))
	http.Handle("/deliverys", middleware.AuthMiddleware(http.HandlerFunc(handler.GetDeliveryBoyByIdHandler)))

	//Categoory router
	http.Handle("/category", middleware.AuthMiddleware(http.HandlerFunc(handler.CategoryHandler)))
	http.Handle("/categorys", middleware.AuthMiddleware(http.HandlerFunc(handler.GetCategoryByIdHandler)))

	//Banner router
	http.Handle("/banner", middleware.AuthMiddleware(http.HandlerFunc(handler.BannerHandler)))
	http.Handle("/banners", middleware.AuthMiddleware(http.HandlerFunc(handler.GetBannerByIdHandler)))

	//Offer router
	http.Handle("/offer", middleware.AuthMiddleware(http.HandlerFunc(handler.OfferHandler)))
	http.Handle("/offers", middleware.AuthMiddleware(http.HandlerFunc(handler.GetOfferByIdHandler)))

	//Address router
	http.Handle("/address", middleware.AuthMiddleware(http.HandlerFunc(handler.AddressHandler)))
	http.Handle("/addresses", middleware.AuthMiddleware(http.HandlerFunc(handler.GetAddressByIdHandler)))

	//Rating router
	http.Handle("/rating", middleware.AuthMiddleware(http.HandlerFunc(handler.RatingHandler)))
	http.Handle("/ratings", middleware.AuthMiddleware(http.HandlerFunc(handler.GetRatingByIdHandler)))

	//Cart router
	http.Handle("/cart", middleware.AuthMiddleware(http.HandlerFunc(handler.CartHandler)))
	http.Handle("/carts", middleware.AuthMiddleware(http.HandlerFunc(handler.GetCartByIdHandler)))

	//Cancel router
	http.Handle("/cancel", middleware.AuthMiddleware(http.HandlerFunc(handler.CancelHandler)))
	http.Handle("/cancels", middleware.AuthMiddleware(http.HandlerFunc(handler.GetCancelByIdHandler)))

	//chat router
	http.Handle("/chat", middleware.AuthMiddleware(http.HandlerFunc(handler.ChatHandler)))
	http.Handle("/chats", middleware.AuthMiddleware(http.HandlerFunc(handler.GetChatByIdHandler)))

	//Payment router
	http.Handle("/payment", middleware.AuthMiddleware(http.HandlerFunc(handler.PaymentHandler)))
	http.Handle("/payments", middleware.AuthMiddleware(http.HandlerFunc(handler.GetPaymentByIdHandler)))

	//Support router
	http.Handle("/support", middleware.AuthMiddleware(http.HandlerFunc(handler.SupportHandler)))
	http.Handle("/supports", middleware.AuthMiddleware(http.HandlerFunc(handler.GetSupportByIdHandler)))

	//Log router
	http.Handle("/log", middleware.AuthMiddleware(http.HandlerFunc(handler.LogHandler)))
	http.Handle("/logs", middleware.AuthMiddleware(http.HandlerFunc(handler.GetLogByIdHandler)))

	http.ListenAndServe(":8080", nil)
}
