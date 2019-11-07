package config

type Conf struct {
	Targets	map[string]Target	`json:"targets"`
}

type Target struct {
	Host	string	`json:"host"`
	Port	uint		`json:"port"`
	User	string	`json:"user"`
	Passw	string	`json:"password"`
}
