// web/tests/e2e/support/parity-mocks.ts
// The parity gate's mocks, rendered at test time from the build the spec serves (design/parity.md):
// the generators (design/mocks/ws_bis.py, ws_sim.py) decide every block, the served build only
// supplies the data. A refresh that changes content (picks, names, numbers, a longer line) changes
// both sides alike; a block the build grows on its own still differs. The committed mocks in
// design/mocks/web stay the human reference; scripts/check-parity-mocks.mjs reports when they
// fall behind what these render.
import { execFileSync } from 'node:child_process';
import { mkdirSync, rmSync } from 'node:fs';
import path from 'node:path';
import { PARITY_OUT, REPO_ROOT } from './parity';

/** Where the test-time mocks land, one folder per page: test-results/parity/mocks/<page>. */
export const PARITY_MOCKS_OUT = path.join(PARITY_OUT, 'mocks');
const GENERATORS = path.join(REPO_ROOT, 'design/mocks');
const PYTHON = process.env.PYTHON ?? 'python3';

/** One generator run, its output kept for the failure message. */
function run(command: string, args: string[], env: NodeJS.ProcessEnv = {}): void {
  try {
    execFileSync(command, args, {
      cwd: REPO_ROOT,
      env: { ...process.env, ...env },
      encoding: 'utf8',
      stdio: ['ignore', 'pipe', 'pipe'],
    });
  } catch (error) {
    const { stdout = '', stderr = '' } = error as { stdout?: string; stderr?: string };
    throw new Error(
      `parity mocks: \`${[command, ...args].join(' ')}\` failed (design/parity.md)\n${stdout}${stderr}`,
      { cause: error },
    );
  }
}

function freshDir(page: string): string {
  const dir = path.join(PARITY_MOCKS_OUT, page);
  rmSync(dir, { recursive: true, force: true });
  mkdirSync(dir, { recursive: true });
  return dir;
}

/**
 * The four BiS mocks from the build at `baseURL` (the spec's own preview, never whatever holds
 * another port): bis_facts.mjs reads the facts off its pages, ws_bis.py renders them.
 * Returns the folder that holds web-bis-<state>.html.
 */
export function renderBisMocks(baseURL: string): string {
  const dir = freshDir('bis');
  const facts = path.join(dir, 'bis-facts.json');
  run('node', [path.join(GENERATORS, 'bis_facts.mjs'), baseURL, facts]);
  run(PYTHON, [path.join(GENERATORS, 'ws_bis.py')], { FS_BIS_FACTS: facts, FS_MOCK_OUT: dir });
  return dir;
}

/**
 * The simulator mocks from the files the spec serves (support/sim-parity.ts: the fixture run
 * and data/builds/<active build>, both in this checkout). Returns the folder that holds them.
 */
export function renderSimMocks(): string {
  const dir = freshDir('sim');
  run(PYTHON, [path.join(GENERATORS, 'ws_sim.py')], { FS_DATA_ROOT: REPO_ROOT, FS_MOCK_OUT: dir });
  return dir;
}

/** The simulator tool mocks (ws_simtools.py) from this checkout's data/builds. Returns the folder. */
export function renderSimToolsMocks(): string {
  const dir = freshDir('simtools');
  run(PYTHON, [path.join(GENERATORS, 'ws_simtools.py')], { FS_DATA_ROOT: REPO_ROOT, FS_MOCK_OUT: dir });
  return dir;
}

/**
 * The logs, reports list, rankings and ratings mocks (ws_logs.py) from the files the spec serves
 * (support/logs-parity.ts): logs_facts.mjs bundles the copy modules and the rankings stubs, and
 * ws_logs.py reads web/src/fixtures/reports and the bundled sample's meta. Returns the folder.
 */
export function renderLogsMocks(): string {
  const dir = freshDir('logs');
  const facts = path.join(dir, 'logs-facts.json');
  run('node', [path.join(GENERATORS, 'logs_facts.mjs'), facts]);
  run(PYTHON, [path.join(GENERATORS, 'ws_logs.py')], { FS_LOGS_FACTS: facts, FS_MOCK_OUT: dir });
  return dir;
}

/**
 * The guide mocks named (ws_guides.py --list), from this checkout's content, guide libs and the
 * fixtures support/guides-page.ts serves: guides_facts.mjs runs the pages' own models, ws_guides.py
 * draws just those mocks. Returns the folder that holds web-guides-*.html.
 */
