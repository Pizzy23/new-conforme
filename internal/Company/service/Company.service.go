package company

import (
	"conforme/db"
	"conforme/internal/interfaces"
	"conforme/util/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateCompany(c *gin.Context, data interfaces.CompanyInput) {
	engine, ok := helpers.GetDBEngineFromContext(c)
	if !ok {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	err := db.Create(engine, data)
	if err != nil {
		c.Set("Error", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Internal Error:", err)
		c.Set("Error", "Internal server error")
		c.Status(http.StatusInternalServerError)
		return
	}

	userData := interfaces.UserInput{
		Name:        data.Name,
		Email:       data.Email,
		Password:    string(hashedPassword),
		CompanyName: data.Company,
		Office:      14,
	}

	dataWithHash := interfaces.UserInputWithHashedPassword{
		UserInput:      userData,
		HashedPassword: string(hashedPassword),
	}

	errU := db.Create(engine, dataWithHash)
	if errU != nil {
		c.Set("Response", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", "Company and owner created")
	c.Status(http.StatusOK)
}

func SearchCompany(c *gin.Context, company string) {
	engine, ok := helpers.GetDBEngineFromContext(c)
	if !ok {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	var result db.Company
	found, err := db.GetByName(engine, &result, company)
	if err != nil {
		c.Set("Response", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if !found {
		c.Set("Response", "Company not found")
		c.Status(http.StatusNotFound)
		return
	}

	c.Set("Response", result)
	c.Status(http.StatusOK)
}
