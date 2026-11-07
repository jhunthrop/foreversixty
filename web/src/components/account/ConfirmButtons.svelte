<!-- web/src/components/account/ConfirmButtons.svelte -->
<!-- The two buttons every destructive ask on /account ends in (premium spec 2026-11-08 section 9:
     Leave guild, Revoke, Unlink, Sign out): the danger confirm and a secondary Cancel. Focus lands
     on Cancel, so a stray Enter backs out; Escape cancels; busy disables both. The question itself
     is the caller's, wherever its row puts it. -->
<script lang="ts">
  import { smallButton } from '../../lib/account/small-button';
  import { button, type ButtonSize } from '../../lib/ui/button';
  import Spinner from '../ui/Spinner.svelte';

  let {
    confirm,
    cancel,
    busy = false,
    size = 'sm',
    onconfirm,
    oncancel,
    testid,
  }: {
    confirm: string;
    cancel: string;
    busy?: boolean;
    size?: ButtonSize;
    onconfirm: () => void;
    oncancel: () => void;
    /** `<testid>-confirm` and `<testid>-cancel`. */
    testid: string;
  } = $props();

  /** Focus on Cancel, drawn as keyboard focus: the ask follows a press, and the ring says where it went. */
  function focusOnMount(node: HTMLElement): void {
    node.focus({ focusVisible: true });
  }

  function onKeydown(event: KeyboardEvent): void {
    if (event.key === 'Escape' && !busy) oncancel();
  }
</script>

<svelte:window onkeydown={onKeydown} />
<button
  type="button"
  class={size === 'sm' ? smallButton('danger') : button('danger', size)}
  disabled={busy}
  aria-busy={busy ? 'true' : undefined}
  onclick={onconfirm}
  data-testid="{testid}-confirm"
  >{#if busy}<Spinner />{/if}{confirm}</button
>
<button
  type="button"
  class={size === 'sm' ? smallButton() : button('secondary', size)}
  disabled={busy}
  onclick={oncancel}
  use:focusOnMount
  data-testid="{testid}-cancel">{cancel}</button
>
