// web/src/lib/sim/character-bootstrap.test.ts
import { describe, expect, it, vi } from 'vitest';
import type { CharacterPath } from '../characters';
import type { CurrentCharacter } from '../current-character';
import {
  decideBootstrap,
  RESTORE_BUSY_KEY,
  runBootstrapRestore,
  settleRestore,
  sourceIdForInstance,
  startBootstrapLoad,
} from './character-bootstrap';
import type { LootSource } from './loot';

function stored(source: CurrentCharacter['source'], ref: string): CurrentCharacter {
  return {
    source,
    ref,
    label: 'Simfury · Fury Warrior',
    classSlug: 'warrior',
    savedAt: '2026-09-21T00:00:00.000Z',
  };
}

describe('decideBootstrap', () => {
  it('prefers a whole request over everything else, including a stored pointer', () => {
    const decision = decideBootstrap(
      { code: 'FS1:1:warrior:orc:0/0/0:', source: 'build', ref: 'b1', hasRequest: true },
      stored('fight', 'abcdefabcdef:1'),
    );
    expect(decision).toEqual({ kind: 'none', restored: false });
  });

  it('prefers ?code= over everything else', () => {
    const decision = decideBootstrap(
      { code: 'FS1:1:warrior:orc:0/0/0:', source: 'build', ref: 'b1' },
      stored('fight', 'abcdefabcdef:1'),
    );
    expect(decision).toEqual({ kind: 'code', code: 'FS1:1:warrior:orc:0/0/0:', restored: false });
  });

  it('loads an addon-sourced ?source=&ref= as an addon paste', () => {
    const decision = decideBootstrap({ code: '', source: 'addon', ref: 'FS1:1:warrior:orc:0/0/0:' }, null);
    expect(decision).toEqual({ kind: 'addon', code: 'FS1:1:warrior:orc:0/0/0:', restored: false });
  });

  it('loads a build-sourced ?source=&ref=', () => {
    const decision = decideBootstrap({ code: '', source: 'build', ref: 'b1' }, null);
    expect(decision).toEqual({ kind: 'build', id: 'b1', restored: false });
  });

  it('loads a fight-sourced ?source=&ref=', () => {
    const decision = decideBootstrap({ code: '', source: 'fight', ref: 'abcdefabcdef:1' }, null);
    expect(decision).toEqual({ kind: 'fight', ref: 'abcdefabcdef:1', restored: false });
  });

  it('loads an armory-sourced ?source=&ref= as a stored CharacterPath', () => {
    const decision = decideBootstrap({ code: '', source: 'armory', ref: 'us/normal/simfury' }, null);
    expect(decision).toEqual({
      kind: 'stored',
      path: { region: 'us', ruleset: 'normal', slug: 'simfury' },
      restored: false,
    });
  });

  it('does not fall through to the stored pointer when ?source=armory&ref= does not parse', () => {
    const decision = decideBootstrap({ code: '', source: 'armory', ref: 'not-a-key' }, stored('build', 'b1'));
    expect(decision).toEqual({ kind: 'none', restored: false });
  });

  it('bootstraps nothing for a ?source=manual link, the same as store.svelte.ts’s own bootstrapSource', () => {
    const decision = decideBootstrap({ code: '', source: 'manual', ref: 'anything' }, null);
    expect(decision).toEqual({ kind: 'none', restored: false });
  });

  it('falls back to the stored pointer when the URL carries no request, no ?code= and no ?source=&ref=', () => {
    const decision = decideBootstrap({ code: '', source: '', ref: '' }, stored('build', 'b1'));
    expect(decision).toEqual({ kind: 'build', id: 'b1', restored: true });
  });

  it('restores an addon-sourced pointer through loadCode, same as a fresh addon paste’s own FS1 string', () => {
    const decision = decideBootstrap(
      { code: '', source: '', ref: '' },
      stored('addon', 'FS1:1:warrior:orc:0/0/0:'),
    );
    expect(decision).toEqual({ kind: 'code', code: 'FS1:1:warrior:orc:0/0/0:', restored: true });
  });

  it('restores a code-sourced pointer through loadCode too', () => {
    const decision = decideBootstrap(
      { code: '', source: '', ref: '' },
      stored('code', 'FS1:1:warrior:orc:0/0/0:'),
    );
    expect(decision).toEqual({ kind: 'code', code: 'FS1:1:warrior:orc:0/0/0:', restored: true });
  });

  it('restores a fight-sourced pointer', () => {
    const decision = decideBootstrap({ code: '', source: '', ref: '' }, stored('fight', 'abcdefabcdef:1'));
    expect(decision).toEqual({ kind: 'fight', ref: 'abcdefabcdef:1', restored: true });
  });

  it('restores an armory-sourced pointer through the parsed CharacterPath', () => {
    const decision = decideBootstrap(
      { code: '', source: '', ref: '' },
      stored('armory', 'us/normal/simfury'),
    );
    expect(decision).toEqual({
      kind: 'stored',
      path: { region: 'us', ruleset: 'normal', slug: 'simfury' },
      restored: true,
    });
  });

  it('is none, not restored, for an armory-sourced pointer whose ref no longer parses', () => {
    const decision = decideBootstrap({ code: '', source: '', ref: '' }, stored('armory', 'not-a-key'));
    expect(decision).toEqual({ kind: 'none', restored: false });
  });

  it('is none when there is no URL bootstrap and no stored pointer', () => {
    expect(decideBootstrap({ code: '', source: '', ref: '' }, null)).toEqual({
      kind: 'none',
      restored: false,
    });
  });

  it('still falls back to the stored pointer when hasRequest is explicitly false', () => {
    const decision = decideBootstrap(
      { code: '', source: '', ref: '', hasRequest: false },
      stored('build', 'b1'),
    );
    expect(decision).toEqual({ kind: 'build', id: 'b1', restored: true });
  });
});

