#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"
source "${ROOT_DIR}/hack/lib/init.sh"

BUILD_DIR="${ROOT_DIR}/.dist/build"
mkdir -p "${BUILD_DIR}"

#
# util
#

function build() {
  local target="$1"
  local task="$2"
  local path="$3"
  shift 3

  local ldflags=(
    "-X main.Name=${target}-${task}"                                # Inject the binary name.
    "-X github.com/seal-io/utils/version.Version=${GIT_VERSION}"    # Inject the git version.
    "-X github.com/seal-io/utils/version.GitCommit=${GIT_COMMIT}"   # Inject the git commit.
    "-X github.com/seal-io/walrus/server/telemetry.APIKey=${WALRUS_TELEMETRY_API_KEY:-}"
    "-w -s"
    "-extldflags '-static'"
  )

  local tags=()
  IFS=" " read -r -a tags <<<"$(seal::target::build_tags "${target}")"

  local platforms=()
  IFS=" " read -r -a platforms <<<"$(seal::target::build_platforms "${target}" "${task}")"

  for platform in "${platforms[@]}"; do
    local os_arch
    IFS="/" read -r -a os_arch <<<"${platform}"
    local os="${os_arch[0]}"
    local arch="${os_arch[1]}"
    local output
    output="${BUILD_DIR}/${target}/${target}-${task}-${os}-${arch}"
    if [[ "${os}" == "windows" ]]; then
      output="${output}.exe"
    fi
    GOOS=${os} GOARCH=${arch} CGO_ENABLED=0 go build \
      -trimpath \
      -ldflags="${ldflags[*]}" \
      -tags="${os} ${tags[*]}" \
      -o="${output}" \
      "${path}"
  done
}

#
# lifecycle
#

function dispatch() {
  local target="$1"
  local path="$2"
  shift 2

  local tasks=()
  IFS=" " read -r -a tasks <<<"$(seal::util::find_subdirs "${path}/cmd")"

  for task in "${tasks[@]}"; do
    seal::log::debug "building ${target} ${task}"
    if [[ "${PARALLELIZE:-true}" == "false" ]]; then
      build "${target}" "${task}" "${path}/cmd/${task}"
    else
      build "${target}" "${task}" "${path}/cmd/${task}" &
    fi
  done
}

#
# main
#

seal::log::info "+++ BUILD +++" "info: ${GIT_VERSION},${GIT_COMMIT:0:7},${GIT_TREE_STATE},${BUILD_DATE}"

dispatch "walrus" "${ROOT_DIR}" "$@"

if [[ "${PARALLELIZE:-true}" == "true" ]]; then
  seal::util::wait_jobs || seal::log::fatal "--- BUILD ---"
fi
seal::log::info "--- BUILD ---"
