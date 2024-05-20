package textbox

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

type Props struct {
	content string
	url     string
	editing bool
}

type TextBox struct {
	cmp   func(p Props) templ.Component
	props Props
}

func (t TextBox) CreateCmp() templ.Component {
	comp := t.cmp(t.props)
	return comp
}

func New(content string, editing bool, url string) *TextBox {
	p := Props{content: content, editing: editing, url: url}
	cmp := textbox
	tb := TextBox{cmp: cmp, props: p}
	http.HandleFunc(tb.props.url, boundHandler(&tb))
	return &tb
}

func boundHandler(tb *TextBox) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			if r.Form.Has("input") {
				tb.props.content = r.Form.Get("input")
				tb.props.editing = false
				tb.CreateCmp().Render(context.Background(), w)
			}
		}
		if r.Method == http.MethodGet {
			tb.props.editing = true
			tb.CreateCmp().Render(context.Background(), w)
		}
	}
	return handler
}
