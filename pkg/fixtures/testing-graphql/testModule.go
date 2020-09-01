package testing_graphql

import (
	"context"
	"encoding/json"
	"github.com/mchirico/agil/pkg/qtypes"
	"github.com/shurcooL/githubv4"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
	// This is how you can test
	// Put in your own function..
	//r := Cards(q)
	//if len(r) != 12874 {
	//	t.Fatalf("Did you pick up the correct file?" +
	//		"\nThis is the wrong size: %v\n",len(r))
	//}

	//var want Q
	//want.Repository.Projects.Edges. = "gopher"
	//want.Viewer.Biography = "The Go gopher."
	//if !reflect.DeepEqual(got, want) {
	//	t.Errorf("client.Query got: %v, want: %v", got, want)
	//}
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
