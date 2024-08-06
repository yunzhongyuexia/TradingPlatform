package tools

import (
	"fmt"
	"server/setETH"
	"testing"
)

func TestETH(t *testing.T) {
	setETH.ConnectETH()
	setETH.QueryTransaction("0xddd67d8f678604be4efa00ffbdabd7f43ab4f71495ff726feb42f2d741411b99")
	defer setETH.Client.Close()
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

func TestGenerateWallet(t *testing.T) {
	setETH.Balance()
}
