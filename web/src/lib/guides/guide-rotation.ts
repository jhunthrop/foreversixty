// web/src/lib/guides/guide-rotation.ts
// The guide's rotation at level 60, for the one build its cards describe: the raid build, or
// the level 60 build when the guide has no separate raid build. The cooldowns, the talent cuts
// to them and to cast times, and the pre-pull timings all read that build's talents and no
// other (2026-11-04: Marksmanship's Rapid Fire read 3 minutes from the leveling build's Rapid
// Killing beside raid text that says 5; Arcane's opener read Frost's 2.5 s Frostbolt), and a
// cast whose talent gate rules that build out is not on its cards (round 18: Arcane's raid
// card listed the level 60 build's Frostbolt, cast only without Heating Up). The web page and
// the app's guide file both build it here.
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
  type GateSettled,
  type GuideRole,
  type LitTalents,
  type RotationLists,
  type SpellLookup,
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
  const lit = litTalentsOf(describedTrees);
  const facts = spellLookupFor(build, classSlug);
  const talentNames = talentNamesOf(build, classSlug);
  const runs = buildRunsGates(facts, talentNames, lit);
  const lines = rotationLinesFor(entry, build, classSlug, spec, top, runs);
  return rotationListsFrom(
    lines,
    loadApl(spec),
    facts,
    specCooldowns(build, classSlug, spec, lit, facts),
    talentCooldownCuts(build, classSlug, lit),
    racialCooldownsFor(build, classSlug),
    role,
    talentCastTimeCuts(build, classSlug, lit),
    runs,
    buildSettlesGate(facts, talentNames),
  );
}

/** Whether a build with `lit` talents makes a cast with these gates: every gate on knowing a
 *  talent (`auraIsKnown`, `spellIsKnown`, or either under `not`) holds for it. A gate on a
 *  trainer spell, or one the data cannot name, says nothing about the build. */
export function buildRunsGates(
  facts: SpellLookup,
  talentNames: ReadonlySet<string>,
  lit: LitTalents,
): (gates: readonly AplGate[]) => boolean {
  return (gates) =>
    gates.every((gate) => {
      const name = gate.kind === 'known' ? facts(gate.spellId)?.name : undefined;
      return name === undefined || !talentNames.has(name) || lit.has(name) === gate.positive;
    });
}

/** Whether a gate is on knowing one of the class's talents: the described build takes it or
 *  leaves it, which `buildRunsGates` has already applied, so a card need not word it. */
export function buildSettlesGate(
  facts: SpellLookup,
  talentNames: ReadonlySet<string>,
): GateSettled {
  return (gate) => {
    const name = gate.kind === 'known' ? facts(gate.spellId)?.name : undefined;
    return name !== undefined && talentNames.has(name);
  };
}

/** The build a guide's level 60 sections describe: the raid build when the guide has one that
 *  differs, else the level 60 build. */
export function describedBuildTrees(
  levelingTrees: readonly DecodedTreeView[],
  raidTrees: readonly DecodedTreeView[] | undefined,
): readonly DecodedTreeView[] {
  return raidTrees ?? levelingTrees;
}
