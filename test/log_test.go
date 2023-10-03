package test

import (
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	f := log.New(os.Stdout, "[Go-Bat]", log.LstdFlags)
	//f.
	//f.SetPrefix("[Go-Bat]")
	f.Println(11212)
	f.Println("dfnjfndjk")
	f.Println("djfjdfjdf")

}
