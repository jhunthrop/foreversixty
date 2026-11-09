<!-- web/src/components/ui/RankList.svelte -->
<!-- A ranked list (tiers spec 2026-11-06 §4.4, owner decision T6): the column heads, then an ordered
     list of RankRows, or of RankGroups when `grouped`. One grid template, set from the list's shape,
     serves the heads and every row. It sits in the page's one Plate and draws nothing of its own.
     Written for the tier list; the log rankings and the guild Top guilds are next. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { rankListStyle, type RankHead } from '../../lib/ui/rank-list';
  import ColumnHeads from './ColumnHeads.svelte';
  import './rank-list.css';

  let {
    heads,
    figures = 1,
    figureWidth,
    links = 0,
    leads = true,
    grouped = false,
    compact = false,
    phoneFigures,
    label,
    testid,
    children,
  }: {
    heads?: readonly RankHead[];
    figures?: number;
    /** A figure column's width in px (112 by default). */
    figureWidth?: number;
    links?: number;
    /** The crest column; off on a list whose rows carry none (a single-class board, K3). */
    leads?: boolean;
    /** The children are RankGroups (a letter column from 1100). */
    grouped?: boolean;
    /** Round 5's compact list (the tier list beside its rail): elastic name and bar tracks that
     *  fit a split page's main column, 48 px rows, the 64 px letter column with the letter in the
     *  title face, and the accent tint on the selected row. */
    compact?: boolean;
    /** `compact` only: below 768, keep this many of the extra figures and set them beside the sub
     *  line, so a row stays two lines (the log rankings' percentile under the DPS). Left out, a
     *  phone row shows every extra figure on its own line, as before. */
    phoneFigures?: 1 | 2;
    /** The list's accessible name. */
    label?: string;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

<div
  class="ui-rank-list"
  data-grouped={grouped ? 'true' : undefined}
  data-compact={compact ? 'true' : undefined}
  data-phone-figures={compact && phoneFigures !== undefined ? phoneFigures : undefined}
  data-leads={leads ? undefined : 'none'}
  style={rankListStyle({ figures, links, figureWidth })}
  data-testid={testid}
>
  {#if heads && heads.length > 0}<ColumnHeads cells={heads} />{/if}
  <ol class="ui-rank-body" aria-label={label}>{@render children()}</ol>
</div>
