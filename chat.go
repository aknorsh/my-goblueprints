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
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
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

	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	err := t.temp1.Execute(w, data)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
	}
}

func main() {

	var addr = flag.String("addr", ":8080", "Address of application")
	var dev = flag.Bool("dev", false, "Activate dev-mode")
	flag.Parse()

	// Gomniauth setup
	gomniauth.SetSecurityKey("bQsCcStprBsKsmasiUrGwTDxw6wfULy9CW8n")
	gomniauth.WithProviders(
		google.New(
			"226288180806-6pbikfu844v2jhhq18pl4t70qae2oelo.apps.googleusercontent.com",
			"hHY7EgrvJZcXpgA6QFXBTOvC",
			"http://localhost:8080/auth/callback/google"),
	)

	r := NewRoom(UseFileSystemAvatar)
	if *dev {
		r.tracer = trace.New(os.Stdout)
	}

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header()["Location"] = []string{"/chat"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	http.Handle("/upload", &templateHandler{filename: "upload.html"})
	http.HandleFunc("/uploader", uploaderHandler)
	http.Handle("/avatars/",
		http.StripPrefix("/avatars/",
			http.FileServer(http.Dir("./avatars"))))

	go r.run()

	log.Println("Starting web server on ", *addr, "...")
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
