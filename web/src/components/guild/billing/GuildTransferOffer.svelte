<!-- web/src/components/guild/billing/GuildTransferOffer.svelte -->
<!-- The receiving officer's block, at the top of #billing (spec 3.2.6): a 2 px gold left rule, the
     offer sentence, Take over the plan (to Checkout in transfer mode) and Not now (declines). -->
<script lang="ts">
  import { startCheckout } from '../../../lib/billing/api';
  import type { GuildPlanView, TransferView } from '../../../lib/billing/contract';
  import { guildBillingCopy } from '../../../lib/billing/guild-billing-copy';
  import { declineTransfer } from '../../../lib/billing/guild-billing-api';
  import { seatsCopy } from '../../../lib/billing/seats-copy';
  import { BILLING_TEXT_BUTTON } from '../../../lib/guild/billing-styles';
  import { PRIMARY_BUTTON_FIXED } from '../../../lib/planner/styles';
  import { BUSY_CLASS } from '../../../lib/ui/busy';

  const copy = seatsCopy.receiver;

  let {
    guildId,
    plan,
    transfer,
    onChanged,
  }: { guildId: number; plan: GuildPlanView; transfer: TransferView; onChanged: () => void } = $props();

  let opening = $state(false);
  let declining = $state(false);
  let error = $state('');

  const interval = $derived(plan.interval ?? transfer.interval);
  const sentence = $derived(
    transfer.starts_at !== null && transfer.amount_cents !== null && interval !== null
      ? copy.offer(transfer.from_name, transfer.starts_at, transfer.amount_cents, interval, transfer.founder)
      : '',
  );

  async function take(): Promise<void> {
    if (interval === null) return;
    opening = true;
    error = '';
    try {
      const result = await startCheckout({ plan: plan.tier, interval, guildId, intent: 'transfer' });
      window.location.assign(result.checkout_url);
    } catch {
      error = guildBillingCopy.chooser.failed;
      opening = false;
    }
  }

  async function notNow(): Promise<void> {
    declining = true;
    error = '';
    try {
      await declineTransfer(guildId);
      onChanged();
    } catch {
      error = seatsCopy.picker.saveFailed;
    } finally {
      declining = false;
    }
  }
</script>

<div
  class="border-gold flex flex-col gap-3 border-l-2 pl-4"
  role="group"
  aria-label={seatsCopy.whoPays.heading}
  data-testid="guild-transfer-offer"
>
  {#if sentence !== ''}
    <p class="text-[14px]" data-testid="guild-transfer-offer-text">{sentence}</p>
  {/if}
  {#if error !== ''}
    <p class="text-[13px]" role="alert" data-testid="guild-transfer-offer-error">{error}</p>
  {/if}
  <div class="flex flex-wrap items-center gap-3">
    <button
      type="button"
      class="{PRIMARY_BUTTON_FIXED} px-4 {opening ? BUSY_CLASS : ''}"
      disabled={opening || declining}
      aria-busy={opening}
      onclick={() => void take()}
      data-testid="guild-transfer-take"
    >
      {opening ? copy.opening : copy.take}
    </button>
    <button
      type="button"
      class="{BILLING_TEXT_BUTTON} {declining ? BUSY_CLASS : ''}"
      disabled={opening || declining}
      onclick={() => void notNow()}
      data-testid="guild-transfer-not-now"
    >
      {copy.notNow}
    </button>
  </div>
</div>
