// web/src/lib/guides/rotation-lists.test.ts
import activeBuild from '../../data/active-build.json';
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import type { AplGate } from '../bis/apl-gates';
import { loadRotations, rotationEntryFor, rotationLinesFor } from '../bis/rotation-view';
import { decodeBuildTrees } from './build-tree';
import type { CollapsedLine } from '../bis/rotation-collapse';
import {
  clauseOf,
  inPageWords,
  joinSpellRuns,
  sayRepeatedRulesOnce,
  cutCooldown,
  litTalentsOf,
  loadApl,
  markRaidBuildRows,
  raidBuildOnlyTalents,
  talentNamesOf,
  specCooldowns,
  prepullCasts,
  talentCooldownCuts,
  rotationListsFrom,
  spellLookupFor,
  usesCooldownPass,
  type SpellFacts,
  type SpellLookup,
} from './rotation-lists';

const BUILD = activeBuild.build;
const REPO = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
const GUIDES_DIR = path.join(REPO, 'web/src/content/guides');

describe('clauseOf', () => {
  it('takes the phrase before the colon and drops the spell’s own name', () => {
    expect(
      clauseOf(
        'Ice Lance on a Fingers of Frost charge: the charge makes the target count as Frozen.',
        'Ice Lance',
      ),
    ).toBe('On a Fingers of Frost charge');
  });

  it('drops a leading "is" and the final stop', () => {
    expect(clauseOf('Frostbolt is the filler between Fingers of Frost charges.', 'Frostbolt')).toBe(
      'The filler between Fingers of Frost charges',
    );
  });

  it('keeps the Fury Death Wish rule whole: it fits the clause cap in one sentence (round 20)', () => {
    // Round 19 found the card cut the rule to "At the pull only if a second one fits"; the
    // curated note now carries the whole rule in one sentence under MAX_WHOLE_CLAUSE.
    const note =
      'At the pull when a second one still fits, otherwise in the execute phase, or in the last 30 seconds when that phase is shorter than the buff.';
    expect(clauseOf(note, 'Death Wish')).toBe(
      'At the pull when a second one still fits, otherwise in the execute phase, or in the last 30 seconds when that phase is shorter than the buff',
    );
  });

  it('cuts at the first sentence or semicolon', () => {
    expect(clauseOf('Moonfire only to clip the last seconds; maintaining it costs more.', 'Moonfire')).toBe(
      'Only to clip the last seconds',
    );
    expect(clauseOf('Seal of Fury kept up. Then something else.', 'Seal of Fury')).toBe('Kept up');
  });

  it('never ends inside a parenthesis the cut left open', () => {
    expect(
      clauseOf(
        'Mana Spring Totem restores 10 mana every 2 seconds for its five-minute life (client data; Era is 60 s).',
        'Mana Spring Totem',
      ),
    ).toBe('Restores 10 mana every 2 seconds for its five-minute life');
  });

  it('drops a leading parenthetical and comma', () => {
    expect(
      clauseOf(
        'Shifting Power (Feral talent, 16 s cooldown) on cooldown whenever there is room',
        'Shifting Power',
      ),
    ).toBe('On cooldown whenever there is room');
    expect(clauseOf('Templar’s Bulwark, an absorb for 8 seconds', 'Templar’s Bulwark')).toBe(
      'An absorb for 8 seconds',
    );
  });

  it('cuts a long clause at a word and ends it with an ellipsis', () => {
    const clause = clauseOf(`Spell ${'word '.repeat(60)}`, 'Spell');
    expect(clause.length).toBeLessThanOrEqual(100);
    expect(clause.endsWith('…')).toBe(true);
  });

  it('keeps a one-sentence clause whole, however it wraps', () => {
    expect(
      clauseOf(
        'Rapid Fire on the pull with Bestial Wrath, as an auto shot leaves and Multi-Shot is about to come back.',
        'Rapid Fire',
      ),
    ).toBe('On the pull with Bestial Wrath, as an auto shot leaves and Multi-Shot is about to come back');
  });

  it('cuts a run-on clause at its last comma when there is one, with no ellipsis', () => {
    const clause = clauseOf(
      `Rapid Fire timed to land with an incoming auto shot just as Aimed Shot is about to come off cooldown, so both benefit ${'and more '.repeat(10)}`,
      'Rapid Fire',
    );
    expect(clause).toBe(
      'Timed to land with an incoming auto shot just as Aimed Shot is about to come off cooldown',
    );
  });

  it('is empty when the note is', () => {
    expect(clauseOf('', 'Spell')).toBe('');
  });
});

