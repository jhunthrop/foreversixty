# Making the repositories private

The site repository (jhunthrop/foreversixty) and the engine fork (jhunthrop/wowsims-forever)
go private after launch, and nothing public may link to them or lead back to the owner. This
is the order that keeps downloads, updates, store publishing and the nightly bots working.
Making a repository private is the owner's call and is done by hand in its Settings; nothing
in this repository flips it.

Do the steps in order. A step that says "wait" gates the next one.

## What already depends on nothing public

- **The site** links to no repository (web/scripts/check-no-owner-links.mjs fails the build if
  `github.com/jhunthrop` reaches `dist/`; src/lib/site/owner-links.test.ts scans the sources).
  "Report a bug" goes to the Discord.
- **Downloads** go through `foreversixty.gg/dl/...` (web/src/lib/downloads). The Worker serves
  them out of the foreversixty-logs bucket under `downloads/`:

  | Path                        | What                                   |
  | --------------------------- | -------------------------------------- |
  | `/dl/<family>/latest.json`  | the family's manifest, cached a minute |
  | `/dl/<family>/<tag>/<name>` | one published file, immutable          |
  | `/dl/<family>/<name>`       | 302 to that file in the latest release |
  | `/dl/addon`, `/dl/data`     | 302 to the latest zip                  |

  Families are `companion`, `addon` and `data`. The manifest's field names match a GitHub
  release object (`tag_name`, `assets[].name`, `assets[].browser_download_url`) plus `schema`,
  `family`, `published_at` and each asset's `size` and `sha256`.

- **Publishing**: addon-release.yml, addon-data-release.yml and companion.yml run
  tools/publish-feed.sh after their GitHub release. It uploads each signed file, then the
  manifest last, and refuses any file without its `.minisig`. It uses the existing
  `CLOUDFLARE_R2_TOKEN` and `CLOUDFLARE_ACCOUNT_ID` secrets, the same token web.yml uses
  for the DuckDB runtime in the same bucket.
- **The companion** (from the release that carries this change) reads
  `foreversixty.gg/dl/companion/latest.json` and `/dl/addon/latest.json`, `/dl/data/latest.json`
  instead of the GitHub API, with the same minisign checks. Older companions still read
  GitHub's releases API, which stops answering them the moment the repository is private.

## 1. Fill the feed (any time after slice B lands)

Each family's feed fills on its next release: the data pack nightly, the addon on its next tag,
the companion on the release that ships the feed updater. To fill it at once from the current
GitHub releases (a signed release only; run from a checkout with the R2 token exported):

```sh
export CLOUDFLARE_API_TOKEN=<the R2 token> CLOUDFLARE_ACCOUNT_ID=<account id>
for family in companion addon data; do
  tag=$(gh release list --limit 100 --json tagName --jq "[.[] | select(.tagName | startswith(\"$family-v\"))][0].tagName")
  dir=$(mktemp -d) && gh release download "$tag" -D "$dir"
  tools/publish-feed.sh "$family" "$tag" "$dir"/*
done
```

Check:

```sh
for f in companion addon data; do curl -fsS https://foreversixty.gg/dl/$f/latest.json | jq -r .tag_name; done
curl -sI https://foreversixty.gg/dl/addon | grep -i '^location: /dl/addon/addon-v'
curl -sI https://foreversixty.gg/dl/companion/foreversixty-companion_windows_amd64.exe | grep -i '^location: /dl/companion/'
```

Then delete `RELEASES_URL` from web/wrangler.jsonc's `vars` (and the `legacyLocation` fallback
in web/src/lib/downloads/serve.ts with its tests) and deploy. From then on no download touches
GitHub.

## 2. Wait for the companions to move

Every companion older than the feed release checks GitHub every six hours and updates itself to
the newest companion release it finds there, which is the feed release or later. Once the
repository is private those companions stop updating for good: the player has to download a new
one from /setup by hand. So the flip waits until most companions have updated.

Measure it two ways:

- The feed release's Windows asset `download_count` on GitHub, which is old companions pulling
  the update: `gh release view <companion-v tag> --json assets --jq '.assets[] | "\(.name) \(.downloadCount)"'`.
- Companions seen lately, from the API database:
  `select count(*) from devices where kind <> 'app' and revoked_at is null and last_seen_at > now() - interval '14 days';`

Flip when the first is at least 90% of the second, or after four weeks, whichever comes first.
After the flip a stale companion keeps working (sign-in, uploads, its installed addon), it only
stops updating; /setup's download is the fix.

## 3. Tokens: what breaks when the repositories are private

Everything that reads a private repository needs a credential; `GITHUB_TOKEN` only covers the
repository its workflow runs in. Set these up and prove each with a dispatch before the flip.

