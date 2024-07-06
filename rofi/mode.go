package rofi

import "bytes"

// Mode is the parsing mode options that needs to be set
type Mode struct {
	Prompt     string // The prompt text.
	Message    string // The message text.
	MarkupRows bool   // If 'true' renders markup in the row.
	// Mark rows as active.
	// It can be of format:
	//
	//	1. an index ('2'),
	//	2. a range of rows (e.g., for last 3 rows: '-3:'
	//	   or for rows from 2 to 4: '2:4' ),
	//	3. a set of rows separated by commas ('1,2,3'),
	//	4.or a combination of all above, like: '2,-2:,3:5'
	Active   string
	Urgent   string // Mark rows as urgent. see: Active for the accepted format
	NoCustom bool   // If set to 'true'; only accept listed entries, ignore custom input.
	// If set to true, it enabled the Custom keybindings for script.
	//
	// **Warning** this breaks the normal rofi flow.
	UseHotKeys bool
	// If set, the selection is not moved to the first entry,
	// but the current position is maintained. The filter is cleared.
	KeepSelection bool
	KeepFilter    bool // If set, the filter is not cleared.
	// If KeepSelection is set, this allows you to override
	//the selected entry (absolute position).
	NewSelection string
	// Passed data to the next execution of the script via
	// the environment variable ROFI_DATA.
	Data string
	// Small theme snippet to f.e. change the background color of a widget.
	//
	// The Theme cannot change the interface while running,
	// it is only usable for small changes in, for example background color,
	// of widgets that get updated during display
	// like the row color of the listview.
	Theme string
}

func (o *Mode) String() string {
	sb := &bytes.Buffer{}
	if o.Prompt != "" {
		appendField(sb, "prompt", o.Prompt)
	}
	if o.Message != "" {
		appendField(sb, "message", o.Message)
	}
	if o.MarkupRows {
		appendField(sb, "markup-rows", o.MarkupRows)
	}
	if o.Urgent != "" {
		appendField(sb, "urgent", o.Urgent)
	}
	if o.Active != "" {
		appendField(sb, "active", o.Active)
	}
	if o.NoCustom {
		appendField(sb, "no-custom", o.NoCustom)
	}
	if o.UseHotKeys {
		appendField(sb, "use-hot-keys", o.UseHotKeys)
	}
	if o.KeepSelection {
		appendField(sb, "keep-selection", o.KeepSelection)
	}
	if o.KeepFilter {
		appendField(sb, "keep-filter", o.KeepFilter)
	}
	if o.NewSelection != "" {
		appendField(sb, "new-selection", o.NewSelection)
	}
	if o.Data != "" {
		appendField(sb, "data", o.Data)
	}
	if o.Theme != "" {
		appendField(sb, "theme", o.Theme)
	}
	return sb.String()
}

func appendField(sb *bytes.Buffer, key string, value interface{}) {
	sb.WriteString(optionStart)
	sb.WriteString(key)
	sb.WriteString(optionSep)
	sb.WriteString(anyToA(value))
	sb.WriteByte('\n')
}
