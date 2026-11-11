<!-- web/src/components/ui/UiSheet.svelte -->
<!-- The component sheet (polish spec 4.2): every shared control in each of its states on one
     page, served at /design-sheet. It is the living reference for design/DESIGN-SYSTEM.md and
     the page the axe, keyboard and capture specs drive (tests/e2e/ui-components.spec.ts).
     Each state is labelled. Hover, focus-visible and active are forced here with static classes
     (buttons, links) or a `data-force` attribute (switch, rows, select) so one capture shows them
     all; on a real page they come from the pointer and the keyboard. -->
<script lang="ts">
  import { button, type ButtonTone } from '../../lib/ui/button';
  import Button from './Button.svelte';
  import { PLATE_SURFACE } from '../../lib/ui/primitives';
  import ActionLink from './ActionLink.svelte';
  import {
    ICON_GUILD,
    ICON_LOGS,
    ICON_MANAGE,
    ICON_PREMIUM,
    ICON_RANKINGS,
    ICON_SETUP,
    ICON_SIM,
  } from '../../lib/ui/ui-icons';
  import SettingsRow from './SettingsRow.svelte';
  import SelectSheet from './SelectSheet.svelte';
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
    danger: 'Remove channel',
  };
  const RING = 'outline-2 outline-offset-2 outline-gold';
  /** The look a real hover or press gives each tone, as static classes. */
  const FORCED: Readonly<Record<ButtonTone, Partial<Record<ForcedState, string>>>> = {
    primary: { hover: 'brightness-[1.08]', 'focus-visible': RING, active: 'translate-y-px brightness-[.92]' },
    secondary: { hover: 'border-gold/30 bg-white/[.08]', 'focus-visible': RING, active: 'bg-raised' },
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
    <h2 id="sheet-buttons" class="text-strong text-[22px] leading-7 font-bold">Buttons</h2>
    {#each TONES as tone (tone)}
      <div class="flex flex-wrap items-end gap-x-6 gap-y-3" data-testid="sheet-buttons-{tone}">
        {#each STATES as state (state)}
          <div class="flex flex-col gap-2">
            {@render label(state)}
            {#if state === 'loading'}
              <Button {tone} busy busyLabel={TONE_WORDS[tone]}>{TONE_WORDS[tone]}</Button>
            {:else if state === 'disabled'}
              <button class={button(tone)} disabled>{TONE_WORDS[tone]}</button>
            {:else}
              <button class="{button(tone)} {FORCED[tone][state] ?? ''}" data-state={state}
                >{TONE_WORDS[tone]}</button
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
    <h2 id="sheet-switch" class="text-strong text-[22px] leading-7 font-bold">Switch</h2>
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
    <h2 id="sheet-rows" class="text-strong text-[22px] leading-7 font-bold">Settings rows</h2>
    <div class="grid gap-6 md:grid-cols-2">
      <div class="flex flex-col gap-2">
        {@render label('default, hover, focus-visible, active')}
        <div class="{PLATE_SURFACE} overflow-hidden" data-testid="sheet-rows">
          <SettingsRow label="Server sims" glyph={ICON_SIM} value="On · 3 runs this month" href="#sim" />
          <div data-force="hover">
            <SettingsRow
              label="History charts"
              glyph={ICON_RANKINGS}
              value="Open on your main character"
              href="#history"
            />
          </div>
          <div data-force="focus-visible">
            <SettingsRow
              label="Supporter mark"
              glyph={ICON_PREMIUM}
              value="Shown on your characters"
              href="#mark"
            />
          </div>
          <div data-force="active">
            <SettingsRow
              label="Discord notifications"
              glyph={ICON_GUILD}
              value="On · 2 events"
              href="#discord"
            />
          </div>
        </div>
      </div>
      <div class="flex flex-col gap-2">
        {@render label('muted, loading, empty, error, disabled')}
        <div class="{PLATE_SURFACE} overflow-hidden">
          <SettingsRow label="Quiet value" glyph={ICON_SETUP} value="Premium only" muted />
          <SettingsRow label="Loading value" glyph={ICON_SIM} loading />
          <SettingsRow label="Empty value" glyph={ICON_RANKINGS} empty />
          <SettingsRow label="Error value" glyph={ICON_LOGS} error="Could not load" />
          <SettingsRow label="Disabled row" glyph={ICON_MANAGE} value="Off" disabled />
        </div>
      </div>
      <div class="flex flex-col gap-2">
        {@render label('with switches, nested')}
        <div class="{PLATE_SURFACE} overflow-hidden">
          <SettingsRow label="A server sim finishes" glyph={ICON_SIM}>
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
            glyph={ICON_MANAGE}
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
    <h2 id="sheet-tags" class="text-strong text-[22px] leading-7 font-bold">Tags</h2>
    <div class="flex flex-wrap gap-3" data-testid="sheet-tags">
      <Tag>Founder</Tag>
      <Tag tone="success" dot>On · 2 events</Tag>
      <Tag tone="muted">Sample</Tag>
      <Tag tone="danger">Off · needs attention</Tag>
    </div>
  </section>

  <section class="flex flex-col gap-4" aria-labelledby="sheet-links">
    <h2 id="sheet-links" class="text-strong text-[22px] leading-7 font-bold">Links</h2>
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
    <h2 id="sheet-stats" class="text-strong text-[22px] leading-7 font-bold">Stat tiles</h2>
    <div class="flex flex-wrap gap-3" data-testid="sheet-stats">
      <StatTile value="07" label="Days" />
      <StatTile value="3" label="Runs this month" />
      <StatTile value="88" label="Forever Score" />
    </div>
  </section>

  <SelectSheet />
</div>

<style>
  /* Forced states for the controls that take them from the pointer or keyboard on a real page. */
  :global([data-force='focus-visible'] [role='switch']) {
    outline: 2px solid var(--color-gold);
    outline-offset: 3px;
  }
  :global([data-force='hover'] [role='switch'] > span:first-child) {
    border-color: color-mix(in srgb, var(--color-gold) 30%, transparent);
  }
  :global([data-force='active'] [role='switch'] > span > span) {
    width: 20px;
  }
  :global([data-force='hover'] a[href]) {
    background: var(--color-hover);
  }
  :global([data-force='focus-visible'] a[href]) {
    outline: 2px solid var(--color-gold);
    outline-offset: -2px;
  }
  :global([data-force='active'] a[href]) {
    background: color-mix(in srgb, var(--color-gold) 10%, transparent);
  }
</style>
