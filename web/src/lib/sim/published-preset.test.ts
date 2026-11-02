// web/src/lib/sim/published-preset.test.ts
// Owner ruling, round 6: below 60 the page's default buffs match the published number at the
// character's band. Bands that publish bare get bare defaults (the class kit); Raid-buffed
// stays a user choice. Which preset a band publishes comes from the ranker's own rule,
// generated into sim/request/published-presets.json, and is checked here against the
// published BiS files themselves.
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { beforeAll, describe, expect, it } from 'vitest';
import activeBuild from '../../data/active-build.json';
import { selectBand, presetOf } from '../bis/presets';
import type { BisFile } from '../bis/types';
import { loadPresetSets } from './preset-sets';
import {
  BUFF_PRESETS,
  defaultSettings,
  publishedPresetAt,
  RAID_BUFFS,
  withPreset,
  withSpecForPreset,
} from './settings';

beforeAll(() => loadPresetSets());

const REPO = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
const GENERATED = path.join(REPO, 'sim', 'request', 'published-presets.json');
const RAID_SETS = path.join(REPO, 'sim', 'request', 'raid-consumes.json');
const BIS_DIR = path.join(REPO, 'data', 'builds', activeBuild.build, 'bis');
const BANDS = [20, 30, 40, 50, 60] as const;

interface LevelSet {
  from_level: number;
  buffs: string[];
  consumes: string[];
}
interface Generated {
  published: { from_level: number; preset: 'bare' | 'raid' }[];
  bare: Record<string, LevelSet[]>;
}

const generated = JSON.parse(readFileSync(GENERATED, 'utf8')) as Generated;

const raidSets = (JSON.parse(readFileSync(RAID_SETS, 'utf8')) as { specs: Record<string, LevelSet[]> }).specs;

function bareAt(spec: string, level: number): LevelSet {
  return generated.bare[spec]!.filter((row) => row.from_level <= level).at(-1)!;
}

/** The raid set the published raid number sends: the kit's buffs layered under the raid list. */
function raidAt(spec: string, level: number): LevelSet {
  return raidSets[spec]!.filter((row) => row.from_level <= level).at(-1)!;
}

const PAGE_ID = { bare: 'bare', raid: 'raid-buffed' } as const;

describe('the page defaults to the preset the band publishes', () => {
  it('offers Raid-buffed, Bare and Custom', () => {
    expect(BUFF_PRESETS.map((row) => [row.id, row.label])).toEqual([
      ['raid-buffed', 'Raid-buffed'],
      ['bare', 'Bare'],
      ['custom', 'Custom'],
    ]);
  });

  it('names the published preset at every level from the generated rule', () => {
    for (let level = 1; level <= 60; level += 1) {
      const want = generated.published.filter((row) => row.from_level <= level).at(-1)!.preset;
      expect(publishedPresetAt(level), `level ${level}`).toBe(PAGE_ID[want]);
    }
    expect(publishedPresetAt(40)).toBe('bare');
    expect(publishedPresetAt(60)).toBe('raid-buffed');
  });

  it('agrees with the headline entry of every published BiS file at every band', () => {
    const files = readdirSync(BIS_DIR).filter((name) => name.endsWith('.json'));
    expect(files.length).toBeGreaterThanOrEqual(27);
    for (const name of files) {
      const file = JSON.parse(readFileSync(path.join(BIS_DIR, name), 'utf8')) as BisFile;
      for (const band of BANDS) {
        const entry = selectBand(file, band, 'horde');
        if (entry === undefined) continue;
        expect(PAGE_ID[presetOf(entry)], `${name} band ${band}`).toBe(publishedPresetAt(band));
      }
    }
  });

  it('defaults every spec below 60 to the bare set (the class kit) and at 60 to Raid-buffed', () => {
    const specs = Object.keys(generated.bare);
    expect(specs.length).toBeGreaterThanOrEqual(27);
    for (const spec of specs) {
      for (const level of [...BANDS, 33, 59]) {
        const settings = defaultSettings(spec, level);
        if (level < 60) {
          expect(settings.preset, `${spec} ${level}`).toBe('bare');
          expect(settings.buffs, `${spec} ${level}`).toEqual(bareAt(spec, level).buffs);
          expect(settings.consumables, `${spec} ${level}`).toEqual(bareAt(spec, level).consumes);
        } else {
          expect(settings.preset, `${spec} ${level}`).toBe('raid-buffed');
          expect(settings.buffs, `${spec} ${level}`).toEqual(raidAt(spec, level).buffs);
          expect(settings.consumables, `${spec} ${level}`).toEqual(raidAt(spec, level).consumes);
        }
      }
    }
  });

  it('gives a level-40 rogue its poisons and a Survival hunter its mark under Bare', () => {
    expect(defaultSettings('rogue-combat', 40).consumables).toEqual([
      'main_hand_imbue:deadly_poison',
      'off_hand_imbue:instant_poison',
    ]);
    expect(defaultSettings('hunter-survival', 40).buffs).toContain('hunters_mark');
  });

  it('re-defaults a character swap to the new band’s preset until the player picks one', () => {
    const atSixty = withSpecForPreset(defaultSettings(), 'warrior-fury', 60);
    expect(atSixty.preset).toBe('raid-buffed');
    const atForty = withSpecForPreset(atSixty, 'warrior-fury', 40);
    expect(atForty.preset).toBe('bare');
    expect(atForty.buffs).toEqual(bareAt('warrior-fury', 40).buffs);
    const backToSixty = withSpecForPreset(atForty, 'mage-fire', 60);
    expect(backToSixty.preset).toBe('raid-buffed');
  });

  it('keeps a Raid-buffed pick below 60: it stays the player’s choice across a swap', () => {
    const picked = withPreset(defaultSettings('warrior-fury', 40), 'raid-buffed', 'warrior-fury', 40);
    expect(picked.preset).toBe('raid-buffed');
    expect(picked.buffs).toEqual([...RAID_BUFFS]);
    const swapped = withSpecForPreset(picked, 'rogue-combat', 30);
    expect(swapped.preset).toBe('raid-buffed');
    expect(swapped.buffs).toEqual([...RAID_BUFFS]);
    const protection = withSpecForPreset(picked, 'paladin-protection', 40);
    expect([...protection.buffs].sort()).toEqual(['devotion_aura', ...RAID_BUFFS].sort());
  });
});
