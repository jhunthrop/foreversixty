// web/src/lib/guild/officer-tools.test.ts
import { describe, expect, it } from 'vitest';
import type { GuildViewerRole } from './api';
import { officerToolsOn, raidGuildToolsOn, refusalGate } from './officer-tools';

const FIELDS: (boolean | undefined)[] = [true, false, undefined];

describe('officerToolsOn', () => {
  it.each(FIELDS)('keeps a member, public viewer and moderator on whatever the field says (%s)', (field) => {
    for (const role of ['public', 'member', 'moderator'] as GuildViewerRole[]) {
      expect(officerToolsOn(role, field, false)).toBe(true);
      expect(officerToolsOn(role, field, true)).toBe(true);
    }
  });

  it('turns an officer read-only when the API says the tools are off', () => {
    expect(officerToolsOn('officer', false, false)).toBe(false);
  });

  it('treats a missing field as on for an officer', () => {
    expect(officerToolsOn('officer', undefined, false)).toBe(true);
    expect(officerToolsOn('officer', true, false)).toBe(true);
  });

  it('turns an officer read-only once a refusal has ended the tools', () => {
    for (const field of FIELDS) expect(officerToolsOn('officer', field, true)).toBe(false);
  });
});

describe('raidGuildToolsOn', () => {
  it.each([true, false])('keeps a member, public viewer and moderator on (%s)', (field) => {
    for (const role of ['public', 'member', 'moderator'] as GuildViewerRole[]) {
      expect(raidGuildToolsOn(role, field, false)).toBe(true);
      expect(raidGuildToolsOn(role, field, true)).toBe(true);
    }
  });

  it('gates an officer when the API says Raid guild is off', () => {
    expect(raidGuildToolsOn('officer', false, false)).toBe(false);
    expect(raidGuildToolsOn('officer', true, false)).toBe(true);
  });

  it('gates an officer once a raid_guild_required refusal has ended it', () => {
    expect(raidGuildToolsOn('officer', true, true)).toBe(false);
  });
});

describe('refusalGate', () => {
  it('maps each Award/Undo refusal code to its gate', () => {
    expect(refusalGate('guild_plan_required')).toBe('plan');
    expect(refusalGate('raid_guild_required')).toBe('raid-guild');
  });

  it('leaves any other code unhandled', () => {
    expect(refusalGate('forbidden')).toBeNull();
    expect(refusalGate('')).toBeNull();
  });
});
