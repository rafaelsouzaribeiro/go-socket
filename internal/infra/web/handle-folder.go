package web

import (
	"fmt"
	"io"
	"net"
	"os"
)

func (h *Iconnection) HandleFolderConnection(conn *net.TCPConn, OutputFolder string, channel chan string) {
	defer conn.Close()

	outFile, err := os.Create(fmt.Sprintf("%s/received.zip", OutputFolder))
	if err != nil {
		channel <- fmt.Sprintf("Error creating file: %v", err)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, conn)
	if err != nil {
		channel <- fmt.Sprintf("Error saving file: %v", err)
		return
	}

	channel <- "File received and saved successfully!"

}
