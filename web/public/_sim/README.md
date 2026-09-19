# `/_sim/<ENGINE_VERSION>/`

`sim.wasm` and `sim.js`, built from the site's own `sim/` Go module (`sim/cmd/wasm`) at one
pinned engine version, served with
`Cache-Control: public, max-age=31536000, immutable` — the version is in the path, so a
new engine is a new directory and nothing is ever revalidated.

The files are **not** committed. CI builds them from the sha pinned in
`web/src/lib/sim/version.ts` and `sim/enginever/version.go` (kept in step by
`make engine-pin`) and writes them here before `astro build`.

Until CI publishes an artifact, `PUBLIC_SIM_ENGINE` is unset, which means `fake`, and
`web/src/fixtures/sim/engine-fake.ts` stands in over the same four functions. Set
`PUBLIC_SIM_ENGINE=wasm` in the build environment to switch; the one branch in
`web/src/lib/sim/engine.ts` is the whole swap, because everything the wasm exchanges with
the page is already JSON.
