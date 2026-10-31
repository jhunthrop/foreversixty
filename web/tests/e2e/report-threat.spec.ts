// web/tests/e2e/report-threat.spec.ts
// The Threat tab's two additions: the target picker, which turns one table of totals into
// "who has the most threat on this enemy", and the taunt list under it, which answers
// "who took the boss off me, and when" in one click.
import { expect, test } from './support/test';
import { choose, fieldBox } from './support/select';

const FIGHT = '/reports/fixture2abcd?fight=3&tab=threat';

/** Warden Kelthas, the only enemy fight 3's threat pairs name. */
const KELTHAS = 'Creature-0-2085-2284-7855-169754-0000AA0002';

test('threat on a target lists each player’s threat on the picked enemy, and the taunts', async ({
  page,
}) => {
  await page.goto(FIGHT);
  const picker = page.getByTestId('threat-target');
  await expect(fieldBox(page, 'threat-target')).toBeVisible();
  await choose(picker, { label: 'Warden Kelthas' });
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

  await choose(page.getByTestId('threat-target'), { label: 'Warden Kelthas' });
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
  await choose(page.getByTestId('threat-target'), { label: 'Warden Kelthas' });
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
  await choose(page.getByTestId('threat-target'), { label: 'Warden Kelthas' });
  await expect(page.getByTestId('threat-on-target')).toContainText('Baelgrim');
});

