// Package activity is the site feed behind home B (design/specs/2026-11-07-home-b.md): the
// activity_events table, the writers that append to it from the paths where things happen
// (a report completes, a guild posts a raid, the hourly board job, the BiS startup diff, a
// guild member's sim finishes), the weekly tier and board snapshots, and the two reads home
// makes: GET /v1/feed and GET /v1/home.
//
// Privacy: a row with a guild id is its guild's only and is read only by the guild's verified
// members; private and flagged reports never write a row; an anonymised account's character
// is never named.
package activity
