package lark

import (
	"encoding/json"
	"strings"
	"testing"
)

// ---- helpers for tests ----

// mustMarshalCard marshals a card and returns the parsed map.
func mustMarshalCard(t *testing.T, card Card) map[string]any {
	t.Helper()
	data, err := card.JSON()
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	return m
}

// mustMarshalElements marshals an Elements slice and returns the parsed array.
func mustMarshalElements(t *testing.T, elems Elements) []any {
	t.Helper()
	data, err := elems.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	var arr []any
	if err := json.Unmarshal(data, &arr); err != nil {
		t.Fatal(err)
	}
	return arr
}

// getBodyElements extracts body.elements from a marshaled card map.
func getBodyElements(t *testing.T, m map[string]any) []any {
	t.Helper()
	body := m["body"].(map[string]any)
	return body["elements"].([]any)
}

// ---- Basic card tests ----

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

	m := mustMarshalCard(t, card)

	if m["schema"] != "2.0" {
		t.Errorf("schema = %v, want 2.0", m["schema"])
	}

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

	elements := getBodyElements(t, m)
	if len(elements) != 1 {
		t.Fatalf("len(elements) = %d, want 1", len(elements))
	}
	md := elements[0].(map[string]any)
	if md["tag"] != "markdown" {
		t.Errorf("element tag = %v, want markdown", md["tag"])
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

	m := mustMarshalCard(t, card)

	if m["schema"] != "2.0" {
		t.Errorf("schema = %v, want 2.0", m["schema"])
	}

	config := m["config"].(map[string]any)
	if config["streaming_mode"] != true {
		t.Errorf("streaming_mode = %v, want true", config["streaming_mode"])
	}
	if config["width_mode"] != "fill" {
		t.Errorf("width_mode = %v, want fill", config["width_mode"])
	}

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

	body := m["body"].(map[string]any)
	if body["vertical_spacing"] != "8px" {
		t.Errorf("vertical_spacing = %v, want 8px", body["vertical_spacing"])
	}

	elements := body["elements"].([]any)
	if len(elements) != 3 {
		t.Fatalf("len(elements) = %d, want 3", len(elements))
	}

	md := elements[0].(map[string]any)
	if md["tag"] != "markdown" {
		t.Errorf("element[0].tag = %v, want markdown", md["tag"])
	}
	if md["element_id"] != "revenue" {
		t.Errorf("element[0].element_id = %v, want revenue", md["element_id"])
	}

	hr := elements[1].(map[string]any)
	if hr["tag"] != "hr" {
		t.Errorf("element[1].tag = %v, want hr", hr["tag"])
	}

	cs := elements[2].(map[string]any)
	if cs["tag"] != "column_set" {
		t.Errorf("element[2].tag = %v, want column_set", cs["tag"])
	}
	columns := cs["columns"].([]any)
	if len(columns) != 2 {
		t.Fatalf("len(columns) = %d, want 2", len(columns))
	}
}

func TestMustJSON(t *testing.T) {
	card := Card{
		Body: Body{
			Elements: Elements{Md("hello")},
		},
	}
	data := card.MustJSON()
	if len(data) == 0 {
		t.Error("MustJSON returned empty data")
	}
}

// ---- Tag injection tests ----

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

// Verify tag injection actually works in marshaled JSON for each element type.
func TestTagInjectionInJSON(t *testing.T) {
	tests := []struct {
		name string
		elem Element
		tag  string
	}{
		{"Div", &Div{Text: PlainText("hello")}, "div"},
		{"Markdown", &Markdown{Content: "test"}, "markdown"},
		{"Image", &Image{ImgKey: "key"}, "img"},
		{"Divider", &Divider{}, "hr"},
		{"Button", &Button{Text: PlainText("ok")}, "button"},
		{"Input", &Input{Name: "x"}, "input"},
		{"Chart", &Chart{AspectRatio: "16:9"}, "chart"},
		{"Table", &Table{PageSize: 5}, "table"},
		{"Checker", &Checker{Name: "c"}, "checker"},
		{"Overflow", &Overflow{Width: "fill"}, "overflow"},
		{"DatePicker", &DatePicker{Name: "d"}, "date_picker"},
		{"TimePicker", &TimePicker{Name: "t"}, "picker_time"},
		{"DateTimePicker", &DateTimePicker{Name: "dt"}, "picker_datetime"},
		{"SelectStatic", &SelectStatic{Name: "s"}, "select_static"},
		{"MultiSelectStatic", &MultiSelectStatic{Name: "ms"}, "multi_select_static"},
		{"SelectPerson", &SelectPerson{Name: "sp"}, "select_person"},
		{"MultiSelectPerson", &MultiSelectPerson{Name: "msp"}, "multi_select_person"},
		{"Person", &Person{UserID: "u1"}, "person"},
		{"PersonList", &PersonList{}, "person_list"},
		{"MultiImageLayout", &MultiImageLayout{CombinationMode: CombinationDouble}, "img_combination"},
		{"ImagePicker", &ImagePicker{Name: "ip"}, "select_img"},
		{"Form", &Form{Name: "f"}, "form"},
		{"ColumnSet", &ColumnSet{}, "column_set"},
		{"InteractiveContainer", &InteractiveContainer{}, "interactive_container"},
		{"CollapsiblePanel", &CollapsiblePanel{}, "collapsible_panel"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elems := Elements{tt.elem}
			arr := mustMarshalElements(t, elems)
			if len(arr) != 1 {
				t.Fatalf("expected 1 element, got %d", len(arr))
			}
			obj := arr[0].(map[string]any)
			if obj["tag"] != tt.tag {
				t.Errorf("tag = %v, want %q", obj["tag"], tt.tag)
			}
		})
	}
}

// ---- Elements edge cases ----

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

func TestElementsEmpty(t *testing.T) {
	elems := Elements{}
	data, err := elems.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "[]" {
		t.Errorf("empty Elements = %s, want []", data)
	}
}

func TestElementsSingleElement(t *testing.T) {
	elems := Elements{Hr()}
	arr := mustMarshalElements(t, elems)
	if len(arr) != 1 {
		t.Fatalf("len = %d, want 1", len(arr))
	}
	if arr[0].(map[string]any)["tag"] != "hr" {
		t.Error("expected hr tag")
	}
}

// ---- omitempty tests ----

func TestOmitEmpty(t *testing.T) {
	card := Card{
		Body: Body{
			Elements: Elements{Hr()},
		},
	}

	m := mustMarshalCard(t, card)

	if _, ok := m["config"]; ok {
		t.Error("config should be omitted when nil")
	}
	if _, ok := m["card_link"]; ok {
		t.Error("card_link should be omitted when nil")
	}
}

func TestOmitEmptyFields(t *testing.T) {
	// Verify that zero-value fields don't appear in JSON.
	elems := Elements{&Button{}}
	arr := mustMarshalElements(t, elems)
	btn := arr[0].(map[string]any)

	// These should all be omitted when empty/zero.
	for _, key := range []string{"element_id", "margin", "type", "size", "width", "name", "disabled", "url"} {
		if _, ok := btn[key]; ok {
			t.Errorf("expected %q to be omitted from empty Button", key)
		}
	}
	// tag should always be present (injected).
	if btn["tag"] != "button" {
		t.Error("tag should be present")
	}
}

// ---- Config tests ----

func TestConfigAllFields(t *testing.T) {
	enableFwd := true
	updateMulti := false
	card := Card{
		Config: &Config{
			StreamingMode: true,
			StreamingConfig: map[string]any{
				"print_frequency_ms": 500,
			},
			Summary: &Summary{
				Content: "Summary text",
				I18nContent: map[string]string{
					"zh_cn": "摘要",
					"en_us": "Summary text",
				},
			},
			Locales:                  []string{"zh_cn", "en_us"},
			EnableForward:            &enableFwd,
			UpdateMulti:              &updateMulti,
			WidthMode:                WidthCompact,
			UseCustomTranslation:     true,
			EnableForwardInteraction: true,
			Style: &Style{
				TextSize: map[string]TextSizeConfig{
					"heading": {
						Default: "heading-0",
						PC:      "heading-1",
						Mobile:  "heading-2",
					},
				},
				Color: map[string]ColorConfig{
					"text_color": {
						LightMode: "#000000",
						DarkMode:  "#ffffff",
					},
				},
			},
		},
		Body: Body{
			Elements: Elements{Md("test")},
		},
	}

	m := mustMarshalCard(t, card)
	config := m["config"].(map[string]any)

	if config["streaming_mode"] != true {
		t.Error("streaming_mode should be true")
	}
	if config["width_mode"] != "compact" {
		t.Errorf("width_mode = %v, want compact", config["width_mode"])
	}
	if config["enable_forward"] != true {
		t.Error("enable_forward should be true")
	}
	if config["update_multi"] != false {
		t.Error("update_multi should be false")
	}
	if config["use_custom_translation"] != true {
		t.Error("use_custom_translation should be true")
	}
	if config["enable_forward_interaction"] != true {
		t.Error("enable_forward_interaction should be true")
	}

	locales := config["locales"].([]any)
	if len(locales) != 2 {
		t.Errorf("locales length = %d, want 2", len(locales))
	}

	summary := config["summary"].(map[string]any)
	i18n := summary["i18n_content"].(map[string]any)
	if i18n["zh_cn"] != "摘要" {
		t.Errorf("zh_cn summary = %v", i18n["zh_cn"])
	}

	style := config["style"].(map[string]any)
	if style["text_size"] == nil {
		t.Error("style.text_size missing")
	}
	if style["color"] == nil {
		t.Error("style.color missing")
	}
}

