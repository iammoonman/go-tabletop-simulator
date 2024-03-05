package tabletopsimulator

type ExhaustiveTransform struct {
	ScaleX float32 `json:"scaleX"`
	ScaleY float32 `json:"scaleY"`
	ScaleZ float32 `json:"scaleZ"`
	RotX   float32 `json:"rotX"`
	RotY   float32 `json:"rotY"`
	RotZ   float32 `json:"rotZ"`
	PosX   float32 `json:"posX"`
	PosY   float32 `json:"posY"`
	PosZ   float32 `json:"posZ"`
}

type ExhaustiveCustomDeck struct {
	FaceURL      string `json:"FaceURL"`
	BackURL      string `json:"BackURL"`
	NumWidth     int    `json:"NumWidth"`
	NumHeight    int    `json:"NumHeight"`
	BackIsHidden bool   `json:"BackIsHidden"`
	UniqueBack   bool   `json:"UniqueBack"`
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

type ExhaustiveComponentTags struct {
	Labels []string `json:"Labels"`
}

type ExhaustiveSnapPoint struct {
	Position ExhaustiveVector `json:"Position"`
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
	SaveName       string                        `json:"SaveName"`
	EpochTime      int                           `json:"EpochTime"`
	GameComplexity string                        `json:"GameComplexity"`
	Tags           []string                      `json:"Tags"`
	GameMode       string                        `json:"GameMode"`
	Gravity        float32                       `json:"Gravity"`
	PlayArea       float32                       `json:"PlayArea"`
	ComponentTags  ExhaustiveComponentTags       `json:"ComponentTags"`
	Date           string                        `json:"Date"`
	Table          string                        `json:"Table"`
	TableURL       string                        `json:"TableURL"`
	Sky            string                        `json:"Sky"`
	SkyURL         string                        `json:"SkyURL"`
	Note           string                        `json:"Note"`
	Rules          string                        `json:"Rules"`
	XmlUI          string                        `json:"XmlUI"`
	CustomUIAssets string                        `json:"CustomUIAssets"`
	LuaScript      string                        `json:"LuaScript"`
	LuaScriptState string                        `json:"LuaScriptState"`
	Grid           ExhaustiveGrid                `json:"Grid"`
	LightingState  ExhaustiveLighting            `json:"LightingState"`
	HandsState     ExhaustiveHands               `json:"HandsState"`
	TurnsState     ExhaustiveTurns               `json:"TurnsState"`
	VectorLines    []ExhaustiveVectorLine        `json:"VectorLines"`
	ObjectStates   []ExhaustiveObjectState       `json:"ObjectStates"`
	SnapPoints     []ExhaustiveSnapPoint         `json:"SnapPoints"`
	DecalPallet    []ExhaustiveCustomDecal       `json:"DecalPallet"`
	Decals         []ExhaustiveDecalState        `json:"Decals"`
	TabStates      map[string]ExhaustiveTabState `json:"TabStates"`
	CameraStates   []ExhaustiveCamera            `json:"CameraStates"`
	VersionNumber  string                        `json:"VersionNumber"`
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
	// Limits            string
	UseMotor bool `json:"UseMotor"`
	// Motor             string
	// Spring            string
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
	Nickname            string                              `json:"Nickname"`
	Description         string                              `json:"Description"`
	ColorDiffuse        ExhaustiveColor                     `json:"ColorDiffuse"`
	Locked              bool                                `json:"Locked"`
	Grid                bool                                `json:"Grid"`
	Snap                bool                                `json:"Snap"`
	Autoraise           bool                                `json:"Autoraise"`
	Sticky              bool                                `json:"Sticky"`
	Tooltip             bool                                `json:"Tooltip"`
	GridProjection      bool                                `json:"GridProjection"`
	HideWhenFaceDown    bool                                `json:"HideWhenFaceDown"`
	Hands               bool                                `json:"Hands"`
	AltSound            bool                                `json:"AltSound"`
	MaterialIndex       int                                 `json:"MaterialIndex"`
	MeshIndex           int                                 `json:"MeshIndex"`
	Layer               int                                 `json:"Layer"`
	Number              int                                 `json:"Number"`
	CardID              int                                 `json:"CardID"`
	SidewaysCard        bool                                `json:"SidewaysCard"`
	RPGmode             bool                                `json:"RPGmode"`
	RPGdead             bool                                `json:"RPGdead"`
	FogColor            string                              `json:"FogColor"`
	FogHidePointers     bool                                `json:"FogHidePointers"`
	FogReverseHiding    bool                                `json:"FogReverseHiding"`
	FogSeethrough       bool                                `json:"FogSeethrough"`
	DeckIDs             []int                               `json:"DeckIDs"`
	CustomDeck          map[string]ExhaustiveCustomDeck     `json:"CustomDeck"`
	CustomMesh          ExhaustiveCustomMeshState           `json:"CustomMesh"`
	CustomImage         ExhaustiveCustomImageState          `json:"CustomImage"`
	CustomAssetbundle   ExhaustiveCustomAssetbundleState    `json:"CustomAssetbundle"`
	FogOfWar            ExhaustiveFogOfWarSaveState         `json:"FogOfWar"`
	FogOfWarRevealer    ExhaustiveFogOfWarRevealerSaveState `json:"FogOfWarRevealer"`
	Clock               ExhaustiveClockSaveState            `json:"Clock"`
	Counter             ExhaustiveCounterState              `json:"Counter"`
	Tablet              ExhaustiveTabletState               `json:"Tablet"`
	Mp3Player           ExhaustiveMp3PlayerState            `json:"Mp3Player"`
	Calculator          ExhaustiveCalculatorState           `json:"Calculator"`
	Text                ExhaustiveTextState                 `json:"Text"`
	XmlUI               string                              `json:"XmlUI"`
	CustomUIAssets      []ExhaustiveCustomAssetState        `json:"CustomUIAssets"`
	LuaScript           string                              `json:"LuaScript"`
	LuaScriptState      string                              `json:"LuaScriptState"`
	ContainedObjects    []ExhaustiveObjectState             `json:"ContainedObjects"`
	PhysicsMaterial     ExhaustivePhysicsMaterialState      `json:"PhysicsMaterial"`
	Rigidbody           ExhaustiveRigidbodyState            `json:"Rigidbody"`
	JointFixed          ExhaustiveJointFixedState           `json:"JointFixed"`
	JointHinge          ExhaustiveJointHingeState           `json:"JointHinge"`
	JointSpring         ExhaustiveJointSpringState          `json:"JointSpring"`
	GUID                string                              `json:"GUID"`
	AttachedSnapPoints  []ExhaustiveSnapPoint               `json:"AttachedSnapPoints"`
	AttachedVectorLines []ExhaustiveVectorLine              `json:"AttachedVectorLines"`
	AttachedDecals      []ExhaustiveDecalState              `json:"AttachedDecals"`
	States              map[string]ExhaustiveObjectState    `json:"States"`
	RotationValues      []ExhaustiveRotationValueState      `json:"RotationValues"`
}
