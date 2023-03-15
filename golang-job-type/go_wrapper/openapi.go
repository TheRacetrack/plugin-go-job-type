package main

import (
	"net/http"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
)

func MountOpenApi(router *gin.Engine, baseUrl string) {
	jobName := os.Getenv("JOB_NAME")
	jobVersion := os.Getenv("JOB_VERSION")

	router.Any(baseUrl+"/static/openapi.json", func(c *gin.Context) {
		tmplt := template.New("openapi.json")
		tmplt, _ = tmplt.ParseFiles("./swaggerui/openapi.json")

		context := map[string]string{
			"jobName":    jobName,
			"jobVersion": jobVersion,
		}
		tmplt.Execute(c.Writer, context)
	})
	router.Any(baseUrl+"/static/{filename}", func(c *gin.Context) {
		filename := c.Param("filename")
		http.ServeFile(c.Writer, c.Request, "./swaggerui/"+filename)
	})

	router.Any(baseUrl+"", func(c *gin.Context) {
		http.Redirect(c.Writer, c.Request, c.Request.URL.Path+"/", http.StatusTemporaryRedirect)
	})
	router.Any(baseUrl+"/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "./swaggerui/index.html")
	})
}
