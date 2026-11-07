<!-- web/src/components/guild/billing/GuildPlanChooser.svelte -->
<!-- The chooser (spec 3.2.2): Yearly/Monthly (the Segmented control), the Guild and Raid guild Rows
     (the one that fits the verified count preselected), the count line, the one gold button to Checkout, then the trial
     and founder lines. Shown for no plan, a lapsed plan and the scoped officer-tools grant. -->
<script lang="ts">
  import { guildBillingCopy as copy } from '../../../lib/billing/guild-billing-copy';
  import type { GuildTier } from '../../../lib/billing/contract';
  import {
    INTERVALS,
    PLAN_ORDER,
    founderOpen,
    formatAmount,
    intervalUnit,
    type Interval,
  } from '../../../lib/billing/plans';
  import { listPriceAfterFounder, priceFor } from '../../../lib/billing/plan-costs';
  import { fitsTier } from '../../../lib/guild/billing-state';
  import { plansPageCopy } from '../../../lib/billing/plans-page-copy';
  import { button } from '../../../lib/ui/button';
  import { PREMIUM_LINKS } from '../../../lib/billing/premium-links';
  import ActionLink from '../../ui/ActionLink.svelte';
  import Figure from '../../ui/Figure.svelte';
  import Row from '../../ui/Row.svelte';
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
  const buttonText = $derived.by(() => {
    if (lapsed) return copy.chooser.subscribeAgain;
    return trialUsed ? copy.chooser.subscribe(tier, interval, founder) : copy.chooser.startTrial;
  });
</script>

<div class="flex max-w-[720px] flex-col gap-3" data-testid="guild-billing-chooser">
  <Segmented
    options={INTERVALS.map((option) => ({ value: option, label: plansPageCopy.interval.labels[option] }))}
    value={interval}
    label={plansPageCopy.interval.groupLabel}
    fill="phone"
    onChange={(next) => (interval = next)}
    testid="guild-billing-interval"
  />
  <!-- One Row per plan: the radio leads, and the name, the lines and the price are its labels, so the
       whole row picks the plan; the picked one is the selected Row (the gold ground and bar). -->
  <div role="radiogroup" aria-label={copy.chooser.planGroup}>
    {#each TIERS as key (key)}
      {@const id = `guild-billing-tier-${key}`}
      <Row title={copy.chooser.tierName(key)} selected={tier === key} testid={`guild-billing-option-${key}`}>
        {#snippet lead()}
          <input
            {id}
            type="radio"
            class="accent-gold focus-visible:outline-gold size-[18px] cursor-pointer focus-visible:outline-2 focus-visible:outline-offset-2"
            name="guild-billing-tier"
            value={key}
            checked={tier === key}
            onchange={() => (picked = key)}
          />
        {/snippet}
        {#snippet titleContent()}<label class="block cursor-pointer" for={id}
            >{copy.chooser.tierName(key)}</label
          >{/snippet}
        {#snippet subContent()}
          <label class="block cursor-pointer" for={id}>
            {copy.chooser.tierLine(key)}
            {#if founder}<span class="block" data-testid={`guild-billing-option-${key}-list`}
                >{listPriceAfterFounder(key, interval)}</span
              >{/if}
          </label>
        {/snippet}
        {#snippet trail()}
          <label class="cursor-pointer" for={id} data-testid={`guild-billing-option-${key}-price`}>
            <Figure
              value={formatAmount(priceFor(key, interval, founder))}
              unit={intervalUnit(interval).trim()}
              size="sm"
            />
          </label>
        {/snippet}
      </Row>
    {/each}
  </div>
  <p class="text-muted text-[13px]" data-testid="guild-billing-count">
    {copy.chooser.verifiedToday(verifiedCount)}
  </p>
  <ActionLink href={PREMIUM_LINKS.guildLeaders} class="w-fit" data-testid="guild-billing-compare"
    >{copy.chooser.compare}</ActionLink
  >
  <a {href} class="{button('primary')} w-full md:w-fit" data-testid="guild-billing-subscribe">{buttonText}</a>
  {#if !trialUsed && !lapsed}
    <p class="text-muted text-[13px]" data-testid="guild-billing-trial-line">{copy.chooser.trialLine}</p>
  {/if}
  {#if founder}
    <p class="text-muted text-[13px]" data-testid="guild-billing-founder-line">{copy.chooser.founderLine}</p>
  {/if}
</div>
