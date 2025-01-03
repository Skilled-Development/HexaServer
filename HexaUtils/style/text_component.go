package style

import (
	"encoding/json"
	"fmt"
)

// TextComponent represents a Minecraft raw JSON text component
type TextComponent struct {
	Text          string          `json:"text,omitempty"`
	Translate     string          `json:"translate,omitempty"`
	Score         *Score          `json:"score,omitempty"`
	Selector      string          `json:"selector,omitempty"`
	Keybind       string          `json:"keybind,omitempty"`
	NBT           string          `json:"nbt,omitempty"`
	Extra         []TextComponent `json:"extra,omitempty"`
	Color         string          `json:"color,omitempty"`
	Bold          bool            `json:"bold,omitempty"`
	Italic        bool            `json:"italic,omitempty"`
	Underlined    bool            `json:"underlined,omitempty"`
	Strikethrough bool            `json:"strikethrough,omitempty"`
	Obfuscated    bool            `json:"obfuscated,omitempty"`
	Insertion     string          `json:"insertion,omitempty"`
	ClickEvent    *ClickEvent     `json:"clickEvent,omitempty"`
	HoverEvent    *HoverEvent     `json:"hoverEvent,omitempty"`
	Font          string          `json:"font,omitempty"`
}

// Score represents the score of a player or entity in the Minecraft scoreboard
type Score struct {
	Name      string `json:"name"`
	Objective string `json:"objective"`
}

// ClickEvent represents the click event with the associated action and value
type ClickEvent struct {
	Action string `json:"action"`
	Value  string `json:"value"`
}

// HoverEvent represents the hover event with the associated action and content
type HoverEvent struct {
	Action   string      `json:"action"`
	Contents interface{} `json:"contents"`
}

// NewTextComponent creates a new text component with optional text content
func NewTextComponent(text string) *TextComponent {
	return &TextComponent{
		Text: text,
	}
}

// SetColor sets the color of the text component
func (tc *TextComponent) SetColor(color string) *TextComponent {
	tc.Color = color
	return tc
}

// SetBold sets the bold formatting of the text component
func (tc *TextComponent) SetBold(bold bool) *TextComponent {
	tc.Bold = bold
	return tc
}

// SetItalic sets the italic formatting of the text component
func (tc *TextComponent) SetItalic(italic bool) *TextComponent {
	tc.Italic = italic
	return tc
}

// SetUnderlined sets the underlined formatting of the text component
func (tc *TextComponent) SetUnderlined(underlined bool) *TextComponent {
	tc.Underlined = underlined
	return tc
}

// SetStrikethrough sets the strikethrough formatting of the text component
func (tc *TextComponent) SetStrikethrough(strikethrough bool) *TextComponent {
	tc.Strikethrough = strikethrough
	return tc
}

// SetObfuscated sets the obfuscated (glitched) text formatting
func (tc *TextComponent) SetObfuscated(obfuscated bool) *TextComponent {
	tc.Obfuscated = obfuscated
	return tc
}

// SetClickEvent sets the click event for the text component
func (tc *TextComponent) SetClickEvent(action, value string) *TextComponent {
	tc.ClickEvent = &ClickEvent{
		Action: action,
		Value:  value,
	}
	return tc
}

// SetHoverEvent sets the hover event for the text component
func (tc *TextComponent) SetHoverEvent(action string, contents interface{}) *TextComponent {
	tc.HoverEvent = &HoverEvent{
		Action:   action,
		Contents: contents,
	}
	return tc
}

// AddExtra adds an additional component to the "extra" array
func (tc *TextComponent) AddExtra(extra *TextComponent) *TextComponent {
	tc.Extra = append(tc.Extra, *extra)
	return tc
}

// SetScore sets the score display for a player
func (tc *TextComponent) SetScore(name, objective string) *TextComponent {
	tc.Score = &Score{
		Name:      name,
		Objective: objective,
	}
	return tc
}

