package lark

// Table displays tabular data (tag: "table").
type Table struct {
	ElementID         string           `json:"element_id,omitempty"`
	Margin            string           `json:"margin,omitempty"`
	PageSize          int              `json:"page_size,omitempty"`
	RowHeight         string           `json:"row_height,omitempty"`
	RowMaxHeight      string           `json:"row_max_height,omitempty"`
	FreezeFirstColumn bool             `json:"freeze_first_column,omitempty"`
	HeaderStyle       *HeaderStyle     `json:"header_style,omitempty"`
	Columns           []TableColumn    `json:"columns,omitempty"`
	Rows              []map[string]any `json:"rows,omitempty"`
}

func (*Table) cardTag() string { return "table" }

// HeaderStyle configures table header appearance.
type HeaderStyle struct {
	TextAlign       string `json:"text_align,omitempty"`
	TextSize        string `json:"text_size,omitempty"`
	BackgroundStyle string `json:"background_style,omitempty"`
	TextColor       string `json:"text_color,omitempty"`
	Bold            *bool  `json:"bold,omitempty"`
	Lines           int    `json:"lines,omitempty"`
}

// TableColumn defines a column in a Table.
type TableColumn struct {
	Name            string        `json:"name,omitempty"`
	DisplayName     string        `json:"display_name,omitempty"`
	Width           string        `json:"width,omitempty"`
	DataType        string        `json:"data_type,omitempty"`
	VerticalAlign   string        `json:"vertical_align,omitempty"`
	HorizontalAlign string        `json:"horizontal_align,omitempty"`
	Format          *NumberFormat `json:"format,omitempty"`
	DateFormat      string        `json:"date_format,omitempty"`
}

// NumberFormat configures number display in table columns.
type NumberFormat struct {
	Precision int    `json:"precision,omitempty"`
	Symbol    string `json:"symbol,omitempty"`
	Separator bool   `json:"separator,omitempty"`
}
