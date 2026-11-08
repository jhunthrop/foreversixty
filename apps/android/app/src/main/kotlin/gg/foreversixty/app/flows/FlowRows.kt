package gg.foreversixty.app.flows

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.offset
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.requiredSize
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Size
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.selected
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.AnnotatedString
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.items.GameIconImage
import gg.foreversixty.app.items.ItemIcon
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemQuality

/** A row's leading image: a game icon (instance, spec, event) or an item in its quality frame. */
sealed interface FlowLead {
    data class Icon(val icon: String?, val build: String?) : FlowLead
    data class Item(val look: ItemLook, val build: String) : FlowLead

    /** A character: the class crest (rosters, the unanswered, the applicant's guildmates); signed-out Home's classes are 44. */
    data class Crest(val className: String, val size: Dp = LEAD) : FlowLead

    /** A checklist mark (Ready check): ticked, missing or a warning; a glyph, not an image. */
    data class Check(val ok: Boolean, val miss: Boolean) : FlowLead
}

/**
 * One row (the mock's Row): an optional 40 lead, the title (15/20 bold) and its line (13/18), and a
 * trailing control that drops under the text at phone width. [slim] rows are 6 above and below and at
 * least 44 tall; others are 10. [unread] puts the gold dot in the gutter. A rule follows unless [last].
 * [inset] pads the row at the sides (`.rgrid .row`, 16); [selected] is the current pick (`.row.cur`): the
 * gold wash, the gold bar at the left and 16 at the sides.
 */
@Composable
fun FlowRow(
    title: String,
    sub: String? = null,
    lead: FlowLead? = null,
    titleColor: Color = FsColor.strong,
    slim: Boolean = false,
    unread: Boolean = false,
    last: Boolean = false,
    inset: Dp = 0.dp,
    selected: Boolean = false,
    trail: (@Composable () -> Unit)? = null,
    onClick: (() -> Unit)? = null,
) {
    val unreadLabel = androidx.compose.ui.res.stringResource(gg.foreversixty.app.R.string.activity_unread)
    val tap = onClick?.let { Modifier.clickable(role = Role.Button, onClick = it) } ?: Modifier
    Column(
        Modifier
            .fillMaxWidth()
            .parityItem(ParityItemKind.Row)
            .then(tap)
            .then(if (selected) Modifier.semantics { this.selected = true }.background(FsColor.selected).drawBehind { drawRect(FsColor.gold, size = Size(SELECTED_BAR.toPx(), size.height)) } else Modifier)
            .heightIn(min = if (slim) SLIM_MIN else 0.dp)
            .padding(horizontal = if (selected) FsSpace.rowPadX else inset)
            .padding(vertical = if (slim) SLIM_PAD else ROW_PAD),
        verticalArrangement = Arrangement.Center,
    ) {
        Row(verticalAlignment = Alignment.CenterVertically) {
            if (unread) {
                Box(Modifier.size(0.dp)) {
                    Box(
                        Modifier
                            .offset(x = -DOT_INSET, y = -DOT_RISE)
                            .requiredSize(DOT)
                            .background(FsColor.gold, CircleShape)
                            .semantics { contentDescription = unreadLabel },
                    )
                }
            }
            if (lead != null) {
                FlowLeadImage(lead)
                Spacer(Modifier.width(FsSpace.space3))
            }
            Column(Modifier.weight(1f)) {
                Text(title, style = FsType.item, color = titleColor)
                if (!sub.isNullOrBlank()) Text(sub, style = SUB, color = FsColor.muted)
            }
        }
        if (trail != null) Box(Modifier.padding(top = FsSpace.space2)) { trail() }
    }
    if (!last) FlowRule()
}

