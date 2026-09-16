import { describe, expect, it } from 'vitest';
import {
  cleanMechanicRows,
  encounterKey,
  mechanicRowKey,
  mostOftenHit,
  nightBossGroups,
  playerMechanics,
  unclassifiedAbilities,
  meleeBucket,
  unjudgedDeaths,
} from './mechanics';
import type { Ability, Actor, MechanicHit, MechanicRow, Summary } from './types';

function hit(guid: string, damage: number, pulls?: number): MechanicHit {
  return { guid, name: guid, hits: 1, damage, first_ms: 0, last_ms: 0, killed: false, pulls };
}

function row(partial: Partial<MechanicRow> & Pick<MechanicRow, 'spell_id' | 'kind'>): MechanicRow {
  return { name: `Spell ${partial.spell_id}`, ...partial };
}

describe('mechanicRowKey', () => {
  it('separates two bosses that list one spell id, and falls back to the boss name', () => {
    expect(mechanicRowKey({ spell_id: 5, encounter_id: 2363 })).not.toBe(
      mechanicRowKey({ spell_id: 5, encounter_id: 2360 }),
    );
    expect(mechanicRowKey({ spell_id: 5, encounter: 'Kaal' })).not.toBe(
      mechanicRowKey({ spell_id: 5, encounter: 'Kryxis' }),
    );
    // A single fight carries neither, so the spell id alone is the identity.
    expect(mechanicRowKey({ spell_id: 5 })).toBe(mechanicRowKey({ spell_id: 5 }));
    expect(encounterKey(2363, 'Kaal')).toBe(encounterKey(2363, 'Something else'));
  });
});

describe('mostOftenHit', () => {
  it('leads on pulls, breaks a tie on damage, and names nobody when nobody leads', () => {
    expect(mostOftenHit([hit('A', 10, 1), hit('B', 5, 3)])?.guid).toBe('B');
    // Equal pulls: the record decides, not the order the fold merged them in.
    expect(mostOftenHit([hit('A', 10, 2), hit('B', 500, 2)])?.guid).toBe('B');
    expect(mostOftenHit([hit('A', 10, 2), hit('B', 10, 2)])).toBeUndefined();
    expect(mostOftenHit(undefined)).toBeUndefined();
    expect(mostOftenHit([])).toBeUndefined();
  });
});

describe('cleanMechanicRows', () => {
  it('excludes by key, so one boss failing a spell does not clear the other boss of it', () => {
    const kaal = row({ spell_id: 5, kind: 'avoidable', encounter_id: 2363, encounter: 'Kaal' });
    const kryxis = row({ spell_id: 5, kind: 'avoidable', encounter_id: 2360, encounter: 'Kryxis' });
    const clean = cleanMechanicRows([kaal, kryxis], [mechanicRowKey(kaal)]);
    expect(clean).toEqual([kryxis]);
    // A copy of the row, not the same object: identity must not be what excludes it.
    expect(cleanMechanicRows([kaal], [mechanicRowKey({ ...kaal })])).toEqual([]);
  });
});

describe('nightBossGroups', () => {
  const kaal = row({
    spell_id: 5,
    kind: 'avoidable',
    name: 'Wicked Gash',
    encounter_id: 2363,
    encounter: 'Kaal',
    pulls_hit: 6,
  });
  const kryxis = row({
    spell_id: 5,
    kind: 'avoidable',
    name: 'Wicked Gash',
    encounter_id: 2360,
    encounter: 'Kryxis',
    pulls_hit: 1,
  });

  it('groups the rows that hit somebody under their boss, with that boss’s pulls', () => {
    const groups = nightBossGroups({
      table_found: true,
      rows: [
        kaal,
        kryxis,
        row({ spell_id: 9, kind: 'avoidable', encounter_id: 2363, encounter: 'Kaal', pulls_hit: 0 }),
        row({ spell_id: 11, kind: 'interrupt', encounter_id: 2363, encounter: 'Kaal', pulls_hit: 4 }),
      ],
      bosses: [
        { encounter_id: 2363, name: 'Kaal', pulls: 18 },
        { encounter_id: 2360, name: 'Kryxis', pulls: 3 },
      ],
    });
    expect(groups.map((group) => [group.name, group.pulls, group.rows.length])).toEqual([
      ['Kaal', 18, 1],
      ['Kryxis', 3, 1],
    ]);
    expect(groups[0].rows[0].pulls_hit).toBe(6);
  });

  it('drops a boss whose mechanics hit nobody rather than listing it at zero', () => {
    expect(
      nightBossGroups({
        table_found: true,
        rows: [row({ spell_id: 5, kind: 'avoidable', encounter_id: 2363, encounter: 'Kaal', pulls_hit: 0 })],
        bosses: [{ encounter_id: 2363, name: 'Kaal', pulls: 18 }],
      }),
    ).toEqual([]);
  });
});

