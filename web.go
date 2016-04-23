package main

import (
    "net"
    "github.com/go-martini/martini"
)

func main() {

  m := martini.Classic()

  interfaces, err := net.Interfaces()
  if err != nil {
      panic("Unable to get interfaces.")
  }

  m.Get("/", func() string {
    return "v1: " + interfaces[1].HardwareAddr.String()
  })

  m.Run()

}
