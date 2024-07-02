package jwt

import (
	"testing"
)

func TestNewHs256Token(t *testing.T) {
	util := NewJwtUtil()
	InitSingKey("vhsdfuivhidhvfevsdf")
	token, err := util.NewHs256Token(User{
		Id:   1,
		Data: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiVXNlclR5cGUiOjEsImV4cCI6MTcxOTc4OTEyOX0.FkgFnN3nuv750Ov0k3p7mqZpHWI8oWf4z-Z2ZQu5p9Q
	t.Log(token)
}

func TestHs256Parse(t *testing.T) {
	InitSingKey("vhsdfuivhidhvfevsdf")
	parse, err := NewJwtUtil().Hs256Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiRGF0YSI6MSwiZXhwIjoxNzE5ODg4NDkyfQ.VruQNqo0v8-BYBZYLI_9PIBI2NOUNYIjB_RwEPOqYSs")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(parse)
	calms, ok := parse.Claims.(*UserClaims)
	if !ok {
		t.Fail()
	}
	t.Log(calms)
}

func TestHs256RefreshToken(t *testing.T) {
	util := NewJwtUtil()
	InitSingKey("vhsdfuivhidhvfevsdf")
	token, err := util.Hs256RefreshToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiRGF0YSI6MSwiZXhwIjoxNzE5ODg3NDEzfQ.SrvbNiSGSOvcdwpYlozm-Hb_1_ETQ8zYci_JCa9r5wc")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}
