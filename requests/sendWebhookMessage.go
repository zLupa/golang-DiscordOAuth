package requests

import (
	"fmt"
	"github.com/imroc/req"
	"os"
	"strings"
	"time"
)

func SendNewUserMessage(user UserInfoStruct) (string, *ErrorStruct) {
	var UserProfileUrl string
	var isNitro string

	type EmbedAuthorStruct struct {
		Name string `json:"name"`
		Url string `json:"url"`
		IconUrl string `json:"icon_url"`
	}

	type EmbedThumbnailStruct struct {
		Url string `json:"url"`
	}

	type EmbedFieldStruct struct {
		Name string `json:"name"`
		Value string `json:"value"`
	}

	type EmbedFooterStruct struct {
		Text string `json:"text"`
		IconUrl string `json:"icon_url"`
	}

	type EmbedStruct struct {
		Title string `json:"title"`
		Author EmbedAuthorStruct `json:"author"`
		Description string `json:"description"`
		Thumbnail EmbedThumbnailStruct `json:"thumbnail"`
		Color int `json:"color"`
		Fields []EmbedFieldStruct `json:"fields"`
		Timestamp time.Time`json:"timestamp"`
		Footer EmbedFooterStruct `json:"footer"`
	}

	type webhookMessageStruct struct {
		Content string `json:"content"`
		Tts bool `json:"tts"`
		Embeds []EmbedStruct `json:"embeds"`
	}

	if strings.HasPrefix(user.Avatar, "a_") {
		UserProfileUrl = fmt.Sprintf("https://cdn.discordapp.com/avatars/%v/%v.gif", user.Id, user.Avatar)
	} else {
		UserProfileUrl = fmt.Sprintf("https://cdn.discordapp.com/avatars/%v/%v.png", user.Id, user.Avatar)
	}

	if user.PremiumType > 0 {
		isNitro = "Yes"
	} else {
		isNitro = "No"
	}

	fmt.Println(UserProfileUrl)

	var webhookMessage webhookMessageStruct = webhookMessageStruct{
		Content: "New user!",
		Tts: false,
		Embeds: []EmbedStruct{
			{
				Title: "Yay! new developer appeared",
				Author: EmbedAuthorStruct{
					Name: user.Username,
					IconUrl: UserProfileUrl,
					Url: "https://github.com/" + user.Username,
				},
				Description: "New developer used the golang OAuth!\nHere's some information about she/he.",
				Thumbnail: EmbedThumbnailStruct{Url: "https://i.imgur.com/MJIAkn5.png"},
				Color: 16744576,
				Fields: []EmbedFieldStruct{
					{
						Name: ":thinking: What's their ID?",
						Value: string(user.Id),
					},
					{
						Name: ":camera: What's their profile picture?",
						Value: fmt.Sprintf("Click [HERE](%v) to see", UserProfileUrl),
					},
					{
						Name: ":face_with_monocle: What's their discriminator?",
						Value: "#" + user.Discriminator,
					},
					{
						Name: ":rolling_eyes: Does she/he have a nitro subscription?",
						Value: isNitro,
					},
					{
						Name: ":flushed: Is she/he a incredible person?",
						Value: "YES!",
					},
				},
				Timestamp: time.Now(),
				Footer: EmbedFooterStruct{
					Text: "Made with :heart: using golang!",
					IconUrl: "https://codespacelab.com/wp-content/uploads/2019/06/Go-Logo_Blue.png",
				},
			},
		},
	}

	resp, err := req.Post(os.Getenv("DISCORD_WEBHOOK_URL"), req.BodyJSON(webhookMessage), req.QueryParam{"wait": true})

	if err != nil {
		return "", &ErrorStruct{Message: "Can't send a message through webhook!", StatusCode: 503, Error: err}
	}

	return resp.String(), nil
}