// web/tests/e2e/character-identity.spec.ts
// Who the character is, in the page head (design/specs/2026-11-06-character.md C3): the crest,
// the name, then the fact line "Level 60 Night Elf", "Discipline Priest", the talent split as a
// planner link, the guild as a link to its page, the realm and the "Hardcore US" chip; the
// owner alone gets the sync line, which re-reads its age each minute. A null field is left out,
// never a placeholder. The API is stubbed from the shared fixture (support/character-fixture.ts);
// the rating read is answered 404 so only the head is in play.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { meAddonFixture } from '../../src/fixtures/me-addon';
import { CHARACTER, CHARACTER_PATH, CHARACTER_URL, OWNER_SCORE } from './support/character-fixture';

const NOT_FOUND = {
  status: 404,
  contentType: 'application/json',
  body: JSON.stringify({ ok: false, data: null, error: 'not_found', request_id: 'r' }),
};

const ok = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

async function stub(page: Page, data: unknown): Promise<void> {
  await page.route(CHARACTER_URL, (route) => route.fulfill(ok(data)));
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(NOT_FOUND));
}

/** The me fixture's addon-synced character, made the owner of this page's character. */
function ownerMe(capturedAt?: string): typeof meAddonFixture {
  const own = meAddonFixture.characters[0];
  if (own.build === undefined) throw new Error('the me fixture character has no build');
  return {
    ...meAddonFixture,
    characters: [
      {
        ...own,
        key: 'us/hardcore/elyra-duskvale',
        region: 'us',
        ruleset: 'hardcore',
        name: 'Elyra Duskvale',
        class: 'Priest',
        build: capturedAt === undefined ? own.build : { ...own.build, captured_at: capturedAt },
      },
    ],
  };
}

const fact = (page: Page, key: string) => page.getByTestId(`character-fact-${key}`);

test('the head says who the character is: level and race, spec, split, guild, realm and the chip', async ({
  page,
}) => {
  await stub(page, { ...CHARACTER.data, ...OWNER_SCORE });
  await page.goto(CHARACTER_PATH);

  const head = page.getByTestId('character-head');
  await expect(head.getByTestId('page-title')).toHaveText('Elyra Duskvale');
  await expect(head.locator('img[src*="/icons/hd/crests/"]')).toHaveAttribute(
    'src',
    /\/icons\/hd\/crests\/priest\.webp$/,
  );
  await expect(fact(page, 'level')).toHaveText('Level 60 Night Elf');
  await expect(fact(page, 'spec')).toHaveText('Discipline Priest');
  await expect(fact(page, 'split')).toHaveText('31/20/0');
  await expect(fact(page, 'split')).toHaveAttribute('href', '/planner');
  await expect(fact(page, 'guild')).toHaveText('<Cx66 Olympus>');
  await expect(fact(page, 'guild')).toHaveAttribute('href', '/guild/us/hardcore/cx66-olympus');
  await expect(fact(page, 'realm')).toHaveText('Doomhowl');
  await expect(fact(page, 'chip')).toHaveText('Hardcore US');
  // A visitor gets no sync line.
  await expect(page.getByTestId('character-sync')).toHaveCount(0);
  await expect(page.getByTestId('character')).not.toContainText(/NaN|null|undefined/);
});

test('a character nobody knows the identity of: no guild, realm, level or split, never a placeholder', async ({
  page,
}) => {
  await stub(page, {
    ...CHARACTER.data,
    ...OWNER_SCORE,
    character: { ...CHARACTER.data.character, guild: null, level: null, race: null, realm: null },
    builds_seen: [],
  });
  await page.goto(CHARACTER_PATH);

  await expect(fact(page, 'spec')).toHaveText('Discipline Priest');
  await expect(fact(page, 'chip')).toHaveText('Hardcore US');
  for (const key of ['level', 'split', 'guild', 'realm']) await expect(fact(page, key), key).toHaveCount(0);
  await expect(page.getByTestId('character')).not.toContainText(/null|undefined|Unknown/);
});

