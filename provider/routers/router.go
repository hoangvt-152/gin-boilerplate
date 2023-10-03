package routers

import (
	"gin-boilerplate/routers/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginI18n "github.com/gin-contrib/i18n"
	"golang.org/x/text/language"
	"encoding/json"
	"os"
	"fmt"

)

func SetupRoute() *gin.Engine {

	environment := viper.GetBool("DEBUG")
	if environment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)  //

	allowedHosts := viper.GetString("ALLOWED_HOSTS")
	router := gin.New()
	router.SetTrustedProxies([]string{allowedHosts})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(ginI18n.Localize(ginI18n.WithBundle(&ginI18n.BundleCfg{
		RootPath:         "/root",
		AcceptLanguage:   []language.Tag{language.German, language.English},
		DefaultLanguage:  language.English,
		UnmarshalFunc:    json.Unmarshal,
		FormatBundleFile: "json",
	  })))
	
	router.Use(middleware.CORSMiddleware())

	RegisterRoutes(router) //routes register

	return router
}
