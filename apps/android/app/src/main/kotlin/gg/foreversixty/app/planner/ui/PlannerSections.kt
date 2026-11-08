package gg.foreversixty.app.planner.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ColumnScope
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.offset
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.text.BasicTextField
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.alpha
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.SolidColor
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.spec.DpColumn
import gg.foreversixty.app.bis.spec.wholeDp
import gg.foreversixty.app.bis.spec.linked
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.flows.FlowButton
import gg.foreversixty.app.flows.FlowLink
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.FlowSkeleton
import gg.foreversixty.app.items.GameIconImage
import gg.foreversixty.app.planner.logic.PlannerShare
import gg.foreversixty.app.planner.logic.TalentIndex
import gg.foreversixty.app.talents.dashedBorder

/** A secondary page section's top (`.layout-wb` sections, 48 apart). */
val PLANNER_SECTION_TOP: Dp = 48.dp

/** The band panel (`BandCompare`): the line, Load and See what differs, and the band's gear. */
class BandSectionModel(
    val title: String,
    val line: String,
    /** `Load the levels 30 to 39 build`, or null when there is nothing to load (a shared build, a match). */
    val load: String?,
    val canDiffer: Boolean,
    val differsShown: Boolean,
    /** `30 to 39`, for `Gear for 30 to 39`. */
    val gearRange: String,
)

class BandSectionActions(val onLoad: () -> Unit = {}, val onDiffers: () -> Unit = {}, val onBis: () -> Unit = {})

@Composable
fun PlannerBandSection(model: BandSectionModel, actions: BandSectionActions) {
    FlowSection(model.title, top = PLANNER_SECTION_TOP) {
        SectionBody {
            Text(model.line, style = FsType.body, color = FsColor.text, modifier = Modifier.wholeDp())
            if (model.load != null || model.canDiffer) {
                Row(horizontalArrangement = Arrangement.spacedBy(GAP_ACTIONS)) {
                    model.load?.let { FlowButton(it, actions.onLoad, small = false) }
                    if (model.canDiffer) {
                        FlowButton(stringResource(R.string.plannerPage_seeWhatDiffers), actions.onDiffers, small = false, selectedState = model.differsShown)
                    }
                }
            }
            Column {
                HorizontalDivider(color = FsColor.lineSoft)
                Row(Modifier.fillMaxWidth().padding(top = GAP_GEAR), verticalAlignment = Alignment.Bottom) {
                    Text(stringResource(R.string.plannerPage_gearFor, model.gearRange), style = FsType.small, color = FsColor.muted, modifier = Modifier.weight(1f).padding(bottom = 2.dp))
                    FlowLink(stringResource(R.string.plannerPage_bisList), onClick = actions.onBis)
                }
            }
        }
    }
}

/** In the game: `Send this build to Felsnap` (the site's send) when the build is the character's class, then Copy. */
class InGameModel(val sendTo: String?, val copied: Boolean, val canCopy: Boolean)

class InGameActions(val send: () -> Unit = {}, val copy: () -> Unit = {}, val addon: () -> Unit = {})

@Composable
fun PlannerInGameSection(model: InGameModel, actions: InGameActions) {
    FlowSection(stringResource(R.string.plannerPage_inGame), top = PLANNER_SECTION_TOP) {
        SectionBody {
            model.sendTo?.let { FlowButton(stringResource(R.string.plannerPage_sendTo, it), actions.send, primary = true, small = false) }
            val copy = stringResource(if (model.copied) R.string.plannerPage_copied else R.string.plannerPage_copyForGame)
            FlowButton(copy, actions.copy, Modifier.alpha(if (model.canCopy) 1f else OFF_ALPHA), primary = model.sendTo == null, small = false)
            Text(stringResource(R.string.plannerPage_gameHint), style = FsType.small, color = FsColor.muted, modifier = Modifier.wholeDp())
            if (model.sendTo == null) FlowLink(stringResource(R.string.plannerPage_getAddon), onClick = actions.addon)
        }
    }
}

/**
 * Point order (`PointOrder`): rows of ten levels under a faint 0 to 9 marker, each socket the talent
 * taken at that level with its rank, the next socket in gold, the rest empty.
 */
