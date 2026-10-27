<!-- web/src/components/guild/billing/GuildSeatRow.svelte -->
<!-- One seat (spec 3.2.5): Held, Offered (waiting), Accepted, Empty or Returned. The right column is
     Move seat, Withdraw or Give seat, and only a reader who may manage seats gets it. Give and Move
     open the picker in place below the row; a locked Move seat is aria-disabled and its date is text. -->
<script lang="ts">
  import { BillingApiError } from '../../../lib/billing/api';
  import type { SeatCandidate, SeatView } from '../../../lib/billing/contract';
  import { giveSeat, moveSeat, seatCandidates, withdrawSeat } from '../../../lib/billing/seats-api';
  import { seatsCopy as copy } from '../../../lib/billing/seats-copy';
  import {
    BILLING_MOVE_BUTTON,
    BILLING_OUTLINE_BUTTON,
    BILLING_TEXT_BUTTON,
  } from '../../../lib/guild/billing-styles';
  import { isSeatLocked, seatErrorLine, seatRowKind } from '../../../lib/guild/seat-view';
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
    onChanged: () => void;
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
  /** An accepted seat's middle is its start; its lock date follows on a line of its own. */
  const lockLine = $derived(
    kind === 'accepted' && locked && seat.movable_at !== null ? copy.canMoveFrom(seat.movable_at) : '',
  );
  const canMove = $derived(kind === 'held' || kind === 'accepted');

  async function open(mode: 'give' | 'move'): Promise<void> {
    if (mode === 'move' && locked) return;
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
  }

  async function submit(candidate: SeatCandidate): Promise<void> {
    if (picking === null) return;
    busy = true;
    error = '';
    try {
      const send = picking === 'give' ? giveSeat : moveSeat;
      await send(guildId, seat.slot, candidate.user_id);
      picking = null;
      onChanged();
    } catch (failure) {
      error = seatErrorLine(failure, candidate.name);
      if (failure instanceof BillingApiError && failure.code === 'seat_locked') onChanged();
    } finally {
      busy = false;
    }
  }

  async function withdraw(): Promise<void> {
    busy = true;
    error = '';
    try {
      await withdrawSeat(guildId, seat.slot);
      onChanged();
    } catch {
      error = copy.picker.saveFailed;
    } finally {
      busy = false;
    }
  }
</script>

<li
  class="border-line-soft flex flex-col gap-3 border-b py-3 last:border-b-0"
  data-testid={`guild-seat-${seat.slot}`}
  data-seat-kind={kind}
>
  <div
    class="grid items-center gap-x-4 gap-y-2 md:min-h-12 {manage
      ? 'md:grid-cols-[minmax(220px,1fr)_minmax(0,1.4fr)_auto]'
      : 'md:grid-cols-[minmax(220px,1fr)_minmax(0,1.4fr)]'}"
  >
    <div class="flex items-center gap-3">
      {#if seat.holder !== null}
        <ClassCrestRing characterClass={seat.holder.class} size={32} />
        <div class="flex min-w-0 flex-col">
          <span
            class="font-display truncate text-[14px] font-bold"
            style={`color: ${classColorVar(seat.holder.class)}`}
            data-testid={`guild-seat-${seat.slot}-name`}>{seat.holder.name}</span
          >
          {#if kind === 'offered'}
            <span><span class="pill pill-sample">{copy.waiting}</span></span>
          {:else if seat.given_at !== null && kind === 'held'}
            <span class="text-muted text-[12px]">{copy.givenOn(seat.given_at)}</span>
          {/if}
        </div>
      {:else}
        <span
          class="border-line inline-block size-8 shrink-0 rounded-full border border-dashed"
          aria-hidden="true"
        ></span>
        <span class="text-muted text-[14px]">{copy.notGiven}</span>
      {/if}
    </div>
    <div class="flex flex-col gap-1 text-[13px]">
      {#if middle !== ''}
        <p
          id={noteId}
          class={kind === 'held' && !locked ? 'text-text' : 'text-muted'}
          data-testid={`guild-seat-${seat.slot}-note`}
        >
          {middle}
        </p>
      {/if}
      {#if lockLine !== ''}
        <p id={lockId} class="text-muted" data-testid={lockId}>{lockLine}</p>
      {/if}
    </div>
    {#if manage}
      <div class="flex items-center">
        {#if canMove}
          <button
            type="button"
            class="{BILLING_OUTLINE_BUTTON} {locked ? 'cursor-not-allowed opacity-50' : ''}"
            aria-disabled={locked}
            aria-describedby={locked ? (lockLine !== '' ? lockId : noteId) : undefined}
            onclick={() => void open('move')}
            data-testid={`guild-seat-${seat.slot}-move`}
          >
            {copy.move}
          </button>
        {:else if kind === 'offered'}
          <button
            type="button"
            class="{BILLING_TEXT_BUTTON} {busy ? BUSY_CLASS : ''}"
            disabled={busy}
            aria-busy={busy}
            onclick={() => void withdraw()}
            data-testid={`guild-seat-${seat.slot}-withdraw`}
          >
            {busy ? copy.withdrawing : copy.withdraw}
          </button>
        {:else}
          <button
            type="button"
            class={BILLING_MOVE_BUTTON}
            onclick={() => void open('give')}
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
