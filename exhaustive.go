package tabletopsimulator

type ExhaustiveTransform struct {
	ScaleX float32 `json:"scaleX"`
	ScaleY float32 `json:"scaleY"`
	ScaleZ float32 `json:"scaleZ"`
	RotX   float32 `json:"rotX,omitempty"`
	RotY   float32 `json:"rotY,omitempty"`
	RotZ   float32 `json:"rotZ,omitempty"`
	PosX   float32 `json:"posX,omitempty"`
	PosY   float32 `json:"posY,omitempty"`
	PosZ   float32 `json:"posZ,omitempty"`
}

var DefaultTransform ExhaustiveTransform = ExhaustiveTransform{ScaleX: 1.0, ScaleY: 1.0, ScaleZ: 1.0}

type ExhaustiveCustomDeck struct {
	FaceURL      string `json:"FaceURL"`
	BackURL      string `json:"BackURL"`
	NumWidth     int    `json:"NumWidth"`
	NumHeight    int    `json:"NumHeight"`
	BackIsHidden bool   `json:"BackIsHidden,omitempty"`
	UniqueBack   bool   `json:"UniqueBack,omitempty"`
}

type ExhaustiveVector struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}
type ExhaustiveColor struct {
	R float32 `json:"R"`
	G float32 `json:"G"`
	B float32 `json:"B"`
}

type ExhaustiveTabState struct {
	Title        string          `json:"Title"`
	Body         string          `json:"Body"`
	Color        ExhaustiveColor `json:"Color"`
	VisibleColor string          `json:"VisibleColor"`
	Id           int             `json:"Id"`
}

type ExhaustiveGrid struct {
	Type         int              `json:"Type"`
	Lines        bool             `json:"Lines"`
	Color        ExhaustiveColor  `json:"Color"`
	Opacity      float32          `json:"Opacity"`
	ThickLines   bool             `json:"ThickLines"`
	Snapping     bool             `json:"Snapping"`
	Offset       bool             `json:"Offset"`
	BothSnapping bool             `json:"BothSnapping"`
	XSize        float32          `json:"xSize"`
	YSize        float32          `json:"ySize"`
	PosOffset    ExhaustiveVector `json:"PosOffset"`
}

type ExhaustiveLighting struct {
	LightIntensity      float32         `json:"LightIntensity"`
	LightColor          ExhaustiveColor `json:"LightColor"`
	AmbientIntensity    float32         `json:"AmbientIntensity"`
	AmbientType         int             `json:"AmbientType"`
	AmbientSkyColor     ExhaustiveColor `json:"AmbientSkyColor"`
	AmbientEquatorColor ExhaustiveColor `json:"AmbientEquatorColor"`
	AmbientGroundColor  ExhaustiveColor `json:"AmbientGroundColor"`
	ReflectionIntensity float32         `json:"ReflectionIntensity"`
	LutIndex            int             `json:"LutIndex"`
	LutContribution     float32         `json:"LutContribution"`
}

type ExhaustiveTurns struct {
	Enable              bool     `json:"Enable"`
	Type                int      `json:"Type"`
	TurnOrder           []string `json:"TurnOrder"`
	Reverse             bool     `json:"Reverse"`
	SkipEmpty           bool     `json:"SkipEmpty"`
	DisableInteractions bool     `json:"DisableInteractions"`
	PassTurns           bool     `json:"PassTurns"`
	TurnColor           string   `json:"TurnColor"`
}

type ExhaustiveComponentTag struct {
	Displayed  string `json:"displayed"`
	Normalized string `json:"normalized"`
}

type ExhaustiveComponentTags struct {
	Labels []ExhaustiveComponentTag `json:"Labels"`
}

type ExhaustiveSnapPoint struct {
	Position ExhaustiveVector `json:"Position"`
	Rotation ExhaustiveVector `json:"Rotation,omitempty"`
}

type ExhaustiveCamera struct {
	Position ExhaustiveVector `json:"Position"`
	Rotation ExhaustiveVector `json:"Rotation"`
	Distance float32          `json:"Distance"`
	Zoomed   bool             `json:"Zoomed"`
}

