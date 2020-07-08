package simpleScript

type Context struct {
	script []rune // 当前待解析的文本串
	index int // 当前待解析文本串下标
	isSucc bool // 当前是否解析成功
	machine *Machine
	token Token
	tokens []Token
}

func NewContext(script []rune) *Context {
	ctx :=  &Context{
		script: script,
		token: &myToken{},
	}
	ctx.machine = NewMachine(STATE_Defualt, map[int]Stater{
		STATE_Defualt : &DefaultState{},
		STATE_Int : &IntState{ctx: ctx},
		STATE_Assignment : &AssignmentState{ctx: ctx},

		STATE_Id : &IdentifierState{ctx: ctx},
	})

	return ctx
}

func (c *Context) AddToken() {
	c.tokens = append(c.tokens, c.token)
	c.token = &myToken{}
	c.isSucc = true
}

func (c *Context) IsEnd() bool {
	return len(c.script) == c.index
}

func (c *Context) Read() rune {
	if c.index >= len(c.script) {
		panic("index out of range")
	}
	c.index++
	return c.script[c.index-1]
}

func (c *Context) Peek() rune {
	if c.index >= len(c.script) {
		panic("index out of range")
	}
	return c.script[c.index]
}

func (c *Context) UnRead() {
	if c.index > 0 {
		c.index--
	}
}

func (c *Context) GetPosition() int {
	return c.index
}

func (c *Context) SetPosition(index int) {
	if index < 0 || index > len(c.script) {
		return
	}
	c.index = index
}