// web/tests/e2e/support/account-page.ts
// /account in the states its parity mocks draw (design/specs/2026-11-08-premium-account.md 7A, 14),
// serving the very /v1/me and /v1/devices bodies the mocks were drawn from: premium_facts.mjs's
// `raw_accounts` and `raw_devices` (the home fixtures with each state's DERIVED fields), read from
// the facts JSON renderAccountMocks writes. `account` is Felsnap (free, one Battle.net character,
// Discord not linked), `account-new` has no character, `account-premium` is Thandrel (own Premium,
// an officer billing Olympus XXVII on Guild, one companion, Discord linked with a channel
// connected). The clock is pinned to the mocks' 2026-11-08T10:00Z. Also renders the account mocks
// at test time for the parity gate (design/parity.md): premium_facts.mjs reads the copy modules,
// ws_premium.py draws the mocks.
import { execFileSync } from 'node:child_process';
import { mkdirSync, readFileSync, rmSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { expect } from './test';
import { PARITY_OUT, REPO_ROOT } from './parity';

/** The mocks' own clock (premium_facts.mjs NOW). */
export const ACCOUNT_NOW = new Date('2026-11-08T10:00:00Z');

export type AccountPageState = 'account' | 'account-new' | 'account-premium';

/** Each state's account and companions, by their keys in the facts. */
const FACTS_KEYS: Readonly<Record<AccountPageState, { account: string; devices: string }>> = {
  account: { account: 'free', devices: 'none' },
  'account-new': { account: 'new', devices: 'none' },
  'account-premium': { account: 'premium', devices: 'one' },
};

const MOCKS_OUT = path.join(PARITY_OUT, 'mocks', 'account');
const FACTS_FILE = path.join(MOCKS_OUT, 'premium-facts.json');

interface RawFacts {
  raw_accounts: Record<string, unknown>;
  raw_devices: Record<string, unknown>;
}

/** The bodies a state serves, from the facts renderAccountMocks wrote. */
function stateBodies(state: AccountPageState): { me: unknown; devices: unknown } {
  const facts = JSON.parse(readFileSync(FACTS_FILE, 'utf8')) as RawFacts;
  const keys = FACTS_KEYS[state];
  return { me: facts.raw_accounts[keys.account], devices: facts.raw_devices[keys.devices] };
}

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'account' }),
});

/** Routes everything `state` reads and pins the clock (after renderAccountMocks has run). */
export async function routeAccountState(page: Page, state: AccountPageState): Promise<void> {
  await page.clock.setFixedTime(ACCOUNT_NOW);
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  const { me, devices } = stateBodies(state);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(me)));
  await page.route('**/v1/devices', (route) => route.fulfill(envelope(devices)));
  // The last-raid line draws nothing until there is a raid: an account with none answers empty.
  await page.route('**/v1/me/coach/recent', (route) =>
    route.fulfill(envelope({ last_raid: null, pending: false })),
  );
}

/** Opens /account in `state`, waiting until the island has mounted and its reads answered. */
export async function openAccountState(page: Page, state: AccountPageState): Promise<void> {
  await routeAccountState(page, state);
  await page.goto('/account');
  await expect(page.getByTestId('account-grid')).toBeVisible();
  await page.waitForLoadState('networkidle');
}

/**
 * The account mocks named, rendered from this checkout's copy modules into
 * test-results/parity/mocks/account (premium_facts.mjs, then ws_premium.py with FS_PREMIUM_FACTS and
 * FS_MOCK_OUT). Returns that folder.
 */
export function renderAccountMocks(mocks: readonly string[]): string {
  const dir = MOCKS_OUT;
  rmSync(dir, { recursive: true, force: true });
  mkdirSync(dir, { recursive: true });
  const facts = FACTS_FILE;
  const generators = path.join(REPO_ROOT, 'design/mocks');
  const run = (command: string, args: string[]) => {
    try {
      execFileSync(command, args, {
        cwd: REPO_ROOT,
        env: { ...process.env, FS_PREMIUM_FACTS: facts, FS_MOCK_OUT: dir },
        encoding: 'utf8',
        stdio: ['ignore', 'pipe', 'pipe'],
      });
    } catch (error) {
      const { stdout = '', stderr = '' } = error as { stdout?: string; stderr?: string };
      throw new Error(`account mocks: \`${[command, ...args].join(' ')}\` failed\n${stdout}${stderr}`, {
        cause: error,
      });
    }
  };
  run('node', [path.join(generators, 'premium_facts.mjs'), facts]);
  run(process.env.PYTHON ?? 'python3', [path.join(generators, 'ws_premium.py'), ...mocks]);
  return dir;
}
