// web/tests/e2e/character-identity.spec.ts
// Task L8.4 and its fix round 2: who the character is and its score over time, on the profile.
// The header reads "Discipline Priest · 31/20/0 · <Cx66 Olympus> · Doomhowl" with
// "Level 60 Night Elf" under it, the "Hardcore US" chip ending the last line, and for the owner
// the Character header's sync line third; the tooltip mock reads as the 1.60 client prints it;
// the Forever Score card draws score_history as a sparkline across its content column with a
// caption. A null field is left out, never a placeholder, and a 404 is still the no-logs page.
// The API is stubbed with page.route from the shared fixture (support/character-fixture.ts),
// which carries the identity and the history; the panel's execution-rating read is answered 404
// so only these parts are in play.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import labels from '../../../docs/contracts/fixtures/score-labels.json' with { type: 'json' };
import { meAddonFixture } from '../../src/fixtures/me-addon';
import {
  CHARACTER,
  CHARACTER_PATH,
  CHARACTER_URL,
  SCORE_HISTORY,
  routeOwner,
} from './support/character-fixture';

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
    'Discipline Priest · 31/20/0 · <Cx66 Olympus> · Doomhowl',
  );
  // "Discipline Priest" in --text, the rest muted.
  await expect(page.getByTestId('character-spec')).toHaveCSS('color', 'rgb(255, 255, 255)');
  await expect(page.getByTestId('character-descriptor')).toHaveCSS('color', 'rgb(154, 148, 132)');
  const guild = page.getByTestId('character-guild-link');
  await expect(guild).toHaveText('<Cx66 Olympus>');
  await expect(guild).toHaveAttribute('href', '/guild/us/hardcore/cx66-olympus');
  await expect(page.getByTestId('character-realm')).toHaveText('Doomhowl');
  await expect(page.getByTestId('character-level-race')).toHaveText('Level 60 Night Elf');
  const chip = page.getByTestId('character-ruleset-chip');
  await expect(chip).toHaveText('Hardcore US');
  await expect(chip).toHaveCSS('color', 'rgb(154, 148, 132)');
  await expect(chip).toHaveCSS('border-top-color', 'rgb(38, 46, 64)');
  // The chip ends the level line, 12px after its text, on that line; on a phone (under 640px) it
  // drops to a line of its own under the race, so the two never run together (W.9).
  const level = await page.getByTestId('character-level-race').boundingBox();
  const chipBox = await chip.boundingBox();
  expect(level).not.toBeNull();
  expect(chipBox).not.toBeNull();
  if (test.info().project.name === 'mobile') {
    expect(chipBox!.y).toBeGreaterThanOrEqual(level!.y + level!.height - 1);
    expect(Math.abs(chipBox!.x - level!.x)).toBeLessThan(1);
  } else {
    expect(Math.round(chipBox!.x - (level!.x + level!.width))).toBe(12);
    expect(Math.abs(chipBox!.y + chipBox!.height / 2 - (level!.y + level!.height / 2))).toBeLessThan(3);
  }
  // A visitor gets no sync line.
  await expect(page.getByTestId('character-sync')).toHaveCount(0);
  // Fix round 1 item 3: the fixture's rows carry execution_score, so the column never reads NaN.
  await expect(page.getByTestId('character-execution').first()).toHaveText('94%');
  await expect(page.getByTestId('character-history')).not.toContainText('NaN');
});

test('a supporter character carries the chip right after the name; any other carries none', async ({
  page,
}) => {
  await stub(page, { ...CHARACTER.data, ...SCORE, supporter: true });
  await page.goto(CHARACTER_PATH);
  const chip = page.getByTestId('supporter-chip');
  await expect(chip).toHaveText('Supporter');
  const name = await page.getByTestId('character-name').boundingBox();
  const chipBox = await chip.boundingBox();
  expect(name).not.toBeNull();
  expect(chipBox).not.toBeNull();
  expect(chipBox!.x).toBeGreaterThanOrEqual(name!.x + name!.width);
  expect(Math.abs(chipBox!.y + chipBox!.height / 2 - (name!.y + name!.height / 2))).toBeLessThan(8);
});

test('a character that is not a supporter has no chip', async ({ page }) => {
  await stub(page, { ...CHARACTER.data, ...SCORE, supporter: false });
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('character-name')).toBeVisible();
  await expect(page.getByTestId('supporter-chip')).toHaveCount(0);
});

