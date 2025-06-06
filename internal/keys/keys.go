package keys

import "github.com/charmbracelet/bubbles/key"

// KeyMap holds the key bindings for the table.
type KeyMap struct {
	Bottom         key.Binding
	Top            key.Binding
	Down           key.Binding
	Up             key.Binding
	CollapseToggle key.Binding
	CollapseAll    key.Binding
	ExpandAll      key.Binding
	Help           key.Binding
	Quit           key.Binding
	Search         key.Binding
	Submit         key.Binding
	Next           key.Binding
}

// Len returns the number of keys in the keymap.
func (KeyMap) Len() int {
	// get number of keys in the keymap
	return 12
}

func NewKeyMap(c *KeyConfig) KeyMap {
	keys := DefaultKeyMap()
	if len(c.BottomKeys) != 0 {
		keys.Bottom.SetKeys(c.BottomKeys...)
	}
	if len(c.TopKeys) != 0 {
		keys.Top.SetKeys(c.TopKeys...)
	}
	if len(c.DownKeys) != 0 {
		keys.Down.SetKeys(c.DownKeys...)
	}
	if len(c.UpKeys) != 0 {
		keys.Up.SetKeys(c.UpKeys...)
	}
	if len(c.CollapseToggleKeys) != 0 {
		keys.CollapseToggle.SetKeys(c.CollapseToggleKeys...)
	}
	if len(c.CollapseAllKeys) != 0 {
		keys.CollapseAll.SetKeys(c.CollapseAllKeys...)
	}
	if len(c.ExpandAllKeys) != 0 {
		keys.ExpandAll.SetKeys(c.ExpandAllKeys...)
	}
	if len(c.HelpKeys) != 0 {
		keys.Help.SetKeys(c.HelpKeys...)
	}
	if len(c.QuitKeys) != 0 {
		keys.Quit.SetKeys(c.QuitKeys...)
	}
	if len(c.SearchKeys) != 0 {
		keys.Search.SetKeys(c.SearchKeys...)
	}
	if len(c.SubmitKeys) != 0 {
		keys.Submit.SetKeys(c.SubmitKeys...)
	}
	if len(c.NextKeys) != 0 {
		keys.Next.SetKeys(c.NextKeys...)
	}
	return keys
}

// DefaultKeyMap is the default key bindings for the table.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Bottom: key.NewBinding(
			key.WithKeys("bottom", "G"),
			key.WithHelp("bottom/G", "bottom"),
		),
		Top: key.NewBinding(
			key.WithKeys("top", "g"),
			key.WithHelp("top/g", "top"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		CollapseToggle: key.NewBinding(
			key.WithKeys("tab", "h", "l"),
			key.WithHelp("tab/h/l", "collapse/expand selected"),
		),
		CollapseAll: key.NewBinding(
			key.WithKeys("<", "H"),
			key.WithHelp("</H", "collapse all"),
		),
		ExpandAll: key.NewBinding(
			key.WithKeys(">", "L"),
			key.WithHelp(">/L", "expand all"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "esc"),
			key.WithHelp("esc", "return"),
		),
		Search: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "search for matching nodes"),
		),
		Submit: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "submit entry"),
		),
		Next: key.NewBinding(
			key.WithKeys("n"),
			key.WithHelp("n", "go to next match"),
		),
	}
}
