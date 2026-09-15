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
import { splitUnitName } from '../characters';
import type { Summary } from './types';

export type EventKind = 'cast' | 'aura-applied' | 'aura-removed' | 'damage' | 'death';

export const EVENT_KINDS: readonly { id: EventKind; label: string }[] = [
  { id: 'cast', label: 'Casts' },
  { id: 'aura-applied', label: 'Auras applied' },
  { id: 'aura-removed', label: 'Auras removed' },
  { id: 'damage', label: 'Damage' },
  { id: 'death', label: 'Deaths' },
];

export interface SummaryEvent {
  atMs: number;
  kind: EventKind;
  /** One line, already phrased; the view prints it rather than assembling it again. */
  text: string;
  /** The actor the line is about, for class colouring. */
  guid: string;
  amount?: number;
}

export function summaryEvents(summary: Summary): SummaryEvent[] {
  const events: SummaryEvent[] = [];

  for (const row of summary.casts) {
    const caster = splitUnitName(row.name).name;
    for (const at of row.sequence) {
      events.push({ atMs: at, kind: 'cast', guid: row.guid, text: `${caster} cast ${row.spell_name}` });
    }
  }

  for (const track of summary.auras) {
    const target = splitUnitName(track.target_name).name;
    for (const segment of track.segments) {
      events.push({
        atMs: segment.start_ms,
        kind: 'aura-applied',
        guid: track.target_guid,
        text: `${track.name} on ${target}`,
      });
      events.push({
        atMs: segment.end_ms,
        kind: 'aura-removed',
        guid: track.target_guid,
        text: `${track.name} off ${target}`,
      });
    }
  }

  for (const death of summary.deaths) {
    const who = splitUnitName(death.name).name;
    for (const hit of death.last) {
      events.push({
        atMs: hit.at_ms,
        kind: 'damage',
        guid: death.guid,
        amount: hit.amount,
        text: `${splitUnitName(hit.source_name).name} hit ${who} with ${hit.spell_name === '' ? 'Melee' : hit.spell_name}`,
      });
    }
    events.push({ atMs: death.at_ms, kind: 'death', guid: death.guid, text: `${who} died` });
  }

  // Stable: equal timestamps keep the order the sources were walked in, which puts the
  // damage that killed someone above the death it caused.
  return events
    .map((event, index) => ({ event, index }))
    .sort((a, b) => a.event.atMs - b.event.atMs || a.index - b.index)
    .map((entry) => entry.event);
}

export function filterEvents(events: SummaryEvent[], kinds: Set<EventKind>, search: string): SummaryEvent[] {
  const needle = search.trim().toLowerCase();
  return events.filter(
    (event) => kinds.has(event.kind) && (needle === '' || event.text.toLowerCase().includes(needle)),
  );
}
