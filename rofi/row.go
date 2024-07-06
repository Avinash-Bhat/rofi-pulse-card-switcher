package rofi

import (
	"bytes"
	"fmt"
)

// Row represents the row options based on the [docs]
//
// [docs]: https://davatorium.github.io/rofi/current/rofi-script.5/#parsing-row-options
type Row struct {
	Value         string
	Icon          string // Set the icon for this row.
	Display       string // Replace the displayed string.
	Meta          string // Specify invisible search terms used for filtering.
	NonSelectable bool   // If true the row cannot be activated.
	Permanent     bool   // If true the row always shows, independent of filter.
	Info          string // Info that, on selection, gets placed in the `ROFI_INFO` environment variable. This entry does not get searched for filtering.
	Urgent        bool   // Set urgent flag on entry
	Active        bool   // Set active flag on entry
}

// String formats the given text based on the row options
func (o *Row) String() string {
	sb := &bytes.Buffer{}
	optStart := true
	sb.WriteString(o.Value)
	optStart = appendVal(sb, optStart, "icon", o.Icon)
	optStart = appendVal(sb, optStart, "display", o.Display)
	optStart = appendVal(sb, optStart, "meta", o.Meta)
	optStart = appendVal(sb, optStart, "nonselectable", o.NonSelectable)
	optStart = appendVal(sb, optStart, "permanent", o.Permanent)
	optStart = appendVal(sb, optStart, "info", o.Info)
	optStart = appendVal(sb, optStart, "urgent", o.Urgent)
	optStart = appendVal(sb, optStart, "active", o.Active)
	return sb.String()
}

func appendVal(sb *bytes.Buffer, appendStart bool, key string, value interface{}) bool {
	b, ok := value.(bool)
	if ok && !b {
		return appendStart
	}
	str, ok := value.(string)
	if ok && str != "" || !ok {
		if appendStart {
			sb.WriteString(fmt.Sprint(optionStart))
			appendStart = false
		} else {
			sb.WriteString(optionSep)
		}
		sb.WriteString(key)
		sb.WriteString(optionSep)
		sb.WriteString(anyToA(value))
	}
	return appendStart
}
