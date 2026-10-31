// web/tests/captures/paperdoll/capture.spec.ts
// Acceptance screenshots for the paperdoll (design/specs/2026-10-31-paperdoll.md §9). Not a test
// of behaviour: it draws each call site with the real build's items and icons, the API stubbed,
// at the widths the spec names, and writes 2x PNGs (the mock boards' scale) plus same-scale
// side-by-sides with the boards to design/reviews/2026-10-31-paperdoll/build/.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page } from '@playwright/test';
import { test } from '../../e2e/support/test';
import { DEAD_API_BASE } from '../../hermetic';
import { CHARACTER, CHARACTER_PATH, CHARACTER_URL } from '../../e2e/support/character-fixture';
import { cropOf, settle, sideBySide, sizeTo, targetIn } from '../support/capture-kit';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const WEB = path.resolve(HERE, '../../..');
const BOARDS = path.resolve(WEB, '../design/mocks/2026-10-31-paperdoll');
const out = targetIn('2026-10-31-paperdoll/build');
const BUILD = (
  JSON.parse(readFileSync(path.join(WEB, 'src/data/active-build.json'), 'utf8')) as { build: string }
).build;
const ITEMS = JSON.parse(
  readFileSync(path.join(WEB, `../data/builds/${BUILD}/items/warrior.json`), 'utf8'),
) as {
  items: { id: number; name: string; item_level: number }[];
};
const byId = new Map(ITEMS.items.map((item) => [item.id, item]));
const API = DEAD_API_BASE;
const WIDTHS = [360, 390, 1024, 1280, 1440, 2000] as const;

/** The raid board's launch-legal Fury set (raid-*.png). */
const RAID: Record<string, number> = {
  head: 12640,
  neck: 12103,
  shoulder: 12927,
  back: 13340,
  chest: 252483,
  wrist: 12936,
  hands: 19049,
  waist: 22232,
  legs: 18736,
  feet: 13967,
  finger1: 18500,
  finger2: 13098,
  trinket1: 249469,
  trinket2: 18537,
  main_hand: 279261,
  off_hand: 279259,
  ranged: 279273,
};

/** The worst-case board's longest launch-legal names (longest-*.png). */
const LONGEST: Record<string, number> = {
  head: 9492,
  neck: 13091,
  shoulder: 18742,
  back: 11812,
  chest: 250520,
  wrist: 17714,
  hands: 10808,
  waist: 252476,
  legs: 250525,
  feet: 22240,
  finger1: 6321,
  finger2: 273030,
  trinket1: 249473,
  trinket2: 11811,
  main_hand: 810,
  off_hand: 7726,
  ranged: 279273,
};

/** The beta board (beta-*.png): worn, pick and gain per slot; null worn name = not in our data. */
const BETA: [slot: string, worn: number | null, unknown: boolean, pick: number | null, gap: number | null][] =
  [
    ['head', 11925, false, 12640, 60],
    ['neck', 12103, false, 12103, 0],
    ['shoulder', 22241, false, 12927, 24.2],
    ['back', 11626, false, 13340, 3.9],
    ['chest', 288892, true, 252483, null],
    ['wrist', 13120, false, 12936, 18.4],
    ['hands', 15063, false, 19049, 4.4],
    ['waist', 288887, true, 22232, null],
    ['legs', 15065, false, 18736, 1.5],
    ['feet', 12555, false, 13967, 11.3],
    ['finger1', 2246, false, 18500, 16.1],
    ['finger2', 18302, false, 13098, 9.1],
    ['trinket1', 249469, false, 249469, 0],
    ['trinket2', 288893, true, 18537, null],
    ['main_hand', 12790, false, null, null],
    ['off_hand', null, false, null, null],
    ['ranged', null, false, null, null],
  ];

function fs1(gear: Record<string, number>, who: string): string {
  const slots = Object.entries(gear)
    .map(([slot, id]) => `${slot}=${id}`)
    .join(',');
  return `FS1:${BUILD}:warrior:human:34320003002/05153105022011501/2:${slots}|level=60|who=${who}:Doomhowl`;
}

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});
const NOT_FOUND = {
  status: 404,
  contentType: 'application/json',
  body: '{"ok":false,"data":null,"error":"not_found"}',
};

