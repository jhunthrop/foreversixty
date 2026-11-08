package gg.foreversixty.app.flows

import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ColumnScope
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.widthIn
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.common.Seal
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType

/**
 * The account and system screens' pieces (design/mocks/ws_app.py `devices`, `delete`, `offline`,
 * `update_required`, `expired`): their head sits in the page container's own 16 inset (`.pc`), so
 * it starts at 32 where the sections start at 16.
 */
@Composable
fun FlowInsetHead(title: String) {
    Box(Modifier.fillMaxWidth().padding(horizontal = FsSpace.space4)) { FlowHead(title, emptyList()) }
}

/** A notice (`.notice`): 15/24, in the container's 16 inset, 14 above and 4 below; gold or plain. */
@Composable
fun FlowNotice(text: String, gold: Boolean = false, color: Color = if (gold) FsColor.gold else FsColor.text) {
    Text(
        text,
        style = NOTICE,
        color = color,
        modifier = Modifier
            .fillMaxWidth()
            .padding(start = FsSpace.space4, end = FsSpace.space4, top = NOTICE_TOP, bottom = FsSpace.space1)
            .semantics { liveRegion = LiveRegionMode.Polite },
    )
}

/** The page's buttons (`.stack`): full width, 48 tall, 12 apart, 20 under what precedes them. */
@Composable
fun FlowStack(content: @Composable ColumnScope.() -> Unit) {
    Column(Modifier.padding(top = FsSpace.space5).fillMaxWidth(), verticalArrangement = Arrangement.spacedBy(FsSpace.space3), content = content)
}

/** One of [FlowStack]'s buttons. */
@Composable
fun FlowStackButton(label: String, onClick: () -> Unit, primary: Boolean = false, danger: Boolean = false, enabled: Boolean = true) {
    FlowButton(label, onClick, Modifier.fillMaxWidth().height(STACK_BUTTON), primary = primary, small = false, danger = danger, enabled = enabled)
}

/**
 * A screen that stands alone (`.center`): the seal 40 under the top, the head, the line centred at
 * most 300 wide, then [buttons] (Update required, Sign in again, Account deleted).
 */
@Composable
fun FlowSealPage(title: String, line: String, buttons: @Composable ColumnScope.() -> Unit) {
    FlowPage {
        Box(Modifier.fillMaxWidth().padding(top = SEAL_TOP, bottom = FsSpace.space2), contentAlignment = Alignment.Center) { Seal(SEAL) }
        FlowInsetHead(title)
        Box(Modifier.fillMaxWidth().padding(top = FsSpace.space2), contentAlignment = Alignment.Center) {
            Text(line, style = FsType.small, color = FsColor.muted, textAlign = TextAlign.Center, modifier = Modifier.widthIn(max = LINE_MAX))
        }
        FlowStack(buttons)
    }
}

/** A small tag (`.pill.tag`): 20 tall, 12/16 semibold, in its tone's outline; muted for `This device`. */
@Composable
fun FlowTag(text: String, tone: Color = FsColor.muted, edge: Color = FsColor.line) {
    Box(
        Modifier
            .heightIn(min = TAG_HEIGHT)
            .border(1.dp, edge, RoundedCornerShape(TAG_RADIUS))
            .padding(horizontal = TAG_PAD),
        contentAlignment = Alignment.Center,
    ) {
        Text(text, style = TAG, color = tone)
    }
}

private val NOTICE = FsType.body.copy(lineHeight = 24.sp)
private val NOTICE_TOP = 14.dp
private val STACK_BUTTON = 48.dp
private val SEAL = 84.dp
private val SEAL_TOP = 40.dp
private val LINE_MAX = 300.dp
private val TAG_HEIGHT = 20.dp
private val TAG_RADIUS = 3.dp
private val TAG_PAD = 6.dp
private val TAG = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.SemiBold)