type ExhaustiveHands struct {
	Enable        bool `json:"Enable"`
	DisableUnused bool `json:"DisableUnused"`
	Hiding        int  `json:"Hiding"`
}

type ExhaustiveVectorLine struct {
	Points3   []ExhaustiveVector `json:"Points3"`
	Color     ExhaustiveColor    `json:"Color"`
	Thickness float32            `json:"Thickness"`
	Rotation  ExhaustiveVector   `json:"Rotation"`
	Loop      bool               `json:"Loop"`
	Square    bool               `json:"Square"`
}

type ExhaustiveDecalState struct {
	Transform   ExhaustiveTransform   `json:"Transform"`
	CustomDecal ExhaustiveCustomDecal `json:"CustomDecal"`
}

type ExhaustiveCustomDecal struct {
	Name     string  `json:"Name"`
	ImageURL string  `json:"ImageURL"`
	Size     float32 `json:"Size"`
}

type ExhaustiveSave struct {
	CameraStates   []ExhaustiveCamera            `json:"CameraStates,omitempty"`
	ComponentTags  ExhaustiveComponentTags       `json:"ComponentTags,omitempty"`
	CustomUIAssets string                        `json:"CustomUIAssets,omitempty"`
	Date           string                        `json:"Date,omitempty"`
	DecalPallet    []ExhaustiveCustomDecal       `json:"DecalPallet,omitempty"`
	Decals         []ExhaustiveDecalState        `json:"Decals,omitempty"`
	EpochTime      int                           `json:"EpochTime,omitempty"`
	GameComplexity string                        `json:"GameComplexity,omitempty"`
	GameMode       string                        `json:"GameMode,omitempty"`
	GameType       string                        `json:"GameType,omitempty"`
	Gravity        float32                       `json:"Gravity,omitempty"`
	Grid           ExhaustiveGrid                `json:"Grid,omitempty"`
	Hands          ExhaustiveHands               `json:"Hands,omitempty"`
	LightingState  ExhaustiveLighting            `json:"LightingState,omitempty"`
	LuaScript      string                        `json:"LuaScript,omitempty"`
	LuaScriptState string                        `json:"LuaScriptState,omitempty"`
	Note           string                        `json:"Note,omitempty"`
	ObjectStates   []ExhaustiveObjectState       `json:"ObjectStates"`
	PlayArea       float32                       `json:"PlayArea,omitempty"`
	PlayerCounts   []int                         `json:"PlayerCounts,omitempty"`
	PlayingTime    []int                         `json:"PlayingTime,omitempty"`
	Rules          string                        `json:"Rules,omitempty"`
	SaveName       string                        `json:"SaveName,omitempty"`
	Sky            string                        `json:"Sky,omitempty"`
	SkyURL         string                        `json:"SkyURL,omitempty"`
	SnapPoints     []ExhaustiveSnapPoint         `json:"SnapPoints,omitempty"`
	Table          string                        `json:"Table,omitempty"`
	TableURL       string                        `json:"TableURL,omitempty"`
	TabStates      map[string]ExhaustiveTabState `json:"TabStates,omitempty"`
	Tags           []string                      `json:"Tags,omitempty"`
	Turns          ExhaustiveTurns               `json:"Turns,omitempty"`
	VectorLines    []ExhaustiveVectorLine        `json:"VectorLines,omitempty"`
	VersionNumber  string                        `json:"VersionNumber,omitempty"`
	XmlUI          string                        `json:"XmlUI,omitempty"`
}

type ExhaustiveRotationValueState struct {
	Value    struct{}         `json:"Value"`
	Rotation ExhaustiveVector `json:"Rotation"`
}

type ExhaustiveCustomImageState struct {
	ImageURL           string                            `json:"ImageURL"`
	ImageSecondaryURL  string                            `json:"ImageSecondaryURL"`
	WidthScale         float32                           `json:"WidthScale"`
	CustomDice         ExhaustiveCustomDiceState         `json:"CustomDice"`
	CustomToken        ExhaustiveCustomTokenState        `json:"CustomToken"`
	CustomJigsawPuzzle ExhaustiveCustomJigsawPuzzleState `json:"CustomJigsawPuzzle"`
	CustomTile         ExhaustiveCustomTileState         `json:"CustomTile"`
}

