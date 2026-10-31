// W.12 ruling 1: the planner never sends another build to a character. Three cases against the
// fixture build with a stubbed account: another class's build (the button names the mismatch and
// stays off), another player's build of the same class (it says what it replaces and asks once),
// and the character's own build (it sends as before). "Sim this build" sits in the same card, under
// the send, in view at 1280 x 900.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';

import { meAddonFixture } from '../../src/fixtures/me-addon';

const ok = (data: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

const KEY = 'us/pvp/obnoxious-yell';
const SAVED = { id: 'k7x2qm4a', url: 'https://foreversixty.gg/b/k7x2qm4a' };
const BUILD = '1.60.1.70450';
/** A named warrior build that nobody on this account holds. */
const ASHENVEIL = `FS1:${BUILD}:warrior:orc:0/5530515/0:head=12640|level=21|who=Ashenveil:Realm`;

function me(characterClass: string) {
  const own = meAddonFixture.characters[0];
  return {
    ...meAddonFixture,
    characters: [
      {
        ...own,
        key: KEY,
        region: 'us',
        ruleset: 'pvp',
        name: 'Obnoxious Yell',
        class: characterClass,
        spec: undefined,
        level: 21,
        build: {
          ...own.build,
          talents: { trees: ['0', '5530515', '0'], points: [0, 29, 0] },
        },
      },
    ],
  };
}

async function signIn(page: Page, characterClass: string) {
  await page.route('**/v1/me', (route) => route.fulfill(ok(me(characterClass))));
  await page.route('**/v1/devices', (route) => route.fulfill(ok([{ id: 'd1', name: 'Desk' }])));
}

test.describe('the planner never sends another build to a character', () => {
  test.skip(({ isMobile }) => isMobile, 'The card is asserted at desktop width');
  test.use({ viewport: { width: 1280, height: 900 } });

  test('another class’s build: the button names the mismatch and stays off', async ({ page }) => {
    await signIn(page, 'Rogue');
    let saved = false;
    await page.route('**/v1/builds', (route) => {
      saved = true;
      return route.fulfill(ok(SAVED, 201));
    });
    await page.goto('/planner?class=warrior');
    await page.getByTestId('talent-1001').click();
    const action = page.getByTestId('character-header-action');
    await expect(action).toHaveText('This is a Warrior build; Obnoxious Yell is a Rogue');
    await expect(action).toBeDisabled();
    await action.click({ force: true });
    expect(saved).toBe(false);
  });

  test('another player’s build of the class says whose it is, what it replaces, and asks once', async ({
    page,
  }) => {
    await signIn(page, 'Warrior');
    const saves: unknown[] = [];
    const queued: unknown[] = [];
    await page.route('**/v1/builds', (route) => {
      saves.push(route.request().postDataJSON());
      return route.fulfill(ok(SAVED, 201));
    });
    await page.route('**/v1/addon/inbox', (route) => {
      queued.push(route.request().postDataJSON());
      return route.fulfill(ok({ build_id: SAVED.id }, 201));
    });
    await page.goto(`/planner?code=${encodeURIComponent(ASHENVEIL)}`);
    await expect(page.getByTestId('planner-build-origin')).toHaveText(/^Ashenveil's build · .*Warrior$/);
    const action = page.getByTestId('character-header-action');
    await expect(action).toHaveText('Send this build to Obnoxious Yell (replaces their 29 points)');

    // The first press asks; cancelling sends nothing.
    await action.click();
    await expect(page.getByTestId('character-header-note')).toHaveText(
      "Send Ashenveil's build to Obnoxious Yell?",
    );
    await expect(action).toHaveText('Yes, send it');
    await page.getByTestId('planner-send-cancel').click();
    await expect(action).toHaveText('Send this build to Obnoxious Yell (replaces their 29 points)');
    expect(saves).toHaveLength(0);

    // Asked again, the yes sends it, once.
    await action.click();
    await action.click();
    await expect(action).toHaveText('Sent');
    expect(saves).toHaveLength(1);
    expect(queued).toEqual([{ build_id: SAVED.id, character_key: KEY }]);
  });

  test('the character’s own build sends without asking, and "Sim this build" sits under it', async ({
    page,
  }) => {
    await signIn(page, 'Warrior');
    await page.route('**/v1/builds', (route) => route.fulfill(ok(SAVED, 201)));
    await page.route('**/v1/addon/inbox', (route) => route.fulfill(ok({ build_id: SAVED.id }, 201)));
    await page.goto('/planner');
    await page.getByTestId('talent-1001').click();
    const action = page.getByTestId('character-header-action');
    await expect(action).toHaveText('Send this build to Obnoxious Yell');
    await expect(page.getByTestId('planner-build-origin')).toHaveCount(0);

    // "Sim this build": in the card, under Send, inside the first screen at 1280 x 900.
    const sim = page.getByTestId('planner-sim-link');
    await expect(sim).toHaveText('Sim this build');
    const sendBox = await action.boundingBox();
    const simBox = await sim.boundingBox();
    expect(sendBox).not.toBeNull();
    expect(simBox).not.toBeNull();
    expect(simBox!.y).toBeGreaterThan(sendBox!.y + sendBox!.height - 1);
    expect(simBox!.y + simBox!.height).toBeLessThan(900);
    expect(Math.abs(simBox!.x - sendBox!.x)).toBeLessThan(2);
    expect(Math.abs(simBox!.width - sendBox!.width)).toBeLessThan(2);
    await expect(page.getByTestId('planner-sim-link')).toHaveCount(1);

    await action.click();
    await expect(action).toHaveText('Sent');
  });
});
