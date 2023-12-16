package main

import (
	"fmt"
  "net/http/httputil"
  
)

type simpleserver struct{
  addr string 
  proxy "httputil.ReverseProxy"  
}



func main() {
	fmt.Println("We want to Build Golang Load Balance")
}
