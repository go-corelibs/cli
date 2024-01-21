// Copyright (c) 2024  The Go-CoreLibs Authors
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
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/go-corelibs/slices"
)

var (
	origFlagStringer = cli.FlagStringer
)

// FlagStringer is a buildable interface for constructing a pre-configured
// cli.FlagStringFunc instances
type FlagStringer interface {
	// PruneEnvVars specifies if `[$ENV_VAR...]` text is removed
	PruneEnvVars(enable bool) FlagStringer
	// PruneDefaults specifies if `(default: ...)` text is removed
	PruneDefaults(enable bool) FlagStringer
	// PruneDefaultBools specifies if only boolean flag defaults are removed
	PruneDefaultBools(enable bool) FlagStringer
	// DetailsOnNewLines specifies if defaults and env vars are places on new
	// lines instead of all in one line
	DetailsOnNewLines(enable bool) FlagStringer
	// Make produces the cli.FlagStringFunc
	Make() cli.FlagStringFunc
}

type cFlagStringer struct {
	pruneEnvVars      bool
	pruneDefaults     bool
	pruneDefaultBools bool
	detailsOnNewLines bool
}

// NewFlagStringer creates a new FlagStringer instance, ready to be configured
// and made into a cli.FlagStringFunc
func NewFlagStringer() FlagStringer {
	return &cFlagStringer{}
}

func (f *cFlagStringer) PruneEnvVars(enable bool) FlagStringer {
	f.pruneEnvVars = enable
	return f
}

func (f *cFlagStringer) PruneDefaults(enable bool) FlagStringer {
	f.pruneDefaults = enable
	return f
}

func (f *cFlagStringer) PruneDefaultBools(enable bool) FlagStringer {
	f.pruneDefaultBools = enable
	return f
}

func (f *cFlagStringer) DetailsOnNewLines(enable bool) FlagStringer {
	f.detailsOnNewLines = enable
	return f
}

func (f *cFlagStringer) Make() cli.FlagStringFunc {
	fs := &CFlagStringer{
		cfg: f,
	}
	return fs.Stringer
}

type CFlagStringer struct {
	cfg *cFlagStringer
}

func (s CFlagStringer) Stringer(flag cli.Flag) (usage string) {
	usage = origFlagStringer(flag)

	if s.cfg.pruneEnvVars {
		if b, _, a, found := slices.Carve([]rune(usage), []rune("[$"), []rune("]")); found {
			if usage = string(b); string(a) != "" {
				usage += " " + string(a)
			}
		}
	}

	if s.cfg.pruneDefaults {
		if b, _, a, found := slices.Carve([]rune(usage), []rune("(default: "), []rune(")")); found {
			before, after := string(b), string(a)
			if bLast := len(before) - 1; bLast >= 0 && before[bLast] == ' ' {
				if after != "" && after[0] == ' ' {
					usage = before + after[1:]
				} else {
					usage = before + after
				}
			} else {
				usage = before + after
			}
		}
	} else if s.cfg.pruneDefaultBools {
		if b, _, _, found := slices.Carve([]rune(usage), []rune("(default: "), []rune("false)")); found {
			usage = string(b)
		}
		if b, _, _, found := slices.Carve([]rune(usage), []rune("(default: "), []rune("true)")); found {
			usage = string(b)
		}
	}

	if s.cfg.detailsOnNewLines {
		// <text> (default: blah) [$...]
		var message, defaults, variables string
		if b, _, found := strings.Cut(usage, "(default: "); found {
			message = b
		} else if b, _, found := strings.Cut(usage, "[$"); found {
			message = b
		}
		if message != "" {
			if _, m, _, found := slices.Carve([]rune(usage), []rune("(default: "), []rune(")")); found {
				defaults = "(default: " + string(m) + ")"
			}
			if _, m, _, found := slices.Carve([]rune(usage), []rune("[$"), []rune("]")); found {
				variables = "[$" + string(m) + "]"
			}
			usage = message
			if defaults != "" {
				usage += "\n\t  " + defaults
			}
			if variables != "" {
				usage += "\n\t  " + variables
			}
		}
	}

	return
}
