// web/src/lib/reports/hero-hook.ts
// Logs landing spec (2026-10-04) §4.A.1's per-player hook, pulled out of LogsHero.svelte for
// the same reason tree-sizes.ts exists for the report page's own At-pull row: the selection
// rules and the dps/label formatting are plain functions, testable without mounting a
// component, and planner-link.ts / sim-link.ts are reused unchanged, never re-implemented.
import { classColorVar } from '../report/format';
import { plannerLinkFor, type PlannerLink } from '../report/planner-link';
import { simLinkFor, type SimLink } from '../report/sim-link';
import type { CombatantRow, FightEntry, RosterRow } from '../report/types';

/** The last kill, falling back to the last fight overall when the report has none (§4.A.1).
 *  Null only for a report with no fights at all. */
export function pickLastKillFight(fights: readonly FightEntry[]): FightEntry | null {
  if (fights.length === 0) return null;
  const kills = fights.filter((fight) => fight.kind === 'encounter' && fight.kill);
  if (kills.length > 0) return kills[kills.length - 1];
  return fights[fights.length - 1];
}

/** Matches the roster against the visitor's own characters by name, case-insensitive; the
 *  first match wins. No match (an officer uploading someone else's log) -> the roster's
 *  own top-dps row. Null only for an empty roster. */
export function selectHookRow(
  roster: readonly RosterRow[],
  myCharacterNames: readonly string[],
): RosterRow | null {
  if (roster.length === 0) return null;
  const mine = new Set(myCharacterNames.map((name) => name.toLowerCase()));
  const matched = roster.find((row) => mine.has(row.name.toLowerCase()));
  if (matched !== undefined) return matched;
  return [...roster].sort((a, b) => b.dps - a.dps)[0];
}

/** Designer ruling, logs mock review (2026-10-04): under 1,000 one decimal, at or above a
 *  whole number with a thousands separator -- both mono. */
export function formatHookDps(dps: number): string {
  return dps < 1000 ? dps.toFixed(1) : Math.round(dps).toLocaleString('en-US');
}

export interface HookLinks {
  name: string;
  classColor: string;
  dpsText: string;
  /** Null exactly when `plannerLinkFor` returns null (the class carries no planner data). */
  plannerLink: PlannerLink | null;
  simLink: SimLink;
}

export interface HookLinksInput {
  reportId: string;
  fightIndex: number;
  dataBuild: string;
  row: RosterRow;
  combatants: readonly CombatantRow[];
  treeSizesFor: (className: string | undefined) => number[];
}

/** Builds the hook's name/dps/links from the selected roster row. A row with no matching
 *  `COMBATANT_INFO` (never seen one for that guid) gets no planner link -- the sim link
 *  always resolves, since it needs only the guid, not a combatant record. */
export function buildHookLinks(input: HookLinksInput): HookLinks {
  const combatant = input.combatants.find((candidate) => candidate.guid === input.row.guid);
  const plannerLink =
    combatant === undefined
      ? null
      : plannerLinkFor({
          dataBuild: input.dataBuild,
          className: input.row.class,
          treeSizes: input.treeSizesFor(input.row.class),
          combatant,
        });
  return {
    name: input.row.name,
    classColor: classColorVar(input.row.class),
    dpsText: formatHookDps(input.row.dps),
    plannerLink,
    simLink: simLinkFor(input.reportId, input.fightIndex, input.row.guid),
  };
}
