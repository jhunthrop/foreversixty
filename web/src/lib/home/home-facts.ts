// web/src/lib/home/home-facts.ts
// Home's density pass (2026-11-09): the short facts the panels carry beside their own figures,
// each read from a field the page already loads. Your standing in the guild (GuildHomeStanding and
// the account's attendance), the guild's own summary when no raid is planned, the upgrades' total
// and best-in-slot count (UpgradesResult), the movers on the spec's board (/v1/home movers) and a
// guild sim's gap to yours. Pure: the home islands read it, the unit tests pin it.
import type { GuildHomeStanding, GuildHomeSummary } from '../guild/api';
import type { AttendanceSource } from '../guild/attendance-contract';
import type { MeCharacterGuild } from '../account/me-types';
import { scoreUnitWord } from '../bis/tank-view';
import type { GuildSimRow, MoverRow, Movers } from './home-answer';
import type { UpgradesResult } from './upgrades';
import { figureText } from './home-b-view';

/** One fact on a line: a strong figure and the words after it (`#2` `of 4 Fury Warriors by item level`). */
export interface HomeFact {
  n: string;
  text: string;
}

export const factsCopy = {
  itemLevel: (n: number) => `Item level ${n}`,
  byItemLevel: (of: number, spec: string) => `of ${of} guild ${spec} by item level`,
  attended: (nights: number, unit: 'raids' | 'raid nights') => `of ${nights} ${unit} attended`,
  raiders: (n: number) => (n === 1 ? 'raider' : 'raiders'),
  bossesDown: (named: number) => `of ${named} bosses down`,
  /** An upgrade row's worn piece: `over Mantle of Thieves`. */
  over: 'over',
  emptySlot: 'Empty slot',
  seeAll: (n: number) => `See all ${n}`,
  slotsBis: 'best in slot',
  climbing: 'Climbing this week',
  top: 'Top this week',
  places: 'places',
  ahead: 'ahead of you',
  behind: 'behind you',
} as const;

/** `#2 of 4 guild Fury Warriors by item level`, `3 of 4 raids attended`: the viewer's place in the guild; a
 *  standing with no peers (one of one) and an account with no attendance are left out. */
export function standingFacts(
  standing: GuildHomeStanding | null | undefined,
  guild: MeCharacterGuild | undefined,
  source: AttendanceSource | undefined,
): HomeFact[] {
  const facts: HomeFact[] = [];
  if (standing && standing.same_spec_count > 1) {
    const spec = `${standing.spec} ${standing.class}s`;
    facts.push({
      n: `#${standing.rank_by_item_level}`,
      text: factsCopy.byItemLevel(standing.same_spec_count, spec),
    });
  }
  const attendance = guild?.attendance;
  if (attendance && attendance.nights > 0) {
    facts.push({
      n: String(attendance.present),
      text: factsCopy.attended(attendance.nights, source === 'raids' ? 'raids' : 'raid nights'),
    });
  }
  return facts;
}

/** `14 raiders · 3 of 9 bosses down`: the guild's summary; bosses only once the tier is pulled. */
export function guildSummaryFacts(summary: GuildHomeSummary | undefined): HomeFact[] {
  if (summary === undefined) return [];
  const facts: HomeFact[] = [{ n: String(summary.raiders), text: factsCopy.raiders(summary.raiders) }];
  if (summary.tier.pulls_total > 0) {
    facts.push({ n: String(summary.tier.down), text: factsCopy.bossesDown(summary.tier.named_encounters) });
  }
  return facts;
}

/** Your upgrades' foot: `+52.1 DPS · 5 best in slot` (See all N beside it counts the upgrades); each
 *  clause only when it is above zero. */
export function upgradesFacts(result: UpgradesResult | undefined): HomeFact[] {
  if (result === undefined || result.upgrades.length === 0) return [];
  const facts: HomeFact[] = [];
  if (result.totalGainDps > 0) {
    facts.push({
      n: `+${figureText(result.totalGainDps)}`,
      text: scoreUnitWord(result.scoreUnit),
    });
  }
  if (result.alreadyBis.length > 0) {
    facts.push({ n: String(result.alreadyBis.length), text: factsCopy.slotsBis });
  }
  return facts;
}

/** The movers block's title: the biggest climbers once a weekly board exists, else the top three. */
export function moversTitle(movers: Movers): string {
  return movers.baseline ? factsCopy.climbing : factsCopy.top;
}

/** A mover's trail: `+6` places climbed; null with no baseline or no climb. */
export function moverDelta(row: MoverRow): string | null {
  return row.delta === null || row.delta <= 0 ? null : `+${row.delta}`;
}

/** A guild sim's gap to yours: `63.7 DPS ahead of you`, `4.2 DPS behind you`; null on your own row or
 *  with no sim. */
export function simGap(row: GuildSimRow, you: GuildSimRow | null, unit: string): string | null {
  if (you === null || row.you) return null;
  const gap = row.value - you.value;
  return `${figureText(Math.abs(gap))} ${unit} ${gap >= 0 ? factsCopy.ahead : factsCopy.behind}`;
}

/** The ladder's bar: the row's value against the top of the ladder, 0 to 1. */
export function simShare(row: GuildSimRow, top: number): number {
  return top > 0 ? row.value / top : 0;
}
