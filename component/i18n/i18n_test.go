package i18n

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {

}

func TestMustLocalize(t *testing.T) {
	tests := []struct {
		name   string
		source string
		result string
		lang   string
	}{
		{
			name:   "testZh",
			source: "测试",
			result: "测试",
			lang:   LANG_ZH,
		},
		{
			name:   "testEN",
			source: "test",
			result: "test",
			lang:   LANG_EN,
		},
	}
	bundle := newI18n(&Config{})
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.result, bundle.MustLocalize(test.lang, test.source))
		})
	}
}
