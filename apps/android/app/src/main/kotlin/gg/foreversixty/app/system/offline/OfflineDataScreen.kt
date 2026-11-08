package gg.foreversixty.app.system.offline

import androidx.activity.compose.BackHandler
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.selection.toggleable
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.AlertDialog
import androidx.compose.material3.Text
import androidx.compose.material3.TextButton
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.compose.ui.window.DialogProperties
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ScaledType
import gg.foreversixty.app.flows.FlowInsetHead
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowNotice
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.FlowStack
import gg.foreversixty.app.flows.FlowStackButton
import java.time.Instant
import java.time.ZoneId
import kotlinx.coroutines.launch

/** Everything the screen draws, so the screenshot tests draw every state without a model. */
data class OfflineDataState(
    val report: OfflineReport?,
    val keepClassesReady: Boolean,
    val wifiOnly: Boolean,
    val online: Boolean,
    val updating: Boolean,
    val note: UpdateNote?,
    val checkedAt: Instant?,
    val confirmingRemove: Boolean = false,
)

class OfflineDataActions(
    val back: () -> Unit = {},
    val keep: (Boolean) -> Unit = {},
    val wifi: (Boolean) -> Unit = {},
    val updateNow: () -> Unit = {},
    val askRemove: () -> Unit = {},
    val remove: () -> Unit = {},
    val cancelRemove: () -> Unit = {},
)

/** You › Offline data, wired to its model. */
@Composable
fun OfflineDataRoute(model: OfflineDataModel, back: () -> Unit) {
    val scope = rememberCoroutineScope()
    var confirming by remember { mutableStateOf(false) }
    BackHandler(onBack = back)
    LaunchedEffect(model) { model.load() }
    val state = OfflineDataState(
        model.report, model.keepClassesReady, model.wifiOnly, model.online, model.updating, model.note, model.checkedAt, confirming,
    )
    OfflineDataScreen(
        state,
        Instant.now(),
        OfflineDataActions(
            back = back,
            keep = { model.keepClassesReady(it) },
            wifi = { model.wifiOnly(it) },
            updateNow = { scope.launch { model.updateNow() } },
            askRemove = { confirming = true },
            remove = {
                confirming = false
                scope.launch { model.remove() }
            },
            cancelRemove = { confirming = false },
        ),
    )
}

/**
 * Offline data (mocks `offline-data`, `offline-data-empty`, `offline-strip`): offline, the strip;
 * On this phone with its size and a row per kind, or the empty line; Updates with the two switches;
 * Update now and Remove downloaded data; when it last checked.
 */
@Composable
fun OfflineDataScreen(state: OfflineDataState, now: Instant, actions: OfflineDataActions, zone: ZoneId = ZoneId.systemDefault()) {
    val copy = OfflineDataCopy(LocalResources.current, zone)
    FlowPage {
        FlowInsetHead(stringResource(R.string.offlineData_title))
        if (!state.online) FlowNotice(copy.offline(state.checkedAt, now))
        val report = state.report
        when {
            report == null -> Unit
            report.isEmpty -> FlowSection(stringResource(R.string.offlineData_sectionEmpty)) {
                FlowNote(stringResource(R.string.offlineData_emptyLine), Modifier.padding(top = FsSpace.space3))
            }
            else -> FlowSection(copy.section(report)) {
                report.rows.forEachIndexed { index, row ->
                    FlowRow(copy.title(row.category), copy.meta(row, now), last = index == report.rows.lastIndex)
                }
            }
        }
        FlowSection(stringResource(R.string.offlineData_updates)) {
            SwitchRow(stringResource(R.string.offlineData_keep), stringResource(R.string.offlineData_keepNote), state.keepClassesReady, actions.keep, last = false)
            SwitchRow(stringResource(R.string.offlineData_wifi), stringResource(R.string.offlineData_wifiNote), state.wifiOnly, actions.wifi, last = true)
        }
        UpdateButtons(state, report?.isEmpty != false, actions)
        FlowNote(copy.footer(state.checkedAt, now), Modifier.padding(top = FsSpace.space4))
    }
    if (state.confirmingRemove) RemoveDialog(actions)
}

/**
 * `Remove downloaded data?`: the sentence in the body face, 24 dp from the screen's sides, and at
 * large text the two buttons stacked at one height so neither label breaks (review round 1).
 */