test('the tooltip mock reads as the client prints it: name, <guild>, "Level 60 Night Elf", class, a blank line, the score', async ({
  page,
}) => {
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);

  const lines = page.getByTestId('character-in-game-tooltip').locator(':scope > span');
  await expect(lines).toHaveText([
    'Elyra Duskvale',
    '<Cx66 Olympus>',
    'Level 60 Night Elf',
    'Priest',
    'Alliance',
    '',
    'Forever Score 1830 · Epic · as of 4 Nov',
  ]);
  // Fix round 3 item 15: the faction comes from the race (a Night Elf is Alliance), so a
  // visitor gets the same line the owner does.
  await expect(page.getByTestId('character-in-game-faction')).toHaveText('Alliance');
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
  const viewport = page.viewportSize()?.width ?? 0;
  if (viewport >= 1280) {
    // Fix round 3 item 13: from 1280 the sparkline is the right column, right of the parts.
    expect(plot.x).toBeGreaterThan(gear.x + gear.width);
    expect(Math.abs(caption.x - plot.x)).toBeLessThan(1);
  } else {
    expect(Math.abs(plot.x - gear.x)).toBeLessThan(1);
    expect(Math.abs(caption.x - gear.x)).toBeLessThan(1);
  }
  expect(Math.abs(last.x + last.width - (plot.x + plot.width))).toBeLessThan(1);
  expect(plot.height).toBe(viewport >= 1024 ? 48 : 40);
  if (viewport >= 768 && viewport < 1280) expect(Math.round(plot.width)).toBe(460);
  else if (viewport < 768) expect(Math.abs(plot.width - gear.width)).toBeLessThan(1);
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

  // "Discipline Priest" then the chip, on the one line (a zero-width word joiner between them,
  // fix round 3 item 2).
  await expect(page.getByTestId('character-descriptor')).toHaveText('Discipline Priest⁠Hardcore US');
  const spec = await page.getByTestId('character-spec').boundingBox();
  const chip = await page.getByTestId('character-ruleset-chip').boundingBox();
  expect(Math.round((chip?.x ?? 0) - ((spec?.x ?? 0) + (spec?.width ?? 0)))).toBe(12);
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
  // One point is not a line: no sparkline and no empty box where it would be, but the date the
  // record starts, in the text colour (fix round 3 item 12).
  await expect(page.getByTestId('forever-score-card')).toBeVisible();
  await expect(page.getByTestId('forever-score-history')).toHaveCount(0);
  const first = page.getByTestId('forever-score-first');
  await expect(first).toHaveText('First score recorded 4 Nov.');
  await expect(first).toHaveCSS('color', 'rgb(233, 228, 216)');
  await expect(first).toHaveCSS('font-size', '12px');
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
  // Fix round 3 item 15: the faction comes from the race, the visitor's mock exactly, not from
  // the /v1/me row (this fixture row says Horde for a Night Elf).
  await expect(page.getByTestId('character-in-game-faction')).toHaveText('Alliance');
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

// ---- L8.4 fix round 3 ----

test('the 404 no-logs page draws the character header above the Forever Score card (item 1, D1)', async ({
  page,
}) => {
  await routeOwner(page);
  await stub(page, 'missing');
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-unlogged')).toBeVisible();
  const name = await page.getByTestId('character-name').boundingBox();
  const card = await page.getByTestId('forever-score-card').boundingBox();
  if (name === null || card === null) throw new Error('a 404 page part is missing');
  expect(name.y + name.height).toBeLessThan(card.y);
});

/** The line box a range of text sits on: its rect's vertical centre. */
async function centreOf(page: Page, testid: string): Promise<number> {
  const box = await page.getByTestId(testid).boundingBox();
  if (box === null) throw new Error(`${testid} is missing`);
  return box.y + box.height / 2;
}

test.describe('a 40-character guild with no realm at 320px (item 2, D2)', () => {
  test.use({ viewport: { width: 320, height: 720 } });

  test('the guild wraps and only the chip is glued to its last word', async ({ page }) => {
    const guild = 'Olympus XXVII Veterans of the Long March'; // 40 characters
    expect(guild).toHaveLength(40);
    await stub(page, {
      ...CHARACTER.data,
      ...SCORE,
      character: {
        ...CHARACTER.data.character,
        guild: { name: guild, rank: 'member' },
        realm: null,
        level: null,
        race: null,
      },
    });
    await page.goto(CHARACTER_PATH);

    const link = page.getByTestId('character-guild-link');
    await expect(link).toHaveText(`<${guild}>`);
    expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(320);
    // The guild's own text wraps: more than one line box.
    const lines = await link.evaluate((element) =>
      [...element.getClientRects()].map((rect) => ({
        right: rect.right,
        top: rect.top,
        bottom: rect.bottom,
      })),
    );
    expect(lines.length).toBeGreaterThan(1);
    // The chip sits on the guild last line, 12px after its last word.
    const lastLine = lines[lines.length - 1];
    const chip = await page.getByTestId('character-ruleset-chip').boundingBox();
    if (chip === null) throw new Error('no chip');
    expect(Math.round(chip.x - lastLine.right)).toBe(12);
    expect(Math.abs(chip.y + chip.height / 2 - (lastLine.top + lastLine.bottom) / 2)).toBeLessThan(1);
  });
});

test('an owner whose build has no gear still gets the sync line, and no "No export yet" (item 3, D3)', async ({
  page,
}) => {
  const me = ownerMe();
  const own = me.characters[0];
  if (own.build === undefined) throw new Error('the me fixture character has no build');
  const { gear: _gear, ...gearless } = own.build;
  me.characters[0] = { ...own, build: gearless };
  await page.route('**/v1/me', (route) => route.fulfill(ok(me)));
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-sync')).toHaveText(/^Addon · \d+ minutes ago$/);
  await expect(page.getByTestId('character')).not.toContainText('No export yet');
});

/** The fixture with a third, older history row so "Best per encounter" renders as well, and one
 *  unscored row in each table. */
function mixedExecution(): Record<string, unknown> {
  const [newest, second] = CHARACTER.data.history;
  const older = {
    ...newest,
    value: 1656,
    percentile: 82.4,
    fought_at: '2026-12-07T20:00:00Z',
    report_id: 'olderreport1',
    fight_index: 2,
    execution_score: 0.81,
  };
  return {
    ...CHARACTER.data,
    best: [CHARACTER.data.best[0], { ...CHARACTER.data.best[1], execution_score: null }],
    history: [newest, { ...second, execution_score: null }, older],
  };
}

test('a scored and an unscored Execution cell sit inside their rows in both tables (items 4 and 6)', async ({
  page,
}) => {
  await stub(page, { ...mixedExecution(), ...SCORE });
  for (const width of [1440, 390]) {
    await page.setViewportSize({ width, height: 900 });
    await page.goto(CHARACTER_PATH);
    await expect(page.getByTestId('character-best')).toBeVisible();
    for (const [table, testid] of [
      ['character-best', 'character-best-execution'],
      ['character-history', 'character-execution'],
    ] as const) {
      // Each row and its cell are read in one pass, so nothing above (a panel settling) can move
      // the page between the two reads.
      const measured = await page.getByTestId(table).evaluate((list, cellTestid) => {
        return [...list.querySelectorAll(':scope > li:not([data-testid="character-table-header"])')].map(
          (li) => {
            const cell = li.querySelector(`[data-testid="${cellTestid}"]`);
            const shown = cell !== null && getComputedStyle(cell).display !== 'none';
            const rect = (element: Element) => {
              const { y, height } = element.getBoundingClientRect();
              return { y, height };
            };
            return { row: rect(li), cell: shown && cell !== null ? rect(cell) : null };
          },
        );
      }, testid);
      expect(measured.length, table).toBeGreaterThan(1);
      const heights: number[] = [];
      for (const [index, { row, cell: box }] of measured.entries()) {
        heights.push(Math.round(row.height));
        // Fix round 4 item 3: both tables keep the column on a phone, so every cell shows.
        expect(box, `${table} row ${index} at ${width}`).not.toBeNull();
        if (box === null) continue;
        const where = `${table} row ${index} at ${width}`;
        expect(box.y, where).toBeGreaterThanOrEqual(row.y - 0.5);
        expect(box.y + box.height, where).toBeLessThanOrEqual(row.y + row.height + 0.5);
        // Centred on the row, not riding at its top.
        expect(Math.abs(box.y + box.height / 2 - (row.y + row.height / 2)), where).toBeLessThan(2);
      }
      // A scored row and an unscored one are the same height (the phone leak made it 117 vs 61).
      expect(new Set(heights).size, `${table} at ${width}: ${heights.join(', ')}`).toBe(1);
    }
  }
});

test('the Execution column is explained under each table title, visibly at every width (item 7, fix round 4 items 6 and 7, fix round 5 item 1)', async ({
  page,
}) => {
  await stub(page, { ...mixedExecution(), ...SCORE });
  await page.goto(CHARACTER_PATH);
  const definitions = page.getByTestId('character-execution-definition');
  await expect(definitions).toHaveCount(2);
  // Damage only, as api/openapi.yaml's ExecutionScore says (fix round 5 item 1).
  const words =
    'Execution: the fight’s damage as a share of what the same gear and talents do in the simulator; healers and tanks have none.';
  for (const definition of await definitions.all()) {
    await expect(definition).toHaveText(words);
    await expect(definition).toBeVisible();
    await expect(definition).toHaveCSS('position', 'static');
    await expect(definition).toHaveCSS('color', 'rgb(154, 148, 132)');
    await expect(definition).toHaveCSS('font-size', '12px');
  }
  // Said once: the header no longer repeats it in a title.
  await expect(page.locator(`[title="${words}"]`)).toHaveCount(0);
});

test.describe('with JavaScript off (item 9)', () => {
  test.use({ javaScriptEnabled: false });

  test('the skeleton goes and one line says why, with the Planner link', async ({ page }) => {
    await page.goto(CHARACTER_PATH);
    await expect(page.getByTestId('character-skeleton')).toBeHidden();
    const line = page.getByTestId('character-noscript');
    await expect(line).toBeVisible();
    await expect(line).toHaveText(
      'Forever Sixty needs JavaScript to show this character. Turn it on and reload, or open the Planner.',
    );
    await expect(line.getByRole('link', { name: 'Planner' })).toHaveAttribute('href', '/planner');
  });
});

/**
 * Line 1's units and dots (fix round 5 item 6): for each "·", the top of its box and of the
 * visible character before it, which must share a line (the dot ends its segment, so no line
 * starts with one); and how many line boxes each unit takes.
 */
async function headerLines(
  page: Page,
): Promise<{ dots: { dot: number; before: number }[]; units: number[] }> {
  return page.getByTestId('character-descriptor').evaluate((element) => {
    const chars: { node: Text; at: number }[] = [];
    const walker = document.createTreeWalker(element, NodeFilter.SHOW_TEXT);
    for (let node = walker.nextNode(); node !== null; node = walker.nextNode()) {
      const text = node as Text;
      for (let at = 0; at < text.data.length; at += 1)
        if (!/[\s\u00a0\u2060]/.test(text.data[at])) chars.push({ node: text, at });
    }
    const top = ({ node, at }: { node: Text; at: number }): number => {
      const range = document.createRange();
      range.setStart(node, at);
      range.setEnd(node, at + 1);
      return range.getBoundingClientRect().top;
    };
    const dots = chars.flatMap((char, index) =>
      char.node.data[char.at] === '·' && index > 0 ? [{ dot: top(char), before: top(chars[index - 1]) }] : [],
    );
    // A unit's line count: the distinct tops of its boxes (Chrome gives an inline unit with a
    // child element and a text node more than one rect on a single line).
    const units = [...element.querySelectorAll('[data-testid="character-segment"]')].map(
      (unit) => new Set([...unit.getClientRects()].map((rect) => Math.round(rect.top))).size,
    );
    return { dots, units };
  });
}

for (const width of [390, 360, 320]) {
  test.describe(`the header at ${width}px (fix round 5 item 6)`, () => {
    test.use({ viewport: { width, height: 740 } });

    test('every dot ends its segment, each segment moves whole, and no line starts with a dot', async ({
      page,
    }) => {
      await stub(page, { ...CHARACTER.data, ...SCORE });
      await page.goto(CHARACTER_PATH);
      await expect(page.getByTestId('character-descriptor')).toHaveText(
        'Discipline Priest · 31/20/0 · <Cx66 Olympus> · Doomhowl',
      );
      const { dots, units } = await headerLines(page);
      expect(dots).toHaveLength(3);
      for (const { dot, before } of dots) expect(Math.abs(dot - before)).toBeLessThan(2);
      // "Discipline Priest", "31/20/0", "<Cx66 Olympus>" and the realm: one line box each.
      expect(units).toEqual([1, 1, 1, 1]);
      // The line does wrap at these widths, after a dot.
      const tops = await page
        .getByTestId('character-segment')
        .evaluateAll((all) => [...new Set(all.map((unit) => Math.round(unit.getBoundingClientRect().top)))]);
      expect(tops.length).toBeGreaterThan(1);
      expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(width);
    });
  });
}

test.describe('a long realm at 360px (item 10, fix round 5 item 6)', () => {
  test.use({ viewport: { width: 360, height: 740 } });

  test('the realm moves whole and no line starts with a dot', async ({ page }) => {
    await stub(page, {
      ...CHARACTER.data,
      ...SCORE,
      character: {
        ...CHARACTER.data.character,
        guild: { name: 'Olympus XXVII Veterans of the Long March', rank: 'member' },
        realm: 'Classic Beta PvE Two Extended',
      },
    });
    await page.goto(CHARACTER_PATH);

    const realm = page.getByTestId('character-realm');
    await expect(realm).toHaveText('Classic Beta PvE Two Extended');
    expect(await realm.evaluate((element) => element.getClientRects().length)).toBe(1);
    // Every "·" shares its line with the character before it: a dot ends its segment's line.
    const { dots } = await headerLines(page);
    expect(dots).toHaveLength(3);
    for (const { dot, before } of dots) expect(Math.abs(dot - before)).toBeLessThan(2);
    expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(360);
  });
});

// L8.5 ruling 3 replaces item 11's offset 0: the ring stands 2px off like every other ring, and
// the guild's 4px margins keep it off the dots (character-performance.spec.ts measures them).
test('the guild link’s focus ring is 2px at a 2px offset (item 11, L8.5 ruling 3)', async ({ page }) => {
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);
  const link = page.getByTestId('character-guild-link');
  await link.focus();
  await expect(link).toBeFocused();
  await expect(link).toHaveCSS('outline-offset', '2px');
  await expect(link).toHaveCSS('outline-width', '2px');
});

