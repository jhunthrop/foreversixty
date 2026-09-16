// web/src/lib/report/events.test.ts
import { describe, expect, it } from 'vitest';
import fixtureSummary from '../../fixtures/report/fights/3/summary.json';
import type { Summary } from './types';
import {
  EVENT_KINDS,
  filterEvents,
  streamEvents,
  summaryEvents,
  type EventKind,
  type SummaryEvent,
} from './events';

const summary = fixtureSummary as Summary;

describe('summaryEvents', () => {
  const events = summaryEvents(summary);

  it('merges casts, auras, deaths and killing hits into one ordered list', () => {
    expect(events.length).toBeGreaterThan(6);
    expect(events.map((event) => event.atMs)).toEqual(
      [...events.map((event) => event.atMs)].sort((a, b) => a - b),
    );
    expect(new Set(events.map((event) => event.kind))).toEqual(
      new Set(['cast', 'aura-applied', 'aura-removed', 'damage', 'death']),
    );
  });

  it('names the six kinds it can produce', () => {
    expect(EVENT_KINDS.map((kind) => kind.id)).toEqual([
      'cast',
      'aura-applied',
      'aura-removed',
      'damage',
      'heal',
      'death',
    ]);
  });

  it('describes the death in words a person reads', () => {
    const death = events.find((event) => event.kind === 'death');
    expect(death?.atMs).toBe(10_100);
    expect(death?.text).toBe('Thalgrit died');
  });

  it('describes a cast and an aura', () => {
    expect(events.find((event) => event.kind === 'cast')?.text).toBe('Morrowlyn cast Frostbolt');
    // The first aura of the fight is one the engine seeds from Baelgrim's COMBATANT_INFO
    // snapshot, which carries spell ids and no names, so it is filed under its id. The
    // Fortitude the priest casts four seconds in is still in the list behind it.
    expect(events.find((event) => event.kind === 'aura-applied')?.text).toBe('Spell #17 on Baelgrim');
    expect(
      events.some(
        (event) => event.kind === 'aura-applied' && event.text === 'Power Word: Fortitude on Baelgrim',
      ),
    ).toBe(true);
  });

  it('prints a stack change as one line, with no removal of the aura that stayed up', () => {
    const track = {
      ...summary.auras[0]!,
      name: 'Wicked Gash',
      appliers: [],
      segments: [
        { start_ms: 1000, end_ms: 2000, stacks: 1 },
        { start_ms: 2000, end_ms: 4000, stacks: 2 },
      ],
    };
    const lines = summaryEvents({ ...summary, auras: [track] })
      .filter((event) => event.text.startsWith('Wicked Gash'))
      .map((event) => `${event.atMs} ${event.text}`);
    const target = track.target_name.split('-')[0];
    expect(lines).toEqual([
      `1000 Wicked Gash on ${target}`,
      `2000 Wicked Gash at 2 stacks on ${target}`,
      `4000 Wicked Gash off ${target}`,
    ]);
  });

  it('names each application by its own applier, not the track’s whole caster set', () => {
    const track = {
      ...summary.auras[0]!,
      name: 'Power Word: Shield',
      appliers: ['Player-A', 'Player-B'],
      segments: [
        { start_ms: 1000, end_ms: 2000, stacks: 1, source_guid: 'Player-B' },
        { start_ms: 3000, end_ms: 4000, stacks: 1, source_guid: 'Player-A' },
      ],
    };
    const names = new Map([
      ['Player-A', 'Hanabanana-Realm'],
      ['Player-B', 'Reglitch-Realm'],
    ]);
    const applied = summaryEvents({ ...summary, auras: [track] }, names)
      .filter((event) => event.kind === 'aura-applied' && event.text.startsWith('Power Word: Shield'))
      .map((event) => event.text);
    expect(applied).toEqual([
      `Power Word: Shield on ${track.target_name.split('-')[0]} by Reglitch`,
      `Power Word: Shield on ${track.target_name.split('-')[0]} by Hanabanana`,
    ]);
  });
});

describe('filterEvents', () => {
  it('finds a line by its amount, comma or no comma', () => {
    const events: SummaryEvent[] = [
      { atMs: 1, kind: 'damage', guid: 'a', guids: ['a'], text: 'Boss hit Tank with Slam', amount: 7554 },
      { atMs: 2, kind: 'damage', guid: 'a', guids: ['a'], text: 'Boss hit Tank with Slam', amount: 120 },
    ];
    const kinds = new Set<EventKind>(['damage']);
    expect(filterEvents(events, kinds, '7,554').map((event) => event.amount)).toEqual([7554]);
    expect(filterEvents(events, kinds, '7554').map((event) => event.amount)).toEqual([7554]);
    expect(filterEvents(events, kinds, 'slam')).toHaveLength(2);
  });

  const events = summaryEvents(summary);

  it('keeps only the chosen kinds', () => {
    const only = filterEvents(events, new Set(['death']), '');
    expect(only).toHaveLength(1);
    expect(only[0].kind).toBe('death');
  });

  it('matches the search against the text, case-insensitively', () => {
    expect(
      filterEvents(events, new Set(EVENT_KINDS.map((kind) => kind.id)), 'frostbolt').length,
    ).toBeGreaterThan(0);
    expect(filterEvents(events, new Set(EVENT_KINDS.map((kind) => kind.id)), 'nothing here')).toEqual([]);
  });
});

describe('streamEvents', () => {
  it('words a miss from the defender’s side and a block beside the hit it took from', () => {
    const base = {
      atMs: 1000,
      sourceGuid: 'boss',
      sourceName: 'General Kaal',
      destGuid: 'tank',
      destName: 'Hobolol-Torghast',
      spellName: '',
      amount: 0,
      overheal: 0,
      absorbed: 0,
      blocked: 0,
      missType: '',
    };
    const [parry, hit] = streamEvents([
      { ...base, kind: 'missed', missType: 'PARRY' },
      { ...base, kind: 'damage', amount: 7554, absorbed: 261, blocked: 1200 },
    ]);
    expect(parry.text).toBe('Hobolol parried General Kaal’s Melee');
    // "parry" is not a substring of "parried": the type rides along for the find.
    expect(filterEvents([parry], new Set<EventKind>(['damage']), 'parry')).toHaveLength(1);
    expect(filterEvents([parry], new Set<EventKind>(['damage']), 'avoided')).toHaveLength(1);
    expect(parry.kind).toBe('damage');
    expect(parry.guid).toBe('tank');
    expect(hit.text).toBe('General Kaal hit Hobolol with Melee (261 absorbed, 1,200 blocked)');
  });
});
