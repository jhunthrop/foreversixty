package gg.foreversixty.app.account

import androidx.annotation.DrawableRes
import androidx.compose.foundation.BorderStroke
import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.CheckCircle
import androidx.compose.material.icons.outlined.Info
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.Icon
import androidx.compose.material3.OutlinedButton
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.TextUnit
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.em
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.BodyFamily
import gg.foreversixty.app.designsystem.DisplayFamily
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsFontWeight
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpacing
import gg.foreversixty.app.designsystem.FsTextSize
import gg.foreversixty.app.designsystem.HeadingFamily

/** The spec's type roles (§4.2), in `sp` so each scales with the system font size. */
object AccountType {
    val title = TextStyle(fontFamily = HeadingFamily, fontWeight = FsFontWeight.headingHeavy, fontSize = FsTextSize.h1Phone)
    val h2 = TextStyle(fontFamily = HeadingFamily, fontWeight = FsFontWeight.headingHeavy, fontSize = FsTextSize.h2)
    val sec = TextStyle(fontFamily = HeadingFamily, fontWeight = FsFontWeight.headingHeavy, fontSize = FsTextSize.cardTitle)
    val name = TextStyle(fontFamily = BodyFamily, fontWeight = FsFontWeight.bodyBold, fontSize = 17.sp)
    val body = TextStyle(fontFamily = BodyFamily, fontWeight = FsFontWeight.bodyRegular, fontSize = FsTextSize.body, lineHeight = FsTextSize.bodyLineHeight)
    val bodyStrong = body.copy(fontWeight = FsFontWeight.bodySemibold)
    val meta = TextStyle(fontFamily = BodyFamily, fontWeight = FsFontWeight.bodyRegular, fontSize = FsTextSize.meta)
    val micro = TextStyle(fontFamily = BodyFamily, fontWeight = FsFontWeight.bodyBold, fontSize = FsTextSize.micro, letterSpacing = 0.14.em)
    val button = TextStyle(fontFamily = BodyFamily, fontWeight = FsFontWeight.bodyBold, fontSize = 13.sp, letterSpacing = 0.07.em)
    val wordmark = TextStyle(fontFamily = DisplayFamily, fontWeight = FontWeight.Bold, fontSize = 22.sp, letterSpacing = 0.1.em)

    /** The selector name stops growing at 22 sp (spec §3.4). */
    @Composable
    fun cappedName(): TextStyle {
        val scale = LocalDensity.current.fontScale
        val size: TextUnit = ((NAME_SIZE * scale).coerceAtMost(NAME_CEILING) / scale).sp
        return name.copy(fontSize = size)
    }

    private const val NAME_SIZE = 17f
    private const val NAME_CEILING = 22f
}

/** The class colours by the API's class word. */
fun classColor(className: String?): Color = when (className?.lowercase()) {
    "warrior" -> FsColor.classWarrior
    "paladin" -> FsColor.classPaladin
    "hunter" -> FsColor.classHunter
    "rogue" -> FsColor.classRogue
    "priest" -> FsColor.classPriest
    "shaman" -> FsColor.classShaman
    "mage" -> FsColor.classMage
    "warlock" -> FsColor.classWarlock
    "druid" -> FsColor.classDruid
    else -> FsColor.text
}

