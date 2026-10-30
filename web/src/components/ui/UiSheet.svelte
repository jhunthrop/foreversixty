<!-- web/src/components/ui/UiSheet.svelte -->
<!-- The component sheet (polish spec 4.2): every shared control in each of its states on one
     page, served at /design-sheet. It is the living reference for design/DESIGN-SYSTEM.md and
     the page the axe, keyboard and capture specs drive (tests/e2e/ui-components.spec.ts).
     Each state is labelled. Hover, focus-visible and active are forced here with static classes
     (buttons, links) or a `data-force` attribute (switch, rows, select) so one capture shows them
     all; on a real page they come from the pointer and the keyboard. -->
<script lang="ts">
  import { button, type ButtonTone } from '../../lib/ui/button';
  import ActionLink from './ActionLink.svelte';
  import CardHead from './CardHead.svelte';
  import Glyph from './Glyph.svelte';
  import SettingsRow from './SettingsRow.svelte';
  import SelectSheet from './SelectSheet.svelte';
  import Spinner from './Spinner.svelte';
  import StatTile from './StatTile.svelte';
  import Switch from './Switch.svelte';
  import Tag from './Tag.svelte';

  type ForcedState = 'default' | 'hover' | 'focus-visible' | 'active' | 'disabled' | 'loading';

  const STATES: readonly ForcedState[] = [
    'default',
    'hover',
    'focus-visible',
    'active',
    'disabled',
    'loading',
  ];
  const TONES: readonly ButtonTone[] = ['primary', 'secondary', 'ghost', 'danger'];
  const TONE_WORDS: Readonly<Record<ButtonTone, string>> = {
    primary: 'Subscribe',
    secondary: 'Manage billing',
    ghost: 'See plans',
    danger: 'Unlink Discord',
  };
  const RING = 'outline-2 outline-offset-2 outline-gold';
  /** The look a real hover or press gives each tone, as static classes. */
  const FORCED: Readonly<Record<ButtonTone, Partial<Record<ForcedState, string>>>> = {
    primary: { hover: 'brightness-[1.08]', 'focus-visible': RING, active: 'translate-y-px brightness-[.92]' },
    secondary: { hover: 'border-gold/30 bg-white/[.08]', 'focus-visible': RING, active: 'bg-card-top' },
    ghost: { hover: 'text-gold-hover bg-gold/[.06]', 'focus-visible': RING, active: 'text-gold-hover' },
    danger: { hover: 'bg-wipe/[.12]', 'focus-visible': RING, active: 'bg-wipe/[.16]' },
  };

  let simSwitch = $state(true);
  let privateSwitch = $state(false);
</script>

