package main

import (
	"fmt"
	"os"

	"github.com/sharki13/dockman/backend/cmd"
	"github.com/sharki13/dockman/backend/wrapper"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func main() {

	columns := []table.Column{
		{Title: "Type", Width: 20},
		{Title: "Name", Width: 20},
		{Title: "State", Width: 10},
	}

	rows := []table.Row{}

	dockerWrapper := wrapper.MakeDockerWrapper()
	dockerInstances, err := dockerWrapper.GetInstances()
	if err != nil && err != cmd.Err_CommandNotFound && err != wrapper.Err_DockerNotRunning {
		fmt.Println(err)
		return
	}

	multipassWrapper := wrapper.MakeMultipassWrapper()
	multipassInstances, err := multipassWrapper.GetInstances()
	if err != nil && err != cmd.Err_CommandNotFound {
		fmt.Println(err)
		return
	}

	instances := append(dockerInstances, multipassInstances...)
	for _, i := range instances {
		rows = append(rows, table.Row{i.Type.String(), i.Name, i.State.String()})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
