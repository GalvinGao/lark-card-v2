// Package main provides an e2e test runner that sends card JSON payloads
// to a real Lark chat via the official Lark SDK to verify they are accepted
// by the server.
//
// Usage:
//
//	cp .env.example .env   # fill in your credentials
//	go run ./internal/test/e2e
//
// Required .env variables:
//
//	LARK_APP_ID       — App ID from Lark Open Platform
//	LARK_APP_SECRET   — App Secret from Lark Open Platform
//	LARK_OPEN_ID      — Open ID of the user to send test cards to
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/joho/godotenv"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"

	larkcard "github.com/GalvinGao/lark-card-v2/lark"
)

const (
	// Lark rate limit error code.
	rateLimitCode = 99991400

	maxRetries   = 5
	baseDelay    = time.Second
	maxDelay     = 30 * time.Second
	delayBetween = 500 * time.Millisecond
)

// testCase defines a named card to send to a Lark chat.
type testCase struct {
	Name string
	Card larkcard.Card
}

func main() {
	_ = godotenv.Load()

	appID := os.Getenv("LARK_APP_ID")
	appSecret := os.Getenv("LARK_APP_SECRET")
	openID := os.Getenv("LARK_OPEN_ID")

	if appID == "" || appSecret == "" || openID == "" {
		fmt.Fprintln(os.Stderr, "LARK_APP_ID, LARK_APP_SECRET, and LARK_OPEN_ID are required")
		fmt.Fprintln(os.Stderr, "Copy .env.example to .env and fill in your credentials")
		os.Exit(1)
	}

	client := lark.NewClient(appID, appSecret,
		lark.WithLogLevel(larkcore.LogLevelWarn),
	)

	cases := buildTestCases()
	var failed int

	for i, tc := range cases {
		fmt.Printf("[%d/%d] %s ... ", i+1, len(cases), tc.Name)

		err := sendCard(client, openID, tc.Card)
		if err != nil {
			fmt.Printf("FAIL: %v\n", err)
			failed++
		} else {
			fmt.Println("OK")
		}

		if i < len(cases)-1 {
			time.Sleep(delayBetween)
		}
	}

	fmt.Printf("\n%d/%d passed\n", len(cases)-failed, len(cases))
	if failed > 0 {
		os.Exit(1)
	}
}

// sendCard marshals the card and sends it via the Lark IM API with retry on rate limit.
func sendCard(client *lark.Client, openID string, card larkcard.Card) error {
	cardJSON, err := card.JSON()
	if err != nil {
		return fmt.Errorf("marshal card: %w", err)
	}

	content := string(cardJSON)

	for attempt := range maxRetries {
		resp, err := client.Im.Message.Create(context.Background(),
			larkim.NewCreateMessageReqBuilder().
				ReceiveIdType(larkim.ReceiveIdTypeOpenId).
				Body(larkim.NewCreateMessageReqBodyBuilder().
					MsgType(larkim.MsgTypeInteractive).
					ReceiveId(openID).
					Content(content).
					Build()).
				Build(),
		)
		if err != nil {
			return fmt.Errorf("send message: %w", err)
		}

		if resp.Success() {
			return nil
		}

		if resp.Code == rateLimitCode {
			delay := backoff(attempt)
			fmt.Printf("rate limited, retrying in %s ... ", delay)
			time.Sleep(delay)
			continue
		}

		return fmt.Errorf("lark error code=%d msg=%s", resp.Code, resp.Msg)
	}

	return errors.New("exceeded max retries due to rate limiting")
}

// backoff returns an exponential backoff duration capped at maxDelay.
func backoff(attempt int) time.Duration {
	delay := time.Duration(float64(baseDelay) * math.Pow(2, float64(attempt)))
	if delay > maxDelay {
		delay = maxDelay
	}
	return delay
}

