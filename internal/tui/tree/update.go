package tree

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/crosleyzack/xplr/internal/nodes"
)

// Update the JSON component
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m == nil {
		return nil, nil
	}
	switch msg := msg.(type) {
	case tea.QuitMsg:
		return m, tea.Batch(tea.ClearScreen, tea.Quit)
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		if m.mode == ModeSearch {
			if key.Matches(msg, m.KeyMap.Submit) {
				m.mode = ModeNormal
				m.getMatchingNodes()
			} else {
				m.searchTerm += msg.String()
			}
			break
		}
		switch {
		case key.Matches(msg, m.KeyMap.Bottom):
			m.cursor = m.NumberOfNodes()
		case key.Matches(msg, m.KeyMap.Top):
			m.cursor = 0
		case key.Matches(msg, m.KeyMap.Down):
			m.NavDown()
		case key.Matches(msg, m.KeyMap.Up):
			m.NavUp()
		case key.Matches(msg, m.KeyMap.Collapse):
			m.InvertCollaped()
		case key.Matches(msg, m.KeyMap.Help):
			m.Help.ShowAll = !m.Help.ShowAll
		case key.Matches(msg, m.KeyMap.Quit):
			return m, tea.Batch(tea.Quit, tea.ClearScreen)
		case key.Matches(msg, m.KeyMap.Search):
			m.searchTerm = ""
			m.mode = ModeSearch
		case key.Matches(msg, m.KeyMap.Next):
			if len(m.searchResults) > 0 {
				m.cursor, m.searchResults = m.searchResults[0], m.searchResults[1:]
			}
		}
	}
	return m, nil
}

// NavUp moves the cursor up in component
func (m *Model) NavUp() {
	m.cursor--
	if m.cursor < 0 {
		m.cursor = 0
		return
	}
}

// NavDown moves the cursor down in component
func (m *Model) NavDown() {
	m.cursor++
	if m.cursor >= m.NumberOfNodes() {
		m.cursor = m.NumberOfNodes() - 1
		return
	}
}

// InvertCollaped inverts the collapsed state of the current node
func (m *Model) InvertCollaped() {
	if m.currentNode != nil && m.currentNode.Children != nil {
		m.currentNode.Expand = !m.currentNode.Expand
	}
}

// getMatchingNodes find nodes which match request
func (m *Model) getMatchingNodes() error {
	count := 0
	f := func(node *nodes.Node, layer int) error {
		if len(node.Children) == 0 && strings.Contains(node.Value, m.searchTerm) {
			// expand parents
			n := node.Parent
			for n != nil {
				n.Expand = true
				n = n.Parent
			}
			m.searchResults = append(m.searchResults, count)
		}
		count++
		return nil
	}
	// todo simplify
	err := nodes.DFS(m.Nodes, f, 0, true)
	if err != nil {
		return err
	}
	return nil
}
