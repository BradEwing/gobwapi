/*
 * gamedata.h — C struct definitions matching BWAPI 4.4.0 shared memory layout.
 *
 * These structs are ABI-compatible with the C++ structs in:
 *   bwapi/include/BWAPI/Client/GameData.h
 *   bwapi/include/BWAPI/Client/UnitData.h
 *   bwapi/include/BWAPI/Client/PlayerData.h
 *   bwapi/include/BWAPI/Client/BulletData.h
 *   bwapi/include/BWAPI/Client/RegionData.h
 *   bwapi/include/BWAPI/Client/ForceData.h
 *   bwapi/include/BWAPI/Client/GameTable.h
 *
 * Field offsets verified against JBWAPI ClientData.java byte offsets.
 * Total GameData size: 33,017,048 bytes.
 */

#ifndef GAMEDATA_H
#define GAMEDATA_H

#include <stdint.h>

/* Boolean type — must be 1 byte to match BWAPI's C++ bool layout. */
typedef unsigned char bw_bool;

/* Array size constants matching BWAPI 4.4.0 */
#define BW_MAX_FORCES           5
#define BW_MAX_PLAYERS          12
#define BW_MAX_UNITS            10000
#define BW_MAX_BULLETS          100
#define BW_MAX_REGIONS          5000
#define BW_MAX_UNIT_TYPES       234
#define BW_MAX_UPGRADE_TYPES    63
#define BW_MAX_TECH_TYPES       47
#define BW_MAX_EVENTS           10000
#define BW_MAX_EVENT_STRINGS    1000
#define BW_MAX_STRINGS          20000
#define BW_MAX_SHAPES           20000
#define BW_MAX_COMMANDS         20000
#define BW_MAX_UNIT_COMMANDS    20000
#define BW_MAX_NUKE_DOTS        200
#define BW_MAX_START_LOCATIONS  8
#define BW_MAX_GAME_INSTANCES   8
#define BW_TRAINING_QUEUE_SIZE  5
#define BW_MAX_NEIGHBORS        256
#define BW_MAP_MAX_TILE_SIZE    256
#define BW_MAP_MAX_WALK_SIZE    1024
#define BW_UNIT_ARRAY_SIZE      1700
#define BW_UNIT_SEARCH_SIZE     3400
#define BW_K_MAX                255
#define BW_M_MAX                3
#define BW_FLAG_MAX             2
#define BW_MAX_SELECTED_UNITS   12
#define BW_FORCE_NAME_SIZE      32
#define BW_PLAYER_NAME_SIZE     25
#define BW_MAP_FILE_NAME_SIZE   261
#define BW_MAP_PATH_NAME_SIZE   261
#define BW_MAP_NAME_SIZE        33
#define BW_MAP_HASH_SIZE        41
#define BW_STRING_SIZE          1024
#define BW_EVENT_STRING_SIZE    256

/* Position — pixel coordinates (8 bytes) */
typedef struct {
    int32_t x;
    int32_t y;
} Position;

/* UnitFinder — unit search index entry (8 bytes) */
typedef struct {
    int32_t unitIndex;
    int32_t searchValue;
} UnitFinder;

/* ForceData — alliance grouping (32 bytes) */
typedef struct {
    char name[BW_FORCE_NAME_SIZE];
} ForceData;

/* RegionData — pathfinding region (1068 bytes) */
typedef struct {
    int32_t  id;
    int32_t  islandID;
    int32_t  center_x;
    int32_t  center_y;
    int32_t  priority;
    int32_t  leftMost;
    int32_t  rightMost;
    int32_t  topMost;
    int32_t  bottomMost;
    int32_t  neighborCount;
    int32_t  neighbors[BW_MAX_NEIGHBORS];
    bw_bool  isAccessible;
    bw_bool  isHigherGround;
} RegionData;

/* BulletData — projectile state (80 bytes) */
typedef struct {
    int32_t  id;
    int32_t  player;
    int32_t  type;
    int32_t  source;
    int32_t  positionX;
    int32_t  positionY;
    double   angle;
    double   velocityX;
    double   velocityY;
    int32_t  target;
    int32_t  targetPositionX;
    int32_t  targetPositionY;
    int32_t  removeTimer;
    bw_bool  exists;
    bw_bool  isVisible[9];
} BulletData;

