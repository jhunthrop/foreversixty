package gg.foreversixty.app.raids.ui

import androidx.compose.foundation.Image
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.designsystem.FsType
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.layout.widthIn
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Warning
import androidx.compose.material3.Icon
import androidx.compose.material3.OutlinedButton
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.runtime.staticCompositionLocalOf
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import coil3.compose.AsyncImage
import gg.foreversixty.app.R
import androidx.compose.ui.graphics.vector.ImageVector
import gg.foreversixty.app.items.GameIconImage
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.designsystem.primaryCard
import gg.foreversixty.app.designsystem.BodyFamily
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsTextSize
import gg.foreversixty.app.designsystem.HeadingFamily
import gg.foreversixty.app.raids.logic.RaidChip
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidFacts
import java.time.Instant

/** The font scale from which rows restack (spec 10: iOS AX1, Android 1.5x and up). */
private const val LARGE_FONT_SCALE = 1.5f

/** True at the large text sizes, where rows restack instead of squeezing. */
@Composable
fun isLargeText(): Boolean = androidx.compose.ui.platform.LocalDensity.current.fontScale >= LARGE_FONT_SCALE

/** A one-word name: one line at every size, shrinking before it would break inside the word. */
@Composable
fun NameText(text: String, style: TextStyle, color: Color, modifier: Modifier = Modifier) {
    androidx.compose.foundation.text.BasicText(
        text = text,
        style = style.copy(color = color),
        maxLines = 1,
        softWrap = false,
        autoSize = androidx.compose.foundation.text.TextAutoSize.StepBased(minFontSize = 11.sp, maxFontSize = style.fontSize),
        modifier = modifier,
    )
}

/** The active data build the raid screens' game icons load from (the shared icon loader). */
val LocalRaidIconBuild = staticCompositionLocalOf<String?> { null }

/**
 * A square game icon on the shared loader: the spec icon (the `#3a4256` line) or a buff. Until the
 * icon loads, or when it fails, the tile shows [fallback] (`MotW`, `Bless`); a [dimmed] tile (a
 * missing buff) greys and fades the art only, so its label stays readable.
 */
@Composable
fun GameIcon(
    stem: String?,
    size: Dp,
    border: Color = GameColors.iconLine,
    borderWidth: Dp = 1.dp,
    fallback: String? = null,
    dimmed: Boolean = false,
) {
    val shape = RoundedCornerShape(FsRadius.pill)
    Box(Modifier.size(size).clip(shape).background(FsColor.cardTop).border(borderWidth, border, shape)) {
        GameIconImage(stem, LocalRaidIconBuild.current, size, fallback, dimmed)
    }
}

/** `micro` uppercase section label, with an optional mono count. */
@Composable
fun SectionLabel(text: String, color: Color = FsColor.muted, trailing: String? = null) {
    val tracking = if (isLargeText()) 0.sp else FsType.micro.letterSpacing
    Row(horizontalArrangement = Arrangement.spacedBy(6.dp), modifier = Modifier.semantics { heading() }) {
        Text(text.uppercase(), style = FsType.micro.copy(letterSpacing = tracking), color = color, maxLines = 1, softWrap = false)
        trailing?.let { Text(it, style = FsType.monoMeta, color = color) }
    }
}

/** The raid status tag: `Planned` gold, `Locked` muted, `Cancelled` wipe, `Ended` muted, `Done` kill. */
@Composable
fun ChipTag(chip: RaidChip) {
    val color = when (chip) {
        RaidChip.Planned -> FsColor.gold
        RaidChip.Cancelled -> FsColor.wipe
        RaidChip.Done -> FsColor.kill
        RaidChip.Locked, RaidChip.Ended -> FsColor.muted
    }
    Tag(stringResource(chip.label), color)
}

