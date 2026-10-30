# Launch-day client move

Forever goes live on 4 November 2026. The live client will be a new build. The client watcher moves the site, the addon and the guides to it without anyone at the keyboard: `.github/workflows/client.yml` at 03:41 UTC, then `tools/activate-build.sh`. This page covers what to watch on the nights of 3 and 4 November, how to finish a move that stopped, how to move by hand when the live client lands somewhere the watcher does not look, and how to roll back.

It is built from the manual checklist in the 2026-10-30 rehearsal (`design/reviews/2026-10-30-launch-rehearsal.md`). Engine follow and pin are covered in `docs/engine-follow.md`.

## The unattended path

| UTC | Workflow | What it does |
|---|---|---|
| 03:41 | `client` | **watch** asks wago.tools for the newest Forever build (`pipeline client-check`, ruled by `data/curated/client.json`). When it is newer than `web/src/data/active-build.json`, **fetch** runs `data.yml`'s fetch for it and **activate** runs `tools/activate-build.sh`. |
| (activate) | | Carries the picks forward, then moves the site, the guides and the addon. Runs the `--check` gates, then the **test gate** (`tools/activation-gate.sh`: `npm run check`, `lint`, `format:check`, the real-data `CF_PAGES=1 npm run build` that prerenders every guide, BiS and planner page from the new build, and the whole vitest suite; beside them, the API packages that read the active build). The gate takes about 4 minutes on a laptop. Only then does it commit, push and tag `addon-v<version>`, and dispatch web, api, sim, addon-release and engine. |
| (activate) | | **The ranking is not dispatched** while the pinned engine is stamped for the old build. The summary says `ranking waits for engine follow: …`. This is the normal state on move day. |
| 05:00 | `addon-data-release` | Nightly addon data. |
| about 04:00, and 05:07 | `engine` | On a move night the activation dispatches it with force right after the push, so the follow starts at about 04:00; the 05:07 schedule then queues behind it in the `engine` concurrency group, and usually skips. Do not wait for 05:07. It follows the new build in the fork. A clean follow pins. A review or red follow opens `engine: follow client <build>` in the fork, or `engine: pin <sha>` on the site. |
| (pin) | | A green pin pushes to main and dispatches web, api, sim and **bis with force**. That is what ranks the new build. |
| 06:23 | `bis` | Nightly ranking. It ranks only when the pinned engine's `FOREVER_BUILD` is the active build; otherwise it writes the same "waits" line and keeps the carried picks. The first ranking on a matching engine also re-measures `dps-per-point.json`. The carry is marked by `dps-per-point.carried-from`, which the activation writes and the re-measure removes. A failed re-measure keeps the carry and the picks still land. On the second night of waiting, or if the engine stamp cannot be read, the run fails red: `ranking stalled on the engine follow`. A ranking redeploys both web and api. |
| 07:15 | `web` | Nightly Lighthouse. |

Every failure in `client` ends with a block in the run summary:

- **Step**: where it stopped.
- **Reason**: the step's output, or the failing tests.
- **To finish or redo the move**: the exact hand command.

The same text appears as an error annotation on the run page. Read that block first; this page explains the commands it gives.

## Night of 3 November (the eve)

1. **Check that main is green.** `web`, `api` and `data` should all pass on main's head. A red main means the gate goes red on move day for a reason that has nothing to do with the client.
2. **Check that main is not half-moved.** `data/builds/<newest>/` exists and the active build is one behind; that is fine. Look for open `engine: follow client …` (fork) or `engine: pin …` (site) pull requests. Each one is hand work that will hold the ranking after the move.
3. **Check that the guides carry.** Run the move locally against the newest fetched build without committing:

   ```
   tools/activate-build.sh <newest build> --no-commit --no-gate
   ```

   A clean pass means the trees carry. A stop at **re-encoding the guide builds** lists the builds a person has to re-spend; do that tonight, not at 03:41. Undo the local run with the `git restore … && git clean …` line it prints.
