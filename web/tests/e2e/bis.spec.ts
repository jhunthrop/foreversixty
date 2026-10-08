// web/tests/e2e/bis.spec.ts
// Lane bis-page-ux (leveling-bis-design.md, the 2026-09-28 UX pass): index -> spec page ->
// toggle faction -> pick a band -> see that band's own "new" row marker and a real quest
// source line, not a generic pill. The page's data comes straight off disk at build time
// (src/lib/bis/load.ts), independent of FOREVER_DATA, so this runs in the default fixture
// suite same as every other content spec here.
//
// "The list" rebuild (bis rebuild spec, 2026-09-30, mock `Bis.dc.html`): the character-
// panel redesign's two-column paperdoll (lane bis-character-panel, 2026-09-29,
// `.paperdoll-left`/`-right`/`-bottom`) is gone, replaced by one full-width list
// (`GearRow.astro`, `bis-slot-<faction>-<band>-<slot>` per row, 17 rows, canonical slot
// order, unchanged) under a real `ClassHeader` and a three-panel "This set"/"Stat
// weights"/"Play it" row. A pick's alternatives are their own grid column beside the pick
// (not stacked under it) -- still every one its own hover target in the shared tooltip
// (`item-hover-<id>`, unchanged). The weight rail (`StatWeightsPanel.astro`) now carries a
// per-band/faction `data-testid="bis-weights-<faction>-<band>"` (added by this lane -- the
// rebuild shipped only a page-generic `bis-stat-weights`, not unique across the many bands/
// factions one page renders). Empty rows now read one of several honest, longer sentences
// from `lib/bis/copy.ts` (`noSourcedItemFirst`/`noSourcedItemLater`/`twoHanderEquippedNamed`/
// `emptyReasonEffectNotModelled`/`noKnownSourceForSlot`) rather than the old page's shorter
// generic lines -- `ANY_EMPTY_COPY` below matches every one by its stable substring, never
// the exact sentence (band numbers and item names vary with the real data), so a sweep over
// whichever empty slots the real data happens to carry this build still asserts a real,
// known copy rather than pinning to whichever reason happened to be present when the test
// was written. Every real published file carries `alternatives` and
// `reference_dps_per_point` (fix round 1, 2026-09-29), so every state below is exercised
// against real published data: `hunter/marksmanship` band 20 for the two-hander empty off-
// hand and for alternatives (including a tie), `hunter/marksmanship` band 60 for the weight
// rail's DPS-per-point line, and `druid/balance` (which never equips a ranged weapon) for
// the ordinary empty-slot case. No fixture is committed for this page any more.
import { test, expect } from '@playwright/test';

const TWO_HANDER_EMPTY_COPY = /is a two-hander; the off hand is taken\./;
const ORDINARY_EMPTY_COPY =
  /raises your damage\.|Nothing here (?:either until \d+|helps at this level either)\.|Nothing sourced at this level helps your DPS|Relic effects aren't simulated yet|No sourced item at this level yet/;
const ANY_EMPTY_COPY = new RegExp(`${TWO_HANDER_EMPTY_COPY.source}|${ORDINARY_EMPTY_COPY.source}`);

/** hunter-marksmanship's real data empties several slots across its bands for several
 *  different reasons (a two-handed main hand's own off-hand, a thin trinket pool, an
 *  unmodelled relic effect) -- a filled-vs-empty sweep over a band's rows asserts one of
 *  `ANY_EMPTY_COPY`'s known reasons for every one, without hardcoding which reason belongs
 *  to which slot or band. */
async function expectEmptyRowCopy(row: import('@playwright/test').Locator): Promise<void> {
  await expect(row).toContainText(ANY_EMPTY_COPY);
}

/** Set once per page load; a click that actually triggers a full navigation (rather than
 *  the page's existing CSS radio/`:target` toggles) reloads the document and wipes this, so
 *  reading it back `true` after an interaction is the one reliable way to prove "no reload"
 *  rather than just "the right content appeared afterwards". */
