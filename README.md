# Introduction

Let's build but not use Make because the OG bros wouldn't like us using Make for simple non-build like commands and Rake uses Ruby, so that statement speaks for itself...

![cardinal](yolo.jpg)

## Setup
Were going to use Makefile to ironically setup Yolofile, one last hoora. The two below commands *should* install Yolofile into your filesystem.

	$> make build
	$> cp bin/yolo /usr/local/bin
	$> yolo
	
## Yolofile
The Yolofile works almost like a Makefile, except you can inline Go code, SUWEET. JK, not yet, that is still a work in progress. But you can execute bash commands and builds scripts inside the file just as you would a Makefile. Here is a full example:

```bash
hello:
	echo "Hello, World!"

test:
	go test ./... -v

bootstrap:
	bash ~/Tmp/bootstrap.sh
```

Execute your commands as follows:

	$> yolo test
	$> yolo bootstrap hello

## Roadmap
- Inline Golang
- Inline scope executed environment variables


#### Disclaimer
I threw this together over a few glasses of 🍷