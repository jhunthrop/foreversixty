// web/src/fixtures/coach/fixtures.test.ts
// The coach's own output, as the e2e stubs and the review captures serve it. The Go side writes
// these files (api/internal/coach TestWebFixtures_Write) and fails when they drift
// (TestWebFixtures_NoDrift); this side proves the web reads every one of them as the API's shape
// and that each state carries the copy the panel expects.
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { coachCopy } from '../../lib/coach/copy';
import { normalizeFightCoach, normalizeRecent } from '../../lib/coach/normalize';
import { viewForAnswer } from '../../lib/coach/states';
import type { CoachFull } from '../../lib/coach/types';

const DIR = path.dirname(fileURLToPath(import.meta.url));
const read = (name: string): unknown => JSON.parse(readFileSync(path.join(DIR, `${name}.json`), 'utf8'));
const full = (name: string): CoachFull => {
  const answer = normalizeFightCoach(read(name));
  if (answer === null || answer.access === 'teaser') throw new Error(`${name} is not a full answer`);
  return answer;
};

const ANSWERS = readdirSync(DIR)
  .filter((file) => file.endsWith('.json') && !file.startsWith('last-raid'))
  .map((file) => file.replace(/\.json$/, ''));

describe('coach fixtures (the coach output the e2e and captures stub)', () => {
  it('every answer file reads as the API shape', () => {
    expect(ANSWERS.length).toBeGreaterThan(10);
    for (const name of ANSWERS) expect(normalizeFightCoach(read(name)), name).not.toBeNull();
  });

  it('each state file lands in its panel state', () => {
    const states: Record<string, string> = {
      'fury-model': 'ready_model',
      'fury-rules': 'ready_rules',
      'fury-writing': 'writing',
      'fury-busy': 'busy',
      'fury-preparing': 'preparing',
      'fury-no-findings': 'no_findings',
      'fury-raid-buffs-only': 'raid_buffs_only',
      'fury-not-reviewable-trash': 'not_reviewable',
      'fury-teaser-several': 'teaser',
      'fury-teaser-empty': 'teaser_empty',
      'holy-model': 'ready_model',
    };
    for (const [name, state] of Object.entries(states)) {
      expect(viewForAnswer(normalizeFightCoach(read(name))).state, name).toBe(state);
    }
  });

  it('no findings opens with rulesNothing; the rule-only review with the boss', () => {
    expect(full('fury-no-findings').review.opening).toBe(coachCopy.rulesNothing);
    expect(full('fury-rules').review.opening).toContain('Onyxia');
  });

  it('the healer carries its own spec, strengths and basis', () => {
    const holy = full('holy-model');
    expect(holy.spec).toBe('Holy');
    expect(holy.role).toBe('healer');
    expect(holy.basis_labels.join(' ')).not.toMatch(/Fury|Warrior/);
    expect(holy.review.doingWell.map((item) => item.text).join(' ')).not.toMatch(/Battle Shout|rotation/);
  });

  it('no time in a review lies after the 3:52 kill', () => {
    const clocks = /\b(\d+):(\d{2})\b/g;
    for (const name of ANSWERS) {
      const text = JSON.stringify((read(name) as { review?: unknown }).review ?? {});
      for (const [, m, s] of text.matchAll(clocks))
        expect(Number(m) * 60 + Number(s), name).toBeLessThanOrEqual(232);
    }
  });

  it('the last-raid nights read as the account line expects', () => {
    const night = normalizeRecent({ last_raid: read('last-raid-molten-core'), pending: false, reviews: [] });
    expect(night.last_raid?.zone).toBe('Molten Core');
    expect(night.last_raid?.first).toMatchObject({ on: 5, of: 8 });
    const onyxia = normalizeRecent({ last_raid: read('last-raid-onyxia'), pending: false, reviews: [] });
    expect(onyxia.last_raid?.first).toMatchObject({ on: 1, of: 1 });
  });
});