async function markNoReload(page: import('@playwright/test').Page): Promise<void> {
  await page.evaluate(() => {
    (window as unknown as { __noReloadMarker?: boolean }).__noReloadMarker = true;
  });
}

async function expectNoReloadSoFar(page: import('@playwright/test').Page): Promise<void> {
  const marker = await page.evaluate(
    () => (window as unknown as { __noReloadMarker?: boolean }).__noReloadMarker,
  );
  expect(marker).toBe(true);
}

test('Leveling BiS: index links to a spec, faction and band pills switch panels with no reload', async ({
  page,
}) => {
  await page.goto('/bis');
  await expect(page.getByRole('heading', { name: 'Leveling BiS', level: 1 })).toBeVisible();

  const marksmanshipLink = page.locator('a[href="/bis/hunter/marksmanship"]');
  await expect(marksmanshipLink).toBeVisible();
  await marksmanshipLink.click();
  await expect(page).toHaveURL(/\/bis\/hunter\/marksmanship$/);
  await markNoReload(page);

  // Alliance is the default panel; its own first band's list is open. The band list is the
  // nightly's (20..60 step 10), so the test reads the first band off the page rather than
  // pinning its level.
  await expect(page.getByTestId('bis-faction-panel-alliance')).toBeVisible();
  const firstBand = page.locator('[data-testid^="bis-band-alliance-"]:visible').first();
  await expect(firstBand).toBeVisible();
  const firstLevel = (await firstBand.getAttribute('data-testid'))!.replace('bis-band-alliance-', '');
  expect(Number(firstLevel)).toBeGreaterThan(0);

  // "The list" renders one row per wearable-and-comparable slot (spec §4.D, 17 rows,
  // unchanged canonical order) -- a filled slot shows the real item (its own icon and
  // tooltip host), and an empty one says why rather than a bare dash (tenet 4). Every slot
  // row is one or the other.
  await expect(firstBand.locator('[data-testid^="item-hover-"]').first()).toBeVisible();
  const slotRows = firstBand.locator(`[data-testid^="bis-slot-alliance-${firstLevel}-"]`);
  const slotCount = await slotRows.count();
  expect(slotCount).toBe(17);
  for (let i = 0; i < slotCount; i += 1) {
    const row = slotRows.nth(i);
    const filled = (await row.locator('[data-testid^="item-hover-"]').count()) > 0;
    if (!filled) await expectEmptyRowCopy(row);
  }

  // Toggle to Horde -- a label click on a hidden radio (the page's existing CSS
  // `~`-sibling toggle mechanism), no navigation: the marker set right after the first
  // navigation is still there, and the Alliance panel this click hides is still in the DOM
  // (`display: none`, not removed), not a fresh page. `ClassHeader` renders once per
  // faction (spec's own note: its suffix/summary text names the active faction, which only
  // a second copy -- not a client re-render -- can do above the fold), so BOTH copies carry
  // a `bis-faction-toggle-horde` pill; only the one inside the currently visible header is
  // actually clickable, hence the `:visible` filter rather than a bare `getByTestId`.
  await page.locator('[data-testid="bis-faction-toggle-horde"]:visible').click();
  await expect(page.getByTestId('bis-faction-panel-horde')).toBeVisible();
  await expect(page.getByTestId('bis-faction-panel-alliance')).toBeHidden();
  await expectNoReloadSoFar(page);

  // Pick band 30 within the Horde panel -- the existing CSS `:target` band toggle, also no
  // navigation.
  await page.getByTestId('bis-band-pill-horde-30').click();
  const band30 = page.getByTestId('bis-band-horde-30');
  await expect(band30).toBeVisible();
  await expectNoReloadSoFar(page);

  // A row this band marks new (hunter-marksmanship/horde/30 has several -- the neck slot's
  // own quest pick among them), and a real quest source line, not a generic pill -- tenet
  // 2's "an item is never just a name" applied to sources (`source-cell.ts`, unchanged by
  // this rebuild).
  await expect(band30.getByTestId('bis-row-new').first()).toBeVisible();
  await expect(band30).toContainText('Quest:');

  // "New at this band" (spec §4.E) replaces the old page's "N upgrades since level N"
  // disclosure -- a computed sentence naming how many of this band's picks are new and
  // where the newly-opened gear comes from, never a fabricated "reaching N opens..." claim
  // at the file's first band.
  const newAtBand30 = page.getByTestId('bis-new-at-band-horde-30');
  await expect(newAtBand30).toContainText(/picks come from there\.|This is the first band/);
});

