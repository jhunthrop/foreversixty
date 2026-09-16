// web/src/fixtures/report/fixture.test.ts
// The engine's JSON is the contract. This parses the checked-in output through the
// TypeScript mirrors and asserts the two invariants the report island relies on: the field
// names match the Go `json:` tags, and no array field is ever null (the engine normalises
// every empty slice to `[]`, and the island's rendering would break on null).
import { describe, expect, it } from 'vitest';
import fight1 from './fights/1/summary.json';
import fight3 from './fights/3/summary.json';
import fight4 from './fights/4/summary.json';
import meta from './meta.json';
import reportFile from './report.json';
import { asArray, type ReportFile, type ReportMeta, type Summary } from '../../lib/report/types';

const report = reportFile as ReportFile;
const one = fight1 as Summary;
const three = fight3 as Summary;
const four = fight4 as Summary;

/** Every array-valued key the island reads, walked to prove none of them is null. */
function everyArray(summary: Summary): unknown[][] {
  return [
    summary.damage_done,
    summary.damage_taken,
    summary.healing,
    summary.healing_taken,
    summary.deaths,
    summary.auras,
    summary.casts,
    summary.interrupts,
    summary.dispels,
    summary.resources,
    summary.threat,
    summary.combatants,
    summary.roster,
    ...summary.damage_done.flatMap((actor) => [actor.abilities, actor.targets, actor.series]),
    ...summary.deaths.flatMap((death) => [death.last, death.auras_held, death.auras_lost]),
    ...summary.auras.flatMap((track) => [track.segments, track.appliers]),
    ...summary.casts.map((cast) => cast.sequence),
    ...summary.resources.map((track) => track.series),
    ...summary.combatants.flatMap((row) => [
      row.gear,
      row.talents,
      row.consumables,
      row.raid_buffs,
      row.missing_buffs,
    ]),
  ];
}

