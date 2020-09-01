package graphics

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"github.com/kennygrant/sanitize"
	"net/http"
)

func Circle(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["text"]
	text := ""
	if !ok || len(keys[0]) < 1 {
		text = "No text"
	} else {
		text = keys[0]
	}

	color1 := "blue"
	color2 := "yellow"
	color3 := "black"
	c1, ok := r.URL.Query()["color1"]
	if ok {
		color1 = sanitize.BaseName(c1[0])
	}
	c2, ok := r.URL.Query()["color2"]
	if ok {
		color2 = sanitize.BaseName(c2[0])
	}
	c3, ok := r.URL.Query()["color3"]
	if ok {
		color3 = sanitize.BaseName(c3[0])
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(150, 50)
	s.Circle(25, 25, 12, fmt.Sprintf("fill:%s;stroke:black", color1))
	s.Rect(38+10, 10, 100, 25, fmt.Sprintf("fill:%s;stroke:black", color2))
	s.Text(80+10, 28, text, fmt.Sprintf("text-anchor:middle;font-size:15px;fill:%s", color3))
	s.End()
}