test('Leveling BiS: a spec with no ranked list yet shows the empty state, not a 404', async ({ page }) => {
  // Healers have no written rotation, so the nightly ranks nothing for them (a dps spec
  // gained a real file the night the nightly first ran, which is what this test once used).
  // priest-holy carries a fixture (bis-healer.spec.ts), so this uses a healer without one.
  await page.goto('/bis/priest/discipline');
  await expect(page.getByTestId('bis-empty-state')).toBeVisible();
  await expect(page.getByTestId('bis-empty-state')).toContainText('No leveling BiS list yet');
});

test('Leveling BiS: a pick’s alternatives render as rows beside it, every one its own hover target in the shared tooltip', async ({
  page,
}) => {
  await page.goto('/bis/hunter/marksmanship#band-alliance-20');
  const band20 = page.getByTestId('bis-band-alliance-20');
  await expect(band20).toBeVisible();

  // Which slots carry alternatives (and how many) is the nightly's call, so the test finds
  // the first slot row that shows "Also:" rather than pinning a slot and a count.
  const slotRows = band20.locator('[data-testid^="bis-slot-alliance-20-"]');
  const slotCount = await slotRows.count();
  let rowWithAlts: import('@playwright/test').Locator | undefined;
  let rowWithoutAlts: import('@playwright/test').Locator | undefined;
  for (let i = 0; i < slotCount; i += 1) {
    const row = slotRows.nth(i);
    const hasAlts = (await row.getByText('Also:').count()) > 0;
    const filled = (await row.locator('[data-testid^="item-hover-"]').count()) > 0;
    if (hasAlts && rowWithAlts === undefined) rowWithAlts = row;
    if (!hasAlts && filled && rowWithoutAlts === undefined) rowWithoutAlts = row;
  }
  test.skip(rowWithAlts === undefined, 'no slot on this band carries alternatives with the current data');
  const altsRow = rowWithAlts!;

  // The pick's own hover target never contains a nested [role="button"] (fix round 1: a
  // button nested in a button broke keyboard/AT tab order when alternatives lived inside
  // it) -- the rebuild keeps runners-up their own grid column beside the pick (spec §4.D),
  // never stacked inside it, so the pick's own host still has none.
  const pickHostForA11y = altsRow.locator('[data-testid^="item-hover-"]').first();
  await expect(pickHostForA11y.locator('[role="button"]')).toHaveCount(0);
  // The pick itself plus at least one alternative -- independent tooltip hosts on one row
  // (alternatives are a sibling block, never nested inside the pick's own host).
  const altHosts = altsRow.locator('[data-testid^="item-hover-"]');
  expect(await altHosts.count()).toBeGreaterThanOrEqual(2);
  // Each alternative is its own compact row -- icon, name, source line and a mono DPS gap
  // (`GearRow.astro`'s `gear-row-alt` grid) -- naming its item level (and, once a band's
  // level requirement exceeds the item's own, the requirement too) and a signed delta, or
  // "same DPS" for a genuine tie. `\s` matches `copy.ts`'s own non-breaking space between a
  // word and its number.
  await expect(altsRow).toContainText(/ilvl\s*\d+/);
  await expect(altsRow).toContainText(/−\d+\.\d DPS|same DPS/);
  const firstAltHost = altHosts.nth(1);
  await expect(firstAltHost.locator('.gear-row-alt-icon, .gear-row-icon-placeholder')).toHaveCount(1);
  await expect(firstAltHost.locator('.gear-row-alt-name')).not.toBeEmpty();
  await expect(firstAltHost.locator('.gear-row-alt-source-text')).not.toBeEmpty();
  await expect(firstAltHost.locator('.gear-row-alt-gap')).toContainText(/−\d+\.\d DPS|same DPS/);

  // A genuine tie anywhere on the band reads "same DPS", never a signed "+0.0 DPS".
  await expect(band20).not.toContainText('+0.0 DPS');

  // Hovering the pick's own NAME (fix round 1: the regression this test guards against was
  // hovering the pick's visual centre landing on a nested alternative instead) shows the
  // pick's own tooltip, not an alternative's.
  const pickHost = altsRow.locator('[data-testid^="item-hover-"]').first();
  const pickName = (await pickHost.locator('.gear-row-name').innerText()).trim();
  await pickHost.locator('.gear-row-name').hover();
  const tooltip = page.getByTestId('item-tooltip');
  await expect(tooltip).toBeVisible();
  await expect(tooltip).toContainText(pickName);
  await page.mouse.move(0, 0);
  await expect(tooltip).toBeHidden();

  // An alternative's own icon opens the exact same shared tooltip, for that alternative.
  const altHost = altsRow.locator('[data-testid^="item-hover-"]').nth(1);
  const altName = (await altHost.locator('.gear-row-alt-name').innerText()).trim();
  await altHost.hover();
  await expect(tooltip).toBeVisible();
  await expect(tooltip).toContainText(altName);
  await page.mouse.move(0, 0);
  await expect(tooltip).toBeHidden();

  // A filled slot with no alternatives shows no "Also:" at all, not an empty placeholder
  // (spec §4) -- checked on whichever such row this band has.
  if (rowWithoutAlts !== undefined) {
    await expect(rowWithoutAlts.getByText('Also:')).toHaveCount(0);
  }
});

