package controller

import (
	"github.com/kumato/kumato/internal/auth"
	"github.com/kumato/kumato/internal/db"
	"github.com/kumato/kumato/internal/gzip"
	"github.com/kumato/kumato/internal/handlers"
	"github.com/kumato/kumato/internal/logger"
	"github.com/kumato/kumato/internal/runtime/controller"
	"github.com/kumato/kumato/internal/types"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

var (
	Cmd = &cobra.Command{
		Use:   "master",
		Short: "Run in master mode",
		Run:   exec,
	}

	addr,
	cert,
	key,
	www,
	data string
	verbose bool
)

func init() {
	Cmd.Flags().StringVarP(&addr, "listen", "l", "", "address this master listens on")
	Cmd.Flags().StringVarP(&cert, "cert", "c", "", "TLS cert file")
	Cmd.Flags().StringVarP(&key, "key", "k", "", "TLS key file")
	Cmd.Flags().StringVar(&www, "www", "", "path to dashboard dist")
	Cmd.Flags().StringVar(&data, "data", "", "path to store data")
	Cmd.MarkFlagRequired("listen")
	Cmd.MarkFlagRequired("cert")
	Cmd.MarkFlagRequired("key")

	if os.Getenv("KUMATO_DEV_MODE") == "" {
		Cmd.MarkFlagRequired("www")
	}

	Cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose log")
}

func exec(cmd *cobra.Command, args []string) {
	db.Connect(data, verbose)
	if verbose {
		logger.EnableDebug()
	}

	var webHandlerFunc func(w http.ResponseWriter, r *http.Request)

	if m := os.Getenv("KUMATO_DEV_MODE"); m == "1" || m == "on" {
		webHandlerFunc = gzip.Gzip(reverseProxy)
	} else {
		webHandlerFunc = gzip.Gzip(spa)
	}

	rpc := grpc.NewServer()
	mux := http.NewServeMux()

	types.RegisterControllerServer(rpc, &controller.Controller{})
	go controller.ReadConfig(data)

	handlers.RegisterHandlers("/api", mux, data)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(
			r.Header.Get("Content-Type"), "application/grpc") {
			if r.Header.Get("INTERNAL_TOKEN") != auth.InternalToken() {
				w.WriteHeader(http.StatusForbidden)
				logger.Fatal("reject forbidden request from", r.RemoteAddr)
				return
			}
			rpc.ServeHTTP(w, r)
			return
		}
		webHandlerFunc(w, r)
	})

	logger.Warn("serve gRPC and web service on:", addr)
	logger.Fatal(http.ListenAndServeTLS(addr, cert, key, mux))
}

func reverseProxy(w http.ResponseWriter, r *http.Request) {
	// redirect www to non-www
	if strings.HasPrefix(r.Host, "www") {
		target := "https://" + r.Host[4:] + r.URL.Path
		if len(r.URL.RawQuery) > 0 {
			target += "?" + r.URL.RawQuery
		}
		http.Redirect(w, r, target, http.StatusFound)
		return
	}

	director := func(req *http.Request) {
		req.Header = r.Header
		req.URL = r.URL
		req.URL.Scheme = "http"
		req.URL.Host = "10.0.0.109:8080"
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(w, r)
}

func spa(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat(www + "/" + r.URL.Path); os.IsNotExist(err) {
		http.ServeFile(w, r, www+"/index.html")
		return
	}
	http.ServeFile(w, r, www+"/"+r.URL.Path)
}
