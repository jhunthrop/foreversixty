<!-- web/src/components/home/HomeCountdowns.svelte -->
<!-- The countdowns under the head (home B spec 2026-11-07 §1): the first raids opening with the three
     raids, and (signed in) the guild's next event with the sign-up state, or the no-guild and
     no-raid lines. Signed out the raid panel stands alone. Both tick in the browser. -->
<script lang="ts">
  import type { NextRaidView } from '../../lib/guild/api';
  import { raidTimeLine } from '../../lib/guild/overview-role-view';
  import { homeBCopy, homeBHrefs } from '../../lib/home/home-b-view';
  import { LINK } from '../../lib/ui/primitives';
  import { REGION } from '../../lib/ui/regions';
  import Countdown from '../ui/Countdown.svelte';
  import RaidsCountdown from '../RaidsCountdown.svelte';

  let {
    raidsOpen,
    build,
    guild,
    raid = null,
    raidHref = null,
    guildHref = null,
  }: {
    /** The opening instant while ahead; null once raids are open. */
    raidsOpen: string | null;
    build: string;
    /** Signed out: undefined (no guild panel). Signed in: the guild's name, or null with none. */
    guild?: string | null;
    raid?: NextRaidView | null;
    raidHref?: string | null;
    guildHref?: string | null;
  } = $props();

  const viewerTz = Intl.DateTimeFormat().resolvedOptions().timeZone;
  const raidLine = $derived.by(() => {
    if (raid === null) return '';
    const state = raid.own === null ? homeBCopy.notSigned : raid.own.state;
    const when = raidTimeLine(raid.starts_at, raid.guild_timezone ?? viewerTz, viewerTz);
    return `${when} · ${homeBCopy.inCount(raid.counts.in)} · ${homeBCopy.you(state)}`;
  });
</script>

<div class="home-countdowns" class:solo={guild === undefined} data-region={REGION.countdowns}>
  <RaidsCountdown {raidsOpen} {build} testid="home-raids-countdown" />
  {#if guild !== undefined}
    {#if guild === null}
      <Countdown eyebrow={homeBCopy.findGuild} target={null} tone="guild" testid="home-guild-countdown">
        <p class="home-guild-what">
          {homeBCopy.findGuildSub}<br /><a class={LINK} href={homeBHrefs.guilds}>{homeBCopy.guildsLink} ›</a>
        </p>
      </Countdown>
    {:else if raid === null}
      <Countdown
        eyebrow={homeBCopy.guildNone(guild)}
        target={null}
        tone="guild"
        testid="home-guild-countdown"
      >
        {#if guildHref !== null}<p class="home-guild-what">
            <a class={LINK} href={guildHref}>{homeBCopy.raidPageLink} ›</a>
          </p>{/if}
      </Countdown>
    {:else}
      <Countdown
        eyebrow={homeBCopy.guildEyebrow(guild)}
        target={raid.starts_at}
        tone="guild"
        testid="home-guild-countdown"
      >
        <p class="home-guild-what">
          <b>{raid.title}</b>{raidLine}<br />
          {#if raidHref !== null}<a class={LINK} href={raidHref}>{homeBCopy.signUpLink} ›</a>{/if}
        </p>
      </Countdown>
    {/if}
  {/if}
</div>

<style>
  .home-countdowns {
    display: grid;
    gap: var(--gap);
    grid-template-columns: minmax(0, 1fr);
    padding: 28px var(--page-inset) 0;
  }
  @media (min-width: 1024px) {
    .home-countdowns:not(.solo) {
      grid-template-columns: minmax(0, 1.3fr) minmax(0, 1fr);
    }
  }
  /* One line per fact, cut with an ellipsis, so the panel keeps one height whatever the zone,
     the raid title or the sign-up line says. */
  .home-guild-what {
    margin: 0;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    text-align: right;
    font-size: 14px;
    line-height: 22px;
    color: var(--color-nav);
  }
  .home-guild-what b {
    display: block;
    overflow: hidden;
    text-overflow: ellipsis;
    font-size: 18px;
    line-height: 24px;
    color: var(--color-strong);
  }
</style>
