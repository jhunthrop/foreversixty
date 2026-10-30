package gg.foreversixty.app.raids.ui

import androidx.compose.foundation.Image
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
import gg.foreversixty.app.designsystem.BodyFamily
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsTextSize
import gg.foreversixty.app.designsystem.HeadingFamily
import gg.foreversixty.app.raids.logic.RaidChip
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidFacts
import java.time.Instant

/** The spec's type roles (4.2), in sp so they follow the font scale. */
object RaidType {
    val pageTitle = TextStyle(fontFamily = HeadingFamily, fontWeight = FontWeight.ExtraBold, fontSize = FsTextSize.h1Phone)
    val h2 = TextStyle(fontFamily = HeadingFamily, fontWeight = FontWeight.ExtraBold, fontSize = FsTextSize.h2)
    val stat = TextStyle(fontFamily = HeadingFamily, fontWeight = FontWeight.ExtraBold, fontSize = FsTextSize.stat)
    val sec = TextStyle(fontFamily = HeadingFamily, fontWeight = FontWeight.ExtraBold, fontSize = FsTextSize.cardTitle)
    val cardHead = TextStyle(fontFamily = HeadingFamily, fontWeight = FontWeight.ExtraBold, fontSize = 22.sp)
    val name = TextStyle(fontFamily = BodyFamily, fontWeight = FontWeight.Bold, fontSize = 17.sp)
    val body = TextStyle(fontFamily = BodyFamily, fontWeight = FontWeight.Normal, fontSize = FsTextSize.body, lineHeight = FsTextSize.bodyLineHeight)
    val bodyStrong = body.copy(fontWeight = FontWeight.SemiBold)
    val meta = TextStyle(fontFamily = BodyFamily, fontWeight = FontWeight.Normal, fontSize = FsTextSize.meta)
    val micro = TextStyle(fontFamily = BodyFamily, fontWeight = FontWeight.Bold, fontSize = FsTextSize.micro, letterSpacing = 1.5.sp)
    val segment = TextStyle(fontFamily = BodyFamily, fontWeight = FontWeight.SemiBold, fontSize = 14.sp)
    val button = TextStyle(fontFamily = BodyFamily, fontWeight = FontWeight.Bold, fontSize = 13.sp, letterSpacing = 0.9.sp)

    /** Times, counts and countdowns (JetBrains Mono is not bundled yet; the system mono stands in). */
    val mono = TextStyle(fontFamily = FontFamily.Monospace, fontSize = FsTextSize.body)
    val monoMeta = TextStyle(fontFamily = FontFamily.Monospace, fontSize = FsTextSize.meta)
    val monoDay = TextStyle(fontFamily = FontFamily.Monospace, fontWeight = FontWeight.SemiBold, fontSize = 20.sp)
}

object RaidColors {
    val strip = Color(0xFF11151F)
    val sheet = Color(0xFF10141D)
    val rowWash = FsColor.gold.copy(alpha = 0.05f)
    val pressed = FsColor.gold.copy(alpha = 0.10f)
    val iconBorder = Color(0xFF3A4256)

    fun classColor(slug: String): Color = when (slug.lowercase()) {
        "warrior" -> FsColor.classWarrior
        "paladin" -> FsColor.classPaladin
        "hunter" -> FsColor.classHunter
        "rogue" -> FsColor.classRogue
        "priest" -> FsColor.classPriest
        "shaman" -> FsColor.classShaman
        "mage" -> FsColor.classMage
        "warlock" -> FsColor.classWarlock
        "druid" -> FsColor.classDruid
        else -> FsColor.muted
    }

    fun qualityBorder(quality: Int): Color = when (quality) {
        0 -> FsColor.rarityPoor
        1 -> FsColor.rarityCommon
        2 -> FsColor.rarityUncommon
        3 -> FsColor.rarityRare
        4 -> FsColor.rarityEpic
        else -> FsColor.rarityLegendary
    }

    fun qualityText(quality: Int): Color = when (quality) {
        3 -> FsColor.rarityRareText
        4 -> FsColor.rarityEpicText
        else -> qualityBorder(quality)
    }

