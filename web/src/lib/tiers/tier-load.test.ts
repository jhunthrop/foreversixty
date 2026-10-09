// web/src/lib/tiers/tier-load.test.ts
import { describe, expect, it } from 'vitest';
import { bisSourceBuild } from '../../../scripts/bis-source-build.mjs';
import { loadBisFile, readSpecCatalog } from '../bis/load';
import activeBuild from '../../data/active-build.json';
import { band, bisFile, catalogEntry } from './tier-test-support';
import { buildTierData, loadTierData } from './tier-load';

const FURY = catalogEntry('warrior', 'fury', 'Fury');
const ARMS = catalogEntry('warrior', 'arms', 'Arms');
const PROT = catalogEntry('paladin', 'protection', 'Protection', 'tank');

function filesFor(): Record<string, ReturnType<typeof bisFile>> {
  const tank = { dtps: 300, tmi: 1, chance_of_death: 0, tps: 600, effective_health: 30000, dps: 1 };
  return {
    [FURY.spec]: bisFile(
      FURY,
      [
        band(FURY, { set_dps: 100 }),
        band(FURY, { set_dps: 90, faction: 'horde', race: 'orc' }),
        band(FURY, { set_dps: 10, preset: 'bare' }),
      ],
      '2026-10-08T10:00:00Z',
    ),
    [ARMS.spec]: bisFile(
      ARMS,
      [band(ARMS, { set_dps: 95 }), band(ARMS, { set_dps: 99, faction: 'horde', race: 'troll' })],
      '2026-10-09T10:00:00Z',
    ),
    [PROT.spec]: {
      ...bisFile(
        PROT,
        [band(PROT, { metrics: tank }), band(PROT, { metrics: tank, faction: 'horde' })],
        '2026-10-07T10:00:00Z',
      ),
    },
  };
}

const deps = (files: Record<string, ReturnType<typeof bisFile>>) => ({
  catalog: [FURY, ARMS, PROT],
  loadFile: (spec: string) => files[spec] ?? null,
  readBossSwingSeconds: () => 2,
});

describe('buildTierData', () => {
  it('ranks each faction on its own numbers and names the race per faction', () => {
    const data = buildTierData(deps(filesFor()));
    expect(data.lists.alliance.dps.map((r) => r.name)).toEqual(['Fury', 'Arms']);
    expect(data.lists.horde.dps.map((r) => r.name)).toEqual(['Arms', 'Fury']);
    expect(data.lists.alliance.dps[0]!.race).toBe('Night Elf');
    expect(data.lists.horde.dps[0]!.race).toBe('Troll');
  });

  it('reads the raid preset, never the bare entry', () => {
    const data = buildTierData(deps(filesFor()));
    expect(data.lists.alliance.dps[0]!.metric).toBe(100);
  });

  it('stamps the newest generated_at and the preset label', () => {
    const data = buildTierData(deps(filesFor()));
    expect(data.updated?.toISOString()).toBe('2026-10-09T10:00:00.000Z');
    expect(data.presetLabel).toBe('Raid-ready, Phase 1');
    expect(data.build).toBe('1.60.1.1');
    expect(data.bossSwingSeconds).toBe(2);
  });

  it('skips a spec with no published file and leaves the others', () => {
    const files = filesFor();
    delete files[ARMS.spec];
    const data = buildTierData(deps(files));
    expect(data.lists.alliance.dps.map((r) => r.name)).toEqual(['Fury']);
  });

  it('gives empty lists and no stamp when nothing is published', () => {
    const data = buildTierData(deps({}));
    expect(data.lists.alliance).toEqual({ dps: [], tank: [], healer: [] });
    expect(data.updated).toBeNull();
    expect(data.presetLabel).toBeNull();
  });
});

describe('loadTierData against the repository data', () => {
  it('reads through the BiS loaders, so a build with no ranked files uses the newest ranked build', () => {
    const data = loadTierData();
    const specs = readSpecCatalog();
    const publishing = specs.filter((s) => loadBisFile(s.spec, activeBuild.build) !== null);
    const listed =
      data.lists.alliance.dps.length + data.lists.alliance.tank.length + data.lists.alliance.healer.length;
    expect(listed).toBe(publishing.length);
    const sample = loadBisFile(publishing[0]!.spec, activeBuild.build);
    const REPO_ROOT = new URL('../../../../', import.meta.url).pathname;
    expect(sample?.build).toBe(bisSourceBuild(REPO_ROOT, activeBuild.build));
    expect(data.bossSwingSeconds).toBeGreaterThan(0);
  });
});
