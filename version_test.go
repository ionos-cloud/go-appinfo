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

package version

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/mod/semver"
)

// verify semantic versioning compatibility of GetVersion
func TestGetVersion(t *testing.T) {
	oldVersion := version
	oldGitCommit := gitCommit
	oldMetadata := metadata
	oldPreRelease := preRelease

	reset := func() {
		version = oldVersion
		gitCommit = oldGitCommit
		metadata = oldMetadata
		preRelease = oldPreRelease
	}
	defer reset()

	cases := []struct {
		name          string
		gitCommit     string
		metadata      string
		preRelease    string
		expPrerelease string
		expBuild      string
	}{
		{
			name:      "no pre-release",
			gitCommit: "sha",
			metadata:  "build-info",
			expBuild:  "+sha.build-info",
		}, {
			name:          "no build info",
			preRelease:    "rc1",
			expPrerelease: "-rc1",
		}, {
			name:          "complete",
			gitCommit:     "sha",
			metadata:      "build-info",
			preRelease:    "rc1",
			expBuild:      "+sha.build-info",
			expPrerelease: "-rc1",
		},
	}
	for _, c := range cases {
		version = "v0.0.0"
		gitCommit = c.gitCommit
		metadata = c.metadata
		preRelease = c.preRelease
		got := GetVersion()
		exp := version + c.expPrerelease + c.expBuild
		require.Equal(t, exp, got)
		require.Equal(t, c.expPrerelease, semver.Prerelease(got))
		require.Equal(t, c.expBuild, semver.Build(got))
		require.True(t, semver.IsValid(got))
		reset()
	}
}

func TestGet(t *testing.T) {
	gitCommit = "git"
	metadata = "buildInfo"
	version = "v0.0.0"
	exp := BuildInfo{
		Version:   GetVersion(),
		GitCommit: gitCommit,
		Metadata:  metadata,
		GoVersion: runtime.Version(),
	}
	got := Get()
	require.Equal(t, exp, got)
}