/** A row's lead on its own (the first raids' icons under the countdown). */
@Composable
fun FlowLeadImage(lead: FlowLead) {
    when (lead) {
        is FlowLead.Icon -> Box(
            Modifier
                .size(LEAD)
                .parityItem(ParityItemKind.Image)
                .clip(RoundedCornerShape(FsRadius.control))
                .background(FsColor.raised)
                .border(1.dp, FsColor.line, RoundedCornerShape(FsRadius.control)),
        ) { GameIconImage(lead.icon, lead.build, LEAD) }
        is FlowLead.Item -> ItemIcon(lead.look.icon, lead.look.quality, lead.build, LEAD, Modifier.parityItem(ParityItemKind.Image))
        is FlowLead.Crest -> gg.foreversixty.app.common.ClassCrest(lead.className, lead.size, Modifier.parityItem(ParityItemKind.Image))
        is FlowLead.Check -> {
            val (mark, color) = when {
                lead.ok -> CHECK_OK to FsColor.kill
                lead.miss -> CHECK_MISS to FsColor.wipe
                else -> CHECK_WARN to FsColor.gold
            }
            Box(Modifier.size(CHECK), contentAlignment = Alignment.Center) {
                Text(mark, style = CHECK_TEXT, color = color)
            }
        }
    }
}

/**
 * One item (the mock's ItemRow): the icon in its quality frame, the name in its quality colour (it opens
 * the item's tooltip), its source, and the gain on the right (`+14.5` with its unit small).
 */
@Composable
fun FlowItemRow(look: ItemLook, build: String, source: String, gain: String?, unit: String?, gainNote: String? = null, onClick: () -> Unit) {
    Row(
        Modifier
            .fillMaxWidth()
            .parityItem(ParityItemKind.Row)
            .clickable(role = Role.Button, onClick = onClick)
            .padding(horizontal = FsSpace.space4, vertical = FsSpace.space2)
            .heightIn(min = ITEM_ROW_BODY),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        ItemIcon(look.icon, look.quality, build, LEAD, Modifier.parityItem(ParityItemKind.Image))
        Spacer(Modifier.width(FsSpace.space3))
        Column(Modifier.weight(1f)) {
            Text(look.name.orEmpty(), style = FsType.item, color = ItemQuality(look.quality).text)
            Text(source, style = SUB, color = FsColor.muted)
        }
        if (gain != null) {
            Spacer(Modifier.width(FsSpace.space3))
            // The gain's basis sits under it at the right (the item sheet's `over your worn`).
            Column(horizontalAlignment = Alignment.End) {
                Text(gainText(gain, unit), style = FsType.figureSm, color = FsColor.strong)
                if (gainNote != null) Text(gainNote, style = SUB, color = FsColor.muted)
            }
        }
    }
}

private fun gainText(gain: String, unit: String?) = buildAnnotatedString {
    append(gain)
    if (!unit.isNullOrEmpty()) {
        withStyle(SpanStyle(fontSize = 13.sp, fontWeight = FontWeight.SemiBold, color = FsColor.muted)) { append(" $unit") }
    }
}

/**
 * The flows' button (`.btn`): 36 tall small, 44 otherwise; [primary] is the gold fill, the selected
 * answer among the raid's four. [busy] holds the button's size with a spinner in place of the label and
 * ignores taps (Sign in while the browser is open).
 */
@Composable
fun FlowButton(
    label: String,
    onClick: () -> Unit,
    modifier: Modifier = Modifier,
    primary: Boolean = false,
    small: Boolean = true,
    selectedState: Boolean? = null,
    busy: Boolean = false,
) {
    val shape = RoundedCornerShape(FsRadius.control)
    val fill = if (primary) Modifier.background(Brush.verticalGradient(listOf(FsColor.goldHover, FsColor.gold)), shape) else Modifier.background(FsColor.raised, shape)
    Box(
        modifier
            .heightIn(min = if (small) SMALL_BUTTON else BUTTON)
            .parityItem(ParityItemKind.Button)
            .clip(shape)
            .then(fill)
            .border(1.dp, if (primary) FsColor.gold else FsColor.line, shape)
            .clickable(enabled = !busy, role = Role.Button, onClick = onClick)
            .then(selectedState?.let { on -> Modifier.semantics { selected = on } } ?: Modifier)
            .padding(horizontal = BUTTON_PAD),
        contentAlignment = Alignment.Center,
    ) {
        val ink = if (primary) FsColor.bg else FsColor.strong
        // The label stays (transparent) under the spinner so the button keeps its width.
        Text(label, style = BUTTON_TEXT, color = if (busy) Color.Transparent else ink, maxLines = 1)
        if (busy) CircularProgressIndicator(Modifier.size(SPINNER), color = ink, strokeWidth = SPINNER_STROKE)
    }
}

