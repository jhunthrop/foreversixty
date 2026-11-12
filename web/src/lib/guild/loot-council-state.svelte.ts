// web/src/lib/guild/loot-council-state.svelte.ts
// The loot council tab's state (web layout spec 2026-11-05 §7.2): the nights and the loot table,
// the chosen night and drop, the drop's candidates (officers) or the viewer's own MS/OS flags
// (members), and the writes (vote, award, undo, flag). The API is injected, so the flow is tested
// without a page. A plan refusal flips the page's gate through `onRefused` and says so once.
import {
  GuildApiError,
  fetchGuildLoot,
  unawardLoot,
  type GuildLootItem,
  type GuildLootPage,
  type GuildLootRaid,
} from './api';
import {
  awardDrop,
  fetchDropCandidates,
  fetchLootInterest,
  fetchLootNights,
  putLootInterest,
  putLootVote,
  type CouncilAward,
  type DropTarget,
  type LootCouncilDrop,
  type LootCouncilVotes,
  type LootNights,
  type MsOs,
  type Vote,
} from './loot-council-api';
import { defaultNight, tableDrop, type DropRow } from './loot-council-view';
import { refusalGate } from './officer-tools';
import { withRetry } from '../ui/copy';

export interface LootCouncilDeps {
  nights: (guildId: number) => Promise<LootNights>;
  table: (guildId: number) => Promise<GuildLootPage>;
  drop: (guildId: number, itemId: number, night: string) => Promise<LootCouncilDrop>;
  interest: (
    guildId: number,
    night: string,
  ) => Promise<{ flags: { character_key: string; item_id: number; ms_os: MsOs }[] }>;
  vote: (guildId: number, target: DropTarget, vote: Vote | null) => Promise<LootCouncilVotes>;
  flag: (guildId: number, target: DropTarget, msOs: MsOs | null) => Promise<MsOs | null>;
  award: (guildId: number, award: CouncilAward) => Promise<unknown>;
  undo: (guildId: number, awardId: number) => Promise<unknown>;
}

export const LOOT_COUNCIL_API: LootCouncilDeps = {
  nights: (id) => fetchLootNights(id),
  table: (id) => fetchGuildLoot(id),
  drop: (id, itemId, night) => fetchDropCandidates(id, itemId, night),
  interest: (id, night) => fetchLootInterest(id, night),
  vote: (id, target, vote) => putLootVote(id, target, vote),
  flag: (id, target, msOs) => putLootInterest(id, target, msOs),
  award: (id, award) => awardDrop(id, award),
  undo: (id, awardId) => unawardLoot(id, awardId),
};

export const WRITE_COPY = {
  vote: {
    ended: "Vote not recorded. This guild's officer tools have ended.",
    failed: withRetry('That vote did not go through.'),
  },
  award: {
    ended: "Not awarded. This guild's officer tools have ended.",
    failed: withRetry('That award did not go through.'),
  },
  undo: {
    ended: "Not undone. This guild's officer tools have ended.",
    failed: withRetry('That did not go through.'),
  },
  flag: {
    ended: "Not saved. This guild's officer tools have ended.",
    failed: withRetry('That did not save.'),
  },
  closed: 'That drop was just awarded.',
} as const;

type WriteKind = 'vote' | 'award' | 'undo' | 'flag';
type Status = 'loading' | 'ready' | 'failed';

export interface LootCouncilOptions {
  guildId: number;
  officer: boolean;
  /** A plan refusal: the page flips Loot to its gate. */
  onRefused: (gate: 'plan' | 'raid-guild') => void;
  deps?: LootCouncilDeps;
}

/** The candidate list with one candidate's tally replaced. */
function withVotes(drop: LootCouncilDrop, key: string, votes: LootCouncilVotes): LootCouncilDrop {
  return {
    ...drop,
    candidates: drop.candidates.map((c) => (c.character_key === key ? { ...c, votes } : c)),
  };
}

export class LootCouncilState {
  status = $state<Status>('loading');
  nights = $state<LootNights | null>(null);
  table = $state<readonly GuildLootItem[]>([]);
  raids = $state<readonly GuildLootRaid[]>([]);
  nightKey = $state<string | null>(null);
  selected = $state<DropRow | null>(null);
  drop = $state<LootCouncilDrop | null>(null);
  dropStatus = $state<Status>('loading');
  /** The viewer's own flags on the chosen night, `${itemId}:${characterKey}` -> MS/OS. */
  flags = $state<Readonly<Record<string, MsOs>>>({});
  confirmingKey = $state<string | null>(null);
  busy = $state(false);
  flagPending = $state<string | null>(null);
  notice = $state('');

  private readonly opts: LootCouncilOptions;
  private readonly deps: LootCouncilDeps;

  constructor(opts: LootCouncilOptions) {
    this.opts = opts;
    this.deps = opts.deps ?? LOOT_COUNCIL_API;
  }

  get night() {
    return this.nights?.nights.find((n) => n.night === this.nightKey) ?? null;
  }

  /** Reads the nights and the loot table; the table is optional (the picker and the slot words). */
  load = async (): Promise<void> => {
    this.status = 'loading';
    try {
      const [nights, table] = await Promise.all([
        this.deps.nights(this.opts.guildId),
        this.deps.table(this.opts.guildId).catch(() => null),
      ]);
      this.nights = nights;
      this.table = table?.items ?? [];
      this.raids = table?.raids ?? [];
      if (this.night === null) this.nightKey = defaultNight(nights.nights)?.night ?? null;
      this.status = 'ready';
      await this.readFlags();
    } catch {
      this.status = 'failed';
    }
  };

