package bodies

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int   `json:"age,omitempty"`
	Address  string `json:"address"`
	Email     string `json:"email"`
	GroupID   int    `json:"group_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
