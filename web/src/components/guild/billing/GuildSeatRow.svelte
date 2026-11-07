<!-- web/src/components/guild/billing/GuildSeatRow.svelte -->
<!-- One seat (spec 3.2.5), an open Row list item: Held, Offered (waiting), Accepted, Empty or
     Returned. The crest leads, the name (class colour) and the seat's lines are the body, the trail is
     Move seat, Withdraw or Give seat, and only a reader who may manage seats gets it. Give and Move
     open the picker in the row's foot; a locked Move seat is aria-disabled and its date is text. -->
<script lang="ts">
  import { tick } from 'svelte';
  import type { SeatCandidate, SeatView } from '../../../lib/billing/contract';
  import { giveSeat, moveSeat, seatCandidates, withdrawSeat } from '../../../lib/billing/seats-api';
  import { seatsCopy as copy } from '../../../lib/billing/seats-copy';
  import { smallButton } from '../../../lib/account/small-button';
  import { isSeatLocked, seatErrorLine, seatRowKind, staleReadRefusal } from '../../../lib/guild/seat-view';
  import { BUSY_CLASS } from '../../../lib/ui/busy';
  import { classTone } from '../../../lib/ui/primitives';
  import ClassCrestRing from '../../character/ClassCrestRing.svelte';
  import Row from '../../ui/Row.svelte';
  import Tag from '../../ui/Tag.svelte';
  import GuildSeatPicker from './GuildSeatPicker.svelte';

  let {
    guildId,
    seat,
    manage,
    now,
    onChanged,
  }: {
    guildId: number;
    seat: SeatView;
    manage: boolean;
    now: Date;
    /** Resolves once the billing read has been reloaded. */
    onChanged: () => Promise<void>;
  } = $props();

  let picking = $state<'give' | 'move' | null>(null);
  let candidates = $state<readonly SeatCandidate[] | null>(null);
  let busy = $state(false);
  let error = $state('');

  const kind = $derived(seatRowKind(seat));
  const locked = $derived(isSeatLocked(seat, now));
  const noteId = $derived(`guild-seat-${seat.slot}-note`);
  const lockId = $derived(`guild-seat-${seat.slot}-lock`);
  const holderName = $derived(seat.holder?.name ?? '');
  const middle = $derived.by(() => {
    if (kind === 'pending' && seat.offeree !== null) {
      return seat.holder === null
        ? copy.offeredOn(seat.offeree.offered_at)
        : copy.movingTo(seat.offeree.name, seat.offeree.offered_at);
    }
    if (kind === 'offered' && seat.given_at !== null) return copy.offeredOn(seat.given_at);
    if (kind === 'accepted' && seat.starts_at !== null) return copy.accepted(seat.starts_at);
    if (kind === 'returned' && seat.returned !== null) {
      return copy.freedSince(seat.returned.at, seat.returned.name);
    }
    if (kind === 'held') {
      return locked && seat.movable_at !== null ? copy.canMoveFrom(seat.movable_at) : copy.canMoveNow;
    }
    return '';
  });
  /** The middle of an accepted seat is not its lock; the lock date follows on a line of its own. An
   *  offer starts no lock (R-P3), so an offered or pending row never shows one. */
  const lockLine = $derived(
    kind === 'accepted' && locked && seat.movable_at !== null ? copy.canMoveFrom(seat.movable_at) : '',
  );
  /** A waiting offer has only Withdraw: the API refuses a move with 409 offer_pending. */
  const canMove = $derived(seat.holder !== null && kind !== 'offered' && kind !== 'pending');
  const moveDisabled = $derived(locked);
  const canWithdraw = $derived(kind === 'offered' || kind === 'pending');
  /** A held seat moving to a payer: line 1 stays the holder, the Waiting pill sits in the middle. */
  const moving = $derived(kind === 'pending' && seat.holder !== null);
  /** The holder, else the member a seat is offered to; null for an empty or returned seat. */
  const who = $derived(seat.holder ?? seat.offeree);
  /** The Waiting tag sits under the name of an offered seat and of an empty seat's offeree. */
  const waitingUnderName = $derived(kind === 'offered' || (seat.holder === null && seat.offeree !== null));

  /** The seat's own Row: the list item the testid names. */
  const rowElement = (): HTMLElement | null =>
    document.querySelector<HTMLElement>(`[data-testid="guild-seat-${seat.slot}"]`);

  /** Puts focus back on the row's action once the opener it came from has re-rendered. */
  async function refocusAction(): Promise<void> {
    await tick();
    rowElement()?.querySelector<HTMLElement>('[data-seat-action]')?.focus();
  }

  async function open(mode: 'give' | 'move'): Promise<void> {
    if (mode === 'move' && moveDisabled) return;
    picking = mode;
    error = '';
    candidates = null;
    try {
      candidates = await seatCandidates(guildId);
    } catch {
      picking = null;
      error = copy.picker.saveFailed;
    }
  }

  function close(): void {
    picking = null;
    error = '';
    void refocusAction();
  }

  /** A refused submit keeps the picker open: focus goes back to its button, which `busy` disabled
   *  (and so blurred), so the alert is read in place and Escape still acts as Cancel. */
  async function refocusSubmit(): Promise<void> {
    await tick();
    rowElement()?.querySelector<HTMLElement>('[data-seat-submit]')?.focus();
  }

  async function submit(candidate: SeatCandidate): Promise<void> {
    if (picking === null) return;
    busy = true;
    error = '';
    let refused = false;
    try {
      const send = picking === 'give' ? giveSeat : moveSeat;
      await send(guildId, seat.slot, candidate.user_id);
      picking = null;
      await onChanged();
      await refocusAction();
    } catch (failure) {
      refused = true;
      error = seatErrorLine(failure, candidate.name, seat.movable_at);
      if (staleReadRefusal(failure)) await onChanged();
    } finally {
      busy = false;
    }
    if (refused) await refocusSubmit();
  }

  async function withdraw(): Promise<void> {
    busy = true;
    error = '';
    try {
      await withdrawSeat(guildId, seat.slot);
      await onChanged();
      await refocusAction();
    } catch {
      error = copy.picker.saveFailed;
    } finally {
      busy = false;
    }
  }
