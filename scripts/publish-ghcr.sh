#!/usr/bin/env bash
set -e

# Configuration
REGISTRY="ghcr.io"
REPO="appeiya/hotel_system2"
TAG="${1:-latest}"
IMAGE_FULL="${REGISTRY}/${REPO}:${TAG}"

echo "=========================================="
echo " Publishing image to GHCR"
echo " Image target: ${IMAGE_FULL}"
echo "=========================================="

if [ -z "$CR_PAT" ]; then
  echo "Error: CR_PAT environment variable is not set."
  echo "Please export your GitHub Personal Access Token with 'write:packages' scope:"
  echo "  export CR_PAT=your_github_pat_token"
  echo "  export GITHUB_USER=your_github_username"
  exit 1
fi

if [ -z "$GITHUB_USER" ]; then
  echo "Error: GITHUB_USER environment variable is not set."
  echo "Please set GITHUB_USER=your_github_username"
  exit 1
fi

echo "Logging in to ${REGISTRY}..."
echo "$CR_PAT" | docker login "${REGISTRY}" -u "$GITHUB_USER" --password-stdin

echo "Building Docker image..."
docker build -t "${IMAGE_FULL}" .

echo "Pushing image to ${REGISTRY}..."
docker push "${IMAGE_FULL}"

echo "Successfully built and pushed ${IMAGE_FULL}"
