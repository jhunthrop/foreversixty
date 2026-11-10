<!-- web/src/components/checkout/CheckoutHead.svelte -->
<!-- The checkout page's head (premium spec 2026-11-08 section 6, decision K1): the PageHeader with
     the plan's name (`Premium`, `Subscribe a guild`, `Checkout` for a bad link) and at its right
     the plan picked on /premium: the price in gold with the Founder mark, its second line (a
     guild plan: the size line and the trial), and `Change plan ›` before payment. It needs only
     the URL, so it draws before /v1/me answers; before the URL is read (the static render) the
     title and answer hold their boxes empty. An account that already has Premium (`owned`) keeps
     the answer's box and draws nothing in it: no price for a player who is not paying one. -->
<script lang="ts">
  import { checkoutGuildCopy as copy } from '../../lib/billing/checkout-guild-copy';
  import { PLANS, type Interval, type PlanKey } from '../../lib/billing/plans';
  import { priceCell } from '../premium/model';
  import { LINK } from '../../lib/ui/primitives';
  import Figure from '../ui/Figure.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Tag from '../ui/Tag.svelte';

  let {
    plan,
    interval,
    founder,
    invalid = false,
    showChange = true,
    owned = false,
  }: {
    /** Null until the launcher has read the query string, and for an invalid link. */
    plan: PlanKey | null;
    interval: Interval;
    founder: boolean;
    invalid?: boolean;
    /** `Change plan ›` shows only before payment. */
    showChange?: boolean;
    /** The account already has Premium: the answer keeps its box, empty. */
    owned?: boolean;
  } = $props();

  const pending = $derived(plan === null && !invalid);
  const guild = $derived(plan !== null && PLANS[plan].guild);
  const title = $derived(
    invalid ? copy.band.invalidTitle : plan === null || guild ? copy.band.title : PLANS[plan].name,
  );
  // Before the URL is read the head draws a guild plan's summary hidden (the taller frame), so a
  // guild checkout, which stays on this page, never moves when the launcher fills it in.
  const shown = $derived<PlanKey | null>(plan ?? (pending ? 'guild' : null));
  const cell = $derived(shown === null ? null : priceCell(shown, interval, founder));
</script>

<div class="checkout-head" class:is-pending={pending} class:is-owned={owned} data-testid="checkout-band">
  <PageHeader {title} tone="strong">
    {#snippet answer()}
      {#if cell !== null && shown !== null}
        <div
          class="checkout-plan"
          data-testid="checkout-plan-summary"
          aria-hidden={pending ? 'true' : undefined}
        >
          <span class="checkout-plan-fig">
            <Figure
              value={cell.amount}
              unit={cell.suffix.trim()}
              label={guild || pending ? PLANS[shown].name : undefined}
              tone="gold"
              align="end"
            />
            {#if founder}<Tag>{copy.planSummary.founder}</Tag>{/if}
          </span>
          {#if guild || pending}
            <span class="text-small text-muted" data-testid="checkout-size-line"
              >{shown === 'raid_guild' ? copy.sizeLine.raidGuild : copy.sizeLine.guild}</span
            >
            <span class="text-small text-strong font-bold">{copy.planSummary.trial}</span>
          {:else}
            <span class="text-small text-muted">{cell.secondLine}</span>
          {/if}
          {#if showChange}<a class={LINK} href="/premium" data-testid="checkout-change-plan"
              >{copy.planSummary.changePlan} ›</a
            >{/if}
        </div>
      {/if}
    {/snippet}
  </PageHeader>
</div>

<style>
  .checkout-plan {
    display: flex;
    flex-direction: column;
    gap: 2px;
  }
  @media (min-width: 900px) {
    .checkout-plan {
      min-width: 300px;
    }
  }
  .checkout-plan-fig {
    display: inline-flex;
    align-items: flex-end;
    gap: 8px;
  }
  .is-pending :global(h1),
  .is-pending .checkout-plan,
  .is-owned :global([data-region='head-answer']) {
    visibility: hidden;
  }
  /* checkout.astro's head script marks a guild checkout before first paint: its title, size line
     and trial (the LCP text) show from the static render; only the price waits for the URL read
     (Raid guild's differs). */
  :global(html[data-checkout-step='guild']) .is-pending :global(h1),
  :global(html[data-checkout-step='guild']) .is-pending .checkout-plan {
    visibility: visible;
  }
  :global(html[data-checkout-step='guild']) .is-pending .checkout-plan-fig {
    visibility: hidden;
  }
  @media (min-width: 900px) {
    .checkout-plan {
      align-items: flex-end;
      max-width: 360px;
    }
  }
</style>
