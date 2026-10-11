// web/src/lib/character-header/character-header.test.ts
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../account/api';
import { planAnswer } from './answer-plan';
import { descriptorLines } from './descriptor';
import { hasGearOnFile, syncText, syncView } from './sync';

const NOW = new Date('2026-10-11T12:00:00Z');
const daysAgo = (days: number): string => new Date(NOW.getTime() - days * 86_400_000).toISOString();

function character(overrides: Partial<MeCharacter> = {}): MeCharacter {
  return {
    key: 'us/normal/tester',
    region: 'us',
    ruleset: 'normal',
    name: 'Tester',
    class: 'Warrior',
    spec: 'Fury',
    level: 60,
    faction: 'alliance',
    build: { source: 'addon', captured_at: daysAgo(0.01), gear: { head: 1 } },
    ...overrides,
  };
}

describe('descriptorLines', () => {
  it('reads race, spec and class on the first line and level and faction on the second', () => {
    expect(descriptorLines(character({ race: 'Gnome' }))).toEqual({
      identity: 'Gnome Fury Warrior',
      level: 'Level 60',
      faction: 'alliance',
    });
  });

  it('leaves the race out when the API has not sent it, never a placeholder', () => {
    expect(descriptorLines(character()).identity).toBe('Fury Warrior');
  });

  it('is the bare class when the spec is not known', () => {
    expect(descriptorLines(character({ spec: undefined })).identity).toBe('Warrior');
  });

  it('accepts a class slug as well as a display name', () => {
    expect(descriptorLines(character({ class: 'warrior' })).identity).toBe('Fury Warrior');
  });

  it('omits the level and the faction it does not know', () => {
    const lines = descriptorLines(character({ level: undefined, faction: undefined }));
    expect(lines.level).toBeNull();
    expect(lines.faction).toBeNull();
  });
});

describe('syncView', () => {
  it('names the source and the age', () => {
    expect(syncView(character(), NOW)).toEqual({ source: 'Addon', age: '14 minutes ago', staleTail: null });
  });

  it('puts the stale words after the age once the addon sync is past its limit', () => {
    const stale = character({ build: { source: 'addon', captured_at: daysAgo(19), gear: { head: 1 } } });
    expect(syncView(stale, NOW)).toEqual({
      source: 'Addon',
      age: '19 days ago',
      staleTail: 'Log out in the game to sync',
    });
  });

  it('never asks a Battle.net sync to log out in the game', () => {
    const bnet = character({ build: { source: 'blizzard', captured_at: daysAgo(19), gear: { head: 1 } } });
    expect(syncView(bnet, NOW)?.staleTail).toBeNull();
  });

  it('goes away when no gear is on file', () => {
    expect(syncView(character({ build: { source: 'addon', captured_at: daysAgo(1) } }), NOW)).toBeNull();
    expect(syncView(character({ build: undefined }), NOW)).toBeNull();
  });
});

describe('syncText', () => {
  it('joins the source and the age with a middle dot', () => {
    expect(syncText({ source: 'Addon', age: '12 minutes ago' })).toBe('Addon \u00b7 12 minutes ago');
  });

  it('is the age alone when the source is not known', () => {
    expect(syncText({ source: '', age: '12 minutes ago' })).toBe('12 minutes ago');
  });
});

describe('stale rule', () => {
  it('uses the player’s own median gap, never under the 3-day floor', () => {
    const build = {
      source: 'addon' as const,
      captured_at: daysAgo(4),
      median_sync_gap_sec: 43_200,
      gear: { head: 1 },
    };
    expect(syncView(character({ build }), NOW)?.staleTail).not.toBeNull();
    expect(syncView(character({ build: { ...build, captured_at: daysAgo(2) } }), NOW)?.staleTail).toBeNull();
  });
});

describe('hasGearOnFile', () => {
  it('treats an empty export as no gear, not as zero worn', () => {
    expect(hasGearOnFile(character({ build: { source: 'addon', captured_at: daysAgo(1), gear: {} } }))).toBe(
      false,
    );
    expect(hasGearOnFile(character())).toBe(true);
  });
});

describe('planAnswer', () => {
  const bis = { levelFloor: 20, needsGear: true };

  it('loads the answer for a character at or above the floor with gear', () => {
    expect(planAnswer(character(), bis)).toBe('load');
    expect(planAnswer(character({ level: 20 }), bis)).toBe('load');
  });

  it('shows the floor sentence below the floor, even with no gear', () => {
    expect(planAnswer(character({ level: 12 }), bis)).toBe('below-floor');
    expect(planAnswer(character({ level: 12, build: undefined }), bis)).toBe('below-floor');
  });

  it('shows the no-gear sentence on a page that reads gear', () => {
    expect(planAnswer(character({ build: undefined }), bis)).toBe('no-gear');
  });

  it('changes nothing for a page with no floor and no gear need', () => {
    const logs = {};
    expect(planAnswer(character({ level: 5, build: undefined }), logs)).toBe('load');
  });

  it('does not guess a floor for an unknown level', () => {
    expect(planAnswer(character({ level: undefined }), bis)).toBe('load');
  });
});
