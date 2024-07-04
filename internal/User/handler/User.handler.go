package user

import (
	user "conforme/internal/User/service"
	"conforme/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Criar um novo usuario
// @Tags User
// @Description Criar um usuario a partir das novas infos
// @Accept json
// @Produce json
// @Param request body interfaces.UserInput true "Dados do campo a ser criado"
// @Param Office header int true "Tipo de usuário (Gerente da Planta 10, Profissional habilitado responsável 11, Responsável pelo PIE 12, Profissional autorizado 13)" Enums(10,11, 12, 13)
// @Success 200 {object} db.User
// @Failure 500 {object} erros.InternalServerError "Error"
// @Router /register [post]
func CreateUser(c *gin.Context) {
	var userData interfaces.UserInput
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	user.CreateUser(c, userData)
}

// @Summary Procurar Usuario
// @Tags User
// @Description Procurar um usuario existente
// @Accept json
// @Produce json
// @Param RG header string true "Rg do usuario"
// @Success 200 {object} db.User
// @Failure 500 {object} erros.InternalServerError "Error"
// @Router /find-user [get]
func SearchUser(c *gin.Context) {
	rg := c.GetHeader("RG")
	user.SearchUser(c, rg)
}
