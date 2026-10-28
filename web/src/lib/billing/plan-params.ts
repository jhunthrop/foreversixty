// web/src/lib/billing/plan-params.ts
// The closed sets a `plan` or `interval` query value is checked against. Every page and link
// builder that reads or writes these params goes through here, so a value outside the set is
// dropped the same way everywhere.
import { INTERVALS, PLAN_ORDER, type Interval, type PlanKey } from './plans';

/** The plan a raw value names, or undefined when it is not one of the three keys. */
export function parsePlan(raw: string | null | undefined): PlanKey | undefined {
  return PLAN_ORDER.find((key) => key === raw);
}

/** The interval a raw value names, or undefined when it is not `monthly` or `yearly`. */
export function parseInterval(raw: string | null | undefined): Interval | undefined {
  return INTERVALS.find((key) => key === raw);
}
