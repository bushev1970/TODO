package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"/home/vladislav/Рабочий стол/Projects/ToDo/main/loadtemplates.go"
)

//tmpl, err := template.New("name").Parse(http.Dir(`{{define}}`)
func Start() {
	r := gin.Default()
	r.StaticFS("/css", http.Dir("../static/css"))
	r.StaticFS("/js", http.Dir("../static/js"))
	r.LoadHTMLGlob("../static/templates/*")
	//tmpl := http.Dir("../static/templates/header.html")
	r.GET("/", Home)
	r.GET("/day.html", Day)
	r.GET("/week.html", Week)
	r.GET("/month.html", Month)
	r.GET("/year.html", Year)
	r.GET("/setting.html", Setting)
	r.Run(":3000")
}
