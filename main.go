// Command ps4-telemetry is a sink for PS4 telemetry hosts: it accepts any
// request, logs it, and answers 200 with an empty body. Nothing is stored.
package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/healthz" {
			fmt.Fprint(w, "ok")
			return
		}
		slog.Info("telemetry", "method", r.Method, "host", r.Host, "path", r.URL.Path, "remote", r.RemoteAddr)
		w.WriteHeader(http.StatusOK)
	})
}

func main() {
	addr := os.Getenv("PS4_TELEMETRY_ADDR")
	if addr == "" {
		addr = ":8541"
	}
	slog.Info("listening", "addr", addr)
	if err := http.ListenAndServe(addr, handler()); err != nil {
		slog.Error("serve", "err", err)
		os.Exit(1)
	}
}
