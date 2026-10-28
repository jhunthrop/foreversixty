import { describe, expect, it } from 'vitest';
import { playerCardLine } from './player-card';

describe('playerCardLine', () => {
  it('reads spec, class, boss with size and the result', () => {
    const fight = { name: 'Onyxia', size: 40, kill: true, duration_ms: 232_000 };
    expect(playerCardLine('Fury', 'Warrior', fight)).toBe('Fury Warrior · Onyxia (40) · Kill 3:52');
  });
  it('leaves out what the report does not know', () => {
    expect(playerCardLine('', 'Warrior', null)).toBe('Warrior');
    const wipe = { name: 'Onyxia', size: undefined, kill: false, duration_ms: 61_000 };
    expect(playerCardLine('Fury', undefined, wipe)).toBe('Fury · Onyxia · Wipe 1:01');
  });
});