describe('the APL readers', () => {
  const apl = {
    rotation: {
      prepullActions: [
        { action: { castSpell: { spellId: { spellId: 2 } } }, doAtValue: { const: { val: '-1s' } } },
        { action: { castSpell: { spellId: { spellId: 1 } } }, doAtValue: { const: { val: '-3.5s' } } },
        { action: { move: {} }, doAtValue: { const: { val: '-15s' } } },
      ],
      priorityList: [{ action: { autocastOtherCooldowns: {} } }],
    },
  };

  it('lists the pre-pull casts, earliest first, and skips what is not a cast', () => {
    expect(prepullCasts(apl)).toEqual([
      { spellId: 1, secondsBefore: 3.5, gates: [] },
      { spellId: 2, secondsBefore: 1, gates: [] },
    ]);
  });

  it('reads a pre-pull cast gated on a known aura, either way round', () => {
    const known = { auraIsKnown: { auraId: { spellId: 9 } } };
    const gated = {
      rotation: {
        prepullActions: [
          {
            action: { condition: known, castSpell: { spellId: { spellId: 3 } } },
            doAtValue: { const: { val: '-1.5s' } },
          },
          {
            action: { condition: { not: { val: known } }, castSpell: { spellId: { spellId: 4 } } },
            doAtValue: { const: { val: '-1.5s' } },
          },
        ],
      },
    };
    expect(prepullCasts(gated)).toEqual([
      { spellId: 3, secondsBefore: 1.5, gates: [{ kind: 'known', spellId: 9, positive: true }] },
      { spellId: 4, secondsBefore: 1.5, gates: [{ kind: 'known', spellId: 9, positive: false }] },
    ]);
    const names: Record<number, SpellFacts> = {
      3: { name: 'Seal of Command', cooldownMs: 0 },
      4: { name: 'Seal of Righteousness', cooldownMs: 0 },
      9: { name: 'Unsealed Wrath', cooldownMs: 0 },
    };
    const { opener } = rotationListsFrom([], gated, (id) => names[id]);
    expect(opener.map((item) => item.clause)).toEqual([
      '1.5 s before the pull, with Unsealed Wrath',
      '1.5 s before the pull, without Unsealed Wrath',
    ]);
  });

  it('reads gates combined under `and`, and lists only the Seals that go out before the pull', () => {
    const known = { auraIsKnown: { auraId: { spellId: 9 } } };
    const judged = {
      auraIsActive: { auraId: { spellId: 7, rank: 6 }, sourceUnit: { type: 'CurrentTarget' } },
    };
    const at = { const: { val: '-1.5s' } };
    const seals = {
      rotation: {
        prepullActions: [
          {
            action: { condition: { not: { val: judged } }, castSpell: { spellId: { spellId: 2 } } },
            doAtValue: at,
          },
          {
            action: { condition: { and: { vals: [known, judged] } }, castSpell: { spellId: { spellId: 3 } } },
            doAtValue: at,
          },
          {
            action: {
              condition: { and: { vals: [{ not: { val: known } }, judged] } },
              castSpell: { spellId: { spellId: 4 } },
            },
            doAtValue: at,
          },
        ],
      },
    };
    const names: Record<number, SpellFacts> = {
      2: { name: 'Seal of the Crusader', cooldownMs: 0 },
      3: { name: 'Seal of Command', cooldownMs: 0 },
      4: { name: 'Seal of Righteousness', cooldownMs: 0 },
      7: { name: 'Judgement of the Crusader', cooldownMs: 0 },
      9: { name: 'Unsealed Wrath', cooldownMs: 0 },
    };
    // Command and Righteousness wait for Judgement of the Crusader, which no pre-pull cast puts
    // on the target: only Crusader goes out before the pull.
    expect(
      rotationListsFrom([], seals, (id) => names[id]).opener.map((item) => [item.name, item.clause]),
    ).toEqual([
      ['Seal of the Crusader', '1.5 s before the pull, while Judgement of the Crusader is not on the target'],
    ]);
    // An aura an earlier pre-pull cast applies is there to gate on.
    const marked = {
      rotation: {
        prepullActions: [
          { action: { castSpell: { spellId: { spellId: 7 } } }, doAtValue: { const: { val: '-3s' } } },
          { action: { condition: judged, castSpell: { spellId: { spellId: 2 } } }, doAtValue: at },
        ],
      },
    };
    expect(rotationListsFrom([], marked, (id) => names[id]).opener.map((item) => item.name)).toEqual([
      'Judgement of the Crusader',
      'Seal of the Crusader',
    ]);
  });

  it('names a talent gate on a priority row its note does not already name', () => {
    const line = (name: string, condition: string, gates: AplGate[]): CollapsedLine => ({
      spellId: name.length,
      name,
      condition,
      truncatedAtSource: false,
      steps: 1,
      gates,
    });
    const heating: AplGate = { kind: 'known', spellId: 9, positive: true };
    const names: Record<number, SpellFacts> = { 9: { name: 'Heating Up', cooldownMs: 0 } };
    const { priority } = rotationListsFrom(
      [
        line('Fireball', '', [heating]),
        line('Pyroblast', 'On 3 stacks of Heating Up', [heating]),
        line('Frostbolt', 'Filler', []),
      ],
      undefined,
      (id) => names[id],
    );
    expect(priority.map((item) => item.clause)).toEqual([
      'With Heating Up',
      'On 3 stacks of Heating Up',
      'Filler',
    ]);
  });

  // Sim-accuracy round 19: a card describes one build and lists only the rows that build
  // casts, so a talent gate that build settles is not worded ("Fireball: Fills in a raid once
  // you stop blasting, with Heating Up" read as a condition on the stacks Fireball builds).
  it('words no talent gate the described build settles', () => {
    const heating: AplGate = { kind: 'known', spellId: 9, positive: true };
    const judged: AplGate = { kind: 'active', spellId: 7, positive: true, onTarget: true };
    const names: Record<number, SpellFacts> = {
      3: { name: 'Pyroblast', cooldownMs: 0 },
      7: { name: 'Judgement of the Crusader', cooldownMs: 0 },
      9: { name: 'Heating Up', cooldownMs: 0 },
    };
    const known = { auraIsKnown: { auraId: { spellId: 9 } } };
    const apl = {
      rotation: {
        prepullActions: [
          {
            action: { condition: known, castSpell: { spellId: { spellId: 3 } } },
            doAtValue: { const: { val: '-6s' } },
          },
        ],
      },
    };
    const line = (name: string, condition: string, gates: AplGate[]): CollapsedLine => ({
      spellId: name.length,
      name,
      condition,
      truncatedAtSource: false,
      steps: 1,
      gates,
    });
    const settled = (gate: AplGate): boolean => gate.kind === 'known' && gate.spellId === 9;
    const { opener, priority } = rotationListsFrom(
      [line('Fireball', 'Fills in a raid', [heating]), line('Seal', 'Filler', [heating, judged])],
      apl,
      (id) => names[id],
      [],
      new Map(),
      [],
      'dps',
      new Map(),
      () => true,
      settled,
    );
    expect(opener.map((item) => item.clause)).toEqual(['6 s before the pull']);
    // An aura gate is not a talent gate: the note explains it, so it is never worded here.
    expect(priority.map((item) => item.clause)).toEqual(['Fills in a raid', 'Filler']);
  });

  it('knows when the cooldown pass is on', () => {
    expect(usesCooldownPass(apl)).toBe(true);
    expect(usesCooldownPass({ rotation: { priorityList: [] } })).toBe(false);
    expect(usesCooldownPass(undefined)).toBe(false);
  });
});

