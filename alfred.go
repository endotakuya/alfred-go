package alfred

import (
	"encoding/json"
	"fmt"
)

type (
	Alfred struct {
		Items []Item
	}

	Item struct {
		UID          string `json:"uid,omitempty"`
		Title        string `json:"title,omitempty"`
		SubTitle     string `json:"subtitle,omitempty"`
		Arg          string `json:"arg,omitempty"`
		Icon         *Icon  `json:"icon,omitempty"`
		Valid        bool   `json:"valid,omitempty"`
		Match        string `json:"match,omitempty"`
		AutoComplete string `json:"autocomplete,omitempty"`
		Type         string `json:"type,omitempty"`
		Mods         *Mods  `json:"mods,omitempty"`
		Text         *Text  `json:"text,omitempty"`
		QuickLookURL string `json:"quicklookurl,omitempty"`
	}

	Icon struct {
		Type string `json:"type"`
		Path string `json:"path"`
	}

	Mods struct {
		ModAlt Mod `json:"alt"`
		ModCmd Mod `json:"cmd"`
	}

	Mod struct {
		Valid    bool   `json:"valid"`
		Arg      string `json:"arg"`
		SubTitle string `json:"subitle"`
	}

	Text struct {
		Copy      string `json:"copy"`
		LargeType string `json:"largetype"`
	}
)

const (
	TYPE_DEFAULT        = "default"
	TYPE_FILE           = "file"
	TYPE_FILE_SKIPCHECK = "file:skipcheck"
	TYPE_FILE_ICON      = "fileicon"
)

// New create item wrapper
func New() *Alfred {
	return &Alfred{
		Items: []Item{},
	}
}

// Add item to item wrapper
func (a *Alfred) Append(item *Item) {
	a.Items = append(a.Items, *item)
}

// Print is ...
func (a *Alfred) Print() {
	m := map[string][]Item{"items": a.Items}
	j, _ := json.Marshal(m)
	fmt.Printf("%s", j)
}
