set -u
GOMAXPROCS=5 nice -n 15 go run ./cmd/addon-builds -repo-root .. 2>&1 | tail -2
echo "=====DIFF"
git diff -- ../data/ | cat
echo "=====END"
GOMAXPROCS=5 nice -n 15 go test -count=1 ./cmd/addon-builds/ 2>&1 | tail -3
