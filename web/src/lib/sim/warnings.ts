// web/src/lib/sim/warnings.ts
// Worth knowing (spec 2026-11-04-sim-results 3.5): the notes that qualify a saved run's number,
// computed only from data the result already has. At most three, in the spec's priority order;
// none, and the card does not render. No note is dismissible, because each one is true.
import { parseActionKey, resolveActionName, type ActionNames } from './action-names';
import { formatMargin } from './estimate';
import { primaryPowerTrack, type PrimaryPower } from './primary-power';
import { grouped, warningsCopy } from './result-copy';
import { specDisplayName, specRow } from './spec-label';
import type { SimSpellFact } from './spell-info';
import type { SimResult } from './types';
import { isStale } from './version';

export type WarningId = 'stale' | 'capstone' | 'missing-talent' | 'at-cap' | 'wide-error';

/** Where a note's trailing link goes. */
export type WarningAction = 'sim' | 'planner';

export interface SimWarning {
  id: WarningId;
  text: string;
  /** The action key whose icon leads the row; absent, the row leads with the alert tile. */
  actionKey?: string;
  action?: WarningAction;
}

/** A talent ability the rotation casts that this build has 0 points in (rule W3). */
export interface TalentGap {
  actionKey: string;
}

export interface WarningInput {
  result: SimResult;
  names: ActionNames | null;
  /** The client's facts for a spell id; undefined until the bundle loads or for an unknown spell. */
  fact: (spellId: number) => SimSpellFact | undefined;
  /** The capstone's spell id; omitted, the spec list's. */
  capstone?: number;
  /** W3's input: null while the talents are not known to fit today's trees, so W3 is not evaluated. */
  talentGaps?: readonly TalentGap[] | null;
}

export const MAX_WARNINGS = 3;
/** A resource lost at the cap at or above this share of what was gained is worth a note. */
const AT_CAP_SHARE = 0.05;
/** An error above this share of the mean is wide enough to be worth a note. */
const WIDE_ERROR_SHARE = 0.005;
/** The level a character can first hold a 31-point talent: below it no capstone is expected. */
const CAPSTONE_LEVEL = 40;

/**
 * The spec's capstone as the spec list publishes it (`capstone_spell`, the level 60 rank's id; spec
 * 11.4). A spec with none has no capstone its rotation is expected to press, so it raises no note.
 */
export function capstoneSpellOf(spec: string): number | undefined {
  return (specRow(spec) as { capstone_spell?: number } | null)?.capstone_spell;
}

function staleWarning(result: SimResult): SimWarning | null {
  return isStale(result.engine_version) ? { id: 'stale', text: warningsCopy.stale, action: 'sim' } : null;
}

/** Whether any rank of the capstone was cast: ranks share a name, the published id is one of them. */
function capstoneCast(result: SimResult, spellId: number, names: ActionNames): boolean {
  const name = names.spell[String(spellId)];
  return result.summary.casts.some((row) => {
    const parsed = parseActionKey(row.spell_name);
    if (parsed?.kind !== 'spell' || row.succeeded <= 0) return false;
    return parsed.id === spellId || (name !== undefined && names.spell[parsed.label] === name);
  });
}

/** W2. It waits for the names: until then a lower rank could not be told from no cast at all. */
function capstoneWarning({ result, names, capstone }: WarningInput): SimWarning | null {
  const spellId = capstone === undefined ? capstoneSpellOf(result.request.spec) : capstone;
  if (spellId === undefined || names === null || result.request.character.level < CAPSTONE_LEVEL) return null;
  if (capstoneCast(result, spellId, names)) return null;
  const actionKey = `spell:${spellId}`;
  return {
    id: 'capstone',
    text: warningsCopy.capstone(resolveActionName(actionKey, names)),
    actionKey,
    action: 'sim',
  };
}

function missingTalentWarning(input: WarningInput, capstone: SimWarning | null): SimWarning | null {
  const gap = input.talentGaps?.[0];
  if (capstone !== null || gap === undefined) return null;
  const { result, names } = input;
  return {
    id: 'missing-talent',
    text: warningsCopy.missingTalent(
      resolveActionName(gap.actionKey, names),
      specDisplayName(result.request.spec),
    ),
    actionKey: gap.actionKey,
    action: 'planner',
  };
}

/** The powers with a cap worth naming. */
const CAPPED_POWER: Readonly<Partial<Record<PrimaryPower, string>>> = {
  rage: warningsCopy.atCapRage,
  energy: warningsCopy.atCapEnergy,
};

function atCapWarning(result: SimResult): SimWarning | null {
  const found = primaryPowerTrack(result);
  if (found === null) return null;
  const { power, track } = found;
  const advice = CAPPED_POWER[power];
  const wasted = track.wasted ?? 0;
  if (advice === undefined || track.gained <= 0 || wasted / track.gained < AT_CAP_SHARE) return null;
  const percent = ((wasted / track.gained) * 100).toFixed(1);
  return {
    id: 'at-cap',
    text: `${warningsCopy.atCap(grouped(wasted), power, percent)} ${advice}`,
  };
}

function wideErrorWarning(result: SimResult): SimWarning | null {
  const { error, mean } = result.dps;
  if (mean <= 0 || error / mean <= WIDE_ERROR_SHARE) return null;
  return { id: 'wide-error', text: warningsCopy.wideError(formatMargin(error)), action: 'sim' };
}

/** The notes for a run, highest priority first, at most `MAX_WARNINGS`. */
export function simWarnings(input: WarningInput): SimWarning[] {
  const { result } = input;
  const capstone = capstoneWarning(input);
  return [
    staleWarning(result),
    capstone,
    missingTalentWarning(input, capstone),
    atCapWarning(result),
    wideErrorWarning(result),
  ]
    .filter((note): note is SimWarning => note !== null)
    .slice(0, MAX_WARNINGS);
}
