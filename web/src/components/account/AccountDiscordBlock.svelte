<!-- web/src/components/account/AccountDiscordBlock.svelte -->
<!-- The account page's Discord block (experience spec 2026-10-25-discord-notifications, sections
     3 to 5): one of five states chosen by blockState(me), drawn from the `/v1/me` the page
     already holds. Flows live in lib/account/discord-actions.ts; every sentence comes from the
     copy module. The typed webhook URL is component state only: cleared after a saved webhook,
     kept after a failed one, never written anywhere else. -->
<script lang="ts">
  import { tick } from 'svelte';
  import type { DiscordEvents, DiscordNotifications, Me } from '../../lib/account/api';
  import {
    removeWebhook,
    saveEvents,
    saveWebhook,
    sendTest,
    type ActionResult,
  } from '../../lib/account/discord-actions';
  import { discordFocus } from '../../lib/account/discord-focus';
  import { LINK_BUTTON } from '../../lib/account/discord-styles';
  import { blockState } from '../../lib/account/notifications';
  import { notificationsCopy } from '../../lib/account/notifications-copy';
  import StatePanel from '../ui/StatePanel.svelte';
  import DiscordEventRows from './discord/DiscordEventRows.svelte';
  import DiscordActions from './discord/DiscordActions.svelte';
  import DiscordFree from './discord/DiscordFree.svelte';
  import DiscordPasteForm from './discord/DiscordPasteForm.svelte';
  import DiscordSummary from './discord/DiscordSummary.svelte';
  import DiscordStatusLine from './discord/DiscordStatusLine.svelte';

  const STATUS_MS = 4000;
  const BOTH_ON: DiscordEvents = { sim_finished: true, upload_parsed: true };

  const INPUT = '#discord-webhook';
  const TEST_BUTTON = '[data-testid="discord-test"]';
  const REMOVE_CANCEL = '[data-testid="discord-remove-cancel"]';
  const FREE_LINK = '[data-testid="discord-free-link"]';

  type Busy = 'save' | 'events' | 'test' | 'remove' | null;

  /** `me` is the page's loaded `/v1/me`; `onChanged` asks the page to re-read it after a change. */
  let { me, onChanged = () => {} }: { me: Me; onChanged?: () => void } = $props();

  const block = $derived(blockState(me));
  let webhook = $state('');
  let formEvents = $state<DiscordEvents>(BOTH_ON);
  let replacing = $state(false);
  let confirming = $state(false);
  let busy = $state<Busy>(null);
  let status = $state<ActionResult | null>(null);
  let fieldFailure = $state<Extract<ActionResult, { ok: false }> | null>(null);
  // Applied by the effect below on the first render where it resolves: a save or remove changes
  // the state only after the page re-reads /v1/me, so the target does not exist yet.
  let pendingFocus = $state<string | null>(discordFocus.get());
  function wantFocus(selector: string | null): void {
    pendingFocus = selector;
    discordFocus.set(selector);
  }
  let statusTimer: ReturnType<typeof setTimeout> | undefined;

  const showForm = $derived(block.state === 'empty' || (replacing && block.state !== 'lapsed'));

  $effect(() => () => clearTimeout(statusTimer));
  $effect(() => {
    const selector = pendingFocus;
    if (selector === null) return;
    void [block.state, showForm, confirming, busy];
    void tick().then(() => {
      const target = document.querySelector<HTMLElement>(selector);
      if (target === null) return;
      target.focus();
      if (pendingFocus === selector) wantFocus(null);
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
    wantFocus(null);
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
      wantFocus(INPUT);
      return;
    }
    webhook = '';
    replacing = false;
    wantFocus(TEST_BUTTON);
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
    wantFocus(block.state === 'lapsed' ? FREE_LINK : INPUT);
  }

  function startReplace(): void {
    if ('discord' in block) formEvents = block.discord.events;
    fieldFailure = null;
    status = null;
    replacing = true;
    wantFocus(INPUT);
  }

  function askRemove(): void {
    confirming = true;
    wantFocus(REMOVE_CANCEL);
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

{#snippet eventRows(discord: DiscordNotifications)}
  <DiscordEventRows
    events={discord.events}
    disabled={busy !== null || block.state === 'lapsed'}
    busy={busy === 'events'}
    showNote={block.state !== 'lapsed'}
    onToggle={changeEvents}
  />
{/snippet}

<div id="discord" class="scroll-mt-20">
  {#if block.state === 'free'}
    <DiscordFree />
  {:else}
    <StatePanel label={notificationsCopy.label} testid="account-discord">
      {#if showForm}
        {#if block.state === 'empty' && status !== null}<DiscordStatusLine {status} />{/if}
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
      {:else if 'discord' in block}
        {@const discord = block.discord}
        <DiscordSummary state={block.state} {discord} />
        {#if block.state !== 'disabled'}{@render eventRows(discord)}{/if}
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
        {#if block.state === 'disabled'}{@render eventRows(discord)}{/if}
      {/if}
    </StatePanel>
  {/if}
</div>
