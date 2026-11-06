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
import { bandModelKey } from './model-key';
import { modelItemIds, type BisPageData } from './page-model';
import { presetsOffered } from './page-context';

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

/** The default state's models alone (first band, Alliance, each preset) with their items. */
export function defaultSlice(data: BisPageData): BisPageData {
  const band = data.bands[0];
  if (band === undefined) return data;
  // Every preset of the default band, so the Profile control renders the same with either set.
  const keys = presetsOffered(data, 'alliance', band).map((preset) => bandModelKey('alliance', band, preset));
  const models = Object.fromEntries(keys.map((key) => [key, data.models[key]!]));
  const ids = modelItemIds(Object.values(models));
  const items = Object.fromEntries(Object.entries(data.items).filter(([id]) => ids.has(Number(id))));
  return { ...data, models, items, partial: true };
}