test('the sparkline caption is 12px in the text colour; a sparse history counts its readings (item 12)', async ({
  page,
}) => {
  await stub(page, { ...CHARACTER.data, ...SCORE, score_history: [SCORE_HISTORY[0], SCORE_HISTORY[4]] });
  await page.goto(CHARACTER_PATH);
  const caption = page.getByTestId('forever-score-history-caption');
  await expect(caption).toHaveText('Up 590 since 12 Oct, 2 readings');
  await expect(caption).toHaveCSS('color', 'rgb(233, 228, 216)');
  await expect(caption).toHaveCSS('font-size', '12px');
  const last = page.getByTestId('forever-score-history-last');
  await expect(last).toHaveCSS('color', 'rgb(154, 148, 132)');
});

test.describe('the score card at 1440 (item 13)', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'a desktop width');

  test('is two columns: the sparkline meets the card’s content edge, over the tier strip', async ({
    page,
  }) => {
    await stub(page, { ...CHARACTER.data, ...SCORE });
    await page.goto(CHARACTER_PATH);
    const card = page.getByTestId('forever-score-card');
    const edge = await card.evaluate((element) => {
      const style = getComputedStyle(element);
      const rect = element.getBoundingClientRect();
      return rect.right - parseFloat(style.borderRightWidth) - parseFloat(style.paddingRight);
    });
    const plot = await page.getByTestId('forever-score-plot').boundingBox();
    const tiers = await page.getByTestId('forever-score-tiers').boundingBox();
    const parts = await page.getByTestId('forever-score-parts').boundingBox();
    if (plot === null || tiers === null || parts === null) throw new Error('a card part is missing');
    expect(Math.abs(plot.x + plot.width - edge)).toBeLessThan(1);
    expect(plot.x).toBeGreaterThan(parts.x + parts.width);
    expect(tiers.y).toBeGreaterThan(plot.y + plot.height);
    expect(Math.abs(tiers.x - plot.x)).toBeLessThan(1);
  });
});

