<!-- web/src/components/guild/billing/GuildSeats.svelte -->
<!-- Premium seats (spec 3.2.5): sub-head with its count, the helper, one row per seat. Officers
     without `seats` see the rows with no right-hand column. Nothing draws before a plan has seats. -->
<script lang="ts">
  import type { SeatView } from '../../../lib/billing/contract';
  import { seatsCopy as copy } from '../../../lib/billing/seats-copy';
  import { givenCount } from '../../../lib/guild/seat-view';
  import GuildSeatRow from './GuildSeatRow.svelte';

  let {
    guildId,
    seats,
    total = seats.length,
    manage,
    now,
    onChanged,
  }: {
    guildId: number;
    seats: readonly SeatView[];
    /** The plan's seat count (3 on Raid guild); a returned row fills a slot, it never adds one. */
    total?: number;
    manage: boolean;
    now: Date;
    onChanged: () => Promise<void>;
  } = $props();
</script>

{#if seats.length > 0}
  <div class="flex flex-col gap-2" data-testid="guild-seats">
    <div class="flex items-baseline gap-3">
      <h3 class="section-title">{copy.heading}</h3>
      <span class="text-muted numeric text-[12px]" data-testid="guild-seats-count">
        {copy.given(Math.min(givenCount(seats), total), total)}
      </span>
    </div>
    <p class="text-muted text-[13px]">{copy.explainer}</p>
    <ul class="border-line-soft border-t">
      {#each seats as seat (seat.slot)}
        <GuildSeatRow {guildId} {seat} {manage} {now} {onChanged} />
      {/each}
    </ul>
  </div>
{/if}
