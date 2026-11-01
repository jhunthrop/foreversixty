import DesignSystem
import DesignTokens
import Models
import SwiftUI

/// F2.7 Your gear: all 17 slots on the Paperdoll (lane F11), a gold dot on every slot the band's
/// list improves. A tap on a socket opens the item's tooltip.
struct GearView: View {
    let rows: [GearRow]
    let intro: String
    let build: String
    /// Opens an item's tooltip; the second value is the item worn in that slot.
    var open: (HomeItemLook, Int?) -> Void = { _, _ in }

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 16) {
                Text(Copy.gearTitle.uppercased()).font(ScaledType.title).foregroundStyle(FSColor.strong)
                    .accessibilityAddTraits(.isHeader)
                Text(intro).font(ScaledType.body).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
                VStack(alignment: .leading, spacing: 10) {
                    PaperdollView(frame: PaperdollModel.frame(PaperdollModel.home(rows)), build: build) { slot in
                        if let look = slot.look { open(look, look.id) }
                    }
                    Text(Copy.charGoldDot).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                        .fixedSize(horizontal: false, vertical: true)
                }
                .padding(.horizontal, 8)
                .padding(.vertical, 14)
                .quietCard()
            }
            .padding(16)
        }
        .background(FSColor.bg)
        .navigationTitle(Copy.gearTitle)
        .navigationBarTitleDisplayMode(.inline)
    }
}