test('the chip is centred on its text line (item 14)', async ({ page }) => {
  test.skip(test.info().project.name === 'mobile', 'on a phone the chip has a line of its own');
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);
  const level = await centreOf(page, 'character-level-race');
  const chip = await centreOf(page, 'character-ruleset-chip');
  expect(Math.abs(chip - level)).toBeLessThan(0.5);
});

test('the owner’s sync age ticks each minute while the tab lives (item 14)', async ({ page }) => {
  const now = new Date('2026-10-12T01:34:53Z');
  await page.clock.install({ time: now });
  const me = ownerMe();
  const own = me.characters[0];
  if (own.build === undefined) throw new Error('the me fixture character has no build');
  me.characters[0] = {
    ...own,
    build: { ...own.build, captured_at: new Date(now.getTime() - 4 * 60_000).toISOString() },
  };
  await page.route('**/v1/me', (route) => route.fulfill(ok(me)));
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);

  const sync = page.getByTestId('character-sync');
  await expect(sync).toHaveText('Addon · 4 minutes ago');
  await page.clock.runFor(60_000);
  await expect(sync).toHaveText('Addon · 5 minutes ago');
});

// ---- L8.4 fix round 4 ----

test('under the score, the points to the next tier: "170 to Legendary (2000)" (fix round 4 item 1)', async ({
  page,
}) => {
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);
  const line = page.getByTestId('forever-score-next-tier');
  await expect(line).toHaveText('170 to Legendary (2000)');
  // Both read in one pass, so a panel settling above cannot move the page between the reads.
  const below = await page.getByTestId('forever-score-headline').evaluate((headline) => {
    const badge = headline.querySelector('[data-testid="score-badge"]')?.getBoundingClientRect();
    const next = headline.querySelector('[data-testid="forever-score-next-tier"]')?.getBoundingClientRect();
    return badge === undefined || next === undefined ? null : next.top - badge.bottom;
  });
  expect(below).not.toBeNull();
  expect(below ?? -1).toBeGreaterThanOrEqual(-0.5);
});

