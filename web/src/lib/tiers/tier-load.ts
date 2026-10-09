// web/src/lib/tiers/tier-load.ts
// Build-time data for the /tiers pages: the level 60 raid band of every written spec, ranked
// per faction and role. Files come through the BiS loaders (`lib/bis/load.ts`), so a build the
// nightly has not ranked yet reads the newest other build's files exactly as the BiS pages do.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import activeBuild from '../../data/active-build.json';
import { latest } from '../dates';
import { bandEntry, loadBisFile, readSpecCatalog } from '../bis/load';
import { presetLabelFor } from '../bis/presets';
import type { BisFile, BisRole, Faction, SpecCatalogEntry } from '../bis/types';
import { TIER_ROLES, rankRole, type TierInput, type TierRow } from './tier-list';

const TIER_BAND = 60;
export const TIER_FACTIONS: readonly Faction[] = ['alliance', 'horde'];

const REPO_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
const TANK_ENCOUNTER_FILE = 'data/curated/tank-encounter.json';

export type RoleLists = Record<BisRole, TierRow[]>;

export interface TierData {
  /** Ranked rows per faction and role. */
  lists: Record<Faction, RoleLists>;
  /** The raid preset's label, e.g. the eyebrow's "Raid-ready, Phase 1"; `null` with no files. */
  presetLabel: string | null;
  /** The newest `generated_at` across the files; `null` with no files. */
  updated: Date | null;
  /** `build` and `engine_version` of the newest file, for the stamp's tooltip. */
  build: string | null;
  engineVersion: string | null;
  /** `heal_profile.duration_sec` from a healer file. */
  healSeconds?: number;
  /** `tank-encounter.json` boss swing speed. */
  bossSwingSeconds?: number;
}

export interface TierLoadDeps {
  catalog: readonly SpecCatalogEntry[];
  loadFile: (spec: string) => BisFile | null;
  readBossSwingSeconds: () => number | undefined;
}

function emptyLists(): RoleLists {
  return { dps: [], tank: [], healer: [] };
}

function inputsFor(
  files: readonly { entry: SpecCatalogEntry; file: BisFile }[],
  faction: Faction,
): TierInput[] {
  return files.flatMap(({ entry, file }) => {
    const band = bandEntry(file, TIER_BAND, faction);
    return band === undefined ? [] : [{ entry, band }];
  });
}

function listsFor(inputs: readonly TierInput[]): RoleLists {
  return Object.fromEntries(TIER_ROLES.map((role) => [role, rankRole(inputs, role)])) as RoleLists;
}

function newestFile(files: readonly BisFile[]): BisFile | undefined {
  return [...files].sort((a, b) => b.generated_at.localeCompare(a.generated_at))[0];
}

function healSecondsOf(files: readonly BisFile[]): number | undefined {
  return files.find((file) => file.heal_profile !== undefined)?.heal_profile?.duration_sec;
}

/** The tier data from the supplied files: pure, so the tests feed it fixtures. */
export function buildTierData(deps: TierLoadDeps): TierData {
  const loaded = deps.catalog.flatMap((entry) => {
    const file = deps.loadFile(entry.spec);
    return file === null ? [] : [{ entry, file }];
  });
  const files = loaded.map(({ file }) => file);
  const newest = newestFile(files);
  const firstBand = loaded
    .map(({ file }) => ({ file, band: bandEntry(file, TIER_BAND, 'alliance') }))
    .find((candidate) => candidate.band !== undefined);
  return {
    lists: {
      alliance: loaded.length === 0 ? emptyLists() : listsFor(inputsFor(loaded, 'alliance')),
      horde: loaded.length === 0 ? emptyLists() : listsFor(inputsFor(loaded, 'horde')),
    },
    presetLabel: firstBand?.band === undefined ? null : presetLabelFor(firstBand.file, firstBand.band),
    updated: files.length === 0 ? null : latest(files.map((file) => new Date(file.generated_at))),
    build: newest?.build ?? null,
    engineVersion: newest?.engine_version ?? null,
    healSeconds: healSecondsOf(files),
    bossSwingSeconds: deps.readBossSwingSeconds(),
  };
}

function readBossSwingSeconds(): number | undefined {
  const parsed = JSON.parse(readFileSync(path.join(REPO_ROOT, TANK_ENCOUNTER_FILE), 'utf8')) as {
    boss?: { swing_speed_sec?: { value?: number } };
  };
  return parsed.boss?.swing_speed_sec?.value;
}

export function loadTierData(): TierData {
  return buildTierData({
    catalog: readSpecCatalog(),
    loadFile: (spec) => loadBisFile(spec, activeBuild.build),
    readBossSwingSeconds,
  });
}
