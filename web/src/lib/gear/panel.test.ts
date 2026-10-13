// web/src/lib/gear/panel.test.ts
// Task L8.6 and its fix round 1: the profile's Gear section's row model and title, the totals
// line under it and the footer link, over the fixture that covers every row state
// (src/fixtures/character-gear.ts).
import { describe, expect, it } from 'vitest';
import { GEAR_FIXTURE, GEAR_PANEL_FIXTURE } from '../../fixtures/character-gear';
import { gearCopy } from './copy';
import { gearGapCount, gearRows, gearTitle, gearTitleSummary, gearTotalsLine, type GearRow } from './panel';
import type { CharacterGearSlot } from './types';

/** The title as read: its no-break spaces (which keep each dot on the line before it) as spaces. */
function title(gear: readonly CharacterGearSlot[], total: number | null): string {
  return gearTitle(gear, total).replace(/\u00a0/g, ' ');
}

/** The worn staff (13000) is two-handed in the build's items file. */
const TWO_HAND = (id: number): boolean => id === 13000;

function row(slot: CharacterGearSlot['slot'], twoHand: (id: number) => boolean = TWO_HAND): GearRow {
  const found = gearRows(GEAR_FIXTURE, twoHand).find((candidate) => candidate.slot === slot);
  if (found === undefined) throw new Error(`no ${slot} row`);
  return found;
}

describe('gearRows (L8.6 ruling 1, fix round 1)', () => {
  it('keeps all 17 rows in the export’s order', () => {
    expect(gearRows(GEAR_FIXTURE).map((r) => r.slot)).toEqual(GEAR_FIXTURE.map((slot) => slot.slot));
  });

  it('labels the pairs "Ring 1", "Ring 2", "Trinket 1", "Trinket 2" in export order (fix 1 ruling 7)', () => {
    expect(gearRows(GEAR_FIXTURE).map((r) => r.label)).toEqual([
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
      'Ring 1',
      'Ring 2',
      'Trinket 1',
      'Trinket 2',
      'Main hand',
      'Off hand',
      'Ranged',
    ]);
  });

  it('an upgrade: worn → pick, the source, "17% drop" and the gap', () => {
    expect(row('head')).toEqual({
      slot: 'head',
      label: 'Head',
      state: 'upgrade',
      worn: { id: 10504, name: 'Green Lens', known: true },
      pick: { id: 16921, name: 'Halo of Transcendence', known: true },
      source: "Onyxia's Lair: Onyxia",
      chance: '17% drop',
      note: null,
      inline: false,
      gap: '+18.4 pts',
      gapLabel: '18.4 points to gain',
      strong: true,
    });
  });

  it('a chance under 1% reads "<1% drop"; a source with no rate shows no chance (fix 1 ruling 8)', () => {
    expect(row('back').chance).toBe('<1% drop');
    expect(row('chest')).toMatchObject({
      state: 'upgrade',
      source: 'Tailoring',
      chance: null,
      gap: '+11.3 pts',
    });
  });

  it('the worn item that is the pick: "Best for your band" on its line, no pick, no gap', () => {
    expect(row('shoulder')).toEqual({
      slot: 'shoulder',
      label: 'Shoulder',
      state: 'best',
      worn: { id: 16924, name: 'Pauldrons of Transcendence', known: true },
      pick: null,
      source: null,
      chance: null,
      note: 'Best for your band',
      inline: true,
      gap: null,
      gapLabel: null,
      strong: false,
    });
    for (const slot of ['hands', 'feet', 'trinket1', 'finger1', 'finger2'] as const)
      expect(row(slot).state).toBe('best');
  });

  it('an empty slot: "Empty" with the pick, its source, chance and gap', () => {
    expect(row('neck')).toMatchObject({
      state: 'empty',
      worn: null,
      pick: { id: 18814, name: 'Choker of the Fire Lord', known: true },
      source: 'Molten Core: Ragnaros',
      chance: '12% drop',
      inline: false,
      gap: '+12.6 pts',
    });
  });

  it('no pick: "No pick for your band yet" beside the worn item, or beside "Empty" (fix 2 ruling 6)', () => {
    expect(row('ranged')).toMatchObject({
      state: 'worn-only',
      worn: { id: 13396, name: "Skul's Ghastly Touch", known: true },
      pick: null,
      note: 'No pick for your band yet',
      inline: true,
      gap: null,
    });
    expect(row('trinket2')).toMatchObject({
      state: 'none',
      worn: null,
      pick: null,
      note: 'No pick for your band yet',
      inline: true,
      gap: null,
    });
  });

  it('a gap the API could not score reads "Not scored"; an unknown worn item "Unknown item"', () => {
    expect(row('legs')).toMatchObject({
      state: 'upgrade',
      worn: { id: 99001, name: 'Unknown item', known: false },
      gap: 'Not scored',
      gapLabel: null,
      strong: false,
    });
  });

  it('the off hand under a worn two-hander: "comes with the main-hand swap", no gap (fix 1 ruling 9)', () => {
    expect(row('off_hand')).toMatchObject({
      state: 'empty',
      source: 'Blackrock Spire: Gyth',
      chance: '9% drop',
      note: 'comes with the main-hand swap',
      inline: false,
      gap: null,
    });
    // The items file says the main hand is one-handed: the API's 0 stands.
    expect(row('off_hand', () => false)).toMatchObject({ note: null, gap: '+0 pts' });
  });

  it('a failed items file: the swap cannot be known, so the empty off hand shows no figure (fix 2 ruling 7)', () => {
    expect(row('off_hand', null as unknown as (id: number) => boolean)).toMatchObject({
      note: null,
      gap: null,
      gapLabel: null,
    });
    const failed = gearRows(GEAR_FIXTURE, null);
    expect(failed.find((r) => r.slot === 'off_hand')).toMatchObject({ note: null, gap: null });
    // Everything else as built.
    expect(failed.find((r) => r.slot === 'head')?.gap).toBe('+18.4 pts');
    expect(failed.find((r) => r.slot === 'legs')?.gap).toBe('Not scored');
  });

  it('the five largest non-zero gaps are strong, ties by export order (fix 1 ruling 6)', () => {
    const strong = gearRows(GEAR_FIXTURE, TWO_HAND)
      .filter((r) => r.strong)
      .map((r) => r.slot);
    // 18.4, 12.6, 11.3, 9.1, 7.6 -- not 6.8 (main hand) or 4.2 (wrist).
    expect(strong).toEqual(['head', 'neck', 'back', 'chest', 'waist']);
    const tied = GEAR_FIXTURE.map((slot) => (slot.gap !== null && slot.gap > 0 ? { ...slot, gap: 5 } : slot));
    expect(
      gearRows(tied)
        .filter((r) => r.strong)
        .map((r) => r.slot),
    ).toEqual(['head', 'neck', 'back', 'chest', 'wrist']);
  });

  it('a whole-number gap drops its ".0"', () => {
    expect(gearRows([{ ...GEAR_FIXTURE[0], gap: 30 }])[0].gap).toBe('+30 pts');
  });
});

