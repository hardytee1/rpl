package controllers

import(
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
    // Delete the cookie by setting its expiration time to a past date
    c.SetCookie("Authorization", "", -1, "", "", false, true)

    // Respond with a success message or redirect the user to a login page
    c.JSON(http.StatusOK, gin.H{
        "message": "Logout successful",
    })
}