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

type cMakeFlagStringer struct {
	pruneEnvVars      bool
	pruneDefaults     bool
	pruneDefaultBools bool
	detailsOnNewLines bool
}

// NewFlagStringer creates a new FlagStringer instance, ready to be configured
// and made into a cli.FlagStringFunc
func NewFlagStringer() FlagStringer {
	return &cMakeFlagStringer{}
}

func (f *cMakeFlagStringer) PruneEnvVars(enable bool) FlagStringer {
	f.pruneEnvVars = enable
	return f
}

func (f *cMakeFlagStringer) PruneDefaults(enable bool) FlagStringer {
	f.pruneDefaults = enable
	return f
}

func (f *cMakeFlagStringer) PruneDefaultBools(enable bool) FlagStringer {
	f.pruneDefaultBools = enable
	return f
}

func (f *cMakeFlagStringer) DetailsOnNewLines(enable bool) FlagStringer {
	f.detailsOnNewLines = enable
	return f
}

func (f *cMakeFlagStringer) Make() cli.FlagStringFunc {
	fs := &cFlagStringer{
		cfg: f,
	}
	return fs.Stringer
}

type cFlagStringer struct {
	cfg *cMakeFlagStringer
}

func (s cFlagStringer) pruneEnvVars(usage string) (pruned string) {
	pruned = usage
	if b, _, a, found := slices.CarveString(pruned, "[$", "]"); found {
		if pruned = b; a != "" {
			pruned += " " + a
		}
	}
	return
}

func (s cFlagStringer) pruneDefaults(usage string) (pruned string) {
	pruned = usage
	if b, _, a, found := slices.CarveString(pruned, "(default: ", ")"); found {
		before, after := b, a
		if bLast := len(before) - 1; bLast >= 0 && before[bLast] == ' ' {
			if after != "" && after[0] == ' ' {
				pruned = before + after[1:]
			} else {
				pruned = before + after
			}
		} else {
			pruned = before + after
		}
	}
	return
}

func (s cFlagStringer) pruneDefaultBools(usage string) (pruned string) {
	pruned = usage
	if b, _, _, found := slices.CarveString(pruned, "(default: ", "false)"); found {
		pruned = b
	}
	if b, _, _, found := slices.CarveString(pruned, "(default: ", "true)"); found {
		pruned = b
	}
	return
}

func (s cFlagStringer) detailsOnNewLines(usage string) (pruned string) {
	pruned = usage
	var message, defaults, variables string
	if before, _, found := strings.Cut(pruned, "(default: "); found {
		message = before
	} else if b, _, ok := strings.Cut(pruned, "[$"); ok {
		message = b
	}
	if message != "" {
		if _, m, _, found := slices.Carve([]rune(pruned), []rune("(default: "), []rune(")")); found {
			defaults = "(default: " + string(m) + ")"
		}
		if _, m, _, found := slices.Carve([]rune(pruned), []rune("[$"), []rune("]")); found {
			variables = "[$" + string(m) + "]"
		}
		pruned = message
		if defaults != "" {
			pruned += "\n\t  " + defaults
		}
		if variables != "" {
			pruned += "\n\t  " + variables
		}
	}
	return
}

func (s cFlagStringer) Stringer(flag cli.Flag) (usage string) {
	usage = origFlagStringer(flag)

	if s.cfg.pruneEnvVars {
		usage = s.pruneEnvVars(usage)
	}

	if s.cfg.pruneDefaults {
		usage = s.pruneDefaults(usage)
	} else if s.cfg.pruneDefaultBools {
		usage = s.pruneDefaultBools(usage)
	}

	if s.cfg.detailsOnNewLines {
		// <text> (default: blah) [$...]
		usage = s.detailsOnNewLines(usage)
	}

	return
}
