#!/usr/bin/env bash

set -eo pipefail

# parse optional flag
if [ "$1" = "--optional" ]; then
  OPTIONAL=true
  echo "Optional flag set, building optional images"
else
  OPTIONAL=false
fi

IMAGES=(
  "ingress-gate"
  "nanolimbo"
)

if $OPTIONAL; then
  IMAGES+=(
  )
fi

for image in "${IMAGES[@]}"; do
  cd $image
  docker buildx build -t prismcloud/$image:latest .
  cd ..
done

# check if k3s is installed
if ! command -v k3s >/dev/null 2>&1; then
  echo "k3s is not installed, not pushing images to cluster"
  exit 1
fi

for image in "${IMAGES[@]}"; do
  docker save prismcloud/$image:latest | sudo k3s ctr images import -
done

echo "Done!"