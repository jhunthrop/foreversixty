// web/tests/e2e/support/tiers-parity.ts
// The tier list states of the parity gate (design/parity.md, design/specs/2026-11-06-tiers.md):
// the pages are prerendered from the published BiS files ws_tiers.py ranks its mocks from, so a
// state only chooses the route, the faction and the session. Signed in is the mock's sample
// character, Thandrel (Fury Warrior, 60): the header and the answer name the same character.
import type { Page } from '@playwright/test';
import { expect } from './test';
import { armoryPointer, signInWith, signedOut, storePointer, type FixtureCharacter } from './selector';

export type TiersParityState = 'dps' | 'dps-you' | 'dps-horde' | 'healer' | 'tank' | 'tank-you';

/** design/mocks/ws_tiers.py `YOU`. */
export const THANDREL: FixtureCharacter = {
  name: 'Thandrel',
  class: 'Warrior',
  spec: 'Fury',
  level: 60,
  faction: 'alliance',
};

const PATH: Record<TiersParityState, string> = {
  dps: '/tiers',
  'dps-you': '/tiers',
  'dps-horde': '/tiers?faction=horde',
  healer: '/tiers/healer',
  tank: '/tiers/tank',
  'tank-you': '/tiers/tank',
};

export async function openTiersState(page: Page, id: TiersParityState): Promise<void> {
  const signedIn = id.endsWith('-you');
  if (signedIn) {
    await signInWith(page, [THANDREL]);
    await storePointer(page, armoryPointer(THANDREL));
  } else {
    await signedOut(page);
  }
  await page.goto(PATH[id]);
  const faction = id === 'dps-horde' ? 'horde' : 'alliance';
  await expect(page.getByTestId(`tier-faction-panel-${faction}`).getByTestId('tier-list')).toBeVisible();
  if (signedIn) await expect(page.getByTestId('tier-answer-body')).toContainText('Fury Warrior');
}
