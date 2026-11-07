<!-- web/src/components/AddonPasteBox.svelte -->
<!-- /setup's paste box (setup spec 2026-11-07 section 4, rail "Try an export"): prove a pasted
     export decodes with the same decoder the planner uses (planner/fs1.ts, read-only import) and
     hand back the two places it goes next. Never builds a talent draft: it only proves the string
     is readable and links on with the raw code. Drawn in the flat system: a FormRow with the mono
     export Field and a secondary Load, the decoder's own refusal under it (role="alert"), then the
     two links and the save block (AddonPasteSave) once an export loads. -->
<script lang="ts">
  import { addonCopy } from '../lib/addon/copy';
  import { ADDON_PASTE_STATUS_MIN_H } from '../lib/addon/paste-layout';
  import { decodeFS1, type FS1Build } from '../lib/planner/fs1';
  import { recordPastedExport } from '../lib/character-selector/paste';
  import { plannerCodeHref, simExportHref } from '../lib/handoff-links';
  import { LINK } from '../lib/ui/primitives';
  import AddonPasteSave from './AddonPasteSave.svelte';
  import Button from './ui/Button.svelte';
  import Field from './ui/Field.svelte';
  import FormRow from './ui/FormRow.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  const FIELD_ROWS = 2;

  /** Whether the visitor is signed in; null while the page's session read is in flight. */
  let { signedIn }: { signedIn: boolean | null } = $props();

  let code = $state('');
  let error = $state<string | null>(null);
  let loaded = $state<string | null>(null);
  // The `who=` section of the loaded code, when it carries one: the save form starts from it.
  let loadedCharacter = $state<FS1Build['character']>(undefined);

  function load(event: SubmitEvent): void {
    event.preventDefault();
    const trimmed = code.trim();
    if (trimmed === '') return;
    const result = decodeFS1(trimmed);
    if (!result.ok) {
      error = result.message;
      loaded = null;
      return;
    }
    error = null;
    loaded = trimmed;
    loadedCharacter = result.build.character;
    // An export that names a character is the player choosing it: it becomes the current
    // character (the account's own, when it holds one by that name). One that names nobody
    // is looked at and changes nothing -- a class word is never made the pointer.
    void recordPastedExport(trimmed);
  }
</script>

<div class="addon-paste flex flex-col gap-3" data-testid="addon-paste-box">
  <FormRow label={addonCopy.setupTryExport} onsubmit={load}>
    <Field
      id="addon-paste-code"
      label={addonCopy.pasteFieldLabel}
      placeholder={addonCopy.pastePlaceholder}
      rows={FIELD_ROWS}
      grow
      bind:value={code}
      testid="addon-paste-code"
    />
    {#snippet submit()}
      <Button type="submit" disabled={code.trim() === ''} testid="addon-paste-submit"
        >{addonCopy.pasteAction}</Button
      >
    {/snippet}
  </FormRow>
  {#if error !== null}
    <p class="text-strong text-[14px] leading-5" data-testid="addon-paste-error" role="alert">{error}</p>
  {/if}
  {#if loaded !== null}
    <div class="reveal flex flex-col gap-3">
      <div class="flex flex-wrap gap-x-4 gap-y-1">
        <a
          class="{LINK} inline-flex min-h-11 items-center md:min-h-0"
          href={plannerCodeHref(loaded)}
          data-testid="addon-paste-planner">{addonCopy.pasteOpenPlanner} ›</a
        >
        <a
          class="{LINK} inline-flex min-h-11 items-center md:min-h-0"
          href={simExportHref(loaded)}
          data-testid="addon-paste-sim">{addonCopy.pasteOpenSim} ›</a
        >
      </div>
      {#if signedIn === null}
        <!-- One row: ADDON_PASTE_STATUS_MIN_H reserves the shorter (sign-in hint) branch's height. -->
        <Skeleton lines={1} minHeight={ADDON_PASTE_STATUS_MIN_H} testid="addon-paste-status-skeleton" />
      {:else}
        {#key loaded}
          <AddonPasteSave {signedIn} code={loaded} character={loadedCharacter} />
        {/key}
      {/if}
    </div>
  {/if}
</div>

<style>
  /* The export is literal text a player pastes: the code face (polish spec 4.4). */
  .addon-paste :global(textarea) {
    min-height: 72px;
    font: 600 14px/20px var(--font-mono);
  }
</style>
