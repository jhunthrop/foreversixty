package gg.foreversixty.app.menu

import androidx.annotation.StringRes
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.WindowInsets
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.statusBars
import androidx.compose.foundation.layout.windowInsetsPadding
import androidx.compose.foundation.layout.wrapContentHeight
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.disabled
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.home.LARGE_TEXT_SCALE
import gg.foreversixty.app.isStale
import gg.foreversixty.app.account.AccountCopy
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsTarget
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.parityContent
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.designsystem.surfaceBand
import gg.foreversixty.app.flows.FlowRule
import gg.foreversixty.app.flows.SUB
import gg.foreversixty.app.model.AccountCharacter

/** The lines the menu's rows carry that come from elsewhere (the plan, the devices, offline data). */
data class MenuLines(
    val notifications: String? = null,
    val plan: String? = null,
    val devices: Int? = null,
    /** The most recently used device other than this one (round 2), for `3 signed in · newest Pixel 9`. */
    val newestDevice: String? = null,
    val offlineData: String? = null,
    val version: String = "",
    /** The operator line's year: the clock's at runtime, fixed in the goldens. */
    val year: Int,
)

/** What the menu does; each closes the menu first (the shell's job). */
data class MenuActions(
    val onSwitch: (AccountCharacter) -> Unit = {},
    val onRefresh: () -> Unit = {},
    val onOpen: (MenuDestination) -> Unit = {},
    val onSignIn: () -> Unit = {},
    val onHelp: () -> Unit = {},
    val onPrivacy: () -> Unit = {},
)

/**
 * The profile menu's content in the round-5 look (spec 2026-11-08 section 3, mock m-android-menu): You on
 * the painted band, then Switch character, Tools (labels alone, owner ruling 2026-11-09), Account and the
 * footer on the drawer's panel. Signed out ([model] null): Sign in with Battle.net and the tools that
 * work signed out.
 */
@Composable
fun ProfileMenu(
    model: MenuModel?,
    lines: MenuLines,
    actions: MenuActions,
    modifier: Modifier = Modifier,
    /** The low-level fold open at first (the menu-low capture). */
    lowLevelOpen: Boolean = false,
    /** Where the Battle.net stamp's Refresh is: at rest, under way or failed. */
    refresh: MenuRefresh = MenuRefresh.Idle,
) {
    // The status bar's inset sits outside the scroll, so nothing scrolls under the clock.
    Column(modifier.windowInsetsPadding(WindowInsets.statusBars).verticalScroll(rememberScrollState()).parityContent()) {
        val you = model?.you
        if (model == null || you == null) {
            Band { SignedOutHead(actions.onSignIn) }
            Body { MenuSection(R.string.menu_tools) { ToolRows(MenuDestination.signedOutTools, actions.onOpen) } }
        } else {
            Band { YouHead(you, model, refresh, actions.onRefresh) }
            Body {
                if (model.switchTo.isNotEmpty() || model.lowLevel.isNotEmpty()) SwitchSection(model, lowLevelOpen, actions.onSwitch)
                MenuSection(R.string.menu_tools) { ToolRows(MenuDestination.tools, actions.onOpen) }
                MenuSection(R.string.menu_account) {
                    MenuDestination.account.forEachIndexed { index, destination ->
                        MenuRow(labelOf(destination), accountLine(destination, lines), last = index == MenuDestination.account.lastIndex) {
                            actions.onOpen(destination)
                        }
                    }
                }
            }
        }
        Body { Footer(lines, actions) }
    }
}

/** The painted band at the drawer's top (the mock's `.you`), under the status bar. */
@Composable
private fun Band(content: @Composable () -> Unit) {
    Column(
        Modifier
            .fillMaxWidth()
            .surfaceBand()
            .padding(start = SIDE, end = SIDE, top = BAND_TOP, bottom = FsSpace.space3),
    ) { content() }
}

/** The drawer's panel under the band, 18 at the sides. */
@Composable
private fun Body(content: @Composable () -> Unit) {
    Column(Modifier.fillMaxWidth().padding(horizontal = SIDE)) { content() }
}

