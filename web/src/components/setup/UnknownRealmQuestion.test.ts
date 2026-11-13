// The unknown-realm question's markup. The vitest config renders components on Svelte's server
// entry, so the tap and its round trip are proved by the API's own tests and the captures.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../../lib/account/me-types';
import { awaitingRealm, unknownRealmCopy } from '../../lib/account/unknown-realm';
import { rulesetLabel } from '../../lib/characters';
import UnknownRealmQuestion from './UnknownRealmQuestion.svelte';

const waiting: MeCharacter = { key: 'us/unknown/newly-rolled', region: 'us', ruleset: 'unknown', name: 'Newly Rolled', class: 'rogue' };
const known: MeCharacter = { key: 'us/pvp/elyra-duskvale', region: 'us', ruleset: 'pvp', name: 'Elyra Duskvale', class: 'mage' };

function body(characters: readonly MeCharacter[]): string {
  return render(UnknownRealmQuestion, { props: { characters, answer: async (key: string) => key } }).body;
}

describe('UnknownRealmQuestion', () => {
  it('asks about each waiting character by name, with the four rulesets as buttons', () => {
    const html = body([waiting]);
    expect(html).toContain(unknownRealmCopy.questionLead);
    expect(html).toContain('Newly Rolled');
    expect(html).toContain(unknownRealmCopy.questionTail);
    for (const [id, label] of [
      ['normal', 'Normal'],
      ['pvp', 'PvP'],
      ['rp', 'Roleplay'],
      ['hardcore', 'Hardcore'],
    ]) {
      expect(html).toMatch(new RegExp(`<button[^>]*data-testid="unknown-realm-choice-${id}"[^>]*>[\\s\\S]*?${label}`));
    }
    expect(html).toContain(unknownRealmCopy.rpPvp);
    expect(html).toMatch(/role="group"[^>]*aria-labelledby="unknown-realm-us-unknown-newly-rolled"/);
  });

  it('draws one block per waiting character and nothing when none waits', () => {
    expect(body([waiting, { ...waiting, key: 'us/unknown/second-roll', name: 'Second Roll' }]).match(/data-testid="unknown-realm-item"/g)).toHaveLength(2);
    expect(body([])).not.toContain('unknown-realm');
  });

  it('asks only about characters on the unknown ruleset', () => {
    expect(awaitingRealm({ characters: [known, waiting] })).toEqual([waiting]);
    expect(awaitingRealm(null)).toEqual([]);
    expect(rulesetLabel('unknown')).toBe('Realm type?');
  });
});
