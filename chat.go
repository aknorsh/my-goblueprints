package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/aknorsh/my-goblueprints/trace"
)

type templateHandler struct {
	once     sync.Once
	filename string
	temp1    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.temp1 = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	err := t.temp1.Execute(w, r)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
	}
}

func main() {

	var addr = flag.String("addr", ":8080", "Address of application")
	var dev = flag.Bool("dev", false, "Activate dev-mode")
	flag.Parse()

	r := NewRoom()
	if *dev {
		r.tracer = trace.New(os.Stdout)
	}

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	go r.run()

	log.Println("Starting web server on ", *addr, "...")
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