export function renderGuidesMocks(mocks: readonly string[]): string {
  const dir = freshDir('guides');
  const facts = path.join(dir, 'guides-facts.json');
  run('node', [path.join(GENERATORS, 'guides_facts.mjs'), facts]);
  run(PYTHON, [path.join(GENERATORS, 'ws_guides.py'), ...mocks], {
    FS_GUIDES_FACTS: facts,
    FS_MOCK_OUT: dir,
  });
  return dir;
}

/**
 * The shared build mocks named (ws_premium.py --list: web-b*), from the copy modules and this
 * checkout's planner data: premium_facts.mjs reads the shared page's copy, planner_facts.mjs the
 * build states ws_premium.py borrows from ws_planner.py, pages_facts.mjs the 404's class picker
 * web-b-unavailable borrows from ws_pages.py. Returns the folder that holds web-b*.html.
 */
export function renderSharedBuildMocks(mocks: readonly string[]): string {
  const dir = freshDir('shared-build');
  const premiumFacts = path.join(dir, 'premium-facts.json');
  const plannerFacts = path.join(dir, 'planner-facts.json');
  const pagesFacts = path.join(dir, 'pages-facts.json');
  run('node', [path.join(GENERATORS, 'premium_facts.mjs'), premiumFacts]);
  run('node', [path.join(GENERATORS, 'planner_facts.mjs'), plannerFacts]);
  run('node', [path.join(GENERATORS, 'pages_facts.mjs'), pagesFacts]);
  run(PYTHON, [path.join(GENERATORS, 'ws_premium.py'), ...mocks], {
    FS_PREMIUM_FACTS: premiumFacts,
    FS_PLANNER_FACTS: plannerFacts,
    FS_PAGES_FACTS: pagesFacts,
    FS_MOCK_OUT: dir,
  });
  return dir;
}

/**
 * The planner mocks named (ws_planner.py --list), from this checkout's data/builds files, the
 * planner libs and the fixtures support/planner-page.ts serves: planner_facts.mjs runs the page's
 * own models, ws_planner.py draws just those mocks. Returns the folder that holds web-planner*.html.
 */
export function renderPlannerMocks(mocks: readonly string[]): string {
  const dir = freshDir('planner');
  const facts = path.join(dir, 'planner-facts.json');
  run('node', [path.join(GENERATORS, 'planner_facts.mjs'), facts]);
  run(PYTHON, [path.join(GENERATORS, 'ws_planner.py'), ...mocks], {
    FS_PLANNER_FACTS: facts,
    FS_MOCK_OUT: dir,
  });
  return dir;
}

/**
 * The changelog, content-page and 404 mocks named (ws_pages.py --list), from this checkout's
 * content collections, copy modules, the changelog model and the fixture account
 * support/pages-page.ts serves: pages_facts.mjs bundles them, ws_pages.py draws just those mocks.
 * Returns the folder that holds them.
 */
export function renderPagesMocks(mocks: readonly string[]): string {
  const dir = freshDir('pages');
  const facts = path.join(dir, 'pages-facts.json');
  run('node', [path.join(GENERATORS, 'pages_facts.mjs'), facts]);
  run(PYTHON, [path.join(GENERATORS, 'ws_pages.py'), ...mocks], { FS_PAGES_FACTS: facts, FS_MOCK_OUT: dir });
  return dir;
}

/**
 * The rankings mocks (ws_rankings.py) from the files the spec serves (support/rankings-parity.ts):
 * rankings_facts.mjs bundles the rankings stubs and reads them through the page's own view helpers.
 * Returns the folder that holds web-rankings*.html.
 */
export function renderRankingsMocks(): string {
  const dir = freshDir('rankings');
  const facts = path.join(dir, 'rankings-facts.json');
  run('node', [path.join(GENERATORS, 'rankings_facts.mjs'), facts]);
  run(PYTHON, [path.join(GENERATORS, 'ws_rankings.py')], { FS_RANKINGS_FACTS: facts, FS_MOCK_OUT: dir });
  return dir;
}

/**
 * The /ratings mock (ws_ratings.py) from the score and rating copy modules, through
 * rankings_facts.mjs. Returns the folder that holds web-ratings.html.
 */
