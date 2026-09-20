import { readFile } from 'node:fs/promises';
import path from 'node:path';
import { describe, expect, it } from 'vitest';
import { createPlannerStore } from '../planner/store.svelte';
import { indexTalents, validateOrder } from '../planner/rules';
import type { ClassRow, Combo, RaceRow, TalentFile } from '../planner/types';
import { simCopy } from './copy';
import { PRESET_BUFFS, PRESET_CONSUMABLES } from './settings';
import {
  SIM_LEVEL,
  characterFromFs1,
  characterFromPlanner,
  codeForCharacterSpec,
  fromBuildDraft,
  gearFromSlots,
  gearSlots,
  plannerGearFor,
  ranksFromTalentsString,
  specForSplit,
  specOf,
  talentLevel,
  talentsString,
  toBuildDraft,
  toCharacterSpec,
  type SimCharacter,
} from './character';
import type { CharacterSource } from './types';

const WEB_ROOT = path.resolve(import.meta.dirname, '../../..');
const BUILD = '1.15.9.69722';

// Reads the checked-in fixture directly (as src/fixtures/planner/fixture.test.ts does)
// rather than public/data/<build>/, which only exists after a sync step and is published
// under src/data/active-build.json's build id, not this fixture's own.
async function json<T>(relative: string): Promise<T> {
  return JSON.parse(await readFile(path.join(WEB_ROOT, 'src/fixtures/planner', relative), 'utf8')) as T;
}
const warriorTalents = (): Promise<TalentFile> => json<TalentFile>('talents/warrior.json');
const classes = (): Promise<ClassRow[]> => json<ClassRow[]>('classes.json');
const races = (): Promise<RaceRow[]> => json<RaceRow[]>('races.json');
const combos = (): Promise<Combo[]> => json<Combo[]>('combos.json');

const source: CharacterSource = { kind: 'addon', ref: '', captured_at: '2026-09-14T10:00:00Z' };
const FURY = 'FS1:1.15.9.69722:warrior:orc:0/5530515/0:head=12640,main_hand=11726';

describe('specForSplit', () => {
  it('names the tree with the most points', () => {
    expect(specForSplit('warrior', [0, 31, 20])).toBe('warrior-fury');
    expect(specForSplit('warrior', [31, 0, 20])).toBe('warrior-arms');
    expect(specForSplit('mage', [0, 0, 31])).toBe('mage-frost');
  });

  it('resolves a tie to the first tree, as the phase 2 spec does', () => {
    expect(specForSplit('warrior', [20, 20, 11])).toBe('warrior-arms');
  });

  it('finds the third tree too, now that the spec list carries tanks and healers', () => {
    expect(specForSplit('warrior', [0, 0, 51])).toBe('warrior-protection');
    expect(specForSplit('druid', [0, 0, 51])).toBe('druid-restoration');
  });

  it('degrades to class-and-index when the spec list has no row for that tree', () => {
    // The fixture warrior file has two trees, but a class the list does not know has none.
    expect(specForSplit('demonhunter', [31, 0, 0])).toBe('demonhunter-0');
  });
});

describe('talentLevel', () => {
  it('is 9 with no points, 10 at the first, and capped at 60', () => {
    expect(talentLevel([])).toBe(9);
    expect(talentLevel([1001])).toBe(10);
    expect(talentLevel(Array.from({ length: 51 }, () => 1001))).toBe(60);
  });

  it('is what the strip shows and never what the engine is sent', () => {
    // api.SimRequest.Validate refuses any level but api.SimLevel.
    expect(SIM_LEVEL).toBe(60);
    expect(talentLevel([1001])).not.toBe(SIM_LEVEL);
  });
});