/* PlayerData — player state (5788 bytes) */
typedef struct {
    char     name[BW_PLAYER_NAME_SIZE];
    /* 3 bytes padding here (compiler-inserted) */
    int32_t  race;
    int32_t  type;
    int32_t  force;
    bw_bool  isAlly[BW_MAX_PLAYERS];
    bw_bool  isEnemy[BW_MAX_PLAYERS];
    bw_bool  isNeutral;
    /* 3 bytes padding here */
    int32_t  startLocationX;
    int32_t  startLocationY;
    bw_bool  isVictorious;
    bw_bool  isDefeated;
    bw_bool  leftGame;
    bw_bool  isParticipating;
    int32_t  minerals;
    int32_t  gas;
    int32_t  gatheredMinerals;
    int32_t  gatheredGas;
    int32_t  repairedMinerals;
    int32_t  repairedGas;
    int32_t  refundedMinerals;
    int32_t  refundedGas;
    int32_t  supplyTotal[3];
    int32_t  supplyUsed[3];
    int32_t  allUnitCount[BW_MAX_UNIT_TYPES];
    int32_t  visibleUnitCount[BW_MAX_UNIT_TYPES];
    int32_t  completedUnitCount[BW_MAX_UNIT_TYPES];
    int32_t  deadUnitCount[BW_MAX_UNIT_TYPES];
    int32_t  killedUnitCount[BW_MAX_UNIT_TYPES];
    int32_t  upgradeLevel[BW_MAX_UPGRADE_TYPES];
    bw_bool  hasResearched[BW_MAX_TECH_TYPES];
    bw_bool  isResearching[BW_MAX_TECH_TYPES];
    bw_bool  isUpgrading[BW_MAX_UPGRADE_TYPES];
    /* 3 bytes padding here */
    int32_t  color;
    int32_t  totalUnitScore;
    int32_t  totalKillScore;
    int32_t  totalBuildingScore;
    int32_t  totalRazingScore;
    int32_t  customScore;
    int32_t  maxUpgradeLevel[BW_MAX_UPGRADE_TYPES];
    bw_bool  isResearchAvailable[BW_MAX_TECH_TYPES];
    bw_bool  isUnitAvailable[BW_MAX_UNIT_TYPES];
} PlayerData;

