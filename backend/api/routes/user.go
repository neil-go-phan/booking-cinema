package routes

import (
	"booking-cinema-backend/api/handler"
	"booking-cinema-backend/api/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userHandler handler.UserHandler
}

func NewUserRoutes(userHandler *handler.UserHandler) *UserRoutes{
	userRoute := &UserRoutes{
		userHandler: *userHandler,
	}
	return userRoute
}

func (userRoutes *UserRoutes)Setup(r *gin.Engine) {
	authRoutes := r.Group("auth")
	{
		authRoutes.GET("check-auth", middlewares.CheckAccessToken(), userRoutes.userHandler.CheckAuth)
		authRoutes.GET("token", middlewares.ExpiredAccessTokenHandler(), userRoutes.userHandler.Token)
		authRoutes.GET("users",middlewares.CheckAccessToken(), userRoutes.userHandler.GetUsers)
		authRoutes.DELETE("delete-user/:username", middlewares.CheckAccessToken(), userRoutes.userHandler.DeleteUser)
		authRoutes.POST("sign-up", userRoutes.userHandler.SignUp)
		authRoutes.POST("sign-in", userRoutes.userHandler.SignIn)
		// authRoutes.PUT("update", middlewares.CheckAccessToken(), userRoutes.userHandler.UpdateUser)
		// authRoutes.GET("search", middlewares.CheckAccessToken(), userRoutes.userHandler.SearchUser)
	}
}