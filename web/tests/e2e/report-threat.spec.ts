// web/tests/e2e/report-threat.spec.ts
// The Threat tab's two additions: the target picker, which turns one table of totals into
// "who has the most threat on this enemy", and the taunt list under it, which answers
// "who took the boss off me, and when" in one click.
import { expect, test } from '@playwright/test';

const FIGHT = '/reports/fixture2abcd?fight=3&tab=threat';

/** Warden Kelthas, the only enemy fight 3's threat pairs name. */
const KELTHAS = 'Creature-0-2085-2284-7855-169754-0000AA0002';

test('threat on a target lists each player’s threat on the picked enemy, and the taunts', async ({
  page,
}) => {
  await page.goto(FIGHT);
  const picker = page.getByTestId('threat-target');
  await expect(picker).toBeVisible();
  await picker.selectOption({ label: 'Warden Kelthas' });
  await expect(page.getByTestId('threat-on-target')).toContainText('Baelgrim');
  await expect(page.getByTestId('threat-on-target').getByTestId('threat-share').first()).toContainText('%');
  await expect(page.getByTestId('threat-taunts')).toContainText('Taunt');
});

test('every enemy keeps the totals table, and picking one swaps in that enemy’s players', async ({
  page,
}) => {
  await page.goto(FIGHT);
  // "Every enemy" is the default, and it is the table that was there before the picker:
  // Warden Kelthas has a row of its own, because the totals count every unit's threat.
  await expect(page.getByTestId('threat-table')).toContainText('Warden Kelthas');
  await expect(page.getByTestId('threat-on-target')).toHaveCount(0);

  await page.getByTestId('threat-target').selectOption({ label: 'Warden Kelthas' });
  const rows = page.getByTestId('threat-on-target').locator('li');
  // The four players who built threat on the boss, largest first; the boss itself is not
  // one of its own attackers, so it has no row here.
  await expect(rows).toHaveCount(4);
  await expect(rows.first()).toContainText('Baelgrim');
  await expect(page.getByTestId('threat-on-target')).not.toContainText('Warden Kelthas');
  // 4,400 of 8,940 built on the boss.
  await expect(rows.first().getByTestId('threat-share')).toContainText('49.2%');
});

test('a source scope narrows the rows but not the share they are measured against', async ({ page }) => {
  // Scoped to Baelgrim alone, he is the only row on the boss -- and he still reads the
  // 49.2% he holds of every player's threat on it, not 100% of himself. The threat the
  // other three built is off screen, not gone.
  await page.goto(`${FIGHT}&source=Player-4184-000000A1&target=${KELTHAS}`);
  const rows = page.getByTestId('threat-on-target').locator('li');
  await expect(rows).toHaveCount(1);
  await expect(rows.first()).toContainText('Baelgrim');
  await expect(rows.first().getByTestId('threat-share')).toContainText('49.2%');
});

test('the picked enemy rides in the URL, so the link opens on it', async ({ page }) => {
  await page.goto(FIGHT);
  await page.getByTestId('threat-target').selectOption({ label: 'Warden Kelthas' });
  await expect(page).toHaveURL(new RegExp(`target=${KELTHAS}`));

  await page.goto(`${FIGHT}&target=${KELTHAS}`);
  await expect(page.getByTestId('threat-target')).toHaveValue(KELTHAS);
  await expect(page.getByTestId('threat-on-target')).toContainText('Baelgrim');
});

test('a taunt says who took what and when, and opens the window around it', async ({ page }) => {
  await page.goto(FIGHT);
  const taunts = page.getByTestId('threat-taunts');
  await expect(taunts).toContainText('8.5s');
  await expect(taunts).toContainText('Baelgrim taunted Warden Kelthas');
  await expect(taunts).toContainText('Taunt');

  await taunts.getByTestId('taunt-window').first().click();
  // Five seconds either side of 8,500ms, snapped out to whole seconds the way every
  // window in the report is.
  await expect(page).toHaveURL(/start=3000&end=14000/);
});

test('a fight with no taunt says so rather than showing an empty list', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=4&tab=threat');
  await expect(page.getByTestId('threat-taunts')).toContainText('No taunts in this pull.');
});

// A fight with no threat still has taunts to answer for: the empty state belongs to the
// table, not to the tab, or "who took the boss off me" goes unanswered on every pull the
// threat model happened to score at nothing.
test('a fight with no threat still says what its taunts were', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=2&tab=threat');
  await expect(page.getByTestId('table-empty')).toContainText('No threat in this window.');
  await expect(page.getByTestId('threat-taunts')).toContainText('No taunts in this pull.');
});

// Baelgrim is a Warrior, not fight 3's tank, but he is the one who taunts Warden Kelthas
// off the tank -- so the timeline names him "the taunter" rather than "the tank".
test('the timeline marks the taunt on the taunter’s lane and names it on hover', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3&view=timelines');
  const mark = page.getByTestId('lane-Player-4184-000000A1').getByTestId('taunt-mark').first();
  await expect(mark).toBeVisible();
  await mark.hover();
  await expect(page.getByTestId('timeline-picked')).toContainText('Taunt');
  // The boss lane mounts on the taunt alone: Warden Kelthas casts nothing in the fixture,
  // so without this guard the taunt had no boss lane to land on.
  await expect(page.getByTestId('lane-boss').getByTestId('taunt-mark').first()).toBeVisible();
});

