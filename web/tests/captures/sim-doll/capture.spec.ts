// web/tests/captures/sim-doll/capture.spec.ts
// Review captures for lane sim-doll (design/reviews/2026-11-11-sim-doll): the sim result's Gear
// section on the shared paperdoll frame at 1440, 2000 and 390, beside the BiS page's doll at 1440
// for one gear look (the character page exists only on the fixture build: its doll is compared from
// the character lane's own captures). Not a test of behaviour. Built with
// FOREVER_DATA=real so the sim doll carries a full set; run from web/ against a served copy of dist
// (see the config). CAPTURE_ONLY=sim-gear (comma list) narrows the shots.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, test, type Page } from '@playwright/test';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const OUT = path.resolve(HERE, '../../../../design/reviews/2026-11-11-sim-doll/after');
const WEB = path.resolve(HERE, '../../..');
const BUILD = (
  JSON.parse(readFileSync(path.join(WEB, 'src/data/active-build.json'), 'utf8')) as { build: string }
).build;
const WIDTHS = [
  { width: 1440, height: 900, scale: 1 },
  { width: 2000, height: 1100, scale: 1 },
  { width: 390, height: 844, scale: 2 },
] as const;
const DESKTOP = [WIDTHS[0]] as const;

/** A level 60 Fury set from the real items file, enchanted as a Fury warrior would (id:enchant). */
const SET: readonly [string, number, number?][] = [
  ['head', 12640, 1506],
  ['neck', 15411],
  ['shoulder', 12927],
  ['back', 13340, 849],
  ['chest', 11726, 1891],
  ['wrist', 13400, 1885],
  ['hands', 14551, 8207],
  ['waist', 13142],
  ['legs', 14554, 1506],
  ['feet', 14616, 911],
  ['finger1', 17713],
  ['finger2', 13098],
  ['trinket1', 11815],
  ['trinket2', 13965],
  ['main_hand', 12784, 1900],
  ['ranged', 12651],
];
const FS1 = `FS1:${BUILD}:warrior:orc:0/5530515/0:${SET.map(([slot, id, ench]) => `${slot}=${id}${ench ? `:${ench}` : ''}`).join(',')}|who=Thandrel:Realm`;

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

async function openSimGear(page: Page): Promise<void> {
  await page.goto('/sim');
  await page.getByTestId('sim-cold-paste-input').fill(FS1);
  await page.getByTestId('sim-cold-paste-run').click();
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 60_000 });
  await page.getByTestId('sim-section-gear').scrollIntoViewIfNeeded();
  await expect(page.getByTestId('sim-gear-card')).toBeVisible();
  await expect(
    page.getByTestId('sim-gear-doll').locator('.pdf-icon').locator('visible=true').first(),
  ).toBeVisible();
}

/** The same character before its run: the doll is the page's one Plate. */
async function openSimLanding(page: Page): Promise<void> {
  await page.goto(`/sim?code=${encodeURIComponent(FS1)}`);
  await expect(page.getByTestId('sim-run-button')).toBeEnabled({ timeout: 60_000 });
  await expect(
    page.getByTestId('sim-gear-doll').locator('.pdf-icon').locator('visible=true').first(),
  ).toBeVisible();
}

async function openBis(page: Page): Promise<void> {
  await page.goto('/bis/warrior/fury/60');
  await expect(page.getByTestId('paperdoll-frame-slot-head')).toBeVisible({ timeout: 30_000 });
}

const SHOTS = [
  { id: 'sim-gear', open: openSimGear, testid: 'sim-gear-card', widths: WIDTHS },
  { id: 'sim-landing', open: openSimLanding, testid: 'sim-gear-card', widths: WIDTHS },
  { id: 'bis-doll', open: openBis, testid: 'bis-page', widths: DESKTOP },
] as const;

const ONLY = process.env.CAPTURE_ONLY?.split(',');

for (const shot of SHOTS.filter((entry) => ONLY === undefined || ONLY.includes(entry.id))) {
  for (const size of shot.widths) {
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
