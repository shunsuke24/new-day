package usecase

import "log"

func (u *usecase) Recommend() (string, error) {
	prompt := "こんにちは！"

	output, err := u.openai.OpenAIChat.Completion(prompt, "gpt-4-1106-preview", 4096, 0.0)
	if err != nil {
		// エラーハンドリングをここで行います
		log.Fatal(err)
	}
	log.Println(output)
	return *output, nil
}
