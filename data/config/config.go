package config

type Conf struct {
	Host		string	`json:"host"`
	Port		int64	`json:"port"`
	Database	string	`json:"database"`
}
