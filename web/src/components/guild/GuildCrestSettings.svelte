<!-- web/src/components/guild/GuildCrestSettings.svelte -->
<!-- Guild crest round (docs/contracts/2026-10-05-guild-crest-api.md): the Settings tab's
     "Guild crest" block, officer only (GuildSettingsTab.svelte's own `role === 'officer'`
     gate -- this component never re-checks that itself, the same trust-the-caller pattern
     GuildOfficerStrip.svelte already uses for its own officer-only controls).

     Tailwind utility classes only, no scoped <style> block: this component only ever
     mounts inside the Settings tab's lazily-loaded chunk (Guild.svelte's
     `createLazyComponent`), which is absent from the page's server-rendered markup the
     same way every other guild/* tab component is (design/reviews/2026-10-04-guild-
     centre-build-report.md, "Header round": the build's CSS-extraction pass only ships a
     component's scoped styles for markup present in the SSR output, silently dropping
     anything that only ever renders after a client-side fetch resolves).

     The 64px ring here is a fixed size, never FactionCrest's own 64/44px responsive pair
     -- this block is a Settings-page control, not the header identity row, so it stays
     one size at every width (matches the contract's own "a 64px ring" wording, singular). -->
<script lang="ts">
  import type { Faction } from '../../lib/faction-mark';
  import { FACTION_BAR_COLOR } from '../../lib/faction-mark';
  import { deleteGuildCrest, GuildApiError, putGuildCrest } from '../../lib/guild/api';
  import { guildCrestCopy, guildHomeCopy } from '../../lib/guild/copy';
  import { CREST_ACCEPT_ATTR, CREST_RULES_LINE, precheckCrestFile } from '../../lib/guild/crest';
  import { guildMarkSrc } from '../../lib/guild/mark';
  import { BUSY_CLASS } from '../../lib/ui/busy';
  import { button } from '../../lib/ui/button';
  import { ROW, ROW_LABEL, ROW_SUB } from '../../lib/ui/row';

  let {
    guildId,
    guild,
    onCrestChanged,
  }: {
    guildId: number;
    guild: { faction?: Faction | null; crest_url?: string | null };
    onCrestChanged: () => void;
  } = $props();

  let fileInputEl = $state<HTMLInputElement | undefined>(undefined);
  let selectedFile = $state<File | null>(null);
  let previewUrl = $state<string | null>(null);
  let fileError = $state('');
  let actionError = $state('');
  let busy = $state<'saving' | 'removing' | null>(null);
  let confirmingRemove = $state(false);

  const hasCrest = $derived(guild.crest_url != null);
  /** The current mark (crest or faction logo), read through the one shared helper -- never
   *  a second `crest_url ?? factionLogoSrc(faction)` copy (the contract's own rule). */
  const currentMark = $derived(guildMarkSrc(guild));
  const ringStyle = $derived(
    guild.faction === 'alliance' || guild.faction === 'horde'
      ? `box-shadow: 0 0 0 2px ${FACTION_BAR_COLOR[guild.faction]};`
      : '',
  );

  function clearPreview(): void {
    if (previewUrl !== null) URL.revokeObjectURL(previewUrl);
    previewUrl = null;
    selectedFile = null;
  }

  function onFileChange(): void {
    fileError = '';
    actionError = '';
    clearPreview();
    const file = fileInputEl?.files?.[0] ?? null;
    if (file === null) return;
    const precheckError = precheckCrestFile(file);
    if (precheckError !== null) {
      fileError = precheckError;
      if (fileInputEl !== undefined) fileInputEl.value = '';
      return;
    }
    selectedFile = file;
    previewUrl = URL.createObjectURL(file);
  }

  async function onSave(): Promise<void> {
    if (selectedFile === null) return;
    busy = 'saving';
    actionError = '';
    try {
      await putGuildCrest(guildId, selectedFile);
      clearPreview();
      if (fileInputEl !== undefined) fileInputEl.value = '';
      onCrestChanged();
    } catch (thrown) {
      actionError = thrown instanceof GuildApiError ? thrown.message : guildCrestCopy.actionFailed;
    } finally {
      busy = null;
    }
  }

  function onRemoveStart(): void {
    confirmingRemove = true;
    actionError = '';
  }

  function onRemoveCancel(): void {
    confirmingRemove = false;
    actionError = '';
  }

  async function onRemoveConfirm(): Promise<void> {
    busy = 'removing';
    actionError = '';
    try {
      await deleteGuildCrest(guildId);
      confirmingRemove = false;
      onCrestChanged();
    } catch (thrown) {
      actionError = thrown instanceof GuildApiError ? thrown.message : guildCrestCopy.actionFailed;
    } finally {
      busy = null;
    }
  }

  // Tests only: jsdom never actually decodes the object URL, so this never leaks there;
  // in a real browser, leaving the Settings tab without saving/cancelling would otherwise
  // hold the blob alive for the rest of the session.
  $effect(() => {
    return () => {
      if (previewUrl !== null) URL.revokeObjectURL(previewUrl);
    };
  });
