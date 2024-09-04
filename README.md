To customize the struct, go to ```pkg/global/index.go``` and modify the Custom struct.
You also need to fix the println statements in the channel.
<br/>

To send data in string<br />

client:<br />

```go
factories := factory.NewClient(factory.FactoryClient{
			TypeString: "Message",
		})
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

```go
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
To send data in slice<br />

client:<br />

 ```go

	people := []global.Custom{
		{Name: "Rafael", Age: 38},
		{Name: "Maria", Age: 30},
		{Name: "João", Age: 25},
	}
	factories := factory.NewClient(factory.FactoryClient{
			TypeSlice: people,
		})
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

 ```go
fmt.Printf("Server started at %s:%s \n", connect.Host, connect.Port)
factories := factory.NewServer(factory.Slice)

channel := make(chan global.Custom)
factories.ChannelCustom = channel

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
To send data in struct<br />

client:<br />
```go

	factories := factory.NewClient(factory.FactoryClient{
			TypeStruct: global.Custom{Name: "Rafael", Age: 38},
		})
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

```go
fmt.Printf("Server started at %s:%s \n", connect.Host, connect.Port)
factories := factory.NewServer(factory.Struct)

channel := make(chan global.Custom)
factories.ChannelCustom = channel

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
To send data in int32<br />

client:<br />
```go
	valueint := 5
	factories := factory.NewClient(factory.FactoryClient{
			TypeInt: &valueint,
		})
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

```go
fmt.Printf("Server started at %s:%s \n", connect.Host, connect.Port)
factories := factory.NewServer(factory.Int32)
channel := make(chan int32)
factories.ChannelInt = channel

	go func() {
		for p := range channel {
			fmt.Printf("Message received: %d\n", p)
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