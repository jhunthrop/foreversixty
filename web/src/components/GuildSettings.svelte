<!-- web/src/components/GuildSettings.svelte -->
<!-- /guild/<region>/<ruleset>/<name>/settings (spec section 4.4). Verified-officer-or-
     leader only; a non-officer sees the honest-copy explanation, never a bare 403. Reads
     GET /v1/guilds/{id}/settings, which itself is officer/leader/moderator-gated server
     side (spec section 2.6) -- a 403 from that call IS the "you are not an officer" signal,
     read directly rather than re-derived from /v1/me client side. -->
<script lang="ts">
  import { GUILD_BANDLESS_TOP } from '../lib/guild/raid-layout';
  import { guildClaimHref, type CharacterPath } from '../lib/characters';
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
  import { claimedByLine, settingsWords } from '../lib/guild/settings-view';
  import { GUILD_LOADING } from '../lib/guild/layout';
  import { guildDiscordReserve } from '../lib/guild/discord-state';
  import { createLazyComponent } from '../lib/report/lazy-component.svelte';
  import { fetchGuild } from '../lib/rankings/api';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import GuildBillingSection from './guild/GuildBillingSection.svelte';
  import GuildInviteLink from './guild/GuildInviteLink.svelte';
  import GuildStatus from './GuildStatus.svelte';
  import { BUSY_CLASS } from '../lib/ui/busy';

  // The Discord section (raids spec 3.8) is its own chunk: GuildShell's entry is at its budget.
  const discordLazy = createLazyComponent(() => import('./guild/GuildDiscordSection.svelte'));

  let {
    path,
    officerTools,
    onPlanRefused,
  }: {
    path: CharacterPath;
    /** From the guild page's Settings tab; the standalone route leaves it out. */
    officerTools?: boolean;
    onPlanRefused?: () => void;
  } = $props();

  let guildId = $state<number | null>(null);
  /** Whether this account may release the claim (the `release` word on its /v1/me guild entry). */
  let mayRelease = $state(false);
  let guildName = $state('');
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

<div class="reveal {GUILD_BANDLESS_TOP} flex flex-col gap-6" data-testid="guild-settings">
  <h1 class="heading-page">
    {guildName === '' ? guildSettingsCopy.heading : `${guildName} · ${guildSettingsCopy.heading}`}
  </h1>
  {#if status === 'loading' || status === 'failed'}
    <GuildStatus
      status={status === 'loading' ? 'loading' : 'failed'}
      error={guildSettingsCopy.failed}
      onRetry={() => void load()}
      lines={GUILD_LOADING.settings.lines}
      minHeight={GUILD_LOADING.settings.minHeight}
      testid="guild-settings"
    />
  {:else if status === 'forbidden'}
    <p class="text-[14px]" data-testid="guild-settings-forbidden">{guildSettingsCopy.forbidden}</p>
  {:else if settings !== null}
    {#if claimedByLine(settings.claim, settings.claimed_by) !== null}
      <p class="text-[13px]" data-testid="guild-settings-claim-state">
        {claimedByLine(settings.claim, settings.claimed_by)}{#if mayRelease}&#32; ·
          <a href={guildClaimHref(path.region, path.ruleset, guildName)} data-testid="guild-settings-release"
            >{settingsWords.releaseOnClaimPage}</a
          >{/if}
      </p>
    {/if}
    {#if settings.claim.state === 'contested' && frozen}
      <p class="text-[13px]" role="alert" data-testid="guild-settings-frozen">
        {guildSettingsCopy.frozenNotice}
      </p>
    {/if}
    {#if guildId !== null}
      <GuildBillingSection {guildId} {path} {guildName} />
    {/if}
    {#if settings.discord !== undefined && guildId !== null}
      {#if discordLazy.current}
        <discordLazy.current
          {guildId}
          discord={settings.discord}
          {frozen}
          {officerTools}
          {onPlanRefused}
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
    <section class="flex flex-col gap-3">
      <label class="quiet-label" for="guild-visibility">{guildSettingsCopy.defaultVisibility}</label>
      <p class="text-muted text-[13px]">{settingsWords.visibilityHelp}</p>
      <select
        id="guild-visibility"
        class={`border-line-warm bg-raised rounded-control text-text h-11 w-fit px-3 text-[14px] ${busy ? BUSY_CLASS : ''}`}
        value={settings.default_visibility}
        onchange={(event) => onSave('default_visibility', (event.currentTarget as HTMLSelectElement).value)}
        disabled={busy || frozen}
        aria-busy={busy}
        data-testid="guild-visibility-select"
      >
        <option value="public">Public</option>
        <option value="unlisted">Unlisted</option>
        <option value="guild">Guild only</option>
      </select>
    </section>
    <section class="flex flex-col gap-3">
      <label class="quiet-label" for="guild-officer-threshold">{settingsWords.officerRank}</label>
      <p class="text-muted text-[13px]">{settingsWords.officerRankHelp}</p>
      <div class="flex flex-wrap items-center gap-3">
        <input
          id="guild-officer-threshold"
          type="number"
          min="0"
          class={`border-line-warm bg-raised rounded-control text-text h-11 w-24 px-3 text-[14px] ${busy ? BUSY_CLASS : ''}`}
          bind:value={rankDraft}
          oninput={() => (rankSaved = false)}
          disabled={busy || frozen}
          aria-busy={busy}
          data-testid="guild-officer-threshold-input"
        />
        <button
          class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4 disabled:opacity-50 ${busy ? BUSY_CLASS : ''}`}
          onclick={() => onSave('officer_max_rank_index', String(rankDraft))}
          disabled={busy || frozen || !rankChanged}
          data-testid="guild-officer-threshold-save"
        >
          Save
        </button>
        {#if rankSaved}
          <span class="text-muted text-[13px]" role="status" data-testid="guild-officer-threshold-saved"
            >{guildSettingsCopy.saved}</span
          >
        {/if}
      </div>
    </section>
    <section class="flex flex-col gap-3">
      <h2 class="section-title">{guildSettingsCopy.inviteHeading}</h2>
      <p class="text-muted text-[13px]">{guildSettingsCopy.inviteWarning}</p>
      {#if guildId !== null}
        <GuildInviteLink
          {guildId}
          rotatedAt={settings.invite.rotated_at}
          disabled={busy || frozen}
          onRotated={(rotatedAt) => {
            if (settings !== null) settings = { ...settings, invite: { rotated_at: rotatedAt } };
          }}
        />
      {/if}
    </section>
    <div class="min-h-[21px]">
      {#if error !== ''}<p class="text-[14px]" role="alert" data-testid="guild-settings-action-error">
          {error}
        </p>{:else if notice !== ''}<p class="text-[14px]" data-testid="guild-settings-notice">
          {notice}
        </p>{/if}
    </div>
  {/if}
</div>
