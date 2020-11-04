package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/hoisie/mustache"
	log "github.com/sirupsen/logrus"
)

func Render(c *gin.Context, filePath string, obj map[string]interface{}) {
	//templateData, err := Asset(filePath)

	templateData, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.WithFields(log.Fields{
			"path": filePath,
		}).Error("Could not find template file")
		c.String(500, "Template not found")
	} else {
		for k, v := range c.Keys {
			obj[k] = v
		}

		//mainTemplate, _ := Asset("main.html")
		mainTemplate, err := ioutil.ReadFile("main.html")
		if err != nil {
			log.Fatalln(err)
		}
		html := mustache.RenderInLayout(string(templateData), string(mainTemplate), obj)

		if c.Writer.Status() == 200 {
			c.Status(200)
		}
		c.Writer.Write([]byte(html))
	}
}

func Serve(bind string) {
	var router = gin.New()
	router.Use(gin.Recovery())
	router.GET("/", Overview)
	router.Run(bind)
}
