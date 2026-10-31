<!-- web/src/components/sim/SimRunBlock.svelte -->
<!-- 2026-09-26 layout pass, review round 1 Finding 1: the one-action row for whichever
     character the header selector calls current -- a single "Run sim" button when that
     character has a build, or the identical "Paste export" link the character list below
     offers when it does not. SimView.svelte resolves `character` with the header's own
     pointer/main precedence, and this block calls the same `hasBuild` check
     LandingState.svelte's row does, so the two can never show a different affordance for
     the same character. -->
<script lang="ts">
  import type { MeCharacter } from '../../lib/account/api';
  import { hasBuild } from '../../lib/account/build-pill';
  import type { CharacterPath } from '../../lib/characters';
  import { landingCopy } from '../../lib/sim/landing-copy';
  import { button } from '../../lib/ui/button';
  import { cardClass } from '../../lib/ui/card';
  import { buildSourcePill } from '../../lib/account/build-pill';
  import { characterIsSimulated, characterSpecName } from '../../lib/character-spec';
  import { classDisplayName } from '../../lib/sim/spec-label';
  import { classColorVar } from '../../lib/report/format';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';

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
  const simulated = $derived(characterIsSimulated(character));
  const classSlug = $derived(character.class?.toLowerCase() ?? '');
  /** "Gnome Holy Priest · 60": the same words and order the simulator's own card uses. */
  const descriptor = $derived(
    [
      [
        character.race,
        characterSpecName(character),
        classSlug === '' ? undefined : classDisplayName(classSlug),
      ]
        .filter((part): part is string => part !== undefined && part !== '')
        .join(' '),
      character.level === undefined ? '' : String(character.level),
    ]
      .filter((part) => part !== '')
      .join(' · '),
  );
  const pill = $derived(buildSourcePill(character.build));
</script>

<section class="{cardClass(false)} mx-[18px] md:mx-0" data-testid="sim-run-block">
  <div class="flex flex-wrap items-center justify-between gap-x-[14px] gap-y-3 px-5 py-4">
    <div class="flex min-w-0 flex-wrap items-center gap-3" data-testid="sim-run-block-identity">
      <ClassCrestRing characterClass={classSlug} size={40} testid="sim-run-block-class-icon" />
      <span class="flex min-w-0 flex-col gap-1">
        <span
          class="text-[17px] leading-6 font-semibold"
          style={`color: ${classColorVar(character.class)}`}
          data-testid="sim-run-block-name">{character.name}</span
        >
        <span class="text-muted text-[13px]" data-testid="sim-run-block-descriptor">{descriptor}</span>
      </span>
      {#if canRun && pill.pillClass !== null}
        <span class="pill {pill.pillClass}" data-testid="sim-run-block-pill">{pill.label}</span>
      {/if}
    </div>
    {#if canRun && !simulated}
      <span class="text-muted text-[13px]" data-testid="sim-run-block-not-simulated"
        >{landingCopy.notSimulatedYet}</span
      >
    {:else if canRun}
      <button
        type="button"
        class="{button('primary', 'md')} shrink-0"
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
          class="{button('ghost', 'sm')} shrink-0"
          href={landingCopy.pasteExportHref}
          data-testid="sim-run-block-paste"
        >
          {landingCopy.pasteExport}
        </a>
      </div>
    {/if}
  </div>
</section>
