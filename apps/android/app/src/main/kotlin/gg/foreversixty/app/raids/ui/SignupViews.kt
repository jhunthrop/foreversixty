package gg.foreversixty.app.raids.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.IntrinsicSize
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.RowScope
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.selection.selectable
import androidx.compose.foundation.selection.selectableGroup
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.text.BasicTextField
import androidx.compose.foundation.text.KeyboardActions
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.KeyboardArrowRight
import androidx.compose.material.icons.filled.Check
import androidx.compose.material.icons.filled.Notifications
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.Text
import androidx.compose.material3.rememberModalBottomSheetState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.alpha
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.focus.onFocusChanged
import androidx.compose.ui.geometry.CornerRadius
import androidx.compose.ui.graphics.PathEffect
import androidx.compose.ui.graphics.SolidColor
import androidx.compose.ui.graphics.drawscope.Stroke
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.semantics.stateDescription
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.input.ImeAction
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.model.RaidSpecOption
import gg.foreversixty.app.model.RaidViewerCharacter
import gg.foreversixty.app.raids.logic.SignupBlock
import gg.foreversixty.app.raids.logic.SignupController
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.logic.SpecAsk
import gg.foreversixty.app.raids.logic.StatusLine
import gg.foreversixty.app.raids.logic.roleLabel
import kotlinx.coroutines.launch

/** The font scale from which the segments go 2 x 2 (spec 10, XL.2). */
private const val GRID_FONT_SCALE = 1.5f

/** The four segments (spec F4.3): In, Late, Bench, Out, in a selectable group; 2 x 2 from 1.5x. */
@Composable
fun Segments(characterName: String, block: SignupBlock, enabled: Boolean, press: (SignupState) -> Unit) {
    val grid = LocalDensity.current.fontScale >= GRID_FONT_SCALE
    val label = stringResource(R.string.raids_a11y_segments, characterName)
    val shape = RoundedCornerShape(FsRadius.control)
    Column(
        Modifier
            .fillMaxWidth()
            .alpha(if (enabled) 1f else 0.5f)
            .background(FsColor.raised, shape)
            .border(1.dp, FsColor.line, shape)
            .semantics { contentDescription = label }
            .selectableGroup(),
    ) {
        val rows = if (grid) SignupState.entries.chunked(2) else listOf(SignupState.entries)
        rows.forEachIndexed { index, row ->
            if (index > 0) HorizontalDivider(color = FsColor.line)
            Row(Modifier.fillMaxWidth().height(IntrinsicSize.Min)) {
                row.forEachIndexed { i, state ->
                    if (i > 0) Box(Modifier.width(1.dp).fillMaxHeight().background(FsColor.line))
                    Segment(state, block, enabled, press)
                }
            }
        }
    }
}

@Composable
private fun RowScope.Segment(state: SignupState, block: SignupBlock, enabled: Boolean, press: (SignupState) -> Unit) {
    val checked = block.shown == state
    val pending = block.pending == state
    val saving = stringResource(R.string.raids_a11y_saving)
    val pendingText = stringResource(R.string.raids_a11y_pending)
    val ground = when {
        !checked -> FsColor.raised
        state.isComing -> FsColor.kill.copy(alpha = 0.15f)
        else -> FsColor.lineWarm
    }
    Box(
        contentAlignment = Alignment.Center,
        modifier = Modifier
            .weight(1f)
            .heightIn(min = 48.dp)
            .background(ground)
            .selectable(selected = checked, enabled = enabled, role = Role.RadioButton) { press(state) }
            .semantics {
                when {
                    checked && block.busy -> stateDescription = saving
                    pending -> stateDescription = pendingText
                }
            }
            .drawBehind {
                if (checked) {
                    val line = 2.dp.toPx()
                    drawRect(
                        color = if (state.isComing) FsColor.kill else FsColor.muted,
                        topLeft = androidx.compose.ui.geometry.Offset(0f, size.height - line),
                        size = androidx.compose.ui.geometry.Size(size.width, line),
                    )
                }
                if (pending) {
                    val inset = 4.dp.toPx()
                    drawRoundRect(
                        color = FsColor.gold,
                        topLeft = androidx.compose.ui.geometry.Offset(inset, inset),
                        size = androidx.compose.ui.geometry.Size(size.width - inset * 2, size.height - inset * 2),
                        cornerRadius = CornerRadius(6.dp.toPx()),
                        style = Stroke(width = 1.5.dp.toPx(), pathEffect = PathEffect.dashPathEffect(floatArrayOf(10f, 7f))),
                    )
                }
            },
    ) {
        Text(stringResource(state.label), style = RaidType.segment, color = if (checked) FsColor.strong else FsColor.muted)
    }
}