describe('the checked-in report fixture', () => {
  it('is the report the engine wrote, with four fights numbered from one', () => {
    expect(report.report_id).toBe('fixture2abcd');
    expect(report.engine_version).toBe('0.3.0');
    expect(report.fights.map((f) => f.index)).toEqual([1, 2, 3, 4]);
    expect(report.health.layout).toBe('retail-v16');
    expect(report.health.advanced_logging).toBe(true);
    expect(report.health.parse_errors).toBe(0);
  });

  it('names the encounter fight and marks it a kill with one death', () => {
    const encounter = report.fights[2];
    expect(encounter.kind).toBe('encounter');
    expect(encounter.name).toBe('Warden Kelthas');
    expect(encounter.encounter_id).toBe(9001);
    expect(encounter.kill).toBe(true);
    expect(encounter.in_progress).toBe(false);
    // 60000, not the 40s from ENCOUNTER_START to the last event: the engine extends a
    // fight's window to the next fight's start when nothing else bounds the gap, and since
    // engine 0.2.6 the summary runs the same wall length, so a pull has one length
    // everywhere and every per-second figure divides by it.
    expect(encounter.duration_ms).toBe(60000);
    expect(encounter.deaths).toBe(1);
  });

  // Added in Task 17's fix round 1 so RankingsMode's `wantedFight` guard half -- previously
  // undrivable because the fixture had only one encounter-kind fight -- can be exercised by
  // an interleaving e2e test (report-tabs.spec.ts).
  it('has a second encounter fight, for tests that need two ranked fights', () => {
    const second = report.fights[3];
    expect(second.kind).toBe('encounter');
    expect(second.name).toBe('Skolex the Insatiable');
    expect(second.encounter_id).toBe(9002);
    expect(second.kill).toBe(true);
    expect(second.duration_ms).toBe(10000);
    expect(second.players).toEqual(['Player-4184-000000A1']);
  });

  it('parses a damage actor with abilities, targets and a per-second series', () => {
    const caster = one.damage_done.find((actor) => actor.name === 'Morrowlyn-Nightslayer');
    expect(caster).toBeDefined();
    expect(caster?.total).toBe(1484);
    expect(caster?.effective).toBe(1484);
    expect(caster?.abilities[0].spell_id).toBe(116);
    expect(caster?.abilities[0].name).toBe('Frostbolt');
    expect(caster?.abilities[0].crits).toBe(1);
    expect(caster?.targets[0].name).toBe('Hollow Sentinel');
    expect(caster?.series).toEqual([1484]);
  });

  it('parses a death with its killing blow, last ten and auras held', () => {
    expect(three.deaths).toHaveLength(1);
    const death = three.deaths[0];
    expect(death.name).toBe('Thalgrit-Nightslayer');
    expect(death.at_ms).toBe(10100);
    expect(death.killing_blow?.spell_name).toBe('Anima Lash');
    expect(death.killing_blow?.overkill).toBe(100);
    // Two hits, not three: the 20:12:08 line is a SWING_DAMAGE_LANDED with no
    // SWING_DAMAGE to repeat, and the engine never counts a landed line as a hit of its
    // own (logs/engine/event/event.go), so the recap is the two Anima Lash hits. The
    // first left the tank on 2300 of 9400.
    expect(death.last).toHaveLength(2);
    expect(death.last[0].spell_name).toBe('Anima Lash');
    expect(death.last[0].hp_after).toBe(2300);
    expect(death.auras_held[0].name).toBe('Necrotic Wound');
  });

  it('counts damage taken as effective, with effective min and max', () => {
    // The boss lands Anima Lash twice on the tank: 2900 at 20:12:09, then 2400 at
    // 20:12:10 of which 100 is overkill. So total is the raw 5300 and effective is 5200,
    // and the extremes are the effective amounts -- min is 2300, the killing hit less
    // its overkill, not the 2400 the log line reports.
    const tank = three.damage_taken.find((actor) => actor.guid === 'Player-4184-000000A4');
    expect(tank?.total).toBe(5300);
    expect(tank?.effective).toBe(5200);
    const lash = tank?.abilities.find((ability) => ability.spell_id === 334660);
    expect(lash).toMatchObject({ total: 5300, effective: 5200, overkill: 100, hits: 2 });
    expect(lash?.min).toBe(2300);
    expect(lash?.max).toBe(2900);
    // And the roster's headline figures are the effective ones over the fight's 60 s wall
    // length (the summary and the fight list agree since engine 0.2.6).
    const row = three.roster.find((entry) => entry.guid === 'Player-4184-000000A4');
    expect(three.duration_ms).toBe(60_000);
    expect(row?.damage_taken).toBe(5200);
    expect(row?.dtps).toBeCloseTo(5200 / 60, 5);
  });

  it('renders the encounter’s mechanics table against what happened', () => {
    // Encounter 9001's curated table (logs/engine/mechanics/tables/9001.json) lists one
    // ability of every kind, so every branch of the mode has something to draw: Anima
    // Lash hit the tank twice and killed him; the boss's Anima Surge cast at 20:12:14
    // was never interrupted; the Wrack Soul it put on the mage at 20:12:16 was never
    // dispelled; Anima Cascade is in its kit but never went out on this pull; and
    // Necrotic Wound is the tank debuff nobody can play around.
    expect(three.mechanics?.table_found).toBe(true);
    const rows = three.mechanics?.rows ?? [];
    expect(rows.map((row) => [row.spell_id, row.kind])).toEqual([
      [334660, 'avoidable'],
      [334653, 'interrupt'],
      [321038, 'dispel'],
      [334661, 'avoidable'],
      [320462, 'unavoidable'],
    ]);
    // The two nobody failed carry no counts at all, which is what puts them under
    // "also in the table" rather than in the problems list.
    expect(rows[3].players ?? []).toEqual([]);
    expect(rows[4].players ?? []).toEqual([]);
    expect(rows[0].players).toEqual([
      {
        guid: 'Player-4184-000000A4',
        name: 'Thalgrit-Nightslayer',
        hits: 2,
        damage: 5200,
        first_ms: 9000,
        last_ms: 10_000,
        killed: true,
      },
    ]);
    expect(rows[1]).toMatchObject({ casts: 1 });
    expect(rows[1].stopped ?? 0).toBe(0);
    expect(rows[2]).toMatchObject({ applied: 1 });
    expect(rows[2].dispelled ?? 0).toBe(0);
    // Skolex has no table at all, which is the mode's empty state.
    expect(four.mechanics).toEqual({ table_found: false, rows: [] });
  });

  it('parses a combatant with gear in the engine’s untagged Go field names', () => {
    const tank = three.combatants.find((row) => row.name === 'Baelgrim-Nightslayer');
    expect(tank?.spec).toBe('Protection');
    expect(tank?.item_level).toBe(183);
    expect(tank?.gear[0].ID).toBe(175850);
    expect(tank?.gear[0].ItemLevel).toBe(183);
    expect(tank?.talents).toContain(202751);
  });

  it('parses the roster with class, role and the ranking metrics', () => {
    const healer = three.roster.find((row) => row.name === 'Sunwick-Nightslayer');
    expect(healer?.role).toBe('healer');
    expect(healer?.healing_done).toBe(1580);
    expect(healer?.hps).toBeCloseTo(1580 / 60, 5);
    const mage = three.roster.find((row) => row.name === 'Morrowlyn-Nightslayer');
    expect(mage?.class).toBe('Mage');
    expect(mage?.class_source).toBe('combatant_info');
  });

  it('keeps a two-part character name intact, segment and all', () => {
    // Forever names are two parts; the log appends the realm-or-ruleset segment with a
    // hyphen. The engine stores the whole logged string, and splitting it is the web's job
    // (src/lib/report/format.ts splitUnitName).
    const elyra = three.roster.find((row) => row.name.startsWith('Elyra'));
    expect(elyra?.name).toBe('Elyra Duskvale-Hardcore');
    expect(three.damage_done.some((actor) => actor.name === 'Elyra Duskvale-Hardcore')).toBe(true);
  });

  it('parses auras, casts, resources and threat', () => {
    // Eight tracks, because the engine opens one for every aura a COMBATANT_INFO
    // snapshot says was already up at the pull (summary.go seedAuras). The warrior and
    // the mage each brought in 17 and 871, which no aura event inside this fight names,
    // so they are filed under their spell ids; the priest brought in her own Fortitude;
    // the warrior's Fortitude is the one the priest casts at 20:12:04; the boss's
    // Necrotic Wound lands on the tank; and its Wrack Soul lands on the mage and is
    // never dispelled, which is the dispel row of the encounter's mechanics table.
    expect(three.auras.map((track) => track.name).sort()).toEqual([
      'Necrotic Wound',
      'Power Word: Fortitude',
      'Power Word: Fortitude',
      'Spell #17',
      'Spell #17',
      'Spell #871',
      'Spell #871',
      'Wrack Soul',
    ]);
    // By target as well as name: two tracks share the name, and which one comes first is
    // the accumulator's ordering, not a fact about the fight. This is the warrior's, the
    // one the priest casts at 20:12:04 and lets fall at 20:12:35.
    expect(
      three.auras.find((t) => t.name === 'Power Word: Fortitude' && t.target_guid === 'Player-4184-000000A1')
        ?.uptime_ms,
    ).toBe(31000);
    // The mage's one cast, five seconds in. Keyed by caster, not by position: the boss's
    // own Anima Surge cast is a row here too.
    expect(three.casts.find((row) => row.guid === 'Player-4184-000000A3')?.sequence).toEqual([5000]);
    expect(three.resources.every((track) => Array.isArray(track.series))).toBe(true);
    expect(three.threat[0].model_version).toBe('base-1');
    expect(three.threat[0].complete).toBe(false);
  });

  it('never emits a null array', () => {
    for (const value of [...everyArray(one), ...everyArray(three)]) {
      expect(Array.isArray(value)).toBe(true);
    }
  });

  it('has a meta.json shaped like GET /v1/reports/{id}', () => {
    const m = meta as ReportMeta;
    expect(m.id).toBe(report.report_id);
    expect(m.visibility).toBe('public');
    expect(m.status).toBe('complete');
    expect(m.data_base_url).toBe('/logs-data/reports/fixture2abcd');
    expect(m.fights.map((f) => f.index)).toEqual(report.fights.map((f) => f.index));
  });

  it('asArray turns the nulls a future engine might emit into empty arrays', () => {
    expect(asArray<number>(null)).toEqual([]);
    expect(asArray<number>(undefined)).toEqual([]);
    expect(asArray([1, 2])).toEqual([1, 2]);
  });
});
