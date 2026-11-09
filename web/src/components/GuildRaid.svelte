<!-- web/src/components/GuildRaid.svelte -->
<!-- The raid page /guild/{region}/{ruleset}/{slug}/raids/{id} (raids spec 3.2, 4 "Page states", R1):
     a raid still ahead, under way or cancelled is the raid planner (guild/planner/RaidPlanner); a
     done raid is the raid report (guild/report/RaidNightReport). A viewer who is
     not a verified member gets the band and the not-visible line. GuildShell loads this file
     as its own chunk; lib/guild/raid-page-controller.svelte.ts owns the data and the 10 s poll. -->
<script lang="ts">
  import { minuteClock } from '../lib/guild/minute-clock.svelte';
  import { viewerLocale } from '../lib/guild/raid-time';
  import { untrack } from 'svelte';
  import { guildClaimHref, guildHref, guildSettingsHref } from '../lib/characters';
  import { gateFactsFrom } from '../lib/guild/plan-gate';
  import { guildClaimCopy } from '../lib/guild/copy';
  import { browserStorage } from '../lib/guild/reserve-keys';
  import { eyebrowOf, kindOf, raidCopy } from '../lib/guild/raid-copy';
  import { raidPageFlags, reservedDoneRows, stripReserveRows } from '../lib/guild/raid-page';

  import { RaidPageController } from '../lib/guild/raid-page-controller.svelte';
  import type { RaidRoute } from '../lib/guild/raid-page-types';
  import { liveRaidPageDeps } from '../lib/guild/raid-page-deps';
  import {
    ATTENDANCE_RESERVE,
    RAID_STRIP_RESERVE,
    REPORT_RESERVE,
    YOUR_NIGHT_RESERVE,
    bandLinesReserve,
    GUILD_BANDLESS_TOP,
  } from '../lib/guild/raid-layout';
  import type { RaidStatusAction } from '../lib/guild/raids-api';
  import FactionCrest from './character/FactionCrest.svelte';
  import PageHeader from './ui/PageHeader.svelte';
  import GuildBandFrame from './guild/GuildBandFrame.svelte';
  import GuildTabs from './guild/GuildTabs.svelte';
  import RaidPlanner from './guild/planner/RaidPlanner.svelte';
  import RaidNightReport from './guild/report/RaidNightReport.svelte';
  import SignInPrompt from './SignInPrompt.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  const GATE_ID = 'officer-tools-gate-raids';

  let {
    path,
    controller,
    now: fixedNow,
    viewerTimezone = new Intl.DateTimeFormat().resolvedOptions().timeZone,
    locale = viewerLocale(),
  }: {
    path: RaidRoute;
    /** A test passes a seeded controller; the page builds its own otherwise. */
    controller?: RaidPageController;
    /** The clock the countdown reads; a test pins it. */
    now?: Date;
    viewerTimezone?: string;
    locale?: string;
  } = $props();

  // The route and the controller are fixed for the island's life: GuildShell mounts one per page.
  // svelte-ignore state_referenced_locally
  const c = controller ?? new RaidPageController(liveRaidPageDeps(path), path);
  // svelte-ignore state_referenced_locally
  const ownController = controller === undefined;

  $effect(() => (ownController ? untrack(() => c.start()) : undefined));

  const clock = minuteClock(() => fixedNow);
  const now = $derived(clock.now);

  const guildName = $derived(c.guild?.name ?? '');
  /** The route's slug stands in for the name until the guild answers, so the loading band's
   *  back link and tabs already point at the guild page. */
  const guildPath = $derived(guildHref(path.guild.region, path.guild.ruleset, guildName || path.guild.slug));
  const settingsHref = $derived(guildSettingsHref(path.guild.region, path.guild.ruleset, guildName));
  /** Raids are a feature of both plans: the plan flags do not change the gate's link. */
  const gateFacts = $derived(
    gateFactsFrom({
      claim: c.claim,
      membership: c.membership,
      officerTools: false,
      raidGuildTools: false,
      settingsHref,
      claimHref: guildClaimHref(path.guild.region, path.guild.ruleset, guildName),
    }),
  );
  const officer = $derived(c.role === 'officer');
  const reserveRows = stripReserveRows(browserStorage());
  /** The raid was last drawn `done` (its Your night rows, else null): the loading page holds the
   *  night, not the sign-up strip. */
  const doneRows = reservedDoneRows(path.raidId, browserStorage());
  const doneReserve = doneRows !== null;
  const visible = $derived(c.access === 'member');
  const raid = $derived(c.raid);
  const flags = $derived(raid === null ? null : raidPageFlags(raid, now));

  /** A done cancel leaves only disabled buttons, so focus moves to the back link. */
  async function runStatus(action: RaidStatusAction, post?: boolean): Promise<boolean> {
    const ok = await c.runStatus(action, post);
    if (ok && action === 'cancel')
      document.querySelector<HTMLAnchorElement>('[data-testid="raid-back-link"]')?.focus();
    return ok;
  }

  function closeEdit(): void {
    c.closeEdit();
    queueMicrotask(() => document.getElementById('raid-officer-edit')?.focus());
  }