test('the owner’s head carries the sync line after the facts, the split to their own planner', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) => route.fulfill(ok(ownerMe())));
  await stub(page, { ...CHARACTER.data, ...OWNER_SCORE });
  await page.goto(CHARACTER_PATH);

  const sync = page.getByTestId('character-sync');
  // The fixture's build was captured minutes before the worker loaded it.
  await expect(sync).toHaveText(/^Addon · \d+ minutes ago$/);
  await expect(page.getByTestId('character-head').getByTestId('character-sync')).toBeVisible();
  await expect(fact(page, 'split')).toHaveAttribute(
    'href',
    /^\/planner\?character=us%2Fhardcore%2Felyra-duskvale$/,
  );
  await expect(page.getByTestId('character')).not.toContainText('Last seen in game');
});

test('a stale addon sync still reads its age in days', async ({ page }) => {
  const stale = new Date(Date.now() - 19 * 86_400_000).toISOString();
  await page.route('**/v1/me', (route) => route.fulfill(ok(ownerMe(stale))));
  await stub(page, { ...CHARACTER.data, ...OWNER_SCORE });
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-sync')).toHaveText(/^Addon · 19 days ago/);
});

test('the owner’s sync age ticks each minute while the tab lives', async ({ page }) => {
  const now = new Date('2026-10-12T01:34:53Z');
  await page.clock.install({ time: now });
  await page.route('**/v1/me', (route) =>
    route.fulfill(ok(ownerMe(new Date(now.getTime() - 4 * 60_000).toISOString()))),
  );
  await stub(page, { ...CHARACTER.data, ...OWNER_SCORE });
  await page.goto(CHARACTER_PATH);

  const sync = page.getByTestId('character-sync');
  await expect(sync).toHaveText('Addon · 4 minutes ago');
  await page.clock.runFor(60_000);
  await expect(sync).toHaveText('Addon · 5 minutes ago');
});

test('the crest sits on the name’s line in the same place for a visitor and the owner', async ({ page }) => {
  /** The crest's offset from the name: the owner's extra sync line must not move one off the other. */
  const crestFromTitle = async () => {
    const head = page.getByTestId('character-head');
    const crest = head.locator('img[src*="/icons/hd/crests/"]');
    await expect(crest).toBeVisible();
    const [c, t] = await Promise.all([crest.boundingBox(), head.getByTestId('page-title').boundingBox()]);
    if (c === null || t === null) throw new Error('the crest or the name is missing');
    return { x: c.x, dx: t.x - c.x };
  };
  await stub(page, { ...CHARACTER.data, ...OWNER_SCORE });
  await page.goto(CHARACTER_PATH);
  const visitor = await crestFromTitle();

  await page.route('**/v1/me', (route) => route.fulfill(ok(ownerMe())));
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('character-sync')).toBeVisible();
  const owner = await crestFromTitle();

  expect(Math.abs(owner.x - visitor.x)).toBeLessThan(1);
  expect(Math.abs(owner.dx - visitor.dx)).toBeLessThan(1);
});

test.describe('a 40-character guild with no realm at 320px', () => {
  test.use({ viewport: { width: 320, height: 640 } });

  test('the head wraps inside the page, nothing scrolls sideways', async ({ page }) => {
    await stub(page, {
      ...CHARACTER.data,
      ...OWNER_SCORE,
      character: {
        ...CHARACTER.data.character,
        guild: { name: 'The Exceedingly Long Guild Name Of Doom', rank: 'member' },
        realm: null,
      },
    });
    await page.goto(CHARACTER_PATH);

    await expect(fact(page, 'guild')).toHaveText('<The Exceedingly Long Guild Name Of Doom>');
    const scrollWidth = await page.evaluate(() => document.documentElement.scrollWidth);
    expect(scrollWidth).toBeLessThanOrEqual(320);
  });
});

test.describe('with JavaScript off', () => {
  test.use({ javaScriptEnabled: false });

  test('one line says why, with the Planner link', async ({ page }) => {
    await page.goto(CHARACTER_PATH);
    const line = page.getByTestId('character-noscript');
    await expect(line).toBeVisible();
    await expect(line).toHaveText(
      'Forever Sixty needs JavaScript to show this character. Turn it on and reload, or open the Planner.',
    );
    await expect(line.getByRole('link', { name: 'Planner' })).toHaveAttribute('href', '/planner');
  });

  // CharacterNoScript.astro hides the skeleton with !important: the skeleton is a ui/Page, whose
  // scoped display rule would otherwise win over the attribute selector.
  test('the loading skeleton goes', async ({ page }) => {
    await page.goto(CHARACTER_PATH);
    await expect(page.getByTestId('character-skeleton')).toBeHidden();
  });
});
