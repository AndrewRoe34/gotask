package ui

import (
	"github.com/EvoSched/gotask/internal/models"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	Main state = iota
	Pending
	Archived
	Detail
	List
	Help
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	main  table.Model
	pend  table.Model
	arch  table.Model
	task  models.Task
	tasks []models.Task
	list  list.Model
	state state
}

func (m model) Init() tea.Cmd {
	return nil
}

func New() {

}
