// web/tests/e2e/parity-bis.spec.ts
// The BiS page against its committed mocks (design/parity.md): four states at 1440 x 1000, each
// region's blocks and box compared (support/parity.ts). A side-by-side per state lands in
// test-results/parity/. Desktop project only: the mocks are 1440 captures.
import { test, expect } from './support/test';
import { openBis } from './support/bis-page';
import {
  captureOutlined,
  captureRegions,
  compareRegions,
  formatFailures,
  openMock,
  prepareParityPage,
  readParityConfig,
  settleParityPage,
  writeSideBySide,
  type ParityFailure,
  type RegionSnapshot,
} from './support/parity';

const CONFIG = readParityConfig('bis-parity.json');

/**
 * Known drift, dated in its entry: these regions (and the regions that contain them) are
 * expected to differ, the regions below them are compared after the drifted height is taken out,
 * and the test fails the day the drift is gone so the entry is deleted with it.
 */
interface ExpectedDrift {
  /** Drifted: every failure here is expected. */
  regions: readonly string[];
  /** Hold the drift: every failure expected, and the regions below move up by its extra height. */
  containers: readonly string[];
  /** Pushed by the drift inside its container (the title sits at the head's bottom edge): their
   *  blocks are still compared, their boxes are not. */
  moved: readonly string[];
  why: string;
}

/** The head answer and what it pushes: the title sits on the head's bottom edge. */
const HEAD_ANSWER = {
  regions: ['head-answer'],
  containers: ['head'],
  moved: ['head-title', 'head-meta'],
} as const;

// land/bis-header (merged 2026-11-06, eb335cb7e) removed the head disclosures. What is left in
// BisAnswer.svelte, each entry dated; delete an entry the day its state matches (the test fails
// then, so it cannot outlive the drift).
const EXPECTED_DRIFT: Readonly<Record<string, ExpectedDrift>> = {
  // TODO(2026-11-06): the trust line and links drop the space before each middot
  // ("consumables.· View ranking sim", "simulator ·Talents"), so the build's trust line fits one
  // line where the mock's wraps (16 px). Fix the spacing in BisAnswer.svelte.
  'dps-60': { ...HEAD_ANSWER, why: 'BisAnswer middot spacing (2026-11-06)' },
  // TODO(2026-11-06): the healer metrics list (Overheal, Mana lasts, Healing per mana: 3 terms,
  // 6 definitions) under the figure. Not in the approved head (one figure, trust line, links);
  // owner decision: drop it, or approve it and add it to the mock (ws_bis.py answer()).
  'healer-60': { ...HEAD_ANSWER, why: 'healer metrics list in the head, decision pending (2026-11-06)' },
  // TODO(2026-11-06): the tank head is four figures (effective health, DTPS, chance of death, TPS)
  // and a TMI line; the mock shows one (DTPS). Owner decision as for the healer metrics.
  'tank-60': { ...HEAD_ANSWER, why: 'tank four-figure head, decision pending (2026-11-06)' },
};

/** The build's regions below the drifted container, moved up by the height the drift added. */
function withoutDriftHeight(
  mock: RegionSnapshot[],
  build: RegionSnapshot[],
  drift: ExpectedDrift,
): RegionSnapshot[] {
  const container = drift.containers[0];
  const m = mock.find((r) => r.name === container);
  const b = build.find((r) => r.name === container);
  if (m === undefined || b === undefined) return build;
  const extra = b.box.height - m.box.height;
  const bottom = b.box.y + b.box.height;
  return build.map((r) => (r.box.y >= bottom ? { ...r, box: { ...r.box, y: r.box.y - extra } } : r));
}

function split(failures: ParityFailure[], drift: ExpectedDrift | undefined) {
  const exempt = new Set([...(drift?.regions ?? []), ...(drift?.containers ?? [])]);
  const moved = new Set(drift?.moved ?? []);
  const isExpected = (f: ParityFailure) => exempt.has(f.region) || (f.kind === 'box' && moved.has(f.region));
  return {
    expected: failures.filter(isExpected),
    unexpected: failures.filter((f) => !isExpected(f)),
  };
}

test.describe('BiS mock parity at 1440', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');
  test.use({ viewport: CONFIG.viewport });

  for (const state of CONFIG.states) {
    test(`${state.id} matches ${state.mock}`, async ({ page, browser }) => {
      test.slow();
      const mockPage = await browser.newPage({ viewport: CONFIG.viewport });
      await openMock(mockPage, state.mock);
      const mock = await captureRegions(mockPage);

      await prepareParityPage(page);
      await openBis(page, state.path);
      await settleParityPage(page);
      const built = await captureRegions(page);

      await writeSideBySide(
        browser,
        `bis-${state.id}`,
        await captureOutlined(mockPage, mock),
        await captureOutlined(page, built),
      );
      await mockPage.close();

      const drift = EXPECTED_DRIFT[state.id];
      const compared = drift === undefined ? built : withoutDriftHeight(mock, built, drift);
      const { expected, unexpected } = split(compareRegions(mock, compared), drift);
      expect(unexpected, `${state.id} differs from ${state.mock}:\n${formatFailures(unexpected)}`).toEqual(
        [],
      );
      if (drift !== undefined) {
        test.info().annotations.push({ type: 'expected drift', description: formatFailures(expected) });
        expect(
          expected.length,
          `${state.id}: the expected drift (${drift.why}) is gone; delete its EXPECTED_DRIFT entry`,
        ).toBeGreaterThan(0);
      }
    });
  }
});
