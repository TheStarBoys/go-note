package simpleScript

func Tokenize(script string) TokenReader {
	ctx := NewContext([]rune(script))
	for ctx.index < len(ctx.script) {
		if isBlank(ctx.Peek()) {
			ctx.Read() // 消耗掉空格
			continue
		}
		state := STATE_Defualt + 1
		for ; !ctx.isSucc && state < STATE_End; state++ {
			ctx.machine.Transition(state)
		}
		if ctx.isSucc == false && state == STATE_End {
			panic("compile err")
		}
		ctx.isSucc = false
	}

	return NewSimpleTokenReader(ctx.tokens)
}