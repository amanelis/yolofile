package api

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	h "github.com/amanelis/yolofile/helpers"
)

// Execute run the high leve parser to interact with the API
func Execute(args []string) {
	if len(args) <= 1 {
		os.Exit(1)
	}

	cmd := args[1]

	file, err := os.Open(h.GetYolofile())
	if err != nil {
	    panic(err)
	}
	defer file.Close()

	var cKey string
	maps := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		method, _ := regexp.MatchString("^(.*):", line)
		if method {
			cKey = strings.Replace(line, ":", "", -1)
			maps[cKey] = []string{}

			continue
		}

		if line == "" {
			continue
		}

		maps[cKey] = append(maps[cKey], strings.TrimSpace(line)  + ";")
	}

	if _, isKeyPresent := maps[cmd]; !isKeyPresent {
		fmt.Println("*command not found in Yolofile")
		os.Exit(1)
  }

	exe := exec.Command("/bin/sh", "-c", strings.Join(maps[cmd][:], " "))

	var stdout, stderr bytes.Buffer
  exe.Stdout = &stdout
  exe.Stderr = &stderr

	if err := exe.Run(); err != nil {
		panic(err)
	}

	fmt.Printf(string(stdout.Bytes()))
}
