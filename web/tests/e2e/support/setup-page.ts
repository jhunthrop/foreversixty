// web/tests/e2e/support/setup-page.ts
// /setup in its states (setup spec 2026-11-07 section 5) from the committed fixtures the setup
// mocks draw (fixtures/setup/README.md): the account is fixtures/home/me-34.json (Felsnap, an addon
// build), its Battle.net-only variant (the build's source set to blizzard, as setup_facts.mjs
// derives it) or me-new.json (no character); the companion is fixtures/setup/devices.json. The
// clock is pinned to the mocks' 2026-11-06T10:00Z. Also renders the /setup mocks at test time for
// the parity gate (design/parity.md): setup_facts.mjs reads the copy modules and these fixtures,
// ws_setup.py draws the mocks from those facts.
import { execFileSync } from 'node:child_process';
import { mkdirSync, readFileSync, rmSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page } from '@playwright/test';
import { expect } from './test';
import { PARITY_OUT, REPO_ROOT } from './parity';

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../fixtures');
/** The mocks' own clock (setup_facts.mjs NOW). */
export const SETUP_NOW = new Date('2026-11-06T10:00:00Z');

export type SetupPageState =
  | 'out'
  | 'mac'
  | 'mobile'
  | 'unknown'
  | 'signin-failed'
  | 'error'
  | 'loading'
  | 'none'
  | 'importing'
  | 'partial'
  | 'import-failed'
  | 'failed'
  | 'bnet'
  | 'addon'
  | 'pairing'
  | 'done'
  | 'claim'
  | 'companion'
  | 'installed'
  | 'outdated';

export const SETUP_STATES: readonly SetupPageState[] = [
  'out',
  'mac',
  'mobile',
  'unknown',
  'signin-failed',
  'error',
  'loading',
  'none',
  'importing',
  'partial',
  'import-failed',
  'failed',
  'bnet',
  'addon',
  'pairing',
  'done',
  'claim',
  'companion',
  'installed',
  'outdated',
];

/** What the companion found in Interface/AddOns, per state that has a companion scan (owner
 *  2026-11-10): none, the feed's newest, or one behind it (fixtures/setup/addon-latest.json is 2.0.37). */
const ADDON_SCANS: Partial<Record<SetupPageState, { folder: string; version: string }[]>> = {
  companion: [],
  installed: [
    { folder: 'ForeverSixty', version: '2.0.37' },
    { folder: 'ForeverSixtyInbox', version: '2.0.37' },
    { folder: 'ForeverSixtyData', version: '2026.11.05' },
  ],
  outdated: [
    { folder: 'ForeverSixty', version: '2.0.30' },
    { folder: 'ForeverSixtyInbox', version: '2.0.30' },
  ],
};

/** The device list each signed-in state reads: none, the companion, or the companion with its scan. */
function devicesOf(state: SetupPageState): unknown[] {
  if (state === 'done') return fixture('setup/devices.json');
  const scan = ADDON_SCANS[state];
  if (scan === undefined) return [];
  return fixture<Record<string, unknown>[]>('setup/devices.json').map((d) => ({ ...d, addons: scan }));
}

function fixture<T = unknown>(name: string): T {
  return JSON.parse(readFileSync(path.join(FIXTURES, name), 'utf8')) as T;
}

interface FixtureMe {
  characters: { build?: Record<string, unknown> }[];
  [key: string]: unknown;
}

/** The Battle.net-only account: me-34 with its build's source set to blizzard (setup_facts.mjs). */
export function meBnet(): FixtureMe {
  const me = fixture<FixtureMe>('home/me-34.json');
  return {
    ...me,
    characters: me.characters.map((c) => ({ ...c, build: { ...c.build, source: 'blizzard' } })),
  };
}

const json = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});
const envelope = (data: unknown) => json({ ok: true, data, error: null, request_id: 'setup' });
const failure = (status: number) =>
  json({ ok: false, data: null, error: { message: 'setup fixture' }, request_id: 'setup' }, status);

/** The system each state's browser reports (lib/setup/visitor-platform.ts): Windows unless the
 *  state is about another one, so every other state draws the Windows rows its mock draws. */
const SYSTEMS = {
  windows: {
    hints: { platform: 'Windows', mobile: false },
    platform: 'Win32',
    ua: 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0 Safari/537.36',
    touch: 0,
  },
  mac: {
    hints: { platform: 'macOS', mobile: false },
    platform: 'MacIntel',
    ua: 'Mozilla/5.0 (Macintosh; Intel Mac OS X 14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0 Safari/537.36',
    touch: 0,
  },
  mobile: {
    hints: { platform: 'Android', mobile: true },
    platform: 'Linux armv81',
    ua: 'Mozilla/5.0 (Linux; Android 15; Pixel 7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0 Mobile Safari/537.36',
    touch: 5,
  },
  unknown: {
    hints: { platform: 'Linux', mobile: false },
    platform: 'Linux x86_64',
    ua: 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0 Safari/537.36',
    touch: 0,
  },
} as const;

function systemOf(state: SetupPageState): keyof typeof SYSTEMS {
  return state === 'mac' || state === 'mobile' || state === 'unknown' ? state : 'windows';
}

