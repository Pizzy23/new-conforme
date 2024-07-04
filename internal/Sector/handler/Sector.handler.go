package sector

import (
	sector "conforme/internal/Sector/service"
	"conforme/internal/interfaces"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Criar setor
// @Tags Sector
// @Description Criar um novo setor para uma empresa
// @Accept json
// @Produce json
// @Param request body interfaces.SectorInput true "Dados do campo a ser criado"
// @Success 200 {object} db.Sector
// @Failure 500 {object} erros.InternalServerError "Error"
// @Router /create-sector [post]
func CreateSector(c *gin.Context) {
	var data interfaces.SectorInput

	if err := c.ShouldBindJSON(&data); err != nil {
		c.Set("Error", "Invalid parameters, need a JSON")
		c.Status(http.StatusBadRequest)
		return
	}

	sector.CreateSectorService(c, data)
}

// @Summary Procurar setor
// @Tags Sector
// @Description Procurar um setor especifico
// @Accept json
// @Produce json
// @Param Company header string true "Company do usuario"
// @Param NameSector header string true "Nome do setor"
// @Param NumberSector header string true "Numero do setor"
// @Success 200 {object} db.Sector
// @Failure 500 {object} erros.InternalServerError "Error"
// @Router /find-sector [get]
func SearchSector(c *gin.Context) {
	nameSector := c.Query("NameSector")
	company := c.Query("Company")
	numberSector := c.Query("NumberSector")

	sector.SearchSectorService(c, nameSector, company, numberSector)
}

// @Summary Criar painel Test
// @Tags Test
// @Description Criar um novo painel para uma empresa
// @Accept json
// @Produce json
// @Param request body interfaces.SectorInputTestNoBytes true "Dados do campo a ser criado"
// @Param content formData file true "Content"
// @Success 200 {object} db.SectorTest
// @Failure 500 {object} erros.InternalServerError "Error"
// @Router /create-painels-test [post]
func CreatePainelTest(c *gin.Context) {
	var input interfaces.SectorInputTestNoBytes

	file, err := c.FormFile("content")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file upload"})
		return
	}

	// Open the file
	fileContent, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer fileContent.Close()

	// Read the file content
	contentBytes, err := ioutil.ReadAll(fileContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	data := interfaces.SectorInputTest{
		Name:        input.Name,
		Number:      input.Number,
		Description: input.Description,
		Content:     contentBytes,
	}

	sector.CreateSectorTestService(c, data)

}

// @Summary Puxa todos os painels
// @Tags Test
// @Description Procurar um Painel existente
// @Accept json
// @Produce json
// @Success 200 {object} db.SectorTest
// @Failure 500 {object} erros.InternalServerError "Error"
// @Router /all-find-painel [get]
func SearchAllTestPainel(c *gin.Context) {
	sector.AllPanel(c)
}
