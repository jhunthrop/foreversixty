from pipeline.models import Item


def normalize_items(
    sparse_rows: list[dict[str, str]], item_rows: list[dict[str, str]]
) -> list[Item]:
    classes = {
        int(r["ID"]): (int(r["ClassID"]), int(r["SubclassID"])) for r in item_rows
    }
    out: list[Item] = []
    for s in sparse_rows:
        item_id = int(s["ID"])
        if item_id not in classes:
            continue
        class_id, subclass_id = classes[item_id]
        out.append(
            Item(
                id=item_id,
                name=s["Display_lang"],
                quality=int(s["OverallQualityID"]),
                item_level=int(s["ItemLevel"]),
                required_level=int(s["RequiredLevel"]),
                class_id=class_id,
                subclass_id=subclass_id,
                inventory_type=int(s["InventoryType"]),
            )
        )
    return out
