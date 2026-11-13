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
import { rankingSimHrefFor } from '../bis/sim-link';
import type { BisBand, BisFile, BisRole, Faction, SpecCatalogEntry } from '../bis/types';
import { loadTalentFile } from '../guides/build-tree';
import { indexTalents, type TalentIndex } from '../planner/rules';
import {
  TIER_ROLES,
  healerBasisOf,
  healerOrderMeasured,
  rankRole,
  type HealerBasis,
  type TierInput,
  type TierRow,
} from './tier-list';

const TIER_BAND = 60;

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
  /** Per faction: every healer band publishes its error and both tuning bounds, so the
   *  healer list can order at all. Without them every healer shows tied. */
  healerOrderMeasured: Record<Faction, boolean>;
  /** The fight both factions' healer lists rank on: the raid fight once every healer band of
   *  both publishes it, the gear fight until then, so the page names one fight. */
  healerBasis: HealerBasis;
  /** `heal_profile.duration_sec` from a healer file: the gear fight's length. */
  healSeconds?: number;
  /** `metrics.raid_fight.duration_sec` from a healer band: the raid fight's length. */
  raidFightSeconds?: number;
  /** `tank-encounter.json` boss swing speed. */
  bossSwingSeconds?: number;
}

export interface TierLoadDeps {
  catalog: readonly SpecCatalogEntry[];
  loadFile: (spec: string) => BisFile | null;
  readBossSwingSeconds: () => number | undefined;
  /** The row's "View ranking sim" href, or `undefined` where /sim cannot rerun the band. */
  rankingSimHref: (entry: SpecCatalogEntry, band: BisBand) => string | undefined;
}

function emptyLists(): RoleLists {
  return { dps: [], tank: [], healer: [] };
}

function inputsFor(
  files: readonly { entry: SpecCatalogEntry; file: BisFile }[],
  faction: Faction,
  rankingSimHref: TierLoadDeps['rankingSimHref'] = () => undefined,
): TierInput[] {
  return files.flatMap(({ entry, file }) => {
    const band = bandEntry(file, TIER_BAND, faction);
    if (band === undefined) return [];
    const simHref = rankingSimHref(entry, band);
    return [simHref === undefined ? { entry, band } : { entry, band, simHref }];
  });
}

function listsFor(inputs: readonly TierInput[], healerBasis: HealerBasis): RoleLists {
  return Object.fromEntries(
    TIER_ROLES.map((role) => [role, rankRole(inputs, role, healerBasis)]),
  ) as RoleLists;
}

function newestFile(files: readonly BisFile[]): BisFile | undefined {
  return [...files].sort((a, b) => b.generated_at.localeCompare(a.generated_at))[0];
}

function healSecondsOf(files: readonly BisFile[]): number | undefined {
  return files.find((file) => file.heal_profile !== undefined)?.heal_profile?.duration_sec;
}

function raidFightSecondsOf(inputs: readonly TierInput[]): number | undefined {
  for (const { band } of inputs) {
    const metrics = band.metrics;
    if (metrics !== null && metrics !== undefined && 'hps' in metrics && metrics.raid_fight !== undefined) {
      return metrics.raid_fight.duration_sec;
    }
  }
  return undefined;
}

/** The tier data from the supplied files: pure, so the tests feed it fixtures. */
export function buildTierData(deps: TierLoadDeps): TierData {
  const loaded = deps.catalog.flatMap((entry) => {
    const file = deps.loadFile(entry.spec);
    return file === null ? [] : [{ entry, file }];
  });
  const files = loaded.map(({ file }) => file);
  const newest = newestFile(files);
  const healerBasis = healerBasisOf([...inputsFor(loaded, 'alliance'), ...inputsFor(loaded, 'horde')]);
  const firstBand = loaded
    .map(({ file }) => ({ file, band: bandEntry(file, TIER_BAND, 'alliance') }))
    .find((candidate) => candidate.band !== undefined);
  return {
    lists: {
      alliance:
        loaded.length === 0
          ? emptyLists()
          : listsFor(inputsFor(loaded, 'alliance', deps.rankingSimHref), healerBasis),
      horde:
        loaded.length === 0
          ? emptyLists()
          : listsFor(inputsFor(loaded, 'horde', deps.rankingSimHref), healerBasis),
    },
    presetLabel: firstBand?.band === undefined ? null : presetLabelFor(firstBand.file, firstBand.band),
    updated: files.length === 0 ? null : latest(files.map((file) => new Date(file.generated_at))),
    build: newest?.build ?? null,
    engineVersion: newest?.engine_version ?? null,
    healerOrderMeasured: {
      alliance: healerOrderMeasured(inputsFor(loaded, 'alliance'), healerBasis),
      horde: healerOrderMeasured(inputsFor(loaded, 'horde'), healerBasis),
    },
    healerBasis,
    healSeconds: healSecondsOf(files),
    raidFightSeconds: raidFightSecondsOf(inputsFor(loaded, 'alliance')),
    bossSwingSeconds: deps.readBossSwingSeconds(),
  };
}

function readBossSwingSeconds(): number | undefined {
  const parsed = JSON.parse(readFileSync(path.join(REPO_ROOT, TANK_ENCOUNTER_FILE), 'utf8')) as {
    boss?: { swing_speed_sec?: { value?: number } };
  };
  return parsed.boss?.swing_speed_sec?.value;
}

/** The active build's talent trees per class, read once each: the sim link's code carries the
 *  band's points these trees can reach, as the BiS page's own links do. */
function talentIndexReader(build: string): (classSlug: string) => TalentIndex {
  const cache = new Map<string, TalentIndex>();
  return (classSlug) => {
    const cached = cache.get(classSlug);
    if (cached !== undefined) return cached;
    const file = loadTalentFile(build, classSlug);
    if (file === undefined) {
      throw new Error(
        `data/builds/${build}/talents/${classSlug}.json is missing; the sim link needs the trees`,
      );
    }
    const index = indexTalents(file);
    cache.set(classSlug, index);
    return index;
  };
}

/** The tier row's sim link over the active build: tank and healer bands get none, since their
 *  DTPS/HPS is not a figure the browser sim runs (`rankingSimHrefFor`). */
export function tierRankingSimHref(build: string): TierLoadDeps['rankingSimHref'] {
  const indexFor = talentIndexReader(build);
  return (entry, band) => rankingSimHrefFor(band, indexFor(entry.class_slug), build);
}

export function loadTierData(): TierData {
  return buildTierData({
    catalog: readSpecCatalog(),
    loadFile: (spec) => loadBisFile(spec, activeBuild.build),
    readBossSwingSeconds,
    rankingSimHref: tierRankingSimHref(activeBuild.build),
  });
}