const SOURCES: LootSource[] = [
  { id: 'raid:molten-core', kind: 'raid', name: 'Molten Core' },
  { id: 'dungeon:hall-of-thanes', kind: 'dungeon', name: 'Hall of Thanes' },
];

describe('sourceIdForInstance', () => {
  it('matches a LootSource.id ending with :<slug>', () => {
    expect(sourceIdForInstance(SOURCES, 'molten-core')).toBe('raid:molten-core');
    expect(sourceIdForInstance(SOURCES, 'hall-of-thanes')).toBe('dungeon:hall-of-thanes');
  });

  it('is null for an empty slug', () => {
    expect(sourceIdForInstance(SOURCES, '')).toBeNull();
  });

  it('is null when nothing matches', () => {
    expect(sourceIdForInstance(SOURCES, 'blackwing-lair')).toBeNull();
  });

  it('refuses a slug carrying anything outside lowercase letters, digits and hyphens', () => {
    expect(sourceIdForInstance(SOURCES, 'Molten-Core')).toBeNull();
    expect(sourceIdForInstance(SOURCES, 'molten_core')).toBeNull();
    expect(sourceIdForInstance(SOURCES, '../../etc')).toBeNull();
    expect(sourceIdForInstance(SOURCES, 'molten core')).toBeNull();
  });

  it('refuses a slug past the bounded length rather than reaching the compare at all', () => {
    expect(sourceIdForInstance(SOURCES, 'a'.repeat(65))).toBeNull();
  });
});

describe('settleRestore', () => {
  it('is inert for a URL-driven load, whether it succeeded or failed', () => {
    expect(settleRestore(false, true)).toEqual({
      clearPointer: false,
      restored: false,
      clearMessage: false,
    });
    expect(settleRestore(false, false)).toEqual({
      clearPointer: false,
      restored: false,
      clearMessage: false,
    });
  });

  it('claims restored only once a restore actually produced a character', () => {
    expect(settleRestore(true, true)).toEqual({
      clearPointer: false,
      restored: true,
      clearMessage: false,
    });
  });

  it('forgets a dead or stale pointer and clears the store’s error, rather than showing "Restored" beside it', () => {
    expect(settleRestore(true, false)).toEqual({
      clearPointer: true,
      restored: false,
      clearMessage: true,
    });
  });
});

function fakeLoaders() {
  return {
    loadCode: vi.fn().mockResolvedValue(undefined),
    loadAddon: vi.fn().mockResolvedValue(undefined),
    loadBuild: vi.fn().mockResolvedValue(undefined),
    loadFight: vi.fn().mockResolvedValue(undefined),
    loadStored: vi.fn().mockResolvedValue(undefined),
    setMessage: vi.fn(),
  };
}

const PATH: CharacterPath = { region: 'us', ruleset: 'normal', slug: 'simfury' };

describe('startBootstrapLoad', () => {
  it('calls loadCode for a code decision', async () => {
    const loaders = fakeLoaders();
    await startBootstrapLoad(loaders, { kind: 'code', code: 'FS1:1:warrior:orc:0/0/0:', restored: true });
    expect(loaders.loadCode).toHaveBeenCalledWith('FS1:1:warrior:orc:0/0/0:');
  });

  it('calls loadAddon for an addon decision', async () => {
    const loaders = fakeLoaders();
    await startBootstrapLoad(loaders, { kind: 'addon', code: 'FS1:1:warrior:orc:0/0/0:', restored: false });
    expect(loaders.loadAddon).toHaveBeenCalledWith('FS1:1:warrior:orc:0/0/0:');
  });

  it('calls loadBuild for a build decision', async () => {
    const loaders = fakeLoaders();
    await startBootstrapLoad(loaders, { kind: 'build', id: 'b1', restored: true });
    expect(loaders.loadBuild).toHaveBeenCalledWith('b1');
  });

  it('calls loadFight for a fight decision', async () => {
    const loaders = fakeLoaders();
    await startBootstrapLoad(loaders, { kind: 'fight', ref: 'abcdefabcdef:1', restored: true });
    expect(loaders.loadFight).toHaveBeenCalledWith('abcdefabcdef:1');
  });

  it('calls loadStored for a stored decision', async () => {
    const loaders = fakeLoaders();
    await startBootstrapLoad(loaders, { kind: 'stored', path: PATH, restored: true });
    expect(loaders.loadStored).toHaveBeenCalledWith(PATH);
  });

  it('calls no loader and returns null for a none decision', () => {
    const loaders = fakeLoaders();
    const load = startBootstrapLoad(loaders, { kind: 'none', restored: false });
    expect(load).toBeNull();
    for (const loader of Object.values(loaders)) expect(loader).not.toHaveBeenCalled();
  });
});

