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
	"fmt"
	"runtime"
	"strings"
)

var (
	name      = ""
	version   = "v0.0.0"
	metadata  = ""
	gitCommit = ""
)

// AppInfo holds all information about the application
type AppInfo struct {
	// Name of the application
	Name string `json:"name,omitempty"`
	// Version is the current semver.
	Version string `json:"version,omitempty"`

	Build BuildInfo
}

// BuildInfo describes the compile time information.
type BuildInfo struct {
	// GitCommit is the git sha1.
	GitCommit string `json:"gitCommit,omitempty"`
	// Metadata is arbitrary metadata provided from the builder
	Metadata string `json:"metadata,omitempty"`
	// GoVersion is the version of the Go compiler used.
	GoVersion string `json:"goVersion,omitempty"`
}

// Get returns the AppInfo
func Get() AppInfo {
	return AppInfo{
		Name:    name,
		Version: version,
		Build:   Build(),
	}
}

// Name returns the application name
func Name() string {
	return name
}

// Version returns the application version
func Version() string {
	return version
}

// UserAgent returns a standard user agent string in the format "name/version"
func UserAgent() string {
	return fmt.Sprintf("%s/%s", name, Version())
}

func (info AppInfo) String() string {
	v := info.Version
	var buildIdentifiers []string
	if info.Build.GitCommit != "" {
		buildIdentifiers = append(buildIdentifiers, info.Build.GitCommit)
	}
	if info.Build.Metadata != "" {
		buildIdentifiers = append(buildIdentifiers, info.Build.Metadata)
	}
	if len(buildIdentifiers) > 0 {
		v += "+" + strings.Join(buildIdentifiers, ".")
	}
	return fmt.Sprintf("%s/%s GoVersion: %s", info.Name, v, info.Build.GoVersion)
}

// Build returns the BuildInfo
func Build() BuildInfo {
	return BuildInfo{
		GitCommit: gitCommit,
		Metadata:  metadata,
		GoVersion: runtime.Version(),
	}
}
