// web/tests/e2e/character-gear.spec.ts
// Task L8.6: the profile's Gear section, from the character answer's `gear` and
// `gear_gap_total` (api/openapi.yaml `CharacterGear`), stubbed with every row state
// (src/fixtures/character-gear.ts) and the build's items/priest.json reduced to the fixture's
// ids and qualities.
//   Ruling 1. Under the score card: 17 rows in the export's order, names in quality colours,
//             the title "Gear · +70 pts across 7 slots"; `gear: []` is one line, no section.
//   Ruling 2. "Most to gain: Gear" lands on the section; the BiS page link is its footer's.
//   Ruling 3. The card's "88 below its top" and the section's "+70 pts across 7 slots" stand
//             apart, with one muted sentence on the difference.
//   Ruling 4. On a phone two lines a row, 44px rows, no sideways scroll.
import { expect, test, type Page } from '@playwright/test';
import { CHARACTER, CHARACTER_PATH, CHARACTER_URL, GEAR, GEAR_ITEMS } from './support/character-fixture';

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

/** The score whose "Most to gain" is Gear, 88 below its top. */
const SCORE = {
  score: 1830,
  tier: 'epic',
  as_of: '2026-11-04',
  components: { gear: 412, experience: 620, parses: 540, attendance: 180, discipline: 78 },
  experience: { dungeons: 14, raid_kills: 0, alts: 2 },
};

const RGB = {
  uncommon: 'rgb(30, 255, 0)',
  rare: 'rgb(61, 148, 240)',
  epic: 'rgb(184, 102, 245)',
  muted: 'rgb(154, 148, 132)',
};

async function open(page: Page, gear: unknown = GEAR): Promise<void> {
  await page.route(CHARACTER_URL, (route) =>
    route.fulfill(ok({ ...CHARACTER.data, ...SCORE, ...(gear as object) })),
  );
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(NOT_FOUND));
  await page.route('**/data/*/items/priest.json', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(GEAR_ITEMS) }),
  );
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('forever-score-card')).toBeVisible();
}

const row = (page: Page, slot: string) => page.getByTestId(`character-gear-row-${slot}`);

/** Every row's cells, read in one pass so a settling layout cannot split them. */
async function rowBoxes(page: Page) {
  return page.evaluate(() =>
    [...document.querySelectorAll('[data-testid^="character-gear-row-"]')].map((li) => {
      const box = (testid: string) => {
        const el = li.querySelector(`[data-testid="${testid}"]`);
        if (el === null) return null;
        const r = el.getBoundingClientRect();
        return { top: r.top, bottom: r.bottom, left: r.left, right: r.right, height: r.height };
      };
      const r = li.getBoundingClientRect();
      return {
        slot: li.getAttribute('data-testid')?.replace('character-gear-row-', '') ?? '',
        height: r.height,
        right: r.right,
        items: box('character-gear-items'),
        meta: box('character-gear-meta'),
        gap: box('character-gear-gap'),
      };
    }),
  );
}

