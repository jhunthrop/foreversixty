<!-- web/src/components/GuildSettings.svelte -->
<!-- /guild/<region>/<ruleset>/<name>/settings (spec section 4.4). Verified-officer-or-
     leader only; a non-officer sees the honest-copy explanation, never a bare 403. Reads
     GET /v1/guilds/{id}/settings, which itself is officer/leader/moderator-gated server
     side (spec section 2.6) -- a 403 from that call IS the "you are not an officer" signal,
     read directly rather than re-derived from /v1/me client side. -->
<script lang="ts">
  import type { CharacterPath } from '../lib/characters';
  import { fetchMeOnce, type GuildBillingView } from '../lib/account/api';
  import { openPortal } from '../lib/billing/api';
  import { billingBlockCopy } from '../lib/billing/copy';
  import {
    GuildApiError,
    fetchGuildSettings,
    rotateInvite,
    updateGuildSettings,
    type GuildSettingsData,
    type GuildVisibility,
  } from '../lib/guild/api';
  import { guildSettingsCopy } from '../lib/guild/copy';
  import { fetchGuild } from '../lib/rankings/api';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';

  let { path }: { path: CharacterPath } = $props();

  let guildId = $state<number | null>(null);
  let guildName = $state('');
  let settings = $state<GuildSettingsData | null>(null);
  let status = $state<'loading' | 'ready' | 'forbidden' | 'failed'>('loading');
  let notice = $state('');
  let error = $state('');
  let rotated = $state<{ token: string; url: string } | null>(null);
  let busy = $state(false);
  /**
   * A guild's billing view is not part of GET /v1/guilds/{id}/settings's own response (the
   * entitlements work landed after that endpoint did, in a lane that could not edit it) --
   * it lives on GET /v1/me's guilds[] entry for this same guild instead (spec 1.4). This
   * page already knows guildId once `load()` resolves, so finding the matching entry is a
   * plain lookup, not a second "which guild" round-trip; the one extra cost is the /v1/me
   * fetch itself, shared via fetchMeOnce's own cache with anything else on the page that
   * already asked.
   */
  let guildBilling = $state<GuildBillingView | null>(null);

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
      settings = await fetchGuildSettings(page.guild.id);
      status = 'ready';
    } catch (thrown) {
      status = thrown instanceof GuildApiError && thrown.status === 403 ? 'forbidden' : 'failed';
      return;
    }
    // Best-effort: a failed /v1/me here should not take down a settings page that already
    // loaded successfully -- the billing section just falls back to "not subscribed" until
    // a retry (page reload) succeeds, same as any other read this page treats as secondary.
    try {
      const me = await fetchMeOnce();
      guildBilling = me?.guilds.find((g) => g.id === guildId)?.plan ?? null;
    } catch {
      guildBilling = null;
    }
  }

  $effect(() => {
    void load();
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

  const onSave = (field: 'default_visibility' | 'officer_max_rank_index', value: string): void =>
    void run(async () => {
      if (guildId === null) return;
      const patch =
        field === 'default_visibility'
          ? { default_visibility: value as GuildVisibility }
          : { officer_max_rank_index: Number(value) };
      settings = await updateGuildSettings(guildId, patch);
      notice = guildSettingsCopy.saved;
    });

  const onRotate = (): void =>
    void run(async () => {
      if (guildId === null) return;
      const result = await rotateInvite(guildId);
      rotated = { token: result.token, url: result.url };
      if (settings !== null) settings = { ...settings, invite: { rotated_at: result.rotated_at } };
    });

  const onManageBilling = (): void =>
    void run(async () => {
      if (guildId === null) return;
      const result = await openPortal(guildId);
      window.location.assign(result.portal_url);
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
</script>

<div class="flex flex-col gap-6" data-testid="guild-settings">
  <h1 class="section-title text-[18px]">
    {guildName === '' ? guildSettingsCopy.heading : `${guildName} · ${guildSettingsCopy.heading}`}
  </h1>
  {#if status === 'loading'}
    <p class="text-muted text-[14px]">{guildSettingsCopy.loading}</p>
  {:else if status === 'forbidden'}
    <p class="text-[14px]" data-testid="guild-settings-forbidden">{guildSettingsCopy.forbidden}</p>
  {:else if status === 'failed'}
    <p class="text-[14px]" role="alert" data-testid="guild-settings-error">{guildSettingsCopy.failed}</p>
  {:else if settings !== null}
    {#if settings.claim.state !== 'unclaimed' && settings.claim.state !== 'contested'}
      <p class="text-[13px]" data-testid="guild-settings-claim-state">
        Claim: {settings.claim.state}
      </p>
    {/if}
    {#if settings.claim.state === 'contested' && frozen}
      <p class="text-[13px]" role="alert" data-testid="guild-settings-frozen">
        {guildSettingsCopy.frozenNotice}
      </p>
    {/if}
    <section class="flex flex-col gap-3" data-testid="guild-settings-billing">
      <h2 class="section-title text-[18px]">Billing</h2>
      {#if guildBilling === null}
        <p class="text-[14px]">{billingBlockCopy.guildNotSubscribed}</p>
        {#if guildId !== null}
          <a
            href={`/premium/checkout?plan=guild&interval=monthly&guild_id=${guildId}`}
            class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4"
            data-testid="guild-subscribe-link"
          >
            {billingBlockCopy.subscribeTheGuild}
          </a>
        {/if}
      {:else}
        <p class="text-[14px]">
          {guildBilling.cancel_at_period_end ? billingBlockCopy.ends : billingBlockCopy.renews}
          {guildBilling.current_period_end
            ? new Date(guildBilling.current_period_end).toLocaleDateString()
            : ''}
          {#if guildBilling.billed_by !== ''}
            · {billingBlockCopy.guildBilledBy(guildBilling.billed_by)}
          {/if}
        </p>
        {#if guildBilling.status === 'past_due'}
          <p class="text-strong text-[13px]" role="alert">{billingBlockCopy.pastDueBanner}</p>
        {/if}
        {#if guildBilling.you_are_billing_contact}
          <button
            class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-4"
            onclick={onManageBilling}
            disabled={busy}
            data-testid="guild-manage-billing"
          >
            {billingBlockCopy.manageBilling}
          </button>
        {/if}
      {/if}
    </section>
    <section class="flex flex-col gap-3">
      <label class="label text-muted" for="guild-visibility">{guildSettingsCopy.defaultVisibility}</label>
      <select
        id="guild-visibility"
        class="border-line-warm bg-raised rounded-control text-text h-11 w-fit px-3 text-[14px]"
        value={settings.default_visibility}
        onchange={(event) => onSave('default_visibility', (event.currentTarget as HTMLSelectElement).value)}
        disabled={busy || frozen}
        data-testid="guild-visibility-select"
      >
        <option value="public">Public</option>
        <option value="unlisted">Unlisted</option>
        <option value="guild">Guild only</option>
      </select>
    </section>
    <section class="flex flex-col gap-3">
      <label class="label text-muted" for="guild-officer-threshold"
        >{guildSettingsCopy.officerThreshold}</label
      >
      <input
        id="guild-officer-threshold"
        type="number"
        min="0"
        class="border-line-warm bg-raised rounded-control text-text h-11 w-24 px-3 text-[14px]"
        value={settings.officer_max_rank_index}
        onchange={(event) =>
          onSave('officer_max_rank_index', (event.currentTarget as HTMLInputElement).value)}
        disabled={busy || frozen}
        data-testid="guild-officer-threshold-input"
      />
    </section>
    <section class="flex flex-col gap-3">
      <h2 class="section-title text-[18px]">{guildSettingsCopy.inviteHeading}</h2>
      <p class="text-muted text-[13px]">{guildSettingsCopy.inviteWarning}</p>
      <button
        class="border-line-warm-strong rounded-control text-strong inline-flex h-11 w-fit items-center border px-4 text-[12px] font-bold tracking-[0.06em] uppercase"
        onclick={onRotate}
        disabled={busy || frozen}
        data-testid="guild-invite-rotate"
      >
        {guildSettingsCopy.rotateButton}
      </button>
      {#if rotated !== null}
        <p class="text-[13px]" data-testid="guild-invite-token">
          {guildSettingsCopy.tokenShownOnce}
          <code class="font-mono">{rotated.url}</code>
        </p>
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
