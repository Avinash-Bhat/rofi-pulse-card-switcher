package rofi

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModeOptions(t *testing.T) {
	type args struct {
		options *Mode
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{"empty opts", args{&Mode{}}, ""},
		{"prompt opts", args{&Mode{Prompt: "Hello"}}, "\x00prompt\x1fHello\n"},
		{"message opts", args{&Mode{Message: "Bye!"}}, "\x00message\x1fBye!\n"},
		{"markup-rows opts", args{&Mode{MarkupRows: true}}, "\x00markup-rows\x1ftrue\n"},
		{"urgent opts", args{&Mode{Urgent: "0,2"}}, "\x00urgent\x1f0,2\n"},
		{"active opts", args{&Mode{Active: "1"}}, "\x00active\x1f1\n"},
		{"no-custom opts", args{&Mode{NoCustom: true}}, "\x00no-custom\x1ftrue\n"},
		{"use-hot-keys opts", args{&Mode{UseHotKeys: true}}, "\x00use-hot-keys\x1ftrue\n"},
		{"keep-selection opts", args{&Mode{KeepSelection: true}}, "\x00keep-selection\x1ftrue\n"},
		{"keep-filter opts", args{&Mode{KeepFilter: true}}, "\x00keep-filter\x1ftrue\n"},
		{"new-selection opts", args{&Mode{NewSelection: "1"}}, "\x00new-selection\x1f1\n"},
		{"data opts", args{&Mode{Data: "123456"}}, "\x00data\x1f123456\n"},
		{"theme opts", args{&Mode{Theme: "element-text { background-color: red;}"}}, "\x00theme\x1felement-text { background-color: red;}\n"},

		{"multiple", args{&Mode{
			Prompt:        "Hello",
			Message:       "Bye!",
			MarkupRows:    true,
			Urgent:        "0,2",
			Active:        "1",
			NoCustom:      true,
			UseHotKeys:    true,
			KeepSelection: true,
			KeepFilter:    true,
			NewSelection:  "1",
			Data:          "123456",
			Theme:         "element-text { background-color: red;}",
		}}, "\x00prompt\x1fHello\n" +
			"\x00message\x1fBye!\n" +
			"\x00markup-rows\x1ftrue\n" +
			"\x00urgent\x1f0,2\n" +
			"\x00active\x1f1\n" +
			"\x00no-custom\x1ftrue\n" +
			"\x00use-hot-keys\x1ftrue\n" +
			"\x00keep-selection\x1ftrue\n" +
			"\x00keep-filter\x1ftrue\n" +
			"\x00new-selection\x1f1\n" +
			"\x00data\x1f123456\n" +
			"\x00theme\x1felement-text { background-color: red;}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fmt.Sprint(tt.args.options)
			assert.Equal(t, tt.expected, got, "Row.String() is not what was expected", tt.args)
		})
	}
}
