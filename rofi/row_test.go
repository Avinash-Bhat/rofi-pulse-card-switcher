package rofi

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRow(t *testing.T) {
	type args struct {
		options *Row
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{"empty opts", args{&Row{Value: "Hello"}}, "Hello"},
		{"icon", args{&Row{Value: "aap", Icon: "folder"}}, "aap\x00icon\x1ffolder"},
		{"display", args{&Row{Value: "hello", Display: "Aloha"}}, "hello\x00display\x1fAloha"},
		{"meta", args{&Row{Value: "hello", Meta: "lang:Hawaiian"}}, "hello\x00meta\x1flang:Hawaiian"},
		{"nonselectable", args{&Row{Value: "hello", NonSelectable: true}}, "hello\x00nonselectable\x1ftrue"},
		{"nonselectable=false", args{&Row{Value: "hello", NonSelectable: false}}, "hello"},
		{"permanent", args{&Row{Value: "hello", Permanent: true}}, "hello\x00permanent\x1ftrue"},
		{"permanent=false", args{&Row{Value: "hello", Permanent: false}}, "hello"},
		{"info", args{&Row{Value: "hello", Info: "greeting"}}, "hello\x00info\x1fgreeting"},
		{"urgent", args{&Row{Value: "hello", Urgent: true}}, "hello\x00urgent\x1ftrue"},
		{"urgent=false", args{&Row{Value: "hello", Urgent: false}}, "hello"},
		{"active", args{&Row{Value: "hello", Active: true}}, "hello\x00active\x1ftrue"},
		{"active=false", args{&Row{Value: "hello", Active: false}}, "hello"},

		{"multiple", args{&Row{
			Value:         "hello",
			Icon:          "folder",
			Display:       "Aloha",
			Meta:          "lang:Hawaiian",
			NonSelectable: true,
			Permanent:     true,
			Info:          "greeting",
			Urgent:        true,
			Active:        true,
		}}, "hello" +
			"\x00icon\x1ffolder" +
			"\x1fdisplay\x1fAloha" +
			"\x1fmeta\x1flang:Hawaiian" +
			"\x1fnonselectable\x1ftrue" +
			"\x1fpermanent\x1ftrue" +
			"\x1finfo\x1fgreeting" +
			"\x1furgent\x1ftrue" +
			"\x1factive\x1ftrue"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fmt.Sprint(tt.args.options)
			assert.Equal(t, tt.expected, got, "Row.String() is not what was expected", tt.args)
		})
	}
}
