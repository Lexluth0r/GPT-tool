package api

import (
	"chat/app/common/errcode"
	"chat/app/common/response"
	"chat/config"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
)

type question struct {
	Question string `json:"question"`
	Key      string `json:"key"`
}

func Chat(c *gin.Context) {

	json := question{}

	c.BindJSON(&json)

	gpt := gogpt.NewClient(config.Conf.Chat.Token)

	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:            gogpt.GPT3TextDavinci003,
		MaxTokens:        500,
		Temperature:      0.9,
		TopP:             1,
		FrequencyPenalty: 0.8,
		PresencePenalty:  0.6,
		Prompt:           json.Question,
	}
	resp, err := gpt.CreateCompletion(ctx, req)

	fmt.Println(resp)
	if err != nil {
		fmt.Println(err)
		response.Error(c, errcode.ERROR)
		return
	}

	text := resp.Choices[0].Text

	response.Success(c, text)

	return
}
