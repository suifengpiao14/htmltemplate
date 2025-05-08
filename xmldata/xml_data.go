package xmldata

import (
	"github.com/clbanning/mxj/v2"
)

func Decode(xmlData string) (mv mxj.Map, err error) {
	if xmlData == "" {
		return nil, nil
	}
	mv, err = mxj.NewMapXml([]byte(xmlData))
	if err != nil {
		return nil, err
	}
	return mv, nil
}
