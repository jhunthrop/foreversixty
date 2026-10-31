// web/tests/captures/guild-raids/capture-429.spec.ts
// The raid form after the API's per-credential rate limit refused Post the raid (429), at 1440 and
// 390, for the 2026-10-31 hermetic-e2e review. PNGs go to design/reviews/2026-10-31-raid-429/.
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, test } from '@playwright/test';
import { gotoHydrated } from '../../e2e/support/hydrated';
import { CHICAGO, GUILD_URL, NOW, stubGuildRaids } from '../../e2e/support/guild-raids';
import { raidCopy } from '../../../src/lib/guild/raid-copy';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const OUT_DIR = path.resolve(HERE, '../../../../design/reviews/2026-10-31-raid-429');
const VIEWPORT_HEIGHT = 1000;
const WIDTHS = [1440, 390];

test.use({ timezoneId: CHICAGO });

for (const width of WIDTHS) {
  test(`raid form 429 @${width}`, async ({ page }) => {
    await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
    const stub = await stubGuildRaids(page, { viewer: 'officer' });
    await page.clock.install({ time: NOW });
    await gotoHydrated(page, `${GUILD_URL}#raids`, 'guild-raids-tab');
    await page.getByTestId('guild-raids-plan').click();
    const form = page.getByTestId('raid-form-create');
    await expect(form).toBeVisible();
    stub.refuseNext = { create: { code: 'rate_limited', status: 429 } };
    await page.getByTestId('raid-form-submit').click();
    await expect(form.getByRole('alert')).toHaveText(raidCopy.rateLimited);
    await form.scrollIntoViewIfNeeded();
    await form.screenshot({ path: path.join(OUT_DIR, `raid-form-429-${width}.png`) });
    await page.screenshot({ path: path.join(OUT_DIR, `raid-form-429-${width}-viewport.png`) });
  });
}