describe('playerMechanics', () => {
  const rows: MechanicRow[] = [
    row({ spell_id: 1, name: 'Piercing Blur', kind: 'avoidable', players: [hit('A', 300)] }),
    row({ spell_id: 2, name: 'Gloom Squall', kind: 'avoidable', players: [] }),
    row({
      spell_id: 3,
      name: 'Vicious Headbutt',
      kind: 'unavoidable',
      players: [hit('A', 12_000), hit('B', 40)],
    }),
    row({ spell_id: 4, name: 'Hungering Drain', kind: 'interrupt', casts: 3 }),
  ];

  it('gives every roster player a card, with their unavoidable damage and what missed them', () => {
    const cards = playerMechanics(rows, [
      { guid: 'A', name: 'Hobolol' },
      { guid: 'B', name: 'Thalgrit' },
      { guid: 'C', name: 'Baelgrim' },
    ]);
    expect(cards.map((card) => card.guid)).toEqual(['A', 'B', 'C']);
    expect(cards[0]).toMatchObject({
      damage: 300,
      unavoidable: { damage: 12_000, names: ['Vicious Headbutt'] },
      avoided: ['Gloom Squall'],
    });
    // B and C never stood in anything, so both avoidable rows are on their "never hit by".
    expect(cards[1]).toMatchObject({
      damage: 0,
      unavoidable: { damage: 40, names: ['Vicious Headbutt'] },
      avoided: ['Piercing Blur', 'Gloom Squall'],
    });
    expect(cards[2].unavoidable).toEqual({ damage: 0, names: [] });
    expect(cards[2].hits).toEqual([]);
  });

  it('never says "never hit by" about a mechanic the night hit them with on another boss', () => {
    const cards = playerMechanics(
      [
        row({
          spell_id: 5,
          name: 'Wicked Gash',
          kind: 'avoidable',
          encounter_id: 2363,
          encounter: 'Kaal',
          players: [hit('A', 100)],
        }),
        row({
          spell_id: 5,
          name: 'Wicked Gash',
          kind: 'avoidable',
          encounter_id: 2360,
          encounter: 'Kryxis',
          players: [],
        }),
      ],
      [{ guid: 'A', name: 'Hobolol' }],
    );
    expect(cards[0].avoided).toEqual([]);
    expect(cards[0].hits).toHaveLength(1);
  });

  it('sorts the hits most costly first', () => {
    const cards = playerMechanics(
      [
        row({ spell_id: 1, name: 'Small', kind: 'avoidable', players: [hit('A', 10)] }),
        row({ spell_id: 2, name: 'Big', kind: 'avoidable', players: [hit('A', 900)] }),
      ],
      [],
    );
    expect(cards[0].hits.map((entry) => entry.row.name)).toEqual(['Big', 'Small']);
  });
});

