package gg.foreversixty.app.menu

import androidx.annotation.StringRes
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.HorizontalDivider
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
import gg.foreversixty.app.R
import gg.foreversixty.app.account.AccountCopy
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.parityContent
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion
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
 * The profile menu's content (spec 2026-11-08 section 3, the X pattern): You, Switch character, Tools,
 * Account and the footer; signed out ([model] null), Sign in with Battle.net and the tools that work
 * signed out.
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
    Column(
        modifier
            .verticalScroll(rememberScrollState())
            .parityContent()
            .padding(horizontal = FsSpace.space4)
            .padding(top = FsSpace.space11, bottom = FsSpace.space6),
    ) {
        val you = model?.you
        if (model == null || you == null) {
            SignedOutHead(actions.onSignIn)
            MenuSection(R.string.menu_tools) { MenuDestination.signedOutTools.forEach { MenuRow(it, null, actions.onOpen) } }
        } else {
            YouHead(you, model, actions.onRefresh)
            if (model.switchTo.isNotEmpty() || model.lowLevel.isNotEmpty()) SwitchSection(model, lowLevelOpen, actions.onSwitch)
            MenuSection(R.string.menu_tools) { MenuDestination.tools.forEach { MenuRow(it, null, actions.onOpen) } }
            MenuSection(R.string.menu_account) { MenuDestination.account.forEach { MenuRow(it, accountLine(it, lines), actions.onOpen) } }
        }
        Footer(lines.version, actions)
    }
}

@Composable
private fun YouHead(you: AccountCharacter, model: MenuModel, onRefresh: () -> Unit) {
    Column(Modifier.padding(horizontal = FsSpace.space4).parityRegion(ParityRegion.HEAD).padding(top = FsSpace.space4, bottom = HEAD_BOTTOM)) {
        Row(Modifier.fillMaxWidth().parityRegion(ParityRegion.HEAD_TITLE), horizontalArrangement = Arrangement.spacedBy(FsSpace.space4)) {
            CharacterPortrait(you.avatarUrl, you.faction, CREST)
            Column(Modifier.weight(1f)) {
                Text(you.name, style = FsType.title, color = classColor(you.className), modifier = Modifier.padding(bottom = NAME_GAP))
                Column(Modifier.fillMaxWidth().parityRegion(ParityRegion.HEAD_META)) {
                    Text(characterLine(you), style = FsType.small, color = FsColor.muted)
                    if (model.isMain) MainTag()
                }
            }
        }
    }
    model.updated?.let { age ->
        Text(
            buildAnnotatedString {
                append(stringResource(R.string.menu_updated, AccountCopy.age(age)))
                append(" · ")
                withStyle(SpanStyle(color = FsColor.gold)) { append(stringResource(R.string.menu_refresh)) }
            },
            style = FsType.small.copy(fontWeight = FsType.bodyStrong.fontWeight),
            color = FsColor.muted,
            modifier = Modifier
                .padding(horizontal = FsSpace.space4)
                .clickable(role = Role.Button, onClick = onRefresh),
        )
    }
    Spacer(Modifier.height(FsSpace.space1))
}

@Composable
private fun MainTag() {
    Text(
        stringResource(R.string.menu_main),
        style = FsType.small.copy(fontWeight = FsType.bodyStrong.fontWeight),
        color = FsColor.gold,
        modifier = Modifier
            .padding(top = FsSpace.space1)
            .border(1.dp, FsColor.gold, RoundedCornerShape(3.dp))
            .padding(horizontal = 6.dp, vertical = 1.dp),
    )
}

@Composable
private fun SignedOutHead(onSignIn: () -> Unit) {
    Column(Modifier.padding(horizontal = FsSpace.space4).parityRegion(ParityRegion.HEAD).padding(top = FsSpace.space4, bottom = HEAD_BOTTOM)) {
        Column(Modifier.fillMaxWidth().parityRegion(ParityRegion.HEAD_TITLE)) {
            Text(stringResource(R.string.menu_out_title), style = FsType.title, color = FsColor.strong, modifier = Modifier.padding(bottom = NAME_GAP))
            Text(stringResource(R.string.signin_body), style = FsType.small, color = FsColor.muted, modifier = Modifier.fillMaxWidth().parityRegion(ParityRegion.HEAD_META))
        }
    }
    Spacer(Modifier.height(FsSpace.space5))
    BrandButton(stringResource(R.string.signin_button), BrandButtonKind.Primary, onSignIn, height = 44.dp)
}

