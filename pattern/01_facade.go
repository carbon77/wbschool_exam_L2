package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/
import (
	"fmt"
)

// Токенайзер - преобразует текст в токены
type Tokenizer struct{}

func (t *Tokenizer) Tokenize(program string) string {
	fmt.Printf("Токенизация: %s\n", program)
	return program
}

// Парсер - преобразует токены в абстрактное синтаксическое дерево (AST)
type Parser struct{}

func (p *Parser) Parse(tokens string) string {
	fmt.Printf("Парсинг: %s\n", tokens)
	return tokens
}

// Вычислитель - обходит синтаксическое дерево и вычисляет результат
type Evaluator struct{}

func (e *Evaluator) Evaluate(ast string) string {
	fmt.Printf("Выполнение: %s\n", ast)
	return ast
}

// Интерпретатор - принимает программу на выполнение последовательно передает ее
// токенизатору, парсеру и вычислителю и печатает результат
type Interpreter struct {
	tokenizer *Tokenizer
	parser    *Parser
	evaluator *Evaluator
}

func (i *Interpreter) Evaluate(program string) {
	tokens := i.tokenizer.Tokenize(program)
	ast := i.parser.Parse(tokens)
	result := i.evaluator.Evaluate(ast)
	fmt.Printf("Результат: %s\n", result)
}

func NewInterpreter() *Interpreter {
	return &Interpreter{&Tokenizer{}, &Parser{}, &Evaluator{}}
}

func TestFacade() {
	interpreter := NewInterpreter()
	interpreter.Evaluate("program")
}
