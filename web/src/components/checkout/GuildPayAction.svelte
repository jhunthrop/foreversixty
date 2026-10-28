<!-- web/src/components/checkout/GuildPayAction.svelte -->
<!-- Spec 2026-10-27-subscribe-and-join §5.3: the pay controls of a guild card, drawn through
     GuildStepBlock's `action` slot. `ready`: the switch note (when the chosen plan does not fit
     the verified count) above the Subscribe button, the page's one fill (a gold outline when two
     or more cards are ready). `on_plan` the viewer pays for: Manage billing, or Move to Raid
     guild when Raid guild is chosen on a Guild plan. The decisions are card-checkout.ts; the
     shell owns the calls, so this draws and reports. Each button carries data-step-action, the
     focus target after a re-read. -->
<script lang="ts">
  import type { MeGuild } from '../../lib/account/me-types';
  import {
    isReadyFill,
    OUTLINE_GOLD_BUTTON,
    payAction,
    type PayBusy,
    PLAIN_BUTTON,
    readyLabel,
    upgradeHref,
  } from '../../lib/billing/card-checkout';
  import { checkoutGuildCopy as copy, type GuildTier } from '../../lib/billing/checkout-guild-copy';
  import { priceFor } from '../../lib/billing/plan-costs';
  import { formatPrice, type Interval } from '../../lib/billing/plans';
  import { switchNote, type StepGuild } from '../../lib/billing/subscribe-step';
  import { INLINE_LINK } from '../../lib/guild/field-styles';
  import { PRIMARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { BUSY_CLASS } from '../../lib/ui/busy';

  let {
    guild,
    plan,
    interval,
    founder,
    readyCards,
    busy,
    locked,
    guildPageHref,
    onSubscribe,
    onManage,
    onSwitch,
  }: {
    guild: MeGuild & StepGuild;
    plan: GuildTier;
    interval: Interval;
    founder: boolean;
    /** How many cards are `ready`: more than one turns every fill into an outline. */
    readyCards: number;
    busy: PayBusy;
    /** Another card is redirecting. */
    locked: boolean;
    guildPageHref: string;
    onSubscribe: () => void;
    onManage: () => void;
    onSwitch: (to: GuildTier) => void;
  } = $props();

  const step = $derived(guild.subscribe?.step ?? null);
  const price = $derived(formatPrice(priceFor(plan, interval, founder), interval));
  const note = $derived(switchNote(plan, guild.verified_count));
  const action = $derived(payAction(guild, plan));
  const subscribeClass = $derived(
    isReadyFill(readyCards) ? `${PRIMARY_BUTTON_FIXED} w-full justify-center` : OUTLINE_GOLD_BUTTON,
  );
</script>

{#if step === 'ready'}
  {#if note !== null}
    <p class="text-muted m-0 text-[13px]" data-testid={`guild-${guild.id}-switch-note`}>
      {note.text}
      <button type="button" class={INLINE_LINK} onclick={() => onSwitch(note.to)}>{note.link}</button>
    </p>
  {/if}
  <button
    type="button"
    class={`${subscribeClass} ${busy === 'checkout' ? BUSY_CLASS : ''}`}
    disabled={busy !== null || locked}
    aria-busy={busy === 'checkout'}
    onclick={onSubscribe}
    data-step-action
    data-testid={`guild-${guild.id}-subscribe`}
  >
    {busy === 'checkout' ? copy.busy.checkout : readyLabel(guild, price)}<span class="sr-only"
      >, {guild.name}</span
    >
  </button>
{:else if action === 'manage'}
  <button
    type="button"
    class={`${PLAIN_BUTTON} ${busy === 'portal' ? BUSY_CLASS : ''}`}
    disabled={busy !== null || locked}
    aria-busy={busy === 'portal'}
    onclick={onManage}
    data-step-action
    data-testid={`guild-${guild.id}-manage`}
  >
    {copy.onPlan.manageBilling}<span class="sr-only">, {guild.name}</span>
  </button>
{:else if action === 'move'}
  <a
    href={upgradeHref(guildPageHref)}
    class={OUTLINE_GOLD_BUTTON}
    data-step-action
    data-testid={`guild-${guild.id}-move`}
  >
    {copy.onPlan.moveToRaidGuild}<span class="sr-only">, {guild.name}</span>
  </a>
{/if}
