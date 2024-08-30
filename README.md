To send a string of data<br />

client:<br />

 ```
factories := factory.NewClient(Connection.Person{}, nil, "Message")
buffer, err := factories.GetClient()

if err != nil {
    panic(err)
}

_, err = conn.Write(buffer)
if err != nil {
    panic(err)
}

 ```

 <br />
 Server:

 ```
fmt.Printf("Server started at %s:%s \n", connect.Host, connect.Port)
factories := factory.NewServer(factory.String)

channel := make(chan string)
factories.ChannelString = channel

go func() {
    for p := range channel {
        fmt.Printf("Message received: %s \n", p)
    }
}()

for {
    conn, err := listener.AcceptTCP()
    if err != nil {
        log.Fatal(err)
    }

    go factories.GetServer(connect, conn)
}


  ```

<br />
To send a slice of data<br />

client:<br />

 ```

	people := []Connection.Person{
		{Name: "Rafael", Age: 38},
		{Name: "Maria", Age: 30},
		{Name: "João", Age: 25},
	}
	factories := factory.NewClient(Connection.Person{}, people, "")
	buffer, err := factories.GetClient()

	if err != nil {
		panic(err)
	}

	_, err = conn.Write(buffer)
	if err != nil {
		panic(err)
	}

 ```

 <br />
 Server:

 ```
fmt.Printf("Server started at %s:%s \n", connect.Host, connect.Port)
	factories := factory.NewServer(factory.Slice)

	channel := make(chan Connection.Person)
	factories.ChannelPerson = channel

	go func() {
		for p := range channel {
			fmt.Printf("Message received. Name: %s, Age: %d\n", p.Name, p.Age)
		}
	}()

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}

		go factories.GetServer(connect, conn)
	}


  ```
<br />
To send a struct of data<br />

client:<br />

 ```

	factories := factory.NewClient(Connection.Person{Name: "Rafael", Age: 38}, nil, "")
	buffer, err := factories.GetClient()

	if err != nil {
		panic(err)
	}

	_, err = conn.Write(buffer)
	if err != nil {
		panic(err)
	}

 ```

 <br />
 Server:

 ```
fmt.Printf("Server started at %s:%s \n", connect.Host, connect.Port)
	factories := factory.NewServer(factory.Struct)

	channel := make(chan Connection.Person)
	factories.ChannelPerson = channel

	go func() {
		for p := range channel {
			fmt.Printf("Message received. Name: %s, Age: %d\n", p.Name, p.Age)
		}
	}()

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}

		go factories.GetServer(connect, conn)
	}

  ```