// The AI-slop copy guard (owner 2026-11-09): every pattern in tools/slop/patterns.txt over every
// copy source in tools/slop/sources.txt outside api/ (api/internal/slopguard covers that tree).
// A real game name or a legal phrase goes in tools/slop/allowlist.txt, never a looser pattern.
import { execFileSync } from 'node:child_process';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';

const SCAN = fileURLToPath(new URL('../../tools/slop/scan.mjs', import.meta.url));
const SCAN_TIMEOUT_MS = 60_000;

function scan(...args: string[]): string {
  return execFileSync(process.execPath, [SCAN, ...args], { encoding: 'utf8' });
}

/** One line per tell the guard must keep catching, with the pattern id it trips. */
const SAMPLES: [string, string][] = [
  ['em-dash', 'Threat — not modeled yet.'],
  ['is-the-colon', 'Fury is the dual-wield warrior: Berserker Stance all fight.'],
  ['nudge-prefix', 'Next: spend your 1 talent point.'],
  ['here-is-what', 'Here is what your log shows for Ragnaros.'],
  ['lets', "Let's get you raid ready."],
  ['buzzword', 'Unlock your class with seamless sims.'],
  ['stock-promise', 'Everything you need, all in one place.'],
  ['simply-just', 'Simply paste your export.'],
  ['exclamation', "'You are all set!'"],
  ['grade-hype', 'Fury is graded D this week'],
  ['explain-obvious', 'This page shows the defaults.'],
  ['whether-youre', "Whether you're leveling or raiding"],
  ['not-just', 'More than just a BiS list'],
];

describe('slop copy guard', () => {
  it('finds no slop pattern in the copy outside the allowlist', { timeout: SCAN_TIMEOUT_MS }, () => {
    expect(scan('--skip', 'api/')).toBe('');
  });

  it.each(SAMPLES)('still catches %s', (id, text) => {
    expect(scan('--probe', text).split('\n')).toContain(id);
  });

  it('lets an empty-cell dash, a real racial and plain copy through', () => {
    expect(scan('--probe', "const NO_FIGURE = '—';")).toBe('');
    expect(scan('--probe', '2 talent points unspent')).toBe('');
    expect(scan('--probe', 'Upgrade: Barbaric Iron Helm, Blacksmithing')).toBe('');
  });
});
