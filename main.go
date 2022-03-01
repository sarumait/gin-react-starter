package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/sarumait/gin-react-starter/ui"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()

	uiFs, err := uiFs()
	if err != nil {
		log.Err(err)
		os.Exit(1)
	}

	router.StaticFS("/ui", uiFs)
	router.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/ui")
	})

	err = router.Run()
	if err != nil {
		log.Err(err)
		os.Exit(1)
	}
}

func uiFs() (http.FileSystem, error) {
	uiFs, err := ui.FS()
	if err != nil {
		return nil, err
	}
	return http.FS(uiFs), nil
}
