// web/src/lib/home/feed-mix.ts
// The signed-in home's "Latest on Forever" (home body pass 2026-11-09): one row per guild member's
// sim (the latest figure, the change since the one before), the guild name left out when it is the
// viewer's own, at most three sims so the other kinds of event keep their place, and the viewer's
// own spec's site events (Best in slot refreshed, the guide updated) from the floor beside them.
// Pure: HomeFeed reads it, the unit tests pin it. The API writes the words; this only regroups the
// segments it sent (api/internal/activity guildSimEvent).
import { FEED_ROWS, FLOOR_MIN, type FeedRow, type FeedSegment } from './feed';

/** Rows home asks for: enough that seven survive the sims folding into one per character. */
export const HOME_FEED_FETCH = 30;
/** Guild sims on the feed at most; the Guild sims ladder below carries the rest. */
export const FEED_SIMS_SHOWN = 3;
/** The floor's first-raids row: the plaque above already counts it down. */
const FLOOR_RAIDS_ID = 'floor:raids';
/** A sim figure as the API writes it: whole numbers from 100, one decimal below (FormatFigure). */
const WHOLE_FROM = 100;

/** `" simmed 512 DPS, #2 of 14 DPS in "`, the sentence between the name and the guild. */
const SIM_HEAD = /^ simmed ([\d.,]+) (\S+)(?:(, #\d+ of \d+) \S+)? in $/;

interface SimLine {
  name: FeedSegment;
  value: number;
  unit: string;
  place: string | null;
  guild: string;
  tail: string;
}

/** A guild sim row's parts, or null for a row in any other words. */
export function simLine(row: FeedRow): SimLine | null {
  if (row.kind !== 'guild_sim' || row.segments.length !== 4) return null;
  const [name, head, guild, tail] = row.segments;
  const match = SIM_HEAD.exec(head.text);
  if (name.style !== 'name' || guild.style !== 'strong' || match === null) return null;
  return {
    name,
    value: Number(match[1].replaceAll(',', '')),
    unit: match[2],
    place: match[3] ?? null,
    guild: guild.text,
    tail: tail.text,
  };
}

/** The same figure rule the API writes sims in. */
export function simFigureText(value: number): string {
  return value >= WHOLE_FROM ? String(Math.round(value)) : value.toFixed(1);
}

/** `+24`, `−3.1`, at the figure's own precision; null when the change rounds to nothing. */
export function simDeltaText(latest: number, previous: number): string | null {
  const change = latest - previous;
  const size = Math.abs(change);
  const text = latest >= WHOLE_FROM ? String(Math.round(size)) : size.toFixed(1);
  if (Number(text) === 0) return null;
  return `${change > 0 ? '+' : '−'}${text}`;
}

function simSegments(line: SimLine, previous: number | null, ownGuild: string | undefined): FeedSegment[] {
  const delta = previous === null ? null : simDeltaText(line.value, previous);
  return [
    line.name,
    { text: ' simmed ' },
    { text: `${simFigureText(line.value)} ${line.unit}`, style: 'strong' },
    ...(delta === null
      ? []
      : [{ text: ' ' }, { text: delta, style: delta.startsWith('+') ? 'gain' : 'loss' } as FeedSegment]),
    ...(line.place === null ? [] : [{ text: line.place }]),
    ...(line.guild === ownGuild
      ? []
      : [{ text: ' in ' }, { text: line.guild, style: 'strong' } as FeedSegment]),
    { text: line.tail },
  ];
}

/** A sim that measured nothing (a healer with nobody to heal, 0.0 HPS): the API never posts one,
 *  and a row that reached the feed before that rule is dropped here. */
function measuredNothing(line: SimLine | null): boolean {
  return line !== null && line.value <= 0;
}

/** One row per character's sim, newest first: the latest figure with the change since the one
 *  before; the viewer's own guild name dropped. A sim that measured nothing is left out. Rows in
 *  other words pass through untouched. */
export function foldSims(rows: readonly FeedRow[], ownGuild: string | undefined): FeedRow[] {
  const newest = [...rows]
    .filter((row) => !measuredNothing(simLine(row)))
    .sort((a, b) => Date.parse(b.at) - Date.parse(a.at));
  const lines = new Map<FeedRow, SimLine>();
  const byCharacter = new Map<string, FeedRow[]>();
  for (const row of newest) {
    const line = simLine(row);
    if (line === null) continue;
    lines.set(row, line);
    const key = `${row.crest ?? ''}:${line.name.text}`;
    byCharacter.set(key, [...(byCharacter.get(key) ?? []), row]);
  }
  const later = new Set([...byCharacter.values()].flatMap((list) => list.slice(1)));
  return newest
    .filter((row) => !later.has(row))
    .map((row) => {
      const line = lines.get(row);
      if (line === undefined) return row;
      const key = `${row.crest ?? ''}:${line.name.text}`;
      const before = byCharacter.get(key)?.[1];
      const previous = before === undefined ? null : (lines.get(before)?.value ?? null);
      return { ...row, segments: simSegments(line, previous, ownGuild) };
    });
}

/** At most FEED_SIMS_SHOWN guild sims, the newest; every other row kept. */
function capSims(rows: readonly FeedRow[]): FeedRow[] {
  let sims = 0;
  return rows.filter((row) => {
    if (row.kind !== 'guild_sim') return true;
    sims += 1;
    return sims <= FEED_SIMS_SHOWN;
  });
}

/** A floor row the API already says: the spec's BiS changed on the feed itself. */
function saidByApi(floorRow: FeedRow, rows: readonly FeedRow[]): boolean {
  return (
    floorRow.tag === 'best_in_slot' &&
    rows.some((row) => row.kind === 'bis_change' && row.href === floorRow.href)
  );
}

export interface SignedInFeed {
  rows: readonly FeedRow[];
  floor: readonly FeedRow[];
  spec: string | undefined;
  ownGuild: string | undefined;
  /** The guild's raid is on the plaque: its post is left out. */
  raidOnPlaque: boolean;
}

/** The signed-in feed: the API's rows with the sims folded and capped, the viewer's spec's site
 *  events beside them, newest first, seven at most; a quiet feed fills from the rest of the floor. */
export function signedInRows(input: SignedInFeed): FeedRow[] {
  const api = capSims(foldSims(input.rows, input.ownGuild)).filter(
    (row) => !(input.raidOnPlaque && row.kind === 'guild_raid'),
  );
  const floor = input.floor.filter((row) => row.id !== FLOOR_RAIDS_ID && !saidByApi(row, api));
  const own = floor
    .filter((row) => input.spec !== undefined && row.spec === input.spec)
    .map((row) => ({ ...row, tag: 'your_spec' as const }));
  const shown = [...api, ...own].sort((a, b) => Date.parse(b.at) - Date.parse(a.at)).slice(0, FEED_ROWS);
  if (shown.length >= FLOOR_MIN) return shown;
  const used = new Set(shown.map((row) => row.id));
  return [...shown, ...floor.filter((row) => !used.has(row.id)).slice(0, FLOOR_MIN - shown.length)];
}
