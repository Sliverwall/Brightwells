package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
	"fmt"
	"log"
)

// ------------------------------ Inventory SYSTEMS -------------------------------

// AddItem adds an item to the inventory at a specified slot.
// If the slot is already occupied, it adds to the count of the existing item.
func AddItem(inv components.InventoryComponent, slot int, item components.ItemComponent) error {
	if slot < 0 || slot >= inv.Slots {
		return fmt.Errorf("invalid slot number")
	}
	if inv.Items[slot].ID == 0 { // Slot is empty
		inv.Items[slot] = item
	} else if inv.Items[slot].ID == item.ID { // Slot already has the same item
		inv.Items[slot].Count += item.Count
	} else {
		return fmt.Errorf("slot already occupied by a different item")
	}
	return nil
}

// RemoveItem removes an item from the inventory from a specified slot.
func RemoveItem(inv components.InventoryComponent, slot int, count int) error {
	if slot < 0 || slot >= inv.Slots {
		return fmt.Errorf("invalid slot number")
	}
	if inv.Items[slot].Count < count {
		return fmt.Errorf("not enough items in slot")
	}
	inv.Items[slot].Count -= count
	if inv.Items[slot].Count == 0 {
		inv.Items[slot] = components.ItemComponent{} // Clear the slot if empty
	}
	return nil
}

// GetItem returns the item in a specified slot.
func GetItem(inv components.InventoryComponent, slot int) (components.ItemComponent, error) {
	if slot < 0 || slot >= inv.Slots {
		return components.ItemComponent{}, fmt.Errorf("invalid slot number")
	}
	return inv.Items[slot], nil
}

// ------------------------------ FOOD SYSTEMS -------------------------------
func (tcs *TriggerCollisionSystem) FoodCollide(entitySlice []*entities.Entity, collisions map[int][]int) {
	for entityID := range collisions {
		entity := entities.GetEntityByID(entitySlice, entityID)
		if entity.HasComponent(components.FoodComponentID) {
			log.Println("Ate apple")
			entity.KillEntity()

		}
	}
}