describe('characterFromFs1', () => {
  it('reads the export into a character with a legal point order', async () => {
    const [file, classRows, raceRows] = await Promise.all([warriorTalents(), classes(), races()]);
    const character = characterFromFs1(FURY, file, classRows, raceRows, source);
    expect(character.ok).toBe(true);
    if (!character.ok) return;
    expect(character.character.class_slug).toBe('warrior');
    expect(character.character.race_slug).toBe('orc');
    expect(character.character.spec).toBe('warrior-fury');
    expect(specOf(indexTalents(file), character.character.point_order)).toBe('warrior-fury');
    expect(character.character.gear).toEqual({ head: 12640, main_hand: 11726 });
    expect(character.character.point_order).toHaveLength(24);
    expect(validateOrder(indexTalents(file), character.character.point_order)).toEqual([]);
  });

  it('spends the prerequisite before the talent that gates on it', async () => {
    const [file, classRows, raceRows] = await Promise.all([warriorTalents(), classes(), races()]);
    const character = characterFromFs1(FURY, file, classRows, raceRows, source);
    if (!character.ok) throw new Error('expected a character');
    const order = character.character.point_order;
    const deathWish = order.indexOf(2006);
    expect(order.slice(0, deathWish).filter((id) => id === 2003)).toHaveLength(3);
  });

  it("passes the planner decoder's own reason through, which names the prefix", async () => {
    const [file, classRows, raceRows] = await Promise.all([warriorTalents(), classes(), races()]);
    const result = characterFromFs1('FS2:1:warrior:orc:0/0/0:', file, classRows, raceRows, source);
    expect(result.ok).toBe(false);
    if (result.ok) return;
    expect(result.message).toBe('That code is FS2; this site reads FS1.');
  });

  it('refuses an export for a class whose talents these are not', async () => {
    const [file, classRows, raceRows] = await Promise.all([warriorTalents(), classes(), races()]);
    const result = characterFromFs1('FS1:1.15.9.69722:mage:gnome:0/0/5:', file, classRows, raceRows, source);
    expect(result.ok).toBe(false);
    if (result.ok) return;
    expect(result.message).toBe(simCopy.classMismatch('mage', 'warrior'));
  });

  it('names the talents it had to drop rather than silently truncating', async () => {
    const [file, classRows, raceRows] = await Promise.all([warriorTalents(), classes(), races()]);
    const result = characterFromFs1(
      'FS1:1.15.9.69722:warrior:orc:0/0000010/0:',
      file,
      classRows,
      raceRows,
      source,
    );
    expect(result.ok).toBe(false);
    if (result.ok) return;
    expect(result.message).toBe(simCopy.unreachableTalents('Death Wish'));
  });

  it('refuses a race this build does not have rather than substituting the first one', async () => {
    const [file, classRows, raceRows] = await Promise.all([warriorTalents(), classes(), races()]);
    const result = characterFromFs1(
      'FS1:1.15.9.69722:warrior:draenei:0/5530515/0:',
      file,
      classRows,
      raceRows,
      source,
    );
    expect(result.ok).toBe(false);
    if (result.ok) return;
    expect(result.message).toBe(simCopy.unknownRace('draenei'));
  });
});

describe('talentsString', () => {
  it('is one decimal digit per talent in tab order, trees joined by a dash', async () => {
    const file = await warriorTalents();
    // Five Booming Voice, five Cruelty, three Unbridled Wrath.
    const order = [2001, 2001, 2001, 2001, 2001, 2002, 2002, 2002, 2002, 2002, 2003, 2003, 2003];
    expect(talentsString(indexTalents(file), order)).toBe('-553');
  });

  it('trims trailing zeros inside a tree but keeps every separator', async () => {
    const file = await warriorTalents();
    expect(talentsString(indexTalents(file), [1001, 2007])).toBe('1-0000001');
  });

  it('is all separators for a build with no points', async () => {
    const file = await warriorTalents();
    expect(talentsString(indexTalents(file), [])).toBe('-');
  });
});

