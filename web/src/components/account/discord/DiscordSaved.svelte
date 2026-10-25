<!-- web/src/components/account/discord/DiscordSaved.svelte -->
<!-- A saved webhook's panel body (saved, switched off, lapsed): destination, notice, the events
     and the actions, in the spec's per-state order (3.2). Switched off puts the events last and
     `Use a different webhook` inside the notice; saved keeps it under the status line. Both hide
     the link while the remove question is open. -->
<script lang="ts">
  import type { DiscordEvents, DiscordNotifications } from '../../../lib/account/api';
  import { removeQuestionFor, type DiscordBlockKey } from '../../../lib/account/notifications';
  import DiscordActions from './DiscordActions.svelte';
  import DiscordEventRows from './DiscordEventRows.svelte';
  import DiscordReplace from './DiscordReplace.svelte';
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
</script>

{#snippet eventRows()}
  <DiscordEventRows
    events={discord.events}
    disabled={busy !== null || state === 'lapsed'}
    busy={busy === 'events'}
    showNote={state !== 'lapsed'}
    onToggle={onToggleEvents}
  />
{/snippet}

<DiscordSummary
  {state}
  {discord}
  replaceHidden={confirming}
  replaceDisabled={busy !== null}
  {onReplace}
  {onStuck}
/>
{#if state !== 'disabled'}{@render eventRows()}{/if}
<!-- 8 px under the actions row; the replace link's own 44 px hit area is the rest -->
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
  <DiscordStatusLine {status} />
  {#if state === 'saved' && !confirming}
    <DiscordReplace {discord} disabled={busy !== null} {onReplace} {onStuck} />
  {/if}
</div>
{#if state === 'disabled'}{@render eventRows()}{/if}
