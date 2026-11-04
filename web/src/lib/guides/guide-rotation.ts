// web/src/lib/guides/guide-rotation.ts
// The guide's rotation at level 60, for the one build its cards describe: the raid build, or
// the level 60 build when the guide has no separate raid build. The cooldowns, the talent cuts
// to them and to cast times, and the pre-pull timings all read that build's talents and no
// other (2026-11-04: Marksmanship's Rapid Fire read 3 minutes from the leveling build's Rapid
// Killing beside raid text that says 5; Arcane's opener read Frost's 2.5 s Frostbolt). One
// exception: a pre-pull cast gated on a talent the described build does not match is the other
// build's, so its cast time is that build's (round 18: Arcane's Frostbolt, cast only without
// Heating Up, is the level 60 build's 2.5 s, not the raid build's 3 s). The web page and the
// app's guide file both build it here.
import { bandTopLevel, loadRotations, rotationEntryFor, rotationLinesFor } from '../bis/rotation-view';
import type { AplGate } from '../bis/apl-gates';
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
  talentNamesOf,
  type GuideRole,
  type LitTalents,
  type RotationLists,
  type SpellLookup,
} from './rotation-lists';

export const GUIDE_ROTATION_LEVEL = 60;

/** `describedTrees`: the build the rotation cards describe (`describedBuildTrees`);
 *  `otherTrees`: the guide's other build (the level 60 build beside a raid build), whose
 *  talents time the pre-pull casts only it runs. Undefined for a spec with no published
 *  rotation at level 60. */
export function guideRotation(
  build: string,
  classSlug: string,
  spec: string,
  role: GuideRole,
  describedTrees: readonly DecodedTreeView[],
  otherTrees?: readonly DecodedTreeView[],
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
    castTimeCutsByGates(
      build,
      classSlug,
      facts,
      lit,
      otherTrees === undefined ? undefined : litTalentsOf(otherTrees),
    ),
  );
}

/** The cast-time cuts for a pre-pull cast's gates: the described build's, unless its talent
 *  gates rule the described build out and admit the other build. Gates on trainer spells (not
 *  talents) say nothing about which build casts. */
function castTimeCutsByGates(
  build: string,
  classSlug: string,
  facts: SpellLookup,
  describedLit: LitTalents,
  otherLit: LitTalents | undefined,
): (gates: readonly AplGate[]) => ReadonlyMap<string, number> {
  const described = talentCastTimeCuts(build, classSlug, describedLit);
  if (otherLit === undefined) return () => described;
  const other = talentCastTimeCuts(build, classSlug, otherLit);
  const talentNames = talentNamesOf(build, classSlug);
  const runs = (lit: LitTalents, gates: readonly AplGate[]): boolean =>
    gates.every((gate) => {
      const name = gate.kind === 'known' ? facts(gate.spellId)?.name : undefined;
      return name === undefined || !talentNames.has(name) || lit.has(name) === gate.positive;
    });
  return (gates) => (!runs(describedLit, gates) && runs(otherLit, gates) ? other : described);
}

/** The build a guide's level 60 sections describe: the raid build when the guide has one that
 *  differs, else the level 60 build. */
export function describedBuildTrees(
  levelingTrees: readonly DecodedTreeView[],
  raidTrees: readonly DecodedTreeView[] | undefined,
): readonly DecodedTreeView[] {
  return raidTrees ?? levelingTrees;
}
