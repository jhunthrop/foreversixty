// web/tests/e2e/support/connect-page.ts
// /connect in each of its states (setup spec 2026-11-07 section 6) from the committed fixtures the
// setup mocks draw (fixtures/setup/README.md): the lookup is fixtures/setup/connect-device.json and
// the account fixtures/home/me-new.json. Also renders the /connect mocks at test time for the
// parity gate (design/parity.md): setup_facts.mjs reads the copy modules and these fixtures,
// ws_setup.py draws the mocks from those facts.
import { execFileSync } from 'node:child_process';
import { mkdirSync, readFileSync, rmSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page } from '@playwright/test';
import { gotoHydrated } from './hydrated';
import { PARITY_OUT, REPO_ROOT } from './parity';

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../fixtures');

export type ConnectPageState =
  | 'signed-out'
  | 'loading'
  | 'pending'
  | 'connecting'
  | 'success'
  | 'cancelled'
  | 'expired'
  | 'used'
  | 'failed';

export const CONNECT_STATES: readonly ConnectPageState[] = [
  'signed-out',
  'loading',
  'pending',
  'connecting',
  'success',
  'cancelled',
  'expired',
  'used',
  'failed',
];

/** The code every state shows; the mocks draw `displayCode('kq47tm')`. */
export const CONNECT_PATH = '/connect?code=kq47tm';

function fixture(name: string): unknown {
  return JSON.parse(readFileSync(path.join(FIXTURES, name), 'utf8'));
}

const json = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});
const envelope = (data: unknown) => json({ ok: true, data, error: null, request_id: 'connect' });
const failure = (status: number) =>
  json({ ok: false, data: null, error: { message: 'connect fixture' }, request_id: 'connect' }, status);

async function routeSignedOut(page: Page): Promise<void> {
  await page.route('**/v1/me', (route) => route.fulfill(failure(401)));
}

/** Signed in as the fixture account (the header's chip too, through the session cookie). */
async function routeSignedIn(page: Page, me: 'answer' | 'pending' = 'answer'): Promise<void> {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) =>
    me === 'pending' ? new Promise<void>(() => {}) : route.fulfill(envelope(fixture('home/me-new.json'))),
  );
  await page.route('**/v1/devices/connect/lookup**', (route) =>
    route.fulfill(envelope(fixture('setup/connect-device.json'))),
  );
}

/** The approve call's answer: the device, a 409, a network failure, or held until never. */
async function routeApprove(page: Page, answer: 'ok' | 'used' | 'abort' | 'hold'): Promise<void> {
  await page.route('**/v1/devices/connect/approve', (route) => {
    if (answer === 'hold') return new Promise<void>(() => {});
    if (answer === 'abort') return route.abort();
    if (answer === 'used') return route.fulfill(failure(409));
    return route.fulfill(envelope(fixture('setup/connect-device.json')));
  });
}

/** Opens /connect in `state`, waiting until the island shows it. */
export async function openConnectState(page: Page, state: ConnectPageState): Promise<void> {
  if (state === 'signed-out' || state === 'expired') await routeSignedOut(page);
  else await routeSignedIn(page, state === 'loading' ? 'pending' : 'answer');
  if (state === 'loading') {
    await gotoHydrated(page, CONNECT_PATH, 'connect');
    return;
  }
  // No code in the query reads as expired.
  await gotoHydrated(page, state === 'expired' ? '/connect' : CONNECT_PATH, 'connect');
  const approve = { success: 'ok', connecting: 'hold', used: 'used', failed: 'abort' } as const;
  if (state in approve) {
    await routeApprove(page, approve[state as keyof typeof approve]);
    await page.getByTestId('connect-approve').click();
  }
  if (state === 'cancelled') {
    await page.route('**/v1/devices/connect/deny', (route) => route.fulfill({ status: 204 }));
    await page.getByTestId('connect-deny').click();
  }
  const shown = state === 'signed-out' ? 'signedOut' : state === 'connecting' ? 'pending' : state;
  await page.locator(`[data-testid="connect"][data-state="${shown}"]`).waitFor();
}

/**
 * The /connect mocks named, rendered from this checkout's copy modules and fixtures into
 * test-results/parity/mocks/connect (setup_facts.mjs, then ws_setup.py with FS_SETUP_FACTS and
 * FS_MOCK_OUT). Returns that folder.
 */
export function renderConnectMocks(mocks: readonly string[]): string {
  const dir = path.join(PARITY_OUT, 'mocks', 'connect');
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
      throw new Error(`connect mocks: \`${[command, ...args].join(' ')}\` failed\n${stdout}${stderr}`, {
        cause: error,
      });
    }
  };
  run('node', [path.join(generators, 'setup_facts.mjs'), facts]);
  run(process.env.PYTHON ?? 'python3', [path.join(generators, 'ws_setup.py'), ...mocks]);
  return dir;
}
