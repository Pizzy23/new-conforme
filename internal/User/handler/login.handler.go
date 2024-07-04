package user

import (
	user "conforme/internal/User/service"

	"github.com/gin-gonic/gin"
)

// @Summary Conectar o usuario
// @Tags Login
// @Description Altera o modo de login para true
// @Accept json
// @Produce json
// @Param Email header string true "Email do usuario"
// @Param Password header string true "Senha do usuario"
// @Success 200 {string} json "{"message": "Result"}"
// @Failure 500 {object} erros.InternalServerError "Error"
// @Router /login [put]
func LoginUser(c *gin.Context) {
	email := c.GetHeader("Email")
	password := c.GetHeader("Password")
	user.Login(c, email, password)
}

// @Summary Desconectar o usuario
// @Tags Login
// @Description Altera o modo de login para false
// @Accept json
// @Produce json
// @Param Email header string true "Email do usuario"
// @Success 200 {string} json "{"message": "Result"}"
// @Failure 500 {object} erros.InternalServerError "Error"
// @Router /loggout [put]
func LoggedUser(c *gin.Context) {
	email := c.GetHeader("Email")
	user.Logged(c, email)
}
