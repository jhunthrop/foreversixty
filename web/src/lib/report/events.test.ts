// web/src/lib/report/events.test.ts
import { describe, expect, it } from 'vitest';
import fixtureSummary from '../../fixtures/report/fights/3/summary.json';
import type { Summary } from './types';
import {
  EVENT_CHIPS,
  EVENT_KINDS,
  chipOn,
  filterEvents,
  sentenceParts,
  toggleChip,
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
    // Ashfang's Bite at 1.5s is now the fight's earliest cast, ahead of Morrowlyn's
    // Frostbolt: a pet's cast row sorts into the timeline by its own timestamp same
    // as any other caster's.
    expect(events.find((event) => event.kind === 'cast')?.text).toBe('Ashfang casts Bite');
    // The first aura of the fight is one the engine seeds from Baelgrim's COMBATANT_INFO
    // snapshot, which carries spell ids and no names, so it is filed under its id. The
    // Fortitude the priest casts four seconds in is still in the list behind it.
    expect(events.find((event) => event.kind === 'aura-applied')?.text).toBe('Baelgrim gains Spell #17');
    expect(
      events.some(
        (event) =>
          event.kind === 'aura-applied' && event.text.startsWith('Baelgrim gains Power Word: Fortitude'),
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
      .filter((event) => event.spellName === 'Wicked Gash')
      .map((event) => `${event.atMs} ${event.text}`);
    const target = track.target_name.split('-')[0];
    const verb = track.type === 'DEBUFF' ? 'is afflicted by' : 'gains';
    expect(lines).toEqual([
      `1000 ${target} ${verb} Wicked Gash`,
      `2000 ${target} ${verb} Wicked Gash (2)`,
      `4000 Wicked Gash fades from ${target}`,
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
      .filter((event) => event.kind === 'aura-applied' && event.spellName === 'Power Word: Shield')
      .map((event) => event.text);
    const verb = track.type === 'DEBUFF' ? 'is afflicted by' : 'gains';
    expect(applied).toEqual([
      `${track.target_name.split('-')[0]} ${verb} Power Word: Shield from Reglitch`,
      `${track.target_name.split('-')[0]} ${verb} Power Word: Shield from Hanabanana`,
    ]);
  });

  it('folds one aura landing on several units in the same instant into one line', () => {
    const base = summary.auras[0]!;
    const hit = (guid: string, name: string) => ({
      ...base,
      type: 'DEBUFF',
      name: 'Thunder Clap',
      spell_id: 8198,
      target_guid: guid,
      target_name: name,
      appliers: ['Player-W'],
      segments: [{ start_ms: 1500, end_ms: 9000, stacks: 1, source_guid: 'Player-W' }],
    });
    const lines = summaryEvents(
      { ...summary, auras: [hit('Creature-1', 'Saltspine'), hit('Creature-2', 'Marsh Skulker')] },
      new Map([['Player-W', 'Kell-Realm']]),
    ).filter((event) => event.spellName === 'Thunder Clap');
    expect(lines.map((event) => event.text)).toEqual([
      'Saltspine and Marsh Skulker are afflicted by Thunder Clap from Kell',
      'Thunder Clap fades from Saltspine and Marsh Skulker',
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

describe('aura refreshes in the full stream', () => {
  it('words a refresh from the caster’s side and files it under auras applied', () => {
    const [line] = streamEvents([
      {
        atMs: 20_000,
        kind: 'aura_refresh',
        sourceGuid: 'Player-2',
        sourceName: 'Sunwick-Nightslayer',
        destGuid: 'Player-1',
        destName: 'Baelgrim-Nightslayer',
        spellName: 'Power Word: Fortitude',
        amount: 0,
        overheal: 0,
        absorbed: 0,
        blocked: 0,
        missType: '',
      },
    ]);
    expect(line.kind).toBe('aura-applied');
    expect(line.text).toBe('Sunwick refreshed Power Word: Fortitude on Baelgrim');
    expect(line.guid).toBe('Player-1');
    expect(line.guids).toEqual(['Player-2', 'Player-1']);
    expect(line.amount).toBeUndefined();
  });

  it('says the aura refreshed itself when the log names no caster', () => {
    const [line] = streamEvents([
      {
        atMs: 1000,
        kind: 'aura_refresh',
        sourceGuid: '',
        sourceName: '',
        destGuid: 'Player-1',
        destName: 'Baelgrim-Nightslayer',
        spellName: 'Rend',
        amount: 0,
        overheal: 0,
        absorbed: 0,
        blocked: 0,
        missType: '',
      },
    ]);
    expect(line.text).toBe('Rend refreshed on Baelgrim');
  });
});

describe('the Show chips', () => {
  const auras = EVENT_CHIPS.find((chip) => chip.id === 'auras')!;

  it('cover every kind once', () => {
    expect(EVENT_CHIPS.flatMap((chip) => chip.kinds).sort()).toEqual(
      EVENT_KINDS.map((kind) => kind.id).sort(),
    );
  });

  it('switch every kind they hold together, keeping the url ids', () => {
    expect(chipOn(auras, [])).toBe(true);
    expect(toggleChip(auras, ['cast'])).toEqual(['cast', 'aura-applied', 'aura-removed']);
    // Half off reads as off; pressing it turns the whole chip back on.
    expect(chipOn(auras, ['aura-removed'])).toBe(false);
    expect(toggleChip(auras, ['aura-removed', 'death'])).toEqual(['death']);
  });
});

describe('sentenceParts', () => {
  it('cuts the line at each name, the longest first', () => {
    expect(
      sentenceParts({
        text: 'Kel hit Kelthas with Melee',
        actors: [
          { guid: 'a', name: 'Kel' },
          { guid: 'b', name: 'Kelthas' },
        ],
      }),
    ).toEqual([
      { text: 'Kel', guid: 'a' },
      { text: ' hit ' },
      { text: 'Kelthas', guid: 'b' },
      { text: ' with Melee' },
    ]);
  });

  it('names the units of the summary’s own lines', () => {
    const death = summaryEvents(summary).find((event) => event.kind === 'death')!;
    const parts = sentenceParts(death);
    expect(parts[0].guid).toBe(death.guid);
    expect(parts.map((part) => part.text).join('')).toBe(death.text);
  });

  it('gives a cast its spell id for the icon', () => {
    const cast = summaryEvents(summary).find((event) => event.kind === 'cast')!;
    expect(cast.spellId).toBe(summary.casts.find((row) => row.sequence.includes(cast.atMs))?.spell_id);
  });
});
