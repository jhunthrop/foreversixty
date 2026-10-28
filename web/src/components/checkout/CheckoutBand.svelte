<!-- web/src/components/checkout/CheckoutBand.svelte -->
<!-- Spec 2026-10-27-subscribe-and-join §5.1: the checkout page's full-bleed band, the /premium
     band's ground with its content in the centred 1344 column. It needs only the URL, so it
     draws on the first client paint, before /v1/me answers. A guild plan carries the plan
     summary (name, price, Founder pill, size line, trial and Change plan); Premium carries its
     h1 only (§4.2, Task 12 fills the body). `plan` null is the moment before the query string
     is read (the static render): the ground and the height reserve only, so nothing shifts
     when the summary arrives. Every plan and the null state reserve the same box. -->
<script lang="ts">
  import { checkoutGuildCopy as copy } from '../../lib/billing/checkout-guild-copy';
  import { CHECKOUT_BAND_RESERVE } from '../../lib/billing/checkout-layout';
  import { priceFor } from '../../lib/billing/plan-costs';
  import { PLANS, formatAmount, intervalUnit, type Interval, type PlanKey } from '../../lib/billing/plans';
  import { PHONE_HIT_AREA } from '../../lib/guild/field-styles';

  let {
    plan,
    interval,
    founder,
  }: {
    /** Null until the launcher has read the query string. */
    plan: PlanKey | null;
    interval: Interval;
    /** The founder window is open (founderOpen(now)). */
    founder: boolean;
  } = $props();

  const PREMIUM_HREF = '/premium';
  const CHANGE_PLAN_LINK = `text-text hover:text-strong focus-visible:outline-gold underline focus-visible:outline-2 focus-visible:outline-offset-2 ${PHONE_HIT_AREA}`;

  const isGuildPlan = $derived(plan !== null && PLANS[plan].guild);
  const title = $derived(plan === 'premium' ? PLANS.premium.name : copy.band.title);
  const amount = $derived(plan === null ? '' : formatAmount(priceFor(plan, interval, founder)));
  const sizeLine = $derived(plan === 'raid_guild' ? copy.sizeLine.raidGuild : copy.sizeLine.guild);
</script>

<div class="full-bleed bg-raised border-line-soft border-b" data-testid="checkout-band">
  <header
    class="page-column flex flex-col gap-2 px-[18px] py-3 md:pt-7 md:pb-6 lg:flex-row lg:items-end lg:justify-between lg:gap-6 {CHECKOUT_BAND_RESERVE}"
  >
    {#if plan !== null}
      <div class="flex min-w-0 flex-col gap-2">
        <p class="label text-gold hidden md:block">
          <i class="bg-gold mr-2.5 inline-block h-px w-7 align-middle" aria-hidden="true"></i>
          {copy.band.label}
        </p>
        <h1 class="font-display text-strong m-0 text-[22px] leading-[1.2] font-bold">{title}</h1>
      </div>
      {#if isGuildPlan}
        <div class="flex flex-col gap-1 text-[13px] lg:max-w-[440px]" data-testid="checkout-plan-summary">
          <p class="m-0 flex flex-wrap items-baseline">
            <span class="font-display text-strong text-[15px] font-bold">{PLANS[plan].name}</span>
            <span class="text-strong ml-2 font-mono text-[20px] font-medium tabular-nums">{amount}</span>
            <span class="text-muted text-[14px]">{intervalUnit(interval)}</span>
            {#if founder}
              <span class="pill pill-site ml-2 self-center">{copy.planSummary.founder}</span>
            {/if}
          </p>
          <p class="text-text m-0" data-testid="checkout-size-line">{sizeLine}</p>
          <p class="text-text m-0">
            <span class="text-strong font-semibold">{copy.planSummary.trial}</span>
            <span aria-hidden="true">&nbsp;·&nbsp;</span>
            <a href={PREMIUM_HREF} class={CHANGE_PLAN_LINK}>{copy.planSummary.changePlan}</a>
          </p>
        </div>
      {/if}
    {/if}
  </header>
</div>
