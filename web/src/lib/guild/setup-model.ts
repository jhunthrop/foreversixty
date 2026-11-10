// web/src/lib/guild/setup-model.ts
// The setup checklist's view of GET …/setup (design/specs/2026-11-09-guild-onboarding.md §3):
// each step's mark, its facts line, its parts, and the texts an officer pastes. Pure, so every
// state is pinned without a page.
import type { StepsMark } from '../../components/ui/StepMark.svelte';
import { GUILD_TIER_MEMBER_LIMIT, PLANS } from '../billing/plans';
import type { GuildSetup, SetupPlan, SetupState, SetupStep, SetupStepKey } from './setup-api';
import { setupCopy } from './setup-copy';
import { visibilityWord } from './settings-view';

/** The steps whose done is an officer's word: they offer a confirm beside Skip. */
export const CONFIRMABLE: ReadonlySet<SetupStepKey> = new Set([
  'officer_ranks',
  'invite_link',
  'report_visibility',
  'linked_role',
]);

/** The optional step: its title carries the word. */
export const OPTIONAL: ReadonlySet<SetupStepKey> = new Set(['linked_role']);

/** The first step still to do: the page's one primary action sits on it. */
export function nextStep(steps: readonly SetupStep[]): SetupStepKey | null {
  return steps.find((s) => s.state === 'todo')?.key ?? null;
}

/** A step's mark in the shared Steps list: done, skipped, the next step, or to do. */
export function stepMark(state: SetupState, isNext: boolean): StepsMark {
  if (state === 'done') return 'ok';
  if (state === 'skipped') return 'skip';
  return isNext ? 'next' : 'todo';
}

/** The answer slot's figure: how many of the steps are done. */
export function doneFigure(setup: Pick<GuildSetup, 'done' | 'total'>): { value: string; unit: string } {
  return { value: String(setup.done), unit: setupCopy.doneFigure(setup.done, setup.total) };
}

/** A sub-item of a step: a part that counts toward it, ticked or not. */
export interface SetupPart {
  key: string;
  label: string;
  done: boolean;
  /** The plan the part needs, when it is not free. */
  plan?: string;
}

export function recruitingParts(setup: Pick<GuildSetup, 'recruiting' | 'discord'>): SetupPart[] {
  const r = setup.recruiting;
  const p = setupCopy.recruiting.parts;
  return [
    { key: 'pitch', label: p.pitch, done: r.pitch },
    { key: 'nights', label: p.nights, done: r.nights !== '' },
    { key: 'wanted', label: p.wanted, done: r.wanted.length > 0 },
    { key: 'role', label: p.role, done: r.role_set },
    { key: 'channel', label: p.channel, done: setup.discord.recruit_channel_name !== null },
  ];
}

export function scheduleParts(setup: Pick<GuildSetup, 'schedule' | 'plan'>): SetupPart[] {
  const p = setupCopy.schedule.parts;
  return [
    { key: 'timezone', label: p.timezone, done: setup.schedule.timezone !== null },
    { key: 'raid', label: p.raid, done: setup.schedule.planned > 0 },
    { key: 'comp', label: p.comp, done: setup.plan !== 'none', plan: setupCopy.schedule.guildPlan },
  ];
}

export function recruitingLine(r: GuildSetup['recruiting']): string {
  const parts: string[] = [r.open ? setupCopy.recruiting.open : setupCopy.recruiting.closed];
  if (r.nights !== '') parts.push(r.nights);
  if (r.wanted.length > 0) parts.push(setupCopy.recruiting.wanted(r.wanted));
  return parts.join(' · ');
}

export function discordLine(d: GuildSetup['discord']): string {
  if (d.installed && d.server_name !== null) {
    if (d.channel_name === null) return setupCopy.discord.pickChannel(d.server_name);
    return `${setupCopy.discord.posts(d.channel_name, d.server_name)} · ${setupCopy.discord.linked(d.linked_members, d.members)}`;
  }
  return d.webhook_connected ? setupCopy.discord.webhook : setupCopy.discord.notInstalled;
}

/** The schedule's line; `when` words the next raid's start in the viewer's zone. */
export function scheduleLine(s: GuildSetup['schedule'], when: (iso: string) => string): string {
  const parts: string[] = [s.timezone ?? setupCopy.schedule.noTimezone];
  if (s.next === null) parts.push(setupCopy.schedule.nothing);
  else
    parts.push(
      setupCopy.schedule.next(s.next.title, when(s.next.starts_at)),
      setupCopy.schedule.planned(s.planned),
    );
  return parts.join(' · ');
}

export function inviteLine(i: GuildSetup['invite']): string {
  return i.made ? setupCopy.invite.joined(i.joined) : setupCopy.invite.none;
}

/** Half the verified accounts synced is done (api setup.go); short of it, the line says how many. */
export function syncNeeded(s: GuildSetup['sync']): number {
  return Math.max(0, Math.ceil(s.accounts / 2) - s.synced);
}

export function syncLine(s: GuildSetup['sync']): string {
  const line = setupCopy.sync.line(s.synced, s.accounts, s.window_days);
  const need = syncNeeded(s);
  return s.accounts > 0 && need > 0 ? `${line} · ${setupCopy.sync.needed(Math.ceil(s.accounts / 2))}` : line;
}

/** The synced share for the bar, 0 with no accounts. */
export function syncShare(s: GuildSetup['sync']): number {
  return s.accounts === 0 ? 0 : s.synced / s.accounts;
}

export function visibilityLine(visibility: string): string {
  return setupCopy.visibility.line(visibilityWord(visibility));
}

/** Whether a rank counts as an officer rank at the guild's setting. */
export function officerRank(rankIndex: number, maxRank: number): boolean {
  return rankIndex <= maxRank;
}

/** The plan rail's current line. */
export function planName(plan: SetupPlan): string {
  if (plan === 'raid_guild') return setupCopy.plan.raidGuild;
  return plan === 'guild' ? setupCopy.plan.guild : setupCopy.plan.none;
}

/** The plan rail's three lines, the seats and the member cap from plans.ts. */
export function planLines(): { label: string; text: string }[] {
  return setupCopy.plan.lines(
    { guild: PLANS.guild.seats, raidGuild: PLANS.raid_guild.seats },
    GUILD_TIER_MEMBER_LIMIT,
  );
}

/** A rank line with the first names on it: `Rank 1 · 2: Kraggor, Mirelle`. */
export function rankLine(line: { rank_index: number; members: number; names: readonly string[] }): string {
  const head = setupCopy.ranks.line(line.rank_index, line.members);
  return line.names.length === 0 ? head : `${head}: ${line.names.join(', ')}`;
}

/** A channel whose name reads as the recruiting channel, for the picker's first choice. */
export const RECRUIT_CHANNEL = /recruit|apply|application|join/i;

/** The steps the settings row and the claim page count: every step done or skipped hides them. */
export function setupOpen(setup: Pick<GuildSetup, 'finished'> | null): boolean {
  return setup !== null && !setup.finished;
}
