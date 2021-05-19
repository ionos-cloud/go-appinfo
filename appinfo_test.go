/*
Copyright Strato AG.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package appinfo

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	oldVersion   = version
	oldName      = name
	oldMetadata  = metadata
	oldGitCommit = gitCommit
)

func resetAppInfo() {
	version = oldVersion
	gitCommit = oldGitCommit
	metadata = oldMetadata
	name = oldName
}

func TestVersion(t *testing.T) {
	defer resetAppInfo()
	version = "v1.0.0"
	require.Equal(t, version, Version())
}

func TestUserAgent(t *testing.T) {
	defer resetAppInfo()
	version = "v1.0.0"
	name = "app"
	require.Equal(t, name+"/"+version, UserAgent())
}

func TestName(t *testing.T) {
	defer resetAppInfo()
	name = "app"
	require.Equal(t, name, Name())
}

func TestAppInfo_String(t *testing.T) {
	defer resetAppInfo()

	cases := []struct {
		name      string
		gitCommit string
		metadata  string
		expBuild  string
	}{
		{
			name: "no metadata",
		}, {
			name:      "complete",
			gitCommit: "sha",
			metadata:  "build-info",
			expBuild:  "+sha.build-info",
		},
	}
	for _, c := range cases {
		name = "app"
		version = "v0.0.0"
		gitCommit = c.gitCommit
		metadata = c.metadata
		got := Get().String()
		exp := name + "/" + version + c.expBuild + " GoVersion: " + runtime.Version()
		require.Equal(t, exp, got)
		resetAppInfo()
	}
}

func TestGet(t *testing.T) {
	defer resetAppInfo()
	name = "app"
	version = "v0.0.0"
	exp := AppInfo{
		Version: version,
		Name:    name,
		Build:   Build(),
	}
	require.Equal(t, exp, Get())
}

func TestBuild(t *testing.T) {
	defer resetAppInfo()
	gitCommit = "sha"
	metadata = "meta"
	exp := BuildInfo{
		GitCommit: gitCommit,
		Metadata:  metadata,
		GoVersion: runtime.Version(),
	}
	require.Equal(t, exp, Build())
}
