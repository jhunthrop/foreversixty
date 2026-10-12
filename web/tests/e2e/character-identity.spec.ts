// web/tests/e2e/character-identity.spec.ts
// Task L8.4 and its fix round 2: who the character is and its score over time, on the profile.
// The header reads "Discipline Priest · 31/20/0 · <Cx66 Olympus> · Classic Beta PvP" with
// "Level 60 Night Elf" under it, the "Hardcore US" chip ending the last line, and for the owner
// the Character header's sync line third; the tooltip mock reads as the 1.60 client prints it;
// the Forever Score card draws score_history as a sparkline across its content column with a
// caption. A null field is left out, never a placeholder, and a 404 is still the no-logs page.
// The API is stubbed with page.route from the shared fixture (support/character-fixture.ts),
// which carries the identity and the history; the panel's execution-rating read is answered 404
// so only these parts are in play.
import { expect, test, type Page } from '@playwright/test';
import labels from '../../../docs/contracts/fixtures/score-labels.json' with { type: 'json' };
import { meAddonFixture } from '../../src/fixtures/me-addon';
import { CHARACTER, CHARACTER_PATH, CHARACTER_URL, SCORE_HISTORY } from './support/character-fixture';

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

const SCORE = {
  score: 1830,
  tier: 'epic',
  as_of: '2026-11-04',
  components: { gear: 412, experience: 620, parses: 540, attendance: 180, discipline: 78 },
  experience: { dungeons: 14, raid_kills: 0, alts: 2 },
};

async function stub(page: Page, data: unknown | 'missing'): Promise<void> {
  await page.route(CHARACTER_URL, (route) => route.fulfill(data === 'missing' ? NOT_FOUND : ok(data)));
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(NOT_FOUND));
}

test('the header says who the character is, the guild links to its page, and the chip ends the level line', async ({
  page,
}) => {
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-descriptor')).toHaveText(
    'Discipline Priest · 31/20/0 · <Cx66 Olympus> · Classic Beta PvP',
  );
  // "Discipline Priest" in --text, the rest muted.
  await expect(page.getByTestId('character-spec')).toHaveCSS('color', 'rgb(233, 228, 216)');
  await expect(page.getByTestId('character-descriptor')).toHaveCSS('color', 'rgb(154, 148, 132)');
  const guild = page.getByTestId('character-guild-link');
  await expect(guild).toHaveText('<Cx66 Olympus>');
  await expect(guild).toHaveAttribute('href', '/guild/us/hardcore/cx66-olympus');
  await expect(page.getByTestId('character-realm')).toHaveText('Classic Beta PvP');
  await expect(page.getByTestId('character-level-race')).toHaveText('Level 60 Night Elf');
  const chip = page.getByTestId('character-ruleset-chip');
  await expect(chip).toHaveText('Hardcore US');
  await expect(chip).toHaveCSS('color', 'rgb(154, 148, 132)');
  await expect(chip).toHaveCSS('border-top-color', 'rgb(38, 46, 64)');
  // The chip is the last item on the level line, 8px after its text, on that line.
  const level = await page.getByTestId('character-level-race').boundingBox();
  const chipBox = await chip.boundingBox();
  expect(level).not.toBeNull();
  expect(chipBox).not.toBeNull();
  expect(Math.round((chipBox?.x ?? 0) - ((level?.x ?? 0) + (level?.width ?? 0)))).toBe(8);
  expect(
    Math.abs((chipBox?.y ?? 0) + (chipBox?.height ?? 0) / 2 - ((level?.y ?? 0) + (level?.height ?? 0) / 2)),
  ).toBeLessThan(3);
  // A visitor gets no sync line.
  await expect(page.getByTestId('character-sync')).toHaveCount(0);
  // Fix round 1 item 3: the fixture's rows carry execution_score, so the column never reads NaN.
  await expect(page.getByTestId('character-execution').first()).toHaveText('94%');
  await expect(page.getByTestId('character-history')).not.toContainText('NaN');
});

test('the tooltip mock reads as the client prints it: name, <guild>, "Level 60 Night Elf (Player)", class, a blank line, the score', async ({
  page,
}) => {
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);

  const lines = page.getByTestId('character-in-game-tooltip').locator(':scope > span');
  await expect(lines).toHaveText([
    'Elyra Duskvale',
    '<Cx66 Olympus>',
    'Level 60 Night Elf (Player)',
    'Priest',
    '',
    'Forever Score 1830 · Epic · as of 4 Nov',
  ]);
  // The visitor's page knows no faction, so there is no faction line.
  await expect(page.getByTestId('character-in-game-faction')).toHaveCount(0);
});

