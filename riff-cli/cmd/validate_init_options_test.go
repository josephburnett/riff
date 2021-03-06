/*
 * Copyright 2018 the original author or authors.
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	projectriff_v1 "github.com/projectriff/riff/kubernetes-crds/pkg/apis/projectriff.io/v1"
	"github.com/projectriff/riff/riff-cli/pkg/options"
	"github.com/projectriff/riff/riff-cli/pkg/osutils"
	"github.com/stretchr/testify/assert"
)

//TODO: These tests should go away
func TestValidateDefaultFunctionResources(t *testing.T) {
	_, initOptions := Init([]projectriff_v1.Invoker{})
	filePath := initOptions.FilePath
	as := assert.New(t)
	opts := options.InitOptions{
		FilePath:     filePath,
		FunctionName: "foo",
		DryRun:       true}
	as.NoError(validateInitOptions(&opts))
}

func TestValidateCleansFunctionResources(t *testing.T) {
	filePath := filepath.Join("../test_data", "python", "demo") + string(os.PathSeparator) + string(os.PathSeparator)
	artifact := "demo.py"
	as := assert.New(t)
	opts := options.InitOptions{FilePath: filePath, Artifact: artifact, DryRun: true}
	as.NoError(validateInitOptions(&opts))
	as.Equal(osutils.Path("../test_data/python/demo"), opts.FilePath)
}

func TestValidateFunctionResourceDoesNotExist(t *testing.T) {
	filePath := osutils.Path("does/not/exist")
	as := assert.New(t)
	opts := options.InitOptions{FilePath: filePath, DryRun: true}
	err := validateInitOptions(&opts)
	as.Equal(fmt.Sprintf("path '%s' does not exist", filePath), err.Error())
}

func TestValidateArtifactIsRegularFile(t *testing.T) {
	filePath := osutils.Path("../test_data/python")
	as := assert.New(t)
	opts := options.InitOptions{FilePath: filePath, Artifact: "demo"}
	err := validateInitOptions(&opts)
	as.Error(err)
	as.Contains(err.Error(), "must be a regular file")
}

func TestValidateArtifactIsInSubDirectory(t *testing.T) {
	filePath := osutils.Path("../test_data")
	as := assert.New(t)
	opts := options.InitOptions{FilePath: filePath, Artifact: "python/demo/demo.py"}
	as.NoError(validateInitOptions(&opts))
}

func TestValidateArtifactIsInCWD(t *testing.T) {
	currentDir := osutils.GetCWD()
	os.Chdir(osutils.Path("../test_data/python/demo"))

	as := assert.New(t)
	opts := options.InitOptions{FilePath: "", Artifact: "demo.py"}
	as.NoError(validateInitOptions(&opts))
	os.Chdir(currentDir)
}

func TestArtifactCannotBeExternalToFilePath(t *testing.T) {
	currentDir := osutils.GetCWD()
	os.Chdir(osutils.Path("../test_data/python/demo"))
	filePath := osutils.GetCWD()
	as := assert.New(t)
	opts := options.InitOptions{FilePath: filePath, Artifact: osutils.Path("../multiple/one.py")}
	err := validateInitOptions(&opts)
	as.Error(err)
	as.Contains(err.Error(), "cannot be external to filepath")
	os.Chdir(currentDir)
}

func TestArtifactRelativeToFilePath(t *testing.T) {
	filePath := osutils.Path("../test_data/python/demo")
	artifact := "demo.py"
	as := assert.New(t)
	opts := options.InitOptions{FilePath: filePath, Artifact: artifact}
	as.NoError(validateInitOptions(&opts))
}

func TestAbsoluteArtifactPathConflctsFilePath(t *testing.T) {
	filePath := osutils.Path("../test_data/python/multiple/one.py")
	artifact := osutils.Path("two.py")
	as := assert.New(t)

	opts := options.InitOptions{FilePath: filePath, Artifact: artifact}
	err := validateInitOptions(&opts)
	as.Contains(err.Error(), "conflicts with filepath")
	as.Error(err)
}

func TestInvalidProtocol(t *testing.T) {
	as := assert.New(t)
	opts := options.InitOptions{Protocol: "grpz"}
	err := validateInitOptions(&opts)
	as.Error(err)
	as.Contains(err.Error(), "unsupported")
}

func TestCleanedProtocol(t *testing.T) {
	as := assert.New(t)
	opts := options.InitOptions{Protocol: "gRPC"}
	err := validateInitOptions(&opts)
	as.NoError(err)
	as.Equal("grpc", opts.Protocol)
}