// The night's Threat picker writes the enemy's NAME into `target` -- an add is a new GUID
// on every pull -- and the damage tabs share that key. They have to read it as a name too:
// read as a GUID it matched nothing, scaled every row to zero, and emptied the table while
// the filter bar still said "Every target".
test('an enemy picked on the night’s Threat tab keeps the damage tabs’ rows', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=all&tab=threat');
  await choose(page.getByTestId('threat-target'), { label: 'Warden Kelthas' });
  await expect(page).toHaveURL(/target=Warden\+Kelthas/);

  const phone = (page.viewportSize()?.width ?? 1280) < 1024;
  if (phone) await choose(page.getByTestId('tab-select'), 'damage-done');
  else await page.getByTestId('tab-damage-done').click();
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

// Inside a brush the table is measured from the pairs' own per-second series: standing at
// the window's end is what decides aggro, built inside it is what the window did. Both
// are measured, so there is no tilde, no "no order to read" note, and the ranking stands.
test('a brushed window reads standing and built from the series, measured', async ({ page }) => {
  await page.goto(`${FIGHT}&start=3000&end=14000&target=${KELTHAS}`);
  await expect(page.getByTestId('threat-approximate-note')).toHaveCount(0);
  await expect(page.getByTestId('threat-window-note')).toContainText('Standing');
  const rows = page.getByTestId('threat-on-target').locator('li');
  await expect(rows.first()).toContainText('Baelgrim');
  await expect(rows.first().getByTestId('threat-standing')).not.toContainText('~');
  await expect(rows.first().getByTestId('threat-built')).toBeVisible();
  await expect(page.getByTestId('threat-share')).not.toHaveCount(0);
  await expect(page.getByTestId('threat-table').locator('ul[data-ranked="true"]')).toHaveCount(1);
  // Shares are of the players' standings, so no row reads past 100% and the top row reads less
  // than everyone put together.
  const shares = await page.getByTestId('threat-share').allTextContents();
  for (const text of shares) expect(Number.parseFloat(text)).toBeLessThanOrEqual(100);
  // The taunt keeps the pull's clock, so the same taunt reads the same time from either side.
  await expect(page.getByTestId('threat-taunts')).toContainText('8.5s');
});

test('the whole fight is unchanged: standing at the end is the total', async ({ page }) => {
  await page.goto(`${FIGHT}&target=${KELTHAS}`);
  const rows = page.getByTestId('threat-on-target').locator('li');
  await expect(rows.first().getByTestId('threat-share')).toContainText('49.2%');
  // No second figure when there is no window: the whole fight has only one number.
  await expect(page.getByTestId('threat-built')).toHaveCount(0);
});

// "Around it" on a taunt sets the window; the table it lands on is measured, not prorated.
test('a taunt’s window lands on a measured table', async ({ page }) => {
  await page.goto(`${FIGHT}&target=${KELTHAS}`);
  await page.getByTestId('threat-taunts').getByTestId('taunt-window').first().click();
  await expect(page).toHaveURL(/start=3000&end=14000/);
  await expect(page.getByTestId('threat-window-note')).toContainText('Standing');
  await expect(page.getByTestId('threat-approximate-note')).toHaveCount(0);
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
  const shareCells = page.getByTestId('threat-share');
  // allInnerTexts() reads whatever is on screen the instant it is called, with no wait of
  // its own -- a bare navigation racing the table's first paint can read zero rows and
  // pass a sum of 0 as "not yet 100%", never getting to fail loudly on a real regression.
  await expect(shareCells.first()).toBeVisible();
  const shares = await shareCells.allInnerTexts();
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
  // One chart per tab (spec §B1.7): the main chart is not drawn on the Threat tab.
  await expect(page.getByTestId('time-chart')).toHaveCount(0);
  // Four players built threat on Warden Kelthas, so four legend chips.
  await expect(chart.getByTestId('legend-chip')).toHaveCount(4);
  await expect(chart).toContainText('Baelgrim');
  // The chart says which enemy it draws, and that it is an estimate.
  // "Every enemy" is picked, so the title says so (fix round 2): the title and the field agree.
  await expect(chart.getByRole('heading')).toHaveText('Estimated threat, every enemy');
  const taunts = chart.locator('[data-testid="lane-marker"][data-kind="taunt"]');
  await expect(taunts).toHaveCount(1);
  await expect(taunts.first()).toHaveAttribute('aria-label', /Taunt · Baelgrim on Warden Kelthas/);
  // The lines are a running total, not a rate: the head says so.
  // (On a phone the title takes the head; the brush carries the window there.)
  if ((page.viewportSize()?.width ?? 1280) >= 768)
    await expect(chart.getByTestId('chart-head-meta')).toHaveText('Cumulative · Whole fight');
  // The caveat line says what the estimate leaves out and links to the glossary's entry.
  await expect(chart.getByTestId('threat-caveat').getByRole('link')).toHaveAttribute(
    'href',
    '#glossary-threat',
  );
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

test('the glossary explains standing and built', async ({ page }) => {
  await page.goto(FIGHT);
  await page.getByTestId('glossary').locator('summary').click();
  await expect(page.getByTestId('glossary')).toContainText('Standing threat');
  await expect(page.getByTestId('glossary')).toContainText('Threat built in a window');
});

// Owner 2026-10-30: the open "More" menu disappeared under the threat chart. The tab row is
// sticky, so it is its own stacking context; the menu's z-index only ranked inside it until
// the row carried one too. The topmost element at each menu item's centre must be that item.
test('the open More menu sits above the threat chart', async ({ page }) => {
  test.skip((page.viewportSize()?.width ?? 1280) < 1024, 'the More menu is the desktop tab row');
  await page.goto(FIGHT);
  await expect(page.getByTestId('threat-chart')).toBeVisible();
  await page.getByTestId('tab-more').click();
  const items = page.getByTestId('tab-more-menu').locator('[role="tab"], a, button');
  const count = await items.count();
  expect(count).toBeGreaterThan(0);
  for (let i = 0; i < count; i += 1) {
    const item = items.nth(i);
    const box = await item.boundingBox();
    if (box === null) continue;
    const onTop = await page.evaluate(
      ({ x, y }) => document.elementFromPoint(x, y)?.closest('[data-testid="tab-more-menu"]') !== null,
      { x: box.x + box.width / 2, y: box.y + box.height / 2 },
    );
    expect(onTop, `menu item ${i} is covered`).toBe(true);
  }
});
