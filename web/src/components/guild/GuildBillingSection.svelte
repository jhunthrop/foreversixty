<!-- web/src/components/guild/GuildBillingSection.svelte -->
<!-- The Billing section of guild settings: the plan line and its explanation for every viewer who
     may see the plan, the price line and buttons only for the claim holder (spec 5.1 to 5.3). -->
<script lang="ts">
  import type { GuildBillingView } from '../../lib/account/api';
  import { billingBlockCopy } from '../../lib/billing/copy';
  import { billedLine, guildPlanLine } from '../../lib/billing/guild-billing-copy';
  import { guildRosterModelCopy } from '../../lib/guild/copy';
  import {
    billingState,
    canManageBilling,
    notOnPlanLines,
    officerToolsExplanation,
    settingsWords,
  } from '../../lib/guild/settings-view';
  import { PRIMARY_BUTTON_FIXED, SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { BUSY_CLASS } from '../../lib/ui/busy';

  let {
    billing,
    mayBill,
    guildId,
    busy = false,
    unavailable = false,
    onManageBilling = () => {},
  }: {
    billing: GuildBillingView | null;
    /** Whether this account holds the claim: only then do the price line and buttons draw. */
    mayBill: boolean;
    guildId: number | null;
    busy?: boolean;
    unavailable?: boolean;
    onManageBilling?: () => void;
  } = $props();

  const state = $derived(billingState(billing, unavailable));
  const checkoutHref = $derived(`/premium/checkout?plan=guild&interval=monthly&guild_id=${guildId}`);
</script>

{#snippet subscribe()}
  {#if guildId !== null}
    <a href={checkoutHref} class="{PRIMARY_BUTTON_FIXED} w-fit px-4" data-testid="guild-subscribe-link">
      {billingBlockCopy.subscribeTheGuild}
    </a>
  {/if}
{/snippet}

<section id="billing" class="flex scroll-mt-20 flex-col gap-3" data-testid="guild-settings-billing">
  <h2 class="section-title text-[18px]">Billing</h2>
  {#if billing === null}
    <p class="text-[14px]" data-testid="guild-billing-not-on-plan">{notOnPlanLines().headline}</p>
    <p class="text-muted text-[13px]" data-testid="guild-billing-price">{notOnPlanLines().detail}</p>
    {@render subscribe()}
  {:else}
    <p class="text-[14px]" data-testid="guild-billing-plan">{guildPlanLine(billing)}</p>
    <p class="text-muted text-[13px]" data-testid="guild-billing-contact">
      {state === 'officer-tools' ? officerToolsExplanation(billing) : billedLine(billing)}
    </p>
    {#if mayBill && state === 'officer-tools'}
      <p class="text-muted text-[13px]" data-testid="guild-billing-price">{notOnPlanLines().detail}</p>
      {@render subscribe()}
    {:else if mayBill}
      {#if billing.status === 'past_due'}
        <p class="text-strong text-[13px]" role="alert">{billingBlockCopy.pastDueBanner}</p>
      {/if}
      {#if state === 'unavailable'}
        <p class="text-[14px]" role="alert" data-testid="guild-billing-unavailable">
          {settingsWords.billingUnavailable}
        </p>
        <button
          class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-4 ${busy ? BUSY_CLASS : ''}`}
          onclick={onManageBilling}
          disabled={busy}
          data-testid="guild-billing-retry"
        >
          {settingsWords.tryAgain}
        </button>
      {:else if canManageBilling(billing)}
        <button
          class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-4 ${busy ? BUSY_CLASS : ''}`}
          onclick={onManageBilling}
          disabled={busy}
          aria-busy={busy}
          data-testid="guild-manage-billing"
        >
          {billingBlockCopy.manageBilling}
        </button>
      {/if}
    {/if}
  {/if}
  {#if !mayBill}
    <p class="text-muted text-[13px]" data-testid="guild-settings-master-only">
      {guildRosterModelCopy.guildMasterOnly}
    </p>
  {/if}
</section>
