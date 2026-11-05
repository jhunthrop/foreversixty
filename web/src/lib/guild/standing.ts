// web/src/lib/guild/standing.ts
// Guild control-centre spec §4.A.1/§4.A.2, week-one spec §3.2: the band's standing line (ranked by
// item level among same-spec peers, or a leveling character's level) and the viewer's needs line.
// The ranked sentence reads `GuildHomeStanding`; the needs line reads the viewer's own row on the
// same /readiness response the board renders, so the two can never disagree.
import type { GuildHomeStanding, GuildReadinessRow } from './api';
import { RAID_LEVEL, isLeveling, passesEveryCheck, readinessFails } from './readiness-view';
import { specName } from './spec-name';

/**
 * The header's own "where do I stand" clause (spec §4.A.1), unchanged mechanism from v1:
 * ranked by item level among same-spec, same-class peers this week. A spec with no other
 * ranked peer says nothing at all (empty): "the only Beast Mastery Hunter in this guild" was
 * a gloss nobody asked for, and its spec came from a different place than the header chip's, so
 * the two read as a contradiction (round 8, P8-1).
 */
export function standingSentence(standing: GuildHomeStanding): string {
  const className = capitalize(standing.class);
  const spec = specName(standing.spec);
  if (standing.same_spec_count <= 1) return '';
  return (
    `${standing.rank_by_item_level} of ${standing.same_spec_count} ${spec} ${className}s ` +
    `this week by item level (${standing.item_level} ilvl)`
  );
}

/**
 * The band's one line for a character under 60 (week-one spec §3.2): it has neither a rank by item
 * level nor readiness checks yet, so the standing line is dropped and this single line says both.
 */
export function levelingLine(name: string, level: number): string {
  return `${name} is level ${level}. Rank and readiness checks start at ${RAID_LEVEL}.`;
}

/** The needs line's three outcomes; `clear` is the green one, `leveling` plain text. */
export interface NeedsLine {
  kind: 'needs' | 'clear' | 'leveling';
  text: string;
}

/**
 * The viewer's own line (week-one spec §3.2.3), its subject the current character. `before` is the
 * next raid's deadline (`raid-copy.ts` `nextRaidDeadline`): `before Thursday`. Null when the row
 * has nothing to say: nothing read, or no export.
 */
export function needsLine(row: GuildReadinessRow, before: string): NeedsLine | null {
  if (isLeveling(row)) {
    return {
      kind: 'leveling',
      text: levelingLine(row.name, row.level ?? 0),
    };
  }
  if (passesEveryCheck(row)) return { kind: 'clear', text: `${row.name} passes every readiness check.` };
  const fails = readinessFails(row);
  return fails.length === 0
    ? null
    : { kind: 'needs', text: `What ${row.name} needs ${before}: ${fails.join('; ')}.` };
}

function capitalize(value: string): string {
  return value.length === 0 ? value : value[0].toUpperCase() + value.slice(1);
}
