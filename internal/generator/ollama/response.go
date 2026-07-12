package ollama

type response struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}
