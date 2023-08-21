package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joaoluizcadore/domain-seller/internal/application"
	"github.com/joaoluizcadore/domain-seller/internal/service"
	"github.com/joaoluizcadore/domain-seller/internal/shared"
	"github.com/joaoluizcadore/domain-seller/internal/utils"
)

func IndexAction(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":  "Domínio à venda!",
		"form":   &shared.Message{},
		"domain": utils.CleanDomain(ctx.Request.Host),
	})

}

func SendMessageAction(ctx *gin.Context) {
	var msg shared.Message
	domain := utils.CleanDomain(ctx.Request.Host)

	if err := ctx.ShouldBind(&msg); err != nil {
		ctx.HTML(http.StatusBadRequest, "index.tmpl", gin.H{
			"title":         "Domínio à venda!",
			"error_message": "Atenção: Todos os campos são necessários, verifique!",
			"form":          msg,
			"domain":        domain,
		})
		return
	}

	msg.Domain = domain
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
			"domain":        domain,
		})
		return
	}

	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":           "Domínio à venda!",
		"success_message": "Mensagem enviada com sucesso! Em breve entraremos em contato.",
		"form":            &shared.Message{},
		"domain":          domain,
	})
}

func ShowVisitCountSummaryAction(ctx *gin.Context) {
	visitsService := application.GetApp().VisitService
	visits, err := visitsService.GetListSummary()

	if err != nil {
		ctx.HTML(http.StatusBadRequest, "visit_count_summary.tmpl", gin.H{
			"title":         "Contagem de visitas",
			"error_message": "Ocorreu um erro ao buscar as visitas, tente novamente mais tarde!",
			"visits":        nil,
		})
	} else {
		ctx.HTML(http.StatusOK, "visit_count_summary.tmpl", gin.H{
			"title":  "Contagem de Visitas",
			"visits": visits,
		})
	}
}