</script>

<section data-testid="guild-crest-settings">
  <div class="{ROW} gap-y-2">
    {#if previewUrl !== null}
      <img
        src={previewUrl}
        alt=""
        aria-hidden="true"
        width="40"
        height="40"
        class="bg-raised box-border h-10 w-10 shrink-0 rounded-full object-contain p-[6px]"
        style={ringStyle}
        data-testid="guild-crest-preview"
      />
    {:else if currentMark !== null}
      <img
        src={currentMark}
        alt=""
        aria-hidden="true"
        width="40"
        height="40"
        class="bg-raised box-border h-10 w-10 shrink-0 rounded-full object-contain p-[6px]"
        style={ringStyle}
        data-testid="guild-crest-current"
      />
    {:else}
      <div
        class="bg-raised box-border h-10 w-10 shrink-0 rounded-full"
        aria-hidden="true"
        data-testid="guild-crest-current"
      ></div>
    {/if}
    <span class="flex min-w-0 flex-1 flex-col">
      <span class="{ROW_LABEL} text-strong">{guildCrestCopy.heading}</span>
      {#if !hasCrest}
        <span class={ROW_SUB} data-testid="guild-crest-default-caption">
          {guildCrestCopy.defaultCaption}
        </span>
      {/if}
    </span>
    <!-- The browser's own file control is hidden (it stays in the page for the file dialog, a keyboard and
         a screen reader); the secondary button is its label, and the focus ring shows on the label. -->
    <label
      class="{button(
        'secondary',
        'sm',
      )} has-[:focus-visible]:outline-gold cursor-pointer has-[:focus-visible]:outline-2 has-[:focus-visible]:outline-offset-2 {busy !==
      null
        ? BUSY_CLASS
        : ''}"
      data-testid="guild-crest-choose"
    >
      {guildCrestCopy.chooseButton}
      <input
        id="guild-crest-file"
        type="file"
        accept={CREST_ACCEPT_ATTR}
        bind:this={fileInputEl}
        onchange={onFileChange}
        disabled={busy !== null}
        class="sr-only"
        aria-label={guildCrestCopy.pickerLabel}
        data-testid="guild-crest-file-input"
      />
    </label>
  </div>

  <div class="flex flex-col gap-3 px-5 pb-4">
    <div class="flex flex-col gap-1">
      <span class="text-muted min-w-0 truncate text-[13px]" data-testid="guild-crest-file-name"
        >{selectedFile === null ? guildCrestCopy.noneChosen : selectedFile.name}</span
      >
      <p class="text-muted text-[12px]">{CREST_RULES_LINE}</p>
    </div>

    {#if fileError !== ''}
      <p class="text-[13px]" role="alert" data-testid="guild-crest-file-error">{fileError}</p>
    {:else if actionError !== ''}
      <p class="text-[13px]" role="alert" data-testid="guild-crest-action-error">{actionError}</p>
    {/if}

    <div class="flex flex-wrap items-center gap-3">
      <button
        type="button"
        class="{button('secondary', 'sm')} {busy === 'saving' ? BUSY_CLASS : ''}"
        onclick={() => void onSave()}
        disabled={selectedFile === null || busy !== null}
        aria-busy={busy === 'saving'}
        data-testid="guild-crest-save"
      >
        {busy === 'saving' ? guildCrestCopy.saving : guildCrestCopy.save}
      </button>

      {#if hasCrest && !confirmingRemove}
        <button
          type="button"
          class={button('ghost', 'sm')}
          onclick={onRemoveStart}
          disabled={busy !== null}
          data-testid="guild-crest-remove"
        >
          {guildCrestCopy.remove}
        </button>
      {/if}
    </div>

    {#if confirmingRemove}
      <div class="border-line-soft flex flex-col gap-3 border p-4" data-testid="guild-crest-remove-confirm">
        <p class="text-[13px]">{guildCrestCopy.removeConfirmLine}</p>
        <div class="flex gap-3">
          <button
            type="button"
            class="{button('danger', 'sm')} {busy === 'removing' ? BUSY_CLASS : ''}"
            onclick={() => void onRemoveConfirm()}
            disabled={busy !== null}
            aria-busy={busy === 'removing'}
            data-testid="guild-crest-remove-confirm-button"
          >
            {busy === 'removing' ? guildCrestCopy.removing : guildCrestCopy.removeConfirmButton}
          </button>
          <button
            type="button"
            class={button('ghost', 'sm')}
            onclick={onRemoveCancel}
            disabled={busy !== null}
          >
            {guildHomeCopy.cancel}
          </button>
        </div>
      </div>
    {/if}
  </div>
</section>