/** The 20 dp status line, always present so nothing shifts; announced politely when it changes. */
@Composable
fun StatusLineText(status: StatusLine?) {
    val color = when (status) {
        is StatusLine.Saved -> FsColor.kill
        is StatusLine.Error -> FsColor.wipe
        is StatusLine.Info, null -> FsColor.muted
    }
    Text(
        status?.text ?: " ",
        style = RaidType.meta,
        color = color,
        modifier = Modifier.fillMaxWidth().heightIn(min = 20.dp).semantics { liveRegion = LiveRegionMode.Polite },
    )
}

/** The who line: crest 36, the name in class colour, the spec icon 18 with `{Spec} · {Role}`. */
@Composable
fun WhoLine(
    name: String,
    className: String,
    specName: String?,
    specIcon: String?,
    role: String,
    trailing: (@Composable () -> Unit)? = null,
) {
    Row(horizontalArrangement = Arrangement.spacedBy(12.dp), verticalAlignment = Alignment.CenterVertically) {
        Crest(className, 36.dp)
        Column(Modifier.weight(1f).semantics(mergeDescendants = true) {}) {
            Text(name, style = RaidType.name, color = RaidColors.classColor(className))
            Row(horizontalArrangement = Arrangement.spacedBy(6.dp), verticalAlignment = Alignment.CenterVertically) {
                GameIcon(specIcon, 18.dp)
                val roleText = if (role.isEmpty()) "" else stringResource(roleLabel(role))
                Text(
                    specName?.let { stringResource(R.string.raids_who, it, roleText) } ?: roleText,
                    style = RaidType.meta,
                    color = FsColor.muted,
                )
            }
        }
        trailing?.invoke()
    }
}

/** One character's block on the sign-up card (spec F4.2 item 5). */
@Composable
fun SignupBlockView(controller: SignupController, character: RaidViewerCharacter, showsNote: Boolean = true) {
    val scope = rememberCoroutineScope()
    val key = character.characterKey
    val block = controller.block(key)
    Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
        WhoLine(
            name = character.name,
            className = character.className,
            specName = character.specName,
            specIcon = character.specIcon,
            role = character.role,
            trailing = if (character.raidSpec != null && controller.segmentsEnabled(key)) {
                { GoldLink(stringResource(R.string.spec_change), { controller.changeSpec(key) }) }
            } else {
                null
            },
        )
        if (character.verified) {
            Segments(character.name, block, controller.segmentsEnabled(key)) { state ->
                scope.launch { controller.press(key, state) }
            }
            StatusLineText(controller.restingStatus(key))
            if (showsNote && block.shown != null) NoteField(controller, key, block)
        } else {
            Text(stringResource(R.string.raid_waitingCharacter), style = RaidType.meta, color = FsColor.muted)
        }
    }
}

@Composable
private fun NoteField(controller: SignupController, key: String, block: SignupBlock) {
    val scope = rememberCoroutineScope()
    var focused by remember { mutableStateOf(false) }
    val text = controller.notes[key].orEmpty()
    val enabled = controller.segmentsEnabled(key)
    val borderColor = when {
        block.noteFailed -> FsColor.wipe
        focused -> FsColor.gold
        else -> FsColor.lineCard
    }
    Column(verticalArrangement = Arrangement.spacedBy(4.dp)) {
        BasicTextField(
            value = text,
            onValueChange = { controller.editNote(key, it) },
            enabled = enabled,
            singleLine = true,
            textStyle = RaidType.body.copy(color = FsColor.text),
            cursorBrush = SolidColor(FsColor.gold),
            keyboardOptions = KeyboardOptions(imeAction = ImeAction.Done),
            keyboardActions = KeyboardActions(onDone = { scope.launch { controller.saveNote(key) } }),
            modifier = Modifier
                .fillMaxWidth()
                .onFocusChanged { state ->
                    if (focused && !state.isFocused) scope.launch { controller.saveNote(key) }
                    focused = state.isFocused
                },
            decorationBox = { inner ->
                Box(
                    contentAlignment = Alignment.CenterStart,
                    modifier = Modifier
                        .fillMaxWidth()
                        .heightIn(min = 48.dp)
                        .background(androidx.compose.ui.graphics.Color.Black.copy(alpha = 0.35f), RoundedCornerShape(FsRadius.control))
                        .border(1.dp, borderColor, RoundedCornerShape(FsRadius.control))
                        .padding(horizontal = 14.dp),
                ) {
                    if (text.isEmpty()) Text(stringResource(R.string.raid_note_placeholder), style = RaidType.body, color = FsColor.muted)
                    inner()
                }
            },
        )
        when {
            block.noteBusy -> Text(stringResource(R.string.raids_note_saving), style = RaidType.meta, color = FsColor.muted)
            block.noteFailed -> Text(stringResource(R.string.raid_note_failed), style = RaidType.meta, color = FsColor.wipe)
            text.length >= SignupController.NOTE_COUNTER_FROM -> Text(
                stringResource(R.string.raid_note_left, SignupController.NOTE_LIMIT - text.length),
                style = RaidType.monoMeta,
                color = FsColor.muted,
            )
        }
    }
}

