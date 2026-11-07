<!-- web/src/components/guild/overview/NextRaidPlate.svelte -->
<!-- The page's one Plate (web layout spec 2026-11-05 §7.1), in its Section: the raid's Row (its
     icon, the guild's time and the viewer's own, the countdown), then the viewer's own Row (`you`:
     RaidBandLine's row, which writes the same sign-up the Discord buttons write). Officers also get
     the In, Late, Bench and Not signed figures and the planner link. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { GuildRaidView, NextRaidView } from '../../../lib/guild/api';
  import activeBuild from '../../../data/active-build.json';
  import { raidZoneIcon } from '../../../lib/first-raids';
  import { raidTimeLine } from '../../../lib/guild/overview-role-view';
  import { countdown } from '../../../lib/guild/raid-copy';
  import Figure from '../../ui/Figure.svelte';
  import ItemIcon from '../../ui/ItemIcon.svelte';
  import Plate from '../../ui/Plate.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';

  let {
    next,
    raid,
    officer,
    planHref,
    viewerTz,
    now,
    you,
  }: {
    next: NextRaidView;
    /** The full raid read, for the officer's counts; null until it answers. */
    raid: GuildRaidView | null;
    officer: boolean;
    planHref: string;
    viewerTz: string;
    now: Date;
    you?: Snippet;
  } = $props();

  const build = activeBuild.build;
  const guildTz = $derived(next.guild_timezone ?? viewerTz);
  const startsAt = $derived(new Date(next.starts_at));
  const endsAt = $derived(new Date(startsAt.getTime() + (raid?.duration_minutes ?? 180) * 60_000));
</script>

<Section
  title="Next raid"
  action={officer ? { label: 'Open the planner', href: planHref } : undefined}
  testid="guild-overview-next-raid"
>
  <Plate>
    <Row title={next.zone_name} open={false}>
      {#snippet lead()}<ItemIcon icon={raidZoneIcon(next.zone_name)} quality={5} {build} size={40} />{/snippet}
      {#snippet subContent()}<span data-testid="guild-next-raid-time"
          >{raidTimeLine(next.starts_at, guildTz, viewerTz)}</span
        >{/snippet}
      {#snippet trail()}
        <Figure
          value={countdown(startsAt, endsAt, now)}
          label="Starts in"
          size="sm"
          tone="gold"
          align="end"
          testid="guild-next-raid-countdown"
        />
      {/snippet}
    </Row>
    {#if you}<div data-testid="guild-next-raid-you">{@render you()}</div>{/if}
    {#if officer && raid !== null}
      <div class="flex flex-wrap gap-x-8 gap-y-3 px-4 py-3" data-testid="guild-next-raid-figures">
        <Figure value={String(raid.counts.in)} label="In" size="sm" tone="kill" />
        <Figure value={String(raid.counts.late)} label="Late" size="sm" tone="gold" />
        <Figure value={String(raid.counts.bench)} label="Bench" size="sm" />
        <Figure
          value={String(raid.not_signed_raiders?.length ?? raid.not_signed.length)}
          label="Not signed"
          size="sm"
          tone="wipe"
        />
      </div>
    {/if}
  </Plate>
</Section>
