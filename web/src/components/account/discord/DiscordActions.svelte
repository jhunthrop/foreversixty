<!-- web/src/components/account/discord/DiscordActions.svelte -->
<!-- The actions row: Send test message (or Test and switch on, the one filled button of a
     switched-off panel) and Remove, or the remove question with its own Remove and Cancel. -->
<script lang="ts">
  import type { DiscordBlockKey } from '../../../lib/account/notifications';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import { PRIMARY_BUTTON_FIXED, SECONDARY_BUTTON_FIXED } from '../../../lib/planner/styles';

  const OUTLINE = `${SECONDARY_BUTTON_FIXED} border-line-warm text-text px-4`;

  let {
    state,
    confirming,
    busy,
    onTest,
    onAskRemove,
    onRemove,
    onCancel,
  }: {
    state: DiscordBlockKey;
    confirming: boolean;
    busy: 'save' | 'events' | 'test' | 'remove' | null;
    onTest: () => void;
    onAskRemove: () => void;
    onRemove: () => void;
    onCancel: () => void;
  } = $props();

  const label = (text: string, pending: boolean, busyText: string): string => (pending ? busyText : text);
</script>

{#if confirming}
  <p class="text-[14px]">{notificationsCopy.removeQuestion}</p>
  <div class="flex flex-wrap gap-3">
    <button class={OUTLINE} disabled={busy !== null} onclick={onRemove} data-testid="discord-remove-confirm">
      {label(notificationsCopy.remove, busy === 'remove', notificationsCopy.removing)}
    </button>
    <button class={OUTLINE} disabled={busy !== null} onclick={onCancel} data-testid="discord-remove-cancel">
      {notificationsCopy.cancel}
    </button>
  </div>
{:else}
  <div class="flex flex-wrap gap-3">
    {#if state === 'disabled'}
      <button
        class="{PRIMARY_BUTTON_FIXED} px-5"
        disabled={busy !== null}
        aria-busy={busy === 'test'}
        onclick={onTest}
        data-testid="discord-test"
      >
        {label(notificationsCopy.testAndSwitchOn, busy === 'test', notificationsCopy.testing)}
      </button>
    {:else if state === 'saved'}
      <button
        class={OUTLINE}
        disabled={busy !== null}
        aria-busy={busy === 'test'}
        onclick={onTest}
        data-testid="discord-test"
      >
        {label(notificationsCopy.test, busy === 'test', notificationsCopy.testing)}
      </button>
    {/if}
    <button class={OUTLINE} disabled={busy !== null} onclick={onAskRemove} data-testid="discord-remove">
      {notificationsCopy.remove}
    </button>
  </div>
{/if}
