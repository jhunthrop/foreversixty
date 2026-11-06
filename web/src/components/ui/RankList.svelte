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
    links = 0,
    wideFigures = false,
    grouped = false,
    label,
    testid,
    children,
  }: {
    heads?: readonly RankHead[];
    figures?: number;
    links?: number;
    wideFigures?: boolean;
    /** The children are RankGroups (a letter column from 1100). */
    grouped?: boolean;
    /** The list's accessible name. */
    label?: string;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

<div
  class="ui-rank-list"
  data-grouped={grouped ? 'true' : undefined}
  style={rankListStyle({ figures, links, wideFigures })}
  data-testid={testid}
>
  {#if heads && heads.length > 0}<ColumnHeads cells={heads} />{/if}
  <ol class="ui-rank-body" aria-label={label}>{@render children()}</ol>
</div>
