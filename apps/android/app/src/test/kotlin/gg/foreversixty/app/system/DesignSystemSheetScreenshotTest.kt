package gg.foreversixty.app.system

import android.graphics.BitmapFactory
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.Shadow
import androidx.compose.ui.graphics.asImageBitmap
import androidx.compose.ui.graphics.painter.BitmapPainter
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onRoot
import androidx.compose.ui.unit.dp
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.captureRoboImage
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.designsystem.FrameCorners
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.Label12
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.PanelFrame
import gg.foreversixty.app.designsystem.PaperdollHeader
import gg.foreversixty.app.designsystem.PaperdollMedallion
import gg.foreversixty.app.designsystem.PaperdollListSlot
import gg.foreversixty.app.designsystem.ScaledType
import gg.foreversixty.app.designsystem.StaleStrip
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.panel
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.designsystem.screenGround
import gg.foreversixty.app.designsystem.surfaceBand
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.items.ItemIcon
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.navigation.TabBadges
import gg.foreversixty.app.testing.STRICT_SNAPSHOT
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode
import java.io.File

/**
 * Lane M0's design-system sheet at Pixel 9's width: every round-5 surface, button, primitive, the
 * paperdoll header and the tab bar (the shell's own) on the painted ground, with the mock's content
 * (design/mocks/directions/mobile.py). The iOS sheet is AppTests/System/DesignSystemSheet.swift;
 * the PNG is the lane's review capture (design/reviews/2026-11-10-m0).
 * Record: ./gradlew :app:recordRoborazziDebug   Verify: ./gradlew :app:verifyRoborazziDebug
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9_SHEET)
class DesignSystemSheetScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    @Test
    fun `the round-5 design system on one sheet`() {
        val portrait = portrait()
        compose.setContent {
            ForeverTheme {
                CompositionLocalProvider(LocalItemIconPainter provides HomeFixtures.repoIcons) {
                    ForeverTabScaffold(selected = ForeverTab.Gear, onSelect = {}, badges = TabBadges(guild = 1, activity = 3)) {
                        Sheet(portrait)
                    }
                }
            }
        }
        compose.onRoot().captureRoboImage("src/test/screenshots/design-system-sheet.png", roborazziOptions = STRICT_SNAPSHOT)
    }

    private fun portrait(): Painter? = File(HomeFixtures.repoRoot, "design/mocks/data/avatars/thandrel.jpg")
        .takeIf { it.exists() }
        ?.let { BitmapFactory.decodeFile(it.path) }
        ?.let { BitmapPainter(it.asImageBitmap()) }
}

private val BUILD: String = File(HomeFixtures.repoRoot, "web/src/data/active-build.json").readText()
    .substringAfter("\"build\"").substringAfter('"').substringBefore('"')

@Composable
private fun Sheet(portrait: Painter?) {
    Column(Modifier.fillMaxSize().screenGround()) {
        Bar(portrait)
        Column(Modifier.padding(horizontal = 14.dp, vertical = FsSpace.space4), verticalArrangement = Arrangement.spacedBy(FsSpace.space5)) {
            Buttons()
            Panels()
            Tags()
            PaperdollHeader(
                title = "Fury Warrior",
                figure = "1004.2",
                portrait = portrait,
                slots = SHEET_SLOTS,
                selected = "head",
                onSelect = {},
            ) { item -> ItemIcon(SHEET_GEAR[item.id].icon, SHEET_GEAR[item.id].quality, BUILD, 38.dp) }
        }
    }
}

/** The band behind the top app bar: the profile medallion and the title in Marcellus (M3's bar). */
@Composable
private fun Bar(portrait: Painter?) {
    Row(
        Modifier.fillMaxWidth().surfaceBand().padding(horizontal = FsSpace.space4, vertical = FsSpace.space4),
        horizontalArrangement = Arrangement.spacedBy(FsSpace.space3),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        PaperdollMedallion(portrait, 32.dp)
        Text("Gear", style = FsType.title.copy(fontSize = ScaledType.cardHead.fontSize, shadow = TITLE_SHADOW), color = FsColor.cream)
    }
}

@Composable
private fun Buttons() {
    Column(verticalArrangement = Arrangement.spacedBy(FsSpace.space3)) {
        Row(horizontalArrangement = Arrangement.spacedBy(FsSpace.space2)) {
            BrandButton("I'm in", BrandButtonKind.Primary, {}, Modifier.weight(1f), height = 48.dp)
            BrandButton("Late", BrandButtonKind.Secondary, {}, Modifier.weight(1f), height = 48.dp)
            BrandButton("Can't come", BrandButtonKind.Secondary, {}, Modifier.weight(1f), height = 48.dp)
        }
        Row(horizontalArrangement = Arrangement.spacedBy(FsSpace.space2)) {
            BrandButton("Turn on", BrandButtonKind.Accent, {}, Modifier.weight(1f), height = 48.dp)
            BrandButton("Not now", BrandButtonKind.Quiet, {}, Modifier.weight(1f), height = 48.dp)
            BrandButton("Leave guild", BrandButtonKind.Destructive, {}, Modifier.weight(1f), height = 48.dp)
        }
    }
}

