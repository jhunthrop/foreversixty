// The planner's golden vectors (apps/shared/fixtures/planner-rules.json) are generated from the
// site planner's own rules by web/scripts/gen-planner-rules.mjs, and both apps replay them. This
// fails when the checked-in file differs from a fresh generation: after changing
// web/src/lib/planner, the talent files or the plan.refuse copy, run
// `node web/scripts/gen-planner-rules.mjs` and commit the result.
import { readFileSync } from 'node:fs';
import { join } from 'node:path';
import { describe, expect, it } from 'vitest';
import { plannerGolden, renderGolden } from '../../../scripts/mobile-codegen/planner-golden';

const REPO_ROOT = join(import.meta.dirname, '..', '..', '..', '..');
const FILE = join(REPO_ROOT, 'apps/shared/fixtures/planner-rules.json');

describe('planner golden vectors', () => {
  const golden = plannerGolden(REPO_ROOT);

  it('match the generator', () => {
    expect(readFileSync(FILE, 'utf8'), 'run `node web/scripts/gen-planner-rules.mjs` and commit').toBe(
      renderGolden(golden),
    );
  });

  it('cover every refusal and a legal guide build per FS1 code', () => {
    const reasons = golden.add.map((d) => d.reason ?? '');
    expect(reasons.some((r) => r.startsWith('Tier '))).toBe(true);
    expect(reasons.some((r) => / needs \d+ points? in /.test(r) && !r.startsWith('Tier '))).toBe(true);
    expect(reasons.some((r) => r.includes('is already at'))).toBe(true);
    expect(reasons).toContain(golden.none_left);
    expect(golden.remove.some((d) => d.reason?.endsWith('has no points to remove'))).toBe(true);
    expect(golden.fs1.length).toBeGreaterThan(0);
    for (const vector of golden.fs1) expect(vector.dropped).toEqual([]);
  });

  it('give a level its level minus nine points, at most 51', () => {
    const at = new Map(golden.points_available.map((p) => [p.level, p.points]));
    expect(at.get(9)).toBe(0);
    expect(at.get(10)).toBe(1);
    expect(at.get(60)).toBe(51);
    expect(at.get(70)).toBe(51);
  });
});