</script>

<Row
  title={who?.name ?? copy.notGiven}
  titleTone={who === null ? undefined : classTone(who.class)}
  item
  open
  testid={`guild-seat-${seat.slot}`}
>
  {#snippet lead()}
    <ClassCrestRing characterClass={who?.class ?? ''} size={32} />
  {/snippet}
  {#snippet titleContent()}
    {#if who === null}<span class="text-muted font-normal">{copy.notGiven}</span>{:else}<span
        data-testid={`guild-seat-${seat.slot}-name`}>{who.name}</span
      >{/if}
  {/snippet}
  {#snippet subContent()}
    {#if waitingUnderName}
      <span class="block pt-1"><Tag tone="muted">{copy.waiting}</Tag></span>
    {:else if seat.given_at !== null && seat.holder !== null && (kind === 'held' || moving)}
      <span class="block">{copy.givenOn(seat.given_at)}</span>
    {/if}
    {#if middle !== ''}
      <span
        id={noteId}
        class="flex flex-wrap items-center gap-2 {kind === 'held' && !locked ? 'text-text' : ''}"
        data-testid={`guild-seat-${seat.slot}-note`}
      >
        {middle}{#if moving}<Tag tone="muted">{copy.waiting}</Tag>{/if}
      </span>
    {/if}
    {#if lockLine !== ''}
      <span id={lockId} class="block" data-testid={lockId}>{lockLine}</span>
    {/if}
  {/snippet}
  {#snippet trail()}
    {#if manage}
      {#if canMove}
        <button
          type="button"
          class={smallButton()}
          aria-disabled={moveDisabled}
          aria-expanded={picking === 'move'}
          aria-describedby={moveDisabled ? (lockLine !== '' ? lockId : noteId) : undefined}
          onclick={() => void open('move')}
          data-seat-action
          data-testid={`guild-seat-${seat.slot}-move`}
        >
          {copy.move}
        </button>
      {/if}
      {#if canWithdraw}
        <button
          type="button"
          class="{smallButton('ghost')} {busy ? BUSY_CLASS : ''}"
          disabled={busy}
          aria-busy={busy}
          onclick={() => void withdraw()}
          data-seat-action={canMove ? undefined : ''}
          data-testid={`guild-seat-${seat.slot}-withdraw`}
        >
          {busy ? copy.withdrawing : copy.withdraw}
        </button>
      {:else if !canMove}
        <button
          type="button"
          class={smallButton()}
          aria-expanded={picking === 'give'}
          onclick={() => void open('give')}
          data-seat-action
          data-testid={`guild-seat-${seat.slot}-give`}
        >
          {copy.give}
        </button>
      {/if}
    {/if}
  {/snippet}
  {#snippet foot()}
    {#if picking === null && error !== ''}
      <p class="text-wipe text-[13px]" role="alert" data-testid={`guild-seat-${seat.slot}-error`}>{error}</p>
    {/if}
    {#if picking !== null}
      <GuildSeatPicker
        slot={seat.slot}
        mode={picking}
        {holderName}
        {candidates}
        {busy}
        {error}
        onSubmit={(candidate) => void submit(candidate)}
        onCancel={close}
      />
    {/if}
  {/snippet}
</Row>
