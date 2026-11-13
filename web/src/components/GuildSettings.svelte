<!-- web/src/components/GuildSettings.svelte -->
<!-- /guild/<region>/<ruleset>/<name>/settings (spec section 4.4). Verified-officer-or-
     leader only; a non-officer sees the honest-copy explanation, never a bare 403. Reads
     GET /v1/guilds/{id}/settings, which itself is officer/leader/moderator-gated server
     side (spec section 2.6) -- a 403 from that call IS the "you are not an officer" signal,
     read directly rather than re-derived from /v1/me client side. -->
<script lang="ts">
  import { guildClaimHref, guildSetupHref, type CharacterPath } from '../lib/characters';
  import { fetchMeFresh, refetchMe, type Me } from '../lib/account/api';
  import { onWindowFocus } from '../lib/data/query';
  import {
    GuildApiError,
    fetchGuildSettings,
    refetchGuildSettings,
    updateGuildSettings,
    type GuildSettingsData,
    type GuildVisibility,
  } from '../lib/guild/api';
  import { guildSettingsCopy } from '../lib/guild/copy';
  import { can } from '../lib/guild/roster-model';
  import { settingsWords, VISIBILITY_OPTIONS } from '../lib/guild/settings-view';
  import { fetchGuildSetup, type GuildSetup } from '../lib/guild/setup-api';
  import { setupCopy } from '../lib/guild/setup-copy';
  import { setupOpen } from '../lib/guild/setup-model';
  import { settingsFacts, stepsLeft } from '../lib/guild/settings-facts';
  import Row from './ui/Row.svelte';
  import Notice from './ui/Notice.svelte';
  import { GUILD_LOADING } from '../lib/guild/layout';
  import { guildDiscordReserve } from '../lib/guild/discord-state';
  import { createLazyComponent } from '../lib/report/lazy-component.svelte';
  import { fetchGuild } from '../lib/rankings/api';
  import type { Snippet } from 'svelte';
  import GuildBillingSection from './guild/GuildBillingSection.svelte';
  import GuildClaimedBy from './guild/GuildClaimedBy.svelte';
  import GuildInviteLink from './guild/GuildInviteLink.svelte';
  import GuildRaiderRanks from './guild/GuildRaiderRanks.svelte';
  import GuildRecords from './guild/billing/GuildRecords.svelte';
  import GuildStatus from './GuildStatus.svelte';
  import Page from './ui/Page.svelte';
  import PageHeader from './ui/PageHeader.svelte';
  import Section from './ui/Section.svelte';
  import Select from './ui/Select.svelte';
  import SettingsRow from './ui/SettingsRow.svelte';
  import { busyCopy } from '../lib/ui/copy';
  import Button from './ui/Button.svelte';

  // The Discord section (raids spec 3.8) is its own chunk: GuildShell's entry is at its budget.
  const discordLazy = createLazyComponent(() => import('./guild/GuildDiscordSection.svelte'));

  let {
    path,
    embedded = false,
    crest,
  }: {
    path: CharacterPath;
    /** Inside the guild page's Settings tab: the page's own head stands, so no h1 of its own. */
    embedded?: boolean;
    /** The guild crest row (officers only), drawn at the foot of the Records card. */
    crest?: Snippet;
  } = $props();

  let guildId = $state<number | null>(null);
  /** Whether this account may release the claim (the `release` word on its /v1/me guild entry). */
  let mayRelease = $state(false);
  let guildName = $state('');
  /** The setup checklist, best-effort: the row above the page shows until every step is settled. */
  let setup = $state<GuildSetup | null>(null);
  /** The setup read failed: the raider ranks row has no rank lines to list. */
  let setupFailed = $state(false);
  let settings = $state<GuildSettingsData | null>(null);
  let status = $state<'loading' | 'ready' | 'forbidden' | 'failed'>('loading');
  let notice = $state('');
  let error = $state('');
  let busy = $state(false);

  /**
   * A 403 from GET .../settings is read directly off the thrown GuildApiError's status --
   * the only signal this page uses to distinguish "you are not an officer" from any other
   * failure. Any other status (network failure, 500, ...) falls through to 'failed' rather
   * than being folded into 'forbidden', so a real outage is never mislabeled as a
   * permissions problem.
   */
  async function load(): Promise<void> {
    status = 'loading';
    try {
      const page = await fetchGuild(path);
      guildId = page.guild.id;
      guildName = page.guild.name;
      settings = await fetchGuildSettings(page.guild.id, undefined, (fresh) => (settings = fresh));
      status = 'ready';
      void fetchGuildSetup(page.guild.id)
        .then((fresh) => {
          setup = fresh;
          setupFailed = false;
        })
        .catch(() => {
          setup = null;
          setupFailed = true;
        });
    } catch (thrown) {
      status = thrown instanceof GuildApiError && thrown.status === 403 ? 'forbidden' : 'failed';
      return;
    }
    // Best-effort: a failed /v1/me here should not take down a settings page that already
    // loaded successfully -- the billing section just falls back to "not subscribed" until
    // a retry (page reload) succeeds, same as any other read this page treats as secondary.
    try {
      const me = await fetchMeFresh(applyBilling);
      applyBilling(me);
    } catch {
      mayRelease = false;
    }
  }

  function applyBilling(me: Me | null): void {
    const entry = me?.guilds.find((g) => g.id === guildId);
    mayRelease = entry === undefined ? false : can(entry.can, 'release');
  }

  $effect(() => {
    void load();
  });

  // A CLI grant or a Stripe webhook cannot notify the browser, so Settings (and the billing
  // block on it) asks again when the player comes back to the window.
  $effect(() => {
    return onWindowFocus(() => {
      if (guildId === null || status !== 'ready') return;
      refetchGuildSettings(guildId, (fresh) => (settings = fresh));
      refetchMe(applyBilling);
    });
  });

  /**
   * Mirrors GuildClaim.svelte's `run()`: every action below goes through here so a
   * rejected `updateGuildSettings`/`rotateInvite` call (network failure, validation
   * error, 500) always ends in a visible `error` message rather than `busy` silently
   * resetting with nothing shown -- the failure-swallowing bug this wrapper exists to
   * rule out.
   */
  async function run(action: () => Promise<void>): Promise<void> {
    busy = true;
    notice = '';
    error = '';
    try {
      await action();
    } catch (thrown) {
      error = thrown instanceof Error ? thrown.message : guildSettingsCopy.actionFailed;
    } finally {
      busy = false;
    }
  }

  /** The officer rank as typed: it saves when Save is pressed, and the line under it says it took. */
  let rankDraft = $derived<number | null>(settings?.officer_max_rank_index ?? null);
  let rankSaved = $state(false);
  const rankChanged = $derived(
    settings !== null &&
      typeof rankDraft === 'number' &&
      rankDraft >= 0 &&
      rankDraft !== settings.officer_max_rank_index,
  );

  const onSave = (field: 'default_visibility' | 'officer_max_rank_index', value: string): void =>
    void run(async () => {
      if (guildId === null) return;
      const patch =
        field === 'default_visibility'
          ? { default_visibility: value as GuildVisibility }
          : { officer_max_rank_index: Number(value) };
      settings = await updateGuildSettings(guildId, patch);
      if (field === 'officer_max_rank_index') rankSaved = true;
      else notice = guildSettingsCopy.saved;
    });

  /** The raider ranks row's last save took; cleared by its next change. */
  let raiderRanksSaved = $state(false);

  const onSaveRaiderRanks = (ranks: number[]): void =>
    void run(async () => {
      if (guildId === null) return;
      settings = await updateGuildSettings(guildId, { raider_rank_indexes: ranks });
      raiderRanksSaved = true;
    });

  /**
   * A contest always freezes officer tools now (a later security-review response
   * simplified the freeze rule): `frozen` is guaranteed true whenever `state ===
   * 'contested'`, so there is no longer a contested-but-not-frozen case to render
   * separately. `settings.claim.frozen` is still read explicitly rather than assumed,
   * since the API's own response shape is the source of truth.
   */
  const frozen = $derived(
    settings !== null && settings.claim.state === 'contested' && settings.claim.frozen === true,
  );

  $effect(() => {
    if (settings?.discord !== undefined) discordLazy.load();
  });