4. **Run the watcher by hand once:** `gh workflow run client.yml -f inspect=true`. If a newer build is listed, this fetches its tables to main, the same as the nightly. It then stops before the move, after the gate. Read its warnings. `unclaimed newer client builds`, `Forever build numbered below the active one` and `Blizzard shipped a client wago.tools has not exported` each mean the next section applies.

## Night of 4 November (launch)

Blizzard serves the live build the hour it ships. wago.tools exports its tables later, a day or more, and the fetch needs those exports. So the first night may move nothing and only warn `wago behind`. That is expected.

Each morning, open the `client` run and check:

- **No warnings, "the active build … is the newest"**: nothing new. Done.
- **`new Forever client` notice, activate green**: the move landed. Then:
  - check `addon-release` published `addon-v<version>` to CurseForge and Wago;
  - check `web` deployed;
  - watch for the `engine` run. A clean follow and green pin rank the same night. Otherwise work the pull request it opens (`docs/engine-follow.md`), and the ranking follows the pin.
  - check `bis` ranked (`chore(bis): nightly leveling BiS for <build>`) and re-measured dps-per-point once the pin landed. Until then the carried picks are live. A carried slot whose item the new client lost reads "no known source"; that is the documented gap until the ranking.
- **activate red**: read the summary block and follow its command. The common ones are covered below.
- **fetch red**: the report job names it. The fetch's own log names the gate it tripped. A shrunken catalogue is the usual one: 70685 lost 1.2% of its items, inside the 2% gate. A truncated ItemSparse export is the other.
- **watch red**: wago.tools or Blizzard did not answer. Nothing moved. The next morning retries, or dispatch `client` once it answers.

Also check the items shrink after any move. The carry prints `carry-bis: … is not in <build>; dropped` once per pick it drops. Many drops mean the new items.json is missing items the ranking will need. Compare the item counts in the two builds' `manifest.json`.

## Finishing a move that stopped

Every command below is what the summary prints, with `<build>` and `<product>` filled in.

**A guide build does not carry** (step: re-encoding the guide builds). A talent was removed, renamed or re-ranked. In 70685 and 70707 that was Arcane Focus, Improved Fireball, Elemental Precision and Critical Mass. Re-spend each named build in its guide (`web/src/content/guides/<class>/<spec>.md`), codes stamped `FS1:<build>`, and fix the prose that names the talent. Then either:

- **Preferred:** from a clean main checkout with the re-spent guides uncommitted, run `tools/activate-build.sh <build>`. It commits the guides with the move and runs the same gate. Then it pushes, tags and dispatches, using your `gh`.
- Land the re-spent guides on main and dispatch `gh workflow run client.yml -f build=<build> -f product=<product>`. Until the move lands, main's guide tests are red (their codes name a build that is not active), so prefer the first option.

**The test gate is red** (step: the test gate). Nothing was committed, pushed or tagged. The reason lists each failure as `web <file> > <test>`, `web (npm run <script>) exited <n>` (check, lint, format or the real build, with its log tail) or `api <package> <Test>`. A red `npm run build` is a page that does not prerender from the new build; its log names the page. On 70707 the frost guide's hit-line test fails on purpose: Precision is not in `SCHOOL_HIT_TALENTS`. Add it there, or add the build to `HIT_LINE_FALLBACK_BUILDS` in `web/src/pages/guides/_spec.test.ts`, on main before the move. A test that pins a name the client renamed should read the name from the build's data, the way `web/src/pages/guides/_spec.test.ts`'s hit line does. Land that fix on main, then dispatch client again.