// ---- CardLink test ----

func TestCardLink(t *testing.T) {
	card := Card{
		CardLink: &CardLink{
			URL:        "https://example.com",
			PCURL:      "https://pc.example.com",
			IOSURL:     "https://ios.example.com",
			AndroidURL: "https://android.example.com",
		},
		Body: Body{
			Elements: Elements{Md("link card")},
		},
	}

	m := mustMarshalCard(t, card)
	cl := m["card_link"].(map[string]any)
	if cl["url"] != "https://example.com" {
		t.Errorf("url = %v", cl["url"])
	}
	if cl["pc_url"] != "https://pc.example.com" {
		t.Errorf("pc_url = %v", cl["pc_url"])
	}
	if cl["ios_url"] != "https://ios.example.com" {
		t.Errorf("ios_url = %v", cl["ios_url"])
	}
	if cl["android_url"] != "https://android.example.com" {
		t.Errorf("android_url = %v", cl["android_url"])
	}
}

// ---- Header tests ----

func TestHeaderTextTagList(t *testing.T) {
	card := Card{
		Header: Header{
			Title:    PlainText("Status"),
			Template: Green,
			TextTagList: []TextTag{
				{Tag: "plain_text", Text: PlainText("Tag1"), Color: Blue},
				{Tag: "plain_text", Text: PlainText("Tag2"), Color: Red},
				{Tag: "plain_text", Text: PlainText("Tag3"), Color: Green},
			},
		},
		Body: Body{
			Elements: Elements{Md("body")},
		},
	}

	m := mustMarshalCard(t, card)
	header := m["header"].(map[string]any)
	tags := header["text_tag_list"].([]any)
	// Lark docs say max 3 text tags are displayed.
	if len(tags) != 3 {
		t.Fatalf("text_tag_list len = %d, want 3", len(tags))
	}
	tag0 := tags[0].(map[string]any)
	if tag0["color"] != "blue" {
		t.Errorf("tag[0].color = %v, want blue", tag0["color"])
	}
}

func TestHeaderI18nTextTagList(t *testing.T) {
	card := Card{
		Header: Header{
			Title: PlainText("Status"),
			I18nTextTagList: map[string][]TextTag{
				"en_us": {
					{Tag: "plain_text", Text: PlainText("English")},
				},
				"zh_cn": {
					{Tag: "plain_text", Text: PlainText("中文")},
				},
			},
		},
		Body: Body{
			Elements: Elements{Md("body")},
		},
	}

	m := mustMarshalCard(t, card)
	header := m["header"].(map[string]any)
	i18n := header["i18n_text_tag_list"].(map[string]any)
	enTags := i18n["en_us"].([]any)
	if len(enTags) != 1 {
		t.Fatalf("en_us tags len = %d, want 1", len(enTags))
	}
}

func TestHeaderSubtitleAndIcon(t *testing.T) {
	card := Card{
		Header: Header{
			Title:    MdText("**Bold Title**"),
			Subtitle: PlainText("Subtitle here"),
			Template: Carmine,
			Icon: &Icon{
				Tag:    CustomIcon,
				ImgKey: "img_abc123",
			},
		},
		Body: Body{
			Elements: Elements{Md("body")},
		},
	}

	m := mustMarshalCard(t, card)
	header := m["header"].(map[string]any)
	if header["template"] != "carmine" {
		t.Errorf("template = %v, want carmine", header["template"])
	}
	subtitle := header["subtitle"].(map[string]any)
	if subtitle["content"] != "Subtitle here" {
		t.Errorf("subtitle content = %v", subtitle["content"])
	}
	icon := header["icon"].(map[string]any)
	if icon["tag"] != "custom_icon" {
		t.Errorf("icon tag = %v, want custom_icon", icon["tag"])
	}
	if icon["img_key"] != "img_abc123" {
		t.Errorf("icon img_key = %v", icon["img_key"])
	}
}

// ---- Helper function tests ----

func TestHelperMd(t *testing.T) {
	md := Md("**bold**")
	if md.Content != "**bold**" {
		t.Errorf("Md content = %q", md.Content)
	}
	if md.cardTag() != "markdown" {
		t.Error("Md should create Markdown element")
	}
}

func TestHelperPlainText(t *testing.T) {
	txt := PlainText("hello")
	if txt.Tag != "plain_text" || txt.Content != "hello" {
		t.Errorf("PlainText = %+v", txt)
	}
}

func TestHelperMdText(t *testing.T) {
	txt := MdText("**bold**")
	if txt.Tag != "lark_md" || txt.Content != "**bold**" {
		t.Errorf("MdText = %+v", txt)
	}
}

func TestHelperBtn(t *testing.T) {
	btn := Btn("Click")
	if btn.Text == nil || btn.Text.Content != "Click" || btn.Text.Tag != "plain_text" {
		t.Errorf("Btn = %+v", btn)
	}
	if btn.cardTag() != "button" {
		t.Error("Btn should create Button element")
	}
}

func TestHelperHr(t *testing.T) {
	hr := Hr()
	if hr.cardTag() != "hr" {
		t.Error("Hr should create Divider element")
	}
}

func TestHelperImg(t *testing.T) {
	img := Img("img_key_123")
	if img.ImgKey != "img_key_123" {
		t.Errorf("Img key = %q", img.ImgKey)
	}
	if img.cardTag() != "img" {
		t.Error("Img should create Image element")
	}
}

// ---- Form tests ----

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

	m := mustMarshalCard(t, card)
	elements := getBodyElements(t, m)
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

