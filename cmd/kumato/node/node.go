package node

import (
	"context"
	"crypto/tls"
	"github.com/kumato/kumato/internal/auth"
	"github.com/kumato/kumato/internal/logger"
	"github.com/kumato/kumato/internal/runtime/worker"
	"github.com/kumato/kumato/internal/types"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net/http"
	"os"
	"time"
)

var (
	Cmd = &cobra.Command{
		Use:   "join",
		Short: "Join cluster as a worker",
		Run:   exec,
	}

	lis,
	addr,
	cert,
	key string
	nanocpu,
	ram,
	gpu int64
	verbose bool
)

func init() {
	Cmd.Flags().StringVarP(&addr, "address", "a", "", "master address")
	Cmd.Flags().StringVarP(&lis, "listen", "l", "", "address this worker listens on")
	Cmd.Flags().Int64Var(&nanocpu, "cpu", 0, "cpu capacity in nano cpus")
	Cmd.Flags().Int64Var(&ram, "ram", 0, "ram capacity in bytes")
	Cmd.Flags().Int64Var(&gpu, "gpu", 0, "ram capacity in bytes")
	Cmd.Flags().StringVarP(&cert, "cert", "c", "", "TLS cert file")
	Cmd.Flags().StringVarP(&key, "key", "k", "", "TLS key file")
	Cmd.MarkFlagRequired("token")
	Cmd.MarkFlagRequired("listen")
	Cmd.MarkFlagRequired("cpu")
	Cmd.MarkFlagRequired("ram")
	Cmd.MarkFlagRequired("cert")
	Cmd.MarkFlagRequired("key")

	Cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose log")
}

func exec(cmd *cobra.Command, args []string) {
	if verbose {
		logger.EnableDebug()
	}

	hn, err := os.Hostname()
	if err != nil {
		logger.Fatal(err)
		os.Exit(-1)
	}

	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})),
		grpc.WithPerRPCCredentials(auth.JWTClaims{}))

	if err != nil {
		logger.Fatal(err)
		os.Exit(-1)
	}

	defer conn.Close()

	client := types.NewControllerClient(conn)

	go register(client, hn)
	serve(client, hn)
}

func serve(client types.ControllerClient, hostname string) {
	rpc := grpc.NewServer()
	mux := http.NewServeMux()

	w := worker.Node{}
	w.ConnectDocker(hostname)
	w.SetCPUMemory(nanocpu, ram, gpu)
	w.SetControllerClient(client, addr)
	types.RegisterWorkerServer(rpc, &w)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("INTERNAL_TOKEN") != auth.InternalToken() {
			w.WriteHeader(http.StatusForbidden)
			logger.Fatal("reject forbidden request from", r.RemoteAddr)
			return
		}
		rpc.ServeHTTP(w, r)
	})

	logger.Warn("serve gRPC on:", lis)
	logger.Fatal(http.ListenAndServeTLS(lis, cert, key, mux))
}

func register(client types.ControllerClient, hostname string) {
	time.Sleep(5 * time.Second)

	for {
		_, err := client.Register(context.Background(), &types.Node{
			Id:      hostname,
			Address: lis,
		})

		if err == nil {
			logger.Info("registration ok")
			return
		}

		logger.Fatal("failed to register:", err.Error())
		time.Sleep(5 * time.Second)
	}
}
