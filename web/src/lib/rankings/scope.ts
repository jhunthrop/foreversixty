// web/src/lib/rankings/scope.ts
// What a rankings board is about, in words: the cohort ("Fury Warrior"), the heading over the
// board ("Fury Warriors", "Every class and spec"), and the one-line summary the phone's filters
// collapse to. Pure, so the page and its tests share them.
import { rulesetLabel } from '../characters';
import { classDisplayName } from '../sim/spec-label';
import { SPECS, type Spec } from '../sim/specs';
import { rankingsCopy } from './copy';
import { RANKING_METRICS, type RankingsState } from './url';

/** The generated spec a class slug and a spec name (any case) name, if the site knows it. */
export function specFor(classSlug: string, specName: string): Spec | undefined {
  const wanted = specName.toLowerCase();
  return SPECS.find((spec) => spec.class_slug === classSlug && spec.name.toLowerCase() === wanted);
}

/** "Fury Warrior", "Warrior", or "" when there is no class. The spec's own name when known. */
export function cohortName(classSlug: string, specName: string): string {
  if (classSlug === '') return '';
  const className = classDisplayName(classSlug);
  if (specName === '') return className;
  return `${specFor(classSlug, specName)?.name ?? specName} ${className}`;
}

/** The board's heading: "Fury Warriors", "Warriors", or "Every class and spec". */
export function scopeHeading(classSlug: string, specName: string): string {
  const cohort = cohortName(classSlug, specName);
  return cohort === '' ? rankingsCopy.everyClassAndSpec : `${cohort}s`;
}

export interface SummaryParts {
  encounterName: string;
  state: Pick<RankingsState, 'ruleset' | 'region' | 'phase' | 'faction' | 'since' | 'metric'>;
  /** The metric the page starts on for this viewer; only a different choice is named. */
  defaultMetric: string;
}

function phaseWord(phase: string, phaseNames: ReadonlyMap<string, string>): string {
  return phase === '' ? rankingsCopy.allPhases : (phaseNames.get(phase) ?? phase);
}

/**
 * The phone's collapsed filters: `Onyxia · Normal · All regions · Phase 1`, then only the
 * narrowing choices that were made (a faction, today, a metric that is not the viewer's own).
 */
export function filterSummary(parts: SummaryParts, phaseNames: ReadonlyMap<string, string>): string {
  const { state } = parts;
  const metric = RANKING_METRICS.find((candidate) => candidate.id === state.metric)?.label ?? state.metric;
  return [
    parts.encounterName,
    state.ruleset === '' ? rankingsCopy.allRulesets : rulesetLabel(state.ruleset),
    state.region === '' ? rankingsCopy.allRegions : state.region.toUpperCase(),
    phaseWord(state.phase, phaseNames),
    state.faction === '' ? '' : state.faction === 'horde' ? 'Horde' : 'Alliance',
    state.since === 'today' ? rankingsCopy.todayOnly : '',
    state.metric === parts.defaultMetric ? '' : metric,
  ]
    .filter((part) => part !== '')
    .join(' · ');
}