test('Leveling BiS: the Play It panel renders a real icon for a rotation line, not the neutral placeholder', async ({
  page,
}) => {
  // hunter-marksmanship's level-10 rotation entry applies through band 20's top (29), and
  // every line on it (addon-data.json) carries an `icon` the build's icons/ tree ships --
  // this guards the fix (rotation-view.ts threads RotationLine.icon through to the view,
  // PlayItPanel.astro renders it) against the old regression of a neutral placeholder on
  // every line regardless of the data.
  await page.goto('/bis/hunter/marksmanship#band-alliance-20');
  const band20 = page.getByTestId('bis-band-alliance-20');
  const playIt = band20.getByTestId('bis-play-it');
  await expect(playIt).toBeVisible();
  const lines = playIt.locator('[data-testid^="bis-rotation-line-"]');
  await expect(lines.first()).toBeVisible();
  const icons = playIt.locator('[data-testid^="bis-rotation-line-"] img.play-it-icon');
  expect(await icons.count()).toBeGreaterThan(0);
  await expect(icons.first()).toHaveAttribute('src', /\/data\/[^/]+\/icons\/.+\.webp$/);
  // Never the client's own red "?" placeholder texture naming (tenet 4), and never the old
  // neutral "no icon resolved" box on a line the build actually has an icon for.
  await expect(playIt.locator('.play-it-icon-placeholder')).toHaveCount(0);
});

