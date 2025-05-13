package htmlenhance

import (
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

/*
	MergeClassAttrs 解析 HTML 并合并同一元素重复的 class 属性

使用场景：动态渲染html模板后，使用该函数整理html class 确保多次设置的class能都生效
*/
func MergeClassAttrs(htmlStr string) (string, error) {
	root, isFullHTMLDocument, err := ParseHTML(htmlStr)
	if err != nil {
		return "", err
	}

	nodes := htmlquery.Find(root, "//*")
	for _, node := range nodes {
		mergeClassAttrs(node)
	}

	newHtmlStr := OutputHTML(root, isFullHTMLDocument)
	return newHtmlStr, nil
}

// mergeClassAttrs 合并某个节点上的重复 class 属性
func mergeClassAttrs(node *html.Node) {
	classValues := []string{}
	newAttrs := []html.Attribute{}

	for _, attr := range node.Attr {
		if attr.Key == "class" {
			classValues = append(classValues, attr.Val)
			continue
		}
		newAttrs = append(newAttrs, attr)
	}

	if len(classValues) > 0 {
		combined := combineClassValues(classValues)
		newAttrs = append(newAttrs, html.Attribute{Key: "class", Val: combined})
	}
	node.Attr = newAttrs
}

// combineClassValues 去重合并 class 值（保留顺序）
func combineClassValues(values []string) string {
	seen := map[string]struct{}{}
	result := []string{}

	for _, val := range values {
		for _, cls := range strings.Fields(val) {
			if _, ok := seen[cls]; !ok {
				seen[cls] = struct{}{}
				result = append(result, cls)
			}
		}
	}
	return strings.Join(result, " ")
}
