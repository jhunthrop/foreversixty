// web/tests/captures/raid-spec-ask/capture.spec.ts
// The raid page's spec ask (guided-cards spec P.10.7), the sign-up row at 1440 and 390: In
// pressed before the pick, after the pick, and the saved spec under `Change spec`. PNGs go to
// design/reviews/2026-10-28-discord-guided-cards/task-10b-raid-spec-*.png.
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test, type Locator, type Page } from '@playwright/test';
import { VIEWER_MEMBER } from '../../../src/fixtures/guild/mock-guild';
import { NOW, ONYXIA_ID, raidUrl, stripId, stubGuildRaids } from '../../e2e/support/guild-raids';
import { gotoHydrated } from '../../e2e/support/hydrated';

const OUT_DIR = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  '../../../../design/reviews/2026-10-28-discord-guided-cards',
);
const WIDTHS = [1440, 390];
const PAD = 16;
const SETTLE_MS = 400;

const row = (page: Page): Locator => page.getByTestId(stripId('row', VIEWER_MEMBER));
const part = (page: Page, name: 'segments' | 'spec' | 'status', suffix = ''): Locator =>
  page.getByTestId(`${stripId(name, VIEWER_MEMBER)}${suffix}`);

async function open(page: Page, width: number, askSpec: boolean): Promise<void> {
  await page.setViewportSize({ width, height: 1000 });
  await stubGuildRaids(page, { viewer: 'member', askSpec });
  await page.clock.install({ time: NOW });
  await gotoHydrated(page, raidUrl(ONYXIA_ID), 'raid-page');
}

async function shoot(page: Page, name: string, width: number): Promise<void> {
  await page.waitForTimeout(SETTLE_MS);
  const box = await row(page).boundingBox();
  if (box === null) throw new Error('no sign-up row');
  await page.screenshot({
    path: path.join(OUT_DIR, `task-10b-raid-spec-${name}-${width}.png`),
    fullPage: true,
    clip: {
      x: Math.max(0, box.x - PAD),
      y: box.y - PAD,
      width: Math.min(width, box.width + PAD * 2),
      height: box.height + PAD * 2,
    },
  });
}

for (const width of WIDTHS) {
  test(`spec ask at ${width}`, async ({ page }) => {
    await open(page, width, true);
    await part(page, 'segments', '-in').click();
    await part(page, 'spec').waitFor();
    await shoot(page, 'before-pick', width);

    await part(page, 'spec').selectOption('hunter-survival');
    await part(page, 'status').getByText('kept').waitFor();
    await shoot(page, 'after-pick', width);
  });

  test(`change spec at ${width}`, async ({ page }) => {
    await open(page, width, false);
    await part(page, 'segments', '-late').click();
    await part(page, 'status').getByText('Saved').waitFor();
    await page.getByTestId(`${stripId('spec', VIEWER_MEMBER)}-toggle`).click();
    await part(page, 'spec').waitFor({ state: 'visible' });
    await shoot(page, 'change', width);
  });
}
