// web/src/lib/home/about-load.ts
// Build-time facts for home's server-rendered words (lib/home/about.ts): the active client build and
// when the pipeline last read it (its manifest), which specs have a BiS file, the bands they cover
// and when they were ranked, the first tier's opening, and the busiest ranked encounters (the
// rankings snapshot). Node only; index.astro calls it.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import activeBuild from '../../data/active-build.json';
import { loadBisFile } from '../bis/load';
import { SPECS } from '../sim/specs';
import { loadRankingsSnapshot } from '../rankings/snapshot-load';
import type { RankingsSnapshot } from '../rankings/snapshot';
import type { AboutFacts } from './about';

const REPO_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');

/** When the data pipeline last read the build's game files, from its manifest; null without one. */
function readAt(build: string): string | null {
  const file = path.join(REPO_ROOT, 'data/builds', build, 'manifest.json');
  if (!existsSync(file)) return null;
  const manifest = JSON.parse(readFileSync(file, 'utf8')) as { fetched_at?: unknown };
  return typeof manifest.fetched_at === 'string' ? manifest.fetched_at : null;
}

export function loadAboutFacts(
  raidsOpen: string | null,
  build: string = activeBuild.build,
  snapshot: RankingsSnapshot = loadRankingsSnapshot(),
): AboutFacts {
  const files = SPECS.flatMap((spec) => {
    const file = loadBisFile(spec.spec, build);
    return file === null ? [] : [{ spec: spec.spec, file }];
  });
  const bandNumbers = files.flatMap(({ file }) => file.bands.map((band) => band.band));
  const ranked = files.map(({ file }) => file.generated_at).sort();
  return {
    build,
    readAt: readAt(build),
    bisRankedAt: ranked.at(-1) ?? null,
    bisSpecs: new Set(files.map(({ spec }) => spec)),
    bands:
      bandNumbers.length === 0 ? null : { first: Math.min(...bandNumbers), last: Math.max(...bandNumbers) },
    raidsOpen,
    ranked: snapshot.encounters.map(({ name, slug, size }) => ({ name, slug, size })),
  };
}
