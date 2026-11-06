// The character page's words and rows (design/specs/2026-11-06-character.md C2 to C11).
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../account/api';
import type { PaperdollSlot } from '../gear/paperdoll';
import type { CharacterPage } from '../rankings/api';
import {
  bestRows,
  buildRows,
  characterNextStep,
  characterPageCopy,
  headFacts,
  pageSpec,
  topUpgrade,
  upgradeTotal,
} from './page-view';

const page = {
  character: {
    name: 'Elyra Duskvale',
    region: 'us',
    ruleset: 'hardcore',
    class: 'Priest',
    guild: { name: 'Cx66 Olympus', rank: 'officer' },
    level: 60,
    race: 'Night Elf',
    realm: 'Doomhowl',
  },
  best: [],
  history: [],
  builds_seen: [{ talent_split: '15/31/5', spec: 'Holy', first_seen: '2026-11-03T20:40:00Z' }],
  gear_spec: 'priest-holy',
} as unknown as CharacterPage;
const path = { region: 'us', ruleset: 'hardcore', slug: 'elyra-duskvale' };

function gain(slot: string, value: number | null): PaperdollSlot {
  return {
    slot,
    state: 'item',
    upgrade: { kind: 'gain', value, unit: 'dps', pick: { id: 1, name: 'Pick', icon: null, quality: 4 } },
  } as PaperdollSlot;
}

describe('headFacts', () => {
  it('reads level and race, spec, the split to the planner, guild, realm and the ruleset', () => {
    const facts = headFacts(page, path, null);
    expect(facts.map((fact) => fact.text)).toEqual([
      'Level 60 Night Elf',
      'Holy Priest',
      '15/31/5',
      '<Cx66 Olympus>',
      'Doomhowl',
      'Hardcore US',
    ]);
    expect(facts.find((fact) => fact.key === 'split')?.href).toBe('/planner');
  });

  it('names the spec from the gear when nothing else does', () => {
    const bare = { ...page, builds_seen: [], gear_spec: 'hunter-beast-mastery' } as CharacterPage;
    expect(pageSpec(bare, null)).toBe('Beast Mastery');
    expect(pageSpec({ ...bare, gear_spec: null } as unknown as CharacterPage, null)).toBeNull();
  });
});

describe('the upgrade total', () => {
  it('sums priced gains and leaves out ties and unpriced slots', () => {
    const slots = [gain('head', 10.25), gain('neck', 0.01), gain('main_hand', null), gain('feet', 4)];
    expect(upgradeTotal(slots)).toBeCloseTo(14.25);
    expect(topUpgrade(slots)?.slot).toBe('head');
    expect(topUpgrade([gain('neck', 0.01)])).toBeNull();
  });
});

describe('characterNextStep', () => {
  const own = {
    key: 'us/normal/felsnap',
    level: 34,
    build: { talents: { points: [23, 0, 0] } },
  } as unknown as MeCharacter;

  it('puts unspent talent points first, then the top upgrade', () => {
    const step = characterNextStep({
      own,
      top: { name: 'Hood', source: 'Wailing Caverns' },
      bisHref: '/bis',
    });
    expect(step.text).toBe('Next: spend your 2 talent points, then get Hood.');
    expect(step.action?.href).toBe('/planner?character=us%2Fnormal%2Ffelsnap');
  });

  it('names the top upgrade when every point is spent', () => {
    const spent = { ...own, build: { talents: { points: [25, 0, 0] } } } as unknown as MeCharacter;
    const step = characterNextStep({
      own: spent,
      top: { name: 'Hood', source: 'Wailing Caverns' },
      bisHref: '/bis',
    });
    expect(step).toEqual({
      text: 'Next: get Hood from Wailing Caverns.',
      action: { label: 'Open best in slot', href: '/bis' },
    });
    expect(characterNextStep({ own: spent, top: null, bisHref: null }).text).toBe('Nothing waiting.');
  });
});

describe('fights and builds', () => {
  it('keeps five best fights, each with its parse colour', () => {
    const row = {
      encounter: 'Lucifron',
      encounter_id: 1,
      difficulty: 9,
      metric: 'hps',
      value: 1012.4,
      percentile: 98.4,
      spec: 'Holy',
      fought_at: '2026-11-03T20:40:00Z',
    };
    const rows = bestRows(
      Array.from({ length: 7 }, (_, index) => ({ ...row, encounter_id: index }) as never),
    );
    expect(rows).toHaveLength(5);
    expect(rows[0]).toMatchObject({
      boss: 'Lucifron',
      sub: '3 Nov · Holy',
      percentile: '98',
      value: '1,012',
    });
    expect(rows[0].color).toBe('var(--color-parse-orange)');
    expect(characterPageCopy.allFights(7)).toBe('All ranked fights (7)');
  });

  it('lists the three newest builds', () => {
    const builds = ['2026-10-01', '2026-11-01', '2026-10-15', '2026-09-01'].map((day, index) => ({
      talent_split: `${index}/0/0`,
      spec: 'Holy',
      first_seen: `${day}T20:00:00Z`,
    }));
    const rows = buildRows('Priest', builds);
    expect(rows.map((row) => row.split)).toEqual(['1/0/0', '2/0/0', '0/0/0']);
    expect(rows[0].sub).toMatch(/first seen 1 Nov 2026$/);
  });
});
