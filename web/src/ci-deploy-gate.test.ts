// web/src/ci-deploy-gate.test.ts
// Pins web.yml's deploy gate: `deploy` waits for `real-engine`, a failed real-engine blocks it,
// and a real-engine that real-engine-scope skipped (the change cannot reach the wasm) does not.
// The `if:` is evaluated, not string-matched, over every result the two jobs can report.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { parse } from 'yaml';

const WORKFLOW = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  '../../.github/workflows/web.yml',
);

interface Job {
  needs?: string | string[];
  if?: string;
}

const jobs = (parse(readFileSync(WORKFLOW, 'utf8')) as { jobs: Record<string, Job> }).jobs;

type JobResult = 'success' | 'failure' | 'cancelled' | 'skipped';

interface Run {
  event: string;
  ref: string;
  cancelled: boolean;
  results: Record<string, JobResult>;
  outputs: Record<string, Record<string, string>>;
}

/**
 * Evaluates a GitHub Actions `if:` over the subset of the expression language deploy's gate
 * uses (status functions, needs.<job>.result / .outputs.<name>, github.ref / event_name,
 * ==, !=, &&, ||, !, parentheses, string literals). Anything else throws, so a gate rewritten
 * past this subset fails here loudly instead of passing by accident.
 */
function evaluate(expression: string, needs: string[], run: Run): boolean {
  const lookups: Record<string, string> = { 'github.ref': run.ref, 'github.event_name': run.event };
  const succeeded = !run.cancelled && needs.every((job) => run.results[job] === 'success');
  // GitHub ANDs an implicit success() onto an if: that calls no status function.
  const gated = /\b(always|success|failure|cancelled)\(\)/.test(expression)
    ? expression
    : `success() && (${expression})`;
  const js = gated
    .replace(/\s+/g, ' ')
    .replace(/\balways\(\)/g, 'true')
    .replace(/\bsuccess\(\)/g, String(succeeded))
    .replace(/\bcancelled\(\)/g, String(run.cancelled))
    .replace(/\bneeds\.([a-z-]+)\.result\b/g, (_, job: string) =>
      JSON.stringify(run.results[job] ?? 'skipped'),
    )
    .replace(/\bneeds\.([a-z-]+)\.outputs\.([a-z-]+)\b/g, (_, job: string, name: string) =>
      JSON.stringify(run.outputs[job]?.[name] ?? ''),
    )
    .replace(/\bgithub\.(ref|event_name)\b/g, (match) => JSON.stringify(lookups[match]))
    .replace(/==/g, '===');
  if (/[a-z_]+\s*\(|\bneeds\.|\bgithub\./i.test(js.replace(/"[^"]*"|'[^']*'/g, ''))) {
    throw new Error(`deploy's if: uses an expression this test does not model: ${js}`);
  }
  return Boolean(new Function(`return (${js});`)());
}

function deployRuns(overrides: Partial<Run> & { results: Record<string, JobResult> }): boolean {
  const run: Run = {
    event: 'push',
    ref: 'refs/heads/main',
    cancelled: false,
    outputs: { 'real-engine-scope': { run: 'true' } },
    ...overrides,
  };
  return evaluate(jobs.deploy.if ?? 'success()', [jobs.deploy.needs ?? []].flat(), run);
}

const GREEN = { build: 'success', 'real-engine-scope': 'success', 'real-engine': 'success' } as const;

describe("web.yml's deploy gate", () => {
  it('deploy needs build, real-engine-scope and real-engine', () => {
    expect([jobs.deploy.needs].flat()).toEqual(
      expect.arrayContaining(['build', 'real-engine-scope', 'real-engine']),
    );
  });

  it('deploys a green push to main', () => {
    expect(deployRuns({ results: GREEN })).toBe(true);
  });

  it('deploys a workflow_dispatch on main, and nothing off main or on a pull request', () => {
    expect(deployRuns({ event: 'workflow_dispatch', results: GREEN })).toBe(true);
    expect(deployRuns({ ref: 'refs/heads/feature', results: GREEN })).toBe(false);
    // A batch landing's checks run on land/<n> (tools/land.sh); only main's run deploys it.
    expect(deployRuns({ ref: 'refs/heads/land/7', results: GREEN })).toBe(false);
    expect(deployRuns({ event: 'pull_request', results: GREEN })).toBe(false);
  });

  it('a failed or cancelled real-engine blocks deploy', () => {
    expect(deployRuns({ results: { ...GREEN, 'real-engine': 'failure' } })).toBe(false);
    expect(deployRuns({ results: { ...GREEN, 'real-engine': 'cancelled' } })).toBe(false);
  });

  it('a real-engine the scope skipped does not block deploy', () => {
    expect(
      deployRuns({
        results: { ...GREEN, 'real-engine': 'skipped' },
        outputs: { 'real-engine-scope': { run: 'false' } },
      }),
    ).toBe(true);
  });

  it('a real-engine skipped for any other reason blocks deploy', () => {
    // The scope job failed, so real-engine never ran: nothing proved the wasm.
    expect(
      deployRuns({
        results: { ...GREEN, 'real-engine-scope': 'failure', 'real-engine': 'skipped' },
        outputs: {},
      }),
    ).toBe(false);
  });

  it('a failed build, or a cancelled run, blocks deploy', () => {
    expect(deployRuns({ results: { ...GREEN, build: 'failure', 'real-engine': 'skipped' } })).toBe(false);
    expect(deployRuns({ cancelled: true, results: GREEN })).toBe(false);
  });
});
