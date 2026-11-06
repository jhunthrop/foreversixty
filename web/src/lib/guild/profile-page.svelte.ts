// web/src/lib/guild/profile-page.svelte.ts
// A member profile's state (web layout spec 2026-11-05 §7.2 Member profile): the guild and who is
// asking, the profile (officer or self), and the officers' notes with Add and Delete. Reads and
// writes are injected, so a test drives it without the network. A member reading their own
// profile never holds notes: the API leaves the field out and this keeps none.
import type { CharacterPath } from '../characters';
import { GuildApiError } from './api';
import { loadGuildView, type GuildViewContext, type GuildViewDeps } from './guild-view-load';
import { withNote, withoutNote } from './profile-view';
import { profileCopy as copy } from './recruit-copy';
import type { MemberNote, MemberProfile } from './recruit-api';

export interface ProfilePageDeps extends GuildViewDeps {
  loadProfile: (guildId: number, account: string) => Promise<MemberProfile>;
  addNote: (guildId: number, account: string, body: string) => Promise<MemberNote>;
  deleteNote: (guildId: number, account: string, noteId: number) => Promise<{ status: 'removed' }>;
}

export type ProfileLoad = 'loading' | 'ready' | 'signed-out' | 'forbidden' | 'not-found' | 'failed';

export class ProfilePage {
  load = $state<ProfileLoad>('loading');
  context = $state<GuildViewContext | null>(null);
  profile = $state<MemberProfile | null>(null);
  /** Officers only; null for a member's own read. */
  notes = $state<MemberNote[] | null>(null);
  saving = $state(false);
  deletingId = $state<number | null>(null);
  noteError = $state('');

  constructor(
    private readonly deps: ProfilePageDeps,
    private readonly path: CharacterPath,
    private readonly account: string,
  ) {}

  async start(): Promise<void> {
    this.load = 'loading';
    try {
      const context = await loadGuildView(this.path, this.deps);
      this.context = context;
      if (context.me === null) {
        this.load = 'signed-out';
        return;
      }
      const profile = await this.deps.loadProfile(context.guild.id, this.account);
      this.profile = profile;
      this.notes = profile.notes ?? null;
      this.load = 'ready';
    } catch (error) {
      this.load = loadOf(error);
    }
  }

  /** The account the notes are about: the profile's own key, so `me` never names a note's subject. */
  private get subject(): string {
    return this.profile?.account_key ?? this.account;
  }

  async addNote(body: string): Promise<boolean> {
    const guildId = this.context?.guild.id;
    const trimmed = body.trim();
    if (guildId === undefined || this.notes === null || trimmed === '' || this.saving) return false;
    this.saving = true;
    this.noteError = '';
    try {
      const note = await this.deps.addNote(guildId, this.subject, trimmed);
      this.notes = withNote(this.notes ?? [], note);
      return true;
    } catch {
      this.noteError = copy.noteFailed;
      return false;
    } finally {
      this.saving = false;
    }
  }

  async deleteNote(noteId: number): Promise<void> {
    const guildId = this.context?.guild.id;
    if (guildId === undefined || this.notes === null || this.deletingId !== null) return;
    this.deletingId = noteId;
    this.noteError = '';
    try {
      await this.deps.deleteNote(guildId, this.subject, noteId);
      this.notes = withoutNote(this.notes ?? [], noteId);
    } catch {
      this.noteError = copy.deleteFailed;
    } finally {
      this.deletingId = null;
    }
  }
}

function loadOf(error: unknown): ProfileLoad {
  if (!(error instanceof GuildApiError)) return 'failed';
  if (error.status === 401) return 'signed-out';
  if (error.status === 403) return 'forbidden';
  if (error.status === 404) return 'not-found';
  return 'failed';
}