type ExhaustiveCustomAssetbundleState struct {
	AssetbundleURL          string `json:"AssetbundleURL"`
	AssetbundleSecondaryURL string `json:"AssetbundleSecondaryURL"`
	/* 0 = Plastic, 1 = Wood, 2 = Metal, 3 = Cardboard */
	MaterialIndex int `json:"MaterialIndex"`
	/* 0 = Generic, 1 = Figurine, 2 = Dice, 3 = Coin, 4 = Board, 5 = Chip, 6 = Bag, 7 = Infinite */
	TypeIndex          int `json:"TypeIndex"`
	LoopingEffectIndex int `json:"LoopingEffectIndex"`
}

type ExhaustiveCustomDiceState struct {
	Type ExhaustiveDiceType `json:"Type"`
}

type ExhaustiveDiceType struct{}

type ExhaustiveCustomTokenState struct {
	Thickness           float32 `json:"Thickness"`
	MergeDistancePixels float32 `json:"MergeDistancePixels"`
	Stackable           bool    `json:"Stackable"`
}

type ExhaustiveCustomTileState struct {
	/* 0 = Box, 1 = Hex, 2 = Circle, 3 = Rounded */
	Type      int     `json:"Type"`
	Thickness float32 `json:"Thickness"`
	Stackable bool    `json:"Stackable"`
	Stretch   bool    `json:"Stretch"`
}

type ExhaustiveCustomJigsawPuzzleState struct {
	NumPuzzlePieces int  `json:"NumPuzzlePieces"`
	ImageOnBoard    bool `json:"ImageOnBoard"`
}

type ExhaustiveCustomMeshState struct {
	MeshURL       string                      `json:"MeshURL"`
	DiffuseURL    string                      `json:"DiffuseURL"`
	NormalURL     string                      `json:"NormalURL"`
	ColliderURL   string                      `json:"ColliderURL"`
	Convex        bool                        `json:"Convex"`
	MaterialIndex int                         `json:"MaterialIndex"`
	TypeIndex     int                         `json:"TypeIndex"`
	CustomShader  ExhaustiveCustomShaderState `json:"CustomShader"`
	CastShadows   bool                        `json:"CastShadows"`
}

type ExhaustiveCustomShaderState struct {
	SpecularColor     ExhaustiveColor `json:"SpecularColor"`
	SpecularIntensity float32         `json:"SpecularIntensity"`
	SpecularSharpness float32         `json:"SpecularSharpness"`
	FresnelStrength   float32         `json:"FresnelStrength"`
}

type ExhaustiveFogOfWarSaveState struct {
	HideGmPointer     bool           `json:"HideGmPointer"`
	HideObjects       bool           `json:"HideObjects"`
	Height            float32        `json:"Height"`
	RevealedLocations map[string]int `json:"RevealedLocations"`
}