describe('rotationListsFrom', () => {
  const lines: CollapsedLine[] = [
    {
      spellId: 10,
      name: 'Bestial Wrath',
      icon: 'icon_bw',
      condition: 'Bestial Wrath on cooldown',
      truncatedAtSource: false,
      steps: 1,
    },
    {
      spellId: 11,
      name: 'Aimed Shot',
      rank: 6,
      icon: 'icon_as',
      condition: 'Aimed Shot on cooldown, timed so its cast does not clip the next auto shot',
      truncatedAtSource: false,
      steps: 1,
    },
  ];
  const facts: Record<number, SpellFacts> = {
    1: { name: 'Aspect of the Hawk', rank: 5, icon: 'icon_hawk', cooldownMs: 0 },
    10: { name: 'Bestial Wrath', icon: 'icon_bw', cooldownMs: 120_000 },
    11: { name: 'Aimed Shot', rank: 6, icon: 'icon_as', cooldownMs: 6_000 },
  };
  const apl = {
    rotation: {
      prepullActions: [
        { action: { castSpell: { spellId: { spellId: 1 } } }, doAtValue: { const: { val: '-10s' } } },
      ],
      priorityList: [{ action: { autocastOtherCooldowns: {} } }],
    },
  };

  it('numbers the priority with an icon, a name, a rank and one clause per line', () => {
    const { priority } = rotationListsFrom(lines, apl, (id) => facts[id]);
    expect(priority).toEqual([
      { spellId: 10, name: 'Bestial Wrath', rank: undefined, icon: 'icon_bw', clause: 'On cooldown' },
      {
        spellId: 11,
        name: 'Aimed Shot',
        rank: 6,
        icon: 'icon_as',
        clause: 'On cooldown, timed so its cast does not clip the next auto shot',
      },
    ]);
  });

  it('builds the opener from the pre-pull casts alone', () => {
    const { opener } = rotationListsFrom(lines, apl, (id) => facts[id]);
    expect(opener.map((item) => item.name)).toEqual(['Aspect of the Hawk']);
    expect(opener[0]).toMatchObject({ rank: 5, icon: 'icon_hawk', clause: '10 s before the pull' });
  });

  it('says "On the pull" for a cast at the pull itself', () => {
    const atPull = {
      rotation: {
        prepullActions: [
          { action: { castSpell: { spellId: { spellId: 11 } } }, doAtValue: { const: { val: '0s' } } },
        ],
      },
    };
    const { opener } = rotationListsFrom(lines, atPull, (id) => facts[id]);
    expect(opener.map((item) => item.clause)).toEqual(['On the pull']);
  });

  it('lists as cooldowns the lines whose spell has a cooldown of half a minute or more', () => {
    const { cooldowns } = rotationListsFrom(lines, apl, (id) => facts[id]);
    expect(cooldowns.map((item) => item.name)).toEqual(['Bestial Wrath', 'Trinkets, racials and potions']);
    expect(cooldowns[0]!.clause).toBe('2 min cooldown');
  });

  it('has no opener, and no cooldowns, with no APL and nothing else to go on', () => {
    const lists = rotationListsFrom([], undefined, () => undefined);
    expect(lists.opener).toEqual([]);
    expect(lists.cooldowns).toEqual([]);
  });

  it('adds the spec’s own cooldown spells after the priority’s, before the trinket line', () => {
    const { cooldowns } = rotationListsFrom(lines, apl, (id) => facts[id], [
      { name: 'Evocation', clause: '8 min cooldown' },
      { name: 'Bestial Wrath', clause: 'ignored: already named by the priority' },
    ]);
    expect(cooldowns.map((item) => item.name)).toEqual([
      'Bestial Wrath',
      'Evocation',
      'Trinkets, racials and potions',
    ]);
  });

  it('names the racial cooldowns of the class’s races on the trinket line', () => {
    const { cooldowns } = rotationListsFrom(lines, apl, (id) => facts[id], [], new Map(), [
      { race: 'Gnome', racial: 'Eureka!' },
      { race: 'Troll', racial: 'Berserking' },
    ]);
    expect(cooldowns.at(-1)).toEqual({
      name: 'Trinkets, racials and potions',
      clause: 'On cooldown, with your biggest damage: Eureka! (Gnome), Berserking (Troll)',
    });
  });

  it('keeps the trinket line plain when no race of the class has a racial cooldown', () => {
    const { cooldowns } = rotationListsFrom(lines, apl, (id) => facts[id], [], new Map(), []);
    expect(cooldowns.at(-1)!.clause).toBe('On cooldown, with your biggest damage');
  });

  it('words the trinket line for a healer by healing, not damage', () => {
    const { cooldowns } = rotationListsFrom(
      lines,
      apl,
      (id) => facts[id],
      [],
      new Map(),
      [{ race: 'Gnome', racial: 'Eureka!' }],
      'healer',
    );
    expect(cooldowns.at(-1)!.clause).toBe('On cooldown, when your healing matters most: Eureka! (Gnome)');
  });
});

