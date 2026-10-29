<!-- web/src/components/account/discord/DiscordSaved.svelte -->
<!-- A saved webhook's card body (saved, switched off, lapsed): destination, notice, the events
     and the actions, in the spec's per-state order (3.2). Switched off puts the events last and
     the replace row right under the notice; saved keeps the replace row last. Both hide the row
     while the remove question is open. -->
<script lang="ts">
  import type { DiscordEvents, DiscordNotifications } from '../../../lib/account/api';
  import { removeQuestionFor, sourceOf, type DiscordBlockKey } from '../../../lib/account/notifications';
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
<div class="border-line-soft flex flex-col border-t px-5 py-4">
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
</div>
{#if state === 'saved' && !confirming}
  <DiscordReplace variant="row" source={sourceOf(discord)} disabled={busy !== null} {onReplace} {onStuck} />
{/if}
{#if state === 'disabled'}{@render eventRows()}{/if}
