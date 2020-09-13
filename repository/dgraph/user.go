package dgraph

type User struct {
	Uid    string `json:"uid,omitempty"`
	Type   string `json:"dgraph.type,omitempty"`
	ID     string `json:"person.id,omitempty"`
	Name   string `json:"name,omitempty"`
	Age    int64  `json:"age,omitempty"`
	Gender string `json:"gender,omitempty"`
	Role   []Role `role:"gender,omitempty"`
}
