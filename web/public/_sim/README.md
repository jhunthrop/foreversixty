# `/_sim/<ENGINE_VERSION>-<sha256 of sim.wasm, 12 hex>/`

`sim.wasm` and `sim.js`, built from the site's own `sim/` Go module (`sim/cmd/wasm`) at one
pinned engine version, served with
`Cache-Control: public, max-age=31536000, immutable` — the directory name carries the engine sha AND a hash of the wasm's own
bytes (`make publish-wasm` writes the name to `artifacts/ARTIFACT_ID`; web.yml hands it to the
page build as `PUBLIC_SIM_ARTIFACT`, which `web/src/lib/sim/version.ts` reads). The site's own
request layer is compiled into the wasm and changes without the pin moving, so the bytes name
the directory, not the pin alone: a
new engine is a new directory and nothing is ever revalidated.

The files are **not** committed. CI builds them from the sha pinned in
`web/src/lib/sim/version.ts` and `sim/enginever/version.go` (kept in step by
`make engine-pin`) and writes them here before `astro build`.

Until CI publishes an artifact, `PUBLIC_SIM_ENGINE` is unset, which means `fake`, and
`web/src/fixtures/sim/engine-fake.ts` stands in over the same four functions. Set
`PUBLIC_SIM_ENGINE=wasm` in the build environment to switch; the one branch in
`web/src/lib/sim/engine.ts` is the whole swap, because everything the wasm exchanges with
the page is already JSON.
