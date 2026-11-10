// web/tests/captures/h-doll/capture.spec.ts
// Review captures for lane h-doll (design/reviews/2026-11-10-h-doll): the three worn dolls that
// left the old grid paperdoll for the shared frame. A report's Player view (Gear at the pull), the
// sim result's Gear section and a saved sim's Character strip, at 1440, 2000 and 390. Not a test
// of behaviour. CAPTURE_VARIANT=before writes before/, else after/. Built with FOREVER_DATA=real
// so the dolls carry a full set; run from web/ against a served copy of dist (see the config).
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, test, type Page } from '@playwright/test';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const VARIANT = process.env.CAPTURE_VARIANT === 'before' ? 'before' : 'after';
const OUT = path.resolve(HERE, '../../../../design/reviews/2026-11-10-h-doll', VARIANT);
const WEB = path.resolve(HERE, '../../..');
const BUILD = (
  JSON.parse(readFileSync(path.join(WEB, 'src/data/active-build.json'), 'utf8')) as { build: string }
).build;
const WIDTHS = [
  { width: 1440, height: 900, scale: 1 },
  { width: 2000, height: 1100, scale: 1 },
  { width: 390, height: 844, scale: 2 },
] as const;

/** A level 60 Fury set from the real items file, a few slots enchanted (id:enchant). */
const SET: readonly [string, number, number?][] = [
  ['head', 12640, 2583],
  ['neck', 15411],
  ['shoulder', 12927, 2717],
  ['back', 13340, 849],
  ['chest', 11726, 1891],
  ['wrist', 13400, 1885],
  ['hands', 14551, 2564],
  ['waist', 13142],
  ['legs', 14554, 2583],
  ['feet', 14616, 911],
  ['finger1', 17713],
  ['finger2', 13098],
  ['trinket1', 11815],
  ['trinket2', 13965],
  ['main_hand', 12784, 1900],
  ['ranged', 12651],
];
const FS1 = `FS1:${BUILD}:warrior:orc:0/5530515/0:${SET.map(([slot, id, ench]) => `${slot}=${id}${ench ? `:${ench}` : ''}`).join(',')}|who=Thandrel:Realm`;

