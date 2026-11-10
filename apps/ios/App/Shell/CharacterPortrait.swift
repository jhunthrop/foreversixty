import DesignSystem
import Models
import SwiftUI
import UIKit

/// Where Battle.net portraits come from (`avatar_url` on `/v1/me` characters): the render CDN
/// through the URL cache and a decoded-image memory cache. Tests inject the repo's stand-ins.
protocol PortraitSource: Sendable {
    func cached(_ url: URL) -> UIImage?
    func load(_ url: URL) async -> UIImage?
}

/// NSCache and URLSession are thread-safe; nothing else here mutates.
final class RemotePortraitSource: PortraitSource, @unchecked Sendable {
    static let shared = RemotePortraitSource()

    private let images = NSCache<NSURL, UIImage>()
    private let session: URLSession

    init(session: URLSession = .shared) {
        self.session = session
    }

    func cached(_ url: URL) -> UIImage? { images.object(forKey: url as NSURL) }

    func load(_ url: URL) async -> UIImage? {
        if let image = cached(url) { return image }
        guard let (data, _) = try? await session.data(from: url), let image = UIImage(data: data) else { return nil }
        images.setObject(image, forKey: url as NSURL)
        return image
    }
}

extension EnvironmentValues {
    @Entry var portraitSource: any PortraitSource = RemotePortraitSource.shared
}

/// A character on the profile button and in the menu (flows spec 3, the owner's exception to the
/// one crest language): the Battle.net portrait in a 2 pt ring of the faction colour, the faction
/// emblem as a badge on its top right (40% of the diameter). Without a portrait the person glyph
/// fills the ring; in the drawer's Switch character rows (`crestClass`, round 5) the class crest does,
/// so a row tells its characters apart.
struct CharacterPortrait: View {
    let avatarURL: URL?
    let faction: String?
    let size: CGFloat
    /// The class whose crest stands in for a missing portrait; nil keeps the person glyph.
    var crestClass: String?

    @Environment(\.portraitSource) private var source
    @State private var image: UIImage?

    static let badgeShare: CGFloat = 0.4
    /// The badge sits out of the circle by these (`.pf-badge { top:-3px; right:-4px }`).
    static let badgeOffset = CGSize(width: 4, height: -3)
    /// The glyph's share of the diameter (`glyph("you", px * 0.6)`).
    static let glyphShare: CGFloat = 0.6

    var body: some View {
        ZStack {
            Circle().fill(FSColor.raised)
            if let shown = image ?? avatarURL.flatMap(source.cached) {
                Image(uiImage: shown).resizable().scaledToFill().parityItem(.image)
            } else if let crestClass, let crest = BrandImages.crest(crestClass) {
                Image(uiImage: crest).resizable().scaledToFill().parityItem(.image)
            } else {
                Image(systemName: "person")
                    .resizable()
                    .scaledToFit()
                    .frame(width: size * Self.glyphShare * 0.6, height: size * Self.glyphShare * 0.6)
                    .foregroundStyle(FSColor.muted)
            }
        }
        .frame(width: size, height: size)
        .clipShape(.circle)
        .overlay(Circle().stroke(ring, lineWidth: FSStroke.crestRing).padding(-FSStroke.crestRing / 2))
        .overlay(alignment: .topTrailing) {
            if let faction, let mark = BrandImages.faction(faction) {
                Image(uiImage: mark)
                    .resizable()
                    .scaledToFit()
                    .frame(width: size * Self.badgeShare, height: size * Self.badgeShare)
                    .shadow(color: FSColor.bg, radius: 1)
                    .parityItem(.image)
                    .offset(Self.badgeOffset)
            }
        }
        .accessibilityHidden(true)
        .task(id: avatarURL) {
            guard let avatarURL else { image = nil; return }
            image = await source.load(avatarURL)
        }
    }

    private var ring: Color {
        switch faction {
        case "alliance": FSColor.allianceBar
        case "horde": FSColor.hordeBar
        default: FSColor.muted
        }
    }
}

extension CharacterPortrait {
    init(_ character: AccountCharacter, size: CGFloat, crestFallback: Bool = false) {
        self.init(
            avatarURL: character.avatarUrl.flatMap(URL.init(string:)), faction: character.faction, size: size,
            crestClass: crestFallback ? character.className : nil
        )
    }
}

/// The profile button, top left on every tab root (flows spec 3): the selected character's
/// portrait at 32 inside the 44 pt target, or the person glyph alone when signed out or with no
/// character. It opens the profile menu.
struct ProfileButton: View {
    let character: AccountCharacter?
    let action: () -> Void

    static let portraitSize: CGFloat = 32
    static let anonGlyph: CGFloat = 20

    var body: some View {
        Button(action: action) {
            Group {
                if let character {
                    CharacterPortrait(character, size: Self.portraitSize)
                } else {
                    Image(systemName: "person")
                        .font(.system(size: Self.anonGlyph * 0.7))
                        .foregroundStyle(FSColor.muted)
                        .frame(width: Self.portraitSize, height: Self.portraitSize)
                        .background(FSColor.raised, in: .circle)
                        .overlay(Circle().stroke(FSColor.muted, lineWidth: FSStroke.crestRing))
                }
            }
            .frame(width: FSTarget.min, height: FSTarget.min)
            .contentShape(.rect)
        }
        .buttonStyle(.plain)
        .accessibilityLabel(Copy.menuOpen)
    }
}