// SetNBT sets the NBT path for the text component
func (tc *TextComponent) SetNBT(nbt string) *TextComponent {
	tc.NBT = nbt
	return tc
}

// ToJSON marshals the TextComponent into a JSON string
func (tc *TextComponent) ToJSON() (string, error) {
	jsonData, err := json.Marshal(tc)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// ToMap converts the TextComponent to a map[string]interface{} manually
func (tc *TextComponent) ToMap() map[string]interface{} {
	result := make(map[string]interface{})

	if tc.Text != "" {
		result["text"] = tc.Text
	}
	if tc.Translate != "" {
		result["translate"] = tc.Translate
	}
	if tc.Score != nil {
		result["score"] = map[string]interface{}{
			"name":      tc.Score.Name,
			"objective": tc.Score.Objective,
		}
	}
	if tc.Selector != "" {
		result["selector"] = tc.Selector
	}
	if tc.Keybind != "" {
		result["keybind"] = tc.Keybind
	}
	if tc.NBT != "" {
		result["nbt"] = tc.NBT
	}
	if len(tc.Extra) > 0 {
		var extraList []map[string]interface{}
		for _, extra := range tc.Extra {
			extraList = append(extraList, extra.ToMap())
		}
		result["extra"] = extraList
	}
	if tc.Color != "" {
		result["color"] = tc.Color
	}
	if tc.Bold {
		result["bold"] = tc.Bold
	}
	if tc.Italic {
		result["italic"] = tc.Italic
	}
	if tc.Underlined {
		result["underlined"] = tc.Underlined
	}
	if tc.Strikethrough {
		result["strikethrough"] = tc.Strikethrough
	}
	if tc.Obfuscated {
		result["obfuscated"] = tc.Obfuscated
	}
	if tc.Insertion != "" {
		result["insertion"] = tc.Insertion
	}
	if tc.ClickEvent != nil {
		result["clickEvent"] = map[string]interface{}{
			"action": tc.ClickEvent.Action,
			"value":  tc.ClickEvent.Value,
		}
	}

	if tc.HoverEvent != nil {
		hoverEventMap := map[string]interface{}{
			"action": tc.HoverEvent.Action,
		}

		if contentTc, ok := tc.HoverEvent.Contents.(*TextComponent); ok {
			hoverEventMap["contents"] = contentTc.ToMap()
		} else {
			hoverEventMap["contents"] = tc.HoverEvent.Contents
		}

		result["hoverEvent"] = hoverEventMap

	}

	if tc.Font != "" {
		result["font"] = tc.Font
	}

	return result
}

// convertContent converts interface{} content of HoverEvent to map[string]interface{} or keep as is
func convertContent(content interface{}) interface{} {
	if tc, ok := content.(*TextComponent); ok {
		return tc.ToMap()
	}
	return content
}

func main() {
	// Example of creating a complex JSON text
	tc := NewTextComponent("Hello, world!").
		SetColor("red").
		SetBold(true).
		SetItalic(true).
		SetClickEvent("open_url", "https://example.com").
		SetHoverEvent("show_text", NewTextComponent("Click here!")).
		AddExtra(NewTextComponent(" More text...").SetColor("blue")).
		SetScore("Test", "test_objective")

	// Convert to JSON
	jsonText, err := tc.ToJSON()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(jsonText)

	// convert to map
	mapText := tc.ToMap()
	fmt.Println(mapText)

	tc2 := NewTextComponent("Hover Test").SetHoverEvent("show_text", NewTextComponent("Click here!").SetColor("blue").SetBold(true))
	mapText2 := tc2.ToMap()
	fmt.Println(mapText2)

	tc3 := NewTextComponent("Complex Hover Test").SetHoverEvent("show_text", NewTextComponent("Click here!").SetColor("blue").SetBold(true).AddExtra(NewTextComponent(" and more")).AddExtra(NewTextComponent(" and another")).SetScore("Test2", "test_objective_2"))
	mapText3 := tc3.ToMap()
	fmt.Println(mapText3)
}