describe('ranksFromTalentsString, the inverse of talentsString', () => {
  it('reads one decimal digit per talent in tab order, per tree', () => {
    expect(ranksFromTalentsString('-5530515-')).toEqual([[], [5, 5, 3, 0, 5, 1, 5], []]);
  });

  it('round-trips a real talentsString exactly, one digit per talent', async () => {
    const file = await warriorTalents();
    // Five Booming Voice, five Cruelty, three Unbridled Wrath -- the same order the
    // talentsString describe block above encodes to '-553'.
    const order = [2001, 2001, 2001, 2001, 2001, 2002, 2002, 2002, 2002, 2002, 2003, 2003, 2003];
    const encoded = talentsString(indexTalents(file), order);
    expect(encoded).toBe('-553');
    // Decoding what H4's saved-sim rerun actually receives (a CharacterSpec.talents
    // string) reaches the same per-talent ranks a fresh encode started from.
    expect(ranksFromTalentsString(encoded)).toEqual([[], [5, 5, 3]]);
  });

  it('is all empty trees for a build with no points', () => {
    expect(ranksFromTalentsString('-')).toEqual([[], []]);
  });
});

describe('gearSlots', () => {
  it('is one entry per equipped slot, in the planner’s own slot order', () => {
    expect(gearSlots({ main_hand: 11726, head: 12640 })).toEqual([
      { slot: 'head', item_id: 12640 },
      { slot: 'main_hand', item_id: 11726 },
    ]);
  });

  it('leaves an empty slot out rather than sending a zero item id', () => {
    expect(gearSlots({})).toEqual([]);
  });
});

describe('gearFromSlots, the inverse of gearSlots', () => {
  it('is the planner’s Gear map, keyed by slot', () => {
    expect(
      gearFromSlots([
        { slot: 'head', item_id: 12640 },
        { slot: 'main_hand', item_id: 11726 },
      ]),
    ).toEqual({
      head: 12640,
      main_hand: 11726,
    });
  });

  it('round-trips gearSlots exactly', () => {
    const gear = { main_hand: 11726, head: 12640 };
    expect(gearFromSlots(gearSlots(gear))).toEqual(gear);
  });

  it('is empty for no gear', () => {
    expect(gearFromSlots([])).toEqual({});
  });
});

describe('plannerGearFor', () => {
  // Minimal filler; each test below overrides only the fields it cares about.
  const base: SimCharacter = {
    name: 'Thrallgar',
    spec: 'warrior-fury',
    class_slug: 'warrior',
    race_slug: 'orc',
    talent_level: 22,
    tree_version: BUILD,
    point_order: [],
    gear: {},
    gear_slots: [],
    buffs: [],
    consumables: [],
    source,
    professions: [],
    bags: [],
    bank: [],
    sets: [],
    loadouts: [],
  };

  it('prefers gear_slots over the id map, mirroring toCharacterSpec’s own precedence (character.ts:281-309)', () => {
    const character: SimCharacter = {
      ...base,
      gear: { head: 1 },
      gear_slots: [
        { slot: 'head', item_id: 12640 },
        { slot: 'main_hand', item_id: 11726 },
      ],
    };
    expect(plannerGearFor(character)).toEqual({ head: 12640, main_hand: 11726 });
  });

  it('falls back to the id map when gear_slots is empty', () => {
    const character: SimCharacter = { ...base, gear: { head: 12640 }, gear_slots: [] };
    expect(plannerGearFor(character)).toEqual({ head: 12640 });
  });

  it('is an empty object when both are empty', () => {
    const character: SimCharacter = { ...base, gear: {}, gear_slots: [] };
    expect(plannerGearFor(character)).toEqual({});
  });

  it('returns a fresh copy from the id map: mutating it does not touch the character', () => {
    const character: SimCharacter = { ...base, gear: { head: 12640 }, gear_slots: [] };
    const gear = plannerGearFor(character);
    gear.head = 99999;
    expect(character.gear.head).toBe(12640);
  });

  it('returns a fresh copy from gear_slots: mutating it does not touch the character', () => {
    const character: SimCharacter = {
      ...base,
      gear: {},
      gear_slots: [{ slot: 'head', item_id: 12640 }],
    };
    const gear = plannerGearFor(character);
    gear.head = 99999;
    expect(character.gear_slots).toEqual([{ slot: 'head', item_id: 12640 }]);
  });
});

