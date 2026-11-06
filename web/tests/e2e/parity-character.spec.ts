// web/tests/e2e/parity-character.spec.ts
// The character page against its committed mocks (design/parity.md, design/specs/2026-11-06-character.md): every
// state at 1440 x 1000, each region's blocks and box compared (support/parity.ts), a side-by-side
// per state in test-results/parity/character-<state>.png. Desktop project only: the mocks are 1440.
import { test, expect } from './support/test';
import { openCharacter, type CharacterState } from './support/character-page';
import {
  captureOutlined,
  captureRegions,
  compareRegions,
  formatFailures,
  MOCKS_DIR,
  openMock,
  prepareParityPage,
  readParityConfig,
  settleParityPage,
  writeSideBySide,
} from './support/parity';

const CONFIG = readParityConfig('character-parity.json');

test.describe('Character mock parity at 1440', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');
  test.use({ viewport: CONFIG.viewport });

  for (const state of CONFIG.states as ((typeof CONFIG.states)[number] & CharacterState)[]) {
    test(`${state.id} matches ${state.mock}`, async ({ page, browser }) => {
      test.slow();
      const mockPage = await browser.newPage({ viewport: CONFIG.viewport });
      // The mocks are the committed ones: both sides read the committed fixtures, never published
      // data (design/parity.md), so a data refresh cannot move either side.
      await openMock(mockPage, MOCKS_DIR, state.mock);
      const mock = await captureRegions(mockPage);

      await prepareParityPage(page);
      await openCharacter(page, state);
      await settleParityPage(page);
      const built = await captureRegions(page);

      await writeSideBySide(
        browser,
        `character-${state.id}`,
        await captureOutlined(mockPage, mock),
        await captureOutlined(page, built),
      );
      await mockPage.close();

      const failures = compareRegions(mock, built);
      expect(failures, `${state.id} differs from ${state.mock}:\n${formatFailures(failures)}`).toEqual([]);
    });
  }
});