function fakeStorage(): Storage {
  const map = new Map<string, string>();
  return {
    getItem: (key) => map.get(key) ?? null,
    setItem: (key, value) => void map.set(key, value),
    removeItem: (key) => void map.delete(key),
    clear: () => map.clear(),
    key: (index) => [...map.keys()][index] ?? null,
    get length() {
      return map.size;
    },
  };
}

describe('runBootstrapRestore', () => {
  it('does nothing and returns false when the URL already won (no double load) — storeHandlesUrl default', async () => {
    const loaders = fakeLoaders();
    const storage = fakeStorage();
    const restored = await runBootstrapRestore(
      loaders,
      { code: 'FS1:1:warrior:orc:0/0/0:', source: '', ref: '' },
      stored('build', 'b1'),
      () => false,
      { storage },
    );
    expect(restored).toBe(false);
    for (const loader of Object.values(loaders)) expect(loader).not.toHaveBeenCalled();
  });

  it('runs the stored pointer through its loader and reports restored on success', async () => {
    const loaders = fakeLoaders();
    const storage = fakeStorage();
    const restored = await runBootstrapRestore(
      loaders,
      { code: '', source: '', ref: '' },
      stored('build', 'b1'),
      () => true,
      { storage },
    );
    expect(restored).toBe(true);
    expect(loaders.loadBuild).toHaveBeenCalledWith('b1');
    expect(loaders.setMessage).not.toHaveBeenCalled();
  });

  it('forgets a dead pointer and clears the message when the restore produced no character', async () => {
    const loaders = fakeLoaders();
    const storage = fakeStorage();
    storage.setItem('fs.currentCharacter', JSON.stringify(stored('build', 'b1')));
    const restored = await runBootstrapRestore(
      loaders,
      { code: '', source: '', ref: '' },
      stored('build', 'b1'),
      () => false,
      { storage },
    );
    expect(restored).toBe(false);
    expect(loaders.setMessage).toHaveBeenCalledWith(null);
    expect(storage.getItem('fs.currentCharacter')).toBeNull();
  });

  /**
   * `storeHandlesUrl: false` (fix round: ToolsView.svelte's own regression -- unlike
   * SimView.svelte's `store.svelte.ts`, `bulk-store.svelte.ts` has no init-time bootstrap of
   * its own, so a direct `?source=&ref=`/`?code=` URL was silently loading nothing at all on
   * the tools island).
   */
  it('storeHandlesUrl false: still runs a URL-driven decision through its loader', async () => {
    const loaders = fakeLoaders();
    const restored = await runBootstrapRestore(
      loaders,
      { code: '', source: 'fight', ref: 'fixture2abcd:3' },
      null,
      () => true,
      { storeHandlesUrl: false },
    );
    // A URL-driven load never claims "restored" -- that word is reserved for the stored
    // pointer (settleRestore's own rule): the chip's "Restored your last character" line
    // would otherwise show for a link the player just followed themselves.
    expect(restored).toBe(false);
    expect(loaders.loadFight).toHaveBeenCalledWith('fixture2abcd:3');
  });

  it('storeHandlesUrl false: a URL-driven `?code=` still loads too', async () => {
    const loaders = fakeLoaders();
    await runBootstrapRestore(
      loaders,
      { code: 'FS1:1:warrior:orc:0/0/0:', source: '', ref: '' },
      null,
      () => true,
      { storeHandlesUrl: false },
    );
    expect(loaders.loadCode).toHaveBeenCalledWith('FS1:1:warrior:orc:0/0/0:');
  });

  it('storeHandlesUrl false: still prefers the stored pointer and reports restored when the URL is bare', async () => {
    const loaders = fakeLoaders();
    const restored = await runBootstrapRestore(
      loaders,
      { code: '', source: '', ref: '' },
      stored('build', 'b1'),
      () => true,
      { storeHandlesUrl: false },
    );
    expect(restored).toBe(true);
    expect(loaders.loadBuild).toHaveBeenCalledWith('b1');
  });

  it('does nothing for a bare URL and no stored pointer, whatever storeHandlesUrl is', async () => {
    const loaders = fakeLoaders();
    const restored = await runBootstrapRestore(
      loaders,
      { code: '', source: '', ref: '' },
      null,
      () => false,
      {
        storeHandlesUrl: false,
      },
    );
    expect(restored).toBe(false);
    for (const loader of Object.values(loaders)) expect(loader).not.toHaveBeenCalled();
  });
});

describe('RESTORE_BUSY_KEY', () => {
  // The sentinel must never collide with a real LandingState character key
  // (`<region>/<ruleset>/<slug>`, always exactly two '/'s).
  it('carries no "/" — never mistakeable for a character key', () => {
    expect(RESTORE_BUSY_KEY).not.toContain('/');
  });
});
