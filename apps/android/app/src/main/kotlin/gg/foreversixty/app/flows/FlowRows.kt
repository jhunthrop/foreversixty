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
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.selected
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.LinkAnnotation
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.TextLinkStyles
import androidx.compose.ui.text.withLink
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsTarget
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ScaledType
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.items.GameIconImage
import gg.foreversixty.app.items.ItemIcon
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemQuality

/** A row's leading image: a game icon (instance, spec, event) or an item in its quality frame. */
sealed interface FlowLead {
    data class Icon(val icon: String?, val build: String?) : FlowLead
    data class Item(val look: ItemLook, val build: String) : FlowLead

    /** A character: the class crest (rosters, the unanswered, the applicant's guildmates). */
    data class Crest(val className: String) : FlowLead

    /** A checklist mark (Ready check): ticked, missing or a warning; a glyph, not an image. */
    data class Check(val ok: Boolean, val miss: Boolean) : FlowLead
}

/**
 * One row (the mock's Row): an optional 40 lead, the title (15/20 bold), its line (13/18) and an optional
 * [meta] line under it, and a
 * trailing control that drops under the text at phone width. [slim] rows are 6 above and below and at
 * least 44 tall; others are 10. [unread] puts the gold dot in the gutter. A rule follows unless [last].
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
    trail: (@Composable () -> Unit)? = null,
    onClick: (() -> Unit)? = null,
    subLink: FlowSubLink? = null,
    meta: String? = null,
) {
    val unreadLabel = androidx.compose.ui.res.stringResource(gg.foreversixty.app.R.string.activity_unread)
    val tap = onClick?.let { Modifier.clickable(role = Role.Button, onClick = it) } ?: Modifier
    Column(
        Modifier
            .fillMaxWidth()
            .parityItem(ParityItemKind.Row)
            .then(tap)
            .heightIn(min = if (slim) SLIM_MIN else 0.dp)
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
                LeadImage(lead)
                Spacer(Modifier.width(FsSpace.space3))
            }
            Column(Modifier.weight(1f)) {
                Text(title, style = FsType.item, color = titleColor)
                if (subLink != null) {
                    Text(subWithLink(sub, subLink), style = SUB, color = LocalSubInk.current)
                } else if (!sub.isNullOrBlank()) {
                    Text(sub, style = SUB, color = LocalSubInk.current)
                }
                // A third line in the same face (Activity's time: `8 h ago`, `6:02 pm`, `Tue`).
                if (meta != null) Text(meta, style = SUB, color = LocalSubInk.current)
            }
        }
        if (trail != null) Box(Modifier.padding(top = FsSpace.space2)) { trail() }
    }
    if (!last) FlowRule()
}

/** A gold link at the end of a row's line (`None in your bags · Auction house`), round 5's inline fix. */
data class FlowSubLink(val label: String, val onClick: () -> Unit)

/** A line ending in its gold [link] (`60 raid · Alliance · Change`), for a row's line or any lead line. */
fun subWithLink(sub: String?, link: FlowSubLink) = buildAnnotatedString {
    if (!sub.isNullOrBlank()) append("$sub · ")
    withLink(LinkAnnotation.Clickable(link.label, TextLinkStyles(SpanStyle(color = FsColor.gold, fontWeight = FontWeight.SemiBold))) { link.onClick() }) {
        append(link.label)
    }
}

@Composable
private fun LeadImage(lead: FlowLead) {
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
        is FlowLead.Crest -> gg.foreversixty.app.common.ClassCrest(lead.className, LEAD, Modifier.parityItem(ParityItemKind.Image))
        is FlowLead.Check -> {
            val (mark, color) = when {
                lead.ok -> CHECK_OK to FsColor.kill
                lead.miss -> CHECK_MISS to FsColor.wipe
                else -> CHECK_WARN to FsColor.gold
            }
            // The mark on its filled disc (round 5, mock m-android-raid): green, red or amber, the glyph dark.
            Box(Modifier.size(CHECK).background(color, CircleShape), contentAlignment = Alignment.Center) {
                Text(mark, style = CHECK_TEXT, color = FsColor.bg)
            }
        }
    }
}

/**
 * One item (the mock's ItemRow): the icon in its quality frame, the name in its quality colour (it opens
 * the item's tooltip), its source, and the gain on the right (`+14.5` with its unit small). [inset]
 * is the row's own side padding: 0 inside a [FlowPanel], which carries it.
 */
@Composable
fun FlowItemRow(
    look: ItemLook,
    build: String,
    source: String,
    gain: String?,
    unit: String?,
    gainNote: String? = null,
    inset: Dp = FsSpace.space4,
    onClick: () -> Unit,
) {
    Row(
        Modifier
            .fillMaxWidth()
            .parityItem(ParityItemKind.Row)
            .clickable(role = Role.Button, onClick = onClick)
            .padding(horizontal = inset, vertical = FsSpace.space2)
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
 * answer among the raid's four.
 */
@Composable
fun FlowButton(
    label: String,
    onClick: () -> Unit,
    modifier: Modifier = Modifier,
    primary: Boolean = false,
    small: Boolean = true,
    selectedState: Boolean? = null,
) {
    val shape = RoundedCornerShape(FsRadius.control)
    val fill = Modifier.background(if (primary) FsColor.teal else FsColor.raised, shape)
    Box(
        modifier
            .heightIn(min = if (small) SMALL_BUTTON else BUTTON)
            .parityItem(ParityItemKind.Button)
            .clip(shape)
            .then(fill)
            .border(1.dp, if (primary) FsColor.tealEdge else FsColor.line, shape)
            .clickable(role = Role.Button, onClick = onClick)
            .then(selectedState?.let { on -> Modifier.semantics { selected = on } } ?: Modifier)
            .padding(horizontal = BUTTON_PAD),
        contentAlignment = Alignment.Center,
    ) {
        Text(label, style = BUTTON_TEXT, color = FsColor.strong, maxLines = 1)
    }
}

/**
 * A choice chip (round 5, spec Mobile): the rule's outline, the selected one filled in the teal edge as a
 * selected segment is. Late's minutes (48 tall) and Activity's filters ([minHeight] 36, 48 to the touch).
 */
@Composable
fun FlowChip(label: String, on: Boolean, onClick: () -> Unit, modifier: Modifier = Modifier, minHeight: Dp = FsTarget.min, role: Role = Role.Button) {
    val shape = RoundedCornerShape(CHIP_RADIUS)
    Box(
        modifier
            .heightIn(min = minHeight)
            .parityItem(ParityItemKind.Button)
            .clip(shape)
            .background(if (on) FsColor.tealEdge else Color.Transparent, shape)
            .border(1.dp, FsColor.rule, shape)
            .clickable(role = role, onClick = onClick)
            .semantics { selected = on }
            .padding(horizontal = FsSpace.space4),
        contentAlignment = Alignment.Center,
    ) {
        Text(label, style = ScaledType.segment, color = FsColor.strong, maxLines = 1)
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

/** The row's line (`.row .sub`, 13/18). */
val SUB = FsType.small.copy(lineHeight = 18.sp)

private val LEAD = 40.dp
private val CHIP_RADIUS = 8.dp
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
private val BUTTON_TEXT = FsType.bodyStrong.copy(fontSize = 13.sp, lineHeight = 16.sp, fontWeight = FontWeight.Bold)
