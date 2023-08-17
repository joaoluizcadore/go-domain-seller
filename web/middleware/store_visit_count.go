package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joaoluizcadore/domain-seller/internal/application"
	"github.com/joaoluizcadore/domain-seller/internal/domain"
)

func StoreVisitCount() gin.HandlerFunc {
	return func(c *gin.Context) {
		visitCountService := application.GetApp().VisitService
		visit := domain.NewVisit(c.Request.Host, c.ClientIP(), nil)
		err := visitCountService.AddVisit(visit)
		if err != nil {
			log.Printf("Error storing visit count: %v", err)
		}
		c.Next()
	}
}
