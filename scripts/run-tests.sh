#!/usr/bin/env bash

# Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
# This program and the accompanying materials are made available under the terms of the under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# http://www.apache.org/licenses/LICENSE-2.0
# Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

set -eu
set -o pipefail

if [[ $# -eq 0 ]]; then
  LIFECYCLE_TESTS_CONFIG=<(echo "[{}]") GO111MODULE=on GOFLAGS="-mod=vendor" go run github.com/onsi/ginkgo/ginkgo -r -dryRun system_tests
fi

GO111MODULE=on GODEBUG="tls13=0" GOFLAGS="-mod=vendor" go run github.com/onsi/ginkgo/ginkgo \
  -randomizeSuites=true \
  -randomizeAllSpecs=true \
  -keepGoing=true \
  -r \
  -cover \
  -trace \
  -race \
  "$@"
