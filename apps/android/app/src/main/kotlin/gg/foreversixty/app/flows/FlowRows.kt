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
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.selected
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.AnnotatedString
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.borderColor
import gg.foreversixty.app.designsystem.containerColor
import gg.foreversixty.app.designsystem.contentColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSize
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsTarget
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ScaledType
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.home.LARGE_TEXT_SCALE
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

    /**
     * A checklist mark (Ready check): ticked, missing or a warning, a glyph; with [icon] the item the
     * check names beside it (28, in the icon frame), when [build] is known.
     */
    data class Check(val ok: Boolean, val miss: Boolean, val icon: String? = null, val build: String? = null) : FlowLead
}

/**
 * One row (the mock's Row): an optional 40 lead, the title (15/20 bold), its line (13/18) and an optional
 * [meta] line under it ([styledSub] in place of [sub] for a line in colours), an optional [value]
 * at the right in tabular figures (an Activity row's DPS), and a
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
    styledSub: AnnotatedString? = null,
    value: String? = null,
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
                    FlowLinkLine(sub, subLink, SUB, LocalSubInk.current, Modifier.fillMaxWidth())
                } else if (styledSub != null) {
                    Text(styledSub, style = SUB, color = LocalSubInk.current)
                } else if (!sub.isNullOrBlank()) {
                    Text(sub, style = SUB, color = LocalSubInk.current)
                }
                // A third line in the same face (Activity's time: `8 h ago`, `6:02 pm`, `Tue`).
                if (meta != null) Text(meta, style = SUB, color = LocalSubInk.current)
            }
            if (value != null) {
                Spacer(Modifier.width(FsSpace.space3))
                Text(value, style = VALUE, color = FsColor.strong, maxLines = 1)
            }
        }
        if (trail != null) Box(Modifier.padding(top = FsSpace.space2)) { trail() }
    }
    if (!last) FlowRule()
}

/** A gold link at the end of a row's line (`None in your bags · Auction house`), round 5's inline fix. */
data class FlowSubLink(val label: String, val onClick: () -> Unit)

/**
 * A line ending in its gold [link] (`60 raid · Alliance · Change`), for a row's line or any lead line.
 * The whole line is the control (spec 3.3's padded link): a 13 px link keeps its look and the layout
 * keeps the line's height, while Compose's minimum touch target takes the tap 48 dp tall around it.
 */
@Composable
fun FlowLinkLine(sub: String?, link: FlowSubLink, style: TextStyle, color: Color, modifier: Modifier = Modifier) {
    Text(
        subWithLink(sub, link.label),
        style = style,
        color = color,
        modifier = modifier.clickable(role = Role.Button, onClickLabel = link.label, onClick = link.onClick),
    )
}

private fun subWithLink(sub: String?, label: String) = buildAnnotatedString {
    if (!sub.isNullOrBlank()) append("$sub · ")
    withStyle(SpanStyle(color = FsColor.gold, fontWeight = FontWeight.SemiBold)) { append(label) }
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
            Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(FsSpace.space2)) {
                Box(Modifier.size(CHECK).background(color, CircleShape), contentAlignment = Alignment.Center) {
                    Text(mark, style = CHECK_TEXT, color = FsColor.bg)
                }
                if (lead.icon != null && lead.build != null) {
                    ItemIcon(lead.icon, null, lead.build, FsSize.icon28, Modifier.parityItem(ParityItemKind.Image))
                }
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
        // From 1.5x a Row's trail moves under its body (spec 3.2), so the name keeps the row's width.
        val gainUnder = LocalDensity.current.fontScale >= LARGE_TEXT_SCALE
        Column(Modifier.weight(1f)) {
            Text(look.name.orEmpty(), style = FsType.item, color = ItemQuality(look.quality).text)
            Text(source, style = SUB, color = FsColor.muted)
            if (gain != null && gainUnder) ItemGain(gain, unit, gainNote, Alignment.Start)
        }
        if (gain != null && !gainUnder) {
            Spacer(Modifier.width(FsSpace.space3))
            ItemGain(gain, unit, gainNote, Alignment.End)
        }
    }
}