@Composable
private fun SwitchSection(model: MenuModel, lowLevelOpen: Boolean, onSwitch: (AccountCharacter) -> Unit) {
    var showLow by rememberSaveable { mutableStateOf(lowLevelOpen) }
    MenuSection(R.string.menu_switch) {
        (model.switchTo + if (showLow) model.lowLevel else emptyList()).forEach { CharacterRow(it, onSwitch) }
        if (model.lowLevel.isNotEmpty()) {
            Text(
                if (showLow) stringResource(R.string.menu_hideLow) else stringResource(R.string.menu_showLow, model.lowLevel.size),
                style = FsType.bodyStrong,
                color = FsColor.gold,
                modifier = Modifier
                    .padding(top = FsSpace.space3)
                    .heightIn(min = 32.dp)
                    .clickable(role = Role.Button) { showLow = !showLow },
            )
        }
    }
}

@Composable
private fun CharacterRow(character: AccountCharacter, onSwitch: (AccountCharacter) -> Unit) {
    Row(
        Modifier
            .fillMaxWidth()
            .heightIn(min = CHARACTER_ROW_HEIGHT)
            .parityItem(ParityItemKind.Row)
            .clickable(role = Role.Button) { onSwitch(character) }
            .padding(vertical = FsSpace.space2),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(FsSpace.space3),
    ) {
        CharacterPortrait(character.avatarUrl, character.faction, PORTRAIT)
        Column {
            Text(character.name, style = FsType.item, color = classColor(character.className))
            Text(characterLine(character), style = FsType.small, color = FsColor.muted)
        }
    }
    HorizontalDivider(color = FsColor.line)
}

@Composable
private fun MenuSection(@StringRes title: Int, rows: @Composable () -> Unit) {
    val text = stringResource(title)
    Column(Modifier.padding(top = FsSpace.space6).parityRegion(ParityRegion.section(text))) {
        Text(text, style = FsType.section, color = FsColor.strong, modifier = Modifier.semantics { heading() }.padding(bottom = FsSpace.space3))
        HorizontalDivider(color = FsColor.line)
        rows()
    }
}

@Composable
private fun MenuRow(destination: MenuDestination, line: String?, onOpen: (MenuDestination) -> Unit) {
    val (title, fixedLine) = labels(destination)
    Column(
        Modifier
            .fillMaxWidth()
            .heightIn(min = ROW_HEIGHT)
            .parityItem(ParityItemKind.Row)
            .clickable(role = Role.Button) { onOpen(destination) }
            .padding(vertical = FsSpace.space2),
        verticalArrangement = Arrangement.Center,
    ) {
        Text(stringResource(title), style = FsType.item, color = FsColor.strong)
        (line ?: fixedLine?.let { stringResource(it) })?.let { Text(it, style = FsType.small, color = FsColor.muted) }
    }
    HorizontalDivider(color = FsColor.line)
}

@Composable
private fun Footer(version: String, actions: MenuActions) {
    Column(Modifier.padding(top = FsSpace.space6)) {
        Text(version, style = FsType.small, color = FsColor.muted)
        Row(Modifier.padding(top = FsSpace.space2)) {
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

private fun labels(destination: MenuDestination): Pair<Int, Int?> = when (destination) {
    MenuDestination.Guides -> R.string.menu_guides to R.string.menu_guides_line
    MenuDestination.Tiers -> R.string.menu_tiers to R.string.menu_tiers_line
    MenuDestination.PatchNotes -> R.string.menu_patchNotes to R.string.menu_patchNotes_line
    MenuDestination.Sims -> R.string.menu_sims to R.string.menu_sims_line
    MenuDestination.Logs -> R.string.menu_logs to R.string.menu_logs_line
    MenuDestination.Planner -> R.string.menu_planner to R.string.menu_planner_line
    MenuDestination.Rankings -> R.string.menu_rankings to R.string.menu_rankings_line
    MenuDestination.Notifications -> R.string.menu_notifications to null
    MenuDestination.Plan -> R.string.menu_plan to null
    MenuDestination.Devices -> R.string.menu_devices to null
    MenuDestination.OfflineData -> R.string.menu_offlineData to null
    MenuDestination.Settings -> R.string.menu_settings to R.string.menu_settings_line
}

@Composable
private fun accountLine(destination: MenuDestination, lines: MenuLines): String? = when (destination) {
    MenuDestination.Notifications -> lines.notifications
    MenuDestination.Plan -> lines.plan
    MenuDestination.Devices -> lines.devices?.let { stringResource(R.string.menu_devices_line, it) }
    MenuDestination.OfflineData -> lines.offlineData?.let { stringResource(R.string.menu_offlineData_line, it) }
    else -> null
}

private val CREST = 56.dp
private val PORTRAIT = 40.dp
private val ROW_HEIGHT = 58.dp
private val CHARACTER_ROW_HEIGHT = 60.dp
private val NAME_GAP = 6.dp
private val HEAD_BOTTOM = 14.dp
