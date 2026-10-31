// web/src/lib/gear/panel.test.ts
// Task L8.6 and its fix round 1: the profile's Gear section's row model and title, the totals
// line under it and the footer link, over the fixture that covers every row state
// (src/fixtures/character-gear.ts).
import { describe, expect, it } from 'vitest';
import { GEAR_FIXTURE, GEAR_PANEL_FIXTURE } from '../../fixtures/character-gear';
import { gearCopy } from './copy';
import { gearGapCount, gearRows, gearSpecLine, gearTitle, gearTitleSummary, type GearRow } from './panel';
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

  it('labels the pairs "Finger 1", "Finger 2", "Trinket 1", "Trinket 2" in export order (fix 1 ruling 7)', () => {
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
      'Finger 1',
      'Finger 2',
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
      gap: '+18.4 points',
      gapLabel: '18.4 points to gain',
      strong: true,
      noteUnknown: false,
    });
  });

  it('a chance under 1% reads "<1% drop"; a source with no rate shows no chance (fix 1 ruling 8)', () => {
    expect(row('back').chance).toBe('<1% drop');
    expect(row('chest')).toMatchObject({
      state: 'upgrade',
      source: 'Tailoring',
      chance: null,
      gap: '+11.3 points',
    });
  });

  it('the worn item that is the pick: "Already your best at this level" on its line, no pick, no gap', () => {
    expect(row('shoulder')).toEqual({
      slot: 'shoulder',
      label: 'Shoulder',
      state: 'best',
      worn: { id: 16924, name: 'Pauldrons of Transcendence', known: true },
      pick: null,
      source: null,
      chance: null,
      note: 'Already your best at this level',
      inline: true,
      gap: null,
      gapLabel: null,
      strong: false,
      noteUnknown: false,
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
      gap: '+12.6 points',
    });
  });

  it('no pick: "No pick for your level range yet" beside the worn item, or beside "Empty" (fix 2 ruling 6)', () => {
    expect(row('ranged')).toMatchObject({
      state: 'worn-only',
      worn: { id: 13396, name: "Skul's Ghastly Touch", known: true },
      pick: null,
      note: 'No pick for your level range yet',
      inline: true,
      gap: null,
    });
    expect(row('trinket2')).toMatchObject({
      state: 'none',
      worn: null,
      pick: null,
      note: 'No pick for your level range yet',
      inline: true,
      gap: null,
    });
  });

  it('a gap the API could not score reads "Not scored"; a worn item the build lacks "Not in our item data yet" (fix 4 item 5)', () => {
    expect(row('legs')).toMatchObject({
      state: 'upgrade',
      worn: { id: 99001, name: 'Not in our item data yet', known: false },
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
    expect(row('off_hand', () => false)).toMatchObject({ note: null, gap: '+0 points' });
  });

  it('an unknown main hand (the file loading or failed): the swap note is held, unknown, and no figure (fix 4 item 4)', () => {
    const failed = gearRows(GEAR_FIXTURE, null);
    expect(failed.find((r) => r.slot === 'off_hand')).toMatchObject({
      note: 'comes with the main-hand swap',
      noteUnknown: true,
      gap: null,
      gapLabel: null,
    });
    // With no worn main hand there is no swap to hold: the API's 0 stands.
    const noMain = GEAR_FIXTURE.map((slot) => (slot.slot === 'main_hand' ? { ...slot, worn: null } : slot));
    expect(gearRows(noMain, null).find((r) => r.slot === 'off_hand')).toMatchObject({
      note: null,
      noteUnknown: false,
      gap: '+0 points',
    });
    // Known, the note is plain.
    expect(row('off_hand')).toMatchObject({ noteUnknown: false });
    // Everything else as built.
    expect(failed.find((r) => r.slot === 'head')?.gap).toBe('+18.4 points');
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
    expect(gearRows([{ ...GEAR_FIXTURE[0], gap: 30 }])[0].gap).toBe('+30 points');
  });
});

