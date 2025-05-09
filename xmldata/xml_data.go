package xmldata

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/clbanning/mxj/v2"
	"github.com/pkg/errors"
)

const (
	xml_root_tag = "root"
)

// 解码xml数据
func Decode(xmlData string) (mv mxj.Map, err error) {
	if xmlData == "" {
		return nil, nil
	}
	xmlData = fmt.Sprintf("<%s>%s</%s>", xml_root_tag, xmlData, xml_root_tag) // 添加根节点，防止解析出错

	mv, err = mxj.NewMapXml([]byte(xmlData))
	if err != nil {
		return nil, err
	}
	mv, err = getValFromMap(mv, xml_root_tag)
	if err != nil {
		return nil, errors.WithMessage(err, "Decode")
	}

	return mv, nil
}

// 从map中获取子map，并返回错误信息
func getValFromMap(mv mxj.Map, path string) (subMv mxj.Map, err error) {
	an, err := mv.ValueForPath(path)
	if err != nil {
		return nil, errors.WithMessage(err, "getValFromMap")
	}
	mv, ok := an.(map[string]any)
	if !ok {
		return nil, errors.Errorf("getValFromMap: %v is not map[string]any", an)
	}
	return mv, nil
}

// 解码数据，支持json和xml格式
func DecodeTplData(dataTpl []byte, data any) (mv mxj.Map, err error) {
	dataTpl = bytes.TrimSpace(dataTpl)
	if len(dataTpl) == 0 {
		return nil, nil
	}

	dateTplByte := []byte(dataTpl)
	isJson := isJsonData(dateTplByte)
	//json 处理
	if isJson {
		mv, err = decodeTplDataJson(dataTpl, data)
		if err != nil {
			err = errors.WithMessage(err, "DecodeTplData decodeTplDataJson")
			return nil, err
		}
		return mv, nil
	}

	xmlData := RenderXmlDataTemplate(string(dataTpl), data)
	mv, err = Decode(xmlData)
	if err != nil {
		err = errors.WithMessage(err, "DecodeTplData RenderXmlDataTemplate")
		return nil, err
	}

	return mv, nil
}

func decodeTplDataJson(dataTpl []byte, data any) (mv mxj.Map, err error) {
	if !hasVarPlaceholder(dataTpl) { // 没有占位符，直接返回json数据
		err = json.Unmarshal(dataTpl, &mv)
		if err != nil {
			err = errors.Wrap(err, "decodeTplDataJson json.Unmarshal")
			return nil, err
		}
		return mv, nil
	}
	// 有占位符，先转为xml再处理占位符替换

	mv, err = mxj.NewMapJson(dataTpl)
	if err != nil {
		err = errors.Wrap(err, "decodeTplDataJson mxj.NewMapJson")
		return nil, err
	}

	dataTpl, err = mv.Xml(xml_root_tag)
	if err != nil {
		err = errors.Wrap(err, "decodeTplDataJson mv.Xml")
		return nil, err
	}

	xmlData := RenderXmlDataTemplate(string(dataTpl), data)
	mv, err = Decode(xmlData)
	if err != nil {
		err = errors.WithMessage(err, "decodeTplDataJson Decode")
		return nil, err
	}
	mv, err = getValFromMap(mv, xml_root_tag)
	if err != nil {
		return nil, errors.WithMessage(err, "decodeTplDataJson")
	}

	return mv, nil
}

// 判断是否为json数据
func isJsonData(data []byte) bool {
	var parsed any
	err := json.Unmarshal([]byte(data), &parsed)
	isJson := err == nil
	return isJson
}

// 判断是否包含变量占位符
func hasVarPlaceholder(data []byte) bool {
	hasVarPlaceHolder := varPlaceholderPattern.Match(data)
	return hasVarPlaceHolder
}
