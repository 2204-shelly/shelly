package main

//go:generate go-bindata -fs -prefix "public/" public/

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/mattn/go-shellwords"
)

type CommandResponse struct {
	Command []string
	Stdout []byte
	Stderr []byte
}

func main() {

	http.HandleFunc("/api/v1/super/secure/sandbox", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Read POST form data
		cmdline := r.PostFormValue("cmd")
		stdin := r.PostFormValue("stdin")

		// Split the command nicely, we're nice people
		//cmds := strings.Split(cmdline, "\x00")
		cmds, err := shellwords.Parse(cmdline)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		log.Printf("%v", cmds)

		var stdout bytes.Buffer
		var stderr bytes.Buffer

		// Run the command
		cmd := exec.Command(cmds[0], cmds[1:]...)
		cmd.Stdin = strings.NewReader(stdin)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err = cmd.Run()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Write data as JSON
		w.WriteHeader(http.StatusOK)
		b, err := json.Marshal(CommandResponse{
			Command: cmds,
			Stdout: stdout.Bytes(),
			Stderr: stderr.Bytes(),
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(b)
		log.Printf("%s", stdout.String())
		log.Printf("%s", stderr.String())
	})

	//http.Handle("/", http.FileServer(http.Dir("public")))
	http.Handle("/", http.FileServer(AssetFile()))

	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}
