// web/src/lib/guild/raider-ranks.ts
// The Settings tab's raider ranks (addon hub L9, migration 0103): which guild ranks the addon's
// Guild and Officer tabs count as raiders. One row per rank, highest (the guild master, 0) first,
// from the synced roster's rank lines, plus any saved rank nobody holds now so a save never drops
// it unseen. Null until an officer says, and the addon then counts every level 60. Pure.
import type { SetupRankLine } from './setup-api';
import { namesWithMore } from './setup-model';

export const raiderRanksCopy = {
  label: 'Raider ranks',
  legend: 'Counts as a raider',
  /** What the setting drives, under the label. */
  explain: 'Ticked ranks count as raiders in the addon’s Guild and Officer tabs.',
  unset: 'Not set. The addon counts every level 60 as a raider.',
  ranks: (ranks: number): string => `${ranks} ${ranks === 1 ? 'rank' : 'ranks'}`,
  /** The summary's `23 characters, 20 at 60`; the level part only when the API sent it. */
  characters: (characters: number, atSixty: number | undefined): string =>
    `${characters} ${characters === 1 ? 'character' : 'characters'}${atSixty === undefined ? '' : `, ${atSixty} at 60`}`,
  /** A rank's holders by level: `13 of 14 at 60`, `all 3 at 60`, `1 at 60`; just the count from an
   *  API that sent no level. */
  holders: (characters: number, atSixty: number | undefined): string => {
    if (atSixty === undefined) return `${characters} on the roster`;
    if (atSixty !== characters) return `${atSixty} of ${characters} at 60`;
    return characters === 1 ? '1 at 60' : `all ${characters} at 60`;
  },
  rank: (index: number): string => `Rank ${index}`,
  guildMaster: 'Guild master',
  officer: 'Officer',
  nobody: 'Nobody on the roster',
  noRanks: 'No ranks synced yet. The addon reads them from your guild roster.',
  ranksFailed: 'Ranks did not load.',
  save: 'Save',
  saved: 'Saved.',
} as const;

export interface RaiderRankRow {
  rankIndex: number;
  title: string;
  /** `Guild master` or `Officer`, empty for the rest. */
  role: string;
  /** How many hold it, how many are 60 and the first names: `3 of 4 at 60 · Kraggor, Mirelle +2`. */
  sub: string;
  checked: boolean;
}

function roleOf(rankIndex: number, officerMax: number): string {
  if (rankIndex === 0) return raiderRanksCopy.guildMaster;
  return rankIndex <= officerMax ? raiderRanksCopy.officer : '';
}

function subOf(line: SetupRankLine | undefined): string {
  if (line === undefined || line.members === 0) return raiderRanksCopy.nobody;
  const holders = raiderRanksCopy.holders(line.members, line.at_sixty);
  if (line.names.length === 0) return holders;
  return `${holders} · ${namesWithMore(line.names, line.members - line.names.length)}`;
}

/** The rows, highest rank first: every rank on the roster and every chosen one. */
export function raiderRankRows(
  lines: readonly SetupRankLine[],
  officerMax: number,
  chosen: readonly number[],
): RaiderRankRow[] {
  const byIndex = new Map(lines.map((line) => [line.rank_index, line]));
  const indexes = [...new Set([...byIndex.keys(), ...chosen])].sort((a, b) => a - b);
  return indexes.map((rankIndex) => ({
    rankIndex,
    title: raiderRanksCopy.rank(rankIndex),
    role: roleOf(rankIndex, officerMax),
    sub: subOf(byIndex.get(rankIndex)),
    checked: chosen.includes(rankIndex),
  }));
}

/** chosen with rankIndex ticked or cleared, sorted; a new list. */
export function toggleRank(chosen: readonly number[], rankIndex: number, on: boolean): number[] {
  const rest = chosen.filter((index) => index !== rankIndex);
  return (on ? [...rest, rankIndex] : rest).sort((a, b) => a - b);
}

/** Whether two rank lists hold the same ranks. */
export function sameRanks(a: readonly number[], b: readonly number[]): boolean {
  const left = [...a].sort((x, y) => x - y);
  const right = [...b].sort((x, y) => x - y);
  return left.length === right.length && left.every((value, i) => value === right[i]);
}

/** What the ticked ranks cover, following the boxes as they change: no box ticked is the
 *  addon's default. Without the roster's rank lines (loading, failed) no character count. */
export function raiderRanksSummary(
  lines: readonly SetupRankLine[] | null,
  ticked: readonly number[],
): string {
  if (ticked.length === 0) return raiderRanksCopy.unset;
  if (lines === null) return raiderRanksCopy.ranks(ticked.length);
  const covered = lines.filter((line) => ticked.includes(line.rank_index));
  const characters = covered.reduce((sum, line) => sum + line.members, 0);
  const atSixty = covered.every((line) => line.at_sixty !== undefined)
    ? covered.reduce((sum, line) => sum + (line.at_sixty ?? 0), 0)
    : undefined;
  return raiderRanksCopy.characters(characters, atSixty);
}
