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
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
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
    val offlineData: String? = null,
    val version: String = "",
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
) {
    Column(modifier.verticalScroll(rememberScrollState()).parityContent()) {
        val you = model?.you
        if (model == null || you == null) {
            Band { SignedOutHead(actions.onSignIn) }
            Body { MenuSection(R.string.menu_tools) { ToolRows(MenuDestination.signedOutTools, actions.onOpen) } }
        } else {
            Band { YouHead(you, model, actions.onRefresh) }
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
        Body { Footer(lines.version, actions) }
    }
}

/** The painted band at the drawer's top (the mock's `.you`), under the status bar. */
@Composable
private fun Band(content: @Composable () -> Unit) {
    Column(
        Modifier
            .fillMaxWidth()
            .surfaceBand()
            .windowInsetsPadding(WindowInsets.statusBars)
            .padding(start = SIDE, end = SIDE, top = BAND_TOP, bottom = FsSpace.space3),
    ) { content() }
}

/** The drawer's panel under the band, 18 at the sides. */
@Composable
private fun Body(content: @Composable () -> Unit) {
    Column(Modifier.fillMaxWidth().padding(horizontal = SIDE)) { content() }
}

@Composable
private fun YouHead(you: AccountCharacter, model: MenuModel, onRefresh: () -> Unit) {
    Column(Modifier.fillMaxWidth().parityRegion(ParityRegion.HEAD)) {
        Row(
            Modifier.fillMaxWidth().parityRegion(ParityRegion.HEAD_TITLE),
            verticalAlignment = Alignment.CenterVertically,
            horizontalArrangement = Arrangement.spacedBy(FsSpace.space4),
        ) {
            CharacterPortrait(you.avatarUrl, you.faction, CREST, className = you.className)
            Column(Modifier.weight(1f)) {
                Text(you.name, style = NAME, color = classColor(you.className), modifier = Modifier.semantics { heading() })
                Column(Modifier.fillMaxWidth().parityRegion(ParityRegion.HEAD_META)) {
                    Text(characterLine(you), style = FsType.small, color = FsColor.text)
                    if (model.isMain) MainTag()
                }
            }
        }
        model.updated?.let { age ->
            Text(
                buildAnnotatedString {
                    append(stringResource(R.string.menu_updated, AccountCopy.age(age)))
                    append(" · ")
                    withStyle(SpanStyle(color = FsColor.gold, fontWeight = FsType.bodyStrong.fontWeight)) { append(stringResource(R.string.menu_refresh)) }
                },
                style = STAMP,
                color = FsColor.text,
                modifier = Modifier
                    .padding(top = FsSpace.space1)
                    .heightIn(min = FsTarget.min)
                    .wrapContentHeight(Alignment.CenterVertically)
                    .clickable(role = Role.Button, onClick = onRefresh),
            )
        }
    }
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
        CharacterPortrait(character.avatarUrl, character.faction, PORTRAIT)
        Column {
            Text(character.name, style = FsType.item, color = classColor(character.className))
            Text(characterLine(character), style = SUB, color = FsColor.muted)
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
private fun Footer(version: String, actions: MenuActions) {
    Column(Modifier.padding(top = FsSpace.space6, bottom = FsSpace.space6)) {
        Text(version, style = FsType.small, color = FsColor.muted)
        Row(Modifier.padding(top = FsSpace.space1)) {
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
        modifier = Modifier.clickable(role = Role.Button, onClick = onClick),
    )
}

/** `Holy Paladin · 60 · Olympus XXVII`: spec and class, level, guild, each when known. */
@Composable
private fun characterLine(character: AccountCharacter): String = listOfNotNull(
    listOfNotNull(character.spec, character.className).joinToString(" ") { AccountCopy.words(it) }.takeIf { it.isNotEmpty() },
    character.level?.toString(),
    character.guild?.name,
).joinToString(" · ")

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
    MenuDestination.Devices -> lines.devices?.let { stringResource(R.string.menu_devices_line, it) }
    MenuDestination.OfflineData -> lines.offlineData?.let { stringResource(R.string.menu_offlineData_line, it) }
    MenuDestination.Settings -> stringResource(R.string.menu_settings_line)
    else -> null
}

private val CREST = 56.dp
private val PORTRAIT = 40.dp
private val SIDE = 18.dp
private val BAND_TOP = 18.dp
private val ROW_PAD = 7.dp
private val SECTION_TOP = 14.dp
private val SECTION_BOTTOM = 6.dp

/** You's name: the display face at 26 (the mock's `.you h2`). */
private val NAME = FsType.title.copy(fontSize = 26.sp, lineHeight = 32.sp, letterSpacing = 0.sp)

/** A section's title: the display face at 16 (`.drawer h3`). */
private val SECTION = FsType.section.copy(fontSize = 16.sp, lineHeight = 22.sp)

/** The stamp under You (12, `Updated 2 h ago from Battle.net · Refresh`). */
private val STAMP = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp)

/** The Main tag (12 semibold). */
private val TAG = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FsType.bodyStrong.fontWeight)