async function atRest(page: Page): Promise<void> {
  await page.evaluate(() => (document.activeElement as HTMLElement | null)?.blur());
  await page.mouse.move(0, 0);
}

async function framed(page: Page, testid: string): Promise<void> {
  await page.getByTestId(testid).waitFor();
  await page.locator('[data-testid^="paperdoll-slot-head"]:not([data-state="loading"])').first().waitFor();
  await page.evaluate(async () => {
    await document.fonts.ready;
    await Promise.all(
      [...document.images].map((img) => {
        img.loading = 'eager';
        return img.complete
          ? null
          : new Promise((resolve) => img.addEventListener('load', resolve, { once: true }));
      }),
    );
  });
  await settle(page);
}

async function openSim(page: Page, width: number, gear: Record<string, number>, who: string): Promise<void> {
  await sizeTo(page, width);
  await page.route(`${API}/**`, (route) => route.fulfill(NOT_FOUND));
  await page.goto('/sim');
  await page.getByTestId('sim-cold-paste-input').fill(fs1(gear, who));
  await page.getByTestId('sim-cold-paste-run').click();
  await framed(page, 'sim-character');
  await atRest(page);
}

function betaGear() {
  return BETA.map(([slot, worn, unknown, pick, gap]) => ({
    slot,
    worn:
      worn === null
        ? null
        : unknown
          ? { id: worn, name: null, item_level: null }
          : {
              id: worn,
              name: byId.get(worn)?.name ?? `Item ${worn}`,
              item_level: byId.get(worn)?.item_level ?? null,
            },
    pick:
      pick === null
        ? null
        : { id: pick, name: byId.get(pick)?.name ?? '', source: 'Stub source', chance: null },
    gap,
  }));
}

interface CharacterOptions {
  itemsDelayMs?: number;
  itemsFail?: boolean;
}

async function openCharacter(page: Page, width: number, options: CharacterOptions = {}): Promise<void> {
  await sizeTo(page, width);
  await page.route(`${API}/**`, (route) => route.fulfill(NOT_FOUND));
  const data = {
    ...CHARACTER.data,
    character: { ...CHARACTER.data.character, name: 'Brakka', class: 'Warrior', race: 'Human' },
    best: CHARACTER.data.best.map((row) => ({ ...row, spec: 'Arms', metric: 'dps' })),
    gear: betaGear(),
    gear_gap_total: 148.8,
    gear_spec: 'warrior-arms',
    gear_spec_source: 'fight',
  };
  await page.route(CHARACTER_URL, (route) => route.fulfill(envelope(data)));
  if (options.itemsDelayMs !== undefined || options.itemsFail === true)
    await page.route('**/data/*/items/warrior.json', async (route) => {
      if (options.itemsFail === true) return route.fulfill({ status: 500, body: '' });
      await new Promise((resolve) => setTimeout(resolve, options.itemsDelayMs));
      return route.continue();
    });
  // A real-data build has no fixture character pages: serve the page shell the deploy rewrites
  // every /character/... address to.
  await page.route(`**${CHARACTER_PATH}`, async (route) =>
    route.fulfill({
      response: await route.fetch({ url: new URL('/character', route.request().url()).href }),
    }),
  );
  await page.goto(CHARACTER_PATH);
  await page.getByTestId('character-gear').waitFor();
  if (options.itemsDelayMs === undefined) await framed(page, 'character-gear');
  await atRest(page);
}

const board = (name: string) => path.join(BOARDS, `${name}.png`);
const exists = (file: string) => {
  try {
    readFileSync(file);
    return true;
  } catch {
    return false;
  }
};

for (const width of WIDTHS) {
  test(`sim raid @${width}`, async ({ page }) => {
    await openSim(page, width, RAID, 'Brakka');
    const file = out(`sim-raid-${width}.png`);
    await cropOf(page, page.getByTestId('sim-character'), file, 4000);
    if (exists(board(`raid-${width}`)))
      await sideBySide(page, board(`raid-${width}`), file, out(`side-raid-${width}.png`), width);
  });

  test(`character beta @${width}`, async ({ page }) => {
    await openCharacter(page, width);
    const file = out(`character-beta-${width}.png`);
    await cropOf(page, page.getByTestId('character-gear'), file, 4000);
    if (exists(board(`beta-${width}`)))
      await sideBySide(page, board(`beta-${width}`), file, out(`side-beta-${width}.png`), width);
  });
}

