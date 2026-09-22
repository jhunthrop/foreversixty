// web/tests/e2e/support/character-fixture.ts
// The GET /v1/characters/<region>/<ruleset>/<name> answer the character-page specs stub.
// One copy, shared: character-phone.spec.ts measures this page's layout at 360px and
// character-retry.spec.ts drives its failed state, and a fixture that drifted between the
// two would let one of them pass against a shape the other no longer sees.
import type { Page } from '@playwright/test';

export const CHARACTER_URL = '**/v1/characters/us/hardcore/elyra-duskvale';
export const CHARACTER_PATH = '/character/us/hardcore/elyra-duskvale';

export const CHARACTER = {
  ok: true,
  data: {
    character: { name: 'Elyra Duskvale', region: 'us', ruleset: 'hardcore', class: 'Priest' },
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
