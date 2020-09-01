package graphQL

import (
	"context"
	"encoding/json"
	"fmt"
	testing_graphql "github.com/mchirico/agil/pkg/fixtures/testing-graphql"
	"github.com/shurcooL/githubv4"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_Cards(t *testing.T) {

	q := QueryGraphQL()
	//fmt.Println(q.Repository.Projects.Edges[0].Node.Columns.Edges[0].Node.Name)
	Columns(q)

	fmt.Printf("%s\n", Cards(q))
}

func Test_Columns(t *testing.T) {

	q := QueryGraphQL()
	//fmt.Println(q.Repository.Projects.Edges[0].Node.Columns.Edges[0].Node.Name)
	Columns(q)

}

func Test_TimeStampIT(t *testing.T) {

	TimeStampIT()
}

func Test_TotalCardsToday(t *testing.T) {
	r := time.Now().Sub(lastUpdate).Seconds()

	fmt.Printf("%v\n", r)
	time.Sleep(3)
	fmt.Printf("%v\n", time.Now().Sub(lastUpdate))

}

func TestClient_Query(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
		if got, want := req.Method, http.MethodPost; got != want {
			t.Errorf("got request method: %v, want: %v", got, want)
		}
		body := mustRead(req.Body)

		q := Q{}
		if err := json.Unmarshal([]byte(body), &q); err != nil {
			t.Errorf("Cannot json.Unmarshal -- This is basic. Must pass")
		}

		w.Header().Set("Content-Type", "application/json")
		mustWrite(w, testing_graphql.Qtest)
	})
	client := githubv4.NewClient(&http.Client{Transport: localRoundTripper{handler: mux}})

	var q Q
	err := client.Query(context.Background(), &q, nil)
	if err != nil {
		t.Fatal(err)
	}

	// This is how you can test
	// Put in your own function..
	r := Cards(q)
	if len(r) != 12813 {
		t.Fatalf("Write size")
	}

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
