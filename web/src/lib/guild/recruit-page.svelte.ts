// web/src/lib/guild/recruit-page.svelte.ts
// The recruiting page's state (web layout spec 2026-11-05 §7.2 Recruitment): the guild and who is
// asking, the public card, the viewer's own applications; Apply and Withdraw. Its reads and writes
// are injected, so a test drives it without the network.
import type { CharacterPath } from '../characters';
import { loadGuildView, type GuildViewContext, type GuildViewDeps } from './guild-view-load';
import { recruitCopy as copy } from './recruit-copy';
import type { ApplicationInput, GuildApplication, RecruitCard } from './recruit-api';
import { applyErrorLine, recruitState, type RecruitState } from './recruit-view';

export interface RecruitPageDeps extends GuildViewDeps {
  loadCard: (guildId: number) => Promise<RecruitCard>;
  loadMine: (guildId: number) => Promise<{ rows: GuildApplication[] }>;
  apply: (guildId: number, input: ApplicationInput) => Promise<GuildApplication>;
  withdraw: (guildId: number, appId: number) => Promise<GuildApplication>;
}

export class RecruitPage {
  load = $state<'loading' | 'ready' | 'failed'>('loading');
  context = $state<GuildViewContext | null>(null);
  card = $state<RecruitCard | null>(null);
  mine = $state<GuildApplication[]>([]);
  busy = $state(false);
  error = $state('');

  readonly state = $derived<RecruitState | null>(
    this.card === null ? null : recruitState(this.card, this.context?.me ?? null, this.mine),
  );

  constructor(
    private readonly deps: RecruitPageDeps,
    private readonly path: CharacterPath,
  ) {}

  async start(): Promise<void> {
    this.load = 'loading';
    try {
      const context = await loadGuildView(this.path, this.deps);
      const [card, mine] = await Promise.all([
        this.deps.loadCard(context.guild.id),
        context.me === null ? { rows: [] } : this.deps.loadMine(context.guild.id).catch(() => ({ rows: [] })),
      ]);
      this.context = context;
      this.card = card;
      this.mine = mine.rows;
      this.load = 'ready';
    } catch {
      this.load = 'failed';
    }
  }

  async apply(input: ApplicationInput): Promise<void> {
    const guildId = this.context?.guild.id;
    if (guildId === undefined || this.busy) return;
    this.busy = true;
    this.error = '';
    try {
      const application = await this.deps.apply(guildId, input);
      this.mine = [application, ...this.mine];
    } catch (error) {
      this.error = applyErrorLine(error);
    } finally {
      this.busy = false;
    }
  }

  async withdraw(): Promise<void> {
    const guildId = this.context?.guild.id;
    const open = this.mine.find((app) => app.status === 'pending');
    if (guildId === undefined || open === undefined || this.busy) return;
    this.busy = true;
    this.error = '';
    try {
      const withdrawn = await this.deps.withdraw(guildId, open.id);
      this.mine = this.mine.map((app) => (app.id === withdrawn.id ? withdrawn : app));
    } catch {
      this.error = copy.mine.withdrawFailed;
    } finally {
      this.busy = false;
    }
  }
}