func TestFormWithAllInteractiveFields(t *testing.T) {
	card := Card{
		Body: Body{
			Elements: Elements{
				&Form{
					Name:              "survey",
					Direction:         "vertical",
					Margin:            "4px 0",
					Padding:           "8px",
					HorizontalSpacing: SpacingMedium,
					VerticalSpacing:   SpacingLarge,
					HorizontalAlign:   "left",
					VerticalAlign:     "top",
					Elements: Elements{
						&Input{
							Name:         "email",
							Required:     true,
							InputType:    InputText,
							MaxLength:    100,
							Placeholder:  PlainText("Email"),
							Label:        PlainText("Email Address"),
							Width:        "fill",
							DefaultValue: "user@example.com",
						},
						&Input{
							Name:      "bio",
							InputType: InputMultilineText,
							Rows:      3,
							MaxRows:   10,
						},
						&SelectStatic{
							Name:          "country",
							Required:      true,
							InitialOption: "us",
							Placeholder:   PlainText("Select country"),
							Options: []SelectOption{
								{Text: PlainText("US"), Value: "us"},
								{Text: PlainText("CN"), Value: "cn"},
								{Text: PlainText("JP"), Value: "jp"},
							},
						},
						&MultiSelectStatic{
							Name:           "interests",
							SelectedValues: []string{"go", "rust"},
							Options: []SelectOption{
								{Text: PlainText("Go"), Value: "go"},
								{Text: PlainText("Rust"), Value: "rust"},
								{Text: PlainText("Python"), Value: "python"},
							},
						},
						&DatePicker{
							Name:        "birthday",
							InitialDate: "2000-01-01",
							Placeholder: PlainText("Birthday"),
						},
						&TimePicker{
							Name:        "alarm",
							InitialTime: "08:00",
						},
						&DateTimePicker{
							Name:            "meeting",
							InitialDatetime: "2026-03-16 14:00",
						},
						&Checker{
							Name:    "agree",
							Checked: false,
							Text:    PlainText("I agree to terms"),
						},
						&Button{
							Text:           PlainText("Submit"),
							Type:           ButtonPrimary,
							FormActionType: FormSubmit,
							Name:           "submit_btn",
						},
						&Button{
							Text:           PlainText("Reset"),
							Type:           ButtonDefault,
							FormActionType: FormReset,
							Name:           "reset_btn",
						},
					},
				},
			},
		},
	}

	m := mustMarshalCard(t, card)
	elements := getBodyElements(t, m)
	form := elements[0].(map[string]any)

	if form["direction"] != "vertical" {
		t.Errorf("form direction = %v", form["direction"])
	}
	if form["horizontal_spacing"] != "medium" {
		t.Errorf("form horizontal_spacing = %v", form["horizontal_spacing"])
	}

	formElems := form["elements"].([]any)
	if len(formElems) != 10 {
		t.Fatalf("form elements = %d, want 10", len(formElems))
	}

	// Verify input with label and max_length.
	emailInput := formElems[0].(map[string]any)
	if emailInput["max_length"] != float64(100) {
		t.Errorf("max_length = %v, want 100", emailInput["max_length"])
	}
	if emailInput["required"] != true {
		t.Error("required should be true")
	}
	if emailInput["default_value"] != "user@example.com" {
		t.Errorf("default_value = %v", emailInput["default_value"])
	}

	// Verify multiline input.
	bioInput := formElems[1].(map[string]any)
	if bioInput["input_type"] != "multiline_text" {
		t.Errorf("input_type = %v", bioInput["input_type"])
	}
	if bioInput["rows"] != float64(3) {
		t.Errorf("rows = %v", bioInput["rows"])
	}

	// Verify select with initial_option.
	sel := formElems[2].(map[string]any)
	if sel["initial_option"] != "us" {
		t.Errorf("initial_option = %v", sel["initial_option"])
	}
	opts := sel["options"].([]any)
	if len(opts) != 3 {
		t.Errorf("options len = %d, want 3", len(opts))
	}

	// Verify multi-select with selected_values.
	msel := formElems[3].(map[string]any)
	sv := msel["selected_values"].([]any)
	if len(sv) != 2 {
		t.Errorf("selected_values len = %d, want 2", len(sv))
	}

	// Verify date picker.
	dp := formElems[4].(map[string]any)
	if dp["initial_date"] != "2000-01-01" {
		t.Errorf("initial_date = %v", dp["initial_date"])
	}

	// Verify time picker.
	tp := formElems[5].(map[string]any)
	if tp["initial_time"] != "08:00" {
		t.Errorf("initial_time = %v", tp["initial_time"])
	}

	// Verify datetime picker.
	dtp := formElems[6].(map[string]any)
	if dtp["initial_datetime"] != "2026-03-16 14:00" {
		t.Errorf("initial_datetime = %v", dtp["initial_datetime"])
	}

	// Verify reset button.
	resetBtn := formElems[9].(map[string]any)
	if resetBtn["form_action_type"] != "reset" {
		t.Errorf("form_action_type = %v, want reset", resetBtn["form_action_type"])
	}
}

// ---- Button tests ----

func TestButtonAllFields(t *testing.T) {
	card := Card{
		Body: Body{
			Elements: Elements{
				&Button{
					ElementID: "btn1",
					Margin:    "4px",
					Type:      ButtonDanger,
					Size:      SizeLarge,
					Width:     "fill",
					Text:      PlainText("Delete"),
					Icon: &Icon{
						Tag:   StandardIcon,
						Token: "delete_outlined",
						Color: Red,
					},
					HoverTips:    PlainText("Click to delete"),
					Disabled:     true,
					DisabledTips: PlainText("Cannot delete"),
					Confirm: &Confirm{
						Title: PlainText("Confirm Delete"),
						Text:  PlainText("Are you sure?"),
					},
					Behaviors: []Behavior{
						{
							Type:       BehaviorCallback,
							Value:      map[string]any{"action": "delete"},
							DefaultURL: "https://example.com/callback",
						},
					},
					URL: "https://example.com",
					MultiURL: &MultiURL{
						URL:        "https://example.com",
						PCURL:      "https://pc.example.com",
						IOSURL:     "https://ios.example.com",
						AndroidURL: "https://android.example.com",
					},
					Value: map[string]any{"id": "123"},
				},
			},
		},
	}

	m := mustMarshalCard(t, card)
	elements := getBodyElements(t, m)
	btn := elements[0].(map[string]any)

	if btn["type"] != "danger" {
		t.Errorf("type = %v, want danger", btn["type"])
	}
	if btn["size"] != "large" {
		t.Errorf("size = %v, want large", btn["size"])
	}
	if btn["width"] != "fill" {
		t.Errorf("width = %v, want fill", btn["width"])
	}
	if btn["disabled"] != true {
		t.Error("disabled should be true")
	}
	if btn["url"] != "https://example.com" {
		t.Errorf("url = %v", btn["url"])
	}

	confirm := btn["confirm"].(map[string]any)
	confirmTitle := confirm["title"].(map[string]any)
	if confirmTitle["content"] != "Confirm Delete" {
		t.Errorf("confirm title = %v", confirmTitle["content"])
	}

	behaviors := btn["behaviors"].([]any)
	if len(behaviors) != 1 {
		t.Fatalf("behaviors len = %d, want 1", len(behaviors))
	}
	beh := behaviors[0].(map[string]any)
	if beh["type"] != "callback" {
		t.Errorf("behavior type = %v, want callback", beh["type"])
	}

	multiURL := btn["multi_url"].(map[string]any)
	if multiURL["pc_url"] != "https://pc.example.com" {
		t.Errorf("multi_url.pc_url = %v", multiURL["pc_url"])
	}

	val := btn["value"].(map[string]any)
	if val["id"] != "123" {
		t.Errorf("value.id = %v", val["id"])
	}
}

func TestButtonTypes(t *testing.T) {
	types := []string{
		ButtonDefault, ButtonPrimary, ButtonDanger, ButtonText,
		ButtonPrimaryText, ButtonDangerText, ButtonPrimaryFilled,
		ButtonDangerFilled, ButtonLaser,
	}
	for _, bt := range types {
		t.Run(bt, func(t *testing.T) {
			elems := Elements{&Button{Type: bt, Text: PlainText("x")}}
			arr := mustMarshalElements(t, elems)
			if arr[0].(map[string]any)["type"] != bt {
				t.Errorf("type = %v, want %v", arr[0].(map[string]any)["type"], bt)
			}
		})
	}
}

// ---- Input tests ----

func TestInputPasswordType(t *testing.T) {
	elems := Elements{
		&Input{
			Name:      "pw",
			InputType: InputPassword,
			MaxLength: 50,
			Width:     "200px",
		},
	}
	arr := mustMarshalElements(t, elems)
	input := arr[0].(map[string]any)
	if input["input_type"] != "password" {
		t.Errorf("input_type = %v", input["input_type"])
	}
	if input["max_length"] != float64(50) {
		t.Errorf("max_length = %v", input["max_length"])
	}
}

func TestInputWithLabelAndBehaviors(t *testing.T) {
	showIcon := true
	elems := Elements{
		&Input{
			Name:          "search",
			Label:         PlainText("Search"),
			LabelPosition: "top",
			ShowIcon:      &showIcon,
			Behaviors: []Behavior{
				{Type: BehaviorCallback},
			},
			Confirm: &Confirm{
				Title: PlainText("Confirm"),
				Text:  PlainText("Submit search?"),
			},
		},
	}
	arr := mustMarshalElements(t, elems)
	input := arr[0].(map[string]any)
	if input["label_position"] != "top" {
		t.Errorf("label_position = %v", input["label_position"])
	}
	if input["show_icon"] != true {
		t.Error("show_icon should be true")
	}
}

// ---- Table tests ----

