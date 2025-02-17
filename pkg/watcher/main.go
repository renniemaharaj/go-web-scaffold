package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	fileChecksums = make(map[string]string)
	excludeDirs   = []string{".git", ".vscode", "pkg/watcher", "node_modules"}
	excludeFiles  = []string{"go.mod", "go.sum", "pkg/watcher/main.go", "static/index.html"}
	mainFilePath  = "cmd/example/main.go"
	wsClients     = make(map[*websocket.Conn]bool)
	mutex         = &sync.Mutex{}
	upgrader      = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	cmd           *exec.Cmd
)

func restartCommand() {
	log.Println("Attempting to re/start Go application...")
	if cmd != nil && cmd.Process != nil {
		log.Println("Killing previous application instance...")
		cmd.Process.Kill()
	}

	log.Println("Starting new application instance...")
	cmd = exec.Command("go", "run", mainFilePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Printf("Error starting application: %v\n", err)
		return
	}

	go func() {
		if err := cmd.Wait(); err != nil {
			log.Printf("Application error: %v\n", err)
			return
		}
		log.Println("Application started successfully!")
		sendUpdatedFiles()
	}()

}

func calculateChecksum(path string) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := md5.Sum(data)
	lastModified := info.ModTime().String()
	return hex.EncodeToString(hash[:]) + lastModified, nil
}

func walkFunc(root string, initial bool) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		absPath, err := filepath.Abs(path)
		if err != nil {
			return err
		}

		// Check if path is in excluded directories
		for _, dir := range excludeDirs {
			absDir, _ := filepath.Abs(dir)
			relPath, _ := filepath.Rel(absDir, absPath)

			// If the path is inside an excluded directory, skip it
			if relPath == "." || (!strings.HasPrefix(relPath, "..") && absPath != root) {
				if initial {
					log.Printf("Excluding directory: %s\n", absDir)
				}
				return filepath.SkipDir
			}
		}

		// Check if the path is in excluded files
		for _, file := range excludeFiles {
			absFile, _ := filepath.Abs(file)
			if absPath == absFile {
				if initial {
					log.Printf("Excluding file: %s\n", absFile)
				}
				return nil
			}
		}

		// If it's a file, check for changes
		if !info.IsDir() {
			checksum, err := calculateChecksum(path)
			if err != nil {
				log.Printf("Error calculating checksum: %v\n", err)
				return err
			}

			if oldChecksum, ok := fileChecksums[path]; !ok || oldChecksum != checksum {
				fileChecksums[path] = checksum
				if !initial {
					log.Printf("File changed: %s\n", path)
					restartCommand()
				}
			}
		}
		return nil
	})
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	wsClients[conn] = true
	log.Printf("Client connected: %s\n", conn.RemoteAddr())
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Client disconnected: %s\n", conn.RemoteAddr())
			delete(wsClients, conn)
			break
		}
	}
}

func sendUpdatedFiles() {
	mutex.Lock()
	defer mutex.Unlock()

	for client := range wsClients {
		log.Printf("Sending file updates to: %s\n", client.RemoteAddr())
		err := client.WriteMessage(websocket.TextMessage, []byte("reload"))
		if err != nil {
			log.Println("Error sending file update:", err)
			delete(wsClients, client)
		}
	}
}

func startServer() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("static")))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/ws", wsHandler)

	server := &http.Server{Addr: ":8080", Handler: mux}
	go func() {
		log.Println("Server running on port 8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("Server error:", err)
		}
	}()
}

func main() {
	go startServer()
	root := "."
	log.Printf("Watching directory: %s\n", root)
	err := walkFunc(root, true)
	if err != nil {
		log.Fatal(err)
	}

	restartCommand()
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if err := walkFunc(root, false); err != nil {
			log.Fatal(err)
		}
	}
}
