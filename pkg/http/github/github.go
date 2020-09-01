package github

// Code from: https://github.com/GitbookIO/go-github-webhook
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type GithubData struct {
	Secret          string
	Fn              func([]byte)
	SecretValidated bool
}

func NewGithubData(secret string, fn func([]byte)) GithubData {

	return GithubData{secret, fn, false}
}

func (g *GithubData) Process(w http.ResponseWriter, req *http.Request) {
	event := req.Header.Get("x-github-event")
	delivery := req.Header.Get("x-github-delivery")
	signature := req.Header.Get("x-hub-signature")

	// Utility funcs
	_fail := func(err error) {
		fail(w, event, err)
	}

	// Ensure headers are all there
	if event == "" || delivery == "" {
		_fail(fmt.Errorf("Missing x-github-* and x-hub-* headers"))
		return
	}

	// No secret provided to github
	if signature == "" && g.Secret != "" {
		_fail(fmt.Errorf("GitHub isn't providing a signature, whilst a secret is being used (please give github's webhook the secret)"))
		return
	}

	// Read body
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		_fail(err)
		return
	}

	// CheckSecret payload (only when secret is provided)
	g.SecretValidated = false
	if g.Secret != "" {
		if err := validePayloadSignature(g.Secret, signature, body); err != nil {
			// Valied validation
			_fail(err)
			return
		}
		g.SecretValidated = true
	}

	if g.SecretValidated {
		g.Fn(body)
	} else {
		log.Printf("\nSECRET NOT VALID: will not run g.Fn(body)\n")
		log.Printf("SECRET LENGTH: %v\n", len(g.Secret))
	}

}

func validePayloadSignature(secret, signatureHeader string, body []byte) error {
	// Check header is valid
	signature_parts := strings.SplitN(signatureHeader, "=", 2)
	if len(signature_parts) != 2 {
		return fmt.Errorf("Invalid signature header: '%s' does not contain two parts (hash type and hash)", signatureHeader)
	}

	// Ensure secret is a sha1 hash
	signature_type := signature_parts[0]
	signature_hash := signature_parts[1]
	if signature_type != "sha1" {
		return fmt.Errorf("Signature should be a 'sha1' hash not '%s'", signature_type)
	}

	// Check that payload came from github
	// skip check if empty secret provided
	if !IsValidPayload(secret, signature_hash, body) {
		return fmt.Errorf("Payload did not come from GitHub")
	}

	return nil
}

func fail(w http.ResponseWriter, event string, err error) {
	w.WriteHeader(500)
	render(w, PayloadPong{
		Ok:    false,
		Event: event,
		Error: err.Error(),
	})
}

func render(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(data)
}
