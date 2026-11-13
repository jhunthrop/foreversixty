set -u
make -C .. simdb >/dev/null || exit 1
FOREVER_UPDATE_GOLDEN=1 GOMAXPROCS=5 nice -n 15 go test -count=1 -run 'TestRotationLadder' ./request/ 2>&1 | tail -3
git status --short
echo "=====GOLDEN-DIFF"
git diff -- request/testdata/ladder/ | cat
echo "=====END"
GOMAXPROCS=5 nice -n 15 go vet ./... 2>&1 | tail -3
GOMAXPROCS=5 nice -n 15 go test -count=1 ./request/ ./specs/ ./leveling/ ./rankchoices/ ./adapter/ ./cmd/spec-breakdown/ 2>&1 | tail -8