    fun parse(percentile: Int): Color = when {
        percentile >= 100 -> FsColor.parseGold
        percentile == 99 -> FsColor.parsePink
        percentile >= 95 -> FsColor.parseOrange
        percentile >= 75 -> FsColor.parsePurple
        percentile >= 50 -> FsColor.parseBlue
        percentile >= 25 -> FsColor.parseGreen
        else -> FsColor.parseGreyText
    }
}

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

/** Where the game icons come from: `/data/{build}/icons/{stem}.webp` on the site. */
data class RaidIcons(val build: String?) {
    fun url(stem: String?): String? =
        if (build == null || stem.isNullOrEmpty()) null else "https://foreversixty.gg/data/$build/icons/$stem.webp"
}

val LocalRaidIcons = staticCompositionLocalOf { RaidIcons(null) }

/** The ringed HD class crest (spec 4.3), bundled from web/public/icons/hd/crests. */
@Composable
fun Crest(className: String, size: Dp) {
    val color = RaidColors.classColor(className)
    val id = crestDrawable(className)
    Box(
        Modifier
            .size(size)
            .shadow(if (size > 32.dp) 7.dp else 4.dp, CircleShape, ambientColor = color, spotColor = color)
            .clip(CircleShape)
            .background(FsColor.raised)
            .border(2.dp, color.copy(alpha = 0.6f), CircleShape),
    ) {
        id?.let { Image(painterResource(it), contentDescription = null, contentScale = ContentScale.Crop, modifier = Modifier.size(size)) }
    }
}

/** The crest drawables the forever.class-crests plugin copies in. */
private fun crestDrawable(className: String): Int? = when (className.lowercase()) {
    "warrior" -> R.drawable.crest_warrior
    "paladin" -> R.drawable.crest_paladin
    "hunter" -> R.drawable.crest_hunter
    "rogue" -> R.drawable.crest_rogue
    "priest" -> R.drawable.crest_priest
    "shaman" -> R.drawable.crest_shaman
    "mage" -> R.drawable.crest_mage
    "warlock" -> R.drawable.crest_warlock
    "druid" -> R.drawable.crest_druid
    else -> null
}

/** A square game icon: the spec icon (border `#3a4256`) or an item icon (quality border). */
@Composable
fun GameIcon(stem: String?, size: Dp, border: Color = RaidColors.iconBorder, borderWidth: Dp = 1.dp) {
    val shape = RoundedCornerShape(FsRadius.pill)
    Box(
        Modifier.size(size).clip(shape).background(FsColor.cardTop).border(borderWidth, border, shape),
    ) {
        LocalRaidIcons.current.url(stem)?.let { url ->
            AsyncImage(model = url, contentDescription = null, contentScale = ContentScale.Crop, modifier = Modifier.size(size))
        }
    }
}

/** `micro` uppercase section label, with an optional mono count. */
@Composable
fun SectionLabel(text: String, color: Color = FsColor.muted, trailing: String? = null) {
    val tracking = if (isLargeText()) 0.sp else RaidType.micro.letterSpacing
    Row(horizontalArrangement = Arrangement.spacedBy(6.dp), modifier = Modifier.semantics { heading() }) {
        Text(text.uppercase(), style = RaidType.micro.copy(letterSpacing = tracking), color = color, maxLines = 1, softWrap = false)
        trailing?.let { Text(it, style = RaidType.monoMeta, color = color) }
    }
}

/** The status tag. */
@Composable
fun StatusTag(text: String, color: Color) {
    Text(
        text.uppercase(),
        style = RaidType.micro.copy(letterSpacing = 0.8.sp),
        color = color,
        maxLines = 1,
        softWrap = false,
        modifier = Modifier
            .heightIn(min = 22.dp)
            .background(color.copy(alpha = 0.12f), RoundedCornerShape(FsRadius.pill))
            .border(1.dp, color.copy(alpha = 0.45f), RoundedCornerShape(FsRadius.pill))
            .padding(horizontal = 7.dp, vertical = 4.dp),
    )
}