func TestTableWithColumnsAndRows(t *testing.T) {
	bold := true
	card := Card{
		Body: Body{
			Elements: Elements{
				&Table{
					ElementID:         "tbl1",
					PageSize:          5,
					RowHeight:         "middle",
					FreezeFirstColumn: true,
					HeaderStyle: &HeaderStyle{
						TextAlign:       "center",
						TextSize:        TextSizeNormal,
						BackgroundStyle: "grey",
						TextColor:       "default",
						Bold:            &bold,
						Lines:           1,
					},
					Columns: []TableColumn{
						{
							Name:        "name",
							DisplayName: "Name",
							Width:       "200px",
							DataType:    DataTypeText,
						},
						{
							Name:            "amount",
							DisplayName:     "Amount",
							Width:           "auto",
							DataType:        DataTypeNumber,
							HorizontalAlign: "right",
							Format: &NumberFormat{
								Precision: 2,
								Symbol:    "$",
								Separator: true,
							},
						},
						{
							Name:        "date",
							DisplayName: "Date",
							DataType:    DataTypeDate,
							DateFormat:  "YYYY-MM-DD",
						},
						{
							Name:        "desc",
							DisplayName: "Description",
							DataType:    DataTypeLarkMd,
						},
					},
					Rows: []map[string]any{
						{"name": "Item A", "amount": 1234.56, "date": 1710547200000, "desc": "**bold** item"},
						{"name": "Item B", "amount": 78.90, "date": 1710633600000, "desc": "normal item"},
					},
				},
			},
		},
	}

	m := mustMarshalCard(t, card)
	elements := getBodyElements(t, m)
	tbl := elements[0].(map[string]any)

	if tbl["tag"] != "table" {
		t.Errorf("tag = %v, want table", tbl["tag"])
	}
	if tbl["page_size"] != float64(5) {
		t.Errorf("page_size = %v, want 5", tbl["page_size"])
	}
	if tbl["freeze_first_column"] != true {
		t.Error("freeze_first_column should be true")
	}

	hs := tbl["header_style"].(map[string]any)
	if hs["bold"] != true {
		t.Error("header bold should be true")
	}
	if hs["text_align"] != "center" {
		t.Errorf("text_align = %v", hs["text_align"])
	}

	cols := tbl["columns"].([]any)
	if len(cols) != 4 {
		t.Fatalf("columns len = %d, want 4", len(cols))
	}

	// Check number format.
	amtCol := cols[1].(map[string]any)
	if amtCol["data_type"] != "number" {
		t.Errorf("data_type = %v, want number", amtCol["data_type"])
	}
	format := amtCol["format"].(map[string]any)
	if format["precision"] != float64(2) {
		t.Errorf("precision = %v", format["precision"])
	}
	if format["symbol"] != "$" {
		t.Errorf("symbol = %v", format["symbol"])
	}

	rows := tbl["rows"].([]any)
	if len(rows) != 2 {
		t.Fatalf("rows len = %d, want 2", len(rows))
	}
}

func TestTableEmptyRows(t *testing.T) {
	elems := Elements{
		&Table{
			Columns: []TableColumn{
				{Name: "col1", DataType: DataTypeText},
			},
			Rows: []map[string]any{},
		},
	}
	arr := mustMarshalElements(t, elems)
	tbl := arr[0].(map[string]any)
	// Empty slices are omitted by Go's omitempty; rows key should be absent.
	if _, ok := tbl["rows"]; ok {
		t.Error("empty rows should be omitted by omitempty")
	}
}

func TestTableWithNilRows(t *testing.T) {
	elems := Elements{
		&Table{
			Columns: []TableColumn{
				{Name: "col1", DataType: DataTypeText},
			},
		},
	}
	arr := mustMarshalElements(t, elems)
	tbl := arr[0].(map[string]any)
	if _, ok := tbl["rows"]; ok {
		t.Error("nil rows should be omitted by omitempty")
	}
}

// ---- ColumnSet / Column tests ----

func TestColumnSetAllFields(t *testing.T) {
	card := Card{
		Body: Body{
			Elements: Elements{
				&ColumnSet{
					ElementID:         "cs1",
					Margin:            "4px",
					HorizontalSpacing: SpacingSmall,
					HorizontalAlign:   "center",
					FlexMode:          FlexBisect,
					BackgroundStyle:   "grey",
					Action: &Action{
						MultiURL: &MultiURL{
							URL: "https://example.com",
						},
					},
					Columns: []Column{
						{
							ElementID:       "col1",
							Width:           "weighted",
							Weight:          3,
							VerticalAlign:   "center",
							Padding:         "4px 8px",
							BackgroundStyle: "default",
							Elements:        Elements{Md("Col 1")},
						},
						{
							Width:    "weighted",
							Weight:   2,
							Elements: Elements{Md("Col 2")},
						},
						{
							Width:    "150px",
							Elements: Elements{Btn("Go")},
						},
					},
				},
			},
		},
	}

	m := mustMarshalCard(t, card)
	elements := getBodyElements(t, m)
	cs := elements[0].(map[string]any)

	if cs["flex_mode"] != "bisect" {
		t.Errorf("flex_mode = %v, want bisect", cs["flex_mode"])
	}
	if cs["horizontal_spacing"] != "small" {
		t.Errorf("horizontal_spacing = %v", cs["horizontal_spacing"])
	}

	action := cs["action"].(map[string]any)
	mu := action["multi_url"].(map[string]any)
	if mu["url"] != "https://example.com" {
		t.Errorf("action url = %v", mu["url"])
	}

	cols := cs["columns"].([]any)
	if len(cols) != 3 {
		t.Fatalf("columns len = %d, want 3", len(cols))
	}

	col0 := cols[0].(map[string]any)
	if col0["weight"] != float64(3) {
		t.Errorf("weight = %v, want 3", col0["weight"])
	}

	col2 := cols[2].(map[string]any)
	if col2["width"] != "150px" {
		t.Errorf("width = %v, want 150px", col2["width"])
	}
}

// Lark docs: max 5 nesting layers for columns.
func TestDeepNestedColumns(t *testing.T) {
	// Build 5 levels of nested ColumnSets.
	innermost := Elements{Md("Level 5")}

	level4 := Elements{
		&ColumnSet{
			Columns: []Column{
				{Width: "auto", Elements: innermost},
			},
		},
	}
	level3 := Elements{
		&ColumnSet{
			Columns: []Column{
				{Width: "auto", Elements: level4},
			},
		},
	}
	level2 := Elements{
		&ColumnSet{
			Columns: []Column{
				{Width: "auto", Elements: level3},
			},
		},
	}
	level1 := Elements{
		&ColumnSet{
			Columns: []Column{
				{Width: "auto", Elements: level2},
			},
		},
	}

	card := Card{
		Body: Body{
			Elements: level1,
		},
	}

	// Should marshal without error even at max nesting depth.
	data, err := card.JSON()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "Level 5") {
		t.Error("innermost content not found in output")
	}
}

// ---- InteractiveContainer tests ----

func TestInteractiveContainerFull(t *testing.T) {
	card := Card{
		Body: Body{
			Elements: Elements{
				&InteractiveContainer{
					ElementID:       "ic1",
					Width:           "fill",
					Height:          "200px",
					Margin:          "8px",
					Direction:       "vertical",
					VerticalSpacing: SpacingMedium,
					BackgroundStyle: "default",
					HasBorder:       true,
					BorderColor:     "blue",
					CornerRadius:    "8px",
					Padding:         "12px",
					Disabled:        false,
					HoverTips:       PlainText("Click me"),
					Behaviors: []Behavior{
						{
							Type:       BehaviorOpenURL,
							DefaultURL: "https://example.com",
							PCURL:      "https://pc.example.com",
						},
					},
					Confirm: &Confirm{
						Title: PlainText("Navigate?"),
						Text:  PlainText("Go to example.com?"),
					},
					Elements: Elements{
						Md("Interactive content"),
						Btn("Action"),
					},
				},
			},
		},
	}

	m := mustMarshalCard(t, card)
	elements := getBodyElements(t, m)
	ic := elements[0].(map[string]any)

	if ic["tag"] != "interactive_container" {
		t.Errorf("tag = %v", ic["tag"])
	}
	if ic["has_border"] != true {
		t.Error("has_border should be true")
	}
	if ic["corner_radius"] != "8px" {
		t.Errorf("corner_radius = %v", ic["corner_radius"])
	}

	behaviors := ic["behaviors"].([]any)
	if len(behaviors) != 1 {
		t.Fatalf("behaviors len = %d", len(behaviors))
	}
	beh := behaviors[0].(map[string]any)
	if beh["type"] != "open_url" {
		t.Errorf("behavior type = %v", beh["type"])
	}

	icElements := ic["elements"].([]any)
	if len(icElements) != 2 {
		t.Fatalf("ic elements = %d, want 2", len(icElements))
	}
}

// ---- CollapsiblePanel tests ----