</script>

{#snippet notes()}
  {#if settings !== null}
    {#if settings.claim.state === 'claimed'}
      <p class="text-[13px]" data-testid="guild-settings-claim-state">
        <GuildClaimedBy holder={settings.claimed_by} />{#if mayRelease}&#32; ·
          <a
            class="underline"
            href={guildClaimHref(path.region, path.ruleset, guildName)}
            data-testid="guild-settings-release">{settingsWords.releaseOnClaimPage}</a
          >{/if}
      </p>
    {/if}
    {#if settings.claim.state === 'contested' && frozen}
      <p class="text-[13px]" role="alert" data-testid="guild-settings-frozen">
        {guildSettingsCopy.frozenNotice}
      </p>
    {/if}
  {/if}
{/snippet}

{#snippet membersSection(current: GuildSettingsData)}
  <Section title={guildSettingsCopy.membersHeading} panel testid="guild-settings-members">
    {@render notes()}
    {#if setup !== null}
      <ul class="pb-2" data-testid="guild-settings-facts">
        {#each settingsFacts(setup) as fact (fact.key)}
          <li><Row title={fact.title} sub={fact.sub === '' ? undefined : fact.sub} density="slim" open /></li>
        {/each}
      </ul>
    {/if}
    {#if guildId !== null}
      <GuildInviteLink
        {guildId}
        rotatedAt={current.invite.rotated_at}
        keptUrl={current.invite.url}
        madeBy={current.invite.made_by}
        disabled={busy || frozen}
        onRotated={(invite) => {
          if (settings !== null) settings = { ...settings, invite };
        }}
      />
    {/if}
    <SettingsRow open label={settingsWords.officerRank} controlBelow>
      {#snippet control(labelId)}
        <div class="flex items-center gap-3">
          <input
            id="guild-officer-threshold"
            type="number"
            min="0"
            aria-labelledby={labelId}
            class="rounded-field text-strong h-11 w-[72px] border border-white/15 bg-white/[.045] px-3 text-[15px] font-semibold disabled:opacity-45"
            bind:value={rankDraft}
            oninput={() => (rankSaved = false)}
            disabled={busy || frozen}
            aria-busy={busy}
            data-testid="guild-officer-threshold-input"
          />
          <Button
            size="sm"
            onclick={() => onSave('officer_max_rank_index', String(rankDraft))}
            busy={busy && rankChanged}
            busyLabel={busyCopy.saving}
            disabled={busy || frozen || !rankChanged}
            testid="guild-officer-threshold-save"
          >
            Save
          </Button>
        </div>
      {/snippet}
      {#snippet below()}
        {#if rankSaved}
          <span class="text-muted text-[13px]" role="status" data-testid="guild-officer-threshold-saved"
            >{guildSettingsCopy.saved}</span
          >
        {/if}
      {/snippet}
    </SettingsRow>
    <GuildRaiderRanks
      lines={setupFailed ? 'failed' : (setup?.ranks.ranks ?? null)}
      officerMax={current.officer_max_rank_index}
      saved={current.raider_rank_indexes}
      {busy}
      disabled={frozen}
      justSaved={raiderRanksSaved}
      onSave={onSaveRaiderRanks}
      onEdit={() => (raiderRanksSaved = false)}
    />
    <SettingsRow open label={guildSettingsCopy.defaultVisibility} controlBelow>
      {#snippet control()}
        <div class="w-40">
          <Select
            id="guild-visibility"
            size="compact"
            label={guildSettingsCopy.defaultVisibility}
            options={VISIBILITY_OPTIONS}
            value={current.default_visibility}
            onchange={(event) =>
              onSave('default_visibility', (event.currentTarget as HTMLSelectElement).value)}
            disabled={busy || frozen}
            testid="guild-visibility-select"
          />
        </div>
      {/snippet}
    </SettingsRow>
  </Section>
{/snippet}

{#snippet recordsSection(id: number)}
  <Section
    panel
    title={crest === undefined ? guildSettingsCopy.recordsOnlyHeading : guildSettingsCopy.recordsHeading}
    testid="guild-settings-records"
  >
    <GuildRecords guildId={id} />
    {@render crest?.()}
  </Section>
{/snippet}

<div class="reveal flex flex-col" data-testid="guild-settings">
  {#if embedded}
    <h2 class="sr-only">{guildSettingsCopy.heading}</h2>
  {:else}
    <PageHeader
      title={guildName === '' ? guildSettingsCopy.heading : guildName}
      tone="strong"
      testid="guild-settings-header"
    >
      {#snippet meta()}<span>{guildSettingsCopy.heading}</span>{/snippet}
    </PageHeader>
  {/if}
  {#if status === 'loading' || status === 'failed'}
    <GuildStatus
      status={status === 'loading' ? 'loading' : 'failed'}
      error={guildSettingsCopy.failed}
      onRetry={() => void load()}
      lines={GUILD_LOADING.settings.lines}
      minHeight={GUILD_LOADING.settings.minHeight}
      testid="guild-settings"
    />
    <!-- The crest does not wait on the settings read: a failed read leaves it usable. -->
    {#if status === 'failed' && guildId !== null && crest !== undefined}
      <div class="max-w-[440px]">{@render recordsSection(guildId)}</div>
    {/if}
  {:else if status === 'forbidden'}
    <p class="text-[14px]" data-testid="guild-settings-forbidden">{guildSettingsCopy.forbidden}</p>
  {:else if settings !== null}
    {@const current = settings}
    {#if setup !== null && setupOpen(setup)}
      <div class="px-(--page-inset) pb-3">
        <Notice
          text={[
            setupCopy.settingsRow.label,
            setupCopy.settingsRow.progress(setup.done, setup.total),
            ...stepsLeft(setup).slice(0, 2),
          ].join(' · ')}
          action={{
            label: setupCopy.settingsRow.link,
            href: guildSetupHref(path.region, path.ruleset, guildName || path.slug),
          }}
          testid="guild-settings-setup-row"
        />
      </div>
    {/if}
    <Page layout="split" panels testid="guild-settings-page">
      {#if guildId !== null}
        <GuildBillingSection {guildId} {path} {guildName} />
      {/if}
      {#if settings.discord !== undefined && guildId !== null}
        {#if discordLazy.current}
          <discordLazy.current
            {guildId}
            discord={settings.discord}
            {frozen}
            onDiscord={(discord) => {
              if (settings !== null) settings = { ...settings, discord };
            }}
            onSettings={(fresh) => (settings = fresh)}
          />
        {:else}
          <div
            id="discord"
            class="scroll-mt-20 {guildDiscordReserve(settings.discord, window.location.search)}"
            data-testid="guild-settings-discord-reserve"
          ></div>
        {/if}
      {/if}
      {#snippet rail()}
        {@render membersSection(current)}
        {#if guildId !== null}{@render recordsSection(guildId)}{/if}
      {/snippet}
    </Page>
    <div class="min-h-[21px] px-(--page-inset) pt-3">
      {#if error !== ''}<p class="text-[14px]" role="alert" data-testid="guild-settings-action-error">
          {error}
        </p>{:else if notice !== ''}<p class="text-[14px]" data-testid="guild-settings-notice">
          {notice}
        </p>{/if}
    </div>
  {/if}
</div>
