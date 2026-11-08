package gg.foreversixty.app.firstrun

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.widthIn
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.layout.layout
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.account.AccountCopy
import gg.foreversixty.app.common.Seal
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.flows.FlowButton
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.FlowTag
import gg.foreversixty.app.model.AccountCharacter

/**
 * F8 step 2 (mock android-first-pick-main): the account's characters at level 10 and up, the highest
 * preselected and tagged Main, and one button that continues with the selection.
 */
@Composable
fun PickMainScreen(
    characters: List<AccountCharacter>,
    selection: String?,
    saving: Boolean,
    failed: Boolean,
    onSelect: (String) -> Unit,
    onContinue: () -> Unit,
) {
    FlowPage {
        FlowHead(stringResource(R.string.first_pickMain_title), listOf(stringResource(R.string.first_pickMain_meta)))
        FlowNote(stringResource(R.string.first_pickMain_lead), Modifier.padding(top = FsSpace.space2))
        FlowSection(stringResource(R.string.first_pickMain_section)) {
            characters.forEachIndexed { index, character ->
                val picked = character.key == selection
                FlowRow(
                    character.name,
                    characterLine(character),
                    FlowLead.Crest(character.className.orEmpty()),
                    titleColor = classColor(character.className),
                    last = index == characters.lastIndex,
                    selected = picked,
                    trail = if (picked) ({ FlowTag(stringResource(R.string.menu_character_main)) }) else null,
                    onClick = { onSelect(character.key) },
                )
            }
        }
        characters.firstOrNull { it.key == selection }?.let { chosen ->
            FlowButton(
                stringResource(R.string.first_pickMain_continue, chosen.name),
                onContinue,
                Modifier.padding(top = FsSpace.space5).fillMaxWidth(),
                primary = true,
                small = false,
                busy = saving,
            )
        }
        if (failed) FlowNote(stringResource(R.string.first_pickMain_failed), Modifier.padding(top = FsSpace.space3), color = FsColor.wipe)
    }
}

/**
 * F8 step 3 (mocks android-first-permission and -long): `You're in` over the guild's name, what the
 * notification does, and the ask. The system's Allow follows Turn on notifications.
 */
@Composable
fun FirstPermissionScreen(guild: String, onTurnOn: () -> Unit, onNotNow: () -> Unit) {
    FlowPage {
        Column(Modifier.fillMaxWidth(), horizontalAlignment = Alignment.CenterHorizontally) {
            Column(Modifier.padding(top = SEAL_TOP, bottom = FsSpace.space2)) { Seal(SEAL) }
            Text(
                stringResource(R.string.first_permission_eyebrow),
                style = FsType.small,
                color = FsColor.muted,
                textAlign = TextAlign.Center,
                modifier = Modifier.pullUp(EYEBROW_PULL),
            )
            FlowHead(guild, emptyList(), balanced = true)
            Text(
                stringResource(R.string.first_permission_body),
                style = FsType.small,
                color = FsColor.muted,
                textAlign = TextAlign.Center,
                modifier = Modifier.padding(top = FsSpace.space2).widthIn(max = MEASURE),
            )
        }
        FlowButton(stringResource(R.string.first_permission_turnOn), onTurnOn, Modifier.padding(top = FsSpace.space5).fillMaxWidth(), primary = true, small = false)
        FlowButton(stringResource(R.string.first_permission_notNow), onNotNow, Modifier.padding(top = FsSpace.space3).fillMaxWidth(), small = false)
    }
}

/** `Fury Warrior · 60 · Olympus XXVII`, as the menu's character rows read. */
@Composable
private fun characterLine(character: AccountCharacter): String {
    val kind = listOfNotNull(character.spec, character.className).joinToString(" ") { AccountCopy.words(it) }
    val level = character.level ?: 0
    return character.guild?.name?.let { stringResource(R.string.menu_character_lineGuild, kind, level, it) }
        ?: stringResource(R.string.menu_character_line, kind, level)
}

/** `.center .eyebrow { margin-bottom: -12px }`: what follows sits [by] higher; the text still draws whole. */
private fun Modifier.pullUp(by: androidx.compose.ui.unit.Dp): Modifier = layout { measurable, constraints ->
    val placeable = measurable.measure(constraints)
    layout(placeable.width, (placeable.height - by.roundToPx()).coerceAtLeast(0)) { placeable.place(0, 0) }
}

private val SEAL = 84.dp
private val SEAL_TOP = 40.dp
private val EYEBROW_PULL = 12.dp
private val MEASURE = 300.dp