/** The account each signed-in state reads. */
function accountOf(state: SetupPageState): unknown {
  switch (state) {
    case 'none':
    case 'failed':
      return fixture('home/me-new.json');
    case 'importing':
      return {
        ...fixture<FixtureMe>('home/me-new.json'),
        bnet_import: { state: 'running', at: '2026-11-06T09:59:30Z' },
      };
    case 'import-failed':
      return {
        ...fixture<FixtureMe>('home/me-new.json'),
        bnet_import: { state: 'failed', at: '2026-11-06T09:00:00Z' },
      };
    case 'partial':
      return { ...meBnet(), bnet_import: { state: 'partial', at: '2026-11-06T09:00:00Z', unresolved: 2 } };
    case 'bnet':
    case 'companion':
    case 'installed':
    case 'outdated':
      return meBnet();
    case 'claim':
      return {
        ...fixture<FixtureMe>('home/me-34.json'),
        guilds: [{ id: 1, region: 'us', ruleset: 'normal', name: 'Employed Millennials', rank: 'leader' }],
      };
    default:
      return fixture('home/me-34.json');
  }
}

const SIGNED_OUT: readonly SetupPageState[] = ['out', 'mac', 'mobile', 'unknown', 'signin-failed'];
/** A claim page a guild claim's `?next=` names (spec 4.6; G7). */
export const CLAIM_NEXT = '/guild/us/normal/employed-millennials/claim';

/** The address each state opens. */
function pathOf(state: SetupPageState): string {
  if (state === 'signin-failed') return '/setup?signin=failed';
  if (state === 'failed') return '/setup?refreshed=1';
  if (state === 'claim') return `/setup?next=${encodeURIComponent(CLAIM_NEXT)}`;
  return '/setup';
}

/** Routes everything `state` reads and pins the clock. */
export async function routeSetupState(page: Page, state: SetupPageState): Promise<void> {
  await page.clock.setFixedTime(SETUP_NOW);
  await page.addInitScript((system) => {
    Object.defineProperty(navigator, 'userAgentData', { get: () => system.hints });
    Object.defineProperty(navigator, 'platform', { get: () => system.platform });
    Object.defineProperty(navigator, 'userAgent', { get: () => system.ua });
    Object.defineProperty(navigator, 'maxTouchPoints', { get: () => system.touch });
  }, SYSTEMS[systemOf(state)]);
  await page.route('**/dl/companion/latest.json', (route) =>
    route.fulfill(json(fixture('setup/companion-latest.json'))),
  );
  await page.route('**/dl/addon/latest.json', (route) =>
    route.fulfill(json(fixture('setup/addon-latest.json'))),
  );
  if (SIGNED_OUT.includes(state)) {
    await page.route('**/v1/me', (route) => route.fulfill(failure(401)));
    return;
  }
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  if (state === 'loading') {
    await page.route('**/v1/me', () => new Promise<void>(() => {}));
    return;
  }
  if (state === 'error') {
    await page.route('**/v1/me', (route) => route.fulfill(failure(500)));
    return;
  }
  const account = accountOf(state);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(account)));
  const devices = devicesOf(state);
  await page.route('**/v1/devices', (route) => route.fulfill(envelope(devices)));
  if (state === 'pairing') {
    await page.route('**/v1/devices/pair', (route) =>
      route.fulfill(envelope(fixture('setup/pairing-code.json'))),
    );
  }
  if (state === 'failed') {
    // Back from Refresh: the snapshot says the import time did not move (compareSnapshot `error`).
    await page.addInitScript(() => {
      sessionStorage.setItem(
        'fs.refresh.snapshot',
        JSON.stringify({ at: '2026-11-06T09:00:00Z', chars: {} }),
      );
    });
  }
}

/** Opens /setup in `state`, waiting until the island has mounted and its reads answered. */
export async function openSetupState(page: Page, state: SetupPageState): Promise<void> {
  await routeSetupState(page, state);
  await page.goto(pathOf(state));
  await expect(page.locator('astro-island:has([data-testid="setup-grid"]):not([ssr])')).toHaveCount(1);
  if (state === 'loading') return;
  await expect(page.getByTestId('setup-steps')).toBeVisible();
  if (state === 'pairing') {
    await page.getByTestId('pairing-start').click();
    await expect(page.getByTestId('pairing-code')).toBeVisible();
  }
  if (state === 'claim') await expect(page.getByTestId('setup-back-to-claim-notice')).toBeVisible();
  await page.waitForLoadState('networkidle');
}

/**
 * The /setup mocks named, rendered from this checkout's copy modules and fixtures into
 * test-results/parity/mocks/setup (setup_facts.mjs, then ws_setup.py with FS_SETUP_FACTS and
 * FS_MOCK_OUT). Returns that folder.
 */
export function renderSetupMocks(mocks: readonly string[]): string {
  const dir = path.join(PARITY_OUT, 'mocks', 'setup');
  rmSync(dir, { recursive: true, force: true });
  mkdirSync(dir, { recursive: true });
  const facts = path.join(dir, 'setup-facts.json');
  const generators = path.join(REPO_ROOT, 'design/mocks');
  const run = (command: string, args: string[]) => {
    try {
      execFileSync(command, args, {
        cwd: REPO_ROOT,
        env: { ...process.env, FS_SETUP_FACTS: facts, FS_MOCK_OUT: dir },
        encoding: 'utf8',
        stdio: ['ignore', 'pipe', 'pipe'],
      });
    } catch (error) {
      const { stdout = '', stderr = '' } = error as { stdout?: string; stderr?: string };
      throw new Error(`setup mocks: \`${[command, ...args].join(' ')}\` failed\n${stdout}${stderr}`, {
        cause: error,
      });
    }
  };
  run('node', [path.join(generators, 'setup_facts.mjs'), facts]);
  run(process.env.PYTHON ?? 'python3', [path.join(generators, 'ws_setup.py'), ...mocks]);
  return dir;
}
