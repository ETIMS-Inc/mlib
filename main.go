package main

import (
	"flag"
	"fmt"
	"io/fs"
	"mlib/pkg/log"
	"net/http"
	"os"
	"time"
)

var version string

func main() {
	addr := flag.String("addr", ":8080", "the address on which the server is started on")
	flag.Parse()
	log.Info("Welcome to MLib v%v", version)

	http.HandleFunc("/workspace", func(writer http.ResponseWriter, request *http.Request) {
		dirFs := os.DirFS("E:\\books\\programming")
		var tree []string
		_ = fs.WalkDir(dirFs, ".", func(path string, d fs.DirEntry, err error) error {
			if path == "." {
				return nil
			}
			tree = append(tree, path)
			return nil
		})
		_, _ = writer.Write([]byte(fmt.Sprint(tree)))
	})

	s := &http.Server{Addr: *addr, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal("could not start http server: %v", err)
	}
}
