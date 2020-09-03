package github

type GithubData struct {
	Secret          string
	Fn              func([]byte)
	SecretValidated bool
}

type PayloadPong struct {
	Ok    bool   `json:"ok"`
	Event string `json:"event"`
	Error string `json:"error,omitempty"`
}
