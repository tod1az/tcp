package test

import (
	"fmt"
	"os"
	"tcp/client"
	"tcp/server"
	"testing"
	"time"
)

var port = "6969"
var host = "localhost"
var userName1 = "Tomas"
var userName2 = "Gonzalo"

func TestServerRun(t *testing.T) {

	s := server.CreateServer(host, port)
	listener, err := s.CreateListener()
	if err != nil {
		t.Error("Error while creating the listener")
	}

	originalStdout := os.Stdout
	r, w := PipeStdout()

	go s.ListenForMessages(listener)

	time.Sleep(100 * time.Millisecond)

	outPut := GetOutput(r, w)
	os.Stdout = originalStdout

	s.StopServer()

	expectedOutput := fmt.Sprintf("Listening on port: %v\n", s.Port)
	if outPut != expectedOutput {
		t.Errorf("expected: %v\nbut got: %v", expectedOutput, outPut)
	}
}

func TestUserConnection(t *testing.T) {
	thisPort := "3001"
	s := server.CreateServer(host, thisPort)
	listener, err := s.CreateListener()
	if err != nil {
		t.Error("Error while creating the listener")
	}

	go s.ListenForMessages(listener)

	time.Sleep(1000 * time.Millisecond)
	c, err := client.CreateClient(userName1, host, thisPort)
	if err != nil {
		t.Error("Error while creating a new client")
	}

	originalStdout := os.Stdout
	r, w := PipeStdout()

	go c.SendName()
	time.Sleep(1000 * time.Millisecond)

	outPut := GetOutput(r, w)
	os.Stdout = originalStdout
	s.StopServer()

	expected := fmt.Sprintf("Addres: %v Name: %v connected\n", c.Conn.LocalAddr(), c.Name)
	if outPut != expected {
		t.Errorf("expected: %v\nbut got: %v", expected, outPut)
	}
}
