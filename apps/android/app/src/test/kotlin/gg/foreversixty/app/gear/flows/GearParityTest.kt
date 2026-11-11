package gg.foreversixty.app.gear.flows

import android.graphics.BitmapFactory
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.asImageBitmap
import androidx.compose.ui.graphics.painter.BitmapPainter
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.unit.dp
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.bis.BisItemRef
import gg.foreversixty.app.bis.BisList
import gg.foreversixty.app.bis.BisPick
import gg.foreversixty.app.bis.BisRow
import gg.foreversixty.app.bis.BisRunnerUp
import gg.foreversixty.app.bis.BisUnit
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.parityContent
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.gear.ItemSource
import gg.foreversixty.app.gear.RunPick
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.items.ItemTooltipModel
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.menu.LocalOpenMenu
import gg.foreversixty.app.menu.CharacterBar
import gg.foreversixty.app.menu.LocalPortraitPainter
import gg.foreversixty.app.menu.MenuFixtures
import gg.foreversixty.app.menu.ProfileBar
import gg.foreversixty.app.model.ItemHolder
import gg.foreversixty.app.model.ItemNeed
import gg.foreversixty.app.model.ItemNeeds
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.navigation.TabBadges
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.testing.ParityCheck
import org.junit.Assert.assertEquals
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode
import java.io.File

