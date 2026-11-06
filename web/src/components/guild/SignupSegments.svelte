<!-- web/src/components/guild/SignupSegments.svelte -->
<!-- In · Late · Bench · Out (spec 3.3, 3.4, 4; owner decision D11): the sign-up's states drawn by
     the one segmented primitive (ui/Segmented), so the site's sign-up and every other in-row choice
     share a look and a keyboard model. Full width on phone. -->
<script lang="ts">
  import type { GuildSignupState } from '../../lib/guild/api';
  import { raidCopy } from '../../lib/guild/raid-copy';
  import { SIGNUP_STATES } from '../../lib/guild/segments-keys';
  import Segmented from '../ui/Segmented.svelte';

  let {
    value,
    pending = null,
    name,
    size,
    disabled = false,
    describedBy,
    onChange,
    testid = 'signup-segments',
  }: {
    value: GuildSignupState | null;
    /** Pressed but not saved (the spec pick is owed): drawn as pending, never as checked. */
    pending?: GuildSignupState | null;
    name: string;
    /** `strip` is 44 px everywhere; `band` is the in-row size from md. */
    size: 'band' | 'strip';
    disabled?: boolean;
    /** The state line's id, read while disabled. */
    describedBy?: string;
    onChange: (state: GuildSignupState) => void;
    testid?: string;
  } = $props();

  const options = SIGNUP_STATES.map((state) => ({ value: state, label: raidCopy.segments[state] }));
</script>

<Segmented
  {options}
  {value}
  {pending}
  label={`${name}'s sign-up`}
  size={size === 'band' ? 'sm' : 'md'}
  fill="phone"
  {disabled}
  {describedBy}
  {onChange}
  {testid}
/>
