// web/src/lib/guides/race-table.test.ts
// Runs against the active build's own race data (data/builds/<build>/, the same under either
// FOREVER_DATA), never the planner fixture: the guide table reads the build, so this does too.
// What the editorial table names is checked against the race text only where the data carries
// that text, so a build whose race sentences move fails a test instead of the page.
import activeBuild from '../../data/active-build.json';
import { describe, expect, it } from 'vitest';
import { buildRace, buildRaces, buildRacesForClass } from './build-races';
import {
  RACIAL_COOLDOWNS,
  RACIAL_EFFECTS_BY_CLASS,
  RACIAL_NAMES,
  RACIAL_PICKS,
  RACIAL_PICKS_BY_CLASS,
  racialCooldownsFor,
  racialEffects,
  racialIconsFor,
  raceNameFor,
  raceTableFor,
  proseNamesRace,
  recommendedRaceLines,
  specKindOf,
  type SpecKind,
} from './race-table';

const BUILD = activeBuild.build;
const raceRows = buildRaces(BUILD);
const mageRaces = buildRacesForClass(BUILD, 'mage');
const raceBy = (slug: string) => buildRace(BUILD, slug);
const CLASSES = ['druid', 'hunter', 'mage', 'paladin', 'priest', 'rogue', 'shaman', 'warlock', 'warrior'];
const classesOf = (slug: string): string[] =>
  CLASSES.filter((classSlug) => buildRacesForClass(BUILD, classSlug).some((race) => race.slug === slug));

describe('the editorial racial tables against the race data', () => {
  const allPicks = (slug: string): string[] => [
    ...Object.values(RACIAL_PICKS[slug] ?? {}).flat(),
    ...Object.values(RACIAL_PICKS_BY_CLASS[slug] ?? {}).flat(),
  ];

  it('names only picks that are among the race’s own racials', () => {
    for (const slug of new Set([...Object.keys(RACIAL_PICKS), ...Object.keys(RACIAL_PICKS_BY_CLASS)])) {
      for (const pick of allPicks(slug)) {
        expect(RACIAL_NAMES[slug], slug).toContain(pick);
      }
    }
  });

  it('finds, with an effect, every racial a pick names wherever the data describes that race’s racials', () => {
    for (const race of raceRows) {
      const effects = racialEffects(race);
      if (effects.size === 0) continue;
      for (const pick of new Set(allPicks(race.slug))) {
        if (race.forever_changes.some((change) => change.text.includes(pick))) {
          expect(effects.get(pick), `${race.slug} ${pick}`).toBeTruthy();
        }
      }
    }
  });
});

describe('racialEffects', () => {
  it('cuts each described racial’s effect out of the data sentence, without the next racial’s name', () => {
    for (const race of raceRows) {
      const effects = racialEffects(race);
      for (const [name, effect] of effects) {
        expect(effect, `${race.slug} ${name}`).not.toBe('');
        for (const other of RACIAL_NAMES[race.slug] ?? []) {
          if (other !== name && !name.includes(other))
            expect(effect, `${race.slug} ${name}`).not.toContain(other);
        }
      }
    }
  });

  it('reads a known racial’s effect in the data’s own words when the data has it', () => {
    const troll = raceBy('troll');
    const berserking = troll === undefined ? undefined : racialEffects(troll).get('Berserking');
    if (berserking !== undefined) expect(berserking).toMatch(/casting and attack speed/);
  });

  it('describes the racials as Blizzard retuned them on 26 and 27 October', () => {
    const effect = (slug: string, name: string) => {
      const race = raceBy(slug);
      return race === undefined ? undefined : racialEffects(race).get(name);
    };
    expect(effect('orc', 'Command')).toMatch(/Hunter and Warlock pets by 5%/);
    expect(effect('troll', 'Berserking')).toMatch(/for 15 seconds/);
    expect(effect('undead', 'Touch of the Grave')).toMatch(/once every 15 seconds/);
    expect(effect('human', 'The Human Spirit')).toMatch(/Spirit by 10%/);
    expect(effect('high-order-skyborne', 'Wind Blessed')).toMatch(/2% haste/);
    expect(effect('dwarf', "Artisan's Precision")).toMatch(/hit chance with spells and attacks by 1%/);
    expect(effect('tauren', 'Endurance')).toMatch(/Health by 5%/);
    expect(effect('tauren', 'Endurance')).not.toMatch(/Hit/);
    expect(effect('windshaper-skyborne', 'Skysight')).toMatch(
      /15% movement and mounted speed for 30 seconds/,
    );
  });
});

