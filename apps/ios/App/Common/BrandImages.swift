import DesignSystem
import SwiftUI
import UIKit

/// The bundled brand images from web/public/icons/hd (project.yml copies the folders in as
/// `crests` and `faction`), so a crest or a faction mark never waits on the network.
enum BrandImages {
    static func crest(_ classSlug: String) -> UIImage? { image("crests", classSlug) }

    static func faction(_ slug: String) -> UIImage? { image("faction", slug) }

    private static func image(_ folder: String, _ name: String) -> UIImage? {
        Bundle.main.url(forResource: name, withExtension: "webp", subdirectory: folder)
            .flatMap { UIImage(contentsOfFile: $0.path) }
    }
}

/// The circular ringed HD crest (spec §4.3, "one crest language"): a 2 pt ring at 60% of the
/// class colour and a soft glow.
struct ClassCrest: View {
    let classSlug: String
    var size: CGFloat = 44

    static let ringOpacity = 0.6
    static let glowOpacity = 0.35
    static let glowRadius: CGFloat = 7

    var body: some View {
        let color = GamePalette.classColor(classSlug)
        ZStack {
            Circle().fill(Theme.Palette.cardTop)
            if let image = BrandImages.crest(classSlug) {
                Image(uiImage: image).resizable().scaledToFill()
            }
        }
        .frame(width: size, height: size)
        .clipShape(Circle())
        .overlay(Circle().stroke(color.opacity(Self.ringOpacity), lineWidth: 2))
        .shadow(color: color.opacity(Self.glowOpacity), radius: Self.glowRadius)
        .accessibilityHidden(true)
    }
}

/// The Horde or Alliance mark.
struct FactionMark: View {
    let faction: String
    var size: CGFloat = 24

    var body: some View {
        Group {
            if let image = BrandImages.faction(faction) {
                Image(uiImage: image).resizable().scaledToFit()
            } else {
                Circle().fill(GamePalette.faction(faction))
            }
        }
        .frame(width: size, height: size)
        .accessibilityHidden(true)
    }
}
