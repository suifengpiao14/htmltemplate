package xmldata

import (
	"strings"

	"regexp"

	"github.com/cbroglie/mustache"
	"github.com/pkg/errors"
)

func RenderXmlDataTemplate(xmlDataTpl string, context ...any) (xmlData string) {
	if xmlDataTpl == "" {
		return ""
	}
	xmlDataTpl = WrapVariableWithCDATA(xmlDataTpl)                   //确保数据填充后，不会引起xml模板解析错误
	renderd, err := mustache.RenderRaw(xmlDataTpl, true, context...) //强制渲染为原始字符串，而不是html转义后的字符串
	if err != nil {
		err := errors.WithMessagef(err, "tp:%s", xmlDataTpl)
		panic(err)
	}
	return renderd
}

// WrapVariableWithCDATA 包裹 {{var}}/{{{var}}} 变量，避免嵌套 CDATA，变量名支持 a-zA-Z0-9_.。
func WrapVariableWithCDATA(tpl string) string {
	// 匹配 {{var}} 或 {{{var}}}，其中 var 仅包含字母、数字、下划线、句点
	varPattern := regexp.MustCompile(`{{{?([a-zA-Z0-9_.]+)}?}}`)

	// 获取 CDATA 区间
	cdataRanges := findCDATARanges(tpl)

	// 构建结果字符串
	var result strings.Builder
	lastIndex := 0

	for _, match := range varPattern.FindAllStringIndex(tpl, -1) {
		start, end := match[0], match[1]

		// 判断是否位于已有 CDATA 中
		if inAnyRange(start, end, cdataRanges) {
			result.WriteString(tpl[lastIndex:end]) // 原样
		} else {
			result.WriteString(tpl[lastIndex:start])
			result.WriteString("<![CDATA[")
			result.WriteString(tpl[start:end])
			result.WriteString("]]>")
		}
		lastIndex = end
	}

	result.WriteString(tpl[lastIndex:])
	return result.String()
}

// findCDATARanges 找到所有 <![CDATA[...]]> 区段位置
func findCDATARanges(s string) [][2]int {
	var ranges [][2]int
	startTag := "<![CDATA["
	endTag := "]]>"
	offset := 0

	for {
		start := strings.Index(s[offset:], startTag)
		if start == -1 {
			break
		}
		start += offset
		end := strings.Index(s[start:], endTag)
		if end == -1 {
			break
		}
		end = start + end + len(endTag)
		ranges = append(ranges, [2]int{start, end})
		offset = end
	}
	return ranges
}

// inAnyRange 判断位置区间是否在任意已知区间内
func inAnyRange(start, end int, ranges [][2]int) bool {
	for _, r := range ranges {
		if start >= r[0] && end <= r[1] {
			return true
		}
	}
	return false
}
