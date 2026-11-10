import { describe, expect, it } from 'vitest';
import type { ItemTooltipModel } from '../items/tooltip';
import type { GuideBis, GuideBisPanel, GuideBisRow } from './guide-bis';
import { dollForFaction, guideDollCells, guideDollTooltips, guideDollViews } from './guide-bis-doll';

function pick(slot: string, itemId: number, name: string): GuideBisRow {
  return {
    slot,
    slotLabel: slot,
    empty: false,
    itemId,
    itemName: name,
    quality: 4,
    icon: `inv_${itemId}`,
    sourceKind: 'raid',
    sourceDetail: 'Molten Core',
    notSimChecked: false,
  };
}

const emptyOffHand: GuideBisRow = {
  slot: 'off_hand',
  slotLabel: 'Off hand',
  empty: true,
  emptyCopy: 'Two-hander equipped',
  notSimChecked: false,
};

function panel(faction: 'alliance' | 'horde', band: number, rows: GuideBisRow[]): GuideBisPanel {
  return {
    faction,
    band,
    label: String(band),
    bisAnchor: `band-${faction}-${band}`,
    raceName: faction === 'alliance' ? 'Human' : 'Orc',
    rows,
    preRaid: false,
  };
}

function tooltip(slotLabel: string): ItemTooltipModel {
  return { slotLabel } as ItemTooltipModel;
}

const TWO_HANDER = 100;
const tips = new Map<number, ItemTooltipModel>([
  [TWO_HANDER, tooltip('Two-Hand')],
  [200, tooltip('Head')],
]);
const lookup = (id: number): ItemTooltipModel | undefined => tips.get(id);

describe('the guide BiS doll', () => {
  it('maps each lane row to a cell with its item, source and slot', () => {
    const cells = guideDollCells(panel('alliance', 60, [pick('head', 200, 'Lionheart Helm')]), lookup);
    expect(cells).toEqual([
      {
        slot: 'head',
        label: 'head',
        item: { id: 200, name: 'Lionheart Helm', quality: 4, icon: 'inv_200' },
        source: 'Molten Core',
        sourceKind: 'raid',
        empty: undefined,
        coveredBy: undefined,
      },
    ]);
  });

  it('greys a two-hander into the empty off hand, and only a two-hander', () => {
    const two = guideDollCells(
      panel('alliance', 60, [pick('main_hand', TWO_HANDER, 'Arcanite Reaper'), emptyOffHand]),
      lookup,
    );
    expect(two[1]).toMatchObject({ item: null, coveredBy: { id: TWO_HANDER }, empty: 'Two-hander equipped' });
    const one = guideDollCells(
      panel('alliance', 60, [pick('main_hand', 200, 'Sword'), emptyOffHand]),
      lookup,
    );
    expect(one[1]?.coveredBy).toBeUndefined();
  });

  it('keeps only level 60, the default faction first, and follows the reader to the other faction', () => {
    const bis: GuideBis = {
      defaultFaction: 'horde',
      factions: ['alliance', 'horde'],
      defaultBand: 60,
      bands: [],
      panels: [
        panel('alliance', 50, [pick('head', 200, 'A50')]),
        panel('alliance', 60, [pick('head', 200, 'A60')]),
        panel('horde', 60, [pick('head', 200, 'H60')]),
      ],
    };
    const views = guideDollViews(bis, lookup);
    expect(views.map((view) => view.faction)).toEqual(['horde', 'alliance']);
    expect(dollForFaction(views, 'alliance')?.raceName).toBe('Human');
    expect(dollForFaction(views, undefined)?.faction).toBe('horde');
    expect(dollForFaction(views, 'neutral')?.faction).toBe('horde');
    expect(guideDollViews(undefined, lookup)).toEqual([]);
  });

  it('carries each pick tooltip once, skipping items with none', () => {
    const views = guideDollViews(
      {
        defaultFaction: 'alliance',
        factions: ['alliance'],
        defaultBand: 60,
        bands: [],
        panels: [panel('alliance', 60, [pick('head', 200, 'Helm'), pick('chest', 300, 'Unknown')])],
      },
      lookup,
    );
    expect(Object.keys(guideDollTooltips(views, lookup))).toEqual(['200']);
  });
});
