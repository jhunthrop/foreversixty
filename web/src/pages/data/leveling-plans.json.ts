// /data/leveling-plans.json: each spec guide's first talent points, from level 10 to the first
// gear range, in the order a player spends them (design/specs/2026-11-04-low-level-home.md).
// Built from the guide's own build through the same walk its Leveling strip uses
// (lib/guides/leveling-builds.ts), so the home page names exactly the talents the guide does.
import type { APIRoute } from 'astro';
import { getCollection } from 'astro:content';
import activeBuild from '../../data/active-build.json';
import { loadTalentFile } from '../../lib/guides/build-tree';
import { guideLadderOrder, guideLevelingBuilds } from '../../lib/guides/leveling-builds';
import { guideTipsUrl } from '../../lib/guides/guide-tips';
import { FIRST_BAND_LEVEL } from '../../lib/bis/landing-bands';
import { jsonResponse } from '../../lib/app-data/endpoint';
import { levelingGuideHref, type LevelingPlan, type LevelingPoint } from '../../lib/leveling/leveling-plan';
import { FIRST_POINT_LEVEL, type TalentFile } from '../../lib/planner/types';
import { SPECS } from '../../lib/sim/specs';

interface TalentCell {
  name: string;
  icon: string;
  maxRank: number;
  tree: number;
  cell: number;
}

/** Every talent of the class by id, with where its digit sits in a talent string. */
function talentCells(file: TalentFile): Map<number, TalentCell> {
  const cells = new Map<number, TalentCell>();
  file.trees.forEach((tree, treeIndex) => {
    [...tree.talents]
      .sort((a, b) => a.tier - b.tier || a.column - b.column)
      .forEach((talent, cell) => {
        cells.set(talent.id, {
          name: talent.name,
          icon: talent.icon,
          maxRank: talent.max_rank,
          tree: treeIndex,
          cell,
        });
      });
  });
  return cells;
}

/** The order as points: each with its level and the talent's rank once it is in. */
function pointsOf(order: readonly number[], cells: ReadonlyMap<number, TalentCell>): LevelingPoint[] {
  const ranks = new Map<number, number>();
  return order.flatMap((talentId, index) => {
    const cell = cells.get(talentId);
    if (cell === undefined) return [];
    const rank = (ranks.get(talentId) ?? 0) + 1;
    ranks.set(talentId, rank);
    return [{ level: FIRST_POINT_LEVEL + index, talentId, rank, ...cell }];
  });
}

function planFor(specKey: string, buildCode: string): LevelingPlan | undefined {
  const spec = SPECS.find((row) => row.spec === specKey);
  if (spec === undefined) return undefined;
  const build = activeBuild.build;
  const first = guideLevelingBuilds(build, spec.class_slug, specKey, spec.tree_index, buildCode).find(
    (row) => row.band === FIRST_BAND_LEVEL,
  );
  const file = loadTalentFile(build, spec.class_slug);
  if (first === undefined || file === undefined) return undefined;
  // The ladder's own order (the guide strip's and the addon's), cut at the first range.
  const order = guideLadderOrder(build, spec.class_slug, specKey, spec.tree_index, buildCode).slice(
    0,
    first.points,
  );
  return {
    specKey,
    classSlug: spec.class_slug,
    specName: spec.name,
    guideHref: levelingGuideHref(specKey) ?? '',
    plannerHref: `/planner?spec=${specKey}&talents=${encodeURIComponent(first.talents)}`,
    tipsUrl: guideTipsUrl(build, spec.class_slug, spec.spec_slug),
    points: pointsOf(order, talentCells(file)),
  };
}

export const GET: APIRoute = async () => {
  const guides = await getCollection('guides');
  const plans: Record<string, LevelingPlan> = {};
  for (const entry of guides) {
    const [classSlug, specSlug] = entry.id.split('/');
    if (entry.data.spec === undefined || entry.data.build === undefined) continue;
    const specKey = `${classSlug}-${specSlug}`;
    const plan = planFor(specKey, entry.data.build);
    if (plan !== undefined) plans[specKey] = plan;
  }
  return jsonResponse(JSON.stringify(plans));
};
