<!-- web/src/components/planner/PlannerCharacterCard.svelte -->
<!-- Rebuild spec §4.B: the Character card paired with the header, signed in only. Reuses
     the BiS rebuild's own data hook (`createCharacterCardState`, `lib/bis/character-card-
     state.svelte.ts`) rather than a second `/v1/me` read, and the same inline crest/faction
     asset recipe `CharacterCard.svelte` already uses in place of importing `ClassCrest.
     astro`/`FactionMark.astro` (a framework component cannot import an `.astro` file at
     all -- see that component's own header comment).

     Deliberately its own small component rather than `CharacterCard.svelte` reused as-is:
     that component's copy and its synced/not-synced branching ("install the addon" vs.
     "send list to the addon") are the BiS page's own job, answering "do you know what I
     wear." The planner never asks that -- "Send this build to the addon" sends the
     planner's own current build and stays live in every signed-in state, synced or not
     (spec §4.B) -- and it drops the BiS card's "Switch" aside too: the planner's own class/
     race selects already let a visitor plan for any character, so a second switch surface
     here would contradict the confirm-free reset `store.selectClass` already enforces. -->
<script lang="ts">
  import { createCharacterCardState } from '../../lib/bis/character-card-state.svelte';
  import { bisCopy } from '../../lib/bis/copy';
  import { classCrestSrc } from '../../lib/class-crest';
  import { factionMarkSrc } from '../../lib/faction-mark';
  import { classColorVar } from '../../lib/report/format';
  import { classSlugFromName } from '../../lib/report/tree-sizes';
  import { relativeTime } from '../../lib/sim/sources';
  import { plannerCharacterCardCopy } from '../../lib/planner/copy';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import Skeleton from '../ui/Skeleton.svelte';
  import LoadError from '../ui/LoadError.svelte';

  let { addonHref = '/addon' }: { addonHref?: string } = $props();

  const state = createCharacterCardState();
  const character = $derived(state.character);
  const synced = $derived(character?.build !== undefined);
  const syncedRelative = $derived(
    character?.build === undefined ? '' : relativeTime(character.build.captured_at),
  );
</script>

{#if state.status === 'loading'}
  <div
    class="panel-box character-card w-full lg:w-[440px] lg:shrink-0"
    data-testid="planner-character-card-loading"
  >
    <Skeleton lines={3} testid="planner-character-card-skeleton" />
    <Skeleton lines={1} rowHeight="h-10" />
  </div>
{:else if state.status === 'failed'}
  <div
    class="panel-box character-card w-full lg:w-[440px] lg:shrink-0"
    data-testid="planner-character-card-error"
  >
    <LoadError message={plannerCharacterCardCopy.loadError} onRetry={() => state.refresh()} />
  </div>
{:else if character !== null}
  {@const color = classColorVar(character.class)}
  {@const faction = character.faction}
  <div class="panel-box character-card w-full lg:w-[440px] lg:shrink-0" data-testid="planner-character-card">
    <span class="label character-card-label">{plannerCharacterCardCopy.label}</span>
    <div class="character-card-identity">
      <img
        src={classCrestSrc(classSlugFromName(character.class ?? ''))}
        alt=""
        width="40"
        height="40"
        class="character-card-crest"
        style={`--c:${color}`}
      />
      <span class="character-card-identity-text">
        <span class="font-display character-card-name" style={`color:${color}`}>{character.name}</span>
        <span class="text-muted character-card-descriptor">
          {#if faction !== undefined}
            <img src={factionMarkSrc(faction)} alt="" width="16" height="16" class="character-card-faction" />
          {/if}
          {bisCopy.characterCardIdentityLine(character.level ?? 0, character.race ?? '')}
          ·
          {synced ? bisCopy.syncedRelative(syncedRelative) : bisCopy.notSyncedYetLabel}
        </span>
      </span>
    </div>
    <a
      class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text justify-center px-4`}
      href={addonHref}
      data-testid="planner-character-card-send"
    >
      {plannerCharacterCardCopy.sendBuildToAddon}
    </a>
  </div>
{/if}

<style>
  .panel-box {
    background: color-mix(in srgb, var(--color-raised) 85%, transparent);
    border: 1px solid var(--color-line);
    border-radius: var(--radius-panel);
  }
  .character-card {
    display: flex;
    flex-direction: column;
    gap: 12px;
    padding: 16px 20px;
    min-height: 140px;
  }
  .character-card-label {
    display: flex;
    align-items: center;
  }
  .character-card-identity {
    display: flex;
    align-items: center;
    gap: 12px;
  }
  .character-card-crest {
    flex-shrink: 0;
    border-radius: 999px;
    box-shadow: 0 0 0 2px var(--c);
    background: var(--color-raised);
    object-fit: cover;
  }
  .character-card-identity-text {
    display: flex;
    flex-direction: column;
    min-width: 0;
  }
  .character-card-name {
    font-size: 15px;
    font-weight: 700;
  }
  .character-card-descriptor {
    display: flex;
    align-items: center;
    gap: 4px;
    font-size: 12px;
  }
  .character-card-faction {
    object-fit: contain;
  }
</style>
