package api

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	h "github.com/amanelis/yolofile/helpers"
)

// Execute run the high leve parser to interact with the API
func Execute(args []string) {
	fmt.Println("Command: ", args[1])
	fmt.Println("Yolofile: ", h.GetYolofile())

	file, err := os.Open(h.GetYolofile())
	if err != nil {
	    panic(err)
	}
	defer file.Close()

	var cKey string
	maps := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		method, _ := regexp.MatchString("^(.*):", line)
		if method {
			cKey = strings.Replace(strings.TrimSpace(line), ":", "", -1)
			maps[cKey] = []string{}

			continue
		}

		maps[cKey] = append(maps[cKey], strings.TrimSpace(line))
	}

	fmt.Println(maps)
	// for _, v := range data {
	// 	fmt.Println(v)
	// }

	if err := scanner.Err(); err != nil {
	    panic(err)
	}
}
