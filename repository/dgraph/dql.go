package dgraph

type Count struct {
	Count int `json:"count"`
}

type UID struct {
	UID string `json:"uid"`
}

type Root struct {
	Count []Count `json:"counts"`
	UID   []UID   `json:"uids"`
}
