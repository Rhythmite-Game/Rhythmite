package classes

type Entity struct {
	// Data
	name      string  // Name of entity, eg. Player
	positionX float64 // Global x position of the entity
	positionY float64 // Global y position of the entity

	// Stats
	health            float64 // Health of the entity, eg. 10HP, value of -404 grants invincibility
	baseAttack        float64 // Base damage dealt by an entity
	attackMultiplier  float64 // Number to multiply damage by, eg. <1 weakens attacks and >1 strengthens attacks
	baseDefence       float64 // Base damage reduction, eg. 0.5 means attacks of 2 damage deal 1.5 damage instead
	defenceMultiplier float64 // Number to multiply defence by, eg. <1 decreases defence and >1 increases defence

	// Movement
	tileSplits     int  // Number of splits an entity sees on a tile, eg. 1 split = 4 subtiles per tile
	legalDiagonals bool // Whether or not diagonal movement is legalised for the entity
}
