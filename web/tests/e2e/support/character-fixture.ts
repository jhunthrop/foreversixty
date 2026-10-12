// web/tests/e2e/support/character-fixture.ts
// The GET /v1/characters/<region>/<ruleset>/<name> answer the character-page specs stub.
// One copy, shared: character-phone.spec.ts measures this page's layout at 360px and
// character-retry.spec.ts drives its failed state, and a fixture that drifted between the
// two would let one of them pass against a shape the other no longer sees.
import type { Page } from '@playwright/test';

export const CHARACTER_URL = '**/v1/characters/us/hardcore/elyra-duskvale';
export const CHARACTER_PATH = '/character/us/hardcore/elyra-duskvale';

/** Five dates of Forever Score, oldest first, rising from rare into epic. */
export const SCORE_HISTORY = [
  { date: '2026-10-12', score: 1240, tier: 'rare' },
  { date: '2026-10-19', score: 1380, tier: 'rare' },
  { date: '2026-10-26', score: 1520, tier: 'epic' },
  { date: '2026-11-01', score: 1690, tier: 'epic' },
  { date: '2026-11-04', score: 1830, tier: 'epic' },
];

export const CHARACTER = {
  ok: true,
  data: {
    // Who the character is (api/openapi.yaml `CharacterIdentity`, Task L8.4): each field is
    // always sent, null when unknown -- a spec that needs a null overrides it.
    character: {
      name: 'Elyra Duskvale',
      region: 'us',
      ruleset: 'hardcore',
      class: 'Priest',
      guild: { name: 'Cx66 Olympus', rank: 'officer' },
      level: 60,
      race: 'Night Elf',
      realm: 'Classic Beta PvP',
    },
    best: [
      {
        encounter: 'Warden Kelthas',
        encounter_id: 9001,
        difficulty: 8,
        metric: 'hps',
        value: 1840,
        percentile: 96.2,
        spec: 'Discipline',
        fought_at: '2026-12-09T22:10:00Z',
        report_id: 'fixture2abcd',
        fight_index: 3,
      },
      {
        encounter: 'Deep Warden',
        encounter_id: 9002,
        difficulty: 8,
        metric: 'hps',
        value: 1420,
        percentile: 71,
        spec: 'Discipline',
        fought_at: '2026-12-08T20:00:00Z',
        report_id: 'otherreport1',
        fight_index: 1,
      },
    ],
    history: [
      {
        encounter: 'Warden Kelthas',
        encounter_id: 9001,
        difficulty: 8,
        metric: 'hps',
        value: 1840,
        percentile: 96.2,
        spec: 'Discipline',
        fought_at: '2026-12-09T22:10:00Z',
        report_id: 'fixture2abcd',
        fight_index: 3,
      },
      {
        encounter: 'Deep Warden',
        encounter_id: 9002,
        difficulty: 8,
        metric: 'hps',
        value: 1420,
        percentile: 71,
        spec: 'Discipline',
        fought_at: '2026-12-08T20:00:00Z',
        report_id: 'otherreport1',
        fight_index: 1,
      },
    ],
    builds_seen: [{ talent_split: '31/20/0', spec: 'Discipline', first_seen: '2026-12-09T22:10:00Z' }],
    // The score over time (api/openapi.yaml `ScorePoint`): one point per UTC date, oldest first.
    // The page draws it only beside a score, which this fixture leaves out; character-score
    // supplies one.
    score_history: SCORE_HISTORY,
  },
  error: null,
  request_id: 'r',
};

/** Answers the character fetch with the fixture above. */
export async function stubCharacter(page: Page): Promise<void> {
  await page.route(CHARACTER_URL, (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(CHARACTER) }),
  );
}
