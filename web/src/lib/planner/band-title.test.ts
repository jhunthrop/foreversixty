// web/src/lib/planner/band-title.test.ts
import { describe, expect, it } from 'vitest';
import { bandTitleForDraft } from './band-compare';
import { bandBuildTitle, titleAfterBandLoad } from './band-title';
import { createPlannerStore } from './store.svelte';

describe('bandBuildTitle', () => {
  it('is the band label alone when the build has no spec', () => {
    expect(bandBuildTitle('20 to 29', null)).toBe('20 to 29');
    expect(bandBuildTitle('20 to 29', '')).toBe('20 to 29');
  });

  it('joins the spec when the band build has one', () => {
    expect(bandBuildTitle('20 to 29', 'Beast Mastery')).toBe('20 to 29 · Beast Mastery');
  });
});

describe('bandTitleForDraft', () => {
  it('names the spec the points sit in, without the class', () => {
    expect(bandTitleForDraft('20 to 29', 'hunter', [11, 0, 0])).toBe('20 to 29 · Beast Mastery');
    expect(bandTitleForDraft('30 to 39', 'hunter', [0, 0, 12])).toBe('30 to 39 · Survival');
  });
});

describe('titleAfterBandLoad', () => {
  it('takes the band title when the draft is untitled', () => {
    expect(titleAfterBandLoad('', '', '20 to 29')).toBe('20 to 29');
  });

  it('replaces the title the previous band load gave', () => {
    expect(titleAfterBandLoad('20 to 29', '20 to 29', '30 to 39')).toBe('30 to 39');
  });

  it('keeps a title the player typed', () => {
    expect(titleAfterBandLoad('My dungeon spec', '20 to 29', '30 to 39')).toBe('My dungeon spec');
  });
});

describe('store.nameBandBuild', () => {
  const make = () =>
    createPlannerStore({ treeVersion: '1.15.9.69722', classSlug: 'warrior', raceSlug: 'human' });

  it('titles the draft, so toDraft carries the band build name', () => {
    const store = make();
    store.nameBandBuild('20 to 29 · Beast Mastery');
    expect(store.title).toBe('20 to 29 · Beast Mastery');
  });

  it('lets a player edit win over a later band load, and keeps an unchanged title', () => {
    const store = make();
    store.nameBandBuild('20 to 29');
    store.nameBandBuild('20 to 29');
    expect(store.title).toBe('20 to 29');
    store.setTitle('Leveling Hunter');
    store.nameBandBuild('30 to 39');
    expect(store.title).toBe('Leveling Hunter');
  });

  it('follows the band while the title is still the band title', () => {
    const store = make();
    store.nameBandBuild('20 to 29');
    store.nameBandBuild('30 to 39');
    expect(store.title).toBe('30 to 39');
  });
});
