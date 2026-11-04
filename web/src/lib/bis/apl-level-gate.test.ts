// web/src/lib/bis/apl-level-gate.test.ts
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import activeBuild from '../../data/active-build.json';
import { canApplyBetween, conditionAt, knownAtFor, type KnownAt } from './apl-level-gate';
import { loadAplSteps } from './apl-steps';
import { bandTopLevel, loadRotations, rotationEntryFor, rotationLinesFor } from './rotation-view';

const ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
const BUILD = activeBuild.build;
/** Every band a BiS page shows; the guide pages show the last. */
const SHOWN_BANDS = [10, 20, 30, 40, 50, 60];

const known = (id: number) => ({ spellIsKnown: { spellId: { spellId: id } } });
const not = (val: unknown) => ({ not: { val } });
const open = { cmp: { op: 'OpLe', lhs: { currentEnergy: {} }, rhs: { const: { val: '60' } } } };

describe('knownAtFor', () => {
  const families = {
    Claw: [
      { id: 1082, level: 20 },
      { id: 9850, level: 58 },
    ],
    Hemorrhage: [{ id: 16511, level: 30 }],
  };
  const knownAt = knownAtFor(families, new Set(['Hemorrhage']));

  it('knows a trained spell once any rank of its family is learned', () => {
    expect(knownAt(9850, 19)).toBe(false);
    expect(knownAt(9850, 20)).toBe(true);
  });

  it('leaves a talent spell open from its learn level, and unknown below it', () => {
    expect(knownAt(16511, 29)).toBe(false);
    expect(knownAt(16511, 60)).toBeUndefined();
  });

  it('leaves an id in no family open', () => {
    expect(knownAt(424242, 60)).toBeUndefined();
  });
});

describe('conditionAt', () => {
  const knownAt: KnownAt = (id, level) => (id === 1 ? level >= 20 : undefined);

  it('reads no condition as true and any other clause as open', () => {
    expect(conditionAt(undefined, 60, knownAt)).toBe(true);
    expect(conditionAt(open, 60, knownAt)).toBeUndefined();
  });

  it('rules out a "before you learn it" step once it is learned', () => {
    expect(conditionAt(not(known(1)), 10, knownAt)).toBe(true);
    expect(conditionAt(not(known(1)), 60, knownAt)).toBe(false);
    expect(conditionAt({ and: { vals: [not(known(1)), open] } }, 60, knownAt)).toBe(false);
    expect(conditionAt({ and: { vals: [not(known(1)), open] } }, 10, knownAt)).toBeUndefined();
  });

  it('keeps an or open unless every branch is ruled out', () => {
    expect(conditionAt({ or: { vals: [not(known(1)), open] } }, 60, knownAt)).toBeUndefined();
    expect(conditionAt({ or: { vals: [not(known(1)), known(1)] } }, 60, knownAt)).toBe(true);
  });
});

// Round 19: a line scoped to a level range with currentLevel (Discipline's 40-59 step-down)
// is decided by the level alone, so the level 60 card leaves it out and a 40-59 band shows it.
describe('currentLevel', () => {
  const noneKnown: KnownAt = () => undefined;
  const level = (op: string, val: number) => ({
    cmp: { op, lhs: { currentLevel: {} }, rhs: { const: { val: String(val) } } },
  });
  const band40to59 = { and: { vals: [level('OpGe', 40), level('OpLt', 60), open] } };

  it('decides a level comparison at a level, either way round', () => {
    expect(conditionAt(level('OpGe', 40), 45, noneKnown)).toBe(true);
    expect(conditionAt(level('OpLt', 60), 60, noneKnown)).toBe(false);
    expect(
      conditionAt(
        { cmp: { op: 'OpGt', lhs: { const: { val: '59' } }, rhs: { currentLevel: {} } } },
        60,
        noneKnown,
      ),
    ).toBe(false);
    expect(conditionAt(band40to59, 45, noneKnown)).toBeUndefined();
    expect(conditionAt(band40to59, 60, noneKnown)).toBe(false);
  });

  it('shows a 40-59 line on the 40 and 50 bands and not on the 60 card', () => {
    expect(canApplyBetween(band40to59, 60, 60, noneKnown)).toBe(false);
    expect(canApplyBetween(band40to59, 50, 59, noneKnown)).toBe(true);
    expect(canApplyBetween(band40to59, 30, 39, noneKnown)).toBe(false);
  });
});

