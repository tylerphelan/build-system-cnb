/*
 * Copyright 2018-2019 the original author or authors.
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

package runner

import (
	"fmt"
	"path/filepath"

	"github.com/buildpack/libbuildpack/application"
	"github.com/cloudfoundry/build-system-cnb/buildsystem"
	"github.com/cloudfoundry/libcfbuildpack/build"
)

// NewRunner creates a new Maven Runner instance.
func NewMavenRunner(build build.Build, buildSystem buildsystem.BuildSystem) Runner {
	return NewRunner(build, mavenBuiltArtifactProvider, buildSystem.Executable(), "-Dmaven.test.skip=true", "package")
}

func mavenBuiltArtifactProvider(application application.Application) (string, error) {
	target := filepath.Join(application.Root, "target")

	candidates, err := filepath.Glob(filepath.Join(target, "*.jar"))
	if err != nil {
		return "", err
	}

	if len(candidates) != 1 {
		return "", fmt.Errorf("unable to determine built artifact in %s, candidates: %s", target, candidates)
	}

	return candidates[0], nil
}
