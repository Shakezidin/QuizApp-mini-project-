package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Shakezidin/pkg/model"
	"github.com/sashabaranov/go-openai"
)

func (a *AdminService) GetQuizLanguagesSVC() []string {
	languages := []string{"Golang", "Python", "Java", "JavaScript"}

	return languages
}

func (a *AdminService) StartQuizSVC(language string) (*model.QuizResponse, error) {
	quizQuestion, err := a.generateQuestions(language)
	if err != nil {
		return nil, err
	}
	err = a.saveToRedis(quizQuestion.Question, quizQuestion.Answer)
	if err != nil {
		return nil, nil
	}

	return quizQuestion, nil
}

func (a *AdminService) CheckAnswerSVC(userAnswer model.UserAnswer) (bool, error) {
	fmt.Println(userAnswer.Question)
	correctAnswer, err := a.getFromRedis(userAnswer.Question)
	if err != nil {
		return false, err
	}

	return userAnswer.Answer == correctAnswer, nil
}

// generateQuestions asynchronously generates quiz questions using the OpenAI API
func (a *AdminService) generateQuestions(language string) (*model.QuizResponse, error) {
	client := openai.NewClient(a.cnfg.QUIZAUTH)

	// Define the prompt
	prompt := `Please provide a JSON stringfied array of 1 quiz response with properties - {question: question string, a: option a string, b: option b string, c: option c string, d: option d string, answer: option string}
	* Important instructions are mentioned below - 
	* Each question should be from these topics - %s.
	* Make sure you only respond with the JSON stringfied data with the format I mentioned above, because I will directly JSON-parse it and I dont want convertion errors.
	* Be interconnected and Do not send duplicate questions.
	* Do not send any text other than the quizz object.
	* please dont sent it as in array of json format,sent it as single json data`

	// Get topics from topicHelper() function
	prompt = fmt.Sprintf(prompt, language)

	// Create a completion request
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo0613,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleSystem, Content: prompt},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	// Parse the response
	var questions model.QuizResponse
	if len(resp.Choices) > 0 {
		content := resp.Choices[0].Message.Content
		err := json.Unmarshal([]byte(content), &questions)
		if err != nil {
			return nil, err
		}
		err = a.saveToRedis(questions.Question, questions.Answer)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return &questions, nil
	}

	return nil, nil
}

func (a *AdminService) saveToRedis(question, answer string) error {
	return a.RedisClient.Set(context.Background(), question, answer, 300*time.Second).Err()
}

func (a *AdminService) getFromRedis(question string) (string, error) {
	return a.RedisClient.Get(context.Background(), question).Result()
}
