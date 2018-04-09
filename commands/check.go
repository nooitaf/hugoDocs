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
	"github.com/spf13/cobra"
)

var _ cmder = (*checkCmd)(nil)

type checkCmd struct {
	cmd *cobra.Command
}

func newCheckCmd() *checkCmd {
	return &checkCmd{cmd: &cobra.Command{
		Use:   "check",
		Short: "Contains some verification checks",
	},
	}
}

func (c *checkCmd) getCommand() *cobra.Command {
	return c.cmd
}