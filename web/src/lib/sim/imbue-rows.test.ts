import { describe, expect, it } from 'vitest';
import { imbueSections, withImbue, type ImbueSection } from './imbue-rows';

/** Every row id a section list offers, flattened. */
function ids(sections: ImbueSection[]): string[] {
  return sections.flatMap((section) => section.kinds.flatMap((kind) => kind.rows.map((row) => row.id)));
}

function slots(sections: ImbueSection[]): string[] {
  return sections.map((section) => section.slot);
}

// Round 10 wow-player review: the weapon section listed all 96 engine ids, so a mage saw
// poisons, shaman imbues, off-hand rows and "Main hand stone: Deadly Poison". The engine's
// stone field is only the second enchant beside a rogue's poison (sim/core stoneBesidePoison).
describe('imbueSections', () => {
  it('gives a caster one main hand of stones, weightstones and oils, and nothing else', () => {
    const mage = imbueSections('mage');
    expect(slots(mage)).toEqual(['main_hand_imbue']);
    expect(mage[0].kinds.map((kind) => kind.kind)).toEqual(['sharpening-stone', 'weightstone', 'oil']);
    expect(ids(mage)).toContain('main_hand_imbue:brilliant_wizard_oil');
    expect(ids(mage).some((id) => /poison|_weapon$|windfury/.test(id))).toBe(false);
  });

  it('never gives a shaman, druid, paladin, priest or warlock an off hand', () => {
    for (const cls of ['shaman', 'druid', 'paladin', 'priest', 'warlock', 'mage']) {
      expect(slots(imbueSections(cls)), cls).toEqual(['main_hand_imbue']);
    }
  });

  it('gives a shaman its own weapon imbues in the main hand, and only a shaman', () => {
    const shaman = ids(imbueSections('shaman'));
    for (const imbue of ['rockbiter_weapon', 'flametongue_weapon', 'frostbrand_weapon', 'windfury_weapon']) {
      expect(shaman).toContain(`main_hand_imbue:${imbue}`);
      expect(ids(imbueSections('warrior'))).not.toContain(`main_hand_imbue:${imbue}`);
    }
  });

  it('gives warriors and hunters both hands, without poisons or a stone-beside slot', () => {
    for (const cls of ['warrior', 'hunter']) {
      const sections = imbueSections(cls);
      expect(slots(sections), cls).toEqual(['main_hand_imbue', 'off_hand_imbue']);
      expect(ids(sections).some((id) => id.includes('poison'))).toBe(false);
    }
  });

  it('gives a rogue poisons in each hand and a stone or oil beside each poison', () => {
    const rogue = imbueSections('rogue');
    expect(slots(rogue)).toEqual(['main_hand_imbue', 'main_hand_stone', 'off_hand_imbue', 'off_hand_stone']);
    expect(ids(rogue)).toContain('off_hand_imbue:instant_poison');
    const beside = rogue.find((section) => section.slot === 'main_hand_stone');
    expect(beside?.kinds.map((kind) => kind.kind)).toEqual(['sharpening-stone', 'weightstone', 'oil']);
  });

  it('lists each item once per slot and never the totem imbue the raid buff already gives', () => {
    const all = ids(imbueSections('rogue'));
    expect(new Set(all).size).toBe(all.length);
    expect(all.some((id) => id.endsWith(':windfury'))).toBe(false);
  });
});

describe('withImbue', () => {
  it('holds one enchant per slot: ticking a second replaces the first', () => {
    const start = {
      buffs: [],
      consumables: ['main_hand_imbue:wizard_oil', 'off_hand_imbue:frost_oil', 'flask_of_supreme_power'],
    };
    const next = withImbue(start, 'main_hand_imbue:brilliant_wizard_oil', true);
    expect(next.consumables).toEqual([
      'off_hand_imbue:frost_oil',
      'flask_of_supreme_power',
      'main_hand_imbue:brilliant_wizard_oil',
    ]);
    expect(withImbue(next, 'main_hand_imbue:brilliant_wizard_oil', false).consumables).toEqual([
      'off_hand_imbue:frost_oil',
      'flask_of_supreme_power',
    ]);
    expect(start.consumables).toHaveLength(3);
  });
});