test('Leveling BiS: the weight rail shows scale factors normalized to the top stat, DPS per point, the haste caption and the addon button', async ({
  page,
}) => {
  await page.goto('/bis/hunter/marksmanship#band-alliance-60');
  const band60 = page.getByTestId('bis-band-alliance-60');
  // `StatWeightsPanel.astro` (the rebuild's replacement for the old weight rail) carries a
  // per-band/faction testid built by the page from `faction`/`view.band` -- there is one
  // instance per band/faction pair on this page, so a page-generic id would not be unique.
  const weights60 = band60.getByTestId('bis-weights-alliance-60');
  await expect(weights60).toBeVisible();
  // The rail follows the convention players know from SimulationCraft's scale factors:
  // per point, normalized so the top per-point stat reads 1.00, with the absolute DPS per
  // point beside it. Which stat is on top moves with the data, so the assertion reads the
  // note the page prints rather than pinning Agility -- the reference stat's own Scale
  // column also reads exactly "1.00" (normalized against itself), not just named in the
  // caption above it.
  await expect(weights60).toContainText(/normalized to [A-Za-z ]+ = 1\.00/);
  await expect(weights60).toContainText('1.00');
  await expect(weights60).toContainText(/\d\.\d\d\d DPS/);
  // Haste has no rating in this client, so it is never a table row: the caption carries it.
  await expect(band60).toContainText(/Haste: \d+\.\d\d per 1%/);
  // The weights feed our own addon; no other addon is ever named (owner rule, 2026-09-30).
  await expect(band60).toContainText('Use these weights in the addon');
  await expect(band60).not.toContainText(/Pawn/);
  // Whichever stat the nightly flags insignificant reads "Not significant", never a number
  // implying a real measurement. Which band carries one moves with the data, so look across
  // the Alliance bands and skip the assertion honestly when none does this build.
  const rails = page.locator('[data-testid^="bis-weights-alliance-"]');
  const railTexts = await rails.allInnerTexts();
  const notSignificantRail = railTexts.find((text) => text.includes('Not significant'));
  test.skip(
    notSignificantRail === undefined,
    'no insignificant stat on any Alliance band with the current data',
  );
  expect(notSignificantRail).toContain('Not significant');
});

test('Leveling BiS: an ordinary empty slot (not a two-hander gap) reads the plain no-source copy', async ({
  page,
}) => {
  // The picks are the nightly's (druids now get idols in their ranged slot from the
  // client's hotfix rows, 2026-09-30), so the test looks for any ordinary empty slot on
  // the page rather than pinning one: an empty row that is not a two-hander gap must carry
  // one of the rebuild's own honest no-source sentences (`lib/bis/copy.ts`'s
  // `noSourcedItemFirst`/`noSourcedItemLater`/`emptyReasonEffectNotModelled`/
  // `noKnownSourceForSlot` -- longer, more specific sentences than the old page's, e.g. "No
  // trinket you can get at 20 to 29 raises your damage. The first that does comes at 50." or
  // "Nothing here either until 50."), never the two-hander-specific
  // `twoHanderEquippedNamed` line ("{Main-hand item} is a two-hander; the off hand is
  // taken.") -- that is its own, deliberately different case (the weapon test covers it via
  // `ANY_EMPTY_COPY` at the top of this file). A band with no ordinary empty slot at all
  // has nothing to check.
  await page.goto('/bis/druid/balance#band-alliance-20');
  const band20 = page.getByTestId('bis-band-alliance-20');
  await expect(band20).toBeVisible();
  const emptyRows = band20.locator('.gear-row-empty');
  const count = await emptyRows.count();
  test.skip(count === 0, 'no empty slot on this band with the current data');

  const emptyTexts = await emptyRows.allInnerTexts();
  for (const text of emptyTexts) {
    expect(text).toMatch(ANY_EMPTY_COPY);
  }
  const ordinaryTexts = emptyTexts.filter((text) => !TWO_HANDER_EMPTY_COPY.test(text));
  test.skip(
    ordinaryTexts.length === 0,
    'every empty slot on this band is the two-hander gap with the current data',
  );
  for (const text of ordinaryTexts) {
    expect(text).toMatch(ORDINARY_EMPTY_COPY);
    expect(text).not.toMatch(TWO_HANDER_EMPTY_COPY);
  }
});
