package components

// complete item list

const ItemComponentID = "ItemComponent"

// Item represents an item in the inventory.
type ItemComponent struct {
	ID    int // Unique identifier for the item
	Count int // Number of items in the slot
}

// InventoryComponentID is the identifier for the InventoryComponent
const InventoryComponentID = "InventoryComponent"

// InventoryComponent represents an inventory with a set number of slots.
type InventoryComponent struct {
	Slots int
	Items []ItemComponent // A slice to hold items for each slot
}

// FoodComponentID is the identifier for the FoodComponent
const FoodComponentID = "FoodComponent"

type FoodComponent struct {
	// flags entity as Food
}