</script>

{#snippet backLinkRow()}
  <a
    href={`${guildPath}#raids`}
    class="text-muted inline-flex min-h-11 items-center self-start text-[13px] underline"
    data-testid="raid-back-link">{raidCopy.backLink(guildName)}</a
  >
{/snippet}

{#snippet raidEyebrow()}
  <span class="text-label text-muted font-bold"
    >{raid === null ? raidCopy.page.eyebrow : eyebrowOf(kindOf(raid))}</span
  >
{/snippet}

{#snippet guildCrest()}
  <FactionCrest faction={c.guild?.faction ?? null} src={c.guild?.mark ?? null} />
{/snippet}

{#snippet guildHeading()}
  <PageHeader
    title={guildName}
    tone={c.guild?.faction ? { faction: c.guild.faction } : 'strong'}
    crest={guildCrest}
  />
{/snippet}

{#if c.load === 'failed'}
  <div class="flex flex-col gap-[22px] md:gap-8 {GUILD_BANDLESS_TOP}">
    <GuildBandFrame>
      {#if c.guild !== null}{@render guildHeading()}{:else}<h1 class="sr-only">
          {raidCopy.page.eyebrow}
        </h1>{/if}
    </GuildBandFrame>
    <LoadError message={raidCopy.loadErrors.raid} onRetry={() => c.retry()} testid="raid-load-error" />
  </div>
{:else if c.load === 'loading' || c.access === 'loading'}
  <div class="flex flex-col gap-[22px] md:gap-8 {GUILD_BANDLESS_TOP}" data-testid="raid-page-loading">
    <span class="sr-only">{raidCopy.page.loading}</span>
    <!-- The back link, eyebrow and tabs need only the route, so they are drawn for real; the
         skeleton holds lines 3–5 (and an officer's wrapped phone row) at their ready height. -->
    <GuildBandFrame>
      {@render backLinkRow()}
      {@render raidEyebrow()}
      <div class={bandLinesReserve(doneReserve, officer)}>
        <Skeleton lines={3} rowHeight="h-5" label={raidCopy.page.loading} testid="raid-band-skeleton" />
      </div>
      <GuildTabs active="raids" role={officer ? 'officer' : 'member'} hrefBase={guildPath} />
    </GuildBandFrame>
    {#if !doneReserve}
      <div
        class={RAID_STRIP_RESERVE}
        style:--strip-rows={reserveRows}
        aria-hidden="true"
        data-testid="raid-strip-reserve"
      ></div>
    {:else}
      {#if !officer}
        <div
          class={YOUR_NIGHT_RESERVE}
          style:--strip-rows={doneRows}
          aria-hidden="true"
          data-testid="raid-your-night-reserve"
        ></div>
      {/if}
      <div class={REPORT_RESERVE} aria-hidden="true" data-testid="raid-report-reserve"></div>
      {#if officer}
        <div class={ATTENDANCE_RESERVE} aria-hidden="true" data-testid="raid-attendance-reserve"></div>
      {/if}
    {/if}
    <div class="min-h-[560px] md:min-h-[420px]" aria-hidden="true" data-testid="raid-roster-reserve"></div>
  </div>
{:else if !visible}
  <div class={GUILD_BANDLESS_TOP}>
    <GuildBandFrame testid="raid-not-visible">
      {@render guildHeading()}
      {#if c.access === 'unverified'}
        <p class="text-text text-[14px]" data-testid="raid-not-visible-line">
          {raidCopy.notVisible.unverified}
        </p>
      {:else if c.access === 'signed-out'}
        <SignInPrompt
          line={raidCopy.notVisible.members(guildName)}
          next={typeof location === 'undefined' ? '/account?signed_in=1' : location.pathname}
          testid="raid-sign-in-prompt"
        />
      {:else}
        <p class="text-text text-[14px]" data-testid="raid-not-visible-line">
          {raidCopy.notVisible.members(guildName)}
        </p>
        <p class="text-muted text-[13px]" data-testid="raid-not-a-member">{guildClaimCopy.notAMember}</p>
      {/if}
    </GuildBandFrame>
  </div>
{:else if raid !== null && flags !== null && !flags.done}
  <p class="sr-only" role="status" data-testid="raid-live-region"></p>
  <RaidPlanner
    {c}
    {raid}
    {flags}
    {now}
    {viewerTimezone}
    {locale}
    {guildName}
    {guildPath}
    {settingsHref}
    {gateFacts}
    gateId={GATE_ID}
    onStatus={runStatus}
    onCloseEdit={closeEdit}
  />
{:else if raid !== null && flags !== null}
  <RaidNightReport
    {c}
    {raid}
    {flags}
    {now}
    {viewerTimezone}
    {locale}
    {guildName}
    {guildPath}
    {settingsHref}
    {gateFacts}
    gateId={GATE_ID}
  />
{/if}
