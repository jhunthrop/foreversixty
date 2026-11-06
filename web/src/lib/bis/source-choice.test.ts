import { describe, expect, it } from 'vitest';
import type { BisBandModel } from './page-model';
import {
  defaultRuleLine,
  setFigureFor,
  sourceChoiceFor,
  sourcesButtonLabel,
  sourcesFieldLabel,
} from './source-choice';
import { SOURCE_FILTER_KINDS } from './source-filter';

const ALL = SOURCE_FILTER_KINDS;
const NO_REP = ALL.filter((kind) => kind !== 'rep');

function model(band: number): BisBandModel {
  return {
    band,
    unit: 'DPS',
    answer: {
      figure: '110.9',
      error: 0.2,
      simmed: 'Simmed 5 Nov',
      trust: [],
      plannerHref: '/planner',
      filtered: {
        dungeon: { dps: 105.1, error: 0.3, changed_slots: 4 },
        leveling_default: { dps: 109.0, error: 0.2, changed_slots: 1 },
      },
    },
  } as unknown as BisBandModel;
}

describe('sourceChoiceFor', () => {
  it('keeps reputation up to Friendly only for the leveling default', () => {
    expect(sourceChoiceFor(NO_REP, 20)).toEqual({ on: NO_REP, repCap: 'friendly' });
  });

  it('turns reputation off entirely for Dungeon only below 60', () => {
    expect(sourceChoiceFor(['dungeon'], 20)).toEqual({ on: ['dungeon'] });
    expect(sourceChoiceFor(['dungeon', 'quest'], 30)).toEqual({ on: ['dungeon', 'quest'] });
  });

  it('never caps at 60', () => {
    expect(sourceChoiceFor(NO_REP, 60)).toEqual({ on: NO_REP });
  });
});

describe('setFigureFor', () => {
  it("uses the sim's dungeon-only set at every band it is published", () => {
    for (const band of [20, 60]) {
      const figure = setFigureFor(model(band), ['dungeon'], 4);
      expect(figure?.value).toBe('105.1');
      expect(figure?.label).toBe('Set: dungeon only');
      expect(figure?.caption).toBe('5.8 DPS under the best set');
    }
  });

  it('uses the leveling default set below 60', () => {
    expect(setFigureFor(model(20), NO_REP, 1)?.value).toBe('109.0');
  });

  it('labels the unfiltered figure when the chosen set is not simmed', () => {
    const figure = setFigureFor(model(20), ['dungeon', 'quest'], 3);
    expect(figure?.value).toBe('110.9');
    expect(figure?.label).toBe('Best set, all sources');
  });
});

describe('labels', () => {
  it('names the selection', () => {
    expect(sourcesFieldLabel(ALL, 60)).toBe('All sources');
    expect(sourcesFieldLabel(['dungeon'], 20)).toBe('Dungeon only');
    expect(sourcesFieldLabel(['dungeon', 'quest'], 20)).toBe('2 of 8 sources');
  });

  it('counts the kinds on for the compact button', () => {
    expect(sourcesButtonLabel(ALL)).toBe('All sources');
    expect(sourcesButtonLabel(NO_REP)).toBe('Sources (7)');
    expect(sourcesButtonLabel(['dungeon'])).toBe('Sources (1)');
  });

  it('states the default rule only for the leveling default', () => {
    expect(defaultRuleLine(NO_REP, 20)).toMatch(/Reputation above Friendly is off/);
    expect(defaultRuleLine(['dungeon'], 20)).toBeUndefined();
  });
});