describe('unclassifiedAbilities', () => {
  function ability(partial: Partial<Ability> & Pick<Ability, 'spell_id' | 'effective'>): Ability {
    return {
      name: `Spell ${partial.spell_id}`,
      total: partial.effective,
      hits: 1,
      crits: 0,
      ticks: 0,
      min: 0,
      max: 0,
      ...partial,
    };
  }
  function actor(guid: string, abilities: Ability[]): Actor {
    const effective = abilities.reduce((sum, a) => sum + a.effective, 0);
    return {
      guid,
      name: guid,
      total: effective,
      effective,
      active_ms: 0,
      abilities,
      targets: [],
      series: [],
    };
  }
  function summaryOf(damage_taken: Actor[], rows?: MechanicRow[]): Summary {
    return {
      engine_version: 't',
      fight_index: 1,
      duration_ms: 1000,
      damage_done: [],
      damage_taken,
      healing: [],
      healing_taken: [],
      deaths: [],
      auras: [],
      casts: [],
      interrupts: [],
      dispels: [],
      resources: [],
      threat: [],
      combatants: [],
      roster: [
        { guid: 'A', name: 'Hobolol', role: 'dps' },
        { guid: 'B', name: 'Thalgrit', role: 'tank' },
      ] as Summary['roster'],
      ...(rows === undefined ? {} : { mechanics: { table_found: true, rows } }),
    };
  }

  it('lists what hit a player and the table does not, most damage first', () => {
    const result = unclassifiedAbilities(
      summaryOf(
        [
          actor('A', [
            ability({ spell_id: 331415, name: 'Wicked Gash', effective: 500 }),
            ability({ spell_id: 999, name: 'Gloom Rot', effective: 300 }),
            ability({ spell_id: 0, name: 'Melee', effective: 9000 }),
          ]),
          actor('B', [ability({ spell_id: 999, name: 'Gloom Rot', effective: 40 })]),
          // Not on the roster: the boss's own damage taken is not a mechanic on anyone.
          actor('Creature-1', [ability({ spell_id: 116, name: 'Frostbolt', effective: 99_000 })]),
        ],
        [row({ spell_id: 331415, name: 'Wicked Gash', kind: 'avoidable' })],
      ),
    );
    expect(result).toEqual([{ spell_id: 999, name: 'Gloom Rot', damage: 340, players: 2 }]);
  });

  it('treats a summary with no table as classifying nothing', () => {
    const result = unclassifiedAbilities(
      summaryOf([actor('A', [ability({ spell_id: 7, name: 'Anima Lash', effective: 10 })])]),
    );
    expect(result).toEqual([{ spell_id: 7, name: 'Anima Lash', damage: 10, players: 1 }]);
  });

  it('leaves out a spell a roster player dealt damage with: friendly fire is nobody’s mechanic', () => {
    const summary = summaryOf([
      actor('A', [
        ability({ spell_id: 116, name: 'Frostbolt', effective: 500 }),
        ability({ spell_id: 999, name: 'Gloom Rot', effective: 40 }),
      ]),
    ]);
    const result = unclassifiedAbilities({
      ...summary,
      damage_done: [actor('B', [ability({ spell_id: 116, name: 'Frostbolt', effective: 500 })])],
    });
    expect(result.map((entry) => entry.spell_id)).toEqual([999]);
  });

  it('buckets the enemies’ melee swings on players, naming who took most', () => {
    const summary = summaryOf([
      actor('A', [
        ability({ spell_id: 0, name: 'Melee', effective: 300 }),
        ability({ spell_id: 7, effective: 10 }),
      ]),
      actor('B', [ability({ spell_id: 0, name: 'Melee', effective: 100 })]),
    ]);
    expect(meleeBucket(summary)).toEqual({
      damage: 400,
      players: 2,
      most: { name: 'A', damage: 300 },
      others: { damage: 100, players: 1 },
    });
    expect(meleeBucket(summaryOf([actor('A', [ability({ spell_id: 7, effective: 10 })])]))).toEqual({
      damage: 0,
      players: 0,
      most: undefined,
      others: { damage: 0, players: 0 },
    });
  });

  it('lists the deaths the table does not explain: a swing, an unlisted spell, or nothing named', () => {
    const summary = summaryOf([], [{ spell_id: 7, name: 'Anima Lash', kind: 'avoidable' }]);
    const death = (
      guid: string,
      at_ms: number,
      blow?: { spell_id: number; spell_name: string; source_name: string },
    ) =>
      ({
        guid,
        name: guid,
        at_ms,
        killing_blow: blow && { ...blow, at_ms, source_guid: 'C', amount: 1 },
        last: [],
        auras_held: [],
        auras_lost: [],
      }) as unknown as Summary['deaths'][number];
    const result = unjudgedDeaths({
      ...summary,
      deaths: [
        death('A', 1000, { spell_id: 7, spell_name: 'Anima Lash', source_name: 'Warden' }),
        death('B', 2000, { spell_id: 0, spell_name: '', source_name: 'Warden' }),
        death('A', 3000, { spell_id: 99, spell_name: 'Gloom', source_name: 'Bat' }),
        death('B', 4000),
      ],
    });
    expect(result.map((entry) => [entry.name, entry.at_ms, entry.by])).toEqual([
      ['B', 2000, 'Warden · melee'],
      ['A', 3000, 'Bat · Gloom'],
      ['B', 4000, 'nothing the log named'],
    ]);
  });
});

describe('a role’s ability landing on the others', () => {
  it('goes on that role’s card as a ranked entry, not under never hit by', () => {
    const gash: MechanicRow = {
      spell_id: 331415,
      name: 'Wicked Gash',
      kind: 'avoidable',
      role: 'tank',
      players: [
        {
          guid: 'dps-1',
          name: 'Mishvamp',
          hits: 3,
          damage: 9000,
          first_ms: 1000,
          last_ms: 5000,
          killed: false,
        },
        {
          guid: 'heal-1',
          name: 'Deadclasslol',
          hits: 1,
          damage: 4000,
          first_ms: 2000,
          last_ms: 2000,
          killed: true,
        },
      ],
    };
    const cards = playerMechanics(
      [gash],
      [
        { guid: 'tank-1', name: 'Hobolol', role: 'tank' },
        { guid: 'dps-1', name: 'Mishvamp', role: 'dps' },
        { guid: 'heal-1', name: 'Deadclasslol', role: 'healer' },
      ],
    );
    const tank = cards.find((card) => card.guid === 'tank-1');
    expect(tank?.avoided).toEqual([]);
    expect(tank?.hits).toHaveLength(1);
    expect(tank?.hits[0]).toMatchObject({ others: 2, hit: { hits: 4, damage: 13000, killed: true } });
    expect(tank?.damage).toBe(0);
    expect(tank?.placed).toBe(13000);
    // The victims' own cards are unchanged.
    expect(cards.find((card) => card.guid === 'dps-1')?.hits[0].others).toBeUndefined();
  });
});
