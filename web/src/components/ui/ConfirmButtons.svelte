<!-- web/src/components/ui/ConfirmButtons.svelte -->
<!-- The two buttons every inline ask ends in (premium spec 2026-11-08 section 9; component audit
     2026-11-10 cluster 14: account, raid clear and cancel, guild billing, the claim): the confirm
     (danger, or primary for a non-destructive ask) and a secondary Cancel. Focus lands
     on Cancel, so a stray Enter backs out; Escape cancels; busy disables both. The question itself
     is the caller's, wherever its row puts it. -->
<script lang="ts">
  import { button, type ButtonSize } from '../../lib/ui/button';
  import { uiCopy } from '../../lib/ui/copy';
  import Spinner from './Spinner.svelte';

  let {
    confirm,
    cancel = uiCopy.cancel,
    busyLabel,
    tone = 'danger',
    busy = false,
    size = 'sm',
    disabled = false,
    focusCancel = true,
    cancelWhileBusy = false,
    width = '',
    confirmTestid,
    cancelTestid,
    onconfirm,
    oncancel,
    testid,
  }: {
    confirm: string;
    cancel?: string;
    /** The confirm's busy word ("Leaving…"); the confirm's own words when left out. */
    busyLabel?: string;
    /** `primary` for an ask that destroys nothing (Upgrade, Claim). */
    tone?: 'danger' | 'primary';
    busy?: boolean;
    size?: ButtonSize;
    /** The confirm waits (a re-read in flight); Cancel stays live. */
    disabled?: boolean;
    /** False where the ask opens on a field that takes focus itself (the transfer's officer pick). */
    focusCancel?: boolean;
    /** Cancel stays live while the confirm is in flight (the raid's Keep it). */
    cancelWhileBusy?: boolean;
    /** Layout only, on both buttons (`w-full` in a card's foot, a fixed width). */
    width?: string;
    /** A caller whose tests already name the two buttons keeps its ids. */
    confirmTestid?: string;
    cancelTestid?: string;
    onconfirm: () => void;
    oncancel: () => void;
    /** `<testid>-confirm` and `<testid>-cancel`, unless the two ids above are given. */
    testid: string;
  } = $props();

  /** Focus on Cancel, drawn as keyboard focus: the ask follows a press, and the ring says where it went. */
  function focusOnMount(node: HTMLElement): void {
    if (focusCancel) node.focus({ focusVisible: true });
  }

  function onKeydown(event: KeyboardEvent): void {
    if (event.key === 'Escape' && (!busy || cancelWhileBusy)) oncancel();
  }
</script>

<svelte:window onkeydown={onKeydown} />
<button
  type="button"
  class={width === '' ? button(tone, size) : `${button(tone, size)} ${width}`}
  disabled={busy || disabled}
  aria-busy={busy ? 'true' : undefined}
  onclick={onconfirm}
  data-testid={confirmTestid ?? `${testid}-confirm`}
  >{#if busy}<Spinner />{busyLabel ?? confirm}{:else}{confirm}{/if}</button
>
<button
  type="button"
  class={width === '' ? button('secondary', size) : `${button('secondary', size)} ${width}`}
  disabled={busy && !cancelWhileBusy}
  onclick={oncancel}
  use:focusOnMount
  data-testid={cancelTestid ?? `${testid}-cancel`}>{cancel}</button
>
