// web/src/lib/character-profile.test.ts
import { describe, expect, it } from 'vitest';
import {
  amountHeader,
  bestRepeatsHistory,
  bisHref,
  characterBuildsCopy,
  characterHeaderCopy,
  characterTableCopy,
  executionCellClass,
  fightRowId,
  latestRatedFight,
  newestBuild,
  newestSpec,
  profileHeader,
  profileSpec,
  seenSplits,
  specClassLabel,
  weakestFightHref,
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
    expect(bisHref('Priest', 'Holy', null, null)).toBe('/bis/priest/holy');
    expect(bisHref('Hunter', 'Beast Mastery', null, null)).toBe('/bis/hunter/beast-mastery');
    expect(bisHref('Priest', null, 60, 'Night Elf')).toBeNull();
    expect(bisHref(undefined, 'Holy', 60, 'Night Elf')).toBeNull();
  });

  // L8.4 fix round 4 item 2: "Gear" opens the spec's page at this character's band (and, when
  // the race gives one, faction), the same route and anchor the /bis landing links use.
  it('opens the spec’s page at the character’s band and faction (fix round 4 item 2)', () => {
    expect(bisHref('Priest', 'Holy', 60, 'Night Elf')).toBe(
      '/bis/priest/holy?faction=alliance#band-alliance-60',
    );
    expect(bisHref('Warrior', 'Fury', 37, 'Orc')).toBe('/bis/warrior/fury?faction=horde#band-horde-30');
    // Below 20 the earliest band; an unknown race keeps the band, with the default panel.
    expect(bisHref('Priest', 'Holy', 12, null)).toBe('/bis/priest/holy#band-alliance-20');
  });
});