/** Gear (design/mocks/ws_app_flows.py `gear`, `item_sheet`; round 5 on Android) inside the tab shell. */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class GearParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun gear() = ParityCheck.state(compose, "gear") { Tab { Gear() } }

    @Test fun gearAt2() = ParityCheck.state(compose, "gear", fontScale = 2f) { Tab { Gear() } }

    @Test fun gearOut() = ParityCheck.state(compose, "gear-out") {
        Tab {
            Column {
                ProfileBar(null)
                GearSignedOut(GearFixtures.classes) {}
            }
        }
    }

    @Test fun itemSheet() = ParityCheck.state(compose, "item-sheet") {
        Icons {
            Column(Modifier.fillMaxSize().parityContent()) {
                Spacer(Modifier.fillMaxWidth().height(DIM))
                Column(Modifier.fillMaxWidth().background(FsColor.raised, RoundedCornerShape(topStart = 16.dp, topEnd = 16.dp))) {
                    Box(Modifier.fillMaxWidth().padding(top = 9.dp, bottom = 9.dp), Alignment.Center) {
                        Box(Modifier.size(36.dp, 5.dp).background(FsColor.line, RoundedCornerShape(3.dp)))
                    }
                    GearSheetBody(GearFixtures.sheet)
                }
            }
        }
    }

    @Test fun placesGroupInstancesThenProfessionsThenReputations() {
        val composer = GearComposer(ApplicationProvider.getApplicationContext<android.content.Context>().resources)
        val rows = composer.where(GearFixtures.picks, GearFixtures.list)
        assertEquals(listOf("Blackrock Spire", "Blacksmithing", "Stormpike Guard"), rows.map { it.title })
        assertEquals("Blackcrow · Shadow Hunter Vosh'gajin, 8%", rows[0].line)
        assertEquals("Lionheart Helm, Greenhammer, Ironwood Blade · ask a blacksmith", rows[1].line)
        assertEquals("Don Julio's Band, needs revered", rows[2].line)
    }

    @Test fun theListLineNamesTheBandTheFactionAndTheEquippedCount() {
        val composer = GearComposer(ApplicationProvider.getApplicationContext<android.content.Context>().resources)
        val content = composer.content(GearFixtures.list, "alliance", emptyList(), emptyList(), setDps = 1004.2)
        assertEquals("Fury Warrior", content.title)
        assertEquals("60 raid · Alliance · 1 of 3 slots match the list", content.facts)
    }

    @Test fun theHeaderFigureIsTheBandsSetFigureInTheHouseFormatAndNoneForATank() {
        val composer = GearComposer(ApplicationProvider.getApplicationContext<android.content.Context>().resources)
        assertEquals("1,004", composer.content(GearFixtures.list, "alliance", emptyList(), emptyList(), setDps = 1004.2).figure)
        assertEquals("182", composer.content(GearFixtures.list, "alliance", emptyList(), emptyList(), setDps = 182.0).figure)
        assertEquals(null, composer.content(GearFixtures.list.copy(unit = BisUnit.Score), "alliance", emptyList(), emptyList(), setDps = 1004.2).figure)
    }

    @Test fun theSheetNamesWhoNeedsItAndWhoHasIt() {
        val res = ApplicationProvider.getApplicationContext<android.content.Context>().resources
        val row = GearSlotRow("head", "Head", ItemLook(2, "Lionheart Helm", 4), "Head", "+14.5", worn = false, row = GearFixtures.list.rows.first())
        val sheet = GearSheetComposer(res, BUILD).content(row, "DPS", null, emptyList(), GearFixtures.needs)
        assertEquals(listOf("Protection Warrior · needs it, the pick on their list", "Enhancement Shaman · has it"), sheet.guild?.map { it.line })
        assertEquals(null, sheet.guildNote)
        assertEquals("+14.5 DPS", sheet.gain)
    }

    @Test fun theSourceLineNamesTheBossThePlaceItsLevelsAndTheChance() {
        val composer = GearSheetComposer(ApplicationProvider.getApplicationContext<android.content.Context>().resources, BUILD)
        assertEquals("Skum, Wailing Caverns (15-24), 50%", composer.sourceLine(ItemSource("dungeon", "Wailing Caverns", 50.0, boss = "Skum", levels = 15..24)))
        assertEquals("Onyxia, Onyxia's Lair, 8%", composer.sourceLine(ItemSource("raid", "Onyxia's Lair", 8.0, boss = "Onyxia", levels = 60..60)))
        assertEquals("Stormpike Guard, revered", composer.sourceLine(ItemSource("rep", "Stormpike Guard", null, standing = "Revered")))
        // Round 8 row 1: the shared source words, with the figure when the data has one.
        assertEquals("Vendor: Pix Xizzix", composer.sourceLine(ItemSource("vendor", "Pix Xizzix", null)))
        assertEquals("Crafted: Blacksmithing (245)", composer.sourceLine(ItemSource("crafted", "Blacksmithing", null, profession = "blacksmithing", skill = 245)))
        assertEquals("Crafted: Blacksmithing", composer.sourceLine(ItemSource("crafted", "Blacksmithing", null, profession = "blacksmithing")))
        assertEquals("Quest: A Collection of Heads", composer.sourceLine(ItemSource("quest", "A Collection of Heads", null)))
        assertEquals("World drop", composer.sourceLine(ItemSource("world_drop", "World drop", null)))
    }

    @Test fun theSheetSaysWhereAPieceTheLootFileDoesNotPlaceComesFrom() {
        val res = ApplicationProvider.getApplicationContext<android.content.Context>().resources
        val crafted = GearFixtures.list.rows.first().let { row -> row.copy(pick = row.pick?.copy(source = "Crafted: Blacksmithing")) }
        val row = GearSlotRow("head", "Head", ItemLook(2, "Lionheart Helm", 4), "Head", "+14.5", worn = false, row = crafted)
        val composer = GearSheetComposer(res, BUILD)
        val where = composer.content(row, "DPS", null, emptyList(), null).sources.single()
        assertEquals("Crafted: Blacksmithing", where.title)
        // Round 8 row 1: no caption and no icon, so the row cannot read as an item.
        assertEquals(null, where.line)
        assertEquals(null, where.lead)
        // The loot file's places win over the list's line when it has them.
        val dropped = composer.content(row, "DPS", null, listOf(ItemSource("raid", "Onyxia's Lair", 8.0, boss = "Onyxia")), null).sources
        assertEquals(listOf("Onyxia, Onyxia's Lair, 8%"), dropped.map { it.title })
    }

    @Test fun aDropChanceIsAWholePercentAndUnderOneReadsLessThanOne() {
        // web/src/lib/items/drop-chance.ts dropPercentLabel.
        assertEquals("8%", GearComposer.chanceWord(8.0))
        assertEquals("1%", GearComposer.chanceWord(1.3))
        assertEquals("20%", GearComposer.chanceWord(19.6))
        assertEquals("3%", GearComposer.chanceWord(2.5))
        assertEquals("<1%", GearComposer.chanceWord(0.3))
    }

    @Test fun aZeroGainReadsNoGainAndTheRunnerUpAlwaysSaysWhatItGives() {
        val res = ApplicationProvider.getApplicationContext<android.content.Context>().resources
        val alt = BisRunnerUp(BisItemRef(5, "Blood Guard's Plate Helm", null), "PvP rank 11 · Blood Guard · Horde", dpsDelta = 0.0, isWorn = false)
        val bisRow = GearFixtures.list.rows.first().let { row -> row.copy(pick = row.pick?.copy(runnersUp = listOf(alt))) }
        val row = GearSlotRow("head", "Head", ItemLook(2, "Lionheart Helm", 4), "Head", "+0.0", worn = false, row = bisRow, gains = false)
        val sheet = GearSheetComposer(res, BUILD).content(row, "DPS", null, emptyList(), null)
        assertEquals("No gain", sheet.gain)
        // Round 9 rows 1 and 3: `Same DPS`, and a PvP source without its faction word.
        assertEquals("Runner-up · PvP rank 11 · Blood Guard · Same DPS", sheet.runnerUp?.line)
        // Placed by the loot file: its boss and its chance, as iOS words it.
        val dropped = GearSheetComposer(res, BUILD).content(
            row, "DPS", null, emptyList(), null, ItemSource("dungeon", "Uldaman", 1.3, boss = "Ancient Treasure"),
        )
        assertEquals("Runner-up · Ancient Treasure · 1% · Same DPS", dropped.runnerUp?.line)
    }

    @Test fun theRunnerUpNamesItsSourceOnlyWhenItDiffersFromThePicks() {
        val res = ApplicationProvider.getApplicationContext<android.content.Context>().resources
        val composer = GearSheetComposer(res, BUILD)
        val alt = BisRunnerUp(BisItemRef(6, "Thorium Cestus", null), "Crafted: Blacksmithing", dpsDelta = -5.0, isWorn = false)
        val crafted = GearFixtures.list.rows.first().let { row -> row.copy(pick = row.pick?.copy(source = "Crafted: Blacksmithing", runnersUp = listOf(alt))) }
        val row = GearSlotRow("off_hand", "Off hand", ItemLook(2, "Dawn's Edge", 3), "Off hand", "+2.8", worn = false, row = crafted)
        val forge = ItemSource("crafted", "Blacksmithing", null, profession = "blacksmithing", skill = 275)
        // Same place as the pick's row (`Crafted: Blacksmithing (275)` above it): the delta alone.
        assertEquals("Runner-up · −5.0 DPS", composer.content(row, "DPS", null, listOf(forge), null).runnerUp?.line)
        assertEquals("Runner-up · −5.0 DPS", composer.content(row, "DPS", null, emptyList(), null).runnerUp?.line)
        // Placed by the loot file at the pick's own place: left out too.
        assertEquals("Runner-up · −5.0 DPS", composer.content(row, "DPS", null, listOf(forge), null, forge).runnerUp?.line)
        // A different place is named, in the source row's words.
        val drop = alt.copy(source = "World drop", dpsDelta = -0.3)
        val vendor = crafted.copy(pick = crafted.pick?.copy(source = "Vendor: Lady Palanseer", runnersUp = listOf(drop)))
        assertEquals("Runner-up · World drop · −0.3 DPS", composer.content(row.copy(row = vendor), "DPS", null, emptyList(), null).runnerUp?.line)
    }

    @Test fun theSheetSaysWhatThePickReplaces() {
        val res = ApplicationProvider.getApplicationContext<android.content.Context>().resources
        val row = GearSlotRow(
            "main_hand", "Main hand", ItemLook(17068, "Deathbringer", 4), "Main hand", "+12.6", worn = false,
            row = GearFixtures.list.rows.first(), equipped = ItemLook(7, "Ironwood Blade", 4),
        )
        val sheet = GearSheetComposer(res, BUILD).content(row, "DPS", null, emptyList(), null)
        assertEquals("Replaces Ironwood Blade", sheet.replaces?.title)
        assertEquals("Equipped", sheet.replaces?.line)
    }

    @Composable
    private fun Gear() {
        GearFlowScreen(GearFixtures.content, BUILD, GearActions(), portrait = portrait, topBar = { CharacterBar(MenuFixtures.thandrel, "Gear") })
    }

    @Composable
    private fun Tab(content: @Composable () -> Unit) {
        ForeverTabScaffold(selected = ForeverTab.Gear, onSelect = {}, badges = TabBadges(raid = 1, activity = 3)) {
            CompositionLocalProvider(LocalOpenMenu provides {}, LocalPortraitPainter provides MenuFixtures.repoPortraits) { Icons(content) }
        }
    }

    /** The game icons from the repo's build data, as Home's parity test paints them. */
    @Composable
    private fun Icons(content: @Composable () -> Unit) {
        CompositionLocalProvider(LocalItemIconPainter provides HomeFixtures.repoIcons) { content() }
    }

    private companion object {
        const val BUILD = "1.60.1.70790"
        val DIM = 120.dp

        /** Thandrel's portrait, the profile button's (design/mocks/data/avatars). */
        val portrait: Painter? = File(HomeFixtures.repoRoot, "design/mocks/data/avatars/thandrel.jpg")
            .takeIf { it.exists() }
            ?.let { BitmapFactory.decodeFile(it.path) }
            ?.let { BitmapPainter(it.asImageBitmap()) }
    }
}

