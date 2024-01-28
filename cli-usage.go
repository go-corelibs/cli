// Copyright (c) 2024  The Go-Curses Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"fmt"
	"os"
	"sort"

	"github.com/maruel/natural"
	"github.com/urfave/cli/v2"
)

// ShowOptions prints each option (and it's aliases), one per line, to
// os.Stderr and with flags named "usage", "help", "version" and "verbose"
// separated and sorted after all other flags
func ShowOptions(ctx *cli.Context) {
	_, _ = fmt.Fprintln(os.Stderr, "options:")
	if len(ctx.App.Flags) > 0 {
		var before, after []string
		for _, f := range ctx.App.Flags {
			var b *BaseFlag
			if b = DecodeBaseFlag(f); b == nil || b.Hidden {
				continue
			}
			switch b.Name {
			case "usage", "help", "version", "verbose":
				after = append(after, b.NameWithAliases())
			default:
				before = append(before, b.NameWithAliases())
			}
		}
		if len(before) > 0 {
			sort.Sort(natural.StringSlice(before))
			for _, l := range before {
				_, _ = fmt.Fprintln(os.Stderr, "\t"+l)
			}
		}
		if len(after) > 0 {
			if len(before) > 0 {
				_, _ = fmt.Fprintf(os.Stderr, "\n")
			}
			sort.Sort(natural.StringSlice(after))
			for _, l := range after {
				_, _ = fmt.Fprintln(os.Stderr, "\t"+l)
			}
		}
	}
}

// ShowUsage prints a very brief usage text line to os.Stderr
func ShowUsage(ctx *cli.Context) {
	message := "usage: "
	if ctx.App.UsageText != "" {
		message += ctx.App.UsageText
	} else {
		message += ctx.App.Name + " [options] [args...]"
	}
	message += "\nuse -h or --help for more details"
	_, _ = fmt.Fprintln(os.Stderr, message)
}

// ShowUsageAndExit calls ShowUsage and then os.Exit with the `exitCode` given
func ShowUsageAndExit(ctx *cli.Context, exitCode int) {
	ShowUsage(ctx)
	os.Exit(exitCode)
}

// ShowUsageOptions is like ShowUsage but includes a listing of all option
// names and aliases
func ShowUsageOptions(ctx *cli.Context) {
	message := "usage: "
	if ctx.App.UsageText != "" {
		message += ctx.App.UsageText
	} else {
		message += ctx.App.Name + " [options] [args...]"
	}
	_, _ = fmt.Fprintln(os.Stderr, message+"\n")
	ShowOptions(ctx)
}

// ShowUsageOptionsAndExit calls ShowUsageOptions and then os.Exit with the
// `exitCode` given
func ShowUsageOptionsAndExit(ctx *cli.Context, exitCode int) {
	ShowUsageOptions(ctx)
	os.Exit(exitCode)
}
