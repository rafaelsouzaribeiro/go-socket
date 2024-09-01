package web

type Iconnection struct {
	Host string
	Port string
}

func New(host, port string) *Iconnection {
	return &Iconnection{
		Host: host,
		Port: port,
	}
}
