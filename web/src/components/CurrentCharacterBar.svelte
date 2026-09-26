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
    currentDoor = null,
    restored = false,
  }: {
    hasOwnPasteBox?: boolean;
    compact?: boolean;
    spine?: boolean;
    /** The door this page is: drawn in gold and marked aria-current, never linked to itself. */
    currentDoor?: 'plan' | 'sim' | 'logs' | 'rankings' | null;
    /** The mounting page restored the pointer from a previous visit (the chip's old note). */
    restored?: boolean;
  } = $props();

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
  {#if classSlug === null}
    <!-- Not wrapped in `.chip-slot`: that class is hidden pre-paint (global.css) whenever
         neither `data-pointer` nor `data-session` is set on `<html>`, which is exactly the
         one signal-free case this branch is FOR -- a visitor with truly nothing to resolve.
         The hiding rule exists to stop a *different* guess (this component's own SSR output,
         made with no client data at all) from flashing before hydration corrects it; for a
         genuinely signed-out, no-pointer visitor there is nothing to correct after hydration
         -- SSR and the hydrated client render the identical line -- so hiding it here would
         make spec 2026-09-25 4.1's signed-out line permanently unreachable for exactly the
         visitor it is written for. `CHIP_HEIGHT` alone still reserves the identical height. -->
    <div class={CHIP_HEIGHT} data-testid="current-character-bar">
      <!-- Sentence above, links below on a phone: as one row the two links were squeezed into
           a side column and wrapped a word per line. One row again from md. -->
      <div
        class="text-muted mx-[18px] flex h-full flex-col justify-center gap-1 overflow-hidden text-[13px] md:mx-0 md:flex-row md:items-center md:gap-3"
      >
        <span data-testid="current-character-bar-signed-out">{currentCharacterCopy.barSignedOutLine}</span>
        <span class="flex min-h-11 items-center gap-4 md:min-h-0 md:gap-3">
          <a class="text-nav inline-flex min-h-11 items-center underline md:min-h-0" href="/login"
            >{currentCharacterCopy.barSignIn}</a
          >
          <a class="text-nav inline-flex min-h-11 items-center underline md:min-h-0" href="/setup#paste"
            >{currentCharacterCopy.barPasteExport}</a
          >
        </span>
      </div>
    </div>
  {:else}
    <!-- This branch alone stays behind `.chip-slot`: it depends on client-only data (the
         stored pointer, the session) the pre-paint script has already hinted at via
         data-pointer/data-session when either is real, so hiding it until hydration
         resolves is what stops it from ever painting a wrong guess. -->
    <div class={`chip-slot ${CHIP_HEIGHT}`} data-testid="current-character-bar">
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
        <!-- Every child is shrink-0: a flex item with an explicit min-width (min-w-11) no longer
             keeps its min-content width, so without it the phone row shrank the links onto
             one another. The row wraps rather than scrolls: a scrolled row hid Forget and
             Switch past the edge with nothing to say so. Six short labels fit one 390px line;
             the restored note is desktop-only so it cannot push the row to a second line. -->
        <div class="flex h-11 flex-wrap items-center gap-x-3 whitespace-nowrap md:h-auto md:flex-1">
          {#each doors as door (door.id)}
            <a
              class="{rowLink} label min-w-11 shrink-0 justify-center {door.id === currentDoor
                ? 'text-gold'
                : 'text-nav'}"
              href={door.href}
              aria-current={door.id === currentDoor ? 'page' : undefined}
              data-testid={door.testid}
              >{#if door.shortLabel !== undefined}<span class="md:hidden">{door.shortLabel}</span><span
                  class="hidden md:inline">{door.label}</span
                >{:else}{door.label}{/if}</a
            >
          {/each}
          {#if current !== null}
            {#if restored}
              <span
                class="text-muted hidden shrink-0 text-[12px] md:inline"
                data-testid="current-character-restored"
              >
                {currentCharacterCopy.restoredNote}
              </span>
            {/if}
            <button
              type="button"
              class="{rowLink} text-muted label min-w-11 shrink-0 justify-center"
              onclick={forget}
              data-testid="current-character-forget"
            >
              {currentCharacterCopy.forget}
            </button>
          {/if}
          {#if me !== null && me.characters.length > 0}
            <div class="relative shrink-0">
              <button
                type="button"
                class="{rowLink} label text-nav min-w-11 justify-center"
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
    </div>
  {/if}
{:else}
  <CurrentCharacterChip {current} hasOwnPasteBox={hasOwnPasteBox || compact} {guildLine} onforget={forget} />
{/if}