/** The mock's Thandrel: the Fury Warrior raid list, his upgrades and where they come from. */
object GearFixtures {
    private data class Doll(val slot: String, val label: String, val name: String, val icon: String, val quality: Int, val source: String, val gain: String?)

    // Greenhammer unlocks with the first raids, so the off hand is the Persuader he wears (mock m-android-gear-frame).
    private val doll = listOf(
        Doll("head", "Head", "Lionheart Helm", "inv_helmet_36", 4, "Blacksmithing", "+14.5"),
        Doll("neck", "Neck", "Skibi's Pendant", "inv_jewelry_amulet_07", 3, "World drop", null),
        Doll("shoulder", "Shoulder", "Theramore Pauldrons", "inv_shoulder_06", 4, "Theramore Expeditionary Force", null),
        Doll("back", "Back", "Howler's Furs", "inv_misc_cape_10", 3, "Pix Xizzix, vendor", "+2.4"),
        Doll("chest", "Chest", "Timbermaw Tunic", "inv_chest_fur", 4, "Leatherworking", "+5.3"),
        Doll("wrist", "Wrist", "Forest Stalker's Bracers", "inv_bracer_12", 4, "Silverwing Sentinels", "+0.1"),
        Doll("hands", "Hands", "Theramore Gauntlets", "inv_gauntlets_29", 4, "Theramore Expeditionary Force", "+1.0"),
        Doll("waist", "Waist", "Valiant's Waistguard", "inv_belt_33", 4, "Pix Xizzix, vendor", null),
        Doll("legs", "Legs", "Sentinel's Plate Legguards", "inv_pants_plate_20", 4, "Illiyana Moonblaze, vendor", "+4.7"),
        Doll("feet", "Feet", "Scalegut Treaders", "inv_boots_chain_11", 3, "Leatherworking", "+0.9"),
        Doll("finger1", "Finger 1", "Don Julio's Band", "inv_jewelry_ring_35", 4, "Stormpike Guard", "+6.8"),
        Doll("finger2", "Finger 2", "Cutthroat's Signet", "inv_jewelry_ring_01", 3, "Pix Xizzix, vendor", "+0.8"),
        Doll("trinket1", "Trinket 1", "Hand of Justice", "inv_jewelry_talisman_01", 3, "Blackrock Depths: Emperor Dagran Thaurissan", null),
        Doll("trinket2", "Trinket 2", "Molten Heart of the Mountain", "inv_misc_gem_ruby_01", 3, "Enchanting", "+3.2"),
        Doll("main_hand", "Main hand", "Ironwood Blade", "inv_sword_07", 4, "Blacksmithing", "+5.9"),
        Doll("off_hand", "Off hand", "Persuader", "inv_hammer_08", 4, "Blacksmithing", null),
        Doll("ranged", "Ranged", "Blackcrow", "inv_weapon_crossbow_04", 3, "Blackrock Spire: Shadow Hunter Vosh'gajin, 8%", "+1.9"),
    )

