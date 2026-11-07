<!-- web/src/components/ui/Tally.svelte -->
<!-- Counts against a target (web layout spec 2026-11-05 §7.3): one entry per group, the label, the
     count over the target toned at, near or short of it, and a Meter. Comp against the target, a
     class in the guild, later the guide's raid roles. Open, no box; entries flow in columns of at
     least 150 px. An item's `tone` overrides the target's (a count that is
     a share, not a goal: a level band of the roster reads muted). -->
<script lang="ts">
  import { CELL_TEXT, tallyTone, type CellTone } from '../../lib/ui/cells';
  import Meter from './Meter.svelte';

  let {
    items,
    testid,
  }: {
    items: readonly {
      key: string;
      label: string;
      have: number;
      want: number;
      note?: string;
      tone?: CellTone;
    }[];
    testid?: string;
  } = $props();
</script>

<ul class="grid grid-cols-[repeat(auto-fill,minmax(150px,1fr))] gap-x-4 gap-y-3" data-testid={testid}>
  {#each items as item (item.key)}
    {@const tone = item.tone ?? tallyTone(item.have, item.want)}
    <li
      class="flex min-w-0 flex-col gap-[6px]"
      data-testid={testid === undefined ? undefined : `${testid}-${item.key}`}
    >
      <span class="flex items-baseline gap-2">
        <span class="text-strong text-[15px] leading-5 font-bold">{item.label}</span>
        <span class="ml-auto text-[15px] leading-5 font-bold tabular-nums {CELL_TEXT[tone]}"
          >{item.have}/{item.want}</span
        >
      </span>
      <Meter
        value={item.want > 0 ? item.have / item.want : 1}
        tone={tone === 'ok' ? 'kill' : tone === 'none' ? 'muted' : 'gold'}
        label={`${item.label}: ${item.have} of ${item.want}`}
      />
      {#if item.note}<span class="text-muted text-[12px] leading-4">{item.note}</span>{/if}
    </li>
  {/each}
</ul>
