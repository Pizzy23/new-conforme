package company

import (
	company "conforme/internal/Company/service"
	"conforme/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Procurar uma companhia
// @Tags Company
// @Description Procurar uma companhia existente
// @Accept json
// @Produce json
// @Param Company header string true "Company do usuario"
// @Success 200 {object} db.Company
// @Failure 500 {object} erros.InternalServerError "Error"
// @Router /find-company [get]
func SearchCompany(c *gin.Context) {
	name := c.GetHeader("Company")

	company.SearchCompany(c, name)
}

// @Summary Criar uma nova companhia
// @Tags Company
// @Description Criar uma nova companhia
// @Accept json
// @Produce json
// @Param request body interfaces.CompanyInput true "Dados do campo a ser criado"
// @Success 200 {object} db.Company
// @Failure 500 {object} erros.InternalServerError "Error"
// @Router /register-company [post]
func CreateCompany(c *gin.Context) {
	var companyData interfaces.CompanyInput
	if err := c.BindJSON(&companyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	company.CreateCompany(c, companyData)
}
