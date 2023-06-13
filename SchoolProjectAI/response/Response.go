package response

import (
	"context"
	"log"

	//"os"

	"github.com/joho/godotenv"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func AiResponse(prompt string) string {
	godotenv.Load()

	apiKey := "" // os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}

	ctx := context.Background()
	client := gogpt.NewClient(apiKey)

	resp, err := client.CreateCompletion(ctx, gogpt.CompletionRequest{
		Model:       "text-davinci-003",
		Temperature: 0.5,
		Prompt:      prompt,
		MaxTokens:   600,
		Stop:        []string{".", "You say:"},
		Echo:        false,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return resp.Choices[0].Text
}
