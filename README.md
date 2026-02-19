# kimani-s-project
building the capstone project
#  Getting Started with Go (Golang) — A Beginner's Web Toolkit Guide

## 1. Title & Objective

**Example:** "Getting Started with Go (Golang) in Web Development — A Beginner's Guide"

- What technology did you choose? **Go (Golang)**
- Why did you choose it? Fast, clean syntax, no external dependencies needed, high job market demand
- What's the end goal (e.g., create a simple HTTP web server)? Build a working web server with routing and a REST API endpoint that returns JSON

---

## 2. Quick Summary of the Technology

> *e.g., "Go is a statically typed, compiled language designed at Google."*

- **What is it?** Go (also called Golang) is an open-source programming language built for simplicity, speed, and reliability
- **Where is it used?** Backend web servers, cloud infrastructure, DevOps tooling — used by Google, Docker, Uber, and Cloudflare
- **One real-world example:** Docker, the containerization platform, is written entirely in Go

---

## 3. System Requirements

- OS: Linux/Mac/Windows
- Tools: Go 1.21+, a terminal (Command Prompt, PowerShell, or bash), a code editor (VS Code recommended)
- Any packages/deps: **None** — this project uses Go's standard library only

---

## 4. Installation & Setup Instructions

*Step-by-step setup with commands (or terminal output if possible):*

**Windows:**
```bash
# Download the installer from https://go.dev/dl/
# Run the .msi file and follow the prompts
# Then verify in a NEW terminal window:
go version
```

**Ubuntu/Linux:**
```bash
sudo apt update
sudo apt install golang-go
go version
```

**Initialize your project:**
```bash
mkdir go-web-toolkit
cd go-web-toolkit
go mod init go-web-toolkit
mkdir -p cmd/server
```

---

## 5. Minimal Working Example

*One clear example that shows the app does something — provide code with inline comments and sample output:*

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

// helloHandler responds with a plain text Hello World message
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World! Welcome to the Go Web Toolkit 🚀")
}

func main() {
    // Register the route and start the server on port 8080
    http.HandleFunc("/", helloHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

**Run it:**
```bash
go run main.go
```

**Expected output in browser at http://localhost:8080:**
```
Hello, World! Welcome to the Go Web Toolkit 🚀
```

---

## 6. AI Prompt Journal

*For each prompt used, record:*

**Prompt 1:**
> *"How do I create a simple HTTP web server in Go using only the standard library?"*

- **AI's response summary:** Explained the `net/http` package, `http.HandleFunc`, and `http.ListenAndServe` with a minimal working example
- **Your evaluation of its helpfulness:** Very clear — the setup only required one import and ran immediately with no errors

---

**Prompt 2:**
> *"Show me how to return a JSON response from a Go HTTP handler"*

- **AI's response summary:** Introduced `encoding/json`, Go structs with JSON field tags, and `json.NewEncoder(w).Encode(response)`
- **Your evaluation of its helpfulness:** Excellent — explained struct tags like `` `json:"message"` `` which was new and very useful

---

**Prompt 3:**
> *"What is http.ServeMux and how do I use it for routing in Go?"*

- **AI's response summary:** Explained that `ServeMux` is Go's built-in router and showed how to register multiple routes with `mux.HandleFunc`
- **Your evaluation of its helpfulness:** Helpful — made it easy to expand from one route to many without any third-party libraries

---

**Prompt 4:**
> *"How do I write unit tests for Go HTTP handlers using httptest?"*

- **AI's response summary:** Introduced `net/http/httptest`, `httptest.NewRequest`, and `httptest.NewRecorder` to simulate HTTP calls in tests
- **Your evaluation of its helpfulness:** Very useful — Go's built-in test tooling was more powerful than expected and required zero setup

---

## 7. Common Issues & Fixes

- **`go: command not found`** — Go is not installed or PATH is not set. On Windows, restart your computer after installing. On Linux, run `source ~/.bashrc` after updating PATH
- **Port already in use** — Another process is using port 8080. Change `:8080` to `:9090` in `http.ListenAndServe` and retry
- **Browser shows nothing** — Make sure you're visiting `http://localhost:8080` (not https) and the terminal shows no errors

---

## 8. References

- **Official docs:** https://pkg.go.dev/net/http
- **Video links:** https://tour.golang.org (interactive browser-based Go tour)
- **Helpful blog posts:** https://gobyexample.com — Go by Example (great for quick reference)