describe('specKindOf', () => {
  it('sorts casters, melee, hunters, healers and tanks', () => {
    expect(specKindOf('mage', 'mage-frost', 'dps')).toBe('caster');
    expect(specKindOf('priest', 'priest-shadow', 'dps')).toBe('caster');
    expect(specKindOf('priest', 'priest-holy', 'healer')).toBe('healer');
    expect(specKindOf('warrior', 'warrior-protection', 'tank')).toBe('tank');
    expect(specKindOf('warrior', 'warrior-fury', 'dps')).toBe('melee');
    expect(specKindOf('hunter', 'hunter-marksmanship', 'dps')).toBe('ranged');
    expect(specKindOf('hunter', 'hunter-survival', 'dps')).toBe('melee');
  });
});

describe('raceTableFor', () => {
  it('has one row per race the class can be, in the data’s order, with the recommended ones marked', () => {
    const slugs = mageRaces.map((race) => race.slug);
    const recommended = slugs.filter((slug) => ['gnome', 'troll'].includes(slug));
    const rows = raceTableFor(BUILD, 'mage', ['gnome', 'troll'], 'caster');
    expect(rows.map((row) => row.slug)).toEqual(slugs);
    expect(rows.filter((row) => row.recommended).map((row) => row.slug)).toEqual(recommended);
  });

  it('lists the build’s races, not the planner fixture’s: Forever’s new mage combos are rows', () => {
    // The two combos Forever added for mages (Deep Dive); the fixture has only the four classic
    // ones, which is what the Frost Mage e2e walk once rendered under FOREVER_DATA=fixture.
    const slugs = raceTableFor(BUILD, 'mage', [], 'caster').map((row) => row.slug);
    expect(slugs).toEqual(expect.arrayContaining(['orc', 'high-order-skyborne']));
  });

  it('gives a row the racials the data describes, and none when it describes none', () => {
    for (const row of raceTableFor(BUILD, 'mage', [], 'caster')) {
      const race = raceBy(row.slug)!;
      const effects = racialEffects(race);
      if (effects.size === 0) expect(row.racials, row.slug).toEqual([]);
      else {
        expect(row.racials.length, row.slug).toBeGreaterThan(0);
        for (const racial of row.racials)
          expect(effects.has(racial.name), `${row.slug} ${racial.name}`).toBe(true);
      }
      expect(raceNameFor(BUILD, row.slug)).toBe(race.name);
    }
  });

  const racialNames = (classSlug: string, kind: SpecKind, slug: string): string[] | undefined =>
    raceTableFor(BUILD, classSlug, [], kind)
      .find((row) => row.slug === slug)
      ?.racials.map((racial) => racial.name);
  const described = (slug: string, name: string): boolean => {
    const race = raceBy(slug);
    return race !== undefined && racialEffects(race).has(name);
  };

  it('gives every Gnome spec Eureka!, its offensive racial, first', () => {
    if (!described('gnome', 'Eureka!')) return;
    expect(racialNames('mage', 'caster', 'gnome')).toEqual(['Eureka!', 'Expansive Mind']);
    expect(racialNames('priest', 'healer', 'gnome')).toEqual(['Eureka!', 'Expansive Mind']);
    expect(racialNames('rogue', 'melee', 'gnome')).toEqual(['Eureka!', 'Expansive Mind']);
    expect(racialNames('warrior', 'tank', 'gnome')).toEqual(['Eureka!', 'Escape Artist']);
  });

  it('describes Eureka! as the client has it, not as the BlizzCon demo read it', () => {
    const gnome = raceBy('gnome');
    const eureka = gnome === undefined ? undefined : racialEffects(gnome).get('Eureka!');
    if (eureka === undefined) return;
    expect(eureka).toMatch(/cost 10% less and deal 10% more for 8 seconds on a 2 minute cooldown/);
    expect(eureka).not.toMatch(/next 3 abilities|50%/);
  });

  const effectOn = (classSlug: string, kind: SpecKind, slug: string, name: string): string | undefined =>
    raceTableFor(BUILD, classSlug, [], kind)
      .find((row) => row.slug === slug)
      ?.racials.find((racial) => racial.name === name)?.effect;

  it('words Eureka! the way each class’s own tooltip does: its resource, and no healing but a priest’s', () => {
    if (!described('gnome', 'Eureka!')) return;
    expect(effectOn('warrior', 'melee', 'gnome', 'Eureka!')).toBe(
      'Makes your non-periodic damaging abilities cost 10% less Rage and deal 10% more damage for 8 seconds on a 2 minute cooldown',
    );
    expect(effectOn('rogue', 'melee', 'gnome', 'Eureka!')).toMatch(
      /non-periodic damaging abilities cost 10% less Energy/,
    );
    expect(effectOn('mage', 'caster', 'gnome', 'Eureka!')).toMatch(
      /^Makes your damaging abilities cost 10% less Mana/,
    );
    expect(effectOn('warlock', 'caster', 'gnome', 'Eureka!')).toMatch(
      /non-periodic damaging abilities cost 10% less Mana/,
    );
    expect(effectOn('priest', 'healer', 'gnome', 'Eureka!')).toMatch(
      /non-periodic damaging and healing abilities cost 10% less Mana and deal 10% more damage or healing/,
    );
    for (const classSlug of ['warrior', 'rogue', 'mage', 'warlock'])
      expect(effectOn(classSlug, 'melee', 'gnome', 'Eureka!'), classSlug).not.toMatch(/heal/i);
  });

  it('keeps every per-class wording to the race data’s duration and cooldown, and covers every class the race can be', () => {
    for (const [slug, byName] of Object.entries(RACIAL_EFFECTS_BY_CLASS)) {
      const race = raceBy(slug);
      if (race === undefined) continue;
      const effects = racialEffects(race);
      for (const [name, byClass] of Object.entries(byName)) {
        const data = effects.get(name);
        if (data === undefined) continue;
        const timing = data.match(/for \d+ seconds on a \d+ minute cooldown/)?.[0];
        expect(timing, `${slug} ${name}`).toBeDefined();
        for (const text of Object.values(byClass)) expect(text, `${slug} ${name}`).toContain(timing);
        expect(Object.keys(byClass).sort(), `${slug} ${name}`).toEqual(classesOf(slug));
      }
    }
  });

  it('cuts a racial’s history off its effect: Artisan’s Precision is the 1% hit alone', () => {
    if (!described('dwarf', "Artisan's Precision")) return;
    const effect = racialEffects(raceBy('dwarf')!).get("Artisan's Precision");
    expect(effect).not.toMatch(/in place of|October/);
  });

  it('shows the racial the sim credits, never a utility one, where the race has one for the spec', () => {
    if (described('human', 'Ruthless')) {
      expect(racialNames('mage', 'caster', 'human')).toEqual(['Ruthless', 'The Human Spirit']);
      expect(racialNames('hunter', 'ranged', 'human')).toEqual(['Ruthless']);
    }
    if (described('undead', 'Touch of the Grave')) {
      expect(racialNames('mage', 'caster', 'undead')).toEqual(['Touch of the Grave']);
      expect(racialNames('warrior', 'tank', 'undead')).toEqual([
        'Touch of the Grave',
        'Will of the Forsaken',
      ]);
      // Touch of the Grave procs on damage only: a healer's credited racial is none, so the
      // fear break leads.
      expect(racialNames('priest', 'healer', 'undead')).toEqual(['Will of the Forsaken']);
    }
    if (described('dwarf', "Artisan's Precision")) {
      expect(racialNames('priest', 'caster', 'dwarf')).toEqual(["Artisan's Precision"]);
      expect(racialNames('priest', 'healer', 'dwarf')).toEqual(['Stoneform']);
    }
    if (described('human', 'Ruthless'))
      for (const kind of ['caster', 'healer', 'melee', 'ranged', 'tank'] as const)
        expect(racialNames('priest', kind, 'human')?.[0], kind).toBe('Ruthless');
    if (described('troll', 'Berserking'))
      expect(racialNames('warrior', 'tank', 'troll')).toEqual(['Berserking']);
    if (described('orc', 'Command')) {
      expect(racialNames('warlock', 'caster', 'orc')).toEqual(['Blood Fury', 'Command']);
      expect(racialNames('hunter', 'ranged', 'orc')).toEqual(['Blood Fury', 'Command']);
      expect(racialNames('mage', 'caster', 'orc')).toEqual(['Blood Fury']);
    }
    if (described('high-order-skyborne', 'Wind Blessed')) {
      expect(racialNames('mage', 'caster', 'high-order-skyborne')).toEqual(['Wind Blessed', 'Read Leyline']);
      expect(racialNames('warrior', 'melee', 'high-order-skyborne')).toEqual(['Wind Blessed']);
    }
  });

  it('falls back to the race’s first described racial when the editorial pick is not in the data', () => {
    const picks = RACIAL_PICKS.troll;
    const troll = raceBy('troll');
    if (picks === undefined || troll === undefined) return;
    const effects = racialEffects(troll);
    if (effects.size === 0) return;
    const saved = [...picks.default];
    picks.default.splice(0, picks.default.length, 'No Such Racial');
    try {
      const row = raceTableFor(BUILD, 'mage', [], 'caster').find((r) => r.slug === 'troll');
      if (row === undefined) return;
      expect(row.racials).toHaveLength(1);
      expect(row.racials[0]!.name).toBe([...effects.keys()][0]);
    } finally {
      picks.default.splice(0, picks.default.length, ...saved);
    }
  });

  it('is empty for a class it does not know', () => {
    expect(raceTableFor(BUILD, 'nonexistent', [], 'melee')).toEqual([]);
  });

  it('shows one race the same racials for one kind of spec on every class’s page', () => {
    const kinds: SpecKind[] = ['caster', 'melee', 'ranged', 'healer', 'tank'];
    for (const kind of kinds) {
      const seen = new Map<string, { classSlug: string; names: string }>();
      for (const classSlug of CLASSES) {
        // A pick written for one class (RACIAL_PICKS_BY_CLASS) is a racial whose client row is
        // that class's alone, Orc Command on Hunter and Warlock pets: it is the one allowed
        // difference, so it is held out of the comparison.
        const gated = new Set(Object.keys(RACIAL_PICKS_BY_CLASS));
        for (const row of raceTableFor(BUILD, classSlug, [], kind)) {
          if (gated.has(row.slug) && RACIAL_PICKS_BY_CLASS[row.slug]?.[classSlug] !== undefined) continue;
          const names = row.racials.map((racial) => racial.name).join(' + ');
          const first = seen.get(row.slug);
          if (first === undefined) seen.set(row.slug, { classSlug, names });
          else
            expect(names, `${row.slug} ${kind}: ${classSlug} against ${first.classSlug}`).toBe(first.names);
        }
      }
    }
  });
});