func TestCollapsiblePanelFull(t *testing.T) {
	card := Card{
		Body: Body{
			Elements: Elements{
				&CollapsiblePanel{
					ElementID:       "cp1",
					Expanded:        true,
					Direction:       "vertical",
					Margin:          "4px",
					Padding:         "8px",
					VerticalSpacing: SpacingSmall,
					BackgroundColor: "default",
					Header: &PanelHeader{
						Title:           PlainText("Details"),
						BackgroundColor: "grey",
						VerticalAlign:   "center",
						Padding:         "4px 8px",
						Width:           "fill",
						Icon: &Icon{
							Tag:   StandardIcon,
							Token: "chevron_down_outlined",
						},
						IconPosition:      "right",
						IconExpandedAngle: 180,
					},
					Border: &BorderConfig{
						Color:        "grey",
						CornerRadius: "4px",
					},
					Elements: Elements{
						Md("Hidden content revealed"),
						&Image{ImgKey: "img_detail"},
					},
				},
			},
		},
	}

	m := mustMarshalCard(t, card)
	elements := getBodyElements(t, m)
	cp := elements[0].(map[string]any)

	if cp["tag"] != "collapsible_panel" {
		t.Errorf("tag = %v", cp["tag"])
	}
	if cp["expanded"] != true {
		t.Error("expanded should be true")
	}

	header := cp["header"].(map[string]any)
	if header["icon_expanded_angle"] != float64(180) {
		t.Errorf("icon_expanded_angle = %v", header["icon_expanded_angle"])
	}
	if header["icon_position"] != "right" {
		t.Errorf("icon_position = %v", header["icon_position"])
	}

	border := cp["border"].(map[string]any)
	if border["corner_radius"] != "4px" {
		t.Errorf("corner_radius = %v", border["corner_radius"])
	}

	cpElements := cp["elements"].([]any)
	if len(cpElements) != 2 {
		t.Fatalf("elements len = %d, want 2", len(cpElements))
	}
}

// ---- Checker tests ----

func TestCheckerFull(t *testing.T) {
	overallCheckable := true
	card := Card{
		Body: Body{
			Elements: Elements{
				&Checker{
					ElementID:        "chk1",
					Name:             "task1",
					Checked:          true,
					Text:             PlainText("Complete task"),
					OverallCheckable: &overallCheckable,
					ButtonArea: &CheckerButtonArea{
						PCDisplayRule: "always",
						Buttons: []Button{
							{
								Text: PlainText("Edit"),
								Type: ButtonText,
								Size: SizeTiny,
							},
							{
								Text: PlainText("Delete"),
								Type: ButtonDangerText,
								Size: SizeTiny,
							},
						},
					},
					CheckedStyle: &CheckedStyle{
						ShowStrikethrough: true,
						Opacity:           0.5,
					},
					Margin:       "4px 0",
					Padding:      "8px",
					HoverTips:    PlainText("Click to toggle"),
					Disabled:     false,
					DisabledTips: PlainText("Disabled"),
					Confirm: &Confirm{
						Title: PlainText("Toggle?"),
						Text:  PlainText("Change status?"),
					},
					Behaviors: []Behavior{
						{Type: BehaviorCallback},
					},
				},
			},
		},
	}

	m := mustMarshalCard(t, card)
	elements := getBodyElements(t, m)
	chk := elements[0].(map[string]any)

	if chk["tag"] != "checker" {
		t.Errorf("tag = %v", chk["tag"])
	}
	if chk["checked"] != true {
		t.Error("checked should be true")
	}
	if chk["overall_checkable"] != true {
		t.Error("overall_checkable should be true")
	}

	ba := chk["button_area"].(map[string]any)
	if ba["pc_display_rule"] != "always" {
		t.Errorf("pc_display_rule = %v", ba["pc_display_rule"])
	}
	buttons := ba["buttons"].([]any)
	// Lark docs: max 3 buttons in checker button area.
	if len(buttons) != 2 {
		t.Fatalf("buttons len = %d, want 2", len(buttons))
	}

	cs := chk["checked_style"].(map[string]any)
	if cs["show_strikethrough"] != true {
		t.Error("show_strikethrough should be true")
	}
	if cs["opacity"] != 0.5 {
		t.Errorf("opacity = %v, want 0.5", cs["opacity"])
	}
}

// ---- Image tests ----

func TestImageAllFields(t *testing.T) {
	preview := false
	elems := Elements{
		&Image{
			ElementID:    "img1",
			Margin:       "4px",
			ImgKey:       "img_abc123",
			Alt:          PlainText("An image"),
			Title:        PlainText("Title"),
			CornerRadius: "8px",
			ScaleType:    ScaleCropCenter,
			Size:         "large",
			Transparent:  true,
			Preview:      &preview,
			Mode:         "crop_center",
			CustomWidth:  "300px",
			CompactWidth: true,
		},
	}
	arr := mustMarshalElements(t, elems)
	img := arr[0].(map[string]any)

	if img["tag"] != "img" {
		t.Errorf("tag = %v", img["tag"])
	}
	if img["img_key"] != "img_abc123" {
		t.Errorf("img_key = %v", img["img_key"])
	}
	if img["corner_radius"] != "8px" {
		t.Errorf("corner_radius = %v", img["corner_radius"])
	}
	if img["scale_type"] != "crop_center" {
		t.Errorf("scale_type = %v", img["scale_type"])
	}
	if img["transparent"] != true {
		t.Error("transparent should be true")
	}
	if img["preview"] != false {
		t.Error("preview should be false")
	}
	if img["compact_width"] != true {
		t.Error("compact_width should be true")
	}
}

func TestMultiImageLayout(t *testing.T) {
	elems := Elements{
		&MultiImageLayout{
			ElementID:              "mil1",
			CombinationMode:        CombinationTrisect,
			CombinationTransparent: true,
			CornerRadius:           "4px",
			ImgList: []ImageItem{
				{ImgKey: "img_1", Transparent: false},
				{ImgKey: "img_2", Transparent: true},
				{ImgKey: "img_3"},
			},
		},
	}
	arr := mustMarshalElements(t, elems)
	mil := arr[0].(map[string]any)

	if mil["tag"] != "img_combination" {
		t.Errorf("tag = %v", mil["tag"])
	}
	if mil["combination_mode"] != "trisect" {
		t.Errorf("combination_mode = %v", mil["combination_mode"])
	}
	imgList := mil["img_list"].([]any)
	if len(imgList) != 3 {
		t.Fatalf("img_list len = %d, want 3", len(imgList))
	}
}

// ---- Person / PersonList tests ----

func TestPerson(t *testing.T) {
	showAvatar := true
	showName := false
	elems := Elements{
		&Person{
			ElementID:  "p1",
			Size:       SizeMedium,
			UserID:     "ou_xxxxxxxxxxxx",
			ShowAvatar: &showAvatar,
			ShowName:   &showName,
			Style:      PersonStyleCapsule,
		},
	}
	arr := mustMarshalElements(t, elems)
	p := arr[0].(map[string]any)

	if p["tag"] != "person" {
		t.Errorf("tag = %v", p["tag"])
	}
	if p["user_id"] != "ou_xxxxxxxxxxxx" {
		t.Errorf("user_id = %v", p["user_id"])
	}
	if p["show_avatar"] != true {
		t.Error("show_avatar should be true")
	}
	if p["show_name"] != false {
		t.Error("show_name should be false")
	}
	if p["style"] != "capsule" {
		t.Errorf("style = %v", p["style"])
	}
}

func TestPersonList(t *testing.T) {
	showName := true
	showAvatar := true
	elems := Elements{
		&PersonList{
			ElementID:         "pl1",
			DropInvalidUserID: true,
			Lines:             2,
			ShowName:          &showName,
			ShowAvatar:        &showAvatar,
			Size:              SizeSmall,
			Persons: []PersonItem{
				{ID: "ou_user1"},
				{ID: "ou_user2"},
				{ID: "ou_user3"},
			},
			Icon: &Icon{
				Tag:   StandardIcon,
				Token: "people_outlined",
			},
			UdIcon: &UdIcon{
				Token: "custom_people",
				Style: map[string]string{
					"color": "blue",
				},
			},
		},
	}
	arr := mustMarshalElements(t, elems)
	pl := arr[0].(map[string]any)

	if pl["tag"] != "person_list" {
		t.Errorf("tag = %v", pl["tag"])
	}
	if pl["drop_invalid_user_id"] != true {
		t.Error("drop_invalid_user_id should be true")
	}
	persons := pl["persons"].([]any)
	if len(persons) != 3 {
		t.Fatalf("persons len = %d, want 3", len(persons))
	}

	udIcon := pl["ud_icon"].(map[string]any)
	if udIcon["token"] != "custom_people" {
		t.Errorf("ud_icon token = %v", udIcon["token"])
	}
}

// ---- Select tests ----

