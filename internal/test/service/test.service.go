package test

import (
	"conforme/db"
	"conforme/internal/interfaces"
	"conforme/util/helpers"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllPanel(c *gin.Context) {
	var Sector []db.SectorTest
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
func AllPdfs(c *gin.Context) {
	var Pdfs []db.PdfTest
	engine, ok := helpers.GetDBEngineFromContext(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to DB engine"})
		return
	}
	err := db.GetAll(engine, &Pdfs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	results := make(chan db.PdfTest, len(Pdfs))

	convertToBase64 := func(pdf db.PdfTest, results chan<- db.PdfTest) {
		pdf.Base64 = base64.StdEncoding.EncodeToString(pdf.Content)
		pdf.Content = nil
		results <- pdf
	}

	for _, pdf := range Pdfs {
		go convertToBase64(pdf, results)
	}

	for i := range Pdfs {
		Pdfs[i] = <-results
	}

	c.JSON(http.StatusOK, Pdfs)
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

func CreatePdfService(c *gin.Context, fileName string, contentBytes []byte) {
	engine, ok := helpers.GetDBEngineFromContext(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to DB engine"})
		return
	}

	pdf := db.PdfTest{
		FileName: fileName,
		Content:  contentBytes,
	}

	err := db.Create(engine, &pdf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Response": pdf})
}
