// web/src/lib/guild/signup-strip.test.ts
import { describe, expect, it } from 'vitest';
import type { GuildRaidViewerCharacter } from './api';
import { raidCopy } from './raid-copy';
import {
  NOTE_MAX,
  defaultPick,
  failureLine,
  noteCounter,
  orderedCharacters,
  pressWaitsForSpec,
  roleOfPick,
} from './signup-strip';

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

describe('defaultPick (raid comp spec 3.2)', () => {
  it('starts on the saved raid spec, else the default spec, else nothing', () => {
    expect(
      defaultPick(character('Ana', { raid_spec: 'warrior-arms', default_spec: 'warrior-arms' }))?.label,
    ).toBe('Arms · DPS');
    expect(defaultPick(character('Bo', { raid_spec: null, default_spec: 'warrior-protection' }))?.label).toBe(
      'Protection · Tank',
    );
    expect(defaultPick(character('Cy', { raid_spec: null, default_spec: null }))).toBeNull();
  });

  it('keeps an off-tank on its tank pair', () => {
    const offTank = character('Brokk', { state: 'in', role: 'tank', raid_spec: 'warrior-fury' });
    expect(defaultPick(offTank)?.label).toBe('Fury · Tank');
    expect(roleOfPick(defaultPick(offTank), offTank)).toBe('tank');
  });

  it('holds In, Late and Bench without a pick, never Out', () => {
    const pick = defaultPick(character('Ana', { raid_spec: 'warrior-arms' }));
    expect(pressWaitsForSpec('in', null)).toBe(true);
    expect(pressWaitsForSpec('bench', null)).toBe(true);
    expect(pressWaitsForSpec('out', null)).toBe(false);
    expect(pressWaitsForSpec('late', pick)).toBe(false);
    expect(roleOfPick(null, character('Cy', { role: 'tank' }))).toBe('tank');
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
    expect(failureLine('raid_ended')).toEqual({ message: 'This raid has ended.', race: true });
    expect(failureLine('invalid')).toEqual({ message: raidCopy.stripStatus.failed, race: false });
  });
});
