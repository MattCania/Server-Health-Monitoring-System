package routes

import (
	"github.com/gin-gonic/gin"
)

// structs
type SignUpRequest struct {
	FirstName       string `json:"firstname"`
	LastName        string `json:"lastname"`
	MiddleName      string `json:"middlename"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type GoogleSignUpRequest struct {
	ClientId string `json:"client_id"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GoogleSignInRequest struct {
	ClientId string `json:"client_id"`
}

// "/auth"
func AuthRoute(r *gin.Engine) {
	r.GET("/auth", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "auth is running",
		})
	})

	r.POST("/auth/signin", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "auth is running",
		})
	})

	r.POST("/auth/google-signin", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "auth is running",
		})
	})

	r.POST("/auth/signup", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "auth is running",
		})
	})

	r.POST("/auth/google-signup", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "auth is running",
		})
	})

}
