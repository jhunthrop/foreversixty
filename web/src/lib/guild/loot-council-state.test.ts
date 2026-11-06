// The loot council's flow with the API injected: load, pick a night and a drop, vote, award, undo,
// a member's flag, and a plan refusal flipping the gate once.
import { describe, expect, it, vi } from 'vitest';
import { buildMockLoot } from '../../fixtures/guild/mock-guild';
import {
  buildMockCouncilDrop,
  buildMockLootNights,
  MOCK_PAST_NIGHT,
} from '../../fixtures/guild/mock-loot-council';
import { GuildApiError } from './api';
import { LootCouncilState, WRITE_COPY, type LootCouncilDeps } from './loot-council-state.svelte';
import { dropRows } from './loot-council-view';

function deps(overrides: Partial<LootCouncilDeps> = {}): LootCouncilDeps {
  return {
    nights: vi.fn(async () => buildMockLootNights()),
    table: vi.fn(async () => buildMockLoot()),
    drop: vi.fn(async (_id: number, itemId: number, night: string) => buildMockCouncilDrop(itemId, night)),
    interest: vi.fn(async () => ({ flags: [] })),
    vote: vi.fn(async () => ({ yes: 3, no: 0, voters: 4, mine: 'yes' as const })),
    flag: vi.fn(async (_id: number, _t: unknown, msOs: 'ms' | 'os' | null) => msOs),
    award: vi.fn(async () => ({ award_id: 5 })),
    undo: vi.fn(async () => ({ status: 'removed' })),
    ...overrides,
  };
}

async function loaded(officer: boolean, api = deps(), onRefused = vi.fn()) {
  const state = new LootCouncilState({ guildId: 7, officer, onRefused, deps: api });
  await state.load();
  return { state, api, onRefused };
}

const firstRow = (state: LootCouncilState) => dropRows(state.night!, state.table)[0];

describe('LootCouncilState', () => {
  it('opens the current night with the loot table, and fails as one state', async () => {
    const { state } = await loaded(true);
    expect(state.status).toBe('ready');
    expect(state.nightKey).toBe('raid:12');
    expect(state.table.length).toBeGreaterThan(0);

    const failed = await loaded(true, deps({ nights: vi.fn(async () => Promise.reject(new Error('down'))) }));
    expect(failed.state.status).toBe('failed');
  });

  it('keeps working without the loot table', async () => {
    const { state } = await loaded(true, deps({ table: vi.fn(async () => Promise.reject(new Error('x'))) }));
    expect(state.status).toBe('ready');
    expect(state.table).toEqual([]);
  });

  it("reads an officer's drop candidates for the chosen night", async () => {
    const { state, api } = await loaded(true);
    const row = firstRow(state);
    state.select(row);
    await vi.waitFor(() => expect(state.dropStatus).toBe('ready'));
    expect(api.drop).toHaveBeenCalledWith(7, row.itemId, 'raid:12');
    expect(state.drop?.candidates.length).toBeGreaterThan(0);

    state.chooseNight(MOCK_PAST_NIGHT);
    expect(state.selected).toBeNull();
    expect(state.drop).toBeNull();
  });

  it('records a vote on the candidate and reads the nights again', async () => {
    const { state, api } = await loaded(true);
    state.select(firstRow(state));
    await vi.waitFor(() => expect(state.dropStatus).toBe('ready'));
    const key = state.drop!.candidates[1].character_key;
    await state.vote(key, 'yes');
    expect(api.vote).toHaveBeenCalledWith(
      7,
      { night: 'raid:12', item_id: firstRow(state).itemId, character_key: key },
      'yes',
    );
    expect(state.drop!.candidates[1].votes.yes).toBe(3);
    expect(api.nights).toHaveBeenCalledTimes(2);
  });

  it('awards with the encounter, MS/OS and reason, then refreshes', async () => {
    const { state, api } = await loaded(true);
    const row = firstRow(state);
    state.select(row);
    state.ask('k');
    await state.award('k', 'os', 'Main spec passed');
    expect(api.award).toHaveBeenCalledWith(7, {
      night: 'raid:12',
      item_id: row.itemId,
      character_key: 'k',
      encounter_id: 1084,
      ms_os: 'os',
      reason: 'Main spec passed',
    });
    expect(state.confirmingKey).toBeNull();
  });

  it('undoes the open drop award', async () => {
    const { state, api } = await loaded(true);
    const awarded = dropRows(state.night!, state.table)[2];
    state.select(awarded);
    await vi.waitFor(() => expect(state.drop?.award).not.toBeNull());
    await state.undo();
    expect(api.undo).toHaveBeenCalledWith(7, state.night!.drops[2].award?.award_id ?? 1);
  });

  it('flips the gate once and says so on a plan refusal', async () => {
    const refused = vi.fn(async () => Promise.reject(new GuildApiError('No', 403, 'raid_guild_required')));
    const { state, onRefused } = await loaded(true, deps({ vote: refused }));
    state.select(firstRow(state));
    await state.vote('k', 'no');
    expect(onRefused).toHaveBeenCalledWith('raid-guild');
    expect(state.notice).toBe(WRITE_COPY.vote.ended);
  });

  it('reads a closed drop again when the vote came too late', async () => {
    const closed = vi.fn(async () => Promise.reject(new GuildApiError('No', 409, 'voting_closed')));
    const { state, api } = await loaded(true, deps({ vote: closed }));
    state.select(firstRow(state));
    await state.vote('k', 'yes');
    expect(state.notice).toBe(WRITE_COPY.closed);
    expect(api.drop).toHaveBeenCalledTimes(2);
  });

  it("gives a failed write the kind's own line", async () => {
    const { state } = await loaded(
      true,
      deps({ award: vi.fn(async () => Promise.reject(new Error('net'))) }),
    );
    state.open(state.table[0].item_id);
    await state.award('k', 'ms', null);
    expect(state.notice).toBe(WRITE_COPY.award.failed);
  });

  it("reads and writes a member's own flags, never the candidates", async () => {
    const api = deps({
      interest: vi.fn(async () => ({ flags: [{ item_id: 1, character_key: 'a', ms_os: 'os' as const }] })),
    });
    const { state } = await loaded(false, api);
    expect(state.flags).toEqual({ '1:a': 'os' });
    const row = firstRow(state);
    state.select(row);
    expect(api.drop).not.toHaveBeenCalled();
    await state.flag('a', 'ms');
    expect(state.flags[`${row.itemId}:a`]).toBe('ms');
    await state.flag('a', null);
    expect(state.flags[`${row.itemId}:a`]).toBeUndefined();
    expect(state.flagPending).toBeNull();
  });

  it('opens a loot table drop by id and ignores an unknown one', async () => {
    const { state } = await loaded(true);
    state.open(-1);
    expect(state.selected).toBeNull();
    state.open(state.table[0].item_id);
    expect(state.selected?.itemId).toBe(state.table[0].item_id);
  });
});
