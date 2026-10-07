// web/src/lib/bis/preset-caption.ts
// The one-line caption under the BiS preset control, from the file's own `presets` block.
import { bisCopy } from './copy';
import type { BisPresetMeta } from './types';

export interface PresetEffectGroup {
  heading: string;
  labels: string[];
}

export function raidCaptionFor(meta: BisPresetMeta): string {
  return bisCopy.presetRaidCaption(meta.label, meta.buffs.length, meta.debuffs.length, meta.consumes.length);
}

/** The disclosure's non-empty groups, in reading order. */
export function effectGroupsFor(meta: BisPresetMeta): PresetEffectGroup[] {
  const groups: PresetEffectGroup[] = [
    { heading: bisCopy.presetBuffsHeading, labels: meta.buffs.map((effect) => effect.label) },
    { heading: bisCopy.presetDebuffsHeading, labels: meta.debuffs.map((effect) => effect.label) },
    { heading: bisCopy.presetConsumesHeading, labels: meta.consumes.map((effect) => effect.label) },
  ];
  return groups.filter((group) => group.labels.length > 0);
}
