// web/tests/e2e/support/premium-page.ts
// /premium and /premium/checkout in their mocked states (design/specs/2026-11-08-premium-account.md
// sections 5, 6, 14). The accounts are the facts reader's own /v1/me bodies (premium_facts.mjs
// `raw_accounts`, DERIVED from the home fixtures), so the build is served exactly what the mocks
// drew; the clock is the mocks' 2026-11-08T10:00Z. Also renders those mocks at test time for the
// parity gate (design/parity.md): premium_facts.mjs reads the copy modules and the models,
// ws_premium.py draws the mocks from those facts.
import { execFileSync } from 'node:child_process';
import { mkdirSync, readFileSync, rmSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { expect } from './test';
import { PARITY_OUT, REPO_ROOT } from './parity';

export const PREMIUM_NOW = new Date('2026-11-08T10:00:00Z');

export type PremiumPageState = 'out' | 'free' | 'seat' | 'checkout-confirmed' | 'checkout-owned';

const json = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});
const envelope = (data: unknown) => json({ ok: true, data, error: null, request_id: 'premium' });
const failure = (status: number) =>
  json({ ok: false, data: null, error: { message: 'premium fixture' }, request_id: 'premium' }, status);

const ACCOUNT_OF: Record<PremiumPageState, string | null> = {
  out: null,
  free: 'free',
  seat: 'seat',
  'checkout-confirmed': 'premium',
  'checkout-owned': 'seat',
};

const PATH_OF: Record<PremiumPageState, string> = {
  out: '/premium',
  free: '/premium',
  seat: '/premium',
  'checkout-confirmed': '/premium/checkout?plan=premium&interval=yearly&status=success&session_id=cs_test_1',
  'checkout-owned': '/premium/checkout?plan=premium&interval=yearly',
};

interface Facts {
  raw_accounts: Record<string, unknown>;
}

function readFacts(dir: string): Facts {
  return JSON.parse(readFileSync(path.join(dir, 'premium-facts.json'), 'utf8')) as Facts;
}

/** Routes what `state` reads and pins the clock. */
export async function routePremiumState(page: Page, state: PremiumPageState, mocks: string): Promise<void> {
  await page.clock.setFixedTime(PREMIUM_NOW);
  const key = ACCOUNT_OF[state];
  if (key === null) {
    await page.route('**/v1/me', (route) => route.fulfill(failure(401)));
    return;
  }
  const account = readFacts(mocks).raw_accounts[key];
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(account)));
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
}

/** Opens `state`, waiting until its account read has drawn. */
export async function openPremiumState(page: Page, state: PremiumPageState, mocks: string): Promise<void> {
  await routePremiumState(page, state, mocks);
  await page.goto(PATH_OF[state]);
  if (state === 'free' || state === 'seat') await expect(page.getByTestId('premium-answer')).toBeVisible();
  if (state === 'checkout-confirmed') await expect(page.getByTestId('checkout-next')).toBeVisible();
  if (state === 'checkout-owned') await expect(page.getByTestId('premium-owned')).toBeVisible();
  await page.waitForLoadState('networkidle');
}

/** The premium mocks named, rendered from this checkout's copy modules into test-results/parity/mocks/premium. */
export function renderPremiumMocks(mocks: readonly string[]): string {
  const dir = path.join(PARITY_OUT, 'mocks', 'premium');
  rmSync(dir, { recursive: true, force: true });
  mkdirSync(dir, { recursive: true });
  const facts = path.join(dir, 'premium-facts.json');
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
      throw new Error(`premium mocks: \`${[command, ...args].join(' ')}\` failed\n${stdout}${stderr}`, {
        cause: error,
      });
    }
  };
  run('node', [path.join(generators, 'premium_facts.mjs'), facts]);
  run(process.env.PYTHON ?? 'python3', [path.join(generators, 'ws_premium.py'), ...mocks]);
  return dir;
}
