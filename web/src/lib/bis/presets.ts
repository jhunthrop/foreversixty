// web/src/lib/bis/presets.ts
// Which published band a reader sees: the level playing field is the Phase 1 raid-ready
// preset, with the bare character (no buffs, debuffs or consumables) kept as the alternative.
// Pure and node-free so the build-time pages (load.ts) and the browser planner island
// (hover.ts) share one selection rule instead of two copies.
import type { BisBand, BisFile, BisPresetId, BisPresetMeta, Faction } from './types';

export const BARE_PRESET: BisPresetId = 'bare';
export const RAID_PRESET: BisPresetId = 'raid';

/** Display and fallback order: the headline preset first. */
const PRESET_ORDER: readonly BisPresetId[] = [RAID_PRESET, BARE_PRESET];

/** A band entry with no `preset` field is a bare band (every file published before presets). */
export function presetOf(band: BisBand): BisPresetId {
  return band.preset ?? BARE_PRESET;
}

/** Narrows an untrusted string (a URL parameter) to a preset id. */
export function parsePresetId(value: string | null | undefined): BisPresetId | undefined {
  return PRESET_ORDER.find((id) => id === value);
}

/**
 * The band entry for `band` + `faction`. `preset` omitted means the default: raid when the
 * file has a raid entry for that band, bare otherwise. An explicit preset the file lacks
 * falls back to bare, so a stale `?preset=raid` link on an older file still resolves.
 */
export function selectBand(
  file: BisFile,
  band: number,
  faction: Faction,
  preset?: BisPresetId,
): BisBand | undefined {
  const candidates = file.bands.filter((entry) => entry.band === band && entry.faction === faction);
  const wanted = preset ?? RAID_PRESET;
  return (
    candidates.find((entry) => presetOf(entry) === wanted) ??
    candidates.find((entry) => presetOf(entry) === BARE_PRESET) ??
    candidates[0]
  );
}

/** The presets one band + faction offers, headline preset first. */
export function bandPresets(file: BisFile, band: number, faction: Faction): BisPresetId[] {
  const offered = new Set(
    file.bands.filter((entry) => entry.band === band && entry.faction === faction).map(presetOf),
  );
  return PRESET_ORDER.filter((id) => offered.has(id));
}

/** The presets a whole file offers anywhere, headline preset first. */
export function filePresets(file: BisFile): BisPresetId[] {
  const offered = new Set(file.bands.map(presetOf));
  return PRESET_ORDER.filter((id) => offered.has(id));
}

/** The file's own description of a preset (label, buffs, debuffs, consumes), if it ships one. */
export function presetMeta(file: BisFile, preset: BisPresetId): BisPresetMeta | undefined {
  return file.presets?.[preset];
}

export const BARE_PRESET_LABEL = 'Bare';

/** The label a surface that cannot carry the toggle shows beside a number it took from
 *  `band`: the file's own label for the preset, "Bare" for the bare character. */
export function presetLabelFor(file: BisFile, band: BisBand): string {
  const preset = presetOf(band);
  return preset === BARE_PRESET ? BARE_PRESET_LABEL : (presetMeta(file, preset)?.label ?? preset);
}
