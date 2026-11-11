@testable import ForeverSixty
import DesignSystem
import SwiftUI
import UIKit

/// Lane M0's review sheet (tests only): every round-5 surface, button, primitive, the
/// paperdoll header and the tab bar on one screen, on the painted ground, with the mock's own
/// content (design/mocks/directions/mobile.py). AppTests/System's sheet snapshot is its capture.
struct DesignSystemSheet: View {
    let build: String
    var portrait: UIImage?

    var body: some View {
        VStack(spacing: 0) {
            bar
            VStack(alignment: .leading, spacing: FSSpace.space5) {
                typeSection
                buttons
                panels
                tags
                paperdoll
                tabBar
            }
            .padding(.horizontal, 14)
            .padding(.vertical, FSSpace.space4)
        }
        .screenGround()
    }

    /// The band behind the bars and the large title.
    private var bar: some View {
        VStack(alignment: .leading, spacing: FSSpace.space2) {
            HStack(spacing: FSSpace.space3) {
                PaperdollMedallion(portrait: portrait, size: 32)
                Text("Gear").font(Font(Theme.Typography.barTitle ?? .systemFont(ofSize: 17))).foregroundStyle(FSColor.strong)
                Spacer()
            }
            Text("Gear").type(FSType.title).foregroundStyle(Theme.Palette.title).shadow(color: .black.opacity(0.45), radius: 2, y: 1)
        }
        .padding(.horizontal, FSSpace.space4)
        .padding(.top, FSSpace.space3)
        .padding(.bottom, 14)
        .frame(maxWidth: .infinity, alignment: .leading)
        .surfaceBand()
    }

    private var typeSection: some View {
        VStack(alignment: .leading, spacing: 0) {
            PageHeader(title: "Onyxia's Lair") { SmallText("Wed 6 Jan · 7:30 pm · 38 signed", color: FSColor.text) }
            FSSection("Next upgrades", trailing: { Text("All gear").font(ScaledType.meta).foregroundStyle(Theme.Palette.accent) }) {
                FSRow(title: "Lionheart Helm", titleColor: ItemQuality(value: 4).text, sub: "Head · Blacksmithing", inset: 0) {
                    ItemIconView(icon: "inv_helmet_36", quality: 4, build: build, size: 40)
                } trail: {
                    Text("+14.5 DPS").font(ScaledType.bodyStrong).foregroundStyle(FSColor.gain)
                }
                FSRow(title: "Don Julio's Band", titleColor: ItemQuality(value: 4).text, sub: "Finger · Stormpike Guard", selected: true) {
                    ItemIconView(icon: "inv_jewelry_ring_35", quality: 4, build: build, size: 40)
                } trail: {
                    Text("+6.8 DPS").font(ScaledType.bodyStrong).foregroundStyle(FSColor.gain)
                }
            }
            .padding(FSSpace.space4)
            .panel(frame: .plain)
        }
    }

    private var buttons: some View {
        VStack(spacing: FSSpace.space3) {
            HStack(spacing: FSSpace.space2) {
                Button("I'm in") {}.buttonStyle(BrandButtonStyle(.primary, height: 44))
                Button("Late") {}.buttonStyle(BrandButtonStyle(.secondary, height: 44))
                Button("Can't come") {}.buttonStyle(BrandButtonStyle(.secondary, height: 44))
            }
            HStack(spacing: FSSpace.space2) {
                Button("Turn on") {}.buttonStyle(BrandButtonStyle(.accent, height: 44))
                Button("Not now") {}.buttonStyle(BrandButtonStyle(.quiet, height: 44))
                Button("Leave guild") {}.buttonStyle(BrandButtonStyle(.destructive, height: 44))
            }
            HStack(spacing: FSSpace.space2) {
                FSButton("Open in the sim", kind: .primary) {}
                FSButton("Change", kind: .secondary) {}
                FSButton("Wowhead", kind: .ghost, fullWidth: false) {}
            }
        }
    }

    private var panels: some View {
        VStack(spacing: FSSpace.space4) {
            VStack(alignment: .leading, spacing: FSSpace.space2) {
                Text("Onyxia's Lair").type(FSType.sectionReading).foregroundStyle(Theme.Palette.title)
                Text("Wed 6 Jan · 7:30 pm").font(ScaledType.meta).foregroundStyle(FSColor.text)
                HStack(spacing: 18) {
                    figure("2", "days")
                    figure("4", "hours")
                    figure("12", "min")
                }
            }
            .padding(18)
            .frame(maxWidth: .infinity, alignment: .leading)
            .panel(frame: .ornate, art: .cave)
            HStack(spacing: FSSpace.space3) {
                Text("Your guild's sims").font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                    .padding(FSSpace.space4).frame(maxWidth: .infinity, alignment: .leading).panel()
                Text("Updated 2 h ago").font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .padding(FSSpace.space4).frame(maxWidth: .infinity, alignment: .leading).quietCard()
            }
        }
    }

