// web/src/lib/billing/coach-sample.test.ts
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

describe('coachSampleView', () => {
  const view = coachSampleView(answer, iconOf);

  it('names the persona and the context from the catalogue', () => {
    expect(view.title).toBe('Your Fury coach');
    expect(view.classSlug).toBe('warrior');
    expect(view.context).toBe(plansPageCopy.coachSample.context('Fury', 'Warrior', 40, true));
    expect(view.context).toBe('Fury Warrior · 40-player kill');
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
