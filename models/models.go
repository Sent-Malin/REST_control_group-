package models

type HumanDataCreate struct {
	Name        string `json:"Name"`
	Surname     string `json:"Surname"`
	YearOfBirth string `json:"Year of birth"`
	GroupNumber int    `json:"Group number"`
}

type GroupDataCreate struct {
	Name     string `json:"Name"`
	ParentId int    `json:"ParentId"`
}

type HumanDataUpdate struct {
	Id          int    `json:"Id"`
	Name        string `json:"Name"`
	Surname     string `json:"Surname"`
	YearOfBirth string `json:"Year of birth"`
	GroupNumber int    `json:"Group number"`
}

type GroupDataUpdate struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name"`
	ParentId int    `json:"ParentId"`
}

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

func NewConfig() *Config {
	return &Config{
		Host:     "",
		Port:     "",
		User:     "",
		Password: "",
		Dbname:   "",
	}
}
