// web/src/lib/guild/applications-page.svelte.ts
// The officers' applications page state (web layout spec 2026-11-05 §7.2 Applications): the list,
// the public card and its Discord role; a vote, Accept or Decline, and the card's save. Reads and
// writes are injected, so a test drives it without the network. The server's 403 is the gate.
import type { CharacterPath } from '../characters';
import { GuildApiError } from './api';
import { appliedDecision, patchOf, withRow, type CardDraft } from './applications-view';
import { loadGuildView, type GuildViewContext, type GuildViewDeps } from './guild-view-load';
import { applicationsCopy as copy } from './recruit-copy';
import type {
  ApplicationDecision,
  ApplicationVote,
  ApplicationVotes,
  OfficerApplication,
  OfficerApplications,
  RecruitCard,
  RecruitPatch,
  RecruitSettings,
} from './recruit-api';

export interface ApplicationsPageDeps extends GuildViewDeps {
  loadCard: (guildId: number) => Promise<RecruitCard>;
  loadList: (guildId: number) => Promise<OfficerApplications>;
  vote: (guildId: number, appId: number, vote: ApplicationVote | null) => Promise<ApplicationVotes>;
  decide: (guildId: number, appId: number, decision: 'accept' | 'decline') => Promise<ApplicationDecision>;
  saveCard: (guildId: number, patch: RecruitPatch) => Promise<RecruitSettings>;
}

export type ApplicationsLoad = 'loading' | 'ready' | 'forbidden' | 'failed';

export class ApplicationsPage {
  load = $state<ApplicationsLoad>('loading');
  context = $state<GuildViewContext | null>(null);
  list = $state<OfficerApplications | null>(null);
  card = $state<RecruitCard | null>(null);
  /** The application whose vote or decision is in flight. */
  busyId = $state<number | null>(null);
  /** A failed vote or decision, by application. */
  rowError = $state<{ id: number; text: string } | null>(null);
  /** The decision just made here: the page says what happened on Discord. */
  lastDecision = $state<OfficerApplication | null>(null);
  saving = $state(false);
  saveLine = $state<{ tone: 'ok' | 'error'; text: string } | null>(null);

  constructor(
    private readonly deps: ApplicationsPageDeps,
    private readonly path: CharacterPath,
  ) {}

  get guildId(): number | null {
    return this.context?.guild.id ?? null;
  }

  async start(): Promise<void> {
    this.load = 'loading';
    try {
      const context = await loadGuildView(this.path, this.deps);
      this.context = context;
      if (context.role !== 'officer') {
        this.load = 'forbidden';
        return;
      }
      const [list, card] = await Promise.all([
        this.deps.loadList(context.guild.id),
        this.deps.loadCard(context.guild.id),
      ]);
      this.list = list;
      this.card = card;
      this.load = 'ready';
    } catch (error) {
      this.load = error instanceof GuildApiError && error.status === 403 ? 'forbidden' : 'failed';
    }
  }

  async vote(appId: number, vote: ApplicationVote | null): Promise<void> {
    const { guildId, list } = this;
    const row = list?.rows.find((r) => r.id === appId);
    if (guildId === null || list === null || row === undefined || this.busyId !== null) return;
    this.busyId = appId;
    this.rowError = null;
    try {
      const votes = await this.deps.vote(guildId, appId, vote);
      this.list = withRow(this.list ?? list, { ...row, votes });
    } catch {
      this.rowError = { id: appId, text: copy.voteFailed };
    } finally {
      this.busyId = null;
    }
  }

  async decide(appId: number, decision: 'accept' | 'decline'): Promise<void> {
    const { guildId, list } = this;
    if (guildId === null || list === null || this.busyId !== null) return;
    this.busyId = appId;
    this.rowError = null;
    try {
      const answer = await this.deps.decide(guildId, appId, decision);
      const current = this.list ?? list;
      // Decided rows lead the record, newest first, as the API orders them on the next read.
      const row = appliedDecision(answer);
      const rest = current.rows.filter((r) => r.id !== appId);
      const pending = rest.filter((r) => r.status === 'pending');
      const decided = rest.filter((r) => r.status !== 'pending');
      this.list = { ...current, rows: [...pending, row, ...decided], pending: pending.length };
      this.lastDecision = row;
    } catch {
      this.rowError = { id: appId, text: copy.decideFailed };
    } finally {
      this.busyId = null;
    }
  }

  async saveCard(draft: CardDraft): Promise<boolean> {
    const { guildId } = this;
    if (guildId === null || this.saving) return false;
    this.saving = true;
    this.saveLine = null;
    try {
      const saved = await this.deps.saveCard(guildId, patchOf(draft));
      this.card = saved.card;
      if (this.list !== null) this.list = { ...this.list, discord_role_id: saved.discord_role_id };
      this.saveLine = { tone: 'ok', text: copy.card.saved };
      return true;
    } catch (error) {
      const field =
        error instanceof GuildApiError && error.fields !== null ? Object.values(error.fields)[0] : '';
      this.saveLine = { tone: 'error', text: field || copy.card.saveFailed };
      return false;
    } finally {
      this.saving = false;
    }
  }
}
