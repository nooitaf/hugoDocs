// Copyright 2015 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"runtime"
	"strings"

	"github.com/gohugoio/hugo/helpers"
	"github.com/gohugoio/hugo/hugolib"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var _ cmder = (*versionCmd)(nil)

type versionCmd struct {
	cmd *cobra.Command
}

func (c *versionCmd) getCommand() *cobra.Command {
	return c.cmd
}

func newVersionCmd() *versionCmd {
	return &versionCmd{
		&cobra.Command{
			Use:   "version",
			Short: "Print the version number of Hugo",
			Long:  `All software has versions. This is Hugo's.`,
			RunE: func(cmd *cobra.Command, args []string) error {
				printHugoVersion()
				return nil
			},
		},
	}
}

func printHugoVersion() {
	if hugolib.CommitHash == "" {
		jww.FEEDBACK.Printf("Hugo Static Site Generator v%s %s/%s BuildDate: %s\n", helpers.CurrentHugoVersion, runtime.GOOS, runtime.GOARCH, hugolib.BuildDate)
	} else {
		jww.FEEDBACK.Printf("Hugo Static Site Generator v%s-%s %s/%s BuildDate: %s\n", helpers.CurrentHugoVersion, strings.ToUpper(hugolib.CommitHash), runtime.GOOS, runtime.GOARCH, hugolib.BuildDate)
	}
}