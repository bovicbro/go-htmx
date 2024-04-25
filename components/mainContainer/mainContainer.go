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
		textbox.New("Some text", false, "/1"),
		textbox.New("Some text", false, "/2"),
		textbox.New("Some text", false, "/3"),
		textbox.New("Some text", false, "/4"),
		textbox.New("Some text", false, "/5"),
		textbox.New("Some text", false, "/6"),
		textbox.New("Some text", false, "/7"),
		textbox.New("Some text", false, "/8"),
		textbox.New("Some text", false, "/9"),
		textbox.New("Some text", false, "/10"),
		textbox.New("Some text", false, "/11"),
		textbox.New("Some text", false, "/12"),
		textbox.New("Some text", false, "/13"),
		textbox.New("Some text", false, "/14"),
		textbox.New("Some text", false, "/15"),
		textbox.New("Some text", false, "/16"),
		textbox.New("Some text", false, "/17"),
		textbox.New("Some text", false, "/18"),
		textbox.New("Some text", false, "/19"),
		textbox.New("Some text", false, "/20"),
		textbox.New("Some text", false, "/21"),
	}
	return tbs
}
