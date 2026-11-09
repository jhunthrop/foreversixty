// web/src/lib/tiers/tier-row-view.test.ts
import { describe, expect, it } from 'vitest';
import { dpsInput, tankInput } from './tier-test-support';
import { rankRole } from './tier-list';
import { rankRowProps } from './tier-row-view';

const dps = rankRole(
  [
    dpsInput('hunter', 'survival', 'Survival', 950.2),
    dpsInput('warrior', 'arms', 'Arms', 942.2),
    { ...dpsInput('mage', 'fire', 'Fire', 800), simHref: '/sim?code=x&run=1' },
  ],
  'dps',
);
const tanks = rankRole(
  [
    tankInput('warrior', 'protection', 'Protection', { dtps: 245, effective_health: 40000, tps: 1633 }),
    tankInput('paladin', 'protection', 'Protection', { dtps: 300, effective_health: 49794, tps: 1200 }),
  ],
  'tank',
);

describe('rankRowProps', () => {
  it('names a DPS row by spec, keys it, and links its name to BiS', () => {
    const arms = rankRowProps(dps[1]!, dps.length);
    expect(arms).toMatchObject({
      rank: '2',
      name: 'Arms',
      sub: 'Warrior · Night Elf',
      href: '/bis/warrior/arms/60',
      key: 'warrior/arms',
      testid: 'tier-row-warrior-arms',
      gap: { text: '−0.8%', label: 'Gap to the top', leader: false },
      tie: { text: '≈ tie with Survival', title: 'Within 1% of Survival Hunter: a tie' },
    });
    expect(arms.figures).toEqual([
      { value: '942.2', unit: 'DPS', label: 'Damage per second', testid: 'tier-number' },
    ]);
  });

  it('gives a DPS row BiS, Guide and Planner, and the Sim only where the number can be rerun', () => {
    const labels = (index: number) => rankRowProps(dps[index]!, dps.length).links.map((link) => link.label);
    expect(labels(1)).toEqual(['BiS', 'Guide', 'Planner']);
    const fire = rankRowProps(dps[2]!, dps.length).links;
    expect(fire.map((link) => link.label)).toEqual(['BiS', 'Guide', 'Planner', 'Sim']);
    expect(fire[0]).toMatchObject({ primary: true, ariaLabel: 'BiS: Fire Mage' });
    expect(fire[3]).toMatchObject({ href: '/sim?code=x&run=1', ariaLabel: 'View ranking sim: Fire Mage' });
    expect(rankRowProps(dps[0]!, dps.length).gap).toMatchObject({ text: 'Top', leader: true });
  });

  it('names a tank in full and marks the best of each column, with the reason in a title', () => {
    const [warrior, paladin] = tanks.map((row) => rankRowProps(row, tanks.length));
    expect(warrior!.name).toBe('Protection Warrior');
    expect(warrior!.sub).toBe('Night Elf');
    expect(warrior!.figures.map((f) => [f.value, f.unit, f.best === true])).toEqual([
      ['245', 'taken / s', true],
      ['40,000', 'HP', false],
      ['1,633', 'threat / s', true],
    ]);
    expect(warrior!.figures[0]!.bestTitle).toBe('Best of the 2 tanks');
    expect(paladin!.figures[1]).toMatchObject({ value: '49,794', best: true });
    expect(paladin!.gap.text).toBe('+22.4%');
    expect('tie' in paladin!).toBe(false);
  });
});
