package worker

import (
	"bytes"
	"context"
	"crypto/tls"
	"github.com/kumato/kumato/internal/auth"
	"github.com/kumato/kumato/internal/logger"
	"github.com/kumato/kumato/internal/runtime/worker/container"
	"github.com/kumato/kumato/internal/types"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Node struct {
	types.UnimplementedWorkerServer
	controller     types.ControllerClient
	controllerAddr string
	docker         container.Client
	hostname       string
}

func (n *Node) StopTask(ctx context.Context, in *types.Task) (*types.Empty, error) {
	return &types.Empty{}, n.docker.Stop(in.GetContainerId())
}

func (n *Node) RunTask(ctx context.Context, in *types.Task) (*types.Task, error) {
	t, err := n.docker.Run(in, auth.InternalToken(), n.controllerAddr)
	if err == nil {
		go n.waitContainer(t)
	}
	return t, err
}

func (n *Node) GetImages(ctx context.Context, in *types.Empty) (*types.Images, error) {
	return n.docker.GetImages()
}

func (n *Node) GetStats(ctx context.Context, in *types.Empty) (*types.Stats, error) {
	return n.getOSStat()
}

func (n *Node) GetLoad(ctx context.Context, in *types.Empty) (*types.Stats, error) {
	return n.getOSLoad()
}

func (n *Node) ConnectDocker(hostname string) {
	logger.Warn("connect docker with hostname:", hostname)
	n.hostname = hostname
	n.docker = container.Connect(hostname)
}

func (n *Node) SetCPUMemory(cpu, ram, gpu int64) {
	logger.Warn("set nano cpu (" + strconv.FormatInt(cpu, 10) + ") and memory (" + strconv.FormatInt(ram, 10) + ")")
	n.docker.SetNanoCPUTotal(cpu)
	n.docker.SetMemoryTotal(ram)
	n.docker.SetGPUTotal(gpu)
}

func (n *Node) SetControllerClient(c types.ControllerClient, addr string) {
	logger.Warn("set controller client to runtime")
	n.controller = c
	n.controllerAddr = addr
}

func (n *Node) waitContainer(t *types.Task) {
	t = n.docker.Wait(t)
	n.uploadLogs(t)
	n.docker.Remove(t.GetContainerId())

	for {
		if _, err := n.controller.TaskDone(context.Background(), t); err != nil {
			logger.Fatal("cannot ask controller to mark container", t.GetNode()+":"+t.GetContainerId(), "done :", err.Error())
			time.Sleep(5 * time.Second)
			continue
		}
		logger.Info("container", t.GetNode()+":"+t.GetContainerId(), "exit")
		break
	}
}

func (n *Node) uploadLogs(t *types.Task) {
	logs, err := n.docker.Logs(t.GetContainerId())
	if err != nil {
		logger.Fatal("cannot read logs from container", t.GetContainerId(), ":", err.Error())
		return
	}
	defer logs.Close()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(logs); err != nil {
		logger.Fatal("fail to read buf from logs", err.Error())
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest("POST", "https://"+n.controllerAddr+"/api/internal/uploadLog", buf)
	if err != nil {
		logger.Fatal("cannot create post request:", err.Error())
		return
	}

	id := strconv.FormatUint(uint64(t.GetId()), 10) + ":" + t.GetFileUri() + ".log"

	req.Header.Set("LOG_FILE", id)
	req.Header.Set("FILE_URI", t.GetFileUri())
	req.Header.Set("INTERNAL_TOKEN", auth.InternalToken())

	resp, err := client.Do(req)
	if err != nil {
		logger.Fatal("error do a post request:", err.Error())
	}

	logger.Info("uploading log file:", id)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Fatal("fail to read body of response:", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		logger.Fatal(string(body))
	}
}
