package requests

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/valyala/fastjson"
)

type UserInfoStruct struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Avatar string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	PublicFlags uint32 `json:"public_flags"`
	Flags uint32 `json:"flags"`
	Locale string `json:"locale"`
	MfaEnabled bool `json:"mfa_enabled"`
	PremiumType uint8 `json:"premium_type"`
}

func GetUserInfo(token string, tokentype string) (UserInfoStruct, *ErrorStruct){
	var url = "https://discord.com/api"
	headers := req.Header{"Authorization": fmt.Sprintf("%s %s", tokentype, token)}

	r, err := req.Get(url + "/users/@me", headers)

	if err != nil {
		return UserInfoStruct{}, &ErrorStruct{Message: "An error occurred when attempt to call DiscordAPI.", StatusCode: 503, Error: err}
	}

	if r.Response().StatusCode != 200 {
		return UserInfoStruct{}, &ErrorStruct{Message: "Unexpected status from DiscordAPI was received.", StatusCode: 400, Error: err}
	}

	var p fastjson.Parser

	parsed, err := p.Parse(r.String())

	if err != nil {
		return UserInfoStruct{}, &ErrorStruct{Message: "Error in parsing the JSON files.", StatusCode: 500, Error: err}
	}

	return UserInfoStruct{
		Id:            string(parsed.GetStringBytes("id")),
		Username:      string(parsed.GetStringBytes("username")),
		Avatar:        string(parsed.GetStringBytes("avatar")),
		Discriminator: string(parsed.GetStringBytes("discriminator")),
		PublicFlags:   uint32(parsed.GetUint("public_flags")),
		Flags:         uint32(parsed.GetUint("flags")),
		Locale:        string(parsed.GetStringBytes("locale")),
		MfaEnabled:    parsed.GetBool("mfa_enabled"),
		PremiumType:   uint8(parsed.GetUint("premium_type")),
	}, nil

}