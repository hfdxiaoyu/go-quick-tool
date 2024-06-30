package jwt

import "testing"

func TestNewHs256Token(t *testing.T) {
	util := NewJwtUtil()
	InitSingKey("vhsdfuivhidhvfevsdf")
	token, err := util.NewHs256Token(User{
		Id:       1,
		UserType: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiVXNlclR5cGUiOjEsImV4cCI6MTcxOTc4OTEyOX0.FkgFnN3nuv750Ov0k3p7mqZpHWI8oWf4z-Z2ZQu5p9Q
	t.Log(token)
}

func TestName(t *testing.T) {
	InitSingKey("vhsdfuivhidhvfevsdf")
	parse, err := NewJwtUtil().Hs256Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiVXNlclR5cGUiOjEsImV4cCI6MTcxOTc4OTEyOX0.FkgFnN3nuv750Ov0k3p7mqZpHWI8oWf4z-Z2ZQu5p9Q")
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
