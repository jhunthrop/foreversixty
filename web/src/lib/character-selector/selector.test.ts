// web/src/lib/character-selector/selector.test.ts
// The nav selector's logic: the stale rule, the list order, the pointer cases and the
// choose-a-character write (design/specs/2026-10-09-nav-character-selector.md decisions 8 to 11).
import { describe, expect, it, vi } from 'vitest';
import type { Me, MeCharacter } from '../account/api';
import { readCurrent, writeCurrent, type CurrentCharacter } from '../current-character';
import { chooseCharacter, urlForChosenPointer, type ChooseDeps } from './choose';
import { CHARACTER_ORDER_STORAGE_KEY, MAX_CHOSEN_KEYS } from './config';
import { orderCharacters, readChosenKeys, recordChosenKey, withChosenFirst } from './order';
import {
  buildSelectorModel,
  closedAriaLabel,
  closedSecondLine,
  filterRows,
  ringMarkOf,
  rowLineThree,
  rowLineTwo,
} from './rows';
import { isFailed, isStale, staleThresholdSec } from './stale';

const NOW = new Date('2026-10-09T12:00:00Z');
const DAY_SEC = 86_400;

function memoryStorage(initial: Record<string, string> = {}): Storage {
  const data = new Map(Object.entries(initial));
  return {
    get length() {
      return data.size;
    },
    clear: () => data.clear(),
    getItem: (key) => data.get(key) ?? null,
    key: (index) => [...data.keys()][index] ?? null,
    removeItem: (key) => void data.delete(key),
    setItem: (key, value) => void data.set(key, value),
  };
}

function daysAgo(days: number): string {
  return new Date(NOW.getTime() - days * DAY_SEC * 1000).toISOString();
}

function character(name: string, overrides: Partial<MeCharacter> = {}): MeCharacter {
  return {
    key: `us/normal/${name.toLowerCase()}`,
    region: 'us',
    ruleset: 'normal',
    name,
    class: 'Warrior',
    spec: 'Fury',
    level: 60,
    realm: 'Living Flame',
    faction: 'alliance',
    ...overrides,
  };
}

function build(
  days: number,
  extra: Partial<NonNullable<MeCharacter['build']>> = {},
): NonNullable<MeCharacter['build']> {
  return { source: 'addon', captured_at: daysAgo(days), ...extra };
}

function me(characters: MeCharacter[], main?: string): Me {
  return {
    user: { id: 1, battletag: null, email: null, role: 'user', anonymize: false },
    characters,
    guilds: [],
    main_character_key: main,
  };
}

describe('the stale rule', () => {
  it('is three times the median gap, clamped to 3 to 14 days', () => {
    expect(staleThresholdSec(DAY_SEC)).toBe(3 * DAY_SEC);
    expect(staleThresholdSec(0.5 * DAY_SEC)).toBe(3 * DAY_SEC);
    expect(staleThresholdSec(2 * DAY_SEC)).toBe(6 * DAY_SEC);
    expect(staleThresholdSec(30 * DAY_SEC)).toBe(14 * DAY_SEC);
  });

  it('falls back to 7 days when the API sends no median', () => {
    expect(staleThresholdSec(null)).toBe(7 * DAY_SEC);
    expect(staleThresholdSec(undefined)).toBe(7 * DAY_SEC);
    expect(staleThresholdSec(0)).toBe(7 * DAY_SEC);
  });

  it('flags a nightly raider after three days and a weekend player not until fourteen', () => {
    const nightly = build(4, { median_sync_gap_sec: DAY_SEC });
    const weekend = build(10, { median_sync_gap_sec: 7 * DAY_SEC });
    expect(isStale(nightly, NOW)).toBe(true);
    expect(isStale(weekend, NOW)).toBe(false);
    expect(isStale(build(15, { median_sync_gap_sec: 7 * DAY_SEC }), NOW)).toBe(true);
  });

  it('uses the 7-day fallback with a null median, and never flags a character with no build', () => {
    expect(isStale(build(6, { median_sync_gap_sec: null }), NOW)).toBe(false);
    expect(isStale(build(8, { median_sync_gap_sec: null }), NOW)).toBe(true);
    expect(isStale(undefined, NOW)).toBe(false);
    expect(isStale({ source: 'addon', captured_at: 'not a date' }, NOW)).toBe(false);
  });

  it('reads failed only from a non-null sync_error', () => {
    expect(isFailed(build(1, { sync_error: 'bnet 502' }))).toBe(true);
    expect(isFailed(build(1, { sync_error: null }))).toBe(false);
    expect(isFailed(build(1))).toBe(false);
    expect(isFailed(undefined)).toBe(false);
  });
});

