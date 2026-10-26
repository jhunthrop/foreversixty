<!-- web/src/components/GuildRaid.svelte -->
<!-- The raid page /guild/{region}/{ruleset}/{slug}/raids/{id} (raids spec 3.2, 4 "Page states", R1):
     the guild band with the back link, eyebrow and status chip, the title, the when and count lines,
     the officer row and the tabs; then the viewer's sign-up strip, the notes and the roster. A viewer
     who is not a verified member gets the band and the not-visible line. GuildShell loads this file
     as its own chunk; lib/guild/raid-page-controller.svelte.ts owns the data and the 10 s poll. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import { guildHref, guildSettingsHref } from '../lib/characters';
  import { guildClaimCopy } from '../lib/guild/copy';
  import { focusOnMount } from '../lib/guild/focus-now';
  import { browserStorage } from '../lib/guild/planned-reserve';
  import { raidCopy } from '../lib/guild/raid-copy';
  import { raidCount, raidPageFlags, raidWhen, stripReserveRows } from '../lib/guild/raid-page';
  import { RaidPageController, type RaidRoute } from '../lib/guild/raid-page-controller.svelte';
  import { liveRaidPageDeps } from '../lib/guild/raid-page-deps';
  import { RAID_BAND_LINES_MIN, RAID_BAND_OFFICER, RAID_STRIP_RESERVE } from '../lib/guild/raid-layout';
  import { raidQuiet } from '../lib/guild/planned-row';
  import type { RaidStatusAction } from '../lib/guild/raids-api';
  import FactionCrest from './character/FactionCrest.svelte';
  import GuildBandFrame from './guild/GuildBandFrame.svelte';
  import GuildTabs from './guild/GuildTabs.svelte';
  import RaidChip from './guild/RaidChip.svelte';
  import RaidForm from './guild/RaidForm.svelte';
  import RaidOfficerRow from './guild/RaidOfficerRow.svelte';
  import RaidRoster from './guild/RaidRoster.svelte';
  import RaidSignupStrip from './guild/RaidSignupStrip.svelte';
  import SignInPrompt from './SignInPrompt.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  const MINUTE_MS = 60_000;
  const GATE_ID = 'officer-tools-gate-raids';
  /** The band's separator, as an expression: Svelte trims the edge spaces of literal text. */
  const SEP = ' · ';

  let {
    path,
    controller,
    now: fixedNow,
    viewerTimezone = new Intl.DateTimeFormat().resolvedOptions().timeZone,
    locale = 'en-US',
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

  let tickNow = $state(new Date());
  $effect(() => {
    if (fixedNow !== undefined) return;
    const timer = setInterval(() => (tickNow = new Date()), MINUTE_MS);
    return () => clearInterval(timer);
  });
  const now = $derived(fixedNow ?? tickNow);

  const guildName = $derived(c.guild?.name ?? '');
  /** The route's slug stands in for the name until the guild answers, so the loading band's
   *  back link and tabs already point at the guild page. */
  const guildPath = $derived(guildHref(path.guild.region, path.guild.ruleset, guildName || path.guild.slug));
  const settingsHref = $derived(guildSettingsHref(path.guild.region, path.guild.ruleset, guildName));
  const officer = $derived(c.role === 'officer');
  const reserveRows = stripReserveRows(browserStorage());
  const visible = $derived(c.access === 'member');
  const raid = $derived(c.raid);
  const flags = $derived(raid === null ? null : raidPageFlags(raid, now));
  const when = $derived(raid === null ? null : raidWhen(raid, now, viewerTimezone, locale));
  const count = $derived(raid === null ? null : raidCount(raid));
  const guildTimezone = $derived(raid?.guild_timezone ?? viewerTimezone);
  /** No gold on a cancelled raid or one already under way (spec 3.2 lines 4–5). */
  const quiet = $derived(raid !== null && raidQuiet(raid, now));
  const goldText = $derived(quiet ? 'text-muted' : 'text-gold');
  const goldFill = $derived(quiet ? 'bg-muted' : 'bg-gold');
  const countText = $derived(raid?.status === 'cancelled' ? 'text-muted' : 'text-strong');

  let backLink = $state<HTMLAnchorElement | null>(null);

  /** A done cancel leaves only disabled buttons, so focus moves to the back link. */
  async function runStatus(action: RaidStatusAction, post?: boolean): Promise<boolean> {
    const ok = await c.runStatus(action, post);
    if (ok && action === 'cancel') backLink?.focus();
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
    bind:this={backLink}
    data-testid="raid-back-link">{raidCopy.backLink(guildName)}</a
  >
{/snippet}

{#snippet raidEyebrow()}
  <span class="label text-gold flex items-center gap-2.5">
    <i class="bg-gold inline-block h-px w-7" aria-hidden="true"></i>{raidCopy.page.eyebrow}
  </span>
{/snippet}

{#snippet guildHeading()}
  <span class="label text-gold flex items-center gap-2.5">
    <i class="bg-gold inline-block h-px w-7" aria-hidden="true"></i>{raidCopy.page.guildEyebrow}
  </span>
  <div class="flex items-center gap-4 lg:gap-[18px]">
    <FactionCrest faction={c.guild?.faction ?? null} src={c.guild?.mark ?? null} />
    <h1 class="font-display text-strong text-[22px] font-bold">{guildName}</h1>
  </div>
{/snippet}

{#if c.load === 'failed'}
  <div class="flex flex-col gap-[22px] md:gap-8">
    <GuildBandFrame faction={c.guild?.faction ?? null}>
      {#if c.guild !== null}{@render guildHeading()}{:else}<h1 class="sr-only">
          {raidCopy.page.eyebrow}
        </h1>{/if}
    </GuildBandFrame>
    <LoadError message={raidCopy.loadErrors.raid} onRetry={() => c.retry()} testid="raid-load-error" />
  </div>
{:else if c.load === 'loading' || c.access === 'loading'}
  <div class="flex flex-col gap-[22px] md:gap-8" data-testid="raid-page-loading">
    <span class="sr-only">{raidCopy.page.loading}</span>
    <!-- The back link, eyebrow and tabs need only the route, so they are drawn for real; the
         skeleton holds lines 3–5 (and an officer's wrapped phone row) at their ready height. -->
    <GuildBandFrame faction={c.guild?.faction ?? null}>
      {@render backLinkRow()}
      {@render raidEyebrow()}
      <div class="{RAID_BAND_LINES_MIN} {officer ? RAID_BAND_OFFICER : ''}">
        <Skeleton lines={3} rowHeight="h-5" label={raidCopy.page.loading} testid="raid-band-skeleton" />
      </div>
      <GuildTabs active="raids" role={officer ? 'officer' : 'member'} hrefBase={guildPath} />
    </GuildBandFrame>
    <div
      class="bg-raised border-line rounded-panel border {RAID_STRIP_RESERVE}"
      style:--strip-rows={reserveRows}
      aria-hidden="true"
      data-testid="raid-strip-reserve"
    ></div>
    <div class="min-h-[560px] md:min-h-[420px]" aria-hidden="true" data-testid="raid-roster-reserve"></div>
  </div>
{:else if !visible}
  <GuildBandFrame faction={c.guild?.faction ?? null} testid="raid-not-visible">
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
{:else if raid !== null && flags !== null && when !== null && count !== null}
  <div class="reveal flex flex-col gap-[22px] md:gap-8" data-testid="raid-page">
    <GuildBandFrame faction={c.guild?.faction ?? null}>
      {@render backLinkRow()}
      <div class="grid grid-cols-1 gap-3 md:grid-cols-[minmax(0,1fr)_auto]">
        <div class="flex flex-wrap items-center gap-3 md:col-start-1 md:row-start-1">
          {@render raidEyebrow()}
          <RaidChip kind="status" status={flags.status} testid="raid-status-chip" />
        </div>
        {#if c.editing}
          <div class="md:col-span-2">
            <RaidForm
              mode="edit"
              zones={c.zones}
              guildTimezone={raid.guild_timezone}
              {viewerTimezone}
              {raid}
              plannedCount={0}
              discordConnected={raid.discord.connected}
              {settingsHref}
              onSubmit={async (payload) => {
                const result = await c.saveEdit(payload);
                if (result.ok) closeEdit();
                return result;
              }}
              onCancel={closeEdit}
              {now}
              {locale}
            />
          </div>
        {:else}
          <h1 class="font-display text-strong text-[22px] font-bold md:col-span-2" data-testid="raid-title">
            {raid.title}{#if raid.title !== raid.zone_name}<span
                class="text-muted font-sans text-[15px] font-normal whitespace-nowrap"
                >{SEP}{raid.zone_name}</span
              >{/if}
          </h1>
          <!-- Phone: three lines (date; times; guild time and countdown); one line from md. The
               separators are entities: Svelte trims a span's edge spaces, which ran the words
               together from md up. -->
          <p class="text-[15px] md:col-span-2" data-testid="raid-when">
            <span class="block md:inline">{when.date}</span><span class="hidden md:inline">{SEP}</span><span
              class="block md:inline"
              ><span class="font-mono">{when.start}</span>&#32;to&#32;<span class="font-mono">{when.end}</span
              ></span
            ><span class="block md:inline"
              ><span class="hidden md:inline">{SEP}</span>{#if when.guildTime !== null}<span
                  class="text-muted">{when.guildTime}</span
                >{SEP}{/if}<span class="{goldText} font-mono">{when.countdown}</span></span
            >
          </p>
          <div class="flex flex-col gap-2 md:col-span-2" data-testid="raid-count">
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
                ><span class="{countText} font-mono text-[15px]">{count.line}</span><span
                  class="text-muted text-[13px]">{SEP}{count.breakdown}</span
                ></span
              >
            </p>
            {#if c.pollPaused}
              <p class="text-muted text-[13px]" data-testid="raid-poll-paused">{raidCopy.poll}</p>
            {/if}
            {#if flags.done}
              <p class="text-[14px]" data-testid="raid-done">
                {raidCopy.raidStates.done}
                <a href={`${guildPath}#raids`} class="text-gold font-bold">{raidCopy.raidStates.openNight}</a>
              </p>
            {/if}
          </div>
          {#if officer}
            <div class="md:col-start-2 md:row-start-1">
              <RaidOfficerRow
                {raid}
                readOnly={!c.officerTools}
                busy={c.busy}
                line={c.line}
                gateId={GATE_ID}
                {settingsHref}
                {guildTimezone}
                onEdit={() => c.openEdit()}
                onStatus={runStatus}
              />
            </div>
          {/if}
        {/if}
      </div>
      <GuildTabs active="raids" role={officer ? 'officer' : 'member'} hrefBase={guildPath} />
    </GuildBandFrame>

    {#if flags.cancelled && c.raceLine !== null}
      <p
        class="text-wipe -mb-2 text-[13px] md:-mb-4"
        role="alert"
        tabindex="-1"
        use:focusOnMount
        data-testid="raid-race-line"
      >
        {c.raceLine}
      </p>
    {/if}
    <RaidSignupStrip
      {raid}
      locked={flags.locked}
      ended={flags.ended}
      cancelled={flags.cancelled}
      onSignup={c.signup}
      currentKey={c.currentKey}
    />

    <section
      class="flex max-w-[720px] flex-col gap-3"
      aria-labelledby="raid-notes-heading"
      data-testid="raid-notes"
    >
      <h2 id="raid-notes-heading" class="section-title text-[18px]">{raidCopy.subheads.notes}</h2>
      {#if raid.notes_html.trim() === ''}
        <p class="text-muted text-[14px]" data-testid="raid-notes-empty">
          {officer ? raidCopy.notesEmpty.officer : raidCopy.notesEmpty.member}
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

    <RaidRoster
      {raid}
      {officer}
      officerTools={c.officerTools}
      gateId={GATE_ID}
      cancelled={flags.cancelled}
      closed={flags.cancelled || flags.done}
      busyKey={c.rosterBusy}
      line={c.line}
      onSetState={(row, state) => c.setRosterState(row, state)}
    />
  </div>
{/if}
