// web/src/lib/planner/score-unit.ts
// What the planner's score figure is called, per role. A healer's published band carries
// HPS and a tank's a tank score (`SlotScoreUnit`, bis/tank-view.ts); the planner's score
// strip, its band card and the "live when ..." line all read the word from here, so the
// two cannot drift into calling a healer's number DPS.
import { specKeyFor } from '../addon/score';
import { scoreUnitWord, type SlotScoreUnit } from '../bis/tank-view';
import { simCopy } from '../sim/copy';
import { specRow } from '../sim/spec-label';
import type { Estimate, SimResult } from '../sim/types';

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
  /** Under the figure while the build is unfinished: every role is simmed live once the
   *  build reaches 51 points, so every role says the same thing. */
  waiting: (): string => simCopy.plannerDpsLiveCaption,
  /** The live run failed or came back without the role's figure: "HPS estimate unavailable
   *  for this build." */
  failed: (unit: SlotScoreUnit): string => `${SCORE_LABEL[unit]} estimate unavailable for this build.`,
  /** The live figure as the strip prints it. A healer's reads to a tenth, like the level
   *  band's "40.9 HPS", so the strip and the band card agree on the digits they share;
   *  damage and a tank score are whole numbers. */
  figure: (value: number, unit: SlotScoreUnit): string =>
    unit === 'hps'
      ? value.toLocaleString('en-US', { minimumFractionDigits: 1, maximumFractionDigits: 1 })
      : Math.round(value).toLocaleString('en-US'),
  /** The level band's headline: "40.9 HPS". */
  bandFigure: (setScore: number, unit: SlotScoreUnit): string =>
    `${setScore.toFixed(1)} ${scoreUnitWord(unit)}`,
} as const;

/** Every role's live run asks the engine for its own figure; damage is the default and
 *  needs no opt-in. */
export function needsRoleMetrics(unit: SlotScoreUnit): boolean {
  return unit !== 'dps';
}

/**
 * The estimate a finished live run headlines, by role: damage per second, a healer's
 * effective HPS, a tank's score. Null when the engine returned a run without the role's
 * block, which the caller treats as a failed run rather than showing a damage number under
 * an HPS label.
 */
export function liveFigureOf(result: SimResult, unit: SlotScoreUnit): Estimate | null {
  switch (unit) {
    case 'dps':
      return result.dps;
    case 'hps':
      return result.healing?.effective_hps ?? null;
    case 'tank_score':
      return result.tank?.score ?? null;
  }
}