/* UnitData — unit state (336 bytes) */
typedef struct {
    int32_t  clearanceLevel;
    int32_t  id;
    int32_t  player;
    int32_t  type;
    int32_t  positionX;
    int32_t  positionY;
    double   angle;
    double   velocityX;
    double   velocityY;
    int32_t  hitPoints;
    int32_t  lastHitPoints;
    int32_t  shields;
    int32_t  energy;
    int32_t  resources;
    int32_t  resourceGroup;
    int32_t  killCount;
    int32_t  acidSporeCount;
    int32_t  scarabCount;
    int32_t  interceptorCount;
    int32_t  spiderMineCount;
    int32_t  groundWeaponCooldown;
    int32_t  airWeaponCooldown;
    int32_t  spellCooldown;
    int32_t  defenseMatrixPoints;
    int32_t  defenseMatrixTimer;
    int32_t  ensnareTimer;
    int32_t  irradiateTimer;
    int32_t  lockdownTimer;
    int32_t  maelstromTimer;
    int32_t  orderTimer;
    int32_t  plagueTimer;
    int32_t  removeTimer;
    int32_t  stasisTimer;
    int32_t  stimTimer;
    int32_t  buildType;
    int32_t  trainingQueueCount;
    int32_t  trainingQueue[BW_TRAINING_QUEUE_SIZE];
    int32_t  tech;
    int32_t  upgrade;
    int32_t  remainingBuildTime;
    int32_t  remainingTrainTime;
    int32_t  remainingResearchTime;
    int32_t  remainingUpgradeTime;
    int32_t  buildUnit;
    int32_t  target;
    int32_t  targetPositionX;
    int32_t  targetPositionY;
    int32_t  order;
    int32_t  orderTarget;
    int32_t  orderTargetPositionX;
    int32_t  orderTargetPositionY;
    int32_t  secondaryOrder;
    int32_t  rallyPositionX;
    int32_t  rallyPositionY;
    int32_t  rallyUnit;
    int32_t  addon;
    int32_t  nydusExit;
    int32_t  powerUp;
    int32_t  transport;
    int32_t  carrier;
    int32_t  hatchery;
    bw_bool  exists;
    bw_bool  hasNuke;
    bw_bool  isAccelerating;
    bw_bool  isAttacking;
    bw_bool  isAttackFrame;
    bw_bool  isBeingGathered;
    bw_bool  isBlind;
    bw_bool  isBraking;
    bw_bool  isBurrowed;
    /* 3 bytes padding here */
    int32_t  carryResourceType;
    bw_bool  isCloaked;
    bw_bool  isCompleted;
    bw_bool  isConstructing;
    bw_bool  isDetected;
    bw_bool  isGathering;
    bw_bool  isHallucination;
    bw_bool  isIdle;
    bw_bool  isInterruptible;
    bw_bool  isInvincible;
    bw_bool  isLifted;
    bw_bool  isMorphing;
    bw_bool  isMoving;
    bw_bool  isParasited;
    bw_bool  isSelected;
    bw_bool  isStartingAttack;
    bw_bool  isStuck;
    bw_bool  isTraining;
    bw_bool  isUnderStorm;
    bw_bool  isUnderDarkSwarm;
    bw_bool  isUnderDWeb;
    bw_bool  isPowered;
    bw_bool  isVisible[9];
    /* 2 bytes padding here */
    int32_t  buttonset;
    int32_t  lastAttackerPlayer;
    bw_bool  recentlyAttacked;
    /* 3 bytes padding here */
    int32_t  replayID;
} UnitData;

/* Event — game event (12 bytes) */
typedef struct {
    int32_t type;
    int32_t v1;
    int32_t v2;
} Event;

/* Command — client command (12 bytes) */
typedef struct {
    int32_t type;
    int32_t value1;
    int32_t value2;
} Command;

/* Shape — debug drawing shape (40 bytes) */
typedef struct {
    int32_t type;
    int32_t ctype;
    int32_t x1;
    int32_t y1;
    int32_t x2;
    int32_t y2;
    int32_t extra1;
    int32_t extra2;
    int32_t color;
    bw_bool isSolid;
} Shape;

/* UnitCommand — unit command (24 bytes) */
typedef struct {
    int32_t type;
    int32_t unitIndex;
    int32_t targetIndex;
    int32_t x;
    int32_t y;
    int32_t extra;
} UnitCommand;

