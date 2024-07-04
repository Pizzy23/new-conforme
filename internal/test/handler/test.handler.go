package test

import (
	"conforme/internal/interfaces"
	test "conforme/internal/test/service"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Criar painel Test
// @Tags Test
// @Description Criar um novo painel para uma empresa
// @Accept multipart/form-data
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

	test.CreateSectorTestService(c, data)

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
	test.AllPanel(c)
}

// @Summary Criar PDF No db
// @Tags Test
// @Description Criar um novo painel para uma empresa
// @Accept multipart/form-data
// @Produce json
// @Param content formData file true "Content"
// @Success 200 {object} []db.PdfTest
// @Failure 500 {object} erros.InternalServerError "Error"
// @Router /pdf-test [post]
func CreatePdfTest(c *gin.Context) {
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
	test.CreatePdfService(c, contentBytes)
}

// @Summary Puxa todos os PDFs
// @Tags Test
// @Description Procurar um Painel existente
// @Accept json
// @Produce json
// @Success 200 {object} []db.PdfTest
// @Failure 500 {object} erros.InternalServerError "Error"
// @Router /all-find-pdfs [get]
func SearchAllTestPdfs(c *gin.Context) {
	test.AllPanel(c)
}
