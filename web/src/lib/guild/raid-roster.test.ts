// web/src/lib/guild/raid-roster.test.ts
// The Sign-ups roster's groups (raids spec 3.7): In rows by role, then Late, Bench, Out and Not
// signed up; role, then class in the game's order, then name.
import { describe, expect, it } from 'vitest';
import { plannedRaid, viewerCharacter } from '../../test-support/raid-fixtures';
import type { GuildRaidNotSigned, GuildRaidSignup } from './api';
import { rosterGroups } from './raid-roster';

function signup(name: string, extra: Partial<GuildRaidSignup> = {}): GuildRaidSignup {
  return {
    character_key: `us/pvp/${name.toLowerCase()}`,
    name,
    class: 'warrior',
    spec: 'Fury',
    role: 'dps',
    state: 'in',
    note: '',
    own: false,
    ...extra,
  };
}

function unsigned(name: string, extra: Partial<GuildRaidNotSigned> = {}): GuildRaidNotSigned {
  return { character_key: `us/pvp/${name.toLowerCase()}`, name, class: 'mage', spec: 'Frost', ...extra };
}

describe('rosterGroups', () => {
  it('splits In rows by role and sorts by class order, then name', () => {
    const groups = rosterGroups(
      plannedRaid({
        signups: [
          signup('Zed', { class: 'druid', spec: 'Balance' }),
          signup('Thorn', { role: 'tank', spec: 'Protection' }),
          signup('Abe', { class: 'rogue', spec: 'Combat' }),
          signup('Bolt'),
          signup('Mend', { role: 'healer', class: 'priest', spec: 'Holy' }),
          signup('Aqua', { role: 'healer', class: 'paladin', spec: 'Holy' }),
          signup('Able', { class: 'rogue', spec: 'Combat' }),
        ],
      }),
    );
    expect(groups.roles.tank.map((r) => r.name)).toEqual(['Thorn']);
    expect(groups.roles.healer.map((r) => r.name)).toEqual(['Aqua', 'Mend']);
    expect(groups.roles.dps.map((r) => r.name)).toEqual(['Bolt', 'Abe', 'Able', 'Zed']);
  });

  it('puts Late, Bench and Out in their own groups, Late in role order', () => {
    const groups = rosterGroups(
      plannedRaid({
        signups: [
          signup('Laggard', { state: 'late' }),
          signup('Tardy', { state: 'late', role: 'tank' }),
          signup('Benched', { state: 'bench' }),
          signup('Gone', { state: 'out' }),
        ],
      }),
    );
    expect(groups.late.map((r) => r.name)).toEqual(['Tardy', 'Laggard']);
    expect(groups.bench.map((r) => r.name)).toEqual(['Benched']);
    expect(groups.out.map((r) => r.name)).toEqual(['Gone']);
    expect(groups.roles.tank).toEqual([]);
  });

  it('lists the not-signed members with their spec role and whether they are the viewer', () => {
    const groups = rosterGroups(
      plannedRaid({
        not_signed: [unsigned('Mira'), unsigned('Bear', { class: 'druid', spec: 'Feral Bear' })],
        viewer: { characters: [viewerCharacter({ character_key: 'us/pvp/mira', state: null })] },
      }),
    );
    expect(groups.notSigned.map((r) => [r.name, r.role, r.state, r.own])).toEqual([
      ['Bear', 'tank', null, false],
      ['Mira', 'dps', null, true],
    ]);
  });
});
