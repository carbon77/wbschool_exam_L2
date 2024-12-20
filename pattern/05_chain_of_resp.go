package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request *SupportRequest) string
}

type BaseHandler struct {
	nextHandler Handler
}

func (bh *BaseHandler) SetNext(handler Handler) Handler {
	bh.nextHandler = handler
	return handler
}

func (bh *BaseHandler) Handle(request *SupportRequest) string {
	if bh.nextHandler != nil {
		return bh.nextHandler.Handle(request)
	}
	return "Никто не может обработать данный запрос"
}

type JuniorSupportHandler struct {
	BaseHandler
}

func (jh *JuniorSupportHandler) Handle(request *SupportRequest) string {
	if request.Level == 1 {
		return fmt.Sprintf("Младший сотрудник обработал запрос: %s", request.Issue)
	}
	return jh.BaseHandler.Handle(request)
}

type SeniorSupportHandler struct {
	BaseHandler
}

func (sh *SeniorSupportHandler) Handle(request *SupportRequest) string {
	if request.Level == 2 {
		return fmt.Sprintf("Старший сотрудник обработал запрос: %s", request.Issue)
	}
	return sh.BaseHandler.Handle(request)
}

type ManagerHandler struct {
	BaseHandler
}

func (mh *ManagerHandler) Handle(request *SupportRequest) string {
	if request.Level == 3 {
		return fmt.Sprintf("Менеджер обработал запрос: %s", request.Issue)
	}
	return mh.BaseHandler.Handle(request)
}

type SupportRequest struct {
	Level int
	Issue string
}

func TestChain() {
	junior := &JuniorSupportHandler{}
	senior := &SeniorSupportHandler{}
	manager := &ManagerHandler{}

	junior.SetNext(senior).SetNext(manager)

	requests := []SupportRequest{
		{Issue: "low level issue", Level: 1},
		{Issue: "high level issue", Level: 3},
		{Issue: "issue", Level: 2},
	}

	for _, request := range requests {
		fmt.Println(junior.Handle(&request))
	}
}
