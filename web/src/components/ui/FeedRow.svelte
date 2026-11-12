<!-- web/src/components/ui/FeedRow.svelte -->
<!-- One feed line (home B spec 2026-11-07 P11): a crest, item or spell icon (from `lead`), a sentence in the body face with
     its names in their colours, and at the right the time and a caps kind (under the sentence on a
     phone). A link to the thing it
     is about, on a --color-line-soft hairline, no box; hover is the S2 ground. `fresh` slides the
     newest line in when the reader allows motion. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { FeedLead } from '../../lib/home/feed';
  import { iconUrl } from '../../lib/ui/assets';
  import Crest from './Crest.svelte';
  import ItemIcon from './ItemIcon.svelte';
  import SpellIcon from './SpellIcon.svelte';

  let {
    href,
    when,
    kind,
    fresh = false,
    lead,
    build,
    leadSize = 32,
    testid,
    children,
  }: {
    href: string;
    when: string;
    /** The row's tag under the time ("Your guild"); absent, the time alone. */
    kind?: string;
    fresh?: boolean;
    /** The row's crest, item or spell icon, as data (the feed row's own fields). */
    lead: FeedLead;
    /** The client build the icon files are under. */
    build: string;
    /** The lead's size; the column is 40 px. */
    leadSize?: number;
    testid?: string;
    /** The sentence. */
    children: Snippet;
  } = $props();
</script>

<a class="ui-feed-row" class:fresh {href} data-testid={testid}>
  <span class="ui-feed-lead"
    >{#if lead.crest !== null}<Crest
        cls={lead.crest}
        size={leadSize}
      />{:else if lead.quality !== null}<ItemIcon
        icon={lead.icon ?? undefined}
        quality={lead.quality}
        {build}
        size={leadSize}
      />{:else}<SpellIcon url={iconUrl(build, lead.icon)} size={leadSize} />{/if}</span
  >
  <span class="ui-feed-text">{@render children()}</span>
  <span class="ui-feed-when"
    >{when}{#if kind}<span class="ui-feed-kind">{kind}</span>{/if}</span
  >
</a>

<style>
  .ui-feed-row {
    display: grid;
    grid-template-columns: 40px minmax(0, 1fr) auto;
    gap: 14px;
    align-items: center;
    padding: 12px 0;
    border-bottom: 1px solid var(--color-line-soft);
    font-size: 15px;
    line-height: 22px;
    color: var(--color-text, inherit);
  }
  .ui-feed-row:hover {
    background: rgb(255 255 255 / 0.04);
  }
  .ui-feed-row:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: -2px;
  }
  .ui-feed-lead {
    display: flex;
    line-height: 0;
  }
  .ui-feed-text {
    min-width: 0;
  }
  .ui-feed-when {
    text-align: right;
    font-size: 13px;
    line-height: 20px;
    white-space: nowrap;
    color: var(--color-muted);
  }
  .ui-feed-kind {
    display: block;
    font-size: 12px;
    line-height: 16px;
    letter-spacing: 0;
    font-weight: 700;
    color: var(--color-muted);
  }
  /* A phone: the time drops under the sentence, which keeps the whole width. */
  @media (max-width: 559px) {
    .ui-feed-row {
      grid-template-columns: 40px minmax(0, 1fr);
      row-gap: 2px;
    }
    .ui-feed-lead {
      grid-row: span 2;
    }
    .ui-feed-when {
      text-align: left;
    }
  }
  @media (prefers-reduced-motion: no-preference) {
    .ui-feed-row.fresh {
      animation: ui-feed-in 0.6s ease-out;
    }
  }
  @keyframes ui-feed-in {
    from {
      opacity: 0;
      transform: translateY(-6px);
    }
    to {
      opacity: 1;
      transform: none;
    }
  }
</style>
