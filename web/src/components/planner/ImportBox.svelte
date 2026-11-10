<!-- web/src/components/planner/ImportBox.svelte -->
<!-- Import from addon (spec 2026-11-08 section 4.9): the paste box for an addon export, open on
     every width. A CodeField (mono: a literal code), Import, then the outcome: an error in
     --wipe, a class switch muted (the planner moves to that class and imports there), the import
     notes muted (the order is always approximated, sometimes a talent could not be placed or the
     export is a build behind); then the way to the addon. -->
<script lang="ts">
  import { addonCopy } from '../../lib/addon/copy';
  import { importFromAddon, type ImportOutcome } from '../../lib/addon/import';
  import type { TalentIndex } from '../../lib/planner/rules';
  import type { Gear } from '../../lib/planner/types';
  import { LINK } from '../../lib/ui/primitives';
  import Button from '../ui/Button.svelte';
  import CodeField from '../ui/CodeField.svelte';
  import Section from '../ui/Section.svelte';
  import PlannerPart from './PlannerPart.svelte';

  let {
    talents,
    activeBuild,
    onimport,
    onwrongclass = undefined,
    panel = false,
    part = false,
  }: {
    /** The heading inside its panel (the standalone planner's rail and column). */
    panel?: boolean;
    /** A part of the planner's take-it-with-you plate (PlannerPart) instead of its own Section. */
    part?: boolean;
    talents: TalentIndex;
    activeBuild: string;
    /**
     * `pastedCode` (Task 10) is the export string as pasted, trimmed the same way `submit`
     * decodes it -- the current-character bridge's own 'addon' pointer is written from
     * this exact string, never a re-encoding of the parsed build, so a caller reading it
     * back later decodes byte-for-byte what the player pasted.
     */
    onimport: (
      build: { classSlug: string; raceSlug: string; order: number[]; gear: Gear; characterName?: string },
      pastedCode: string,
    ) => void;
    /**
     * An export for a class other than the loaded one. When the mounting page can switch
     * class it takes the code here and imports once that class's talents arrive, instead
     * of the box telling the player to switch by hand; without this callback the box shows
     * the wrong-class message as before (Top Gear's inline box has one class).
     */
    onwrongclass?: (classSlug: string, pastedCode: string) => void;
  } = $props();

  let code = $state('');
  let outcome = $state<ImportOutcome | null>(null);

  const notes = $derived(outcome !== null && outcome.ok ? outcome.notes : []);
  const failure = $derived(outcome !== null && !outcome.ok ? outcome.message : null);
  /** The class switch is news, not an error: it reads muted, the error in --wipe. */
  let switching = $state(false);

  function submit(): void {
    const trimmed = code.trim();
    const result = importFromAddon(trimmed, talents, activeBuild);
    switching = false;
    if (!result.ok && result.wrongClass !== undefined && onwrongclass !== undefined) {
      outcome = { ok: false, message: addonCopy.importSwitchingClass(result.wrongClass) };
      switching = true;
      onwrongclass(result.wrongClass, trimmed);
      return;
    }
    outcome = result;
    if (result.ok) {
      onimport(
        {
          classSlug: result.classSlug,
          raceSlug: result.raceSlug,
          order: result.order,
          gear: result.gear,
          ...(result.characterName === undefined ? {} : { characterName: result.characterName }),
        },
        trimmed,
      );
    }
  }
</script>

{#snippet body()}
  <div class="pl-sec">
    <CodeField
      bind:value={code}
      placeholder={addonCopy.importPlaceholder}
      label={addonCopy.importTitle}
      invalid={failure !== null && !switching}
      testid="import-code"
    />
    <div class="pl-acts">
      <Button disabled={code.trim() === ''} onclick={submit} testid="import-submit"
        >{addonCopy.importAction}</Button
      >
    </div>
    {#if failure !== null}
      <p
        class="pl-note {switching ? 'text-muted' : 'text-wipe'}"
        role={switching ? 'status' : 'alert'}
        data-testid="import-error"
      >
        {failure}
      </p>
    {/if}
    {#each notes as note (note)}
      <p class="pl-note text-muted" data-testid="import-note">{note}</p>
    {/each}
    <p class="text-small text-muted">
      {addonCopy.importNoExport}
      <a
        class="{LINK} [@media(pointer:coarse)]:inline-block [@media(pointer:coarse)]:min-h-11"
        href="/setup#addon"
        data-testid="import-get-addon">{addonCopy.importGetAddon} ›</a
      >
    </p>
  </div>
{/snippet}

{#if part}
  <PlannerPart title={addonCopy.importTitle} testid="import-box">{@render body()}</PlannerPart>
{:else}
  <Section title={addonCopy.importTitle} {panel} testid="import-box">{@render body()}</Section>
{/if}
