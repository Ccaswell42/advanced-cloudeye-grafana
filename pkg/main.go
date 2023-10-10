package main

import (
	"os"

	"github.com/cloudru/cloudeye-grafana/pkg/plugin"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func main() {
	if err := datasource.Serve(plugin.NewCloudEye()); err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}
