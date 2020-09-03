package testing_graphql

import (
	"context"
	"encoding/json"
	"github.com/mchirico/agil/pkg/qtypes"
	"github.com/shurcooL/githubv4"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MockOpt(gh4 *qtypes.GH4) error {

	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
		if got, want := req.Method, http.MethodPost; got != want {
			log.Fatalf("got request method: %v, want: %v", got, want)
		}
		body := mustRead(req.Body)

		q := qtypes.Q{}
		if err := json.Unmarshal([]byte(body), &q); err != nil {
			log.Fatalf("Cannot json.Unmarshal -- This is basic. Must pass")
		}

		w.Header().Set("Content-Type", "application/json")
		mustWrite(w, Qtest())
	})
	gh4.Client = githubv4.NewClient(&http.Client{Transport: localRoundTripper{handler: mux}})

	return nil
}

func MockQueryGraphQL(t *testing.T) qtypes.Q {
	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
		if got, want := req.Method, http.MethodPost; got != want {
			t.Errorf("got request method: %v, want: %v", got, want)
		}
		body := mustRead(req.Body)

		q := qtypes.Q{}
		if err := json.Unmarshal([]byte(body), &q); err != nil {
			t.Errorf("Cannot json.Unmarshal -- This is basic. Must pass")
		}

		w.Header().Set("Content-Type", "application/json")
		mustWrite(w, Qtest())
	})
	client := githubv4.NewClient(&http.Client{Transport: localRoundTripper{handler: mux}})

	var q qtypes.Q
	err := client.Query(context.Background(), &q, nil)
	if err != nil {
		t.Fatal(err)
	}

	return q

}

// localRoundTripper is an http.RoundTripper that executes HTTP transactions
// by using handler directly, instead of going over an HTTP connection.
type localRoundTripper struct {
	handler http.Handler
}

func (l localRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	l.handler.ServeHTTP(w, req)
	return w.Result(), nil
}

func mustRead(r io.Reader) string {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func mustWrite(w io.Writer, s string) {
	_, err := io.WriteString(w, s)
	if err != nil {
		panic(err)
	}
}
