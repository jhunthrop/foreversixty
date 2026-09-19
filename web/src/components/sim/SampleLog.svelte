<!-- web/src/components/sim/SampleLog.svelte -->
<!-- Design 5.1. One iteration, in order, with the pre-pull ruled off from the fight and
     the resources after each cast. The note above it is not decoration: a table of one
     iteration's casts looks exactly like a rotation guide, and it is not one. -->
<script lang="ts">
  import type { ActionNames } from '../../lib/sim/action-names';
  import { simCopy } from '../../lib/sim/copy';
  import { sampleLog } from '../../lib/sim/sample-log';
  import type { SampleCast } from '../../lib/sim/types';

  let { sample, actionNames }: { sample: SampleCast[] | undefined; actionNames: ActionNames | null } =
    $props();

  const log = $derived(sampleLog(sample, actionNames));
</script>

{#if log.rows.length === 0}
  <p class="text-muted p-4 text-[14px]" data-testid="sim-sample-empty">{simCopy.sampleEmpty}</p>
{:else}
  <div class="flex flex-col gap-2 p-2" data-testid="sim-sample-log">
    <p class="text-muted px-2 text-[12px]">{simCopy.sampleNote}</p>
    <!-- The resource columns make this the one table on the lane that can genuinely be
         wider than a phone, so it gets its own scroller rather than the page getting one. -->
    <div class="overflow-x-auto">
      <table class="w-full min-w-[420px] text-[13px]">
        <thead>
          <tr class="text-muted label">
            <th class="px-2 py-1 text-left">{simCopy.sampleTimeHeading}</th>
            <th class="px-2 py-1 text-left">{simCopy.sampleCastHeading}</th>
            <th class="px-2 py-1 text-left">{simCopy.sampleTargetHeading}</th>
            {#each log.columns as column (column)}
              <th class="px-2 py-1 text-right">{simCopy.resourceLabel[column] ?? column}</th>
            {/each}
          </tr>
        </thead>
        <tbody>
          {#each log.rows as row, index (row.key)}
            <!-- One rule, under the last pre-pull cast: the pull is the only boundary in
                 this table and a heading row would break the column alignment to say it. -->
            <tr
              class={`border-line-soft border-b ${
                index + 1 === log.prePullCount ? 'border-b-line-warm-strong' : ''
              }`}
              data-testid={`sim-sample-row-${row.key}`}
            >
              <td class="tabular px-2 py-1 font-mono">
                {row.time}{#if row.prePull}<span class="text-muted ml-1 text-[11px]"
                    >{simCopy.samplePrePull}</span
                  >{/if}
              </td>
              <td class="px-2 py-1 font-semibold">{row.name}</td>
              <td class="text-muted px-2 py-1">{row.target}</td>
              {#each log.columns as column (column)}
                <td class="tabular px-2 py-1 text-right font-mono">{row.resources[column] ?? ''}</td>
              {/each}
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  </div>
{/if}
