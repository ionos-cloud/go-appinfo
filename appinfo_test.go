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
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	version = "(devel)" // version available during tests

	oldName      = name
	oldMetadata  = metadata
	oldBuildDate = buildDate
)

func resetAppInfo() {
	name = oldName
	metadata = oldMetadata
	buildDate = oldBuildDate
}

func TestVersion(t *testing.T) {
	require.Equal(t, version, Version())
}

func TestUserAgent(t *testing.T) {
	t.Cleanup(resetAppInfo)
	name = "app"
	require.Equal(t, name+"/"+version, UserAgent())
}

func TestName(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		require.Equal(t, "go-appinfo", Name())
	})

	t.Run("custom name", func(t *testing.T) {
		t.Cleanup(resetAppInfo)
		name = "app"
		require.Equal(t, name, Name())
	})
}

func TestAppInfo_String(t *testing.T) {
	cases := []struct {
		name     string
		metadata string
		expBuild string
		date     string
	}{
		{
			name:     "no metadata",
			expBuild: "+hash.dirty",
		}, {
			name:     "complete",
			metadata: "build-info",
			date:     "2021-05-19T15:24:12Z",
			expBuild: "+hash.dirty.buildDate:2021-05-19T15:24:12Z.build-info",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Cleanup(resetAppInfo)
			name = "app"
			metadata = c.metadata
			buildDate = c.date
			info := AppInfo{
				Name:    Name(),
				Version: Version(),
				Build:   Build(),
			}
			info.Build.GitCommit = "hash"
			info.Build.GitTreeState = "dirty"
			got := info.String()
			exp := name + "/" + version + c.expBuild + " GoVersion: " + runtime.Version() +
				" GoPlatform: " + runtime.GOARCH + " GoCompiler: " + runtime.Compiler
			require.Equal(t, exp, got)
		})
	}
}

func TestGet(t *testing.T) {
	t.Cleanup(resetAppInfo)
	name = "app"
	exp := AppInfo{
		Version: version,
		Name:    name,
		Build:   Build(),
	}
	require.Equal(t, exp, Get())
}

func TestBuild(t *testing.T) {
	t.Cleanup(resetAppInfo)
	metadata = "meta"
	exp := BuildInfo{
		Date:      buildDate,
		Metadata:  metadata,
		GoVersion: runtime.Version(),
		Compiler:  runtime.Compiler,
		Platform:  runtime.GOARCH,
	}
	require.Equal(t, exp, Build())
}

func TestBuildInfo_populateFromSettings(t *testing.T) {
	b := BuildInfo{}
	settings := []debug.BuildSetting{
		{Key: "vcs.revision", Value: "hash"},
		{Key: "vcs.modified", Value: "true"},
	}
	b.populateFromSettings(settings)
	require.Equal(t, "hash", b.GitCommit)
	require.Equal(t, "dirty", b.GitTreeState)
}
