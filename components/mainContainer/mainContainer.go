package mainContainer

import (
	"context"
	"fmt"
	"htmx/components/textbox"
	"net/http"

	"github.com/a-h/templ"
)

type Props struct {
	tbs []*textbox.TextBox
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

func New() *MainContainer {
	tbs := createMockData()
	p := Props{tbs, "/addItem"}
	mc := MainContainer{cmp: mainContainer, props: p}
	http.HandleFunc(mc.props.url, boundHandler(&mc))
	return &mc
}

func boundHandler(mc *MainContainer) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		thing := append(mc.props.tbs, textbox.New("", true, fmt.Sprintf("/%d", len(mc.props.tbs)+1)))
		mc.props.tbs = thing
		mc.CreateCmp().Render(context.Background(), w)
	}
	return handler
}

func createMockData() []*textbox.TextBox {
	tbs := []*textbox.TextBox{
		textbox.New("This is another ", true, "/1"),
		textbox.New("This is another component", false, "/2"),
		textbox.New("This is another component", false, "/3"),
		textbox.New("This is another component", false, "/4"),
	}
	return tbs
}
