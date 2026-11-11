<!-- web/src/components/guild/overview/NextRaidRow.svelte -->
<!-- The next raid as one Row: the zone's icon, the guild's time and the viewer's own, the countdown.
     Next raid, Wednesday and Your week lead with it. On a phone the countdown ("in 4 days, 4 hours")
     sits on its own line under the time instead of the trail, so it never squeezes the title. -->
<script lang="ts">
  import ZoneIcon from '../../ui/ZoneIcon.svelte';
  import activeBuild from '../../../data/active-build.json';
  import type { NextRaidView } from '../../../lib/guild/api';
  import { raidTimeLine } from '../../../lib/guild/overview-role-view';
  import { countdown } from '../../../lib/guild/raid-copy';
  import Figure from '../../ui/Figure.svelte';
  import Row from '../../ui/Row.svelte';

  let {
    next,
    durationMinutes = 180,
    viewerTz,
    now,
  }: {
    next: Pick<NextRaidView, 'zone_name' | 'starts_at' | 'guild_timezone'>;
    durationMinutes?: number;
    viewerTz: string;
    now: Date;
  } = $props();

  const build = activeBuild.build;
  const guildTz = $derived(next.guild_timezone ?? viewerTz);
  const startsAt = $derived(new Date(next.starts_at));
  const endsAt = $derived(new Date(startsAt.getTime() + durationMinutes * 60_000));
  const left = $derived(countdown(startsAt, endsAt, now));
</script>

<Row title={next.zone_name} open={false}>
  {#snippet lead()}<ZoneIcon zone={next.zone_name} {build} size={40} />{/snippet}
  {#snippet subContent()}<span data-testid="guild-next-raid-time"
      >{raidTimeLine(next.starts_at, guildTz, viewerTz)}</span
    ><span class="next-raid-phone">{left}</span>{/snippet}
  {#snippet trail()}
    <span class="next-raid-trail">
      <Figure value={left} size="sm" tone="gold" align="end" testid="guild-next-raid-countdown" />
    </span>
  {/snippet}
</Row>

<style>
  .next-raid-phone {
    display: block;
    color: var(--color-gold);
    font-weight: 600;
  }
  .next-raid-trail {
    display: none;
  }
  @media (min-width: 640px) {
    .next-raid-phone {
      display: none;
    }
    .next-raid-trail {
      display: contents;
    }
  }
</style>
