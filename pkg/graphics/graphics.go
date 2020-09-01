package graphics

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"github.com/kennygrant/sanitize"
	"net/http"
)

func Circle(w http.ResponseWriter, r *http.Request) {

	text := "No text"
	t1, ok := r.URL.Query()["text"]
	if ok {
		text = sanitize.BaseName(t1[0])
	}
	text2 := "No text"
	t2, ok := r.URL.Query()["text2"]
	if ok {
		text2 = sanitize.BaseName(t2[0])
	}

	color1 := "firebrick"
	color2 := "yellow"
	color3 := "black"
	color4 := color1
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
	c4, ok := r.URL.Query()["color4"]
	if ok {
		color4 = sanitize.BaseName(c4[0])
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(300, 50)
	s.Circle(25, 20, 12, fmt.Sprintf("fill:%s;stroke:black", color1))
	s.Circle(31, 20, 6, fmt.Sprintf("fill:%s;stroke:black", color2))
	s.Circle(25, 26, 6, fmt.Sprintf("fill:%s;stroke:black", color2))
	s.Circle(19, 20, 6, fmt.Sprintf("fill:%s;stroke:black", color2))
	s.Circle(25, 14, 6, fmt.Sprintf("fill:%s;stroke:black", color2))
	s.Rect(38+5, 10, 150, 25, fmt.Sprintf("fill:%s;stroke:%s", color2, color2))
	s.Rect(38+5+75, 10, 75, 25, fmt.Sprintf("fill:%s;stroke:%s", color4, color4))
	s.Text(70, 28, text, fmt.Sprintf("text-anchor:middle;font-size:15px;fill:%s", color3))
	s.Text(70+75, 28, text2, fmt.Sprintf("text-anchor:middle;font-size:15px;fill:%s", color2))
	s.End()
}
