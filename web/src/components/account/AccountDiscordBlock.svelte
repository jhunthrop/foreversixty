<!-- web/src/components/account/AccountDiscordBlock.svelte -->
<!-- The account page's Discord block (experience spec 2026-10-25-discord-notifications, sections
     3 to 5): one of five states chosen by blockState(me), drawn from the `/v1/me` the page
     already holds. Flows live in lib/account/discord-actions.ts; every sentence comes from the
     copy module. The typed webhook URL is component state only: cleared after a saved webhook,
     kept after a failed one, never written anywhere else. -->
<script lang="ts">
  import { tick } from 'svelte';
  import type { DiscordEvents, Me } from '../../lib/account/api';
  import {
    removeWebhook,
    saveEvents,
    saveWebhook,
    sendTest,
    type ActionResult,
  } from '../../lib/account/discord-actions';
  import { LINK_BUTTON } from '../../lib/account/discord-styles';
  import { blockState } from '../../lib/account/notifications';
  import { notificationsCopy } from '../../lib/account/notifications-copy';
  import StatePanel from '../ui/StatePanel.svelte';
  import DiscordEventRows from './discord/DiscordEventRows.svelte';
  import DiscordActions from './discord/DiscordActions.svelte';
  import DiscordPasteForm from './discord/DiscordPasteForm.svelte';
  import DiscordSummary from './discord/DiscordSummary.svelte';
  import DiscordStatusLine from './discord/DiscordStatusLine.svelte';

  const STATUS_MS = 4000;
  const BOTH_ON: DiscordEvents = { sim_finished: true, upload_parsed: true };

  type Busy = 'save' | 'events' | 'test' | 'remove' | null;

  /** `me` is the page's loaded `/v1/me`; `onChanged` asks the page to re-read it after a change. */
  let { me, onChanged = () => {} }: { me: Me; onChanged?: () => void } = $props();

  const block = $derived(blockState(me));
  let webhook = $state('');
  let formEvents = $state<DiscordEvents>(BOTH_ON);
  let replacing = $state(false);
  let confirming = $state(false);
  let busy = $state<Busy>(null);
  let status = $state<{ ok: boolean; message: string } | null>(null);
  let fieldFailure = $state<Extract<ActionResult, { ok: false }> | null>(null);
  let focusSelector = $state<string | null>(null);
  let statusTimer: ReturnType<typeof setTimeout> | undefined;

  const showForm = $derived(block.state === 'empty' || (replacing && block.state !== 'lapsed'));

  $effect(() => () => clearTimeout(statusTimer));
  $effect(() => {
    if (focusSelector === null) return;
    void showForm;
    void tick().then(() => {
      document.querySelector<HTMLElement>(focusSelector!)?.focus();
      focusSelector = null;
    });
  });

  function show(result: ActionResult): void {
    clearTimeout(statusTimer);
    status = result;
    if (result.ok) statusTimer = setTimeout(() => (status = null), STATUS_MS);
  }

  async function run(kind: Exclude<Busy, null>, action: () => Promise<ActionResult>): Promise<ActionResult> {
    busy = kind;
    status = null;
    const result = await action();
    busy = null;
    if (result.ok) onChanged();
    return result;
  }

  async function submit(): Promise<void> {
    fieldFailure = null;
    const result = await run('save', () => saveWebhook(webhook, formEvents));
    if (!result.ok) {
      if (result.target === 'status') show(result);
      else fieldFailure = result;
      return;
    }
    webhook = '';
    replacing = false;
    focusSelector = '[data-testid="discord-test"]';
    show(result);
  }

  async function changeEvents(next: DiscordEvents): Promise<boolean> {
    const result = await run('events', () => saveEvents(next));
    show(result);
    return result.ok;
  }

  async function test(): Promise<void> {
    show(await run('test', () => sendTest(block.state === 'disabled')));
  }

  async function remove(): Promise<void> {
    const result = await run('remove', removeWebhook);
    show(result);
    if (!result.ok) return;
    confirming = false;
    focusSelector = '#discord-webhook';
  }

  function startReplace(): void {
    if ('discord' in block) formEvents = block.discord.events;
    fieldFailure = null;
    status = null;
    replacing = true;
    focusSelector = '#discord-webhook';
  }

  function askRemove(): void {
    confirming = true;
    focusSelector = '[data-testid="discord-remove-cancel"]';
  }

  function cancel(): void {
    replacing = false;
    confirming = false;
    webhook = '';
    fieldFailure = null;
  }

  function onKeydown(event: KeyboardEvent): void {
    if (event.key === 'Escape' && (replacing || confirming) && busy === null) cancel();
  }
</script>

<svelte:window onkeydown={onKeydown} />

<div id="discord" class="scroll-mt-20">
  {#if block.state === 'free'}
    <p class="text-muted min-h-11 text-[13px]" data-testid="discord-free">
      {notificationsCopy.freeSentence}
      <a class="text-text hover:text-gold-hover inline-flex min-h-11 items-center underline" href="/premium"
        >{notificationsCopy.seePlans}</a
      >
    </p>
  {:else}
    <StatePanel label={notificationsCopy.label} testid="account-discord">
      {#if showForm}
        <DiscordPasteForm
          bind:webhook
          bind:events={formEvents}
          error={fieldFailure?.message ?? null}
          offerPlans={fieldFailure?.target === 'plans'}
          saving={busy === 'save'}
          {replacing}
          onSubmit={submit}
          onCancel={cancel}
        />
        {#if block.state === 'empty'}<DiscordStatusLine {status} />{/if}
      {:else if 'discord' in block}
        {@const discord = block.discord}
        <DiscordSummary state={block.state} {discord} />
        <DiscordEventRows
          events={discord.events}
          disabled={busy !== null || block.state === 'lapsed'}
          busy={busy === 'events'}
          showNote={block.state !== 'lapsed'}
          onToggle={changeEvents}
        />
        <DiscordActions
          state={block.state}
          {confirming}
          {busy}
          onTest={test}
          onAskRemove={askRemove}
          onRemove={remove}
          onCancel={cancel}
        />
        <DiscordStatusLine {status} />
        {#if block.state !== 'lapsed'}
          <button
            class="{LINK_BUTTON} w-fit"
            disabled={busy !== null}
            onclick={startReplace}
            data-testid="discord-replace"
          >
            {notificationsCopy.replace}
          </button>
        {/if}
      {/if}
    </StatePanel>
  {/if}
</div>
