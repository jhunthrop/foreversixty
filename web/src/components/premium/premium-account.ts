// web/src/components/premium/premium-account.ts
// /premium's signed-in read (premium spec 2026-11-08 section 5, gap G1): imported only when the
// pre-paint `data-session` says someone is signed in, so a signed-out visit loads none of it. One
// `/v1/me` (the shared cache), then the head answer, the `Your plan` marks, the state line and a
// guild column's on-plan line are drawn into the server-rendered slots.
import { mount } from 'svelte';
import { fetchMeOnce, type Me } from '../../lib/account/api';
import { writePremiumHint } from '../../lib/account/premium-hint';
import { premiumSource } from '../../lib/account/plan-line';
import type { GuildTier } from '../../lib/billing/checkout-guild-copy';
import {
  guildOnPlan,
  premiumAnswerView,
  premiumMine,
  premiumStateLine,
} from '../../lib/billing/premium-page';
import GuildOnPlan from './GuildOnPlan.svelte';
import PremiumAnswer from './PremiumAnswer.svelte';
import PremiumStateLine from './PremiumStateLine.svelte';

const GUILD_TIERS: readonly GuildTier[] = ['guild', 'raid_guild'];

function slot(name: string, key?: string): HTMLElement | null {
  const selector =
    key === undefined ? `[data-premium-slot="${name}"]` : `[data-premium-slot="${name}"][data-tier="${key}"]`;
  return document.querySelector<HTMLElement>(selector);
}

function drawGuilds(me: Me): void {
  for (const tier of GUILD_TIERS) {
    const view = guildOnPlan(me, tier);
    const column = document.querySelector<HTMLElement>(`[data-guild-column="${tier}"]`);
    const target = slot('guild', tier);
    if (view === null || column === null || target === null) continue;
    column.dataset.onPlan = 'true';
    const tag = column.querySelector<HTMLElement>('[data-guild-tag]');
    if (tag !== null) tag.textContent = view.guildName;
    mount(GuildOnPlan, { target, props: { view } });
  }
}

/** Draws every signed-in part of the page from one account. */
export function drawAccount(me: Me, now: Date = new Date(), doc: Document = document): void {
  writePremiumHint(premiumSource(me, now), doc);
  doc.documentElement.dataset.planMine = premiumMine(me, now);
  const answer = slot('answer');
  if (answer !== null) mount(PremiumAnswer, { target: answer, props: { view: premiumAnswerView(me, now) } });
  const line = premiumStateLine(me, now);
  const state = slot('state');
  if (line !== null && state !== null) mount(PremiumStateLine, { target: state, props: { line } });
  drawGuilds(me);
}

export async function bootPremiumAccount(): Promise<void> {
  let me: Me | null;
  try {
    me = await fetchMeOnce();
  } catch {
    // The page is complete signed out; a failed read leaves it so (the hint's look stays).
    return;
  }
  if (me === null) {
    // A stale session cookie: the answer's reserved box closes.
    document.documentElement.dataset.planMine = 'out';
    return;
  }
  drawAccount(me);
}
