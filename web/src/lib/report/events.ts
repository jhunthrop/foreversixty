// web/src/lib/report/events.ts
// The Events view's list, built from what the summary already timestamps.
//
// The complete event stream is events.parquet and belongs to the Queries view: it is two
// to ten megabytes and a DuckDB instance. Everything the summary timestamps -- every cast,
// every aura application and removal, every death and the hits that caused it -- answers
// "what happened at 1:12?" with no download at all, which is the question an event list is
// actually opened for.
//
// Every source read here is exact under a window (window.ts's scopeSummary): cast
// sequence, aura segments and deaths are none of the fields that scoping only scales, so
// this list never needs an approximate mark.
import type { StreamLine } from './exact';
import { splitUnitName } from '../characters';
import type { Summary } from './types';

export type EventKind = 'cast' | 'aura-applied' | 'aura-removed' | 'damage' | 'heal' | 'death';

export const EVENT_KINDS: readonly { id: EventKind; label: string }[] = [
  { id: 'cast', label: 'Casts' },
  { id: 'aura-applied', label: 'Auras applied' },
  { id: 'aura-removed', label: 'Auras removed' },
  { id: 'damage', label: 'Hits' },
  { id: 'heal', label: 'Heals' },
  { id: 'death', label: 'Deaths' },
];

export interface SummaryEvent {
  atMs: number;
  kind: EventKind;
  /** One line, already phrased; the view prints it rather than assembling it again. */
  text: string;
  /** The actor the line is about, for class colouring. */
  guid: string;
  /** Every actor the line involves, for scoping by source: a heal is the healer's and the healed's. */
  guids: string[];
  amount?: number;
  /** Words the find matches besides the text: a miss's own type ("parry"), so the noun finds the verb. */
  tags?: string[];
}

/** How a miss reads, by the log's miss type: the hit that did not land, and why. */
const MISS_VERBS: Record<string, string> = {
  PARRY: 'parried',
  DODGE: 'dodged',
  BLOCK: 'blocked',
  MISS: 'was missed by',
  IMMUNE: 'was immune to',
  ABSORB: 'absorbed',
  DEFLECT: 'deflected',
  EVADE: 'evaded',
  RESIST: 'resisted',
  REFLECT: 'reflected',
};

/**
 * The full stream's lines in the view's shape: every hit, heal and avoided hit, said the
 * same way as a death's. A miss is filed under the hits it is a member of, worded from
 * the defender's side ("Hobolol parried General Kaal's Melee"), so a search for "parry"
 * finds the moment the Damage Taken row counted.
 */
export function streamEvents(lines: StreamLine[]): SummaryEvent[] {
  return lines.map((line) => {
    const who = splitUnitName(line.sourceName).name || 'Something';
    const whom = splitUnitName(line.destName).name;
    const spell = line.spellName === '' ? 'Melee' : line.spellName;
    if (line.kind === 'missed') {
      const verb = MISS_VERBS[line.missType] ?? `avoided (${line.missType.toLowerCase()})`;
      const soaked =
        line.missType === 'ABSORB' && line.amount > 0 ? ` (${line.amount.toLocaleString()})` : '';
      return {
        atMs: line.atMs,
        kind: 'damage',
        guid: line.destGuid,
        guids: [line.sourceGuid, line.destGuid],
        text: `${whom} ${verb} ${who}’s ${spell}${soaked}`,
        tags: [line.missType.toLowerCase(), 'avoided'],
      };
    }
    const parts =
      line.kind === 'heal'
        ? line.overheal > 0
          ? [`${line.overheal.toLocaleString()} over`]
          : []
        : [
            ...(line.absorbed > 0 ? [`${line.absorbed.toLocaleString()} absorbed`] : []),
            ...(line.blocked > 0 ? [`${line.blocked.toLocaleString()} blocked`] : []),
          ];
    const detail = parts.length === 0 ? '' : ` (${parts.join(', ')})`;
    return {
      atMs: line.atMs,
      kind: line.kind,
      guid: line.kind === 'heal' ? line.sourceGuid : line.destGuid,
      guids: [line.sourceGuid, line.destGuid],
      amount: line.kind === 'heal' ? line.amount - line.overheal : line.amount,
      text: `${who} ${line.kind === 'heal' ? 'healed' : 'hit'} ${whom} with ${spell}${detail}`,
    };
  });
}

