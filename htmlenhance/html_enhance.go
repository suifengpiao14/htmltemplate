package htmlenhance

import (
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/cbroglie/mustache"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/net/html"
)

/*
功能目标：
为 HTML 模板中的所有非 <script> 节点设置：
一个唯一的 data-node-key
一个动态属性占位符：{{<nodeKey>Attrs}}=""(html_render.go 会写入 {{<nodeKey>Attrs}})
*/

const (
	DataNodeKeyAttr    = "data-node-key"
	attrPlaceholderFmt = "%sAttrs"
	ignoredTagScript   = "script"
)

/*
	InjectNodeIdentityAttributes 为 HTML 模板中每个节点添加唯一 node-key 和属性占位符

使用场景: 新增/修改html模板时，为每个节点注入唯一标识和动态属性占位符，方便后续渲染时动态填充属性值。
*/
func InjectNodeIdentityAttributes(tpl string) (string, error) {
	tpl = strings.TrimSpace(tpl)
	if tpl == "" {
		return "", nil
	}
	root, isFullHTMLDocument, err := parseHTML(tpl)
	if err != nil {
		return "", errors.WithMessagef(err, "failed to parse template")
	}

	nodes := htmlquery.Find(root, "//*")
	for _, node := range nodes {
		if node.Type != html.ElementNode || strings.EqualFold(node.Data, ignoredTagScript) {
			continue
		}

		// 跳过已有 data-node-key 的节点
		if getAttrValue(node, DataNodeKeyAttr) != "" {
			continue
		}

		// 设置唯一 node-key
		nodeKey := uuidWithoutDash()
		setAttr(node, DataNodeKeyAttr, nodeKey)

		// 添加动态属性占位符（Key = {{nodeKeyAttrs}}）
		placeholder := AttrPlaceholderName(nodeKey)
		setAttr(node, placeholder, "")
	}
	return OutputHTML(root, isFullHTMLDocument), nil
}

func AttrPlaceholderName(nodeKey string) string {
	placeholder := fmt.Sprintf("{{%s}}", fmt.Sprintf(attrPlaceholderFmt, nodeKey))
	return placeholder
}

// isFullHTMLDocument 判断是否为完整的 HTML 文档（包含 <!DOCTYPE> 或 <html>）
func isFullHTMLDocument(htmlStr string) bool {
	htmlStr = strings.TrimSpace(strings.ToLower(htmlStr))
	return strings.HasPrefix(htmlStr, "<!doctype") || strings.Contains(htmlStr, "<html")
}

// parseHTML 将模板片段 包裹在 div 中，便于解析多根节点模板
func parseHTML(tpl string) (root *html.Node, isFullHtmlDoc bool, err error) {
	isFullHtmlDoc = isFullHTMLDocument(tpl)
	if isFullHtmlDoc {
		root, err = htmlquery.Parse(strings.NewReader(tpl))
		if err != nil {
			return nil, false, err
		}
		return root, isFullHtmlDoc, nil
	}
	docID := fmt.Sprintf("wrapper-id-%s", uuidWithoutDash())
	wrapped := fmt.Sprintf(`<div id="%s">%s</div>`, docID, tpl)
	doc, err := htmlquery.Parse(strings.NewReader(wrapped))
	if err != nil {
		return nil, false, err
	}
	root = htmlquery.FindOne(doc, fmt.Sprintf(`//*[@id='%s']`, docID))
	if root == nil {
		return nil, false, fmt.Errorf("wrapper node not found")
	}
	return root, isFullHtmlDoc, nil
}

// getAttrValue 获取属性值（忽略大小写）
func getAttrValue(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if strings.EqualFold(attr.Key, key) {
			return attr.Val
		}
	}
	return ""
}

// setAttr 设置属性（如果已存在则更新，否则追加）
func setAttr(node *html.Node, key, val string) {
	for i := range node.Attr {
		if strings.EqualFold(node.Attr[i].Key, key) {
			node.Attr[i].Val = val
			return
		}
	}
	node.Attr = append(node.Attr, html.Attribute{Key: key, Val: val})
}

// uuidWithoutDash 返回无短横线 UUID 字符串
func uuidWithoutDash() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}

func OutputHTML(n *html.Node, self bool) string {
	var b strings.Builder
	if self {
		Render(&b, n)
	} else {
		for n := n.FirstChild; n != nil; n = n.NextSibling {
			Render(&b, n)
		}
	}
	return b.String()
}

//RenderHtmlTpl 渲染html模板，并返回渲染后的字符串。如果传入空字符串则直接返回""

func RenderHtmlTpl(htmlTpl string, context ...any) (renderdHtml string, err error) {
	if htmlTpl == "" {
		return "", nil
	}
	renderd, err := mustache.RenderRaw(htmlTpl, true, context...) //强制渲染为原始字符串，而不是html转义后的字符串
	if err != nil {
		err := errors.WithMessagef(err, "RenderHtmlTpl:%s", htmlTpl)
		return "", err
	}
	return renderd, nil
}
