// web/src/lib/guides/rotation-lists.test.ts
import activeBuild from '../../data/active-build.json';
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { loadRotations, rotationEntryFor, rotationLinesFor } from '../bis/rotation-view';
import type { CollapsedLine } from '../bis/rotation-collapse';
import {
  clauseOf,
  formatCooldown,
  loadApl,
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
});

describe('the Frost Mage lists, from the real data', () => {
  const entry = rotationEntryFor(loadRotations(BUILD, 'mage-frost') ?? [], 60);

  it('reads the level 60 priority off the published lines', () => {
    if (entry === undefined) throw new Error('mage-frost has no level 60 rotation');
    const lines = rotationLinesFor(entry, BUILD, 'mage', 'mage-frost');
    const lists = rotationListsFrom(lines, loadApl('mage-frost'), spellLookupFor(BUILD, 'mage'));
    // A floor, not the exact list: the published lines are the nightly's, and they gain the
    // rotation's newer lines (Ice Barrier, Cone of Cold) when it next republishes.
    const names = lists.priority.map((item) => item.name);
    expect(names).toEqual(expect.arrayContaining(['Ice Lance', 'Frostbolt']));
    expect(names.indexOf('Ice Lance')).toBeLessThan(names.indexOf('Frostbolt'));
    expect(lists.priority.every((item) => item.icon !== undefined && item.clause !== '')).toBe(true);
    // The opener comes straight off the APL's pre-pull casts.
    expect(lists.opener.map((item) => item.name)).toEqual(['Ice Barrier']);
    expect(lists.opener[0]?.clause).toBe('1.5 s before the pull');
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
  const tree = (name: string, rank: number, descriptions: string[]) => ({
    name: 'Demonology',
    position: 1,
    points: rank,
    talents: [
      {
        id: 1,
        name,
        icon: 'x',
        max_rank: descriptions.length,
        tier: 3,
        column: 0,
        prereq_talent_id: null,
        prereq_rank: null,
        spell_id: 1,
        ranks: descriptions.map((description) => ({ spell_id: 1, description })),
        rank,
      },
    ],
  });

  it('reads a lit talent’s cooldown cut off its own rank text', () => {
    const cuts = talentCooldownCuts([
      tree('Decimation', 2, [
        'Reduces the cooldown of your Soul Fire spell by 45%. When you cast…',
        'Reduces the cooldown of your Soul Fire spell by 90%. When you cast…',
      ]),
    ] as never);
    expect(cuts.get('Soul Fire')).toBe(90);
  });

  it('ignores a talent with no points', () => {
    const cuts = talentCooldownCuts([
      tree('Decimation', 0, ['Reduces the cooldown of your Soul Fire spell by 90%.']),
    ] as never);
    expect(cuts.size).toBe(0);
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
    expect(
      rotationListsFrom([line] as never, apl, facts, [], new Map([['Soul Fire', 90]])).cooldowns,
    ).toEqual([]);
  });
});

describe('specCooldowns for Frost Mage, from the client data', () => {
  const items = specCooldowns(
    BUILD,
    'mage',
    new Set(['Cold Snap', 'Ice Block']),
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
    const shaman = specCooldowns(BUILD, 'shaman', new Set(), spellLookupFor(BUILD, 'shaman')).map(
      (item) => item.name,
    );
    expect(shaman).not.toContain('Astral Recall');
  });

  it('leaves out race-locked and unlearnable class spells, and keeps a baseline one', () => {
    const priest = specCooldowns(BUILD, 'priest', new Set(), spellLookupFor(BUILD, 'priest')).map(
      (item) => item.name,
    );
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
      new Set(['Adrenaline Rush', 'Cold Blood']),
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
      const lit = new Set(talents.trees.flatMap((tree) => tree.talents.map((talent) => talent.name)));
      const blank = specCooldowns(BUILD, classSlug, lit, spellLookupFor(BUILD, classSlug))
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
