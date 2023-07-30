package config

import (
	"os"
)

func OpenAIApiKey() string {
	return os.Getenv("OPENAI_API_KEY")
}
