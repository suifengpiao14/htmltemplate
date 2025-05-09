package xmldata_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suifengpiao14/htmltemplate/xmldata"
)

func TestDecode(t *testing.T) {

	var xmlData = `

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


`
	mv, err := xmldata.Decode(xmlData)
	require.NoError(t, err)
	jsonData, err := mv.JsonIndent("", "  ")
	require.NoError(t, err)
	s := string(jsonData)
	fmt.Println(s)
}

func TestDecode2(t *testing.T) {
	data := `
	<orders>
	<orderType><![CDATA[黄金到店]]></orderType>
	<orderId><![CDATA[16690403]]></orderId>
	<orderTime><![CDATA[2024-12-26 17:00:07]]></orderTime>
	<className><![CDATA[金条]]></className>
	<productPictrue><![CDATA[/abc.png]]></productPictrue>
	<orderStatus><![CDATA[待跟进]]></orderStatus>
	<price><![CDATA[Y361121.19]]></price>
	<operatorName><![CDATA[超级管理员]]></operatorName>
	<userPhone><![CDATA[159***2341]]></userPhone>
</orders>
<orders>
	<orderType><![CDATA[黄金到店]]></orderType>
	<orderId><![CDATA[16690407]]></orderId>
	<orderTime><![CDATA[2024-05-18 17:00:07]]></orderTime>
	<className><![CDATA[金条]]></className>
	<productPictrue><![CDATA[/abc.png]]></productPictrue>
	<orderStatus><![CDATA[已下单]]></orderStatus>
	<price><![CDATA[Y361121.19]]></price>
	<operatorName><![CDATA[超级管理员]]></operatorName>
	<userPhone><![CDATA[159***2333]]></userPhone>
</orders>
	`
	mv, err := xmldata.Decode(data)
	require.NoError(t, err)
	b, err := mv.JsonIndent("", " ")
	require.NoError(t, err)
	s := string(b)
	fmt.Println(s)

}
