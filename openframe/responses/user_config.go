package responses

type UserConfig struct {
	Config struct {
		PubSubUrl string `json:"pubsub_url"`
	}
}