/** The log's slot order (lib/report/planner-link.ts LOG_PAPERDOLL_SLOTS). */
const LOG_ORDER = [
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

const SHELL_HTML = `<!doctype html>
<html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>Saved sim</title>
<link rel="stylesheet" href="/sim-island.css"></head>
<body><main id="main">
<div id="sim" data-sim-mount></div>
<script type="module" src="/sim-island.js"></script>
</main></body></html>`;

async function shoot(page: Page, id: string, width: number, testid: string): Promise<void> {
  const target = page.getByTestId(testid).first();
  await target.scrollIntoViewIfNeeded();
  await page.waitForTimeout(700);
  // From the top, so a sticky bar sits where it starts and not over the section.
  await page.evaluate(() => window.scrollTo(0, 0));
  await page.waitForTimeout(300);
  const box = await target.evaluate((element) => {
    const rect = element.getBoundingClientRect();
    return { x: rect.x + window.scrollX, y: rect.y + window.scrollY, width: rect.width, height: rect.height };
  });
  await page.screenshot({ path: path.join(OUT, `${id}-${width}.png`), fullPage: true, clip: box });
  await page.screenshot({ path: path.join(OUT, `${id}-page-${width}.png`), fullPage: true });
}

/** The real build's item facts, for the fixture build's report (whose warrior file has 7 items). */
const REAL = path.resolve(WEB, '..', 'data/builds', BUILD);
const REAL_FILES = /^(items\/|icons\/|loot\.json|enchants\.json|sets\.json|itemnames\.json)/;

/** Each set piece's real item level, as the log writes it. */
const ITEM_LEVEL = new Map(
  (
    JSON.parse(readFileSync(path.join(REAL, 'items/warrior.json'), 'utf8')) as {
      items: { id: number; item_level: number }[];
    }
  ).items.map((item) => [item.id, item.item_level]),
);

async function realItems(page: Page): Promise<void> {
  await page.route(`**/data/${BUILD}/**`, async (route) => {
    const rel = new URL(route.request().url()).pathname.split(`/data/${BUILD}/`)[1] ?? '';
    const file = path.join(REAL, rel);
    if (!REAL_FILES.test(rel) || !existsSync(file)) return route.continue();
    return route.fulfill({ path: file });
  });
}

async function openReportPlayer(page: Page): Promise<void> {
  await realItems(page);
  await page.route('**/fights/3/summary.json*', async (route) => {
    const response = await route.fetch();
    const summary = (await response.json()) as { combatants: { gear: unknown[] }[] };
    const bySlot = new Map(SET.map(([slot, id, ench]) => [slot, { id, ench }]));
    summary.combatants[0]!.gear = LOG_ORDER.map((slot) => {
      const piece = bySlot.get(slot);
      return {
        ID: piece?.id ?? 0,
        ItemLevel: piece ? (ITEM_LEVEL.get(piece.id) ?? 0) : 0,
        Enchants: piece?.ench ? [piece.ench] : null,
        BonusIDs: null,
        Gems: null,
      };
    });
    await route.fulfill({ response, json: summary });
  });
  await page.goto('/reports/fixture2abcd?fight=3');
  await page.getByTestId('at-pull-links').first().getByRole('link', { name: 'Gear ›' }).click();
  await expect(page.getByTestId('report-gear')).toBeVisible();
}

async function openSimGear(page: Page): Promise<void> {
  await page.goto('/sim');
  await page.getByTestId('sim-cold-paste-input').fill(FS1);
  await page.getByTestId('sim-cold-paste-run').click();
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 60_000 });
  await page.getByTestId('sim-section-gear').scrollIntoViewIfNeeded();
  await expect(page.getByTestId('sim-gear-card')).toBeVisible();
}

async function openSavedStrip(page: Page): Promise<void> {
  const id = 'simbulk23456';
  const fixture = JSON.parse(readFileSync(path.join(WEB, 'src/fixtures/sim/bulk-result.json'), 'utf8')) as {
    request: { character: { name: string; gear: unknown[] } };
  };
  fixture.request.character.name = 'Thandrel';
  fixture.request.character.gear = SET.map(([slot, id_, ench]) => ({
    slot,
    item_id: id_,
    ...(ench ? { enchant: ench } : {}),
  }));
  await page.route(`**/sim/${id}`, (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: SHELL_HTML }),
  );
  await page.route(`**/v1/sims/${id}`, (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        data: { ...fixture, sim_id: id },
        error: null,
        request_id: 'req-capture',
      }),
    }),
  );
  await page.goto(`/sim/${id}`);
  await expect(page.getByTestId('sim-character')).toBeVisible({ timeout: 30_000 });
}

const SHOTS = [
  { id: 'report-player', open: openReportPlayer, testid: 'report-gear' },
  { id: 'sim-gear', open: openSimGear, testid: 'sim-gear-card' },
  { id: 'sim-strip', open: openSavedStrip, testid: 'sim-character' },
] as const;

/** CAPTURE_ONLY=report-player (comma list): the report needs a fixture build, the sim a real one. */
const ONLY = process.env.CAPTURE_ONLY?.split(',');

for (const shot of SHOTS.filter((entry) => ONLY === undefined || ONLY.includes(entry.id))) {
  for (const size of WIDTHS) {
    test(`${shot.id} at ${size.width}`, async ({ browser }) => {
      const context = await browser.newContext({
        viewport: { width: size.width, height: size.height },
        deviceScaleFactor: size.scale,
        isMobile: size.width < 768,
        hasTouch: size.width < 768,
      });
      const page = await context.newPage();
      await shot.open(page);
      await page.waitForLoadState('networkidle');
      await shoot(page, shot.id, size.width, shot.testid);
      await context.close();
    });
  }
}
