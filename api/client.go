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

	exe := exec.Command("/bin/sh", "-c", strings.Join(maps[cmd][:], " "))

	var stdout, stderr bytes.Buffer
  exe.Stdout = &stdout
  exe.Stderr = &stderr

	if err := exe.Run(); err != nil {
		panic(err)
	}

	outStr, _ := string(stdout.Bytes()), string(stderr.Bytes())

	fmt.Printf(outStr)

	if err := scanner.Err(); err != nil {
	    panic(err)
	}
}
