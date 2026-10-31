import DesignSystem
import SwiftUI
import UIKit

/// The bundled brand images from web/public (project.yml copies the `crests` and `faction`
/// folders and seal-400.png into the bundle), so none of them ever waits on the network.
enum BrandImages {
    static func crest(_ className: String?) -> UIImage? {
        className.flatMap { image("crests", $0.lowercased(), "webp") }
    }

    static func faction(_ slug: String) -> UIImage? { image("faction", slug, "webp") }

    static var seal: UIImage? { image(nil, "seal-400", "png") }

    private static func image(_ folder: String?, _ name: String, _ type: String) -> UIImage? {
        Bundle.main.url(forResource: name, withExtension: type, subdirectory: folder)
            .flatMap { UIImage(contentsOfFile: $0.path) }
    }
}

/// The circular ringed HD crest (spec §4.3, apps/shared/README.md "one crest language"): the
/// bundled crest clipped to a circle over the raised backing, a 2 pt ring in the class colour
/// at 60%, and its glow. Never a square class icon.
struct ClassCrest: View {
    let className: String?
    let size: CGFloat

    var body: some View {
        let color = ClassColor.of(className)
        ZStack {
            Circle().fill(FSColor.raised)
            if let image = BrandImages.crest(className) {
                Image(uiImage: image).resizable().scaledToFill()
            }
        }
        .frame(width: size, height: size)
        .clipShape(.circle)
        .overlay(Circle().stroke(color.opacity(0.6), lineWidth: 2))
        .shadow(color: color.opacity(0.35), radius: 7)
        .accessibilityHidden(true)
    }
}

/// The Horde or Alliance mark; a dot in the faction colour for an unknown slug.
struct FactionMark: View {
    let faction: String
    var size: CGFloat = 24

    var body: some View {
        Group {
            if let image = BrandImages.faction(faction) {
                Image(uiImage: image).resizable().scaledToFit()
            } else {
                Circle().fill(GameColors.faction(faction))
            }
        }
        .frame(width: size, height: size)
        .accessibilityHidden(true)
    }
}

/// The LX seal.
struct Seal: View {
    let size: CGFloat

    var body: some View {
        Group {
            if let image = BrandImages.seal {
                Image(uiImage: image).resizable().scaledToFit()
            } else {
                Circle().stroke(FSColor.gold, lineWidth: 2)
            }
        }
        .frame(width: size, height: size)
        .accessibilityHidden(true)
    }
}
