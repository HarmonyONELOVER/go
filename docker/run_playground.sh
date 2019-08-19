#!/bin/bash
cd "$(dirname "$0")"

docker build -t alephzero_go_playground -f playground.Dockerfile .

docker run                              \
  --rm                                  \
  -it                                   \
  -v "$(realpath ..)":/go/src/alephzero \
  -p 3000:3000                          \
  alephzero_go_playground