/** The gain and its basis under it (the item sheet's `over your worn`). */
@Composable
private fun ItemGain(gain: String, unit: String?, note: String?, align: Alignment.Horizontal) {
    Column(horizontalAlignment = align) {
        Text(gainText(gain, unit), style = FsType.figureSm, color = FsColor.strong)
        if (note != null) Text(note, style = SUB, color = FsColor.muted)
    }
}

private fun gainText(gain: String, unit: String?) = buildAnnotatedString {
    append(gain)
    if (!unit.isNullOrEmpty()) {
        withStyle(FsType.figureUnit.toSpanStyle().copy(fontWeight = FontWeight.SemiBold, color = FsColor.muted)) { append(" $unit") }
    }
}

/**
 * The flows' button (`.btn`): 36 tall small, 44 otherwise, in the site's and iOS's colours (round 5
 * row 5): [primary] the flat teal with the teal edge (the selected answer among the raid's four,
 * Accept), else the secondary blue with the blue edge (Preview mention, Decline), white labels.
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
    val kind = if (primary) BrandButtonKind.Primary else BrandButtonKind.Secondary
    Box(
        modifier
            .heightIn(min = if (small) SMALL_BUTTON else BUTTON)
            .parityItem(ParityItemKind.Button)
            .clip(shape)
            .background(containerColor(kind) ?: Color.Transparent, shape)
            .border(1.dp, borderColor(kind), shape)
            .clickable(role = Role.Button, onClick = onClick)
            .then(selectedState?.let { on -> Modifier.semantics { selected = on } } ?: Modifier)
            .padding(horizontal = BUTTON_PAD),
        contentAlignment = Alignment.Center,
    ) {
        Text(label, style = BUTTON_TEXT, color = contentColor(kind), maxLines = 1)
    }
}

/**
 * A choice chip (round 5, spec Mobile): the rule's outline, the selected one filled in the teal edge as a
 * selected segment is. Late's minutes (48 tall) and Activity's filters ([minHeight] 36, 48 to the touch).
 * Not [enabled] it takes [FlowSegmented]'s quiet colours.
 */
@Composable
fun FlowChip(
    label: String,
    on: Boolean,
    onClick: () -> Unit,
    modifier: Modifier = Modifier,
    minHeight: Dp = FsTarget.min,
    role: Role = Role.Button,
    enabled: Boolean = true,
) {
    val shape = RoundedCornerShape(CHIP_RADIUS)
    val colors = ChoiceColors.of(enabled)
    Box(
        modifier
            .heightIn(min = minHeight)
            .parityItem(ParityItemKind.Button)
            .clip(shape)
            .background(if (on) colors.selected else Color.Transparent, shape)
            .border(1.dp, if (on) colors.selectedEdge else colors.edge, shape)
            .clickable(enabled = enabled, role = role, onClick = onClick)
            .semantics { selected = on }
            .padding(horizontal = FsSpace.space4),
        contentAlignment = Alignment.Center,
    ) {
        Text(label, style = ScaledType.segment, color = colors.text, maxLines = 1)
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
    Text(text, style = FsType.bodyStrong, color = FsColor.gold, modifier = modifier.clickable(role = Role.Button, onClick = onClick))
}

/** The small chip's height (Activity's filters, Home's Late minutes); Compose takes the tap 48 tall. */
val FLOW_CHIP_SMALL: Dp = 36.dp

/** A row's trailing figure (iOS PanelRow's value): the strong body role in tabular figures. */
private val VALUE = FsType.bodyStrong.copy(fontFeatureSettings = "tnum")

/** The row's line (`.row .sub`, 13/18). */
val SUB = FsType.rowSub

private val LEAD = 40.dp
private val CHIP_RADIUS = 8.dp
private val CHECK = 22.dp
private const val CHECK_OK = "\u2713"
private const val CHECK_MISS = "\u00d7"
private const val CHECK_WARN = "!"
private val CHECK_TEXT = FsType.buttonSmall
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
private val BUTTON_TEXT = FsType.buttonSmall
