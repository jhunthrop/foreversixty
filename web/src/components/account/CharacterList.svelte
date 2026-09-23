<!-- web/src/components/account/CharacterList.svelte -->
<!-- /account's Characters panel (brief 2026-09-22 §B3): the design system's Character row
     inside a StatePanel, in place of the plain text rows the previous lane shipped. A pure
     render of the two props Account.svelte already has in hand -- `me.characters` and
     `me.bnet_imported_at` -- it makes no fetch of its own, and neither does each row's
     `CharacterRowLink` (spec 2026-09-22 §3.4): unlike `CharacterHandoffLinks` (still used,
     one row at a time, by the hero band and the character page), a list of N characters
     must not cost N `sim-input` fetches just to draw its rows. -->
<script lang="ts">
  import type { MeCharacter } from '../../lib/account/api';
  import { battlenetStartUrl } from '../../lib/account/api';
  import { characterDescriptor, classSquare, classIconUrl } from '../../lib/account/character-descriptor';
  import { characterListCopy } from '../../lib/account/character-list-copy';
  import { characterHref, guildRankLabel } from '../../lib/characters';
  import { relativeTime } from '../../lib/dates';
  import { classColorVar } from '../../lib/report/format';
  import CharacterRowLink from './CharacterRowLink.svelte';
  import { buildSourcePill } from '../../lib/account/build-pill';
  import EmptyState from '../ui/EmptyState.svelte';
  import StatePanel from '../ui/StatePanel.svelte';

  let { characters, bnetImportedAt }: { characters: MeCharacter[]; bnetImportedAt?: string } = $props();

  /** `?next=` on the refresh link (spec §7.2/§7.3): a second Battle.net login re-runs the
   *  import and lands back here with the toast query param. */
  const REFRESH_HREF = battlenetStartUrl('/account?refreshed=1');
</script>

{#snippet importedAside()}
  {#if bnetImportedAt !== undefined}
    <span data-testid="bnet-imported">
      {characterListCopy.importedFrom(relativeTime(new Date(bnetImportedAt)))}
      <a href={REFRESH_HREF}>{characterListCopy.refreshFromBattlenet}</a>
    </span>
  {/if}
{/snippet}

<StatePanel
  label={characterListCopy.heading}
  aside={bnetImportedAt === undefined ? undefined : importedAside}
  testid="account-characters"
>
  {#if characters.length === 0}
    <EmptyState
      message={characterListCopy.empty}
      action={{ label: characterListCopy.refreshFromBattlenet, href: REFRESH_HREF }}
      testid="account-characters-empty"
    />
  {:else}
    <ul class="flex flex-col">
      {#each characters as character (character.key)}
        {@const square = classSquare(character)}
        {@const classIcon = classIconUrl(character)}
        {@const pill = buildSourcePill(character.build)}
        <li class="border-line-soft flex flex-wrap items-center gap-3 border-b py-3 text-[14px]">
          {#if character.avatar_url !== undefined}
            <img
              class="h-9 w-9 shrink-0 rounded-[3px] object-cover"
              src={character.avatar_url}
              alt=""
              loading="lazy"
              data-testid="character-avatar"
            />
          {:else}
            <!-- The class icon sits over the letter square, so a failed load (a blank,
                 transparent image) still shows the letter underneath. -->
            <span
              class="relative flex h-9 w-9 shrink-0 items-center justify-center rounded-[3px] text-[15px] font-bold"
              style={`background-color: color-mix(in srgb, ${square.color} 22%, transparent); color: ${square.color}`}
              data-testid="character-avatar-fallback"
            >
              {square.letter}
              {#if classIcon !== undefined}
                <img
                  class="absolute inset-0 h-9 w-9 rounded-[3px] object-cover"
                  src={classIcon}
                  alt=""
                  loading="lazy"
                  data-testid="character-class-icon"
                />
              {/if}
            </span>
          {/if}
          <div class="flex min-w-0 flex-1 flex-col gap-0.5">
            <a
              class="w-fit [font-family:var(--font-display)] text-[15px] font-semibold"
              style:color={classColorVar(character.class)}
              href={characterHref(character.region, character.ruleset, character.name)}
            >
              {character.name}
            </a>
            <span class="text-muted text-[13px]" data-testid="character-descriptor">
              {characterDescriptor(character)}
            </span>
            <span
              class={pill.pillClass === null ? 'text-muted text-[12px]' : `pill ${pill.pillClass} w-fit`}
              data-testid="character-build-pill"
            >
              {pill.label}
            </span>
            {#if character.guild !== undefined}
              <span class="text-muted text-[13px]" data-testid="character-guild-line">
                {character.guild.name}
                {#if character.guild.rank !== undefined}· {guildRankLabel(character.guild.rank)}{/if}
                {#if character.guild.verified}
                  <span class="text-strong" data-testid="character-guild-verified"
                    >{characterListCopy.verified}</span
                  >
                {/if}
              </span>
            {/if}
          </div>
          <CharacterRowLink {character} />
        </li>
      {/each}
    </ul>
  {/if}
  <p class="text-muted text-[13px]">
    Gear and talents come from Battle.net and refresh nightly.
    <a class="text-text underline" href="/addon">{characterListCopy.introInstallAddonLink}</a>
    to include bags and bank and to update right after a session;
    <a class="text-text underline" href="/addon#paste">{characterListCopy.introPasteLink}</a>
    for a character Battle.net has no data for.
  </p>
</StatePanel>
