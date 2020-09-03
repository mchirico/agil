package http

import (
	"github.com/mchirico/agil/pkg/graphQL"
	"github.com/mchirico/agil/pkg/graphics"
	"github.com/mchirico/agil/pkg/http/handles"
	"github.com/mchirico/agil/pkg/http/handles/actions"
	"log"
	"net/http"
	_ "time/tzdata"
)

func SetupHandles() {

	handles.Count = 0

	go graphQL.UpdateEveryNtime()

	http.HandleFunc("/cards", func(w http.ResponseWriter, r *http.Request) {
		q := graphQL.QueryGraphQL()
		w.Write([]byte(graphQL.Cards(q)))

	})

	http.HandleFunc("/circle", graphics.Circle)

	http.HandleFunc("/timestampit", handles.TimeStampIt())

	http.HandleFunc("/", handles.BaseRoot)

	p := actions.GithubActions()
	http.HandleFunc("/github", p.Process)

	fs := http.FileServer(http.Dir("/static/dir"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

}

func Server() {
	SetupHandles()
	log.Println("starting server...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
