package govern

type Task struct {
	Module string            `json:"module"`
	Args   map[string]string `json:"args"`
	Items  []string          `json:"items,omitempty"`
}