@Composable
fun PlannerPointOrderSection(index: TalentIndex, order: List<Int>, build: String) {
    FlowSection(stringResource(R.string.plannerPage_pointOrder), meta = stringResource(R.string.plannerPage_spent, order.size), top = PLANNER_SECTION_TOP) {
        // On a whole-dp grid (the BiS page's DpColumn): six rows of 73 do not drift down a pixel each.
        DpColumn(Modifier.fillMaxWidth()) {
            Row(Modifier.padding(top = FsSpace.space1).clearAndSetSemantics {}, horizontalArrangement = Arrangement.spacedBy(SOCKET_GAP)) {
                repeat(PER_ROW) { Text(it.toString(), style = MARKER, color = MARKER_COLOR, textAlign = TextAlign.Center, modifier = Modifier.width(SOCKET).height(MARKER_LINE)) }
            }
            // Each point with the rank it brings its talent to.
            val taken = order.mapIndexed { i, id -> id to order.subList(0, i + 1).count { it == id } }
            DECADES.forEachIndexed { row, first ->
                if (row > 0) HorizontalDivider(color = FsColor.lineSoft)
                DpColumn(Modifier.padding(vertical = FsSpace.space2), spacing = 6.dp) {
                    Text(decade(first), style = FsType.small, color = FsColor.muted, modifier = Modifier.height(SMALL_LINE))
                    Row(horizontalArrangement = Arrangement.spacedBy(SOCKET_GAP)) {
                        val sockets = if (first >= LAST_LEVEL) 1 else PER_ROW
                        repeat(sockets) { i ->
                            val point = first - FIRST_LEVEL + i
                            Socket(taken.getOrNull(point), index, build, level = first + i, next = point == order.size)
                        }
                    }
                }
            }
        }
    }
}

@Composable
private fun decade(first: Int): String =
    if (first >= LAST_LEVEL) stringResource(R.string.plannerPage_bandLevel, first) else stringResource(R.string.plannerPage_bandLevels, first, first + PER_ROW - 1)

@Composable
private fun Socket(point: Pair<Int, Int>?, index: TalentIndex, build: String, level: Int, next: Boolean) {
    val talent = point?.let { index.byId[it.first] }
    if (point == null || talent == null) {
        val description = if (next) stringResource(R.string.plannerPage_nextPoint, level) else stringResource(R.string.plannerPage_bandLevel, level)
        Box(
            Modifier.size(SOCKET).dashedBorder(if (next) FsColor.gold else FsColor.line, SOCKET_RADIUS, width = 1.dp)
                .semantics { contentDescription = description },
        )
        return
    }
    val description = stringResource(R.string.plannerPage_pointAt, level, talent.name, point.second)
    Box(Modifier.size(SOCKET).semantics(mergeDescendants = true) { contentDescription = description }) {
        Box(Modifier.size(SOCKET).clip(RoundedCornerShape(SOCKET_RADIUS)).parityItem(ParityItemKind.Image)) {
            GameIconImage(talent.icon, build, SOCKET, fallback = talent.name.take(2))
        }
        Text(
            point.second.toString(), style = SOCKET_RANK, color = FsColor.muted,
            modifier = Modifier.align(Alignment.BottomEnd).offset(4.dp, 5.dp).background(FsColor.bg, CircleShape).padding(horizontal = 3.dp),
        )
    }
}

/** Share this build: the title and Share, the build as a line with Copy as text, and Your builds. */
class ShareModel(val title: String, val summary: String, val copied: Boolean, val signedIn: Boolean, val problem: String? = null)

class ShareActions(
    val onTitle: (String) -> Unit = {},
    val share: () -> Unit = {},
    val copyText: () -> Unit = {},
    val builds: () -> Unit = {},
    val signIn: () -> Unit = {},
)

@Composable
fun PlannerShareSection(model: ShareModel, actions: ShareActions) {
    FlowSection(stringResource(R.string.plannerPage_shareTitle), top = PLANNER_SECTION_TOP) {
        SectionBody {
            Row(verticalAlignment = Alignment.Bottom, horizontalArrangement = Arrangement.spacedBy(GAP_ACTIONS)) {
                Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(6.dp)) {
                    Text(stringResource(R.string.plannerPage_titleLabel), style = LABEL, color = FsColor.muted)
                    InputField(
                        model.title, { actions.onTitle(it.take(PlannerShare.TITLE_MAX)) }, stringResource(R.string.plannerPage_titlePlaceholder),
                        FIELD_TEXT, Modifier.parityItem(ParityItemKind.Field).heightIn(min = FIELD_HEIGHT), singleLine = true,
                    )
                }
                FlowButton(stringResource(R.string.plannerPage_share), actions.share, small = false)
            }
            model.problem?.let { Text(it, style = FsType.small, color = FsColor.wipe) }
            Column(verticalArrangement = Arrangement.spacedBy(GAP_ACTIONS)) {
                Text(model.summary, style = FsType.small, color = FsColor.muted, modifier = Modifier.wholeDp())
                FlowButton(stringResource(if (model.copied) R.string.plannerPage_copied else R.string.plannerPage_copyAsText), actions.copyText, small = false)
            }
            val (label, open) = if (model.signedIn) R.string.plannerPage_yourBuilds to actions.builds else R.string.plannerPage_signInForBuilds to actions.signIn
            Box(Modifier.fillMaxWidth().heightIn(min = FIELD_HEIGHT), contentAlignment = Alignment.CenterStart) { FlowLink(stringResource(label), onClick = open) }
        }
    }
}

