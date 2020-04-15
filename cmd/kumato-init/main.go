package main

import (
	"compress/flate"
	"crypto/tls"
	"github.com/kumato/kumato/internal/runtime/worker/container"
	"github.com/mholt/archiver/v3"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	buildDir  = "/BUILD"
	resultDir = "/RESULT"
)

var (
	internalToken     = ""
	controllerAddress = ""
	fileID            = ""
	taskID            = ""
)

func init() {
	internalToken = os.Getenv(container.EnvKeyInternalToken)
	controllerAddress = os.Getenv(container.EnvKeyControllerAddress)
	fileID = os.Getenv(container.EnvKeyFileID)
	taskID = os.Getenv(container.EnvKeyTaskID)
	if err := os.Setenv("AUTORUN_BUILD", "/BUILD"); err != nil {
		panic(err)
	}
	if err := os.Setenv("AUTORUN_RESULT", "/RESULT"); err != nil {
		panic(err)
	}
	if internalToken == "" || controllerAddress == "" || fileID == "" || taskID == "" {
		panic("environment variables are not full filled")
	}

	if err := os.Mkdir(buildDir, 0755); err != nil {
		panic(err)
	}
	if err := os.Mkdir(resultDir, 0755); err != nil {
		panic(err)
	}
}

func main() {
	getFile()

	if err := archiver.Unarchive("/"+fileID, buildDir); err != nil {
		panic(err)
	}

	var sh string

	_ = filepath.Walk(buildDir, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() && f.Name() == "AUTORUN.sh" {
			sh = path
			return io.EOF
		}
		return nil
	})

	if sh == "" {
		panic("AUTORUN.sh is not found")
	}

	cmd := exec.Command("bash", sh)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
	}

	if err := (&archiver.Zip{
		CompressionLevel:       flate.BestCompression,
		MkdirAll:               true,
		SelectiveCompression:   true,
		ContinueOnError:        false,
		OverwriteExisting:      false,
		ImplicitTopLevelFolder: false,
	}).Archive([]string{resultDir}, "/result-"+taskID+".zip"); err != nil {
		panic(err)
	}

	uploadFile()

}

func getFile() {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest("GET", "https://"+controllerAddress+"/api/internal/getFile/"+fileID, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set(container.EnvKeyInternalToken, internalToken)

	var resp *http.Response

	for {
		var err error

		resp, err = client.Do(req)
		if err == nil {
			break
		}

		log.Println("fail to download task archive file:", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("fail to read body of response:", err)
	}

	err = ioutil.WriteFile("/"+fileID, body, 0644)
	if err != nil {
		panic(err)
	}
}

func uploadFile() {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	f, err := os.Open("/result-" + taskID + ".zip")
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "https://"+controllerAddress+"/api/internal/uploadResult", f)
	if err != nil {
		panic(err)
	}

	req.Header.Set("RESULT_FILE", taskID+":"+fileID+":RESULT.zip")
	req.Header.Set("INTERNAL_TOKEN", internalToken)

	var resp *http.Response

	for {
		var err error

		resp, err = client.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			break
		}

		log.Println("fail to upload result archive file:", err)
	}

	defer resp.Body.Close()
}