describe('the Frost Mage lists, from the real data', () => {
  const entry = rotationEntryFor(loadRotations(BUILD, 'mage-frost') ?? [], 60);

  it('reads the level 60 priority off the published lines', () => {
    if (entry === undefined) throw new Error('mage-frost has no level 60 rotation');
    const lines = rotationLinesFor(entry, BUILD, 'mage', 'mage-frost', 60);
    const lists = rotationListsFrom(lines, loadApl('mage-frost'), spellLookupFor(BUILD, 'mage'));
    // A floor, not the exact list: the published lines are the nightly's, and they gain the
    // rotation's newer lines (Ice Barrier, Cone of Cold) when it next republishes.
    const names = lists.priority.map((item) => item.name);
    expect(names).toEqual(expect.arrayContaining(['Ice Lance', 'Frostbolt']));
    expect(names.indexOf('Ice Lance')).toBeLessThan(names.indexOf('Frostbolt'));
    expect(lists.priority.every((item) => item.icon !== undefined && item.clause !== '')).toBe(true);
    // The opener comes straight off the APL's pre-pull casts.
    expect(lists.opener.map((item) => item.name)).toEqual(['Ice Barrier', 'Frostbolt']);
    // With no build's talents passed, Frostbolt starts its full base cast before the pull; the
    // guide's own build (Improved Frostbolt) is guide-rotation.test.ts's.
    expect(lists.opener.map((item) => item.clause)).toEqual(['4 s before the pull', '3 s before the pull']);
  });

  it('names the pre-pull casts of a spec that has them, with their icons', () => {
    const hunter = rotationEntryFor(loadRotations(BUILD, 'hunter-beast-mastery') ?? [], 60);
    if (hunter === undefined) throw new Error('hunter-beast-mastery has no level 60 rotation');
    const lines = rotationLinesFor(hunter, BUILD, 'hunter', 'hunter-beast-mastery', 60);
    const lists = rotationListsFrom(lines, loadApl('hunter-beast-mastery'), spellLookupFor(BUILD, 'hunter'));
    const names = lists.opener.map((item) => item.name);
    expect(names).toContain('Aspect of the Hawk');
    expect(names).toContain('Aimed Shot');
    expect(lists.opener.find((item) => item.name === 'Aimed Shot')?.icon).toBeDefined();
    expect(lists.cooldowns.map((item) => item.name)).toContain('Bestial Wrath');
  });
});

describe('the Demonology opener, from the real data', () => {
  it('names every pre-pull cast, talent spells included', () => {
    const lists = rotationListsFrom([], loadApl('warlock-demonology'), spellLookupFor(BUILD, 'warlock'));
    // Demonic Sacrifice and Soul Link are talent spells with no spell-constant row; the
    // opener must still name them, in time order, with their talent icons.
    expect(lists.opener.map((item) => item.name)).toEqual([
      'Summon Imp',
      'Demonic Sacrifice',
      'Summon Succubus',
      'Soul Link',
    ]);
    const soulLink = lists.opener.find((item) => item.name === 'Soul Link');
    expect(soulLink?.icon).toBeDefined();
    expect(soulLink?.clause).toBe('2 s before the pull');
  });
});

