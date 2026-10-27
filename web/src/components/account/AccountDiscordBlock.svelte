<!-- web/src/components/account/AccountDiscordBlock.svelte -->
<!-- The account page's Discord block (experience spec 2026-10-25-discord-notifications, sections
     3 to 5): one of five states chosen by blockState(me), drawn from the `/v1/me` the page
     already holds. Flows live in lib/account/discord-actions.ts; every sentence comes from the
     copy module. The typed webhook URL is component state only: cleared after a saved webhook,
     kept after a failed one, never written anywhere else. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import type { DiscordEvents, Me } from '../../lib/account/api';
  import {
    awaitsConnectRefresh,
    NEW_WEBHOOK_EVENTS,
    removeWebhook,
    saveEvents,
    saveWebhook,
    sendTest,
    type ActionResult,
    type ReturnOutcome,
  } from '../../lib/account/discord-actions';
  import { createFocusQueue, consumeReturn } from '../../lib/account/discord-return.svelte';
  import { blockState, connectAvailable } from '../../lib/account/notifications';
  import { notificationsCopy } from '../../lib/account/notifications-copy';
  import StatePanel from '../ui/StatePanel.svelte';
  import DiscordConnect from './discord/DiscordConnect.svelte';
  import DiscordFree from './discord/DiscordFree.svelte';
  import DiscordPasteForm from './discord/DiscordPasteForm.svelte';
  import DiscordSaved from './discord/DiscordSaved.svelte';
  import DiscordStatusLine from './discord/DiscordStatusLine.svelte';

  // How long `Saved.` and `Removed.` stay; a result with `stays` holds until the next press (spec 4.3).
  const STATUS_MS = 4000;
  // The longest the block waits for the `connected` refresh before drawing from the `me` it has.
  const CONNECT_SETTLE_MS = 4000;

  const INPUT = '#discord-webhook';
  const CONNECT = '[data-testid="discord-connect"]';
  const TEST_BUTTON = '[data-testid="discord-test"]';
  const REMOVE_CANCEL = '[data-testid="discord-remove-cancel"]';
  const FREE_LINK = '[data-testid="discord-free-link"]';

  type Busy = 'save' | 'events' | 'test' | 'remove' | null;

  /** `me` is the page's loaded `/v1/me`; a change patches the cache, so `me` updates in place. */
  let {
    me,
    quiet = false,
  }: {
    me: Me;
    /** Another block holds the page's one fill (a seat offer): this block draws no fill of its own. */
    quiet?: boolean;
  } = $props();

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
  let statusTimer: ReturnType<typeof setTimeout> | undefined;

  const connectOn = $derived(connectAvailable(me));
  const showForm = $derived(block.state === 'empty' || (replacing && block.state !== 'lapsed'));
  const showConnect = $derived(showForm && connectOn && !pasteView);

  $effect(() => () => clearTimeout(statusTimer));

  const focus = createFocusQueue(() => [block.state, showForm, confirming, busy]);

  // After a `connected` return the held `me` predates the webhook: draw nothing (the reserve holds
  // the space) until the fresh `me` carries it, so the empty state never flashes (spec 4.3).
  let settling = $state(typeof window !== 'undefined' && awaitsConnectRefresh(window.location.search));
  let waitingFor = $state<ReturnOutcome | null>(null);
  let settleTimer: ReturnType<typeof setTimeout> | undefined;

  function reveal(outcome: ReturnOutcome): void {
    clearTimeout(settleTimer);
    waitingFor = null;
    settling = false;
    status = outcome.status;
    focus.request(outcome.focus);
  }

  // The callback redirects here with ?discord=<outcome>: show it once, move focus, drop the flag.
  onMount(() => {
    let outcome: ReturnOutcome | null = null;
    outcome = consumeReturn(() => outcome !== null && reveal(outcome));
    if (outcome === null) {
      settling = false;
      return;
    }
    if (!outcome.refresh) return reveal(outcome);
    const waiting = outcome;
    waitingFor = waiting;
    settleTimer = setTimeout(() => reveal(waiting), CONNECT_SETTLE_MS);
    return () => clearTimeout(settleTimer);
  });

  $effect(() => {
    if (waitingFor !== null && 'discord' in block) reveal(waitingFor);
  });

  function show(result: ActionResult): void {
    clearTimeout(statusTimer);
    status = result;
    if (result.ok && result.stays !== true) statusTimer = setTimeout(() => (status = null), STATUS_MS);
  }

  /** The connect click never navigated: say so in the panel's status line (spec 4.3 item 5). */
  function showStuck(): void {
    show({ ok: true, message: notificationsCopy.connectStuck });
  }

  async function run(kind: Exclude<Busy, null>, action: () => Promise<ActionResult>): Promise<ActionResult> {
    busy = kind;
    status = null;
    focus.request(null);
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
      focus.request(INPUT);
      return;
    }
    webhook = '';
    replacing = false;
    pasteView = false;
    focus.request(TEST_BUTTON);
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
    focus.request(afterRemove);
  }

  function startReplace(): void {
    if ('discord' in block) formEvents = block.discord.events;
    fieldFailure = null;
    status = null;
    replacing = true;
    pasteView = false;
    focus.request(connectOn ? CONNECT : INPUT);
  }

  function showPaste(): void {
    status = null;
    pasteView = true;
    focus.request(INPUT);
  }

  function showConnectView(): void {
    status = null;
    pasteView = false;
    webhook = '';
    fieldFailure = null;
    focus.request(CONNECT);
  }

  function askRemove(): void {
    status = null;
    confirming = true;
    focus.request(REMOVE_CANCEL);
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

{#if block.state !== 'off' && !settling}
  <div id="discord" class="scroll-mt-20" onkeydown={onKeydown} role="presentation">
    {#if block.state === 'free'}
      <DiscordFree />
    {:else}
      <StatePanel label={notificationsCopy.label} testid="account-discord">
        {#if showConnect}
          <DiscordConnect {status} {replacing} onStuck={showStuck} onPaste={showPaste} onCancel={cancel} />
        {:else if showForm}
          {#if block.state === 'empty' && status !== null}<DiscordStatusLine {status} />{/if}
          <DiscordPasteForm
            bind:webhook
            bind:events={formEvents}
            error={fieldFailure?.message ?? null}
            offerPlans={fieldFailure?.target === 'plans'}
            saving={busy === 'save'}
            {replacing}
            secondary={quiet}
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
            onStuck={showStuck}
          />
        {/if}
      </StatePanel>
    {/if}
  </div>
{/if}
