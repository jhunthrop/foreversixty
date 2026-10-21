// web/src/lib/planner/send-rule.test.ts
// W.12 ruling 1: the planner never sends another build to a character. The button is on only for
// the character's own build; another class's build leaves it off and says why; another player's
// build of the same class says what it replaces and asks once.
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../account/api';
import { plannerCharacterCardCopy as copy, plannerCopy } from './copy';
import { sendActionFor } from './send-state';

const BOW = {
  key: 'us/pvp/bow-jackzon',
  region: 'us',
  ruleset: 'pvp',
  name: 'Bow Jackzon',
  class: 'hunter',
  build: { source: 'addon', captured_at: 'x', talents: { trees: ['1', '1', '1'], points: [31, 16, 4] } },
} as MeCharacter;

const PALADIN = { classSlug: 'paladin', className: 'Paladin' };
const HUNTER = { classSlug: 'hunter', className: 'Hunter' };

describe('sendActionFor: whose build is on screen', () => {
  it('another class’s build: "This is a Paladin build; Bow Jackzon is a Hunter", disabled', () => {
    const action = sendActionFor(BOW, { kind: 'idle' }, 26, 51, PALADIN);
    expect(action).toMatchObject({
      kind: 'picker',
      label: 'This is a Paladin build; Bow Jackzon is a Hunter',
      disabled: true,
    });
  });

  it('another class’s build stays off with nothing spent, as under a class tile', () => {
    const action = sendActionFor(BOW, { kind: 'idle' }, 0, 51, { classSlug: 'mage', className: 'Mage' });
    expect(action).toMatchObject({
      label: 'This is a Mage build; Bow Jackzon is a Hunter',
      disabled: true,
    });
  });

  it('another class’s build stays off whoever it belongs to', () => {
    const action = sendActionFor(BOW, { kind: 'idle' }, 26, 51, { ...PALADIN, owner: 'Ashenveil' });
    expect(action).toMatchObject({
      label: 'This is a Paladin build; Bow Jackzon is a Hunter',
      disabled: true,
    });
  });

  it('another player’s build of the same class says what it replaces, and is live', () => {
    const action = sendActionFor(BOW, { kind: 'idle' }, 51, 51, { ...HUNTER, owner: 'Stranger' });
    expect(action.label).toBe('Send this build to Bow Jackzon (replaces their 51 points)');
    expect('disabled' in action && action.disabled).toBeFalsy();
  });

  it('a character with no points on file has nothing to replace', () => {
    const action = sendActionFor(BOW, { kind: 'idle' }, 51, null, { ...HUNTER, owner: '' });
    expect(action.label).toBe('Send this build to Bow Jackzon');
  });

  it('asks once: "Send Ashenveil’s build to Bow Jackzon?", with a yes and a cancel', () => {
    const action = sendActionFor(BOW, { kind: 'confirming' }, 51, 51, { ...HUNTER, owner: 'Ashenveil' });
    expect(action).toMatchObject({
      label: 'Yes, send it',
      secondary: { label: 'Cancel' },
      note: { text: "Send Ashenveil's build to Bow Jackzon?" },
    });
    expect(copy.confirmQuestion('', 'Bow Jackzon')).toBe('Send this build to Bow Jackzon?');
  });

  it('the character’s own build, same class and no owner, sends as before', () => {
    const action = sendActionFor(BOW, { kind: 'idle' }, 51, 51, HUNTER);
    expect(action.label).toBe('Send this build to Bow Jackzon');
    expect('disabled' in action && action.disabled).toBeFalsy();
  });

  it('carries "Sim this build" under the send, in every state', () => {
    for (const state of [{ kind: 'idle' }, { kind: 'sending' }, { kind: 'sent' }] as const) {
      const action = sendActionFor(BOW, state, 51, 51, { ...HUNTER, simHref: '/sim?code=x' });
      expect(action).toMatchObject({ secondary: { label: 'Sim this build', href: '/sim?code=x' } });
    }
  });

  it('says whose build it is in the planner header line', () => {
    expect(copy.originLine('Ashenveil', 'Holy Paladin')).toBe("Ashenveil's build · Holy Paladin");
    expect(copy.originLine('', 'Paladin')).toBe('A linked build · Paladin');
  });
});

describe('the planner’s hand-off words', () => {
  it('a nameless link says it carries no talent ranks', () => {
    expect(plannerCopy.noTalentRanksNote).toBe('This link carries no talent ranks.');
  });

  it('a record that reaches further than a legal order names the count it kept', () => {
    expect(plannerCopy.keptNote(25, 51, 'Bow Jackzon')).toContain("Loaded 25 of Bow Jackzon's 51 points");
  });

  it('the short-build note says both counts and no promise about the addon', () => {
    expect(copy.shorterThanSpend(25, 'Bow Jackzon', 51)).toBe(
      'This build has 25 points; Bow Jackzon has 51.',
    );
  });
});
