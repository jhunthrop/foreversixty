<!-- web/src/components/guild/GuildPlannedRaids.svelte -->
<!-- The Raids tab's `Planned` section (raids spec 3.1): the sub-head with its count, the next four
     rows by start and `Show all {n} planned raids`, which expands in place; the empty line says
     who plans raids (member) or what planning does (officer). -->
<script lang="ts">
  import Card from '../ui/Card.svelte';
  import type { GuildRaidView } from '../../lib/guild/api';
  import { guildRaidHref } from '../../lib/characters';
  import { sortByStart } from '../../lib/guild/planned-row';
  import { raidCopy, showAll } from '../../lib/guild/raid-copy';
  import PlannedRaidRow from './PlannedRaidRow.svelte';

  /** Rows shown before `Show all`. */
  const FIRST_ROWS = 4;

  let {
    planned,
    officer,
    guildPath,
    now,
    viewerTimezone,
  }: {
    planned: readonly GuildRaidView[];
    officer: boolean;
    guildPath: string;
    now: Date;
    viewerTimezone: string;
  } = $props();

  let expanded = $state(false);
  const sorted = $derived(sortByStart(planned));
  const shown = $derived(expanded ? sorted : sorted.slice(0, FIRST_ROWS));
</script>

<Card
  glyph="bell"
  title={raidCopy.subheads.planned}
  meta={sorted.length > 0 ? String(sorted.length) : undefined}
  body="none"
  testid="guild-planned"
>
  {#if sorted.length === 0}
    <p class="border-line-soft text-muted border-t px-5 py-4 text-[13px]" data-testid="guild-planned-empty">
      {officer ? raidCopy.emptyPlanned.officer : raidCopy.emptyPlanned.member}
    </p>
  {:else}
    <div class="border-line-soft flex flex-col border-t px-5" data-testid="guild-planned-rows">
      {#each shown as raid (raid.id)}
        <PlannedRaidRow {raid} href={guildRaidHref(guildPath, raid.id)} {now} {viewerTimezone} />
      {/each}
    </div>
    {#if !expanded && sorted.length > FIRST_ROWS}
      <div class="border-line-soft border-t px-5">
        <button
          type="button"
          class="text-text focus-visible:outline-gold h-11 text-[13px] underline focus-visible:outline-2"
          onclick={() => (expanded = true)}
          data-testid="guild-planned-show-all">{showAll(sorted.length)}</button
        >
      </div>
    {/if}
  {/if}
</Card>