describe('toCharacterSpec', () => {
  it('is everything the engine needs, as plain JSON and nothing else', async () => {
    const file = await warriorTalents();
    const order = [2001, 2001, 2001, 2001, 2001, 2002, 2002, 2002, 2002, 2002, 2003, 2003, 2003];
    const character: SimCharacter = {
      name: 'Thrallgar',
      spec: 'warrior-fury',
      class_slug: 'warrior',
      race_slug: 'orc',
      talent_level: 22,
      tree_version: BUILD,
      point_order: order,
      gear: { head: 12640 },
      gear_slots: [],
      professions: [],
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      buffs: [],
      consumables: [],
      source,
    };
    expect(
      toCharacterSpec(character, indexTalents(file), ['battle_shout'], ['elixir_of_the_mongoose']),
    ).toEqual({
      name: 'Thrallgar',
      race: 'orc',
      class: 'warrior',
      // Always 60, whatever the talent spend says: sim/request refuses any other level.
      level: 60,
      talents: '-553',
      gear: [{ slot: 'head', item_id: 12640 }],
      buffs: ['battle_shout'],
      consumes: ['elixir_of_the_mongoose'],
    });
  });

  it('sends ids in the engine’s own snake case, because a kebab id is an error there', () => {
    // sim/request/buffs.go reads ids off the protobuf descriptors: RaidBuffs.battle_shout
    // is "battle_shout", AgilityElixir.ElixirOfTheMongoose is "elixir_of_the_mongoose", and
    // an off-hand imbue is slot-qualified. An unknown id fails the whole run.
    for (const id of [...PRESET_BUFFS['raid-buffed'], ...PRESET_CONSUMABLES['raid-buffed']]) {
      expect(id).toMatch(/^[a-z0-9_]+(:[a-z0-9_]+)?$/);
    }
  });

  it('survives JSON, which is the only form the wasm ever sees it in', async () => {
    const file = await warriorTalents();
    const character: SimCharacter = {
      name: 'Thrallgar',
      spec: 'warrior-fury',
      class_slug: 'warrior',
      race_slug: 'orc',
      talent_level: 11,
      tree_version: BUILD,
      point_order: [2001, 2002],
      gear: { head: 12640 },
      gear_slots: [],
      professions: [],
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      buffs: [],
      consumables: [],
      source,
    };
    const spec = toCharacterSpec(character, indexTalents(file), [], []);
    expect(JSON.parse(JSON.stringify(spec))).toEqual(spec);
  });
});

describe('toCharacterSpec cooldowns', () => {
  it('omits the key entirely when there are none, rather than sending an empty list', async () => {
    const file = await warriorTalents();
    const character: SimCharacter = {
      name: 'Thrallgar',
      spec: 'warrior-fury',
      class_slug: 'warrior',
      race_slug: 'orc',
      talent_level: 11,
      tree_version: BUILD,
      point_order: [2001, 2002],
      gear: { head: 12640 },
      gear_slots: [],
      professions: [],
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      buffs: [],
      consumables: [],
      source,
    };
    const spec = toCharacterSpec(character, indexTalents(file), [], [], []);
    expect(spec.cooldowns).toBeUndefined();
  });

  it('carries the rows through, copied rather than shared', async () => {
    const file = await warriorTalents();
    const character: SimCharacter = {
      name: 'Thrallgar',
      spec: 'warrior-fury',
      class_slug: 'warrior',
      race_slug: 'orc',
      talent_level: 11,
      tree_version: BUILD,
      point_order: [2001, 2002],
      gear: { head: 12640 },
      gear_slots: [],
      professions: [],
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      buffs: [],
      consumables: [],
      source,
    };
    const rows = [{ id: 'major_mana_potion', at_sec: [0] }];
    const spec = toCharacterSpec(character, indexTalents(file), [], [], rows);
    expect(spec.cooldowns).toEqual(rows);
    expect(spec.cooldowns).not.toBe(rows);
  });
});