test('"+88 pts" on one line, and Gear opens the spec’s best-in-slot page at this band (fix round 4 item 2)', async ({
  page,
}) => {
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);
  const cue = page.getByTestId('forever-score-gain-cue');
  await expect(cue).toHaveText('+88 pts');
  const cueBox = await cue.boundingBox();
  expect(cueBox?.height ?? 99).toBeLessThan(20);
  const sentence = page.getByTestId('forever-score-most-to-gain');
  await expect(sentence).toHaveText('Most to gain: Gear, 88 below its top.');
  // The fixture's newest build is Discipline, the character a level 60 Night Elf.
  await expect(sentence.getByRole('link', { name: 'Gear' })).toHaveAttribute(
    'href',
    '/bis/priest/discipline?faction=alliance#band-alliance-60',
  );
});

test('the Performance rating defines the part it names, visibly at every width (fix round 4 item 3)', async ({
  page,
}) => {
  await page.route(CHARACTER_URL, (route) => route.fulfill(ok({ ...CHARACTER.data, ...SCORE })));
  await page.route(`${CHARACTER_URL}/rating`, (route) =>
    route.fulfill(
      ok({
        player_key: 'us/hardcore/elyra-duskvale',
        sample_size: 6,
        trend: Array.from({ length: 6 }, (_, i) => ({
          fought_at: `2026-12-0${i + 1}T00:00:00Z`,
          overall: 40 + i * 5,
          report_id: 'fixture2abcd',
          fight_index: i + 1,
        })),
        best_component: 'preparation',
        worst_component: 'activity',
        latest: {
          player_key: 'us/hardcore/elyra-duskvale',
          player_name: 'Elyra Duskvale',
          class: 'Priest',
          spec: 'Holy',
          role: 'healer',
          overall: 65,
          overall_uncapped: 65,
          overall_capped: false,
          basis: 'percentile',
          components: [],
        },
      }),
    ),
  );
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('rating-card-weakest')).toHaveText('Weakest part of a fight: Activity.');
  const definition = page.getByTestId('rating-card-weakest-definition');
  await expect(definition).toHaveText(
    'Activity: the share of the fight you spent doing something rather than idle, leaving out the times the fight expects you to wait.',
  );
  await expect(definition).toBeVisible();
  await expect(definition).toHaveCSS('position', 'static');
  await expect(definition).toHaveCSS('font-size', '12px');
  await expect(definition).toHaveCSS('color', 'rgb(154, 148, 132)');
});

