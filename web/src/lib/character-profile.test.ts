// web/src/lib/character-profile.test.ts
import { describe, expect, it } from 'vitest';
import {
  amountHeader,
  bestRepeatsHistory,
  bisHref,
  characterHeaderCopy,
  characterTableCopy,
  newestSpec,
  profileHeader,
  specClassLabel,
} from './character-profile';
import type { CharacterFight, CharacterPage } from './rankings/api';

const fight = (spec?: string): CharacterFight => ({
  encounter: 'Warden Kelthas',
  encounter_id: 9001,
  difficulty: 8,
  metric: 'hps',
  value: 1840,
  spec,
  fought_at: '2026-12-09T22:10:00Z',
  report_id: 'fixture2abcd',
  fight_index: 3,
  execution_score: null,
});

describe('newestSpec', () => {
  it('reads the newest fight’s spec, then the best rows', () => {
    expect(newestSpec({ history: [fight('Holy'), fight('Discipline')], best: [] })).toBe('Holy');
    expect(newestSpec({ history: [fight()], best: [fight(), fight('Shadow')] })).toBe('Shadow');
  });

  it('is null when no fight names one', () => {
    expect(newestSpec({ history: [], best: [] })).toBeNull();
    expect(newestSpec({ history: [fight()], best: [fight()] })).toBeNull();
  });
});

describe('specClassLabel', () => {
  it('reads "Holy Priest" with a spec and "Priest" without', () => {
    expect(specClassLabel('Priest', 'Holy')).toBe('Holy Priest');
    expect(specClassLabel('Priest', null)).toBe('Priest');
    expect(specClassLabel(undefined, 'Holy')).toBe('');
  });
});

describe('bisHref', () => {
  it('links the spec’s best-in-slot page when class and spec are known', () => {
    expect(bisHref('Priest', 'Holy')).toBe('/bis/priest/holy');
    expect(bisHref('Hunter', 'Beast Mastery')).toBe('/bis/hunter/beast-mastery');
    expect(bisHref('Priest', null)).toBeNull();
    expect(bisHref(undefined, 'Holy')).toBeNull();
  });
});

describe('the fight tables (fix round 2 item 6)', () => {
  it('heads the amount column with the metric every row shares, else both', () => {
    expect(amountHeader([{ metric: 'hps' }, { metric: 'hps' }])).toBe('HPS');
    expect(amountHeader([{ metric: 'dps' }])).toBe('DPS');
    expect(amountHeader([{ metric: 'dps' }, { metric: 'hps' }])).toBe('DPS or HPS');
  });

  it('knows when "Best per encounter" only repeats every ranked fight', () => {
    const a = { report_id: 'r1', fight_index: 3 };
    const b = { report_id: 'r2', fight_index: 1 };
    const c = { report_id: 'r3', fight_index: 2 };
    expect(bestRepeatsHistory([a, b], [b, a])).toBe(true);
    expect(bestRepeatsHistory([a], [a, b])).toBe(false);
    expect(bestRepeatsHistory([a, c], [a, b])).toBe(false);
  });

  it('names the columns Boss, Spec, Parse %, Date, Execution', () => {
    expect(Object.values(characterTableCopy)).toEqual(['Boss', 'Spec', 'Parse %', 'Date', 'Execution']);
  });
});

