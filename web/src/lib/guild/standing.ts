// web/src/lib/guild/standing.ts
// Guild control-centre spec §4.A.1/§4.A.2: the Overview header's standing line (ranked by
// item level among same-spec peers) and the member-only "before Thursday" sentence. Both
// read straight off `GuildHomeStanding` (the contract's own `home.standing`) -- no second
// fetch, matching spec §8's "no new round trip from Overview."
import type { GuildHomeStanding } from './api';

/** Exact copy, spec §6: the all-clear line is a stated fact, never a silent absence. */
export const EVERY_CHECK_PASSES_LINE = 'Every readiness check passes. Nothing needed before Thursday.';

/**
 * The header's own "where do I stand" clause (spec §4.A.1), unchanged mechanism from v1:
 * ranked by item level among same-spec, same-class peers this week. A spec with no other
 * ranked peer reads "the only {spec} {class} in this guild this week" rather than a
 * meaningless "1 of 1".
 */
export function standingSentence(standing: GuildHomeStanding): string {
  const className = capitalize(standing.class);
  if (standing.same_spec_count <= 1) {
    return `the only ${standing.spec} ${className} in this guild this week`;
  }
  return (
    `${standing.rank_by_item_level} of ${standing.same_spec_count} ${standing.spec} ${className}s ` +
    `this week by item level (${standing.item_level} ilvl)`
  );
}

/**
 * Spec §4.A.2/§6, verbatim: joins the standing's own `needs_before_next_raid` (server-
 * phrased, max 3) with "; ", or the all-clear line when the array is empty -- `undefined`/
 * `null` (a pre-contract response, or a viewer with no standing at all) is treated the
 * same as empty, never thrown.
 */
export function needsBeforeThursdaySentence(needs: readonly string[] | null | undefined): string {
  const fails = needs ?? [];
  if (fails.length === 0) return EVERY_CHECK_PASSES_LINE;
  return `What the guild needs from you before Thursday: ${fails.join('; ')}.`;
}

function capitalize(value: string): string {
  return value.length === 0 ? value : value[0].toUpperCase() + value.slice(1);
}
