<!-- web/src/components/ui/SelectSheet.svelte -->
<!-- The dropdown's part of the component sheet (design/specs/2026-10-30-controls-charts.md §A10
     row 42, capture 5): every closed-field state at 52, 44 and 36 px with hover, focus-visible,
     open and pressed forced by `data-force`, the lead-icon fields, and one live field of every
     variant (plain with descriptions, icons and searchable, grouped searchable, multi-select,
     right slot, item icons) that the e2e, axe and capture specs open. -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import { classCrestSrc } from '../../lib/class-crest';
  import { button } from '../../lib/ui/button';
  import { dataUrl } from '../../lib/planner/load';
  import { selectCopy } from '../../lib/ui/copy';
  import type { SelectOption } from '../../lib/ui/listbox-model';
  import { simCopy } from '../../lib/sim/copy';
  import Select, { type SelectSize } from './Select.svelte';

  const SIZES: readonly { size: SelectSize; name: string }[] = [
    { size: 'default', name: 'Default · 52 px · forms' },
    { size: 'compact', name: 'Compact · 44 px · filter bars, card heads, every field on a coarse pointer' },
    { size: 'dense', name: 'Dense · 36 px · report and sim tool bars, fine pointer from 768 px' },
  ];
  const FORCED = ['default', 'hover', 'focus-visible', 'open', 'active'] as const;
  const STYLES: SelectOption[] = Object.entries(simCopy.fightStyleOptions).map(([value, description]) => ({
    value,
    label: value.replace(/-/g, ' ').replace(/^./, (c) => c.toUpperCase()),
    description,
  }));
  const PATCHWERK: SelectOption[] = [{ value: 'patchwerk', label: 'Patchwerk' }, ...STYLES.slice(1)];

  const crest = (cls: string) => ({ kind: 'crest' as const, src: classCrestSrc(cls) });
  const player = (
    name: string,
    cls: string,
    group: string,
    extra: Partial<SelectOption> = {},
  ): SelectOption => ({
    value: name.toLowerCase(),
    label: name,
    group,
    icon: crest(cls),
    colour: `--color-class-${cls}`,
    ...extra,
  });
  const PLAYERS: SelectOption[] = [
    { value: 'friendlies', label: 'All friendlies' },
    { value: 'enemies', label: 'All enemies' },
    player('Thandril', 'warrior', 'Tanks'),
    player('Brokk', 'warrior', 'Tanks'),
    player('Corwin', 'priest', 'Healers'),
    player('Selwyn', 'paladin', 'Healers'),
    player('Garrosk', 'warrior', 'Damage'),
    player('Velyn', 'rogue', 'Damage'),
    player('Morrigan', 'mage', 'Damage'),
    player('Ysolde', 'mage', 'Damage'),
    player('Ashka', 'warlock', 'Damage'),
    player('Kael', 'hunter', 'Damage'),
    player('Durgan', 'warrior', 'Damage'),
    player('Kargath', 'warrior', 'Damage'),
    player('Zaxis', 'rogue', 'Damage'),
    player('Lightningbolltz', 'shaman', 'Damage', { you: true }),
  ];
  const SOURCES: SelectOption[] = [
    { value: 'anywhere', label: 'Anywhere' },
    { value: 'mc', label: 'Molten Core', group: 'Raids' },
    { value: 'bwl', label: 'Blackwing Lair', group: 'Raids' },
    { value: 'ony', label: 'Onyxia’s Lair', group: 'Raids' },
    { value: 'brd', label: 'Blackrock Depths', group: 'Dungeons' },
    { value: 'lbrs', label: 'Lower Blackrock Spire', group: 'Dungeons' },
    { value: 'ubrs', label: 'Upper Blackrock Spire', group: 'Dungeons' },
    { value: 'strat', label: 'Stratholme', group: 'Dungeons' },
    { value: 'kazzak', label: 'Lord Kazzak', group: 'World bosses' },
  ];
  const ENEMIES: SelectOption[] = [
    { value: 'all', label: 'Every enemy' },
    { value: 'rag', label: 'Ragnaros' },
    { value: 'sons', label: 'Son of Flame', meta: '×8' },
  ];
  const COMPARE: SelectOption[] = [
    {
      value: 'median',
      label: 'Raid median',
      group: 'Compare',
      description: 'Median DPS of the 28 damage dealers',
    },
    ...PLAYERS.slice(2).map((option) => ({ ...option, meta: String(100 + option.label.length * 37) })),
  ];
  const icon = (stem: string): string => dataUrl(activeBuild.build, `icons/${stem}.webp`);
  const ITEMS: SelectOption[] = [
    {
      value: 'a',
      label: 'Axe Specialization',
      icon: { kind: 'item', src: icon('fixture_axe_specialization'), quality: 'epic' },
    },
    {
      value: 'b',
      label: 'Booming Voice',
      icon: { kind: 'item', src: icon('fixture_booming_voice'), quality: 'rare' },
    },
    { value: 'c', label: 'Cruelty', icon: { kind: 'spec', src: icon('fixture_cruelty') } },
  ];

  let style = $state('patchwerk');
  let source = $state('morrigan');
  let where = $state('anywhere');
  let enemy = $state('all');
  let item = $state('a');
  let dialog = $state<HTMLDialogElement | null>(null);
  let dialogStyle = $state('patchwerk');
  let drawn = $state(['median', 'garrosk', 'lightningbolltz']);
  const TOP = ['median', 'garrosk', 'lightningbolltz'];
