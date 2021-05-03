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
	"strings"
)

var (
	version    = "v0"
	preRelease = ""
	metadata   = ""
	gitCommit  = ""
)

// BuildInfo describes the compile time information.
type BuildInfo struct {
	// Version is the current semver.
	Version string `json:"version,omitempty"`
	// GitCommit is the git sha1.
	GitCommit string `json:"git_commit,omitempty"`
	// Metadata is arbitrary metadata provided from the builder
	Metadata string `json:"metadata,omitempty"`
	// GoVersion is the version of the Go compiler used.
	GoVersion string `json:"go_version,omitempty"`
}

// GetVersion returns the semver string of the version
func GetVersion() string {
	v := version
	if preRelease != "" {
		v += "-" + preRelease
	}
	var buildIdentifiers []string
	if gitCommit != "" {
		buildIdentifiers = append(buildIdentifiers, gitCommit)
	}
	if metadata != "" {
		buildIdentifiers = append(buildIdentifiers, metadata)
	}
	if len(buildIdentifiers) > 0 {
		v += "+" + strings.Join(buildIdentifiers, ".")
	}
	return v
}

// Get returns build info
func Get() BuildInfo {
	return BuildInfo{
		Version:   GetVersion(),
		GitCommit: gitCommit,
		Metadata:  metadata,
		GoVersion: runtime.Version(),
	}
}
