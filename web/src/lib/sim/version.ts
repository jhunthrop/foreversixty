// web/src/lib/sim/version.ts
// One string identifies the engine build everywhere (simulator interface contract, "Engine
// version"): the short sha of wowsims-forever the artifacts were built from. It names the
// immutable asset directory, it goes into every SimRequest and SimResult, and a result that
// carries a different one is labelled stale and never silently re-run.
//
// It is a literal rather than an import from sim/enginever/version.go because the web build
// has no Go toolchain; `make engine-pin` writes both, version.test.ts asserts they agree,
// and web.yml runs that test.

export const ENGINE_VERSION = '8b2169e61';

/** Where a version's browser artifacts live. Cached immutably, so the path carries the sha. */
export function engineAssetUrl(file: 'sim.wasm' | 'sim.js', version: string = ENGINE_VERSION): string {
  return `/_sim/${version}/${file}`;
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