describe('the race text against the client’s spell rows', () => {
  const effect = (slug: string, name: string): string | undefined => {
    const race = raceBy(slug);
    return race === undefined ? undefined : racialEffects(race).get(name);
  };

  it('reads the numbers the 70753 client carries, not the demo’s or an outlet’s', () => {
    if (effect('orc', 'Shatter Curse') === undefined) return;
    expect(effect('orc', 'Shatter Curse')).toMatch(/by 5% for 8 seconds on a 3 minute cooldown/);
    expect(effect('tauren', 'War Stomp')).toMatch(/for 2 seconds on a 90 second cooldown/);
    expect(effect('troll', 'Regeneration')).toMatch(/by 20%/);
    expect(effect('night-elf', 'Quickness')).toMatch(/dodge chance by 1% and movement speed by 2%/);
    expect(effect('night-elf', 'Shadowmeld')).toMatch(/10 second cooldown/);
    expect(effect('gnome', 'Escape Artist')).toMatch(/for 3 seconds on a 2 minute cooldown/);
    expect(effect('dwarf', 'Stoneform')).toMatch(/by 10% for 8 seconds on a 3 minute cooldown$/);
    expect(effect('dwarf', 'Big Game Hunter')).toMatch(/Beasts by 5%/);
  });

  it('reads Read Leyline as the client spells it, whichever spacing the race data uses', () => {
    if (raceBy('high-order-skyborne') === undefined) return;
    expect(effect('high-order-skyborne', 'Read Leyline')).toMatch(
      /^50% health and mana regeneration for 15 seconds/,
    );
  });
});

