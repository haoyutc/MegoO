package main

import (
	"flag"
	"fmt"
	"github.com/sevlyar/go-daemon"
	"html"
	"log"
	"net/http"
	"os"
	"syscall"
	"time"
)

// 要终止守护进程，使用:
// kill `cat sample.pid`
func main() {
	//commonUse()
	signalUse()
}

func commonUse() {
	ctx := daemon.Context{
		PidFileName: "sample.pid",
		PidFilePerm: 0644,
		LogFileName: "sample.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Args:        []string{"[go-daemon sample]"},
		Umask:       027,
	}
	d, err := ctx.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer ctx.Release()

	log.Println("----------------------------------")
	log.Println("daemon start")
	serveHTTP()
}
func serveHTTP() {
	http.HandleFunc("/", httpHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request from %s: %s %q", r.RemoteAddr, r.Method, r.URL)
	fmt.Fprintf(w, "go-daemon: %q", html.EscapeString(r.URL.Path))
}

var (
	signal = flag.String("s", "", `Send signal to the daemon:
  quit — graceful shutdown
  stop — fast shutdown
  reload — reloading the configuration file`)
)

func signalUse() {
	flag.Parse()
	daemon.AddCommand(daemon.StringFlag(signal, "quit"), syscall.SIGQUIT, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "stop"), syscall.SIGTERM, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "reload"), syscall.SIGHUP, reloadHandler)

	ctx := daemon.Context{
		PidFileName: "sample.pid",
		PidFilePerm: 0644,
		LogFileName: "sample.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Args:        []string{"[go-daemon sample]"},
		Umask:       027,
	}
	if len(daemon.ActiveFlags()) > 0 {
		d, err := ctx.Search()
		if err != nil {
			log.Fatalf("Unable send signal to the daemon: %s", err.Error())

		}
		daemon.SendCommands(d)
		return
	}
	d, err := ctx.Reborn()
	if err != nil {
		log.Fatalln(err)
	}
	if d != nil {
		return
	}
	defer ctx.Release()

	log.Println("- - - - - - - - - - - - - - -")
	log.Println("daemon started")

	go worker()

	err = daemon.ServeSignals()
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}
	log.Println("daemon terminated")
}

var (
	stop = make(chan struct{})
	done = make(chan struct{})
)

func worker() {
LOOP:
	for {
		log.Println("this is for test")
		time.Sleep(time.Second) // this work to be done by worker
		select {
		case <-stop:
			break LOOP
		default:

		}
	}
	done <- struct{}{}
}
func termHandler(sig os.Signal) error {
	log.Println("Terminating...")
	stop <- struct {
	}{}
	if sig == syscall.SIGQUIT {
		<-done
	}
	return daemon.ErrStop
}
func reloadHandler(sig os.Signal) error {
	log.Println("server reloaded")
	return nil
}
