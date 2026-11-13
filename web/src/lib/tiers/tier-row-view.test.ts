// web/src/lib/tiers/tier-row-view.test.ts
import { describe, expect, it } from 'vitest';
import { dpsInput, healerInput, tankInput } from './tier-test-support';
import { rankRole } from './tier-list';
import { barFraction, rankRowProps } from './tier-row-view';

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
  it('names a row by its full spec name, keys it, and makes the name its one link, to BiS', () => {
    const arms = rankRowProps(dps[1]!, dps);
    expect(arms).toMatchObject({
      rank: '2',
      name: 'Arms Warrior',
      href: '/bis/warrior/arms/60',
      barTone: 'accent',
      key: 'warrior/arms',
      testid: 'tier-row-warrior-arms',
      gap: { text: '−0.8%', label: 'Gap to the top', leader: false },
    });
    expect('sub' in arms || 'links' in arms || 'tie' in arms).toBe(false);
    // The column head names the unit.
    expect(arms.figures).toEqual([
      { value: '942.2', label: 'Damage per second', testid: 'tier-number' },
      // Below 768 the gap column goes: the vs #1 rides under the figure.
      { value: '−0.8%', label: 'Gap to the top', bare: true, phoneOnly: true, testid: 'tier-gap-phone' },
    ]);
  });

  it("links a DPS row's figure to the ranking sim run behind it, keeping the name on BiS", () => {
    const fire = rankRowProps(dps[2]!, dps);
    expect(fire.href).toBe('/bis/mage/fire/60');
    expect(fire.figures[0]).toEqual({
      value: '800.0',
      label: 'Damage per second',
      testid: 'tier-number',
      href: '/sim?code=x&run=1',
      title: 'Open in the sim: Fire Mage',
      ariaLabel: '800.0 DPS, open Fire Mage in the sim',
    });
    // No sim run behind a tank's number: its figure is plain text.
    expect(rankRowProps(tanks[0]!, tanks).figures[0]!.href).toBeUndefined();
  });

  it("reads the leader's gap as best: the column reads against it", () => {
    expect(rankRowProps(dps[0]!, dps).gap).toMatchObject({ text: 'best', leader: true });
    expect(rankRowProps(tanks[0]!, tanks).gap).toMatchObject({ text: 'best', leader: true });
  });

  it("marks a tank's best of each column, with the reason in a title", () => {
    const [warrior, paladin] = tanks.map((row) => rankRowProps(row, tanks));
    expect(warrior!.name).toBe('Protection Warrior');
    expect(warrior!.figures.map((f) => [f.value, f.best === true])).toEqual([
      ['245', true],
      ['40,000', false],
      ['1,633', true],
    ]);
    expect(warrior!.figures[0]!.bestTitle).toBe('Best of the 2 tanks');
    expect(paladin!.figures[1]).toMatchObject({ value: '49,794', best: true });
    expect(paladin!.gap.text).toBe('+22.4%');
  });
});

describe('a tied healer', () => {
  it('names its partner on the line under the name, so the shared rank explains itself', () => {
    const bounds = (hps: number) => [
      { scale: 1.5, hps, hps_error: 1 },
      { scale: 1.65, hps, hps_error: 1 },
    ];
    const healers = rankRole(
      [
        healerInput('druid', 'restoration', 'Restoration', 700, 1, bounds(700)),
        healerInput('priest', 'holy', 'Holy', 600, 1, bounds(600)),
        healerInput('shaman', 'restoration', 'Restoration', 599.5, 1, bounds(599.5)),
      ],
      'healer',
    );
    const holy = rankRowProps(healers[1]!, healers);
    expect(holy.rank).toBe('=2');
    expect(holy.note).toMatchObject({ text: '≈ tie with Restoration Shaman' });
    expect(rankRowProps(healers[0]!, healers).note).toBeUndefined();
  });
});

describe('barFraction', () => {
  it('starts the bar at 90% of the lowest spec and fills it at the top', () => {
    const rows = [{ fraction: 1 }, { fraction: 0.9 }, { fraction: 0.8 }];
    expect(barFraction(rows[0]!, rows)).toBe(1);
    expect(barFraction(rows[2]!, rows)).toBeCloseTo((0.8 - 0.72) / 0.28);
    expect(barFraction(rows[1]!, rows)).toBeGreaterThan(barFraction(rows[2]!, rows));
  });
});
