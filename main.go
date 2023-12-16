package main

import (
	"fmt"
  "net/http/httputil"
  
)

type simpleserver struct{
  addr string 
  proxy "httputil.ReverseProxy"  
}

func newsimpleserver(addr string) *simpleserver{
  serverUrl, err:=url.Parse(addr)
  handleErr(err)

  return &simpleserver{
    addr: addr,
    proxy: httputil.NewSingleHostReverseProxy(serverUrl),
}
}

  func handleErr(err error){
  if err!=nil{
  fmt.Println("Error: %v/n ", err)
}    
  }   
func main() {
	fmt.Println("We want to Build Golang Load Balance")
}
