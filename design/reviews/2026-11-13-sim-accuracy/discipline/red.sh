#!/usr/bin/env bash
# red.sh: the lane's guard test alone, from sim/ on the box (shows red before the card change).
cd "$(dirname "$0")/../../../../sim" && go test ./request/ -run TestPriestTankFlashHealWaitsOnTheManaPace -count=1
