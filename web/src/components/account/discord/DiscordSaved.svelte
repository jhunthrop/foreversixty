<!-- web/src/components/account/discord/DiscordSaved.svelte -->
<!-- A saved webhook's block (saved, switched off, lapsed; premium spec 2026-11-08 section 7.6): where
     it posts, the warning or notice, the event switches, Send test message and Remove, and the way to
     a different channel, in the spec's per-state order (3.2): switched off puts the events last and
     the replace link right under the notice. The replace link hides while the remove question is
     open. -->
<script lang="ts">
  import type { DiscordEvents, DiscordNotifications } from '../../../lib/account/api';
  import { removeQuestionFor, sourceOf, type DiscordBlockKey } from '../../../lib/account/notifications';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import { LINK } from '../../../lib/ui/primitives';
  import DiscordActions from './DiscordActions.svelte';
  import DiscordConnectLink from './DiscordConnectLink.svelte';
  import DiscordEventRows from './DiscordEventRows.svelte';
  import DiscordStatusLine from './DiscordStatusLine.svelte';
  import DiscordSummary from './DiscordSummary.svelte';

  let {
    state,
    discord,
    confirming,
    busy,
    status,
    onToggleEvents,
    onTest,
    onAskRemove,
    onRemove,
    onCancel,
    onReplace,
    onStuck,
  }: {
    state: Exclude<DiscordBlockKey, 'off' | 'free' | 'empty'>;
    discord: DiscordNotifications;
    confirming: boolean;
    busy: 'save' | 'events' | 'test' | 'remove' | null;
    status: { ok: boolean; message: string; target?: string } | null;
    onToggleEvents: (next: DiscordEvents) => Promise<boolean>;
    onTest: () => void;
    onAskRemove: () => void;
    onRemove: () => void;
    onCancel: () => void;
    onReplace: () => void;
    onStuck: () => void;
  } = $props();

  const connected = $derived(sourceOf(discord) === 'connected');
  const replaceShown = $derived(!confirming && (state === 'saved' || state === 'disabled'));
</script>

{#snippet eventRows()}
  <DiscordEventRows
    events={discord.events}
    disabled={busy !== null}
    inactive={state === 'lapsed' || state === 'disabled'}
    busy={busy === 'events'}
    showNote={state !== 'lapsed'}
    onToggle={onToggleEvents}
  />
{/snippet}

{#snippet replace()}
  {#if connected}
    <p class="flex flex-wrap items-baseline gap-x-3">
      <DiscordConnectLink
        label={notificationsCopy.reconnect}
        class="{LINK} text-[15px]"
        testid="discord-replace"
        busy={busy !== null}
        describedBy="discord-replace-note"
        {onStuck}
      >
        {#snippet content(text)}{text} ›{/snippet}
      </DiscordConnectLink>
      <span id="discord-replace-note" class="text-small text-muted" data-testid="discord-replace-note"
        >{notificationsCopy.reconnectAside}</span
      >
    </p>
  {:else}
    <p>
      <button
        type="button"
        class="{LINK} text-[15px]"
        disabled={busy !== null}
        onclick={onReplace}
        data-testid="discord-replace">{notificationsCopy.replace} ›</button
      >
    </p>
  {/if}
{/snippet}

<DiscordSummary {state} {discord} />
{#if state === 'disabled' && replaceShown}{@render replace()}{/if}
{#if state !== 'disabled'}{@render eventRows()}{/if}
<div class="flex flex-col">
  <DiscordActions
    {state}
    {confirming}
    removeQuestion={removeQuestionFor(discord)}
    {busy}
    {onTest}
    {onAskRemove}
    {onRemove}
    {onCancel}
  />
  <DiscordStatusLine {status} reserve={false} />
</div>
{#if state === 'saved' && replaceShown}{@render replace()}{/if}
{#if state === 'disabled'}{@render eventRows()}{/if}
