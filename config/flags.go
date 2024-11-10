package config

import "flag"

var (
	Jobs     bool
	Database string
)

func Flags() {
	flag.BoolVar(&Jobs, "jobs", false, "Specify using background jobs")
	flag.StringVar(&Database, "database", "dev", "Specify database to use")
	flag.Parse()
}