describe('the list order', () => {
  const current = character('Current', { build: build(0.01) });
  const chosenRecent = character('Recent', { build: build(9) });
  const chosenOlder = character('Older', { build: build(8) });
  const newest = character('Newest', { build: build(3) });
  const oldest = character('Oldest', { build: build(19) });
  const unsynced = character('Unsynced');
  const all = [oldest, newest, unsynced, chosenOlder, current, chosenRecent];

  it('puts the current character first, then the chosen ones in choice order, then newest sync', () => {
    const ordered = orderCharacters(all, current.key, [chosenRecent.key, chosenOlder.key]);
    expect(ordered.map((c) => c.name)).toEqual([
      'Current',
      'Recent',
      'Older',
      'Newest',
      'Oldest',
      'Unsynced',
    ]);
  });

  it('does not let a sync move a character the player has chosen', () => {
    const synced = { ...chosenRecent, build: build(0.001) };
    const ordered = orderCharacters(
      all.map((c) => (c.key === synced.key ? synced : c)),
      current.key,
      [chosenOlder.key, chosenRecent.key],
    );
    expect(ordered.slice(0, 3).map((c) => c.name)).toEqual(['Current', 'Older', 'Recent']);
  });

  it('ignores chosen keys that are no longer on the account and the current key among them', () => {
    const ordered = orderCharacters(all, current.key, ['gone', current.key, newest.key]);
    expect(ordered.map((c) => c.name).slice(0, 2)).toEqual(['Current', 'Newest']);
  });

  it('keeps fs.characterOrder to twenty keys, most recent first, without duplicates', () => {
    const storage = memoryStorage();
    for (let i = 0; i < MAX_CHOSEN_KEYS + 5; i += 1) recordChosenKey(`k${i}`, storage);
    recordChosenKey('k10', storage);
    const keys = readChosenKeys(storage);
    expect(keys).toHaveLength(MAX_CHOSEN_KEYS);
    expect(keys[0]).toBe('k10');
    expect(new Set(keys).size).toBe(MAX_CHOSEN_KEYS);
    expect(withChosenFirst(['a', 'b'], 'b')).toEqual(['b', 'a']);
  });

  it('reads malformed storage as empty', () => {
    expect(readChosenKeys(memoryStorage({ [CHARACTER_ORDER_STORAGE_KEY]: '{oops' }))).toEqual([]);
    expect(readChosenKeys(memoryStorage({ [CHARACTER_ORDER_STORAGE_KEY]: '{"a":1}' }))).toEqual([]);
    expect(readChosenKeys(memoryStorage({ [CHARACTER_ORDER_STORAGE_KEY]: '["a",3]' }))).toEqual(['a']);
  });
});