describe('the planner conversion, both ways', () => {
  it('round-trips a character through BuildDraft without losing a point or a slot', async () => {
    const [file, classRows, raceRows] = await Promise.all([warriorTalents(), classes(), races()]);
    const loaded = characterFromFs1(FURY, file, classRows, raceRows, source, 'Thrallgar');
    if (!loaded.ok) throw new Error(loaded.message);
    const order = loaded.character.point_order;
    const character: SimCharacter = {
      name: 'Thrallgar',
      spec: 'warrior-fury',
      class_slug: 'warrior',
      race_slug: 'orc',
      talent_level: 33,
      tree_version: BUILD,
      point_order: order,
      gear: { head: 12640, main_hand: 11726 },
      gear_slots: gearSlots({ head: 12640, main_hand: 11726 }),
      professions: [],
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      buffs: ['battle_shout'],
      consumables: ['elixir_of_the_mongoose'],
      source,
    };

    const draft = toBuildDraft(character, classRows, raceRows);
    expect(draft).not.toBeNull();
    expect(draft).toEqual({
      class_id: 1,
      race_id: 2,
      tree_version: BUILD,
      point_order: order,
      gear: { head: 12640, main_hand: 11726 },
    });

    const back = fromBuildDraft(draft!, file, classRows, raceRows, {
      name: 'Thrallgar',
      buffs: ['battle_shout'],
      consumables: ['elixir_of_the_mongoose'],
      source,
    });
    expect(back).toEqual({ ...character, talent_level: 33 });
  });

  it('refuses a draft whose talent file is for another class', async () => {
    const [file, classRows, raceRows] = await Promise.all([warriorTalents(), classes(), races()]);
    const draft = { class_id: 3, race_id: 2, tree_version: BUILD, point_order: [], gear: {} };
    expect(fromBuildDraft(draft, file, classRows, raceRows, { source })).toBeNull();
  });

  it('refuses a character whose race the reference data does not know', async () => {
    const [classRows, raceRows] = await Promise.all([classes(), races()]);
    const character: SimCharacter = {
      name: 'Nobody',
      spec: 'warrior-fury',
      class_slug: 'warrior',
      race_slug: 'skyborne-unknown',
      talent_level: 60,
      tree_version: BUILD,
      point_order: [],
      gear: {},
      gear_slots: [],
      professions: [],
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      buffs: [],
      consumables: [],
      source,
    };
    expect(toBuildDraft(character, classRows, raceRows)).toBeNull();
  });
});

// Task 17. The plan's test block names fictional helpers (fixtureCharacter, fixtureIndex,
// FURY_CODE, fixtureTalents, fixtureClasses, fixtureRaces, SOURCE); rewritten here against
// this file's own real, async helpers (task-rulings.md, Task 17), keeping the plan's
// assertions.
describe('characterFromFs1 with version 2 sections', () => {
  it('carries bags, bank, sets, loadouts and professions onto the character', async () => {
    const [file, classRows, raceRows] = await Promise.all([warriorTalents(), classes(), races()]);
    const code = `${FURY}|bags=16963|bank=19360:2505:1820|sets=AQ=head=21329|loadouts=Arms=5530515/0/0|professions=engineering,alchemy`;
    const result = characterFromFs1(code, file, classRows, raceRows, source);
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    expect(result.character.bags).toEqual([{ itemId: 16963 }]);
    expect(result.character.bank).toEqual([{ itemId: 19360, enchant: 2505, suffix: 1820 }]);
    expect(result.character.sets).toEqual([{ name: 'AQ', gear: [{ slot: 'head', itemId: 21329 }] }]);
    expect(result.character.loadouts[0].name).toBe('Arms');
    expect(result.character.professions).toEqual(['engineering', 'alchemy']);
  });

  it('carries a gear entry’s enchant and suffix (contract 10.5) onto gear_slots', async () => {
    const [file, classRows, raceRows] = await Promise.all([warriorTalents(), classes(), races()]);
    const result = characterFromFs1(
      'FS1:1.15.9.69722:warrior:orc:0/5530515/0:head=12640:2504:1820',
      file,
      classRows,
      raceRows,
      source,
    );
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    expect(result.character.gear_slots).toEqual([
      { slot: 'head', item_id: 12640, enchant: 2504, suffix: 1820 },
    ]);
    // The planner-facing map keeps the id alone, as every reader of it expects.
    expect(result.character.gear).toEqual({ head: 12640 });
  });

  it('gives a version 1 export the empty lists rather than leaving them undefined', async () => {
    const [file, classRows, raceRows] = await Promise.all([warriorTalents(), classes(), races()]);
    const result = characterFromFs1(FURY, file, classRows, raceRows, source);
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    expect(result.character.bags).toEqual([]);
    expect(result.character.sets).toEqual([]);
    expect(result.character.professions).toEqual([]);
    expect(result.character.gear_slots.length).toBeGreaterThan(0);
  });
});