test.describe('the fight tables on a 390px phone (fix round 4 items 3 and 7)', () => {
  test.use({ viewport: { width: 390, height: 844 } });

  test('both keep Execution, 44px wide and 16px after the amount; Spec and Date stay off', async ({
    page,
  }) => {
    await stub(page, { ...mixedExecution(), ...SCORE });
    await page.goto(CHARACTER_PATH);
    for (const [table, testid] of [
      ['character-best', 'character-best-execution'],
      ['character-history', 'character-execution'],
    ] as const) {
      const list = page.getByTestId(table);
      await expect(list).toBeVisible();
      const header = list.getByTestId('character-table-header');
      await expect(header.getByText('Execution', { exact: true })).toBeVisible();
      await expect(header.getByText('Spec', { exact: true })).toBeHidden();
      await expect(header.getByText('Date', { exact: true })).toBeHidden();
      // The caption under the title is shown (fix round 4 item 7).
      const caption = list.locator(
        'xpath=preceding-sibling::p[@data-testid="character-execution-definition"]',
      );
      await expect(caption).toBeVisible();
      const measured = await list.evaluate((element, cellTestid) => {
        return [...element.querySelectorAll(':scope > li:not([data-testid="character-table-header"])')].map(
          (li) => {
            const cell = li.querySelector(`[data-testid="${cellTestid}"]`);
            const amount = li.querySelector('a[href^="/reports/"][aria-label]');
            if (cell === null || amount === null) return null;
            const c = cell.getBoundingClientRect();
            const a = amount.getBoundingClientRect();
            return {
              width: c.width,
              gap: c.left - a.right,
              shown: getComputedStyle(cell).display !== 'none',
            };
          },
        );
      }, testid);
      expect(measured.length, table).toBeGreaterThan(1);
      for (const [index, row] of measured.entries()) {
        const where = `${table} row ${index}`;
        expect(row, where).not.toBeNull();
        expect(row?.shown, where).toBe(true);
        expect(row?.width ?? 0, where).toBeGreaterThanOrEqual(44);
        expect(Math.abs((row?.gap ?? 0) - 16), where).toBeLessThan(0.5);
      }
    }
  });
});

test.describe('the fight tables at 1440 (fix round 4 item 8)', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'a desktop width');

  test('Date sits 20px after the amount, 8px more than the 12px column gap', async ({ page }) => {
    await stub(page, { ...mixedExecution(), ...SCORE });
    await page.goto(CHARACTER_PATH);
    const gaps = await page.getByTestId('character-history').evaluate((element) =>
      [...element.querySelectorAll(':scope > li:not([data-testid="character-table-header"])')].map((li) => {
        const amount = li.querySelector('a[href^="/reports/"][aria-label]')?.getBoundingClientRect();
        const date = li.querySelector('[data-testid="character-fight-date"]')?.getBoundingClientRect();
        return amount === undefined || date === undefined ? null : date.left - amount.right;
      }),
    );
    expect(gaps.length).toBeGreaterThan(1);
    for (const gap of gaps) expect(Math.abs((gap ?? 0) - 20)).toBeLessThan(0.5);
  });
});

test('the builds footnote says it in plain words (fix round 4 item 4)', async ({ page }) => {
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('character-builds-footnote')).toHaveText(
    'Seen in logs as a 31/20/0 split. Logs don’t record the order points were spent, so there is no saved build to open.',
  );
});

test('a visitor on the 404 page gets "Get the addon" and "Upload logs"; the owner keeps its button (fix round 4 item 5)', async ({
  page,
}) => {
  await stub(page, 'missing');
  await page.goto(CHARACTER_PATH);
  const next = page.getByTestId('character-not-found-next');
  await expect(next).toBeVisible();
  await expect(next.getByRole('link', { name: 'Get the addon' })).toHaveAttribute('href', '/setup');
  await expect(next.getByRole('link', { name: 'Upload logs' })).toHaveAttribute('href', '/logs#upload');

  await page.route('**/v1/me', (route) => route.fulfill(ok(ownerMe())));
  await page.reload();
  await expect(page.getByTestId('forever-score-empty-action')).toHaveText('Upload logs');
  await expect(page.getByTestId('forever-score-empty-next')).toHaveCount(0);
});

