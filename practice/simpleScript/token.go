package simpleScript

type TokenType int

const (
	TOKEN_TYPE_Plus = iota   // +
	TOKEN_TYPE_Minus  // -
	TOKEN_TYPE_Star   // *
	TOKEN_TYPE_Slash  // /

	TOKEN_TYPE_GE     // >=
	TOKEN_TYPE_GT     // >
	TOKEN_TYPE_EQ     // ==
	TOKEN_TYPE_LE     // <=
	TOKEN_TYPE_LT     // <

	TOKEN_TYPE_SemiColon // ;
	TOKEN_TYPE_LeftParen // (
	TOKEN_TYPE_RightParen// )

	TOKEN_TYPE_Assignment// =

	TOKEN_TYPE_If
	TOKEN_TYPE_Else

	TOKEN_TYPE_Int

	TOKEN_TYPE_Identifier    //标识符

	TOKEN_TYPE_IntLiteral    //整型字面量
	TOKEN_TYPE_StringLiteral   //字符串字面量
)

type Token interface {
	SetType(tokenType TokenType)
	GetType() TokenType
	SetText(data []rune)
	GetText() []rune
}

type TokenReader interface {
	IsEnd() bool
	Read() Token
	UnRead()
	Peek() Token
	GetPosition() int
	SetPosition(pos int)
}

type myToken struct {
	tokenType TokenType
	text []rune
}

func (t *myToken) SetType(tokenType TokenType) {
	t.tokenType = tokenType
}

func (t *myToken) GetType() TokenType {
	return t.tokenType
}

func (t *myToken) SetText(data []rune) {
	t.text = data
}

func (t *myToken) GetText() []rune {
	return t.text
}

type SimpleTokenReader struct {
	tokens []Token
	pos int
}

func NewSimpleTokenReader(tokens []Token) *SimpleTokenReader {
	return &SimpleTokenReader{
		tokens: tokens,
	}
}

func (r *SimpleTokenReader) IsEnd() bool {
	return r.pos == len(r.tokens)
}

func (r *SimpleTokenReader) Read() Token {
	if r.pos >= len(r.tokens) {
		return nil
	}
	r.pos++
	return r.tokens[r.pos-1]
}

func (r *SimpleTokenReader) Peek() Token {
	if r.pos >= len(r.tokens) {
		return nil
	}
	return r.tokens[r.pos]
}

func (r *SimpleTokenReader) UnRead() {
	if r.pos > 0 {
		r.pos--
	}
}

func (r *SimpleTokenReader) GetPosition() int {
	return r.pos
}

func (r *SimpleTokenReader) SetPosition(pos int) {
	if pos < 0 || pos > len(r.tokens) {
		return
	}
	r.pos = pos
}