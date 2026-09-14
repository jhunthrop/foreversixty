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
    expect(events.map((event) => event.atMs)).toEqual([...events.map((event) => event.atMs)].sort((a, b) => a - b));
    expect(new Set(events.map((event) => event.kind))).toEqual(
      new Set(['cast', 'aura-applied', 'aura-removed', 'damage', 'death']),
    );
  });

  it('names the five kinds it can produce', () => {
    expect(EVENT_KINDS.map((kind) => kind.id)).toEqual([
      'cast', 'aura-applied', 'aura-removed', 'damage', 'death',
    ]);
  });

  it('describes the death in words a person reads', () => {
    const death = events.find((event) => event.kind === 'death');
    expect(death?.atMs).toBe(10_100);
    expect(death?.text).toBe('Thalgrit died');
  });

  it('describes a cast and an aura', () => {
    expect(events.find((event) => event.kind === 'cast')?.text).toBe('Morrowlyn cast Frostbolt');
    expect(events.find((event) => event.kind === 'aura-applied')?.text).toBe(
      'Power Word: Fortitude on Baelgrim',
    );
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
    expect(filterEvents(events, new Set(EVENT_KINDS.map((kind) => kind.id)), 'frostbolt').length).toBeGreaterThan(0);
    expect(filterEvents(events, new Set(EVENT_KINDS.map((kind) => kind.id)), 'nothing here')).toEqual([]);
  });
});
