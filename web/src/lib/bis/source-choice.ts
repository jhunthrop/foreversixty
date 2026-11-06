// web/src/lib/bis/source-choice.ts
// The words and the set figure around the BiS page's source filters (web layout spec 2026-11-05
// §5.11). Filtering is re-picking, never hiding (source-filter.ts resolves the picks); this module
// says what was chosen, states the default rule in one line, and decides which set figure the
// head shows: the sim's own figure for the chosen sources when the file publishes one
// (`filtered_set_dps`), otherwise the unfiltered figure, labelled as the unfiltered one.
import type { BisBandModel } from './page-model';
import { MAX_BAND } from './page-state';
import { SOURCE_FILTER_KINDS } from './source-filter';
import type { SourceChoice } from './doll-view';
import type { SourceFilterKind } from './types';

/** The checklist's words for each kind. */
export function sourceKindLabel(kind: SourceFilterKind): string {
  switch (kind) {
    case 'dungeon':
      return 'Dungeon';
    case 'raid':
      return 'Raid';
    case 'world':
      return 'World / BoE';
    case 'quest':
      return 'Quest';
    case 'crafted':
      return 'Crafted';
    case 'rep':
      return 'Reputation';
    case 'vendor':
      return 'Vendor';
    case 'pvp':
      return 'PvP';
  }
}

/** The state's kinds as the filter resolves them. Only the leveling default (below 60, every
 *  kind but Reputation) keeps standings up to Friendly, the set the ranker measures as
 *  `leveling_default`; any other selection without Reputation turns it off entirely, so
 *  "Dungeon only" is dungeon drops only at every band. */
export function sourceChoiceFor(on: readonly SourceFilterKind[], band: number): SourceChoice {
  return isLevelingDefault(on, band) ? { on, repCap: 'friendly' } : { on };
}

function isLevelingDefault(on: readonly SourceFilterKind[], band: number): boolean {
  return band < MAX_BAND && on.length === SOURCE_FILTER_KINDS.length - 1 && !on.includes('rep');
}

/** The closed field's text: `All sources`, `Dungeon only`, `7 of 8 sources`. */
export function sourcesFieldLabel(on: readonly SourceFilterKind[], band: number): string {
  if (on.length === SOURCE_FILTER_KINDS.length) return 'All sources';
  if (isLevelingDefault(on, band)) return 'No rep above Friendly';
  if (on.length === 1) return `${sourceKindLabel(on[0]!)} only`;
  if (on.length === 0) return 'No sources';
  return `${on.length} of ${SOURCE_FILTER_KINDS.length} sources`;
}

/** The control bar's compact Sources button: the count of kinds on, or "All sources". */
export function sourcesButtonLabel(on: readonly SourceFilterKind[]): string {
  return on.length === SOURCE_FILTER_KINDS.length ? 'All sources' : `Sources (${on.length})`;
}

/** The words inside "Re-picked (…)". */
export function choiceWords(on: readonly SourceFilterKind[], band: number): string {
  if (isLevelingDefault(on, band)) return 'no reputation above Friendly';
  if (on.length === 1) return `${sourceKindLabel(on[0]!).toLowerCase()} only`;
  return 'your sources';
}

/** The default rule, stated once under the summary when it applies (§5.11 leveling defaults).
 *  Crafts are not gated by skill (no rule for a leveler's likely skill is decided yet). */
export function defaultRuleLine(on: readonly SourceFilterKind[], band: number): string | undefined {
  if (!isLevelingDefault(on, band)) return undefined;
  return 'Reputation above Friendly is off below level 60, so those slots are re-picked without it. Turn it on in Sources to include it.';
}

export interface SetFigure {
  value: string;
  /** The label over the figure. */
  label: string;
  /** One quiet line under the label: how far under the best set, or why it is the unfiltered one. */
  caption?: string;
  error?: number;
}

function filteredKey(
  on: readonly SourceFilterKind[],
  band: number,
): 'all' | 'dungeon' | 'leveling_default' | undefined {
  if (on.length === SOURCE_FILTER_KINDS.length) return 'all';
  if (isLevelingDefault(on, band)) return 'leveling_default';
  if (on.length === 1 && on[0] === 'dungeon') return 'dungeon';
  return undefined;
}

/** The head's figure for the chosen sources (`changed` slots differ from the best set). */
export function setFigureFor(
  model: BisBandModel,
  on: readonly SourceFilterKind[],
  changed: number,
): SetFigure | undefined {
  const answer = model.answer;
  if (answer.figure === undefined) return undefined;
  const best: SetFigure = {
    value: answer.figure,
    label: `This set, ${answer.simmed.toLowerCase()}`,
    error: answer.error,
  };
  if (changed === 0) return best;
  const key = filteredKey(on, model.band);
  const measured = key === undefined ? undefined : answer.filtered?.[key];
  if (measured !== undefined) {
    const under = Number(answer.figure) - measured.dps;
    return {
      value: measured.dps.toFixed(1),
      label: `Set: ${choiceWords(on, model.band)}`,
      caption: `${Math.abs(under).toFixed(1)} ${model.unit} ${under >= 0 ? 'under' : 'over'} the best set`,
      error: measured.error,
    };
  }
  return {
    ...best,
    label: 'Best set, all sources',
    caption: `Your sources change ${changed} ${changed === 1 ? 'slot' : 'slots'}; that set has not been simmed yet`,
  };
}