describe('gearTitle (L8.6 fix round 1 ruling 3)', () => {
  it('"Gear · +70 points across 7 slots · 1 not scored" from gear_gap_total, the gaps and the nulls', () => {
    expect(gearGapCount(GEAR_FIXTURE)).toBe(7);
    expect(title(GEAR_PANEL_FIXTURE.gear, GEAR_PANEL_FIXTURE.gear_gap_total)).toBe(
      'Gear · 6 of 15 slots best in slot · +70 points across 7 slots · 1 not scored',
    );
  });

  it('with every gap scored there is no "not scored"; one slot is "slot"', () => {
    const scored = GEAR_FIXTURE.filter((slot) => slot.slot !== 'legs');
    expect(title(scored, 70)).toBe('Gear · 6 of 14 slots best in slot · +70 points across 7 slots');
    expect(title([GEAR_FIXTURE[0], GEAR_FIXTURE[2]], 18.4)).toBe(
      'Gear · 1 of 2 slots best in slot · +18.4 points across 1 slot',
    );
  });

  it('a total of 0 with an arrow left: "+0 points from the picks you can wear" (fix 2 ruling 5)', () => {
    // The fixture with every scored gap 0: Neck and Off hand are still empty, their picks shown.
    const zero = GEAR_FIXTURE.map((slot) => ({ ...slot, gap: slot.gap === null ? null : 0 }));
    expect(title(zero, 0)).toBe(
      'Gear · 6 of 15 slots best in slot · +0 points from the picks you can wear · 1 not scored',
    );
    expect(
      title(
        zero.filter((slot) => slot.slot !== 'legs'),
        0,
      ),
    ).toBe('Gear · 6 of 14 slots best in slot · +0 points from the picks you can wear');
  });

  it('a total of 0 with no arrow anywhere: "every slot is the best you can wear" (fix 2 ruling 5)', () => {
    // Every pick worn: no row shows an arrow.
    const worn = GEAR_FIXTURE.filter((slot) => slot.pick === null || slot.worn?.id === slot.pick.id);
    expect(title(worn, 0)).toBe('Gear · every slot is the best you can wear');
    expect(title([GEAR_FIXTURE[13], GEAR_FIXTURE[16]], null)).toBe('Gear');
  });

  it('the summary after "Gear" alone, its own dots glued to the word before them (fix 2 ruling 4)', () => {
    expect(gearTitleSummary(GEAR_PANEL_FIXTURE.gear, 70)).toBe(
      '6 of 15 slots best in slot · +70 points across 7 slots · 1 not scored',
    );
    expect(gearTitleSummary([GEAR_FIXTURE[13]], null)).toBeNull();
  });

  it('nulls are counted with any total: "Gear · 2 not scored"', () => {
    const unscored = [GEAR_FIXTURE[8], { ...GEAR_FIXTURE[0], gap: null }];
    expect(title(unscored, null)).toBe('Gear · 2 not scored');
  });
});

describe('gearCopy (L8.6 fix round 1)', () => {
  it('the no-gear line', () => {
    expect(gearCopy.noGear).toBe('No gear on file yet. Sync the addon, or sign in to claim this character.');
    const { lead, sync, or, signIn, claimEnd, end } = gearCopy.noGearParts;
    expect(`${lead}${sync}${or}${signIn}${claimEnd}`).toBe(gearCopy.noGear);
    // The owner has no account to sign in to: sync, and nothing else.
    expect(`${lead}${sync}${end}`).toBe('No gear on file yet. Sync the addon.');
  });

  it('the footer link: "Full Holy BiS list at 60", a band below 60 by its range (fix 1 ruling 3)', () => {
    expect(gearCopy.bisLink('Holy', 60)).toBe('Full Holy BiS list at 60');
    expect(gearCopy.bisLink('Arms', 30)).toBe('Full Arms BiS list at 30 to 39');
    expect(gearCopy.bisLink('Holy', null)).toBe('Full Holy BiS list');
    expect(gearCopy.bisArrow).toBe('→');
  });
});

describe('gearSpecLine (A.2)', () => {
  it('names the spec and says the talents named it', () => {
    expect(gearSpecLine('warrior-fury', 'talents')).toBe('Fury · from your talents');
  });

  it('says no points are spent when the class default answers', () => {
    expect(gearSpecLine('warrior-arms', 'default')).toBe('Arms · no points spent yet');
  });

  it('title-cases a spec of several words', () => {
    expect(gearSpecLine('hunter-beast-mastery', 'talents')).toBe('Beast Mastery · from your talents');
  });

  it('says nothing when a fight named the spec, or the API sent none', () => {
    expect(gearSpecLine('warrior-fury', 'fight')).toBeNull();
    expect(gearSpecLine(null, null)).toBeNull();
    expect(gearSpecLine(undefined, undefined)).toBeNull();
  });
});
