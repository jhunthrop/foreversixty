<!-- web/src/components/ui/Card.svelte -->
<!-- The card (site alignment spec 2026-10-29 §3.2 to §3.4, S1): the account card surface (quiet
     drops the gold top edge), a CardHead, one body whose padding follows what it holds (`text`
     16/20 with a top rule, `table` 0 20 8 with a top rule, `rows` and `none` full-bleed: rows draw
     their own rules), and an optional foot. Card.astro is the static twin; both read
     lib/ui/card.ts. A card is never a link as a whole. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { CARD_BODY, CARD_FOOT, cardClass, type CardBody } from '../../lib/ui/card';
  import type { GlyphName } from '../../lib/ui/glyphs';
  import CardHead from './CardHead.svelte';
  import type { TileTone } from './IconTile.svelte';

  let {
    glyph,
    tone = 'gold',
    lead,
    title,
    label,
    titleSnippet,
    meta,
    action,
    aside,
    quiet = false,
    body = 'text',
    foot,
    headingLevel = 2,
    compact = false,
    testid,
    id,
    class: className = '',
    children,
  }: {
    glyph?: GlyphName;
    tone?: TileTone;
    /** Replaces the tile: an item icon, a spec icon, a crest. */
    lead?: Snippet;
    title: string;
    label?: string;
    titleSnippet?: Snippet;
    meta?: string;
    action?: { label: string; href: string };
    /** The head's right slot: a Tag or one small button. */
    aside?: Snippet;
    /** A secondary card (prose aside, table in prose, canvas frame): no gold top edge. */
    quiet?: boolean;
    body?: CardBody;
    foot?: Snippet;
    headingLevel?: 2 | 3;
    /** The 48 px head for dense data cards. */
    compact?: boolean;
    testid?: string;
    id?: string;
    /** Layout classes from the parent (a grid span, a scroll margin). */
    class?: string;
    children: Snippet;
  } = $props();
</script>

<section {id} class={[cardClass(quiet), className]} data-testid={testid}>
  <CardHead
    {glyph}
    {tone}
    {lead}
    {label}
    {title}
    {titleSnippet}
    {meta}
    {action}
    {aside}
    {headingLevel}
    {compact}
  />
  {#if body === 'none'}
    {@render children()}
  {:else}
    <div class={CARD_BODY[body]} data-card-body={body}>{@render children()}</div>
  {/if}
  {#if foot !== undefined}<div class={CARD_FOOT} data-card-foot>{@render foot()}</div>{/if}
</section>
