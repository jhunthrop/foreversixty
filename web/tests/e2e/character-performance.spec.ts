// web/tests/e2e/character-performance.spec.ts
// Task L8.5: the Performance Rating card and the profile's leftovers.
//   Ruling 1. The rating card takes the Forever Score card's treatment: two columns from 1280,
//             the sparkline filling the right one, "Up 9 since 4 Dec" under it from the same
//             model, one line with the latest fight's rating and the count (fix round 1: no
//             claim about the Forever Score's Parses), and the
//             weakest part linking to its fight's row in "Every ranked fight".
//   Ruling 2. Experience notes what counts.
//   Ruling 3. The 404 card is capped at 720px; the guild link's focus ring stands 2px off and
//             the dots either side stay clear of it.
//   Ruling 4. Below 400px a part's row reads "412/500" with an 80px label column.
//   Ruling 5. /setup says "type /fs copy, press Ctrl+C".
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { CHARACTER, CHARACTER_PATH, CHARACTER_URL, routeOwner } from './support/character-fixture';

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
  // Fix round 1 item 3: Experience matches its own inputs: 14 dungeons at 25 and 7 raid bosses
  // at 40 make 630, and the five parts add up to the score.
  score: 1840,
  tier: 'epic',
  as_of: '2026-12-09',
  components: { gear: 412, experience: 630, parses: 540, attendance: 180, discipline: 78 },
  experience: { dungeons: 14, raid_kills: 7, alts: 2 },
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

  test('says the latest fight’s rating and how many are rated, with no claim about Parses (fix round 1 items 1 and 2)', async ({
    page,
  }) => {
    await open(page);
    // Fix round 2 item 1: it names the fight, from "Every ranked fight" and the trend.
    await expect(page.getByTestId('rating-card-verdict')).toHaveText(
      'Latest fight (Warden Kelthas, 9 Dec) rated 65 of 100 · 6 fights rated',
    );
    // Fix round 2 item 2: what a parse is next to what the rating is, muted under it.
    const note = page.getByTestId('rating-card-parse-vs-rating');
    await expect(note).toHaveText(
      'A parse is your rank against other players; the rating is how well the fight was played.',
    );
    await expect(note).toHaveCSS('color', 'rgb(154, 148, 132)');
    await expect(page.getByTestId('rating-card')).not.toContainText('counts toward');
    await expect(page.getByTestId('rating-card')).not.toContainText('Parses');
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
  // Fix round 1 item 3: the points are what the breakdown under them makes, 14 x 25 + 7 x 40.
  await expect(page.getByTestId('forever-score-experience-detail')).toHaveText(
    '14 dungeons · 7 raid bosses · 2 alts',
  );
  await expect(page.getByTestId('forever-score-experience')).toContainText('630 of 700');
  const note = page.getByTestId('forever-score-experience-note');
  // Fix round 2 items 4 and 6: "points" named, on one line.
  await expect(note).toHaveText('25 points a dungeon, 40 a raid boss, up to 700.');
  await expect(note).toBeVisible();
  await expect(note).toHaveCSS('font-size', '12px');
  await expect(note).toHaveCSS('color', 'rgb(154, 148, 132)');
  // Both boxes in one pass, once the web fonts have settled: two separate reads let a font
  // swap reflow the card between them and put the note "above" a detail line that had just
  // shrunk (one run in four on CI).
  const boxes = await page.evaluate(async () => {
    await document.fonts.ready;
    const rect = (testid: string) =>
      document.querySelector(`[data-testid="${testid}"]`)?.getBoundingClientRect() ?? null;
    const detail = rect('forever-score-experience-detail');
    const note = rect('forever-score-experience-note');
    return detail === null || note === null
      ? null
      : { detailBottom: detail.bottom, noteTop: note.top, noteHeight: note.height };
  });
  expect(boxes).not.toBeNull();
  expect(boxes!.noteTop).toBeGreaterThanOrEqual(boxes!.detailBottom - 0.5);
  expect(boxes!.noteHeight).toBeLessThan(20);
});

test.describe('the Experience note on a 390px phone (fix round 2 item 6)', () => {
  test.use({ viewport: { width: 390, height: 844 } });

  test('stays on one line', async ({ page }) => {
    await open(page);
    const box = await page.getByTestId('forever-score-experience-note').boundingBox();
    expect(box?.height ?? 99).toBeLessThan(20);
  });
});

