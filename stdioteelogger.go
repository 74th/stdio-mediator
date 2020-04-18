package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
)

type reader struct {
	name   string
	reader *bufio.Reader
}

func main() {
	var logFilePath string
	logFilePath, ok := os.LookupEnv("STDIO_TEE_LOGGER_LOG")
	if !ok {
		logFilePath = "stdio_tee_logger.log"
	}

	f, err := os.Create(logFilePath)
	if err != nil {
		log.Printf("cannot open %s", err.Error())
		os.Exit(1)
	}
	defer f.Close()
	log.SetFlags(log.Lmicroseconds)
	log.SetOutput(f)

	exe := exec.Command(os.Args[1], os.Args[2:]...)
	exeStdout, _ := exe.StdoutPipe()
	exeStdin, _ := exe.StdinPipe()
	exeStderr, _ := exe.StderrPipe()
	err = exe.Start()
	if err != nil {
		log.Printf("cannot start %s", err.Error())
		os.Exit(1)
	}

	// readers := []reader{
	readers := []struct {
		mark   string
		reader *bufio.Reader
	}{
		{
			">",
			bufio.NewReader(io.TeeReader(os.Stdin, exeStdin)),
		},
		{
			"<",
			bufio.NewReader(io.TeeReader(exeStdout, os.Stdout)),
		},
		{
			"E<",
			bufio.NewReader(io.TeeReader(exeStderr, os.Stderr)),
		},
	}

	var done sync.WaitGroup
	done.Add(len(readers))
	for _, r := range readers {
		go func(mark string, reader *bufio.Reader) {
			for {
				l, _, err := reader.ReadLine()
				if err == io.EOF {
					log.Printf("%s: EOF", mark)
					break
				}
				log.Printf("%s: %s", mark, l)
			}
			done.Done()
		}(r.mark, r.reader)
	}
	done.Wait()
}
