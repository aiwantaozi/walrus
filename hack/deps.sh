#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"
source "${ROOT_DIR}/hack/lib/init.sh"

function mod() {
  local path="$1"

  [[ "${path}" == "${ROOT_DIR}" ]] || pushd "${path}" >/dev/null 2>&1

  go mod tidy
  go mod download

  [[ "${path}" == "${ROOT_DIR}" ]] || popd >/dev/null 2>&1
}

function dispatch() {
  local target="$1"
  local path="$2"

  shift 2
  local specified_targets="$*"
  if [[ -n ${specified_targets} ]] && [[ ! ${specified_targets} =~ ${target} ]]; then
    return
  fi

  seal::log::debug "modding ${target}"
  if [[ "${PARALLELIZE:-true}" == "false" ]]; then
    mod "${path}"
  else
    mod "${path}" &
  fi
}

#
# main
#

seal::log::info "+++ MOD +++"

dispatch "utils" "${ROOT_DIR}/staging/utils" "$@"
dispatch "walrus" "${ROOT_DIR}" "$@"

if [[ "${PARALLELIZE:-true}" == "true" ]]; then
  seal::util::wait_jobs || seal::log::fatal "--- MOD ---"
fi
seal::log::info "--- MOD ---"