test('the guild link’s phone box is 44px, and a press turns it --gold-hover (fix round 4 items 6 and 8)', async ({
  page,
}) => {
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);
  const link = page.getByTestId('character-guild-link');
  const box = await link.boundingBox();
  if (box === null) throw new Error('the guild link is missing');
  if (test.info().project.name === 'mobile') expect(Math.abs(box.height - 44)).toBeLessThan(0.5);
  // Hold the press, read the colour, then release off the link so the page does not navigate.
  await page.mouse.move(box.x + box.width / 2, box.y + box.height / 2);
  await page.mouse.down();
  await expect(link).toHaveCSS('color', 'rgb(245, 210, 122)');
  await page.mouse.move(2, 2);
  await page.mouse.up();
});

test('the crest tops the name block, in the same place for a visitor and the owner (fix round 4 item 8)', async ({
  page,
}) => {
  await stub(page, { ...CHARACTER.data, ...SCORE });
  const measure = async (): Promise<{ crest: number; name: number }> => {
    await expect(page.getByTestId('character-name')).toBeVisible();
    // Read in one pass, relative to the page (scroll included), so nothing settling moves one read.
    const measured = await page.locator('header:has(h1)').evaluate((header) => {
      const crest = header.querySelector('[data-testid^="character-avatar"]')?.getBoundingClientRect();
      const name = header.querySelector('h1')?.getBoundingClientRect();
      return crest === undefined || name === undefined
        ? null
        : { crest: crest.top + window.scrollY, name: name.top + window.scrollY };
    });
    if (measured === null) throw new Error('the crest or the name is missing');
    return measured;
  };
  await page.goto(CHARACTER_PATH);
  const visitor = await measure();
  expect(Math.abs(visitor.crest - visitor.name)).toBeLessThan(0.5);

  await page.route('**/v1/me', (route) => route.fulfill(ok(ownerMe())));
  // An owner whose session already settled its pointer (the first answer of a session writes the
  // main): this measures the header alone.
  await page.addInitScript(() => window.sessionStorage.setItem('fs.pointerHealed', '1'));
  await page.reload();
  await expect(page.getByTestId('character-sync')).toBeVisible();
  const owner = await measure();
  expect(Math.abs(owner.crest - visitor.crest)).toBeLessThan(0.5);
});

// ---- L8.4 fix round 5 ----

test('the score card’s links rest gold and take --gold-hover on hover and press (fix round 5 item 2)', async ({
  page,
}) => {
  test.skip(test.info().project.name === 'mobile', 'hover is a pointer state');
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);
  const card = page.getByTestId('forever-score-card');
  // The page settles first (the rating panel's skeleton goes, the reveal finishes), so the
  // pointer stays over the link it hovers.
  await expect(page.getByTestId('character-rating-loading')).toHaveCount(0);
  await page.waitForFunction(() =>
    document.getAnimations().every((animation) => animation.playState !== 'running'),
  );
  for (const link of [
    page.getByTestId('forever-score-most-to-gain').getByRole('link', { name: 'Gear' }),
    card.getByRole('link', { name: 'What the score counts' }),
  ]) {
    await expect(link).toHaveCSS('color', 'rgb(229, 185, 85)');
    await expect(async () => {
      await link.hover();
      await expect(link).toHaveCSS('color', 'rgb(245, 210, 122)', { timeout: 500 });
    }).toPass();
    await page.mouse.down();
    await expect(link).toHaveCSS('color', 'rgb(245, 210, 122)');
    await page.mouse.move(0, 0);
    await page.mouse.up();
    await expect(link).toHaveCSS('color', 'rgb(229, 185, 85)');
  }
});

test('"+700 pts", the widest cue, fits its column on one line (fix round 5 item 3)', async ({ page }) => {
  await stub(page, {
    ...CHARACTER.data,
    ...SCORE,
    components: { ...SCORE.components, experience: 0 },
  });
  await page.goto(CHARACTER_PATH);
  const cue = page.getByTestId('forever-score-gain-cue');
  await expect(cue).toHaveText('+700 pts');
  const fit = await cue.evaluate((element) => {
    const range = document.createRange();
    range.selectNodeContents(element);
    const text = range.getBoundingClientRect();
    const box = element.getBoundingClientRect();
    return {
      text: text.width,
      box: box.width,
      height: box.height,
      scroll: element.scrollWidth,
      client: element.clientWidth,
    };
  });
  expect(fit.box).toBe(60);
  expect(fit.text).toBeLessThanOrEqual(fit.box);
  expect(fit.scroll).toBeLessThanOrEqual(fit.client);
  expect(fit.height).toBeLessThan(20);
});

test('the builds footnote names every split when there is more than one (fix round 5 item 4)', async ({
  page,
}) => {
  await stub(page, {
    ...CHARACTER.data,
    ...SCORE,
    builds_seen: [
      { talent_split: '31/20/0', spec: 'Discipline', first_seen: '2026-12-01T22:10:00Z' },
      { talent_split: '0/31/20', spec: 'Holy', first_seen: '2026-12-09T22:10:00Z' },
    ],
  });
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('character-builds-footnote')).toHaveText(
    'Seen in logs as 0/31/20 and 31/20/0 splits. Logs don’t record the order points were spent, so there is no saved build to open.',
  );
});

