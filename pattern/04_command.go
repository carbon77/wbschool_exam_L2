package pattern

import (
	"fmt"
)

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Editor struct{}

func (e *Editor) Save() {
	fmt.Println("Сохранение состояние редактора")
}

func (e *Editor) GetSelection() {
	fmt.Println("Возврат выбранного текста")
}

func (e *Editor) DeleteSelection() {
	fmt.Println("Удаление выбранного текста")
}

func (e *Editor) ReplaceSelection() {
	fmt.Println("Вставка текста из буфера обмена")
}

type Command interface {
	Execute()
}

type CopyCommand struct {
	editor *Editor
}

func (cc *CopyCommand) Execute() {
	cc.editor.GetSelection()
}

type DeleteCommand struct {
	editor *Editor
}

func (dc *DeleteCommand) Execute() {
	dc.editor.DeleteSelection()
}

type PasteCommand struct {
	editor *Editor
}

func (pc *PasteCommand) Execute() {
	pc.editor.ReplaceSelection()
}

type Invoker interface {
	SetCommand(command Command)
	ExecuteCommand()
}

type invokerImpl struct {
	command Command
}

func (impl *invokerImpl) SetCommand(command Command) {
	impl.command = command
}
func (impl *invokerImpl) ExecuteCommand() {
	impl.command.Execute()
}

type Button struct {
	invokerImpl
}
type Shortcut struct {
	invokerImpl
}

func TestCommand() {
	editor := &Editor{}

	copyButton := &Button{}
	copyButton.SetCommand(&CopyCommand{editor})

	deleteButton := &Button{}
	deleteButton.SetCommand(&DeleteCommand{editor})

	deleteShortcut := &Shortcut{}
	deleteShortcut.SetCommand(&DeleteCommand{editor})

	copyButton.ExecuteCommand()
	deleteShortcut.ExecuteCommand()
	deleteButton.ExecuteCommand()
}
