import { describe, expect, it } from 'vitest';
import { guildJoinCopy } from './copy';
import { holdsErrorRoom, joiningRow, showsInvitedLine } from './join-view';

describe('showsInvitedLine', () => {
  it('shows the invite only while it can be taken', () => {
    expect(showsInvitedLine(false, false)).toBe(true);
    expect(showsInvitedLine(false, true)).toBe(false);
    expect(showsInvitedLine(true, false)).toBe(false);
    expect(showsInvitedLine(true, true)).toBe(false);
  });
});

describe('holdsErrorRoom', () => {
  it('holds the error line only while joining or after a failure', () => {
    expect(holdsErrorRoom(false, '')).toBe(false);
    expect(holdsErrorRoom(true, '')).toBe(true);
    expect(holdsErrorRoom(false, 'That did not work. Try again.')).toBe(true);
  });
});

const ROGUE = { name: 'Thornhide', class: 'rogue', race: 'Night Elf', level: 52, on_roster: true };

describe('joiningRow', () => {
  it('a character on the roster joins, in class colour, with what is known of them', () => {
    expect(joiningRow(ROGUE, 'Olympus')).toEqual({
      crestClass: 'rogue',
      label: 'You join as Thornhide',
      labelClass: 'rogue',
      sub: 'Level 52 Night Elf Rogue',
    });
    expect(joiningRow({ ...ROGUE, race: null, level: null }, 'Olympus').sub).toBe('Rogue');
    expect(joiningRow({ ...ROGUE, class: null, race: null }, 'Olympus')).toMatchObject({
      crestClass: '',
      sub: 'Level 52',
    });
  });

  it('characters, none on the roster: says who officers see', () => {
    expect(joiningRow({ ...ROGUE, on_roster: false }, 'Olympus')).toEqual({
      crestClass: 'rogue',
      label: 'None of your characters is in Olympus in game yet.',
      labelClass: null,
      sub: 'Officers see Thornhide until one of yours shows on the roster.',
    });
  });

  it('no characters: the account joins, on the neutral ring', () => {
    for (const none of [null, undefined]) {
      expect(joiningRow(none, 'Olympus')).toEqual({
        crestClass: '',
        label: 'You join with your account',
        labelClass: null,
        sub: 'Your characters show once you sync the addon or import from Battle.net.',
      });
    }
  });
});

describe('the join card words', () => {
  it('the guild line counts its members', () => {
    expect(guildJoinCopy.guildLine('PvP US', 1)).toBe('PvP US · 1 member on Forever Sixty');
    expect(guildJoinCopy.guildLine('PvP US', 12)).toBe('PvP US · 12 members on Forever Sixty');
    expect(guildJoinCopy.guildLine('PvP US', undefined)).toBe('PvP US');
  });

  it('the invite line does not name an inviter, and verified is spelled out', () => {
    expect(guildJoinCopy.invited('Olympus')).toBe('You are invited to join Olympus.');
    expect(guildJoinCopy.verifiedAtOnce).toBe(
      'Joining by this link verifies you at once. No officer has to approve you.',
    );
  });
});
