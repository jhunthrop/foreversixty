// The loot council's words (web layout spec 2026-11-05 §7.2).
import { describe, expect, it } from 'vitest';
import { buildMockLoot } from '../../fixtures/guild/mock-guild';
import { buildMockCouncilDrop, buildMockLootNights } from '../../fixtures/guild/mock-loot-council';
import type { LootCouncilCandidate, LootDecision } from './loot-council-api';
import {
  REASON_MAX,
  awardAsk,
  awardReason,
  candidateFacts,
  candidateGain,
  candidateSub,
  decisionSub,
  decisionTitle,
  defaultNight,
  dropRows,
  dropWhat,
  msOsTag,
  nightDay,
  nightOptions,
  nightSummary,
  openable,
  ownUsers,
  reasonError,
  tableDrop,
} from './loot-council-view';

const table = buildMockLoot().items;
const nights = buildMockLootNights().nights;
const [current, past] = nights;
const firstCandidate = (): LootCouncilCandidate =>
  buildMockCouncilDrop(current.drops[0].item_id).candidates[0];

describe('nights', () => {
  it('reads the raid day as the guild has it, never shifted by the viewer zone', () => {
    expect(nightDay('2026-12-15')).toBe('Tue 15 Dec');
    expect(nightDay('nonsense')).toBe('nonsense');
  });

  it('offers every night, opens the current one, and sums its drops', () => {
    expect(nightOptions(nights)).toEqual([
      { value: 'raid:12', label: 'Tue 15 Dec' },
      { value: 'raid:11', label: 'Tue 8 Dec' },
    ]);
    expect(defaultNight(nights)?.night).toBe('raid:12');
    expect(defaultNight([past])?.night).toBe('raid:11');
    expect(defaultNight([])).toBeNull();
    expect(nightSummary(current)).toBe('3 drops, 2 open');
    expect(nightSummary({ ...current, drops: [] })).toBe('No drops yet');
  });
});

describe('drop rows', () => {
  it('reads the slot words and who can use an open drop, the winner on an awarded one', () => {
    const rows = dropRows(current, table);
    expect(rows.map((row) => row.state)).toEqual(['open', 'open', 'awarded']);
    expect(rows[0].sub).toMatch(/· \d+ candidates?$/);
    expect(rows[2].sub).toContain(`Awarded to ${current.drops[2].award?.name}`);
  });

  it('names a non-gear drop by its kind', () => {
    expect(dropWhat('quest', undefined)).toBe('Quest item');
    expect(dropWhat('bag', undefined)).toBe('Bag');
    expect(dropWhat('gear', undefined)).toBe('');
  });

  it('opens loot table drops only on the current night, and only what a member can use', () => {
    expect(openable(past, table, null)).toEqual([]);
    const officer = openable(current, table, null);
    expect(officer.some((item) => current.drops.some((d) => d.item_id === item.item_id))).toBe(false);
    expect(officer.some((item) => item.kind === 'skinned')).toBe(false);
    const someone = officer.find((item) => item.kind === 'gear')!.candidates[0].character_key;
    const mine = openable(current, table, [someone]);
    expect(mine.length).toBeGreaterThan(0);
    expect(mine.every((item) => ownUsers(item, [someone]).length > 0)).toBe(true);
    expect(openable(current, table, [])).toEqual([]);
  });

  it('turns a loot table item into an open row', () => {
    const row = tableDrop(table[0], 1084);
    expect(row).toMatchObject({ itemId: table[0].item_id, state: 'open', award: null, encounterId: 1084 });
  });

  it("lists the viewer's own characters once", () => {
    const item = table.find((row) => row.kind === 'gear' && row.candidates.length > 0)!;
    const key = item.candidates[0].character_key;
    expect(
      ownUsers({ ...item, eligible: [{ character_key: key, name: 'x', class: 'warrior', item_id: 1 }] }, [
        key,
      ]),
    ).toEqual([{ key, name: item.candidates[0].name }]);
    expect(ownUsers(undefined, [key])).toEqual([]);
  });
});

describe('a candidate', () => {
  it("reads the gain in the BiS page's unit over what it replaces, then attendance and loot this tier", () => {
    const c = firstCandidate();
    expect(candidateGain(c)).toBe(`+14.2\u00a0DPS over ${c.replaces}`);
    expect(candidateGain({ ...c, fills_empty_slot: true, tier: 1, ilvl_delta: 70 })).toBe(
      'Fills an empty slot',
    );
    expect(candidateGain({ ...c, replaces: '' })).toBe('+14.2\u00a0DPS');
    expect(candidateFacts({ ...c, attendance: { present: 3, nights: 4 }, items_won: 1 })).toEqual([
      '3 of 4 raid nights',
      '1 item this tier',
    ]);
    expect(candidateFacts({ ...c, attendance: { present: 0, nights: 0 }, items_won: 0 })).toEqual([
      '0 items this tier',
    ]);
    expect(candidateSub(c).split(' · ').length).toBeGreaterThanOrEqual(2);
  });

  it('marks MS/OS as declared or derived, and none for a quest drop', () => {
    const c = firstCandidate();
    expect(msOsTag({ ...c, ms_os: 'ms', ms_os_source: 'declared' })).toEqual({
      label: 'MS',
      declared: true,
      title: 'They flagged it main spec',
    });
    expect(msOsTag({ ...c, ms_os: 'os', ms_os_source: 'derived' })?.title).toBe(
      'Not flagged. Off spec by their BiS list',
    );
    expect(msOsTag({ ...c, ms_os: null, ms_os_source: null })).toBeNull();
  });
});

describe('the decision record', () => {
  const d: LootDecision = current.decisions[0];

  it('reads the item to whom, MS/OS, the votes, by whom and the day', () => {
    expect(decisionTitle(d)).toBe(`${d.item_name} to ${d.name}`);
    expect(decisionSub(d, 'UTC')).toBe('Main spec · 3 yes, 0 no of 4 · by Kraggor · Tue 15 Dec');
  });

  it('leaves out what an award never kept, and the votes for a member', () => {
    const member = buildMockLootNights(false).nights[0].decisions[0];
    expect(decisionSub(member, 'UTC')).toBe('Main spec · by Kraggor · Tue 15 Dec');
    expect(decisionSub({ ...d, ms_os: null, votes_yes: null, votes_no: null, awarded_by: '' }, 'UTC')).toBe(
      'Tue 15 Dec',
    );
  });
});

describe('the award confirm', () => {
  it('asks by name and keeps the reason within the cap', () => {
    expect(awardAsk('Deathbringer', 'Seraphine')).toBe('Award Deathbringer to Seraphine?');
    expect(awardReason('  Best upgrade  ')).toBe('Best upgrade');
    expect(awardReason('   ')).toBeNull();
    expect(reasonError('x'.repeat(REASON_MAX))).toBe('');
    expect(reasonError('x'.repeat(REASON_MAX + 1))).toBe('Keep the reason to 200 characters.');
  });
});
