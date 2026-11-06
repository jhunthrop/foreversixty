<!-- web/src/components/GuildRaid.svelte -->
<!-- The raid page /guild/{region}/{ruleset}/{slug}/raids/{id} (raids spec 3.2, 4 "Page states", R1):
     a raid still ahead, under way or cancelled is the raid planner (guild/planner/RaidPlanner); a
     done raid keeps the guild band, the report, the notes and the closed roster. A viewer who is
     not a verified member gets the band and the not-visible line. GuildShell loads this file
     as its own chunk; lib/guild/raid-page-controller.svelte.ts owns the data and the 10 s poll. -->
<script lang="ts">
  import { minuteClock } from '../lib/guild/minute-clock.svelte';
  import { viewerLocale } from '../lib/guild/raid-time';
  import { untrack } from 'svelte';
  import { guildClaimHref, guildHref, guildSettingsHref } from '../lib/characters';
  import { gateFactsFrom } from '../lib/guild/plan-gate';
  import { attendanceCopy } from '../lib/guild/attendance-copy';
  import { guildClaimCopy } from '../lib/guild/copy';
  import { browserStorage } from '../lib/guild/reserve-keys';
  import { SEP, eyebrowOf, kindOf, notesEmptyOf, raidCopy, zoneLineOf } from '../lib/guild/raid-copy';
  import {
    raidCount,
    raidPageFlags,
    raidWhen,
    reservedDoneRows,
    stripReserveRows,
  } from '../lib/guild/raid-page';
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
  import { raidQuiet } from '../lib/guild/planned-row';
  import { classRow } from '../lib/guild/raid-comp';
  import type { RaidStatusAction } from '../lib/guild/raids-api';
  import FactionCrest from './character/FactionCrest.svelte';
  import RaidDoneBand from './guild/RaidDoneBand.svelte';
  import GuildBandFrame from './guild/GuildBandFrame.svelte';
  import GuildTabs from './guild/GuildTabs.svelte';
  import RaidChip from './guild/RaidChip.svelte';
  import RaidClassRow from './guild/RaidClassRow.svelte';
  import Figures from './guild/Figures.svelte';
  import RaidDisclosure from './guild/RaidDisclosure.svelte';
  import RaidDoneSections from './guild/RaidDoneSections.svelte';
  import RaidRoster from './guild/RaidRoster.svelte';
  import RaidPlanner from './guild/planner/RaidPlanner.svelte';
  import ZoneKeyNote from './guild/ZoneKeyNote.svelte';
  import RaidYourNight from './guild/RaidYourNight.svelte';
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
  const when = $derived(raid === null ? null : raidWhen(raid, now, viewerTimezone, locale));
  const count = $derived(raid === null ? null : raidCount(raid));
  const classItems = $derived(raid === null ? [] : classRow(raid));
  const guildTimezone = $derived(raid?.guild_timezone ?? viewerTimezone);
  /** No gold on a cancelled raid or one already under way (spec 3.2 lines 4–5). */
  const quiet = $derived(raid !== null && raidQuiet(raid, now));
  const goldText = $derived(quiet ? 'text-muted' : 'text-gold');
  const goldFill = $derived(quiet ? 'bg-muted' : 'bg-gold');
  /** `Your night` takes the strip's slot for a member with an own character in the log or signed
   *  up; officers read their rows in the full sections. */
  const showYourNight = $derived(!officer && (raid?.attendance ?? []).some((row) => row.own));
  const countText = $derived(raid?.status === 'cancelled' ? 'text-muted' : 'text-strong');
  /** A done or ended raid sets its when line in two lines on phone (attendance spec 3.1). */
  const pastLines = $derived(flags !== null && (flags.done || flags.ended));

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
  <span class="label text-gold flex items-center gap-2.5">
    <i class="bg-gold inline-block h-px w-7" aria-hidden="true"></i>{raid === null
      ? raidCopy.page.eyebrow
      : eyebrowOf(kindOf(raid))}
  </span>
{/snippet}

