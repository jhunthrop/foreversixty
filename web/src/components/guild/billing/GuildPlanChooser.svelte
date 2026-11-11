<!-- web/src/components/guild/billing/GuildPlanChooser.svelte -->
<!-- The chooser (spec 3.2.2): Yearly/Monthly, the Guild and Raid guild cards (the one that fits the
     verified count preselected), the count line, the one gold button to Checkout, then the trial
     and founder lines. Shown for no plan, a lapsed plan and the scoped officer-tools grant. -->
<script lang="ts">
  import { button } from '../../../lib/ui/button';
  import { guildBillingCopy as copy } from '../../../lib/billing/guild-billing-copy';
  import type { GuildTier } from '../../../lib/billing/contract';
  import {
    PLAN_ORDER,
    founderOpen,
    formatUsd,
    INTERVALS,
    intervalUnit,
    type Interval,
  } from '../../../lib/billing/plans';
  import { plansPageCopy } from '../../../lib/billing/plans-page-copy';
  import { listPriceAfterFounder, priceFor } from '../../../lib/billing/plan-costs';
  import { fitsTier } from '../../../lib/guild/billing-state';
  import { PREMIUM_LINKS } from '../../../lib/billing/premium-links';
  import ActionLink from '../../ui/ActionLink.svelte';
  import Segmented from '../../ui/Segmented.svelte';

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
  const intervalOptions = INTERVALS.map((option) => ({
    value: option,
    label: plansPageCopy.interval.labels[option],
  }));
  const buttonText = $derived.by(() => {
    if (lapsed) return copy.chooser.subscribeAgain;
    return trialUsed ? copy.chooser.subscribe(tier, interval, founder) : copy.chooser.startTrial;
  });
</script>

<div class="flex max-w-[720px] flex-col items-start gap-3" data-testid="guild-billing-chooser">
  <Segmented
    options={intervalOptions}
    value={interval}
    label={plansPageCopy.interval.groupLabel}
    onChange={(next) => (interval = next)}
    testid="guild-billing-interval"
  />
  <!-- The two plans as open rows, a radio each (no tiles): name and price on one line, its line under. -->
  <div role="radiogroup" aria-label="Plan" class="flex w-full flex-col gap-1">
    {#each TIERS as key (key)}
      <label
        class="has-[:focus-visible]:outline-gold flex cursor-pointer items-start gap-3 py-2 has-[:focus-visible]:outline-2"
        data-testid={`guild-billing-option-${key}`}
      >
        <input
          type="radio"
          class="accent-gold mt-[3px] size-4 shrink-0"
          name="guild-billing-tier"
          value={key}
          checked={tier === key}
          onchange={() => (picked = key)}
        />
        <span class="flex min-w-0 flex-col gap-[2px]">
          <span class="flex flex-wrap items-baseline gap-x-3">
            <span class="text-strong text-[15px] leading-5 font-bold">{copy.chooser.tierName(key)}</span>
            <span class="text-muted text-[13px]" data-testid={`guild-billing-option-${key}-price`}>
              <span class="text-strong text-[15px] font-bold tabular-nums"
                >{formatUsd(priceFor(key, interval, founder))}</span
              ><span>{intervalUnit(interval)}</span>
            </span>
            {#if founder}
              <span class="text-muted text-[13px]" data-testid={`guild-billing-option-${key}-list`}
                >{listPriceAfterFounder(key, interval)}</span
              >
            {/if}
          </span>
          <span class="text-muted text-[13px]">{copy.chooser.tierLine(key)}</span>
        </span>
      </label>
    {/each}
  </div>
  <p class="text-muted text-[13px]" data-testid="guild-billing-count">
    {copy.chooser.verifiedToday(verifiedCount)}
  </p>
  <ActionLink href={PREMIUM_LINKS.guildLeaders} class="w-fit" data-testid="guild-billing-compare"
    >{copy.chooser.compare}</ActionLink
  >
  <a {href} class="{button('primary')} w-fit" data-testid="guild-billing-subscribe">{buttonText}</a>
  {#if !trialUsed && !lapsed}
    <p class="text-muted text-[13px]" data-testid="guild-billing-trial-line">{copy.chooser.trialLine}</p>
  {/if}
  {#if founder}
    <p class="text-muted text-[13px]" data-testid="guild-billing-founder-line">{copy.chooser.founderLine}</p>
  {/if}
</div>
