// web/tests/e2e/support/rankings-parity.ts
// The /rankings and /rankings/<slug> states of the parity gate (design/parity.md), served from the
// rankings stubs (support/rankings-fixture.ts), the same files design/mocks/ws_rankings.py draws
// through rankings_facts.mjs.
import type { Page } from '@playwright/test';
import { expect } from './test';
import { EMPTY_BOARD, guildSpeedRows, stubRankedEncounters, stubRankingsApi } from './rankings-fixture';
import { OBNOXIOUS, armoryPointer, signInWith, signedOut, storePointer } from './selector';

export type RankingsParityState =
  | 'rankings'
  | 'rankings-empty'
  | 'rankings-board'
  | 'rankings-board-empty'
  | 'rankings-guilds'
  | 'rankings-you'
  | 'rankings-off'
  | 'rankings-unranked';

type SignedOutState =
  'rankings' | 'rankings-empty' | 'rankings-board' | 'rankings-board-empty' | 'rankings-guilds';

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

const SIGNED_OUT: Record<SignedOutState, { path: string; ready: string }> = {
  rankings: { path: '/rankings', ready: 'rankings-encounters' },
  'rankings-empty': { path: '/rankings', ready: 'rankings-no-encounters' },
  'rankings-board': { path: '/rankings/warden-kelthas', ready: 'ranking-1' },
  'rankings-board-empty': {
    path: '/rankings/warden-kelthas?class=warrior&spec=Fury',
    ready: 'rankings-empty',
  },
  'rankings-guilds': { path: '/rankings/warden-kelthas?board=guild&kind=speed', ready: 'guild-ranking-1' },
};

async function openSignedOut(page: Page, id: SignedOutState): Promise<void> {
  await signedOut(page);
  await stubRankingsApi(page);
  // A later route wins: these answer over the stub's own.
  if (id === 'rankings-empty') await stubRankedEncounters(page, []);
  if (id === 'rankings-board-empty') {
    await page.route('**/v1/rankings?**', (route) => route.fulfill(envelope(EMPTY_BOARD)));
  }
  if (id === 'rankings-guilds') {
    await page.route('**/v1/rankings/guilds?**', (route) =>
      route.fulfill(envelope({ rows: guildSpeedRows() })),
    );
  }
  await page.goto(SIGNED_OUT[id].path);
  await expect(page.getByTestId(SIGNED_OUT[id].ready)).toBeVisible();
}

/** Opens one state, every read answered from the stubs. */
export async function openRankingsState(page: Page, id: RankingsParityState): Promise<void> {
  if (id in SIGNED_OUT) return openSignedOut(page, id as SignedOutState);
  await signInWith(page, [OBNOXIOUS]);
  await storePointer(page, armoryPointer(OBNOXIOUS));
  await stubRankingsApi(page, {
    ownOffPage: id === 'rankings-off',
    unranked: id === 'rankings-unranked',
  });
  // Signed in, /rankings opens the encounter the character last killed (Onyxia).
  await page.goto('/rankings');
  await expect(page.getByTestId('ranking-1')).toBeVisible();
  await expect(
    page.getByTestId(id === 'rankings-unranked' ? 'rankings-answer-unranked' : 'rankings-answer-rank'),
  ).toBeVisible();
}
