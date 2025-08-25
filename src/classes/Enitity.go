package classes

type Entity struct {
	// Data
	Name      string  // Name of entity, eg. Player
	PositionX float32 // Global x position of the entity
	PositionY float32 // Global y position of the entity

	// Stats
	Health            float32 // Health of the entity, eg. 10HP, value of -404 grants invincibility
	BaseAttack        float32 // Base damage dealt by an entity
	AttackMultiplier  float32 // Number to multiply damage by, eg. <1 weakens attacks and >1 strengthens attacks
	BaseDefense       float32 // Base damage reduction, eg. 0.5 means attacks of 2 damage deal 1.5 damage instead
	DefenseMultiplier float32 // Number to multiply defense by, eg. <1 decreases defence and >1 increases defence
	SpeedMultiplier   int     // Determine the number of tiles travelled in a single movement

	// Movement
	TileSplits     int  // Number of splits an entity sees on a tile, eg. 1 split = 4 subtiles per tile
	LegalDiagonals bool // Whether or not diagonal movement is legalised for the entity
	DashEnabled    bool // Whether or not the speed multiplier can be manipulated via dashing
}
