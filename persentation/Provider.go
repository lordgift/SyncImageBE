package persentation

import (
	"SyncImageBE/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Provider struct {
	service *service.BizService
}

func CreateService() *Provider {
	s := &Provider{
		service: &service.BizService{},
	}
	return s
}
func CreateRouter(p *Provider) *gin.Engine {
	r := gin.Default()
	r.GET("/health", p.health)
	return r
}

func (p *Provider) health(c *gin.Context) {
	err := p.service.SaveFile()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}
