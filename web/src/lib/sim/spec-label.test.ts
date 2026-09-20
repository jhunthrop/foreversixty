import { readFile } from 'node:fs/promises';
import path from 'node:path';
import { describe, expect, it } from 'vitest';
import { SPECS } from './specs';
import {
  classOfSpec,
  dpsSpecs,
  isSimulatedSpec,
  nonDpsSpecs,
  specDisplayName,
  specLabel,
  specRow,
} from './spec-label';

const CURATED = path.resolve(import.meta.dirname, '../../../../data/curated/specs.json');

describe('the generated spec list', () => {
  it('is exactly what data/curated/specs.json holds, so the data lane is the only author', async () => {
    const curated = JSON.parse(await readFile(CURATED, 'utf8')) as unknown;
    expect(JSON.parse(JSON.stringify(SPECS))).toEqual(curated);
  });

  it('names every spec as <class-slug>-<spec-slug>', () => {
    for (const row of SPECS) {
      expect(row.spec).toBe(`${row.class_slug}-${row.spec_slug}`);
      expect(row.spec).toMatch(/^[a-z]+(-[a-z]+)+$/);
    }
  });

  it('carries the two launch specs the engine can actually build', () => {
    expect(specRow('warrior-fury')).not.toBeNull();
    expect(specRow('mage-frost')).not.toBeNull();
  });
});

describe('specLabel', () => {
  it('is the spec and its class, because "Fury" alone is not a spec', () => {
    expect(specLabel('warrior-fury')).toBe('Fury Warrior');
    expect(specLabel('mage-frost')).toBe('Frost Mage');
    expect(specLabel('hunter-beast-mastery')).toBe('Beast Mastery Hunter');
  });

  it('degrades to the slug itself for a spec the list does not know', () => {
    expect(specLabel('warrior-gladiator')).toBe('warrior-gladiator');
    expect(classOfSpec('warrior-gladiator')).toBe('');
    expect(specRow('warrior-gladiator')).toBeNull();
  });
});

describe('specDisplayName', () => {
  // Fix round 1, Finding 1: extracted from SettingsBar.svelte's own rotationName, which
  // duplicated this exact fallback. Shared now by the settings bar and the rotation card.
  it('is the bare spec name, never the class joined in -- that is specLabel’s job', () => {
    expect(specDisplayName('warrior-fury')).toBe('Fury');
    expect(specDisplayName('mage-frost')).toBe('Frost');
  });

  it('degrades to the slug itself for a spec the list does not know, not a raised-case guess', () => {
    expect(specDisplayName('warrior-gladiator')).toBe('warrior-gladiator');
  });
});

describe('dpsSpecs', () => {
  it('is every dps spec and no tank or healer, since launch is Quick Sim for dps only', () => {
    const rows = dpsSpecs();
    expect(rows.length).toBeGreaterThan(0);
    expect(rows.every((row) => row.role === 'dps')).toBe(true);
    expect(rows.map((row) => row.spec)).toContain('warrior-fury');
    expect(rows.map((row) => row.spec)).not.toContain('warrior-protection');
    expect(rows.map((row) => row.spec)).not.toContain('druid-restoration');
  });

  it('is a subset of the generated list and never a second copy of it', () => {
    expect(dpsSpecs().length).toBeLessThan(SPECS.length);
    for (const row of dpsSpecs()) expect(SPECS).toContain(row);
  });
});

describe('nonDpsSpecs', () => {
  it('is every healer and tank spec, the seven launch scope excludes', () => {
    const rows = nonDpsSpecs();
    expect(rows).toHaveLength(7);
    expect(rows.every((row) => row.role !== 'dps')).toBe(true);
    expect(rows.map((row) => row.spec)).toContain('warrior-protection');
    expect(rows.map((row) => row.spec)).toContain('druid-restoration');
    expect(rows.map((row) => row.spec)).not.toContain('warrior-fury');
  });

  it('is the exact complement of dpsSpecs(): together they are every spec, exactly once', () => {
    const combined = [...dpsSpecs(), ...nonDpsSpecs()].map((row) => row.spec).sort();
    expect(combined).toEqual([...SPECS].map((row) => row.spec).sort());
  });
});

describe('isSimulatedSpec', () => {
  // Task 3: every rotation claim, run gate and engine-error translation reads this one
  // predicate. Checked against all 27 canonical specs, not a handful, so a future data
  // change that moves a spec's role cannot silently disagree with dpsSpecs()/nonDpsSpecs().
  it('is true for every dps spec and false for every healer or tank spec', () => {
    for (const row of dpsSpecs()) expect(isSimulatedSpec(row.spec)).toBe(true);
    for (const row of nonDpsSpecs()) expect(isSimulatedSpec(row.spec)).toBe(false);
  });

  it('fails closed for a spec string the canonical list does not carry', () => {
    expect(isSimulatedSpec('warrior-gladiator')).toBe(false);
    expect(isSimulatedSpec('')).toBe(false);
  });
});
