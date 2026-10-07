import { describe, expect, it } from 'vitest';
import { bandEntry } from './load';
import { effectGroupsFor, raidCaptionFor } from './preset-caption';
import { bandPresets, filePresets, parsePresetId, presetLabelFor, presetMeta, selectBand } from './presets';
import type { BisBand, BisFile, BisPresetMeta, Faction } from './types';

function band(level: number, faction: Faction, setDps: number, preset?: 'bare' | 'raid'): BisBand {
  return {
    spec: 'warrior-fury',
    band: level,
    faction,
    race: 'human',
    talents: '',
    talent_points: 0,
    weights: [],
    slots: [],
    set_dps: setDps,
    no_source_count: 0,
    ...(preset === undefined ? {} : { preset }),
  } as BisBand;
}

const RAID_META: BisPresetMeta = {
  label: 'Raid-ready, Phase 1',
  buffs: [{ id: 1, label: 'Buff A' }],
  debuffs: [],
  consumes: [
    { id: 2, label: 'Flask' },
    { id: 3, label: 'Elixir' },
  ],
  notes: 'Invented for the test.',
};

function fileOf(bands: BisBand[], presets?: BisFile['presets']): BisFile {
  return {
    spec: 'warrior-fury',
    build: 'b',
    engine_version: 'e',
    generated_at: '2026-10-07',
    bands,
    presets,
  };
}

/** Today's published shape: no `preset` field anywhere, no `presets` block. */
const LEGACY = fileOf([band(50, 'alliance', 100), band(60, 'alliance', 200), band(60, 'horde', 201)]);

const WITH_RAID = fileOf(
  [
    band(50, 'alliance', 100),
    band(60, 'alliance', 200),
    band(60, 'horde', 201),
    band(60, 'alliance', 300, 'raid'),
    band(60, 'horde', 301, 'raid'),
  ],
  { raid: RAID_META },
);

describe('selectBand on a file without presets', () => {
  it('reads the unlabelled entry whatever the preference', () => {
    expect(selectBand(LEGACY, 60, 'alliance')?.set_dps).toBe(200);
    expect(selectBand(LEGACY, 60, 'alliance', 'raid')?.set_dps).toBe(200);
    expect(selectBand(LEGACY, 60, 'horde', 'bare')?.set_dps).toBe(201);
  });

  it('returns undefined for a band the file lacks', () => {
    expect(selectBand(LEGACY, 40, 'alliance')).toBeUndefined();
  });
});

describe('selectBand on a file with raid entries', () => {
  it('defaults to raid where it exists', () => {
    expect(selectBand(WITH_RAID, 60, 'alliance')?.set_dps).toBe(300);
    expect(selectBand(WITH_RAID, 60, 'horde')?.set_dps).toBe(301);
  });

  it('honours an explicit bare choice', () => {
    expect(selectBand(WITH_RAID, 60, 'alliance', 'bare')?.set_dps).toBe(200);
  });

  it('falls back to bare below level 60, where no raid entry exists', () => {
    expect(selectBand(WITH_RAID, 50, 'alliance')?.set_dps).toBe(100);
  });

  it('treats an entry explicitly marked bare as bare', () => {
    const marked = fileOf([band(60, 'alliance', 5, 'bare'), band(60, 'alliance', 6, 'raid')]);
    expect(selectBand(marked, 60, 'alliance', 'bare')?.set_dps).toBe(5);
  });

  it('is what load.ts bandEntry returns', () => {
    expect(bandEntry(WITH_RAID, 60, 'alliance')?.set_dps).toBe(300);
    expect(bandEntry(WITH_RAID, 60, 'alliance', 'bare')?.set_dps).toBe(200);
  });
});

describe('preset listings', () => {
  it('lists the presets a band offers, headline first', () => {
    expect(bandPresets(WITH_RAID, 60, 'alliance')).toEqual(['raid', 'bare']);
    expect(bandPresets(WITH_RAID, 50, 'alliance')).toEqual(['bare']);
    expect(bandPresets(LEGACY, 60, 'alliance')).toEqual(['bare']);
    expect(bandPresets(LEGACY, 40, 'alliance')).toEqual([]);
  });

  it('lists the presets a file offers', () => {
    expect(filePresets(WITH_RAID)).toEqual(['raid', 'bare']);
    expect(filePresets(LEGACY)).toEqual(['bare']);
  });

  it('parses only known ids from a URL parameter', () => {
    expect(parsePresetId('bare')).toBe('bare');
    expect(parsePresetId('raid')).toBe('raid');
    expect(parsePresetId('nope')).toBeUndefined();
    expect(parsePresetId(null)).toBeUndefined();
  });
});

describe('preset labels and captions', () => {
  it('labels a raid band with the file label and a bare band as Bare', () => {
    expect(presetLabelFor(WITH_RAID, selectBand(WITH_RAID, 60, 'alliance')!)).toBe('Raid-ready, Phase 1');
    expect(presetLabelFor(WITH_RAID, selectBand(WITH_RAID, 60, 'alliance', 'bare')!)).toBe('Bare');
    expect(presetLabelFor(LEGACY, selectBand(LEGACY, 60, 'alliance')!)).toBe('Bare');
  });

  it('exposes the file preset block', () => {
    expect(presetMeta(WITH_RAID, 'raid')).toBe(RAID_META);
    expect(presetMeta(LEGACY, 'raid')).toBeUndefined();
  });

  it('counts the includes in the caption and drops empty groups from the disclosure', () => {
    expect(raidCaptionFor(RAID_META)).toBe('Raid-ready, Phase 1: 1 buff, 0 debuffs and 2 consumables.');
    expect(effectGroupsFor(RAID_META).map((group) => group.heading)).toEqual(['Buffs', 'Consumables']);
  });
});
