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
  import { missingPasteSaveField, type PasteSaveField } from '../lib/addon/paste-save';
  import { REGIONS, RULESETS, rulesetFromRealm } from '../lib/characters';
  import type { FS1Build } from '../lib/planner/fs1';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { BUSY_CLASS } from '../lib/ui/busy';
  import Select from './ui/Select.svelte';

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
  // The field the last press found empty, so it can be marked and focused.
  let missing = $state<PasteSaveField | null>(null);
  /** Each field's id, so the one a press found empty takes focus. */
  const FIELD_IDS: Record<PasteSaveField, string> = {
    name: 'addon-paste-name',
    region: 'addon-paste-region-field',
    ruleset: 'addon-paste-ruleset-field',
  };

  const NEEDS: Record<PasteSaveField, string> = {
    name: addonCopy.pasteNeedsName,
    region: addonCopy.pasteNeedsRegion,
    ruleset: addonCopy.pasteNeedsRuleset,
  };

  async function onSave(): Promise<void> {
    missing = missingPasteSaveField({ name, region, ruleset });
    if (missing !== null) {
      error = NEEDS[missing];
      document.getElementById(FIELD_IDS[missing])?.focus();
      return;
    }
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
    <label class="quiet-label" for="addon-paste-name">{addonCopy.pasteNameLabel}</label>
    <input
      id={FIELD_IDS.name}
      class="border-line-warm bg-raised rounded-control text-text h-11 border px-3 text-[14px]"
      bind:value={name}
      aria-invalid={missing === 'name' ? 'true' : undefined}
      disabled={saved}
      data-testid="addon-paste-name"
    />
    <div class="flex flex-wrap gap-3">
      <Select
        id={FIELD_IDS.region}
        class="w-[140px]"
        size="compact"
        label={addonCopy.pasteRegionLabel}
        bind:value={region}
        placeholder={addonCopy.pasteRegionLabel}
        options={REGIONS.map((r) => ({ value: r, label: r.toUpperCase() }))}
        invalid={missing === 'region'}
        disabled={saved}
        testid="addon-paste-region"
      />
      <Select
        id={FIELD_IDS.ruleset}
        class="w-[180px]"
        size="compact"
        label={addonCopy.pasteRulesetLabel}
        bind:value={ruleset}
        placeholder={addonCopy.pasteRulesetLabel}
        options={RULESETS.map((r) => ({ value: r.id, label: r.label }))}
        invalid={missing === 'ruleset'}
        disabled={saved}
        testid="addon-paste-ruleset"
      />
    </div>
    <button
      type="button"
      class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4 ${busy ? BUSY_CLASS : ''}`}
      disabled={busy || saved}
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
