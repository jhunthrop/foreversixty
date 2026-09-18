module github.com/jhunthrop/foreversixty/sim

go 1.25.11

require github.com/jhunthrop/foreversixty/logs v0.0.0

replace github.com/jhunthrop/foreversixty/logs => ../logs

// Development pin. CI REWRITES this line (it never deletes it) to
//   replace github.com/wowsims/classic => github.com/jhunthrop/wowsims-forever <pseudo-version>
// at the sha in sim/enginever/version.go. See .github/workflows/sim.yml,
// written by Task 13. Deleting it instead would leave the require above
// pointing at v0.0.0-00010101000000-000000000000, which resolves to
// nothing.
replace github.com/wowsims/classic => /Users/jh/code/wowsims-forever
