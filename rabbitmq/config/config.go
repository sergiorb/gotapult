package config

type Config struct {
  Servers	map[string]RabbitMqConf	`json:servers`
}

type RabbitMqConf struct {
	Host	string	`json:"host"`
	Port	uint	`json:"port"`
	User	string	`json:"user"`
	Passw	string	`json:"password"`
}