describe('talentCooldownCuts', () => {
  it('reads a lit talent’s percent cut off its own rank text, at the rank taken', () => {
    expect(talentCooldownCuts(BUILD, 'warlock', new Map([['Decimation', 2]])).get('Soul Fire')).toEqual({
      ms: 0,
      percent: 90,
    });
    expect(talentCooldownCuts(BUILD, 'warlock', new Map([['Decimation', 1]])).get('Soul Fire')).toEqual({
      ms: 0,
      percent: 45,
    });
  });

  it('ignores a talent with no points', () => {
    expect(talentCooldownCuts(BUILD, 'warlock', new Map()).size).toBe(0);
  });

  it('reads every wording the client uses: "on your", "abilities", a three-spell list, "its cooldown"', () => {
    const hunter = talentCooldownCuts(
      BUILD,
      'hunter',
      new Map([
        ['Rapid Killing', 2],
        ["Survivalist's Discipline", 2],
      ]),
    );
    expect(hunter.get('Rapid Fire')).toEqual({ ms: 120_000, percent: 0 });
    expect(hunter.get('Deterrence')).toEqual({ ms: 0, percent: 40 });
    const paladin = talentCooldownCuts(BUILD, 'paladin', new Map([['Sacred Duty', 2]]));
    expect(paladin.get("Templar's Bulwark")).toEqual({ ms: 60_000, percent: 0 });
    expect(paladin.get('Divine Shield')).toEqual({ ms: 60_000, percent: 0 });
    expect(
      talentCooldownCuts(BUILD, 'druid', new Map([['Improved Tranquility', 2]])).get('Tranquility'),
    ).toEqual({
      ms: 0,
      percent: 60,
    });
  });

  it('drops a spell whose cut cooldown falls under the floor from the Cooldowns list', () => {
    const facts: SpellLookup = (id) =>
      id === 17924 ? { name: 'Soul Fire', rank: 2, cooldownMs: 60_000 } : undefined;
    const line = {
      spellId: 17924,
      name: 'Soul Fire',
      rank: 2,
      condition: 'Soul Fire while Decimation is up.',
    };
    const apl = { rotation: { priorityList: [] } };
    expect(rotationListsFrom([line] as never, apl, facts).cooldowns.map((c) => c.name)).toEqual([
      'Soul Fire',
    ]);
    const cuts = new Map([['Soul Fire', { ms: 0, percent: 90 }]]);
    expect(rotationListsFrom([line] as never, apl, facts, [], cuts).cooldowns).toEqual([]);
  });
});

describe('specCooldowns for Frost Mage, from the client data', () => {
  const items = specCooldowns(
    BUILD,
    'mage',
    'mage-frost',
    new Map([
      ['Cold Snap', 1],
      ['Ice Block', 1],
    ]),
    spellLookupFor(BUILD, 'mage'),
  );

  it('names the baseline and the lit talent cooldowns, each with its length', () => {
    const names = items.map((item) => item.name);
    expect(names).toContain('Evocation');
    expect(names).toContain('Cold Snap');
    expect(names).toContain('Ice Block');
    expect(items.find((item) => item.name === 'Cold Snap')?.clause).toBe('10 min cooldown');
    expect(items.find((item) => item.name === 'Evocation')?.clause).toBe('8 min cooldown');
  });

  it('leaves out a talent the build does not take and a short baseline cooldown', () => {
    const names = items.map((item) => item.name);
    expect(names).not.toContain('Presence of Mind');
    expect(names).not.toContain('Arcane Power');
    expect(names).not.toContain('Blink');
  });

  it('never lists a long-cast utility spell as a fight cooldown', () => {
    const shaman = specCooldowns(
      BUILD,
      'shaman',
      'shaman-elemental',
      new Map(),
      spellLookupFor(BUILD, 'shaman'),
    ).map((item) => item.name);
    expect(shaman).not.toContain('Astral Recall');
  });

  it('leaves out race-locked and unlearnable class spells, and keeps a baseline one', () => {
    const priest = specCooldowns(
      BUILD,
      'priest',
      'priest-shadow',
      new Map(),
      spellLookupFor(BUILD, 'priest'),
    ).map((item) => item.name);
    for (const name of ['Shadowfiend', 'Feedback', 'Divine Grace', 'Dark Sacrifice', "Elune's Grace"]) {
      expect(priest).not.toContain(name);
    }
    expect(priest).toContain('Fear Ward');
  });

  it('gives Evocation its icon from the spell icon data', () => {
    expect(items.find((item) => item.name === 'Evocation')?.icon).toBe('spell_nature_purge');
  });

  it('reads a lit talent the trainer data does not list from the talent file', () => {
    const rogue = specCooldowns(
      BUILD,
      'rogue',
      'rogue-combat',
      new Map([
        ['Adrenaline Rush', 1],
        ['Cold Blood', 1],
      ]),
      spellLookupFor(BUILD, 'rogue'),
    );
    expect(rogue.find((item) => item.name === 'Adrenaline Rush')?.clause).toBe('5 min cooldown');
    expect(rogue.find((item) => item.name === 'Cold Blood')?.clause).toBe('3 min cooldown');
    expect(rogue.map((item) => item.name)).not.toContain('Preparation');
  });

  it('takes a lit talent’s icon from the talent file', () => {
    expect(items.find((item) => item.name === 'Cold Snap')?.icon).toBe('spell_frost_wizardmark');
  });
});