@Composable
private fun Panels() {
    Column(verticalArrangement = Arrangement.spacedBy(FsSpace.space4)) {
        Box(Modifier.fillMaxWidth().panel(PanelFrame.Ornate, cave = true)) {
            Column(Modifier.padding(18.dp)) {
                Text("Onyxia's Lair", style = FsType.sectionReading, color = FsColor.cream)
                Text("Wed 6 Jan · 7:30 pm", style = ScaledType.meta, color = FsColor.text)
                Row(Modifier.padding(top = FsSpace.space2), horizontalArrangement = Arrangement.spacedBy(18.dp)) {
                    Figure("2", "days")
                    Figure("4", "hours")
                    Figure("12", "min")
                }
            }
            FrameCorners(Modifier.matchParentSize())
        }
        Row(horizontalArrangement = Arrangement.spacedBy(FsSpace.space3)) {
            Text(
                "Your guild's sims",
                Modifier.weight(1f).panel(PanelFrame.Plain).padding(FsSpace.space4),
                style = ScaledType.bodyStrong,
                color = FsColor.strong,
            )
            Text("Updated 2 h ago", Modifier.weight(1f).quietCard().padding(FsSpace.space4), style = ScaledType.meta, color = FsColor.muted)
        }
    }
}

@Composable
private fun Figure(value: String, unit: String) {
    Column {
        Text(value, style = FsType.figureBig, color = FsColor.strong)
        Text(unit, style = ScaledType.meta, color = FsColor.muted)
    }
}

@Composable
private fun Tags() {
    Column(verticalArrangement = Arrangement.spacedBy(FsSpace.space3)) {
        Row(horizontalArrangement = Arrangement.spacedBy(FsSpace.space2)) {
            Tag("Officer", FsColor.gold)
            Tag("Killed", FsColor.kill)
            Tag("Wiped", FsColor.wipe)
            Tag("In", FsColor.kill, checked = true)
        }
        Label12("New for you", color = FsColor.mutedGround)
        OfflineStrip("Offline · showing Tuesday's raid")
        StaleStrip("Couldn't refresh the roster", action = "Retry", onAction = {})
    }
}

private val TITLE_SHADOW = Shadow(Color.Black.copy(alpha = 0.45f), Offset(0f, 1f), 2f)

private class SheetGear(val slot: String, val label: String, val name: String, val icon: String, val quality: Int, val gain: String?)

/** The mock's Fury list (design/mocks/ws_app_flows.py `DOLL`), Greenhammer left out until the first raids. */
private val SHEET_GEAR = listOf(
    SheetGear("head", "Head", "Lionheart Helm", "inv_helmet_36", 4, "+14.5"),
    SheetGear("neck", "Neck", "Skibi's Pendant", "inv_jewelry_amulet_07", 3, null),
    SheetGear("shoulder", "Shoulder", "Theramore Pauldrons", "inv_shoulder_06", 4, null),
    SheetGear("back", "Back", "Howler's Furs", "inv_misc_cape_10", 3, "+2.4"),
    SheetGear("chest", "Chest", "Timbermaw Tunic", "inv_chest_fur", 4, "+5.3"),
    SheetGear("wrist", "Wrist", "Forest Stalker's Bracers", "inv_bracer_12", 4, "+0.1"),
    SheetGear("hands", "Hands", "Theramore Gauntlets", "inv_gauntlets_29", 4, "+1.0"),
    SheetGear("waist", "Waist", "Valiant's Waistguard", "inv_belt_33", 4, null),
    SheetGear("legs", "Legs", "Sentinel's Plate Legguards", "inv_pants_plate_20", 4, "+4.7"),
    SheetGear("feet", "Feet", "Scalegut Treaders", "inv_boots_chain_11", 3, "+0.9"),
    SheetGear("finger1", "Finger", "Don Julio's Band", "inv_jewelry_ring_35", 4, "+6.8"),
    SheetGear("finger2", "Finger", "Cutthroat's Signet", "inv_jewelry_ring_01", 3, "+0.8"),
    SheetGear("trinket1", "Trinket", "Hand of Justice", "inv_jewelry_talisman_01", 3, null),
    SheetGear("trinket2", "Trinket", "Molten Heart of the Mountain", "inv_misc_gem_ruby_01", 3, "+3.2"),
    SheetGear("main_hand", "Main Hand", "Ironwood Blade", "inv_sword_07", 4, "+5.9"),
    SheetGear("off_hand", "Off Hand", "Persuader", "inv_hammer_08", 4, null),
    SheetGear("ranged", "Ranged", "Blackcrow", "inv_weapon_crossbow_04", 3, "+1.9"),
)

private val SHEET_SLOTS = SHEET_GEAR.mapIndexed { index, gear ->
    PaperdollListSlot(
        id = gear.slot,
        label = gear.label,
        item = PaperdollListSlot.Item(
            id = index,
            name = gear.name,
            quality = ItemQuality(gear.quality).text,
            detail = gear.gain?.let { "$it DPS" } ?: "Equipped",
            isGain = gear.gain != null,
        ),
    )
}

/** Pixel 9's width, tall enough for the whole sheet above the shell's navigation bar. */
private const val PIXEL_9_SHEET = "w412dp-h1340dp-normal-long-notround-any-420dpi-keyshidden-nonav"
