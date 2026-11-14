// web/src/lib/bis/page-server.ts
// The BiS page's build-time data for one spec, built once per build process and shared by the
// page (which inlines only the default state's slice) and its JSON endpoint (the full set, which
// the island fetches after mount). Keeping the full set out of the HTML keeps it off the first
// paint's critical path.
import activeBuild from '../../data/active-build.json';
import { loadTalentFile } from '../guides/build-tree';
import { indexTalents } from '../planner/rules';
import { loadBisFile } from './load';
import { buildBisPageData, type BisServerPage } from './page-data';
import type { ItemTooltipModel } from '../items/tooltip';
import { bandModelKey } from './model-key';
import { modelItemIds, type BisPageData } from './page-model';
import { presetsOffered } from './page-context';
import { MAX_BAND } from './band-levels';
import { listIsPreRaid } from './pre-raid';

const memo = new Map<string, BisServerPage | null>();

/** The spec's page data, or null when the spec has no ranked BiS file yet. */
export function bisServerPage(spec: string, classSlug: string): BisServerPage | null {
  const key = `${activeBuild.build}:${spec}`;
  if (!memo.has(key)) memo.set(key, buildFor(spec, classSlug));
  return memo.get(key)!;
}

function buildFor(spec: string, classSlug: string): BisServerPage | null {
  const build = activeBuild.build;
  const file = loadBisFile(spec, build);
  if (file === null) return null;
  const talentFile = loadTalentFile(build, classSlug);
  if (talentFile === undefined) {
    throw new Error(
      `data/builds/${build}/talents/${classSlug}.json is missing; the sim link needs the trees`,
    );
  }
  return buildBisPageData({ spec, classSlug, build, file, talentIndex: indexTalents(talentFile) });
}

const preRaidMemo = new Map<string, boolean>();

/** Whether the spec's level 60 list holds no raid or PvP pick, read off the ranked file itself
 *  (pre-raid.ts listIsPreRaid): then it is the pre-raid list, the page's title says so
 *  (seo-copy.ts) and /60/pre-raid is canonical to /60. Raids opening put raid drops in the list on
 *  the next nightly run, and the next build drops the word and gives /60/pre-raid its own page. */
export function bisPreRaid(spec: string): boolean {
  const key = `${activeBuild.build}:${spec}`;
  if (!preRaidMemo.has(key)) {
    const file = loadBisFile(spec, activeBuild.build);
    const sixty = file?.bands.filter((band) => band.band >= MAX_BAND) ?? [];
    preRaidMemo.set(key, sixty.length > 0 && listIsPreRaid(sixty.flatMap((band) => band.slots)));
  }
  return preRaidMemo.get(key)!;
}

/** The default state's models alone (the page's band, Alliance, each preset) with their items:
 *  the first band on /bis/<class>/<spec>, a band page's own band (lib/bis/band-path.ts). */
export function defaultSlice(data: BisPageData, band: number | undefined = data.bands[0]): BisPageData {
  if (band === undefined) return data;
  // Every preset of the default band, so the Profile control renders the same with either set.
  const keys = presetsOffered(data, 'alliance', band).map((preset) => bandModelKey('alliance', band, preset));
  const models = Object.fromEntries(keys.map((key) => [key, data.models[key]!]));
  const ids = modelItemIds(Object.values(models));
  const items = Object.fromEntries(Object.entries(data.items).filter(([id]) => ids.has(Number(id))));
  return { ...data, models, items, partial: true };
}

/** The tooltips of a slice's items alone (2026-11-13, page weight): the page inlines these and the
 *  island takes the rest with the full data (lib/bis/page-registry.ts addTooltipModels). Every
 *  band's tooltips inline were 167 KB of a feral page's 585 KB. */
export function sliceTooltips(
  tooltips: Readonly<Record<number, ItemTooltipModel>>,
  slice: BisPageData,
): Record<number, ItemTooltipModel> {
  return Object.fromEntries(Object.entries(tooltips).filter(([id]) => id in slice.items));
}
