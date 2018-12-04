package Easy_web_c_s

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	"os"
)


import (
"net/http"
"os"
"github.com/codegangsta/negroni"
"github.com/gorilla/mux"
"github.com/unrolled/render"
)
func NewServer() *negroni.Negroni {//copy自上课案例
	formatter := render.New()
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}
func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("web server directory error")
		} else {
			webRoot = root
		}
	}
	mx.HandleFunc("/unknown", unknownPage).Methods("GET")
	mx.HandleFunc("/", indexPage).Methods("GET")
	mx.HandleFunc("/index", indexPage).Methods("GET")
	mx.HandleFunc("/calculator", calculatorPage).Methods("GET")
	mx.HandleFunc("/", submit).Methods("POST")
	mx.PathPrefix("/file").Handler(http.StripPrefix("/file/", http.FileServer(http.Dir(webRoot+"/assets/"))))

}