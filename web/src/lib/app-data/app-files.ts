// web/src/lib/app-data/app-files.ts
// The four W3-C1 files, built for one build, and their size budgets (spec F9 "Sync with the
// nightly data"): a guide at most 60 KB gzipped, tiers.json 40 KB, the guide index 10 KB. The
// site build fails over a budget, so a file the phones download never grows unnoticed.
import { gzipSync } from 'node:zlib';
import { readSpecCatalog } from '../bis/load';
import { loadTierData } from '../tiers/tier-load';
import { buildChangelog, type AppChangelog } from './changelog-file';
import { readSpecGuides } from './guide-content';
import { buildGuideFile, type AppGuideFile } from './guide-file';
import { buildGuideIndex, type AppGuideIndex } from './guide-index';
import { buildTiersFile, type AppTiersFile } from './tiers-file';

const KB = 1024;

export const APP_FILE_BUDGETS = {
  guide: 60 * KB,
  tiers: 40 * KB,
  guideIndex: 10 * KB,
  changelog: 60 * KB,
} as const;

export type AppFileKind = keyof typeof APP_FILE_BUDGETS;

/** The body as compact JSON; throws when its gzipped size passes the kind's budget. */
export function encodeAppFile(kind: AppFileKind, name: string, body: unknown): string {
  const text = JSON.stringify(body);
  const size = gzipSync(text).length;
  const budget = APP_FILE_BUDGETS[kind];
  if (size > budget) throw new Error(`${name} is ${size} bytes gzipped, over its ${budget}-byte budget`);
  return text;
}

const guideCache = new Map<string, AppGuideFile[]>();

/** Every spec guide's file on `build`, built once per build. */
export function appGuides(build: string): AppGuideFile[] {
  const cached = guideCache.get(build);
  if (cached !== undefined) return cached;
  const catalog = readSpecCatalog();
  const guides = readSpecGuides().map((source) => buildGuideFile(build, source, catalog));
  guideCache.set(build, guides);
  return guides;
}

export function appGuideIndex(build: string): AppGuideIndex {
  return buildGuideIndex(build, appGuides(build), readSpecCatalog());
}

export function appTiers(build: string): AppTiersFile {
  return buildTiersFile(build, loadTierData());
}

export function appChangelog(): AppChangelog {
  return buildChangelog();
}
