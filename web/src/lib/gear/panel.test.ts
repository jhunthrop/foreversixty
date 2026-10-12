// web/src/lib/gear/panel.test.ts
// Task L8.6: the profile's Gear section's row model and title (ruling 1), its words apart from
// the score card's (ruling 3) and its footer link (ruling 2), over the fixture that covers every
// row state (src/fixtures/character-gear.ts).
import { describe, expect, it } from 'vitest';
import { GEAR_FIXTURE, GEAR_PANEL_FIXTURE } from '../../fixtures/character-gear';
import { scoreCopy } from '../score/copy';
import { gearCopy } from './copy';
import { gearGapCount, gearRows, gearTitle, type GearRow } from './panel';
import type { CharacterGearSlot } from './types';

function row(slot: CharacterGearSlot['slot']): GearRow {
  const found = gearRows(GEAR_FIXTURE).find((candidate) => candidate.slot === slot);
  if (found === undefined) throw new Error(`no ${slot} row`);
  return found;
}

describe('gearRows (L8.6 ruling 1)', () => {
  it('keeps all 17 rows in the export’s order, each labelled as the client labels its slot', () => {
    const rows = gearRows(GEAR_FIXTURE);
    expect(rows.map((r) => r.slot)).toEqual([
      'head',
      'neck',
      'shoulder',
      'back',
      'chest',
      'wrist',
      'hands',
      'waist',
      'legs',
      'feet',
      'finger1',
      'finger2',
      'trinket1',
      'trinket2',
      'main_hand',
      'off_hand',
      'ranged',
    ]);
    expect(rows.map((r) => r.label)).toEqual([
      'Head',
      'Neck',
      'Shoulder',
      'Back',
      'Chest',
      'Wrist',
      'Hands',
      'Waist',
      'Legs',
      'Feet',
      'Ring',
      'Ring',
      'Trinket',
      'Trinket',
      'Main hand',
      'Off hand',
      'Ranged',
    ]);
  });

  it('an upgrade: worn → pick · source · chance · +N pts', () => {
    expect(row('head')).toEqual({
      slot: 'head',
      label: 'Head',
      state: 'upgrade',
      worn: { id: 10504, name: 'Green Lens', known: true },
      pick: { id: 16921, name: 'Halo of Transcendence', known: true },
      source: 'Onyxia’s Lair: Onyxia',
      chance: '17%',
      gap: '+18.4 pts',
      gapLabel: '18.4 points to gain',
    });
  });

  it('a chance under 1% reads "<1%"; a source with no rate shows no chance', () => {
    expect(row('back').chance).toBe('<1%');
    expect(row('chest')).toMatchObject({
      state: 'upgrade',
      source: 'Tailoring',
      chance: null,
      gap: '+11.3 pts',
    });
  });

  it('a worn item that is the pick: "Best for your band", no arrow (no pick shown), no gap', () => {
    expect(row('shoulder')).toEqual({
      slot: 'shoulder',
      label: 'Shoulder',
      state: 'best',
      worn: { id: 16924, name: 'Pauldrons of Transcendence', known: true },
      pick: null,
      source: null,
      chance: null,
      gap: null,
      gapLabel: null,
    });
    for (const slot of ['hands', 'feet', 'trinket1'] as const) expect(row(slot).state).toBe('best');
  });

  it('rings worn the other way round (gap 0, ids apart) read as best too: the API matched them', () => {
    expect(row('finger1')).toMatchObject({ state: 'best', worn: { name: 'Blackstone Ring' }, pick: null });
    expect(row('finger2')).toMatchObject({ state: 'best', worn: { name: 'Cauterizing Band' }, pick: null });
  });

  it('an empty slot: "Empty" with the pick, its source, chance and gap', () => {
    expect(row('neck')).toEqual({
      slot: 'neck',
      label: 'Neck',
      state: 'empty',
      worn: null,
      pick: { id: 18814, name: 'Choker of the Fire Lord', known: true },
      source: 'Molten Core: Ragnaros',
      chance: '12%',
      gap: '+12.6 pts',
      gapLabel: '12.6 points to gain',
    });
    // Empty under a worn two-hander: the main hand carries the pair's gap, this one reads 0.
    expect(row('off_hand')).toMatchObject({ state: 'empty', worn: null, gap: '+0 pts' });
  });

  it('a slot with no pick reads only the worn item; with nothing worn either, only "Empty"', () => {
    expect(row('ranged')).toEqual({
      slot: 'ranged',
      label: 'Ranged',
      state: 'worn-only',
      worn: { id: 13396, name: 'Skul’s Ghastly Touch', known: true },
      pick: null,
      source: null,
      chance: null,
      gap: null,
      gapLabel: null,
    });
    expect(row('trinket2')).toMatchObject({ state: 'none', worn: null, pick: null, gap: null });
  });

  it('a worn item the build does not have is "Unknown item", and a gap it cannot score is a dash', () => {
    expect(row('legs')).toEqual({
      slot: 'legs',
      label: 'Legs',
      state: 'upgrade',
      worn: { id: 99001, name: 'Unknown item', known: false },
      pick: { id: 16814, name: 'Pants of Prophecy', known: true },
      source: 'Molten Core: Magmadar',
      chance: '18%',
      gap: '–',
      gapLabel: 'Not scored',
    });
  });

  it('a whole-number gap drops its ".0"', () => {
    const rows = gearRows([{ ...GEAR_FIXTURE[0], gap: 30 }]);
    expect(rows[0].gap).toBe('+30 pts');
  });
});