type ExhaustiveFogOfWarRevealerSaveState struct {
	Active bool    `json:"Active"`
	Range  float32 `json:"Range"`
	Color  string  `json:"Color"`
}
type ExhaustiveClockSaveState struct {
	ClockState    string `json:"ClockState"`
	SecondsPassed int    `json:"SecondsPassed"`
	Paused        bool   `json:"Paused"`
}
type ExhaustiveCounterState struct {
	Value int `json:"Value"`
}
type ExhaustiveTabletState struct {
	PageURL string `json:"PageURL"`
}
type ExhaustiveMp3PlayerState struct {
	SongTitle string  `json:"SongTitle"`
	Genre     string  `json:"Genre"`
	Volume    float32 `json:"Volume"`
	IsPlaying bool    `json:"IsPlaying"`
	LoopOne   bool    `json:"LoopOne"`
	MenuTitle string  `json:"MenuTitle"`
	Menu      string  `json:"Menu"`
}
type ExhaustiveCalculatorState struct {
	Value  string  `json:"Value"`
	Memory float32 `json:"Memory"`
}
type ExhaustiveTextState struct {
	Text       string          `json:"Text"`
	ColorState ExhaustiveColor `json:"ColorState"`
	FontSize   int             `json:"FontSize"`
}
type ExhaustivePhysicsMaterialState struct {
	StaticFriction  float32 `json:"StaticFriction"`
	DynamicFriction float32 `json:"DynamicFriction"`
	Bounciness      float32 `json:"Bounciness"`
	FrictionCombine int     `json:"FrictionCombine"`
	BounceCombine   int     `json:"BounceCombine"`
}
type ExhaustiveRigidbodyState struct {
	Mass        float32 `json:"Mass"`
	Drag        float32 `json:"Drag"`
	AngularDrag float32 `json:"AngularDrag"`
	UseGravity  bool    `json:"UseGravity"`
}
type ExhaustiveLimits struct {
	MinBounce         float32 `json:"minBounce"`
	MaxBounce         float32 `json:"maxBounce"`
	Min               float32 `json:"min"`
	Max               float32 `json:"max"`
	Bounciness        float32 `json:"bounciness"`
	BounceMinVelocity float32 `json:"bounceMinVelocity"`
	ContactDistance   float32 `json:"contactDistance"`
}
type ExhaustiveMotor struct {
	TargetVelocity float32 `json:"targetVelocity"`
	Force          float32 `json:"force"`
	FreeSpin       bool    `json:"freeSpin"`
}
type ExhaustiveSpring struct {
	Spring         float32 `json:"spring"`
	Damper         float32 `json:"damper"`
	TargetPosition float32 `json:"targetPosition"`
}
type ExhaustiveJointFixedState struct {
	ConnectedBodyGUID string           `json:"ConnectedBodyGUID"`
	EnableCollision   bool             `json:"EnableCollision"`
	Axis              ExhaustiveVector `json:"Axis"`
	Anchor            ExhaustiveVector `json:"Anchor"`
	ConnectedAnchor   ExhaustiveVector `json:"ConnectedAnchor"`
	BreakForce        float32          `json:"BreakForce"`
	BreakTorgue       float32          `json:"BreakTorgue"`
}
type ExhaustiveJointHingeState struct {
	ConnectedBodyGUID string           `json:"ConnectedBodyGUID"`
	EnableCollision   bool             `json:"EnableCollision"`
	Axis              ExhaustiveVector `json:"Axis"`
	Anchor            ExhaustiveVector `json:"Anchor"`
	ConnectedAnchor   ExhaustiveVector `json:"ConnectedAnchor"`
	BreakForce        float32          `json:"BreakForce"`
	BreakTorgue       float32          `json:"BreakTorgue"`
	UseLimits         bool             `json:"UseLimits"`
	UseMotor          bool             `json:"UseMotor"`
	Limits            ExhaustiveLimits `json:"Limits"`
	Motor             ExhaustiveMotor  `json:"Motor"`
	Spring            ExhaustiveSpring `json:"Spring"`
}
type ExhaustiveJointSpringState struct {
	ConnectedBodyGUID string           `json:"ConnectedBodyGUID"`
	EnableCollision   bool             `json:"EnableCollision"`
	Axis              ExhaustiveVector `json:"Axis"`
	Anchor            ExhaustiveVector `json:"Anchor"`
	ConnectedAnchor   ExhaustiveVector `json:"ConnectedAnchor"`
	BreakForce        float32          `json:"BreakForce"`
	BreakTorgue       float32          `json:"BreakTorgue"`
	Damper            float32          `json:"Damper"`
	MaxDistance       float32          `json:"MaxDistance"`
	MinDistance       float32          `json:"MinDistance"`
	Spring            float32          `json:"Spring"`
}
type ExhaustiveCustomAssetState struct {
	Name string `json:"Name"`
	URL  string `json:"URL"`
}

