module github.com/jhunthrop/foreversixty/sim

go 1.25.11

require (
	github.com/jhunthrop/foreversixty/logs v0.0.0
	google.golang.org/protobuf v1.36.6
)

require github.com/wowsims/classic v1.0.14

require golang.org/x/exp v0.0.0-20250620022241-b7579e27df2b // indirect

replace github.com/jhunthrop/foreversixty/logs => ../logs

// Development pin. CI REWRITES this line (it never deletes it) to
//   replace github.com/wowsims/classic => github.com/jhunthrop/wowsims-forever <pseudo-version>
// at the sha in sim/enginever/version.go. See .github/workflows/sim.yml,
// written by Task 13.
//
// sim/request and sim/adapter import the engine, so the `require
// github.com/wowsims/classic` above is real and deleting this replace
// instead of rewriting it would leave it unresolved.
replace github.com/wowsims/classic => /Users/jh/code/wowsims-forever