    private func figure(_ value: String, _ unit: String) -> some View {
        VStack(alignment: .leading, spacing: 2) {
            Text(value).type(FSType.figureBig).foregroundStyle(FSColor.strong)
            Text(unit).font(ScaledType.meta).foregroundStyle(FSColor.muted)
        }
    }

    private var tags: some View {
        VStack(alignment: .leading, spacing: FSSpace.space3) {
            HStack(spacing: FSSpace.space2) {
                FSTag("Officer", tone: .gold)
                FSTag("Killed", tone: .kill)
                FSTag("Wiped", tone: .wipe)
                FSTag("Bench", tone: .muted)
                TagView(text: "In", color: FSColor.kill, systemImage: "checkmark")
            }
            Label12("New for you", color: Theme.Palette.mutedOnGround)
            OfflineStrip(text: "Offline · showing Tuesday's raid")
        }
    }

    private var paperdoll: some View {
        PaperdollHeader(title: "Fury Warrior", figure: "1004.2", portrait: portrait, slots: SheetGear.slots, selected: "head", onSelect: { _ in }) { item in
            let row = SheetGear.row(item.id)
            ItemIconView(icon: row.icon, quality: row.quality, build: build, size: 38)
        }
    }

    private var tabBar: some View {
        TabView(selection: .constant(AppTab.gear)) {
            ForEach(AppTab.allCases) { tab in
                Tab(tab.title, image: tab.icon, value: tab) { Color.clear }
                    .badge(tab == .raid ? 1 : tab == .activity ? 3 : 0)
            }
        }
        .tint(Theme.Palette.tabSelected)
        .frame(height: 120)
        .clipShape(.rect(cornerRadius: Theme.Radius.panel))
    }
}

/// The mock's Fury list (design/mocks/ws_app_flows.py `DOLL`), Greenhammer left out until the
/// first raids, so the off hand shows Persuader worn.
private enum SheetGear {
    static let rows: [(slot: String, label: String, name: String, icon: String, quality: Int, gain: String?)] = [
        ("head", "Head", "Lionheart Helm", "inv_helmet_36", 4, "+14.5"),
        ("neck", "Neck", "Skibi's Pendant", "inv_jewelry_amulet_07", 3, nil),
        ("shoulder", "Shoulder", "Theramore Pauldrons", "inv_shoulder_06", 4, nil),
        ("back", "Back", "Howler's Furs", "inv_misc_cape_10", 3, "+2.4"),
        ("chest", "Chest", "Timbermaw Tunic", "inv_chest_fur", 4, "+5.3"),
        ("wrist", "Wrist", "Forest Stalker's Bracers", "inv_bracer_12", 4, "+0.1"),
        ("hands", "Hands", "Theramore Gauntlets", "inv_gauntlets_29", 4, "+1.0"),
        ("waist", "Waist", "Valiant's Waistguard", "inv_belt_33", 4, nil),
        ("legs", "Legs", "Sentinel's Plate Legguards", "inv_pants_plate_20", 4, "+4.7"),
        ("feet", "Feet", "Scalegut Treaders", "inv_boots_chain_11", 3, "+0.9"),
        ("finger1", "Finger", "Don Julio's Band", "inv_jewelry_ring_35", 4, "+6.8"),
        ("finger2", "Finger", "Cutthroat's Signet", "inv_jewelry_ring_01", 3, "+0.8"),
        ("trinket1", "Trinket", "Hand of Justice", "inv_jewelry_talisman_01", 3, nil),
        ("trinket2", "Trinket", "Molten Heart of the Mountain", "inv_misc_gem_ruby_01", 3, "+3.2"),
        ("main_hand", "Main Hand", "Ironwood Blade", "inv_sword_07", 4, "+5.9"),
        ("off_hand", "Off Hand", "Persuader", "inv_hammer_08", 4, nil),
        ("ranged", "Ranged", "Blackcrow", "inv_weapon_crossbow_04", 3, "+1.9"),
    ]

    /// The row behind an item's id (its index), for its icon and quality.
    static func row(_ id: Int) -> (icon: String, quality: Int) { (rows[id].icon, rows[id].quality) }

    static let slots: [PaperdollListSlot] = rows.enumerated().map { index, row in
        PaperdollListSlot(id: row.slot, label: row.label, item: PaperdollListSlot.Item(
            id: index, name: row.name, quality: ItemQuality(value: row.quality).text,
            detail: row.gain.map { "\($0) DPS" } ?? "Worn", isGain: row.gain != nil
        ))
    }
}
