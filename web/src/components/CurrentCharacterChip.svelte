<!-- web/src/components/CurrentCharacterChip.svelte -->
<!-- Design 4 (architect review), fixed-height revision (fix round 1): the visible half of
     the current-character pointer. Reads the pointer it is given (never localStorage
     itself -- the mounting page owns when to read/restore, so this stays a pure render of
     one prop) and shows up to three links plus Forget, in class colour. The wrapper's
     height is a FIXED constant in every rendered state -- CHIP_HEIGHT, `h-[88px] md:h-11`
     with `overflow-hidden` -- not a min-height: two 44px rows on phone (label, then a
     non-wrapping, horizontally scrollable action row), collapsing to one 44px row on
     desktop where label and actions sit side by side. Both the loaded chip and the
     no-character line share the exact same wrapper height, and neither ever wraps its
     action row, so resolving from "no character" to "loaded" -- or the reverse, after
     Forget -- never moves content below it (Global Constraint: no layout shift).

     Copy addon code (spec section 1) is a third, optional link: the chip itself only ever
     holds a pointer, never the planner's talent index the addon-code string is built from,
     so the mounting page (the planner) passes the already-built string down and this stays
     a pure render. `copied` has no timer, same idiom as SharePanel.svelte's `copiedFrom`:
     it stands until the character changes underneath it, which a fresh `current` prop
     already does by remounting this component's state from scratch.

     `restored` (fix round 1, Task 4's review): true when the mounting page's own load came
     from the stored pointer rather than a fresh paste or a URL. The tools island reserves
     a slot for this chip before hydration (`bulk-skeleton.ts`'s own `chipSlot`) rather than
     rendering its own ad-hoc restored line, so the "Restored your last character" wording
     lives here, in the one place every /sim* page's chip already renders. -->
<script lang="ts">
  import { plannerHrefFor, simHrefFor, type CurrentCharacter } from '../lib/current-character';
  import { currentCharacterCopy } from '../lib/current-character-copy';
  import { CHIP_HEIGHT } from '../lib/current-character-layout';
  import { classColorVar, rowLink } from '../lib/report/format';

  let {
    current,
    onforget,
    hasOwnPasteBox = false,
    addonCode = '',
    restored = false,
    guildLine = '',
  }: {
    current: CurrentCharacter | null;
    onforget: () => void;
    hasOwnPasteBox?: boolean;
    /** The planner's current build, already encoded as an FS1 string. Empty when the
     *  planner has no build to copy -- the chip then renders no copy-addon link at all,
     *  rather than a link that would fail to copy anything. */
    addonCode?: string;
    /** True once the mounting page's own load came from the stored pointer rather than
     *  the URL. Only shown alongside a loaded `current` -- a page that passes this true
     *  with no character loaded (a dead pointer, say) sees the ordinary empty state. */
    restored?: boolean;
    /** "Iron Vanguard · Officer", already formatted -- spec 2026-09-22 §7.5. Built by
     *  CurrentCharacterBar.svelte, which is the one place with `/v1/me` in hand; this stays
     *  a pure render of the string it is given, same as `addonCode`. Empty when the current
     *  character has no known guild (not an 'armory' pointer, no match, or unguilded). */
    guildLine?: string;
  } = $props();

  let copied = $state(false);

  async function copyAddonCode(): Promise<void> {
    try {
      await navigator.clipboard.writeText(addonCode);
      copied = true;
    } catch {
      copied = false;
    }
  }

  /** The 44px hit target (`rowLink`) plus the nav link colour, shared by the three
   *  colour-bearing actions -- planner, simulator, copy addon code -- so the two classes
   *  stay paired in one place instead of three call sites that could drift apart (fix
   *  round 2: `rowLink` alone carries no colour, and the planner/sim links and the copy
   *  button lost `text-nav` in the round-1 swap). Forget is deliberately not one of these:
   *  it keeps its own `text-muted`, unchanged. */
  const ACTION = `${rowLink} text-nav`;
</script>

{#if current !== null}
  <div
    class={`border-line bg-raised rounded-panel mx-[18px] overflow-hidden border px-3 text-[13px] md:mx-0 ${CHIP_HEIGHT}`}
    data-testid="current-character-chip"
  >
    <div class="flex h-full flex-col md:flex-row md:items-center md:gap-3">
      <span
        class="flex h-11 shrink-0 items-center gap-1.5 truncate font-semibold md:h-auto"
        data-testid="current-character-label"
      >
        {#if restored}
          <span class="text-muted font-normal" data-testid="current-character-restored">
            {currentCharacterCopy.restoredNote}
          </span>
        {/if}
        <span style:color={classColorVar(current.classSlug)}>{current.label}</span>
        {#if guildLine !== ''}
          <span class="text-muted truncate font-normal" data-testid="current-character-guild">
            {guildLine}
          </span>
        {/if}
      </span>
      <div
        class="flex h-11 flex-nowrap items-center gap-3 overflow-x-auto whitespace-nowrap md:h-auto md:flex-1"
      >
        <a class={ACTION} href={plannerHrefFor(current)} data-testid="current-character-planner">
          {currentCharacterCopy.openInPlanner}
        </a>
        <a class={ACTION} href={simHrefFor(current)} data-testid="current-character-sim">
          {currentCharacterCopy.openInSimulator}
        </a>
        {#if addonCode !== ''}
          <button
            type="button"
            class={ACTION}
            onclick={copyAddonCode}
            data-testid="current-character-copy-addon"
          >
            {copied ? currentCharacterCopy.copiedAddonCode : currentCharacterCopy.copyAddonCode}
          </button>
        {/if}
        <button
          type="button"
          class={`${rowLink} text-muted min-w-11 justify-center md:ml-auto md:justify-start`}
          onclick={onforget}
          data-testid="current-character-forget"
        >
          {currentCharacterCopy.forget}
        </button>
      </div>
    </div>
  </div>
{:else if !hasOwnPasteBox}
  <p
    class={`text-muted mx-[18px] flex items-center overflow-hidden text-[13px] md:mx-0 ${CHIP_HEIGHT}`}
    data-testid="current-character-chip"
  >
    {currentCharacterCopy.noCharacterLine}
  </p>
{/if}
