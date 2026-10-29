// web/src/lib/billing/coach-sample.test.ts
import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { coachSampleView, loadCoachSampleAnswer, type IconOf } from './coach-sample';
import { plansPageCopy } from './plans-page-copy';

const answer = loadCoachSampleAnswer();
const iconOf: IconOf = (id) => (id === null ? null : `/data/b/icons/${id}.webp`);

describe('loadCoachSampleAnswer', () => {
  it('reads the Fury model fixture as a full review', () => {
    expect(answer.spec).toBe('Fury');
    expect(answer.access).toBe('full');
  });
});

describe('the sample file against the coach fixture', () => {
  const read = (url: string): Record<string, unknown> =>
    JSON.parse(readFileSync(new URL(url, import.meta.url), 'utf8')) as Record<string, unknown>;
  const fixture = read('../../fixtures/coach/fury-model.json');
  const sample = (read('../../data/premium-coach-sample.json') as { answer: Record<string, unknown> }).answer;

  /** Every leaf path whose value differs, as `a.b[0].c`. */
  function differences(a: unknown, b: unknown, at = ''): string[] {
    if (typeof a !== 'object' || a === null || typeof b !== 'object' || b === null) {
      return a === b ? [] : [at];
    }
    const keys = new Set([...Object.keys(a), ...Object.keys(b)]);
    return [...keys].flatMap((key) =>
      differences((a as Record<string, unknown>)[key], (b as Record<string, unknown>)[key], `${at}.${key}`),
    );
  }

  it('differs only by the player ruling that the figures agree', () => {
    expect(differences(fixture, sample).sort()).toEqual(
      [
        '.review.fixFirst.why',
        '.review.fixFirst.links.0.sentence',
        '.review.thenFix.1.links.0.sentence',
        '.anchors.a3.to_ms',
        '.anchors.a5.to_ms',
      ].sort(),
    );
  });

  it('has the gap and the full-bar window agree: 2:44 to 2:58 is 14 seconds', () => {
    expect(answer.review.fixFirst?.why).toContain('At 2:44–2:58 you sat on a full bar');
    expect(answer.review.fixFirst?.why).toContain('so that one gap cost you two Bloodthirsts.');
    expect(answer.review.thenFix[1]?.title).toBe('Rage full for 14 seconds');
    expect(answer.anchors.a5?.to_ms).toBe(179000);
    expect((answer.anchors.a5?.to_ms ?? 0) - (answer.anchors.a5?.at_ms ?? 0)).toBe(14_000);
  });
});

describe('coachSampleView', () => {
  const view = coachSampleView(answer, iconOf);

  it('names the persona and the context from the catalogue', () => {
    expect(view.title).toBe('Your Fury coach');
    expect(view.classSlug).toBe('warrior');
    expect(view.context).toBe(plansPageCopy.coachSample.context('Fury', 'Warrior', 'Molten Core', 40, true));
    expect(view.context).toBe('Fury Warrior · Molten Core · 40-player kill');
  });

  it('carries the review words verbatim and drops drill, links and unchecked lines', () => {
    expect(view.opening).toBe(answer.review.opening);
    expect(view.fixFirst?.title).toBe(answer.review.fixFirst?.title);
    expect(view.fixFirst?.why).toBe(answer.review.fixFirst?.why);
    expect(view.then.map((item) => item.title)).toEqual(answer.review.thenFix.map((item) => item.title));
    expect(view.doingWell.map((item) => item.text)).toEqual(answer.review.doingWell.map((item) => item.text));
    expect(JSON.stringify(view)).not.toContain(answer.review.fixFirst?.drillText);
  });

  it('takes each icon from the first anchor that carries a spell', () => {
    expect(view.fixFirst?.icon).toBe('/data/b/icons/23894.webp');
    expect(view.then[0]?.icon).toBeNull();
    expect(view.then[1]?.icon).toBeNull();
    expect(view.doingWell[0]?.icon).toBeNull();
    expect(view.doingWell[1]?.icon).toBe('/data/b/icons/6673.webp');
  });

  it('says wipe for a wipe', () => {
    expect(coachSampleView({ ...answer, kill: false }, iconOf).context).toContain('40-player wipe');
  });

  it('fails loudly for a spec the catalogue lacks', () => {
    expect(() => coachSampleView({ ...answer, spec: 'Nonesuch' }, iconOf)).toThrow(/Nonesuch/);
  });
});
