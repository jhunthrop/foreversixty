<!-- web/src/components/character/ProfileDescriptor.svelte -->
<!-- The character profile's lines under the name (Task L8.4 ruling 1, fix round 2 items 3 to 5,
     fix round 3 items 2, 3, 11 and 14, fix round 5 item 6), all on the name's left edge:
       1. "Holy Priest · 31/20/0 · <Olympus XXVII> · Classic Beta PvP": "Holy Priest" in --text,
          the rest muted; each piece only when known. Each segment is one nowrap unit that ends
          in its "·" (a no-break space, then the dot), and plain spaces between the units are the
          only breaks, so a wrap never starts a line with a dot; the realm moves whole (fix
          round 5 item 6). The guild links to its page: its unit is an inline-block at most the
          line's width, so it moves whole when it fits a line and a guild longer than the line
          still wraps inside it (fix round 3 item 2); its 44px phone tap target is vertical
          padding, its focus ring at no offset so it does not reach over the dots.
       2. "Level 60 Night Elf", when known.
       The "HARDCORE US" chip ends the last of those lines, 8px after its text, aligned to the
          line box so it is centred on the text, and never wraps away from the last word: a word
          joiner in a nowrap span, inside the last unit, glues it there.
       3. The owner's sync line, the Character header's own: "Addon · 4 minutes ago", and past
          the nav spec's stale limit "Log out in the game to sync" in ember. Plain data in:
          Character.svelte works out ownership and the sync (lib/character-header/sync.ts). -->
<script lang="ts">
  import { characterHeaderCopy, type ProfileHeader } from '../../lib/character-profile';
  import { STALE_SEPARATOR, syncText, type SyncView } from '../../lib/character-header/sync';

  let { header, sync = null }: { header: ProfileHeader; sync?: SyncView | null } = $props();

  /** One unit of line 1: "Holy Priest", the split, the guild or the realm. */
  interface Segment {
    key: 'spec' | 'split' | 'guild' | 'realm';
    text: string;
    href?: string;
  }

  /** The pieces as line 1's units: the spec piece's split is a unit of its own. */
  const segments = $derived(
    header.pieces.flatMap((piece): Segment[] =>
      piece.key === 'spec' && piece.split !== undefined
        ? [
            { key: 'spec', text: piece.text },
            { key: 'split', text: piece.split },
          ]
        : [{ key: piece.key, text: piece.text, href: piece.href }],
    ),
  );

  /** Muted text and one border token, whichever line it ends (fix round 2 item 3); top-aligned,
   *  a 22px pill in a 22px line box, so its centre is the line's (fix round 3 item 14). */
  const CHIP = 'pill text-muted border-line align-top ml-2';
  /** In-text links are underlined: colour alone does not mark a link (WCAG 1.4.1). Inline and
   *  wrapping inside its unit, so a guild longer than the line wraps (fix round 3 item 2); on a
   *  phone 14px of padding above and below the 16px content box make the 44px target (15px made
   *  46: fix round 4 items 6 and 8), and an inline box's padding never moves the line. The focus
   *  ring keeps its 2px size at no offset (fix round 3 item 11); a press takes --gold-hover, the
   *  design system's link :active. */
  const GUILD_LINK =
    'text-muted hover:text-strong active:text-gold-hover underline underline-offset-2 focus-visible:outline-offset-0 max-md:py-[14px] whitespace-normal';
  /** A unit of line 1 (fix round 5 item 6): never broken inside. */
  const SEGMENT = 'whitespace-nowrap';
  /** The guild's unit: at most the line's width, so it moves whole when it fits and wraps inside
   *  itself only when it is longer than the line. */
  const GUILD_SEGMENT = 'inline-block max-w-full whitespace-nowrap';
  const LINE = 'text-muted text-[13px] leading-[22px]';
</script>

{#snippet chip()}<span class={CHIP} data-testid="character-ruleset-chip">{header.chip}</span>{/snippet}

{#snippet words(item: Segment)}{#if item.key === 'spec'}<span class="text-text" data-testid="character-spec"
      >{item.text}</span
    >{:else if item.key === 'split'}<span class="tabular" data-testid="character-split">{item.text}</span
    >{:else if item.href !== undefined}<a
      class={GUILD_LINK}
      href={item.href}
      data-testid="character-guild-link">{item.text}</a
    >{:else}<span data-testid={`character-${item.key}`}>{item.text}</span>{/if}{/snippet}

{#if segments.length > 0 || header.chipOn === 'pieces'}
  <p class={LINE} data-testid="character-descriptor">
    {#each segments as item, index (item.key)}{#if index > 0}{characterHeaderCopy.segmentBreak}{/if}<span
        class={item.key === 'guild' ? GUILD_SEGMENT : SEGMENT}
        data-testid="character-segment"
        >{@render words(
          item,
        )}{#if index < segments.length - 1}{characterHeaderCopy.segmentTail}{:else if header.chipOn === 'pieces'}<span
            class="whitespace-nowrap">{characterHeaderCopy.chipGlue}{@render chip()}</span
          >{/if}</span
      >{/each}{#if segments.length === 0 && header.chipOn === 'pieces'}{@render chip()}{/if}
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
