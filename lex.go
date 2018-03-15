package artmatlang

func lex(src []byte) []token {
	lexed := make([]token, len(src))
	tokenCount := 0
	buffer := 0

	for pos, val := range src {
		if buffer == 0 {
			switch val {
			case byte('('):
				lexed[tokenCount] = token{token: byte('(')}
				tokenCount += 1
			case byte(')'):
				lexed[tokenCount] = token{token: byte(')')}
				tokenCount += 1
			case byte('+'):
				lexed[tokenCount] = token{token: byte('+')}
				tokenCount += 1
			case byte('-'):
				lexed[tokenCount] = token{token: byte('-')}
				tokenCount += 1
			case byte('*'):
				lexed[tokenCount] = token{token: byte('*')}
				tokenCount += 1
			case byte('/'):
				lexed[tokenCount] = token{token: byte('/')}
				tokenCount += 1
			case byte('^'):
				lexed[tokenCount] = token{token: byte('^')}
				tokenCount += 1
			case byte('x'):
				lexed[tokenCount] = token{token: byte('x')}
				tokenCount += 1
			case byte(' '):
			default:

				lexed[tokenCount] = token{token: byte('\x00')}
				lexed[tokenCount].lexeme, buffer = wordTilBreak(src[pos:])
				tokenCount += 1

			}
		} else {
			buffer -= 1
		}
	}
	return lexed[:tokenCount]

}

var special = [9]byte{byte('('), byte(')'), byte('+'), byte('-'), byte('*'), byte('/'), byte('^'), byte('x'), byte(' ')}

func isSpecial(char byte) (bool, int) {
	for index, val := range special {
		if char == val {
			return true, index
		}
	}
	return false, 8
}

// wordTilBreak finds the end of the word, assuming that
func wordTilBreak(input []byte) ([]byte, int) {
	len := 0
	for _, val := range input {
		//		fmt.Println(int(val))
		specialChar, _ := isSpecial(val)
		if specialChar {
			return input[:len], len
		} else {
			len += 1
		}
	}
	return input, len
}
