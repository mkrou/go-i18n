package i18n

import (
	"testing"
)

var translate = Map{
	"key1": "val1",
	"subcat1": Map{
		"key2": "val2",
		"subcat2": Map{
			"key3": "val3",
		},
	},
}

func TestT(t *testing.T) {
	if err := AddLanguage("en", translate); err != nil {
		t.Error(err)
	}
	if err := Current([]string{"en"}); err != nil {
		t.Error(err)
	}
	if result := T("subcat1.key2"); result != "val2" {
		t.Error(internal.hash)
	}
}
func TestErrT(t *testing.T) {

}
