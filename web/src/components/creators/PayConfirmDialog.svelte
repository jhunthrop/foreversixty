<!-- web/src/components/creators/PayConfirmDialog.svelte -->
<!-- The payout confirm on /admin/creators/payouts (mock ws_prims `Dialog`, web-admin-creators-confirm):
     the account deletion dialog's sheet and geometry (AccountDeleteDialog), one line and two actions.
     Focus goes to the title on open and stays inside; Escape, the scrim and Cancel close it (not while
     the transfer runs). -->
<script lang="ts">
  import { tick } from 'svelte';
  import { payoutsCopy as copy } from '../../lib/creators/copy';
  import { button } from '../../lib/ui/button';
  import { REGION } from '../../lib/ui/regions';
  import Button from '../ui/Button.svelte';
  import Overlay from '../ui/Overlay.svelte';
  import Spinner from '../ui/Spinner.svelte';

  let {
    title,
    amount,
    busy,
    balance,
    blocked,
    problem,
    onconfirm,
    onclose,
  }: {
    title: string;
    amount: string;
    busy: boolean;
    /** The balance line (`Available in Stripe: …`), or the checking line while it loads. */
    balance: string;
    /** Pay is off: the balance is short, still loading, or failed to load. */
    blocked: boolean;
    /** Why Pay is off or what went wrong, in plain words; '' for nothing. */
    problem: string;
    onconfirm: () => void;
    onclose: () => void;
  } = $props();

  let panel = $state<HTMLElement | undefined>(undefined);
  let heading = $state<HTMLElement | undefined>(undefined);

  $effect(() => {
    void tick().then(() => heading?.focus());
  });

  function close(): void {
    if (!busy) onclose();
  }

  function onkeydown(event: KeyboardEvent): void {
    if (event.key === 'Escape') {
      event.preventDefault();
      close();
      return;
    }
    if (event.key !== 'Tab' || panel === undefined) return;
    const items = [...panel.querySelectorAll<HTMLElement>('button:not([disabled])')];
    if (items.length === 0) return;
    const first = items[0];
    const last = items[items.length - 1];
    if (event.shiftKey && (document.activeElement === first || document.activeElement === heading)) {
      event.preventDefault();
      last.focus();
    } else if (!event.shiftKey && document.activeElement === last) {
      event.preventDefault();
      first.focus();
    }
  }
</script>

<div class="pay-scrim" aria-hidden="true" onclick={close}></div>
<div
  bind:this={panel}
  class="pay-dialog"
  role="dialog"
  aria-modal="true"
  aria-labelledby="pay-title"
  tabindex="-1"
  data-region={REGION.dialog}
  data-testid="pay-dialog"
  {onkeydown}
>
  <Overlay kind="sheet" role="none" class="pay-box">
    <div class="flex flex-col gap-1 px-5 pt-5 pb-4">
      <h2
        bind:this={heading}
        id="pay-title"
        class="text-strong m-0 text-[20px] leading-6 font-bold outline-none"
        tabindex="-1"
      >
        {title}
      </h2>
      <p class="m-0 text-[15px] leading-[22px]">{copy.confirmBody(amount)}</p>
      <p class="text-muted m-0 mt-2 text-[14px] leading-5" data-testid="pay-balance">{balance}</p>
      {#if problem !== ''}<p
          class="text-wipe m-0 mt-1 text-[14px] leading-5"
          role="alert"
          data-testid="pay-problem"
        >
          {problem}
        </p>{/if}
    </div>
    <div class="pay-actions ui-overlay-rule-t">
      <Button onclick={close} disabled={busy} testid="pay-cancel">{copy.cancel}</Button>
      <button
        type="button"
        class={button(blocked ? 'secondary' : 'primary')}
        disabled={busy || blocked}
        aria-busy={busy ? 'true' : undefined}
        onclick={onconfirm}
        data-testid="pay-confirm"
        >{#if busy}<Spinner />{copy.approving}{:else}{copy.confirmPay(amount)}{/if}</button
      >
    </div>
  </Overlay>
</div>

<style>
  .pay-scrim {
    position: fixed;
    inset: 0;
    z-index: 60;
    background: #05070bcc;
  }
  .pay-dialog {
    position: fixed;
    z-index: 61;
    right: 0;
    bottom: 0;
    left: 0;
    display: flex;
    flex-direction: column;
    outline: none;
  }
  .pay-actions {
    display: flex;
    gap: 12px;
    padding: 16px 20px calc(16px + env(safe-area-inset-bottom));
  }
  @media (max-width: 767px) {
    .pay-actions > :global(*) {
      flex: 1;
    }
  }
  @media (min-width: 768px) {
    .pay-dialog {
      top: 50%;
      right: auto;
      bottom: auto;
      left: 50%;
      width: 560px;
      max-width: calc(100% - 48px);
      transform: translate(-50%, -50%);
    }
    .pay-dialog > :global(.pay-box) {
      border-bottom-width: 1px;
      border-radius: var(--radius-plate);
    }
    .pay-actions {
      padding-bottom: 16px;
    }
  }
</style>
