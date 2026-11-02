# Sim-accuracy round 8: the WASM lane (2026-11-02)

Owner ruling 4: sim.wasm.gz is at 94.7% of its 4.5 MiB budget. Find cheap wins and record what is left. The budget is not raised. This lane did no per-spec audit. It changes no DPS and makes no engine commit.

## The build, measured

Built the way deploy builds it: engine at the pin cd4dfdcc0 (a worktree, through go.work), go1.25.11 from sim/go.mod's `go` line, `make simdb`, `make publish-wasm` (`GOOS=js GOARCH=wasm go build -trimpath -ldflags="-s -w"`). The budget is 4.5 MiB = 4,718,592 bytes.

| | raw sim.wasm | sim.wasm.gz | budget | headroom |
|---|---|---|---|---|
| live, deploy run 41374713528 (170b7b84c), `curl` 2026-11-02 | (not stored) | 4,485,806 | 95.1% | 232,786 |
| local, origin/main 170b7b84c (gzip -9) | 24,920,373 | 4,481,935 | 95.0% | 236,657 |
| **local, wt/r8-wasm (pigz -11 -I 3)** | 24,920,373 (identical bytes) | **4,276,658** | **90.6%** | **441,934** |

The live file answers `content-type: application/gzip` with no `Content-Encoding`, whatever Accept-Encoding asks for. So the page's DecompressionStream inflate path is the one in use, and the stored file is exactly what visitors download.

## What is in it

Sizes come from a symbol build (same flags without `-s -w`) and the wasm name section. The gzip column is each part compressed on its own, which is close to its share of the whole.

