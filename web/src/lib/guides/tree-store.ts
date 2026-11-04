// web/src/lib/guides/tree-store.ts
// The read-only planner store a guide's talent tree renders from (GuideBuildTree,
// GuideLevelingTree): the code decoded, its class's talent file fetched once, and a store for
// a point order -- the code's own order reconstructed, or one the caller gives (the leveling
// tree's first N ladder points). Never a second tree renderer: TreeGrid draws the store.
import { decodeFS1, orderFromRanks } from '../planner/fs1';
import { loadTalents } from '../planner/load';
import { indexTalents } from '../planner/rules';
import { createPlannerStore, type PlannerStore } from '../planner/store.svelte';
import type { TalentFile } from '../planner/types';

export interface GuideTreeSource {
  treeVersion: string;
  classSlug: string;
  raceSlug: string;
  talentFile: TalentFile;
  /** The code's own point order, reconstructed legally from its ranks. */
  order: number[];
}

/** The code's talent file and order; throws when the code does not decode or the file does not load. */
export async function loadGuideTreeSource(code: string): Promise<GuideTreeSource> {
  const decoded = decodeFS1(code);
  if (!decoded.ok) throw new Error(`guide tree code does not decode: ${code}`);
  const talentFile = await loadTalents(decoded.build.dataBuild, decoded.build.classSlug);
  const { order } = orderFromRanks(indexTalents(talentFile), decoded.build.treeRanks);
  return {
    treeVersion: decoded.build.dataBuild,
    classSlug: decoded.build.classSlug,
    raceSlug: decoded.build.raceSlug,
    talentFile,
    order,
  };
}

/** A read-only store for `order` (the source's own order by default). */
export function guideTreeStore(
  source: GuideTreeSource,
  order: readonly number[] = source.order,
): PlannerStore {
  const store = createPlannerStore({
    treeVersion: source.treeVersion,
    classSlug: source.classSlug,
    raceSlug: source.raceSlug,
    order: [...order],
    readOnly: true,
  });
  store.setTalents(source.talentFile);
  return store;
}