/** Buttons sharing a row equally, 8 apart (the raid's four answers, Late's three chips). */
@Composable
fun FlowButtonRow(labels: List<String>, selected: String?, onPress: (Int) -> Unit) {
    Row(Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.spacedBy(FsSpace.space2)) {
        labels.forEachIndexed { index, label ->
            FlowButton(label, { onPress(index) }, Modifier.weight(1f), primary = label == selected, selectedState = label == selected)
        }
    }
}

/** A full-width button (`.vstack .btn`, 44 tall), 20 under what precedes it. */
@Composable
fun FlowWideButton(label: String, primary: Boolean, onClick: () -> Unit) {
    FlowButton(label, onClick, Modifier.padding(top = FsSpace.space5).fillMaxWidth(), primary = primary, small = false)
}

/** An inline link (`.link`, 15/24 semibold gold): `Open raid night ›`, `Where to get it`. */
@Composable
fun FlowLink(text: String, modifier: Modifier = Modifier, onClick: () -> Unit) {
    Text(text, style = FsType.bodyStrong.copy(lineHeight = 24.sp), color = FsColor.gold, modifier = modifier.clickable(role = Role.Button, onClick = onClick))
}

/** A small gold tag (`.tag`): `Main` under the picked character. */
@Composable
fun FlowTag(text: String) {
    Text(
        text,
        style = TAG,
        color = FsColor.gold,
        modifier = Modifier
            .border(1.dp, FsColor.gold.copy(alpha = TAG_BORDER), RoundedCornerShape(TAG_RADIUS))
            .padding(horizontal = TAG_PAD_X, vertical = TAG_PAD_Y),
    )
}

/**
 * One line of a feed (the site's FeedRow, `Forever right now`): the 40 lead, the sentence (15/22, its
 * names and items styled by the caller) and, on the right, when and its kind; a rule under every line.
 */
@Composable
fun FlowFeedRow(lead: FlowLead, sentence: AnnotatedString, time: String, kind: String, onClick: () -> Unit) {
    Row(
        Modifier
            .fillMaxWidth()
            .parityItem(ParityItemKind.Row)
            .clickable(role = Role.Button, onClick = onClick)
            .padding(vertical = FsSpace.space3),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        FlowLeadImage(lead)
        Spacer(Modifier.width(FEED_GAP))
        Text(sentence, style = FsType.body, color = FsColor.text, modifier = Modifier.weight(1f))
        Spacer(Modifier.width(FEED_GAP))
        Column(horizontalAlignment = Alignment.End) {
            Text(time, style = FsType.small, color = FsColor.muted, maxLines = 1)
            Text(kind, style = FEED_KIND, color = FsColor.muted, maxLines = 1)
        }
    }
    FlowRule()
}

/** The row's line (`.row .sub`, 13/18). */
val SUB = FsType.small.copy(lineHeight = 18.sp)

private val LEAD = 40.dp
private val CHECK = 22.dp
private const val CHECK_OK = "\u2713"
private const val CHECK_MISS = "\u00d7"
private const val CHECK_WARN = "!"
private val CHECK_TEXT = FsType.bodyStrong.copy(fontSize = 13.sp, lineHeight = 22.sp, fontWeight = FontWeight.Bold)
private val ROW_PAD = 10.dp
private val SLIM_PAD = 6.dp
private val SLIM_MIN = 44.dp
private val DOT = 6.dp
private val DOT_INSET = 10.dp
private val DOT_RISE = 3.dp
private val ITEM_ROW_BODY: Dp = 46.dp
private val SMALL_BUTTON = 36.dp
private val BUTTON = 44.dp
private val BUTTON_PAD = 18.dp
private val SELECTED_BAR = 3.dp
private val SPINNER = 18.dp
private val SPINNER_STROKE = 2.dp
private val TAG_RADIUS = 4.dp
private val TAG_PAD_X = 7.dp
private val TAG_PAD_Y = 1.dp
private const val TAG_BORDER = 0.6f
private val TAG = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.SemiBold)
private val FEED_GAP = 14.dp
private val FEED_KIND = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.Bold)
private val BUTTON_TEXT = FsType.bodyStrong.copy(fontSize = 13.sp, lineHeight = 16.sp, fontWeight = FontWeight.Bold)