@Composable
private fun YouHead(you: AccountCharacter, model: MenuModel, refresh: MenuRefresh, onRefresh: () -> Unit) {
    Column(Modifier.fillMaxWidth().parityRegion(ParityRegion.HEAD)) {
        Row(
            Modifier.fillMaxWidth().parityRegion(ParityRegion.HEAD_TITLE),
            verticalAlignment = Alignment.CenterVertically,
            horizontalArrangement = Arrangement.spacedBy(FsSpace.space4),
        ) {
            CharacterPortrait(you.avatarUrl, you.faction, CREST)
            Column(Modifier.weight(1f)) {
                Text(you.name, style = NAME, color = classColor(you.className), modifier = Modifier.semantics { heading() })
                Column(Modifier.fillMaxWidth().parityRegion(ParityRegion.HEAD_META)) {
                    Text(headLine(you), style = FsType.small, color = FsColor.text)
                    if (model.isMain) MainTag()
                }
            }
        }
        model.updated?.let { age ->
            BattleNetStamp(stringResource(R.string.menu_updated, AccountCopy.age(age)), age.isStale(), you.build?.syncError != null, refresh, onRefresh)
        }
    }
}

/**
 * `Updated 11 d ago from Battle.net · Refresh` (spec section 3). A tap says so at once (row 12): the
 * action reads `Refreshing…` and is disabled until the call returns, then the new stamp, or the error
 * line under it. Past [isStale] the stamp takes the warning tone. After a failed Battle.net refresh
 * ([syncFailed], the build's `sync_error`) the stamp reads it in wipe red with Try again. From font
 * scale 1.5 the action takes its own 48 dp row (a4).
 */
@Composable
private fun BattleNetStamp(updated: String, stale: Boolean, syncFailed: Boolean, refresh: MenuRefresh, onRefresh: () -> Unit) {
    val refreshing = refresh == MenuRefresh.Refreshing
    val stamp = if (syncFailed) stringResource(R.string.menu_syncError) else updated
    val stampColor = when {
        syncFailed -> FsColor.wipe
        stale -> FsColor.warn
        else -> FsColor.text
    }
    val action = stringResource(
        when {
            refreshing -> R.string.menu_refreshing
            syncFailed -> R.string.common_tryAgain
            else -> R.string.menu_refresh
        },
    )
    val actionColor = if (refreshing) FsColor.muted else FsColor.gold
    // Round 4 row 10: only the action takes the press, inside its own 48 dp target; the stamp is text.
    // Round 5 a2: while refreshing the press is gone with its ripple (no patch behind `Refreshing…`).
    val press = if (refreshing) Modifier.semantics { disabled() } else Modifier.clickable(role = Role.Button, onClick = onRefresh)
    val tap = Modifier
        .heightIn(min = FsTarget.min)
        .then(press)
        .semantics { liveRegion = LiveRegionMode.Polite }
        .wrapContentHeight(Alignment.CenterVertically)
    val actionText: @Composable (Modifier) -> Unit = { modifier ->
        Text(action, style = STAMP, color = actionColor, fontWeight = FsType.bodyStrong.fontWeight, modifier = modifier.then(tap))
    }
    if (LocalDensity.current.fontScale >= LARGE_TEXT_SCALE) {
        Text(stamp, style = STAMP, color = stampColor, modifier = Modifier.padding(top = FsSpace.space1))
        actionText(Modifier)
    } else {
        Row(Modifier.padding(top = FsSpace.space1), verticalAlignment = Alignment.CenterVertically) {
            Text(
                buildAnnotatedString {
                    withStyle(SpanStyle(color = stampColor)) { append(stamp) }
                    append(STAMP_DOT)
                },
                style = STAMP,
                color = FsColor.text,
                modifier = Modifier.weight(1f, fill = false),
            )
            actionText(Modifier)
        }
    }
    if (refresh == MenuRefresh.Failed) Text(stringResource(R.string.menu_refreshFailed), style = STAMP, color = FsColor.wipe)
}

@Composable
private fun MainTag() {
    Text(
        stringResource(R.string.menu_main),
        style = TAG,
        color = FsColor.gold,
        modifier = Modifier
            .padding(top = FsSpace.space1)
            .border(1.dp, FsColor.gold, RoundedCornerShape(FsRadius.control))
            .padding(horizontal = FsSpace.space2, vertical = 1.dp),
    )
}

