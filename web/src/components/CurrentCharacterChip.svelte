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
     already does by remounting this component's state from scratch. -->
<script lang="ts">
  import { plannerHrefFor, simHrefFor, type CurrentCharacter } from '../lib/current-character';
  import { currentCharacterCopy } from '../lib/current-character-copy';
  import { classColorVar, rowLink } from '../lib/report/format';

  let {
    current,
    onforget,
    hasOwnPasteBox = false,
    addonCode = '',
  }: {
    current: CurrentCharacter | null;
    onforget: () => void;
    hasOwnPasteBox?: boolean;
    /** The planner's current build, already encoded as an FS1 string. Empty when the
     *  planner has no build to copy -- the chip then renders no copy-addon link at all,
     *  rather than a link that would fail to copy anything. */
    addonCode?: string;
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

  /** Two 44px rows on phone, one on desktop -- the one fixed height both rendered states
   *  share, so resolving between them never moves anything below the chip. */
  const CHIP_HEIGHT = 'h-[88px] md:h-11';
</script>

{#if current !== null}
  <div
    class={`border-line bg-raised rounded-panel overflow-hidden border px-3 text-[13px] ${CHIP_HEIGHT}`}
    data-testid="current-character-chip"
  >
    <div class="flex h-full flex-col md:flex-row md:items-center md:gap-3">
      <span
        class="flex h-11 shrink-0 items-center truncate font-semibold md:h-auto"
        style:color={classColorVar(current.classSlug)}
        data-testid="current-character-label"
      >
        {current.label}
      </span>
      <div
        class="flex h-11 flex-nowrap items-center gap-3 overflow-x-auto whitespace-nowrap md:h-auto md:flex-1"
      >
        <a class={rowLink} href={plannerHrefFor(current)} data-testid="current-character-planner">
          {currentCharacterCopy.openInPlanner}
        </a>
        <a class={rowLink} href={simHrefFor(current)} data-testid="current-character-sim">
          {currentCharacterCopy.openInSimulator}
        </a>
        {#if addonCode !== ''}
          <button
            type="button"
            class={rowLink}
            onclick={copyAddonCode}
            data-testid="current-character-copy-addon"
          >
            {copied ? currentCharacterCopy.copiedAddonCode : currentCharacterCopy.copyAddonCode}
          </button>
        {/if}
        <button
          type="button"
          class={`${rowLink} text-muted md:ml-auto`}
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
    class={`text-muted flex items-center overflow-hidden text-[13px] ${CHIP_HEIGHT}`}
    data-testid="current-character-chip"
  >
    {currentCharacterCopy.noCharacterLine}
  </p>
{/if}