describe('racialCooldownsFor', () => {
  it('names only racials among the race’s own', () => {
    for (const [slug, racial] of Object.entries(RACIAL_COOLDOWNS))
      expect(RACIAL_NAMES[slug]).toContain(racial);
  });

  it('lists the racial cooldown of each race the class can be, in the build’s race order', () => {
    const cooldowns = racialCooldownsFor(BUILD, 'mage');
    const expected = mageRaces
      .filter((race) => RACIAL_COOLDOWNS[race.slug] !== undefined)
      .map((race) => ({ race: race.name, racial: RACIAL_COOLDOWNS[race.slug]! }));
    expect(cooldowns).toEqual(expected);
    expect(cooldowns).toContainEqual({ race: 'Gnome', racial: 'Eureka!' });
  });

  it('gives a warrior Eureka!, Blood Fury, Berserking and Elune’s Light', () => {
    expect(racialCooldownsFor(BUILD, 'warrior').map((entry) => entry.racial)).toEqual(
      expect.arrayContaining(['Eureka!', 'Blood Fury', 'Berserking', "Elune's Light"]),
    );
  });

  it('is empty for a class it does not know', () => {
    expect(racialCooldownsFor(BUILD, 'nonexistent')).toEqual([]);
  });
});

describe('recommendedRaceLines', () => {
  it('lists each recommended race that has racials, with their names', () => {
    const rows = raceTableFor(BUILD, 'mage', ['gnome', 'troll'], 'caster');
    const lines = recommendedRaceLines(rows);
    for (const line of lines) {
      const row = rows.find((r) => r.slug === line.slug)!;
      expect(row.recommended).toBe(true);
      expect(line.names).toEqual(row.racials.map((racial) => racial.name));
    }
    expect(lines.length).toBe(rows.filter((row) => row.recommended && row.racials.length > 0).length);
  });

  it('leaves out a race the guide’s own prose already names, so the pick is said once', () => {
    const rows = raceTableFor(BUILD, 'mage', ['gnome', 'undead'], 'caster');
    const prose = '**Gnome** is the best Alliance Arcane mage: Eureka! is 10% more damage.';
    expect(recommendedRaceLines(rows, prose).map((line) => line.slug)).toEqual(['undead']);
  });
});

