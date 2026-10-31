// web/tests/e2e/support/character-fixture.ts
// The GET /v1/characters/<region>/<ruleset>/<name> answer the character-page specs stub.
// One copy, shared: character-phone.spec.ts measures this page's layout at 360px and
// character-retry.spec.ts drives its failed state, and a fixture that drifted between the
// two would let one of them pass against a shape the other no longer sees.
import type { Page } from '@playwright/test';
import { characterHistoryFixture } from '../../../src/fixtures/character-history';
import type { CharacterHistory } from '../../../src/lib/character/history-api';
import { meAddonFixture } from '../../../src/fixtures/me-addon';
import { readFileSync, existsSync } from 'node:fs';
import {
  GEAR_FIXTURE_ITEMS,
  GEAR_FIXTURE_LONGEST,
  GEAR_PANEL_FIXTURE,
} from '../../../src/fixtures/character-gear';

/** The gear panel (api/openapi.yaml `CharacterGear`, Task L8.6): every row state, from
 *  src/fixtures/character-gear.ts. The shared CHARACTER below leaves it out, as an API that
 *  predates it does; character-gear.spec.ts adds it. */
export const GEAR = GEAR_PANEL_FIXTURE;

/** The paperdoll's worst case: the longest names in every slot. */
export const GEAR_LONGEST = { gear: GEAR_FIXTURE_LONGEST, gear_gap_total: null };

/** The build's items/priest.json for the fixture's ids: the real rows (quality, icon,
 *  two-handed, stats), so the names take their colours and icons and the tooltip has a row. */
export const GEAR_ITEMS = { build: 'fixture', class_slug: 'priest', items: GEAR_FIXTURE_ITEMS };

const ICONS = new URL('../../../../data/builds/1.60.1.70291/icons/', import.meta.url);

/** Serves an item icon from the repository's real build, which the fixture build does not
 *  carry; 404 for one it lacks, and for each name in `missing`. */
export async function routeGearIcons(page: Page, missing: readonly string[] = []): Promise<void> {
  await page.route('**/data/*/icons/*.webp', (route) => {
    const name = new URL(route.request().url()).pathname.split('/').pop() ?? '';
    const file = new URL(name, ICONS);
    return existsSync(file) && !missing.includes(name)
      ? route.fulfill({ status: 200, contentType: 'image/webp', body: readFileSync(file) })
      : route.fulfill({ status: 404, body: '' });
  });
}

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

/** Every best and history row carries `execution_score` (fix round 1 item 3), as the API
 *  always sends it (a number, or null for an unscored fight): a row without one read "NaN%". */
export const CHARACTER = {
  ok: true,
  data: {
    // The owning account shows the Supporter mark (api `Character.supporter`); a spec sets it.
    supporter: false,
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
      realm: 'Doomhowl',
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
        execution_score: 0.94,
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
        execution_score: 0.81,
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
        execution_score: 0.94,
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
        execution_score: 0.81,
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

/** A `/v1/me` answer whose one character is the fixture's own (Elyra Duskvale, hardcore US), so
 *  the page reads as the owner's: the 404 no-logs card with its upload action, not the visitor's
 *  "No character called ... yet" page. */
export function ownerMeAnswer(): { ok: true; data: unknown; error: null; request_id: string } {
  const own = meAddonFixture.characters[0];
  return {
    ok: true,
    error: null,
    request_id: 'r',
    data: {
      ...meAddonFixture,
      characters: [
        {
          ...own,
          key: 'us/hardcore/elyra-duskvale',
          region: 'us',
          ruleset: 'hardcore',
          name: 'Elyra Duskvale',
          class: 'Priest',
        },
      ],
    },
  };
}

/** Routes `/v1/me` to `ownerMeAnswer()`. */
export async function routeOwner(page: Page): Promise<void> {
  await page.route('**/v1/me', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(ownerMeAnswer()) }),
  );
}

/** Answers the owner-only history route (experience spec 2026-10-24-character-history section
 *  10) with `history`, or with a status (402 lapsed, 404 not the owner, 500). */
export async function routeCharacterHistory(
  page: Page,
  history: CharacterHistory | number = characterHistoryFixture(),
): Promise<void> {
  await page.route(`${CHARACTER_URL}/history`, (route) =>
    typeof history === 'number'
      ? route.fulfill({
          status: history,
          contentType: 'application/json',
          body: JSON.stringify({ ok: false, data: null, error: { message: 'history failed' } }),
        })
      : route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({ ok: true, data: history, error: null }),
        }),
  );
}
