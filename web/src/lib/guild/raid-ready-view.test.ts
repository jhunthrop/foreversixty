import { describe, expect, it } from 'vitest';
import type { GuildRaidSignup, GuildReadinessPage, GuildReadinessRow } from './api';
import { ofLine, raidReadyView, tankFireLine, tanksMeta, unknownLine } from './raid-ready-view';

const resist = (fire: number) => ({ fire, nature: 0, frost: 0, shadow: 0, arcane: 0 });

function row(key: string, extra: Partial<GuildReadinessRow>): GuildReadinessRow {
  return {
    character_key: key,
    attuned: null,
    resistances: null,
    flask_food: null,
    ...extra,
  } as GuildReadinessRow;
}

function signup(
  key: string,
  role: GuildRaidSignup['role'],
  state: GuildRaidSignup['state'],
): GuildRaidSignup {
  return { character_key: key, name: key, class: 'warrior', spec: '', role, state, note: '', own: false };
}

function page(rows: GuildReadinessRow[], attunement = true, fire = true): GuildReadinessPage {
  return {
    median_item_level: 0,
    generated_at: '',
    rows,
    raid: { id: 1, title: '', starts_at: '', cutoff_at: '', attunement, fire_resistance: fire },
  };
}

const rows = [
  row('a', { attuned: true, flask_food: { flask_or_elixir: true, food: true }, resistances: resist(142) }),
  row('b', { attuned: false, flask_food: { flask_or_elixir: true, food: false } }),
  row('c', { attuned: true }),
  row('d', {}),
];
const signups = [
  signup('a', 'tank', 'in'),
  signup('d', 'tank', 'late'),
  signup('c', 'tank', 'bench'),
  signup('b', 'dps', 'in'),
];

describe('raidReadyView', () => {
  it('counts the attuned and the flask and food carriers among those whose export says', () => {
    const view = raidReadyView(page(rows), signups)!;
    expect(view.attuned).toEqual({ yes: 2, known: 3, unknown: 1 });
    expect(view.flaskFood).toEqual({ yes: 1, known: 2, unknown: 2 });
    expect(ofLine(view.attuned!)).toBe('2 of 3');
    expect(unknownLine(view.attuned!, 'quests')).toBe('1 has not shared their quests');
    expect(unknownLine(view.flaskFood, 'bags')).toBe('2 have not shared their bags');
    expect(unknownLine({ yes: 1, known: 1, unknown: 0 }, 'bags')).toBe('');
  });

  it("lists each In or Late tank's fire resistance, the unshared ones too", () => {
    const tanks = raidReadyView(page(rows), signups)!.tanks!;
    expect(tanks.map((t) => [t.key, t.fire])).toEqual([
      ['a', 142],
      ['d', null],
    ]);
    expect(tanks.map(tankFireLine)).toEqual(['142 fire resistance', 'Not shared yet']);
    expect(tanksMeta(tanks)).toBe('2 tanks signed');
    expect(tanksMeta([])).toBe('No tank signed yet');
  });

  it('asks for attunement and fire resistance only where the raid does', () => {
    const view = raidReadyView(page(rows, false, false), signups)!;
    expect(view.attuned).toBeNull();
    expect(view.tanks).toBeNull();
    expect(view.flaskFood.known).toBe(2);
  });

  it('hides when nothing is tracked or the page is not cut to a raid', () => {
    expect(raidReadyView(page([row('x', {}), row('y', {})]), signups)).toBeNull();
    expect(raidReadyView({ ...page(rows), raid: undefined }, signups)).toBeNull();
    expect(raidReadyView(page([row('x', { attuned: true })], true, false), [])).not.toBeNull();
  });
});
