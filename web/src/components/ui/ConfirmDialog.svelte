<!-- web/src/components/ui/ConfirmDialog.svelte -->
<!-- The one modal confirm (mock ws_prims `Dialog`; the account deletion dialog's geometry): a sheet
     over a scrim, a 560 px box centred from 768, a full-width bottom sheet below. A title, one line
     that says what the confirm does, any lines the caller adds (a balance, a problem), then Cancel
     and the confirm Button, which reads busy (spinner plus the busy word) while the call runs and
     is off while `disabled`. Focus goes to the title on open and stays inside (lib/ui/focus-trap);
     Escape, the scrim and Cancel close it, never while busy. The caller gives focus back to the
     control that opened it. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { tick } from 'svelte';
  import { type ButtonTone } from '../../lib/ui/button';
  import { busyCopy, uiCopy } from '../../lib/ui/copy';
  import { trapTab } from '../../lib/ui/focus-trap';
  import { REGION } from '../../lib/ui/regions';
  import Button from './Button.svelte';
  import Overlay from './Overlay.svelte';

  let {
    title,
    line,
    confirmLabel,
    cancelLabel = uiCopy.cancel,
    tone = 'primary',
    busy = false,
    busyLabel = busyCopy.working,
    disabled = false,
    onconfirm,
    onclose,
    testid = 'confirm',
    children,
  }: {
    title: string;
    /** What confirming does, in one sentence. */
    line: string;
    confirmLabel: string;
    cancelLabel?: string;
    /** The confirm button's tone: primary, or danger for something that cannot be undone. */
    tone?: ButtonTone;
    /** The call is running: the confirm reads busy and nothing closes the dialog. */
    busy?: boolean;
    busyLabel?: string;
    /** The confirm is off (a check has not passed yet); the caller says why in `children`. */
    disabled?: boolean;
    onconfirm: () => void;
    onclose: () => void;
    /** The base of the test ids: `{testid}-dialog`, `-title`, `-cancel`, `-confirm`. */
    testid?: string;
    /** Lines under the sentence (a balance, a problem), each with its own test id. */
    children?: Snippet;
  } = $props();

  let panel = $state<HTMLElement | undefined>(undefined);
  let heading = $state<HTMLElement | undefined>(undefined);
  const titleId = $derived(`${testid}-title`);

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
    if (panel !== undefined) trapTab(event, panel, heading);
  }
</script>

<div class="ui-confirm-scrim" aria-hidden="true" onclick={close}></div>
<div
  bind:this={panel}
  class="ui-confirm"
  role="dialog"
  aria-modal="true"
  aria-labelledby={titleId}
  tabindex="-1"
  data-region={REGION.dialog}
  data-testid={`${testid}-dialog`}
  {onkeydown}
>
  <Overlay kind="sheet" role="none" class="ui-confirm-box">
    <div class="flex flex-col gap-1 px-5 pt-5 pb-4">
      <h2
        bind:this={heading}
        id={titleId}
        class="text-strong m-0 text-[20px] leading-6 font-bold outline-none"
        tabindex="-1"
        data-testid={titleId}
      >
        {title}
      </h2>
      <p class="m-0 text-[15px] leading-[22px]">{line}</p>
      {#if children}{@render children()}{/if}
    </div>
    <div class="ui-confirm-actions ui-overlay-rule-t">
      <Button onclick={close} disabled={busy} testid={`${testid}-cancel`}>{cancelLabel}</Button>
      <Button
        tone={disabled ? 'secondary' : tone}
        {busy}
        {busyLabel}
        {disabled}
        onclick={onconfirm}
        testid={`${testid}-confirm`}>{confirmLabel}</Button
      >
    </div>
  </Overlay>
</div>

<style>
  .ui-confirm-scrim {
    position: fixed;
    inset: 0;
    z-index: 60;
    background: #05070bcc;
  }
  .ui-confirm {
    position: fixed;
    z-index: 61;
    right: 0;
    bottom: 0;
    left: 0;
    display: flex;
    flex-direction: column;
    outline: none;
  }
  .ui-confirm-actions {
    display: flex;
    gap: 12px;
    padding: 16px 20px calc(16px + env(safe-area-inset-bottom));
  }
  @media (max-width: 767px) {
    .ui-confirm-actions > :global(*) {
      flex: 1;
    }
  }
  @media (min-width: 768px) {
    .ui-confirm {
      top: 50%;
      right: auto;
      bottom: auto;
      left: 50%;
      width: 560px;
      max-width: calc(100% - 48px);
      transform: translate(-50%, -50%);
    }
    .ui-confirm > :global(.ui-confirm-box) {
      border-bottom-width: 1px;
      border-radius: var(--radius-plate);
    }
    .ui-confirm-actions {
      padding-bottom: 16px;
    }
  }
</style>