**Coach rules** (step: the log coach's rule files). A spell the coach rules name is gone from the new tables. Edit `data/curated/coach/<spec>.json` for the new spell, land it, and dispatch again.

**The push failed** after five rebases. Nothing was pushed or tagged. Dispatch again.

**The tag or a dispatch failed.** The move is on main. Dispatch `client` again with the same build: `activate-build.sh` finds its own move commit, sees that `addon-v<version>` is not on origin (or was never released), and resumes. It tags the move's commit and dispatches. The nightly does the same on its own when the TOC version has no tag. A run with nothing left says `nothing to do`. If only some dispatches failed after the release went out, run the `gh workflow run …` lines the summary prints. A move that timed out or was cancelled mid-step gets a summary naming the step it was on; a rerun resumes it if it got past the push.

**The ranking never runs.** The summary said `ranking waits for engine follow`, and no green pin followed. From the second night, `bis` fails red with `ranking stalled on the engine follow`. Open the fork's `engine: follow client <build>` pull request, or the site's `engine: pin <sha>` one, and finish it (`docs/engine-follow.md`). A green pin, or the auto-merge of a hand-fixed pin pull request, dispatches `bis` with force. Never force a ranking by hand on an engine stamped for another build: `bis.yml` refuses it, and with reason. That engine reads the new client's talent strings by the old trees' indexes.

## Moving by hand to a build the watcher will not take

The watcher activates the newest build under a product in `client.json`'s `products`, or a 1.x at or above `min_version` under any key it does not ignore. It never activates:

- **A new product key carrying 2.0 or higher.** The version rule covers only 1.x. It warns `unclaimed newer client builds` and names the key.
- **A renumbering down** (live as 1.16.0, say). It warns `Forever build numbered below the active one`.
- **A PTR or test key** (`test_product_pattern`: `ptr`, `test`, `wowt`-shaped). A Forever PTR is numbered like live and must never move the site. It is not warned about.
- **A 1.60+ build under an ignored key** (Era's own, say). It warns as unclaimed.

To move to one of these:

1. Confirm it is the live client: the launcher's build, or Blizzard's `us.version.battle.net/<key>/versions`.
2. For a new key that will carry future builds, add it to `products` in `data/curated/client.json` and land it. The next nightly then follows it on its own. (`load_config` refuses a listed key that matches the PTR pattern.)
3. Dispatch with the build and the key:

   ```
   gh workflow run client.yml -f build=<build> -f product=<key>
   ```

   The fetch runs for that key, then the move.

A 2.0 renumber needs nothing else in the script: the addon TOCs take `## Interface: 20000`, and the summary's Interface line shows it. For a renumbering down, check the new number reads as intended everywhere a version is compared: the addon version bump is independent of the client number. Then dispatch the same way.

## Rolling back to the previous build

Roll back when the new build's data is wrong in a way that cannot be fixed forward the same day. Examples: a truncated fetch landed, or the live client turned out to be a different build.

**Preferred: move forward to the old build.** `activate-build.sh` moves to any build whose tables are committed, not only a newer one:

```
tools/activate-build.sh <previous build>
```

Run it locally from a clean main. The old build's directory already has its own `bis/` and `dps-per-point.json`, so nothing is carried. The guides re-encode back by talent name. The addon gets a new patch version, so CurseForge and Wago see an upgrade rather than a downgrade, and is tagged and released. The site deploys. The ranking waits until the pinned engine is stamped for the old build again.

If the guides were re-spent for talents the old trees do not have, the re-encode stops and names them. Then use the revert below.

**Fallback: revert the activation commit.**

```
git revert <sha of "chore(client): activate build <build>">
```

Then bump `## Version:` in `addon/ForeverSixty/ForeverSixty.toc` past the released `addon-v<version>`, add a CHANGELOG line, push, tag `addon-v<new version>` and dispatch:

```
gh workflow run addon-release.yml --ref addon-v<new version> -f dry_run=false -f publish=true
gh workflow run web.yml --ref main
gh workflow run api.yml --ref main
gh workflow run sim.yml --ref main
```

Never move the addon's version backwards.

**Either way:**

- If the engine was already pinned for the new build, the ranking waits until it follows the old one. `gh workflow run engine.yml -f force=true` follows the active build.
- Hold the watcher so it does not move forward again the next morning: `gh workflow disable client.yml`, then `gh workflow enable client.yml` once the fix lands. Adding the key to `ignore_products` does not help for a key in `products`, because a listed key always counts.
