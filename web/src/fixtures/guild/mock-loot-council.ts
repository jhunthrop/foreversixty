// web/src/fixtures/guild/mock-loot-council.ts
// The loot council's API answers for tests and the mocked e2e (api/openapi.yaml LootNights,
// LootCouncilDrop), built over buildMockLoot's loot table so names, icons and candidates agree.
// Two nights: the current one with two open drops and one award, and last week's with one award.
import type { GuildLootItem } from '../../lib/guild/api';
import type {
  LootCouncilCandidate,
  LootCouncilDrop,
  LootDecision,
  LootNightDrop,
  LootNights,
} from '../../lib/guild/loot-council-api';
import { buildMockLoot } from './mock-guild';

const ENCOUNTER = 1084;
export const MOCK_CURRENT_NIGHT = 'raid:12';
export const MOCK_PAST_NIGHT = 'raid:11';

function gearWithCandidates(): GuildLootItem[] {
  return buildMockLoot().items.filter((item) => item.kind === 'gear' && item.candidates.length > 0);
}

function decision(item: GuildLootItem, awardId: number, at: string, officer: boolean): LootDecision {
  const winner = item.candidates[0];
  return {
    award_id: awardId,
    item_id: item.item_id,
    item_name: item.name,
    icon: item.icon,
    quality: item.quality,
    character_key: winner.character_key,
    name: winner.name,
    class: winner.class,
    ms_os: 'ms',
    reason: awardId === 1 ? 'Best upgrade, full attendance' : null,
    votes_yes: officer ? 3 : null,
    votes_no: officer ? 0 : null,
    votes_officers: officer ? 4 : null,
    awarded_by: 'Kraggor',
    at,
  };
}

function drop(item: GuildLootItem, award: LootDecision | null): LootNightDrop {
  return {
    item_id: item.item_id,
    name: item.name,
    icon: item.icon,
    quality: item.quality,
    kind: item.kind,
    encounter_id: ENCOUNTER,
    boss: 'Onyxia',
    state: award === null ? 'open' : 'awarded',
    award,
  };
}

/** The nights as the API answers them; a member's copy has no vote counts. */
export function buildMockLootNights(officer = true): LootNights {
  const [first, second, third, fourth] = gearWithCandidates();
  const awarded = decision(third, 1, '2026-12-15T21:10:00Z', officer);
  const past = decision(fourth, 2, '2026-12-08T21:30:00Z', officer);
  return {
    officers: 4,
    nights: [
      {
        night: MOCK_CURRENT_NIGHT,
        raid_id: 12,
        report_id: null,
        date: '2026-12-15',
        at: '2026-12-15T19:30:00Z',
        zone: "Onyxia's Lair",
        current: true,
        drops: [drop(first, null), drop(second, null), drop(third, awarded)],
        decisions: [awarded],
      },
      {
        night: MOCK_PAST_NIGHT,
        raid_id: 11,
        report_id: null,
        date: '2026-12-08',
        at: '2026-12-08T19:30:00Z',
        zone: "Onyxia's Lair",
        current: false,
        drops: [drop(fourth, past)],
        decisions: [past],
      },
    ],
  };
}

function candidate(row: GuildLootItem['candidates'][number], index: number): LootCouncilCandidate {
  return {
    ...row,
    fills_empty_slot: row.fills_empty_slot ?? false,
    tier: index === 0 ? 0 : row.tier,
    gain_dps: index === 0 ? 14.2 : row.gain_dps,
    gain_unit: index === 0 ? 'dps' : row.gain_unit,
    bis_for_them: index === 0,
    ms_os: index < 2 ? 'ms' : 'os',
    ms_os_source: index === 0 ? 'declared' : 'derived',
    votes: { yes: index === 0 ? 2 : 0, no: index === 1 ? 1 : 0, voters: 4, mine: index === 0 ? 'yes' : null },
  };
}

/** One drop's candidates on a night, open unless the night awarded it. */
export function buildMockCouncilDrop(itemId: number, night = MOCK_CURRENT_NIGHT): LootCouncilDrop {
  const item = buildMockLoot().items.find((row) => row.item_id === itemId);
  if (item === undefined) throw new Error(`no mock drop ${itemId}`);
  const award =
    buildMockLootNights()
      .nights.find((n) => n.night === night)
      ?.drops.find((d) => d.item_id === itemId)?.award ?? null;
  return {
    night,
    item_id: item.item_id,
    name: item.name,
    icon: item.icon,
    quality: item.quality,
    kind: item.kind,
    officers: 4,
    closed: award !== null,
    award,
    candidates: item.candidates.map(candidate),
  };
}
