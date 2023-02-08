package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"runtime/pprof"

	"xmirror.cn/iast/goat/http/beegoserver"
	"xmirror.cn/iast/goat/http/chiserver"
	"xmirror.cn/iast/goat/http/echoserve"
	"xmirror.cn/iast/goat/http/ginserve"
	"xmirror.cn/iast/goat/http/irisserve"
	"xmirror.cn/iast/goat/targets/sqltargets"
	"xmirror.cn/iast/goat/targets/weakpass"
	"xmirror.cn/iast/goat/util"
)

const DefaultAddr = ":8080"

var (
	addr        = flag.String("addr", DefaultAddr, "listen on this `:port`")
	profCpuFile = flag.String("cpu", "", "record runtime CPU profiling records into given file")
	profMemFile = flag.String("mem", "", "enable runtime memory profiling records into given file")
)

func main() {
	// parse arguments
	flag.Parse()

	// CPU profiling enable
	if len(*profCpuFile) > 0 {
		f, err := os.Create(*profCpuFile)
		if err != nil {
			log.Fatal("error create CPU profile: ", err)
		}
		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("error starting CPU profiling: ", err)
		}
	}

	//graceful shutdown to clean up database file
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		log.Println("Shutting down")
		// clean up
		err := sqltargets.DBCloseAndRemove()
		if err != nil {
			log.Fatal(err)
		}
		// CPU profiling ends
		if len(*profCpuFile) > 0 {
			pprof.StopCPUProfile()
		}
		// memory profiling enable
		if len(*profMemFile) > 0 {
			f, err := os.Create(*profMemFile)
			if err != nil {
				log.Fatal("error create memory profile: ", err)
			}
			defer f.Close()

			runtime.GC()
			if err := pprof.WriteHeapProfile(f); err != nil {
				log.Fatal("error writing memory profile: ", err)
			}
		}

		os.Exit(0)
	}()

	// server
	AddFront()
	TargetService()

	log.Printf("Server startup at: %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func TargetService() {
	weakpass.WeakPass()
	ginRouter := ginserve.Setup()
	http.HandleFunc("/gin/", func(w http.ResponseWriter, r *http.Request) {
		ginRouter.ServeHTTP(w, r)
	})

	echoRouter := echoserve.Setup()
	http.HandleFunc("/echo/", func(w http.ResponseWriter, r *http.Request) {
		echoRouter.ServeHTTP(w, r)
	})

	irisRouter, err := irisserve.Setup()
	if err != nil {
		log.Fatalf("Error iris setup")
	}
	http.HandleFunc("/iris/", func(w http.ResponseWriter, r *http.Request) {
		irisRouter.ServeHTTP(w, r)
	})

	chiRouter := chiserver.Setup()
	http.HandleFunc("/chi/", func(w http.ResponseWriter, r *http.Request) {
		chiRouter.ServeHTTP(w, r)
	})

	beegoRouter := beegoserver.Setup()
	http.HandleFunc("/beego/", func(w http.ResponseWriter, r *http.Request) {
		beegoRouter.ServeHTTP(w, r)
	})
}

func AddFront() {
	exPath := util.StartDir()
	staticPath, _ := filepath.Abs(filepath.Join(exPath, "front/static/"))
	fs := http.FileServer(http.Dir(staticPath))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexPath, err := filepath.Abs(filepath.Join(exPath, "front/index.html"))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		t, err := template.ParseFiles(indexPath)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		t.Execute(w, nil)
	})
}
