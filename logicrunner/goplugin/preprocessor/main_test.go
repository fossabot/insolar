/*
 *    Copyright 2018 INS Ecosystem
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/insolar/insolar/logicrunner/goplugin/testutil"
)

func Test_generateContractWrapper(t *testing.T) {
	buf := bytes.Buffer{}
	err := generateContractWrapper("../testplugins/secondary/main.go", &buf)
	assert.NoError(t, err)
	// io.Copy(os.Stdout, w)
	code, err := ioutil.ReadAll(&buf)
	assert.NoError(t, err)
	if len(code) == 0 {
		t.Fatal("generator returns zero length code")
	}
}

func Test_generateContractProxy(t *testing.T) {
	buf := bytes.Buffer{}
	err := generateContractProxy("../testplugins/secondary/main.go", &buf)
	assert.NoError(t, err)

	code, err := ioutil.ReadAll(&buf)
	assert.NoError(t, err)
	if len(code) == 0 {
		t.Fatal("generator returns zero length code")
	}
}

func TestCompileContractProxy(t *testing.T) {
	cwd, err := os.Getwd()
	assert.NoError(t, err)
	defer os.Chdir(cwd) // nolint: errcheck

	tmpDir, err := ioutil.TempDir("", "test-")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir) // nolint: errcheck

	err = os.MkdirAll(tmpDir+"/src/secondary/", 0777)
	assert.NoError(t, err)

	// XXX: dirty hack to make `dep` installed packages available in generated code
	err = os.Symlink(cwd+"/../../../vendor/", tmpDir+"/src/secondary/vendor")
	assert.NoError(t, err)

	proxyFh, err := os.OpenFile(tmpDir+"/src/secondary/main.go", os.O_WRONLY|os.O_CREATE, 0644)
	assert.NoError(t, err)

	err = generateContractProxy("../testplugins/secondary/main.go", proxyFh)
	assert.NoError(t, err)

	err = proxyFh.Close()
	assert.NoError(t, err)

	err = testutil.WriteFile(tmpDir, "/test.go", `
package test

import "secondary"

func main() {
	_ = secondary.GetObject("some")
}
	`)

	err = os.Chdir(tmpDir)
	assert.NoError(t, err)

	origGoPath, err := testutil.ChangeGoPath(tmpDir)
	assert.NoError(t, err)
	defer os.Setenv("GOPATH", origGoPath) // nolint: errcheck

	out, err := exec.Command("go", "build", "test.go").CombinedOutput()
	assert.NoError(t, err, string(out))
}
