package user

import (
	"conforme/db"
	"conforme/util/helpers"
	"conforme/util/validations"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func Login(c *gin.Context, email string, password string) {
	engine, ok := helpers.GetDBEngineFromContext(c)
	if !ok {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	resultUser, errUser := FindUserByEmail(engine, email)
	if errUser != nil {
		c.Set("Error", errUser.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if validations.ValidPasswordIt2Pass(resultUser.Password, password) {
		if resultUser.IsLogged {
			c.Set("Error", "User is already connected")
			c.Status(http.StatusInternalServerError)
			return
		}

		resultUser.IsLogged = true
		err := db.Update(engine, &resultUser, resultUser)
		if err != nil {
			c.Set("Error", err.Error())
			c.Status(http.StatusInternalServerError)
			return
		}

		c.Set("Response", resultUser)
		c.Status(http.StatusOK)
		return
	}

	c.Set("Error", "Password Incorrect")
	c.Status(http.StatusNotAcceptable)
}

func Logged(c *gin.Context, email string) {
	engine, ok := helpers.GetDBEngineFromContext(c)
	if !ok {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	resultUser, errUser := FindUserByEmail(engine, email)
	if errUser != nil {
		c.Set("Error", errUser.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if !resultUser.IsLogged {
		c.Set("Error", "User is already disconnected")
		c.Status(http.StatusInternalServerError)
		return
	}

	resultUser.IsLogged = false
	err := db.Update(engine, &resultUser, resultUser)
	if err != nil {
		c.Set("Error", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Set("Response", resultUser)
	c.Status(http.StatusOK)
}

func FindUserByEmail(engine *xorm.Engine, email string) (db.User, error) {
	var user db.User
	found, err := db.GetByName(engine, &user, email)
	if err != nil {
		return db.User{}, err
	}
	if !found {
		return db.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}
