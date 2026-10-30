<!-- web/src/components/ui/CardHead.svelte -->
<!-- The card head (polish spec 4.2 CARD_HEAD, site alignment spec 2026-10-29 §3.3): an IconTile
     36 (or the card's own thing as the `lead`: an item icon, a crest), a micro label over a
     heading-face title, and an optional right slot (a meta line, an ActionLink, or a snippet
     holding a Tag or one small button). The title is the card's heading (`h2`, or `h3` inside a
     section with its own `h2`). Under a micro label the title is the card's value (Plan:
     "Premium") and keeps its copy's case; alone it reads uppercase. `compact` is the 48 px head
     for dense data cards (§3.3.1): a 28 px tile and the 18 px card title. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { CARD_HEAD, cardHeadTileSize, cardHeadTitleClass } from '../../lib/ui/card';
  import type { GlyphName } from '../../lib/ui/glyphs';
  import ActionLink from './ActionLink.svelte';
  import IconTile, { type TileTone } from './IconTile.svelte';

  let {
    glyph,
    tone = 'gold',
    lead,
    label,
    title,
    titleSnippet,
    titleContent,
    titleId,
    titleTestid,
    meta,
    action,
    aside,
    headingLevel = 2,
    compact = false,
    class: className = '',
    testid,
  }: {
    /** The tile's glyph; omitted (and no `lead`), the head has no tile (a prose aside). */
    glyph?: GlyphName;
    tone?: TileTone;
    /** Replaces the tile: the item's own icon or a crest, when the card is about one thing. */
    lead?: Snippet;
    /** The micro label over the title; omitted, the title stands alone. */
    label?: string;
    title: string;
    /** Beside the title (a Founder tag). */
    titleSnippet?: Snippet;
    /** Replaces the title's text inside the heading (an item's name in its quality colour, with its tooltip). */
    titleContent?: Snippet;
    /** The heading's id, for a section's aria-labelledby. */
    titleId?: string;
    /** A test id on the heading itself. */
    titleTestid?: string;
    /** A 13 px muted line at the right ("Updated 3 hours ago", "20 to 29"). */
    meta?: string;
    /** An ActionLink at the right ("See all"). */
    action?: { label: string; href: string };
    /** The right slot for anything richer: a Tag or one small button. */
    aside?: Snippet;
    headingLevel?: 2 | 3;
    compact?: boolean;
    /** Layout classes from the card's parent (a phone fold that hides the head). */
    class?: string;
    testid?: string;
  } = $props();

  const titleClass = $derived(cardHeadTitleClass(compact, label !== undefined));
  const hasRight = $derived(meta !== undefined || action !== undefined || aside !== undefined);
</script>

<div
  class={[compact ? CARD_HEAD.compact : CARD_HEAD.root, className]}
  data-testid={testid}
  data-compact={compact || undefined}
>
  {#if lead !== undefined}
    {@render lead()}
  {:else if glyph !== undefined}
    <IconTile {glyph} {tone} size={cardHeadTileSize(compact)} />
  {/if}
  <div class={CARD_HEAD.titleBlock}>
    {#if label !== undefined && !compact}
      <span class={CARD_HEAD.label}>{label}</span>
    {/if}
    <div class={CARD_HEAD.titleRow}>
      {#if headingLevel === 3}
        <h3 class={titleClass} id={titleId} data-testid={titleTestid}>
          {#if titleContent}{@render titleContent()}{:else}{title}{/if}
        </h3>
      {:else}
        <h2 class={titleClass} id={titleId} data-testid={titleTestid}>
          {#if titleContent}{@render titleContent()}{:else}{title}{/if}
        </h2>
      {/if}
      {@render titleSnippet?.()}
    </div>
  </div>
  {#if hasRight}
    <div class={CARD_HEAD.aside}>
      {#if meta !== undefined}<span class={CARD_HEAD.meta} data-testid="card-head-meta">{meta}</span>{/if}
      {#if action !== undefined}<ActionLink href={action.href}>{action.label}</ActionLink>{/if}
      {@render aside?.()}
    </div>
  {/if}
</div>