test.describe('the 404 page at 1440 (L8.5 ruling 3)', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'a desktop width');

  test('caps the no-logs card at 720px', async ({ page }) => {
    await routeOwner(page);
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

// Fix round 1 item 4: each value is measured against its own grid cell (the bounding width of
// the text against the cell's), and 320 is measured too, where the track is about 32px.
for (const [width, track] of [
  [390, { min: 60 }],
  [360, { min: 60 }],
] as const) {
  test.describe(`the score parts at ${width} (L8.5 ruling 4)`, () => {
    test.use({ viewport: { width, height: 844 } });

    test('read "412/500" in an 80px label column, every value inside its own cell', async ({ page }) => {
      await open(page);
      const gear = page.getByTestId('forever-score-gear');
      await expect(gear.getByText('412/500', { exact: true })).toBeVisible();
      await expect(gear.getByText('412 of 500', { exact: true })).toBeHidden();
      await expect(page.getByTestId('forever-score-gain-cue')).toHaveText('+88 pts');
      const row = await gear.evaluate((element) => {
        const cells = [...element.children].slice(0, 4) as HTMLElement[];
        const box = (node: Node) => {
          const range = document.createRange();
          range.selectNodeContents(node);
          const r = range.getBoundingClientRect();
          return { left: r.left, right: r.right, width: r.width, height: r.height };
        };
        const value = [...cells[2].children].find((child) => getComputedStyle(child).display !== 'none')!;
        return {
          cells: cells.map((cell) => {
            const r = cell.getBoundingClientRect();
            return { left: r.left, right: r.right, width: r.width };
          }),
          value: box(value),
          cue: box(cells[3]),
          rowRight: element.getBoundingClientRect().right,
        };
      });
      expect(row.cells[0].width).toBe(80);
      expect(row.cells[1].width).toBeGreaterThanOrEqual(track.min);
      // "412/500" sits inside its 56px cell, on one line.
      expect(row.value.width).toBeLessThanOrEqual(row.cells[2].width + 0.5);
      expect(row.value.left).toBeGreaterThanOrEqual(row.cells[2].left - 0.5);
      expect(row.value.right).toBeLessThanOrEqual(row.cells[2].right + 0.5);
      expect(row.value.height).toBeLessThan(25);
      // "+88 pts" sits inside its 60px cell, on one line, inside the row.
      expect(row.cue.width).toBeLessThanOrEqual(row.cells[3].width + 0.5);
      expect(row.cue.left).toBeGreaterThanOrEqual(row.cells[3].left - 0.5);
      expect(row.cue.right).toBeLessThanOrEqual(row.cells[3].right + 0.5);
      expect(row.cue.height).toBeLessThan(25);
      expect(row.cells[3].right).toBeLessThanOrEqual(row.rowRight + 0.5);
      expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(width);
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

test.describe('the /ratings sample at 1440 (fix round 1 item 5)', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'a desktop width');

  test('stays one column in its 560px container, the sparkline under the line', async ({ page }) => {
    await page.goto('/ratings');
    const sample = page.getByTestId('ratings-rating-sample');
    await expect(sample.getByTestId('rating-card')).toBeVisible();
    const boxes = await sample.evaluate((element) => {
      const box = (testid: string) =>
        element.querySelector(`[data-testid="${testid}"]`)!.getBoundingClientRect();
      const main = box('rating-card-main');
      const side = box('rating-card-side');
      return { mainBottom: main.bottom, mainLeft: main.left, sideTop: side.top, sideLeft: side.left };
    });
    expect(boxes.sideTop).toBeGreaterThanOrEqual(boxes.mainBottom - 0.5);
    expect(Math.abs(boxes.sideLeft - boxes.mainLeft)).toBeLessThan(1);
  });
});

// Fix round 2 item 8: below 340px the cue moves under the Gear label and the track takes its
// 60px column back (about 100px at 320, was about 32).
test.describe('the score parts at 320 (fix round 2 item 8)', () => {
  test.use({ viewport: { width: 320, height: 568 } });

  test('put "+88 pts" under the Gear label and give the track its column', async ({ page }) => {
    await open(page);
    const gear = page.getByTestId('forever-score-gear');
    await expect(gear.getByText('412/500', { exact: true })).toBeVisible();
    const cue = page.getByTestId('forever-score-gain-cue');
    await expect(cue).toHaveText('+88 pts');
    const row = await gear.evaluate((element) => {
      const cells = [...element.children].slice(0, 4) as HTMLElement[];
      const r = (cell: Element) => {
        const b = cell.getBoundingClientRect();
        return { left: b.left, right: b.right, top: b.top, bottom: b.bottom, width: b.width };
      };
      const range = document.createRange();
      range.selectNodeContents(cells[3]);
      const text = range.getBoundingClientRect();
      return {
        label: r(cells[0]),
        track: r(cells[1]),
        value: r(cells[2]),
        cue: r(cells[3]),
        cueText: text.width,
      };
    });
    expect(row.label.width).toBe(80);
    expect(row.track.width).toBeGreaterThanOrEqual(90);
    expect(row.cue.top).toBeGreaterThanOrEqual(row.label.bottom - 0.5);
    expect(Math.abs(row.cue.left - row.label.left)).toBeLessThan(1);
    expect(row.cueText).toBeLessThanOrEqual(row.cue.width + 0.5);
    expect(row.value.right).toBeLessThanOrEqual(row.track.right + 8 + 56 + 0.5);
    // An unmarked row has no second line.
    const parses = await page.getByTestId('forever-score-parses').boundingBox();
    expect(parses?.height ?? 99).toBeLessThan(25);
    expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(320);
  });

  // Fix round 2 item 7: under 360px the line breaks at the dot.
  test('breaks the rating line at the dot: "6 fights rated" on its own line', async ({ page }) => {
    await open(page);
    const verdict = page.getByTestId('rating-card-verdict');
    const parts = await verdict.evaluate((element) =>
      [...element.querySelectorAll('span')].map((span) => {
        const b = span.getBoundingClientRect();
        return { text: span.textContent, top: b.top, display: getComputedStyle(span).display };
      }),
    );
    const rated = parts.find((part) => part.text === 'rated 65 of 100');
    const count = parts.find((part) => part.text === '6 fights rated');
    const dot = parts.find((part) => part.text?.trim() === '·');
    expect(dot?.display).toBe('none');
    expect(count?.display).toBe('block');
    expect(count?.top ?? 0).toBeGreaterThan(rated?.top ?? 0);
  });
});

// Fix round 2 items 3 and 5: the latest rated fight's row carries "rated 65" under its parse;
// as the address's target it takes a 2px gold left edge and a gold tint, 96px under the top.
for (const width of [1440, 390]) {
  test.describe(`the weakest part’s row at ${width} (fix round 2 items 3 and 5)`, () => {
    test.use({ viewport: { width, height: width === 1440 ? 900 : 844 } });

    test('carries "rated 65" and, reached by Enter on the link, a gold left edge and tint', async ({
      page,
    }) => {
      await open(page);
      const row = page.getByTestId('character-history').locator('#fight-fixture2abcd-3');
      await expect(row.getByTestId('character-fight-rated')).toHaveText('rated 65');
      await expect(page.getByTestId('character-history').getByTestId('character-fight-rated')).toHaveCount(1);
      const link = page.getByTestId('rating-card-weakest').getByRole('link', { name: 'Activity' });
      await link.focus();
      await page.keyboard.press('Enter');
      await expect(page).toHaveURL(/#fight-fixture2abcd-3$/);
      // Tailwind's shadow stack draws the edge as the last layer, the others transparent.
      await expect(row).toHaveCSS('box-shadow', /(^|, )rgb\(229, 185, 85\) 2px 0px 0px 0px inset$/);
      await expect(row).toHaveCSS('background-color', 'rgba(229, 185, 85, 0.08)');
      await expect(row).toHaveCSS('scroll-margin-top', '96px');
      await expect(row).toBeInViewport();
      const top = await row.evaluate((element) => element.getBoundingClientRect().top);
      expect(top).toBeGreaterThanOrEqual(95);
      // The other row is not marked: no gold in its shadow, no tint.
      const other = page.getByTestId('character-history').locator('#fight-otherreport1-1');
      const otherShadow = await other.evaluate((element) => getComputedStyle(element).boxShadow);
      expect(otherShadow).not.toContain('229, 185, 85');
      await expect(other).toHaveCSS('background-color', 'rgba(0, 0, 0, 0)');
    });
  });
}