type ExhaustiveObjectState struct {
	Name                string                              `json:"Name"`
	Transform           ExhaustiveTransform                 `json:"Transform"`
	XmlUI               string                              `json:"XmlUI,omitempty"`
	AltLookAngle        ExhaustiveVector                    `json:"AltLookAngle,omitempty"`
	AltSound            bool                                `json:"AltSound,omitempty"`
	AttachedDecals      []ExhaustiveDecalState              `json:"AttachedDecals,omitempty"`
	AttachedSnapPoints  []ExhaustiveSnapPoint               `json:"AttachedSnapPoints,omitempty"`
	AttachedVectorLines []ExhaustiveVectorLine              `json:"AttachedVectorLines,omitempty"`
	Autoraise           bool                                `json:"Autoraise,omitempty"`
	Calculator          ExhaustiveCalculatorState           `json:"Calculator,omitempty"`
	CardID              int                                 `json:"CardID,omitempty"`
	Clock               ExhaustiveClockSaveState            `json:"Clock,omitempty"`
	ColorDiffuse        ExhaustiveColor                     `json:"ColorDiffuse,omitempty"`
	ContainedObjects    []ExhaustiveObjectState             `json:"ContainedObjects,omitempty"`
	Counter             ExhaustiveCounterState              `json:"Counter,omitempty"`
	CustomAssetbundle   ExhaustiveCustomAssetbundleState    `json:"CustomAssetbundle,omitempty"`
	CustomDeck          map[string]ExhaustiveCustomDeck     `json:"CustomDeck,omitempty"`
	CustomImage         ExhaustiveCustomImageState          `json:"CustomImage,omitempty"`
	CustomMesh          ExhaustiveCustomMeshState           `json:"CustomMesh,omitempty"`
	CustomUIAssets      []ExhaustiveCustomAssetState        `json:"CustomUIAssets,omitempty"`
	DeckIDs             []int                               `json:"DeckIDs,omitempty"`
	Description         string                              `json:"Description,omitempty"`
	FogColor            string                              `json:"FogColor,omitempty"`
	FogHidePointers     bool                                `json:"FogHidePointers,omitempty"`
	FogOfWar            ExhaustiveFogOfWarSaveState         `json:"FogOfWar,omitempty"`
	FogOfWarRevealer    ExhaustiveFogOfWarRevealerSaveState `json:"FogOfWarRevealer,omitempty"`
	FogReverseHiding    bool                                `json:"FogReverseHiding,omitempty"`
	FogSeethrough       bool                                `json:"FogSeethrough,omitempty"`
	Grid                bool                                `json:"Grid,omitempty"`
	GridProjection      bool                                `json:"GridProjection,omitempty"`
	GUID                string                              `json:"GUID,omitempty"`
	Hands               bool                                `json:"Hands,omitempty"`
	HideWhenFaceDown    bool                                `json:"HideWhenFaceDown,omitempty"`
	JointFixed          ExhaustiveJointFixedState           `json:"JointFixed,omitempty"`
	JointHinge          ExhaustiveJointHingeState           `json:"JointHinge,omitempty"`
	JointSpring         ExhaustiveJointSpringState          `json:"JointSpring,omitempty"`
	Layer               int                                 `json:"Layer,omitempty"`
	Locked              bool                                `json:"Locked,omitempty"`
	LuaScript           string                              `json:"LuaScript,omitempty"`
	LuaScriptState      string                              `json:"LuaScriptState,omitempty"`
	MaterialIndex       int                                 `json:"MaterialIndex,omitempty"`
	MeshIndex           int                                 `json:"MeshIndex,omitempty"`
	Mp3Player           ExhaustiveMp3PlayerState            `json:"Mp3Player,omitempty"`
	Nickname            string                              `json:"Nickname,omitempty"`
	Number              int                                 `json:"Number,omitempty"`
	PhysicsMaterial     ExhaustivePhysicsMaterialState      `json:"PhysicsMaterial,omitempty"`
	Rigidbody           ExhaustiveRigidbodyState            `json:"Rigidbody,omitempty"`
	RotationValues      []ExhaustiveRotationValueState      `json:"RotationValues,omitempty"`
	RPGdead             bool                                `json:"RPGdead,omitempty"`
	RPGmode             bool                                `json:"RPGmode,omitempty"`
	SidewaysCard        bool                                `json:"SidewaysCard,omitempty"`
	Snap                bool                                `json:"Snap,omitempty"`
	States              map[string]ExhaustiveObjectState    `json:"States,omitempty"`
	Sticky              bool                                `json:"Sticky,omitempty"`
	Tablet              ExhaustiveTabletState               `json:"Tablet,omitempty"`
	Text                ExhaustiveTextState                 `json:"Text,omitempty"`
	Tooltip             bool                                `json:"Tooltip,omitempty"`
}