@Composable
fun ChipTag(chip: RaidChip) {
    val color = when (chip) {
        RaidChip.Planned -> FsColor.gold
        RaidChip.Cancelled -> FsColor.wipe
        RaidChip.Done -> FsColor.kill
        RaidChip.Locked, RaidChip.Ended -> FsColor.muted
    }
    StatusTag(stringResource(chip.label), color)
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
        Text(weekday, style = RaidType.micro.copy(letterSpacing = 0.sp), color = FsColor.muted)
        Text(day, style = RaidType.monoDay, color = FsColor.strong)
        Text(month, style = RaidType.micro.copy(letterSpacing = 0.sp), color = FsColor.muted)
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
            Text(line, style = RaidType.bodyStrong, color = FsColor.strong)
            facts.fill?.let { FillBar(it, null, muted) }
        } else {
            Row(horizontalArrangement = Arrangement.spacedBy(12.dp), verticalAlignment = Alignment.CenterVertically) {
                facts.fill?.let { FillBar(it, 120.dp, muted) }
                Text(line, style = RaidType.bodyStrong, color = FsColor.strong)
            }
        }
        Text(breakdown, style = RaidType.meta, color = FsColor.muted)
    }
}

/** The offline strip (spec 5.0): 36 dp, the glyph, `Offline. Showing what you had at {time}.` */
@Composable
fun OfflineStrip(text: String) {
    Row(
        horizontalArrangement = Arrangement.spacedBy(8.dp),
        verticalAlignment = Alignment.CenterVertically,
        modifier = Modifier
            .fillMaxWidth()
            .heightIn(min = 36.dp)
            .background(RaidColors.strip)
            .padding(horizontal = 16.dp)
            .semantics(mergeDescendants = true) {},
    ) {
        Icon(Icons.Filled.Warning, contentDescription = null, tint = FsColor.nav, modifier = Modifier.size(16.dp))
        Text(text, style = RaidType.meta, color = FsColor.nav)
    }
    Box(Modifier.fillMaxWidth().size(1.dp).background(FsColor.line))
}

/** The primary card surface (`--surface-card`). */
fun Modifier.primaryCard(): Modifier = this
    .fillMaxWidth()
    .clip(RoundedCornerShape(FsRadius.panel))
    .background(Brush.verticalGradient(listOf(Color(0xFF1A1712), FsColor.raised)))
    .border(1.dp, FsColor.lineCard, RoundedCornerShape(FsRadius.panel))
    .padding(16.dp)

/** A quiet card: raised, line-soft border. */
fun Modifier.quietCard(): Modifier = this
    .fillMaxWidth()
    .clip(RoundedCornerShape(FsRadius.panel))
    .background(FsColor.raised)
    .border(1.dp, FsColor.lineSoft, RoundedCornerShape(FsRadius.panel))

/** The secondary button (4.5% white, gold 17% border, 48 dp, 13 sp 700 uppercase). */
@Composable
fun SecondaryButton(
    text: String,
    onClick: () -> Unit,
    enabled: Boolean = true,
    gold: Boolean = false,
    leading: (@Composable () -> Unit)? = null,
) {
    OutlinedButton(
        onClick = onClick,
        enabled = enabled,
        shape = RoundedCornerShape(FsRadius.control),
        border = androidx.compose.foundation.BorderStroke(1.dp, if (gold) FsColor.gold else FsColor.lineCard),
        colors = ButtonDefaults.outlinedButtonColors(
            containerColor = Color.White.copy(alpha = 0.045f),
            contentColor = if (gold) FsColor.gold else FsColor.strong,
        ),
        modifier = Modifier.heightIn(min = 48.dp),
    ) {
        leading?.let {
            it()
            Box(Modifier.width(8.dp))
        }
        Text(text.uppercase(), style = RaidType.button, maxLines = 1, softWrap = false)
    }
}

/** The gold text link, 48 dp target. */
@Composable
fun GoldLink(text: String, onClick: () -> Unit, trailing: Boolean = false) {
    Row(
        verticalAlignment = Alignment.CenterVertically,
        modifier = Modifier.heightIn(min = 48.dp).clickable(onClick = onClick),
    ) {
        Text(text + if (trailing) " ›" else "", style = RaidType.bodyStrong, color = FsColor.gold, maxLines = 1, softWrap = false)
    }
}

/** An error region (spec 5.0): the sentence in wipe red and Try again. */
@Composable
fun ErrorRegion(message: String, retry: () -> Unit) {
    Column(verticalArrangement = Arrangement.spacedBy(12.dp), modifier = Modifier.quietCard().padding(16.dp)) {
        Text(message, style = RaidType.body, color = FsColor.wipe)
        SecondaryButton(stringResource(R.string.common_tryAgain), retry)
    }
}
