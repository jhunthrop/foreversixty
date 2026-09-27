<!-- web/src/components/AddonPasteSave.svelte -->
<!-- The signed-in half of AddonPasteBox.svelte's successful decode (spec 2026-09-22 §7.4,
     plan Ruling 1): POST /v1/me/exports needs a character name, region and ruleset beside
     the export. The addon's `who=` section supplies the name (first and last: Forever
     characters have both) and the realm, which names the ruleset on every realm seen so
     far, so those two start filled and the player checks them; the region has no source in
     the string, so that one is asked. Split out of AddonPasteBox.svelte (rather than inlined there) so the
     signed-in and signed-out branches are each a pure render of an explicit `signedIn`
     prop and can be exercised with svelte/server's render() the way GuildJoin.svelte's own
     effect-driven fetchMeOnce() cannot be (its SSR test only ever sees the pre-effect
     loading state) -- AddonPasteBox.svelte owns the fetchMeOnce() call and passes the
     resolved boolean down once it is known. -->
<script lang="ts">
  import { ACCOUNT_FAILED, AccountError, postMyExports } from '../lib/account/api';
  import { addonCopy } from '../lib/addon/copy';
  import { REGIONS, RULESETS, rulesetFromRealm } from '../lib/characters';
  import type { FS1Build } from '../lib/planner/fs1';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { BUSY_CLASS } from '../lib/ui/busy';

  let { signedIn, code, character }: { signedIn: boolean; code: string; character?: FS1Build['character'] } =
    $props();

  // Initial values only, on purpose: AddonPasteBox.svelte keys this component on the code,
  // so a new paste mounts a fresh form rather than overwriting what the player has typed.
  // svelte-ignore state_referenced_locally
  let name = $state(character?.name ?? '');
  let region = $state('');
  // svelte-ignore state_referenced_locally
  let ruleset = $state<string>(character === undefined ? '' : rulesetFromRealm(character.realm));
  let busy = $state(false);
  let saved = $state(false);
  let error = $state('');

  const canSave = $derived(name.trim() !== '' && region !== '' && ruleset !== '');

  async function onSave(): Promise<void> {
    busy = true;
    error = '';
    try {
      await postMyExports([{ name: name.trim(), region, ruleset, export: code }]);
      saved = true;
    } catch (thrown) {
      error = thrown instanceof AccountError ? thrown.message : ACCOUNT_FAILED;
    } finally {
      busy = false;
    }
  }
</script>

{#if signedIn}
  <div class="flex flex-col gap-2" data-testid="addon-paste-save">
    <label class="label text-muted" for="addon-paste-name">{addonCopy.pasteNameLabel}</label>
    <input
      id="addon-paste-name"
      class="border-line-warm bg-raised rounded-control text-text h-11 border px-3 text-[14px]"
      bind:value={name}
      disabled={saved}
      data-testid="addon-paste-name"
    />
    <div class="flex flex-wrap gap-3">
      <select
        aria-label={addonCopy.pasteRegionLabel}
        class="border-line-warm bg-raised rounded-control text-text h-11 border px-3 text-[14px]"
        bind:value={region}
        disabled={saved}
        data-testid="addon-paste-region"
      >
        <option value="">{addonCopy.pasteRegionLabel}</option>
        {#each REGIONS as r (r)}
          <option value={r}>{r.toUpperCase()}</option>
        {/each}
      </select>
      <select
        aria-label={addonCopy.pasteRulesetLabel}
        class="border-line-warm bg-raised rounded-control text-text h-11 border px-3 text-[14px]"
        bind:value={ruleset}
        disabled={saved}
        data-testid="addon-paste-ruleset"
      >
        <option value="">{addonCopy.pasteRulesetLabel}</option>
        {#each RULESETS as r (r.id)}
          <option value={r.id}>{r.label}</option>
        {/each}
      </select>
    </div>
    <button
      type="button"
      class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4 ${busy ? BUSY_CLASS : ''}`}
      disabled={busy || saved || !canSave}
      aria-busy={busy}
      onclick={onSave}
      data-testid="addon-paste-save-button"
    >
      {saved ? addonCopy.pasteSaved : addonCopy.pasteSaveAction}
    </button>
    {#if error !== ''}
      <p class="text-strong text-[13px]" role="alert" data-testid="addon-paste-save-error">{error}</p>
    {/if}
  </div>
{:else}
  <p class="text-muted text-[13px]" data-testid="addon-paste-signin-hint">{addonCopy.pasteSignInHint}</p>
{/if}