    private fun look(i: Int, d: Doll) = ItemLook(1000 + i, d.name, d.quality, d.icon)

    val content = GearContent(
        title = "Fury Warrior",
        figure = "1,004",
        facts = "60 raid · Alliance · 5 of 17 slots match the list",
        where = listOf(
            GearWhereRow("Blackrock Spire", "Blackcrow · Shadow Hunter Vosh'gajin, 8%", ItemLook(1, "Blackcrow", 3, "inv_weapon_crossbow_04"), "1 upgrade · +1.9 DPS"),
            GearWhereRow("Blacksmithing", "Lionheart Helm, Ironwood Blade · ask a blacksmith", ItemLook(2, "Lionheart Helm", 4, "inv_helmet_36"), "2 upgrades · +20.4 DPS"),
            GearWhereRow("Stormpike Guard", "Don Julio's Band, needs revered", ItemLook(3, "Don Julio's Band", 4, "inv_jewelry_ring_35"), "1 upgrade · +6.8 DPS"),
        ),
        slots = doll.mapIndexed { i, d ->
            GearSlotRow(d.slot, d.label, look(i, d), "${d.label} · ${d.source}", d.gain, worn = d.gain == null, row = BisRow(d.slot, d.label, null, null))
        },
        unitWord = "DPS",
        header = GearHeader("+47.5 DPS from 12 upgrades", "synced 2 h ago", syncedStale = false, notModelled = null),
    )

