<!-- web/src/components/account/discord/DiscordActions.svelte -->
<!-- The actions row: Send test message (or Test and switch on, the one filled button of a
     switched-off panel) and Remove, or the remove question with its own Remove and Cancel.
     `gateId` (the guild's officer-read-only) disables both test buttons and describes them by
     that sentence; Remove stays live. `disabled` (a frozen claim) turns every button off. -->
<script lang="ts">
  import type { DiscordBlockKey } from '../../../lib/account/notifications';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import { PRIMARY_BUTTON_FIXED, SECONDARY_BUTTON_FIXED } from '../../../lib/planner/styles';

  const OUTLINE = `${SECONDARY_BUTTON_FIXED} border-line-warm text-text px-4`;

  let {
    state,
    confirming,
    removeQuestion,
    busy,
    gateId,
    disabled = false,
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
    onTest: () => void;
    onAskRemove: () => void;
    onRemove: () => void;
    onCancel: () => void;
  } = $props();

  const label = (text: string, pending: boolean, busyText: string): string => (pending ? busyText : text);
  const off = $derived(busy !== null || disabled);
  const testOff = $derived(off || gateId !== undefined);
  // The gate's and the freeze's off state is the shared disabled look; a busy row keeps its look.
  const DIM = 'cursor-not-allowed opacity-50';
  const testDim = $derived(disabled || gateId !== undefined ? DIM : '');
  const allDim = $derived(disabled ? DIM : '');
</script>

{#if confirming}
  <p class="text-[14px]">{removeQuestion}</p>
  <div class="flex flex-wrap gap-3">
    <button class="{OUTLINE} {allDim}" disabled={off} onclick={onRemove} data-testid="discord-remove-confirm">
      {label(notificationsCopy.remove, busy === 'remove', notificationsCopy.removing)}
    </button>
    <button class="{OUTLINE} {allDim}" disabled={off} onclick={onCancel} data-testid="discord-remove-cancel">
      {notificationsCopy.cancel}
    </button>
  </div>
{:else}
  <div class="flex flex-wrap gap-3">
    {#if state === 'disabled'}
      <button
        class="{PRIMARY_BUTTON_FIXED} px-5 {testDim}"
        disabled={testOff}
        aria-describedby={gateId}
        aria-busy={busy === 'test'}
        onclick={onTest}
        data-testid="discord-test"
      >
        {label(notificationsCopy.testAndSwitchOn, busy === 'test', notificationsCopy.testing)}
      </button>
    {:else if state === 'saved'}
      <button
        class="{OUTLINE} {testDim}"
        disabled={testOff}
        aria-describedby={gateId}
        aria-busy={busy === 'test'}
        onclick={onTest}
        data-testid="discord-test"
      >
        {label(notificationsCopy.test, busy === 'test', notificationsCopy.testing)}
      </button>
    {/if}
    <button class="{OUTLINE} {allDim}" disabled={off} onclick={onAskRemove} data-testid="discord-remove">
      {notificationsCopy.remove}
    </button>
  </div>
{/if}
