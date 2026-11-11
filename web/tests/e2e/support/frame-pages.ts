// web/tests/e2e/support/frame-pages.ts
// One page of every type the site serves, each opened in a settled state from the committed
// fixtures (the parity openers where a page has one), for the page-frame checks
// (page-frame.spec.ts) and the frame captures (frame-capture.spec.ts): the site header bar, the
// page head band, the content column and the footer share one frame on every page
// (owner 2026-11-10; tokens.css --page-max and --page-gutter).
import type { Page } from '@playwright/test';
import { apiUnreachable } from './test';
import { openBis } from './bis-page';
import { openBisLanding } from './bis-landing-page';
import { openCharacter, type CharacterState } from './character-page';
import { openConnectState } from './connect-page';
import { openGuidesState } from './guides-page';
import { openGuildState } from './guild-parity';
import { openHome } from './home-page';
import { openLogsState } from './logs-parity';
import { openPagesState } from './pages-page';
import { readParityConfig, waitForChromeHydrated } from './parity';
import { openPlanner } from './planner-page';
import { openRankingsState } from './rankings-parity';
import { openReportState } from './report-parity';
import { OBNOXIOUS, armoryPointer, signInWith, storePointer } from './selector';
import { openSetupState } from './setup-page';
import { SHARED_BUILD_ID, openSharedBuild } from './shared-build-page';
import { openSimState, openToolState } from './sim-parity';
import { openTiersState } from './tiers-parity';
import { PANEL_URL, STATES, routeCoach, signInAs } from './coach-stubs';
import { gotoHydrated } from './hydrated';

export interface FramePage {
  /** A short name: the capture's file name and the test's title. */
  id: string;
  open: (page: Page) => Promise<void>;
}

const CHARACTER = readParityConfig('character-parity.json').states[0] as unknown as CharacterState;

/** A page with no parity opener of its own: the static page signed out, its reads unanswered. */
const plain =
  (pathname: string) =>
  (page: Page): Promise<void> =>
    openPagesState(page, pathname, 'out');

/** /account signed in with one character (the account parity states read rendered mock facts). */
async function openAccount(page: Page): Promise<void> {
  await signInWith(page, [OBNOXIOUS]);
  await storePointer(page, armoryPointer(OBNOXIOUS));
  await page.goto('/account');
  await page.waitForLoadState('networkidle');
}

export const FRAME_PAGES: readonly FramePage[] = [
  { id: 'home', open: (page) => openHome(page, {}) },
  { id: 'bis-landing', open: (page) => openBisLanding(page, '/bis', 'out') },
  { id: 'bis', open: (page) => openBis(page, '/bis/hunter/marksmanship/60') },
  { id: 'guides', open: (page) => openGuidesState(page, '/guides', 'out') },
  { id: 'guides-class', open: (page) => openGuidesState(page, '/guides/hunter', 'out') },
  { id: 'guides-spec', open: (page) => openGuidesState(page, '/guides/hunter/survival', 'out') },
  { id: 'tiers', open: (page) => openTiersState(page, 'dps') },
  { id: 'sim', open: (page) => openSimState(page, 'landing-out') },
  { id: 'sim-result', open: (page) => openSimState(page, 'result') },
  { id: 'sim-gear', open: (page) => openToolState(page, 'gear') },
  { id: 'planner', open: (page) => openPlanner(page, '/planner') },
  { id: 'rankings', open: (page) => openRankingsState(page, 'rankings') },
  { id: 'logs', open: (page) => openLogsState(page, 'logs') },
  { id: 'reports', open: (page) => openLogsState(page, 'reports') },
  { id: 'report', open: (page) => openReportState(page, 'report-summary', '/reports/sample?tab=summary') },
  { id: 'character', open: (page) => openCharacter(page, CHARACTER) },
  { id: 'guild', open: (page) => openGuildState(page, 'dense-officer') },
  { id: 'guild-roster', open: (page) => openGuildState(page, 'dense-roster') },
  { id: 'guild-loot', open: (page) => openGuildState(page, 'dense-loot') },
  { id: 'account', open: openAccount },
  { id: 'setup', open: (page) => openSetupState(page, 'out') },
  { id: 'connect', open: (page) => openConnectState(page, 'signed-out') },
  { id: 'login', open: plain('/login') },
  { id: 'premium', open: plain('/premium') },
  { id: 'characters', open: plain('/characters') },
  { id: 'ratings', open: plain('/ratings') },
  { id: 'creators', open: plain('/creators') },
  { id: 'app', open: plain('/app') },
  { id: 'changelog', open: plain('/changelog') },
  { id: 'legal', open: plain('/privacy') },
  { id: 'about', open: plain('/about') },
  { id: '404', open: plain('/this-page-does-not-exist') },
  { id: 'b', open: (page) => openSharedBuild(page, 'survival.html', SHARED_BUILD_ID) },
  { id: 'b-unavailable', open: plain('/b-unavailable') },
];

