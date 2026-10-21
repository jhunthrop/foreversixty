// web/src/lib/character-header/answer-view.test.ts
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../account/api';
import { answerViewFor, type AnswerPhase } from './answer-view';
import type { AnswerSpec, HeaderAnswer } from './types';

const FIGURES: HeaderAnswer = {
  kind: 'figures',
  label: 'Upgrades for your level range',
  staleLabel: 'Slots you wore at the last sync',
  figures: [
    { value: '9', unit: 'upgrades' },
    { value: '+11.2', unit: 'DPS' },
  ],
  note: { text: 'Best in slot for 60, together' },
};

const SPEC: AnswerSpec = {
  levelFloor: 20,
  needsGear: true,
  strings: {
    belowLevelFloor: 'Best in slot starts at 20.',
    noGear: { sentence: 'Add the addon.', link: { label: 'Get the addon', href: '/setup' } },
    failed: 'Your upgrades did not load.',
  },
  load: () => Promise.resolve(FIGURES),
};

const READY: AnswerPhase = { phase: 'ready', answer: FIGURES };

function character(overrides: Partial<MeCharacter> = {}): MeCharacter {
  return {
    key: 'us/normal/tester',
    region: 'us',
    ruleset: 'normal',
    name: 'Tester',
    class: 'Warrior',
    level: 60,
    build: { source: 'addon', captured_at: '2026-10-11T00:00:00Z', gear: { head: 1 } },
    ...overrides,
  };
}

describe('answerViewFor', () => {
  it('carries both figures of a two-figure answer', () => {
    const view = answerViewFor(character(), SPEC, READY, false);
    expect(view).toMatchObject({ kind: 'figures', label: 'Upgrades for your level range' });
    expect(view?.kind === 'figures' && view.answer.figures.map((figure) => figure.value)).toEqual([
      '9',
      '+11.2',
    ]);
  });

  it('is a skeleton while the answer loads', () => {
    expect(answerViewFor(character(), SPEC, { phase: 'loading' }, false)).toEqual({ kind: 'loading' });
  });

  it('uses the page’s stale label only when the sync is stale', () => {
    expect(answerViewFor(character(), SPEC, READY, true)).toMatchObject({
      label: 'Slots you wore at the last sync',
    });
    expect(answerViewFor(character(), SPEC, READY, false)).toMatchObject({
      label: 'Upgrades for your level range',
    });
  });

  it('shows the level-floor sentence below the floor without waiting for the answer', () => {
    expect(answerViewFor(character({ level: 12 }), SPEC, { phase: 'loading' }, false)).toEqual({
      kind: 'sentence',
      text: 'Best in slot starts at 20.',
    });
  });

  it('shows the no-gear sentence and link, never a zero', () => {
    const view = answerViewFor(character({ build: undefined }), SPEC, { phase: 'loading' }, false);
    expect(view).toEqual({
      kind: 'no-gear',
      sentence: 'Add the addon.',
      link: { label: 'Get the addon', href: '/setup' },
    });
  });

  it('shows the failed sentence the page supplied', () => {
    expect(answerViewFor(character(), SPEC, { phase: 'failed' }, false)).toEqual({
      kind: 'failed',
      text: 'Your upgrades did not load.',
    });
  });

  it('renders an empty answer as one sentence', () => {
    const empty: AnswerPhase = {
      phase: 'ready',
      answer: { kind: 'sentence', text: 'Every slot matches the 60 list.' },
    };
    expect(answerViewFor(character(), SPEC, empty, false)).toEqual({
      kind: 'sentence',
      text: 'Every slot matches the 60 list.',
    });
  });

  it('shows nothing for a page with no answer, or an answer that says none', () => {
    expect(answerViewFor(character(), undefined, READY, false)).toBeNull();
    expect(answerViewFor(character(), SPEC, { phase: 'ready', answer: { kind: 'none' } }, false)).toBeNull();
  });

  it('leaves a page with no floor and no gear need unchanged for a level 5 with no build', () => {
    const logs: AnswerSpec = { strings: { failed: 'x' }, load: () => Promise.resolve(FIGURES) };
    expect(answerViewFor(character({ level: 5, build: undefined }), logs, READY, false)?.kind).toBe(
      'figures',
    );
  });
});
