// web/src/lib/character-header/answer-view.ts
// What the answer slot draws, as data: the header's state selection for the slot in one pure
// function (spec 2026-10-11 section 3). Ready, loading (the answer only), the level-floor
// sentence, the no-gear sentence and link, the empty sentence, failed, and nothing at all.
import type { MeCharacter } from '../account/api';
import { planAnswer } from './answer-plan';
import type { AnswerLink, AnswerSpec, FiguresAnswer, HeaderAnswer } from './types';

/** Where the page's `load` is for the current character. */
export type AnswerPhase =
  { phase: 'loading' } | { phase: 'ready'; answer: HeaderAnswer } | { phase: 'failed' };

export type AnswerView =
  | { kind: 'loading' }
  | { kind: 'figures'; answer: FiguresAnswer; label: string }
  | { kind: 'sentence'; text: string }
  | { kind: 'no-gear'; sentence: string; link: AnswerLink }
  | { kind: 'failed'; text: string };

/** The label a figures answer wears: the page's stale label when its answer reads gear and the sync is stale. */
export function figuresLabel(answer: FiguresAnswer, stale: boolean): string {
  return stale && answer.staleLabel !== undefined ? answer.staleLabel : answer.label;
}

function viewOfLoaded(answer: HeaderAnswer, stale: boolean): AnswerView | null {
  if (answer.kind === 'none') return null;
  if (answer.kind === 'sentence') return { kind: 'sentence', text: answer.text };
  return { kind: 'figures', answer, label: figuresLabel(answer, stale) };
}

/** Null when the slot shows nothing (no answer spec, or the page said `none`, or the page has no words for a state). */
export function answerViewFor(
  character: MeCharacter,
  spec: AnswerSpec | undefined,
  phase: AnswerPhase,
  stale: boolean,
): AnswerView | null {
  if (spec === undefined) return null;
  const plan = planAnswer(character, spec);
  if (plan === 'below-floor') {
    const text = spec.strings.belowLevelFloor;
    return text === undefined ? null : { kind: 'sentence', text };
  }
  if (plan === 'no-gear') {
    const noGear = spec.strings.noGear;
    return noGear === undefined ? null : { kind: 'no-gear', sentence: noGear.sentence, link: noGear.link };
  }
  if (phase.phase === 'loading') return { kind: 'loading' };
  if (phase.phase === 'failed') return { kind: 'failed', text: spec.strings.failed };
  return viewOfLoaded(phase.answer, stale);
}
