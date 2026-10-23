// web/src/lib/guides/rotation-lists.test.ts
import activeBuild from '../../data/active-build.json';
import { describe, expect, it } from 'vitest';
import { loadRotations, rotationEntryFor, rotationLinesFor } from '../bis/rotation-view';
import type { CollapsedLine } from '../bis/rotation-collapse';
import {
  clauseOf,
  formatCooldown,
  loadApl,
  prepullCasts,
  rotationListsFrom,
  spellLookupFor,
  usesCooldownPass,
  type SpellFacts,
} from './rotation-lists';

const BUILD = activeBuild.build;

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

  it('cuts at the first sentence or semicolon', () => {
    expect(clauseOf('Moonfire only to clip the last seconds; maintaining it costs more.', 'Moonfire')).toBe(
      'Only to clip the last seconds',
    );
    expect(clauseOf('Seal of Fury kept up. Then something else.', 'Seal of Fury')).toBe('Kept up');
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

  it('cuts a long clause at its last comma when there is one, with no ellipsis', () => {
    const clause = clauseOf(
      'Rapid Fire timed to land with an incoming auto shot just as Aimed Shot is about to come off cooldown, so both benefit',
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

describe('formatCooldown', () => {
  it('writes seconds under a minute and minutes after', () => {
    expect(formatCooldown(30_000)).toBe('30 s');
    expect(formatCooldown(120_000)).toBe('2 min');
    expect(formatCooldown(90_000)).toBe('1 min 30 s');
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
      { spellId: 1, secondsBefore: 3.5 },
      { spellId: 2, secondsBefore: 1 },
    ]);
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

  it('builds the opener from the pre-pull casts, the cooldown pass and the priority list', () => {
    const { opener } = rotationListsFrom(lines, apl, (id) => facts[id]);
    expect(opener.map((item) => item.name)).toEqual([
      'Aspect of the Hawk',
      'On-use cooldowns',
      'Priority list',
    ]);
    expect(opener[0]).toMatchObject({ rank: 5, icon: 'icon_hawk', clause: '10 s before the pull' });
  });

  it('lists as cooldowns the lines whose spell has a cooldown of half a minute or more', () => {
    const { cooldowns } = rotationListsFrom(lines, apl, (id) => facts[id]);
    expect(cooldowns.map((item) => item.name)).toEqual(['Bestial Wrath', 'Trinkets, racials and potions']);
    expect(cooldowns[0]!.clause).toBe('2 min cooldown');
  });

  it('with no APL still ends the opener on the priority list and has no cooldowns', () => {
    const lists = rotationListsFrom([], undefined, () => undefined);
    expect(lists.opener.map((item) => item.name)).toEqual(['Priority list']);
    expect(lists.cooldowns).toEqual([]);
  });
});

describe('the Frost Mage lists, from the real data', () => {
  const entry = rotationEntryFor(loadRotations(BUILD, 'mage-frost') ?? [], 60);

  it('reads the level 60 priority off the published lines', () => {
    if (entry === undefined) throw new Error('mage-frost has no level 60 rotation');
    const lines = rotationLinesFor(entry, BUILD, 'mage', 'mage-frost');
    const lists = rotationListsFrom(lines, loadApl('mage-frost'), spellLookupFor(BUILD, 'mage'));
    expect(lists.priority.map((item) => item.name)).toEqual(['Ice Lance', 'Frostbolt']);
    expect(lists.priority.every((item) => item.icon !== undefined && item.clause !== '')).toBe(true);
    expect(lists.opener.at(-1)?.name).toBe('Priority list');
  });

  it('names the pre-pull casts of a spec that has them, with their icons', () => {
    const hunter = rotationEntryFor(loadRotations(BUILD, 'hunter-beast-mastery') ?? [], 60);
    if (hunter === undefined) throw new Error('hunter-beast-mastery has no level 60 rotation');
    const lines = rotationLinesFor(hunter, BUILD, 'hunter', 'hunter-beast-mastery');
    const lists = rotationListsFrom(lines, loadApl('hunter-beast-mastery'), spellLookupFor(BUILD, 'hunter'));
    const names = lists.opener.map((item) => item.name);
    expect(names).toContain('Aspect of the Hawk');
    expect(names).toContain('Aimed Shot');
    expect(lists.opener.find((item) => item.name === 'Aimed Shot')?.icon).toBeDefined();
    expect(lists.cooldowns.map((item) => item.name)).toContain('Bestial Wrath');
  });
});
