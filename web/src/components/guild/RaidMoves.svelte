<!-- web/src/components/guild/RaidMoves.svelte -->
<!-- Biggest changes (attendance spec 3.4, officers): the five largest moves in a raider's best DPS or
     HPS on the same boss, against the previous done raid of the zone. A rise is `--kill`, a fall stays
     muted. No previous raid: `First {Zone} raid logged. Changes show from the next one.` -->
<script lang="ts">
  import type { RaidMoves } from '../../lib/guild/attendance-contract';
  import {
    attendanceCopy,
    changesSince,
    firstRaidLine,
    moveDelta,
    moveFigures,
    onBoss,
  } from '../../lib/guild/attendance-copy';
  import { SEP } from '../../lib/guild/raid-copy';
  import { OWN_ROW_WASH, ROW_DIVIDER } from '../../lib/guild/raid-report-styles';
  import { dateLabel } from '../../lib/guild/raid-time';
  import { biggestMoves } from '../../lib/guild/report-view';
  import { classColorVar } from '../../lib/report/format';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import RaidBlock from './RaidBlock.svelte';

  let {
    moves,
    zoneName,
    guildTimezone,
    locale,
    ownKeys = [],
  }: {
    moves: RaidMoves;
    zoneName: string;
    guildTimezone: string;
    locale: string;
    ownKeys?: readonly string[];
  } = $props();

  const previous = $derived(moves.previous);
  const sinceTail = $derived(
    previous === null ? '' : changesSince(dateLabel(new Date(previous.starts_at), guildTimezone, locale)),
  );
  const rows = $derived(biggestMoves(moves.rows));
</script>

{#if previous === null}
  <RaidBlock heading={attendanceCopy.report.changes} testid="raid-moves">
    <p class="text-muted text-[14px]" data-testid="raid-moves-first">{firstRaidLine(zoneName)}</p>
  </RaidBlock>
{:else}
  <RaidBlock heading={attendanceCopy.report.changes} tail={sinceTail} testid="raid-moves">
    <ul>
      {#each rows as move (`${move.character_key}-${move.metric}-${move.encounter}`)}
        {@const delta = moveDelta(move.before, move.after)}
        <!-- From 768: who, `on {Boss}`, the figures and the delta right-aligned. Below: who and the
             delta, then `on Onyxia · DPS 238 → 291`. -->
        <li
          class="{ROW_DIVIDER} grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-0.5 px-1 py-1.5 md:grid-cols-[minmax(0,1fr)_auto_auto_56px]"
          style={ownKeys.includes(move.character_key) ? `background:${OWN_ROW_WASH}` : undefined}
          data-testid="raid-move-row"
        >
          <span class="flex min-w-0 items-center gap-2">
            <ClassCrestRing characterClass={move.class} size={20} />
            <span class="truncate text-[13px] font-semibold" style={`color:${classColorVar(move.class)}`}
              >{move.name}</span
            >
          </span>
          <span
            class="{delta.rising ? 'text-kill' : 'text-muted'} numeric text-right text-[13px] md:order-last"
            >{delta.text}</span
          >
          <span class="text-muted col-span-2 pl-7 text-[13px] whitespace-nowrap md:col-span-1 md:pl-0"
            >{onBoss(move.encounter)}<span class="md:hidden">{SEP}</span><span
              class="text-text numeric text-[14px] md:hidden"
              >{moveFigures(move.metric, move.before, move.after)}</span
            ></span
          >
          <span class="text-text numeric hidden text-[14px] whitespace-nowrap md:inline"
            >{moveFigures(move.metric, move.before, move.after)}</span
          >
        </li>
      {/each}
    </ul>
  </RaidBlock>
{/if}
