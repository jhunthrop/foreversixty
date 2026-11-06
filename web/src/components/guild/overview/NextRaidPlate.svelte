<!-- web/src/components/guild/overview/NextRaidPlate.svelte -->
<!-- The page's one Plate (web layout spec 2026-11-05 §7.1): the next raid in the guild's time and the
     viewer's own, the countdown, then the viewer's own row (`you`: the band line, which writes the
     same sign-up the Discord buttons write). Officers also get the In, Late, Bench and Not signed
     figures and the planner link. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { GuildRaidView, NextRaidView } from '../../../lib/guild/api';
  import { raidTimeLine } from '../../../lib/guild/overview-role-view';
  import { countdown } from '../../../lib/guild/raid-copy';
  import { LINK } from '../../../lib/ui/primitives';
  import Figure from '../../ui/Figure.svelte';
  import Plate from '../../ui/Plate.svelte';

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

  const guildTz = $derived(next.guild_timezone ?? viewerTz);
  const startsAt = $derived(new Date(next.starts_at));
  const endsAt = $derived(new Date(startsAt.getTime() + (raid?.duration_minutes ?? 180) * 60_000));
</script>

<section
  class="flex flex-col gap-3"
  aria-labelledby="guild-next-raid-head"
  data-testid="guild-overview-next-raid"
>
  <div class="flex items-baseline justify-between gap-4">
    <h2 id="guild-next-raid-head" class="text-strong text-[18px] leading-6 font-bold">Next raid</h2>
    {#if officer}<a class={LINK} href={planHref}>Open the planner ›</a>{/if}
  </div>
  <Plate flush={false} label="Next raid">
    <div class="flex flex-col gap-4 p-4">
      <div class="flex flex-wrap items-end justify-between gap-3">
        <div class="min-w-0">
          <p class="text-strong text-[22px] leading-7 font-bold">{next.zone_name}</p>
          <p class="text-muted text-small" data-testid="guild-next-raid-time">
            {raidTimeLine(next.starts_at, guildTz, viewerTz)}
          </p>
        </div>
        <Figure
          value={countdown(startsAt, endsAt, now)}
          size="sm"
          tone="gold"
          align="end"
          testid="guild-next-raid-countdown"
        />
      </div>
      {#if you}<div data-testid="guild-next-raid-you">{@render you()}</div>{/if}
      {#if officer && raid !== null}
        <div class="flex flex-wrap gap-x-8 gap-y-3" data-testid="guild-next-raid-figures">
          <Figure value={String(raid.counts.in)} label="In" size="sm" tone="kill" />
          <Figure value={String(raid.counts.late)} label="Late" size="sm" />
          <Figure value={String(raid.counts.bench)} label="Bench" size="sm" />
          <Figure
            value={String(raid.not_signed_raiders?.length ?? raid.not_signed.length)}
            label="Not signed"
            size="sm"
            tone="wipe"
          />
        </div>
      {/if}
    </div>
  </Plate>
</section>
