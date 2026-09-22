// api/internal/dataaddon/doc.go
// Package dataaddon is the nightly job that writes the Forever Sixty Data
// addon's ForeverSixtyData global: every character with at least one public
// rated fight in the last 90 days, and every guild with at least one
// verified member, serialized to Data.lua (or Data-<region>.lua, when the
// combined file would exceed 4 MiB) for the addon-data-release.yml workflow
// to package and publish. addon/ForeverSixty/Ratings.lua is the reader this
// package's output must satisfy exactly; see
// docs/superpowers/plans/2026-09-21-data-addon.md for the key-format
// reconciliation and aggregation rule this package implements.
package dataaddon

import "time"

// JobCommand is this lane's Cloud Run job name, dispatched the same way
// rating.BackfillJobCommand already is (api/cmd/api/main.go's os.Args[1]
// switch).
const JobCommand = "data-addon"

// ratingWindow is "the last 90 days" the dispatch's aggregation rule names,
// for both a character's rating and a guild's night count.
const ratingWindow = 90 * 24 * time.Hour

// maxSingleFileBytes is the dispatch's own split threshold: a rendered
// Data.lua at or under this size ships as one file; larger, it ships as
// one file per region instead (see write.go's Render).
const maxSingleFileBytes = 4 << 20 // 4 MiB
