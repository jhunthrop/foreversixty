// web/src/lib/sim/result-band.ts
// The saved run's band as data (spec 2026-11-04-sim-results 3.3): which heading it wears, the
// descriptor, the error line and the facts row. Pure, so the figures are pinned by a test and
// the talent split is computed once for the whole page.
import { classRows, raceRows } from '../planner/reference';
import { PLACEHOLDER_CHARACTER_NAMES, ranksFromTalentsString } from './character';
import { durationLabel } from './encounter';
import { formatMargin } from './estimate';
import { relativeError } from './precision';
import { bandCopy, grouped } from './result-copy';
import { specDisplayName, specLabel } from './spec-label';
import type { SimResult } from './types';

export interface BandFact {
  id: 'fight' | 'targets' | 'buffs' | 'consumables' | 'range' | 'runs' | 'saved';
  label: string;
  value: string;
}

/** Points per tree, from the stored talents string: "17/32/2". The page's one split. */
export function talentSplit(talents: string): string {
  if (talents === '') return '';
  return ranksFromTalentsString(talents)
    .map((tree) => tree.reduce((sum, rank) => sum + rank, 0))
    .join('/');
}

/** The h1: the member's title, a character's real name, or the spec. */
export function bandHeading(result: SimResult): string {
  if (result.title !== undefined && result.title !== '') return result.title;
  const { name } = result.request.character;
  if (name !== '' && !PLACEHOLDER_CHARACTER_NAMES.has(name)) return name;
  return specLabel(result.request.spec);
}

/** "Troll Fury Warrior", then "Level 60", then "17/32/2 talents": each only when known. */
export function bandDescriptor(result: SimResult): string[] {
  const { character, spec } = result.request;
  const race = raceRows.find((row) => row.slug === character.race)?.name;
  const className = classRows.find((row) => row.slug === character.class)?.name;
  const who = [race, specDisplayName(spec), className].filter((part) => part !== undefined).join(' ');
  const split = talentSplit(character.talents);
  return [
    who,
    ...(character.level > 0 ? [bandCopy.level(character.level)] : []),
    ...(split === '' ? [] : [bandCopy.talents(split)]),
  ];
}

/** The margin and the relative error of the mean, as the band prints them. */
export function bandError(result: SimResult): { margin: string; percent: string } | null {
  const { dps } = result;
  if (dps.mean <= 0 || dps.error <= 0) return null;
  return { margin: `± ${formatMargin(dps.error)}`, percent: `${(relativeError(dps) * 100).toFixed(2)}%` };
}

/** Whole percent of a fight's variation: 0.2 is 20. */
export function variationPercent(variation: number): number {
  return Math.round(variation * 100);
}

/** "4 Nov 2026" in UTC, so every viewer reads one date for one save. Null when unreadable. */
export function savedDate(iso: string | undefined): string | null {
  if (iso === undefined || iso === '') return null;
  const when = new Date(iso);
  if (Number.isNaN(when.getTime())) return null;
  return when.toLocaleDateString('en-GB', {
    day: 'numeric',
    month: 'short',
    year: 'numeric',
    timeZone: 'UTC',
  });
}

/**
 * Raid when the list equals the spec's raid preset, None when empty, else Custom. `raidBuffs` is
 * null until the preset table's chunk has loaded; the request's own preset answers until then.
 */
export function buffsFact(result: SimResult, raidBuffs: readonly string[] | null): string {
  const { buffs } = result.request.character;
  if (buffs.length === 0) return bandCopy.none;
  if (raidBuffs === null) return result.request.preset === 'raid' ? bandCopy.buffsRaid : bandCopy.buffsCustom;
  const wanted = new Set(raidBuffs);
  const same = buffs.length === wanted.size && buffs.every((buff) => wanted.has(buff));
  return same ? bandCopy.buffsRaid : bandCopy.buffsCustom;
}

export function bandFacts(result: SimResult, raidBuffs: readonly string[] | null): BandFact[] {
  const { encounter, character } = result.request;
  const { facts } = bandCopy;
  const saved = savedDate(result.created_at);
  return [
    {
      id: 'fight',
      label: facts.fight,
      value: bandCopy.fight(durationLabel(encounter.duration_sec), variationPercent(encounter.variation)),
    },
    { id: 'targets', label: facts.targets, value: String(encounter.targets) },
    { id: 'buffs', label: facts.buffs, value: buffsFact(result, raidBuffs) },
    {
      id: 'consumables',
      label: facts.consumables,
      value: character.consumes.length === 0 ? bandCopy.none : String(character.consumes.length),
    },
    {
      id: 'range',
      label: facts.range,
      value: bandCopy.range(grouped(result.dps.min), grouped(result.dps.max)),
    },
    { id: 'runs', label: facts.runs, value: grouped(result.iterations_run) },
    ...(saved === null ? [] : [{ id: 'saved' as const, label: facts.saved, value: saved }]),
  ];
}
