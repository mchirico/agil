package http

import (
	"fmt"
	"github.com/mchirico/agil/pkg/graphQL"
	"github.com/mchirico/agil/pkg/graphics"
	"github.com/mchirico/agil/pkg/http/github"
	"log"
	"net/http"
	"os"
	"time"
	_ "time/tzdata"
)

func Static() {

	log.Println("starting server...")
	count := 0

	go graphQL.UpdateEvery10min()

	http.HandleFunc("/cards", func(w http.ResponseWriter, r *http.Request) {
		q := graphQL.QueryGraphQL()
		w.Write([]byte(graphQL.Cards(q)))

	})

	http.Handle("/circle", http.HandlerFunc(graphics.Circle))

	http.HandleFunc("/timestampit", func(w http.ResponseWriter, r *http.Request) {
		graphQL.TimeStampIT()

		loc, err := time.LoadLocation("America/New_York")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%v %v", loc, err)))
		}
		w.Write([]byte(time.Now().In(loc).Format("01-02 15:04:05 pm")))

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		version := "v0.0.3a"
		switch r.Method {
		case "GET":
			msg := fmt.Sprintf("\nversion: %v\ngithub: %v\n", version, count)
			w.Write([]byte(msg))
		case "POST":
			msg := fmt.Sprintf("Hello world: POST: %v", r.FormValue("user"))
			w.Write([]byte(msg))
		default:
			w.Write([]byte(`"Sorry, only GET and POST methods are supported."`))
		}

	})

	var secret = os.Getenv("GITHUB_WEBHOOK_SECRET")
	g := github.GithubData{secret, func(payload *github.GitHubPayload) {
		count += 1

	}}
	http.HandleFunc("/github", g.Process)

	fs := http.FileServer(http.Dir("/static/dir"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	log.Fatal(http.ListenAndServe(":3000", nil))

}
