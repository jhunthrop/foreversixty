import { describe, expect, it } from 'vitest';
import { noTauntsLine, orderedTaunts, pickedEnemyName } from './threat-taunts';
import type { Taunt, ThreatPair } from './types';

const taunt = (at: number, target: string): Taunt => ({
  at_ms: at,
  source_guid: 'Player-1',
  source_name: 'Baelgrim-Realm',
  target_guid: `Creature-${target}`,
  target_name: target,
  spell_id: 355,
  spell_name: 'Taunt',
});

describe('the Threat tab taunts', () => {
  it('names the picked enemy from the pair on its GUID', () => {
    const pairs = [
      { guid: 'Player-1', name: 'B', target_guid: 'C-1', target_name: 'Warden Kelthas', threat: 1 },
    ];
    expect(pickedEnemyName(pairs as ThreatPair[], 'C-1')).toBe('Warden Kelthas');
    expect(pickedEnemyName(pairs as ThreatPair[], '')).toBeUndefined();
    expect(pickedEnemyName(pairs as ThreatPair[], 'C-9')).toBeUndefined();
  });

  it('orders by time and narrows to the picked enemy', () => {
    const list = [taunt(9000, 'Add'), taunt(3000, 'Boss'), taunt(5000, 'Add')];
    expect(orderedTaunts(list, undefined).map((t) => t.at_ms)).toEqual([3000, 5000, 9000]);
    expect(orderedTaunts(list, 'Add').map((t) => t.at_ms)).toEqual([5000, 9000]);
  });

  it('says what emptied the list', () => {
    expect(
      noTauntsLine({ sourceName: undefined, enemy: undefined, windowed: false, scopeNoun: 'pull' }),
    ).toBe('No taunts in this pull.');
    expect(
      noTauntsLine({ sourceName: 'Baelgrim-Realm', enemy: 'Boss', windowed: true, scopeNoun: 'pull' }),
    ).toBe('No taunts by Baelgrim on Boss in this window.');
  });
});
