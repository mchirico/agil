package handles

import (
	"fmt"
	"github.com/mchirico/agil/pkg/graphQL"
	"log"
	"net/http"
	"time"
)

var Count = 0

func BaseRoot(w http.ResponseWriter, r *http.Request) {

	version := "v0.0.4a"
	switch r.Method {
	case "GET":
		msg := fmt.Sprintf("\nversion: %v\ngithub: %v\n", version, Count)
		w.Write([]byte(msg))
	case "POST":
		// msg := fmt.Sprintf("Hello world: POST: %v", r.FormValue("user"))
		w.Write([]byte("post"))
	default:
		w.Write([]byte(`"Sorry, only GET and POST methods are supported."`))
	}

}

type FN struct {
	fn func() error
}

func DefaultFN() *FN {
	_f := &FN{}
	_f.fn = graphQL.TimeStampIT
	return _f
}

func TimeStampIt(options ...func(*FN) error) func(w http.ResponseWriter, r *http.Request) {

	_f := DefaultFN()
	for _, op := range options {
		err := op(_f)
		if err != nil {
			log.Fatalf("Invalid Option Setup")
		}
	}

	timeStampIt := func(w http.ResponseWriter, r *http.Request) {
		_f.fn()

		loc, err := time.LoadLocation("America/New_York")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%v %v", loc, err)))
		}
		w.Write([]byte(time.Now().In(loc).Format("01-02 15:04:05 pm")))

	}
	return timeStampIt
}
