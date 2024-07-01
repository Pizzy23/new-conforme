package sector

import (
	sector "conforme/internal/Sector/service"
	"conforme/internal/interfaces"
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

// @Summary Criar painel
// @Tags Sector
// @Description Criar um novo painel para uma empresa
// @Accept json
// @Produce json
// @Param request body interfaces.PainelInput true "Dados do campo a ser criado"
// @Success 200 {object} db.Painel
// @Failure 500 {object} erros.InternalServerError "Error"
// @Router /create-painels [post]
func CreatePainel(c *gin.Context) {
	var data interfaces.PainelInput

	if err := c.ShouldBindJSON(&data); err != nil {
		c.Set("Error", "Invalid parameters, need a JSON")
		c.Status(http.StatusBadRequest)
		return
	}

	sector.CreatePainelService(c, data)
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

// @Summary Procurar um Painel
// @Tags Company
// @Description Procurar um Painel existente
// @Accept json
// @Produce json
// @Param Painel header string true "Painel que deseja buscar"
// @Success 200 {object} db.Painel
// @Failure 500 {object} erros.InternalServerError "Error"
// @Router /find-painel [get]
func SearchPainel(c *gin.Context) {
	name := c.GetHeader("Painel")
	sector.SearchPanelService(c, name)
}
