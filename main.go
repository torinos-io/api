package main

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/creasty/configo"
	"github.com/honeybadger-io/honeybadger-go"
	"github.com/mitchellh/panicwrap"

	"github.com/torinos-io/api/server"
	"github.com/torinos-io/api/store"
	"github.com/torinos-io/api/type/system"
)

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	if path.Base(dir) == "bin" {
		dir = path.Join(dir, "..")
	}
	os.Chdir(dir)
}

func main() {
	c := getConfig()

	honeybadger.Configure(honeybadger.Configuration{
		Env:    c.Env,
		APIKey: c.HoneybadgerAPIKey,
	})
	defer honeybadger.Monitor()

	supervisePanic()
	start(c)
}

func start(c *system.Config) {
	ctx := &system.AppContext{Config: c}

	ctx.MainDB = store.NewDatabase(c.DatabaseURL, true)
	defer ctx.MainDB.Close()

	server.Run(ctx)
}

func getConfig() *system.Config {
	c := &system.Config{}
	if err := configo.Load(c, configo.Option{
		Dir: "./data/config",
	}); err != nil {
		panic(err)
	}
	return c
}

func supervisePanic() {
	pw := &panicwrap.WrapConfig{
		Handler: func(output string) {
			honeybadger.Notify(
				errors.New(output),
				honeybadger.ErrorClass{Name: "panicwrap"},
				honeybadger.Fingerprint{Content: output},
				honeybadger.Context{
					"Stack": strings.Split(output, "\n"),
				},
			)
			honeybadger.Flush()
			os.Exit(1)
		},
	}

	exitStatus, err := panicwrap.Wrap(pw)
	if err != nil {
		panic(err)
	}
	if exitStatus >= 0 {
		os.Exit(exitStatus)
	}
	if !panicwrap.Wrapped(pw) {
		os.Exit(0)
	}
}