// buildTestCases returns all cards that will be sent for validation.
func buildTestCases() []testCase {
	overallCheckable := true
	bold := true

	return []testCase{
		{
			Name: "simple markdown card",
			Card: larkcard.Card{
				Header: larkcard.Header{
					Title:    larkcard.PlainText("E2E Test: Simple Card"),
					Template: larkcard.Blue,
				},
				Body: larkcard.Body{
					Elements: larkcard.Elements{
						larkcard.Md("This is a **simple** e2e test card."),
					},
				},
			},
		},
		{
			Name: "card with config and card_link",
			Card: larkcard.Card{
				Config: &larkcard.Config{
					WidthMode: larkcard.WidthFill,
				},
				CardLink: &larkcard.CardLink{
					URL: "https://github.com/GalvinGao/lark-card-v2",
				},
				Header: larkcard.Header{
					Title:    larkcard.PlainText("E2E Test: Config & Link"),
					Subtitle: larkcard.PlainText("Click card to open repo"),
					Template: larkcard.Indigo,
				},
				Body: larkcard.Body{
					Elements: larkcard.Elements{
						larkcard.Md("Card with `config.width_mode=fill` and a `card_link`."),
					},
				},
			},
		},
		{
			Name: "header with text tags",
			Card: larkcard.Card{
				Header: larkcard.Header{
					Title:    larkcard.PlainText("E2E Test: Header Tags"),
					Template: larkcard.Green,
					TextTagList: []larkcard.TextTag{
						{Text: larkcard.PlainText("Tag1"), Color: larkcard.Blue},
						{Text: larkcard.PlainText("Tag2"), Color: larkcard.Red},
					},
				},
				Body: larkcard.Body{
					Elements: larkcard.Elements{
						larkcard.Md("Header should show two colored tags."),
					},
				},
			},
		},
		{
			Name: "column set layout",
			Card: larkcard.Card{
				Header: larkcard.Header{
					Title:    larkcard.PlainText("E2E Test: ColumnSet"),
					Template: larkcard.Violet,
				},
				Body: larkcard.Body{
					Elements: larkcard.Elements{
						&larkcard.ColumnSet{
							HorizontalSpacing: larkcard.SpacingMedium,
							Columns: []larkcard.Column{
								{
									Width:  "weighted",
									Weight: 2,
									Elements: larkcard.Elements{
										larkcard.Md("**Left** (weight=2)"),
									},
								},
								{
									Width:  "weighted",
									Weight: 1,
									Elements: larkcard.Elements{
										larkcard.Md("**Right** (weight=1)"),
									},
								},
								{
									Width: "auto",
									Elements: larkcard.Elements{
										&larkcard.Button{Text: larkcard.PlainText("Auto"), Type: larkcard.ButtonPrimary, Size: larkcard.SizeSmall},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Name: "table with columns and rows",
			Card: larkcard.Card{
				Header: larkcard.Header{
					Title:    larkcard.PlainText("E2E Test: Table"),
					Template: larkcard.Turquoise,
				},
				Body: larkcard.Body{
					Elements: larkcard.Elements{
						&larkcard.Table{
							PageSize:  5,
							RowHeight: "middle",
							HeaderStyle: &larkcard.HeaderStyle{
								Bold:            &bold,
								BackgroundStyle: "grey",
								TextAlign:       "center",
							},
							Columns: []larkcard.TableColumn{
								{Name: "name", DisplayName: "Name", DataType: larkcard.DataTypeText},
								{Name: "value", DisplayName: "Value", DataType: larkcard.DataTypeNumber, Format: &larkcard.NumberFormat{Precision: 2, Symbol: "$", Separator: true}},
							},
							Rows: []map[string]any{
								{"name": "Item A", "value": 1234.56},
								{"name": "Item B", "value": 78.90},
							},
						},
					},
				},
			},
		},
		{
			Name: "form with inputs and selects",
			Card: larkcard.Card{
				Header: larkcard.Header{
					Title:    larkcard.PlainText("E2E Test: Form"),
					Template: larkcard.Purple,
				},
				Body: larkcard.Body{
					Elements: larkcard.Elements{
						&larkcard.Form{
							Name: "e2e_form",
							Elements: larkcard.Elements{
								&larkcard.Input{
									Name:        "username",
									Placeholder: larkcard.PlainText("Enter name"),
									MaxLength:   100,
								},
								&larkcard.SelectStatic{
									Name:        "priority",
									Placeholder: larkcard.PlainText("Select"),
									Options: []larkcard.SelectOption{
										{Text: larkcard.PlainText("High"), Value: "high"},
										{Text: larkcard.PlainText("Low"), Value: "low"},
									},
								},
								&larkcard.DatePicker{
									Name:        "due_date",
									Placeholder: larkcard.PlainText("Pick date"),
								},
								&larkcard.Button{
									Text:           larkcard.PlainText("Submit"),
									Type:           larkcard.ButtonPrimary,
									FormActionType: larkcard.FormSubmit,
									Name:           "submit",
								},
							},
						},
					},
				},
			},
		},
		{
			Name: "interactive container with behavior",
			Card: larkcard.Card{
				Header: larkcard.Header{
					Title:    larkcard.PlainText("E2E Test: Interactive Container"),
					Template: larkcard.Wathet,
				},
				Body: larkcard.Body{
					Elements: larkcard.Elements{
						&larkcard.InteractiveContainer{
							Width:           "fill",
							BackgroundStyle: "default",
							HasBorder:       true,
							CornerRadius:    "8px",
							Padding:         "12px",
							Behaviors: []larkcard.Behavior{
								{
									Type:       larkcard.BehaviorOpenURL,
									DefaultURL: "https://github.com/GalvinGao/lark-card-v2",
								},
							},
							Elements: larkcard.Elements{
								larkcard.Md("Click this container to open a URL."),
							},
						},
					},
				},
			},
		},
		{
			Name: "collapsible panel",
			Card: larkcard.Card{
				Header: larkcard.Header{
					Title:    larkcard.PlainText("E2E Test: Collapsible Panel"),
					Template: larkcard.Yellow,
				},
				Body: larkcard.Body{
					Elements: larkcard.Elements{
						&larkcard.CollapsiblePanel{
							Expanded: true,
							Header: &larkcard.PanelHeader{
								Title:             larkcard.PlainText("Click to collapse"),
								IconExpandedAngle: 180,
							},
							Border: &larkcard.BorderConfig{
								Color:        "grey",
								CornerRadius: "4px",
							},
							Elements: larkcard.Elements{
								larkcard.Md("Hidden content inside the panel."),
								larkcard.Hr(),
								larkcard.Md("More content after a divider."),
							},
						},
					},
				},
			},
		},
		{
			Name: "checker task list",
			Card: larkcard.Card{
				Header: larkcard.Header{
					Title:    larkcard.PlainText("E2E Test: Checkers"),
					Template: larkcard.Carmine,
				},
				Body: larkcard.Body{
					Elements: larkcard.Elements{
						&larkcard.Checker{
							Name:             "done_task",
							Checked:          true,
							Text:             larkcard.PlainText("Completed task"),
							OverallCheckable: &overallCheckable,
							CheckedStyle: &larkcard.CheckedStyle{
								ShowStrikethrough: true,
								Opacity:           0.5,
							},
						},
						&larkcard.Checker{
							Name:             "pending_task",
							Checked:          false,
							Text:             larkcard.PlainText("Pending task"),
							OverallCheckable: &overallCheckable,
							ButtonArea: &larkcard.CheckerButtonArea{
								Buttons: []larkcard.Button{
									{Text: larkcard.PlainText("Edit"), Type: larkcard.ButtonPrimaryText, Size: larkcard.SizeTiny},
								},
							},
						},
					},
				},
			},
		},
		{
			Name: "overflow menu",
			Card: larkcard.Card{
				Header: larkcard.Header{
					Title:    larkcard.PlainText("E2E Test: Overflow"),
					Template: larkcard.Grey,
				},
				Body: larkcard.Body{
					Elements: larkcard.Elements{
						larkcard.Md("Card with an overflow menu:"),
						&larkcard.Overflow{
							Options: []larkcard.OverflowOption{
								{Text: larkcard.PlainText("Option A"), Value: "a"},
								{Text: larkcard.PlainText("Option B"), Value: "b"},
								{Text: larkcard.PlainText("Option C"), Value: "c"},
							},
						},
					},
				},
			},
		},
		{
			Name: "all button types",
			Card: larkcard.Card{
				Header: larkcard.Header{
					Title:    larkcard.PlainText("E2E Test: Button Types"),
					Template: larkcard.Blue,
				},
				Body: larkcard.Body{
					Elements: larkcard.Elements{
						&larkcard.ColumnSet{
							FlexMode:          larkcard.FlexFlow,
							HorizontalSpacing: larkcard.SpacingSmall,
							Columns: []larkcard.Column{
								{Width: "auto", Elements: larkcard.Elements{&larkcard.Button{Text: larkcard.PlainText("default"), Type: larkcard.ButtonDefault, Size: larkcard.SizeSmall}}},
								{Width: "auto", Elements: larkcard.Elements{&larkcard.Button{Text: larkcard.PlainText("primary"), Type: larkcard.ButtonPrimary, Size: larkcard.SizeSmall}}},
								{Width: "auto", Elements: larkcard.Elements{&larkcard.Button{Text: larkcard.PlainText("danger"), Type: larkcard.ButtonDanger, Size: larkcard.SizeSmall}}},
								{Width: "auto", Elements: larkcard.Elements{&larkcard.Button{Text: larkcard.PlainText("text"), Type: larkcard.ButtonText, Size: larkcard.SizeSmall}}},
								{Width: "auto", Elements: larkcard.Elements{&larkcard.Button{Text: larkcard.PlainText("laser"), Type: larkcard.ButtonLaser, Size: larkcard.SizeSmall}}},
							},
						},
					},
				},
			},
		},
		{
			Name: "deeply nested columns (3 levels)",
			Card: larkcard.Card{
				Header: larkcard.Header{
					Title:    larkcard.PlainText("E2E Test: Nested Columns"),
					Template: larkcard.Red,
				},
				Body: larkcard.Body{
					Elements: larkcard.Elements{
						&larkcard.ColumnSet{
							Columns: []larkcard.Column{
								{
									Width:  "weighted",
									Weight: 1,
									Elements: larkcard.Elements{
										&larkcard.ColumnSet{
											Columns: []larkcard.Column{
												{
													Width:  "weighted",
													Weight: 1,
													Elements: larkcard.Elements{
														&larkcard.ColumnSet{
															Columns: []larkcard.Column{
																{Width: "auto", Elements: larkcard.Elements{larkcard.Md("**Level 3**")}},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Width:    "auto",
									Elements: larkcard.Elements{larkcard.Md("Sibling column")},
								},
							},
						},
					},
				},
			},
		},
		{
			Name: "div with icon",
			Card: larkcard.Card{
				Header: larkcard.Header{
					Title:    larkcard.PlainText("E2E Test: Div"),
					Template: larkcard.Default,
				},
				Body: larkcard.Body{
					Elements: larkcard.Elements{
						&larkcard.Div{
							Text: &larkcard.Text{
								Tag:       "plain_text",
								Content:   "Status: Running",
								TextSize:  larkcard.TextSizeNormal,
								TextColor: "green",
							},
							Icon: &larkcard.Icon{
								Tag:   larkcard.StandardIcon,
								Token: "status-doing_outlined",
								Color: larkcard.Green,
							},
						},
					},
				},
			},
		},
		{
			Name: "person list",
			Card: larkcard.Card{
				Header: larkcard.Header{
					Title:    larkcard.PlainText("E2E Test: Person List"),
					Template: larkcard.Orange,
				},
				Body: larkcard.Body{
					Elements: larkcard.Elements{
						larkcard.Md("Person elements require valid user IDs to render."),
					},
				},
			},
		},
		{
			Name: "kitchen sink — mixed elements",
			Card: buildKitchenSinkCard(),
		},
	}
}

// buildKitchenSinkCard constructs a complex card that exercises multiple component types.
func buildKitchenSinkCard() larkcard.Card {
	return larkcard.Card{
		Config: &larkcard.Config{
			WidthMode: larkcard.WidthFill,
		},
		Header: larkcard.Header{
			Title:    larkcard.PlainText("E2E Test: Kitchen Sink"),
			Subtitle: larkcard.PlainText("Multiple component types in one card"),
			Template: larkcard.Purple,
			TextTagList: []larkcard.TextTag{
				{Text: larkcard.PlainText("e2e"), Color: larkcard.Green},
			},
		},
		Body: larkcard.Body{
			Elements: larkcard.Elements{
				// Section 1: text
				larkcard.Md("**Section 1:** Rich text with `code`, **bold**, and *italic*."),
				larkcard.Hr(),

				// Section 2: columns
				&larkcard.ColumnSet{
					HorizontalSpacing: larkcard.SpacingSmall,
					Columns: []larkcard.Column{
						{
							Width:  "weighted",
							Weight: 1,
							Elements: larkcard.Elements{
								&larkcard.Div{
									Text: larkcard.MdText("Left column content"),
								},
							},
						},
						{
							Width:  "weighted",
							Weight: 1,
							Elements: larkcard.Elements{
								&larkcard.Div{
									Text: larkcard.MdText("Right column content"),
								},
							},
						},
					},
				},
				larkcard.Hr(),

				// Section 3: interactive
				&larkcard.CollapsiblePanel{
					Expanded: true,
					Header: &larkcard.PanelHeader{
						Title: larkcard.PlainText("Collapsible Section"),
					},
					Elements: larkcard.Elements{
						larkcard.Md("Inside the panel."),
						&larkcard.Button{Text: larkcard.PlainText("Action"), Type: larkcard.ButtonPrimary},
					},
				},

				// Section 4: select
				&larkcard.SelectStatic{
					Name:        "kitchen_select",
					Placeholder: larkcard.PlainText("Pick one"),
					Options: []larkcard.SelectOption{
						{Text: larkcard.PlainText("Alpha"), Value: "a"},
						{Text: larkcard.PlainText("Beta"), Value: "b"},
					},
				},
			},
		},
	}
}

func init() {
	// Validate that our cards produce valid JSON before sending.
	for _, tc := range buildTestCases() {
		data, err := tc.Card.JSON()
		if err != nil {
			panic(fmt.Sprintf("test case %q: card.JSON() failed: %v", tc.Name, err))
		}
		var m map[string]any
		if err := json.Unmarshal(data, &m); err != nil {
			panic(fmt.Sprintf("test case %q: invalid JSON: %v", tc.Name, err))
		}
	}
}
