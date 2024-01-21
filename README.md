[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/cli)](https://goreportcard.com/report/github.com/go-corelibs/cli)

# cli - command-line interface utilities

a collection of utilities for working with github.com/urfave/cli

# Installation

``` shell
> go get github.com/go-corelibs/cli@latest
```

# Examples

## FlagStringer

``` go
import (
    "github.com/urfave/cli/v2"

    clcli "github.com/go-corelibs/cli"
)

func init() {
    // - remove default text from boolean flag usage
    // - place defaults and env vars on newlines
    cli.FlagStringer = clcli.NewFlagStringer().
		PruneDefaultBools(true).
		DetailsOnNewLines(true).
		Make()
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2024 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
