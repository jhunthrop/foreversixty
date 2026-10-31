// web/src/lib/rankings/options.ts
// The choices each rankings filter offers, built once so the page and its tests agree. The first
// option of an optional filter is its "all" (empty value), named in the page's own words.
import { REGIONS, RULESETS } from '../characters';
import { classRows } from '../planner/reference';
import { titleize } from '../report/og-meta';
import { specIconSrc } from '../guild/raid-spec';
import { SPECS } from '../sim/specs';
import type { SelectOption } from '../ui/listbox-model';
import { classOption } from '../ui/select-options';
import type { EncounterOption } from './api';
import { rankingsCopy } from './copy';
import { PHASES } from './phases';
import { FACTIONS, RANKING_METRICS } from './url';

/** One filter choice: the shared dropdown's option (a resting `disabled` one is the raid
 *  roster's `—`, raids spec 3.7; `ariaLabel` names an option whose label is a mark). */
export type FilterOption = SelectOption;

const all = (label: string): FilterOption => ({ value: '', label });

/** The listed encounters; a linked encounter the list lacks (or has not loaded yet) is still
 *  offered, named from its slug, so the select never shows a blank. */
export function encounterOptions(encounters: readonly EncounterOption[], current: string): FilterOption[] {
  const listed = encounters.map((encounter) => ({ value: encounter.slug, label: encounter.name }));
  const missing = current !== '' && !listed.some((option) => option.value === current);
  return missing ? [...listed, { value: current, label: titleize(current) }] : listed;
}

export const rulesetOptions = (): FilterOption[] => [
  all(rankingsCopy.allRulesets),
  ...RULESETS.map((ruleset) => ({ value: ruleset.id, label: ruleset.label })),
];

export const regionOptions = (): FilterOption[] => [
  all(rankingsCopy.allRegions),
  ...REGIONS.map((region) => ({ value: region, label: region.toUpperCase() })),
];

export const phaseOptions = (): FilterOption[] => [
  all(rankingsCopy.allPhases),
  ...PHASES.map((phase) => ({ value: phase.id, label: phase.label })),
];

export const factionOptions = (): FilterOption[] => [
  all(rankingsCopy.bothFactions),
  ...FACTIONS.map((faction) => ({ value: faction, label: faction === 'horde' ? 'Horde' : 'Alliance' })),
];

export const metricOptions = (): FilterOption[] =>
  RANKING_METRICS.map((metric) => ({ value: metric.id, label: metric.label }));

export const classOptions = (): FilterOption[] => [
  all(rankingsCopy.everyClass),
  ...classRows.map((row) => classOption({ value: row.slug, label: row.name, className: row.slug })),
];

/** The chosen class's specs, by their own names (what the API files a build under). */
export const specOptions = (classSlug: string): FilterOption[] => [
  all(rankingsCopy.everySpec),
  ...SPECS.filter((spec) => spec.class_slug === classSlug).map((spec) => ({
    value: spec.name,
    label: spec.name,
    icon: { kind: 'spec' as const, src: specIconSrc(spec) },
  })),
];
