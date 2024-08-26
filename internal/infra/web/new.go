package web

type Iconnection struct {
	Host string
	Port string
}

type Person struct {
	Name string
	Age  int
}

func New(host, port string) *Iconnection {
	return &Iconnection{
		Host: host,
		Port: port,
	}
}
