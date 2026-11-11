// web/tests/captures/setup-steps/capture.spec.ts
// The review captures for /setup's reordered checklist (design/reviews/2026-11-10-setup-steps): signed
// out; signed in from Battle.net without a companion; a companion that found no addon; a companion
// that found it; and an addon behind the feed. Each opens the page as setup-states.spec.ts does (the
// committed fixtures and the mocks' clock) and shoots the full page and the plate alone. Not a test of
// behaviour.
import { mkdirSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test } from '@playwright/test';
import { prepareParityPage, settleParityPage } from '../../e2e/support/parity';
import { openSetupState, type SetupPageState } from '../../e2e/support/setup-page';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const OUT = path.resolve(HERE, '../../../../design/reviews/2026-11-10-setup-steps');
mkdirSync(OUT, { recursive: true });

const STATES: readonly { id: string; state: SetupPageState }[] = [
  { id: 'signed-out', state: 'out' },
  { id: 'no-companion', state: 'bnet' },
  { id: 'companion-no-addon', state: 'companion' },
  { id: 'companion-addon', state: 'installed' },
  { id: 'addon-outdated', state: 'outdated' },
];
const WIDTHS = [1440, 390] as const;
const HEIGHT = 1000;

for (const { id, state } of STATES) {
  for (const width of WIDTHS) {
    test(`${id} at ${width}`, async ({ page }) => {
      await page.setViewportSize({ width, height: HEIGHT });
      await prepareParityPage(page);
      await openSetupState(page, state);
      await settleParityPage(page);
      // A full-page shot never scrolls: load the lazy pictures below the fold first.
      await page.evaluate(async () => {
        const images = [...document.querySelectorAll('img')];
        for (const image of images) image.loading = 'eager';
        await Promise.all(images.map((image) => image.decode().catch(() => undefined)));
      });
      await page.screenshot({ path: path.join(OUT, `${id}-${width}.png`), fullPage: true });
      await page
        .getByRole('region', { name: 'Your setup' })
        .screenshot({ path: path.join(OUT, `${id}-${width}-plate.png`) });
    });
  }
}
