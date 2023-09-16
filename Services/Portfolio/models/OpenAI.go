package models

type GPT3TurboMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GPT3TurboChoice struct {
	Index        int              `json:"index"`
	Message      GPT3TurboMessage `json:"message"`
	FinishReason string           `json:"finish_reason"`
}

type GPT3TurboUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type GPT3TurboResponse struct {
	ID      string            `json:"id"`
	Object  string            `json:"object"`
	Created int               `json:"created"`
	Choices []GPT3TurboChoice `json:"choices"`
	Usage   GPT3TurboUsage    `json:"usage"`
}

type GPT3TurboRequest struct {
	MaxTokens   int                 `json:"max_tokens"`
	Temperature float64             `json:"temperature"`
	Model       string              `json:"model"`
	TopP        float64             `json:"top_p"`
	N           int                 `json:"n"`
	PresencePen float64             `json:"presence_penalty"`
	FrequencyPe float64             `json:"frequency_penalty"`
	Messages    []*GPT3TurboMessage `json:"messages"`
}

type AdaEmbedRequest struct {
	Input []string `json:"input"`
	Model string   `json:"model"`
}

type AdaEmbedResponseData struct {
	Object    string    `json:"object"`
	Index     int       `json:"index"`
	Embedding []float64 `json:"embedding"`
}

type AdaEmbedResponseUsage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

type AdaEmbedResponse struct {
	Object string                  `json:"object"`
	Data   []*AdaEmbedResponseData `json:"data"`
	Usage  AdaEmbedResponseUsage   `json:"usage"`
}

func CreateGPT3TurboRequest() *GPT3TurboRequest {
	var new_request *GPT3TurboRequest = new(GPT3TurboRequest)
	new_request.MaxTokens = 100
	new_request.Temperature = 1
	new_request.Model = "gpt-3.5-turbo"
	new_request.TopP = 1
	new_request.N = 1
	new_request.PresencePen = 0
	new_request.FrequencyPe = 0
	new_request.Messages = make([]*GPT3TurboMessage, 0)

	return new_request
}

func CreateAdaEmbedRequest() *AdaEmbedRequest {
	var new_request *AdaEmbedRequest = new(AdaEmbedRequest)
	new_request.Model = "text-embedding-ada-002"
	new_request.Input = make([]string, 0)

	return new_request
}