/** The reminder ask (spec F4.4), inset in the sign-up card after the first In or Late. */
@Composable
fun ReminderAsk(turnOn: () -> Unit, notNow: () -> Unit) {
    Row(horizontalArrangement = Arrangement.spacedBy(12.dp), modifier = Modifier.quietCard().padding(14.dp)) {
        Box(
            contentAlignment = Alignment.Center,
            modifier = Modifier.size(32.dp).background(FsColor.gold.copy(alpha = 0.12f), RoundedCornerShape(FsRadius.control)),
        ) {
            Icon(Icons.Filled.Notifications, contentDescription = null, tint = FsColor.gold, modifier = Modifier.size(18.dp))
        }
        Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
            Text(stringResource(R.string.ask_raid_title), style = RaidType.bodyStrong, color = FsColor.strong)
            Text(stringResource(R.string.ask_raid_body), style = RaidType.meta, color = FsColor.muted)
            Row(horizontalArrangement = Arrangement.spacedBy(12.dp), verticalAlignment = Alignment.CenterVertically) {
                SecondaryButton(stringResource(R.string.ask_turnOn_raid), turnOn, gold = true)
                Text(
                    stringResource(R.string.ask_notNow).uppercase(),
                    style = RaidType.button,
                    color = FsColor.nav,
                    modifier = Modifier.heightIn(min = 48.dp).clickable(onClick = notNow).padding(vertical = 15.dp),
                )
            }
        }
    }
}

/** "Which spec tonight?" (spec F4.3), a Material bottom sheet. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun SpecSheet(ask: SpecAsk, character: RaidViewerCharacter, raidTitle: String, date: String, pick: (RaidSpecOption) -> Unit, close: () -> Unit) {
    ModalBottomSheet(
        onDismissRequest = close,
        sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = true),
        containerColor = RaidColors.sheet,
    ) {
        SpecSheetContent(ask, character, raidTitle, date, pick)
    }
}

/** The sheet's body, apart from the sheet so screenshot tests draw it. */
@Composable
fun SpecSheetContent(ask: SpecAsk, character: RaidViewerCharacter, raidTitle: String, date: String, pick: (RaidSpecOption) -> Unit) {
    val sentence = when (ask) {
        is SpecAsk.Ask -> stringResource(R.string.spec_lead, character.name, stringResource(ask.state.label), raidTitle, date)
        is SpecAsk.Change -> stringResource(R.string.spec_lead_change, character.name, raidTitle, date)
    }
    val lead = buildAnnotatedString {
        val at = sentence.indexOf(character.name)
        if (at < 0) {
            append(sentence)
        } else {
            append(sentence.substring(0, at))
            withStyle(SpanStyle(color = FsColor.strong, fontWeight = RaidType.bodyStrong.fontWeight)) { append(character.name) }
            append(sentence.substring(at + character.name.length))
        }
    }
    Column(verticalArrangement = Arrangement.spacedBy(14.dp), modifier = Modifier.padding(16.dp)) {
        Text(stringResource(R.string.spec_title).uppercase(), style = RaidType.sec, color = FsColor.strong, modifier = Modifier.semantics { heading() })
        Text(lead, style = RaidType.body, color = FsColor.text)
        Column(Modifier.quietCard()) {
            val options = character.specOptions.orEmpty()
            options.forEachIndexed { index, option ->
                if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
                val checked = ask is SpecAsk.Change && character.raidSpec == option.spec
                Row(
                    horizontalArrangement = Arrangement.spacedBy(12.dp),
                    verticalAlignment = Alignment.CenterVertically,
                    modifier = Modifier
                        .fillMaxWidth()
                        .heightIn(min = 64.dp)
                        .selectable(selected = checked, role = Role.Button) { pick(option) }
                        .padding(horizontal = 14.dp),
                ) {
                    GameIcon(option.icon, 40.dp)
                    Column(Modifier.weight(1f)) {
                        Text(option.name, style = RaidType.bodyStrong, color = FsColor.strong)
                        Text(stringResource(roleLabel(option.role)), style = RaidType.meta, color = FsColor.muted)
                    }
                    Icon(
                        if (checked) Icons.Filled.Check else Icons.AutoMirrored.Filled.KeyboardArrowRight,
                        contentDescription = null,
                        tint = if (checked) FsColor.kill else FsColor.muted,
                    )
                }
            }
        }
        Text(stringResource(R.string.spec_closeNote), style = RaidType.meta, color = FsColor.muted)
        Spacer(Modifier.height(8.dp))
    }
}