func TestSelectStaticWithOptions(t *testing.T) {
	elems := Elements{
		&SelectStatic{
			ElementID:     "sel1",
			Name:          "priority",
			Required:      true,
			InitialOption: "medium",
			InitialIndex:  1,
			Placeholder:   PlainText("Select priority"),
			Width:         "200px",
			Options: []SelectOption{
				{
					Text:  PlainText("High"),
					Value: "high",
					Icon:  &Icon{Tag: StandardIcon, Token: "fire_outlined"},
				},
				{Text: PlainText("Medium"), Value: "medium"},
				{Text: PlainText("Low"), Value: "low"},
			},
			Confirm: &Confirm{
				Title: PlainText("Change?"),
				Text:  PlainText("Change priority?"),
			},
		},
	}
	arr := mustMarshalElements(t, elems)
	sel := arr[0].(map[string]any)

	if sel["tag"] != "select_static" {
		t.Errorf("tag = %v", sel["tag"])
	}
	if sel["initial_index"] != float64(1) {
		t.Errorf("initial_index = %v", sel["initial_index"])
	}
	opts := sel["options"].([]any)
	if len(opts) != 3 {
		t.Fatalf("options len = %d", len(opts))
	}
	// First option should have icon.
	opt0 := opts[0].(map[string]any)
	if opt0["icon"] == nil {
		t.Error("first option should have icon")
	}
}

func TestSelectPersonAndMulti(t *testing.T) {
	elems := Elements{
		&SelectPerson{
			Name:          "reviewer",
			InitialOption: "ou_reviewer1",
			Options: []PersonOption{
				{Value: "ou_reviewer1"},
				{Value: "ou_reviewer2"},
			},
		},
		&MultiSelectPerson{
			Name:           "watchers",
			SelectedValues: []string{"ou_user1", "ou_user2"},
			Options: []PersonOption{
				{Value: "ou_user1"},
				{Value: "ou_user2"},
				{Value: "ou_user3"},
			},
		},
	}
	arr := mustMarshalElements(t, elems)
	if len(arr) != 2 {
		t.Fatalf("len = %d, want 2", len(arr))
	}

	sp := arr[0].(map[string]any)
	if sp["tag"] != "select_person" {
		t.Errorf("tag = %v", sp["tag"])
	}
	if sp["initial_option"] != "ou_reviewer1" {
		t.Errorf("initial_option = %v", sp["initial_option"])
	}

	msp := arr[1].(map[string]any)
	if msp["tag"] != "multi_select_person" {
		t.Errorf("tag = %v", msp["tag"])
	}
	sv := msp["selected_values"].([]any)
	if len(sv) != 2 {
		t.Errorf("selected_values len = %d", len(sv))
	}
}

// ---- Overflow tests ----

func TestOverflow(t *testing.T) {
	elems := Elements{
		&Overflow{
			ElementID: "of1",
			Width:     "fill",
			Options: []OverflowOption{
				{
					Text:  PlainText("Option A"),
					Value: "a",
					MultiURL: &MultiURL{
						URL: "https://a.example.com",
					},
				},
				{
					Text:  PlainText("Option B"),
					Value: "b",
				},
			},
			Value: map[string]any{"key": "val"},
			Confirm: &Confirm{
				Title: PlainText("Select?"),
			},
		},
	}
	arr := mustMarshalElements(t, elems)
	of := arr[0].(map[string]any)

	if of["tag"] != "overflow" {
		t.Errorf("tag = %v", of["tag"])
	}
	opts := of["options"].([]any)
	if len(opts) != 2 {
		t.Fatalf("options len = %d", len(opts))
	}
	opt0 := opts[0].(map[string]any)
	mu := opt0["multi_url"].(map[string]any)
	if mu["url"] != "https://a.example.com" {
		t.Errorf("url = %v", mu["url"])
	}
}

// ---- ImagePicker tests ----

func TestImagePicker(t *testing.T) {
	canPreview := true
	elems := Elements{
		&ImagePicker{
			ElementID:   "ip1",
			MultiSelect: true,
			Layout:      LayoutTrisect,
			Name:        "photos",
			Required:    true,
			CanPreview:  &canPreview,
			AspectRatio: "1:1",
			Value:       []string{"val1", "val2"},
			Options: []ImagePickerOption{
				{
					ImgKey:    "img_opt1",
					Value:     "val1",
					HoverTips: PlainText("Photo 1"),
				},
				{
					ImgKey:       "img_opt2",
					Value:        "val2",
					Disabled:     true,
					DisabledTips: PlainText("Unavailable"),
				},
				{ImgKey: "img_opt3", Value: "val3"},
			},
			Behaviors: []Behavior{
				{Type: BehaviorCallback},
			},
		},
	}
	arr := mustMarshalElements(t, elems)
	ip := arr[0].(map[string]any)

	if ip["tag"] != "select_img" {
		t.Errorf("tag = %v", ip["tag"])
	}
	if ip["multi_select"] != true {
		t.Error("multi_select should be true")
	}
	if ip["layout"] != "trisect" {
		t.Errorf("layout = %v", ip["layout"])
	}
	if ip["aspect_ratio"] != "1:1" {
		t.Errorf("aspect_ratio = %v", ip["aspect_ratio"])
	}
	opts := ip["options"].([]any)
	if len(opts) != 3 {
		t.Fatalf("options len = %d", len(opts))
	}
}

// ---- Chart test ----

func TestChart(t *testing.T) {
	spec := map[string]any{
		"type": "bar",
		"data": []map[string]any{
			{"x": "A", "y": 10},
			{"x": "B", "y": 20},
		},
	}
	elems := Elements{
		&Chart{
			ElementID:   "chart1",
			AspectRatio: "16:9",
			ColorTheme:  "brand",
			Height:      "200px",
			ChartSpec:   spec,
		},
	}
	arr := mustMarshalElements(t, elems)
	ch := arr[0].(map[string]any)

	if ch["tag"] != "chart" {
		t.Errorf("tag = %v", ch["tag"])
	}
	if ch["aspect_ratio"] != "16:9" {
		t.Errorf("aspect_ratio = %v", ch["aspect_ratio"])
	}
	if ch["height"] != "200px" {
		t.Errorf("height = %v", ch["height"])
	}
	if ch["chart_spec"] == nil {
		t.Error("chart_spec should be present")
	}
}

// ---- Div tests ----

func TestDivWithTextAndIcon(t *testing.T) {
	elems := Elements{
		&Div{
			ElementID: "div1",
			Margin:    "4px 8px",
			Width:     "fill",
			Text: &Text{
				Tag:       "plain_text",
				Content:   "Status: OK",
				TextSize:  TextSizeNormal,
				TextColor: "green",
				TextAlign: "left",
				Lines:     2,
			},
			Icon: &Icon{
				Tag:   StandardIcon,
				Token: "check_outlined",
				Color: Green,
			},
		},
	}
	arr := mustMarshalElements(t, elems)
	div := arr[0].(map[string]any)

	if div["tag"] != "div" {
		t.Errorf("tag = %v", div["tag"])
	}

	text := div["text"].(map[string]any)
	if text["lines"] != float64(2) {
		t.Errorf("lines = %v, want 2", text["lines"])
	}
	if text["text_size"] != "normal" {
		t.Errorf("text_size = %v", text["text_size"])
	}

	icon := div["icon"].(map[string]any)
	if icon["token"] != "check_outlined" {
		t.Errorf("icon token = %v", icon["token"])
	}
}

// ---- Markdown tests ----

func TestMarkdownAllFields(t *testing.T) {
	elems := Elements{
		&Markdown{
			ElementID: "md1",
			Margin:    "8px 0",
			Content:   "# Heading\n**Bold** and *italic*\n- List item",
			TextSize:  TextSizeHeading1,
			TextColor: "default",
			TextAlign: "center",
			Icon: &Icon{
				Tag:   StandardIcon,
				Token: "info_outlined",
			},
		},
	}
	arr := mustMarshalElements(t, elems)
	md := arr[0].(map[string]any)

	if md["text_size"] != "heading-1" {
		t.Errorf("text_size = %v", md["text_size"])
	}
	if md["text_align"] != "center" {
		t.Errorf("text_align = %v", md["text_align"])
	}
	if md["icon"] == nil {
		t.Error("icon should be present")
	}
}

// ---- Large element count (approaching 200 limit) ----

