// web/src/lib/guild/signup-strip.test.ts
import { describe, expect, it } from 'vitest';
import type { GuildRaidViewerCharacter } from './api';
import { raidCopy } from './raid-copy';
import { NOTE_MAX, failureLine, initialRole, noteCounter, orderedCharacters } from './signup-strip';

function character(name: string, extra: Partial<GuildRaidViewerCharacter> = {}): GuildRaidViewerCharacter {
  return {
    character_key: name.toLowerCase(),
    name,
    class: 'warrior',
    spec: 'Fury',
    verified: true,
    state: null,
    role: 'dps',
    ...extra,
  };
}

describe('orderedCharacters', () => {
  it('puts the current character first and keeps the rest in their order', () => {
    const list = [character('Ana'), character('Bo'), character('Cy')];
    expect(orderedCharacters(list, 'cy').map((c) => c.name)).toEqual(['Cy', 'Ana', 'Bo']);
    expect(orderedCharacters(list, null).map((c) => c.name)).toEqual(['Ana', 'Bo', 'Cy']);
    expect(orderedCharacters(list, 'nobody').map((c) => c.name)).toEqual(['Ana', 'Bo', 'Cy']);
  });

  it('never mutates the input', () => {
    const list = [character('Ana'), character('Bo')];
    orderedCharacters(list, 'bo');
    expect(list.map((c) => c.name)).toEqual(['Ana', 'Bo']);
  });
});

describe('initialRole', () => {
  it('keeps a signed-up role and prefills an unsigned one from the spec', () => {
    expect(initialRole(character('Ana', { state: 'in', role: 'tank', spec: 'Fury' }))).toBe('tank');
    expect(initialRole(character('Bo', { spec: 'Holy', role: 'dps' }))).toBe('healer');
    expect(initialRole(character('Cy', { spec: 'Unknown', role: 'tank' }))).toBe('tank');
  });
});

describe('noteCounter', () => {
  it('appears from 120 characters and counts down to the 140 cap', () => {
    expect(noteCounter('x'.repeat(119))).toBeNull();
    expect(noteCounter('x'.repeat(120))).toBe(raidCopy.stripStatus.left(NOTE_MAX - 120));
    expect(noteCounter('x'.repeat(NOTE_MAX))).toBe('0 left');
  });
});

describe('failureLine', () => {
  it('names the races, and anything else is the generic failure', () => {
    expect(failureLine('raid_locked')).toEqual({ message: raidCopy.races.locked, race: true });
    expect(failureLine('raid_cancelled')).toEqual({ message: raidCopy.races.cancelled, race: true });
    expect(failureLine('invalid')).toEqual({ message: raidCopy.stripStatus.failed, race: false });
  });
});
