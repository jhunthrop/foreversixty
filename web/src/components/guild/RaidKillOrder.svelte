<!-- web/src/components/guild/RaidKillOrder.svelte -->
<!-- Kill order (attendance spec 3.4): one 44 px row per named boss in kill order, each linking to its
     fight on the report page; a boss never killed reads `Not killed · n wipes` with its name muted;
     pulls with no named encounter take one line. Only `first pull` is `--kill`. -->
<script lang="ts">
  import type { RaidReport } from '../../lib/guild/attendance-contract';
  import { attendanceCopy, killRowText, unnamedLine } from '../../lib/guild/attendance-copy';
  import { ROW_DIVIDER, ROW_LINK } from '../../lib/guild/raid-report-styles';
  import RaidBlock from './RaidBlock.svelte';

  let { report, reportId }: { report: RaidReport; reportId: string } = $props();

  const fightHref = (fightIndex: number): string => `/reports/${reportId}?fight=${fightIndex}`;
</script>

<RaidBlock heading={attendanceCopy.report.kills} count={report.kill_order.length} testid="raid-kill-order">
  <ul>
    {#each report.kill_order as entry, index (entry.encounter_id)}
      {@const text = killRowText(entry)}
      <li class={ROW_DIVIDER}>
        <a
          href={fightHref(entry.fight_index)}
          class="{ROW_LINK} flex min-h-11 items-baseline gap-3 px-1 py-2.5"
          data-testid="raid-kill-row"
        >
          <span class="text-muted numeric w-4 text-[13px]">{index + 1}</span>
          <span class="{text.time === null ? 'text-muted' : 'text-strong'} text-[14px] font-semibold"
            >{entry.name}</span
          >
          {#if text.time !== null}<span class="text-text numeric ml-auto text-[14px]">{text.time}</span>{/if}
          <span class="text-muted text-[13px] {text.time === null ? 'ml-auto' : ''}"
            >{#if text.firstPull}<span class="text-kill">{text.note}</span
              >{:else}{text.note}{/if}{text.deaths}</span
          >
        </a>
      </li>
    {/each}
  </ul>
  {#if report.unnamed !== null}
    <p class="text-muted text-[13px]" data-testid="raid-unnamed-line">{unnamedLine(report.unnamed)}</p>
  {/if}
</RaidBlock>