export function renderRatingsMocks(): string {
  const dir = freshDir('ratings');
  const facts = path.join(dir, 'rankings-facts.json');
  run('node', [path.join(GENERATORS, 'rankings_facts.mjs'), facts]);
  run(PYTHON, [path.join(GENERATORS, 'ws_ratings.py')], { FS_RANKINGS_FACTS: facts, FS_MOCK_OUT: dir });
  return dir;
}

/**
 * The tier list mocks (ws_tiers.py) from this checkout's data/builds/<active build>/bis, the files
 * the fixture build prerenders /tiers from. Returns the folder that holds web-tiers*.html.
 */
export function renderTiersMocks(): string {
  const dir = freshDir('tiers');
  run(PYTHON, [path.join(GENERATORS, 'ws_tiers.py')], { FS_DATA_ROOT: REPO_ROOT, FS_MOCK_OUT: dir });
  return dir;
}

/**
 * The report mocks named (ws_report.py --list), from the bundled real log the spec serves
 * (web/public/sample-report/sample) and the copy modules: logs_facts.mjs bundles the copy,
 * ws_report.py draws just those mocks, each as the sample page the spec serves. Returns the folder that holds web-report-*.html.
 */
export function renderReportMocks(mocks: readonly string[]): string {
  const dir = freshDir('report');
  const facts = path.join(dir, 'logs-facts.json');
  run('node', [path.join(GENERATORS, 'logs_facts.mjs'), facts]);
  // FS_REPORT_SAMPLE: every report mock drawn as /reports/sample, the page the spec opens on this log.
  run(PYTHON, [path.join(GENERATORS, 'ws_report.py'), ...mocks], {
    FS_LOGS_FACTS: facts,
    FS_MOCK_OUT: dir,
    FS_REPORT_SAMPLE: '1',
  });
  return dir;
}

/**
 * The /bis landing mocks named (ws_bis_landing.py --list), from the landing's own models, this
 * checkout's data/builds/<active build>/bis and the fixtures support/bis-landing-page.ts serves:
 * bis_landing_facts.mjs runs the models, ws_bis_landing.py draws just those mocks. Returns the
 * folder that holds web-bis-landing*.html.
 */
export function renderBisLandingMocks(mocks: readonly string[]): string {
  const dir = freshDir('bis-landing');
  const facts = path.join(dir, 'bis-landing-facts.json');
  run('node', [path.join(GENERATORS, 'bis_landing_facts.mjs'), facts]);
  run(PYTHON, [path.join(GENERATORS, 'ws_bis_landing.py'), ...mocks], {
    FS_BIS_LANDING_FACTS: facts,
    FS_MOCK_OUT: dir,
  });
  return dir;
}

/**
 * The public /creators mock (ws_creators.py) from this checkout's lib/creators/copy.ts:
 * creators_facts.mjs bundles the copy, ws_creators.py draws web-creators.html. Returns the folder.
 */
export function renderCreatorsMocks(): string {
  const dir = freshDir('creators');
  const facts = path.join(dir, 'creators-page-facts.json');
  run('node', [path.join(GENERATORS, 'creators_facts.mjs'), facts]);
  run(PYTHON, [path.join(GENERATORS, 'ws_creators.py')], { FS_CREATORS_FACTS: facts, FS_MOCK_OUT: dir });
  return dir;
}

/** The /app/auth mocks (ws_app_auth.py), from the page's copy module. Returns the folder. */
export function renderAppAuthMocks(): string {
  const dir = freshDir('app-auth');
  run(PYTHON, [path.join(GENERATORS, 'ws_app_auth.py')], { FS_MOCK_OUT: dir });
  return dir;
}

/**
 * The app beta mocks (ws_app_beta.py: /app in each signup state, /admin/testers) from this
 * checkout's copy modules and QR (app_beta_facts.mjs). Returns the folder that holds them.
 */
export function renderAppBetaMocks(): string {
  const dir = freshDir('app-beta');
  const facts = path.join(dir, 'app-beta-facts.json');
  run('node', [path.join(GENERATORS, 'app_beta_facts.mjs'), facts]);
  run(PYTHON, [path.join(GENERATORS, 'ws_app_beta.py')], { FS_APP_BETA_FACTS: facts, FS_MOCK_OUT: dir });
  return dir;
}
