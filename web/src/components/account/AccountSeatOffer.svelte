<!-- web/src/components/account/AccountSeatOffer.svelte -->
<!-- A guild's offer of a Premium seat to an account that pays for its own Premium. First in the
     Plan panel; it focuses nothing on load, and after Accept or Decline it is a status line.
     The calls are injectable so the states test without the network. -->
<script lang="ts">
  import { accountPlanCopy } from '../../lib/billing/account-plan-copy';
  import { acceptSeatOffer, declineSeatOffer } from '../../lib/billing/seats-api';
  import type { MeSeatOffer } from '../../lib/billing/contract';
  import { PRIMARY_BUTTON_FIXED, SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';

  import {
    answerOffer,
    offerStatusLine,
    type OfferCall,
    type OfferOutcome,
  } from '../../lib/account/seat-offer-flow';

  let {
    offer,
    zone,
    onAccepted,
    accept = acceptSeatOffer,
    decline = declineSeatOffer,
  }: {
    offer: MeSeatOffer;
    zone?: string;
    /** Called once the API took the accept, so the page can show the new plan line. */
    onAccepted?: () => void;
    accept?: OfferCall;
    decline?: OfferCall;
  } = $props();

  let outcome = $state<OfferOutcome>('idle');
  const busy = $derived(outcome === 'accepting' || outcome === 'declining');
  const copy = accountPlanCopy.offer;

  async function respond(
    call: OfferCall,
    pending: 'accepting' | 'declining',
    done: 'accepted' | 'declined',
  ): Promise<void> {
    outcome = pending;
    outcome = await answerOffer(call, offer.guild_id, done);
    if (outcome === 'accepted') onAccepted?.();
  }

  const statusLine = $derived(offerStatusLine(outcome, offer, zone));
</script>

<div class="flex flex-col gap-2" data-testid="account-seat-offer">
  {#if statusLine !== ''}
    <p class="text-[14px]" role="status" data-testid="account-seat-offer-status">{statusLine}</p>
  {:else}
    <p class="text-strong text-[14px] font-bold">{copy.title(offer.guild_name)}</p>
    <p class="text-muted text-[13px]">{copy.from(offer.from_name, offer.offered_at, zone)}</p>
    {#if offer.own_ends_at !== null && offer.own_amount_cents !== null && offer.own_interval !== null}
      <p class="text-[14px]">
        {copy.consequence(offer.own_ends_at, offer.own_amount_cents, offer.own_interval, zone)}
      </p>
    {/if}
    <div class="flex flex-wrap gap-3">
      <button
        class="{PRIMARY_BUTTON_FIXED} px-4"
        type="button"
        disabled={busy}
        onclick={() => void respond(accept, 'accepting', 'accepted')}
        data-testid="account-seat-offer-accept"
      >
        {outcome === 'accepting' ? copy.accepting : copy.accept}
      </button>
      <button
        class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-4"
        type="button"
        disabled={busy}
        onclick={() => void respond(decline, 'declining', 'declined')}
        data-testid="account-seat-offer-decline"
      >
        {copy.decline}
      </button>
    </div>
    {#if outcome === 'failed'}
      <p class="text-strong text-[13px]" role="alert" data-testid="account-seat-offer-error">
        {copy.failed}
      </p>
    {/if}
  {/if}
</div>