describe('specCooldowns keeps to the spec’s own play', () => {
  const names = (classSlug: string, spec: string): string[] =>
    specCooldowns(BUILD, classSlug, spec, new Map(), spellLookupFor(BUILD, classSlug)).map(
      (item) => item.name,
    );

  it('lists only the baseline cooldowns the spec’s curated list names', () => {
    for (const spec of ['warrior-arms', 'warrior-fury']) {
      const arms = names('warrior', spec);
      for (const name of ['Intimidating Shout', 'Challenging Shout', 'Retaliation'])
        expect(arms).not.toContain(name);
    }
    const protection = names('warrior', 'warrior-protection');
    expect(protection).not.toContain('Recklessness');
    expect(protection).toContain('Shield Wall');
  });

  // Voice of Truth (20) and Lightwell (40) are trainer spells on 70753, not removed talents:
  // they stay off only because no spec's curated row names them (Lightwell is the healing
  // priests' prepull cast, so their rotation list carries it).
  it('leaves a long baseline cooldown no curated row names off (Voice of Truth, Lightwell)', () => {
    for (const spec of ['paladin-holy', 'paladin-protection', 'paladin-retribution']) {
      expect(names('paladin', spec)).not.toContain('Voice of Truth');
    }
    for (const spec of ['priest-discipline', 'priest-holy', 'priest-shadow']) {
      expect(names('priest', spec)).not.toContain('Lightwell');
    }
  });

  it('names only trainer spells the class learns by level 60 in the curated list', () => {
    const curated = JSON.parse(
      readFileSync(path.join(REPO, 'data/curated/guide-cooldowns.json'), 'utf8'),
    ) as {
      specs: Record<string, string[]>;
    };
    for (const [spec, list] of Object.entries(curated.specs)) {
      const classSlug = spec.split('-')[0]!;
      const trainables = JSON.parse(
        readFileSync(path.join(REPO, 'data/builds', BUILD, 'trainables', `${classSlug}.json`), 'utf8'),
      ) as { trainables: { name: string; ranks: { level: number }[] }[] };
      const learnable = new Set(
        trainables.trainables.filter((t) => t.ranks.some((rank) => rank.level <= 60)).map((t) => t.name),
      );
      for (const name of list) expect(learnable.has(name), `${spec}: ${name}`).toBe(true);
    }
  });
});

describe('talent cooldown cuts', () => {
  it('reads Guardian’s Favor at the rank the build takes', () => {
    expect(
      talentCooldownCuts(BUILD, 'paladin', new Map([["Guardian's Favor", 2]])).get('Blessing of Protection'),
    ).toEqual({ ms: 120_000, percent: 0 });
    expect(
      talentCooldownCuts(BUILD, 'paladin', new Map([["Guardian's Favor", 1]])).get('Blessing of Protection'),
    ).toEqual({ ms: 60_000, percent: 0 });
    expect(talentCooldownCuts(BUILD, 'paladin', new Map()).size).toBe(0);
  });

  it('shows Holy’s Blessing of Protection at 3 min with Guardian’s Favor 2/2', () => {
    const holy = specCooldowns(
      BUILD,
      'paladin',
      'paladin-holy',
      new Map([["Guardian's Favor", 2]]),
      spellLookupFor(BUILD, 'paladin'),
    );
    expect(holy.find((item) => item.name === 'Blessing of Protection')?.clause).toBe('3 min cooldown');
  });

  it('never cuts below zero', () => {
    expect(cutCooldown(60_000, { ms: 120_000, percent: 0 })).toBe(0);
    expect(cutCooldown(60_000, undefined)).toBe(60_000);
    expect(cutCooldown(300_000, { ms: 60_000, percent: 50 })).toBe(120_000);
  });

  it('litTalentsOf keeps the highest rank either build takes', () => {
    const lit = litTalentsOf([
      {
        talents: [
          { name: 'A', rank: 1 },
          { name: 'B', rank: 0 },
        ],
      },
      { talents: [{ name: 'A', rank: 3 }] },
    ]);
    expect([...lit]).toEqual([['A', 3]]);
  });
});

/** Every spell id a rotation casts, pre-pull included, at any depth. */
function castIds(node: unknown, into: Set<number> = new Set()): Set<number> {
  if (Array.isArray(node)) {
    for (const item of node) castIds(item, into);
  } else if (typeof node === 'object' && node !== null) {
    const record = node as Record<string, unknown>;
    const cast = record.castSpell as { spellId?: { spellId?: unknown } } | undefined;
    if (typeof cast?.spellId?.spellId === 'number') into.add(cast.spellId.spellId);
    for (const value of Object.values(record)) castIds(value, into);
  }
  return into;
}

