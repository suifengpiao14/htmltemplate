package htmlenhance

import (
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
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
	dataNodeKeyAttr    = "data-node-key"
	attrPlaceholderFmt = "%sAttrs"
	ignoredTagScript   = "script"
)

// InjectNodeIdentityAttributes 为 HTML 模板中每个节点添加唯一 node-key 和属性占位符
func InjectNodeIdentityAttributes(tpl string) (string, error) {
	tpl = strings.TrimSpace(tpl)
	if tpl == "" {
		return "", nil
	}

	root, err := parseHTMLWithWrapper(tpl)
	if err != nil {
		return "", errors.WithMessagef(err, "failed to parse template")
	}

	nodes := htmlquery.Find(root, "//*")
	for _, node := range nodes {
		if node.Type != html.ElementNode || strings.EqualFold(node.Data, ignoredTagScript) {
			continue
		}

		// 跳过已有 data-node-key 的节点
		if getAttrValue(node, dataNodeKeyAttr) != "" {
			continue
		}

		// 设置唯一 node-key
		nodeKey := uuidWithoutDash()
		setAttr(node, dataNodeKeyAttr, nodeKey)

		// 添加动态属性占位符（Key = {{nodeKeyAttrs}}）
		placeholder := fmt.Sprintf("{{%s}}", fmt.Sprintf(attrPlaceholderFmt, nodeKey))
		setAttr(node, placeholder, "")
	}
	return OutputHTML(root, false), nil
}

// parseHTMLWithWrapper 将模板包裹在 div 中，便于解析多根节点模板
func parseHTMLWithWrapper(tpl string) (*html.Node, error) {
	docID := fmt.Sprintf("wrapper-id-%s", uuidWithoutDash())
	wrapped := fmt.Sprintf(`<div id="%s">%s</div>`, docID, tpl)
	doc, err := htmlquery.Parse(strings.NewReader(wrapped))
	if err != nil {
		return nil, err
	}
	root := htmlquery.FindOne(doc, fmt.Sprintf(`//*[@id='%s']`, docID))
	if root == nil {
		return nil, fmt.Errorf("wrapper node not found")
	}
	return root, nil
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