describe('gearTitle (L8.6 fix round 1 ruling 3)', () => {
  it('"Gear · +70 pts across 7 slots · 1 not scored" from gear_gap_total, the gaps and the nulls', () => {
    expect(gearGapCount(GEAR_FIXTURE)).toBe(7);
    expect(title(GEAR_PANEL_FIXTURE.gear, GEAR_PANEL_FIXTURE.gear_gap_total)).toBe(
      'Gear · +70 pts across 7 slots · 1 not scored',
    );
  });

  it('with every gap scored there is no "not scored"; one slot is "slot"', () => {
    const scored = GEAR_FIXTURE.filter((slot) => slot.slot !== 'legs');
    expect(title(scored, 70)).toBe('Gear · +70 pts across 7 slots');
    expect(title([GEAR_FIXTURE[0], GEAR_FIXTURE[2]], 18.4)).toBe('Gear · +18.4 pts across 1 slot');
  });

  it('a total of 0 with an arrow left: "+0 pts from the picks you can wear" (fix 2 ruling 5)', () => {
    // The fixture with every scored gap 0: Neck and Off hand are still empty, their picks shown.
    const zero = GEAR_FIXTURE.map((slot) => ({ ...slot, gap: slot.gap === null ? null : 0 }));
    expect(title(zero, 0)).toBe('Gear · +0 pts from the picks you can wear · 1 not scored');
    expect(
      title(
        zero.filter((slot) => slot.slot !== 'legs'),
        0,
      ),
    ).toBe('Gear · +0 pts from the picks you can wear');
  });

  it('a total of 0 with no arrow anywhere: "every slot is the best you can wear" (fix 2 ruling 5)', () => {
    // Every pick worn: no row shows an arrow.
    const worn = GEAR_FIXTURE.filter((slot) => slot.pick === null || slot.worn?.id === slot.pick.id);
    expect(title(worn, 0)).toBe('Gear · every slot is the best you can wear');
    expect(title([GEAR_FIXTURE[13], GEAR_FIXTURE[16]], null)).toBe('Gear');
  });

  it('the summary after "Gear" alone, its own dots glued to the word before them (fix 2 ruling 4)', () => {
    expect(gearTitleSummary(GEAR_PANEL_FIXTURE.gear, 70)).toBe('+70 pts across 7 slots · 1 not scored');
    expect(gearTitleSummary([GEAR_FIXTURE[13]], null)).toBeNull();
  });

  it('nulls are counted with any total: "Gear · 2 not scored"', () => {
    const unscored = [GEAR_FIXTURE[8], { ...GEAR_FIXTURE[0], gap: null }];
    expect(title(unscored, null)).toBe('Gear · 2 not scored');
  });
});

describe('gearTotalsLine (L8.6 fix round 1 ruling 3)', () => {
  it('names both numbers: the card’s gear gap and gear_gap_total', () => {
    expect(gearTotalsLine(88, 70)).toBe(
      'The card’s 88 counts the band’s best set. These 70 count the picks you can wear.',
    );
    // A total of 0 has its own wording (fix 2 ruling 5).
    expect(gearTotalsLine(26, 0)).toBe(
      'The card’s 26 counts the band’s best set. The picks you can wear add 0.',
    );
    expect(gearTotalsLine(88, 18.4)).toContain('These 18.4 count');
  });

  it('is omitted when the card has no gear gap, or the API scored no slot', () => {
    expect(gearTotalsLine(null, 70)).toBeNull();
    expect(gearTotalsLine(0, 70)).toBeNull();
    expect(gearTotalsLine(88, null)).toBeNull();
  });
});

describe('gearCopy (L8.6 fix round 1)', () => {
  it('the no-gear line', () => {
    expect(gearCopy.noGear).toBe('No gear on file. Upload a log or sync the addon.');
    const { lead, upload, or, sync, end } = gearCopy.noGearParts;
    expect(`${lead}${upload}${or}${sync}${end}`).toBe(gearCopy.noGear);
  });

  it('the footer link: "Full Holy BiS list at 60", a band below 60 by its range (fix 1 ruling 3)', () => {
    expect(gearCopy.bisLink('Holy', 60)).toBe('Full Holy BiS list at 60');
    expect(gearCopy.bisLink('Arms', 30)).toBe('Full Arms BiS list at 30 to 39');
    expect(gearCopy.bisLink('Holy', null)).toBe('Full Holy BiS list');
    expect(gearCopy.bisArrow).toBe('→');
  });
});
