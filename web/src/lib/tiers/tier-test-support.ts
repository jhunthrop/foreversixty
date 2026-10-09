// web/src/lib/tiers/tier-test-support.ts
// Shared builders for the tier list unit tests.
import type { BisBand, BisFile, BisRole, Faction, SpecCatalogEntry } from '../bis/types';
import type { TierInput } from './tier-list';

export function catalogEntry(
  classSlug: string,
  specSlug: string,
  name: string,
  role: BisRole = 'dps',
): SpecCatalogEntry {
  return {
    spec: `${classSlug}-${specSlug}`,
    class_slug: classSlug,
    spec_slug: specSlug,
    name,
    role,
    tree_index: 0,
    reference_stat: 'attack_power',
    weight_stats: [],
  };
}

export function band(
  entry: SpecCatalogEntry,
  overrides: Partial<BisBand> & { faction?: Faction } = {},
): BisBand {
  return {
    spec: entry.spec,
    band: 60,
    preset: 'raid',
    faction: 'alliance',
    race: 'night-elf',
    talents: '0-0-0',
    talent_points: 51,
    weights: [],
    slots: [],
    set_dps: 100,
    no_source_count: 0,
    coverage: {},
    new_at_band: [],
    weights_run_seconds: 1,
    verify_run_seconds: 1,
    role: entry.role as BisRole,
    ...overrides,
  };
}

export function dpsInput(classSlug: string, specSlug: string, name: string, setDps: number): TierInput {
  const entry = catalogEntry(classSlug, specSlug, name);
  return { entry, band: band(entry, { set_dps: setDps }) };
}

export function tankInput(
  classSlug: string,
  specSlug: string,
  name: string,
  metrics: { dtps: number; effective_health: number; tps: number },
): TierInput {
  const entry = catalogEntry(classSlug, specSlug, name, 'tank');
  return {
    entry,
    band: band(entry, { metrics: { ...metrics, tmi: 1, chance_of_death: 0, dps: 1 } }),
  };
}

export function bisFile(entry: SpecCatalogEntry, bands: BisBand[], generatedAt: string): BisFile {
  return {
    spec: entry.spec,
    build: '1.60.1.1',
    engine_version: 'e1',
    generated_at: generatedAt,
    bands,
    presets: { raid: { label: 'Raid-ready, Phase 1', buffs: [], debuffs: [], consumes: [] } },
  };
}
