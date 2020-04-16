# a3
Go modules and plugins, a use case

The repository contains three modules: `api`, `impl`, and `plugin`. The latter two modules import a package from the first one.
The `impl` module loads the plugin and passes the implementation to it. The plugin's job is to pretty-print the implementation.

Build the plugin from the `plugin` directory with `go build -buildmode=plugin -o p.so`. Run the use case from the `impl` directory with `go run main.go`.
