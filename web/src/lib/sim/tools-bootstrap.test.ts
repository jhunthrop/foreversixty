// web/src/lib/sim/tools-bootstrap.test.ts
import { describe, expect, it } from 'vitest';
import type { CurrentCharacter } from '../current-character';
import type { LootSource } from './loot';
import { decideToolsBootstrap, settleToolsRestore, sourceIdForInstance } from './tools-bootstrap';

function stored(source: CurrentCharacter['source'], ref: string): CurrentCharacter {
  return {
    source,
    ref,
    label: 'Simfury · Fury Warrior',
    classSlug: 'warrior',
    savedAt: '2026-09-21T00:00:00.000Z',
  };
}

describe('decideToolsBootstrap', () => {
  it('prefers ?code= over everything else', () => {
    const decision = decideToolsBootstrap(
      { code: 'FS1:1:warrior:orc:0/0/0:', source: 'build', ref: 'b1' },
      stored('fight', 'abcdefabcdef:1'),
    );
    expect(decision).toEqual({ kind: 'code', code: 'FS1:1:warrior:orc:0/0/0:', restored: false });
  });

  it('loads an addon-sourced ?source=&ref= as an addon paste', () => {
    const decision = decideToolsBootstrap(
      { code: '', source: 'addon', ref: 'FS1:1:warrior:orc:0/0/0:' },
      null,
    );
    expect(decision).toEqual({ kind: 'addon', code: 'FS1:1:warrior:orc:0/0/0:', restored: false });
  });

  it('loads a build-sourced ?source=&ref=', () => {
    const decision = decideToolsBootstrap({ code: '', source: 'build', ref: 'b1' }, null);
    expect(decision).toEqual({ kind: 'build', id: 'b1', restored: false });
  });

  it('loads a fight-sourced ?source=&ref=', () => {
    const decision = decideToolsBootstrap({ code: '', source: 'fight', ref: 'abcdefabcdef:1' }, null);
    expect(decision).toEqual({ kind: 'fight', ref: 'abcdefabcdef:1', restored: false });
  });

  it('loads an armory-sourced ?source=&ref= as a stored CharacterPath', () => {
    const decision = decideToolsBootstrap({ code: '', source: 'armory', ref: 'us/normal/simfury' }, null);
    expect(decision).toEqual({
      kind: 'stored',
      path: { region: 'us', ruleset: 'normal', slug: 'simfury' },
      restored: false,
    });
  });

  it('does not fall through to the stored pointer when ?source=armory&ref= does not parse', () => {
    const decision = decideToolsBootstrap(
      { code: '', source: 'armory', ref: 'not-a-key' },
      stored('build', 'b1'),
    );
    expect(decision).toEqual({ kind: 'none', restored: false });
  });

  it('bootstraps nothing for a ?source=manual link, the same as store.svelte.ts’s own bootstrapSource', () => {
    const decision = decideToolsBootstrap({ code: '', source: 'manual', ref: 'anything' }, null);
    expect(decision).toEqual({ kind: 'none', restored: false });
  });

  it('falls back to the stored pointer when the URL carries neither ?code= nor ?source=&ref=', () => {
    const decision = decideToolsBootstrap({ code: '', source: '', ref: '' }, stored('build', 'b1'));
    expect(decision).toEqual({ kind: 'build', id: 'b1', restored: true });
  });

  it('restores an addon-sourced pointer through loadCode, same as a fresh addon paste’s own FS1 string', () => {
    const decision = decideToolsBootstrap(
      { code: '', source: '', ref: '' },
      stored('addon', 'FS1:1:warrior:orc:0/0/0:'),
    );
    expect(decision).toEqual({ kind: 'code', code: 'FS1:1:warrior:orc:0/0/0:', restored: true });
  });

  it('restores a code-sourced pointer through loadCode too', () => {
    const decision = decideToolsBootstrap(
      { code: '', source: '', ref: '' },
      stored('code', 'FS1:1:warrior:orc:0/0/0:'),
    );
    expect(decision).toEqual({ kind: 'code', code: 'FS1:1:warrior:orc:0/0/0:', restored: true });
  });

  it('restores a fight-sourced pointer', () => {
    const decision = decideToolsBootstrap(
      { code: '', source: '', ref: '' },
      stored('fight', 'abcdefabcdef:1'),
    );
    expect(decision).toEqual({ kind: 'fight', ref: 'abcdefabcdef:1', restored: true });
  });

  it('restores an armory-sourced pointer through the parsed CharacterPath', () => {
    const decision = decideToolsBootstrap(
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
    const decision = decideToolsBootstrap({ code: '', source: '', ref: '' }, stored('armory', 'not-a-key'));
    expect(decision).toEqual({ kind: 'none', restored: false });
  });

  it('is none when there is no URL bootstrap and no stored pointer', () => {
    expect(decideToolsBootstrap({ code: '', source: '', ref: '' }, null)).toEqual({
      kind: 'none',
      restored: false,
    });
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

describe('settleToolsRestore', () => {
  it('is inert for a URL-driven load, whether it succeeded or failed', () => {
    expect(settleToolsRestore(false, true)).toEqual({
      clearPointer: false,
      restored: false,
      clearMessage: false,
    });
    expect(settleToolsRestore(false, false)).toEqual({
      clearPointer: false,
      restored: false,
      clearMessage: false,
    });
  });

  it('claims restored only once a restore actually produced a character', () => {
    expect(settleToolsRestore(true, true)).toEqual({
      clearPointer: false,
      restored: true,
      clearMessage: false,
    });
  });

  it('forgets a dead or stale pointer and clears the store’s error, rather than showing "Restored" beside it', () => {
    expect(settleToolsRestore(true, false)).toEqual({
      clearPointer: true,
      restored: false,
      clearMessage: true,
    });
  });
});