/** `@drawable/crest_<class>`, copied from web/public by the forever.brand-art plugin. */
@DrawableRes
private fun crestOf(className: String?): Int? = when (className?.lowercase()) {
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

/** The circular ringed HD crest (spec §4.3): never a square class icon. */
@Composable
fun ClassCrest(className: String?, size: Dp, modifier: Modifier = Modifier) {
    val color = classColor(className)
    Box(
        modifier
            .size(size)
            .shadow(7.dp, CircleShape, ambientColor = color, spotColor = color)
            .clip(CircleShape)
            .background(FsColor.raised)
            .border(2.dp, color.copy(alpha = RING_ALPHA), CircleShape),
    ) {
        crestOf(className)?.let {
            Image(painterResource(it), contentDescription = null, contentScale = ContentScale.Crop, modifier = Modifier.fillMaxSize())
        }
    }
}

private const val RING_ALPHA = 0.6f

/** The LX seal. */
@Composable
fun Seal(size: Dp) {
    Image(painterResource(R.drawable.brand_seal), contentDescription = null, modifier = Modifier.size(size))
}

/** The hero band ground (`--surface-band`): the ember radial over the warm dark. */
fun Modifier.bandBackground(): Modifier = background(FsColor.bg).background(
    Brush.radialGradient(listOf(FsColor.ember.copy(alpha = 0.28f), Color.Transparent), radius = 1_400f, center = androidx.compose.ui.geometry.Offset(600f, 0f)),
)

/** The toast at the top of F1.1: failures in wipe red, a finished deletion neutral. */
@Composable
fun NoticeToast(text: String, modifier: Modifier = Modifier, isFailure: Boolean = true) {
    val border = if (isFailure) FsColor.wipe.copy(alpha = 0.45f) else FsColor.line
    Row(
        modifier
            .fillMaxWidth()
            .background(FsColor.raised, RoundedCornerShape(FsRadius.panel))
            .border(1.dp, border, RoundedCornerShape(FsRadius.panel))
            .padding(FsSpacing.space4)
            .semantics { liveRegion = LiveRegionMode.Assertive },
        horizontalArrangement = Arrangement.spacedBy(FsSpacing.space3),
    ) {
        Icon(
            if (isFailure) Icons.Outlined.Info else Icons.Outlined.CheckCircle,
            contentDescription = null,
            tint = if (isFailure) FsColor.wipe else FsColor.gold,
        )
        Text(text, style = AccountType.body, color = FsColor.strong)
    }
}

/** The site's buttons (spec §4.5), full width. */
enum class BrandButtonKind { Primary, Secondary, Destructive, Quiet }

@Composable
fun BrandButton(
    label: String,
    kind: BrandButtonKind,
    onClick: () -> Unit,
    modifier: Modifier = Modifier,
    busy: Boolean = false,
    height: Dp = 52.dp,
    fill: Boolean = true,
) {
    val shape = RoundedCornerShape(FsRadius.control)
    val size = if (fill) modifier.fillMaxWidth().heightIn(min = height) else modifier.heightIn(min = height)
    val content: @Composable () -> Unit = {
        if (busy) {
            CircularProgressIndicator(Modifier.size(20.dp), color = contentColor(kind), strokeWidth = 2.dp)
        } else {
            Text(label.uppercase(), style = AccountType.button, textAlign = TextAlign.Center)
        }
    }
    if (kind == BrandButtonKind.Primary) {
        Button(
            onClick = onClick,
            enabled = !busy,
            shape = shape,
            modifier = size.background(Brush.verticalGradient(listOf(FsColor.goldHover, FsColor.gold)), shape),
            colors = ButtonDefaults.buttonColors(
                containerColor = Color.Transparent,
                contentColor = FsColor.bg,
                disabledContainerColor = Color.Transparent,
                disabledContentColor = FsColor.bg,
            ),
        ) { content() }
    } else {
        OutlinedButton(
            onClick = onClick,
            enabled = !busy,
            shape = shape,
            modifier = size,
            border = BorderStroke(1.dp, borderColor(kind)),
            colors = ButtonDefaults.outlinedButtonColors(
                containerColor = Color.White.copy(alpha = 0.045f),
                contentColor = contentColor(kind),
                disabledContentColor = contentColor(kind),
            ),
        ) { content() }
    }
}

private fun contentColor(kind: BrandButtonKind): Color = when (kind) {
    BrandButtonKind.Primary -> FsColor.bg
    BrandButtonKind.Secondary -> FsColor.strong
    BrandButtonKind.Destructive -> FsColor.wipe
    BrandButtonKind.Quiet -> FsColor.nav
}

private fun borderColor(kind: BrandButtonKind): Color = when (kind) {
    BrandButtonKind.Destructive -> FsColor.wipe.copy(alpha = 0.45f)
    else -> FsColor.gold.copy(alpha = 0.17f)
}

