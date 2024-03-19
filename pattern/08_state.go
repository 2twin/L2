package pattern

import "fmt"

/*
	Состояние — это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости
	от своего состояния. Извне создаётся впечатление, что изменился класс объекта.

	Плюсы:
	- Избавляет от множества больших условных операторов машины состояний.
	- Концентрирует в одном месте код, связанный с определённым состоянием.
	- Упрощает код контекста.

	Минусы:
	- Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

type State interface {
	GetName() string
	Freeze(ctx *StateContext)
	Heat(ctx *StateContext)
}

// ------------------------------------
type SolidState struct {
	name string
}

func NewSolidState() *SolidState {
	return &SolidState{
		name: "Solid",
	}
}

func (s *SolidState) GetName() string {
	return s.name
}

func (s *SolidState) Freeze(ctx *StateContext) {
	fmt.Println("Nothing happens...")
}

func (s *SolidState) Heat(ctx *StateContext) {
	ctx.SetState(NewLiquidState())
}

// ------------------------------------
type LiquidState struct {
	name string
}

func NewLiquidState() *LiquidState {
	return &LiquidState{
		name: "Liquid",
	}
}

func (s *LiquidState) GetName() string {
	return s.name
}

func (s *LiquidState) Freeze(ctx *StateContext) {
	ctx.SetState(NewSolidState())
}

func (s *LiquidState) Heat(ctx *StateContext) {
	ctx.SetState(NewGaseousState())
}

// ------------------------------------
type GaseousState struct {
	name string
}

func NewGaseousState() *GaseousState {
	return &GaseousState{
		name: "Gaseous",
	}
}

func (s *GaseousState) GetName() string {
	return s.name
}

func (s *GaseousState) Freeze(ctx *StateContext) {
	ctx.SetState(NewLiquidState())
}

func (s *GaseousState) Heat(ctx *StateContext) {
	fmt.Println("Nothing happens...")
}

// ------------------------------------
type StateContext struct {
	state State
}

func NewStateContext() *StateContext {
	return &StateContext{
		state: NewSolidState(),
	}
}

func (s *StateContext) Freeze() {
	fmt.Printf("Freezing %s substance...\n", s.state.GetName())
	s.state.Freeze(s)
}

func (s *StateContext) Heat() {
	fmt.Printf("Heating %s substance...\n", s.state.GetName())
	s.state.Heat(s)
}

func (s *StateContext) SetState(state State) {
	fmt.Println("Changing state to:", state.GetName())
	s.state = state
}

func (s *StateContext) GetState() State {
	return s.state
}


func main() {
	stateCtx := NewStateContext()
	stateCtx.Heat()
	stateCtx.Heat()
	stateCtx.Heat()
	stateCtx.Freeze()
	stateCtx.Freeze()
	stateCtx.Freeze()
}