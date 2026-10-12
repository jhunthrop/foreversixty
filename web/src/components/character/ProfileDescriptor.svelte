<!-- web/src/components/character/ProfileDescriptor.svelte -->
<!-- The character profile's lines under the name (Task L8.4 ruling 1, fix round 2 items 3 to 5),
     all on the name's left edge:
       1. "Holy Priest · 31/20/0 · <Olympus XXVII> · Classic Beta PvP": "Holy Priest" in --text,
          the rest muted; each piece only when known; the guild links to its page, with a 44px
          tap target on a phone (padding, taken back out of the layout by a negative margin).
       2. "Level 60 Night Elf", when known.
       The "HARDCORE US" chip ends the last of those lines, 8px after its text, and never wraps
       away from it.
       3. The owner's sync line, the Character header's own: "Addon · 4 minutes ago", and past
          the nav spec's stale limit "Log out in the game to sync" in ember. Plain data in:
          Character.svelte works out ownership and the sync (lib/character-header/sync.ts). -->
<script lang="ts">
  import {
    characterHeaderCopy,
    type ProfileHeader,
    type ProfileHeaderPiece,
  } from '../../lib/character-profile';
  import { STALE_SEPARATOR, syncText, type SyncView } from '../../lib/character-header/sync';

  let { header, sync = null }: { header: ProfileHeader; sync?: SyncView | null } = $props();

  /** Muted text and one border token, whichever line it ends (fix round 2 item 3). */
  const CHIP = 'pill text-muted border-line align-middle ml-2';
  /** In-text links are underlined: colour alone does not mark a link (WCAG 1.4.1). On a phone
   *  11px of padding above and below the 22px line make a 44px target, and the matching
   *  negative margins keep the line where it was. */
  const GUILD_LINK =
    'text-muted hover:text-strong underline underline-offset-2 max-md:inline-block max-md:py-[11px] max-md:-my-[11px]';
  const LINE = 'text-muted text-[13px] leading-[22px]';

  const last = $derived(header.pieces.length - 1);
</script>

{#snippet chip()}<span class={CHIP} data-testid="character-ruleset-chip">{header.chip}</span>{/snippet}

{#snippet piece(item: ProfileHeaderPiece)}{#if item.key === 'spec'}<span
      class="text-text"
      data-testid="character-spec">{item.text}</span
    >{#if item.split !== undefined}{characterHeaderCopy.separator}<span
        class="tabular"
        data-testid="character-split">{item.split}</span
      >{/if}{:else if item.href !== undefined}<a
      class={GUILD_LINK}
      href={item.href}
      data-testid="character-guild-link">{item.text}</a
    >{:else}<span data-testid={`character-${item.key}`}>{item.text}</span>{/if}{/snippet}

{#if header.pieces.length > 0 || header.chipOn === 'pieces'}
  <p class={LINE} data-testid="character-descriptor">
    {#each header.pieces as item, index (item.key)}{#if index > 0}{characterHeaderCopy.separator}{/if}{#if index === last && header.chipOn === 'pieces'}<span
          class="whitespace-nowrap">{@render piece(item)}{@render chip()}</span
        >{:else}{@render piece(item)}{/if}{/each}{#if header.pieces.length === 0}{@render chip()}{/if}
  </p>
{/if}
{#if header.levelRace !== null}
  <p class={`${LINE} whitespace-nowrap`} data-testid="character-level-line">
    <span data-testid="character-level-race">{header.levelRace}</span
    >{#if header.chipOn === 'levelRace'}{@render chip()}{/if}
  </p>
{/if}
{#if sync !== null}
  <p class="text-muted line-clamp-2 text-[12px] leading-[18px]" data-testid="character-sync">
    {syncText(sync)}{#if sync.staleTail !== null}{STALE_SEPARATOR}<span
        class="text-ember"
        data-testid="character-sync-stale">{sync.staleTail}</span
      >{/if}
  </p>
{/if}
