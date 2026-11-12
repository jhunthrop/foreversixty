<!-- web/src/components/ui/ShowAll.svelte -->
<!-- The one show-all toggle (component audit 2026-11-10 cluster 16): a gold link-weight button under
     a shortened list, "Show all 24" while it is shortened and "Show fewer" once open, announcing its
     state (aria-expanded) and naming the list it opens (aria-controls). Nothing shows when the list
     already fits. -->
<script lang="ts">
  import { LINK } from '../../lib/ui/primitives';
  import { showAllLabel, uiCopy } from '../../lib/ui/copy';

  let {
    shown,
    total,
    expanded,
    onToggle,
    label,
    fewerLabel = uiCopy.showFewer,
    collapsible = true,
    controls,
    testid,
  }: {
    /** Rows on screen while shortened. */
    shown: number;
    total: number;
    expanded: boolean;
    onToggle: () => void;
    /** The shortened words; "Show all {total}" when left out. */
    label?: string;
    fewerLabel?: string;
    /** False: once open the list stays open (no "Show fewer"). */
    collapsible?: boolean;
    /** The list's element id. */
    controls?: string;
    testid?: string;
  } = $props();

  const visible = $derived(total > shown && (!expanded || collapsible));
</script>

{#if visible}
  <button
    type="button"
    class="{LINK} inline-flex min-h-11 items-center self-start text-[14px]"
    aria-expanded={expanded}
    aria-controls={controls}
    data-testid={testid}
    onclick={onToggle}>{expanded ? fewerLabel : (label ?? showAllLabel(total))}</button
  >
{/if}