/** The report needs a fixture report, which only the fixture build carries: run it against that
 *  build (FIXTURE_PORT, default 4483), with the fixture's own warrior items and two ids the
 *  data does not have. */
const FIXTURE_ORIGIN = `http://localhost:${process.env.FIXTURE_PORT ?? '4483'}`;
const FIXTURE_LEVELS = new Map(
  (
    JSON.parse(readFileSync(path.join(WEB, 'src/fixtures/planner/items/warrior.json'), 'utf8')) as {
      items: { id: number; item_level: number }[];
    }
  ).items.map((item) => [item.id, item.item_level]),
);
const AT_PULL: Record<string, number> = {
  head: 12640,
  shoulder: 16966,
  chest: 288892,
  finger1: 19325,
  trinket1: 13968,
  trinket2: 21550,
  main_hand: 12784,
  ranged: 288893,
};

test.describe('report at pull', () => {
  test.use({ baseURL: FIXTURE_ORIGIN });

  for (const width of WIDTHS) {
    test(`report at pull @${width}`, async ({ page }) => {
      await sizeTo(page, width);
      await page.route(`${API}/**`, (route) => route.fulfill(NOT_FOUND));
      await page.route('**/fights/3/summary.json*', async (route) => {
        const response = await route.fetch();
        const summary = (await response.json()) as { combatants: { gear: unknown[] }[] };
        const slots = [
          'head',
          'neck',
          'shoulder',
          'shirt',
          'chest',
          'waist',
          'legs',
          'feet',
          'wrist',
          'hands',
        ];
        const tail = [
          'finger1',
          'finger2',
          'trinket1',
          'trinket2',
          'back',
          'main_hand',
          'off_hand',
          'ranged',
          'tabard',
        ];
        // The log's own item level for each piece is the item's (the fixture rows'); the two ids
        // the data lacks take a plausible 65. The row's header figure is the log's mean of them.
        const level = (id: number) => FIXTURE_LEVELS.get(id) ?? 65;
        const worn = Object.values(AT_PULL);
        const first = summary.combatants[0]! as { gear: unknown[]; item_level?: number };
        first.item_level = Math.round(worn.reduce((sum, id) => sum + level(id), 0) / worn.length);
        first.gear = [...slots, ...tail].map((slot, index) => {
          const id = AT_PULL[slot] ?? 0;
          return {
            ID: id,
            ItemLevel: id > 0 ? level(id) : 0,
            Enchants: index % 3 === 0 && id > 0 ? [1900] : null,
            BonusIDs: null,
            Gems: null,
          };
        });
        await route.fulfill({ response, json: summary });
      });
      await page.goto('/reports/fixture2abcd?fight=3');
      await page.getByTestId('combatant-gear').first().click();
      await framed(page, 'gear-list');
      await atRest(page);
      await cropOf(
        page,
        page.getByTestId('combatants').getByRole('listitem').first(),
        out(`report-${width}.png`),
        4000,
      );
      if (width === 1280) {
        // The fixture rows carry `bonding` and `armor_type`, which the real build gains from the
        // next nightly: the tooltip's "Binds when equipped" and "Plate" lines.
        await page.getByTestId('gear-list').getByTestId('paperdoll-item-head').hover();
        await page.getByTestId('item-tooltip').waitFor();
        await settle(page);
        await cropOf(page, page.getByTestId('gear-list'), out('tip-bind-1280.png'), 4000);
      }
    });

    test(`saved sim @${width}`, async ({ page }) => {
      await sizeTo(page, width);
      await page.goto('/sim/simfixtureab');
      await framed(page, 'sim-character');
      await atRest(page);
      await cropOf(page, page.getByTestId('sim-character'), out(`sim-saved-${width}.png`), 4000);
    });
  }
});