const GUIDE_SPECS = readdirSync(GUIDES_DIR, { withFileTypes: true })
  .filter((entry) => entry.isDirectory())
  .flatMap((entry) =>
    readdirSync(path.join(GUIDES_DIR, entry.name))
      .filter((file) => file.endsWith('.md') && file !== 'index.md')
      .map((file) => ({ classSlug: entry.name, spec: `${entry.name}-${file.slice(0, -'.md'.length)}` })),
  )
  .filter(({ spec }) => loadApl(spec) !== undefined);

describe('every spell a guide renders has the client’s icon', () => {
  it.each(GUIDE_SPECS.map(({ classSlug, spec }) => [spec, classSlug] as const))(
    '%s: every cast its rotation names, pre-pull included',
    (spec, classSlug) => {
      const lookup = spellLookupFor(BUILD, classSlug);
      const blank = [...castIds((loadApl(spec) as { rotation?: unknown }).rotation)]
        .map((id) => ({ id, facts: lookup(id) }))
        .filter(({ facts }) => facts !== undefined && facts.icon === undefined)
        .map(({ id, facts }) => `${facts!.name} (${id})`);
      expect(blank).toEqual([]);
    },
  );

  it.each([...new Set(GUIDE_SPECS.map(({ classSlug }) => classSlug))])(
    '%s: every cooldown the guide can list, every talent lit',
    (classSlug) => {
      const talents = JSON.parse(
        readFileSync(path.join(REPO, 'data/builds', BUILD, 'talents', `${classSlug}.json`), 'utf8'),
      ) as { trees: { talents: { name: string }[] }[] };
      const lit = new Map(
        talents.trees.flatMap((tree) => tree.talents.map((talent) => [talent.name, 1] as const)),
      );
      const blank = GUIDE_SPECS.filter((guide) => guide.classSlug === classSlug)
        .flatMap(({ spec }) => specCooldowns(BUILD, classSlug, spec, lit, spellLookupFor(BUILD, classSlug)))
        .filter((item) => item.icon === undefined)
        .map((item) => item.name);
      expect(blank).toEqual([]);
    },
  );

  it('gives the warlock pet summons their own icons', () => {
    const lookup = spellLookupFor(BUILD, 'warlock');
    expect(lookup(688)?.icon).toBe('spell_shadow_summonimp');
    expect(lookup(712)?.icon).toBe('spell_shadow_summonsuccubus');
  });
});

describe('the guide builds’ own cooldowns', () => {
  function guideCode(file: string, key: 'build' | 'raidBuild'): string {
    const raw = readFileSync(path.join(GUIDES_DIR, file), 'utf8');
    return new RegExp(`^${key}: '([^']+)'`, 'm').exec(raw)?.[1] ?? '';
  }

  it('lists Adrenaline Rush for Combat, which its builds take', () => {
    const trees = (['build', 'raidBuild'] as const).flatMap(
      (key) => decodeBuildTrees(BUILD, guideCode('rogue/combat.md', key)) ?? [],
    );
    const lit = litTalentsOf(trees);
    const names = specCooldowns(BUILD, 'rogue', 'rogue-combat', lit, spellLookupFor(BUILD, 'rogue'));
    expect(names.map((item) => item.name)).toContain('Adrenaline Rush');
  });

  it('shows Holy’s Blessing of Protection less what the Holy build’s Guardian’s Favor takes off', () => {
    const lit = litTalentsOf(decodeBuildTrees(BUILD, guideCode('paladin/holy.md', 'build')) ?? []);
    const rank = lit.get("Guardian's Favor") ?? 0;
    const holy = specCooldowns(BUILD, 'paladin', 'paladin-holy', lit, spellLookupFor(BUILD, 'paladin'));
    const want = ['5 min cooldown', '4 min cooldown', '3 min cooldown'][rank];
    expect(holy.find((item) => item.name === 'Blessing of Protection')?.clause).toBe(want);
  });
});

