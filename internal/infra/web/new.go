package web

type Iconnection struct {
	Host string
	Port string
}

type Send struct {
	Message string
}

func New(host, port string) *Iconnection {
	return &Iconnection{
		Host: host,
		Port: port,
	}
}
