<!-- web/src/components/sim/SimRunBlock.svelte -->
<!-- 2026-09-26 layout pass, review round 1 Finding 1: the one-action row directly under the
     spine bar, for whichever character the spine bar itself calls current -- a single
     "Run sim" button when that character has a build, or the identical "Paste export" link
     the character list below offers when it does not. SimView.svelte resolves `character`
     with the exact same pointer/main precedence CurrentCharacterBar.svelte's own spine
     mode uses, and this block calls the same `hasBuild` check LandingState.svelte's row
     does, so the two can never show a different affordance for the same character. -->
<script lang="ts">
  import type { MeCharacter } from '../../lib/account/api';
  import { hasBuild } from '../../lib/account/build-pill';
  import type { CharacterPath } from '../../lib/characters';
  import { landingCopy } from '../../lib/sim/landing-copy';
  import { PRIMARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import CharacterIdentity from '../character/CharacterIdentity.svelte';

  let {
    character,
    path,
    busy,
    onrun,
  }: {
    character: MeCharacter;
    /** Null only when the character's own key fails to parse -- never in practice, since
     *  `MeCharacter.key` is always `<region>/<ruleset>/<slug>` -- and the button disables
     *  rather than picking a path it cannot build. */
    path: CharacterPath | null;
    busy: boolean;
    onrun: (path: CharacterPath) => void;
  } = $props();

  const canRun = $derived(hasBuild(character));
</script>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-wrap items-center justify-between gap-3 border p-4 md:mx-0"
  data-testid="sim-run-block"
>
  <CharacterIdentity {character} size="md" descriptor="realm" testid="sim-run-block-identity" />
  {#if canRun}
    <button
      type="button"
      class="{PRIMARY_BUTTON_FIXED} shrink-0 px-5 disabled:opacity-50"
      disabled={busy || path === null}
      aria-busy={busy}
      onclick={() => path !== null && onrun(path)}
      data-testid="sim-run-block-action"
    >
      {landingCopy.runSim}
    </button>
  {:else}
    <div class="flex flex-wrap items-center gap-3">
      <span class="text-muted text-[13px]" data-testid="sim-run-block-no-build">
        {landingCopy.runBlockNoBuild(character.name)}
      </span>
      <a
        class="text-nav label inline-flex min-h-11 shrink-0 items-center underline md:min-h-9"
        href={landingCopy.pasteExportHref}
        data-testid="sim-run-block-paste"
      >
        {landingCopy.pasteExport}
      </a>
    </div>
  {/if}
</section>
