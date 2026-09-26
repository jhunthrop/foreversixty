<!-- web/src/components/CurrentCharacterBar.svelte -->
<!-- Two things live here. The default mode (compact/hasOwnPasteBox) is the current-character
     chip for a page that has no character state of its own (account, character, setup) --
     unchanged from before this lane. `spine: true` is the new band spec 2026-09-25 section
     4.1 describes: one row under the header, the resolved character's identity, four door
     links, and a Switch control, mounted on every tool and guide page (Tasks 5-8). The two
     modes share nothing but the pointer read and CHIP_HEIGHT -- a page never sets both. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import { fetchMeOnce, type Me, type MeCharacter } from '../lib/account/api';
  import { createQueryState } from '../lib/data/query.svelte';
  import { API_BASE_URL } from '../lib/planner/config';
  import {
    CURRENT_CHARACTER_CHANGED,
    clearCurrent,
    readCurrent,
    writeCurrent,
    type CurrentCharacter,
  } from '../lib/current-character';
  import { resolveSpineClassSlug, spineDoorsFor } from '../lib/current-character-bar';
  import { CHIP_HEIGHT } from '../lib/current-character-layout';
  import { currentCharacterCopy } from '../lib/current-character-copy';
  import { guildRankLabel } from '../lib/characters';
  import { heroCharacter } from '../lib/account/hero-character';
  import { mainCharacter, pointerForCharacter } from '../lib/account/main-character';
  import { classColorVar, rowLink } from '../lib/report/format';
  import CharacterIdentity from './character/CharacterIdentity.svelte';
  import CharacterSwitchList from './character/CharacterSwitchList.svelte';
  import CurrentCharacterChip from './CurrentCharacterChip.svelte';

  let {
    hasOwnPasteBox = false,
    compact = false,
    spine = false,
  }: { hasOwnPasteBox?: boolean; compact?: boolean; spine?: boolean } = $props();

  let current = $state<CurrentCharacter | null>(null);
  let guildLine = $state('');

  $effect(() => {
    const read = (): void => {
      const next = readCurrent();
      current = next;
      guildLine = '';
      if (next === null || next.source !== 'armory') return;
      const key = next.ref;
      void fetchMeOnce()
        .then((me) => {
          if (untrack(() => current)?.ref !== key) return;
          const character = me?.characters.find((c) => c.key === key);
          if (character?.guild === undefined) return;
          const rank = character.guild.rank === undefined ? '' : ` · ${guildRankLabel(character.guild.rank)}`;
          guildLine = `${character.guild.name}${rank}`;
        })
        .catch(() => {
          guildLine = '';
        });
    };
    read();
    window.addEventListener(CURRENT_CHARACTER_CHANGED, read);
    return () => window.removeEventListener(CURRENT_CHARACTER_CHANGED, read);
  });

  function forget(): void {
    clearCurrent();
    current = null;
    guildLine = '';
  }

  // Spine-only: the session, read the same way home-hero.svelte.ts and AccountMenu.svelte
  // already do -- one cached /v1/me shared across every island on the page.
  const session = createQueryState<Me | null>(`${API_BASE_URL}/v1/me`, () => fetchMeOnce(), {
    scope: 'private',
    ttlMs: 10 * 60 * 1000,
  });
  const me = $derived(spine ? session.data : null);
  const main = $derived(me !== null ? mainCharacter(me.characters, me.main_character_key) : null);
  const pointerCharacter = $derived(me !== null ? heroCharacter(current, me.characters) : null);
  // Pointer wins; the main is the fallback only once there is no pointer at all (spec 4.1).
  const displayCharacter = $derived<MeCharacter | null>(pointerCharacter ?? (current === null ? main : null));
  const classSlug = $derived(resolveSpineClassSlug(pointerCharacter, current, main));
  const doors = $derived(
    spineDoorsFor(classSlug, current, displayCharacter?.region ?? null, displayCharacter?.ruleset ?? null),
  );

  let switchOpen = $state(false);
  let switchRoot: HTMLElement | undefined = $state();

  function onSwitch(character: MeCharacter): void {
    writeCurrent(pointerForCharacter(character));
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    switchOpen = false;
  }

  function onDocumentPointerDown(event: PointerEvent): void {
    if (!switchOpen || !(event.target instanceof Node) || switchRoot?.contains(event.target)) return;
    switchOpen = false;
  }

  function onDocumentKeydown(event: KeyboardEvent): void {
    if (switchOpen && event.key === 'Escape') switchOpen = false;
  }

  $effect(() => {
    if (!spine) return;
    document.addEventListener('pointerdown', onDocumentPointerDown);
    document.addEventListener('keydown', onDocumentKeydown);
    return () => {
      document.removeEventListener('pointerdown', onDocumentPointerDown);
      document.removeEventListener('keydown', onDocumentKeydown);
    };
  });
</script>

{#if spine}
  <div class={`chip-slot ${CHIP_HEIGHT}`} data-testid="current-character-bar">
    {#if classSlug === null}
      <p class="text-muted mx-[18px] flex h-full items-center gap-3 overflow-hidden text-[13px] md:mx-0">
        <span data-testid="current-character-bar-signed-out">{currentCharacterCopy.barSignedOutLine}</span>
        <a class="text-nav underline" href="/login">{currentCharacterCopy.barSignIn}</a>
        <a class="text-nav underline" href="/setup#paste">{currentCharacterCopy.barPasteExport}</a>
      </p>
    {:else}
      <div class="flex h-full flex-col md:flex-row md:items-center md:gap-4" bind:this={switchRoot}>
        <div class="flex h-11 shrink-0 items-center md:h-auto" data-testid="current-character-bar-identity">
          {#if displayCharacter !== null}
            <CharacterIdentity character={displayCharacter} size="md" descriptor="realm" />
          {:else if current !== null}
            <span class="truncate font-semibold" style:color={classColorVar(current.classSlug)}>
              {current.label}
            </span>
          {/if}
        </div>
        <div
          class="flex h-11 flex-nowrap items-center gap-3 overflow-x-auto whitespace-nowrap md:h-auto md:flex-1"
        >
          {#each doors as door (door.id)}
            <a class="{rowLink} text-nav min-w-11 justify-center" href={door.href} data-testid={door.testid}
              >{door.label}</a
            >
          {/each}
          {#if me !== null && me.characters.length > 0}
            <div class="relative">
              <button
                type="button"
                class="{rowLink} text-nav min-w-11 justify-center"
                data-testid="current-character-bar-switch"
                onclick={() => (switchOpen = !switchOpen)}
              >
                {currentCharacterCopy.barSwitch}
              </button>
              {#if switchOpen}
                <div
                  class="bg-raised border-line rounded-panel absolute top-full left-0 z-30 mt-2 w-[240px] border py-2"
                >
                  <CharacterSwitchList
                    characters={me.characters}
                    currentKey={current?.source === 'armory' ? current.ref : (main?.key ?? null)}
                    onswitch={onSwitch}
                  />
                </div>
              {/if}
            </div>
          {/if}
        </div>
      </div>
    {/if}
  </div>
{:else}
  <CurrentCharacterChip {current} hasOwnPasteBox={hasOwnPasteBox || compact} {guildLine} onforget={forget} />
{/if}
