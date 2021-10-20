package explorer

import (
	"fmt"
	"github.com/4kang0624/LyricsTranslationAssistant/lyrics"
	"log"
	"net/http"
	"text/template"
)

const (
	port string = ":4000"
	templateDir string = "explorer/templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Lyrics []*lyrics.Lyric
}

func home (rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", lyrics.Lyrics().ShowAllLyrics()}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("lyricsData")
		lyrics.Lyrics().AddLyric(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}


func Start(){
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening in on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}