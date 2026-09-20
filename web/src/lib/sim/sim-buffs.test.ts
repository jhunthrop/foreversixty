// web/src/lib/sim/sim-buffs.test.ts
import { describe, expect, it } from 'vitest';
import simbuffsJson from '../../fixtures/planner/simbuffs.json';
import { buffIcon, buffName, type SimBuffFile } from './sim-buffs';

const file = simbuffsJson as unknown as SimBuffFile;

describe('buffName and buffIcon', () => {
  it('name and illustrate an id the file has', () => {
    expect(buffName(file, 'flask_of_supreme_power')).toBe('Flask of Supreme Power');
    expect(buffIcon(file, 'flask_of_supreme_power')).toBe('fixture_item_snakestone_charm');
  });

  it('fall back to a legible form of the id rather than to an empty row', () => {
    expect(buffName(file, 'juju_power')).toBe('juju power');
    expect(buffIcon(file, 'juju_power')).toBe('');
  });

  it('read an empty file without throwing', () => {
    expect(buffName({ entries: {} }, 'battle_shout')).toBe('battle shout');
  });
});