describe('the model', () => {
  const pointerFor = (overrides: Partial<CurrentCharacter>): CurrentCharacter => ({
    source: 'armory',
    ref: 'us/normal/tester',
    label: 'Tester · Fury Warrior',
    classSlug: 'warrior',
    savedAt: daysAgo(2),
    ...overrides,
  });
  const tester = character('Tester', { build: build(0.01) });
  const frost = character('Frost', { class: 'Mage', spec: 'Frost', build: build(2, { source: 'blizzard' }) });

  it('marks the pointed account character current, first, and reads addon characters', () => {
    const model = buildSelectorModel({
      me: me([frost, tester]),
      pointer: pointerFor({}),
      chosenKeys: [],
      nowMs: NOW.getTime(),
    });
    expect(model.session).toBe('signed-in');
    expect(model.rows.map((row) => [row.name, row.current])).toEqual([
      ['Tester', true],
      ['Frost', false],
    ]);
    expect(model.hasAddonCharacter).toBe(true);
  });

  it('falls back to the main when there is no pointer, and offers the addon with no addon character', () => {
    const model = buildSelectorModel({
      me: me([character('Solo', { build: build(1, { source: 'blizzard' }) })]),
      pointer: null,
      chosenKeys: [],
      nowMs: NOW.getTime(),
    });
    expect(model.current?.name).toBe('Solo');
    expect(model.hasAddonCharacter).toBe(false);
  });

  it('shows a code or addon pointer matching no account character as a Pasted export row', () => {
    const model = buildSelectorModel({
      me: me([tester]),
      pointer: pointerFor({ source: 'code', ref: 'FS1.abc', label: 'Pasty · Frost Mage', classSlug: 'mage' }),
      chosenKeys: [],
      nowMs: NOW.getTime(),
    });
    expect(model.rows[0]).toMatchObject({
      name: 'Pasty',
      spec: 'Frost',
      className: 'Mage',
      sourceWord: 'Pasted export',
      current: true,
      kind: 'pointer',
    });
    expect(model.rows[1]).toMatchObject({ name: 'Tester', current: false });
  });

  it('matches a pointer whose label carries the key slug to the account row and shows its display name', () => {
    const bow = character('Bow Jackzon', { key: 'us/normal/bow-jackzon', class: 'Hunter', spec: 'Beast Mastery' });
    const model = buildSelectorModel({
      me: me([tester, bow]),
      pointer: pointerFor({
        source: 'addon',
        ref: 'FS1.def',
        label: 'bow-jackzon · Beast Mastery Hunter',
        classSlug: 'hunter',
      }),
      chosenKeys: [],
      nowMs: NOW.getTime(),
    });
    expect(model.rows).toHaveLength(2);
    expect(model.current).toMatchObject({ name: 'Bow Jackzon', kind: 'account' });
  });

  it('treats a code or addon pointer naming an account character as that character, never twice', () => {
    const model = buildSelectorModel({
      me: me([tester]),
      pointer: pointerFor({
        source: 'addon',
        ref: 'FS1.def',
        label: 'Tester · Arms Warrior',
        classSlug: 'warrior',
      }),
      chosenKeys: [],
      nowMs: NOW.getTime(),
    });
    expect(model.rows).toHaveLength(1);
    expect(model.rows[0]).toMatchObject({ name: 'Tester', kind: 'account', current: true });
    expect(model.current?.key).toBe(tester.key);
  });

  it('is signed out with no rows and no pointer, and with the pasted row when a paste is stored', () => {
    const empty = buildSelectorModel({ me: null, pointer: null, chosenKeys: [], nowMs: NOW.getTime() });
    expect(empty).toMatchObject({ session: 'signed-out', rows: [], current: null });
    const pasted = buildSelectorModel({
      me: null,
      pointer: pointerFor({ source: 'addon', label: 'Warrior' }),
      chosenKeys: [],
      nowMs: NOW.getTime(),
    });
    expect(pasted.session).toBe('signed-out');
    expect(pasted.current).toMatchObject({ name: 'Warrior', sourceWord: 'Pasted export' });
  });

  it('is expired when an armory pointer outlives the session', () => {
    const model = buildSelectorModel({
      me: null,
      pointer: pointerFor({}),
      chosenKeys: [],
      nowMs: NOW.getTime(),
    });
    expect(model.session).toBe('expired');
    expect(ringMarkOf(model)).toBe('session');
    expect(model.current?.name).toBe('Tester');
  });

  it('shows the pointer while /v1/me is still loading', () => {
    const model = buildSelectorModel({
      me: undefined,
      pointer: pointerFor({}),
      chosenKeys: [],
      nowMs: NOW.getTime(),
    });
    expect(model.session).toBe('loading');
    expect(model.current?.name).toBe('Tester');
  });

  it('marks stale and failed rows, and the ring mark follows the current one', () => {
    const stale = character('Stale', { build: build(19) });
    const failed = character('Failed', {
      build: build(3, { source: 'blizzard', sync_error: 'refresh failed' }),
    });
    const model = buildSelectorModel({
      me: me([stale, failed]),
      pointer: pointerFor({ ref: stale.key }),
      chosenKeys: [],
      nowMs: NOW.getTime(),
    });
    expect(model.rows.map((row) => row.status)).toEqual(['stale', 'failed']);
    expect(ringMarkOf(model)).toBe('stale');
    expect(rowLineThree(model.rows[0])).toMatchObject({
      text: 'Addon · 19 days ago',
      tail: 'Log in to the game to update',
    });
    expect(rowLineThree(model.rows[1])).toMatchObject({
      text: 'Battle.net refresh failed · last good 3 days ago',
      failed: true,
    });
  });
});

