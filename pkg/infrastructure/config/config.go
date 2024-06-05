package config

type PortManager struct {
	RunnerPort     string `mapstructure:"PORTNO"`
	PostNrelSvcUrl string `mapstructure:"POSTNREL_SVC_URL"`
}

type DataBase struct {
	DBUser     string `mapstructure:"DBUSER"`
	DBName     string `mapstructure:"DBNAME"`
	DBPassword string `mapstructure:"DBPASSWORD"`
	DBHost     string `mapstructure:"DBHOST"`
	DBPort     string `mapstructure:"DBPORT"`
}

type Config struct {
	PortMngr PortManager
	DB       DataBase
}

func LoadConfig() (*Config, error) {
	var portmngr PortManager
	var db DataBase

	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&portmngr)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&db)
	if err != nil {
		return nil, err
	}

	config := Config{PortMngr: portmngr, DB: db}
	return &config, nil

}