describe('gearTitle (L8.6 ruling 1)', () => {
  it('reads "Gear · +70 pts across 7 slots" from gear_gap_total and the count of non-zero gaps', () => {
    expect(gearGapCount(GEAR_FIXTURE)).toBe(7);
    expect(gearTitle(GEAR_PANEL_FIXTURE.gear, GEAR_PANEL_FIXTURE.gear_gap_total)).toBe(
      'Gear · +70 pts across 7 slots',
    );
    expect(gearTitle(GEAR_FIXTURE, 88)).toBe('Gear · +88 pts across 7 slots');
  });

  it('one slot is "slot"; a decimal total keeps its tenth', () => {
    const one = [GEAR_FIXTURE[0], GEAR_FIXTURE[2]];
    expect(gearTitle(one, 18.4)).toBe('Gear · +18.4 pts across 1 slot');
  });

  it('a total of 0 says there is nothing to gain over the picks; a null total is the word alone', () => {
    const allPicks = GEAR_FIXTURE.map((slot) => ({ ...slot, gap: slot.gap === null ? null : 0 }));
    expect(gearTitle(allPicks, 0)).toBe('Gear · nothing to gain over the picks');
    expect(
      gearTitle(
        GEAR_FIXTURE.map((slot) => ({ ...slot, gap: null })),
        null,
      ),
    ).toBe('Gear');
  });
});

describe('gearCopy (L8.6 rulings 1 to 3)', () => {
  it('the no-gear line', () => {
    expect(gearCopy.noGear).toBe('No gear on file. Upload a log or sync the addon.');
    const { lead, upload, or, sync, end } = gearCopy.noGearParts;
    expect(`${lead}${upload}${or}${sync}${end}`).toBe(gearCopy.noGear);
  });

  it('says in one sentence how the slot gaps differ from the card’s "below its top"', () => {
    expect(gearCopy.difference).toBe(
      'Below its top measures against the band’s best set; the slot gaps measure against the picks you can wear.',
    );
    // The card keeps its own sentence (ruling 3).
    expect(`${scoreCopy.mostToGainLead} Gear${scoreCopy.mostToGainRest(88)}`).toBe(
      'Most to gain: Gear, 88 below its top.',
    );
    // The section never says "below its top" of its own total; the card never says "across".
    expect(gearTitle(GEAR_FIXTURE, 70)).not.toContain('below');
    expect(scoreCopy.mostToGainRest(88)).not.toContain('across');
  });

  it('the footer link: "Every pick for Holy at 60", a band below 60 by its range', () => {
    expect(gearCopy.bisLink('Holy', 60)).toBe('Every pick for Holy at 60');
    expect(gearCopy.bisLink('Arms', 30)).toBe('Every pick for Arms at 30 to 39');
    expect(gearCopy.bisLink('Holy', null)).toBe('Every pick for Holy');
    expect(gearCopy.bisArrow).toBe('→');
  });
});
