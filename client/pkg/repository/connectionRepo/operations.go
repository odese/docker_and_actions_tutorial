package connectionRepo

import (
	"bufio"
	"net"
	"strings"
)

// SendMessage sends a message to the server
func SendMessage(conn net.Conn, message string) (err error) {
	_, err = conn.Write([]byte(message + "\n"))
	return err
}

// ReadResponse reads the server response
func ReadResponse(conn net.Conn) (response string, err error) {
	serverReader := bufio.NewReader(conn)
	response, err = serverReader.ReadString('\n')
	response = strings.TrimSpace(response)
	return response, err
}
