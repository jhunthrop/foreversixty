<!-- web/src/components/checkout/CheckoutConfirmed.svelte -->
<!-- Back from Stripe with Premium on (redesign review 2026-11-09, replacing the finished three-step
     list): one framed panel, `Premium is on.` and the renewal line, then what Premium turned on as
     rows that go there: the coach, the server sims, and the plan itself. Utility classes only: the
     launcher draws this after hydration, so a scoped style would never reach the page's CSS. -->
<script lang="ts">
  import { checkoutCopy } from '../../lib/billing/copy';
  import { plansPageCopy } from '../../lib/billing/plans-page-copy';
  import { premiumSourceCopy } from '../../lib/billing/premium-source-copy';
  import Plate from '../ui/Plate.svelte';
  import Row from '../ui/Row.svelte';

  let { planName, renewal }: { planName: string; renewal: string | null } = $props();

  const [coach, sims] = plansPageCopy.premium.features;
  const NEXT = [
    { title: checkoutCopy.nextCoach, sub: coach?.text ?? '', href: '/logs' },
    { title: checkoutCopy.nextSim, sub: sims?.text ?? '', href: '/sim' },
    { title: premiumSourceCopy.buttons.seePlan, sub: '', href: '/account#plan' },
  ];
</script>

<Plate label="Checkout" testid="checkout-plate" frame="framed">
  <div class="flex flex-col gap-1 px-6 pt-5 pb-3">
    <p
      class="text-strong m-0 [font-family:var(--font-heading)] text-[24px] leading-[30px] font-normal tracking-[var(--tracking-heading)]"
      role="status"
      data-testid="checkout-message"
    >
      {checkoutCopy.planOn(planName)}
    </p>
    {#if renewal !== null}<p class="text-small text-muted m-0" data-testid="checkout-renewal">
        {renewal}
      </p>{/if}
  </div>
  <nav class="flex flex-col px-2 pb-3" aria-label={checkoutCopy.nextLabel} data-testid="checkout-next">
    {#each NEXT as row (row.href)}
      <Row title={row.title} sub={row.sub || undefined} href={row.href} />
    {/each}
  </nav>
</Plate>
