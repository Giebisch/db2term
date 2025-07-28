package main

import (
	"db2term/internal/config"
	"db2term/internal/tui"
	"log/slog"
)

func main() {
	conf_file, err := config.GetConfigFile()
	if err != nil {
		slog.Error("Error getConfigFile", "error", err)
	}
	conf, err := config.ReadConfig(conf_file)
	if err != nil {
		slog.Error("Error readConfig", "error", err)
	}

	tui.NewTUI(conf)
}
