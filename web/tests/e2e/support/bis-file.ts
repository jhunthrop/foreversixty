// web/tests/e2e/support/bis-file.ts
// The BiS file a /bis/<class>/<spec> page was built from, so a spec asserts what the page
// renders from the file instead of pinning numbers that every nightly re-rank moves. The
// precedence is the page loader's (src/lib/bis/load.ts): the published
// data/builds/<build>/bis/<spec>.json wins over the committed fixture. Every written spec
// now has a published file, so the fixture is a fallback these specs never reach.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { bisSourceBuild } from '../../../scripts/bis-source-build.mjs';
import { ACTIVE_BUILD } from './active-build';

const WEB_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../..');
const REPO_ROOT = path.resolve(WEB_ROOT, '..');

export type BisPresetId = 'raid' | 'bare';

export interface BisHitToCapFile {
  baseline: number;
  specials: number;
  white?: number;
}

export interface TankMetricsFile {
  effective_health: number;
  dtps: number;
  chance_of_death: number;
  tps: number;
  tmi: number;
  dps: number;
}

export interface BisWeightFile {
  stat: string;
  rating_factor?: number;
  insignificant?: boolean;
}

export interface BisBandFile {
  spec: string;
  faction: 'alliance' | 'horde';
  band: number;
  preset?: BisPresetId;
  hit_to_cap?: (BisHitToCapFile & { kind?: string }) | null;
  metrics?: TankMetricsFile | null;
  weights?: BisWeightFile[];
  role?: 'dps' | 'tank' | 'healer';
  race?: string;
  set_dps?: number;
  weights_low_confidence?: boolean;
  generated_at?: string;
}

interface PresetFile {
  label: string;
  buffs: { id: string; label: string }[];
  debuffs: { id: string; label: string }[];
  consumes: { id: string; label: string }[];
}

export interface BisFileData {
  bands: BisBandFile[];
  presets?: Partial<Record<BisPresetId, PresetFile>>;
  heal_profile?: { label: string };
}

/** The published file for a spec, from the active build or, before the nightly has ranked
 *  it, the newest ranked build (the page loader's own fallback). */
function publishedBisPath(spec: string): string {
  const source = bisSourceBuild(REPO_ROOT, ACTIVE_BUILD);
  return path.join(REPO_ROOT, 'data/builds', source, 'bis', `${spec}.json`);
}

export function readBisFile(spec: string): BisFileData {
  const published = publishedBisPath(spec);
  const fixture = path.join(WEB_ROOT, 'src/data/fixtures/bis', `${spec}.json`);
  const file = existsSync(published) ? published : fixture;
  return JSON.parse(readFileSync(file, 'utf8')) as BisFileData;
}

/** The band the page shows for a faction, level and preset; level 60 publishes both presets. */
export function bisBand(
  spec: string,
  faction: BisBandFile['faction'],
  band: number,
  preset: BisPresetId = 'raid',
): BisBandFile {
  const found = readBisFile(spec).bands.find(
    (candidate) =>
      candidate.faction === faction &&
      candidate.band === band &&
      (candidate.preset === undefined || candidate.preset === preset),
  );
  if (found === undefined) throw new Error(`${spec} publishes no ${faction} band ${band} (${preset})`);
  return found;
}

export function bisPreset(spec: string, preset: BisPresetId): PresetFile {
  const found = readBisFile(spec).presets?.[preset];
  if (found === undefined) throw new Error(`${spec} publishes no ${preset} preset`);
  return found;
}

/** "Hit to cap: 7% for specials[, 27% for white swings]" for a whole-percent melee cap. */
export function hitCapText(hitToCap: BisHitToCapFile): string {
  const white = hitToCap.white === undefined ? '' : `, ${hitToCap.white}% for white swings`;
  return `Hit to cap: ${hitToCap.specials}% for specials${white}`;
}

const WHOLE = new Intl.NumberFormat('en-US', { maximumFractionDigits: 0 });
const ONE_DECIMAL = new Intl.NumberFormat('en-US', { minimumFractionDigits: 1, maximumFractionDigits: 1 });
/** A non-zero chance of death under this share reads "<0.1%" rather than "0.0%". */
const CHANCE_RESOLUTION = 0.001;

export interface TankFigureTexts {
  /** Effective health, damage taken per second, chance of death, threat per second. */
  headline: [string, string, string, string];
  secondary: string;
}

/** What the tank panel prints for a band's metrics, in the order the panel shows them. */
export function tankFigureTexts(metrics: TankMetricsFile): TankFigureTexts {
  const chance =
    metrics.chance_of_death > 0 && metrics.chance_of_death < CHANCE_RESOLUTION
      ? '<0.1%'
      : `${(metrics.chance_of_death * 100).toFixed(1)}%`;
  return {
    headline: [
      WHOLE.format(metrics.effective_health),
      ONE_DECIMAL.format(metrics.dtps),
      chance,
      ONE_DECIMAL.format(metrics.tps),
    ],
    secondary: `TMI ${WHOLE.format(metrics.tmi)} · own damage ${ONE_DECIMAL.format(metrics.dps)} DPS`,
  };
}

export function tankMetrics(spec: string, preset: BisPresetId): TankMetricsFile {
  const { metrics } = bisBand(spec, 'alliance', 60, preset);
  if (metrics === null || metrics === undefined) throw new Error(`${spec} band 60 publishes no tank metrics`);
  return metrics;
}

export interface SpecCatalogRow {
  class_slug: string;
  spec_slug: string;
  spec: string;
  name: string;
  role: 'dps' | 'tank' | 'healer';
}

/** data/curated/specs.json: every written spec, in catalogue order. */
export function readSpecCatalog(): SpecCatalogRow[] {
  return JSON.parse(
    readFileSync(path.join(REPO_ROOT, 'data/curated/specs.json'), 'utf8'),
  ) as SpecCatalogRow[];
}

/** `/bis/<class>/<spec>` for the first written spec no file (published or fixture) covers, or
 *  `undefined` when every spec has one, as it does once the nightly ranks every role. */
export function unrankedBisRoute(): string | undefined {
  const bare = readSpecCatalog().find(
    ({ spec }) =>
      !existsSync(publishedBisPath(spec)) &&
      !existsSync(path.join(WEB_ROOT, 'src/data/fixtures/bis', `${spec}.json`)),
  );
  return bare === undefined ? undefined : `/bis/${bare.class_slug}/${bare.spec_slug}`;
}
