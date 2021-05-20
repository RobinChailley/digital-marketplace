package conf

type Configuration struct {
	Host        string      		`yaml:"host"`
	Port        string      		`yaml:"port"`
	Credentials Credential  		`yaml:"credential"`
	DatabaseConfig DatabaseConfig 	`yaml:"database-config"`
	AccountsApi AccountsApi 		`yaml:"accounts-api"`
}

type AccountsApi struct {
	URL 		string				`yaml:"url"`
}

type Credential struct {
	ApiKey 		string 				`yaml:"api_key"`
	JwtSecret  	string 				`yaml:"jwt_secret"`
}

type DatabaseConfig struct {
	User 		string 				`yaml:"user"`
	Database	string 				`yaml:"database"`
	Addr 		string 				`yaml:"addr"`
	Password	string 				`yaml:"password"`
}
