<!-- web/src/components/ui/Grid.svelte -->
<!-- A coverage matrix (web layout spec 2026-11-05 §7.3): a label column, then one centred column per
     header, each cell toned ok / warn / miss / none. Buff and debuff cover, attendance over the last
     raids, later readiness and consumables by pull. Open hairline rows, no box. A table, so a screen
     reader reads each cell against its column; `lead` draws a row's icon or crest. -->
<script lang="ts" generics="R extends GridRowModel">
  import type { Snippet } from 'svelte';
  import { CELL_TEXT, type GridRowModel } from '../../lib/ui/cells';

  let {
    caption,
    columns,
    rows,
    lead,
    testid,
  }: {
    /** The table's name for assistive tech ("Buff and debuff cover"); not drawn. */
    caption: string;
    columns: readonly string[];
    rows: readonly R[];
    lead?: Snippet<[R]>;
    testid?: string;
  } = $props();

  const CELL = 'border-line-soft border-b px-[6px] py-[10px]';
</script>

<table class="w-full border-collapse" data-testid={testid}>
  <caption class="sr-only">{caption}</caption>
  <thead>
    <tr>
      <th scope="col" class="border-line-soft border-b px-[6px] py-2"
        ><span class="sr-only">{caption}</span></th
      >
      {#each columns as column (column)}
        <th
          scope="col"
          class="text-label tracking-label text-muted border-line-soft border-b px-[6px] py-2 text-center font-bold uppercase"
          >{column}</th
        >
      {/each}
    </tr>
  </thead>
  <tbody>
    {#each rows as row (row.key)}
      <tr data-testid={testid === undefined ? undefined : `${testid}-row`}>
        <th scope="row" class="{CELL} min-w-0 text-left font-normal">
          <span class="flex min-w-0 items-center gap-2">
            {#if lead}{@render lead(row)}{/if}
            <span class="min-w-0">
              <span class="text-strong block text-[14px] leading-[18px] font-bold">{row.label}</span>
              {#if row.sub}<span class="text-muted block text-[12px] leading-4">{row.sub}</span>{/if}
            </span>
          </span>
        </th>
        {#each row.cells as cell, index (index)}
          <td
            class="{CELL} text-center text-[14px] leading-[18px] font-bold tabular-nums {CELL_TEXT[
              cell.tone
            ]}"
            >{#if cell.label}<span aria-hidden="true">{cell.text}</span><span class="sr-only"
                >{cell.label}</span
              >{:else}{cell.text}{/if}</td
          >
        {/each}
      </tr>
    {/each}
  </tbody>
</table>
