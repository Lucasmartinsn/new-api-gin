package handlers

import (
	"strconv"

	models "github.com/Lucasmartinsn/new-api-gin/models/records"
	"github.com/gin-gonic/gin"
)

func Upfoto(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"Error":   true,
			"Mensage": "erro ao fazer parse do id %v" + err.Error(),
		})
		return
	}

	var records models.Putfoto

	err = c.ShouldBindJSON(&records)
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   true,
			"Mensage": "erro ao fazer decode do json: %v" + err.Error(),
		})
		return
	}

	rows, err := models.Upfoto(int64(id), records)
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   true,
			"Mensage": "erro ao atualizar register: %v" + err.Error(),
		})
		return
	}

	if rows > 1 {
		c.JSON(400, gin.H{
			"Error":   true,
			"Mensage": "erro: foram atualizador %d registros",
		})
		return
	}

	resp := map[string]any{
		"Error":   false,
		"mensage": "dados atualizados com sucesso",
	}

	c.JSON(200, resp)

}
