package xmldata_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suifengpiao14/htmltemplate/xmldata"
)

func TestWrapVariableWithCDATA(t *testing.T) {
	tpl := `<desc>{{hello.world}}</desc><raw>{{{raw_value}}}</raw><inside><![CDATA[ already has {{skip_me}} ]]></inside>`
	newTpl := xmldata.WrapVariableWithCDATA(tpl)
	excepted := `<desc><![CDATA[{{hello.world}}]]></desc><raw><![CDATA[{{{raw_value}}}]]></raw><inside><![CDATA[ already has {{skip_me}} ]]></inside>`
	require.Equal(t, excepted, newTpl)
}
