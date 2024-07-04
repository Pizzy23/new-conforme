package sector

import (
	"conforme/db"
	"conforme/internal/interfaces"
	"conforme/util/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchSectorService(c *gin.Context, name string, company string, number string) {

	engine, ok := helpers.GetDBEngineFromContext(c)
	if !ok {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	var newCompany db.Company
	found, err := db.GetByName(engine, &newCompany, company)
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

	var result db.Sector
	found, err = db.SearchSector(engine, &result, name, newCompany.ID, name)
	if err != nil {
		c.Set("Response", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if !found {
		c.Set("Response", "Sector not found")
		c.Status(http.StatusNotFound)
		return
	}

	c.Set("Response", result)
	c.Status(http.StatusOK)
}

func CreatePainelService(c *gin.Context, data interfaces.PainelInput) {

	engine, ok := helpers.GetDBEngineFromContext(c)
	if !ok {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}
	painel := db.Painel{
		SectorID:    data.SectorID,
		Name:        data.Name,
		Number:      data.Number,
		Review:      data.Review,
		Description: data.Description,
		PanelType:   data.PanelType,
		ControlCopy: data.ControlCopy,
	}

	err := db.Create(engine, &painel)
	if err != nil {
		c.Set("Response", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Set("Response", painel)
	c.Status(http.StatusOK)
}

func CreateSectorService(c *gin.Context, data interfaces.SectorInput) {

	engine, ok := helpers.GetDBEngineFromContext(c)
	if !ok {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}
	sector := db.Sector{
		CompanyID:   data.CompanyID,
		Name:        data.Name,
		Number:      data.Number,
		Description: data.Description,
	}

	err := db.Create(engine, &sector)
	if err != nil {
		c.Set("Response", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Set("Response", sector)
	c.Status(http.StatusOK)
}

func SearchPanelService(c *gin.Context, name string) {
	var painel db.Painel
	engine, ok := helpers.GetDBEngineFromContext(c)
	if !ok {
		c.Set("Response", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}
	found, err := db.GetByName(engine, &painel, name)
	if err != nil {
		c.Set("Response", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	if !found {
		c.Set("Response", "Panel not found")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", painel)
	c.Status(http.StatusOK)
}

func AllPanel(c *gin.Context) {
	var Sector db.SectorTest
	engine, ok := helpers.GetDBEngineFromContext(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to DB engine"})
		return
	}
	err := db.GetAll(engine, &Sector)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Sector)
}

func CreateSectorTestService(c *gin.Context, data interfaces.SectorInputTest) {
	var Sector db.SectorTest
	engine, ok := helpers.GetDBEngineFromContext(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to DB engine"})
		return
	}

	Sector = db.SectorTest{
		Name:        data.Name,
		Number:      data.Number,
		Description: data.Description,
		Content:     data.Content,
	}

	err := db.Create(engine, &Sector)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Response": Sector})
}
