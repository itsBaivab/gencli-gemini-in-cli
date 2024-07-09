package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

var chatCmd = &cobra.Command{
	Use:   "search [your question]",
	Short: "Ask a question and get a response",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := getApiRespone(args)
		fmt.Println(res)
	},
}

func getApiRespone(args []string) string {

	userArgs := strings.Join(args[1:], " ")

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(userArgs+"in 100-120 words."))
	if err != nil {
		log.Fatal(err)
	}

	finalResponse := resp.Candidates[0].Content.Parts[0]

	return fmt.Sprint(finalResponse)
}