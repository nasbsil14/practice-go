package web

import (
    "log"
    "net/http"
	"sync"
	"html/template"
	"path/filepath"
)

type templateHandler struct {
	once sync.Once
	filename string
	templ *template.Template
}
//上記構造体に定義されたメソッド（レシーバーで紐付き）
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
		template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}


func Server() {

	//アドレス演算子でポインタ渡し
	http.Handle("/", &templateHandler{filename: "index.html"})
    //http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //    w.Write([]byte(`
    //    <html>
    //      <head>
    //        <title>チャット</title>
    //      </head>
    //      <body>
    //        チャットしましょう
    //      </body>
    //    </html>
    //    `))
    //})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
