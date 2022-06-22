package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port string
	root string
)

func init() {
	flag.StringVar(&port, "port", "8080", "-port 8080")
	flag.StringVar(&root, "root", "./", "-root ./")
}

func main() {
	flag.Parse()

	log.Println(port, root)

	http.Handle("/", http.FileServer(http.Dir(root)))
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%s", port), "./cert/server.crt", "./cert/server.key", nil)) //server.crt:服务端证书 包含服务端公钥信息 server.key：服务端私钥
}
