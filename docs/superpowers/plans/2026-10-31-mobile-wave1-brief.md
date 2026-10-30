# Mobile Wave 1: shared brief for every feature lane

Each lane builds ONE feature on BOTH platforms, iOS (SwiftUI) and Android (Compose), so the two stay in parity.

## Read first
- `docs/superpowers/plans/2026-10-30-mobile-apps-lanes.md`: the plan, its decisions, the lifecycle catalogue and Wave 1.
- `design/specs/2026-10-30-mobile-apps.md`: the binding spec. It is the "SHIP" revision; your feature's section, its states and its copy keys are what you build.
- `design/mocks/2026-10-30-mobile-apps/renders/`: your feature's boards, on both platforms.
- The reports for P1 (`docs/superpowers/reports/2026-10-30-mobile-p1.md`) and P2 (`…-mobile-p2.md`): the generated models, the copy catalog, the link table, AppConfig and the telemetry already on main.
- `docs/contracts/app-v1.md` and `api/openapi.yaml`: the contract. Build against the fixtures. The live API for sign-in, push and home is dark until after Nov 6.

## Rules
- **One source of truth.**
  - Models and endpoints come from `api/openapi.yaml` through the generators. Never hand-write a model.
  - Strings live in `apps/shared/copy/en.json`, under your feature's key prefix. Never inline a string.
  - Deep links live in `apps/shared/links.json`.
  - If you need a spec change, edit openapi.yaml plus a fixture, regenerate, and keep the contract tests green.
- **Shared files.** Other feature lanes run at the same time.
  - Touch shared files (the app graph, tab and navigation setup, en.json, links.json, openapi.yaml) only additively, in your own block, and keep the edits small.
  - Put the feature's screens in their own module or folder on each platform.
  - Gate the feature behind its AppConfig feature key, which defaults to on in debug and in fixture builds.
- **Parity.** One test list, implemented on both platforms. Screenshot tests on both, at the default size and the largest text size, in the spec's states.
- **House rules.**
  - Never cd outside your worktree, never stash, never run a broad pkill, never dispatch subagents.
  - Commit per platform and per step with `git commit -q -F <file>` as its own command, with the trailer `Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>`. Do NOT push.
  - core.fileMode is off: any new `.sh` file needs `git update-index --chmod=+x`.
- **Verify, with exit codes checked and never through a pipe:**
  - `make -C apps/ios test`
  - in `apps/android`: `./gradlew test :app:verifyRoborazziDebug lint assembleDebug`, with `JAVA_HOME=/opt/homebrew/opt/openjdk@21/libexec/openjdk.jdk/Contents/Home`
  - in web, `npm run check`, `lint` and `format:check`, whenever you touched anything under web/ (the drift tests live there, and CI's type check has failed twice on them)
  - the codegen drift checks, if you touched a source: `npx vitest run src/lib/mobile src/styles/mobile-tokens.test.ts` in web, with web/node_modules symlinked from `/Users/jh/code/forever/.claude/worktrees/lane-discord-bot/web/node_modules` and removed before committing
  - the Go contract tests, if you touched openapi.yaml or fixtures, on a private database you create and then drop
- **Captures.** Simulator and emulator screenshots of every screen and state, both platforms, saved to `design/reviews/2026-10-31-mobile-<feature>/` as `<screen>-<state>-ios.png` and `-and.png`. The two reviewers judge from these.
- **Report.** Write it to `docs/superpowers/reports/2026-10-31-mobile-<feature>.md`. Return only STATUS, commits, a test summary and concerns.
