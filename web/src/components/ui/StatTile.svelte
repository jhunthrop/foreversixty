<!-- web/src/components/ui/StatTile.svelte -->
<!-- A headline figure (polish spec 4.2): the figure in the heading face at 28 px over a micro
     label, on a quiet tile. Replaces a mono headline number. The label names the figure for a
     screen reader; the figure carries tabular numerals.
     Two optional extras for a tile that holds more than a number: `note`, a line under the
     label, and `figureTestid`, so a test reads the figure on its own rather than the whole tile.
     `figure` replaces `value` when the figure is more than text (a button that asks for it).
     `min-w-0` lets the tile shrink inside a grid column, which keeps a phone rail in the viewport. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    value = '',
    label,
    testid,
    figureTestid,
    figureClass = '',
    tone = 'strong',
    figure,
    note,
  }: {
    value?: string;
    label: string;
    testid?: string;
    figureTestid?: string;
    /** Extra classes for the figure row (a tone colour, or a reserved height). */
    figureClass?: string;
    /** The figure's colour: the text tone, or gold for the one figure a page wants to catch the eye. */
    tone?: 'strong' | 'gold';
    figure?: Snippet;
    note?: Snippet;
  } = $props();

  const TONE_CLASS = { strong: 'text-strong', gold: 'text-gold' } as const;
</script>

<div
  class="rounded-control border-line-soft flex min-w-0 flex-col gap-1 border bg-white/[.03] px-[14px] py-3"
  data-testid={testid}
>
  <div class="figure-stat flex items-center {TONE_CLASS[tone]} {figureClass}">
    {#if figure !== undefined}
      {@render figure()}
    {:else}
      <span data-testid={figureTestid}>{value}</span>
    {/if}
  </div>
  <span class="text-muted text-[11px] font-bold tracking-[.08em] uppercase">{label}</span>
  {#if note !== undefined}{@render note()}{/if}
</div>
