// web/src/lib/bis/tank-test-support.ts
// Shared builders for the tank-band unit tests.
import type { PanelViewDeps } from './panel-view';
import type { BisBand, BisFile, BisSlot } from './types';

const TANK_METRICS = {
  dtps: 301.2,
  tmi: 1523,
  chance_of_death: 0.0003,
  tps: 442.7,
  effective_health: 11235,
  dps: 260,
} as const;

const HEAD: BisSlot = {
  slot: 'head',
  item_id: 1,
  item_name: 'Tank Helm',
  source: 'Vendor: Someone',
  source_kind: 'vendor',
  score: 40,
  sim_dps: 61.4,
  dps_delta: 1.3,
  swap_note: 'confirmed by the sim against Spare Helm (id 2): kept the pick, 61.4 vs 60.1 set DPS',
  verified: true,
  alternatives: [
    {
      item_id: 2,
      item_name: 'Spare Helm',
      source: 'x',
      source_kind: 'dungeon',
      dps_delta: -1.3,
      verified: true,
    },
  ],
};

export function tankBand(overrides: Partial<BisBand> = {}): BisBand {
  return {
    spec: 'warrior-protection',
    band: 60,
    faction: 'alliance',
    race: 'human',
    talents: '',
    talent_points: 51,
    weights: [
      { stat: 'stamina', weight: 1, error: 0.01 },
      { stat: 'armor', weight: 0.4, error: 0.02 },
      { stat: 'defense', weight: 1.6, error: 0.05, unit: 'rating', rating_factor: 2.4 },
    ],
    slots: [HEAD],
    set_dps: 260,
    no_source_count: 0,
    new_at_band: [],
    coverage: {},
    weights_run_seconds: 0,
    verify_run_seconds: 0,
    reference_dps_per_point: 0.84,
    role: 'tank',
    metrics: { ...TANK_METRICS },
    score_unit: 'tank_score',
    ...overrides,
  };
}

export function fileOf(bands: BisBand[]): BisFile {
  return {
    spec: 'warrior-protection',
    build: 'test',
    engine_version: 'test',
    generated_at: new Date(0).toISOString(),
    bands,
  };
}

export function bandsOf(file: BisFile): BisBand[] {
  return file.bands;
}

export function depsWithSpec(spec: string): PanelViewDeps {
  return { itemDetails: new Map(), loot: { sources: [], quests: {} }, tooltipFor: () => undefined, spec };
}