@Composable
private fun SignedOutHead(onSignIn: () -> Unit) {
    Column(Modifier.fillMaxWidth().parityRegion(ParityRegion.HEAD)) {
        Column(Modifier.fillMaxWidth().parityRegion(ParityRegion.HEAD_TITLE)) {
            Text(stringResource(R.string.menu_out_title), style = NAME, color = FsColor.cream, modifier = Modifier.semantics { heading() })
            Text(
                stringResource(R.string.signin_body),
                style = FsType.small,
                color = FsColor.text,
                modifier = Modifier.padding(top = FsSpace.space1).fillMaxWidth().parityRegion(ParityRegion.HEAD_META),
            )
        }
        BrandButton(
            stringResource(R.string.signin_button),
            BrandButtonKind.Primary,
            onSignIn,
            Modifier.padding(top = FsSpace.space4, bottom = FsSpace.space1).parityItem(ParityItemKind.Button),
            height = FsTarget.min,
        )
    }
}

@Composable
private fun SwitchSection(model: MenuModel, lowLevelOpen: Boolean, onSwitch: (AccountCharacter) -> Unit) {
    var showLow by rememberSaveable { mutableStateOf(lowLevelOpen) }
    MenuSection(R.string.menu_switch) {
        val shown = model.switchTo + if (showLow) model.lowLevel else emptyList()
        shown.forEachIndexed { index, character -> CharacterRow(character, last = index == shown.lastIndex, onSwitch) }
        if (model.lowLevel.isNotEmpty()) {
            Text(
                if (showLow) stringResource(R.string.menu_hideLow) else stringResource(R.string.menu_showLow, model.lowLevel.size),
                style = FsType.bodyStrong,
                color = FsColor.gold,
                modifier = Modifier
                    .heightIn(min = FsTarget.min)
                    .wrapContentHeight(Alignment.CenterVertically)
                    .clickable(role = Role.Button) { showLow = !showLow },
            )
        }
    }
}

@Composable
private fun CharacterRow(character: AccountCharacter, last: Boolean, onSwitch: (AccountCharacter) -> Unit) {
    Row(
        Modifier
            .fillMaxWidth()
            .parityItem(ParityItemKind.Row)
            .clickable(role = Role.Button) { onSwitch(character) }
            .padding(vertical = ROW_PAD),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(FsSpace.space3),
    ) {
        CharacterPortrait(character.avatarUrl, character.faction, PORTRAIT, className = character.className)
        Column {
            Text(character.name, style = FsType.item, color = classColor(character.className))
            Text(switchLine(character), style = SUB, color = FsColor.muted)
        }
    }
    if (!last) FlowRule()
}

/** A section of the drawer: the cream display title (16, the mock's `.drawer h3`), then its rows. */
@Composable
private fun MenuSection(@StringRes title: Int, rows: @Composable () -> Unit) {
    val text = stringResource(title)
    Column(Modifier.parityRegion(ParityRegion.section(text))) {
        Text(
            text,
            style = SECTION,
            color = FsColor.cream,
            modifier = Modifier.semantics { heading() }.padding(top = SECTION_TOP, bottom = SECTION_BOTTOM),
        )
        rows()
    }
}

/** The tools as labels alone (owner ruling 2026-11-09): the ported site tools are secondary entry points. */
@Composable
private fun ToolRows(tools: List<MenuDestination>, onOpen: (MenuDestination) -> Unit) {
    tools.forEachIndexed { index, tool -> MenuRow(labelOf(tool), null, last = index == tools.lastIndex) { onOpen(tool) } }
}

@Composable
private fun MenuRow(@StringRes title: Int, line: String?, last: Boolean, onClick: () -> Unit) {
    Column(
        Modifier
            .fillMaxWidth()
            .heightIn(min = FsTarget.min)
            .parityItem(ParityItemKind.Row)
            .clickable(role = Role.Button, onClick = onClick)
            .padding(vertical = ROW_PAD),
        verticalArrangement = Arrangement.Center,
    ) {
        Text(stringResource(title), style = FsType.item, color = FsColor.strong)
        line?.let { Text(it, style = SUB, color = FsColor.muted) }
    }
    if (!last) FlowRule()
}

@Composable
private fun Footer(lines: MenuLines, actions: MenuActions) {
    Column(Modifier.padding(top = FsSpace.space6, bottom = FsSpace.space4)) {
        // The site footer's fan line and operator line (web/src/lib/footer-copy.ts), word for word.
        Text(stringResource(R.string.menu_legal_fan), style = FsType.small, color = FsColor.muted)
        Text(stringResource(R.string.menu_legal_operator, lines.year.toString()), style = FsType.small, color = FsColor.muted)
        Text(lines.version, style = FsType.small, color = FsColor.muted, modifier = Modifier.padding(top = FsSpace.space3))
        Row(verticalAlignment = Alignment.CenterVertically) {
            FooterLink(R.string.menu_help, actions.onHelp)
            Text(" · ", style = FsType.small, color = FsColor.muted)
            FooterLink(R.string.menu_privacy, actions.onPrivacy)
        }
    }
}

