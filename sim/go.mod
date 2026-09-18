module github.com/jhunthrop/foreversixty/sim

go 1.25.11

require github.com/jhunthrop/foreversixty/logs v0.0.0

replace github.com/jhunthrop/foreversixty/logs => ../logs

// Development pin. CI REWRITES this line (it never deletes it) to
//   replace github.com/wowsims/classic => github.com/jhunthrop/wowsims-forever <pseudo-version>
// at the sha in sim/enginever/version.go. See .github/workflows/sim.yml,
// written by Task 13.
//
// There is no matching `require github.com/wowsims/classic ...` above:
// nothing in this module imports the engine yet, so `go mod tidy` removed
// it. Task 3's sim/adapter package imports the engine for real, which
// brings the require back — at that point deleting this replace instead
// of rewriting it would leave that require unresolved.
replace github.com/wowsims/classic => /Users/jh/code/wowsims-forever
