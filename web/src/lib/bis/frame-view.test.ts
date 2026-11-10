import { describe, expect, it } from 'vitest';
import { placeRange, rowFigure, rowLine, simWeightsHref, stackView } from './frame-view';
import type { BisBandModel } from './page-model';
import type { ScaleRow } from './scale-view';
import type { DollCell } from '../ui/paperdoll';

function row(stat: string, label: string, scaleFactor: number, significant = true): ScaleRow {
  return {
    stat,
    label,
    scaleFactor,
    dpsPerPoint: scaleFactor * 0.316,
    scaleError: 0.01,
    perPointUnit: 'DPS',
    significant,
    reading: significant ? 'value' : 'not-significant',
    barPercent: scaleFactor * 100,
  } as ScaleRow;
}

function model(rows: ScaleRow[]): BisBandModel {
  return { weights: { rows, note: '' } } as unknown as BisBandModel;
}

const cell = (source: string, label = 'Back'): DollCell => ({
  slot: 'back',
  label,
  item: { id: 1, name: 'Glowing Lizardscale Cloak', quality: 3 },
  source,
});

describe('stackView', () => {
  it('keeps the four loudest significant stats and says what 1.00 is worth', () => {
    const view = stackView(
      model([
        row('agility', 'Agility', 1),
        row('spirit', 'Spirit', 0.9, false),
        row('intellect', 'Intellect', 0.51),
        row('rap', 'Ranged attack power', 0.32),
        row('hit', 'Hit rating', 0.31),
        row('crit', 'Crit rating', 0.21),
      ]),
    );
    expect(view.rows.map((r) => r.stat)).toEqual(['agility', 'intellect', 'rap', 'hit']);
    expect(view.unitLine).toBe('1.00 = 0.316 DPS');
  });

  it('has no unit line when nothing is significant', () => {
    expect(stackView(model([row('agility', 'Agility', 1, false)]))).toEqual({ rows: [] });
  });
});

describe('the itinerary lines', () => {
  it('takes the stop range from the instance the heading names', () => {
    expect(placeRange('Skum · Wailing Caverns (15-24) · 50%', 'Wailing Caverns')).toBe('15 to 24');
  });

  it('leaves a world drop range on its row', () => {
    expect(placeRange('World drop (20-26) · BoE · on the AH', 'World / BoE')).toBeUndefined();
  });

  it('says the slot and only what the heading does not', () => {
    expect(rowLine(cell('Skum · Wailing Caverns (15-24) · 50%'), 'Wailing Caverns')).toBe(
      'Back · Skum · 50%',
    );
    expect(rowLine(cell('Crafted · Leatherworking', 'Head'), 'Crafted · Leatherworking')).toBe('Head');
    expect(rowLine(cell('Crafted · Leatherworking (skill 300)', 'Head'), 'Crafted · Leatherworking')).toBe(
      'Head · skill 300',
    );
    expect(
      rowLine(cell('Silverwing Sentinels reputation · Honored', 'Neck'), 'Silverwing Sentinels reputation'),
    ).toBe('Neck · Honored');
  });
});

describe('simWeightsHref', () => {
  it('opens the same set on the weights page', () => {
    expect(simWeightsHref('/sim?code=abc')).toBe('/sim/weights?code=abc');
    expect(simWeightsHref(undefined)).toBeUndefined();
  });
});

describe('rowFigure', () => {
  it('shows a gain, a worn pick or the row state, and nothing for a tie', () => {
    expect(rowFigure({ value: '+0.3', unit: 'DPS', tone: 'strong' })).toEqual({
      value: '+0.3',
      unit: 'DPS',
      tone: 'gain',
    });
    expect(rowFigure({ value: 'Same', tone: 'muted' })).toBeUndefined();
    expect(rowFigure({ value: 'Worn', tone: 'kill' })).toEqual({ value: 'Worn', tone: 'kill' });
    expect(rowFigure({ value: '\u2014', basis: { prefix: 'Only option' }, tone: 'muted' })).toEqual({
      value: 'Only option',
      tone: 'muted',
    });
    expect(rowFigure(undefined)).toBeUndefined();
  });
});
