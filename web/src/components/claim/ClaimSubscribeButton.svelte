<!-- web/src/components/claim/ClaimSubscribeButton.svelte -->
<!-- Spec 2026-10-27 §6.2: after a claim made for a plan, the `ready` button leads `Next`: the
     page's one gold fill. It starts Stripe Checkout for the claimed guild. -->
<script lang="ts">
  import { startCheckout } from '../../lib/billing/api';
  import { checkoutGuildCopy } from '../../lib/billing/checkout-guild-copy';
  import type { ClaimPlan } from '../../lib/guild/claim-plan';
  import { PRIMARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { BUSY_CLASS } from '../../lib/ui/busy';

  let { guildId, claimPlan, label }: { guildId: number; claimPlan: ClaimPlan; label: string } = $props();

  let busy = $state(false);
  let failed = $state(false);

  async function onClick(): Promise<void> {
    busy = true;
    failed = false;
    try {
      const result = await startCheckout({ plan: claimPlan.plan, interval: claimPlan.interval, guildId });
      window.location.assign(result.checkout_url);
    } catch {
      failed = true;
      busy = false;
    }
  }
</script>

<div class="flex flex-col gap-2">
  <button
    class={`${PRIMARY_BUTTON_FIXED} w-full justify-center px-4 md:w-fit ${busy ? BUSY_CLASS : ''}`}
    onclick={() => void onClick()}
    disabled={busy}
    aria-busy={busy}
    data-testid="guild-claim-next-subscribe"
  >
    {busy ? checkoutGuildCopy.busy.checkout : label}
  </button>
  <!-- Drawn only on a failure, after the player's own click: an empty reserve would push the
       next item 40 px down instead of the list's 12 (row 16). -->
  {#if failed}
    <p class="m-0 text-[14px]" role="alert" data-testid="guild-claim-subscribe-error">
      {checkoutGuildCopy.checkoutErrors.generic}
    </p>
  {/if}
</div>
