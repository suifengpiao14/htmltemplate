package htmlenhance_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suifengpiao14/htmltemplate/htmlenhance"
)

func TestMergeClassAttrs(t *testing.T) {
	tpl := `<div class="a b c" class="d e f">hello world</div>`
	newTpl, err := htmlenhance.MergeClassAttrs(tpl)
	require.NoError(t, err)
	excepted := `<div class="a b c d e f">hello world</div>`
	require.Equal(t, excepted, newTpl)

}
