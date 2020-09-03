package handles

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_TimeStampIt(t *testing.T) {
	result := false
	op := func(_f *FN) error {
		_f.fn = func() error {
			result = true
			fmt.Printf("test")
			return nil
		}
		return nil
	}

	TimeStampItHandler := TimeStampIt(op)

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	TimeStampItHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	if resp.StatusCode != 200 {
		t.Fatalf("Bad response status code")
	}

	if !result {
		t.Fatalf("Couldn't swap new function")
	}

}
