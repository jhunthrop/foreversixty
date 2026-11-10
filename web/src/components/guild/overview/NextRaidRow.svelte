<!-- web/src/components/guild/overview/NextRaidRow.svelte -->
<!-- The next raid as one Row: the zone's icon, the guild's time and the viewer's own, the countdown.
     Next raid, Wednesday and Your week lead with it. -->
<script lang="ts">
  import activeBuild from '../../../data/active-build.json';
  import type { NextRaidView } from '../../../lib/guild/api';
  import { raidZoneIcon } from '../../../lib/first-raids';
  import { raidTimeLine } from '../../../lib/guild/overview-role-view';
  import { countdown } from '../../../lib/guild/raid-copy';
  import Figure from '../../ui/Figure.svelte';
  import ItemIcon from '../../ui/ItemIcon.svelte';
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
</script>

<Row title={next.zone_name} open={false}>
  {#snippet lead()}<ItemIcon icon={raidZoneIcon(next.zone_name)} quality={5} {build} size={40} />{/snippet}
  {#snippet subContent()}<span data-testid="guild-next-raid-time">{raidTimeLine(next.starts_at, guildTz, viewerTz)}</span
    >{/snippet}
  {#snippet trail()}
    <Figure value={countdown(startsAt, endsAt, now)} size="sm" tone="gold" align="end" testid="guild-next-raid-countdown" />
  {/snippet}
</Row>
