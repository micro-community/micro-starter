package dgraph

type Role struct {
	Uid      string     `json:"uid,omitempty"`
	Type     string     `json:"dgraph.type,omitempty"`
	ID       string     `json:"role.id,omitempty"`
	Name     string     `json:"role.name,omitempty"`
	Resource []Resource `role:"resource,omitempty"`
}
