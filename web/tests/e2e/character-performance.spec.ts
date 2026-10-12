// web/tests/e2e/character-performance.spec.ts
// Task L8.5: the Performance Rating card and the profile's leftovers.
//   Ruling 1. The rating card takes the Forever Score card's treatment: two columns from 1280,
//             the sparkline filling the right one, "Up 9 since 4 Dec" under it from the same
//             model, one verdict line that agrees with the score card's Parses row, and the
//             weakest part linking to its fight's row in "Every ranked fight".
//   Ruling 2. Experience notes what counts.
//   Ruling 3. The 404 card is capped at 720px; the guild link's focus ring stands 2px off and
//             the dots either side stay clear of it.
//   Ruling 4. Below 400px a part's row reads "412/500" with an 80px label column.
//   Ruling 5. /setup says "type /fs copy, press Ctrl+C".
import { expect, test, type Page } from '@playwright/test';
import { CHARACTER, CHARACTER_PATH, CHARACTER_URL } from './support/character-fixture';

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
  as_of: '2026-12-09',
  components: { gear: 412, experience: 620, parses: 540, attendance: 180, discipline: 78 },
  experience: { dungeons: 14, raid_kills: 0, alts: 2 },
};

/** Six rated fights, 4 to 9 Dec, 56 up to 65; the newest is the fixture's Warden Kelthas row in
 *  "Every ranked fight" (fixture2abcd, fight 3), the one before it its Deep Warden row. */
const TREND = [
  { fought_at: '2026-12-04T20:00:00Z', overall: 56, report_id: 'olderreport1', fight_index: 1 },
  { fought_at: '2026-12-05T20:00:00Z', overall: 58, report_id: 'olderreport1', fight_index: 2 },
  { fought_at: '2026-12-06T20:00:00Z', overall: 60, report_id: 'olderreport2', fight_index: 1 },
  { fought_at: '2026-12-07T20:00:00Z', overall: 61, report_id: 'olderreport2', fight_index: 2 },
  { fought_at: '2026-12-08T20:00:00Z', overall: 63, report_id: 'otherreport1', fight_index: 1 },
  { fought_at: '2026-12-09T22:10:00Z', overall: 65, report_id: 'fixture2abcd', fight_index: 3 },
];

const RATING = {
  player_key: 'us/hardcore/elyra-duskvale',
  sample_size: 6,
  // Newest first, as the API sends it.
  trend: [...TREND].reverse(),
  best_component: 'preparation',
  worst_component: 'activity',
  latest: {
    player_key: 'us/hardcore/elyra-duskvale',
    player_name: 'Elyra Duskvale',
    class: 'Priest',
    spec: 'Discipline',
    role: 'healer',
    overall: 65,
    overall_uncapped: 65,
    overall_capped: false,
    basis: 'percentile',
    components: [],
  },
};

async function open(page: Page): Promise<void> {
  await page.route(CHARACTER_URL, (route) => route.fulfill(ok({ ...CHARACTER.data, ...SCORE })));
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(ok(RATING)));
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('rating-card')).toBeVisible();
  await expect(page.getByTestId('character-history')).toBeVisible();
}

/** The rating card's parts, read in one pass so a settling layout cannot split them. */
async function cardBoxes(page: Page) {
  return page.evaluate(() => {
    const box = (testid: string) => {
      const element = document.querySelector(`[data-testid="${testid}"]`);
      if (element === null) return null;
      const r = element.getBoundingClientRect();
      return { x: r.x, y: r.y, width: r.width, height: r.height, right: r.right, bottom: r.bottom };
    };
    const card = document.querySelector('[data-testid="rating-card"]') as HTMLElement;
    const style = getComputedStyle(card);
    const r = card.getBoundingClientRect();
    const inset = (side: 'left' | 'right') =>
      parseFloat(style.getPropertyValue(`padding-${side}`)) +
      parseFloat(style.getPropertyValue(`border-${side}-width`));
    return {
      content: { left: r.left + inset('left'), right: r.right - inset('right') },
      main: box('rating-card-main'),
      side: box('rating-card-side'),
      plot: box('rating-card-plot'),
    };
  });
}

for (const width of [1440, 1280]) {
  test.describe(`the rating card at ${width} (L8.5 ruling 1)`, () => {
    test.use({ viewport: { width, height: 900 } });
    test.skip(() => test.info().project.name !== 'desktop', 'a desktop width');

    test('is two columns, the sparkline filling the right one to the card’s content edge', async ({
      page,
    }) => {
      await open(page);
      const { content, main, side, plot } = await cardBoxes(page);
      if (main === null || side === null || plot === null) throw new Error('a rating card part is missing');
      expect(side.x).toBeGreaterThanOrEqual(main.right);
      expect(Math.abs(side.y - main.y)).toBeLessThan(1);
      expect(Math.abs(plot.width - side.width)).toBeLessThan(1);
      expect(Math.abs(side.right - content.right)).toBeLessThan(1);
      expect(Math.abs(main.x - content.left)).toBeLessThan(1);
      // The score card's plot height from 1024px.
      expect(plot.height).toBe(48);
    });
  });
}