describe('profileHeader (Task L8.4 ruling 1)', () => {
  const PATH = { region: 'us', ruleset: 'hardcore' };
  const page = (
    character: Partial<CharacterPage['character']>,
    spec?: string,
  ): Parameters<typeof profileHeader>[0] => ({
    character: { name: 'Elyra Duskvale', region: 'us', ruleset: 'hardcore', ...character },
    history: spec === undefined ? [] : [fight(spec)],
    best: [],
    builds_seen: [],
  });
  const FULL = {
    class: 'Priest',
    guild: { name: 'Cx66 Olympus', rank: 'officer' as const },
    realm: 'Classic Beta PvP',
    level: 60,
    race: 'Night Elf',
  };
  const text = (header: ReturnType<typeof profileHeader>): string =>
    header.pieces.map((piece) => piece.text).join(characterHeaderCopy.separator);

  it('reads "Holy Priest · <Guild> · Realm", then "Level 60 Night Elf", with "Hardcore US" as the chip', () => {
    const header = profileHeader(page(FULL, 'Holy'), PATH);
    expect(text(header)).toBe('Holy Priest · <Cx66 Olympus> · Classic Beta PvP');
    expect(header.levelRace).toBe('Level 60 Night Elf');
    expect(header.chip).toBe('Hardcore US');
  });

  it('links the guild to its page on the character’s region and ruleset', () => {
    const header = profileHeader(page(FULL, 'Holy'), PATH);
    expect(header.pieces.map((piece) => piece.key)).toEqual(['spec', 'guild', 'realm']);
    expect(header.pieces[1].href).toBe('/guild/us/hardcore/cx66-olympus');
    expect(header.pieces[0].href).toBeUndefined();
    expect(header.pieces[2].href).toBeUndefined();
  });

  it('drops the guild when it is null, never a placeholder', () => {
    const header = profileHeader(page({ ...FULL, guild: null }, 'Holy'), PATH);
    expect(text(header)).toBe('Holy Priest · Classic Beta PvP');
  });

  it('drops the realm when it is null or empty', () => {
    expect(text(profileHeader(page({ ...FULL, realm: null }, 'Holy'), PATH))).toBe(
      'Holy Priest · <Cx66 Olympus>',
    );
    expect(text(profileHeader(page({ ...FULL, realm: '  ' }, 'Holy'), PATH))).toBe(
      'Holy Priest · <Cx66 Olympus>',
    );
  });

  it('reads the class alone without a spec, and drops it with no class', () => {
    expect(text(profileHeader(page(FULL), PATH))).toBe('Priest · <Cx66 Olympus> · Classic Beta PvP');
    expect(text(profileHeader(page({ ...FULL, class: undefined }, 'Holy'), PATH))).toBe(
      '<Cx66 Olympus> · Classic Beta PvP',
    );
  });

  it('has no pieces at all when nothing is known, and the chip still stands', () => {
    const header = profileHeader(page({}), PATH);
    expect(header.pieces).toEqual([]);
    expect(header.levelRace).toBeNull();
    expect(header.chip).toBe('Hardcore US');
  });

  it('writes the second line from whichever of level and race is known', () => {
    expect(profileHeader(page({ ...FULL, race: null }), PATH).levelRace).toBe('Level 60');
    expect(profileHeader(page({ ...FULL, level: null }), PATH).levelRace).toBe('Night Elf');
    expect(profileHeader(page({ ...FULL, level: null, race: null }), PATH).levelRace).toBeNull();
  });

  it('reads a page that predates the identity fields (all of them omitted) as all unknown', () => {
    const header = profileHeader(page({ class: 'Priest' }, 'Holy'), PATH);
    expect(text(header)).toBe('Holy Priest');
    expect(header.levelRace).toBeNull();
  });

  it('echoes a ruleset the site has not heard of in the chip', () => {
    expect(profileHeader(page(FULL), { region: 'eu', ruleset: 'seasonal' }).chip).toBe('seasonal EU');
  });

  // Fix round 2 item 4: the guild in the client's own angle brackets, and the spec piece carries
  // the newest build's talent split.
  it('writes the guild the way the client does, in angle brackets, still linked', () => {
    const header = profileHeader(page(FULL, 'Holy'), PATH);
    expect(header.pieces[1]).toEqual({
      key: 'guild',
      text: '<Cx66 Olympus>',
      href: '/guild/us/hardcore/cx66-olympus',
    });
  });

  it('carries the newest build’s talent split on the spec piece, and none without a build', () => {
    const builds = [
      { talent_split: '21/30/0', spec: 'Holy', first_seen: '2026-11-01T20:00:00Z' },
      { talent_split: '31/20/0', spec: 'Holy', first_seen: '2026-12-09T22:10:00Z' },
      { talent_split: '0/21/30', spec: 'Shadow', first_seen: '2026-11-20T20:00:00Z' },
    ];
    const header = profileHeader({ ...page(FULL, 'Holy'), builds_seen: builds }, PATH);
    expect(header.pieces[0]).toEqual({ key: 'spec', text: 'Holy Priest', split: '31/20/0' });
    expect(profileHeader(page(FULL, 'Holy'), PATH).pieces[0]).toEqual({ key: 'spec', text: 'Holy Priest' });
    // A page from an API that predates builds_seen reads as no build.
    const { builds_seen: _omitted, ...older } = page(FULL, 'Holy');
    expect(profileHeader(older, PATH).pieces[0]).toEqual({ key: 'spec', text: 'Holy Priest' });
  });

  // Fix round 2 item 3: the chip is the last item on the last text line.
  it('puts the chip on the level line when there is one, else on the first line', () => {
    expect(profileHeader(page(FULL, 'Holy'), PATH).chipOn).toBe('levelRace');
    expect(profileHeader(page({ ...FULL, level: null, race: null }, 'Holy'), PATH).chipOn).toBe('pieces');
    expect(profileHeader(page({ class: 'Priest' }, 'Holy'), PATH).chipOn).toBe('pieces');
  });
});
