package pubsub

//Option for options setting
type Option func(*Options)

//Options for pubsub
type Options struct {
	PubTopics []string `json:"pub_topics"`
	SubTopics []string `json:"sub_topics"`
}
