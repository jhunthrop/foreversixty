// web/src/lib/rankings/defaults.test.ts
import { describe, expect, it, vi } from 'vitest';
import type { MeCharacter } from '../account/api';
import {
  applyDefaults,
  pickDefaultEncounter,
  rankingsDefaults,
  reaim,
  resolveDefaultEncounter,
} from './defaults';
import { defaultRankingsState, parseRankingsState, rankingsSearch } from './url';

const NOW = new Date('2026-12-13T12:00:00Z');
const FURY: MeCharacter = {
  key: 'us/normal/obnoxious-yell',
  region: 'us',
  ruleset: 'normal',
  name: 'Obnoxious Yell',
  class: 'Warrior',
  spec: 'Fury',
};

describe('rankingsDefaults', () => {
  it('aims a character at its own class, spec, ruleset and the current phase', () => {
    expect(rankingsDefaults(FURY, NOW)).toEqual({
      class: 'warrior',
      spec: 'Fury',
      ruleset: 'normal',
      phase: 'raids-1',
      metric: 'dps',
    });
  });

  it('picks the metric by role', () => {
    const holy = { ...FURY, class: 'Priest', spec: 'Holy' };
    const protection = { ...FURY, spec: 'Protection' };
    expect(rankingsDefaults(holy, NOW).metric).toBe('hps');
    expect(rankingsDefaults(protection, NOW).metric).toBe('damage_taken');
  });

  it('aims a signed-out visitor at nothing but the current phase', () => {
    expect(rankingsDefaults(null, NOW)).toEqual({
      class: '',
      spec: '',
      ruleset: '',
      phase: 'raids-1',
      metric: 'dps',
    });
  });

  it('keeps a spec the site does not know, and a missing spec, as given', () => {
    expect(rankingsDefaults({ ...FURY, spec: 'Mystery' }, NOW).spec).toBe('Mystery');
    expect(rankingsDefaults({ ...FURY, spec: undefined }, NOW).spec).toBe('');
  });
});

describe('applyDefaults', () => {
  const defaults = rankingsDefaults(FURY, NOW);

  it('fills every aimed filter a link did not name', () => {
    const filled = applyDefaults(defaultRankingsState(), '', defaults);
    expect(filled).toMatchObject({ class: 'warrior', spec: 'Fury', ruleset: 'normal', phase: 'raids-1' });
  });

  it('lets a link that names a filter win, including an explicit all', () => {
    const search = '?class=mage&phase=all';
    const filled = applyDefaults(parseRankingsState(search), search, defaults);
    expect(filled.class).toBe('mage');
    expect(filled.phase).toBe('');
    expect(filled.ruleset).toBe('normal');
  });

  it("reads a link naming a class alone as that whole class, not the viewer's spec of it", () => {
    const search = '?class=mage';
    const filled = applyDefaults(parseRankingsState(search), search, defaults);
    expect(filled).toMatchObject({ class: 'mage', spec: '', ruleset: 'normal' });
  });

  it('writes the choice back so a reload keeps it', () => {
    const state = { ...applyDefaults(defaultRankingsState(), '', defaults), phase: '' };
    expect(rankingsSearch(state, defaults)).toBe('?phase=all');
    expect(rankingsSearch(applyDefaults(defaultRankingsState(), '', defaults), defaults)).toBe('');
  });
});

describe('reaim', () => {
  it('re-aims at the new character on page one and keeps the other filters', () => {
    const before = { ...defaultRankingsState(), region: 'eu', faction: 'horde', page: 4, class: 'mage' };
    const after = reaim(before, rankingsDefaults(FURY, NOW));
    expect(after).toMatchObject({ class: 'warrior', spec: 'Fury', region: 'eu', faction: 'horde', page: 1 });
  });
});

describe('resolveDefaultEncounter', () => {
  const listed = [
    { id: 9001, name: 'Warden Kelthas', slug: 'warden-kelthas' },
    { id: 9002, name: 'Onyxia', slug: 'onyxia' },
  ];
  const reads = (history: { encounter_id: number }[] | Error) => ({
    encounters: () => Promise.resolve({ rows: listed }),
    character: vi.fn(() =>
      history instanceof Error ? Promise.reject(history) : Promise.resolve({ history: history as never }),
    ),
  });

  it("opens on the character's most recently killed encounter", async () => {
    const result = await resolveDefaultEncounter(FURY, reads([{ encounter_id: 9002 }]));
    expect(result).toEqual({ slug: 'onyxia', encounters: listed });
  });

  it('reads the character page at its own path', async () => {
    const source = reads([]);
    await resolveDefaultEncounter(FURY, source);
    expect(source.character).toHaveBeenCalledWith({
      region: 'us',
      ruleset: 'normal',
      slug: 'obnoxious-yell',
    });
  });

  it('does not wait on a character page that fails, or on a signed-out visitor', async () => {
    expect((await resolveDefaultEncounter(FURY, reads(new Error('down')))).slug).toBe('warden-kelthas');
    const signedOut = reads([{ encounter_id: 9002 }]);
    expect((await resolveDefaultEncounter(null, signedOut)).slug).toBe('warden-kelthas');
    expect(signedOut.character).not.toHaveBeenCalled();
  });
});

describe('the encounter default', () => {
  const encounters = [
    { id: 9001, name: 'Warden Kelthas', slug: 'warden-kelthas' },
    { id: 9002, name: 'Onyxia', slug: 'onyxia' },
  ];

  it('is the encounter the character most recently killed', () => {
    const history = [{ encounter_id: 9002 }, { encounter_id: 9001 }];
    expect(pickDefaultEncounter(history, encounters)).toBe('onyxia');
  });

  it('skips a wipe and an encounter the API does not list', () => {
    const history = [{ encounter_id: 9002, kill: false }, { encounter_id: 1234 }, { encounter_id: 9001 }];
    expect(pickDefaultEncounter(history, encounters)).toBe('warden-kelthas');
  });

  it('falls back to the first listed encounter, and to none when nothing is ranked', () => {
    expect(pickDefaultEncounter([], encounters)).toBe('warden-kelthas');
    expect(pickDefaultEncounter([], [])).toBe('');
  });
});
