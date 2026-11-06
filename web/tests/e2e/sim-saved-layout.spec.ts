// web/tests/e2e/sim-saved-layout.spec.ts
// The saved result's layout after its lazy sections load (spec 2026-11-04-sim-results 3.7, 4.2,
// 8; build reviews 2026-11-05; the open Sections of spec 2026-11-06-sim §4.5): every lazy holder
// is its section's own height once the section renders or fails, the timeline's loading head is its ready head at its ready height, and on a phone no
// lane name sits over the plot, so every event at the pull can be seen.
import type { Page, Route } from '@playwright/test';
import { expect, test } from './support/test';
import { gotoSavedSim } from './support/sim-page';

const FIXTURES = ['simfullfixt1', 'simfixtureab'] as const;
const WIDTHS = [390, 1440] as const;
/** The timeline's lazy chunk is a hashed `sim-island-*.js`: it is the one that draws the toolbar. */
const CHUNK = '**/sim-island-*.js';
const TIMELINE_MARK = 'timeline-toolbar';

/** Scrolls the page top to bottom so every card nears the viewport and loads, then settles. */
async function loadEveryCard(page: Page): Promise<void> {
  const height = await page.evaluate(() => document.documentElement.scrollHeight);
  for (let y = 0; y <= height + 1200; y += 500) {
    await page.evaluate((top) => window.scrollTo(0, top), y);
    await page.waitForTimeout(30);
  }
  await expect(page.locator('[data-testid$="-skeleton"]')).toHaveCount(0);
}

/** The lazy holders (LazySection's `sim-section-<id>`), not their skeleton or error inside. */
const HOLDERS =
  'div[data-testid^="sim-section-"]:not([data-testid$="-skeleton"]):not([data-testid$="-error"])';

/** Each lazy holder's box against its section's, and what is left of its reserve. */
function sectionBoxes(page: Page): Promise<{ id: string; box: number; card: number; minHeight: string }[]> {
  return page.evaluate(
    (selector) =>
      [...document.querySelectorAll<HTMLElement>(selector)].map((holder) => {
        const card = holder.firstElementChild?.matches('.reveal')
          ? holder.firstElementChild.firstElementChild
          : holder.firstElementChild;
        return {
          id: (holder.dataset.testid ?? '').replace('sim-section-', ''),
          box: holder.getBoundingClientRect().height,
          card: card?.getBoundingClientRect().height ?? 0,
          minHeight: getComputedStyle(holder).minHeight,
        };
      }),
    HOLDERS,
  );
}

/** Serves every chunk, the timeline's only once `decide` says so: or fails it. */
async function routeTimelineChunk(page: Page, decide: () => Promise<'serve' | 'fail'>): Promise<void> {
  await page.route(CHUNK, async (route: Route) => {
    const response = await route.fetch();
    const body = await response.text();
    if (body.includes(TIMELINE_MARK) && (await decide()) === 'fail') await route.abort();
    else await route.fulfill({ response, body });
  });
}

// A chunk still being fetched through the route when a test ends must not fail it.
test.afterEach(({ page }) => page.unrouteAll({ behavior: 'ignoreErrors' }));

for (const id of FIXTURES) {
  for (const width of WIDTHS) {
    test(`${id} at ${width}: every lazy section is its card's height after load`, async ({ page }) => {
      await page.setViewportSize({ width, height: 900 });
      await page.goto(`/sim/${id}`);
      await expect(page.getByTestId('sim-saved')).toBeVisible();
      await loadEveryCard(page);

      const boxes = await sectionBoxes(page);
      expect(boxes.map((row) => row.id)).toEqual(expect.arrayContaining(['talents', 'buffs', 'settings']));
      for (const row of boxes) {
        expect(
          Math.abs(row.box - row.card),
          `${row.id}: section ${row.box} vs card ${row.card}`,
        ).toBeLessThanOrEqual(2);
        expect(row.minHeight, `${row.id} keeps no reserve`).toMatch(/^(0px|auto)$/);
      }
    });
  }
}

