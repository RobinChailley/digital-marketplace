package conf

type Configuration struct {
	Host        					string      				`yaml:"host"`
	Port        					string      				`yaml:"port"`
	Credentials 					Credential  				`yaml:"credential"`
	DatabaseConfig 				DatabaseConfig 			`yaml:"database-config"`
	AdsService 						Service							`yaml:"ads-service"`
	AccountsService 			Service							`yaml:"accounts-service"`
}

type Credential struct {
	ApiKey 								string 							`yaml:"api_key"`
	JwtSecret  						string 							`yaml:"jwt_secret"`
}

type DatabaseConfig struct {
	User 									string 							`yaml:"user"`
	Database							string 							`yaml:"database"`
	Addr 									string 							`yaml:"addr"`
	Password							string 							`yaml:"password"`
}

type Service struct {
	URL    								string 							`yaml:"url"`
	ApiKey 								string 							`yaml:"api_key"`
}