  chooseNight = (key: string): void => {
    if (key === this.nightKey) return;
    this.nightKey = key;
    this.clearSelection();
    void this.readFlags();
  };

  select = (row: DropRow): void => {
    this.clearSelection();
    this.selected = row;
    if (this.opts.officer) void this.readDrop();
  };

  /** Opens a loot table drop on the current night. */
  open = (itemId: number): void => {
    const item = this.table.find((row) => row.item_id === itemId);
    if (item !== undefined) this.select(tableDrop(item));
  };

  ask = (characterKey: string): void => {
    this.confirmingKey = characterKey;
  };

  cancel = (): void => {
    this.confirmingKey = null;
  };

  vote = async (characterKey: string, vote: Vote | null): Promise<void> => {
    const target = this.target(characterKey);
    if (target === null) return;
    await this.write('vote', async (id) => {
      const votes = await this.deps.vote(id, target, vote);
      if (this.drop !== null) this.drop = withVotes(this.drop, characterKey, votes);
    });
    // A first vote puts a loot table drop on the night.
    await this.reloadNights();
  };

  award = async (characterKey: string, msOs: MsOs | null, reason: string | null): Promise<void> => {
    const target = this.target(characterKey);
    if (target === null || this.selected === null) return;
    const body: CouncilAward = { ...target, encounter_id: this.selected.encounterId, ms_os: msOs, reason };
    const done = await this.write('award', (id) => this.deps.award(id, body));
    if (done) await this.refresh();
  };

  undo = async (): Promise<void> => {
    const awardId = this.drop?.award?.award_id;
    if (awardId === undefined) return;
    const done = await this.write('undo', (id) => this.deps.undo(id, awardId));
    if (done) await this.refresh();
  };

  flag = async (characterKey: string, msOs: MsOs | null): Promise<void> => {
    const target = this.target(characterKey);
    if (target === null) return;
    this.flagPending = characterKey;
    const key = `${target.item_id}:${characterKey}`;
    await this.write('flag', async (id) => {
      const saved = await this.deps.flag(id, target, msOs);
      const { [key]: _dropped, ...rest } = this.flags;
      this.flags = saved === null ? rest : { ...rest, [key]: saved };
    });
    this.flagPending = null;
    await this.reloadNights();
  };

  readDrop = async (): Promise<void> => {
    const night = this.nightKey;
    const row = this.selected;
    if (night === null || row === null) return;
    this.dropStatus = 'loading';
    try {
      const drop = await this.deps.drop(this.opts.guildId, row.itemId, night);
      if (this.selected?.itemId === row.itemId && this.nightKey === night) {
        this.drop = drop;
        this.dropStatus = 'ready';
      }
    } catch {
      this.dropStatus = 'failed';
    }
  };

  private clearSelection(): void {
    this.selected = null;
    this.drop = null;
    this.confirmingKey = null;
    this.notice = '';
  }

  private target(characterKey: string): DropTarget | null {
    if (this.nightKey === null || this.selected === null) return null;
    return { night: this.nightKey, item_id: this.selected.itemId, character_key: characterKey };
  }

  private async readFlags(): Promise<void> {
    if (this.opts.officer || this.nightKey === null) return;
    try {
      const page = await this.deps.interest(this.opts.guildId, this.nightKey);
      this.flags = Object.fromEntries(page.flags.map((f) => [`${f.item_id}:${f.character_key}`, f.ms_os]));
    } catch {
      this.flags = {};
    }
  }

  private async reloadNights(): Promise<void> {
    try {
      this.nights = await this.deps.nights(this.opts.guildId);
    } catch {
      // The tab keeps the nights it has; the next visit reads them again.
    }
  }

  /** After an award or undo: the nights (the drop's state, the record) and the drop. */
  private async refresh(): Promise<void> {
    this.confirmingKey = null;
    const row = this.selected;
    await this.reloadNights();
    const fresh = this.night?.drops.find((d) => d.item_id === row?.itemId);
    if (row !== null && fresh !== undefined)
      this.selected = { ...row, state: fresh.state, award: fresh.award };
    await this.readDrop();
  }

  /** Runs a write; a refusal or failure becomes the one notice. True when it went through. */
  private async write(kind: WriteKind, run: (guildId: number) => Promise<unknown>): Promise<boolean> {
    this.busy = true;
    this.notice = '';
    try {
      await run(this.opts.guildId);
      return true;
    } catch (thrown) {
      this.notice = this.refusal(kind, thrown);
      if (thrown instanceof GuildApiError && thrown.code === 'voting_closed') await this.refresh();
      return false;
    } finally {
      this.busy = false;
    }
  }

  private refusal(kind: WriteKind, thrown: unknown): string {
    if (!(thrown instanceof GuildApiError)) return WRITE_COPY[kind].failed;
    if (thrown.code === 'voting_closed') return WRITE_COPY.closed;
    const gate = refusalGate(thrown.code);
    if (gate === null) return thrown.message || WRITE_COPY[kind].failed;
    this.opts.onRefused(gate);
    return WRITE_COPY[kind].ended;
  }
}
