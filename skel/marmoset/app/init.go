package appname

import (
	"net/http"

	"github.com/otiai10/marmoset"
	"{{.PackagePath}}/server/controllers"
)

func init() {
	marmoset.LoadViews("./views")
	router := marmoset.NewRouter()
	router.GET("/", controllers.Index)
	router.Static("/public", "./public")
	http.Handle("/", router)
}
