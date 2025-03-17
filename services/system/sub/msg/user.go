package msg

type UserMsg struct {
	ID    int    `json:"user_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
