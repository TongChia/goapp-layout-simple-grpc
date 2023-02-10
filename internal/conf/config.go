package conf

type HTTP struct {
	Addr    string
	Timeout string
}

type GRPC struct {
	Addr    string
	Timeout string
}

type Server struct {
	HTTP *HTTP
	GRPC *GRPC
}

type Database struct {
	Driver string
	Source string
}

type Redis struct {
	Addr         string
	WriteTimeout string
	ReadTimeout  string
}

type Data struct {
	Database *Database
	Redis    *Redis
}

type Config struct {
	Server *Server
	Data   *Data
}
