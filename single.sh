#!/bin/bash

set -e

tree="$1"
step="$2"

./bin/ptool -type=memory -plots=$tree=./_assets/mem-$tree-$step.out -step=$step

./bin/ptool -type=time -plots=$tree=./_assets/time-$tree-$step.out -step=$step