func TestManyElements(t *testing.T) {
	var elems Elements
	for i := 0; i < 200; i++ {
		elems = append(elems, Md("item"))
	}

	card := Card{
		Body: Body{Elements: elems},
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
	if len(elements) != 200 {
		t.Errorf("expected 200 elements, got %d", len(elements))
	}
}

// ---- Mixed element types in single body ----

func TestMixedElementTypes(t *testing.T) {
	card := Card{
		Body: Body{
			Elements: Elements{
				Md("Markdown"),
				Hr(),
				&Div{Text: PlainText("Div")},
				&Image{ImgKey: "img1"},
				&Button{Text: PlainText("Btn"), Type: ButtonPrimary},
				&ColumnSet{
					Columns: []Column{
						{Width: "auto", Elements: Elements{Md("Col")}},
					},
				},
				&Person{UserID: "ou_1"},
			},
		},
	}

	m := mustMarshalCard(t, card)
	elements := getBodyElements(t, m)
	if len(elements) != 7 {
		t.Fatalf("expected 7 elements, got %d", len(elements))
	}

	expectedTags := []string{"markdown", "hr", "div", "img", "button", "column_set", "person"}
	for i, tag := range expectedTags {
		elem := elements[i].(map[string]any)
		if elem["tag"] != tag {
			t.Errorf("element[%d].tag = %v, want %v", i, elem["tag"], tag)
		}
	}
}

// ---- Body layout fields ----

func TestBodyLayoutFields(t *testing.T) {
	card := Card{
		Body: Body{
			Direction:         "vertical",
			Padding:           "16px",
			HorizontalSpacing: SpacingLarge,
			VerticalSpacing:   SpacingMedium,
			HorizontalAlign:   "center",
			VerticalAlign:     "top",
			Elements:          Elements{Md("content")},
		},
	}

	m := mustMarshalCard(t, card)
	body := m["body"].(map[string]any)
	if body["direction"] != "vertical" {
		t.Errorf("direction = %v", body["direction"])
	}
	if body["padding"] != "16px" {
		t.Errorf("padding = %v", body["padding"])
	}
	if body["horizontal_spacing"] != "large" {
		t.Errorf("horizontal_spacing = %v", body["horizontal_spacing"])
	}
	if body["horizontal_align"] != "center" {
		t.Errorf("horizontal_align = %v", body["horizontal_align"])
	}
}

// ---- Special characters in content ----

func TestSpecialCharactersInContent(t *testing.T) {
	card := Card{
		Header: Header{
			Title: PlainText(`Title with "quotes" & <html> chars`),
		},
		Body: Body{
			Elements: Elements{
				Md("Content with `code` and\nnewlines\tand tabs"),
				Md("Unicode: 日本語 中文 한국어 العربية emoji 🎉🚀"),
				Md("Math: 2 > 1 && 1 < 2"),
			},
		},
	}

	data, err := card.JSON()
	if err != nil {
		t.Fatal(err)
	}

	// Verify it round-trips correctly.
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}

	header := m["header"].(map[string]any)
	title := header["title"].(map[string]any)
	if title["content"] != `Title with "quotes" & <html> chars` {
		t.Errorf("title content = %v", title["content"])
	}

	elements := getBodyElements(t, m)
	md1 := elements[1].(map[string]any)
	if !strings.Contains(md1["content"].(string), "日本語") {
		t.Error("unicode content not preserved")
	}
}

// ---- Column weight boundaries (1-5 per docs) ----

func TestColumnWeightBoundaries(t *testing.T) {
	tests := []struct {
		name   string
		weight int
	}{
		{"min_weight", 1},
		{"max_weight", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elems := Elements{
				&ColumnSet{
					Columns: []Column{
						{Width: "weighted", Weight: tt.weight, Elements: Elements{Md("x")}},
					},
				},
			}
			arr := mustMarshalElements(t, elems)
			cs := arr[0].(map[string]any)
			cols := cs["columns"].([]any)
			col := cols[0].(map[string]any)
			if col["weight"] != float64(tt.weight) {
				t.Errorf("weight = %v, want %v", col["weight"], tt.weight)
			}
		})
	}
}

// ---- Enum constant values ----

func TestEnumConstants(t *testing.T) {
	// Header colors.
	colors := map[string]string{
		"Blue": Blue, "Wathet": Wathet, "Turquoise": Turquoise,
		"Green": Green, "Yellow": Yellow, "Orange": Orange,
		"Red": Red, "Carmine": Carmine, "Violet": Violet,
		"Purple": Purple, "Indigo": Indigo, "Grey": Grey,
		"Default": Default,
	}
	for name, val := range colors {
		if val == "" {
			t.Errorf("color %s is empty", name)
		}
	}

	// Flex modes.
	flexModes := []string{FlexNone, FlexStretch, FlexFlow, FlexBisect, FlexTrisect}
	expected := []string{"none", "stretch", "flow", "bisect", "trisect"}
	for i, fm := range flexModes {
		if fm != expected[i] {
			t.Errorf("FlexMode %d = %q, want %q", i, fm, expected[i])
		}
	}

	// Image scale types.
	if ScaleCropCenter != "crop_center" {
		t.Errorf("ScaleCropCenter = %v", ScaleCropCenter)
	}
	if ScaleCropTop != "crop_top" {
		t.Errorf("ScaleCropTop = %v", ScaleCropTop)
	}
	if ScaleFitHorizontal != "fit_horizontal" {
		t.Errorf("ScaleFitHorizontal = %v", ScaleFitHorizontal)
	}

	// Input types.
	if InputText != "text" {
		t.Errorf("InputText = %v", InputText)
	}
	if InputMultilineText != "multiline_text" {
		t.Errorf("InputMultilineText = %v", InputMultilineText)
	}
	if InputPassword != "password" {
		t.Errorf("InputPassword = %v", InputPassword)
	}

	// Table data types.
	dtypes := []string{DataTypeText, DataTypeLarkMd, DataTypeOptions, DataTypeNumber, DataTypePersons, DataTypeDate, DataTypeMarkdown}
	expectedDt := []string{"text", "lark_md", "options", "number", "persons", "date", "markdown"}
	for i, dt := range dtypes {
		if dt != expectedDt[i] {
			t.Errorf("DataType %d = %q, want %q", i, dt, expectedDt[i])
		}
	}
}

// ---- Realistic end-to-end card scenarios ----

func TestDeployNotificationCard(t *testing.T) {
	card := Card{
		Config: &Config{
			WidthMode: WidthFill,
		},
		Header: Header{
			Title:    PlainText("Deploy Notification"),
			Subtitle: PlainText("Production"),
			Template: Green,
			TextTagList: []TextTag{
				{Tag: "plain_text", Text: PlainText("Success"), Color: Green},
			},
		},
		Body: Body{
			Elements: Elements{
				Md("Service **user-api** deployed to production"),
				Hr(),
				&ColumnSet{
					HorizontalSpacing: SpacingMedium,
					Columns: []Column{
						{
							Width:  "weighted",
							Weight: 1,
							Elements: Elements{
								&Div{Text: &Text{Tag: "plain_text", Content: "Version", TextColor: "grey"}},
								Md("**v2.3.1**"),
							},
						},
						{
							Width:  "weighted",
							Weight: 1,
							Elements: Elements{
								&Div{Text: &Text{Tag: "plain_text", Content: "Duration", TextColor: "grey"}},
								Md("**45s**"),
							},
						},
						{
							Width:  "weighted",
							Weight: 1,
							Elements: Elements{
								&Div{Text: &Text{Tag: "plain_text", Content: "Operator", TextColor: "grey"}},
								&Person{UserID: "ou_deployer", Size: SizeSmall},
							},
						},
					},
				},
				Hr(),
				&Button{
					Text: PlainText("View Logs"),
					Type: ButtonPrimary,
					Size: SizeMedium,
					URL:  "https://logs.example.com",
				},
			},
		},
	}

	m := mustMarshalCard(t, card)
	if m["schema"] != "2.0" {
		t.Error("missing schema")
	}

	elements := getBodyElements(t, m)
	// md, hr, column_set, hr, button
	if len(elements) != 5 {
		t.Fatalf("expected 5 elements, got %d", len(elements))
	}
}

