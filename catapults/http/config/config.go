package config

type Conf struct {
	Targets	map[string]Target	`json:"targets"`
}

type Target struct {
	Schema	string 	`json:"schema"`
	Host	string	`json:"host"`
	Port	uint	`json:"port"`
}