/** The coach's panel on the worked Onyxia pull, a Premium viewer with the full review. */
async function openCoach(page: Page): Promise<void> {
  await signInAs(page, 'premium');
  await routeCoach(page, STATES['ready-model']);
  await gotoHydrated(page, PANEL_URL(), 'coach-panel');
}

/** Every frame page plus the report tabs, guild views and coach panel that the frame checks leave
 *  out, for the captures (frame-capture.spec.ts) that review a change across the whole site. */
export const CAPTURE_PAGES: readonly FramePage[] = [
  ...FRAME_PAGES,
  {
    id: 'report-damage-done',
    open: (page) => openReportState(page, 'report-damage-done', '/reports/sample?tab=damage-done'),
  },
  {
    id: 'report-deaths',
    open: (page) => openReportState(page, 'report-deaths', '/reports/sample?fight=20&tab=deaths'),
  },
  { id: 'coach', open: openCoach },
  { id: 'guild-raids', open: (page) => openGuildState(page, 'dense-raids') },
  { id: 'guild-report', open: (page) => openGuildState(page, 'report') },
];

/** Opens a frame page and waits for the header to settle; reads its opener leaves unanswered fail
 *  as the API out of reach, which every page renders without. */
export async function openFramePage(page: Page, framePage: FramePage): Promise<void> {
  await apiUnreachable(page, '**/v1/**');
  await framePage.open(page);
  await waitForChromeHydrated(page);
  await page.locator('[data-site-header][data-header-wired]').waitFor({ state: 'attached' });
}

/** A box's left and right edges, in CSS px. */
export interface Edges {
  left: number;
  right: number;
}

/** The frame as the page draws it: the header bar's edges and the content box of every visible
 *  `.page-column` (the page head band, the page body and the footer each open one). */
export interface FrameMeasure {
  header: Edges;
  columns: readonly (Edges & { where: string })[];
}

export async function measureFrame(page: Page): Promise<FrameMeasure> {
  return page.evaluate(() => {
    const bar = document.querySelector('.site-header')!.getBoundingClientRect();
    const columns = Array.from(document.querySelectorAll<HTMLElement>('.page-column'))
      .filter((el) => el.getBoundingClientRect().width > 0 && getComputedStyle(el).visibility !== 'hidden')
      .map((el) => {
        const box = el.getBoundingClientRect();
        const style = getComputedStyle(el);
        const where = el.closest('footer')
          ? 'footer'
          : el.hasAttribute('data-page-body')
            ? 'body'
            : (el.parentElement?.className.split(' ')[0] ?? 'band');
        return {
          where,
          left: Math.round(box.left + parseFloat(style.paddingLeft)),
          right: Math.round(box.right - parseFloat(style.paddingRight)),
        };
      });
    return { header: { left: Math.round(bar.left), right: Math.round(bar.right) }, columns };
  });
}