test('the whole night names the pull each taunt came from', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=all&tab=threat');
  await expect(page.getByTestId('threat-taunts')).toContainText('Warden Kelthas · pull 1');
  // Over the night a pair's target is the enemy's name, one row per name across pulls.
  await page.getByTestId('threat-target').selectOption({ label: 'Warden Kelthas' });
  await expect(page.getByTestId('threat-on-target')).toContainText('Baelgrim');
});

// The night's Threat picker writes the enemy's NAME into `target` -- an add is a new GUID
// on every pull -- and the damage tabs share that key. They have to read it as a name too:
// read as a GUID it matched nothing, scaled every row to zero, and emptied the table while
// the filter bar still said "Every target".
test('an enemy picked on the night’s Threat tab keeps the damage tabs’ rows', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=all&tab=threat');
  await page.getByTestId('threat-target').selectOption({ label: 'Warden Kelthas' });
  await expect(page).toHaveURL(/target=Warden\+Kelthas/);

  await page.getByTestId('tab-damage-done').click();
  const rows = page.getByTestId('actor-table').locator('li');
  await expect(rows.first()).toBeVisible();
  await expect(page.getByTestId('table-empty')).toHaveCount(0);
  // And the bar names the enemy rather than reading "Every target" over a full table.
  await expect(page.locator('#filter-target')).not.toHaveValue('');
  await expect(page.locator('#filter-target').locator('option:checked')).toHaveText('Warden Kelthas');
});

test('the link alone carries the night’s named enemy into Damage Done', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=all&tab=damage-done&target=Warden+Kelthas');
  await expect(page.getByTestId('actor-table').locator('li').first()).toBeVisible();
  await expect(page.locator('#filter-target').locator('option:checked')).toHaveText('Warden Kelthas');
});

// Inside a brush the figures are the fight's totals scaled, which is not a standing: the
// table holds its bars and shares back and says why, and the taunt keeps the pull's clock
// so the same taunt reads the same time from either side of the link.
test('a brushed window shows greyed totals with no ranking, and the taunt keeps its time', async ({
  page,
}) => {
  await page.goto(`${FIGHT}&start=3000&end=14000`);
  await expect(page.getByTestId('threat-approximate-note')).toContainText('not measured yet');
  await expect(page.getByTestId('threat-share')).toHaveCount(0);
  await expect(page.getByTestId('threat-table').locator('ul[data-ranked="false"]')).toHaveCount(1);
  await expect(page.getByTestId('threat-taunts')).toContainText('8.5s');
});

test('a source scope shares the totals table against everyone too', async ({ page }) => {
  await page.goto(`${FIGHT}&source=Player-4184-000000A1`);
  const shares = page.getByTestId('threat-share');
  await expect(shares).toHaveCount(1);
  await expect(shares.first()).not.toContainText('100.0%');
});

// The totals table lists the enemies too, but the share is the raid's: every player's
// threat sums to 100% and the boss's own row reads a dash, or the raid would read half
// of itself with the rest held by the enemy it is fighting.
test('the totals table shares the players to 100% and leaves the enemies out of it', async ({ page }) => {
  await page.goto(FIGHT);
  const shares = await page.getByTestId('threat-share').allInnerTexts();
  const sum = shares.reduce((total, text) => total + Number.parseFloat(text), 0);
  expect(Math.round(sum * 10) / 10).toBeCloseTo(100, 0);
  // The fixture's boss builds no threat row of its own; the sample log's bosses do, and
  // the harness smoke checks their rows read a dash where a share would be.
  await expect(page.getByTestId('threat-share-note')).toContainText('every player’s threat');
});

// The chart is the answer to "did I pull it off the tank at 3:54": one cumulative line
// per player on the picked enemy, the taunts marked on the axis, and the same brush the
// rest of the page uses.
test('the threat chart draws one line per player and marks the taunts', async ({ page }) => {
  await page.goto(FIGHT);
  const chart = page.getByTestId('threat-chart');
  await expect(chart).toBeVisible();
  // Four players built threat on Warden Kelthas, so four legend entries.
  await expect(chart.getByTestId('chart-line-label')).toHaveCount(4);
  await expect(chart).toContainText('Baelgrim');
  // With no enemy picked the chart says which one it is drawing.
  await expect(page.getByTestId('threat-chart-note')).toContainText('Warden Kelthas');
  await expect(chart.getByTestId('chart-mark')).toHaveCount(1);
  await expect(chart.getByTestId('chart-mark').first()).toHaveAttribute('title', /Taunt/);
});

test('the threat chart’s own brush sets the report’s window', async ({ page }) => {
  await page.goto(FIGHT);
  await page.getByTestId('threat-chart').getByTestId('window-start').fill('6000');
  await expect(page).toHaveURL(/start=6000/);
});

test('the night has no threat chart, and says why', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=all&tab=threat');
  await expect(page.getByTestId('threat-chart')).toHaveCount(0);
  await expect(page.getByTestId('threat-chart-note')).toContainText('no clock');
});
