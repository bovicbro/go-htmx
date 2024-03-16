package textboxList

import (
	"context"
	"github.com/a-h/templ"
	"htmx/components/textbox"
	"net/http"
)

type Props struct {
	url string
	tbs []*textbox.TextBox
}

type TextboxList struct {
	cmp   func(p Props) templ.Component
	props Props
}

func New(url string) TextboxList {
	p := Props{url: url}
	cmp := textboxList
	tbl := TextboxList{cmp: cmp, props: p}
	return tbl
}

func (t TextboxList) CreateCmp() templ.Component {
	comp := t.cmp(t.props)
	return comp
}

func (tbl TextboxList) addTextbox(tb *textbox.TextBox) {
	tbl.props.tbs = append(tbl.props.tbs, tb)
}

func (tbl TextboxList) deleteTextbox(index int) {
	tbl.props.tbs = append(tbl.props.tbs[:index], tbl.props.tbs[index+1:]...)
}

func boundHandler(tbs *TextboxList) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tbs.CreateCmp().Render(context.Background(), w)
		}
	}
	return handler
}