for (const width of [1440, 390]) {
  test.describe(`the Gear section at ${width} (L8.6)`, () => {
    test.use({ viewport: { width, height: width === 1440 ? 900 : 844 } });

    test('sits under the score card, titled from gear_gap_total and the count of gaps (ruling 1)', async ({
      page,
    }) => {
      await open(page);
      const section = page.getByTestId('character-gear');
      await expect(page.getByTestId('character-gear-title')).toHaveText('Gear · +70 pts across 7 slots');
      // Nothing between the card and the section: the section is the next thing down, one
      // section gap (22px on a phone, 32px from md) under it.
      const next = await page.evaluate(() => {
        const card = document.querySelector('[data-testid="forever-score"]')!.getBoundingClientRect();
        const gear = document.querySelector('[data-testid="character-gear"]')!.getBoundingClientRect();
        return gear.top - card.bottom;
      });
      expect(next).toBeGreaterThan(0);
      expect(next).toBeLessThanOrEqual(33);
      await expect(section.locator('[data-testid^="character-gear-row-"]')).toHaveCount(17);
    });

    test('draws every row state in quality colours (ruling 1)', async ({ page }) => {
      await open(page);
      // An upgrade: uncommon worn, epic pick, the source and chance, the gap.
      const head = row(page, 'head');
      await expect(head.getByTestId('character-gear-slot')).toHaveText('Head');
      await expect(head.getByTestId('character-gear-worn')).toHaveText('Green Lens');
      await expect(head.getByTestId('character-gear-worn')).toHaveCSS('color', RGB.uncommon);
      await expect(head.getByTestId('character-gear-pick')).toHaveText('Halo of Transcendence');
      await expect(head.getByTestId('character-gear-pick')).toHaveCSS('color', RGB.epic);
      await expect(head.getByTestId('character-gear-meta')).toContainText('Onyxia’s Lair: Onyxia · 17%');
      await expect(head.getByTestId('character-gear-gap')).toHaveText('+18.4 pts');
      // A chance under 1%, and a source with no rate.
      await expect(row(page, 'back').getByTestId('character-gear-meta')).toContainText(
        'Dire Maul: Tendris Warpwood · <1%',
      );
      await expect(row(page, 'chest').getByTestId('character-gear-meta')).toContainText('Tailoring');
      await expect(row(page, 'chest').getByTestId('character-gear-meta')).not.toContainText('%');
      // The worn item is the pick: its name in its colour, "Best for your band" muted, no arrow.
      const shoulder = row(page, 'shoulder');
      await expect(shoulder.getByTestId('character-gear-worn')).toHaveCSS('color', RGB.epic);
      await expect(shoulder.getByTestId('character-gear-meta')).toContainText('Best for your band');
      await expect(shoulder.getByTestId('character-gear-meta')).toHaveCSS('color', RGB.muted);
      await expect(shoulder.getByTestId('character-gear-pick')).toHaveCount(0);
      await expect(shoulder.getByTestId('character-gear-gap')).toHaveText('');
      // The rings worn the other way round read the same.
      await expect(row(page, 'finger1').getByTestId('character-gear-meta')).toContainText(
        'Best for your band',
      );
      // An empty slot: "Empty" muted, with the pick.
      const neck = row(page, 'neck');
      await expect(neck.getByTestId('character-gear-empty')).toHaveText('Empty');
      await expect(neck.getByTestId('character-gear-empty')).toHaveCSS('color', RGB.muted);
      await expect(neck.getByTestId('character-gear-pick')).toHaveText('Choker of the Fire Lord');
      await expect(neck.getByTestId('character-gear-gap')).toHaveText('+12.6 pts');
      // An unknown worn item and an unscored gap.
      const legs = row(page, 'legs');
      await expect(legs.getByTestId('character-gear-worn')).toHaveText('Unknown item');
      await expect(legs.getByTestId('character-gear-worn')).toHaveCSS('color', RGB.muted);
      await expect(legs.getByTestId('character-gear-gap')).toHaveText('–');
      await expect(legs.getByTestId('character-gear-gap')).toHaveAttribute('aria-label', 'Not scored');
      // No pick: only the worn item; nothing at all: only "Empty".
      await expect(row(page, 'ranged').getByTestId('character-gear-worn')).toHaveText('Skul’s Ghastly Touch');
      await expect(row(page, 'ranged').getByTestId('character-gear-worn')).toHaveCSS('color', RGB.rare);
      await expect(row(page, 'ranged').getByTestId('character-gear-pick')).toHaveCount(0);
      await expect(row(page, 'ranged').getByTestId('character-gear-meta')).toHaveCount(0);
      await expect(row(page, 'trinket2').getByTestId('character-gear-empty')).toHaveText('Empty');
      await expect(row(page, 'trinket2').getByTestId('character-gear-pick')).toHaveCount(0);
    });

    test('right-aligns the gaps in one mono column, every row at least 44px (rulings 1, 4)', async ({
      page,
    }) => {
      await open(page);
      await expect(row(page, 'head').getByTestId('character-gear-gap')).toHaveCSS(
        'font-family',
        /JetBrains Mono/,
      );
      const rows = await rowBoxes(page);
      expect(rows).toHaveLength(17);
      for (const r of rows) expect(r.height, r.slot).toBeGreaterThanOrEqual(44);
      const rights = rows.flatMap((r) => (r.gap === null || r.gap.right === r.gap.left ? [] : [r.gap.right]));
      expect(rights.length).toBeGreaterThanOrEqual(9);
      for (const right of rights) expect(Math.abs(right - rights[0])).toBeLessThan(1);
    });

    test('lays a row out as the width asks, with no sideways scroll (ruling 4)', async ({ page }) => {
      await open(page);
      const rows = await rowBoxes(page);
      const head = rows.find((r) => r.slot === 'head')!;
      if (width === 390) {
        // Two lines: worn -> pick, then source · chance with the gap at its end.
        expect(head.meta!.top).toBeGreaterThanOrEqual(head.items!.bottom - 1);
        expect(Math.abs(head.gap!.bottom - head.meta!.bottom)).toBeLessThan(4);
      } else {
        // One line: the source follows the pick, the gap at the row's end.
        expect(Math.abs(head.meta!.top - head.items!.top)).toBeLessThan(4);
        expect(head.meta!.left).toBeGreaterThan(head.items!.right - 1);
        await expect(row(page, 'head').getByTestId('character-gear-meta')).toContainText(
          '· Onyxia’s Lair: Onyxia · 17%',
        );
      }
      const scroll = await page.evaluate(() => ({
        width: document.documentElement.scrollWidth,
        client: document.documentElement.clientWidth,
      }));
      expect(scroll.width).toBeLessThanOrEqual(scroll.client);
    });

    test('"Most to gain: Gear" lands on the section; the BiS page is in its footer (ruling 2)', async ({
      page,
    }) => {
      await open(page);
      const link = page.getByTestId('forever-score-most-to-gain').getByRole('link', { name: 'Gear' });
      await expect(link).toHaveAttribute('href', '#gear');
      await link.click();
      await expect(page).toHaveURL(/#gear$/);
      const section = page.getByTestId('character-gear');
      await expect(section).toBeInViewport();
      await expect(section).toHaveCSS('scroll-margin-top', '96px');
      const top = await section.evaluate((element) => element.getBoundingClientRect().top);
      expect(top).toBeGreaterThanOrEqual(95);
      expect(top).toBeLessThan(110);
      const bis = page.getByTestId('character-gear-bis');
      await expect(bis).toHaveText('Every pick for Discipline at 60 →');
      await expect(bis).toHaveAttribute('href', '/bis/priest/discipline?faction=alliance#band-alliance-60');
    });

    test('the card’s "88 below its top" and the section’s "+70 pts across 7 slots" stand apart (ruling 3)', async ({
      page,
    }) => {
      await open(page);
      await expect(page.getByTestId('forever-score-most-to-gain')).toHaveText(
        'Most to gain: Gear, 88 below its top.',
      );
      await expect(page.getByTestId('character-gear-title')).toHaveText('Gear · +70 pts across 7 slots');
      const difference = page.getByTestId('character-gear-difference');
      await expect(difference).toHaveText(
        'Below its top measures against the band’s best set; the slot gaps measure against the picks you can wear.',
      );
      await expect(difference).toHaveCSS('color', RGB.muted);
    });

    test('`gear: []`: one line in place of the section, and the Gear link lands on it (ruling 1)', async ({
      page,
    }) => {
      await open(page, { gear: [], gear_gap_total: null });
      await expect(page.getByTestId('character-gear')).toHaveCount(0);
      const line = page.getByTestId('character-gear-none');
      await expect(line).toHaveText('No gear on file. Upload a log or sync the addon.');
      await page.getByTestId('forever-score-most-to-gain').getByRole('link', { name: 'Gear' }).click();
      await expect(page).toHaveURL(/#gear$/);
      await expect(line).toBeInViewport();
    });
  });
}
