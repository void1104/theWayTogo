package main

import (
	"bytes"
	"expvar"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// hello world, the web server
/**
包 expvar 可以创建（Int，Float 和 String 类型）变量，并将它们发布为公共变量。
它会在 HTTP URL /debug/vars 上以 JSON 格式公布。通常它被用于服务器操作计数。
*/
var helloRequests = expvar.NewInt("hello-requests")

// flags:
var webroot = flag.String("root", "/home/user", "web root directory")

// simple flag server
var booleanflag = flag.Bool("boolean", true, "another flag for testing")

// Simple counter server. POSTing to it will set the value.
type Counter struct {
	n int
}

// a channel
type Chan chan int

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(Logger))
	http.Handle("/go/hello", http.HandlerFunc(HelloServer))
	// The counter is published as a variable directly
	ctr := new(Counter)
	expvar.Publish("counter", ctr)
	http.Handle("/counter", ctr)

}

/**
日志服务
*/
func Logger(w http.ResponseWriter, req *http.Request) {
	log.Print(req.URL.String())
	w.WriteHeader(404)
	w.Write([]byte("oops"))
}

/**
hello服务
*/
func HelloServer(w http.ResponseWriter, req *http.Request) {
	helloRequests.Add(1)
	io.WriteString(w, "hello world!\n")
}

/**
VisitAll函数迭代所有的标签(flag),打印它们的名称，值和默认值
*/
func FlagServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Flags:\n")
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() != f.DefValue {
			fmt.Fprintf(w, "%s = %s [default = %s]\n", f.Name, f.Value.String(), f.DefValue)
		} else {
			fmt.Fprintf(w, "%s = %s\n", f.Name, f.Value.String())
		}
	})
}

// simple argument server
func ArgServer(w http.ResponseWriter, req *http.Request) {
	for _, s := range os.Args {
		fmt.Fprintf(w, s, " ")
	}
}

func ChanCreate() Chan {
	c := make(Chan)
	go func(c Chan) {
		for x := 0; ; x++ {
			c <- x
		}
	}(c)
	return c
}

// This makes Counter satisfy the expvar.Var interface, so we can export
// it directly
func (ctr *Counter) String() string {
	return fmt.Sprintf("%d", ctr.n)
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		ctr.n++
	case "POST":
		buf := new(bytes.Buffer)
		io.Copy(buf, req.Body)
		body := buf.String()
		if n, err := strconv.Atoi(body); err != nil {
			fmt.Fprintf(w, "bad POST: %v\nbody: [%v]\n", err, body)
		} else {
			ctr.n = n
			fmt.Fprintf(w, "counter reset\n")
		}
	}
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}
