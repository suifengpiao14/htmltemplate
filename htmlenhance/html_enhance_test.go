package htmlenhance_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suifengpiao14/htmltemplate/htmlenhance"
)

var htmlTpl = `
<div>
  <span>Hello {{user_name}}</span>
  <img src="{{img_url}}" />
</div>
`

func TestInjectNodeIdentityAttributes(t *testing.T) {
	newHtml, err := htmlenhance.InjectNodeIdentityAttributes(htmlTpl)
	require.NoError(t, err)
	println(newHtml)

}
