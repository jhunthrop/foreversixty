// web/src/lib/guild/roster-best.ts
// A Progression "Roster bests" row in a player's words (lifecycle round 8, P8-7): the spec carries
// its class and the encounter follows it, "Holy Paladin · Onyxia", and the number carries its unit,
// "488 HPS". The row is the raider's best KILL of THAT boss; the Roster tab's chip is their best
// boss kill across the tier's bosses, so for a raider the chip is the largest figure on this list.
// The execution score reads `92% of sim` (the kill against what that gear does in the simulator);
// with no score the row ends at its number, never a bare dash (guild round 3).
import { outputFigure, outputUnit } from './output-figure';
import type { GuildRosterBest } from '../rankings/api';
import { executionLabel, executionTitle } from '../sim/execution';
import { specClassName } from './spec-name';

export interface RosterBestLabels {
  what: string;
  value: string;
  hover: string;
  /** `92% of sim`, `null` when the fight carries no score. */
  execution: string | null;
  /** The execution's title: what the percentage measures. */
  executionHover: string;
}

export function rosterBestLabels(row: GuildRosterBest): RosterBestLabels {
  const unit = outputUnit(row.metric);
  const score = row.execution_score;
  return {
    what: `${specClassName(row.player.spec, row.player.class)} · ${row.encounter}`,
    value: `${outputFigure(row.value)} ${unit}`,
    hover: `Top ${unit} in any ${row.encounter} kill`,
    execution: score === null ? null : `${executionLabel(score)} of sim`,
    executionHover: executionTitle(score),
  };
}
