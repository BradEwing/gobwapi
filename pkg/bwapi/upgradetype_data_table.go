package bwapi

type fullUpgradeTypeData struct {
	mineralPriceBase   int32
	mineralPriceFactor int32
	gasPriceBase       int32
	gasPriceFactor     int32
	upgradeTimeBase    int32
	upgradeTimeFactor  int32
	maxRepeats         int32
	whatUpgrades       int32 // UnitType ID
	race               int32 // Race ID
	whatsRequired0     int32 // UnitType needed at level 1
	whatsRequired1     int32 // UnitType needed at level 2
	whatsRequired2     int32 // UnitType needed at level 3
}

// fullUpgradeTypeTable contains static data for all 63 BWAPI upgrade types, indexed by UpgradeType ID.
var fullUpgradeTypeTable = [63]fullUpgradeTypeData{
	// 0: Terran_Infantry_Armor
	{mineralPriceBase: 100, mineralPriceFactor: 75, gasPriceBase: 100, gasPriceFactor: 75, upgradeTimeBase: 4000, upgradeTimeFactor: 480, maxRepeats: 3, whatUpgrades: 122, race: 1, whatsRequired0: 228, whatsRequired1: 116, whatsRequired2: 116},
	// 1: Terran_Vehicle_Plating
	{mineralPriceBase: 100, mineralPriceFactor: 75, gasPriceBase: 100, gasPriceFactor: 75, upgradeTimeBase: 4000, upgradeTimeFactor: 480, maxRepeats: 3, whatUpgrades: 123, race: 1, whatsRequired0: 228, whatsRequired1: 116, whatsRequired2: 116},
	// 2: Terran_Ship_Plating
	{mineralPriceBase: 150, mineralPriceFactor: 75, gasPriceBase: 150, gasPriceFactor: 75, upgradeTimeBase: 4000, upgradeTimeFactor: 480, maxRepeats: 3, whatUpgrades: 123, race: 1, whatsRequired0: 228, whatsRequired1: 116, whatsRequired2: 116},
	// 3: Zerg_Carapace
	{mineralPriceBase: 150, mineralPriceFactor: 75, gasPriceBase: 150, gasPriceFactor: 75, upgradeTimeBase: 4000, upgradeTimeFactor: 480, maxRepeats: 3, whatUpgrades: 139, race: 0, whatsRequired0: 228, whatsRequired1: 132, whatsRequired2: 133},
	// 4: Zerg_Flyer_Carapace
	{mineralPriceBase: 150, mineralPriceFactor: 75, gasPriceBase: 150, gasPriceFactor: 75, upgradeTimeBase: 4000, upgradeTimeFactor: 480, maxRepeats: 3, whatUpgrades: 141, race: 0, whatsRequired0: 228, whatsRequired1: 132, whatsRequired2: 133},
	// 5: Protoss_Ground_Armor
	{mineralPriceBase: 100, mineralPriceFactor: 75, gasPriceBase: 100, gasPriceFactor: 75, upgradeTimeBase: 4000, upgradeTimeFactor: 480, maxRepeats: 3, whatUpgrades: 166, race: 2, whatsRequired0: 228, whatsRequired1: 165, whatsRequired2: 165},
	// 6: Protoss_Air_Armor
	{mineralPriceBase: 150, mineralPriceFactor: 75, gasPriceBase: 150, gasPriceFactor: 75, upgradeTimeBase: 4000, upgradeTimeFactor: 480, maxRepeats: 3, whatUpgrades: 164, race: 2, whatsRequired0: 228, whatsRequired1: 169, whatsRequired2: 169},
	// 7: Terran_Infantry_Weapons
	{mineralPriceBase: 100, mineralPriceFactor: 75, gasPriceBase: 100, gasPriceFactor: 75, upgradeTimeBase: 4000, upgradeTimeFactor: 480, maxRepeats: 3, whatUpgrades: 122, race: 1, whatsRequired0: 228, whatsRequired1: 116, whatsRequired2: 116},
	// 8: Terran_Vehicle_Weapons
	{mineralPriceBase: 100, mineralPriceFactor: 75, gasPriceBase: 100, gasPriceFactor: 75, upgradeTimeBase: 4000, upgradeTimeFactor: 480, maxRepeats: 3, whatUpgrades: 123, race: 1, whatsRequired0: 228, whatsRequired1: 116, whatsRequired2: 116},
	// 9: Terran_Ship_Weapons
	{mineralPriceBase: 100, mineralPriceFactor: 50, gasPriceBase: 100, gasPriceFactor: 50, upgradeTimeBase: 4000, upgradeTimeFactor: 480, maxRepeats: 3, whatUpgrades: 123, race: 1, whatsRequired0: 228, whatsRequired1: 116, whatsRequired2: 116},
	// 10: Zerg_Melee_Attacks
	{mineralPriceBase: 100, mineralPriceFactor: 50, gasPriceBase: 100, gasPriceFactor: 50, upgradeTimeBase: 4000, upgradeTimeFactor: 480, maxRepeats: 3, whatUpgrades: 139, race: 0, whatsRequired0: 228, whatsRequired1: 132, whatsRequired2: 133},
	// 11: Zerg_Missile_Attacks
	{mineralPriceBase: 100, mineralPriceFactor: 50, gasPriceBase: 100, gasPriceFactor: 50, upgradeTimeBase: 4000, upgradeTimeFactor: 480, maxRepeats: 3, whatUpgrades: 139, race: 0, whatsRequired0: 228, whatsRequired1: 132, whatsRequired2: 133},
	// 12: Zerg_Flyer_Attacks
	{mineralPriceBase: 100, mineralPriceFactor: 75, gasPriceBase: 100, gasPriceFactor: 75, upgradeTimeBase: 4000, upgradeTimeFactor: 480, maxRepeats: 3, whatUpgrades: 141, race: 0, whatsRequired0: 228, whatsRequired1: 132, whatsRequired2: 133},
	// 13: Protoss_Ground_Weapons
	{mineralPriceBase: 100, mineralPriceFactor: 50, gasPriceBase: 100, gasPriceFactor: 50, upgradeTimeBase: 4000, upgradeTimeFactor: 480, maxRepeats: 3, whatUpgrades: 166, race: 2, whatsRequired0: 228, whatsRequired1: 165, whatsRequired2: 165},
	// 14: Protoss_Air_Weapons
	{mineralPriceBase: 100, mineralPriceFactor: 75, gasPriceBase: 100, gasPriceFactor: 75, upgradeTimeBase: 4000, upgradeTimeFactor: 480, maxRepeats: 3, whatUpgrades: 164, race: 2, whatsRequired0: 228, whatsRequired1: 169, whatsRequired2: 169},
	// 15: Protoss_Plasma_Shields
	{mineralPriceBase: 200, mineralPriceFactor: 100, gasPriceBase: 200, gasPriceFactor: 100, upgradeTimeBase: 4000, upgradeTimeFactor: 480, maxRepeats: 3, whatUpgrades: 166, race: 2, whatsRequired0: 228, whatsRequired1: 164, whatsRequired2: 164},
	// 16: U_238_Shells
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 1500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 112, race: 1, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 17: Ion_Thrusters
	{mineralPriceBase: 100, mineralPriceFactor: 0, gasPriceBase: 100, gasPriceFactor: 0, upgradeTimeBase: 1500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 120, race: 1, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 18: gap_18
	{mineralPriceBase: 0, mineralPriceFactor: 0, gasPriceBase: 0, gasPriceFactor: 0, upgradeTimeBase: 0, upgradeTimeFactor: 0, maxRepeats: 0, whatUpgrades: 228, race: 7, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 19: Titan_Reactor
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 116, race: 1, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 20: Ocular_Implants
	{mineralPriceBase: 100, mineralPriceFactor: 0, gasPriceBase: 100, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 117, race: 1, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 21: Moebius_Reactor
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 117, race: 1, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 22: Apollo_Reactor
	{mineralPriceBase: 200, mineralPriceFactor: 0, gasPriceBase: 200, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 115, race: 1, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 23: Colossus_Reactor
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 118, race: 1, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 24: Ventral_Sacs
	{mineralPriceBase: 200, mineralPriceFactor: 0, gasPriceBase: 200, gasPriceFactor: 0, upgradeTimeBase: 2400, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 132, race: 0, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 25: Antennae
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 2000, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 132, race: 0, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 26: Pneumatized_Carapace
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 2000, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 132, race: 0, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 27: Metabolic_Boost
	{mineralPriceBase: 100, mineralPriceFactor: 0, gasPriceBase: 100, gasPriceFactor: 0, upgradeTimeBase: 1500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 142, race: 0, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 28: Adrenal_Glands
	{mineralPriceBase: 200, mineralPriceFactor: 0, gasPriceBase: 200, gasPriceFactor: 0, upgradeTimeBase: 1500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 142, race: 0, whatsRequired0: 133, whatsRequired1: 228, whatsRequired2: 228},
	// 29: Muscular_Augments
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 1500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 135, race: 0, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 30: Grooved_Spines
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 1500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 135, race: 0, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 31: Gamete_Meiosis
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 138, race: 0, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 32: Metasynaptic_Node
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 136, race: 0, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 33: Singularity_Charge
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 164, race: 2, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 34: Leg_Enhancements
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 2000, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 163, race: 2, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 35: Scarab_Damage
	{mineralPriceBase: 200, mineralPriceFactor: 0, gasPriceBase: 200, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 171, race: 2, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 36: Reaver_Capacity
	{mineralPriceBase: 200, mineralPriceFactor: 0, gasPriceBase: 200, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 171, race: 2, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 37: Gravitic_Drive
	{mineralPriceBase: 200, mineralPriceFactor: 0, gasPriceBase: 200, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 171, race: 2, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 38: Sensor_Array
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 2000, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 159, race: 2, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 39: Gravitic_Boosters
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 2000, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 159, race: 2, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 40: Khaydarin_Amulet
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 165, race: 2, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 41: Apial_Sensors
	{mineralPriceBase: 100, mineralPriceFactor: 0, gasPriceBase: 100, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 169, race: 2, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 42: Gravitic_Thrusters
	{mineralPriceBase: 200, mineralPriceFactor: 0, gasPriceBase: 200, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 169, race: 2, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 43: Carrier_Capacity
	{mineralPriceBase: 100, mineralPriceFactor: 0, gasPriceBase: 100, gasPriceFactor: 0, upgradeTimeBase: 1500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 169, race: 2, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 44: Khaydarin_Core
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 170, race: 2, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 45: gap_45
	{mineralPriceBase: 0, mineralPriceFactor: 0, gasPriceBase: 0, gasPriceFactor: 0, upgradeTimeBase: 0, upgradeTimeFactor: 0, maxRepeats: 0, whatUpgrades: 228, race: 7, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 46: gap_46
	{mineralPriceBase: 0, mineralPriceFactor: 0, gasPriceBase: 0, gasPriceFactor: 0, upgradeTimeBase: 0, upgradeTimeFactor: 0, maxRepeats: 0, whatUpgrades: 228, race: 7, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 47: Argus_Jewel
	{mineralPriceBase: 100, mineralPriceFactor: 0, gasPriceBase: 100, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 169, race: 2, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 48: gap_48
	{mineralPriceBase: 0, mineralPriceFactor: 0, gasPriceBase: 0, gasPriceFactor: 0, upgradeTimeBase: 0, upgradeTimeFactor: 0, maxRepeats: 0, whatUpgrades: 228, race: 7, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 49: Argus_Talisman
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 165, race: 2, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 50: gap_50
	{mineralPriceBase: 0, mineralPriceFactor: 0, gasPriceBase: 0, gasPriceFactor: 0, upgradeTimeBase: 0, upgradeTimeFactor: 0, maxRepeats: 0, whatUpgrades: 228, race: 7, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 51: Caduceus_Reactor
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 2500, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 112, race: 1, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 52: Chitinous_Plating
	{mineralPriceBase: 150, mineralPriceFactor: 0, gasPriceBase: 150, gasPriceFactor: 0, upgradeTimeBase: 2000, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 140, race: 0, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 53: Anabolic_Synthesis
	{mineralPriceBase: 200, mineralPriceFactor: 0, gasPriceBase: 200, gasPriceFactor: 0, upgradeTimeBase: 2000, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 140, race: 0, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 54: Charon_Boosters
	{mineralPriceBase: 100, mineralPriceFactor: 0, gasPriceBase: 100, gasPriceFactor: 0, upgradeTimeBase: 2000, upgradeTimeFactor: 0, maxRepeats: 1, whatUpgrades: 120, race: 1, whatsRequired0: 123, whatsRequired1: 228, whatsRequired2: 228},
	// 55: gap_55
	{mineralPriceBase: 0, mineralPriceFactor: 0, gasPriceBase: 0, gasPriceFactor: 0, upgradeTimeBase: 0, upgradeTimeFactor: 0, maxRepeats: 0, whatUpgrades: 228, race: 7, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 56: gap_56
	{mineralPriceBase: 0, mineralPriceFactor: 0, gasPriceBase: 0, gasPriceFactor: 0, upgradeTimeBase: 0, upgradeTimeFactor: 0, maxRepeats: 0, whatUpgrades: 228, race: 7, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 57: gap_57
	{mineralPriceBase: 0, mineralPriceFactor: 0, gasPriceBase: 0, gasPriceFactor: 0, upgradeTimeBase: 0, upgradeTimeFactor: 0, maxRepeats: 0, whatUpgrades: 228, race: 7, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 58: gap_58
	{mineralPriceBase: 0, mineralPriceFactor: 0, gasPriceBase: 0, gasPriceFactor: 0, upgradeTimeBase: 0, upgradeTimeFactor: 0, maxRepeats: 0, whatUpgrades: 228, race: 7, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 59: gap_59
	{mineralPriceBase: 0, mineralPriceFactor: 0, gasPriceBase: 0, gasPriceFactor: 0, upgradeTimeBase: 0, upgradeTimeFactor: 0, maxRepeats: 0, whatUpgrades: 228, race: 7, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 60: Upgrade_60
	{mineralPriceBase: 0, mineralPriceFactor: 0, gasPriceBase: 0, gasPriceFactor: 0, upgradeTimeBase: 0, upgradeTimeFactor: 0, maxRepeats: 0, whatUpgrades: 228, race: 7, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 61: None
	{mineralPriceBase: 0, mineralPriceFactor: 0, gasPriceBase: 0, gasPriceFactor: 0, upgradeTimeBase: 0, upgradeTimeFactor: 0, maxRepeats: 0, whatUpgrades: 228, race: 7, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
	// 62: Unknown
	{mineralPriceBase: 0, mineralPriceFactor: 0, gasPriceBase: 0, gasPriceFactor: 0, upgradeTimeBase: 0, upgradeTimeFactor: 0, maxRepeats: 0, whatUpgrades: 228, race: 8, whatsRequired0: 228, whatsRequired1: 228, whatsRequired2: 228},
}
