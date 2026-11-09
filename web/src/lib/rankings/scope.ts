// web/src/lib/rankings/scope.ts
// What a rankings board is about, in words: the cohort ("Fury Warrior"), the heading over the
// board ("Fury Warriors", "Every class and spec"). Pure, so the page and its tests share them.
import { classDisplayName } from '../sim/spec-label';
import { SPECS, type Spec } from '../sim/specs';
import { rankingsCopy } from './copy';

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
  return specFor(classSlug, specName)?.full_name ?? `${specName} ${className}`;
}

/** The board's heading: "Fury Warriors", "Warriors", or "Every class and spec". */
export function scopeHeading(classSlug: string, specName: string): string {
  const cohort = cohortName(classSlug, specName);
  return cohort === '' ? rankingsCopy.everyClassAndSpec : `${cohort}s`;
}
