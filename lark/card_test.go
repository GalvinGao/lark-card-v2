package lark

import (
	"encoding/json"
	"testing"
)

func TestSimpleCard(t *testing.T) {
	card := Card{
		Header: Header{
			Title:    PlainText("Deploy Notice"),
			Template: Blue,
		},
		Body: Body{
			Elements: Elements{
				Md("Service **user-api** deployed to production ✅"),
			},
		},
	}

	data, err := card.JSON()
	if err != nil {
		t.Fatal(err)
	}

	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}

	// Check schema
	if m["schema"] != "2.0" {
		t.Errorf("schema = %v, want 2.0", m["schema"])
	}

	// Check header
	header := m["header"].(map[string]any)
	if header["template"] != "blue" {
		t.Errorf("template = %v, want blue", header["template"])
	}
	title := header["title"].(map[string]any)
	if title["tag"] != "plain_text" {
		t.Errorf("title.tag = %v, want plain_text", title["tag"])
	}
	if title["content"] != "Deploy Notice" {
		t.Errorf("title.content = %v, want Deploy Notice", title["content"])
	}

	// Check body elements
	body := m["body"].(map[string]any)
	elements := body["elements"].([]any)
	if len(elements) != 1 {
		t.Fatalf("len(elements) = %d, want 1", len(elements))
	}
	md := elements[0].(map[string]any)
	if md["tag"] != "markdown" {
		t.Errorf("element tag = %v, want markdown", md["tag"])
	}
	if md["content"] != "Service **user-api** deployed to production ✅" {
		t.Errorf("unexpected content: %v", md["content"])
	}
}

func TestComplexCard(t *testing.T) {
	card := Card{
		Config: &Config{
			StreamingMode: true,
			WidthMode:     WidthFill,
			Summary: &Summary{
				Content: "Generating...",
			},
		},
		Header: Header{
			Title:    PlainText("Report"),
			Subtitle: PlainText("2026-03-16"),
			Template: Indigo,
			Icon: &Icon{
				Tag:   StandardIcon,
				Token: "robot_outlined",
				Color: Blue,
			},
			Padding: "4px 8px",
		},
		Body: Body{
			VerticalSpacing: "8px",
			Padding:         "12px",
			Elements: Elements{
				&Markdown{
					Content:   "**Revenue:** ¥1,234,567",
					ElementID: "revenue",
					TextColor: "grey",
				},
				Hr(),
				&ColumnSet{
					Columns: []Column{
						{
							Width:    "weighted",
							Weight:   1,
							Elements: Elements{Md("Left")},
						},
						{
							Width:    "auto",
							Elements: Elements{Btn("Action")},
						},
					},
				},
			},
		},
	}

	data, err := card.JSON()
	if err != nil {
		t.Fatal(err)
	}

	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}

	// Check schema
	if m["schema"] != "2.0" {
		t.Errorf("schema = %v, want 2.0", m["schema"])
	}

	// Check config
	config := m["config"].(map[string]any)
	if config["streaming_mode"] != true {
		t.Errorf("streaming_mode = %v, want true", config["streaming_mode"])
	}
	if config["width_mode"] != "fill" {
		t.Errorf("width_mode = %v, want fill", config["width_mode"])
	}
	summary := config["summary"].(map[string]any)
	if summary["content"] != "Generating..." {
		t.Errorf("summary.content = %v, want Generating...", summary["content"])
	}

	// Check header
	header := m["header"].(map[string]any)
	if header["template"] != "indigo" {
		t.Errorf("template = %v, want indigo", header["template"])
	}
	if header["padding"] != "4px 8px" {
		t.Errorf("padding = %v, want 4px 8px", header["padding"])
	}
	icon := header["icon"].(map[string]any)
	if icon["tag"] != "standard_icon" {
		t.Errorf("icon.tag = %v, want standard_icon", icon["tag"])
	}

	// Check body
	body := m["body"].(map[string]any)
	if body["vertical_spacing"] != "8px" {
		t.Errorf("vertical_spacing = %v, want 8px", body["vertical_spacing"])
	}

	elements := body["elements"].([]any)
	if len(elements) != 3 {
		t.Fatalf("len(elements) = %d, want 3", len(elements))
	}

	// First element: markdown
	md := elements[0].(map[string]any)
	if md["tag"] != "markdown" {
		t.Errorf("element[0].tag = %v, want markdown", md["tag"])
	}
	if md["element_id"] != "revenue" {
		t.Errorf("element[0].element_id = %v, want revenue", md["element_id"])
	}
	if md["text_color"] != "grey" {
		t.Errorf("element[0].text_color = %v, want grey", md["text_color"])
	}

	// Second element: hr
	hr := elements[1].(map[string]any)
	if hr["tag"] != "hr" {
		t.Errorf("element[1].tag = %v, want hr", hr["tag"])
	}

	// Third element: column_set
	cs := elements[2].(map[string]any)
	if cs["tag"] != "column_set" {
		t.Errorf("element[2].tag = %v, want column_set", cs["tag"])
	}
	columns := cs["columns"].([]any)
	if len(columns) != 2 {
		t.Fatalf("len(columns) = %d, want 2", len(columns))
	}

	// Check first column has nested markdown element with tag injected
	col0 := columns[0].(map[string]any)
	col0Elements := col0["elements"].([]any)
	if len(col0Elements) != 1 {
		t.Fatalf("len(col0.elements) = %d, want 1", len(col0Elements))
	}
	nestedMd := col0Elements[0].(map[string]any)
	if nestedMd["tag"] != "markdown" {
		t.Errorf("nested element tag = %v, want markdown", nestedMd["tag"])
	}

	// Check second column has nested button
	col1 := columns[1].(map[string]any)
	col1Elements := col1["elements"].([]any)
	if len(col1Elements) != 1 {
		t.Fatalf("len(col1.elements) = %d, want 1", len(col1Elements))
	}
	btn := col1Elements[0].(map[string]any)
	if btn["tag"] != "button" {
		t.Errorf("nested element tag = %v, want button", btn["tag"])
	}
	btnText := btn["text"].(map[string]any)
	if btnText["content"] != "Action" {
		t.Errorf("button text = %v, want Action", btnText["content"])
	}
}

