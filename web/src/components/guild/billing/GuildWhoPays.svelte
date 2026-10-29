<!-- web/src/components/guild/billing/GuildWhoPays.svelte -->
<!-- Who pays (spec 3.2.6), the subscriber only: transfer the plan to another officer as an offer, the
     waiting line with Withdraw, and the declined line with the button back. The consequence sentence
     reads from the current plan, then what stays with the guild; nothing changes until the officer
     accepts. After a send or a withdraw, focus lands on the block's new action once it is enabled. -->
<script lang="ts">
  import { tick } from 'svelte';
  import type { GuildPlanView, TransferCandidate, TransferView } from '../../../lib/billing/contract';
  import { sendTransfer, withdrawTransfer } from '../../../lib/billing/guild-billing-api';
  import { seatsCopy } from '../../../lib/billing/seats-copy';
  import { transferErrorLine } from '../../../lib/guild/billing-errors';
  import {
    BILLING_OUTLINE_BUTTON,
    BILLING_SAFE_TEXT_BUTTON,
    CONFIRM_ACTION_WIDTH,
    CONFIRM_ACTIONS,
  } from '../../../lib/guild/billing-styles';
  import { focusOnMount } from '../../../lib/guild/focus-now';
  import { PRIMARY_BUTTON_FIXED } from '../../../lib/planner/styles';
  import { BUSY_CLASS } from '../../../lib/ui/busy';
  import SelectChevron from '../../ui/SelectChevron.svelte';

  const copy = seatsCopy.whoPays;

  let {
    guildId,
    plan,
    transfer,
    candidates,
    onChanged,
  }: {
    guildId: number;
    plan: GuildPlanView;
    transfer: TransferView | null;
    candidates: readonly TransferCandidate[];
    /** Resolves once the billing read has been reloaded. */
    onChanged: () => Promise<void>;
  } = $props();

  let opening = $state(false);
  let chosenId = $state<number | null>(null);
  let busy = $state(false);
  let error = $state('');
  let rootEl = $state<HTMLElement | null>(null);

  const waiting = $derived(transfer !== null && transfer.state === 'offered' && !transfer.to_you);
  const declined = $derived(transfer !== null && transfer.state === 'declined');
  const chosen = $derived(candidates.find((c) => c.user_id === chosenId) ?? candidates[0] ?? null);
  const startsAt = $derived(plan.current_period_end ?? plan.trial_end);
  const consequence = $derived(
    chosen !== null && startsAt !== null && plan.amount_cents !== null && plan.interval !== null
      ? copy.offerNote(chosen.name, startsAt, plan.amount_cents, plan.interval)
      : '',
  );

  /** Puts focus back on the block's action (Transfer the plan or Withdraw) once it has re-rendered. */
  async function refocusAction(): Promise<void> {
    await tick();
    rootEl?.querySelector<HTMLElement>('[data-whopays-action]')?.focus();
  }

  const openOfferTo = $derived(transfer !== null && transfer.state === 'offered' ? transfer.to_name : null);

  /** Focus moves only after `busy` clears: the new action (Withdraw) is disabled while busy, and a
   *  disabled button cannot take focus, which let it fall through to Records (UX 9). */
  async function run(action: () => Promise<void>, after: () => void = () => {}): Promise<void> {
    busy = true;
    error = '';
    let done = false;
    try {
      await action();
      after();
      await onChanged();
      done = true;
    } catch (thrown) {
      error = transferErrorLine(thrown, openOfferTo);
    } finally {
      busy = false;
    }
    if (done) await refocusAction();
  }

  const send = (): Promise<void> => {
    const target = chosen;
    if (target === null) return Promise.resolve();
    return run(
      () => sendTransfer(guildId, target.user_id),
      () => (opening = false),
    );
  };

  function close(): void {
    opening = false;
    error = '';
    void refocusAction();
  }

  function onKeydown(event: KeyboardEvent): void {
    if (event.key !== 'Escape') return;
    event.preventDefault();
    close();
  }
</script>

