<!-- web/src/components/guild/RaidMoves.svelte -->
<!-- Biggest changes (attendance spec 3.4, officers): the five largest moves in a raider's best DPS or
     HPS on the same boss, against the previous done raid of the zone. A rise is `--kill`, a fall stays
     muted. No previous raid: `First {Zone} raid logged. Changes show from the next one.` -->
<script lang="ts">
  import type { RaidMoves } from '../../lib/guild/attendance-contract';
  import {
    changesHead,
    firstRaidLine,
    moveDelta,
    moveFigures,
    onBoss,
  } from '../../lib/guild/attendance-copy';
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
  const heading = $derived(
    previous === null ? '' : changesHead(dateLabel(new Date(previous.starts_at), guildTimezone, locale)),
  );
  const rows = $derived(biggestMoves(moves.rows));
</script>

{#if previous === null}
  <p class="text-muted text-[14px]" data-testid="raid-moves-first">{firstRaidLine(zoneName)}</p>
{:else}
  <RaidBlock {heading} testid="raid-moves">
    <ul>
      {#each rows as move (`${move.character_key}-${move.metric}-${move.encounter}`)}
        {@const delta = moveDelta(move.before, move.after)}
        <li
          class="{ROW_DIVIDER} flex min-h-11 flex-wrap items-center gap-x-3 gap-y-0.5 px-1 py-1.5"
          style={ownKeys.includes(move.character_key) ? `background:${OWN_ROW_WASH}` : undefined}
          data-testid="raid-move-row"
        >
          <ClassCrestRing characterClass={move.class} size={20} />
          <span class="text-[13px] font-semibold" style={`color:${classColorVar(move.class)}`}
            >{move.name}</span
          >
          <span class="text-muted text-[13px]">{onBoss(move.encounter)}</span>
          <span class="text-text font-mono text-[14px]"
            >{moveFigures(move.metric, move.before, move.after)}</span
          >
          <span class="{delta.rising ? 'text-kill' : 'text-muted'} font-mono text-[13px]">{delta.text}</span>
        </li>
      {/each}
    </ul>
  </RaidBlock>
{/if}
