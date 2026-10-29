<!-- web/src/components/checkout/CheckoutBand.svelte -->
<!-- Spec 2026-10-27-subscribe-and-join §5.1: the checkout page's full-bleed band, the /premium
     band's ground with its content in the centred 1344 column. It needs only the URL, so it
     draws on the first client paint, before /v1/me answers. A guild plan carries the plan
     summary (name, price, Founder pill, size line, trial and Change plan); Premium carries its
     h1 only (§4.2, Task 12 fills the body); an invalid link carries the h1 `Checkout`. `plan`
     null without `invalid` is the moment before the query string is read (the static render):
     it draws the Guild summary's frame (h1, size line, trial line) so the LCP text is in the
     static HTML, hidden by CSS unless checkout.astro's head script marked the URL a guild
     checkout (html[data-checkout-step]). Name and price fill in on hydration into reserved
     lines, so nothing shifts. Every state reserves the same box. -->
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
    invalid = false,
  }: {
    /** Null until the launcher has read the query string, and for an invalid link. */
    plan: PlanKey | null;
    interval: Interval;
    /** The founder window is open (founderOpen(now)). */
    founder: boolean;
    /** The URL was judged invalid: the band names the page `Checkout`, with no summary. */
    invalid?: boolean;
  } = $props();

  const PREMIUM_HREF = '/premium';
  const CHANGE_PLAN_LINK = `text-text hover:text-strong focus-visible:outline-gold underline focus-visible:outline-2 focus-visible:outline-offset-2 ${PHONE_HIT_AREA}`;

  const pending = $derived(plan === null && !invalid);
  const isGuildPlan = $derived(plan !== null && PLANS[plan].guild);
  const showSummary = $derived(pending || isGuildPlan);
  const title = $derived(
    pending || plan === 'guild' || plan === 'raid_guild'
      ? copy.band.title
      : plan === 'premium'
        ? PLANS.premium.name
        : copy.band.invalidTitle,
  );
  const planName = $derived(plan === null ? '' : PLANS[plan].name);
  const amount = $derived(plan === null ? '' : formatAmount(priceFor(plan, interval, founder)));
  const sizeLine = $derived(plan === 'raid_guild' ? copy.sizeLine.raidGuild : copy.sizeLine.guild);
</script>

<div class="full-bleed bg-raised border-line-soft border-b" data-testid="checkout-band">
  <header
    class="checkout-band-row page-column flex flex-col gap-2 px-[18px] py-3 md:pt-7 md:pb-6 lg:flex-row lg:items-start lg:justify-between lg:gap-6 {CHECKOUT_BAND_RESERVE}"
  >
    <div class="contents" class:checkout-pending={pending}>
      <div class="flex min-w-0 flex-col gap-2">
        <p class="label text-gold hidden md:block">
          <i class="bg-gold mr-2.5 inline-block h-px w-7 align-middle" aria-hidden="true"></i>
          {copy.band.label}
        </p>
        <h1 class="heading-page m-0">{title}</h1>
      </div>
      {#if showSummary}
        <div class="flex flex-col gap-1 text-[13px] lg:max-w-[440px]" data-testid="checkout-plan-summary">
          <p class="m-0 flex min-h-[26px] flex-wrap items-baseline">
            <span class="heading-plan">{planName}</span>
            {#if !pending}
              <!-- One inline box: a flex child would drop the unit's leading space ("$120a year"). -->
              <span class="ml-2" data-testid="checkout-plan-price"
                ><span class="text-strong numeric text-[20px] font-medium">{amount}</span><span
                  class="text-muted text-[14px]">{intervalUnit(interval)}</span
                ></span
              >
              {#if founder}
                <span class="pill pill-site ml-2 self-center">{copy.planSummary.founder}</span>
              {/if}
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
    </div>
  </header>
</div>

<style>
  /* The static render's frame is the Guild checkout's: show it only when the head script read a
     guild URL, so a Premium or invalid link never flashes it. */
  :global(html:not([data-checkout-step='guild'])) .checkout-pending {
    display: none;
  }
  /* plan=premium carries only its h1, so the band hugs it (row 2) instead of keeping the guild
     summary's reserve. The head script marks the URL before first paint, so the static band and
     the hydrated one are the same box: checkout-layout.ts PREMIUM_BAND_PX (phone 12 + h1 34 +
     12; from 768 28 + label 16.5 + 8 + h1 44 + 24). */
  :global(html[data-checkout-band='premium']) .checkout-band-row {
    min-height: 58px;
  }
  @media (min-width: 768px) {
    :global(html[data-checkout-band='premium']) .checkout-band-row {
      min-height: 121px;
    }
  }
</style>
