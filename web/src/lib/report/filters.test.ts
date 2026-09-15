// web/src/lib/report/filters.test.ts
import { describe, expect, it } from 'vitest';
import fixtureReport from '../../fixtures/report/report.json';
import fixtureSummary from '../../fixtures/report/fights/3/summary.json';
import type { Actor, ReportFile, Summary, Unit } from './types';
import {
  bossGuidsOf,
  friendlyGuids,
  DEFAULT_FILTERS,
  abilityOptions,
  applyActorFilters,
  bossGuids,
  playerGuids,
  targetOptions,
} from './filters';

const report = fixtureReport as ReportFile;
const summary = fixtureSummary as Summary;
const context = {
  bosses: bossGuids(report.units, 'Warden Kelthas'),
  players: playerGuids(report.units),
  deaths: summary.deaths,
};

describe('unit sets from report.json', () => {
  it('finds the encounter boss by the fight’s own name', () => {
    expect([...context.bosses]).toEqual(['Creature-0-2085-2284-7855-169754-0000AA0002']);
    expect(bossGuids(report.units, 'Trash').size).toBe(0);
  });

  it('knows which units are players, and which are on their side', () => {
    expect(context.players.size).toBe(5);
    // Five players and the hunter's pet: a pet is its owner's, not an enemy.
    expect(friendlyGuids(report.units).size).toBe(6);
    expect(context.players.has('Player-4184-000000A1')).toBe(true);
    expect(context.players.has('Creature-0-2085-2284-7855-169754-0000AA0002')).toBe(false);
  });
});

describe('filter options', () => {
  it('lists every ability and every target present, deduplicated and sorted by size', () => {
    expect(abilityOptions(summary.damage_done).map((option) => option.name)).toEqual([
      'Anima Lash',
      'Slam',
      'Frostbolt',
      'Melee',
      'Shadow Word: Pain',
    ]);
    expect(targetOptions(summary.damage_done).map((option) => option.name)).toContain('Warden Kelthas');
  });
});

describe('applyActorFilters', () => {
  it('changes nothing by default', () => {
    expect(applyActorFilters(summary.damage_done, DEFAULT_FILTERS, context)).toEqual(summary.damage_done);
  });

  it('keeps only the named target, and rescales the row to it', () => {
    const filtered = applyActorFilters(
      summary.damage_done,
      { ...DEFAULT_FILTERS, target: 'Creature-0-2085-2284-7855-169754-0000AA0002' },
      context,
    );
    expect(filtered.map((actor) => actor.name)).not.toContain('Warden Kelthas');
    const mage = filtered.find((actor) => actor.name === 'Morrowlyn-Nightslayer');
    expect(mage?.total).toBe(3110);
  });

  it('keeps only the named ability', () => {
    const filtered = applyActorFilters(summary.damage_done, { ...DEFAULT_FILTERS, ability: 1464 }, context);
    expect(filtered).toHaveLength(1);
    expect(filtered[0].name).toBe('Baelgrim-Nightslayer');
    expect(filtered[0].abilities.map((ability) => ability.name)).toEqual(['Slam']);
  });

  it('boss damage only drops damage aimed anywhere else', () => {
    const filtered = applyActorFilters(summary.damage_done, { ...DEFAULT_FILTERS, bossOnly: true }, context);
    expect(filtered.map((actor) => actor.name)).not.toContain('Warden Kelthas');
    expect(filtered.map((actor) => actor.name)).toContain('Baelgrim-Nightslayer');
  });

  it('hides NPC rows when asked, leaving the players', () => {
    const filtered = applyActorFilters(
      summary.damage_done,
      { ...DEFAULT_FILTERS, playersOnly: true },
      context,
    );
    expect(filtered.every((actor) => context.players.has(actor.guid))).toBe(true);
  });

  it('counting overkill raises a row that had some', () => {
    const plain = applyActorFilters(summary.damage_taken, DEFAULT_FILTERS, context);
    const withOverkill = applyActorFilters(
      summary.damage_taken,
      { ...DEFAULT_FILTERS, countOverkill: true },
      context,
    );
    const victim = (rows: typeof plain): number =>
      rows.find((actor) => actor.name === 'Thalgrit-Nightslayer')?.total ?? 0;
    expect(victim(withOverkill)).toBe(victim(plain) + 100);
  });

  it('ignoring events after a death clips that actor’s series at the second they died', () => {
    const filtered = applyActorFilters(
      summary.damage_taken,
      { ...DEFAULT_FILTERS, ignoreAfterDeath: true },
      context,
    );
    const victim = filtered.find((actor) => actor.name === 'Thalgrit-Nightslayer');
    expect(victim?.series.length).toBeLessThanOrEqual(11);
  });

  it('finds every night’s boss by name, and none of the players', () => {
    const units = [
      { guid: 'Creature-1', name: 'Kaal', kind: 'creature' },
      { guid: 'Creature-2', name: 'Kryxis', kind: 'creature' },
      { guid: 'Player-1', name: 'Kaal', kind: 'player' },
    ] as Unit[];
    expect([...bossGuidsOf(units, ['Kaal', 'Kryxis', ''])]).toEqual(['Creature-1', 'Creature-2']);
  });

  it('folds same-named units into one target option that matches all of them', () => {
    const actors: Actor[] = [
      {
        guid: 'P1',
        name: 'One',
        total: 30,
        effective: 30,
        active_ms: 0,
        abilities: [
          {
            spell_id: 1,
            name: 'Bolt',
            total: 30,
            effective: 30,
            hits: 3,
            crits: 0,
            ticks: 0,
            min: 10,
            max: 10,
          },
        ],
        targets: [
          { guid: 'C1', name: 'Mindbender', total: 10 },
          { guid: 'C2', name: 'Mindbender', total: 10 },
          { guid: 'C3', name: 'Boss', total: 10 },
        ],
        series: [30],
      },
    ];
    const options = targetOptions(actors);
    expect(options.map((option) => option.name)).toEqual(['Mindbender ×2', 'Boss']);
    const filtered = applyActorFilters(
      actors,
      { ...DEFAULT_FILTERS, target: options[0].id },
      {
        bosses: new Set(),
        players: new Set(['P1']),
        deaths: [],
      },
    );
    expect(filtered[0].targets.map((target) => target.guid)).toEqual(['C1', 'C2']);
    expect(filtered[0].effective).toBe(20);
  });

  it('finds a boss whose unit name is not the encounter’s to the letter', () => {
    const units = [
      { guid: 'C1', name: 'Halkias', kind: 'creature' },
      { guid: 'C2', name: 'Amarth', kind: 'creature' },
      { guid: 'C3', name: 'Surgeon Stitchflesh', kind: 'creature' },
      { guid: 'C4', name: 'Stitchflesh Abomination', kind: 'creature' },
      { guid: 'C5', name: 'Grand Overseer', kind: 'creature' },
    ] as Unit[];
    expect([...bossGuidsOf(units, ['Halkias, the Sin-Stained Goliath'])]).toEqual(['C1']);
    expect([...bossGuidsOf(units, ['Amarth, The Harvester'])]).toEqual(['C2']);
    expect([...bossGuidsOf(units, ['Stichflesh'])]).toEqual(['C3', 'C4']);
    expect([...bossGuidsOf(units, ['General Kaal'])]).toEqual([]);
  });
});
