package handlers

import (
        "fmt"
        "io"
        "log"
        "net/http"
        "os"
        "path/filepath"
        "strings"
        "time"
)

func ClientHandler(rootDir string, verbose bool) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
                requestedPath := r.URL.Path
                if requestedPath == "" || requestedPath == "/" {
                        requestedPath = "/"
                }
                fullPath := filepath.Join(rootDir, filepath.Clean(requestedPath))

                // Prevent directory traversal
                absRoot, _ := filepath.Abs(rootDir)
                absPath, _ := filepath.Abs(fullPath)
                if len(absPath) < len(absRoot) || absPath[:len(absRoot)] != absRoot {
                        http.Error(w, "Invalid path", http.StatusForbidden)
                        return
                }

                info, err := os.Stat(fullPath)
                if err != nil {
                        http.Error(w, "Not found", http.StatusNotFound)
                        return
                }

                if info.IsDir() {
                        // Try to serve index.html if it exists
                        indexPath := filepath.Join(fullPath, "index.html")
                        if _, err := os.Stat(indexPath); err == nil {
                                serveFile(indexPath, verbose, w)
                                return
                        }

                        // Otherwise, generate directory listing with links
                        entries, err := os.ReadDir(fullPath)
                        if err != nil {
                                http.Error(w, "Failed to read directory", http.StatusInternalServerError)
                                return
                        }

                        w.Header().Set("Content-Type", "text/html; charset=utf-8")
                        fmt.Fprintf(w, "<html><head><meta charset=\"utf-8\"><title>Index of %s</title></head><body><h1>Index of %s</h1><hr><ul>", requestedPath, requestedPath)
                        for _, e := range entries {
                                name := e.Name()
                                if e.IsDir() {
					fmt.Fprintf(w, `<li><a href="%s/%s/">%s</a></li>`, requestedPath, name, name)
					continue
                                }
                                fmt.Fprintf(w, `<li><a href="%s">%s</a></li>`, name, name)
                        }
                        fmt.Fprint(w, "</ul></body></html>")

                        if verbose {
                                log.Printf("Listed directory: %s (%d entries)", fullPath, len(entries))
                        }
                        return
                }

                serveFile(fullPath, verbose, w)
        }
}

func serveFile(fullPath string, verbose bool, w http.ResponseWriter) {
        start := time.Now()

        // Since switch expression has default so mime => application/octet-stream will always be overwritten
        mime := "application/octet-stream"
        ext := strings.ToLower(filepath.Ext(fullPath))
        switch ext {
        case ".txt":
                mime = "text/plain"
        case ".html":
                mime = "text/html"
        case ".css":
                mime = "text/css"
        case ".js":
                mime = "application/javascript"
        case ".png":
                mime = "image/png"
        case ".jpg", ".jpeg":
                mime = "image/jpeg"
        case ".gif":
                mime = "image/gif"
        case ".svg":
                mime = "image/svg+xml"
        case ".json":
                mime = "application/json"
        case ".pdf":
                mime = "application/pdf"
        case ".mp4":
                mime = "video/mp4"
        case ".mp3":
                mime = "audio/mpeg"
        case ".zip":
                mime = "application/zip"
        default:
                mime = "text/plain"
        }

        w.Header().Set("Content-Type", mime)

        f, err := os.Open(fullPath)
        if err != nil {
                http.Error(w, "Failed to open file", http.StatusNotFound)
                return
        }
        defer f.Close()

        info, _ := f.Stat()
        fileSize := info.Size()

        _, _ = io.Copy(w, f)

        elapsed := time.Since(start)
        if verbose {
                log.Printf(
                        "Served file: %s (%d bytes, %s)",
                        fullPath,
                        fileSize,
                        elapsed,
                        )
        }
}
