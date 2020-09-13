package dgraph

type Resource struct {
	Uid  string `json:"uid,omitempty"`
	Type string `json:"dgraph.type,omitempty"`
	ID   string `json:"resource.id,omitempty"`
	Name string `json:"resource.name,omitempty"`
}
