// web/src/lib/bis/preset-caption.ts
// The one-line caption under the BiS preset control, from the file's own `presets` block.
import { bisCopy } from './copy';
import type { BisPresetMeta } from './types';

interface PresetEffectGroup {
  heading: string;
  labels: string[];
}

export function raidCaptionFor(meta: BisPresetMeta): string {
  return bisCopy.presetRaidCaption(meta.label, meta.buffs.length, meta.debuffs.length, meta.consumes.length);
}

/** The non-empty groups, in reading order. */
function effectGroupsFor(meta: BisPresetMeta): PresetEffectGroup[] {
  const groups: PresetEffectGroup[] = [
    { heading: bisCopy.presetBuffsHeading, labels: meta.buffs.map((effect) => effect.label) },
    { heading: bisCopy.presetDebuffsHeading, labels: meta.debuffs.map((effect) => effect.label) },
    { heading: bisCopy.presetConsumesHeading, labels: meta.consumes.map((effect) => effect.label) },
  ];
  return groups.filter((group) => group.labels.length > 0);
}

/** The groups as one plain line: "Buffs: a, b. Debuffs: c. Consumables: d." */
export function effectLineFor(meta: BisPresetMeta): string {
  return effectGroupsFor(meta)
    .map((group) => `${group.heading}: ${group.labels.join(', ')}.`)
    .join(' ');
}