test('focus-visible @1280', async ({ page }) => {
  await openCharacter(page, 1280);
  await page.getByTestId('paperdoll-item-head').focus();
  await page.keyboard.press('Tab');
  await page.keyboard.press('Tab');
  await page.getByTestId('item-tooltip').waitFor();
  await settle(page);
  await cropOf(page, page.getByTestId('character-gear'), out('focus-1280.png'), 1400);
});

for (const width of [360, 390, 1024]) {
  test(`longest names @${width}`, async ({ page }) => {
    await openSim(page, width, LONGEST, 'Brakka');
    const file = out(`longest-${width}.png`);
    await cropOf(page, page.getByTestId('sim-character'), file, 4000);
    if (exists(board(`longest-${width}`)))
      await sideBySide(page, board(`longest-${width}`), file, out(`side-longest-${width}.png`), width);
  });
}

for (const width of [390, 1280]) {
  test(`loading and failed @${width}`, async ({ page }) => {
    await openCharacter(page, width, { itemsDelayMs: 60_000 });
    await page.getByTestId('paperdoll-slot-head').waitFor();
    await settle(page);
    await cropOf(page, page.getByTestId('character-gear'), out(`loading-${width}.png`), 4000);
    const failed = await page.context().newPage();
    await openCharacter(failed, width, { itemsFail: true });
    await cropOf(failed, failed.getByTestId('character-gear'), out(`failed-${width}.png`), 4000);
  });
}

test('tooltips @1280', async ({ page }) => {
  await openCharacter(page, 1280);
  const shoot = async (testid: string, file: string) => {
    await page.getByTestId(testid).scrollIntoViewIfNeeded();
    await page.getByTestId(testid).hover();
    await page.getByTestId('item-tooltip').waitFor();
    await page.evaluate(() => document.fonts.ready);
    await settle(page);
    await page.screenshot({ path: out(file) });
    await page.mouse.move(0, 0);
  };
  await page.evaluate(() => window.scrollTo(0, document.getElementById('gear')!.offsetTop - 80));
  await shoot('paperdoll-item-head', 'tip-left-1280.png');
  await shoot('paperdoll-item-hands', 'tip-right-1280.png');
  await page.evaluate(() => {
    const cell = document.querySelector('[data-testid="paperdoll-item-main_hand"]')!;
    window.scrollTo(0, cell.getBoundingClientRect().top + window.scrollY - window.innerHeight + 90);
  });
  await shoot('paperdoll-item-main_hand', 'tip-flip-1280.png');
});

test.describe('phone tooltips', () => {
  test.use({ hasTouch: true, isMobile: true });

  for (const width of [360, 390]) {
    test(`tooltips @${width}`, async ({ page }) => {
      await openCharacter(page, width);
      const tap = async (testid: string, file: string, bottom = false) => {
        const target = page.getByTestId(testid);
        if (bottom)
          await target.evaluate((el) =>
            window.scrollTo(0, el.getBoundingClientRect().top + window.scrollY - window.innerHeight + 70),
          );
        else
          await target.evaluate((el) =>
            window.scrollTo(0, el.getBoundingClientRect().top + window.scrollY - 120),
          );
        await target.tap();
        await page.getByTestId('item-tooltip').waitFor();
        await settle(page);
        await page.screenshot({ path: out(file) });
        await target.tap();
      };
      await tap('paperdoll-item-head', `tip-phone-${width}.png`);
      await tap('paperdoll-uprow-head', `tip-phoneup-${width}.png`);
      await tap('paperdoll-item-main_hand', `tip-phoneflip-${width}.png`, true);
      // Closing: a tap elsewhere and Escape each leave no panel.
      const head = page.getByTestId('paperdoll-item-head');
      await head.evaluate((el) => window.scrollTo(0, el.getBoundingClientRect().top + window.scrollY - 120));
      await head.tap();
      await page.getByTestId('item-tooltip').waitFor();
      await page.getByTestId('paperdoll-note').tap();
      await page.keyboard.press('Escape');
      await page.getByTestId('item-tooltip').waitFor({ state: 'detached' });
      await settle(page);
      await page.screenshot({ path: out(`tip-phoneclosed-${width}.png`) });
    });
  }
});
