import { describe, expect, it } from 'vitest';
import { coachModelReview, coachTeaser } from '../../test-support/coach-fixtures';
import { normalizeFightCoach, normalizeHistory, normalizeNightLine, normalizeRecent } from './normalize';

describe('normalizeFightCoach', () => {
  it('rejects anything without a known access', () => {
    expect(normalizeFightCoach(null)).toBeNull();
    expect(normalizeFightCoach([])).toBeNull();
    expect(normalizeFightCoach({ access: 'root' })).toBeNull();
  });

  it('round-trips a teaser and a full review, dropping unknown fields', () => {
    expect(normalizeFightCoach({ ...coachTeaser(), extra: 1 })).toEqual(coachTeaser());
    expect(normalizeFightCoach({ ...coachModelReview(), extra: 1 })).toEqual(coachModelReview());
  });

  it('keeps a known account cap and drops an unknown one', () => {
    const cap = { period: 'month', until: '2026-11-01' };
    expect(normalizeFightCoach({ ...coachModelReview(), busy: true, cap })).toMatchObject({
      busy: true,
      cap,
    });
    expect(normalizeFightCoach({ ...coachModelReview(), cap: { period: 'year' } })).toMatchObject({
      cap: null,
    });
    expect(normalizeFightCoach({ ...coachModelReview(), cap: undefined })).toMatchObject({ cap: null });
  });

  it('a teaser without a spell id keeps null', () => {
    const got = normalizeFightCoach({ access: 'teaser', fix_first: { title: 't', code: 'c' } });
    expect(got).toMatchObject({ fix_first: { spell_id: null }, finding_count: 0, status: 'ready' });
  });

  it('keeps string links as ids and object links with their sentence', () => {
    const got = normalizeFightCoach({
      access: 'full',
      review: { thenFix: [{ title: 't', links: ['l1', { id: 'l2', sentence: 'S.' }, 4] }] },
    });
    expect(got).toMatchObject({
      review: {
        thenFix: [
          {
            links: [
              { id: 'l1', sentence: null },
              { id: 'l2', sentence: 'S.' },
            ],
          },
        ],
      },
    });
  });

  it('keeps the debuff flag on an aura anchor and drops unknown anchor kinds', () => {
    const got = normalizeFightCoach({
      access: 'full',
      anchors: { a: { kind: 'aura', spell_id: 3, debuff: true }, b: { kind: 'nope' } },
    });
    expect(got).toMatchObject({ anchors: { a: { kind: 'aura', spell_id: 3, debuff: true } } });
    expect(Object.keys((got as { anchors: object }).anchors)).toEqual(['a']);
  });

  it("keeps a debuff anchor's NPC target", () => {
    const boss = 'Creature-0-1-1084-1-10184-0000000001';
    const got = normalizeFightCoach({
      access: 'full',
      anchors: {
        a: { kind: 'aura', spell_id: 3, debuff: true, target_guid: boss },
        b: { kind: 'aura', spell_id: 4, target_guid: '' },
      },
    });
    expect(got).toMatchObject({ anchors: { a: { target_guid: boss } } });
    expect((got as { anchors: Record<string, object> }).anchors.b).not.toHaveProperty('target_guid');
  });
});

describe('normalizeRecent and normalizeHistory', () => {
  it('a non-object is an empty recent', () => {
    expect(normalizeRecent('x')).toEqual({ last_raid: null, pending: false, reviews: [] });
  });

  it('history items carry boss, kill and the {title} fix first', () => {
    const page = normalizeHistory({
      reviews: [{ boss: 'Ragnaros', kill: false, fix_first: { title: 'T', code: 'x' } }],
      next_cursor: '',
    });
    expect(page.next_cursor).toBeNull();
    expect(page.reviews[0]).toMatchObject({ boss: 'Ragnaros', kill: false, fix_first: { title: 'T' } });
    expect(page.reviews[0]?.fix_first).toEqual({ title: 'T' });
  });
});

describe('normalizeNightLine', () => {
  it('needs a stem', () => {
    expect(normalizeNightLine({ on: 1 })).toBeNull();
    expect(normalizeNightLine({ stem: 's', on: 1, of: 2 })).toMatchObject({ stem: 's', on: 1, of: 2 });
  });
});