for (const width of [360, 390]) {
  test(`at ${width} every event at the pull shows, and no lane name covers the plot`, async ({ page }) => {
    await page.setViewportSize({ width, height: 800 });
    await gotoSavedSim(page, '/sim/simfullfixt1');
    await page.getByTestId('sim-section-timeline').scrollIntoViewIfNeeded();
    const plot = page.getByTestId('timeline-plot');
    await expect(plot.locator('g[data-row] rect').first()).toBeAttached();
    // The plot at its start, as a player first sees it: 0:00 just right of the icon column.
    await page.getByTestId('timeline-scroller').evaluate((element) => element.scrollTo({ left: 0 }));

    const report = await plot.evaluate((element) => {
      const plotBox = element.getBoundingClientRect();
      const pull = plotBox.left + 10; // 0:00 sits 2.5 s after the plot's start at 4 px a second.
      const all = [...element.querySelectorAll('g[data-row] rect')];
      const atPull = all.filter((rect) => {
        const box = rect.getBoundingClientRect();
        return box.left <= pull + 2 && box.right >= pull - 2;
      });
      // Each mark at the pull, brought on screen, is what the page shows at its pull end.
      const hidden = atPull.filter((rect) => {
        // Vertically only: the plot stays scrolled to its start, the pull at its left edge.
        window.scrollBy(0, rect.getBoundingClientRect().top - window.innerHeight / 2);
        const box = rect.getBoundingClientRect();
        const hit = document.elementFromPoint(
          box.left + Math.min(2, box.width / 2),
          box.top + box.height / 2,
        );
        return hit !== rect;
      });
      const rects = all.map((rect) => rect.getBoundingClientRect());
      const chips = [...element.querySelectorAll('[data-testid="timeline-lane-chip"]')].map((chip) =>
        chip.getBoundingClientRect(),
      );
      const overlaps = chips.filter((chip) =>
        rects.some(
          (box) =>
            box.left < chip.right && chip.left < box.right && box.top < chip.bottom && chip.top < box.bottom,
        ),
      );
      return { atPull: atPull.length, hidden: hidden.length, chips: chips.length, overlaps: overlaps.length };
    });

    // Berserking, Frozen Heart, Mighty Rage Potion and Bloodthirst presses and bars at 0:00, at least.
    expect(report.atPull).toBeGreaterThanOrEqual(8);
    expect(report.hidden).toBe(0);
    expect(report.chips).toBeGreaterThan(0);
    expect(report.overlaps).toBe(0);

    // The whole fight scrolls into view: the plot's end reaches the scroller's right edge.
    const scroller = page.getByTestId('timeline-scroller');
    await scroller.evaluate((element) => element.scrollTo({ left: element.scrollWidth }));
    const ends = await page.evaluate(() => {
      const box = document.querySelector('[data-testid="timeline-scroller"]')!.getBoundingClientRect();
      const plotBox = document.querySelector('[data-testid="timeline-plot"]')!.getBoundingClientRect();
      return { plotRight: plotBox.right, scrollerRight: box.right };
    });
    expect(Math.abs(ends.plotRight - ends.scrollerRight)).toBeLessThanOrEqual(1);
  });
}

test('the timeline loads under its ready head at its ready height', async ({ page, isMobile }) => {
  // A coarse pointer at 1440 grows the zoom buttons to 44 px: a touch screen this wide is a
  // tablet whose card mounts below the fold, where 15 px moves nothing on screen.
  test.skip(isMobile, 'the wide reserve is measured with a fine pointer');
  await page.setViewportSize({ width: 1440, height: 900 });
  let release = (): void => {};
  const gate = new Promise<void>((resolve) => (release = resolve));
  await routeTimelineChunk(page, async () => {
    await gate;
    return 'serve';
  });
  await gotoSavedSim(page, '/sim/simfullfixt1');
  const section = page.getByTestId('sim-section-timeline');
  await section.scrollIntoViewIfNeeded();
  await expect(page.getByTestId('sim-section-timeline-skeleton')).toBeVisible();

  const head = (): Promise<{ text: string; glyph: string; height: number }> =>
    section.evaluate((element) => {
      const card = element.querySelector('section')!;
      const top = card.firstElementChild as HTMLElement;
      return {
        text: top.innerText.replace(/\s+/g, ' ').trim(),
        glyph: top.querySelector('svg')?.innerHTML ?? '',
        height: element.getBoundingClientRect().height,
      };
    });
  const loading = await head();
  release();
  await expect(page.getByTestId('timeline-toolbar')).toBeVisible();
  const ready = await head();

  expect(loading.text).toBe(ready.text);
  expect(loading.text).toContain('The median run · 2:53 long');
  expect(loading.glyph).toBe(ready.glyph);
  expect(Math.abs(loading.height - ready.height)).toBeLessThanOrEqual(2);
});

test('a timeline chunk that fails shows the compact error, with no reserve left behind', async ({ page }) => {
  await page.setViewportSize({ width: 1440, height: 900 });
  await routeTimelineChunk(page, () => Promise.resolve('fail'));
  await gotoSavedSim(page, '/sim/simfullfixt1');
  await page.getByTestId('sim-section-timeline').scrollIntoViewIfNeeded();
  await expect(page.getByTestId('sim-section-timeline-error')).toBeVisible();

  const [timeline] = (await sectionBoxes(page)).filter((row) => row.id === 'timeline');
  expect(timeline!.box).toBeLessThan(300);
  expect(Math.abs(timeline!.box - timeline!.card)).toBeLessThanOrEqual(2);
});