describe('row copy', () => {
  const row = (overrides: Partial<MeCharacter> = {}) =>
    buildSelectorModel({
      me: me([character('Obnoxious Yell', { build: build(0.01), ...overrides })]),
      pointer: null,
      chosenKeys: [],
      nowMs: NOW.getTime(),
    }).rows[0];

  it('writes line 2 as spec class, level and realm, with the ruleset unless normal', () => {
    expect(rowLineTwo(row())).toBe('Fury Warrior · 60 · Living Flame');
    expect(rowLineTwo(row({ ruleset: 'hardcore' }))).toBe('Fury Warrior · 60 · Living Flame · Hardcore');
    expect(rowLineTwo(row({ realm: undefined, level: undefined, spec: undefined }))).toBe('Warrior');
  });

  it('writes the closed second line with and without the class', () => {
    expect(closedSecondLine(row(), true)).toBe('Fury Warrior · 60');
    expect(closedSecondLine(row(), false)).toBe('Fury · 60');
    expect(closedSecondLine(row({ spec: undefined }), false)).toBe('Warrior · 60');
  });

  it('writes the aria label with the mark sentence before the action', () => {
    expect(closedAriaLabel(row(), null)).toBe('Obnoxious Yell, Fury Warrior, level 60. Change character.');
    expect(closedAriaLabel(row(), 'stale')).toBe(
      'Obnoxious Yell, Fury Warrior, level 60. Sync is stale. Change character.',
    );
  });

  it('keeps the plain age and the source on a Battle.net stale row and links the refresh', () => {
    const stale = row({ build: build(19, { source: 'blizzard' }) });
    expect(rowLineThree(stale)).toEqual({
      text: 'Battle.net · 19 days ago',
      tail: 'Refresh from your account',
      failed: false,
    });
  });

  it('filters by name, ignoring case', () => {
    const rows = buildSelectorModel({
      me: me([character('Alpha'), character('Beta'), character('Alphonse')]),
      pointer: null,
      chosenKeys: [],
      nowMs: NOW.getTime(),
    }).rows;
    expect(
      filterRows(rows, ' alph ')
        .map((r) => r.name)
        .sort(),
    ).toEqual(['Alphonse', 'Alpha'].sort());
    expect(filterRows(rows, '')).toHaveLength(3);
  });
});

describe('choosing a character', () => {
  function deps(pathname: string, storage: Storage): ChooseDeps & { events: Event[]; urls: string[] } {
    const events: Event[] = [];
    const urls: string[] = [];
    return {
      storage,
      pathname,
      events,
      urls,
      dispatch: (event) => void events.push(event),
      replaceUrl: (href) => void urls.push(href),
      reload: vi.fn(),
      now: () => NOW,
    };
  }
  const tester = character('Tester', { spec: 'Fury' });

  it('writes an armory pointer keyed by the character, remembers it and announces it', () => {
    const storage = memoryStorage();
    const wired = deps('/tiers', storage);
    chooseCharacter(tester, wired);
    expect(readCurrent(storage)).toMatchObject({
      source: 'armory',
      ref: tester.key,
      label: 'Tester · Fury Warrior',
      classSlug: 'warrior',
      savedAt: NOW.toISOString(),
    });
    expect(readChosenKeys(storage)).toEqual([tester.key]);
    expect(wired.events.map((event) => event.type)).toEqual(['fs:current-character']);
    expect(wired.urls).toEqual([]);
    expect(wired.reload).not.toHaveBeenCalled();
  });

  it('points the planner and simulator URLs at the choice and loads them once more', () => {
    const planner = deps('/planner', memoryStorage());
    chooseCharacter(tester, planner);
    expect(planner.urls).toEqual(['/planner?class=warrior']);
    expect(planner.reload).toHaveBeenCalledTimes(1);

    const sim = deps('/sim/gear', memoryStorage());
    chooseCharacter(tester, sim);
    expect(sim.urls).toEqual([`/sim/gear?source=armory&ref=${encodeURIComponent(tester.key)}`]);
    expect(sim.reload).toHaveBeenCalledTimes(1);
  });

  it('leaves other pages alone', () => {
    const pointer = writePointerFor('/guides');
    expect(urlForChosenPointer('/guides', pointer)).toBeNull();
    expect(urlForChosenPointer('/bis', pointer)).toBeNull();
  });

  function writePointerFor(path: string): CurrentCharacter {
    const storage = memoryStorage();
    chooseCharacter(tester, deps(path, storage));
    const pointer = readCurrent(storage);
    if (pointer === null) throw new Error('pointer not written');
    writeCurrent(pointer, storage);
    return pointer;
  }
});