test.describe('the rating card at 1440 (L8.5 ruling 1)', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'a desktop width');

  test('captions the sparkline "Up 9 since 4 Dec" and dates its right edge "9 Dec"', async ({ page }) => {
    await open(page);
    const caption = page.getByTestId('rating-card-history-caption');
    await expect(caption).toHaveText('Up 9 since 4 Dec');
    await expect(caption).toHaveCSS('font-size', '12px');
    await expect(caption).toHaveCSS('color', 'rgb(233, 228, 216)');
    await expect(page.getByTestId('rating-card-history-last')).toHaveText('9 Dec');
    await expect(page.getByTestId('rating-trend')).toHaveAttribute(
      'aria-label',
      'Performance Rating, last 6 fights: 56 on 4 Dec, 65 on 9 Dec',
    );
  });

  test('says one verdict that agrees with the score card’s Parses row', async ({ page }) => {
    await open(page);
    await expect(page.getByTestId('rating-card-verdict')).toHaveText(
      'Rated 65 of 100 over the last 6 fights · counts toward Parses 540 of 600',
    );
    await expect(page.getByTestId('forever-score-parses')).toContainText('Parses');
    await expect(page.getByTestId('forever-score-parses')).toContainText('540 of 600');
  });

  test('takes "Weakest part of a fight: Activity." to its fight’s row in "Every ranked fight"', async ({
    page,
  }) => {
    await open(page);
    const weakest = page.getByTestId('rating-card-weakest');
    await expect(weakest).toHaveText('Weakest part of a fight: Activity.');
    const link = weakest.getByRole('link', { name: 'Activity' });
    await expect(link).toHaveAttribute('href', '#fight-fixture2abcd-3');
    await link.click();
    await expect(page).toHaveURL(/#fight-fixture2abcd-3$/);
    const row = page.getByTestId('character-history').locator('#fight-fixture2abcd-3');
    await expect(row).toContainText('Warden Kelthas');
    await expect(row).toBeInViewport();
  });
});

test.describe('the rating card on a 390px phone (L8.5 ruling 1)', () => {
  test.use({ viewport: { width: 390, height: 844 } });

  test('is one column, the sparkline the card’s whole content width under the verdict', async ({ page }) => {
    await open(page);
    const { content, main, side, plot } = await cardBoxes(page);
    if (main === null || side === null || plot === null) throw new Error('a rating card part is missing');
    expect(side.y).toBeGreaterThanOrEqual(main.bottom - 0.5);
    expect(Math.abs(plot.x - content.left)).toBeLessThan(1);
    expect(Math.abs(plot.right - content.right)).toBeLessThan(1);
    expect(plot.height).toBe(40);
    await expect(page.getByTestId('rating-card-history-caption')).toHaveText('Up 9 since 4 Dec');
    await expect(page.getByTestId('rating-card-verdict')).toBeVisible();
    expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(390);
  });

  test('gives the weakest part’s link a 44px target that is the link’s own', async ({ page }) => {
    await open(page);
    const link = page.getByTestId('rating-card-weakest').getByRole('link', { name: 'Activity' });
    await link.scrollIntoViewIfNeeded();
    const hits = await link.evaluate((element) => {
      const r = element.getBoundingClientRect();
      const x = r.left + r.width / 2;
      return {
        height: r.height,
        top: element.contains(document.elementFromPoint(x, r.top + 0.5)),
        bottom: element.contains(document.elementFromPoint(x, r.bottom - 0.5)),
      };
    });
    expect(hits.height).toBeGreaterThanOrEqual(44);
    expect(hits.top).toBe(true);
    expect(hits.bottom).toBe(true);
  });
});

test('notes under Experience what counts, muted at 12px (L8.5 ruling 2)', async ({ page }) => {
  await open(page);
  const note = page.getByTestId('forever-score-experience-note');
  await expect(note).toHaveText(
    'Counts dungeon and raid kills, alts included: 25 a different dungeon, 40 a different raid boss, up to 700.',
  );
  await expect(note).toBeVisible();
  await expect(note).toHaveCSS('font-size', '12px');
  await expect(note).toHaveCSS('color', 'rgb(154, 148, 132)');
  const detail = await page.getByTestId('forever-score-experience-detail').boundingBox();
  const box = await note.boundingBox();
  expect(box?.y ?? 0).toBeGreaterThanOrEqual((detail?.y ?? 0) + (detail?.height ?? 0) - 0.5);
});