describe('toCharacterSpec with per-slot enchants and professions', () => {
  it('sends gear_slots verbatim rather than rebuilding the list from the id map', async () => {
    const file = await warriorTalents();
    const character: SimCharacter = {
      name: 'Thrallgar',
      spec: 'warrior-fury',
      class_slug: 'warrior',
      race_slug: 'orc',
      talent_level: 22,
      tree_version: BUILD,
      point_order: [],
      gear: { head: 12640 },
      gear_slots: [{ slot: 'head', item_id: 12640, enchant: 2504, suffix: 1820 }],
      professions: ['engineering'],
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      buffs: [],
      consumables: [],
      source,
    };
    const spec = toCharacterSpec(character, indexTalents(file), [], []);
    expect(spec.gear).toEqual([{ slot: 'head', item_id: 12640, enchant: 2504, suffix: 1820 }]);
    expect(spec.professions).toEqual(['engineering']);
  });

  it('falls back to the id map for a character that has no slot list', async () => {
    const file = await warriorTalents();
    const character: SimCharacter = {
      name: 'Thrallgar',
      spec: 'warrior-fury',
      class_slug: 'warrior',
      race_slug: 'orc',
      talent_level: 22,
      tree_version: BUILD,
      point_order: [],
      gear: { head: 12640 },
      gear_slots: [],
      professions: [],
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      buffs: [],
      consumables: [],
      source,
    };
    const spec = toCharacterSpec(character, indexTalents(file), [], []);
    expect(spec.gear).toEqual([{ slot: 'head', item_id: 12640 }]);
    expect(spec.professions).toBeUndefined();
  });
});

describe('codeForCharacterSpec', () => {
  // Finding 1, final whole-branch review: SimView.svelte's "Run this yourself" and
  // store-request.ts's Apply used to build this FS1 v2 code by hand, in two places that
  // drifted once already. This is the one shared conversion now; the round trip through
  // characterFromFs1 is what proves it is lossless the same way the two call sites relied
  // on their own hand-written versions being.
  it('round-trips a CharacterSpec through an FS1 v2 code, enchants, suffixes and professions included', async () => {
    const file = await warriorTalents();
    const character: SimCharacter = {
      name: 'Thrallgar',
      spec: 'warrior-fury',
      class_slug: 'warrior',
      race_slug: 'orc',
      talent_level: 22,
      tree_version: BUILD,
      point_order: [],
      gear: { head: 12640 },
      gear_slots: [{ slot: 'head', item_id: 12640, enchant: 2504, suffix: 1820 }],
      professions: ['engineering', 'alchemy'],
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      buffs: [],
      consumables: [],
      source,
    };
    const spec = toCharacterSpec(character, indexTalents(file), [], []);

    const code = codeForCharacterSpec(spec, BUILD);

    const [classRows, raceRows] = await Promise.all([classes(), races()]);
    const decoded = characterFromFs1(code, file, classRows, raceRows, source);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.character.class_slug).toBe('warrior');
    expect(decoded.character.race_slug).toBe('orc');
    expect(decoded.character.gear_slots).toEqual([
      { slot: 'head', item_id: 12640, enchant: 2504, suffix: 1820 },
    ]);
    expect(decoded.character.professions).toEqual(['engineering', 'alchemy']);
    expect(talentsString(indexTalents(file), decoded.character.point_order)).toBe(spec.talents);
  });

  it('omits professions from the code when the spec has none', async () => {
    const file = await warriorTalents();
    const character: SimCharacter = {
      name: 'Thrallgar',
      spec: 'warrior-fury',
      class_slug: 'warrior',
      race_slug: 'orc',
      talent_level: 22,
      tree_version: BUILD,
      point_order: [],
      gear: { head: 12640 },
      gear_slots: [],
      professions: [],
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      buffs: [],
      consumables: [],
      source,
    };
    const spec = toCharacterSpec(character, indexTalents(file), [], []);

    const code = codeForCharacterSpec(spec, BUILD);

    expect(code.includes('professions=')).toBe(false);
  });
});