| Where                                                                                                                                                                             | What it does with GitHub today                                                                                           | After the flip                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| bis.yml, data.yml, sim.yml: `git clone https://github.com/jhunthrop/wowsims-forever .engine`                                                                                      | anonymous clone of the fork                                                                                              | needs a token: clone with `https://x-access-token:${{ secrets.ENGINE_REPO_TOKEN }}@github.com/...`. ENGINE_REPO_TOKEN must then grant contents: read on the fork to these workflows too (it already has contents: write for engine.yml).                                                                                                                                                                                                                             |
| .github/actions/pin-engine (api.yml, bis.yml, client.yml, engine.yml, sim.yml, web.yml): `go mod edit -replace ...=github.com/jhunthrop/wowsims-forever@<sha>` then `go mod tidy` | Go fetches the fork through proxy.golang.org                                                                             | the public proxy cannot fetch a private module. Give the action a `token` input and, before its `go mod edit`, set `GOPRIVATE=github.com/jhunthrop/*` (Go then asks GitHub directly, not the proxy or sum.golang.org) and `git config --global url."https://x-access-token:${TOKEN}@github.com/jhunthrop/".insteadOf "https://github.com/jhunthrop/"`. Every caller passes `token: ${{ secrets.ENGINE_REPO_TOKEN }}`: one change in the action, one line per caller. |
| api.yml: docker/build-push-action with api/Dockerfile (`ENGINE_MODULE=github.com/jhunthrop/wowsims-forever`, `go mod tidy` inside the build)                                      | the build fetches the fork through the proxy                                                                             | pass the token as a BuildKit secret (`secrets: engine_token=${{ secrets.ENGINE_REPO_TOKEN }}` on the action, `RUN --mount=type=secret,id=engine_token ...` in the Dockerfile) and set `GOPRIVATE` and the same `insteadOf` inside that `RUN`. Never an `ARG`: build args are kept in the image history.                                                                                                                                                              |
| engine.yml: `actions/checkout` of `FORK_REPO` (two jobs), `gh pr list/create/edit -R` the fork, push to the fork                                                                  | checkout is anonymous ("the fork is public, so the checkout needs no token"); push and PRs already use ENGINE_REPO_TOKEN | add `token: ${{ secrets.ENGINE_REPO_TOKEN }}` to both fork checkouts.                                                                                                                                                                                                                                                                                                                                                                                                |
| tools/engine/pin-conditions.sh and pinned-build.sh (engine.yml): `gh run list -R` / `gh api repos/<fork>/contents`                                                                | `GH_TOKEN` is the site repo's `github.token`, which can read a public fork                                               | run them with `GH_TOKEN: ${{ secrets.ENGINE_REPO_TOKEN }}`; the token needs actions: read and contents: read on the fork.                                                                                                                                                                                                                                                                                                                                            |
| bis.yml, data.yml, client.yml: commit and `git push` to this repository, `gh workflow run`, `gh issue create`                                                                     | `GITHUB_TOKEN` with `contents: write` / `actions: write` / `issues: write`                                               | unchanged: `GITHUB_TOKEN` works the same in a private repository.                                                                                                                                                                                                                                                                                                                                                                                                    |
| addon-release.yml, addon-data-release.yml: BigWigsMods/packager to CurseForge (`CF_API_KEY`) and Wago (`WAGO_API_TOKEN`)                                                          | the packager reads the throwaway checkout the workflow builds, not GitHub                                                | unchanged. On the CurseForge and Wago project pages, remove any "Source" or "Issues" link to GitHub by hand.                                                                                                                                                                                                                                                                                                                                                         |
| companion.yml: Azure Trusted Signing over OIDC (environment `signing`, subject `repo:jhunthrop/foreversixty:environment:signing`)                                                 |                                                                                                                          | unchanged: the subject names the repository, not its visibility.                                                                                                                                                                                                                                                                                                                                                                                                     |
| tools/land.sh, the integrator's `gh`                                                                                                                                              | the owner's `gh` login                                                                                                   | unchanged.                                                                                                                                                                                                                                                                                                                                                                                                                                                           |

Before the flip, make the fork private first and run, by dispatch: sim.yml, bis.yml (dry run),
data.yml, engine.yml (`dry_run`), and a branch push that runs api.yml's image build. All green
means the tokens are right. Then the site repository.

## 4. GitHub Actions minutes

Public repositories run Actions free; private ones spend the account's included minutes (2,000
a month on Free, 3,000 on Pro), and Windows minutes count double and macOS ten times. The week
to 2026-11-06 used more than 4,000 minutes of run time across 1,000+ runs (web 1,500, bis
1,170, api 630; job minutes are higher still, since jobs run in parallel). Before the flip,
the owner either:

- sets an Actions budget on the account (Settings, Billing and plans) large enough for a
  month at that rate, or
