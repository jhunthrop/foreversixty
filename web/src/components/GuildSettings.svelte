<!-- web/src/components/GuildSettings.svelte -->
<!-- /guild/<region>/<ruleset>/<name>/settings (spec section 4.4). Verified-officer-or-
     leader only; a non-officer sees the honest-copy explanation, never a bare 403. Reads
     GET /v1/guilds/{id}/settings, which itself is officer/leader/moderator-gated server
     side (spec section 2.6) -- a 403 from that call IS the "you are not an officer" signal,
     read directly rather than re-derived from /v1/me client side. -->
<script lang="ts">
  import type { CharacterPath } from '../lib/characters';
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

  let { path }: { path: CharacterPath } = $props();

  let guildId = $state<number | null>(null);
  let guildName = $state('');
  let settings = $state<GuildSettingsData | null>(null);
  let status = $state<'loading' | 'ready' | 'forbidden' | 'failed'>('loading');
  let notice = $state('');
  let rotated = $state<{ token: string; url: string } | null>(null);
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
      settings = await fetchGuildSettings(page.guild.id);
      status = 'ready';
    } catch (thrown) {
      status = thrown instanceof GuildApiError && thrown.status === 403 ? 'forbidden' : 'failed';
    }
  }

  $effect(() => {
    void load();
  });

  async function onSave(
    field: 'default_visibility' | 'officer_max_rank_index',
    value: string,
  ): Promise<void> {
    if (guildId === null) return;
    busy = true;
    notice = '';
    try {
      const patch =
        field === 'default_visibility'
          ? { default_visibility: value as GuildVisibility }
          : { officer_max_rank_index: Number(value) };
      settings = await updateGuildSettings(guildId, patch);
      notice = guildSettingsCopy.saved;
    } finally {
      busy = false;
    }
  }

  async function onRotate(): Promise<void> {
    if (guildId === null) return;
    busy = true;
    try {
      const result = await rotateInvite(guildId);
      rotated = { token: result.token, url: result.url };
      if (settings !== null) settings = { ...settings, invite: { rotated_at: result.rotated_at } };
    } finally {
      busy = false;
    }
  }
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
    <section class="flex flex-col gap-3">
      <label class="label text-muted" for="guild-visibility">{guildSettingsCopy.defaultVisibility}</label>
      <select
        id="guild-visibility"
        class="border-line-warm bg-raised rounded-control text-text h-11 w-fit px-3 text-[14px]"
        value={settings.default_visibility}
        onchange={(event) =>
          void onSave('default_visibility', (event.currentTarget as HTMLSelectElement).value)}
        disabled={busy}
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
          void onSave('officer_max_rank_index', (event.currentTarget as HTMLInputElement).value)}
        disabled={busy}
        data-testid="guild-officer-threshold-input"
      />
    </section>
    <section class="flex flex-col gap-3">
      <h2 class="section-title text-[18px]">{guildSettingsCopy.inviteHeading}</h2>
      <p class="text-muted text-[13px]">{guildSettingsCopy.inviteWarning}</p>
      <button
        class="border-line-warm-strong rounded-control text-strong inline-flex h-11 w-fit items-center border px-4 text-[12px] font-bold tracking-[0.06em] uppercase"
        onclick={() => void onRotate()}
        disabled={busy}
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
      {#if notice !== ''}<p class="text-[14px]" data-testid="guild-settings-notice">{notice}</p>{/if}
    </div>
  {/if}
</div>
