package gg.foreversixty.app.account

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.selection.selectable
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Check
import androidx.compose.material.icons.filled.KeyboardArrowDown
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.Text
import androidx.compose.material3.rememberModalBottomSheetState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.hapticfeedback.HapticFeedbackType
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.platform.LocalHapticFeedback
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.CustomAccessibilityAction
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.customActions
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.style.Hyphens
import androidx.compose.ui.text.style.LineBreak
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.SessionManager
import gg.foreversixty.app.auth.session.me
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpacing
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.Me

/** Font scale from which the selector drops its short line (Android 1.3x, spec §3.4). */
private const val DROP_DETAIL_SCALE = 1.3f

/**
 * The selector (spec §3.4): crest 34, the full name in the class colour, the short line and a
 * chevron. The one way to switch; Home, Raids and BiS put it in their TopAppBar title slot. The
 * name never truncates; it stops growing at 22 sp and the short line drops from 1.3x.
 */
@Composable
fun CharacterSelector(character: AccountCharacter, open: Boolean, onClick: () -> Unit) {
    val shape = RoundedCornerShape(50)
    val label = AccountCopy.selectorLabel(character)
    val hint = stringResource(R.string.account_switcher_open)
    val action = stringResource(R.string.account_switcher_action)
    val compact = LocalDensity.current.fontScale >= DROP_DETAIL_SCALE
    Row(
        Modifier
            .shadow(if (open) 8.dp else 0.dp, shape, ambientColor = FsColor.gold, spotColor = FsColor.gold)
            .background(FsColor.raised, shape)
            .border(1.dp, if (open) FsColor.gold else FsColor.gold.copy(alpha = 0.17f), shape)
            .clickable(role = Role.Button, onClickLabel = hint, onClick = onClick)
            .semantics {
                contentDescription = label
                customActions = listOf(CustomAccessibilityAction(action) { onClick(); true })
            }
            .padding(start = 4.dp, end = FsSpacing.space3, top = 4.dp, bottom = 4.dp),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(FsSpacing.space2),
    ) {
        ClassCrest(character.className, 34.dp)
        Column {
            Text(character.name, style = FsType.cappedName(), color = classColor(character.className), softWrap = false)
            val detail = AccountCopy.selectorDetail(character)
            if (!compact && detail != null) Text(detail, style = FsType.meta, color = FsColor.muted, maxLines = 1)
        }
        Icon(Icons.Filled.KeyboardArrowDown, contentDescription = null, tint = FsColor.muted)
    }
}

/** F1.5's content: one 68 dp row per character, the current one washed and checked. */
@Composable
fun SwitcherSheetContent(me: Me, currentKey: String?, onPick: (AccountCharacter) -> Unit) {
    Column(Modifier.fillMaxWidth().verticalScroll(rememberScrollState())) {
        Column(Modifier.padding(FsSpacing.space4), verticalArrangement = Arrangement.spacedBy(FsSpacing.space1)) {
            Text(
                stringResource(R.string.switcher_title).uppercase(),
                style = FsType.sec,
                color = FsColor.strong,
                modifier = Modifier.semantics { heading() },
            )
            me.user.battletag?.let {
                Text(stringResource(R.string.switcher_who, it), style = FsType.meta, color = FsColor.muted)
            }
        }
        me.characters.forEach { character ->
            SwitcherRow(character, character.key == currentKey) { onPick(character) }
            HorizontalDivider(color = FsColor.lineSoft)
        }
        Text(
            stringResource(R.string.switcher_footer),
            style = FsType.meta,
            color = FsColor.muted,
            modifier = Modifier.padding(FsSpacing.space4),
        )
    }
}

@Composable
private fun SwitcherRow(character: AccountCharacter, current: Boolean, onClick: () -> Unit) {
    Row(
        Modifier
            .fillMaxWidth()
            .heightIn(min = 68.dp)
            .background(if (current) Color.White.copy(alpha = 0.04f) else Color.Transparent)
            .selectable(selected = current, role = Role.RadioButton, onClick = onClick)
            .padding(horizontal = FsSpacing.space4, vertical = FsSpacing.space2),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(FsSpacing.space3),
    ) {
        ClassCrest(character.className, 44.dp)
        Column(Modifier.weight(1f)) {
            // `Thrall · US · PvP`: same-named characters differ by region and ruleset.
            val place = AccountCopy.place(character)
            Text(
                buildAnnotatedString {
                    withStyle(SpanStyle(color = classColor(character.className))) { append(character.name) }
                    withStyle(FsType.meta.toSpanStyle().copy(color = FsColor.muted)) { append(" · $place") }
                },
                style = FsType.name.copy(lineBreak = LineBreak.Simple, hyphens = Hyphens.None),
            )
            Text(AccountCopy.switcherDetail(character), style = FsType.meta, color = FsColor.text)
            Text(character.guild?.name ?: stringResource(R.string.switcher_noGuild), style = FsType.meta, color = FsColor.muted)
        }
        if (current) Icon(Icons.Filled.Check, contentDescription = null, tint = FsColor.gold)
    }
}

/**
 * The selector with its sheet, for a root screen's TopAppBar title slot. Draws nothing when the
 * account has no characters (the sheet is never empty).
 */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun CharacterSelectorWithSheet(
    me: Me,
    current: AccountCharacter?,
    open: Boolean,
    onOpenChange: (Boolean) -> Unit,
    onPick: (AccountCharacter) -> Unit,
) {
    current ?: return
    val haptics = LocalHapticFeedback.current
    CharacterSelector(current, open) { onOpenChange(true) }
    if (open) {
        ModalBottomSheet(
            onDismissRequest = { onOpenChange(false) },
            sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = false),
            containerColor = FsColor.raised,
        ) {
            SwitcherSheetContent(me, current.key) {
                haptics.performHapticFeedback(HapticFeedbackType.SegmentTick)
                onPick(it)
                onOpenChange(false)
            }
        }
    }
}

/**
 * The one character switcher bound to the session, for the Home, Raids and BiS root screens
 * (spec §3.4). Draws nothing when signed out or with no characters.
 */
@Composable
fun CharacterSelectorBar(session: SessionManager, modifier: Modifier = Modifier) {
    val phase by session.phase.collectAsStateWithLifecycle()
    val picked by session.pickedKey.collectAsStateWithLifecycle()
    var open by rememberSaveable { mutableStateOf(false) }
    val me = phase.me ?: return
    Row(modifier.fillMaxWidth().padding(horizontal = FsSpacing.space4, vertical = FsSpacing.space2)) {
        CharacterSelectorWithSheet(me, session.currentCharacter(phase, picked), open, { open = it }) { session.pick(it.key) }
    }
}

