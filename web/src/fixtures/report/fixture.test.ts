// web/src/fixtures/report/fixture.test.ts
// The engine's JSON is the contract. This parses the checked-in output through the
// TypeScript mirrors and asserts the two invariants the report island relies on: the field
// names match the Go `json:` tags, and no array field is ever null (the engine normalises
// every empty slice to `[]`, and the island's rendering would break on null).
import { describe, expect, it } from 'vitest';
import fight1 from './fights/1/summary.json';
import fight3 from './fights/3/summary.json';
import meta from './meta.json';
import reportFile from './report.json';
import { asArray, type ReportFile, type ReportMeta, type Summary } from '../../lib/report/types';

const report = reportFile as ReportFile;
const one = fight1 as Summary;
const three = fight3 as Summary;

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
    expect(report.engine_version).toBe('0.1.0');
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
    // 60000, not the encounter's own 40s ENCOUNTER_START-to-ENCOUNTER_END span: the engine
    // extends a fight's report.json window to the next fight's start when nothing else
    // bounds the gap between them. fights/3/summary.json's own duration_ms stays exactly
    // 40000 -- the combat-derived figure every table and DPS/HPS calculation reads -- so
    // this is a report.json-list-only figure (Task 17 fix round 1 discovered this while
    // adding fight 4 below; see task-17-report.md).
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
    expect(death.last).toHaveLength(3);
    expect(death.last[0].hp_after).toBe(5200);
    expect(death.auras_held[0].name).toBe('Necrotic Wound');
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
    expect(healer?.hps).toBeCloseTo(39.5, 5);
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
    expect(three.auras.map((track) => track.name).sort()).toEqual([
      'Necrotic Wound',
      'Power Word: Fortitude',
    ]);
    expect(three.auras.find((t) => t.name === 'Power Word: Fortitude')?.uptime_ms).toBe(31000);
    expect(three.casts[0].sequence).toEqual([5000]);
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