@Composable
private fun RemoveDialog(actions: OfflineDataActions) {
    val large = LocalDensity.current.fontScale >= STACK_BUTTONS_SCALE
    val remove = @Composable { modifier: Modifier ->
        TextButton(actions.remove, modifier.heightIn(min = 48.dp)) {
            Text(stringResource(R.string.offlineData_removeAction), style = ScaledType.bodyStrong, color = FsColor.wipe, textAlign = TextAlign.Center)
        }
    }
    val cancel = @Composable { modifier: Modifier ->
        TextButton(actions.cancelRemove, modifier.heightIn(min = 48.dp)) {
            Text(stringResource(R.string.common_cancel), style = ScaledType.bodyStrong, color = FsColor.nav, textAlign = TextAlign.Center)
        }
    }
    AlertDialog(
        onDismissRequest = actions.cancelRemove,
        properties = DialogProperties(usePlatformDefaultWidth = false),
        modifier = Modifier.padding(horizontal = FsSpace.space6),
        containerColor = FsColor.raised,
        text = { Text(stringResource(R.string.offlineData_removeConfirm), style = ScaledType.body, color = FsColor.strong) },
        confirmButton = {
            if (large) {
                Column(Modifier.fillMaxWidth()) {
                    remove(Modifier.fillMaxWidth())
                    cancel(Modifier.fillMaxWidth())
                }
            } else {
                Row {
                    cancel(Modifier)
                    remove(Modifier)
                }
            }
        },
    )
}

private const val STACK_BUTTONS_SCALE = 1.3f

@Composable
private fun UpdateButtons(state: OfflineDataState, empty: Boolean, actions: OfflineDataActions) {
    FlowStack {
        val label = stringResource(if (state.updating) R.string.offlineData_updating else R.string.offlineData_updateNow)
        FlowStackButton(label, actions.updateNow, enabled = state.online && !state.updating)
        // Shown on both platforms; with nothing saved it stays, disabled (review round 1).
        FlowStackButton(stringResource(R.string.offlineData_remove), actions.askRemove, danger = true, enabled = !empty)
    }
    val reason = when {
        !state.online -> stringResource(R.string.offline_needsConnection, stringResource(R.string.offlineData_updateNow))
        state.note == UpdateNote.UpToDate -> stringResource(R.string.offlineData_upToDate)
        state.note == UpdateNote.Failed -> stringResource(R.string.offline_updateFailed)
        else -> null
    }
    reason?.let {
        FlowNote(
            it,
            Modifier.padding(top = FsSpace.space3).semantics { liveRegion = LiveRegionMode.Polite },
            if (state.note == UpdateNote.Failed) FsColor.wipe else FsColor.muted,
        )
    }
}

/**
 * A setting (the mock's Row with a Switch trail): the title and its line, then the label beside the
 * switch on its own 44 line under them; the whole row toggles.
 */
@Composable
private fun SwitchRow(title: String, note: String, checked: Boolean, onChange: (Boolean) -> Unit, last: Boolean) {
    Column(Modifier.fillMaxWidth().toggleable(value = checked, role = Role.Switch, onValueChange = onChange)) {
        FlowRow(title, note, last = last, trail = {
            Row(Modifier.heightIn(min = SWITCH_LINE), verticalAlignment = Alignment.CenterVertically) {
                Text(title, style = SWITCH_LABEL, color = FsColor.text)
                Box(Modifier.padding(start = FsSpace.space3)) { Toggle(checked) }
            }
        })
    }
}

/** The mock's switch (`.swt i`): a 40 by 22 track, gold when on, the knob at the end it points to. */
@Composable
private fun Toggle(on: Boolean) {
    val shape = RoundedCornerShape(TRACK_HEIGHT / 2)
    Box(
        Modifier
            .size(TRACK_WIDTH, TRACK_HEIGHT)
            .background(if (on) FsColor.gold else FsColor.lineSoft, shape)
            .border(1.dp, if (on) FsColor.gold else FsColor.line, shape)
            .padding(KNOB_INSET),
        contentAlignment = if (on) Alignment.CenterEnd else Alignment.CenterStart,
    ) {
        Box(Modifier.size(KNOB).background(if (on) FsColor.bg else FsColor.muted, CircleShape))
    }
}

private val SWITCH_LINE = 44.dp
private val TRACK_WIDTH = 40.dp
private val TRACK_HEIGHT = 22.dp
private val KNOB = 16.dp
private val KNOB_INSET = 3.dp
private val SWITCH_LABEL = FsType.small.copy(fontSize = 14.sp)
