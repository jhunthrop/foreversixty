<!-- web/src/components/account/AccountDiscordLink.svelte -->
<!-- The You panel's Discord link row (experience spec 2026-10-27-discord-bot §3.2): one of
     unlinked (with a lead, the button a fill only when a guild of the account signs up in
     Discord), linked (the display name and Unlink, which asks in place) or absent when the
     deployment has the bot off. Draws from the `/v1/me` the page already holds. Every sentence
     comes from the copy module; the callback's outcome is read from `?discord_link=` once. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { fetchMeFresh, type Me } from '../../lib/account/api';
  import { discordLinkCopy } from '../../lib/account/discord-link-copy';
  import { discordLinkStartUrl, unlinkDiscord } from '../../lib/account/discord-link-api';
  import {
    discordLinkReturn,
    LINK_BUTTON_SELECTOR,
    takeDiscordLinkFlag,
    UNLINK_BUTTON_SELECTOR,
    type DiscordLinkOutcome,
  } from '../../lib/account/discord-link-return';
  import { ACCOUNT_DISCORD_LINK_RESERVE } from '../../lib/account/layout';
  import { createFocusQueue } from '../../lib/account/discord-return.svelte';
  import { PRIMARY_BUTTON_FIXED, SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import DiscordConnectLink from './discord/DiscordConnectLink.svelte';

  /** `Unlinked.` stays this long; a return's line stays until the next action. */
  const STATUS_MS = 4000;
  /** The longest the row waits for the fresh `/v1/me` after a linked return. */
  const SETTLE_MS = 4000;
  /** The display name is shown at most this long. */
  const NAME_MAX = 32;
  const ANCHOR = '#you';
  const CANCEL = '[data-testid="discord-unlink-cancel"]';
  const GAP_PX = 16;
  const OUTLINE = `${SECONDARY_BUTTON_FIXED} border-line-warm text-text px-4`;

  interface Line {
    message: string;
    role: 'status' | 'alert';
    help?: string;
  }

  let { me }: { me: Me } = $props();

  const linkedName = $derived(me.discord_link?.display_name ?? null);
  const suggested = $derived(me.discord_link_suggested ?? null);
  let confirming = $state(false);
  let busy = $state(false);
  let line = $state<Line | null>(null);
  let lineTimer: ReturnType<typeof setTimeout> | undefined;
  let settling = $state(false);
  let settleTimer: ReturnType<typeof setTimeout> | undefined;

  const focus = createFocusQueue(() => [linkedName, confirming, settling]);
  $effect(() => () => {
    clearTimeout(lineTimer);
    clearTimeout(settleTimer);
  });

  function show(next: Line | null, expires = false): void {
    clearTimeout(lineTimer);
    line = next;
    if (expires && next !== null) lineTimer = setTimeout(() => (line = null), STATUS_MS);
  }

  function reveal(outcome: DiscordLinkOutcome): void {
    clearTimeout(settleTimer);
    settling = false;
    show({ message: outcome.message, role: outcome.role, help: outcome.help });
    focus.request(outcome.focus);
  }

  /** Reads the callback's outcome once. `start` leaves for the start route; `linked` holds for the fresh `me`. */
  onMount(() => {
    scrollToAnchor();
    const flag = takeDiscordLinkFlag();
    if (flag === null) return;
    const outcome = discordLinkReturn(flag);
    if (outcome.navigate) {
      window.location.assign(discordLinkStartUrl());
      return;
    }
    if (!outcome.refresh) return reveal(outcome);
    settling = true;
    // A link that never shows up in `me` is a failed link: Unlink would never exist to take focus.
    const settle = (): void => reveal(linkedName !== null ? outcome : discordLinkReturn('failed'));
    settleTimer = setTimeout(settle, SETTLE_MS);
    fetchMeFresh().then(
      (fresh) => reveal(fresh?.discord_link ? outcome : discordLinkReturn('failed')),
      settle,
    );
  });

  /** The panel renders after the browser's own fragment scroll has run, so `#you` is honoured here. */
  function scrollToAnchor(): void {
    if (window.location.hash === ANCHOR) document.querySelector(ANCHOR)?.scrollIntoView();
  }

  function showStuck(): void {
    show({ message: discordLinkCopy.stuck, role: 'status' });
  }

  function askUnlink(): void {
    show(null);
    confirming = true;
    focus.request(CANCEL, { visible: true });
  }

  function cancel(): void {
    confirming = false;
    focus.request(UNLINK_BUTTON_SELECTOR);
  }

  async function unlink(): Promise<void> {
    busy = true;
    show(null);
    try {
      await unlinkDiscord(me);
      confirming = false;
      show({ message: discordLinkCopy.unlinked, role: 'status' }, true);
      focus.request(LINK_BUTTON_SELECTOR);
    } catch {
      show({ message: discordLinkCopy.unlinkFailed, role: 'alert' });
    } finally {
      busy = false;
    }
  }

  function onKeydown(event: KeyboardEvent): void {
    if (event.key === 'Escape' && confirming && !busy) cancel();
  }

  const linkClass = $derived(suggested === null ? OUTLINE : `${PRIMARY_BUTTON_FIXED} px-5`);
  const settledMinHeight = ACCOUNT_DISCORD_LINK_RESERVE.linked.lg - GAP_PX;
</script>

{#if me.discord_link_available === true}
  <div
    class="flex flex-col items-start gap-2"
    onkeydown={onKeydown}
    role="presentation"
    data-testid="account-discord-link"
  >
    {#if settling}
      <div style="min-height: {settledMinHeight}px" aria-hidden="true"></div>
    {:else if linkedName !== null}
      <p class="text-[14px] [overflow-wrap:anywhere]" data-testid="discord-linked-as">
        <span class="text-muted">{discordLinkCopy.linkedLead}</span>
        <strong class="text-strong font-semibold">{Array.from(linkedName).slice(0, NAME_MAX).join('')}</strong
        >
      </p>
      {#if confirming}
        <p class="text-text text-[14px]" data-testid="discord-unlink-question">
          {discordLinkCopy.unlinkQuestion}
        </p>
        <div class="flex flex-wrap items-center gap-3">
          <button
            type="button"
            class={OUTLINE}
            disabled={busy}
            onclick={unlink}
            data-testid="discord-unlink-confirm"
          >
            {busy ? discordLinkCopy.unlinking : discordLinkCopy.unlink}
          </button>
          <button
            type="button"
            class={OUTLINE}
            disabled={busy}
            onclick={cancel}
            data-testid="discord-unlink-cancel"
          >
            {discordLinkCopy.cancel}
          </button>
        </div>
      {:else}
        <button type="button" class={OUTLINE} onclick={askUnlink} data-testid="discord-unlink">
          {discordLinkCopy.unlink}
        </button>
      {/if}
    {:else}
      <p class={suggested === null ? 'text-muted text-[13px]' : 'text-text text-[14px]'}>
        {suggested === null ? discordLinkCopy.lead : discordLinkCopy.leadSuggested(suggested)}
      </p>
      <DiscordConnectLink
        label={discordLinkCopy.link}
        class={linkClass}
        testid="discord-link-start"
        href={discordLinkStartUrl()}
        onStuck={showStuck}
      />
    {/if}
    <div class="min-h-[21px] text-[13px]" role={line?.role ?? 'status'} data-testid="discord-link-status">
      {#if line !== null}
        <p class={line.role === 'alert' ? 'text-wipe' : 'text-text'}>{line.message}</p>
        {#if line.help !== undefined}
          <p class="text-muted mt-1 text-[13px]">{line.help}</p>
        {/if}
      {/if}
    </div>
  </div>
{/if}
