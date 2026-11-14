# Applying the identity and listings

Everything here is prepared in the repo and nothing has been uploaded. Each platform is an
outward-facing change, so the owner applies it by hand. Rebuild the images first if any source
changed:

```sh
node design/community/discord/build-assets.mjs   # Discord
node design/listing/build-store.mjs              # store icons, feature graphic, store screenshots
node design/listing/check-store-text.mjs         # store text limits and voice
python3 design/listing/gen_captures.py           # addon gallery captures, from the mocks
node design/listing/gen_gallery.mjs              # addon gallery plates
```

Each command prints one line per file and ends each line in `ok`.

## Discord

Files are in `design/community/discord/assets/`.

- **Server icon and banner: no change.** The owner kept the live seal v4 icon and its banner (2026-11-11); `icon.png` and `banner.jpg` are copies of what is live.
- [ ] **Invite background.** Server Settings > Overview > Invite background, if the server has it: upload `splash.jpg`, the same art with "Forever Sixty" on it.
- [ ] **Bot avatar.** Developer Portal > the Forever Sixty application > General Information > App icon: upload `avatar.png` (1024 px, the painted tile). Then Bot > Icon: the same file.
- [ ] **Premium role icon.** Server Settings > Roles > Premium > Role icon (needs boost level 2): upload `premium.png` (the flat seal's tiny bold drawing; Discord shows it at about 20 px).
- [ ] **Class crests and faction emoji** are unchanged in look and only rebuilt; re-upload only if Discord shows an old one.
- [ ] Or run `node design/community/discord/apply.mjs --images` with the bot token. It sets the server icon, banner and role icons (the icon and banner re-send the live seal v4), but not the invite background or the bot avatar.
- [ ] Check: the server list still shows the seal v4 icon.

## CurseForge and Wago

The text is in `curseforge-description.md` (Forever Sixty) and `data-addon-description.md`
(Forever Sixty Data). Their "Uploading the gallery" section has the step-by-step gallery upload.

- [ ] **Project icon and account avatar**, on both projects and both sites: upload `icons/listing-icon-400.png`.
- [ ] **Summary or tagline:** paste the "One line" field from each file.
- [ ] **Description.** For CurseForge, paste `curseforge-description.html` (or `data-addon-description.html`) into the editor's HTML source view. For Wago, paste the Markdown under "## Description".
- [ ] **Gallery.** The eleven plates in `gallery/` come from the addon's mocks (`RECAPTURE.md`); follow "Uploading the gallery" in `curseforge-description.md`, deleting the old plates first.

## Before the store screenshots go up

The store plates show the apps exactly as they are on the fixtures today. Four things in the apps
should change first; each is the apps' own work, and none is an image fix. After them, rerun
`design/listing/capture-apps.sh` and `build-store.mjs`.

- [ ] Android: the profile button and the gear header's medallion draw an empty ring when the
      portrait does not load (`GearRoute.kt` passes a portrait painter with no fallback). iOS falls
      back to the LX seal.
- [ ] The retired forced caps on the raid page ("ONYXIA'S LAIR", "YOUR SIGN-UP", "NOTES", "SIGN-UPS"
      on iOS, the title on Android): queued lane F.
- [ ] Home's activity row says "Sign up In, Late or Out." Bench is missing.
- [ ] The fixture character: a level 58 Fury Warrior with hunter pants and an engineering gun as its
      next upgrades, "2 levels to band 60", and "0 of 17 slots on the list" on Gear. A level 60
      character with real picks and a part-done list would read right to a player.

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
- [ ] **Graphics:** the app icon is `play/icon-512.png` (the full square; Play rounds it, and the round one on the review sheet is only a preview). The feature graphic is `play/feature-graphic.png`.
- [ ] **Phone screenshots:** upload `play/screenshots/01-home.jpg` to `04-gear.jpg` in order.
- [ ] **Store settings:** category, contact email, website and privacy policy.

## When something changes

- **The app's look changes:** rerun `design/listing/capture-apps.sh`, which drives the shared simulator and emulator under their locks, then `build-store.mjs`.
- **A seal, the tile or the ground changes:** rerun `web/scripts/build-app-icon.mjs`, `build-assets.mjs` and `build-store.mjs`. They all draw from `web/scripts/brand-art.mjs`.
