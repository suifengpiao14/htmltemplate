package xmldata_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suifengpiao14/htmltemplate/xmldata"
)

func TestParseXml(t *testing.T) {
	mv, err := xmldata.Decode(xmlData)
	require.NoError(t, err)
	jsonData, err := mv.JsonIndent("", "  ")
	require.NoError(t, err)
	s := string(jsonData)
	fmt.Println(s)
}

var xmlData = `
<div>
	<orders>
		<orderType>type1</orderType>
		<orderId>1001</orderId>
		<orderTime>2025-05-08</orderTime>
		<className>A</className>
		<productPictrue>pic1.jpg</productPictrue>
		<orderStatus>shipped</orderStatus>
		<price>88.88</price>
		<operatorName>Alice</operatorName>
		<user>
			<name>first user</name>
			<phone>1111111111</phone>
		</user>
		
		<description><![CDATA[
			<div id="order-1">xml data html</div>
			<script>
			if (1<2){
				console.log("ok")
			}
			</script>
		]]></description>
	</orders>
	<orders>
		<orderType>type2</orderType>
		<orderId>1002</orderId>
		<orderTime>2025-05-09</orderTime>
		<className>B</className>
		<productPictrue>pic2.jpg</productPictrue>
		<orderStatus>pending</orderStatus>
		<price>66.66</price>
		<operatorName>Bob</operatorName>
		<user>
			<name>second user</name>
			<phone>2222222</phone>
		</user>
		<description><![CDATA[
			<div  id="order-2">xml data html</div>
			<script>
			if (1<2){
				console.log("ok")
			}
			</script>
		]]></description>
	</orders>
</div>

`
