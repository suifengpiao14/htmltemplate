package htmlcomponent_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
)

func TestComponent(t *testing.T) {
	var component = htmlcomponent.Component{
		Name:  "xyxzapi/index",
		Nodes: htmlcomponent.ComponentNodes{},
	}
	data := map[string]any{}
	html, err := component.Render(data)
	require.NoError(t, err)
	fmt.Println(html)
}
