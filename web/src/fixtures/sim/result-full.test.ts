// The real published Fury run the saved result page is built against, held to the spec's
// accuracy checks (design/specs/2026-11-04-sim-results.md section 10). Regenerate it with
// forever-sim from design/mocks/data/simview/fury-request.json; these checks say whether the
// regenerated run still agrees with itself.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { describe, expect, it } from 'vitest';
import { resolveActionName, type ActionNames } from '../../lib/sim/action-names';
import { repoRoot } from '../../lib/data/build-files';
import type { SampleOutcome, SimResult } from '../../lib/sim/types';
import activeBuild from '../../data/active-build.json';
import fixture from './result-full.json';

const result = fixture as unknown as SimResult;
const DATA = path.join(repoRoot(), 'data');
const readJson = <T>(file: string): T => JSON.parse(readFileSync(path.join(DATA, file), 'utf8')) as T;

const BUILD = activeBuild.build;
const abilities = result.summary.damage_done[0]!.abilities;
const ability = (key: string) => {
  const row = abilities.find((a) => a.name === key);
  if (row === undefined) throw new Error(`no ability row ${key}`);
  return row;
};

/** Every action key a section of the page draws. */
function actionKeys(): string[] {
  const keys = new Set<string>();
  for (const row of result.sample ?? []) keys.add(row.action);
  for (const aura of result.sample_auras ?? []) keys.add(aura.action);
  for (const aura of result.summary.auras) keys.add(aura.name);
  for (const row of abilities) keys.add(row.name);
  for (const row of result.summary.casts) keys.add(row.spell_name);
  for (const resource of result.resource_sources ?? []) {
    for (const row of [...resource.gained, ...resource.spent]) keys.add(row.action);
  }
  return [...keys].sort();
}

const spellLabel = (key: string) => /^spell:(\d+)/.exec(key)?.[1];

describe('the fixture is a pure SimResult from one real run', () => {
  it('carries the fixed identity and no note', () => {
    expect(result.sim_id).toBe('simfullfixt1');
    expect(result.iterations_run).toBe(10000);
    expect(fixture).not.toHaveProperty('_fixture_note');
  });

  it('pins the ability table the spec calculates by hand', () => {
    const execute = ability('spell:20662');
    expect(execute.total).toBe(34337);
    expect(execute.hits).toBe(23);
    expect(execute.crits).toBe(12);
    const mainHand = ability('other:attack/1');
    expect(mainHand.total).toBe(23802);
    expect(mainHand.hits).toBe(47);
    const deepWounds = ability('spell:12867');
    expect(deepWounds.ticks).toBe(59);
    expect(Math.round(deepWounds.total / deepWounds.ticks)).toBe(141);
  });
});

describe('the DPS spread', () => {
  it('sums to the iterations run, in contiguous bins that cover the run', () => {
    const hist = result.dps_histogram!;
    expect(hist.counts.reduce((a, b) => a + b, 0)).toBe(result.iterations_run);
    expect(hist.start).toBeLessThanOrEqual(result.dps.min);
    expect(hist.start + (hist.counts.length - 1) * hist.bin_dps + hist.bin_dps).toBeGreaterThanOrEqual(
      result.dps.max,
    );
  });
});

describe('the resource sources', () => {
  const rage = result.resource_sources!.find((r) => r.resource === 'rage')!;
  const total = result.summary.resources[0]!;
  const sum = (rows: { amount: number }[]) => rows.reduce((a, r) => a + r.amount, 0);

  it('sum to what the summary gained and spent', () => {
    expect(Math.abs(sum(rage.gained) - total.gained)).toBeLessThanOrEqual(0.5);
    expect(Math.abs(sum(rage.spent) - total.spent)).toBeLessThanOrEqual(0.5);
  });

  it('book Bloodrage as a gain (ruling R5)', () => {
    expect(rage.gained.map((r) => r.action)).toContain('spell:2687');
    expect(rage.spent.map((r) => r.action)).not.toContain('spell:2687');
  });
});

