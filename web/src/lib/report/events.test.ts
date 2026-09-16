// web/src/lib/report/events.test.ts
import { describe, expect, it } from 'vitest';
import fixtureSummary from '../../fixtures/report/fights/3/summary.json';
import type { Summary } from './types';
import { EVENT_KINDS, filterEvents, summaryEvents } from './events';

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
