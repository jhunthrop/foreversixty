// web/src/lib/planner/score-unit.ts
// What the planner's score figure is called, per role. A healer's published band carries
// HPS and a tank's a tank score (`SlotScoreUnit`, bis/tank-view.ts); the planner's score
// strip, its band card and the "live when ..." line all read the word from here, so the
// two cannot drift into calling a healer's number DPS.
import { specKeyFor } from '../addon/score';
import { scoreUnitWord, type SlotScoreUnit } from '../bis/tank-view';
import { simCopy } from '../sim/copy';
import { specRow } from '../sim/spec-label';

const UNIT_BY_ROLE = { dps: 'dps', healer: 'hps', tank: 'tank_score' } as const;

/** The unit a spec's headline figure is in; an unknown spec is damage, the common case. */
export function scoreUnitForSpec(spec: string): SlotScoreUnit {
  const role = specRow(spec)?.role;
  return role === undefined ? 'dps' : UNIT_BY_ROLE[role];
}

/** The unit of the spec a build leans toward (the same inference the h1 suffix uses). */
export function scoreUnitForBuild(classSlug: string, pointsPerTree: readonly number[]): SlotScoreUnit {
  return scoreUnitForSpec(specKeyFor(classSlug, pointsPerTree));
}

const SCORE_LABEL: Record<SlotScoreUnit, string> = {
  dps: simCopy.plannerDpsLabel,
  hps: 'HPS',
  tank_score: 'Tank score',
};

const SCORE_RATE_PHRASE: Record<SlotScoreUnit, string> = {
  dps: 'damage per second',
  hps: 'healing per second',
  tank_score: 'tank score',
};

export const plannerScoreCopy = {
  /** The caption over the figure: "DPS", "HPS", "Tank score". */
  label: (unit: SlotScoreUnit): string => SCORE_LABEL[unit],
  /** The unit in plain words, for titles and prose. */
  ratePhrase: (unit: SlotScoreUnit): string => SCORE_RATE_PHRASE[unit],
  /** The ask-first button on a constrained device. */
  show: (unit: SlotScoreUnit): string =>
    unit === 'dps' ? simCopy.plannerDpsShow : `Show ${SCORE_LABEL[unit]}`,
  /** Under the figure while the build is unfinished. Only damage is simmed live (see
   *  `isLiveSimmed`), so a healer or tank says what is true instead of promising a figure. */
  waiting: (unit: SlotScoreUnit): string =>
    isLiveSimmed(unit) ? simCopy.plannerDpsLiveCaption : plannerScoreCopy.notSimmed(unit),
  /** The browser sim runs the damage model only: no raid damage to heal and no boss to
   *  tank, so it cannot produce HPS or a tank score. The band card carries those. */
  notSimmed: (unit: SlotScoreUnit): string =>
    `no live ${SCORE_LABEL[unit].toLowerCase()} yet; see the level band`,
  /** The level band's headline: "40.9 HPS". */
  bandFigure: (setScore: number, unit: SlotScoreUnit): string =>
    `${setScore.toFixed(1)} ${scoreUnitWord(unit)}`,
} as const;

/** Only a damage spec has a live number that is what its label says. */
export function isLiveSimmed(unit: SlotScoreUnit): boolean {
  return unit === 'dps';
}
