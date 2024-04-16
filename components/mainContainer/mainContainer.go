package mainContainer

import (
	"context"
	"htmx/components/textbox"
	"net/http"

	"github.com/a-h/templ"
)

type Props struct {
	tbs []textbox.TextBox
	url string
}

type MainContainer struct {
	cmp   func(p Props) templ.Component
	props Props
}

func (t MainContainer) CreateCmp() templ.Component {
	comp := t.cmp(t.props)
	return comp
}

func New() MainContainer {
	tbs := []textbox.TextBox{
		textbox.New("This is another component", false, "/1"),
		textbox.New("This is another component", false, "/2"),
		textbox.New("This is another component", false, "/3"),
		textbox.New("This is another component", false, "/4"),
	}
	p := Props{tbs, "/mc"}
	cmp := mainContainer
	mc := MainContainer{cmp: cmp, props: p}
	// http.HandleFunc(mc.props.url, boundHandler(&tb))
	return mc
}

func boundHandler(mc *MainContainer) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		mc.props.tbs = append(mc.props.tbs, textbox.New("", false, "/thing"))
		mc.CreateCmp().Render(context.Background(), w)
	}
	return handler
}