@Composable
private fun FooterLink(@StringRes text: Int, onClick: () -> Unit) {
    Text(
        stringResource(text),
        style = FsType.small.copy(fontWeight = FsType.bodyStrong.fontWeight),
        color = FsColor.gold,
        modifier = Modifier
            .heightIn(min = FsTarget.min)
            .wrapContentHeight(Alignment.CenterVertically)
            .clickable(role = Role.Button, onClick = onClick),
    )
}

/** The head's line (round 2): `Fury Warrior · 58 · Officer`, the guild rank in place of the guild. */
@Composable
private fun headLine(character: AccountCharacter): String {
    val rank = character.guild?.rank?.takeIf { it.isNotBlank() } ?: return characterLine(character)
    val kind = kindOf(character) ?: return characterLine(character)
    val level = character.level ?: return characterLine(character)
    return stringResource(R.string.menu_character_lineRank, kind, level, AccountCopy.words(rank))
}

/** A switch row's line: the character's line, then `item level 62` when Battle.net reports it. */
@Composable
private fun switchLine(character: AccountCharacter): String {
    val itemLevel = character.itemLevel?.let { stringResource(R.string.menu_character_itemLevel, it) }
    return listOfNotNull(characterLine(character), itemLevel).joinToString(LINE_DOT)
}

private fun kindOf(character: AccountCharacter): String? =
    listOfNotNull(character.spec, character.className).joinToString(" ") { AccountCopy.words(it) }.takeIf { it.isNotEmpty() }

/** `Holy Paladin · 60 · Olympus XXVII`: spec and class, level, guild, each when known. */
@Composable
private fun characterLine(character: AccountCharacter): String = listOfNotNull(
    kindOf(character),
    character.level?.toString(),
    character.guild?.name,
).joinToString(LINE_DOT)

@StringRes
private fun labelOf(destination: MenuDestination): Int = when (destination) {
    MenuDestination.Guides -> R.string.menu_guides
    MenuDestination.Tiers -> R.string.menu_tiers
    MenuDestination.PatchNotes -> R.string.menu_patchNotes
    MenuDestination.Sims -> R.string.menu_sims
    MenuDestination.Logs -> R.string.menu_logs
    MenuDestination.Planner -> R.string.menu_planner
    MenuDestination.Rankings -> R.string.menu_rankings
    MenuDestination.Notifications -> R.string.menu_notifications
    MenuDestination.Plan -> R.string.menu_plan
    MenuDestination.Devices -> R.string.menu_devices
    MenuDestination.OfflineData -> R.string.menu_offlineData
    MenuDestination.Settings -> R.string.menu_settings
}

@Composable
private fun accountLine(destination: MenuDestination, lines: MenuLines): String? = when (destination) {
    MenuDestination.Notifications -> lines.notifications
    MenuDestination.Plan -> lines.plan
    MenuDestination.Devices -> lines.devices?.let { count ->
        lines.newestDevice?.let { stringResource(R.string.menu_devices_lineNewest, count, it) } ?: stringResource(R.string.menu_devices_line, count)
    }
    MenuDestination.OfflineData -> lines.offlineData?.let { stringResource(R.string.menu_offlineData_line, it) }
    MenuDestination.Settings -> stringResource(R.string.menu_settings_line)
    else -> null
}

/** The Battle.net stamp's break before Refresh: plain spaces, so it can wrap there at large text. */
private const val STAMP_DOT = " · "

/** The line's separator, held to the word before it so a wrap never starts a line on a dot. */
private const val LINE_DOT = " · "

private val CREST = 56.dp
private val PORTRAIT = 40.dp
private val SIDE = 18.dp
private val BAND_TOP = 18.dp
private val ROW_PAD = 7.dp
private val SECTION_TOP = 14.dp
private val SECTION_BOTTOM = 6.dp

/** You's name: the display face at 26 (the mock's `.you h2`). */
private val NAME = FsType.drawerName

/** A section's title: the display face at 16 (`.drawer h3`). */
private val SECTION = FsType.drawerSection

/** The stamp under You (12, `Updated 2 h ago from Battle.net · Refresh`). */
private val STAMP = FsType.figureUnit

/** The Main tag (12 semibold). */
private val TAG = FsType.tag