describe('proseNamesRace', () => {
  it('matches the race’s whole name only', () => {
    expect(proseNamesRace('On the Horde, **Undead** is the pick.', 'Undead')).toBe(true);
    expect(proseNamesRace('Night Elf is the Alliance pick.', 'Night Elf')).toBe(true);
    expect(proseNamesRace('An elf is not named here.', 'Night Elf')).toBe(false);
    expect(proseNamesRace('Orcish axes', 'Orc')).toBe(false);
  });
});

describe('raceNameFor', () => {
  it('turns a slug into the race’s name, and passes an unknown slug through', () => {
    const [first] = raceRows;
    expect(raceNameFor(BUILD, first!.slug)).toBe(first!.name);
    expect(raceNameFor(BUILD, 'who')).toBe('who');
  });
});

describe('racial icons', () => {
  const classes = ['druid', 'hunter', 'mage', 'paladin', 'priest', 'rogue', 'shaman', 'warlock', 'warrior'];
  const kinds = ['caster', 'melee', 'ranged', 'healer', 'tank'] as const;

  it('gives every racial a Races table shows the client’s own icon', () => {
    const blank = new Set<string>();
    for (const classSlug of classes) {
      for (const kind of kinds) {
        for (const row of raceTableFor(BUILD, classSlug, [], kind)) {
          for (const racial of row.racials)
            if (racial.icon === undefined) blank.add(`${row.slug}: ${racial.name}`);
        }
      }
    }
    expect([...blank]).toEqual([]);
  });

  it('matches a name the client spells differently (Read Leyline)', () => {
    expect(racialIconsFor(BUILD).get('readleyline')).toBe('ability_mage_incantersabsorbtion');
    expect(racialIconsFor(BUILD).get('eureka')).toBe('inv_gnometoy');
  });
});
