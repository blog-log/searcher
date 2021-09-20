package config

type Algolia struct {
	AppId     string `envconfig:"APP_ID" required:"true"`
	ApiKey    string `envconfig:"API_KEY" required:"true"`
	IndexName string `envconfig:"SEARCH_INDEX" required:"true"`
}
