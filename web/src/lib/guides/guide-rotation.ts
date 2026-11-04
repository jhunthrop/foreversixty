// web/src/lib/guides/guide-rotation.ts
// The guide's rotation at level 60, for the one build its cards describe: the raid build, or
// the level 60 build when the guide has no separate raid build. The cooldowns, the talent cuts
// to them and to cast times, and the pre-pull timings all read that build's talents and no
// other (2026-11-04: Marksmanship's Rapid Fire read 3 minutes from the leveling build's Rapid
// Killing beside raid text that says 5; Arcane's opener read Frost's 2.5 s Frostbolt). The web
// page and the app's guide file both build it here.
import { bandTopLevel, loadRotations, rotationEntryFor, rotationLinesFor } from '../bis/rotation-view';
import type { DecodedTreeView } from './build-tree';
import { racialCooldownsFor } from './race-table';
import {
  litTalentsOf,
  loadApl,
  rotationListsFrom,
  specCooldowns,
  spellLookupFor,
  talentCastTimeCuts,
  talentCooldownCuts,
  type GuideRole,
  type RotationLists,
} from './rotation-lists';

export const GUIDE_ROTATION_LEVEL = 60;

/** `describedTrees`: the build the rotation cards describe (`describedBuildTrees`). Undefined
 *  for a spec with no published rotation at level 60. */
export function guideRotation(
  build: string,
  classSlug: string,
  spec: string,
  role: GuideRole,
  describedTrees: readonly DecodedTreeView[],
): RotationLists | undefined {
  const top = bandTopLevel(GUIDE_ROTATION_LEVEL);
  const entry = rotationEntryFor(loadRotations(build, spec) ?? [], top);
  if (entry === undefined) return undefined;
  const lines = rotationLinesFor(entry, build, classSlug, spec, top);
  const lit = litTalentsOf(describedTrees);
  const facts = spellLookupFor(build, classSlug);
  return rotationListsFrom(
    lines,
    loadApl(spec),
    facts,
    specCooldowns(build, classSlug, spec, lit, facts),
    talentCooldownCuts(build, classSlug, lit),
    racialCooldownsFor(build, classSlug),
    role,
    talentCastTimeCuts(build, classSlug, lit),
  );
}

/** The build a guide's level 60 sections describe: the raid build when the guide has one that
 *  differs, else the level 60 build. */
export function describedBuildTrees(
  levelingTrees: readonly DecodedTreeView[],
  raidTrees: readonly DecodedTreeView[] | undefined,
): readonly DecodedTreeView[] {
  return raidTrees ?? levelingTrees;
}
