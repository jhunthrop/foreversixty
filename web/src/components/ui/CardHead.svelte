<!-- web/src/components/ui/CardHead.svelte -->
<!-- The card head (polish spec 4.2 CARD_HEAD): an IconTile 36, a micro label over a heading-face
     title, and an optional right slot (a Tag or a 13 px meta). The title is the card's heading
     (`h2`), so the page keeps its outline; `titleTestid` and `labelTestid` keep today's hooks.
     Under a micro label the title is the card's value (Plan: "Premium") and keeps its copy's
     case; alone it is the card's title ("Devices") and reads uppercase, as every section title
     does (polish spec 4.5 rule 2). -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { GlyphName } from '../../lib/ui/glyphs';
  import IconTile, { type TileTone } from './IconTile.svelte';

  let {
    glyph,
    tone = 'gold',
    label,
    title,
    titleSnippet,
    aside,
    testid,
  }: {
    glyph: GlyphName;
    tone?: TileTone;
    /** The micro label over the title; omitted, the title stands alone. */
    label?: string;
    title: string;
    /** Beside the title (a Founder tag). */
    titleSnippet?: Snippet;
    /** The head's right slot. */
    aside?: Snippet;
    testid?: string;
  } = $props();
</script>

<div class="flex flex-wrap items-center gap-x-3 gap-y-2 px-5 py-4" data-testid={testid}>
  <IconTile {glyph} {tone} />
  <div class="flex min-w-0 flex-1 flex-col gap-1">
    {#if label !== undefined}
      <span class="text-muted text-[11px] leading-none font-bold tracking-[.14em] uppercase">{label}</span>
    {/if}
    <div class="flex flex-wrap items-center gap-x-2 gap-y-1">
      <h2 class={label === undefined ? 'heading-plan uppercase' : 'heading-plan'}>{title}</h2>
      {@render titleSnippet?.()}
    </div>
  </div>
  {#if aside !== undefined}<div class="ml-auto min-w-0 text-right text-[13px]">{@render aside()}</div>{/if}
</div>