{#snippet guildHeading()}
  <span class="label text-gold flex items-center gap-2.5">
    <i class="bg-gold inline-block h-px w-7" aria-hidden="true"></i>{raidCopy.page.guildEyebrow}
  </span>
  <div class="flex items-center gap-4 lg:gap-[18px]">
    <FactionCrest faction={c.guild?.faction ?? null} src={c.guild?.mark ?? null} />
    <h1 class="heading-page">{guildName}</h1>
  </div>
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
        class="bg-raised border-line rounded-panel border {RAID_STRIP_RESERVE}"
        style:--strip-rows={reserveRows}
        aria-hidden="true"
        data-testid="raid-strip-reserve"
      ></div>
    {:else}
      {#if !officer}
        <div
          class="bg-raised border-line rounded-panel border {YOUR_NIGHT_RESERVE}"
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
{:else if raid !== null && flags !== null && when !== null && count !== null}
  <!-- A done raid: the report path (slice 4 rebuilds it). -->
  <div class="reveal flex flex-col gap-[22px] md:gap-8 {GUILD_BANDLESS_TOP}" data-testid="raid-page">
    <p class="sr-only" role="status" data-testid="raid-live-region">
      {c.logMatched ? attendanceCopy.liveRegion : ''}
    </p>
    <GuildBandFrame>
      {@render backLinkRow()}
      <div class="grid grid-cols-1 gap-3 md:grid-cols-[minmax(0,1fr)_auto]">
        <div class="flex flex-wrap items-center gap-3 md:col-start-1 md:row-start-1">
          {@render raidEyebrow()}
          <RaidChip kind="status" status={flags.status} testid="raid-status-chip" />
        </div>
        <h1 class="heading-page md:col-span-2" data-testid="raid-title">
          <!-- The zone line in its own case, as the planned row writes it: `· Guild event · 20 players`,
                 never the h1's uppercase (guild round 2). -->
          {raid.title}{#if raid.title !== raid.zone_name}<span
              class="text-muted font-sans text-[15px] font-normal tracking-normal whitespace-nowrap normal-case"
              data-testid="raid-title-zone">{SEP}{zoneLineOf(raid)}</span
            >{/if}
        </h1>
        <!-- Phone: three lines (date; times; guild time and countdown); one line from md. The
               separators are entities: Svelte trims a span's edge spaces, which ran the words
               together from md up. -->
        <p class="text-[15px] md:col-span-2" data-testid="raid-when">
          <span class="block md:inline">{when.date}</span><span class="hidden md:inline">{SEP}</span><span
            class="{pastLines ? 'inline' : 'block md:inline'} whitespace-nowrap"
            ><span class="numeric">{when.start}</span>&#32;{raidCopy.page.timeTo}&#32;<span class="numeric"
              >{when.end}</span
            ></span
          ><span class={pastLines ? 'inline' : 'block md:inline'}
            ><span class={pastLines ? 'inline' : 'hidden md:inline'}>{SEP}</span
            >{#if when.yourTime !== null}<span class="text-muted whitespace-nowrap">{when.yourTime}</span
              >{SEP}{/if}<Figures text={when.countdown} class={goldText} /></span
          >
        </p>
        {#if raid.key_item != null}
          <div class="md:col-span-2"><ZoneKeyNote item={raid.key_item} testid="raid-key" /></div>
        {/if}
        <div class="flex flex-col gap-2 md:col-span-2" data-testid="raid-count">
          {#if raid.night === null}
            <p class="flex flex-wrap items-center gap-x-3 gap-y-1">
              {#if count.fill !== null}
                <span
                  class="bg-line-soft relative inline-block h-2 w-[120px] overflow-hidden rounded-full md:w-[160px]"
                  aria-hidden="true"
                  ><span class="{goldFill} absolute inset-y-0 left-0" style:width={`${count.fill}%`}
                  ></span></span
                >
              {/if}
              <span
                ><span class="{countText} numeric text-[15px]">{count.line}</span
                >{#if count.breakdown !== ''}<span class="text-muted text-[13px]">{SEP}{count.breakdown}</span
                  >{/if}</span
              >
            </p>
            <RaidClassRow items={classItems} {quiet} />
          {/if}
          {#if c.pollPaused}
            <p class="text-muted text-[13px]" data-testid="raid-poll-paused">{raidCopy.poll}</p>
          {/if}
          <RaidDoneBand
            {raid}
            {now}
            {guildPath}
            {viewerTimezone}
            {guildTimezone}
            {locale}
            {officer}
            readOnly={!c.officerTools}
            gateId={GATE_ID}
            loadCandidates={c.loadCandidates}
            onMatch={c.matchLog}
            onClear={c.clearMatch}
          />
        </div>
      </div>
      <GuildTabs active="raids" role={officer ? 'officer' : 'member'} hrefBase={guildPath} />
    </GuildBandFrame>

    {#if showYourNight}
      <RaidYourNight {raid} currentKey={c.currentKey} {locale} />
    {/if}
    <RaidDoneSections
      {raid}
      actions={c.night}
      {guildName}
      {guildPath}
      {settingsHref}
      {officer}
      officerTools={c.officerTools}
      gateId={GATE_ID}
      {gateFacts}
      {viewerTimezone}
      {locale}
    />

    <RaidDisclosure
      heading={raidCopy.subheads.notes}
      headingId="raid-notes-heading"
      testid="raid-notes-disclosure"
    >
      <section
        class="flex max-w-[720px] flex-col gap-3"
        aria-labelledby="raid-notes-heading"
        data-testid="raid-notes"
      >
        {#if raid.notes_html.trim() === ''}
          <p class="text-muted text-[14px]" data-testid="raid-notes-empty">
            {notesEmptyOf(kindOf(raid), officer)}
          </p>
        {:else}
          <div
            class="prose-forever [&_a]:text-gold [&_li]:marker:text-muted text-[14px] leading-[1.6] [&_ol]:list-decimal [&_ul]:list-disc"
            data-testid="raid-notes-prose"
          >
            <!-- The API renders and sanitises the notes (paragraphs, bold, italic, lists, links). -->
            <!-- eslint-disable-next-line svelte/no-at-html-tags -->
            {@html raid.notes_html}
          </div>
        {/if}
      </section>
    </RaidDisclosure>

    <RaidRoster
      {raid}
      {officer}
      officerTools={c.officerTools}
      gateId={GATE_ID}
      closed
      collapsed
      busyKey={c.rosterBusy}
      line={c.line}
      onSetState={(row, state) => c.setRosterState(row, state)}
    />
  </div>
{/if}
