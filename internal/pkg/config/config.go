include (
	"gopkg.in/yaml.v2"
)

type Database struct {
	Hostname string `yaml:"hostname"`
	Database string `yaml:"database"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type ServerData struct 
{
	Port string `yaml:"port"`
}

type Config struct {
	db Database
	server ServerData
}

func Configure() {

}



