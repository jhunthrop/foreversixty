<!-- web/src/components/account/AccountDiscordLink.svelte -->
<!-- The Discord section's link row (experience spec 2026-10-27-discord-bot §3.2; premium spec
     2026-11-08 section 7.6): one Row, unlinked (the lead and Link Discord, a fill only when a guild
     of the account signs up in Discord), linked (the display name and Unlink, which asks in the row:
     danger Unlink and Cancel, focus on Cancel, Escape cancels) or absent when the deployment has the
     bot off. Draws from the `/v1/me` the page already holds. Every sentence
     comes from the copy module; the callback's outcome is read from `?discord_link=` once. -->
<script lang="ts">
  import { smallButton } from '../../lib/account/small-button';
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
  import { createFocusQueue } from '../../lib/account/discord-return.svelte';
  import Row from '../ui/Row.svelte';
  import ConfirmButtons from '../ui/ConfirmButtons.svelte';
  import DiscordConnectLink from './discord/DiscordConnectLink.svelte';

  /** `Unlinked.` stays this long; a return's line stays until the next action. */
  const STATUS_MS = 4000;
  /** The longest the row waits for the fresh `/v1/me` after a linked return. */
  const SETTLE_MS = 4000;
  /** The display name is shown at most this long. */
  const NAME_MAX = 32;
  /** The bot's callback still lands on `#you` (api discordbot); the row now lives under `#discord`. */
  const ANCHORS = ['#you', '#discord'];
  const SECTION = '#discord';
  const CANCEL = '[data-testid="discord-unlink-cancel"]';

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

  /** The section renders after the browser's own fragment scroll has run, so the anchor is honoured here. */
  function scrollToAnchor(): void {
    if (ANCHORS.includes(window.location.hash)) document.querySelector(SECTION)?.scrollIntoView();
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

  const linkClass = $derived(smallButton(suggested === null ? 'secondary' : 'primary'));
  const shownName = $derived(linkedName === null ? '' : Array.from(linkedName).slice(0, NAME_MAX).join(''));
</script>

{#if me.discord_link_available === true}
  <div data-testid="account-discord-link">
    {#if settling}
      <div class="min-h-[var(--row-h)]" aria-hidden="true"></div>
    {:else if linkedName !== null && confirming}
      <Row title={discordLinkCopy.unlinkQuestion} open wrapTrail>
        {#snippet titleContent()}<span data-testid="discord-unlink-question"
            >{discordLinkCopy.unlinkQuestion}</span
          >{/snippet}
        {#snippet trail()}
          <ConfirmButtons
            confirm={busy ? discordLinkCopy.unlinking : discordLinkCopy.unlink}
            cancel={discordLinkCopy.cancel}
            {busy}
            onconfirm={() => void unlink()}
            oncancel={cancel}
            testid="discord-unlink"
          />
        {/snippet}
      </Row>
    {:else if linkedName !== null}
      <Row title={`${discordLinkCopy.linkedLead} ${shownName}`} open wrapTrail testid="discord-linked-as">
        {#snippet titleContent()}
          <span class="[overflow-wrap:anywhere]">{discordLinkCopy.linkedLead} <b>{shownName}</b></span>
        {/snippet}
        {#snippet trail()}
          <button type="button" class={smallButton()} onclick={askUnlink} data-testid="discord-unlink"
            >{discordLinkCopy.unlink}</button
          >
        {/snippet}
      </Row>
    {:else}
      <Row
        title={suggested === null ? discordLinkCopy.lead : discordLinkCopy.leadSuggested(suggested)}
        open
        wrapTrail
      >
        {#snippet trail()}
          <DiscordConnectLink
            label={discordLinkCopy.link}
            class={linkClass}
            testid="discord-link-start"
            href={discordLinkStartUrl()}
            onStuck={showStuck}
          />
        {/snippet}
      </Row>
    {/if}
    {#if line !== null}
      <div class="pt-2 text-[13px]" role={line.role} data-testid="discord-link-status">
        <p class={line.role === 'alert' ? 'text-wipe' : 'text-text'}>{line.message}</p>
        {#if line.help !== undefined}<p class="text-muted mt-1">{line.help}</p>{/if}
      </div>
    {/if}
  </div>
{/if}
