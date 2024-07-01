package user

import (
	"conforme/db"
	"conforme/internal/interfaces"
	"conforme/util/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context, data interfaces.UserInput) {
	engine, ok := helpers.GetDBEngineFromContext(c)
	if !ok {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}
	passwordBytes := []byte(data.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Internal Error:", err)
		c.Set("Error", "Internal server error")
		c.Status(http.StatusInternalServerError)
		return
	}

	dataWithHash := interfaces.UserInputWithHashedPassword{
		UserInput:      data,
		HashedPassword: string(hashedPassword),
	}

	errU := db.Create(engine, dataWithHash)
	if errU != nil {
		c.Set("Response", errU.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", "User Create")
	c.Status(http.StatusOK)
}
