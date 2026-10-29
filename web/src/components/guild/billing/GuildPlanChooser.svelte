<!-- web/src/components/guild/billing/GuildPlanChooser.svelte -->
<!-- The chooser (spec 3.2.2): Yearly/Monthly, the Guild and Raid guild cards (the one that fits the
     verified count preselected), the count line, the one gold button to Checkout, then the trial
     and founder lines. Shown for no plan, a lapsed plan and the scoped officer-tools grant. -->
<script lang="ts">
  import { guildBillingCopy as copy } from '../../../lib/billing/guild-billing-copy';
  import type { GuildTier } from '../../../lib/billing/contract';
  import {
    PLAN_ORDER,
    founderOpen,
    formatAmount,
    intervalUnit,
    type Interval,
  } from '../../../lib/billing/plans';
  import { listPriceAfterFounder, priceFor } from '../../../lib/billing/plan-costs';
  import { fitsTier } from '../../../lib/guild/billing-state';
  import { PRIMARY_BUTTON_FIXED } from '../../../lib/planner/styles';
  import IntervalSegments from './IntervalSegments.svelte';

  let {
    guildId,
    verifiedCount,
    trialUsed,
    lapsed,
    now,
    initialTier = null,
    initialInterval = 'yearly',
  }: {
    guildId: number;
    verifiedCount: number;
    trialUsed: boolean;
    lapsed: boolean;
    now: Date;
    /** `?plan` off the settings link (spec 2026-10-27 §6.1): the card that starts picked. */
    initialTier?: GuildTier | null;
    /** `?interval` off the same link. */
    initialInterval?: Interval;
  } = $props();

  const TIERS: readonly GuildTier[] = PLAN_ORDER.filter(
    (key): key is GuildTier => key === 'guild' || key === 'raid_guild',
  );

  // svelte-ignore state_referenced_locally
  let interval = $state<Interval>(initialInterval);
  // svelte-ignore state_referenced_locally
  let picked = $state<GuildTier | null>(initialTier);
  const tier = $derived(picked ?? fitsTier(verifiedCount));
  const founder = $derived(founderOpen(now));
  const href = $derived(`/premium/checkout?plan=${tier}&interval=${interval}&guild_id=${guildId}`);
  const buttonText = $derived.by(() => {
    if (lapsed) return copy.chooser.subscribeAgain;
    return trialUsed ? copy.chooser.subscribe(tier, interval, founder) : copy.chooser.startTrial;
  });
</script>

<div class="flex max-w-[720px] flex-col gap-3" data-testid="guild-billing-chooser">
  <IntervalSegments value={interval} name="guild-billing-interval" onChange={(next) => (interval = next)} />
  <div role="radiogroup" aria-label="Plan" class="grid gap-3 md:grid-cols-2">
    {#each TIERS as key (key)}
      <label
        class="border-line has-[:checked]:border-gold has-[:checked]:bg-card-top rounded-control has-[:focus-visible]:outline-gold flex min-h-24 cursor-pointer flex-col gap-1 border p-4 has-[:focus-visible]:outline-2"
        data-testid={`guild-billing-option-${key}`}
      >
        <input
          type="radio"
          class="sr-only"
          name="guild-billing-tier"
          value={key}
          checked={tier === key}
          onchange={() => (picked = key)}
        />
        <span class="section-title">{copy.chooser.tierName(key)}</span>
        <!-- Only the figures are mono (spec 3.2.2); the unit is 13 px muted prose and the founder
             list price is plain text below. -->
        <span class="text-muted text-[13px]" data-testid={`guild-billing-option-${key}-price`}>
          <span class="text-strong numeric text-[20px]">{formatAmount(priceFor(key, interval, founder))}</span
          ><span>{intervalUnit(interval)}</span>
        </span>
        {#if founder}
          <span class="text-muted text-[13px]" data-testid={`guild-billing-option-${key}-list`}
            >{listPriceAfterFounder(key, interval)}</span
          >
        {/if}
        <span class="text-muted text-[13px]">{copy.chooser.tierLine(key)}</span>
      </label>
    {/each}
  </div>
  <p class="text-muted text-[13px]" data-testid="guild-billing-count">
    {copy.chooser.verifiedToday(verifiedCount)}
  </p>
  <a {href} class="{PRIMARY_BUTTON_FIXED} w-fit px-4" data-testid="guild-billing-subscribe">{buttonText}</a>
  {#if !trialUsed && !lapsed}
    <p class="text-muted text-[13px]" data-testid="guild-billing-trial-line">{copy.chooser.trialLine}</p>
  {/if}
  {#if founder}
    <p class="text-muted text-[13px]" data-testid="guild-billing-founder-line">{copy.chooser.founderLine}</p>
  {/if}
</div>