<div bind:this={rootEl} class="flex flex-col gap-2" data-testid="guild-whopays">
  <h3 class="section-title">{copy.heading}</h3>
  <p class="text-muted text-[13px]">{copy.help}</p>
  {#if waiting && transfer !== null}
    <div class="flex flex-wrap items-center gap-3">
      <p class="text-[14px]" data-testid="guild-transfer-waiting">
        {copy.offered(transfer.to_name, transfer.at)}
      </p>
      <button
        type="button"
        class="{BILLING_SAFE_TEXT_BUTTON} {busy ? BUSY_CLASS : ''}"
        disabled={busy}
        aria-busy={busy}
        onclick={() => void run(() => withdrawTransfer(guildId))}
        data-whopays-action
        data-testid="guild-transfer-withdraw"
      >
        {copy.withdraw}
      </button>
    </div>
    {#if error !== ''}
      <p class="text-[13px]" role="alert" data-testid="guild-transfer-error">{error}</p>
    {/if}
  {:else if opening && candidates.length > 0}
    <div
      class="border-line-soft flex flex-col gap-3 border p-4"
      role="group"
      aria-labelledby="guild-transfer-label"
      onkeydown={onKeydown}
      data-testid="guild-transfer-confirm"
    >
      <label class="flex flex-col gap-1" for="guild-transfer-select">
        <span id="guild-transfer-label" class="quiet-label">{copy.transferTo}</span>
        <span class="relative block">
          <select
            id="guild-transfer-select"
            class="border-line-warm bg-raised text-text rounded-control hover:border-gold-hover h-11 w-full min-w-0 cursor-pointer appearance-none truncate border pr-8 pl-3 text-[13px]"
            use:focusOnMount
            value={chosen?.user_id}
            onchange={(event) => (chosenId = Number((event.currentTarget as HTMLSelectElement).value))}
            data-testid="guild-transfer-select"
          >
            {#each candidates as candidate (candidate.user_id)}
              <option value={candidate.user_id}>{candidate.name}</option>
            {/each}
          </select>
          <SelectChevron />
        </span>
      </label>
      {#if consequence !== '' && chosen !== null}
        <p class="text-[14px]" data-testid="guild-transfer-consequence">{consequence}</p>
        <p class="text-muted text-[13px]" data-testid="guild-transfer-stays">
          {copy.staysWithGuild(chosen.name)}
        </p>
      {/if}
      {#if error !== ''}
        <p class="text-[13px]" role="alert" data-testid="guild-transfer-error">{error}</p>
      {/if}
      <div class={CONFIRM_ACTIONS}>
        <button
          type="button"
          class="{PRIMARY_BUTTON_FIXED} {CONFIRM_ACTION_WIDTH} px-4 {busy ? BUSY_CLASS : ''}"
          disabled={busy}
          aria-busy={busy}
          onclick={() => void send()}
          data-testid="guild-transfer-send"
        >
          {busy ? copy.sending : copy.send}
        </button>
        <button type="button" class="{BILLING_SAFE_TEXT_BUTTON} {CONFIRM_ACTION_WIDTH}" onclick={close}
          >{copy.cancel}</button
        >
      </div>
    </div>
  {:else}
    {#if declined && transfer !== null}
      <p class="text-muted text-[13px]" data-testid="guild-transfer-declined">
        {copy.declined(transfer.to_name)}
      </p>
    {/if}
    <div class="flex flex-wrap items-center gap-3">
      <button
        type="button"
        class={BILLING_OUTLINE_BUTTON}
        disabled={candidates.length === 0}
        aria-describedby={candidates.length === 0 ? 'guild-transfer-none' : undefined}
        onclick={() => (opening = true)}
        data-whopays-action
        data-testid="guild-transfer-open"
      >
        {copy.transfer}
      </button>
      {#if candidates.length === 0}
        <p class="text-muted text-[13px]" id="guild-transfer-none" data-testid="guild-transfer-none">
          {copy.noCandidates}
        </p>
      {/if}
    </div>
    {#if error !== ''}
      <p class="text-[13px]" role="alert" data-testid="guild-transfer-error">{error}</p>
    {/if}
  {/if}
</div>
