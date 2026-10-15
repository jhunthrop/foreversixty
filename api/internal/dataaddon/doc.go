// api/internal/dataaddon/doc.go
// Package dataaddon is the nightly job that writes the Forever Sixty Data addon's
// ForeverSixtyData global in contract 5's format 2 (docs/contracts/addon-v2.md; the
// fixture docs/contracts/fixtures/data-addon-v2.lua is its golden file): a record for
// every character with a public rated fight in the last 90 days, each the character's
// rating.ScoreV2 (the formula's only home; this package copies its figures and never
// works one out), a no-logs record for every verified guild member without one, and a
// record for every guild with at least one verified member. It renders Data.lua (or
// Data-<region>.lua, when the combined file would exceed 4 MiB) for the
// addon-data-release.yml workflow to package and publish. The addon's Score module
// (addon/ForeverSixty/Score.lua) is the reader.
package dataaddon

import "time"

// JobCommand is this lane's Cloud Run job name, dispatched the same way
// rating.BackfillJobCommand already is (api/cmd/api/main.go's os.Args[1]
// switch).
const JobCommand = "data-addon"

// ratingWindow is "the last 90 days": a character is in the file when it has a rated
// fight in it, and its best parses and its guild's nights and kills are read from it.
const ratingWindow = 90 * 24 * time.Hour

// maxSingleFileBytes is the split threshold: a rendered Data.lua at or under this size
// ships as one file; larger, it ships as one file per region instead (see write.go's
// Render).
const maxSingleFileBytes = 4 << 20 // 4 MiB
