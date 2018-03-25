package artmatlang

type token struct {
	token  byte
	lexeme []byte
}

//
// Token types :
// byte('(') : (
// byte(')') : )
// byte('+') : +
// byte('-') : -
// byte('*') : *
// byte('/') : /
// byte('N') : number
//
