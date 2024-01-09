package main

import "fmt"

type Human struct {
	Name string
	Job  string
}

func NewHuman(name string, job string) *Human {
	return &Human{
		Name: name,
		Job:  job,
	}
}

func (h *Human) Hire(job string) {
	h.Job = job
}

func (h *Human) Info() {
	fmt.Println(h.Name, h.Job)
}

type Action struct {
	*Human
	Action string
}

func NewAction(human *Human, action string) Action {
	return Action{
		Human:  human,
		Action: action,
	}
}

func main() {
	john := NewHuman("John Doe", "Developer")
	act := NewAction(john, "die")
	act.Info()
	act.Hire("Engineer")
	act.Info()
}
