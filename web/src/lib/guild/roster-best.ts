// web/src/lib/guild/roster-best.ts
// A Progression "Roster bests" row in a player's words (lifecycle round 8, P8-7): the spec carries
// its class and the encounter follows it, "Holy Paladin · Onyxia", and the number carries its unit,
// "488 HPS". The row is the raider's best KILL of THAT boss; the Roster tab's chip is their best
// boss kill across the tier's bosses, so for a raider the chip is the largest figure on this list.
import { formatAmount } from '../report/format';
import type { GuildRosterBest } from '../rankings/api';
import { specClassName } from './spec-name';

export function rosterBestLabels(row: GuildRosterBest): { what: string; value: string; hover: string } {
  const unit = row.metric.toUpperCase();
  return {
    what: `${specClassName(row.player.spec, row.player.class)} · ${row.encounter}`,
    value: `${formatAmount(Math.round(row.value))} ${unit}`,
    hover: `Best ${unit} in any ${row.encounter} kill`,
  };
}
