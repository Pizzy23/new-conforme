package NC

import (
	"conforme/db"
	"conforme/internal/interfaces"
	"conforme/util/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateNC(c *gin.Context, input interfaces.NotConformInput) {
	engine, ok := helpers.GetDBEngineFromContext(c)
	if !ok {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}
	NC := db.NotConform{
		Number:      input.Number,
		Title:       input.Title,
		Desc:        input.Desc,
		Tech:        input.Tech,
		Legal:       input.Legal,
		Recommended: input.Recommended,
	}
	err := db.Create(engine, NC)
	if err != nil {
		c.Set("Error", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", "NC Create :) ")
	c.Status(http.StatusOK)
}

func PullNC(c *gin.Context, input string) {
	var NC db.NotConform
	engine, ok := helpers.GetDBEngineFromContext(c)
	if !ok {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	id, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		c.Set("Error", "Invalid ID format")
		c.Status(http.StatusBadRequest)
		return
	}

	found, err := db.GetByID(engine, &NC, id)
	if err != nil {
		c.Set("Response", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if !found {
		c.Set("Response", "Not Conformity not found")
		c.Status(http.StatusNotFound)
		return
	}

	c.Set("Response", NC)
	c.Status(http.StatusOK)
}
