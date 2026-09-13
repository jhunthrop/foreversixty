/// <reference types="astro/client" />

// Type augmentation for the custom `client:interaction` directive registered in
// astro.config.mjs (see src/directives/interaction.ts for the runtime behavior).
declare namespace Astro {
  interface ClientDirectives {
    'client:interaction'?: boolean;
  }
}
