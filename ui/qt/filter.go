package qt

import (
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/ui"
	"github.com/obscuren/qml"
)

func NewFilterFromMap(object map[string]interface{}, eth core.Backend) *core.Filter {
	filter := ui.NewFilterFromMap(object, eth)

	if object["topics"] != nil {
		filter.SetTopics(makeTopics(object["topics"]))
	}

	return filter
}

func makeTopics(v interface{}) (d [][]byte) {
	if qList, ok := v.(*qml.List); ok {
		var s []string
		qList.Convert(&s)

		d = ui.MakeTopics(s)
	} else if str, ok := v.(string); ok {
		d = ui.MakeTopics(str)
	}

	return
}