{#snippet label(text: string)}
  <span class="text-muted numeric text-[11px]">{text}</span>
{/snippet}

<div class="flex flex-col gap-12 pb-16" data-testid="ui-sheet">
  <section class="flex flex-col gap-5" aria-labelledby="sheet-buttons">
    <h2 id="sheet-buttons" class="heading-section">Buttons</h2>
    {#each TONES as tone (tone)}
      <div class="flex flex-wrap items-end gap-x-6 gap-y-3" data-testid="sheet-buttons-{tone}">
        {#each STATES as state (state)}
          <div class="flex flex-col gap-2">
            {@render label(state)}
            {#if state === 'loading'}
              <button class={button(tone)} disabled aria-busy="true"><Spinner />{TONE_WORDS[tone]}</button>
            {:else if state === 'disabled'}
              <button class={button(tone)} disabled>{TONE_WORDS[tone]}</button>
            {:else}
              <button class="{button(tone)} {FORCED[tone][state] ?? ''}" data-state={state}
                ><Glyph name="card" />{TONE_WORDS[tone]}</button
              >
            {/if}
          </div>
        {/each}
        <div class="flex flex-col gap-2">
          {@render label('small (36 px from 1024, fine pointer)')}
          <button class={button(tone, 'sm')}>{TONE_WORDS[tone]}</button>
        </div>
      </div>
    {/each}
  </section>

  <section class="flex flex-col gap-4" aria-labelledby="sheet-switch">
    <h2 id="sheet-switch" class="heading-section">Switch</h2>
    <div class="flex flex-wrap items-end gap-x-8 gap-y-4" data-testid="sheet-switches">
      {#each [{ name: 'off', on: false }, { name: 'on', on: true }] as one (one.name)}
        <div class="flex flex-col gap-2">
          {@render label(one.name)}
          <span id="sw-{one.name}" class="sr-only">{one.name}</span>
          <Switch
            checked={one.on}
            labelledby="sw-{one.name}"
            onchange={() => {}}
            testid="sheet-switch-{one.name}"
          />
        </div>
      {/each}
      {#each ['hover', 'focus-visible', 'active'] as forced (forced)}
        <div class="flex flex-col gap-2" data-force={forced}>
          {@render label(forced)}
          <span id="sw-{forced}" class="sr-only">{forced}</span>
          <Switch checked={true} labelledby="sw-{forced}" onchange={() => {}} />
        </div>
      {/each}
      <div class="flex flex-col gap-2">
        {@render label('disabled')}
        <span id="sw-disabled" class="sr-only">Disabled</span>
        <Switch checked={false} disabled labelledby="sw-disabled" onchange={() => {}} />
      </div>
      <div class="flex flex-col gap-2">
        {@render label('saving')}
        <span id="sw-saving" class="sr-only">Saving</span>
        <Switch checked={true} saving labelledby="sw-saving" onchange={() => {}} />
      </div>
    </div>
  </section>

  <section class="flex flex-col gap-4" aria-labelledby="sheet-rows">
    <h2 id="sheet-rows" class="heading-section">Settings rows</h2>
    <div class="grid gap-6 md:grid-cols-2">
      <div class="flex flex-col gap-2">
        {@render label('default, hover, focus-visible, active')}
        <div class="card-surface overflow-hidden" data-testid="sheet-rows">
          <SettingsRow label="Server sims" glyph="server" value="On · 3 runs this month" href="#sim" />
          <div data-force="hover">
            <SettingsRow
              label="History charts"
              glyph="chart"
              value="Open on your main character"
              href="#history"
            />
          </div>
          <div data-force="focus-visible">
            <SettingsRow label="Supporter mark" glyph="star" value="Shown on your characters" href="#mark" />
          </div>
          <div data-force="active">
            <SettingsRow label="Discord notifications" glyph="bell" value="On · 2 events" href="#discord" />
          </div>
        </div>
      </div>
      <div class="flex flex-col gap-2">
        {@render label('muted, loading, empty, error, disabled')}
        <div class="card-surface overflow-hidden">
          <SettingsRow label="Quiet value" glyph="bell" value="Premium only" muted />
          <SettingsRow label="Loading value" glyph="star" loading />
          <SettingsRow label="Empty value" glyph="star" empty />
          <SettingsRow label="Error value" glyph="star" error="Could not load" />
          <SettingsRow label="Disabled row" glyph="star" value="Off" disabled />
        </div>
      </div>
      <div class="flex flex-col gap-2">
        {@render label('with switches, nested')}
        <div class="card-surface overflow-hidden">
          <SettingsRow label="A server sim finishes" glyph="server">
            {#snippet control(labelId)}
              <Switch
                checked={simSwitch}
                labelledby={labelId}
                onchange={(next) => (simSwitch = next)}
                testid="sheet-switch-live"
              />
            {/snippet}
          </SettingsRow>
          <SettingsRow
            label="Include private logs"
            glyph="lock"
            sub="Off: private logs are never posted."
            nested
          >
            {#snippet control(labelId)}
              <Switch
                checked={privateSwitch}
                labelledby={labelId}
                onchange={(next) => (privateSwitch = next)}
              />
            {/snippet}
          </SettingsRow>
        </div>
      </div>
    </div>
  </section>

  <section class="flex flex-col gap-4" aria-labelledby="sheet-tags">
    <h2 id="sheet-tags" class="heading-section">Tags</h2>
    <div class="flex flex-wrap gap-3" data-testid="sheet-tags">
      <Tag>Founder</Tag>
      <Tag tone="success" dot>On · 2 events</Tag>
      <Tag tone="muted">Sample</Tag>
      <Tag tone="danger">Off · needs attention</Tag>
    </div>
  </section>

  <section class="flex flex-col gap-4" aria-labelledby="sheet-links">
    <h2 id="sheet-links" class="heading-section">Links</h2>
    <div class="flex flex-wrap items-end gap-x-8 gap-y-3">
      <div class="flex flex-col gap-1">
        {@render label('default')}
        <ActionLink href="#specs">All 28 specs</ActionLink>
      </div>
      <div class="flex flex-col gap-1">
        {@render label('hover')}
        <ActionLink href="#specs" class="text-gold-hover [&>svg]:translate-x-0.5">All 28 specs</ActionLink>
      </div>
      <div class="flex flex-col gap-1">
        {@render label('focus-visible')}
        <ActionLink href="#specs" class={RING}>All 28 specs</ActionLink>
      </div>
      <div class="flex flex-col gap-1">
        {@render label('active')}
        <ActionLink href="#specs" class="text-gold-hover">All 28 specs</ActionLink>
      </div>
    </div>
    <p class="text-[15px]">
      An inline link in prose reads like <a class="text-gold underline" href="#refunds">the refund policy</a>,
      with a soft gold underline.
    </p>
  </section>

  <section class="flex flex-col gap-4" aria-labelledby="sheet-stats">
    <h2 id="sheet-stats" class="heading-section">Stat tiles</h2>
    <div class="flex flex-wrap gap-3" data-testid="sheet-stats">
      <StatTile value="07" label="Days" />
      <StatTile value="3" label="Runs this month" />
      <StatTile value="88" label="Forever Score" />
    </div>
  </section>

  <SelectSheet />

  <section class="flex flex-col gap-4" aria-labelledby="sheet-head">
    <h2 id="sheet-head" class="heading-section">Card head</h2>
    <div class="card-surface max-w-[520px]" data-testid="sheet-card-head">
      <CardHead glyph="crown" label="Plan" title="Premium">
        {#snippet titleSnippet()}<Tag>Founder</Tag>{/snippet}
        {#snippet aside()}<Tag tone="success" dot>On · 2 events</Tag>{/snippet}
      </CardHead>
    </div>
  </section>
</div>

<style>
  /* Forced states for the controls that take them from the pointer or keyboard on a real page. */
  :global([data-force='focus-visible'] [role='switch']) {
    outline: 2px solid var(--color-gold);
    outline-offset: 3px;
  }
  :global([data-force='hover'] [role='switch'] > span:first-child) {
    border-color: rgba(229, 185, 85, 0.3);
  }
  :global([data-force='active'] [role='switch'] > span > span) {
    width: 20px;
  }
  :global([data-force='hover'] a[href]) {
    background: var(--row-hover);
  }
  :global([data-force='focus-visible'] a[href]) {
    outline: 2px solid var(--color-gold);
    outline-offset: -2px;
  }
  :global([data-force='active'] a[href]) {
    background: rgba(229, 185, 85, 0.1);
  }
</style>
