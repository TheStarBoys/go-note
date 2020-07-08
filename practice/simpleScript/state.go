package simpleScript

import "log"

// Iter 迭代器，匹配一个状态的所有可能
func Iter(ctx *Context, tokenType TokenType, rules []rune) {
	beforeIndx := ctx.index
	data := []rune{}
	for _, r := range rules {
		if ctx.Peek() != r {
			ctx.machine.Transition(STATE_Defualt)
			ctx.SetPosition(beforeIndx)
			return
		}
		data = append(data, ctx.Read())
	}
	if len(rules) == 0 { // 如果没有对应规则，视为标识符匹配
		if !isAlpha(ctx.Peek()) { // 以字母开头的才认为是一个标识符
			ctx.machine.Transition(STATE_Defualt)
			ctx.SetPosition(beforeIndx)
			return
		}
		for !ctx.IsEnd() && !isBlank(ctx.Peek()) {
			r := ctx.Read()
			if r == '=' {
				ctx.UnRead()
				break
			}
			if !isAlpha(r) && !isDigit(r) {
				ctx.machine.Transition(STATE_Defualt)
				ctx.SetPosition(beforeIndx)
				return
			}
			data = append(data, r)
		}
	}
	ctx.token.SetType(tokenType)
	ctx.token.SetText(data)
	ctx.AddToken()
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isAlpha(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

func isBlank(r rune) bool {
	return r == '\n' || r == '\t' || r == ' '
}


type DefaultState struct {}

func (state *DefaultState) Enter() {}

func (state *DefaultState) Exit() {}

// IntState 整型状态
type IntState struct {
	ctx *Context
}

func (i *IntState) Enter() {
	log.Println("Enter IntState")
	Iter(i.ctx, TOKEN_TYPE_Int, []rune{
		0 : 'i',
		1 : 'n',
		2 : 't',
		3 : ' ',
	})
}

func (state *IntState) Exit() {}

// IdentifierState 标识符状态
type IdentifierState struct {
	ctx *Context
}

func (state *IdentifierState) Enter() {
	log.Println("Enter IdentifierState")
	Iter(state.ctx, TOKEN_TYPE_Identifier, nil)
}

func (state *IdentifierState) Exit() {}

// AssignmentState 赋值状态
type AssignmentState struct {
	ctx *Context
}

func (state *AssignmentState) Enter() {
	log.Println("Enter AssignmentState")
	Iter(state.ctx, TOKEN_TYPE_Assignment, []rune{
		0 : '=',
	})
}

func (state *AssignmentState) Exit() {}

// IntLiteralState 赋值状态
type IntLiteralState struct {
	ctx *Context
}

func (state *IntLiteralState) Enter() {
	log.Println("Enter IntLiteralState")
	Iter(state.ctx, TOKEN_TYPE_IntLiteral, []rune{
		0 : '=',
	})
}

func (state *IntLiteralState) Exit() {}
