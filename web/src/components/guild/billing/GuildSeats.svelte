<!-- web/src/components/guild/billing/GuildSeats.svelte -->
<!-- Premium seats (spec 3.2.5): a level 3 Section with its count, the helper, one open Row per seat. Officers
     without `seats` see the rows with no right-hand column. Nothing draws before a plan has seats. -->
<script lang="ts">
  import type { SeatView } from '../../../lib/billing/contract';
  import { seatsCopy as copy } from '../../../lib/billing/seats-copy';
  import { givenCount } from '../../../lib/guild/seat-view';
  import Section from '../../ui/Section.svelte';
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
  <Section
    title={copy.heading}
    level={3}
    meta={copy.given(Math.min(givenCount(seats), total), total)}
    metaTestid="guild-seats-count"
    testid="guild-seats"
  >
    <p class="text-muted pt-2 text-[13px]">{copy.explainer}</p>
    <ul class="m-0 list-none p-0">
      {#each seats as seat (seat.slot)}
        <GuildSeatRow {guildId} {seat} {manage} {now} {onChanged} />
      {/each}
    </ul>
  </Section>
{/if}
