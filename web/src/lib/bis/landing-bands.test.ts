// web/src/lib/bis/landing-bands.test.ts
import { describe, expect, it } from 'vitest';
import {
  DEFAULT_LANDING_BAND,
  LANDING_BANDS,
  bandCaptions,
  bandLabel,
  parseBandParam,
  selectedLandingBand,
  topSourceWords,
} from './landing-bands';

const picks = (...kinds: string[]) => kinds.map((source_kind) => ({ source_kind }));

describe('the strip’s bands', () => {
  it('are the five bands, labelled as ranges with a plain top band', () => {
    expect(LANDING_BANDS).toEqual([20, 30, 40, 50, 60]);
    expect(LANDING_BANDS.map(bandLabel)).toEqual(['20 to 29', '30 to 39', '40 to 49', '50 to 59', '60']);
  });

  it('default to the first band, where launch players are', () => {
    expect(DEFAULT_LANDING_BAND).toBe(20);
  });
});

describe('parseBandParam', () => {
  it('accepts only a band the strip has', () => {
    expect(parseBandParam('40')).toBe(40);
    expect(parseBandParam('45')).toBeNull();
    expect(parseBandParam('70')).toBeNull();
    expect(parseBandParam('abc')).toBeNull();
    expect(parseBandParam('')).toBeNull();
    expect(parseBandParam(null)).toBeNull();
    expect(parseBandParam(undefined)).toBeNull();
  });
});

describe('selectedLandingBand', () => {
  it('prefers the viewer’s own choice, then the character’s band, then 20 to 29', () => {
    expect(selectedLandingBand(50, 40)).toBe(50);
    expect(selectedLandingBand(null, 40)).toBe(40);
    expect(selectedLandingBand(null, undefined)).toBe(20);
  });
});

describe('topSourceWords', () => {
  it('names the two most common kinds, merging both world-drop spellings and skipping unsourced rows', () => {
    const slots = picks('dungeon', 'world', 'world_drop', 'dungeon', 'world', 'quest', '', 'crafted');
    expect(topSourceWords(slots)).toEqual(['world drops', 'dungeons']);
  });

  it('keeps first-seen order on a tie', () => {
    expect(topSourceWords(picks('vendor', 'quest', 'rep'))).toEqual(['vendors', 'quests']);
  });

  it('is empty when nothing is sourced', () => {
    expect(topSourceWords(picks('', 'unknown_kind'))).toEqual([]);
  });
});

describe('bandCaptions', () => {
  const byBand = new Map([
    [20, picks('dungeon', 'dungeon', 'crafted')],
    [30, picks('dungeon', 'crafted', 'dungeon')],
    [40, picks('dungeon', 'dungeon', 'crafted')],
    [50, picks('vendor', 'vendor', 'dungeon')],
    [60, picks('vendor', 'quest', 'vendor')],
  ]);
  const captions = bandCaptions(byBand);

  it('gives the first band its sources as the baseline', () => {
    expect(captions[0]).toMatchObject({ label: '20 to 29', text: 'Top sources: dungeons, crafting.', same: false });
  });

  it('says "Same sources as" the band before where the sources match, so no pill is bare', () => {
    expect(captions[1]).toMatchObject({ text: 'Same sources as 20 to 29', same: true });
    expect(captions[2]).toMatchObject({ text: 'Same sources as 30 to 39', same: true });
  });

  it('gives a band that differs its own sources again', () => {
    expect(captions[3]).toMatchObject({ text: 'Top sources: vendors, dungeons.', same: false });
    expect(captions[4]).toMatchObject({ label: '60', text: 'Top sources: vendors, quests.', same: false });
  });

  it('has one caption per band', () => {
    expect(captions.map((caption) => caption.band)).toEqual([...LANDING_BANDS]);
  });
});