describe('canApplyBetween', () => {
  const knownAt: KnownAt = (id, level) => (id === 1 ? level >= 24 : undefined);

  it('keeps a step that can fire anywhere in the shown range', () => {
    expect(canApplyBetween(known(1), 20, 29, knownAt)).toBe(true);
    expect(canApplyBetween(not(known(1)), 20, 29, knownAt)).toBe(true);
  });

  it('drops a step its level gate rules out across the whole range', () => {
    expect(canApplyBetween(known(1), 10, 19, knownAt)).toBe(false);
    expect(canApplyBetween(not(known(1)), 30, 39, knownAt)).toBe(false);
  });
});

// --- the sweep: every spec, every level a page shows ----------------------------------------

interface SpellRanksFile {
  classes: Record<string, Record<string, { id: number; level: number }[]>>;
}
interface TalentsFile {
  trees: { talents: { name: string }[] }[];
}

function readJson<T>(...parts: string[]): T {
  return JSON.parse(readFileSync(path.join(ROOT, ...parts), 'utf8')) as T;
}

const hasBuild = existsSync(path.join(ROOT, 'data/builds', BUILD, 'addon-data.json'));
const specs = hasBuild
  ? Object.keys(
      readJson<{ rotations: Record<string, unknown> }>('data/builds', BUILD, 'addon-data.json').rotations,
    )
  : [];

describe.skipIf(!hasBuild)('no shown rotation row is ruled out by its level gate', () => {
  const spellRanks = hasBuild ? readJson<SpellRanksFile>('data/builds', BUILD, 'spellranks.json') : undefined;

  it.each(specs)('%s', (spec) => {
    const [classSlug] = spec.split('-');
    const families = spellRanks?.classes[classSlug] ?? {};
    const talentFile = path.join('data/builds', BUILD, 'talents', `${classSlug}.json`);
    const talentNames = existsSync(path.join(ROOT, talentFile))
      ? new Set(readJson<TalentsFile>(talentFile).trees.flatMap((tree) => tree.talents.map((t) => t.name)))
      : new Set<string>();
    const knownAt = knownAtFor(families, talentNames);
    const entries = loadRotations(BUILD, spec) ?? [];
    const offenders: string[] = [];
    for (const band of SHOWN_BANDS) {
      const top = bandTopLevel(band);
      const entry = rotationEntryFor(entries, top);
      if (entry === undefined) continue;
      const steps = loadAplSteps(spec, families, entry.level) ?? [];
      const ruledOut = (spellId: number) => {
        const own = steps.filter((step) => step.resolvedSpellId === spellId);
        return (
          own.length > 0 && own.every((step) => !canApplyBetween(step.condition, entry.level, top, knownAt))
        );
      };
      const shown = rotationLinesFor(entry, BUILD, classSlug, spec, top);
      offenders.push(...shown.filter((row) => ruledOut(row.spellId)).map((row) => `${row.name} at ${top}`));
    }
    expect(offenders).toEqual([]);
  });

  it('drops the pre-Cat-Form Moonfire and Wrath from the level 60 Feral list', () => {
    const entry = rotationEntryFor(loadRotations(BUILD, 'druid-feral') ?? [], 60);
    if (entry === undefined) throw new Error('druid-feral has no level 60 rotation');
    const names = rotationLinesFor(entry, BUILD, 'druid', 'druid-feral', 60).map((row) => row.name);
    expect(names).not.toContain('Moonfire');
    expect(names).not.toContain('Wrath');
    expect(names).toEqual(expect.arrayContaining(['Shred', 'Rip', 'Ferocious Bite']));
  });

  it('keeps them at level 10, before Cat Form', () => {
    const entry = rotationEntryFor(loadRotations(BUILD, 'druid-feral') ?? [], bandTopLevel(10));
    if (entry === undefined) throw new Error('druid-feral has no level 10 rotation');
    const names = rotationLinesFor(entry, BUILD, 'druid', 'druid-feral', bandTopLevel(10)).map(
      (row) => row.name,
    );
    expect(names).toEqual(expect.arrayContaining(['Moonfire', 'Wrath']));
  });
});
