package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joaoluizcadore/domain-seller/internal/application"
	"github.com/joaoluizcadore/domain-seller/internal/service"
	"github.com/joaoluizcadore/domain-seller/internal/shared"
)

func IndexAction(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Domínio à venda!",
		"form":  &shared.Message{},
	})

}

func SendMessageAction(ctx *gin.Context) {
	var msg shared.Message

	if err := ctx.ShouldBind(&msg); err != nil {
		ctx.HTML(http.StatusBadRequest, "index.tmpl", gin.H{
			"title":         "Domínio à venda!",
			"error_message": "Atenção: Todos os campos são necessários, verifique!",
			"form":          msg,
		})
		return
	}

	msg.Domain = ctx.Request.Host
	msg.IP = ctx.ClientIP()
	now := time.Now()
	msg.Date = &now

	notificationService := service.NewNotificationService()
	err := notificationService.SendNotification(msg)

	if err != nil {
		ctx.HTML(http.StatusBadRequest, "index.tmpl", gin.H{
			"title":         "Domínio à venda!",
			"error_message": "Ocorreu um erro ao enviar a mensagem, tente novamente mais tarde!",
			"form":          msg,
		})
		return
	}

	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":           "Domínio à venda!",
		"success_message": "Mensagem enviada com sucesso! Em breve entraremos em contato.",
		"form":            &shared.Message{},
	})
}

func ShowVisitsAction(ctx *gin.Context) {
	visitsService := application.GetApp().VisitService
	visits, err := visitsService.GetListSummary()

	if err != nil {
		ctx.HTML(http.StatusBadRequest, "visits.tmpl", gin.H{
			"title":         "Visitas",
			"error_message": "Ocorreu um erro ao buscar as visitas, tente novamente mais tarde!",
			"visits":        nil,
		})
	} else {
		ctx.HTML(http.StatusOK, "visits.tmpl", gin.H{
			"title":  "Visitas",
			"visits": visits,
		})
	}
}
