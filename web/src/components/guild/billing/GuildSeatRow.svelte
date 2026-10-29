<!-- web/src/components/guild/billing/GuildSeatRow.svelte -->
<!-- One seat (spec 3.2.5): Held, Offered (waiting), Accepted, Empty or Returned. The right column is
     Move seat, Withdraw or Give seat, and only a reader who may manage seats gets it. Give and Move
     open the picker in place below the row; a locked Move seat is aria-disabled and its date is text. -->
<script lang="ts">
  import { tick } from 'svelte';
  import type { SeatCandidate, SeatView } from '../../../lib/billing/contract';
  import { giveSeat, moveSeat, seatCandidates, withdrawSeat } from '../../../lib/billing/seats-api';
  import { seatsCopy as copy } from '../../../lib/billing/seats-copy';
  import {
    BILLING_MOVE_BUTTON,
    BILLING_OUTLINE_BUTTON,
    BILLING_TEXT_BUTTON,
    OPENER_PRESSED,
  } from '../../../lib/guild/billing-styles';
  import { isSeatLocked, seatErrorLine, seatRowKind, staleReadRefusal } from '../../../lib/guild/seat-view';
  import { BUSY_CLASS } from '../../../lib/ui/busy';
  import ClassCrestRing from '../../character/ClassCrestRing.svelte';
  import { classColorVar } from '../../../lib/report/format';
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
  let rowEl = $state<HTMLElement | null>(null);

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

  /** Puts focus back on the row's action once the opener it came from has re-rendered. */
  async function refocusAction(): Promise<void> {
    await tick();
    rowEl?.querySelector<HTMLElement>('[data-seat-action]')?.focus();
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

  function onKeydown(event: KeyboardEvent): void {
    if (event.key !== 'Escape' || picking === null) return;
    event.preventDefault();
    close();
  }

  /** A refused submit keeps the picker open: focus goes back to its button, which `busy` disabled
   *  (and so blurred), so the alert is read in place and Escape still acts as Cancel. */
  async function refocusSubmit(): Promise<void> {
    await tick();
    rowEl?.querySelector<HTMLElement>('[data-seat-submit]')?.focus();
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

<li
  bind:this={rowEl}
  onkeydown={onKeydown}
  class="border-line-soft flex flex-col gap-3 border-b py-3 last:border-b-0"
  data-testid={`guild-seat-${seat.slot}`}
  data-seat-kind={kind}
>
  <div
    class="grid items-center gap-x-4 gap-y-2 md:min-h-12 {manage
      ? 'grid-cols-[minmax(0,1fr)_auto] md:grid-cols-[minmax(220px,1fr)_minmax(0,1.4fr)_auto]'
      : 'grid-cols-1 md:grid-cols-[minmax(220px,1fr)_minmax(0,1.4fr)]'}"
    data-testid={`guild-seat-${seat.slot}-grid`}
  >
    <!-- Below 768 the row is two lines (spec 7): who and the button on line 1, the middle on line 2. -->
    <div class="col-start-1 row-start-1 flex min-w-0 items-center gap-3">
      {#if seat.holder !== null}
        <ClassCrestRing characterClass={seat.holder.class} size={32} />
        <div class="flex min-w-0 flex-col">
          <span
            class="heading-name truncate text-[14px]"
            style={`color: ${classColorVar(seat.holder.class)}`}
            data-testid={`guild-seat-${seat.slot}-name`}>{seat.holder.name}</span
          >
          {#if kind === 'offered'}
            <span><span class="pill pill-sample">{copy.waiting}</span></span>
          {:else if seat.given_at !== null && (kind === 'held' || moving)}
            <span class="text-muted text-[12px]">{copy.givenOn(seat.given_at)}</span>
          {/if}
        </div>
      {:else if seat.offeree !== null}
        <ClassCrestRing characterClass={seat.offeree.class} size={32} />
        <div class="flex min-w-0 flex-col">
          <span
            class="heading-name truncate text-[14px]"
            style={`color: ${classColorVar(seat.offeree.class)}`}
            data-testid={`guild-seat-${seat.slot}-name`}>{seat.offeree.name}</span
          >
          <span><span class="pill pill-sample">{copy.waiting}</span></span>
        </div>
      {:else}
        <span
          class="border-line inline-block size-8 shrink-0 rounded-full border border-dashed"
          aria-hidden="true"
        ></span>
        <span class="text-muted text-[14px]">{copy.notGiven}</span>
      {/if}
    </div>
    <div
      class="col-span-full row-start-2 flex flex-col gap-1 text-[13px] md:col-span-1 md:col-start-2 md:row-start-1"
    >
      {#if middle !== ''}
        <p
          id={noteId}
          class="{kind === 'held' && !locked ? 'text-text' : 'text-muted'} {moving
            ? 'flex flex-wrap items-center gap-2'
            : ''}"
          data-testid={`guild-seat-${seat.slot}-note`}
        >
          {middle}{#if moving}<span class="pill pill-sample">{copy.waiting}</span>{/if}
        </p>
      {/if}
      {#if lockLine !== ''}
        <p id={lockId} class="text-muted" data-testid={lockId}>{lockLine}</p>
      {/if}
    </div>
    {#if manage}
      <div class="col-start-2 row-start-1 flex items-center justify-end gap-2 md:col-start-3">
        {#if canMove}
          <button
            type="button"
            class="{BILLING_OUTLINE_BUTTON} {picking === 'move' ? OPENER_PRESSED : ''}"
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
            class="{BILLING_TEXT_BUTTON} {busy ? BUSY_CLASS : ''}"
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
            class="{BILLING_MOVE_BUTTON} {picking === 'give' ? OPENER_PRESSED : ''}"
            aria-expanded={picking === 'give'}
            onclick={() => void open('give')}
            data-seat-action
            data-testid={`guild-seat-${seat.slot}-give`}
          >
            {copy.give}
          </button>
        {/if}
      </div>
    {/if}
  </div>
  {#if picking === null && error !== ''}
    <p class="text-[13px]" role="alert" data-testid={`guild-seat-${seat.slot}-error`}>{error}</p>
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
</li>