describe('characterFromPlanner', () => {
  it('is null before the talent file has loaded', () => {
    const store = createPlannerStore({ treeVersion: BUILD, classSlug: 'warrior', raceSlug: 'human' });
    expect(characterFromPlanner(store)).toBeNull();
  });

  // Task 20 review, HIGH: store.toDraft() throws when classRow or raceRow cannot resolve
  // (store.svelte.ts:319-321), and classSlug is untrusted input -- an unvalidated ?class=
  // query string, or a decoded FS1 code naming a class the reference data does not have.
  // Nothing repairs classSlug the way repairRaceForClass repairs raceSlug, so classRow stays
  // null forever once the reference data has loaded. Before the fix, this called
  // store.toDraft() anyway and threw, uncaught, inside Planner.svelte's automatic $effect.
  it('is null, not a throw, when the class slug names no loaded class', async () => {
    const [file, classRows, raceRows, comboRows] = await Promise.all([
      warriorTalents(),
      classes(),
      races(),
      combos(),
    ]);
    const store = createPlannerStore({ treeVersion: BUILD, classSlug: 'doesnotexist', raceSlug: 'human' });
    store.setReference({ classes: classRows, races: raceRows, combos: comboRows });
    store.setTalents(file);

    expect(store.classRow).toBeNull();
    expect(() => characterFromPlanner(store)).not.toThrow();
    expect(characterFromPlanner(store)).toBeNull();
  });

  // The same failure mode, for a race the reference data does not have -- reachable the same
  // way (an unvalidated ?race=, or a decoded FS1 code's race slug). setReference's own
  // repairRaceForClass would otherwise silently move an unknown raceSlug to the first legal
  // race for the class, so this passes no combos: with nothing legal to repair to, raceSlug
  // stays unresolved and raceRow stays null, the way it would for a class/race pair this
  // build's own combos.json genuinely has no entry for.
  it('is null, not a throw, when the race slug names no loaded race', async () => {
    const [file, classRows, raceRows] = await Promise.all([warriorTalents(), classes(), races()]);
    const store = createPlannerStore({
      treeVersion: BUILD,
      classSlug: 'warrior',
      raceSlug: 'doesnotexist',
    });
    store.setReference({ classes: classRows, races: raceRows, combos: [] });
    store.setTalents(file);

    expect(store.classRow).not.toBeNull();
    expect(store.raceRow).toBeNull();
    expect(() => characterFromPlanner(store)).not.toThrow();
    expect(characterFromPlanner(store)).toBeNull();
  });

  it('carries the planner build through once the class and race both resolve', async () => {
    const [file, classRows, raceRows, comboRows] = await Promise.all([
      warriorTalents(),
      classes(),
      races(),
      combos(),
    ]);
    const store = createPlannerStore({ treeVersion: BUILD, classSlug: 'warrior', raceSlug: 'orc' });
    store.setReference({ classes: classRows, races: raceRows, combos: comboRows });
    store.setTalents(file);
    store.applyOrder([1001], {});

    const character = characterFromPlanner(store);
    expect(character).not.toBeNull();
    expect(character?.class_slug).toBe('warrior');
    expect(character?.race_slug).toBe('orc');
    expect(character?.point_order).toEqual([1001]);
  });
});
