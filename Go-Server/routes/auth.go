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

// Won't be used (yet)
type GoogleSignUpRequest struct {
	ClientId string `json:"client_id"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Won't be used (yet)
type GoogleSignInRequest struct {
	ClientId string `json:"client_id"`
}

// "/auth"
const route = "/auth"

func AuthRoute(r *gin.Engine) {
	r.GET(route, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "auth is running",
		})
	})

	r.POST(route+"/signin", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "auth is running",
		})
	})

	// Won't be used (yet)
	r.POST(route+"/google-signin", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "auth is running",
		})
	})

	r.POST("/auth/signup", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "auth is running",
		})
	})

	// Won't be used (yet)
	r.POST("/auth/google-signup", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "auth is running",
		})
	})

}
