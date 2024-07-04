package NC

import (
	NC "conforme/internal/NC/service"
	"conforme/internal/interfaces"
	erros "conforme/middleware/interfaces/errors"

	"github.com/gin-gonic/gin"
)

// @Summary Inserir NC
// @Tags NC
// @Description Insere um novo registro NC
// @Accept json
// @Produce json
// @Param input body interfaces.NotConformInput true "Dados de entrada para inserção"
// @Success 200 {string} string "File uploaded successfully"
// @Failure 500 {object} erros.InternalServerError
// @Router /insert-nc [post]
func InsertNC(c *gin.Context) {
	var input interfaces.NotConformInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Set("Response", "Paramets is invalid, need a json")
		c.Status(erros.StatusNotAcceptable)
		return
	}
	NC.CreateNC(c, input)
}

// @Summary Obter NC
// @Tags NC
// @Description Obtém registros NC com base nos parâmetros fornecidos
// @Accept json
// @Produce json
// @Param Email header string true "Email do usuario"
// @Success 200 {string} json "{"message": "Result"}"
// @Failure 500 {object} erros.InternalServerError
// @Router /pull-nc [get]
func PullNC(c *gin.Context) {
	number := c.GetHeader("Number")
	NC.PullNC(c, number)
}
