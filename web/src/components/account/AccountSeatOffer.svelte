<!-- web/src/components/account/AccountSeatOffer.svelte -->
<!-- A guild's offer of a Premium seat to an account that pays for its own Premium: the confirm
     block with a 2px gold left rule, first in the Plan panel. It focuses nothing on load; once
     answered it is a status line and its buttons never return. The calls are injectable. -->
<script lang="ts">
  import {
    answerOffer,
    isAnswered,
    isBusy,
    offerStatusLine,
    type OfferCall,
    type OfferOutcome,
  } from '../../lib/account/seat-offer-flow';
  import { accountPlanCopy } from '../../lib/billing/account-plan-copy';
  import type { MeSeatOffer } from '../../lib/billing/contract';
  import { PRIMARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { acceptSeatOffer, declineSeatOffer } from '../../lib/billing/seats-api';

  let {
    offer,
    zone,
    onAnswered,
    accept = acceptSeatOffer,
    decline = declineSeatOffer,
  }: {
    offer: MeSeatOffer;
    zone?: string;
    /** Called once the API took an answer, so the page can read the account again. */
    onAnswered?: () => void;
    accept?: OfferCall;
    decline?: OfferCall;
  } = $props();

  let outcome = $state<OfferOutcome>('idle');
  const busy = $derived(isBusy(outcome));
  const answered = $derived(isAnswered(outcome));
  const copy = accountPlanCopy.offer;

  async function respond(
    call: OfferCall,
    pending: 'accepting' | 'declining',
    done: 'accepted' | 'declined',
  ): Promise<void> {
    if (busy || answered) return;
    outcome = pending;
    outcome = await answerOffer(call, offer.guild_id, done);
    if (isAnswered(outcome)) onAnswered?.();
  }

  const statusLine = $derived(offerStatusLine(outcome, offer, zone));
</script>

<div class="border-gold flex flex-col gap-2 border-l-2 pl-3" data-testid="account-seat-offer">
  <div role="status" data-testid="account-seat-offer-status">
    {#if answered}<p class="text-muted text-[13px]">{statusLine}</p>{/if}
  </div>
  {#if !answered}
    <p class="text-strong text-[15px] font-bold">{copy.title(offer.guild_name)}</p>
    <p class="text-muted text-[13px]">{copy.from(offer.from_name, offer.offered_at, zone)}</p>
    {#if offer.own_ends_at !== null && offer.own_amount_cents !== null && offer.own_interval !== null}
      <p class="text-[14px]">
        {copy.consequence(offer.own_ends_at, offer.own_amount_cents, offer.own_interval, zone)}
      </p>
    {/if}
    <div class="flex flex-wrap items-center gap-4">
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
        class="text-text min-h-11 text-[14px] underline disabled:opacity-60"
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
