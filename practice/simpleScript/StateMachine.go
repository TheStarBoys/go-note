package simpleScript

import (
	"fmt"
)

// 状态机
type Machine struct {
	currentState int // 当前状态
	defaultState int // 默认状态
	states map[int]Stater // 状态表
}

func NewMachine(defaultState int, states map[int]Stater) *Machine {
	m :=  &Machine{
		currentState: defaultState,
		defaultState: defaultState,
		states: states,
	}
	m.states[defaultState].Enter()

	return m
}

// Transition 状态迁移
func (m *Machine) Transition(nextState int) error {
	err := m.transition(nextState)
	if err != nil {
		fmt.Printf("machine transition err: %v\n", err)
	}

	return err
}

func (m *Machine) transition(nextState int) error {
	// 判断当前状态和切换的下一个状态是否一致，相同则不切换
	if m.currentState == nextState {
		return fmt.Errorf("machine: transition: currentState is equal to nextState:%v", nextState)
	}
	// 判断状态是否存在
	nextStateObj, ok := m.states[nextState]
	if !ok {
		return fmt.Errorf("machine: transition: state(%v) not define", nextState)
	}

	// 退出当前状态
	currentStateObj, _ := m.states[m.currentState]
	currentStateObj.Exit()

	// 设置当前状态为下一状态
	m.currentState = nextState

	// 进入下一状态
	nextStateObj.Enter()
	return nil
}

// Stater 状态机接口
type Stater interface {
	Enter() // 进入状态
	Exit() // 退出状态
}

const (
	// 有限状态机状态定义

	STATE_Defualt = iota

	// 整型关键字
	STATE_Int

	// 赋值
	STATE_Assignment

	// 条件语句关键字
	//STATE_If
	//STATE_Else

	// 逻辑运算符
	//STATE_GE // >=
	//STATE_GT // >

	// 算数运算符
	//STATE_Plus
	//STATE_Minus
	//STATE_Star
	//STATE_Slash
	//
	//
	//STATE_SemiColon
	//STATE_LeftParen
	//STATE_RightParen

	// 字面量
	STATE_IntLiteral

	// 标识符
	STATE_Id

	// 状态终止, 相当于所有状态总数目
	STATE_End
)