- cuts the schedule first (bis.yml and data.yml nightlies, web.yml's per-push lanes), or
- moves the heavy jobs to self-hosted runners.

Without one of these, every workflow, the releases and the nightly bots among them, stops when
the minutes run out.

## 5. What else leads back

- **proxy.golang.org and sum.golang.org** keep every module version they have ever served. The
  fork's pinned versions (`github.com/jhunthrop/wowsims-forever@<sha>`) stay downloadable there
  after the fork is private. The Go team removes a module from the proxy only on request
  (see the FAQ at https://proxy.golang.org); file it after the flip if
  that matters.
- **Git history and release metadata**: commits carry the owner's name and email. They are
  private with the repository; nothing published carries them (check the binaries, below).
- **Published binaries and the engine module** carry Go package paths. See "Public artifacts"
  below.

## 6. The flip

1. Fork: Settings, General, Danger zone, Change visibility, Private. Run the dispatch set in
   section 3.
2. Site repository: the same, once section 2's measure is met and section 4 is settled.
3. Check: `curl -fsS https://foreversixty.gg/dl/companion/latest.json`, /setup's three download
   buttons, the next nightly data release reaching the feed, and the next bis.yml run.

## Public artifacts

Scanned 2026-11-06 for `jhunthrop`, `hunthrop`, the owner's full name and email, and
`wowsims-forever`, as ASCII and UTF-16LE, decompressed. The owner's name and email are in no
public file. What remains after slices A and B:

| Artifact                                            | What it carries                                                                                                                                                                                                                                                                                              | Fix                                                                                                                                                                                |
| --------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| The engine module `/_sim/<artifact>/sim.wasm.gz`    | 673 `github.com/jhunthrop/foreversixty/...` strings (function and type names, `.go` paths, build info) and one `=> github.com/jhunthrop/wowsims-forever` build-info line. It is already built `-trimpath -ldflags="-s -w"`; Go keeps package paths in the runtime's function table, so no flag removes them. | Rename the site's Go modules (below).                                                                                                                                              |
| The companion binaries (Windows exe, Linux)         | about 1,100 `github.com/jhunthrop/foreversixty/companion/...` strings, the same way. The release before the feed also carries `api.github.com` and `jhunthrop/foreversixty` in its updater. Version info and the Authenticode signer say Commish LLC only.                                                   | Rename the site's Go modules; the feed release removes the updater strings.                                                                                                        |
| Data zips up to data-v2026.11.05                    | `CHANGELOG.md` linked `github.com/jhunthrop/foreversixty/tree/<tag>`                                                                                                                                                                                                                                         | Fixed: addon-data-release.yml writes its own changelog, and both release workflows refuse a zip that names the repository or owner. Published zips on CurseForge and Wago keep it. |
| Earlier releases' GitHub bodies                     | "Full Changelog" compare links                                                                                                                                                                                                                                                                               | Private with the repository.                                                                                                                                                       |
| proxy.golang.org                                    | 15 versions of `github.com/jhunthrop/wowsims-forever`, each `.info` naming the fork's URL and commit                                                                                                                                                                                                         | Only a removal request to the Go team takes them down. A directory replace for the engine (below) stops new ones.                                                                  |
| Older `/_sim/<artifact>/` directories on Cloudflare | not checked                                                                                                                                                                                                                                                                                                  | They go with the next deploys' asset set.                                                                                                                                          |

Measured on the pinned engine (26.5 MB wasm; counts of `jhunthrop` / `wowsims-forever`):

| Build                                                              | Size       | jhunthrop | wowsims-forever |
| ------------------------------------------------------------------ | ---------- | --------- | --------------- |
| no flags                                                           | 27,486,707 | 1,526     | 516             |
| `-trimpath`                                                        | 27,422,221 | 1,074     | 1               |
| `-ldflags="-s -w"`                                                 | 26,552,303 | 1,125     | 516             |
| `-trimpath -ldflags="-s -w"` (today)                               | 26,487,791 | 673       | 1               |
| today's flags, site modules renamed, engine as a directory replace | 26,476,348 | 0         | 0               |

The engine fork needs no rename: its module path is already `github.com/wowsims/classic`.
The smallest change that clears the wasm and the companion is renaming the site's four Go
modules (`github.com/jhunthrop/foreversixty/{api,sim,logs,companion}`) to a path that never
resolves, such as `foreversixty.gg/x/...` (all four are `replace`d locally, so nothing
fetches them), and pinning the engine as a directory replace
(`replace github.com/wowsims/classic => ../.engine`, checked out at the pinned sha) instead of
`=> github.com/jhunthrop/wowsims-forever@<sha>` through the proxy. Blast radius: 1,111 files
and 3,290 lines, nearly all Go imports (api 708 files, sim 234, logs 67, companion 51), the 4
go.mod files, the pin-engine action, api/Dockerfile, the Makefile and a handful of docs; a
mechanical rewrite that builds cleanly. It also removes the need for `GOPRIVATE` in section 3,
since nothing then goes through the proxy. Not done here: it is the coordinator's call.
