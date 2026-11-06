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
