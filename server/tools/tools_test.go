package tools

import (
	"fmt"
	"testing"
)

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEyMTMzMTIzLCJuYW1lIjoieWUiLCJleHAiOjE3MjI1ODIzNTEsImlzcyI6Inl1bnpob25neXVleGlhIn0.uqtAz2ooEsulQjBoO_FAInh9-xgED3GoRL7GKm5oE2k"
	claims, _ := ParseToken(token)
	fmt.Println(claims)
}
