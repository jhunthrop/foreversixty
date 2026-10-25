<!-- web/src/components/account/AccountDiscordBlock.svelte -->
<!-- The account page's Discord block (experience spec 2026-10-25-discord-notifications, sections
     3 to 5): one of five states chosen by blockState(me), drawn from the `/v1/me` the page
     already holds. Flows live in lib/account/discord-actions.ts; every sentence comes from the
     copy module. The typed webhook URL is component state only: cleared after a saved webhook,
     kept after a failed one, never written anywhere else. -->
<script lang="ts">
  import { onMount, tick } from 'svelte';
  import type { DiscordEvents, Me } from '../../lib/account/api';
  import {
    refreshAfterConnect,
    removeWebhook,
    returnOutcome,
    saveEvents,
    saveWebhook,
    sendTest,
    takeReturnFlag,
    type ActionResult,
  } from '../../lib/account/discord-actions';
  import { blockState, connectAvailable } from '../../lib/account/notifications';
  import { notificationsCopy } from '../../lib/account/notifications-copy';
  import StatePanel from '../ui/StatePanel.svelte';
  import DiscordConnect from './discord/DiscordConnect.svelte';
  import DiscordFree from './discord/DiscordFree.svelte';
  import DiscordPasteForm from './discord/DiscordPasteForm.svelte';
  import DiscordSaved from './discord/DiscordSaved.svelte';
  import DiscordStatusLine from './discord/DiscordStatusLine.svelte';

  const STATUS_MS = 4000;
  const NEW_WEBHOOK_EVENTS: DiscordEvents = {
    sim_finished: true,
    upload_parsed: true,
    upload_private: false,
  };

  const INPUT = '#discord-webhook';
  const CONNECT = '[data-testid="discord-connect"]';
  const TEST_BUTTON = '[data-testid="discord-test"]';
  const REMOVE_CANCEL = '[data-testid="discord-remove-cancel"]';
  const FREE_LINK = '[data-testid="discord-free-link"]';

  type Busy = 'save' | 'events' | 'test' | 'remove' | null;

  /** `me` is the page's loaded `/v1/me`; a change patches the cache, so `me` updates in place. */
  let { me }: { me: Me } = $props();

  const block = $derived(blockState(me));
  let webhook = $state('');
  let formEvents = $state<DiscordEvents>(NEW_WEBHOOK_EVENTS);
  let replacing = $state(false);
  // The member's swap from the Connect view to the paste view; only meaningful where connect is on.
  let pasteView = $state(false);
  let confirming = $state(false);
  let busy = $state<Busy>(null);
  let status = $state<ActionResult | null>(null);
  let fieldFailure = $state<Extract<ActionResult, { ok: false }> | null>(null);
  // Applied by the effect below on the first render where it resolves: a save or remove changes
  // the state when the patched `me` arrives, so the target may not exist yet.
  let pendingFocus = $state<string | null>(null);
  let statusTimer: ReturnType<typeof setTimeout> | undefined;

  const connectOn = $derived(connectAvailable(me));
  const showForm = $derived(block.state === 'empty' || (replacing && block.state !== 'lapsed'));
  const showConnect = $derived(showForm && connectOn && !pasteView);

  $effect(() => () => clearTimeout(statusTimer));

  // The callback redirects here with ?discord=<outcome>: show it once, move focus, drop the flag.
  onMount(() => {
    const flag = takeReturnFlag(window.location, window.history);
    if (flag === null) return;
    const outcome = returnOutcome(flag);
    status = outcome.status;
    pendingFocus = outcome.focus;
    if (outcome.refresh) refreshAfterConnect();
  });
  $effect(() => {
    const selector = pendingFocus;
    if (selector === null) return;
    void [block.state, showForm, confirming, busy];
    void tick().then(() => {
      const target = document.querySelector<HTMLElement>(selector);
      if (target === null) return;
      target.focus();
      if (pendingFocus === selector) pendingFocus = null;
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
    pendingFocus = null;
    const result = await action();
    busy = null;
    return result;
  }

  async function submit(): Promise<void> {
    fieldFailure = null;
    const result = await run('save', () => saveWebhook(me, webhook, formEvents));
    if (!result.ok) {
      if (result.target === 'status') show(result);
      else fieldFailure = result;
      pendingFocus = INPUT;
      return;
    }
    webhook = '';
    replacing = false;
    pasteView = false;
    pendingFocus = TEST_BUTTON;
    show(result);
  }

  async function changeEvents(next: DiscordEvents): Promise<boolean> {
    const result = await run('events', () => saveEvents(me, next));
    show(result);
    return result.ok;
  }

  async function test(): Promise<void> {
    show(await run('test', () => sendTest(me, block.state === 'disabled')));
  }

  async function remove(): Promise<void> {
    const afterRemove = block.state === 'lapsed' ? FREE_LINK : connectOn ? CONNECT : INPUT; // the patched `me` flips the state
    const result = await run('remove', () => removeWebhook(me));
    show(result);
    if (!result.ok) return;
    confirming = false;
    pasteView = false;
    pendingFocus = afterRemove;
  }

  function startReplace(): void {
    if ('discord' in block) formEvents = block.discord.events;
    fieldFailure = null;
    status = null;
    replacing = true;
    pasteView = false;
    pendingFocus = connectOn ? CONNECT : INPUT;
  }

  function showPaste(): void {
    pasteView = true;
    pendingFocus = INPUT;
  }

  function showConnectView(): void {
    pasteView = false;
    webhook = '';
    fieldFailure = null;
    pendingFocus = CONNECT;
  }

  function askRemove(): void {
    confirming = true;
    pendingFocus = REMOVE_CANCEL;
  }

  function cancel(): void {
    replacing = false;
    pasteView = false;
    confirming = false;
    webhook = '';
    fieldFailure = null;
  }

  function onKeydown(event: KeyboardEvent): void {
    if (event.key === 'Escape' && (replacing || confirming) && busy === null) cancel();
  }
</script>

{#if block.state !== 'off'}
  <div id="discord" class="scroll-mt-20" onkeydown={onKeydown} role="presentation">
    {#if block.state === 'free'}
      <DiscordFree />
    {:else}
      <StatePanel label={notificationsCopy.label} testid="account-discord">
        {#if showConnect}
          <DiscordConnect {status} {replacing} onPaste={showPaste} onCancel={cancel} />
        {:else if showForm}
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
            onConnect={connectOn ? showConnectView : undefined}
          />
        {:else if 'discord' in block}
          <DiscordSaved
            state={block.state}
            discord={block.discord}
            {confirming}
            {busy}
            {status}
            onToggleEvents={changeEvents}
            onTest={test}
            onAskRemove={askRemove}
            onRemove={remove}
            onCancel={cancel}
            onReplace={startReplace}
          />
        {/if}
      </StatePanel>
    {/if}
  </div>
{/if}
