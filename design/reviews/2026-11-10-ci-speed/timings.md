# CI speed: where ios.yml and android.yml spend their time

Lane `lane/ci-speed`, 2026-11-10. Step timings come from `gh run view <id> --json jobs`, and the
phases inside a step come from that step's log timestamps.

## Before

### ios.yml (`xcode-27` image, Xcode 27.1, iOS 26.5 23F77 simulator)

| run | total | checkout | simulator step | SPM cache restore | build and test |
|---|---|---|---|---|---|
| 42842739728 (main) | 39m15s | 4m40s | 10m17s | <5s (miss, 52s save) | 22m15s |
| 42835770463 (main) | 39m37s | 1m57s | 12m05s | <5s | 24m52s |
| 42829021401 (land/batch-i) | 36m46s | 3m11s | 15m56s | 2m01s (hit) | 15m00s |
| 42817958624 (land/batch-i) | 35m36s | 2m52s | 15m56s | 2m18s (hit) | 11m23s |

Inside run 42842739728:

- **Checkout, 2 to 5 min.** This was a full checkout of a tree with 3.1 GB of design/reviews and
  2.4 GB of data/builds.
- **Simulator, 10 min.** Downloading iOS 26.5 took 1m20s. The first boot took 8m43s, nearly all
  of it data migration (LaunchServicesMigrator).
- **The booted device was lost.** By the time the test step ran, the device was "no longer
  available and booted". A second device was made and booted, which cost 3m20s more.
- **SwiftPM, about 4 min, uncached.** This covers git clones of sentry-cocoa, swift-snapshot-testing,
  swift-syntax, swift-custom-dump and swift-issue-reporting, plus the Sentry binary. The cache is
  1 GB. On a hit it takes about 30s to download and 1m40s to extract, then resolving takes about
  45s. Its key was the whole of project.yml, which every lane edits, so it missed often and then
  saved another 1 GB.
- **Compile, about 7 min.** Nearly all of this is the ForeverSixty app target. Sentry is a binary,
  so it is not compiled.
- **Tests, about 6 min.** There are 862 Swift Testing tests in 98 suites, which take 281s. The rest
  is launch and install. The snapshot suites are the bulk: Rankings 47s, Raids 41s, Home flow
  parity 28s.

### android.yml (`ubuntu-latest`)

| run | total | tests + Roborazzi verify | lint | debug APK | setup-gradle post |
|---|---|---|---|---|---|
| 42842746217 (main) | 9m12s | 4m07s | 2m25s | 39s | 1m03s |
| 42835776664 (main) | 9m23s | 5m28s | 1m51s | 1m03s | 26s |
| 42829026579 (land/batch-i) | 8m46s | 5m14s | 1m47s | 1m04s | n/a |
| 42824007575 (land/batch-i) | 7m58s | 4m46s | 1m40s | 57s | n/a |

- The snapshots are Roborazzi on the JVM (Robolectric), so no emulator is involved.
- The configuration cache and build cache are already on: 98 of the 201 test tasks come FROM-CACHE.
- Each of the three Gradle steps recalculated the task graph ("no cached configuration is available").
- Lint is almost all `app:lintAnalyzeDebug` and `app:lintAnalyzeDebugUnitTest` (1m50s), and it ran
  only after the tests had finished.

## Changes

- iOS:
  - Sparse checkout without design/reviews, and data/builds comes from the cache (data-builds action).
  - The simulator is created first and booted in the background (`tools/ios-ci-simulator.sh
    boot-start`). The app is built for testing at the same time, and the job waits for the boot
    only before `test-without-building`.
  - The SwiftPM cache is keyed on project.yml's `packages:` section and the Xcode version, and is
    saved only by main.
  - Xcode compilation caching (`COMPILATION_CACHE_ENABLE_CACHING`) is on. It is content-addressed,
    so it survives a fresh checkout's new mtimes. It is restored from main and starts afresh each
    month.
  - XcodeGen is cached.
- Android: tests, Roborazzi verify, lint and the debug APK run in one Gradle invocation with
  `--continue`.

## After

(filled in from the lane's dispatch runs)
