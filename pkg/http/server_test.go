package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func HandleRun(t *testing.T, url string, expected string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, req)

	status := w.Code
	fmt.Println(status)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Header.Get("Content-Type"))
	if !strings.Contains(string(body), expected) {
		t.Logf("Expected: %v", expected)
		t.Fatalf("\n\n Got: %v\n", string(body))

	}

}

func Test_Handlers(t *testing.T) {
	SetupHandles()
	HandleRun(t, "/", "version:")
	HandleRun(t, "/cards", "/cmd :agil:todo")
	HandleRun(t, "/circle", "Generated by SVGo")
	HandleRun(t, "/static", "Moved Permanently")
	HandleRun(t, "/github", "Missing x-github-* and x-hub-* headers")
	HandleRun(t, "/timestampit", ":")

}