</script>

<section class="flex flex-col gap-6" aria-labelledby="sheet-select" data-testid="sheet-selects">
  <h2 id="sheet-select" class="heading-section">Select</h2>

  {#each SIZES as row (row.size)}
    <div class="flex flex-col gap-3" data-testid="sheet-select-{row.size}">
      <span class="text-muted numeric text-[11px]">{row.name}</span>
      <div class="grid grid-cols-2 gap-4 md:grid-cols-5">
        {#each FORCED as forced (forced)}
          <div class="flex flex-col gap-2" data-force={forced}>
            <label class="quiet-label" for="sheet-{row.size}-{forced}">{forced}</label>
            <Select id="sheet-{row.size}-{forced}" size={row.size} value="patchwerk" options={PATCHWERK} />
          </div>
        {/each}
        <div class="flex flex-col gap-2">
          <label class="quiet-label" for="sheet-{row.size}-disabled">disabled</label>
          <Select
            id="sheet-{row.size}-disabled"
            size={row.size}
            value="patchwerk"
            options={PATCHWERK}
            disabled
          />
        </div>
        <div class="flex flex-col gap-2">
          <label class="quiet-label" for="sheet-{row.size}-error">error</label>
          <Select
            id="sheet-{row.size}-error"
            size={row.size}
            value="patchwerk"
            options={PATCHWERK}
            error="Pick a fight style."
          />
        </div>
        <div class="flex flex-col gap-2">
          <label class="quiet-label" for="sheet-{row.size}-loading">loading</label>
          <Select id="sheet-{row.size}-loading" size={row.size} options={PATCHWERK} loading />
        </div>
        <div class="flex flex-col gap-2">
          <label class="quiet-label" for="sheet-{row.size}-empty">empty</label>
          <Select
            id="sheet-{row.size}-empty"
            size={row.size}
            options={[]}
            emptyLabel={selectCopy.noEnemies}
          />
        </div>
        <div class="flex flex-col gap-2">
          <label class="quiet-label" for="sheet-{row.size}-placeholder">placeholder</label>
          <Select
            id="sheet-{row.size}-placeholder"
            size={row.size}
            value=""
            options={PATCHWERK}
            placeholder="Choose a spec"
          />
        </div>
      </div>
    </div>
  {/each}

  <div class="flex flex-col gap-3">
    <span class="text-muted numeric text-[11px]">Live: every variant</span>
    <div class="grid gap-6 md:grid-cols-3">
      <div class="flex flex-col gap-2">
        <label class="quiet-label" for="sheet-fight-style">Fight style</label>
        <Select
          id="sheet-fight-style"
          size="compact"
          bind:value={style}
          options={STYLES}
          testid="sheet-fight-style"
        />
      </div>
      <div class="flex items-center gap-2">
        <label class="quiet-label" for="sheet-source">Source</label>
        <Select
          id="sheet-source"
          class="w-[220px]"
          size="dense"
          bind:value={source}
          options={PLAYERS}
          searchable
          searchPlaceholder={selectCopy.filterPlayers}
          testid="sheet-source"
        />
      </div>
      <div class="flex flex-col gap-2">
        <label class="quiet-label" for="sheet-where">Source</label>
        <Select
          id="sheet-where"
          size="compact"
          bind:value={where}
          options={SOURCES}
          searchable
          searchPlaceholder={selectCopy.filterSources}
          testid="sheet-where"
        />
      </div>
      <div class="flex flex-col gap-2">
        <label class="quiet-label" for="sheet-item">Item</label>
        <Select id="sheet-item" size="compact" bind:value={item} options={ITEMS} testid="sheet-item" />
      </div>
      <div class="flex items-center justify-end gap-2">
        <label class="quiet-label" for="sheet-enemy">Enemy</label>
        <Select
          id="sheet-enemy"
          class="w-[180px]"
          size="dense"
          align="end"
          bind:value={enemy}
          options={ENEMIES}
          testid="sheet-enemy"
        />
      </div>
      <div class="flex items-center justify-end gap-2">
        <label class="quiet-label" for="sheet-players">Players</label>
        <Select
          id="sheet-players"
          class="w-[150px]"
          size="dense"
          align="end"
          multiple
          bind:values={drawn}
          options={COMPARE}
          fieldLabel={drawn.join() === TOP.join()
            ? `Top 8 of ${COMPARE.length - 1}`
            : `${drawn.length} of ${COMPARE.length - 1}`}
          searchPlaceholder={selectCopy.filterPlayers}
          footActions={[
            { id: 'top', label: 'Top 8', pressed: drawn.join() === TOP.join() },
            { id: 'custom', label: 'Custom', pressed: drawn.join() !== TOP.join(), inert: true },
          ]}
          footCount="{drawn.length} drawn"
          onfoot={() => (drawn = [...TOP])}
          testid="sheet-players"
        />
      </div>
      <div class="flex flex-col gap-2">
        <span class="quiet-label">In a dialog</span>
        <button
          type="button"
          class={button('secondary')}
          data-testid="sheet-dialog-open"
          onclick={() => dialog?.showModal()}>Open a dialog</button
        >
        <dialog
          bind:this={dialog}
          class="card-surface text-text m-auto w-[360px] p-5 backdrop:bg-black/60"
          aria-labelledby="sheet-dialog-title"
          data-testid="sheet-dialog"
        >
          <h3 id="sheet-dialog-title" class="heading-section mb-3">Settings</h3>
          <label class="quiet-label" for="sheet-dialog-style">Fight style</label>
          <Select
            id="sheet-dialog-style"
            size="compact"
            bind:value={dialogStyle}
            options={PATCHWERK}
            testid="sheet-dialog-style"
          />
        </dialog>
      </div>
    </div>
  </div>
</section>

<style>
  /* Forced states: the look the pointer and keyboard give a real field (§A5). */
  :global([data-force='hover'] .sel-face) {
    border-color: rgb(229 185 85 / 0.45);
    background: rgb(255 255 255 / 0.06);
  }
  :global([data-force='focus-visible'] .sel-face) {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  :global([data-force='open'] .sel-face) {
    border-color: var(--color-gold);
    background: rgb(255 255 255 / 0.06);
    box-shadow: var(--glow-gold);
  }
  :global([data-force='open'] .sel-chev) {
    rotate: 180deg;
    color: var(--color-strong);
  }
  :global([data-force='active'] .sel-face) {
    border-color: var(--color-gold);
    background: #101624;
  }
</style>