test.describe('the 404 page at 1440 (L8.5 ruling 3)', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'a desktop width');

  test('caps the no-logs card at 720px', async ({ page }) => {
    await page.route(CHARACTER_URL, (route) => route.fulfill(NOT_FOUND));
    await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(NOT_FOUND));
    await page.goto(CHARACTER_PATH);
    await expect(page.getByTestId('character-unlogged')).toBeVisible();
    const card = await page.getByTestId('forever-score-card').boundingBox();
    expect(card?.width).toBe(720);
  });
});

test('the guild link’s 2px-offset focus ring leaves the dots either side clear (L8.5 ruling 3)', async ({
  page,
}) => {
  await open(page);
  const link = page.getByTestId('character-guild-link');
  await link.focus();
  await expect(link).toBeFocused();
  await expect(link).toHaveCSS('outline-offset', '2px');
  const clear = await link.evaluate((element) => {
    const style = getComputedStyle(element);
    const reach = parseFloat(style.outlineOffset) + parseFloat(style.outlineWidth);
    const r = element.getBoundingClientRect();
    /** The "·" glyph boxes in the header line, in reading order. */
    const line = element.closest('[data-testid="character-descriptor"]') as HTMLElement;
    const walker = document.createTreeWalker(line, NodeFilter.SHOW_TEXT);
    const dots: DOMRect[] = [];
    for (let node = walker.nextNode(); node !== null; node = walker.nextNode()) {
      const text = node.textContent ?? '';
      for (let i = text.indexOf('·'); i !== -1; i = text.indexOf('·', i + 1)) {
        const range = document.createRange();
        range.setStart(node, i);
        range.setEnd(node, i + 1);
        dots.push(range.getBoundingClientRect());
      }
    }
    const before = dots.filter((dot) => dot.right <= r.left + 1).at(-1);
    const after = dots.find((dot) => dot.left >= r.right - 1);
    return {
      reach,
      gapBefore: before === undefined ? null : r.left - reach - before.right,
      gapAfter: after === undefined ? null : after.left - (r.right + reach),
    };
  });
  expect(clear.reach).toBe(4);
  expect(clear.gapBefore).not.toBeNull();
  expect(clear.gapAfter).not.toBeNull();
  expect(clear.gapBefore ?? -1).toBeGreaterThanOrEqual(1);
  expect(clear.gapAfter ?? -1).toBeGreaterThanOrEqual(1);
});

for (const width of [390, 360]) {
  test.describe(`the score parts at ${width} (L8.5 ruling 4)`, () => {
    test.use({ viewport: { width, height: 844 } });

    test('read "412/500" in an 80px label column, leaving the track room', async ({ page }) => {
      await open(page);
      const gear = page.getByTestId('forever-score-gear');
      await expect(gear.getByText('412/500', { exact: true })).toBeVisible();
      await expect(gear.getByText('412 of 500', { exact: true })).toBeHidden();
      const columns = await gear.evaluate((row) =>
        [...row.children].slice(0, 3).map((cell) => cell.getBoundingClientRect().width),
      );
      expect(columns[0]).toBe(80);
      expect(columns[1]).toBeGreaterThanOrEqual(60);
      // "412/500" fits its column on one line.
      const value = await gear.getByText('412/500', { exact: true }).evaluate((element) => ({
        scroll: element.scrollWidth,
        client: element.clientWidth,
        height: element.getBoundingClientRect().height,
      }));
      expect(value.scroll).toBeLessThanOrEqual(value.client);
      expect(value.height).toBeLessThan(25);
    });
  });
}

test.describe('the score parts at 400 (L8.5 ruling 4)', () => {
  test.use({ viewport: { width: 400, height: 844 } });

  test('keep "412 of 500" and the 88px label column from 400px', async ({ page }) => {
    await open(page);
    const gear = page.getByTestId('forever-score-gear');
    await expect(gear.getByText('412 of 500', { exact: true })).toBeVisible();
    await expect(gear.getByText('412/500', { exact: true })).toBeHidden();
    const label = await gear.evaluate((row) => row.children[0].getBoundingClientRect().width);
    expect(label).toBe(88);
  });
});

test('/setup says "type /fs copy, press Ctrl+C" (L8.5 ruling 5)', async ({ page }) => {
  await page.goto('/setup');
  await expect(
    page.getByText(
      'In game, type /fs copy, press Ctrl+C, then paste it into the planner’s Import from addon box.',
    ),
  ).toHaveCount(1);
  await expect(page.getByText('/fs export')).toHaveCount(0);
});
