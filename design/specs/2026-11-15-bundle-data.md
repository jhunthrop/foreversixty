# Bundling Forever Sixty Data into the main zip

Proposal only, 2026-11-15. Nothing here is built.

## The question

Should the ForeverSixty zip carry the ForeverSixtyData folder too, the way some large addons ship
several folders in one download, so a player installs one thing and has the ratings?

## Where things stand

- `addon/ForeverSixty` is CurseForge project 1706689 and Wago XKqEdoKy. Releases come from
  `addon-v*` tags through `.github/workflows/addon-release.yml` and the BigWigs packager. The zip
  already holds two top-level folders: `.pkgmeta`'s `move-folders` lifts `ForeverSixtyInbox/` out of
  the packager's topdir, and the workflow checks for exactly those two folders.
- `addon/ForeverSixtyData` is a separate project, CurseForge 1706707 and Wago ANz7R564.
  `addon-data-release.yml` runs at 05:00 UTC every night. It downloads the `Data.lua` files the
  `data-addon` Cloud Run job wrote to the bucket (GCP auth through Workload Identity), builds a zip
  with its own throwaway `.pkgmeta`, and publishes `data-v<date>`. The checked-in `Data.lua` is only
  the empty shape.
- The main addon lists it as `## OptionalDeps: ForeverSixtyData` and reads the global
  `ForeverSixtyData` (Score.lua, ChatBadges.lua). It marks the file stale after 3 days
  (`Score.STALE_AFTER_SECONDS`) and says when the pack is missing, broken, or newer than it reads.

## What bundling would take in the packager

1. In addon-release.yml's "give the packager its own checkout" step, add the data folder beside the
   inbox: `git archive HEAD -- addon/ForeverSixtyData | tar -x -C "$pkgroot" --strip-components=1`.
   Then overwrite its `Data.lua` with the night's file from the bucket. The main release job would
   need the same `google-github-actions/auth` step and `id-token: write` as the data job. Without that
   step the bundle only ships the empty shape, which is worse than nothing because it hides the
   "not installed" prompt.
2. In `.pkgmeta`, add `ForeverSixty/ForeverSixtyData: ForeverSixtyData` to `move-folders`, and make
   the zip check expect three folders.
3. The data TOC's `## Version: @project-version@` would take the main addon's tag (`2.0.56`) instead
   of the date, so Score.lua's dated `generated` field becomes the only freshness signal.

## The cost: freshness

The data changes every night. The main addon ships when there is something to ship, sometimes
days apart. Bundled data is as old as the last main release, so it goes stale three days after any
quiet spell. There are two ways around that, and both are bad:

- **Republish the main addon every night with the new data.** Every player gets a full main-addon
  update every day, the CurseForge and Wago apps show a new version daily, and the changelog fills
  with empty entries. CHANGELOG.md has player-facing lines by rule, so this needs a nightly heading
  with no lines, which breaks the "the changelog names a real version" gate. It also ties every
  data night to the main release pipeline (tests, the data-check job, signing).
- **Keep both projects and bundle a snapshot.** Two projects would ship a folder with the same
  name. A WoW AddOns directory can hold only one `ForeverSixtyData` folder, so the last install wins.
  The CurseForge app tracks folders by fingerprint against one owning project, so it reports the
  folder as modified or unknown, and updating either project overwrites the other's files. Players
  who have both would see their data go back to the main release's snapshot whenever the main addon
  updates.

## CurseForge and Wago relations

- CurseForge's file relations include `include` (another project's files are packed inside this
  one), `embeddedLibrary`, `optionalDependency`, `requiredDependency`, `tool` and `incompatible`.
  Bundling honestly means marking the main file `include` of project 1706707. CurseForge treats
  that as "you already have it", so the app stops offering the data project separately. That is only
  correct if the bundled copy is as fresh as the standalone one, which, as above, it is not.
- The packager writes relations from `.pkgmeta` (`required-dependencies`, `optional-dependencies`,
  `embedded-libraries`, `tools-used`, `incompatible`). Today the main `.pkgmeta` declares none, so
  CurseForge does not know the two projects are related, even though the TOC names OptionalDeps.
- Wago reads the same `.pkgmeta` relations from the packager. It has no separate "include" kind, so
  a bundled folder looks to Wago like part of the main addon, and the standalone data project would
  conflict on install.

## Recommendation

Do not bundle. Get most of the benefit with no folder clash:

1. In `addon/ForeverSixty/.pkgmeta`, add an `optional-dependencies` entry with the CurseForge slug
   of project 1706707 (look it up on the project page; the main project's is `foreversixty`). Wago
   reads the same file. The CurseForge app and
   Wago then offer the data pack when the main addon is installed, and both projects keep their own
   cadence: the main addon on release, the data every night.
2. Keep the in-game prompt that already shows when the pack is missing (Score.status "not
   installed"), and link the data project from /setup's addon step next to the main download.
3. Check that the CurseForge slug is right before the first release that carries the relation. A
   wrong slug fails the upload (the packager stops with an API error), and does not produce a silent
   bad relation.

If one download is still wanted later, the only clean form is to retire project 1706707: mark its
last file as superseded, publish the main addon nightly as described above, and accept the daily
update noise. That needs a decision on the daily changelog first.
