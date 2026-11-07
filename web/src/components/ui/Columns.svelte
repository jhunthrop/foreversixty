<!-- web/src/components/ui/Columns.svelte -->
<!-- PlateColumns' columns without the Plate (premium spec 2026-11-08 section 3A): each child is one
     column, hairlines between, side by side from 1024 and stacked below. `open` (default) draws no
     side padding, for columns on the page ground; `open={false}` is the padded form a caller puts in
     its own Plate when each column opens with its own head (PlanHead). `secondFirst` stacks the
     second column first below 1024: the plan a page sells leads on a phone. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    count,
    open = true,
    secondFirst = false,
    label,
    testid,
    children,
  }: {
    /** How many columns (1 to 3). */
    count: number;
    open?: boolean;
    secondFirst?: boolean;
    /** An accessible name, when the columns are a group of their own. */
    label?: string;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

<div
  class="ui-cols"
  data-open={open ? 'true' : 'false'}
  data-swap={secondFirst ? 'true' : undefined}
  style:--cols={count}
  role={label ? 'group' : undefined}
  aria-label={label}
  data-testid={testid}
>
  {@render children()}
</div>

<style>
  .ui-cols {
    display: grid;
    grid-template-columns: minmax(0, 1fr);
  }
  .ui-cols > :global(*) {
    display: flex;
    flex-direction: column;
    gap: 12px;
    min-width: 0;
    padding: 16px 20px 20px;
  }
  .ui-cols[data-open='true'] > :global(*) {
    padding: 16px 0 0;
  }
  .ui-cols > :global(* + *) {
    border-top: 1px solid var(--color-line-soft);
  }
  @media (max-width: 1023px) {
    .ui-cols[data-swap='true'] {
      display: flex;
      flex-direction: column;
    }
    .ui-cols[data-swap='true'] > :global(:nth-child(2)) {
      order: -1;
      border-top: 0;
    }
    .ui-cols[data-swap='true'] > :global(:first-child) {
      border-top: 1px solid var(--color-line-soft);
    }
  }
  @media (min-width: 1024px) {
    .ui-cols {
      grid-template-columns: repeat(var(--cols), minmax(0, 1fr));
    }
    .ui-cols > :global(* + *) {
      border-top: 0;
      border-left: 1px solid var(--color-line-soft);
    }
    .ui-cols[data-open='true'] > :global(*) {
      padding: 4px 24px 0;
    }
    .ui-cols[data-open='true'] > :global(:first-child) {
      padding-left: 0;
    }
    .ui-cols[data-open='true'] > :global(:last-child) {
      padding-right: 0;
    }
  }
</style>
