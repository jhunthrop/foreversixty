// web/tests/e2e/character-identity.spec.ts
// Task L8.4: who the character is and its score over time, on the profile. The header reads
// "Discipline Priest · <Guild> · <Realm>" with "Level 60 Night Elf" under it and "Hardcore US"
// as a muted trailing chip; the tooltip mock gains the guild and level lines in the game's
// order; the Forever Score card draws score_history as a sparkline. A null field is left out,
// never a placeholder, and a 404 is still the no-logs page. The API is stubbed with page.route
// from the shared fixture (support/character-fixture.ts), which carries the identity and the
// history; the panel's execution-rating read is answered 404 so only these parts are in play.
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

test('the header says who the character is, the guild links to its page, and the ruleset is a chip', async ({
  page,
}) => {
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-descriptor')).toHaveText(
    'Discipline Priest · Cx66 Olympus · Classic Beta PvPHardcore US',
  );
  const guild = page.getByTestId('character-guild-link');
  await expect(guild).toHaveText('Cx66 Olympus');
  await expect(guild).toHaveAttribute('href', '/guild/us/hardcore/cx66-olympus');
  await expect(page.getByTestId('character-realm')).toHaveText('Classic Beta PvP');
  await expect(page.getByTestId('character-level-race')).toHaveText('Level 60 Night Elf');
  const chip = page.getByTestId('character-ruleset-chip');
  await expect(chip).toHaveText('Hardcore US');
  await expect(chip).toHaveCSS('color', 'rgb(154, 148, 132)');
  // Fix round 1 item 3: the fixture's rows carry execution_score, so the column never reads NaN.
  await expect(page.getByTestId('character-execution').first()).toHaveText('94%');
  await expect(page.getByTestId('character-history')).not.toContainText('NaN');
});

test('the tooltip mock reads name, guild, "Level 60 Night Elf Priest", a blank line, then the score', async ({
  page,
}) => {
  await stub(page, { ...CHARACTER.data, ...SCORE });
  await page.goto(CHARACTER_PATH);

  const lines = page.getByTestId('character-in-game-tooltip').locator(':scope > span');
  await expect(lines).toHaveText([
    'Elyra Duskvale',
    '<Cx66 Olympus>',
    'Level 60 Night Elf Priest',
    '',
    'Forever Score 1830 · Epic · as of 4 Nov',
  ]);
});

test('the score card draws the history as a gold sparkline, newest point in the tier colour', async ({
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
  await expect(page.getByTestId('forever-score-sparkline-last')).toHaveAttribute(
    'fill',
    labels.tiers.epic.color,
  );
  await expect(page.getByTestId('forever-score-history-first')).toHaveText('12 Oct');
  await expect(page.getByTestId('forever-score-history-last')).toHaveText('4 Nov');

  // The table a screen reader gets instead: out of sight, one row per date, oldest first.
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
    score_history: [SCORE_HISTORY[4]],
  });
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-descriptor')).toHaveText('Discipline PriestHardcore US');
  await expect(page.getByTestId('character-guild-link')).toHaveCount(0);
  await expect(page.getByTestId('character-realm')).toHaveCount(0);
  await expect(page.getByTestId('character-level-race')).toHaveCount(0);
  await expect(page.getByTestId('character-in-game-guild')).toHaveCount(0);
  await expect(page.getByTestId('character-in-game-descriptor')).toHaveCount(0);
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