/** The date block, `THU` / `10` / `DEC` (hidden from TalkBack; the title row reads the date). */
@Composable
fun DateBlock(at: Instant, clock: RaidClock, width: Dp = 52.dp, height: Dp = 56.dp) {
    val (weekday, day, month) = clock.dateBlock(at)
    Column(
        horizontalAlignment = Alignment.CenterHorizontally,
        verticalArrangement = Arrangement.Center,
        modifier = Modifier
            .widthIn(min = width)
            .heightIn(min = height)
            .background(FsColor.bg.copy(alpha = 0.5f), RoundedCornerShape(FsRadius.control))
            .border(1.dp, FsColor.line, RoundedCornerShape(FsRadius.control))
            .padding(vertical = 4.dp, horizontal = 4.dp)
            .clearAndSetSemantics {},
    ) {
        Text(weekday, style = FsType.micro.copy(letterSpacing = 0.sp), color = FsColor.muted)
        Text(day, style = FsType.monoDay, color = FsColor.strong)
        Text(month, style = FsType.micro.copy(letterSpacing = 0.sp), color = FsColor.muted)
    }
}

/** The 8 dp progress bar: line-soft track, gold fill. A null width takes the full row. */
@Composable
fun FillBar(fill: Float, width: Dp?, muted: Boolean = false) {
    val track = if (width == null) Modifier.fillMaxWidth() else Modifier.width(width)
    Box(track.height(8.dp).clip(CircleShape).background(FsColor.lineSoft)) {
        Box(Modifier.fillMaxHeight().fillMaxWidth(fill).clip(CircleShape).background(if (muted) FsColor.muted else FsColor.gold))
    }
}

/** The count block (spec F4.1): the bar, `24 of 40 signed up`, and the role breakdown. */
@Composable
fun CountBlock(facts: RaidFacts, muted: Boolean = false) {
    val res = androidx.compose.ui.platform.LocalResources.current
    val line = facts.countLine(res)
    val breakdown = facts.breakdown(res)
    val a11y = stringResource(R.string.raids_a11y_count, line, breakdown.replace(RaidFacts.SEPARATOR, ", "))
    Column(verticalArrangement = Arrangement.spacedBy(6.dp), modifier = Modifier.clearAndSetSemantics { contentDescription = a11y }) {
        if (isLargeText()) {
            // Large text: the count above a full-width bar (spec 10).
            Text(line, style = FsType.bodyStrong, color = FsColor.strong)
            facts.fill?.let { FillBar(it, null, muted) }
        } else {
            Row(horizontalArrangement = Arrangement.spacedBy(12.dp), verticalAlignment = Alignment.CenterVertically) {
                facts.fill?.let { FillBar(it, 120.dp, muted) }
                Text(line, style = FsType.bodyStrong, color = FsColor.strong)
            }
        }
        Text(breakdown, style = FsType.meta, color = FsColor.muted)
    }
}

/** A raid card in the primary surface: full width, 16 dp inside. */
fun Modifier.raidPrimaryCard(): Modifier = fillMaxWidth().primaryCard().padding(16.dp)

/** A raid list in the quiet surface: full width, its rows edge to edge. */
fun Modifier.raidQuietCard(): Modifier = fillMaxWidth().quietCard()

/** The secondary button at its label's width, 48 dp (spec §4.5, sentence case). */
@Composable
fun RaidSecondaryButton(
    text: String,
    onClick: () -> Unit,
    enabled: Boolean = true,
    kind: BrandButtonKind = BrandButtonKind.Secondary,
    leadingIcon: ImageVector? = null,
    busy: Boolean = false,
) {
    BrandButton(text, kind, onClick, enabled = enabled, busy = busy, height = 48.dp, fill = false, leadingIcon = leadingIcon)
}

/** The gold text link, 48 dp target. */
@Composable
fun GoldLink(text: String, onClick: () -> Unit, trailing: Boolean = false) {
    Row(
        verticalAlignment = Alignment.CenterVertically,
        modifier = Modifier.heightIn(min = 48.dp).clickable(onClick = onClick),
    ) {
        Text(text + if (trailing) " ›" else "", style = FsType.bodyStrong, color = FsColor.gold, maxLines = 1, softWrap = false)
    }
}

/** An error region (spec 5.0): the sentence in wipe red and Try again. */
@Composable
fun ErrorRegion(message: String, retry: () -> Unit) {
    Column(verticalArrangement = Arrangement.spacedBy(12.dp), modifier = Modifier.raidQuietCard().padding(16.dp)) {
        Text(message, style = FsType.body, color = FsColor.wipe)
        RaidSecondaryButton(stringResource(R.string.common_tryAgain), retry)
    }
}
