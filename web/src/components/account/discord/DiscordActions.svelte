<!-- web/src/components/account/discord/DiscordActions.svelte -->
<!-- The actions row: Send test message (or Test and switch on, the one filled button of a
     switched-off panel) and Remove, or the remove question with its own Remove and Cancel.
     `gateId` (the guild's officer-read-only) disables both test buttons and describes them by
     that sentence; Remove stays live. `disabled` (a frozen claim) turns every button off. The
     buttons are the shared recipes: test is secondary (primary when it switches posts back on),
     Remove is danger, the question's Cancel is ghost. -->
<script lang="ts">
  import type { DiscordBlockKey } from '../../../lib/account/notifications';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import { button } from '../../../lib/ui/button';
  import Glyph from '../../ui/Glyph.svelte';
  import Spinner from '../../ui/Spinner.svelte';

  let {
    state,
    confirming,
    removeQuestion,
    busy,
    gateId,
    disabled = false,
    removeLabel = notificationsCopy.remove,
    removingLabel = notificationsCopy.removing,
    sendGlyph = true,
    onTest,
    onAskRemove,
    onRemove,
    onCancel,
  }: {
    state: DiscordBlockKey;
    confirming: boolean;
    /** The question for the webhook's source (connected or pasted). */
    removeQuestion: string;
    busy: 'save' | 'events' | 'test' | 'remove' | null;
    gateId?: string;
    disabled?: boolean;
    /** The remove button's words; the guild section keeps Remove, the account card says Unlink Discord. */
    removeLabel?: string;
    removingLabel?: string;
    /** The send glyph on the test button; the guild section is too narrow for it at 360 px. */
    sendGlyph?: boolean;
    onTest: () => void;
    onAskRemove: () => void;
    onRemove: () => void;
    onCancel: () => void;
  } = $props();

  const label = (text: string, pending: boolean, busyText: string): string => (pending ? busyText : text);
  const off = $derived(busy !== null || disabled);
  const testOff = $derived(off || gateId !== undefined);
</script>

{#if confirming}
  <p class="text-[14px]">{removeQuestion}</p>
  <div class="flex flex-wrap gap-3">
    <button
      class={button('danger')}
      disabled={off}
      aria-busy={busy === 'remove' ? 'true' : undefined}
      onclick={onRemove}
      data-testid="discord-remove-confirm"
    >
      {#if busy === 'remove'}<Spinner />{/if}{label(removeLabel, busy === 'remove', removingLabel)}
    </button>
    <button class={button('ghost')} disabled={off} onclick={onCancel} data-testid="discord-remove-cancel">
      {notificationsCopy.cancel}
    </button>
  </div>
{:else}
  <div class="flex flex-wrap gap-3">
    {#if state === 'disabled' || state === 'saved'}
      <button
        class={button(state === 'disabled' ? 'primary' : 'secondary')}
        disabled={testOff}
        aria-describedby={gateId}
        aria-busy={busy === 'test' ? 'true' : undefined}
        onclick={onTest}
        data-testid="discord-test"
      >
        {#if busy === 'test'}<Spinner />{:else if sendGlyph}<Glyph name="send" />{/if}{label(
          state === 'disabled' ? notificationsCopy.testAndSwitchOn : notificationsCopy.test,
          busy === 'test',
          notificationsCopy.testing,
        )}
      </button>
    {/if}
    <button class={button('danger')} disabled={off} onclick={onAskRemove} data-testid="discord-remove">
      {removeLabel}
    </button>
  </div>
{/if}