export function summaryEvents(
  summary: Summary,
  names: ReadonlyMap<string, string> = new Map(),
): SummaryEvent[] {
  const events: SummaryEvent[] = [];

  for (const row of summary.casts) {
    const caster = splitUnitName(row.name).name;
    for (const at of row.sequence) {
      events.push({
        atMs: at,
        kind: 'cast',
        guid: row.guid,
        guids: [row.guid],
        text: `${caster} cast ${row.spell_name}`,
      });
    }
  }

  for (const track of summary.auras) {
    const target = splitUnitName(track.target_name).name;
    const applier = track.appliers
      .map((guid) => splitUnitName(names.get(guid) ?? '').name)
      .filter(Boolean)
      .slice(0, 2)
      .join(', ');
    let previous: { end_ms: number; stacks: number } | undefined;
    const joins = (later: { start_ms: number } | undefined, earlier: { end_ms: number }): boolean =>
      later !== undefined && later.start_ms <= earlier.end_ms + 50;
    track.segments.forEach((segment, index) => {
      // A segment that begins where the last one ended with a different stack count is a
      // stack change, not a new application, and the verb says so; the aura never came
      // off, so the earlier segment prints no removal. The engine keeps no refresh
      // events, so nothing here claims a refresh either way.
      const continues = previous !== undefined && joins(segment, previous);
      const continued = joins(track.segments[index + 1], segment);
      const verb =
        continues && segment.stacks !== previous?.stacks
          ? `at ${segment.stacks} ${segment.stacks === 1 ? 'stack' : 'stacks'} on`
          : 'on';
      previous = { end_ms: segment.end_ms, stacks: segment.stacks };
      // The segment's own applier when the engine kept it; the track's set only when it
      // did not, so two priests' shields each say which priest.
      const own =
        segment.source_guid === undefined ? '' : splitUnitName(names.get(segment.source_guid) ?? '').name;
      const by = own === '' ? applier : own;
      const guids =
        segment.source_guid === undefined
          ? [track.target_guid, ...track.appliers]
          : [track.target_guid, segment.source_guid];
      events.push({
        atMs: segment.start_ms,
        kind: 'aura-applied',
        guid: track.target_guid,
        guids,
        text: `${track.name} ${verb} ${target}${by === '' ? '' : ` by ${by}`}`,
      });
      if (!continued)
        events.push({
          atMs: segment.end_ms,
          kind: 'aura-removed',
          guid: track.target_guid,
          guids,
          text: `${track.name} off ${target}${by === '' ? '' : ` (by ${by})`}`,
        });
    });
  }

  for (const death of summary.deaths) {
    const who = splitUnitName(death.name).name;
    for (const hit of death.last) {
      events.push({
        atMs: hit.at_ms,
        kind: 'damage',
        guid: death.guid,
        guids: [death.guid, hit.source_guid],
        amount: hit.amount,
        text: `${splitUnitName(hit.source_name).name} hit ${who} with ${hit.spell_name === '' ? 'Melee' : hit.spell_name}`,
      });
    }
    for (const heal of death.heals ?? []) {
      events.push({
        atMs: heal.at_ms,
        kind: 'heal',
        guid: death.guid,
        guids: [death.guid, heal.source_guid],
        amount: heal.amount - (heal.overheal ?? 0),
        text: `${splitUnitName(heal.source_name).name} healed ${who} with ${heal.spell_name}`,
      });
    }
    events.push({
      atMs: death.at_ms,
      kind: 'death',
      guid: death.guid,
      guids: [death.guid],
      text: `${who} died`,
    });
  }

  // Stable: equal timestamps keep the order the sources were walked in, which puts the
  // damage that killed someone above the death it caused.
  return events
    .map((event, index) => ({ event, index }))
    .sort((a, b) => a.event.atMs - b.event.atMs || a.index - b.index)
    .map((entry) => entry.event);
}

/**
 * The find matches a line's words and its amount: a reader who just read "7,554" off the
 * Damage Taken tab types it here to find the hit, with or without the comma.
 */
export function filterEvents(events: SummaryEvent[], kinds: Set<EventKind>, search: string): SummaryEvent[] {
  const needle = search.trim().toLowerCase();
  const digits = needle.replaceAll(',', '');
  const byAmount = digits !== '' && /^\d+$/.test(digits);
  return events.filter(
    (event) =>
      kinds.has(event.kind) &&
      (needle === '' ||
        event.text.toLowerCase().includes(needle) ||
        (event.tags?.some((tag) => tag.includes(needle)) ?? false) ||
        (byAmount && event.amount !== undefined && String(event.amount).includes(digits))),
  );
}