func TestMustJSON(t *testing.T) {
	card := Card{
		Body: Body{
			Elements: Elements{
				Md("hello"),
			},
		},
	}
	data := card.MustJSON()
	if len(data) == 0 {
		t.Error("MustJSON returned empty data")
	}
}

func TestAllElementTags(t *testing.T) {
	tests := []struct {
		name string
		elem Element
		tag  string
	}{
		{"Div", &Div{}, "div"},
		{"Markdown", &Markdown{}, "markdown"},
		{"Image", &Image{}, "img"},
		{"MultiImageLayout", &MultiImageLayout{}, "img_combination"},
		{"Divider", &Divider{}, "hr"},
		{"Person", &Person{}, "person"},
		{"PersonList", &PersonList{}, "person_list"},
		{"Chart", &Chart{}, "chart"},
		{"Table", &Table{}, "table"},
		{"ColumnSet", &ColumnSet{}, "column_set"},
		{"Form", &Form{}, "form"},
		{"InteractiveContainer", &InteractiveContainer{}, "interactive_container"},
		{"CollapsiblePanel", &CollapsiblePanel{}, "collapsible_panel"},
		{"Button", &Button{}, "button"},
		{"Input", &Input{}, "input"},
		{"SelectStatic", &SelectStatic{}, "select_static"},
		{"MultiSelectStatic", &MultiSelectStatic{}, "multi_select_static"},
		{"SelectPerson", &SelectPerson{}, "select_person"},
		{"MultiSelectPerson", &MultiSelectPerson{}, "multi_select_person"},
		{"DatePicker", &DatePicker{}, "date_picker"},
		{"TimePicker", &TimePicker{}, "picker_time"},
		{"DateTimePicker", &DateTimePicker{}, "picker_datetime"},
		{"Overflow", &Overflow{}, "overflow"},
		{"ImagePicker", &ImagePicker{}, "select_img"},
		{"Checker", &Checker{}, "checker"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.elem.cardTag(); got != tt.tag {
				t.Errorf("%s.cardTag() = %q, want %q", tt.name, got, tt.tag)
			}
		})
	}
}

func TestElementsNil(t *testing.T) {
	var elems Elements
	data, err := elems.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "null" {
		t.Errorf("nil Elements = %s, want null", data)
	}
}

func TestFormCard(t *testing.T) {
	card := Card{
		Body: Body{
			Elements: Elements{
				&Form{
					Name: "myform",
					Elements: Elements{
						&Input{
							Name:        "username",
							Placeholder: PlainText("Enter name"),
						},
						&Button{
							Text:           PlainText("Submit"),
							Type:           ButtonPrimary,
							FormActionType: FormSubmit,
							Name:           "btn_submit",
						},
					},
				},
			},
		},
	}

	data, err := card.JSON()
	if err != nil {
		t.Fatal(err)
	}

	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}

	body := m["body"].(map[string]any)
	elements := body["elements"].([]any)
	form := elements[0].(map[string]any)
	if form["tag"] != "form" {
		t.Errorf("form tag = %v, want form", form["tag"])
	}
	if form["name"] != "myform" {
		t.Errorf("form name = %v, want myform", form["name"])
	}

	formElements := form["elements"].([]any)
	if len(formElements) != 2 {
		t.Fatalf("len(form.elements) = %d, want 2", len(formElements))
	}

	input := formElements[0].(map[string]any)
	if input["tag"] != "input" {
		t.Errorf("input tag = %v, want input", input["tag"])
	}

	btn := formElements[1].(map[string]any)
	if btn["tag"] != "button" {
		t.Errorf("button tag = %v, want button", btn["tag"])
	}
	if btn["form_action_type"] != "submit" {
		t.Errorf("form_action_type = %v, want submit", btn["form_action_type"])
	}
}

func TestOmitEmpty(t *testing.T) {
	card := Card{
		Body: Body{
			Elements: Elements{
				Hr(),
			},
		},
	}

	data, err := card.JSON()
	if err != nil {
		t.Fatal(err)
	}

	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}

	// Config should not be present
	if _, ok := m["config"]; ok {
		t.Error("config should be omitted when nil")
	}
	// card_link should not be present
	if _, ok := m["card_link"]; ok {
		t.Error("card_link should be omitted when nil")
	}
}
