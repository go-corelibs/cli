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

// Package cli provides command-line interface utilities
package cli

import (
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/maps"
)

func ClearEmptyCategories(flags []cli.Flag) {
	// this function is not easily maintained, needs upstream fix
	// see: https://github.com/urfave/cli/issues/1860

	cats := make(map[string][]int)
	visible := make(map[string][]int)

	for idx, flag := range flags {
		switch t := flag.(type) {
		case *cli.DurationFlag:
			if cats[t.Category] = append(cats[t.Category], idx); !t.Hidden {
				visible[t.Category] = append(visible[t.Category], idx)
			}
		case *cli.Float64Flag:
			if cats[t.Category] = append(cats[t.Category], idx); !t.Hidden {
				visible[t.Category] = append(visible[t.Category], idx)
			}
		case *cli.Float64SliceFlag:
			if cats[t.Category] = append(cats[t.Category], idx); !t.Hidden {
				visible[t.Category] = append(visible[t.Category], idx)
			}
		case *cli.GenericFlag:
			if cats[t.Category] = append(cats[t.Category], idx); !t.Hidden {
				visible[t.Category] = append(visible[t.Category], idx)
			}
		case *cli.Int64Flag:
			if cats[t.Category] = append(cats[t.Category], idx); !t.Hidden {
				visible[t.Category] = append(visible[t.Category], idx)
			}
		case *cli.Int64SliceFlag:
			if cats[t.Category] = append(cats[t.Category], idx); !t.Hidden {
				visible[t.Category] = append(visible[t.Category], idx)
			}
		case *cli.IntFlag:
			if cats[t.Category] = append(cats[t.Category], idx); !t.Hidden {
				visible[t.Category] = append(visible[t.Category], idx)
			}
		case *cli.IntSliceFlag:
			if cats[t.Category] = append(cats[t.Category], idx); !t.Hidden {
				visible[t.Category] = append(visible[t.Category], idx)
			}
		case *cli.PathFlag:
			if cats[t.Category] = append(cats[t.Category], idx); !t.Hidden {
				visible[t.Category] = append(visible[t.Category], idx)
			}
		case *cli.StringFlag:
			if cats[t.Category] = append(cats[t.Category], idx); !t.Hidden {
				visible[t.Category] = append(visible[t.Category], idx)
			}
		case *cli.StringSliceFlag:
			if cats[t.Category] = append(cats[t.Category], idx); !t.Hidden {
				visible[t.Category] = append(visible[t.Category], idx)
			}
		case *cli.TimestampFlag:
			if cats[t.Category] = append(cats[t.Category], idx); !t.Hidden {
				visible[t.Category] = append(visible[t.Category], idx)
			}
		case *cli.Uint64Flag:
			if cats[t.Category] = append(cats[t.Category], idx); !t.Hidden {
				visible[t.Category] = append(visible[t.Category], idx)
			}
		case *cli.Uint64SliceFlag:
			if cats[t.Category] = append(cats[t.Category], idx); !t.Hidden {
				visible[t.Category] = append(visible[t.Category], idx)
			}
		case *cli.UintFlag:
			if cats[t.Category] = append(cats[t.Category], idx); !t.Hidden {
				visible[t.Category] = append(visible[t.Category], idx)
			}
		case *cli.UintSliceFlag:
			if cats[t.Category] = append(cats[t.Category], idx); !t.Hidden {
				visible[t.Category] = append(visible[t.Category], idx)
			}
		default:
		}
	}

	for _, category := range maps.Keys(cats) {
		if len(visible[category]) == 0 {
			for _, idx := range cats[category] {
				switch t := flags[idx].(type) {
				case *cli.DurationFlag:
					t.Category = ""
				case *cli.Float64Flag:
					t.Category = ""
				case *cli.Float64SliceFlag:
					t.Category = ""
				case *cli.GenericFlag:
					t.Category = ""
				case *cli.Int64Flag:
					t.Category = ""
				case *cli.Int64SliceFlag:
					t.Category = ""
				case *cli.IntFlag:
					t.Category = ""
				case *cli.IntSliceFlag:
					t.Category = ""
				case *cli.PathFlag:
					t.Category = ""
				case *cli.StringFlag:
					t.Category = ""
				case *cli.StringSliceFlag:
					t.Category = ""
				case *cli.TimestampFlag:
					t.Category = ""
				case *cli.Uint64Flag:
					t.Category = ""
				case *cli.Uint64SliceFlag:
					t.Category = ""
				case *cli.UintFlag:
					t.Category = ""
				case *cli.UintSliceFlag:
					t.Category = ""
				default:
				}
			}
		}
	}

	return
}
