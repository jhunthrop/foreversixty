<!-- web/src/components/CurrentCharacterBar.svelte -->
<!-- Two things live here. The default mode (compact/hasOwnPasteBox) is the current-character
     chip for a page that has no character state of its own (account, character, setup) --
     unchanged from before this lane. `spine: true` is the new band spec 2026-09-25 section
     4.1 describes: one row under the header, the resolved character's identity, four door
     links and Forget, mounted on every tool and guide page (Tasks 5-8); switching character
     is the header selector's job alone (nav spec 2026-10-09). The two modes share nothing but the pointer read and CHIP_HEIGHT -- a page never sets both. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import { fetchMeOnce, type Me, type MeCharacter } from '../lib/account/api';
  import { createQueryState } from '../lib/data/query.svelte';
  import { API_BASE_URL } from '../lib/planner/config';
  import {
    clearCurrent,
    onCurrentCharacterChange,
    readCurrent,
    type CurrentCharacter,
  } from '../lib/current-character';
  import { resolveSpineClassSlug, spineDoorsFor } from '../lib/current-character-bar';
  import { CHIP_HEIGHT, STRIP_HEIGHT } from '../lib/current-character-layout';
  import { currentCharacterCopy } from '../lib/current-character-copy';
  import { guildRankLabel } from '../lib/characters';
  import { heroCharacter } from '../lib/account/hero-character';
  import { mainCharacter } from '../lib/account/main-character';
  import { classColorVar } from '../lib/report/format';
  import { cardClass } from '../lib/ui/card';
  import { ROW, ROW_LABEL } from '../lib/ui/row';
  import IconTile from './ui/IconTile.svelte';
  import { button } from '../lib/ui/button';
  import CharacterIdentity from './character/CharacterIdentity.svelte';
  import CurrentCharacterChip from './CurrentCharacterChip.svelte';

  let {
    hasOwnPasteBox = false,
    compact = false,
    spine = false,
    currentDoor = null,
    restored = false,
    profile = false,
    quietSignedOut = false,
    cardName = '',
    inCard = false,
  }: {
    /** The row sits inside another card (the sim landing's Your characters) instead of in a card of its own. */
    inCard?: boolean;
    /** The name of the character the page's own card shows. When it is the one this strip speaks
     *  for, the strip shows only its doors: the card carries the identity, said once. */
    cardName?: string;
    /** The signed-out sentence and its two links are left out: the page already asks one thing
     *  (an empty character page, W.10). */
    quietSignedOut?: boolean;
    hasOwnPasteBox?: boolean;
    compact?: boolean;
    spine?: boolean;
    /** The door this page is: drawn in gold and marked aria-current, never linked to itself. */
    currentDoor?: 'plan' | 'sim' | 'logs' | 'rankings' | null;
    /** The mounting page restored the pointer from a previous visit (the chip's old note). */
    restored?: boolean;
    /** The character page (Forever Score fix round 3 item 4, L8.4 fix round 2 item 7): the
     *  signed-out line sits below the score card at every width (one row of its two links on a
     *  phone), while a signed-in chip stays on top. Character.svelte orders the header and the
     *  score card around it. */
    profile?: boolean;
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
    return onCurrentCharacterChange(read);
  });

  function forget(): void {
    clearCurrent();
    current = null;
    guildLine = '';
  }

  // Spine-only: the session, read the same way home-hero.svelte.ts and the header selector
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
  /** The page's card already names the character this strip speaks for. */
  const identityOnCard = $derived(
    cardName !== '' && displayCharacter !== null && displayCharacter.name === cardName,
  );
  const doors = $derived(
    spineDoorsFor(classSlug, current, displayCharacter?.region ?? null, displayCharacter?.ruleset ?? null),
  );
</script>

{#if spine}
  {#if classSlug === null && quietSignedOut}
    <!-- Nothing: the page's one prompt stands alone. -->
  {:else if classSlug === null}
    <!-- Not wrapped in `.chip-slot`: that class is hidden pre-paint (global.css) whenever
         neither `data-pointer` nor `data-session` is set on `<html>`, which is exactly the
         one signal-free case this branch is FOR -- a visitor with truly nothing to resolve.
         The hiding rule exists to stop a *different* guess (this component's own SSR output,
         made with no client data at all) from flashing before hydration corrects it; for a
         genuinely signed-out, no-pointer visitor there is nothing to correct after hydration
         -- SSR and the hydrated client render the identical line -- so hiding it here would
         make spec 2026-09-25 4.1's signed-out line permanently unreachable for exactly the
         visitor it is written for. `CHIP_HEIGHT` alone still reserves the identical height. -->
    <div class={profile ? '-order-1 h-11' : CHIP_HEIGHT} data-testid="current-character-bar">
      <!-- Sentence above, links below on a phone: as one row the two links were squeezed into
           a side column and wrapped a word per line. One row again from md. -->
      <div
        class="text-muted mx-[18px] flex h-full flex-col justify-center gap-1 overflow-hidden text-[13px] md:mx-0 md:flex-row md:items-center md:justify-start md:gap-3"
      >
        <span class={profile ? 'max-md:sr-only' : undefined} data-testid="current-character-bar-signed-out"
          >{currentCharacterCopy.barSignedOutLine}</span
        >
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
    <div
      class="chip-slot {inCard ? '' : `${STRIP_HEIGHT} ${cardClass(false)} mx-[18px] md:mx-0`} {profile
        ? '-order-4'
        : ''}"
      data-testid="current-character-bar"
    >
      <!-- One account-style row (polish spec 4.2): the crest and name on the left, the doors and
           Forget as buttons on the right; the identity takes the whole first line on a phone and
           the doors wrap under it. In a card of its own the row drops its top rule. -->
      <div
        class="{ROW} {inCard ? '' : 'h-full'} gap-y-2 py-2"
        style:border-top-width={inCard ? undefined : '0'}
        data-testid="current-character-bar-row"
      >
        {#if !identityOnCard}
          <div class="min-w-0 basis-full md:flex-1 md:basis-0" data-testid="current-character-bar-identity">
            {#if displayCharacter !== null}
              <CharacterIdentity character={displayCharacter} size="row" descriptor="realm" />
            {:else if current !== null}
              <span class="truncate font-semibold" style:color={classColorVar(current.classSlug)}>
                {current.label}
              </span>
            {/if}
          </div>
        {:else}
          <div class="flex min-w-0 basis-full items-center gap-[14px] md:flex-1 md:basis-0">
            <IconTile glyph="user" size={32} />
            <span class="{ROW_LABEL} text-strong">{currentCharacterCopy.barDoorsLabel}</span>
          </div>
        {/if}
        <!-- Every child is shrink-0 and the row wraps rather than scrolls: a scrolled row hid
             Forget past the edge with nothing to say so. -->
        <div class="flex flex-wrap items-center gap-x-1 whitespace-nowrap md:ml-auto md:justify-end">
          {#each doors as door (door.id)}
            <a
              class="{button('ghost', 'sm')} min-w-11 shrink-0 {door.id === currentDoor
                ? 'underline underline-offset-4'
                : ''}"
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
              class="{button('ghost', 'sm')} shrink-0"
              onclick={forget}
              title={displayCharacter === null
                ? undefined
                : `Forget ${displayCharacter.name} and use your main`}
              data-testid="current-character-forget"
            >
              {currentCharacterCopy.forget}
            </button>
          {/if}
        </div>
      </div>
    </div>
  {/if}
{:else}
  <CurrentCharacterChip {current} hasOwnPasteBox={hasOwnPasteBox || compact} {guildLine} onforget={forget} />
{/if}
