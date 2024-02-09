package lexer

type Lexer struct {
	Content []byte
	Ip      int
	Count   int
}

func (l *Lexer) Next() byte {
	for l.Ip < l.Count && !isToken(l.Content[l.Ip]) {
		l.Ip += 1
	}

	if l.Ip >= l.Count {
		return 0
	}

	l.Ip += 1
	return l.Content[l.Ip-1]
}

const TOKENS = "<>+-.,[]"

func isToken(symbol byte) bool {
	for _, token := range []byte(TOKENS) {
		if token == symbol {
			return true
		}
	}
	return false
}
