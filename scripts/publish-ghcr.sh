#!/usr/bin/env bash

set -euo pipefail

# ==========================
# Configuration
# ==========================
REGISTRY="ghcr.io"
OWNER="AppeiYA"
IMAGE_NAME="hotel_system2"
TAG="${1:-latest}"

IMAGE="${REGISTRY}/${OWNER}/${IMAGE_NAME}:${TAG}"

# ==========================
# Validate Environment
# ==========================
: "${GITHUB_USER:?GITHUB_USER is not set}"
: "${CR_PAT:?CR_PAT is not set}"

if ! command -v docker >/dev/null 2>&1; then
    echo "Docker is not installed."
    exit 1
fi

echo
echo "=========================================="
echo "Publishing Docker image to GHCR"
echo "Image: ${IMAGE}"
echo "=========================================="
echo

# ==========================
# Authenticate
# ==========================
echo "Logging in to ${REGISTRY}..."

echo "${CR_PAT}" | docker login "${REGISTRY}" \
    --username "${GITHUB_USER}" \
    --password-stdin

# ==========================
# Build
# ==========================
echo
echo "Building image..."
docker build -t "${IMAGE}" .

# ==========================
# Push
# ==========================
echo
echo "Pushing image..."
docker push "${IMAGE}"

# ==========================
# Cleanup
# ==========================
docker logout "${REGISTRY}" >/dev/null 2>&1 || true

echo
echo "✅ Successfully published!"
echo "Image: ${IMAGE}"