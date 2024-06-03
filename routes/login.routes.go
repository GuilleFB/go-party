package routes

import (
	"github.com/GuilleFB/go-party/handlers"
	"github.com/GuilleFB/go-party/middlewares"
	"github.com/gin-gonic/gin"
)

func RoutesLogin(rgin *gin.Engine) {
	rgin.POST("/gin/auth/signup", handlers.PostUserHandler) // Whit Gin
	rgin.POST("/gin/auth/login", handlers.Login)
	rgin.GET("/gin/user/profile", middlewares.CheckAuth, handlers.GetUserProfile)
}

// func RoutesSignup(r *mux.Router) {
// 	r.HandleFunc("/mux/auth/signup", handlers.PostUserHandler).Methods("POST")
// }