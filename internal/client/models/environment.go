package models

type Environment struct {
	Environment string `json:"environment"`
}

type EnvironmentGetRequest struct {
	Key string `json:"key"`
}

type EnvironmentDeleteRequest struct {
	Key string `json:"key"`
}

type EnvironmentUpdateRequest struct {
	Key      string        `json:"key"`
	Contents string        `json:"contents"`
	Servers  []interface{} `json:"servers"`
}