test('a long guild name wraps inside the mock and never pushes the score line out', async ({ page }) => {
  const long = 'The Exceedingly Long Guild Name Of Olympus';
  await stub(page, {
    ...CHARACTER.data,
    ...SCORE,
    character: { ...CHARACTER.data.character, guild: { name: long, rank: 'member' } },
  });
  await page.goto(CHARACTER_PATH);

  const tooltip = page.getByTestId('character-in-game-tooltip');
  const box = await tooltip.boundingBox();
  expect(box?.width ?? 0).toBeLessThanOrEqual(320);
  const score = await page.getByTestId('character-in-game-score').boundingBox();
  expect((score?.x ?? 0) + (score?.width ?? 0)).toBeLessThanOrEqual((box?.x ?? 0) + (box?.width ?? 0));
  await expect(page.getByTestId('character-in-game-guild')).toHaveText(`<${long}>`);
});

test('the score card draws the history across its content column, with a caption and the newest date', async ({
  page,
}) => {
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);

  const history = page.getByTestId('forever-score-history');
  await expect(history).toBeVisible();
  await expect(page.getByTestId('forever-score-sparkline').locator('path')).toHaveAttribute(
    'stroke',
    'var(--color-gold)',
  );
  const dot = page.getByTestId('forever-score-sparkline-last');
  await expect(dot).toHaveCSS('background-color', 'rgb(163, 53, 238)');
  expect(labels.tiers.epic.color.toLowerCase()).toBe('#a335ee');
  await expect(page.getByTestId('forever-score-history-caption')).toHaveText('Up 590 since 12 Oct');
  await expect(page.getByTestId('forever-score-history-last')).toHaveText('4 Nov');
  await expect(
    page.getByRole('img', { name: 'Forever Score history: 1240 on 12 Oct, 1830 on 4 Nov' }),
  ).toBeVisible();

  // Fix round 2 item 1: the plot and the date row start on the Gear label's left edge; the plot
  // is the tier strip's width from md (460px) and the content width on a phone; 40px high,
  // 48px from 1024px; the 8px end dot is centred on the right edge.
  const plot = await page.getByTestId('forever-score-plot').boundingBox();
  const gear = await page.getByTestId('forever-score-gear').boundingBox();
  const caption = await page.getByTestId('forever-score-history-caption').boundingBox();
  const last = await page.getByTestId('forever-score-history-last').boundingBox();
  const dotBox = await dot.boundingBox();
  if (plot === null || gear === null || caption === null || last === null || dotBox === null)
    throw new Error('a sparkline part is missing');
  expect(Math.abs(plot.x - gear.x)).toBeLessThan(1);
  expect(Math.abs(caption.x - gear.x)).toBeLessThan(1);
  expect(Math.abs(last.x + last.width - (plot.x + plot.width))).toBeLessThan(1);
  const viewport = page.viewportSize()?.width ?? 0;
  expect(plot.height).toBe(viewport >= 1024 ? 48 : 40);
  if (viewport >= 768) expect(Math.round(plot.width)).toBe(460);
  else expect(Math.abs(plot.width - gear.width)).toBeLessThan(1);
  expect([dotBox.width, dotBox.height]).toEqual([8, 8]);
  expect(Math.abs(dotBox.x + dotBox.width / 2 - (plot.x + plot.width))).toBeLessThan(1);

  // The table a screen reader gets as well: out of sight, one row per date, oldest first.
  const table = page.getByRole('table', { name: 'Forever Score by date, oldest first' });
  await expect(table).toBeAttached();
  await expect(table).not.toBeInViewport();
  const rows = table.locator('tbody tr');
  await expect(rows).toHaveCount(SCORE_HISTORY.length);
  await expect(rows.first()).toHaveText('12 Oct 20261240');
  await expect(rows.last()).toHaveText('4 Nov 20261830');
});

