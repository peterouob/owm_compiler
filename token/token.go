package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" //遇到未知詞法或字符
	EOF = "EOF" //文件結尾

	//標示符＋字面量
	IDENT = "IDENT" //add,foobar,x,y...
	INT = "INT" //123456

	//運算符
	ASSIGN = "="
	PLUS = "+"

	//分隔符
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//關鍵字
	FUNCTION = "FUNCTION"
	LET = "LET"
)