describe('every action key resolves to a name and an icon (ruling R1)', () => {
  const spells = readJson<{ id: number; name: string }[]>(`builds/${BUILD}/spells.json`);
  const engineNames = readJson<{ spells: Record<string, { name: string }> }>(
    'curated/engine-spell-names.json',
  );
  const names: ActionNames = {
    spell: {
      ...Object.fromEntries(spells.map((s) => [String(s.id), s.name])),
      ...Object.fromEntries(Object.entries(engineNames.spells).map(([id, row]) => [id, row.name])),
    },
    item: { '13442': 'Mighty Rage Potion' },
  };

  // spellicons.json is regenerated from the curated list, so an id on that list has its
  // client icon on the next `python -m pipeline icons`; each was checked against the client's
  // SpellMisc and ManifestInterfaceData for this build by hand.
  const spellIcons = readJson<Record<string, string>>(`builds/${BUILD}/spellicons.json`);
  const curatedIcons = readJson<{ spell_ids: number[]; icon_rows: Record<string, number> }>(
    'curated/extra-spell-icons.json',
  );
  const hasSpellIcon = (id: string) =>
    id in spellIcons || curatedIcons.spell_ids.includes(Number(id)) || id in curatedIcons.icon_rows;

  const simBuffs = readJson<{ entries: Record<string, { name: string; icon: string }> }>(
    `builds/${BUILD}/simbuffs.json`,
  );
  const weaponSlots = new Set(result.request.character.gear?.map((g) => g.slot));

  /** A swing draws the weapon, as the Attack button does; the refund draws the page's glyph. */
  const nonSpellIcon: Record<string, boolean> = {
    'other:attack/1': weaponSlots.has('main_hand'),
    'other:attack/2': weaponSlots.has('off_hand'),
    'other:attack/3': weaponSlots.has('main_hand'),
    'other:refund': true,
  };

  it.each(actionKeys())('%s', (key) => {
    const name = resolveActionName(key, names);
    expect(name).not.toMatch(/unnamed/i);
    expect(name).not.toBe(key);

    const spell = spellLabel(key);
    if (spell !== undefined) {
      expect(hasSpellIcon(spell), `no icon for ${key}`).toBe(true);
    } else if (key.startsWith('item:')) {
      expect(Object.values(simBuffs.entries).map((e) => e.name)).toContain(name);
    } else {
      expect(nonSpellIcon[key], `no icon for ${key}`).toBe(true);
    }
  });

  it('never draws the engine extra attack counter it files under Thrash', () => {
    expect(actionKeys()).not.toContain('spell:21919');
  });

  it('names a Windfury source as the buff the request asked for', () => {
    expect(result.extra_attack_sources!.map((s) => s.source)).toEqual(['windfury_totem']);
    expect(simBuffs.entries.windfury_totem).toBeDefined();
  });
});

describe('the sample agrees with the summary (section 10 check 8)', () => {
  const rows = result.sample!;
  const tally = (key: string, outcomes: SampleOutcome[]) =>
    rows.filter((r) => r.action === key && r.outcome !== undefined && outcomes.includes(r.outcome)).length;

  const POINTS = 15;
  const direct = ['spell:20662', 'spell:23894', 'spell:1680', 'other:attack/1', 'other:attack/2'];

  it.each(direct)('%s lands its crit rate and its miss rate within 15 points', (key) => {
    const row = ability(key);
    const landed: SampleOutcome[] = ['hit', 'crit', 'glance', 'block'];
    const sampledLanded = tally(key, landed);
    const sampledMissed = tally(key, ['miss']);
    const summaryMissed = row.misses?.MISS ?? 0;
    const attempts = row.hits + summaryMissed + (row.misses?.DODGE ?? 0) + (row.misses?.PARRY ?? 0);
    const sampledAttempts = sampledLanded + sampledMissed + tally(key, ['dodge', 'parry']);

    expect(
      Math.abs((100 * tally(key, ['crit'])) / sampledLanded - (100 * row.crits) / row.hits),
    ).toBeLessThan(POINTS);
    expect(Math.abs((100 * sampledMissed) / sampledAttempts - (100 * summaryMissed) / attempts)).toBeLessThan(
      POINTS,
    );
  });

  it('counts Deep Wounds ticks as ticks, and applications as casts', () => {
    const ticks = tally('spell:12867', ['tick', 'tick_crit']);
    const summary = ability('spell:12867');
    expect(Math.abs(ticks - summary.ticks)).toBeLessThan(0.15 * summary.ticks);
    expect(
      Math.abs((100 * tally('spell:12867', ['tick_crit'])) / ticks - (100 * summary.crits) / summary.ticks),
    ).toBeLessThan(POINTS);
    const applications = rows.filter((r) => r.action === 'spell:12867' && r.outcome === undefined);
    expect(applications.length).toBeGreaterThan(ticks);
    expect(applications.every((r) => r.amount === undefined)).toBe(true);
  });

  it('agrees a debuff uptime with its periodic row', () => {
    const TICK_MS = 3000;
    const track = result.summary.auras.find((a) => a.name === 'spell:12867' && a.type === 'DEBUFF')!;
    expect(track.target_guid).toBe('target-1');
    expect(track.target_name).toBe('Target 1');
    expect(Math.abs(track.uptime_ms - ability('spell:12867').ticks * TICK_MS)).toBeLessThan(
      0.02 * track.uptime_ms,
    );
  });

  it('keeps every sample aura window on the sample run’s own clock', () => {
    const end = result.sample_duration_ms!;
    for (const aura of result.sample_auras!) {
      for (const [from, to] of aura.segments) {
        expect(from).toBeGreaterThanOrEqual(0);
        expect(to).toBeGreaterThan(from);
        expect(to).toBeLessThanOrEqual(end);
      }
    }
    const sunder = result.sample_auras!.find((a) => a.action === 'spell:11597')!;
    expect(sunder.on).toBe('target');
    expect(sunder.stacks!.at(-1)![1]).toBe(5);
  });
});

describe('the sheet', () => {
  it('fills attack power and an explicit no haste (ruling R3)', () => {
    expect(result.final_stats!.attack_power).toBeGreaterThan(0);
    expect(result.final_stats!.melee_haste_pct).toBe(0);
    expect(result.final_stats).not.toHaveProperty('spell_power');
  });

  it('shows the target at its level with less armor after debuffs (ruling R4)', () => {
    const target = result.target_stats!;
    expect(target.level).toBe(63);
    expect(target.armor_after_debuffs).toBeLessThan(target.armor!);
  });
});