test('a character nobody knows the identity of: no guild, realm or level line, never a placeholder', async ({
  page,
}) => {
  await stub(page, {
    ...CHARACTER.data,
    ...SCORE,
    character: { ...CHARACTER.data.character, guild: null, level: null, race: null, realm: null },
    builds_seen: [],
    score_history: [SCORE_HISTORY[4]],
  });
  await page.goto(CHARACTER_PATH);

  // "Discipline Priest" then the chip, on the one line.
  await expect(page.getByTestId('character-descriptor')).toHaveText('Discipline PriestHardcore US');
  const spec = await page.getByTestId('character-spec').boundingBox();
  const chip = await page.getByTestId('character-ruleset-chip').boundingBox();
  expect(Math.round((chip?.x ?? 0) - ((spec?.x ?? 0) + (spec?.width ?? 0)))).toBe(8);
  await expect(page.getByTestId('character-split')).toHaveCount(0);
  await expect(page.getByTestId('character-guild-link')).toHaveCount(0);
  await expect(page.getByTestId('character-realm')).toHaveCount(0);
  await expect(page.getByTestId('character-level-race')).toHaveCount(0);
  await expect(page.getByTestId('character-in-game-guild')).toHaveCount(0);
  await expect(page.getByTestId('character-in-game-descriptor')).toHaveCount(0);
  await expect(page.getByTestId('character-in-game-class')).toHaveCount(0);
  // No guild and no level line: no blank line either, the score follows the name.
  await expect(page.getByTestId('character-in-game-tooltip').locator(':scope > span')).toHaveText([
    'Elyra Duskvale',
    'Forever Score 1830 · Epic · as of 4 Nov',
  ]);
  // One point is not a line: no sparkline and no empty box where it would be.
  await expect(page.getByTestId('forever-score-card')).toBeVisible();
  await expect(page.getByTestId('forever-score-history')).toHaveCount(0);
  await expect(page.getByTestId('character')).not.toContainText(/null|undefined|Unknown/);
});

/** The me fixture's addon-synced character, made the owner of this page's character. */
function ownerMe(): typeof meAddonFixture {
  const own = meAddonFixture.characters[0];
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
      },
    ],
  };
}

test('the owner gets the Character header’s sync line third, on the name’s left edge, and the faction in the mock', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) => route.fulfill(ok(ownerMe())));
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);

  const sync = page.getByTestId('character-sync');
  // The fixture's build was captured 4 minutes before the worker loaded it.
  await expect(sync).toHaveText(/^Addon · \d+ minutes ago$/);
  await expect(sync).toHaveCSS('font-size', '12px');
  await expect(page.getByTestId('character-sync-stale')).toHaveCount(0);
  await expect(page.getByTestId('character')).not.toContainText('Last seen in game');
  const descriptor = await page.getByTestId('character-descriptor').boundingBox();
  const level = await page.getByTestId('character-level-line').boundingBox();
  const syncBox = await sync.boundingBox();
  expect(syncBox?.x).toBe(descriptor?.x);
  expect(syncBox?.x).toBe(level?.x);
  expect((syncBox?.y ?? 0) > (level?.y ?? 0)).toBe(true);
  await expect(page.getByTestId('character-in-game-faction')).toHaveText('Horde');
});

test('a stale addon sync puts only "Log out in the game to sync" in ember', async ({ page }) => {
  const me = ownerMe();
  const own = me.characters[0];
  if (own.build === undefined) throw new Error('the me fixture character has no build');
  me.characters[0] = {
    ...own,
    build: { ...own.build, captured_at: new Date(Date.now() - 19 * 86_400_000).toISOString() },
  };
  await page.route('**/v1/me', (route) => route.fulfill(ok(me)));
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-sync')).toHaveText(
    'Addon · 19 days ago · Log out in the game to sync',
  );
  const stale = page.getByTestId('character-sync-stale');
  await expect(stale).toHaveText('Log out in the game to sync');
  await expect(stale).toHaveCSS('color', 'rgb(214, 110, 40)');
  await expect(page.getByTestId('character-sync')).toHaveCSS('color', 'rgb(154, 148, 132)');
});

test('a visitor’s sign-in band sits below the Forever Score card at every width', async ({ page }) => {
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);

  const header = await page.getByTestId('character-name').boundingBox();
  const card = await page.getByTestId('forever-score-card').boundingBox();
  const band = await page.getByTestId('current-character-bar').boundingBox();
  if (header === null || card === null || band === null) throw new Error('a profile part is missing');
  expect(band.y).toBeGreaterThan(card.y + card.height);
  expect(header.y).toBeLessThan(card.y);
});

test('a 404 is still the no-logs page, and its owner still gets "Upload logs" (ruling 4)', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) => route.fulfill(ok(ownerMe())));
  await stub(page, 'missing');
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-unlogged')).toBeVisible();
  await expect(page.getByTestId('character-ruleset-chip')).toHaveText('Hardcore US');
  await expect(page.getByTestId('forever-score-empty')).toContainText('No logs yet');
  const action = page.getByTestId('forever-score-empty-action');
  await expect(action).toHaveText('Upload logs');
  await expect(action).toHaveAttribute('href', '/setup');
  await expect(page.getByTestId('forever-score-history')).toHaveCount(0);
  await expect(page.getByTestId('character-error')).toHaveCount(0);
});
