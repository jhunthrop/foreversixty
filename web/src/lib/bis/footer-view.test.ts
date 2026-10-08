// web/src/lib/bis/footer-view.test.ts
import { describe, expect, it } from 'vitest';
import { newAtBandLinesFor, sourceGroupsFor } from './footer-view';
import type { BandInfo, RowView } from './panel-view';

function row(overrides: Partial<RowView>): RowView {
  return { slot: 'head', empty: false, ...overrides };
}

function bandInfo(overrides: Partial<BandInfo>): BandInfo {
  return {
    band: 20,
    bandIndex: 0,
    rows: [],
    newSlots: new Set(),
    changed: undefined,
    previousBand: undefined,
    upgradesCount: 0,
    unit: 'DPS',
    healer: undefined,
    setDps: 100,
    dpsDelta: undefined,
    setDpsPartial: false,
    setDpsPartialCount: 0,
    race: 'troll',
    talentPoints: 0,
    scaleRows: [],
    scaleNoteLine: '',
    hasteCaptionLine: undefined,
    hitCap: undefined,
    totalSlots: 17,
    ...overrides,
  };
}

describe('sourceGroupsFor (spec §4.E "Where to get it")', () => {
  it('groups dungeon rows by instance, not by boss, sorted by count then alphabetically', () => {
    const rows: RowView[] = [
      row({
        slot: 'shoulder',
        itemId: 1,
        itemName: 'A',
        sourceCell: { kind: 'dungeon', instance: 'Wailing Caverns', boss: 'Lady Anacondra' },
      }),
      row({
        slot: 'back',
        itemId: 2,
        itemName: 'B',
        sourceCell: { kind: 'dungeon', instance: 'Wailing Caverns', boss: 'Skum' },
      }),
      row({
        slot: 'head',
        itemId: 3,
        itemName: 'C',
        sourceCell: { kind: 'crafted', name: 'Leatherworking' },
      }),
    ];
    const groups = sourceGroupsFor(rows);
    expect(groups).toEqual([
      {
        label: 'Wailing Caverns',
        items: [
          { itemId: 1, itemName: 'A' },
          { itemId: 2, itemName: 'B' },
        ],
      },
      { label: 'Leatherworking', items: [{ itemId: 3, itemName: 'C' }] },
    ]);
  });

  it('groups every world_drop source under the auction-house label, regardless of item', () => {
    const rows: RowView[] = [
      row({ slot: 'wrist', itemId: 1, itemName: 'A', sourceCell: { kind: 'world_drop' } }),
      row({ slot: 'feet', itemId: 2, itemName: 'B', sourceCell: { kind: 'world_drop' } }),
    ];
    const groups = sourceGroupsFor(rows);
    expect(groups).toEqual([
      {
        label: 'World drop (auction house)',
        items: [
          { itemId: 1, itemName: 'A' },
          { itemId: 2, itemName: 'B' },
        ],
      },
    ]);
  });

  it('skips an empty row entirely -- never a group for "nothing"', () => {
    const rows: RowView[] = [row({ slot: 'trinket1', empty: true, emptyCopy: 'x' })];
    expect(sourceGroupsFor(rows)).toEqual([]);
  });
});

describe('newAtBandLinesFor (spec §4.E "New at this band")', () => {
  it('is the fixed first-band sentence at bandIndex 0, never the templated "reaching" sentence', () => {
    const current = bandInfo({ bandIndex: 0, rows: [] });
    const { sentence1, sentence2 } = newAtBandLinesFor(
      current,
      20,
      ['head: X'],
      undefined,
      undefined,
      undefined,
    );
    expect(sentence1).toBe('This is the first band — every pick here is new.');
    expect(sentence2).toBeUndefined();
  });

  it('names the real dungeon and rep sources for a later band, with the rep as "the X rewards"', () => {
    const rows: RowView[] = [
      row({ slot: 'shoulder', sourceCell: { kind: 'dungeon', instance: 'Gnomeregan' } }),
      row({ slot: 'neck', sourceCell: { kind: 'rep', faction: 'Warsong Outriders' } }),
    ];
    const current = bandInfo({ bandIndex: 1, rows, totalSlots: 17 });
    const { sentence1 } = newAtBandLinesFor(
      current,
      30,
      ['shoulder: Forest Tracker Epaulets', 'neck: Scout’s Medallion'],
      undefined,
      undefined,
      undefined,
    );
    expect(sentence1).toBe(
      'Reaching 30 opens Gnomeregan and the Warsong Outriders rewards: 2 of the 17 picks come from there.',
    );
  });

  it('names the first slot that clears between this band and the next, never a deeper scan', () => {
    const current = bandInfo({
      bandIndex: 1,
      rows: [
        row({ slot: 'off_hand', empty: true, emptyCopy: 'x' }),
        row({ slot: 'trinket1', empty: true, emptyCopy: 'y' }),
      ],
    });
    const next = bandInfo({
      bandIndex: 2,
      rows: [row({ slot: 'off_hand', empty: false }), row({ slot: 'trinket1', empty: true, emptyCopy: 'y' })],
    });
    const { sentence2 } = newAtBandLinesFor(current, 30, [], next, '40 to 49', []);
    expect(sentence2).toBe('Next band (40 to 49): the first off hand that helps.');
  });

  it('drops the whole next-band sentence at the top band', () => {
    const current = bandInfo({ bandIndex: 4, rows: [] });
    const { sentence2 } = newAtBandLinesFor(current, 60, [], undefined, undefined, undefined);
    expect(sentence2).toBeUndefined();
  });
});
