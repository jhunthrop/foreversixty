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
  import { browserStorage } from '../lib/guild/planned-reserve';
  import { raidCopy } from '../lib/guild/raid-copy';
  import { raidCount, raidPageFlags, raidWhen, stripReserveRows } from '../lib/guild/raid-page';
  import { RaidPageController, type RaidRoute } from '../lib/guild/raid-page-controller.svelte';
  import { liveRaidPageDeps } from '../lib/guild/raid-page-deps';
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
  const BAND_MIN = 'min-h-[300px] md:min-h-[232px]';

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
  const guildPath = $derived(guildHref(path.guild.region, path.guild.ruleset, guildName));
  const settingsHref = $derived(guildSettingsHref(path.guild.region, path.guild.ruleset, guildName));
  const officer = $derived(c.role === 'officer');
  const reserveRows = stripReserveRows(browserStorage());
  const visible = $derived(c.access === 'member');
  const raid = $derived(c.raid);
  const flags = $derived(raid === null ? null : raidPageFlags(raid, now));
  const when = $derived(raid === null ? null : raidWhen(raid, now, viewerTimezone, locale));
  const count = $derived(raid === null ? null : raidCount(raid));
  const guildTimezone = $derived(raid?.guild_timezone ?? viewerTimezone);

  function closeEdit(): void {
    c.closeEdit();
    queueMicrotask(() => document.getElementById('raid-officer-edit')?.focus());
  }
</script>

{#snippet guildHeading()}
  <span class="label text-gold flex items-center gap-2.5">
    <i class="bg-gold inline-block h-px w-7" aria-hidden="true"></i>Guild
  </span>
  <div class="flex items-center gap-4 lg:gap-[18px]">
    <FactionCrest faction={c.guild?.faction ?? null} src={c.guild?.mark ?? null} />
    <h1 class="font-display text-strong text-[22px] font-bold">{guildName}</h1>
  </div>
{/snippet}

{#if c.load === 'loading' || c.access === 'loading'}
  <div class="flex flex-col gap-[22px] md:gap-8" data-testid="raid-page-loading">
    <span class="sr-only">{raidCopy.page.loading}</span>
    <GuildBandFrame faction={c.guild?.faction ?? null} minHeight={BAND_MIN}>
      <Skeleton lines={3} rowHeight="h-5" label={raidCopy.page.loading} testid="raid-band-skeleton" />
    </GuildBandFrame>
    <div
      class="bg-raised border-line rounded-panel h-[calc(56px+var(--strip-rows)*196px)] border md:h-[calc(60px+var(--strip-rows)*68px)]"
      style:--strip-rows={reserveRows}
      aria-hidden="true"
      data-testid="raid-strip-reserve"
    ></div>
    <div class="min-h-[560px] md:min-h-[420px]" aria-hidden="true" data-testid="raid-roster-reserve"></div>
  </div>
{:else if c.load === 'failed'}
  <div class="flex flex-col gap-[22px] md:gap-8">
    <GuildBandFrame faction={c.guild?.faction ?? null}>
      {#if c.guild !== null}{@render guildHeading()}{:else}<h1 class="sr-only">
          {raidCopy.page.eyebrow}
        </h1>{/if}
    </GuildBandFrame>
    <LoadError message={raidCopy.loadErrors.raid} onRetry={() => c.retry()} testid="raid-load-error" />
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
      <a
        href={`${guildPath}#raids`}
        class="text-muted inline-flex min-h-11 items-center self-start text-[13px] underline"
        data-testid="raid-back-link">{raidCopy.backLink(guildName)}</a
      >
      <div class="grid grid-cols-1 gap-3 md:grid-cols-[minmax(0,1fr)_auto]">
        <div class="flex flex-wrap items-center gap-3 md:col-start-1 md:row-start-1">
          <span class="label text-gold flex items-center gap-2.5">
            <i class="bg-gold inline-block h-px w-7" aria-hidden="true"></i>{raidCopy.page.eyebrow}
          </span>
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
              >
                · {raid.zone_name}</span
              >{/if}
          </h1>
          <p class="text-[15px] md:col-span-2" data-testid="raid-when">
            <span class="block md:inline">{when.date}</span><span class="hidden md:inline"> · </span><span
              class="block md:inline"
              ><span class="font-mono">{when.start}</span> to <span class="font-mono">{when.end}</span></span
            ><span class="block md:inline"
              >{#if when.guildTime !== null}<span class="text-muted"
                  ><span class="hidden md:inline"> · </span>{when.guildTime}</span
                >
              {/if}<span class="text-gold font-mono"
                ><span class="hidden md:inline"> · </span>{when.countdown}</span
              ></span
            >
          </p>
          <div class="flex flex-col gap-2 md:col-span-2" data-testid="raid-count">
            <p class="flex flex-wrap items-center gap-x-3 gap-y-1">
              {#if count.fill !== null}
                <span
                  class="bg-line-soft relative inline-block h-2 w-[120px] overflow-hidden rounded-full md:w-[160px]"
                  aria-hidden="true"
                  ><span class="bg-gold absolute inset-y-0 left-0" style:width={`${count.fill}%`}
                  ></span></span
                >
              {/if}
              <span
                ><span class="text-strong font-mono text-[15px]">{count.line}</span><span
                  class="text-muted text-[13px]"
                >
                  · {count.breakdown}</span
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
                onStatus={(action, post) => c.runStatus(action, post)}
              />
            </div>
          {/if}
        {/if}
      </div>
      <GuildTabs active="raids" role={officer ? 'officer' : 'member'} hrefBase={guildPath} />
    </GuildBandFrame>

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
        <div class="prose-forever [&_a]:text-gold text-[14px] leading-[1.6]">
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
      busyKey={c.rosterBusy}
      line={c.line}
      onSetState={(row, state) => c.setRosterState(row, state)}
    />
  </div>
{/if}
