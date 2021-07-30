package presentation

import (
	"SyncImageBE/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
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
	r.POST("/upload", p.upload)
	return r
}

func (p *Provider) upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uploaded": true,
	})
}
