package requests

import (
	"github.com/imroc/req"
	"github.com/joho/godotenv"
	"github.com/valyala/fastjson"
	"os"
)

type responseStruct struct {
	AccessToken string
	ExpiresIn int
	RefreshToken string
	Scope string
	TokenType string
}

type ErrorStruct struct {
	StatusCode int
	Message string
	Error error
}

func GetToken(code string) (responseStruct, *ErrorStruct) {
	err := godotenv.Load()

	if err != nil {
		return responseStruct{}, &ErrorStruct{Message: "Can't load environment variables.", StatusCode: 500, Error: err}
	}

	request := req.New()

	var url = "https://discord.com/api/v6"
	var clientId = os.Getenv("DISCORD_CLIENT_ID")
	var clientSecret = os.Getenv("DISCORD_CLIENT_SECRET")
	var redirectUri = os.Getenv("REDIRECT_URI")

	params := req.Param{
		"client_id": clientId,
		"client_secret": clientSecret,
		"grant_type": "authorization_code",
		"code": code,
		"redirect_uri": redirectUri,
	}

	r, err := request.Post(url + "/oauth2/token", params)

	if err != nil {
		return responseStruct{}, &ErrorStruct{Message: "An error occurred when attempt to call DiscordAPI.", StatusCode: 503, Error: err}
	}

	if r.Response().StatusCode != 200 {
		return responseStruct{}, &ErrorStruct{Message: "Unexpected status from DiscordAPI was received.", StatusCode: 503, Error: err}
	}

	var p fastjson.Parser
	parsed, err := p.Parse(r.String())

	if err != nil {
		return responseStruct{}, &ErrorStruct{Message: "Error in parsing the JSON.", StatusCode: 500, Error: err}
	}

	ResponseObject := responseStruct{
		AccessToken:  string(parsed.GetStringBytes("access_token")),
		ExpiresIn:    parsed.GetInt("expires_in"),
		RefreshToken: string(parsed.GetStringBytes("refresh_token")),
		Scope:        string(parsed.GetStringBytes("scope")),
		TokenType:    string(parsed.GetStringBytes("token_type")),
	}

	return ResponseObject, nil
}
