package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Name    string `form:"name" json:"name" binding:"required"`
	Email   string `form:"email" json:"email" binding:"required"`
	Message string `form:"message" json:"message" binding:"required"`
	Phone   string `form:"phone" json:"phone" binding:"required"`
}

func IndexAction(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Domínio à venda!",
		"form":  &Message{},
	})
}

func SendMessageAction(ctx *gin.Context) {
	var msg Message

	if err := ctx.ShouldBind(&msg); err != nil {
		ctx.HTML(http.StatusBadRequest, "index.tmpl", gin.H{
			"title":         "Domínio à venda!",
			"error_message": "Atenção: Todos os campos são necessários, verifique!",
			"form":          msg,
		})
	}

	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":           "Domínio à venda!",
		"success_message": "Mensagem enviada com sucesso! Em breve entraremos em contato.",
		"form":            &Message{},
	})

}
