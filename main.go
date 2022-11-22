package main

import (
	humanize "github.com/dustin/go-humanize"
	"github.com/minio/minio/pkg/disk"
)

func printUsage(path string) error {
	di, err := disk.GetInfo(path)
	if err != nil {
		return err
	}
	percentage := (float64(di.Total-di.Free) / float64(di.Total)) * 100
	fmt.Printf("%s of %s disk space used (%0.2f%%)\n",
		humanize.Bytes(di.Total-di.Free),
		humanize.Bytes(di.Total),
		percentage,
	)
}
