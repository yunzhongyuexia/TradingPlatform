package tools

import (
	"fmt"
	"testing"
)

func TestReadABIFromFile(t *testing.T) {

}

func TestGenerateToken(t *testing.T) {
	uid := 12133123
	name := "ye"
	token, _ := GenerateToken(int64(uid), name)
	fmt.Println(token)
}
func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEyMTMzMTIzLCJuYW1lIjoieWUiLCJleHAiOjE3MjI1ODIzNTEsImlzcyI6Inl1bnpob25neXVleGlhIn0.uqtAz2ooEsulQjBoO_FAInh9-xgED3GoRL7GKm5oE2k"
	claims, _ := ParseToken(token)
	fmt.Println(claims)
}
