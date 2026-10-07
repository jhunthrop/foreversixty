// web/src/lib/sim/version.ts
// One string identifies the engine build everywhere (simulator interface contract, "Engine
// version"): the short sha of wowsims-forever the artifacts were built from. It names the
// immutable asset directory, it goes into every SimRequest and SimResult, and a result that
// carries a different one is labelled stale and never silently re-run.
//
// It is a literal rather than an import from sim/enginever/version.go because the web build
// has no Go toolchain; `make engine-pin` writes both, version.test.ts asserts they agree,
// and web.yml runs that test.

export const ENGINE_VERSION = '91f45d755';

/**
 * The directory the browser artifacts were published under for THIS page build: the engine
 * sha plus twelve hex digits of the wasm's own sha256 (`make publish-wasm`), handed to the
 * build by web.yml as PUBLIC_SIM_ARTIFACT. The site's request layer is compiled into the
 * wasm and changes without the pin moving, and the directory is served immutable for a
 * year, so the name has to change whenever the bytes do. A build with no artifact (the
 * fake engine, tests) falls back to the bare version.
 */
// `import.meta.env` is Vite's: astro/vitest define it, but this module is also loaded by
// plain Node (Playwright's config and the e2e specs import it), where it is undefined.
const viteEnv: Record<string, string | undefined> | undefined = import.meta.env;
export const ENGINE_ARTIFACT: string = viteEnv?.PUBLIC_SIM_ARTIFACT ?? ENGINE_VERSION;

/** Where the browser artifacts live. Cached immutably, so the path carries the bytes' name. */
export function engineAssetUrl(file: 'sim.wasm' | 'sim.js', artifact: string = ENGINE_ARTIFACT): string {
  return `/_sim/${artifact}/${file}`;
}

/**
 * True when a stored result was produced by some other engine build. An empty version counts
 * as stale: a result with no provenance is exactly the one a player should be told about.
 */
export function isStale(resultVersion: string): boolean {
  return resultVersion !== ENGINE_VERSION;
}

/** How a version reads in the UI. */
export function engineLabel(version: string): string {
  return `Engine ${version}`;
}
