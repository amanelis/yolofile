# Introduction

I'll admit, I was once a bro that used GNU Make incorrectly. Sining in every way possible, running bash scripts, alias other make tasks just to write less in the shell, executing laughably simple tasks using only Make. It's time to repent my sins. I had a dream, it was Drake yelling at me from a dark tunnel, all I could hear is the simple phrase, YOLO.

In college I had a professor whom I greatly admired [(Hugh B. Maynard)](https://cs.utsa.edu/people/faculty/maynard-hugh-157). I know deep down in my soul that if he saw how I had been using Make in the past he would have closed his Newsgroup terminal windows, LaTeX editor, looked at me and said: that is not how you use GNU Make.

---

Let's build but not use Make because the OG bros wouldn't like us using Make for simple non-build like commands and Rake uses Ruby, so that statement speaks for itself... let's begin!

![cardinal](yolo.jpg)

## Setup
Were going to use Make to ironically setup Yolofile, one last hoora. The two below commands *should* install Yolofile into your filesystem.

	$> make build
	$> cp bin/yolo /usr/local/bin
	$> yolo
	
## Yolofile
The Yolofile works almost like a Makefile, except you can inline Go code, SUWEET. JK, not yet, that is still a work in progress. But you can execute bash commands and builds scripts inside the file just as you would a Makefile. Here is a full example:

```bash
hello:
	echo "Hello, World!"

test:
	go test ./...
```

Execute your commands as follows:

	$> yolo test
	?   	github.com/amanelis/yolofile	[no test files]
	?   	github.com/amanelis/yolofile/api [no test files]
	?   	github.com/amanelis/yolofile/cmd [no test files]
	ok  	github.com/amanelis/yolofile/helpers 0.005s
	
	$> yolo hello
	Hello, World!

## Roadmap
- Inline Golang
- Inline scope executed environment variables
- One liner to install Yolo, cause we don't have time to run Make anymore


## Disclaimer
I threw this together over a few glasses of 🍷