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
	"github.com/urfave/cli/v2"
)

// BaseFlag is a type containing only the common variables for any given
// cli.Flag type
type BaseFlag struct {
	Name        string
	Category    string
	DefaultText string
	Usage       string
	Required    bool
	Hidden      bool
	Aliases     []string
	EnvVars     []string
	original    cli.Flag
}

// Original returns the actual cli.Flag this BaseFlag instance was derived
// from
func (b *BaseFlag) Original() (f cli.Flag) {
	f = b.original
	return
}

// NameWithAliases returns the flag name and any aliases, with leading dashes
func (b *BaseFlag) NameWithAliases() (name string) {
	if len(b.Name) == 1 {
		name = "-" + b.Name
	} else {
		name = "--" + b.Name
	}
	if len(b.Aliases) > 0 {
		name += " ("
		for idx, alias := range b.Aliases {
			if idx > 0 {
				name += ", "
			}
			if len(alias) == 1 {
				name += "-" + alias
			} else {
				name += "--" + alias
			}
		}
		name += ")"
	}
	return
}

// NewBaseFlag constructs a new BaseFlag instance
func NewBaseFlag(name, category, defaultText, usage string, required, hidden bool, aliases, envVars []string, original cli.Flag) *BaseFlag {
	return &BaseFlag{
		Name:        name,
		Category:    category,
		DefaultText: defaultText,
		Usage:       usage,
		Required:    required,
		Hidden:      hidden,
		Aliases:     aliases,
		EnvVars:     envVars,
		original:    original,
	}
}

// DecodeBaseFlag returns a BaseFlag representation of the given cli.Flag
//
//gocyclo:ignore
func DecodeBaseFlag(flag cli.Flag) (base *BaseFlag) {
	switch t := flag.(type) {

	case *cli.BoolFlag:
		return NewBaseFlag(t.Name, t.Category, t.DefaultText, t.Usage, t.Required, t.Hidden, t.Aliases, t.EnvVars, t)

	case *cli.DurationFlag:
		return NewBaseFlag(t.Name, t.Category, t.DefaultText, t.Usage, t.Required, t.Hidden, t.Aliases, t.EnvVars, t)

	case *cli.Float64Flag:
		return NewBaseFlag(t.Name, t.Category, t.DefaultText, t.Usage, t.Required, t.Hidden, t.Aliases, t.EnvVars, t)

	case *cli.Float64SliceFlag:
		return NewBaseFlag(t.Name, t.Category, t.DefaultText, t.Usage, t.Required, t.Hidden, t.Aliases, t.EnvVars, t)

	case *cli.GenericFlag:
		return NewBaseFlag(t.Name, t.Category, t.DefaultText, t.Usage, t.Required, t.Hidden, t.Aliases, t.EnvVars, t)

	case *cli.Int64Flag:
		return NewBaseFlag(t.Name, t.Category, t.DefaultText, t.Usage, t.Required, t.Hidden, t.Aliases, t.EnvVars, t)

	case *cli.Int64SliceFlag:
		return NewBaseFlag(t.Name, t.Category, t.DefaultText, t.Usage, t.Required, t.Hidden, t.Aliases, t.EnvVars, t)

	case *cli.IntFlag:
		return NewBaseFlag(t.Name, t.Category, t.DefaultText, t.Usage, t.Required, t.Hidden, t.Aliases, t.EnvVars, t)

	case *cli.IntSliceFlag:
		return NewBaseFlag(t.Name, t.Category, t.DefaultText, t.Usage, t.Required, t.Hidden, t.Aliases, t.EnvVars, t)

	case *cli.PathFlag:
		return NewBaseFlag(t.Name, t.Category, t.DefaultText, t.Usage, t.Required, t.Hidden, t.Aliases, t.EnvVars, t)

	case *cli.StringFlag:
		return NewBaseFlag(t.Name, t.Category, t.DefaultText, t.Usage, t.Required, t.Hidden, t.Aliases, t.EnvVars, t)

	case *cli.StringSliceFlag:
		return NewBaseFlag(t.Name, t.Category, t.DefaultText, t.Usage, t.Required, t.Hidden, t.Aliases, t.EnvVars, t)

	case *cli.TimestampFlag:
		return NewBaseFlag(t.Name, t.Category, t.DefaultText, t.Usage, t.Required, t.Hidden, t.Aliases, t.EnvVars, t)

	case *cli.Uint64Flag:
		return NewBaseFlag(t.Name, t.Category, t.DefaultText, t.Usage, t.Required, t.Hidden, t.Aliases, t.EnvVars, t)

	case *cli.Uint64SliceFlag:
		return NewBaseFlag(t.Name, t.Category, t.DefaultText, t.Usage, t.Required, t.Hidden, t.Aliases, t.EnvVars, t)

	case *cli.UintFlag:
		return NewBaseFlag(t.Name, t.Category, t.DefaultText, t.Usage, t.Required, t.Hidden, t.Aliases, t.EnvVars, t)

	case *cli.UintSliceFlag:
		return NewBaseFlag(t.Name, t.Category, t.DefaultText, t.Usage, t.Required, t.Hidden, t.Aliases, t.EnvVars, t)

	default:
	}

	return
}