// Sim-accuracy round 16: the lists come from the level 60 raid priority, and the page shows
// the level 60 build beside the raid build. A row for a talent only the raid build takes is
// marked, so a player on the level 60 build knows it is not theirs.
describe('raid build rows', () => {
  it('names the talents only the raid build takes', () => {
    const only = raidBuildOnlyTalents(
      new Set(['Reckless Swings', 'Mortal Strike', 'Improved Slam']),
      new Map([
        ['Mortal Strike', 1],
        ['Improved Slam', 5],
      ]),
      new Map([
        ['Mortal Strike', 1],
        ['Reckless Swings', 1],
      ]),
    );
    expect([...only]).toEqual(['Reckless Swings']);
  });

  it('marks every list row whose spell is one of them, and no other', () => {
    const lists = {
      opener: [{ name: 'Charge', clause: '' }],
      priority: [
        { name: 'Mortal Strike', clause: 'On cooldown' },
        { name: 'Reckless Swings', clause: '' },
      ],
      cooldowns: [{ name: 'Reckless Swings', clause: '1 min cooldown' }],
    };
    const marked = markRaidBuildRows(lists, new Set(['Reckless Swings']));
    expect(marked.priority.map((item) => item.raidBuildOnly === true)).toEqual([false, true]);
    expect(marked.cooldowns[0]!.raidBuildOnly).toBe(true);
    expect(marked.opener[0]!.raidBuildOnly).toBeUndefined();
    expect(lists.priority[1]).not.toHaveProperty('raidBuildOnly');
  });

  // Round 19: with the "with Heating Up" words gone from the raid card, a row that needs a
  // raid-only talent (Arcane's Fire Blast, Fireball and raid Frostbolt need Heating Up) says
  // so with the tag, as a raid-only talent's own row does.
  it('marks a row that needs one of them, and carries what each row needs', () => {
    const heating: AplGate = { kind: 'known', spellId: 9, positive: true };
    const pom: AplGate = { kind: 'known', spellId: 8, positive: true };
    const names: Record<number, SpellFacts> = {
      8: { name: 'Presence of Mind', cooldownMs: 0 },
      9: { name: 'Heating Up', cooldownMs: 0 },
    };
    const line = (name: string, gates: AplGate[]): CollapsedLine => ({
      spellId: name.length,
      name,
      condition: '',
      truncatedAtSource: false,
      steps: 1,
      gates,
    });
    const lists = rotationListsFrom(
      [line('Fireball', [heating]), line('Frostbolt', [pom]), line('Arcane Blast', [])],
      undefined,
      (id) => names[id],
    );
    expect(lists.priority.map((item) => item.needs ?? [])).toEqual([
      ['Heating Up'],
      ['Presence of Mind'],
      [],
    ]);
    const marked = markRaidBuildRows(lists, new Set(['Heating Up']));
    expect(marked.priority.map((item) => item.raidBuildOnly === true)).toEqual([true, false, false]);
  });

  it('marks Arms’ Reckless Swings, which only the raid build takes, and not Spearing Strike, which both take', () => {
    const arms = readFileSync(path.join(GUIDES_DIR, 'warrior/arms.md'), 'utf8');
    const code = (key: string): string => new RegExp(`^${key}: '([^']+)'`, 'm').exec(arms)![1]!;
    const only = raidBuildOnlyTalents(
      talentNamesOf(BUILD, 'warrior'),
      litTalentsOf(decodeBuildTrees(BUILD, code('build')) ?? []),
      litTalentsOf(decodeBuildTrees(BUILD, code('raidBuild')) ?? []),
    );
    expect(only.has('Reckless Swings')).toBe(true);
    expect(only.has('Spearing Strike')).toBe(false);
    expect(only.has('Mortal Strike')).toBe(false);
  });
});

describe('sayRepeatedRulesOnce', () => {
  it('says a rule the row above already says once, on the first row', () => {
    const rules = 'On the tank below 80% · on a party member below 65%';
    const rows = sayRepeatedRulesOnce([
      { name: 'Greater Heal', rank: 4, clause: rules },
      { name: 'Greater Heal', rank: 1, clause: rules },
      { name: 'Heal', rank: 4, clause: rules },
      { name: 'Renew', rank: 7, clause: 'When Renew is not on the tank' },
    ]);
    expect(rows.map((row) => row.clause)).toEqual([
      rules,
      'At the same thresholds',
      'At the same thresholds',
      'When Renew is not on the tank',
    ]);
  });

  it('says a short rule again', () => {
    const rows = sayRepeatedRulesOnce([
      { name: 'Strider Kick', clause: 'On cooldown' },
      { name: 'Immolation Trap', rank: 5, clause: 'On cooldown' },
    ]);
    expect(rows.map((row) => row.clause)).toEqual(['On cooldown', 'On cooldown']);
  });
});

describe('inPageWords', () => {
  it("writes the panel's mana-against-boss shorthand out", () => {
    expect(inPageWords('on the tank below 80% while mana % is over boss health %')).toBe(
      "on the tank below 80% while your mana is ahead of the boss's health",
    );
    expect(inPageWords('when mana % is under boss health % and Innervate is down')).toBe(
      "when your mana is behind the boss's health and Innervate is down",
    );
  });
});

describe('joinSpellRuns', () => {
  it('joins a run of one spell and rank into one row, every rule kept in order', () => {
    const rows = joinSpellRuns([
      { name: 'Flash Heal', rank: 7, clause: 'On the tank below 35%' },
      { name: 'Greater Heal', rank: 4, clause: 'On the tank below 80%' },
      { name: 'Greater Heal', rank: 4, clause: 'On a party member below 65%' },
      { name: 'Greater Heal', rank: 1, clause: 'On the tank below 50%' },
      { name: 'Greater Heal', rank: 4, clause: 'On a party member below 35%' },
    ]);
    expect(rows.map((row) => `${row.name} ${row.rank}: ${row.clause}`)).toEqual([
      'Flash Heal 7: On the tank below 35%',
      'Greater Heal 4: On the tank below 80% · on a party member below 65%',
      'Greater Heal 1: On the tank below 50%',
      'Greater Heal 4: On a party member below 35%',
    ]);
  });

  it('never joins a raid-build-only row to one every build casts', () => {
    const rows = joinSpellRuns([
      { name: 'Slam', clause: 'On cooldown' },
      { name: 'Slam', clause: 'With Reckless Swings up', raidBuildOnly: true },
    ]);
    expect(rows).toHaveLength(2);
  });
});
