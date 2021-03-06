/*
 * Copyright 2017 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"github.com/projectriff/riff/riff-cli/pkg/docker"
	"github.com/projectriff/riff/riff-cli/pkg/initializer"
	"github.com/projectriff/riff/riff-cli/pkg/kubectl"
	"github.com/spf13/cobra"
)

// CreateAndWireRootCommand creates all riff commands and sub commands, as well as the top-level 'root' command,
// wires them together and returns the root command, ready to execute.
func CreateAndWireRootCommand(realDocker docker.Docker, dryRunDocker docker.Docker,
	realKubeCtl kubectl.KubeCtl, dryRunKubeCtl kubectl.KubeCtl) (*cobra.Command, error) {

	invokers, err := initializer.LoadInvokers(realKubeCtl)
	if err != nil {
		return nil, err
	}

	rootCmd := Root()

	initCmd, initOptions := Init(invokers)
	initInvokerCmds, err := InitInvokers(invokers, initOptions)
	if err != nil {
		return nil, err
	}
	initCmd.AddCommand(initInvokerCmds...)

	buildCmd, _ := Build(realDocker, dryRunDocker)

	applyCmd, _ := Apply(realKubeCtl, dryRunKubeCtl)

	createCmd := Create(initCmd, buildCmd, applyCmd)
	createInvokerCmds := CreateInvokers(invokers, initInvokerCmds, buildCmd, applyCmd)
	createCmd.AddCommand(createInvokerCmds...)

	deleteCmd, _ := Delete()

	rootCmd.AddCommand(
		applyCmd,
		buildCmd,
		createCmd,
		deleteCmd,
		initCmd,
		List(),
		Logs(),
		Publish(),
		Update(buildCmd, applyCmd),
		Version(),
	)

	rootCmd.AddCommand(
		Completion(rootCmd),
		Docs(rootCmd),
	)

	return rootCmd, nil
}