func TestTaskBoardCard(t *testing.T) {
	overallCheckable := true
	card := Card{
		Header: Header{
			Title:    PlainText("Task Board"),
			Template: Violet,
		},
		Body: Body{
			Elements: Elements{
				&Checker{
					Name:             "task_1",
					Checked:          true,
					Text:             PlainText("Set up CI/CD pipeline"),
					OverallCheckable: &overallCheckable,
					CheckedStyle: &CheckedStyle{
						ShowStrikethrough: true,
						Opacity:           0.6,
					},
				},
				&Checker{
					Name:             "task_2",
					Checked:          false,
					Text:             PlainText("Write unit tests"),
					OverallCheckable: &overallCheckable,
				},
				&Checker{
					Name:             "task_3",
					Checked:          false,
					Text:             PlainText("Deploy to staging"),
					OverallCheckable: &overallCheckable,
					ButtonArea: &CheckerButtonArea{
						Buttons: []Button{
							{Text: PlainText("Assign"), Type: ButtonPrimaryText, Size: SizeTiny},
						},
					},
				},
				Hr(),
				&CollapsiblePanel{
					Expanded: false,
					Header: &PanelHeader{
						Title: PlainText("Completed Tasks (5)"),
					},
					Elements: Elements{
						Md("- ~~Design review~~\n- ~~API spec~~\n- ~~Database schema~~\n- ~~Auth integration~~\n- ~~Load testing~~"),
					},
				},
			},
		},
	}

	m := mustMarshalCard(t, card)
	elements := getBodyElements(t, m)
	if len(elements) != 5 {
		t.Fatalf("expected 5 elements, got %d", len(elements))
	}

	// Verify first checker is checked.
	chk1 := elements[0].(map[string]any)
	if chk1["tag"] != "checker" {
		t.Errorf("tag = %v", chk1["tag"])
	}
	if chk1["checked"] != true {
		t.Error("first checker should be checked")
	}

	// Verify collapsible panel.
	cp := elements[4].(map[string]any)
	if cp["tag"] != "collapsible_panel" {
		t.Errorf("tag = %v", cp["tag"])
	}
}

func TestDataDashboardCard(t *testing.T) {
	bold := true
	card := Card{
		Config: &Config{
			WidthMode: WidthFill,
		},
		Header: Header{
			Title:    PlainText("Sales Dashboard"),
			Template: Indigo,
		},
		Body: Body{
			Elements: Elements{
				&Chart{
					AspectRatio: "2:1",
					Height:      "300px",
					ChartSpec: map[string]any{
						"type": "line",
						"data": map[string]any{
							"values": []map[string]any{
								{"month": "Jan", "revenue": 100},
								{"month": "Feb", "revenue": 150},
							},
						},
					},
				},
				Hr(),
				&Table{
					PageSize:  10,
					RowHeight: "middle",
					HeaderStyle: &HeaderStyle{
						Bold:            &bold,
						BackgroundStyle: "grey",
					},
					Columns: []TableColumn{
						{Name: "product", DisplayName: "Product", DataType: DataTypeText},
						{Name: "sales", DisplayName: "Sales", DataType: DataTypeNumber, Format: &NumberFormat{Precision: 0, Symbol: "$", Separator: true}},
						{Name: "date", DisplayName: "Date", DataType: DataTypeDate, DateFormat: "YYYY-MM-DD"},
					},
					Rows: []map[string]any{
						{"product": "Widget A", "sales": 50000, "date": 1710547200000},
						{"product": "Widget B", "sales": 35000, "date": 1710633600000},
					},
				},
			},
		},
	}

	m := mustMarshalCard(t, card)
	elements := getBodyElements(t, m)
	if len(elements) != 3 {
		t.Fatalf("expected 3 elements, got %d", len(elements))
	}

	chart := elements[0].(map[string]any)
	if chart["tag"] != "chart" {
		t.Errorf("tag = %v", chart["tag"])
	}

	table := elements[2].(map[string]any)
	if table["tag"] != "table" {
		t.Errorf("tag = %v", table["tag"])
	}
}

// ---- Disabled component tests ----

func TestDisabledComponents(t *testing.T) {
	card := Card{
		Body: Body{
			Elements: Elements{
				&Button{
					Text:         PlainText("Disabled Button"),
					Disabled:     true,
					DisabledTips: PlainText("Cannot click"),
				},
				&Input{
					Name:         "disabled_input",
					Disabled:     true,
					DisabledTips: PlainText("Read only"),
				},
				&SelectStatic{
					Name:     "disabled_select",
					Disabled: true,
				},
				&DatePicker{
					Name:     "disabled_date",
					Disabled: true,
				},
				&InteractiveContainer{
					Disabled:     true,
					DisabledTips: PlainText("Locked"),
					Elements:     Elements{Md("Locked content")},
				},
			},
		},
	}

	m := mustMarshalCard(t, card)
	elements := getBodyElements(t, m)

	for i, elem := range elements {
		obj := elem.(map[string]any)
		if obj["disabled"] != true {
			t.Errorf("element[%d] disabled should be true", i)
		}
	}
}

// ---- Card with no header ----

func TestCardWithoutHeader(t *testing.T) {
	card := Card{
		Body: Body{
			Elements: Elements{Md("No header card")},
		},
	}

	m := mustMarshalCard(t, card)
	if m["schema"] != "2.0" {
		t.Error("schema missing")
	}

	// Header is a non-pointer struct, so it will be present but empty.
	if header, ok := m["header"]; ok {
		h := header.(map[string]any)
		if len(h) != 0 {
			t.Errorf("expected empty header, got %d fields", len(h))
		}
	}
	// Either absent or empty is fine for a zero-value header.
}

// ---- Containers nested inside each other (valid combos) ----

func TestColumnSetInsideCollapsiblePanel(t *testing.T) {
	card := Card{
		Body: Body{
			Elements: Elements{
				&CollapsiblePanel{
					Expanded: true,
					Header: &PanelHeader{
						Title: PlainText("Expand"),
					},
					Elements: Elements{
						&ColumnSet{
							Columns: []Column{
								{Width: "auto", Elements: Elements{Md("A")}},
								{Width: "auto", Elements: Elements{Md("B")}},
							},
						},
					},
				},
			},
		},
	}

	m := mustMarshalCard(t, card)
	elements := getBodyElements(t, m)
	cp := elements[0].(map[string]any)
	cpElements := cp["elements"].([]any)
	cs := cpElements[0].(map[string]any)
	if cs["tag"] != "column_set" {
		t.Errorf("nested tag = %v, want column_set", cs["tag"])
	}
}

func TestInteractiveContainerInsideColumn(t *testing.T) {
	card := Card{
		Body: Body{
			Elements: Elements{
				&ColumnSet{
					Columns: []Column{
						{
							Width: "weighted",
							Weight: 1,
							Elements: Elements{
								&InteractiveContainer{
									Behaviors: []Behavior{
										{Type: BehaviorOpenURL, DefaultURL: "https://example.com"},
									},
									Elements: Elements{
										Md("Click anywhere"),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	m := mustMarshalCard(t, card)
	elements := getBodyElements(t, m)
	cs := elements[0].(map[string]any)
	cols := cs["columns"].([]any)
	col := cols[0].(map[string]any)
	colElements := col["elements"].([]any)
	ic := colElements[0].(map[string]any)
	if ic["tag"] != "interactive_container" {
		t.Errorf("tag = %v", ic["tag"])
	}
}

// ---- Empty card ----

func TestMinimalCard(t *testing.T) {
	card := Card{}
	data, err := card.JSON()
	if err != nil {
		t.Fatal(err)
	}

	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	if m["schema"] != "2.0" {
		t.Error("schema should always be 2.0")
	}
}

// ---- JSON round-trip stability ----

func TestJSONRoundTrip(t *testing.T) {
	card := Card{
		Header: Header{
			Title:    PlainText("Test"),
			Template: Blue,
		},
		Body: Body{
			Elements: Elements{
				Md("hello"),
				Hr(),
				Btn("click"),
			},
		},
	}

	data1, err := card.JSON()
	if err != nil {
		t.Fatal(err)
	}

	// Unmarshal and re-marshal raw JSON to check stability.
	var raw any
	if err := json.Unmarshal(data1, &raw); err != nil {
		t.Fatal(err)
	}
	data2, err := json.Marshal(raw)
	if err != nil {
		t.Fatal(err)
	}

	// Both should produce valid JSON with same structure.
	var m1, m2 map[string]any
	json.Unmarshal(data1, &m1)
	json.Unmarshal(data2, &m2)

	// Compare key counts at top level.
	if len(m1) != len(m2) {
		t.Errorf("key count mismatch: %d vs %d", len(m1), len(m2))
	}
}

// ---- Flexible spacing values ----

func TestAllSpacingConstants(t *testing.T) {
	spacings := map[string]string{
		"small":       SpacingSmall,
		"medium":      SpacingMedium,
		"large":       SpacingLarge,
		"extra_large": SpacingExtraLarge,
	}

	for expected, val := range spacings {
		if val != expected {
			t.Errorf("Spacing %q = %q, want %q", expected, val, expected)
		}
	}
}
