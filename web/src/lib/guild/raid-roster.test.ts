// web/src/lib/guild/raid-roster.test.ts
// The Sign-ups roster's groups (raids spec 3.7, raid comp spec 3.3): In rows by the API's group,
// then Late, Bench, Out and Not signed up; class in the game's order, then name.
import { describe, expect, it } from 'vitest';
import { plannedRaid, viewerCharacter } from '../../test-support/raid-fixtures';
import type { GuildRaidNotSigned, GuildRaidSignup } from './api';
import { alsoLine, officersSeeLine, rosterGroups } from './raid-roster';

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
  it('splits In rows by the API group and sorts by class order, then name', () => {
    const groups = rosterGroups(
      plannedRaid({
        signups: [
          signup('Zed', { class: 'druid', spec: 'Balance', group: 'ranged' }),
          signup('Thorn', { role: 'tank', spec: 'Protection', group: 'tank' }),
          signup('Brokk', { role: 'tank', spec: 'Fury', group: 'tank' }),
          signup('Abe', { class: 'rogue', spec: 'Combat', group: 'melee' }),
          signup('Bolt', { group: 'melee' }),
          signup('Mend', { role: 'healer', class: 'priest', spec: 'Holy', group: 'healer' }),
          signup('Aqua', { role: 'healer', class: 'paladin', spec: 'Holy', group: 'healer' }),
          signup('Able', { class: 'rogue', spec: 'Combat', group: 'melee' }),
          signup('Kelthar', { class: 'hunter', spec: 'Survival', group: 'melee' }),
          signup('Zulmara', { class: 'hunter', spec: 'Marksmanship', group: 'ranged' }),
        ],
      }),
    );
    expect(groups.groups.tank.map((r) => r.name)).toEqual(['Brokk', 'Thorn']);
    expect(groups.groups.healer.map((r) => r.name)).toEqual(['Aqua', 'Mend']);
    expect(groups.groups.melee.map((r) => r.name)).toEqual(['Bolt', 'Kelthar', 'Abe', 'Able']);
    expect(groups.groups.ranged.map((r) => r.name)).toEqual(['Zulmara', 'Zed']);
  });

  it('reads the group the API sent, never a class or spec rule (a DPS row with none is ranged)', () => {
    const groups = rosterGroups(plannedRaid({ signups: [signup('Corvane', { class: 'rogue', spec: '' })] }));
    expect(groups.groups.ranged.map((r) => r.name)).toEqual(['Corvane']);
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
    expect(groups.groups.tank).toEqual([]);
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

describe('one row per account (raids spec 3.7)', () => {
  const ROOTBOUND = {
    character_key: 'us/pvp/rootbound',
    name: 'Rootbound',
    class: 'druid',
    spec: 'Restoration',
    spec_key: 'druid-restoration',
    state: 'bench' as const,
  };
  const QUILL = signup('Quillfeather', {
    class: 'priest',
    spec: 'Holy',
    spec_key: 'priest-holy',
    role: 'healer',
    state: 'late',
    own: true,
    also: [ROOTBOUND],
  });

  it('names the account’s other answered characters under its row', () => {
    expect(alsoLine([ROOTBOUND])).toBe('Also: Rootbound (Restoration Druid, Bench)');
    expect(
      alsoLine([
        ROOTBOUND,
        { ...ROOTBOUND, name: 'Brew', class: 'mage', spec_key: 'mage-fire', state: 'out' },
      ]),
    ).toBe('Also: Rootbound (Restoration Druid, Bench); Brew (Fire Mage, Out)');
    expect(alsoLine([])).toBeNull();
    const groups = rosterGroups(plannedRaid({ signups: [QUILL, signup('Bolt')], not_signed: [] }));
    expect(groups.late.map((row) => [row.name, row.also])).toEqual([
      ['Quillfeather', 'Also: Rootbound (Restoration Druid, Bench)'],
    ]);
    expect(groups.bench).toEqual([]);
    expect(groups.groups.ranged.map((row) => row.also)).toEqual([null]);
  });

  it('tells the member which character the officers see, on the other one only', () => {
    const raid = plannedRaid({ signups: [QUILL] });
    expect(officersSeeLine(raid, 'us/pvp/rootbound')).toBe('Officers see Quillfeather (Late).');
    expect(officersSeeLine(raid, 'us/pvp/quillfeather')).toBeNull();
    expect(officersSeeLine(raid, 'us/pvp/nobody')).toBeNull();
  });
});