/** Import from addon: the code field, Import, and the way to the addon. */
class ImportModel(val code: String, val unreadable: Boolean)

class ImportActions(val onCode: (String) -> Unit = {}, val import: () -> Unit = {}, val addon: () -> Unit = {})

@Composable
fun PlannerImportSection(model: ImportModel, actions: ImportActions) {
    FlowSection(stringResource(R.string.plannerPage_importTitle), top = PLANNER_SECTION_TOP) {
        SectionBody {
            InputField(
                model.code, actions.onCode, stringResource(R.string.plannerPage_importPlaceholder), CODE_TEXT,
                Modifier.heightIn(min = CODE_HEIGHT).background(CODE_FILL, RoundedCornerShape(FsRadius.control)),
                border = if (model.unreadable) FsColor.wipe else FsColor.line,
            )
            FlowButton(stringResource(R.string.plannerPage_import), actions.import, Modifier.alpha(if (model.code.isBlank()) OFF_ALPHA else 1f), small = false)
            if (model.unreadable) Text(stringResource(R.string.plannerPage_importBad), style = FsType.small, color = FsColor.wipe)
            Text(linked(stringResource(R.string.plannerPage_noExport) + " ", stringResource(R.string.plannerPage_addonLink), actions.addon), style = FsType.small, color = FsColor.muted, modifier = Modifier.wholeDp())
        }
    }
}

/** A section that has not loaded: its title, a line and a button's place. */
@Composable
fun PlannerSkeletonSection(title: String) {
    FlowSection(title, top = PLANNER_SECTION_TOP) {
        SectionBody {
            FlowSkeleton(SKELETON_LINE, FsSpace.space5)
            FlowSkeleton(SKELETON_BUTTON, FIELD_HEIGHT)
        }
    }
}

/** A planner section's body (`.pl-sec`): 14 under the rule, 12 between its parts. */
@Composable
private fun SectionBody(content: @Composable ColumnScope.() -> Unit) {
    Column(Modifier.fillMaxWidth().padding(top = SECTION_BODY), verticalArrangement = Arrangement.spacedBy(FsSpace.space3), content = content)
}

/** A bordered text field (`.field`, `.codef`) with its placeholder. */
@Composable
private fun InputField(
    value: String,
    onValue: (String) -> Unit,
    placeholder: String,
    style: TextStyle,
    modifier: Modifier,
    singleLine: Boolean = false,
    border: Color = FsColor.line,
) {
    val shape = RoundedCornerShape(FsRadius.control)
    BasicTextField(
        value = value,
        onValueChange = onValue,
        singleLine = singleLine,
        textStyle = style.copy(color = FsColor.strong),
        cursorBrush = SolidColor(FsColor.gold),
        modifier = modifier.fillMaxWidth().border(1.dp, border, shape),
        decorationBox = { inner ->
            Box(Modifier.padding(horizontal = FsSpace.space3, vertical = FIELD_PAD), contentAlignment = Alignment.CenterStart) {
                if (value.isEmpty()) Text(placeholder, style = style, color = FsColor.muted)
                inner()
            }
        },
    )
}

/** Every level a point is spent at: 10 to 60, in rows of ten. */
private const val FIRST_LEVEL = 10
private const val LAST_LEVEL = 60
private const val PER_ROW = 10
private val DECADES = (FIRST_LEVEL..LAST_LEVEL step PER_ROW).toList()
private const val OFF_ALPHA = 0.5f
private val SECTION_BODY = 14.dp
private val GAP_ACTIONS = 10.dp
private val GAP_GEAR = 10.dp
private val SOCKET = 30.dp
private val MARKER_LINE = 14.dp
private val SMALL_LINE = 20.dp
private val SOCKET_GAP = 4.dp
private val SOCKET_RADIUS = 5.dp
private val FIELD_HEIGHT = 44.dp
private val FIELD_PAD = 10.dp
private val CODE_HEIGHT = 68.dp
private val SKELETON_LINE = 260.dp
private val SKELETON_BUTTON = 200.dp
/** The faint level marker over the first row (`.porder-m`). */
private val MARKER_COLOR = Color(0xFF5D6577)
/** The code field's well (`.codef`). */
private val CODE_FILL = Color(0xFF0A0D14)
private val MARKER = FsType.small.copy(fontSize = 11.sp, lineHeight = 14.sp)
private val SOCKET_RANK = FsType.small.copy(fontSize = 12.sp, lineHeight = 14.sp, fontWeight = FontWeight.Bold)
private val LABEL = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.Bold)
private val FIELD_TEXT = FsType.body.copy(fontSize = 14.sp, lineHeight = 20.sp)
private val CODE_TEXT = FsType.small.copy(fontFamily = FontFamily.Monospace, fontWeight = FontWeight.Medium)
