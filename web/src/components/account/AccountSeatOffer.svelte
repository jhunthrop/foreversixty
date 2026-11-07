<!-- web/src/components/account/AccountSeatOffer.svelte -->
<!-- A guild's offer of a Premium seat to an account that pays for its own Premium (premium spec
     2026-11-08 section 7.3): one selected Row in the Plan section (the gold ground and bar) with the
     offer, who sent it and what accepting does to the account's own plan, Accept the seat and
     Decline in its foot. It focuses nothing on load; once answered it is one status line and its
     buttons never return. The calls are injectable. -->
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
  import { acceptSeatOffer, declineSeatOffer } from '../../lib/billing/seats-api';
  import { button } from '../../lib/ui/button';
  import Row from '../ui/Row.svelte';

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

  const consequence = $derived(
    offer.own_ends_at !== null && offer.own_amount_cents !== null && offer.own_interval !== null
      ? copy.consequence(offer.own_ends_at, offer.own_amount_cents, offer.own_interval, zone)
      : null,
  );
</script>

{#if answered}
  <p class="text-small text-muted" role="status" data-testid="account-seat-offer-status">
    {offerStatusLine(outcome, offer, zone)}
  </p>
{:else}
  <div data-testid="account-seat-offer">
    <Row title={copy.title(offer.guild_name)} selected>
      {#snippet subContent()}
        {copy.from(offer.from_name, offer.offered_at, zone)}
        {#if consequence !== null}<br /><span class="text-text text-[15px] leading-6">{consequence}</span
          >{/if}
      {/snippet}
      {#snippet foot()}
        <button
          type="button"
          class={button('primary')}
          disabled={busy}
          aria-busy={outcome === 'accepting' ? 'true' : undefined}
          onclick={() => void respond(accept, 'accepting', 'accepted')}
          data-testid="account-seat-offer-accept"
          >{outcome === 'accepting' ? copy.accepting : copy.accept}</button
        >
        <button
          type="button"
          class={button('secondary')}
          disabled={busy}
          onclick={() => void respond(decline, 'declining', 'declined')}
          data-testid="account-seat-offer-decline">{copy.decline}</button
        >
      {/snippet}
    </Row>
    {#if outcome === 'failed'}
      <p class="text-wipe pt-2 text-[13px]" role="alert" data-testid="account-seat-offer-error">
        {copy.failed}
      </p>
    {/if}
  </div>
{/if}
