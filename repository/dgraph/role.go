package dgraph

type Role struct {
	Uid       string     `json:"uid,omitempty"`
	Type      string     `json:"dgraph.type,omitempty"`
	ID        string     `json:"role.id,omitempty"`
	Name      string     `json:"role.name,omitempty"`
	Resources []Resource `role:"resource,omitempty"`
}