/* GameData — main shared memory structure (~33 MB) */
typedef struct {
    int32_t    client_version;
    int32_t    revision;
    bw_bool    isDebug;
    /* 3 bytes padding */
    int32_t    instanceID;
    int32_t    botAPM_noselects;
    int32_t    botAPM_selects;

    int32_t    forceCount;
    ForceData  forces[BW_MAX_FORCES];

    int32_t    playerCount;
    PlayerData players[BW_MAX_PLAYERS];

    int32_t    initialUnitCount;
    /* 4 bytes padding (UnitData requires 8-byte alignment for double) */
    UnitData   units[BW_MAX_UNITS];

    int32_t    unitArray[BW_UNIT_ARRAY_SIZE];

    BulletData bullets[BW_MAX_BULLETS];

    int32_t    nukeDotCount;
    Position   nukeDots[BW_MAX_NUKE_DOTS];

    int32_t    gameType;
    int32_t    latency;
    int32_t    latencyFrames;
    int32_t    latencyTime;
    int32_t    remainingLatencyFrames;
    int32_t    remainingLatencyTime;
    bw_bool    hasLatCom;
    bw_bool    hasGUI;
    /* 2 bytes padding */
    int32_t    replayFrameCount;
    uint32_t   randomSeed;
    int32_t    frameCount;
    int32_t    elapsedTime;
    int32_t    countdownTimer;
    int32_t    fps;
    double     averageFPS;

    int32_t    mouseX;
    int32_t    mouseY;
    bw_bool    mouseState[BW_M_MAX];
    bw_bool    keyState[BW_K_MAX];
    /* 2 bytes padding */
    int32_t    screenX;
    int32_t    screenY;

    bw_bool    flags[BW_FLAG_MAX];
    /* 2 bytes padding */
    int32_t    mapWidth;
    int32_t    mapHeight;
    char       mapFileName[BW_MAP_FILE_NAME_SIZE];
    char       mapPathName[BW_MAP_PATH_NAME_SIZE];
    char       mapName[BW_MAP_NAME_SIZE];
    char       mapHash[BW_MAP_HASH_SIZE];

    int32_t    getGroundHeight[BW_MAP_MAX_TILE_SIZE][BW_MAP_MAX_TILE_SIZE];
    bw_bool    isWalkable[BW_MAP_MAX_WALK_SIZE][BW_MAP_MAX_WALK_SIZE];
    bw_bool    isBuildable[BW_MAP_MAX_TILE_SIZE][BW_MAP_MAX_TILE_SIZE];
    bw_bool    isVisible[BW_MAP_MAX_TILE_SIZE][BW_MAP_MAX_TILE_SIZE];
    bw_bool    isExplored[BW_MAP_MAX_TILE_SIZE][BW_MAP_MAX_TILE_SIZE];
    bw_bool    hasCreep[BW_MAP_MAX_TILE_SIZE][BW_MAP_MAX_TILE_SIZE];
    bw_bool    isOccupied[BW_MAP_MAX_TILE_SIZE][BW_MAP_MAX_TILE_SIZE];

    uint16_t   mapTileRegionId[BW_MAP_MAX_TILE_SIZE][BW_MAP_MAX_TILE_SIZE];
    uint16_t   mapSplitTilesMiniTileMask[BW_MAX_REGIONS];
    uint16_t   mapSplitTilesRegion1[BW_MAX_REGIONS];
    uint16_t   mapSplitTilesRegion2[BW_MAX_REGIONS];

    int32_t    regionCount;
    RegionData regions[BW_MAX_REGIONS];

    int32_t    startLocationCount;
    Position   startLocations[BW_MAX_START_LOCATIONS];

    bw_bool    isInGame;
    bw_bool    isMultiplayer;
    bw_bool    isBattleNet;
    bw_bool    isPaused;
    bw_bool    isReplay;
    /* 3 bytes padding */
    int32_t    selectedUnitCount;
    int32_t    selectedUnits[BW_MAX_SELECTED_UNITS];

    int32_t    self;
    int32_t    enemy;
    int32_t    neutral;

    int32_t    eventCount;
    Event      events[BW_MAX_EVENTS];

    int32_t    eventStringCount;
    char       eventStrings[BW_MAX_EVENT_STRINGS][BW_EVENT_STRING_SIZE];

    int32_t    stringCount;
    char       strings[BW_MAX_STRINGS][BW_STRING_SIZE];

    int32_t    shapeCount;
    Shape      shapes[BW_MAX_SHAPES];

    int32_t    commandCount;
    Command    commands[BW_MAX_COMMANDS];

    int32_t    unitCommandCount;
    UnitCommand unitCommands[BW_MAX_UNIT_COMMANDS];

    int32_t    unitSearchSize;
    UnitFinder xUnitSearch[BW_UNIT_SEARCH_SIZE];
    UnitFinder yUnitSearch[BW_UNIT_SEARCH_SIZE];
} GameData;

/* GameInstance — entry in the game table (12 bytes) */
typedef struct {
    uint32_t serverProcessID;
    bw_bool  isConnected;
    /* 3 bytes padding */
    uint32_t lastKeepAliveTime;
} GameInstance;

/* GameTable — list of available game instances (96 bytes) */
typedef struct {
    GameInstance gameInstances[BW_MAX_GAME_INSTANCES];
} GameTable;

#endif /* GAMEDATA_H */
