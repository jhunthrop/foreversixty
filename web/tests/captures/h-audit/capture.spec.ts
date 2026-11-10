// web/tests/captures/h-audit/capture.spec.ts
// The 2026-11-09 h-audit captures: every page the redesign reviewed before verdicts were saved, at
// 1440 and 390, opened through the parity gate's own state openers (tests/e2e/support/*-parity.ts
// and friends), so the signed-in states carry the same fixture characters the gate uses.
//   CAPTURE_OUT=<dir> CAPTURE_ONLY=<regex> CAPTURE_PORT=<port> \
//     npx playwright test -c tests/captures/h-audit/playwright.capture.config.ts
import { mkdirSync } from 'node:fs';
import path from 'node:path';
import { test, type Page } from '@playwright/test';
import { prepareParityPage, readParityConfig, REPO_ROOT, settleParityPage } from '../../e2e/support/parity';
import {
  openSimState,
  openToolState,
  type SimParityState,
  type ToolParityState,
} from '../../e2e/support/sim-parity';
import { openLogsState, type LogsParityState } from '../../e2e/support/logs-parity';
import { openReportState, type ReportParityState } from '../../e2e/support/report-parity';
import { openRankingsState, type RankingsParityState } from '../../e2e/support/rankings-parity';
import { openTiersState, type TiersParityState } from '../../e2e/support/tiers-parity';
import { openCharacter, type CharacterState } from '../../e2e/support/character-page';
import { openHome, type HomeState } from '../../e2e/support/home-page';
import { openBis } from '../../e2e/support/bis-page';
import { openBisLanding, type BisLandingState } from '../../e2e/support/bis-landing-page';
import { openGuidesState, type GuidesState } from '../../e2e/support/guides-page';
import {
  openPlanner,
  plannerAccount,
  PLANNER_CODES,
  type PlannerState,
} from '../../e2e/support/planner-page';

const OUT = process.env.CAPTURE_OUT ?? path.join(REPO_ROOT, 'design/reviews/2026-11-09-h-audit/first');
const ONLY = process.env.CAPTURE_ONLY ? new RegExp(process.env.CAPTURE_ONLY) : null;
const WIDTHS = [1440, 390] as const;
const HEIGHT = 1000;

const stateOf = <T>(file: string, id: string): T => {
  const state = readParityConfig(file).states.find((entry) => entry.id === id);
  if (!state) throw new Error(`no state ${id} in ${file}`);
  return state as unknown as T;
};

const plannerCode = (code: string, spec: string): string =>
  `/planner?code=${encodeURIComponent(code)}&spec=${spec}`;
const felsnapKey = plannerAccount().characters[0].key;

const PAGES: readonly [name: string, open: (page: Page) => Promise<void>][] = [
  ...(['landing-out', 'landing-in', 'result', 'saved'] as const).map(
    (id): [string, (page: Page) => Promise<void>] => [
      `sim-${id}`,
      (page) => openSimState(page, id as SimParityState),
    ],
  ),
  ...(['gear', 'talents', 'drops', 'weights', 'specs'] as const).map(
    (id): [string, (page: Page) => Promise<void>] => [
      `sim-${id}`,
      (page) => openToolState(page, id as ToolParityState),
    ],
  ),
  ...(['logs-you', 'reports-list'] as const).map((id): [string, (page: Page) => Promise<void>] => [
    id,
    (page) => openLogsState(page, id as LogsParityState),
  ]),
  ...(
    [
      'report-summary',
      'report-damage-done',
      'report-healing',
      'report-damage-taken',
      'report-deaths',
    ] as const
  ).map((id): [string, (page: Page) => Promise<void>] => [
    id,
    (page) =>
      openReportState(
        page,
        id as ReportParityState,
        stateOf<{ path: string }>('report-parity.json', id).path,
      ),
  ]),
  ...(['rankings', 'rankings-you', 'rankings-board'] as const).map(
    (id): [string, (page: Page) => Promise<void>] => [
      id,
      (page) => openRankingsState(page, id as RankingsParityState),
    ],
  ),
  ['ratings', async (page) => void (await page.goto('/ratings'))],
  ...(['visitor-60', 'owner-60-premium'] as const).map((id): [string, (page: Page) => Promise<void>] => [
    `character-${id}`,
    (page) => openCharacter(page, stateOf<CharacterState>('character-parity.json', id)),
  ]),
  ['planner', (page) => openPlanner(page, '/planner', 'out' as PlannerState)],
  [
    'planner-you',
    (page) => openPlanner(page, `/planner?character=${encodeURIComponent(felsnapKey)}`, 'felsnap'),
  ],
  [
    'planner-guide',
    (page) => openPlanner(page, plannerCode(PLANNER_CODES.survival, 'hunter-survival'), 'out'),
  ],
  ...(['dps', 'dps-you'] as const).map((id): [string, (page: Page) => Promise<void>] => [
    `tiers-${id}`,
    (page) => openTiersState(page, id as TiersParityState),
  ]),
  ...(['out', '60'] as const).map((id): [string, (page: Page) => Promise<void>] => [
    `bis-landing-${id}`,
    (page) => openBisLanding(page, '/bis', id as BisLandingState),
  ]),
  ['bis-spec', (page) => openBis(page, '/bis/hunter/marksmanship/60')],
  ...(['out', '60'] as const).map((id): [string, (page: Page) => Promise<void>] => [
    `home-${id}`,
    (page) => openHome(page, stateOf<HomeState>('home-parity.json', id)),
  ]),
  ['guides', (page) => openGuidesState(page, '/guides', 'out' as GuidesState)],
  ['guides-you', (page) => openGuidesState(page, '/guides', 'hunter34')],
  ['guides-class', (page) => openGuidesState(page, '/guides/hunter', 'out')],
  ['guides-spec', (page) => openGuidesState(page, '/guides/hunter/survival', 'out')],
];

mkdirSync(OUT, { recursive: true });

for (const [name, open] of PAGES) {
  if (ONLY && !ONLY.test(name)) continue;
  for (const width of WIDTHS) {
    test(`${name} at ${width}`, async ({ page }) => {
      test.slow();
      await page.setViewportSize({ width, height: HEIGHT });
      await prepareParityPage(page);
      await open(page);
      await settleParityPage(page);
      await page.screenshot({ path: path.join(OUT, `${name}-${width}.png`), fullPage: true });
    });
  }
}
