package parser

import (
	"kdo.com/go/monkey/lexer"
	"kdo.com/go/monkey/token"
)


type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peerToken token.Token
}