// L8.4 fix round 4 item 4: the builds footnote in plain words. Fix round 5 item 4: every
// distinct split seen when there is more than one, and no footnote from a blank split.
describe('the builds footnote (fix round 4 item 4, fix round 5 item 4)', () => {
  const TAIL = 'Logs don’t record the order points were spent, so there is no saved build to open.';

  it('names the one split and says why there is no build to open', () => {
    expect(characterBuildsCopy.footnote(['0/31/20'])).toBe(`Seen in logs as a 0/31/20 split. ${TAIL}`);
  });

  it('names every distinct split when there is more than one (fix round 5 item 4)', () => {
    expect(characterBuildsCopy.footnote(['0/31/20', '20/31/0'])).toBe(
      `Seen in logs as 0/31/20 and 20/31/0 splits. ${TAIL}`,
    );
    expect(characterBuildsCopy.footnote(['0/31/20', '20/31/0', '5/31/15'])).toBe(
      `Seen in logs as 0/31/20, 20/31/0 and 5/31/15 splits. ${TAIL}`,
    );
  });

  it('is absent with no split to name: the blank-split fallback is gone (fix round 5 item 4)', () => {
    expect(characterBuildsCopy.footnote([])).toBeNull();
  });

  it('reads the distinct, non-blank splits, newest first', () => {
    expect(
      seenSplits([
        { talent_split: '20/31/0', spec: 'Holy', first_seen: '2026-10-01T00:00:00Z' },
        { talent_split: '0/31/20', spec: 'Holy', first_seen: '2026-10-12T00:00:00Z' },
        { talent_split: ' 20/31/0 ', spec: 'Holy', first_seen: '2026-10-05T00:00:00Z' },
        { talent_split: '  ', spec: 'Holy', first_seen: '2026-10-11T00:00:00Z' },
      ]),
    ).toEqual(['0/31/20', '20/31/0']);
    // Two builds with the one split name it once: a single split.
    expect(
      seenSplits([
        { talent_split: '0/31/20', spec: 'Holy', first_seen: '2026-10-01T00:00:00Z' },
        { talent_split: '0/31/20', spec: 'Discipline', first_seen: '2026-10-02T00:00:00Z' },
      ]),
    ).toEqual(['0/31/20']);
    expect(seenSplits([{ talent_split: '', first_seen: '2026-10-01T00:00:00Z' }])).toEqual([]);
    expect(seenSplits(undefined)).toEqual([]);
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
    const {
      executionDefinition: _definition,
      executionUnscored: _unscored,
      ratedNote: _rated,
      ...columns
    } = characterTableCopy;
    expect(Object.values(columns)).toEqual(['Boss', 'Spec', 'Parse %', 'Date', 'Execution']);
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

// L8.4 fix round 3 item 8: the split shown and the spec word beside it come from one record.
describe('the header’s spec and split (fix round 3 item 8)', () => {
  const PATH = { region: 'us', ruleset: 'hardcore' };
  const base = {
    character: { name: 'Elyra Duskvale', region: 'us', ruleset: 'hardcore', class: 'Priest' },
    history: [fight('Holy')],
    best: [],
  };

  it('takes the spec word from the newest build when it names one, over the fights’ spec', () => {
    const header = profileHeader(
      {
        ...base,
        builds_seen: [
          { talent_split: '0/31/20', spec: 'Holy', first_seen: '2026-11-01T20:00:00Z' },
          { talent_split: '31/20/0', spec: 'Discipline', first_seen: '2026-12-09T22:10:00Z' },
        ],
      },
      PATH,
    );
    expect(header.pieces[0]).toEqual({ key: 'spec', text: 'Discipline Priest', split: '31/20/0' });
    expect(
      profileSpec({
        ...base,
        builds_seen: [{ talent_split: '31/20/0', spec: 'Discipline', first_seen: 'x' }],
      }),
    ).toBe('Discipline');
  });

  it('falls back to the character’s spec when the newest build names none', () => {
    const header = profileHeader(
      { ...base, builds_seen: [{ talent_split: '0/31/20', first_seen: '2026-12-09T22:10:00Z' }] },
      PATH,
    );
    expect(header.pieces[0]).toEqual({ key: 'spec', text: 'Holy Priest', split: '0/31/20' });
    expect(newestBuild([])).toBeNull();
    expect(profileSpec({ ...base, builds_seen: [] })).toBe('Holy');
  });
});

// L8.4 fix round 5 item 6 (replacing fix round 3 item 10's leading dot): a "·" ends the
// segment before it, glued by a no-break space, and a plain breakable space parts the
// segments, so a wrap never starts a line with a dot.
describe('the header’s segment separator (fix round 5 item 6)', () => {
  it('is the dot glued to the end of the segment before it, then a breakable space', () => {
    expect(characterHeaderCopy.segmentBreak).toBe(' ');
    expect(characterHeaderCopy).not.toHaveProperty('segmentLead');
    expect(characterHeaderCopy.segmentTail).toBe('\u00a0·');
  });
});

// L8.4 fix round 3 items 4, 6 and 7: one Execution cell style for both tables and both states,
// and the column explained.
describe('the Execution cell (fix round 3 items 4 and 6)', () => {
  const tokens = (classes: string): string[] => classes.split(/\s+/).filter((token) => token !== '');

  // Fix round 4 items 3 and 7: the column stays on phones in both tables, so no state is hidden.
  it('shows a scored cell and an unscored one at every width (fix round 4 item 3)', () => {
    for (const scored of [true, false]) {
      const cell = tokens(executionCellClass(scored));
      expect(cell).toContain('flex');
      expect(cell.filter((token) => /(^|:)hidden$/.test(token))).toEqual([]);
    }
  });

  it('centres both states in their row with the same box: 44px high and wide, right-aligned, 16px after the amount on a phone', () => {
    const shape = (classes: string): string[] => tokens(classes).filter((token) => token !== 'text-muted');
    expect(shape(executionCellClass(true))).toEqual(shape(executionCellClass(false)));
    expect(tokens(executionCellClass(true))).toEqual(
      expect.arrayContaining(['min-h-11', 'min-w-11', 'items-center', 'justify-end', 'max-md:ml-1']),
    );
    expect(tokens(executionCellClass(false))).toContain('text-muted');
    expect(tokens(executionCellClass(true))).not.toContain('text-muted');
  });

  // Fix round 5 item 1: damage only, as api/openapi.yaml's ExecutionScore says (L6.7): the
  // simulator scores the dps role alone, so a healer's or a tank's fight has none.
  it('defines the column in one line, damage only (item 7, fix round 5 item 1)', () => {
    expect(characterTableCopy.executionDefinition).toBe(
      'Execution: the fight’s damage as a share of what the same gear and talents do in the simulator; healers and tanks have none.',
    );
    expect(characterTableCopy.executionDefinition).not.toContain('healing');
  });

  // Fix round 5 item 7: what the dash is, read aloud and on hover.
  it('names an unscored cell "Not scored yet" (fix round 5 item 7)', () => {
    expect(characterTableCopy.executionUnscored).toBe('Not scored yet');
  });
});

// L8.5 ruling 1: "Weakest part of a fight: Activity." goes to its fight's row in "Every ranked
// fight". The rating API reads worst_component from its newest card only (api/internal/rating/
// dto.go's bestWorstComponent), and neither the trend nor the history rows carry a per-fight
// Activity, so the fight the part was read from is the newest rated one.
describe('the weakest part’s fight (L8.5 ruling 1)', () => {
  const ranked = (report: string, index: number, at: string): CharacterFight => ({
    ...fight('Holy'),
    report_id: report,
    fight_index: index,
    fought_at: at,
  });
  const TREND = [
    { fought_at: '2026-12-08T20:00:00Z', overall: 61, report_id: 'otherreport1', fight_index: 1 },
    { fought_at: '2026-12-09T22:10:00Z', overall: 65, report_id: 'fixture2abcd', fight_index: 3 },
    { fought_at: '2026-12-04T20:00:00Z', overall: 56, report_id: 'older', fight_index: 2 },
  ];
  const HISTORY = [
    ranked('fixture2abcd', 3, '2026-12-09T22:10:00Z'),
    ranked('otherreport1', 1, '2026-12-08T20:00:00Z'),
  ];

  it('names a row by its report and fight: "fight-fixture2abcd-3"', () => {
    expect(fightRowId({ report_id: 'fixture2abcd', fight_index: 3 })).toBe('fight-fixture2abcd-3');
  });

  it('is the newest rated fight’s row, whatever order the trend comes in', () => {
    expect(weakestFightHref(TREND, HISTORY)).toBe('#fight-fixture2abcd-3');
  });

  it('is null when that fight is not a row of "Every ranked fight", or there is no trend', () => {
    expect(weakestFightHref(TREND, HISTORY.slice(1))).toBeNull();
    expect(weakestFightHref([], HISTORY)).toBeNull();
    expect(weakestFightHref(TREND, [])).toBeNull();
  });

  // Fix round 2 items 1 and 3: the latest rated fight named for the rating card's line (its
  // encounter from "Every ranked fight", its day-month date from the trend) and matched to its
  // row by report and fight index. The rating endpoint's `latest` card carries neither, so it is
  // the trend's newest point -- the card `latest` is built from.
  it('names the latest rated fight: encounter, "9 Dec", and its row (fix round 2 items 1 and 3)', () => {
    expect(latestRatedFight(TREND, HISTORY)).toEqual({
      report_id: 'fixture2abcd',
      fight_index: 3,
      encounter: 'Warden Kelthas',
      date: '9 Dec',
      rowId: 'fight-fixture2abcd-3',
    });
    expect(latestRatedFight(TREND, HISTORY.slice(1))).toEqual({
      report_id: 'fixture2abcd',
      fight_index: 3,
      encounter: null,
      date: '9 Dec',
      rowId: null,
    });
    expect(latestRatedFight([], HISTORY)).toBeNull();
  });
});