    val classes = listOf("warrior", "paladin", "hunter", "rogue", "priest", "shaman", "mage", "warlock", "druid")
        .map { it to it.replaceFirstChar(Char::titlecase) }

    private const val BUILD = "1.60.1.70790"
    private val deathbringer = ItemLook(17068, "Deathbringer", 4, "inv_axe_09")
    private val composer = GearSheetComposer(ApplicationProvider.getApplicationContext<android.content.Context>().resources, BUILD)

    /** Deathbringer from Onyxia (the 1.60.1.70790 files: 8%, one-hand axe), over the Ironwood Blade he wears. */
    val sheet: GearSheetContent
        get() = GearSheetContent(
            tooltip = ItemTooltipModel(
                look = deathbringer,
                headerMeta = "",
                known = true,
                binds = "Binds when picked up",
                slotName = "One-Hand",
                typeName = "Axe",
                weaponDamage = "114 - 213 Damage",
                weaponSpeed = "Speed 2.90",
                weaponDps = "(56.4 damage per second)",
                effect = "Chance on hit: Sends a shadowy bolt at the enemy causing 125 Shadow damage.",
                requiresLevel = "Requires Level 60",
            ),
            gain = "+12.6 DPS",
            sources = listOf(
                GearSheetLine(
                    composer.sourceLine(ItemSource("raid", "Onyxia's Lair", 8.0, boss = "Onyxia", levels = 60..60)), null, null,
                ),
            ),
            replaces = GearSheetLine(
                "Replaces Ironwood Blade", "Equipped",
                FlowLead.Item(ItemLook(7, "Ironwood Blade", 4, "inv_sword_07"), BUILD),
            ),
            runnerUp = GearSheetLine(
                "Vis'kag the Bloodletter", "Runner-up · −1.8 DPS",
                FlowLead.Item(ItemLook(17075, "Vis'kag the Bloodletter", 4, "inv_sword_18"), BUILD),
                titleColor = ItemQuality(4).text,
            ),
            guild = listOf(
                GearSheetLine("Grunvald", "Protection Warrior · needs it, the pick on their list", FlowLead.Crest("warrior"),
                    classColor("warrior")),
                GearSheetLine("Torvak", "Enhancement Shaman · has it", FlowLead.Crest("shaman"),
                    classColor("shaman")),
            ),
            guildNote = null,
        )

    private fun pick(id: Int, name: String, worn: Boolean = false) =
        BisPick(BisItemRef(id, name, null), "", "", isUnmodelled = false, isWorn = worn, runnersUp = emptyList<BisRunnerUp>(), source = "")

    val list = BisList(
        setDpsLine = null,
        rows = listOf(
            BisRow("head", "Head", pick(2, "Lionheart Helm"), null),
            BisRow("ranged", "Ranged", pick(1, "Blackcrow"), null),
            BisRow("finger_1", "Finger", pick(3, "Don Julio's Band", worn = true), null),
        ),
        unit = BisUnit.Dps,
        specLabel = "Fury Warrior",
        rangeLabel = "60 raid",
        setFigure = "1004.2",
    )

    private val smith = ItemSource("crafted", "Blacksmithing", null, profession = "blacksmithing")
    val picks = listOf(
        RunPick("head", 2, "Lionheart Helm", 14.5, smith),
        RunPick("off_hand", 5, "Greenhammer", 9.7, smith),
        RunPick("finger_1", 3, "Don Julio's Band", 6.8, ItemSource("rep", "Stormpike Guard", null, standing = "Revered")),
        RunPick("main_hand", 4, "Ironwood Blade", 5.9, smith),
        RunPick("ranged", 1, "Blackcrow", 1.9, ItemSource("dungeon", "Blackrock Spire", 8.0, boss = "Shadow Hunter Vosh'gajin")),
        RunPick("legs", 6, "Sentinel's Plate Legguards", 4.7, ItemSource("vendor", "Illiyana Moonblaze", null)),
    )

    val needs = ItemNeeds(
        17068,
        listOf(ItemNeed("us/pvp/grunvald", "Grunvald", "Warrior", "Protection", 60, 60, "main_hand", true)),
        listOf(ItemHolder("us/pvp/torvak", "Torvak", "Shaman", "Enhancement")),
        0,
    )
}