| part | raw | gzip -9 alone |
|---|---|---|
| code section | 11,673,166 | 2,294,528 |
| - engine (sim/core 1.99 MB, core/proto 1.08 MB, common 0.38 MB, nine classes 0.16 to 0.28 MB each) | 5,241,894 | 909,555 |
| - Go std (reflect, encoding/json, regexp, time, fmt, strconv, slices...) | 2,010,829 | 459,771 |
| - google.golang.org/protobuf | 2,063,122 | 381,292 |
| - runtime | 1,083,970 | 273,011 |
| - site (sim/request, api, bulk, adapter, combine, simdb, spellranks, logs/engine) | 397,493 | 91,597 |
| - package init functions (map literals; the two largest are sim/core `map.init.10` 184 KB and `map.init.9` 100 KB) | 849,308 | 126,371 |
| data section | 13,204,389 | 2,158,235 |
| - pclntab (function names, file table, pc tables) | 3,031,952 | 842,620 |
| - the embeds: simdb.bin 2,972,619, spellranks.json 430,002, enchants.json 54,003, apl/*.json 281,286, heal profile and tank encounter 11 KB | 3.75 MB | about 410,000 (simdb 342,833, spellranks 38,086, apls 22,755) |
| - a 2.85 MB repeating table | 2,850,165 | 26,851 |
| - type descriptors, strings, protobuf raw descriptors (about 100,000 small segments) | about 3.5 MB | about 840,000 |

What is not in it: net/http, the engine's db.bin (with_db), the adapter's testdata/*.result.pb (embedded in sim/adapter/fixture.go but dead-stripped: their bytes are absent from the module), and any duplicated embed. `testing` and `flag` are linked through engine sim/core and paladin/retribution non-test files, at 19 KB of code (3 KB gzip). That is not worth an engine change. regexp and encoding/json cannot be dropped from the site side, because google.golang.org/protobuf's own JSON and impl packages import them.

## Findings

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| W1 | **The deployed gzip was `gzip -9`, and a zopfli-class deflate of the same module is 4.6% smaller.** `pigz -11 -I 3 -n` gives 4,276,658 bytes against gzip -9's 4,481,935 (-205,277). It is still a plain gzip stream, so the page's `DecompressionStream('gzip')` and every standard inflater read it. The module does not change, so the sim cannot change. Variants: pigz -11 (15 iterations) 4,272,225; zopfli --i5 4,282,659; pigz -11 -b 1024 -I 5 4,286,606. -I 3 keeps 98% of the gain for about a fifth of the default time (72 s CPU, 39 s wall on this loaded laptop, about 20 s on a 4-core runner). pigz 2.8 is on GitHub's ubuntu-24.04 image (actions/runner-images Ubuntu2404-Readme.md). | `wc -c` on each output; `gzip -dc out \| cmp - sim.wasm`; live headers above | CONFIRMED (win) | S 2c0943d1c |
| W2 | **The budget measured a second compression, not the deployed file.** `make artifacts` gated `gzip -9 -c sim.wasm \| wc -c`, publish-wasm gzipped again, and sim.yml's artifact job had a third copy of the measurement. Now tools/wasm-gzip.sh writes artifacts/sim.wasm.gz, inflates it back with the system gzip and compares it with the module, then gates WASM_BUDGET_MB on that file. `make artifacts` and sim.yml's budget step call it, and publish-wasm copies the file as is. The deploy job's separate `make artifacts` step is gone, because publish-wasm runs it, so the module was built and measured twice. The real-engine cache key, the real-engine scope regex and the web.yml/sim.yml path filters now include the script, so a change to it cannot be served from a stale cached gzip. | Makefile `artifacts`/`publish-wasm`; .github/workflows/sim.yml "wasm size budget"; web.yml `real-engine` and `deploy` | CONFIRMED | S 2c0943d1c |
| W3 | GOWASM=satconv,signext saves nothing: 4,482,328 against 4,481,935 (satconv alone 4,481,084). | three builds, `gzip -9 \| wc -c` | Aligned (no change) | - |
| W4 | `-gcflags` without inlining (`-l`) shrinks the module, but every useful scope costs speed. All variants gave bit-identical results (W6 harness, 8 specs at 1,000 iterations). Node CPU time over 4 runs of 4 level-60 specs at 1,000 iterations: base 6,862 ms. Results: protobuf only -90 KB, +3.2% CPU. Protobuf plus std, keeping runtime/math/time/sort/slices inlined: -189 KB, +2.9%. Protobuf plus all std: -236 KB, +16.6%. engine sim/core only: -290 KB, +25% wall or worse in every round. Everything (`all=-l`): -599 KB, slower still. Not landed, because a slower sim on every visitor's machine is not a cheap win. The owner can trade it later (Still open). | builds v-pb, v-pbstd2, v-pbstd, v-enginecore, gc-l; scratch timing2.txt | OPEN (owner trade) | - |
| W5 | binaryen `wasm-opt -Oz` takes 2.7 MB off the raw module but only 39 KB off the gzip (-O2: 28 KB). It needs a new build tool in CI, and binaryen rewriting Go's wasm is a correctness risk for 0.9%. `-gcflags=all=-B` (no bounds checks) saves 130 KB but turns an out-of-range index into silent memory corruption. Neither landed. | wasm-opt 133 (brew), sizes above | Aligned (rejected) | - |
| W6 | **Proof of bit-identical results.** (a) The module is byte-identical before and after (`cmp`), and the new gzip inflates to it (`gzip -dc \| cmp`). (b) The wasm harness (node, the published sim.js, `simRun`) ran all 28 specs at levels 60 and 40, 56 requests at 300 iterations, from the old gzip -9 file and the new pigz file. All 56 canonical result JSONs are identical (combined hash 957d8d134a59bc13 both times). (c) The four real-engine browser specs (sim-real-engine, sim-gear-real-engine, sim-weights-real-engine, sim-talents-real-engine), `npm run test:e2e:real-engine -- --workers=1`, against a `FOREVER_DATA=real PUBLIC_SIM_ENGINE=wasm` build: 4 passed with the new file and 4 passed with the old one. (d) spec-breakdown (native, origin/main against wt/r8-wasm) at seeds 1 and 50001, 10,000 iterations: identical DPS and error to the last digit for mage-fire 659.977 and 661.032, warrior-fury 767.841 and 768.518, warlock-affliction 574.523 and 575.001, hunter-survival 772.638 and 772.773, priest-shadow 714.105 and 714.242. No server left running (port 4341 free). | scratch hc-before.txt / hc-after.txt, e2e.txt, sb-spot.txt | CONFIRMED (no change) | - |
| W7 | Side note from the harness: the same request gives the same numbers, but some summary rows with equal sort keys come back in a different order run to run (druid-balance: item:13444, other:mana_gain and spell:1249110 swap). This is map order, not a sim difference. The harness compares order-insensitively for that reason. A stable tiebreak (spell id) would make result JSON reproducible byte for byte. | two runs of the same base module, `diff` of the JSON | OPEN (minor, not a sim bug) | - |
| W8 | With 4 Playwright workers on this laptop (load average 60 to 85 across the shared lanes), sim-talents-real-engine timed out on "2 valid combinations" with both the old and the new gzip. Alone, or with `--workers=1`, it passes, and CI passed it at 170b7b84c. That is contention, not the wasm. | three runs | Aligned (environment) | - |

## Before/after

| | before (origin/main) | after (wt/r8-wasm) | change |
|---|---|---|---|
| sim.wasm raw | 24,920,373 | 24,920,373 | 0 (identical bytes) |
| sim.wasm.gz | 4,481,935 | 4,276,658 | -205,277 (-4.6%) |
| share of the 4.5 MiB budget | 95.0% | 90.6% | -4.4 points |
| headroom | 236,657 | 441,934 | +205,277 |
| DPS, every spec | - | - | 0 (W6) |

The integrator should expect the deployed figure to read about 4.28 MB at this pin. CI's module differs from a go.work build by a few KB of build info.

## What is left: next options

Ordered by saving. All estimates are against the new 4,276,658-byte file.

| option | estimated saving | cost and risk |
|---|---|---|
| Serve brotli: store `sim.wasm.br` (q11: 3,080,096 bytes) and serve it with `Content-Encoding: br` from src/worker.ts (or decode it in the page) | about -1.2 MB (28%), to about 65% of budget | DecompressionStream does not take brotli in every browser the site supports, so this needs either a Worker response with Content-Encoding: br (unproven: Cloudflare must pass the stored encoding through without re-encoding; check with curl on a preview deploy before switching) or a JS/wasm brotli decoder in the page (about 30 to 60 KB, and the inflate is no longer native). It changes wasm-fetch.ts and the probe in support/real-engine.ts. This is the one large lever left. |
| No inlining in protobuf plus non-hot std (W4 pbstd2) | about -190 KB (4 points) | about +3% CPU per sim in the browser |
| No inlining in protobuf only (W4) | about -90 KB (2 points) | about +3% CPU |
| simdb.bin re-encoding (2.97 MB raw, 343 KB gzip, the largest embed) | maybe -100 to -170 KB if columnar or delta-coded | pipeline change to a nightly-owned artifact plus a decoder in sim/internal/simdb; not cheap |
| Engine: turn the big package-level map literals (sim/core `map.init.10`/`map.init.9`, likely `baseStatsByClassLevel` and the client set/spell rows in client_sets_gen.go) into static arrays or embedded tables | about -25 to -30 KB | engine generator change and fork goldens; lookups must keep map semantics for missing keys |
| binaryen `wasm-opt -Oz` (W5) | about -40 KB | new CI tool, correctness risk on Go wasm |
| pigz -11 with the default 15 iterations instead of 3 | -4.4 KB | about 5x the compression time |
| pclntab (843 KB gzip, the largest single part: function and file names for tracebacks) | up to several hundred KB with a renaming build (garble -tiny) | invasive. It renames symbols that protobuf and reflection read, and breaks panic traces. Not recommended. |

Things measured to be worth nothing: GOWASM (W3), dropping `testing`/`flag` (3 KB), and dropping the site's regexp or encoding/json (protobuf imports both).

## For other lanes

- Any lane: `make artifacts` now needs `pigz` locally (`brew install pigz`). It fails fast with that message if pigz is missing.
- Integrator: the deploy job's step list changed (one publish step, no separate `make artifacts`), and the real-engine cache key now hashes tools/wasm-gzip.sh, so the first run after merge is a cache miss.

## Still open

- W4 and brotli are owner trades (speed, or a serving-path change that needs a preview deploy to prove).
- W7: an order-stable summary for byte-reproducible result JSON (adapter owner).