test('the builds footnote is absent when no build carries a split (fix round 5 item 4)', async ({ page }) => {
  await stub(page, {
    ...CHARACTER.data,
    ...SCORE,
    builds_seen: [{ talent_split: ' ', spec: 'Discipline', first_seen: '2026-12-09T22:10:00Z' }],
  });
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('character-builds')).toBeVisible();
  await expect(page.getByTestId('character-builds-footnote')).toHaveCount(0);
});

test('under "Most to gain", every part’s gap to its top, largest first (fix round 5 item 5)', async ({
  page,
}) => {
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);
  const gaps = page.getByTestId('forever-score-gaps');
  await expect(gaps).toHaveText('Gear +88 · Experience +80 · Parses +60 · Preparation +22 · Attendance +20');
  await expect(gaps).toBeVisible();
  await expect(gaps).toHaveCSS('color', 'rgb(154, 148, 132)');
  const under = await page.getByTestId('forever-score-card').evaluate((card) => {
    const sentence = card
      .querySelector('[data-testid="forever-score-most-to-gain"]')
      ?.getBoundingClientRect();
    const line = card.querySelector('[data-testid="forever-score-gaps"]')?.getBoundingClientRect();
    return sentence === undefined || line === undefined ? null : line.top - sentence.bottom;
  });
  expect(under).not.toBeNull();
  expect(under ?? -1).toBeGreaterThanOrEqual(0);
  expect(under ?? 99).toBeLessThan(8);
});

test('the "—" Execution cell is named "Not scored yet" in both tables (fix round 5 item 7)', async ({
  page,
}) => {
  await stub(page, { ...mixedExecution(), ...SCORE });
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('character-best')).toBeVisible();
  for (const testid of ['character-best-execution', 'character-execution']) {
    const dash = page.getByTestId(testid).filter({ hasText: '—' });
    await expect(dash).toHaveCount(1);
    await expect(dash).toHaveAttribute('title', 'Not scored yet');
    await expect(dash).toHaveAttribute('aria-label', 'Not scored yet');
  }
});

test.describe('the score card’s foot and the fight rows on a 390px phone (fix round 5 item 8)', () => {
  test.use({ viewport: { width: 390, height: 844 } });

  test('"What the score counts" sits 16px under the line above it; Gear keeps a 44px target (fix round 6)', async ({
    page,
  }) => {
    await stub(page, { ...CHARACTER.data, ...SCORE });
    await page.goto(CHARACTER_PATH);
    await expect(page.getByTestId('forever-score-gaps')).toBeVisible();
    await page.getByTestId('forever-score-gaps').scrollIntoViewIfNeeded();
    const foot = await page.getByTestId('forever-score-card').evaluate((card) => {
      const textOf = (element: Element | null | undefined): { top: number; bottom: number } | null => {
        if (element === null || element === undefined) return null;
        const range = document.createRange();
        range.selectNodeContents(element);
        const rects = [...range.getClientRects()];
        return { top: Math.min(...rects.map((r) => r.top)), bottom: Math.max(...rects.map((r) => r.bottom)) };
      };
      const gaps = textOf(card.querySelector('[data-testid="forever-score-gaps"]'));
      const explain = textOf(
        [...card.querySelectorAll('a')].find((link) => link.textContent?.trim() === 'What the score counts'),
      );
      const link = card.querySelector('[data-testid="forever-score-most-to-gain"] a');
      const gear = link?.getBoundingClientRect();
      if (gaps === null || explain === null || link === null || gear === undefined) return null;
      // What a tap at the top and bottom edges of the 44px target lands on (fix round 6): the
      // link itself, not the gaps line that follows it.
      const x = gear.left + gear.width / 2;
      const hit = (y: number): boolean => {
        const target = document.elementFromPoint(x, y);
        return target !== null && (target === link || link.contains(target));
      };
      return {
        space: explain.top - gaps.bottom,
        gear: gear.height,
        top: hit(gear.top + 0.5),
        bottom: hit(gear.bottom - 0.5),
      };
    });
    expect(foot).not.toBeNull();
    expect(Math.abs((foot?.space ?? 0) - 16)).toBeLessThanOrEqual(1);
    expect(foot?.gear ?? 0).toBeGreaterThanOrEqual(44);
    expect(foot?.top).toBe(true);
    expect(foot?.bottom).toBe(true);
  });

  test('a 1px --border-soft line parts the fight rows', async ({ page }) => {
    await stub(page, { ...mixedExecution(), ...SCORE });
    await page.goto(CHARACTER_PATH);
    await expect(page.getByTestId('character-best')).toBeVisible();
    for (const table of ['character-best', 'character-history']) {
      const rows = page.getByTestId(table).locator(':scope > li');
      expect(await rows.count()).toBeGreaterThan(2);
      for (const row of await rows.all()) {
        await expect(row).toHaveCSS('border-bottom-width', '1px');
        await expect(row).toHaveCSS('border-bottom-style', 'solid');
        await expect(row).toHaveCSS('border-bottom-color', 'rgb(28, 34, 48)');
      }
    }
  });
});
