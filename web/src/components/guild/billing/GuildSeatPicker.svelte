<!-- web/src/components/guild/billing/GuildSeatPicker.svelte -->
<!-- Give and Move, in place under the row (spec 3.2.5): a native select of verified members who hold no
     seat here, the payer note when a member who pays for Premium is chosen, focus on the select,
     Escape is handled by the row. Presentational: the row loads the candidates and sends the choice. -->
<script lang="ts">
  import type { SeatCandidate } from '../../../lib/billing/contract';
  import { seatsCopy as copy } from '../../../lib/billing/seats-copy';
  import {
    BILLING_SAFE_TEXT_BUTTON,
    CONFIRM_ACTION_WIDTH,
    CONFIRM_ACTIONS,
  } from '../../../lib/guild/billing-styles';
  import { focusOnMount } from '../../../lib/guild/focus-now';
  import { sortedByName } from '../../../lib/guild/seat-view';
  import { PRIMARY_BUTTON_FIXED } from '../../../lib/planner/styles';
  import { BUSY_CLASS } from '../../../lib/ui/busy';
  import Skeleton from '../../ui/Skeleton.svelte';
  import Select from '../../ui/Select.svelte';

  let {
    slot,
    mode,
    holderName,
    candidates,
    busy,
    error,
    onSubmit,
    onCancel,
  }: {
    slot: number;
    mode: 'give' | 'move';
    /** The current holder's name, for the move note. */
    holderName: string;
    /** Null while the candidates load. */
    candidates: readonly SeatCandidate[] | null;
    busy: boolean;
    error: string;
    onSubmit: (candidate: SeatCandidate) => void;
    onCancel: () => void;
  } = $props();

  const options = $derived(candidates === null ? [] : sortedByName(candidates));
  let chosenId = $state<number | null>(null);
  const chosen = $derived(options.find((c) => c.user_id === chosenId) ?? options[0] ?? null);
  const title = $derived(mode === 'give' ? copy.picker.giveTitle : copy.picker.moveTitle);
  const submitText = $derived.by(() => {
    if (busy) return mode === 'give' ? copy.giving : copy.moving;
    return mode === 'give' ? copy.picker.give : copy.picker.move;
  });
  const selectId = $derived(`guild-seat-${slot}-select`);
</script>

<div
  class="border-line-soft flex flex-col gap-3 border p-4"
  role="group"
  aria-labelledby={`${selectId}-label`}
  aria-busy={candidates === null || busy}
  data-testid={`guild-seat-${slot}-picker`}
>
  {#if candidates === null}
    <Skeleton lines={1} rowHeight="h-11" label="Loading members" testid={`guild-seat-${slot}-loading`} />
  {:else if options.length === 0}
    <p class="text-[14px]" data-testid={`guild-seat-${slot}-empty`}>{copy.picker.everyoneHasPremium}</p>
    <button type="button" class={BILLING_SAFE_TEXT_BUTTON} use:focusOnMount onclick={onCancel}>
      {copy.picker.cancel}
    </button>
  {:else}
    <div class="flex flex-col gap-1">
      <label id={`${selectId}-label`} class="quiet-label" for={selectId}>{title}</label>
      <Select
        id={selectId}
        value={chosen === null ? '' : String(chosen.user_id)}
        options={options.map((option) => ({
          value: String(option.user_id),
          label: copy.picker.option(option.name, option.class, option.pays_premium),
        }))}
        focusOnMount
        testid={`guild-seat-${slot}-select`}
        onchange={(event) => (chosenId = Number((event.currentTarget as HTMLSelectElement).value))}
      />
    </div>
    {#if chosen?.pays_premium}
      <p class="text-muted text-[13px]" data-testid={`guild-seat-${slot}-pays-note`}>
        {copy.picker.paysNote}
      </p>
    {/if}
    {#if mode === 'move'}
      <p class="text-muted text-[13px]">{copy.picker.keepsNote(holderName)}</p>
    {/if}
    {#if error !== ''}
      <p class="text-[13px]" role="alert" data-testid={`guild-seat-${slot}-error`}>{error}</p>
    {/if}
    <div class={CONFIRM_ACTIONS}>
      <button
        type="button"
        class="{PRIMARY_BUTTON_FIXED} {CONFIRM_ACTION_WIDTH} px-4 {busy ? BUSY_CLASS : ''}"
        disabled={busy || chosen === null}
        aria-busy={busy}
        onclick={() => chosen !== null && onSubmit(chosen)}
        data-seat-submit
        data-testid={`guild-seat-${slot}-submit`}
      >
        {submitText}
      </button>
      <button type="button" class="{BILLING_SAFE_TEXT_BUTTON} {CONFIRM_ACTION_WIDTH}" onclick={onCancel}
        >{copy.picker.cancel}</button
      >
    </div>
  {/if}
</div>
