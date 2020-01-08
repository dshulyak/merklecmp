#!/bin/bash

set -e

step="$1"

./bin/ptool -type=memory -plots=iavl=./_assets/mem-iavlTree-$step.out,patricia=_assets/mem-patricia-$step.out,urkel=_assets/mem-urkel-$step.out -step=$step

./bin/ptool -type=time -plots=iavl=./_assets/time-iavlTree-$step.out,patricia=_assets/time-patricia-$step.out,urkel=_assets/time-urkel-$step.out -step=$step
