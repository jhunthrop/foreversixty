# Applying the identity and listings

Everything here is prepared in the repo and nothing has been uploaded. Each platform is an
outward-facing change, so the owner applies it by hand. Rebuild the images first if any source
changed:

```sh
node design/community/discord/build-assets.mjs   # Discord
node design/listing/build-store.mjs              # store icons, feature graphic, store screenshots
node design/listing/check-store-text.mjs         # store text limits and voice
node design/listing/gen_gallery.mjs              # addon gallery (needs the in-game captures)
```

Each command prints one line per file and ends each line in `ok`.

## Discord

Files are in `design/community/discord/assets/`.

- [ ] **Server icon.** Server Settings > Overview > Server icon: upload `icon.png` (512 px, the painted tile).
- [ ] **Server banner.** Server Settings > Overview > Server banner (needs boost level 2): upload `banner.jpg` (1920x1080).
- [ ] **Invite background.** Server Settings > Overview > Invite background, if the server has it: upload `banner.jpg`.
- [ ] **Bot avatar.** Developer Portal > the Forever Sixty application > General Information > App icon: upload `avatar.png` (1024 px, the same tile). Then Bot > Icon: the same file.
- [ ] **Premium role icon.** Server Settings > Roles > Premium > Role icon (needs boost level 2): upload `premium.png` (the flat seal; Discord shows it at about 20 px).
- [ ] **Class crests and faction emoji** are unchanged in look and only rebuilt; re-upload only if Discord shows an old one.
- [ ] Or run `node design/community/discord/apply.mjs --images` with the bot token. It sets the server icon, banner and role icons, but not the bot avatar.
- [ ] Check: the server in the sidebar shows the painted seal, and the banner shows "Forever Sixty" in sentence case.

## CurseForge and Wago

The text is in `curseforge-description.md` (Forever Sixty) and `data-addon-description.md`
(Forever Sixty Data). Their "Uploading the gallery" section has the step-by-step gallery upload.

- [ ] **Project icon and account avatar**, on both projects and both sites: upload `icons/listing-icon-400.png`.
- [ ] **Summary or tagline:** paste the "One line" field from each file.
- [ ] **Description.** For CurseForge, paste `curseforge-description.html` (or `data-addon-description.html`) into the editor's HTML source view. For Wago, paste the Markdown under "## Description".
- [ ] **Gallery: wait.** It needs two things first. BA1 must have shipped the addon's new look, and the owner must take the twelve in-game shots in `RECAPTURE.md`. Then run `gen_gallery.mjs` and follow "Uploading the gallery". Delete the old navy plates from both dashboards when the new ones go up, not before.

## App Store Connect

The text is in `app-store/metadata.md` and the images are in `app-store/`.

- [ ] **App Information:** name, subtitle, primary and secondary category, Privacy Policy URL.
- [ ] **Version page:** promotional text, description, keywords, support URL, marketing URL, copyright, and What's New.
- [ ] **iPhone 6.9" screenshots:** upload `app-store/screenshots/01-home.jpg` to `04-gear.jpg` in order. There is no iPad set; the app is iPhone only.
- [ ] **App icon.** It comes from the build (`apps/shared/brand/ios`, the same tile). `app-store/icon-1024.png` is the same image, for anywhere a separate upload is asked.
- [ ] Answer the age-rating questionnaire.

## Google Play Console

The text is in `play/metadata.md` and the images are in `play/`.

- [ ] **Main store listing:** app name, short description and full description.
- [ ] **Graphics:** the app icon is `play/icon-512.png`, and the feature graphic is `play/feature-graphic.png`.
- [ ] **Phone screenshots:** upload `play/screenshots/01-home.jpg` to `04-gear.jpg` in order.
- [ ] **Store settings:** category, contact email, website and privacy policy.

## When something changes

- **The app's look changes:** rerun `design/listing/capture-apps.sh`, which drives the shared simulator and emulator under their locks, then `build-store.mjs`.
- **A seal, the tile or the ground changes:** rerun `web/scripts/build-app-icon.mjs`, `build-assets.mjs` and `build-store.mjs`. They all draw from `web/scripts/brand-art.mjs`.
