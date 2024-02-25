package tabletopsimulator

type Transform struct {
	ScaleX float32 `json:"scaleX"`
	ScaleY float32 `json:"scaleY"`
	ScaleZ float32 `json:"scaleZ"`
}

type Save struct {
	ObjectStates []Deck `json:"ObjectStates"`
}

type Deck struct {
	Name             string               `json:"Name"`
	Transform        Transform            `json:"Transform"`
	DeckIDs          []int                `json:"DeckIDs"`
	ContainedObjects []Card               `json:"ContainedObjects"`
	CustomDeck       map[string]CardImage `json:"CustomDeck"`
}

type CardImage struct {
	FaceURL      string `json:"FaceURL"`
	BackURL      string `json:"BackURL"`
	NumWidth     uint8  `json:"NumWidth"`
	NumHeight    uint8  `json:"NumHeight"`
	BackIsHidden bool   `json:"BackIsHidden"`
}

type CardState struct {
	CustomDeck  map[string]CardImage `json:"CustomDeck"`
	Name        string               `json:"Name"`
	Transform   Transform            `json:"Transform"`
	Nickname    string               `json:"Nickname"`
	Description string               `json:"Description"`
	Memo        string               `json:"Memo"`
	CardID      uint32               `json:"CardID"`
	LuaScript   string               `json:"LuaScript"`
}

type Card struct {
	Name        string               `json:"Name"`
	Transform   Transform            `json:"Transform"`
	Nickname    string               `json:"Nickname"`
	Description string               `json:"Description"`
	Memo        string               `json:"Memo"`
	States      map[string]CardState `json:"States"`
	LuaScript   string               `json:"LuaScript"`
}

func NewDeckObject() Deck {
	var w = Deck{Name: "Deck", Transform: Transform{ScaleX: 1.0, ScaleY: 1.0, ScaleZ: 1.0}, DeckIDs: []int{}, ContainedObjects: []Card{}, CustomDeck: make(map[string]CardImage)}
	return w
}

func NewCardEntry(nickname string, description string, memo string) Card {
	var c = Card{Name: "Card", Transform: Transform{ScaleX: 1.0, ScaleY: 1.0, ScaleZ: 1.0}, Nickname: nickname, Description: description, Memo: memo, States: map[string]CardState{}, LuaScript: ""}
	return c
}

func NewImageEntry(f string, b string) CardImage {
	var i = CardImage{FaceURL: f, BackURL: b, NumWidth: 1, NumHeight: 1, BackIsHidden: true}
	return i
}

func NewStateEntry(nickname string, description string, memo string, luaScript string, image CardImage) CardState {
	var s = CardState{Name: "Card", Transform: Transform{ScaleX: 1.0, ScaleY: 1.0, ScaleZ: 1.0}, Nickname: nickname, CustomDeck: map[string]CardImage{}, Description: description, Memo: memo, CardID: 100, LuaScript: luaScript}
	s.CustomDeck["1"] = image
	return s
}

func NewTransform() Transform {
	return Transform{ScaleX: 1.0, ScaleY: 1.0, ScaleZ: 1.0}
}

type MaximalTransform struct {
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

type MaximalCustomDeck struct {
	FaceURL      string `json:"FaceURL"`
	BackURL      string `json:"BackURL"`
	NumWidth     int    `json:"NumWidth"`
	NumHeight    int    `json:"NumHeight"`
	BackIsHidden bool   `json:"BackIsHidden"`
	UniqueBack   bool   `json:"UniqueBack"`
}

type MaximalVector struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type MaximalColor struct {
	R float32 `json:"R"`
	G float32 `json:"G"`
	B float32 `json:"B"`
}

type MaximalTabState struct {
	Title        string       `json:"Title"`
	Body         string       `json:"Body"`
	Color        MaximalColor `json:"Color"`
	VisibleColor string       `json:"VisibleColor"`
	Id           int          `json:"Id"`
}

type MaximalGrid struct {
	Type         int           `json:"Type"`
	Lines        bool          `json:"Lines"`
	Color        MaximalColor  `json:"Color"`
	Opacity      float32       `json:"Opacity"`
	ThickLines   bool          `json:"ThickLines"`
	Snapping     bool          `json:"Snapping"`
	Offset       bool          `json:"Offset"`
	BothSnapping bool          `json:"BothSnapping"`
	XSize        float32       `json:"xSize"`
	YSize        float32       `json:"ySize"`
	PosOffset    MaximalVector `json:"PosOffset"`
}

type MaximalLighting struct {
	LightIntensity      float32      `json:"LightIntensity"`
	LightColor          MaximalColor `json:"LightColor"`
	AmbientIntensity    float32      `json:"AmbientIntensity"`
	AmbientType         int          `json:"AmbientType"`
	AmbientSkyColor     MaximalColor `json:"AmbientSkyColor"`
	AmbientEquatorColor MaximalColor `json:"AmbientEquatorColor"`
	AmbientGroundColor  MaximalColor `json:"AmbientGroundColor"`
	ReflectionIntensity float32      `json:"ReflectionIntensity"`
	LutIndex            int          `json:"LutIndex"`
	LutContribution     float32      `json:"LutContribution"`
}

type MaximalTurns struct {
	Enable              bool     `json:"Enable"`
	Type                int      `json:"Type"`
	TurnOrder           []string `json:"TurnOrder"`
	Reverse             bool     `json:"Reverse"`
	SkipEmpty           bool     `json:"SkipEmpty"`
	DisableInteractions bool     `json:"DisableInteractions"`
	PassTurns           bool     `json:"PassTurns"`
	TurnColor           string   `json:"TurnColor"`
}

type MaximalComponentTags struct {
	Labels []string `json:"Labels"`
}

type MaximalSnapPoint struct {
	Position MaximalVector `json:"Position"`
}

type MaximalCamera struct {
	Position MaximalVector `json:"Position"`
	Rotation MaximalVector `json:"Rotation"`
	Distance float32       `json:"Distance"`
	Zoomed   bool          `json:"Zoomed"`
}

type MaximalHands struct {
	Enable        bool `json:"Enable"`
	DisableUnused bool `json:"DisableUnused"`
	Hiding        int  `json:"Hiding"`
}

type MaximalVectorLine struct {
	Points3   []MaximalVector `json:"Points3"`
	Color     MaximalColor    `json:"Color"`
	Thickness float32         `json:"Thickness"`
	Rotation  MaximalVector   `json:"Rotation"`
	Loop      bool            `json:"Loop"`
	Square    bool            `json:"Square"`
}

type MaximalDecalState struct {
	Transform   MaximalTransform   `json:"Transform"`
	CustomDecal MaximalCustomDecal `json:"CustomDecal"`
}

type MaximalCustomDecal struct {
	Name     string  `json:"Name"`
	ImageURL string  `json:"ImageURL"`
	Size     float32 `json:"Size"`
}

type MaximalSave struct {
	SaveName       string                     `json:"SaveName"`
	EpochTime      int                        `json:"EpochTime"`
	GameComplexity string                     `json:"GameComplexity"`
	Tags           []string                   `json:"Tags"`
	GameMode       string                     `json:"GameMode"`
	Gravity        float32                    `json:"Gravity"`
	PlayArea       float32                    `json:"PlayArea"`
	ComponentTags  MaximalComponentTags       `json:"ComponentTags"`
	Date           string                     `json:"Date"`
	Table          string                     `json:"Table"`
	TableURL       string                     `json:"TableURL"`
	Sky            string                     `json:"Sky"`
	SkyURL         string                     `json:"SkyURL"`
	Note           string                     `json:"Note"`
	Rules          string                     `json:"Rules"`
	XmlUI          string                     `json:"XmlUI"`
	CustomUIAssets string                     `json:"CustomUIAssets"`
	LuaScript      string                     `json:"LuaScript"`
	LuaScriptState string                     `json:"LuaScriptState"`
	Grid           MaximalGrid                `json:"Grid"`
	LightingState  MaximalLighting            `json:"LightingState"`
	HandsState     MaximalHands               `json:"HandsState"`
	TurnsState     MaximalTurns               `json:"TurnsState"`
	VectorLines    []MaximalVectorLine        `json:"VectorLines"`
	ObjectStates   []MaximalObjectState       `json:"ObjectStates"`
	SnapPoints     []MaximalSnapPoint         `json:"SnapPoints"`
	DecalPallet    []MaximalCustomDecal       `json:"DecalPallet"`
	Decals         []MaximalDecalState        `json:"Decals"`
	TabStates      map[string]MaximalTabState `json:"TabStates"`
	CameraStates   []MaximalCamera            `json:"CameraStates"`
	VersionNumber  string                     `json:"VersionNumber"`
}

type MaximalRotationValueState struct {
	Value    struct{}      `json:"Value"`
	Rotation MaximalVector `json:"Rotation"`
}

type MaximalCustomImageState struct {
	ImageURL           string                         `json:"ImageURL"`
	ImageSecondaryURL  string                         `json:"ImageSecondaryURL"`
	WidthScale         float32                        `json:"WidthScale"`
	CustomDice         MaximalCustomDiceState         `json:"CustomDice"`
	CustomToken        MaximalCustomTokenState        `json:"CustomToken"`
	CustomJigsawPuzzle MaximalCustomJigsawPuzzleState `json:"CustomJigsawPuzzle"`
	CustomTile         MaximalCustomTileState         `json:"CustomTile"`
}

type MaximalCustomAssetbundleState struct {
	AssetbundleURL          string `json:"AssetbundleURL"`
	AssetbundleSecondaryURL string `json:"AssetbundleSecondaryURL"`
	/* 0 = Plastic, 1 = Wood, 2 = Metal, 3 = Cardboard */
	MaterialIndex int `json:"MaterialIndex"`
	/* 0 = Generic, 1 = Figurine, 2 = Dice, 3 = Coin, 4 = Board, 5 = Chip, 6 = Bag, 7 = Infinite */
	TypeIndex          int `json:"TypeIndex"`
	LoopingEffectIndex int `json:"LoopingEffectIndex"`
}

type MaximalCustomDiceState struct {
	Type MaximalDiceType `json:"Type"`
}

type MaximalDiceType struct{}

type MaximalCustomTokenState struct {
	Thickness           float32 `json:"Thickness"`
	MergeDistancePixels float32 `json:"MergeDistancePixels"`
	Stackable           bool    `json:"Stackable"`
}

type MaximalCustomTileState struct {
	/* 0 = Box, 1 = Hex, 2 = Circle, 3 = Rounded */
	Type      int     `json:"Type"`
	Thickness float32 `json:"Thickness"`
	Stackable bool    `json:"Stackable"`
	Stretch   bool    `json:"Stretch"`
}

type MaximalCustomJigsawPuzzleState struct {
	NumPuzzlePieces int  `json:"NumPuzzlePieces"`
	ImageOnBoard    bool `json:"ImageOnBoard"`
}

type MaximalCustomMeshState struct {
	MeshURL       string                   `json:"MeshURL"`
	DiffuseURL    string                   `json:"DiffuseURL"`
	NormalURL     string                   `json:"NormalURL"`
	ColliderURL   string                   `json:"ColliderURL"`
	Convex        bool                     `json:"Convex"`
	MaterialIndex int                      `json:"MaterialIndex"`
	TypeIndex     int                      `json:"TypeIndex"`
	CustomShader  MaximalCustomShaderState `json:"CustomShader"`
	CastShadows   bool                     `json:"CastShadows"`
}

type MaximalCustomShaderState struct {
	SpecularColor     MaximalColor `json:"SpecularColor"`
	SpecularIntensity float32      `json:"SpecularIntensity"`
	SpecularSharpness float32      `json:"SpecularSharpness"`
	FresnelStrength   float32      `json:"FresnelStrength"`
}

type MaximalFogOfWarSaveState struct {
	HideGmPointer     bool           `json:"HideGmPointer"`
	HideObjects       bool           `json:"HideObjects"`
	Height            float32        `json:"Height"`
	RevealedLocations map[string]int `json:"RevealedLocations"`
}

type MaximalFogOfWarRevealerSaveState struct {
	Active bool    `json:"Active"`
	Range  float32 `json:"Range"`
	Color  string  `json:"Color"`
}
type MaximalClockSaveState struct {
	ClockState    string `json:"ClockState"`
	SecondsPassed int    `json:"SecondsPassed"`
	Paused        bool   `json:"Paused"`
}
type MaximalCounterState struct {
	Value int `json:"Value"`
}
type MaximalTabletState struct {
	PageURL string `json:"PageURL"`
}
type MaximalMp3PlayerState struct {
	SongTitle string  `json:"SongTitle"`
	Genre     string  `json:"Genre"`
	Volume    float32 `json:"Volume"`
	IsPlaying bool    `json:"IsPlaying"`
	LoopOne   bool    `json:"LoopOne"`
	MenuTitle string  `json:"MenuTitle"`
	Menu      string  `json:"Menu"`
}
type MaximalCalculatorState struct {
	Value  string  `json:"Value"`
	Memory float32 `json:"Memory"`
}
type MaximalTextState struct {
	Text       string       `json:"Text"`
	ColorState MaximalColor `json:"ColorState"`
	FontSize   int          `json:"FontSize"`
}
type MaximalPhysicsMaterialState struct {
	StaticFriction  float32 `json:"StaticFriction"`
	DynamicFriction float32 `json:"DynamicFriction"`
	Bounciness      float32 `json:"Bounciness"`
	FrictionCombine int     `json:"FrictionCombine"`
	BounceCombine   int     `json:"BounceCombine"`
}
type MaximalRigidbodyState struct {
	Mass        float32 `json:"Mass"`
	Drag        float32 `json:"Drag"`
	AngularDrag float32 `json:"AngularDrag"`
	UseGravity  bool    `json:"UseGravity"`
}
type MaximalJointFixedState struct {
	ConnectedBodyGUID string        `json:"ConnectedBodyGUID"`
	EnableCollision   bool          `json:"EnableCollision"`
	Axis              MaximalVector `json:"Axis"`
	Anchor            MaximalVector `json:"Anchor"`
	ConnectedAnchor   MaximalVector `json:"ConnectedAnchor"`
	BreakForce        float32       `json:"BreakForce"`
	BreakTorgue       float32       `json:"BreakTorgue"`
}
type MaximalJointHingeState struct {
	ConnectedBodyGUID string        `json:"ConnectedBodyGUID"`
	EnableCollision   bool          `json:"EnableCollision"`
	Axis              MaximalVector `json:"Axis"`
	Anchor            MaximalVector `json:"Anchor"`
	ConnectedAnchor   MaximalVector `json:"ConnectedAnchor"`
	BreakForce        float32       `json:"BreakForce"`
	BreakTorgue       float32       `json:"BreakTorgue"`
	UseLimits         bool          `json:"UseLimits"`
	// Limits            string
	UseMotor bool `json:"UseMotor"`
	// Motor             string
	// Spring            string
}
type MaximalJointSpringState struct {
	ConnectedBodyGUID string        `json:"ConnectedBodyGUID"`
	EnableCollision   bool          `json:"EnableCollision"`
	Axis              MaximalVector `json:"Axis"`
	Anchor            MaximalVector `json:"Anchor"`
	ConnectedAnchor   MaximalVector `json:"ConnectedAnchor"`
	BreakForce        float32       `json:"BreakForce"`
	BreakTorgue       float32       `json:"BreakTorgue"`
	Damper            float32       `json:"Damper"`
	MaxDistance       float32       `json:"MaxDistance"`
	MinDistance       float32       `json:"MinDistance"`
	Spring            float32       `json:"Spring"`
}
type MaximalCustomAssetState struct {
	Name string `json:"Name"`
	URL  string `json:"URL"`
}

type MaximalObjectState struct {
	Name                string                           `json:"Name"`
	Transform           MaximalTransform                 `json:"Transform"`
	Nickname            string                           `json:"Nickname"`
	Description         string                           `json:"Description"`
	ColorDiffuse        MaximalColor                     `json:"ColorDiffuse"`
	Locked              bool                             `json:"Locked"`
	Grid                bool                             `json:"Grid"`
	Snap                bool                             `json:"Snap"`
	Autoraise           bool                             `json:"Autoraise"`
	Sticky              bool                             `json:"Sticky"`
	Tooltip             bool                             `json:"Tooltip"`
	GridProjection      bool                             `json:"GridProjection"`
	HideWhenFaceDown    bool                             `json:"HideWhenFaceDown"`
	Hands               bool                             `json:"Hands"`
	AltSound            bool                             `json:"AltSound"`
	MaterialIndex       int                              `json:"MaterialIndex"`
	MeshIndex           int                              `json:"MeshIndex"`
	Layer               int                              `json:"Layer"`
	Number              int                              `json:"Number"`
	CardID              int                              `json:"CardID"`
	SidewaysCard        bool                             `json:"SidewaysCard"`
	RPGmode             bool                             `json:"RPGmode"`
	RPGdead             bool                             `json:"RPGdead"`
	FogColor            string                           `json:"FogColor"`
	FogHidePointers     bool                             `json:"FogHidePointers"`
	FogReverseHiding    bool                             `json:"FogReverseHiding"`
	FogSeethrough       bool                             `json:"FogSeethrough"`
	DeckIDs             []int                            `json:"DeckIDs"`
	CustomDeck          map[string]MaximalCustomDeck     `json:"CustomDeck"`
	CustomMesh          MaximalCustomMeshState           `json:"CustomMesh"`
	CustomImage         MaximalCustomImageState          `json:"CustomImage"`
	CustomAssetbundle   MaximalCustomAssetbundleState    `json:"CustomAssetbundle"`
	FogOfWar            MaximalFogOfWarSaveState         `json:"FogOfWar"`
	FogOfWarRevealer    MaximalFogOfWarRevealerSaveState `json:"FogOfWarRevealer"`
	Clock               MaximalClockSaveState            `json:"Clock"`
	Counter             MaximalCounterState              `json:"Counter"`
	Tablet              MaximalTabletState               `json:"Tablet"`
	Mp3Player           MaximalMp3PlayerState            `json:"Mp3Player"`
	Calculator          MaximalCalculatorState           `json:"Calculator"`
	Text                MaximalTextState                 `json:"Text"`
	XmlUI               string                           `json:"XmlUI"`
	CustomUIAssets      []MaximalCustomAssetState        `json:"CustomUIAssets"`
	LuaScript           string                           `json:"LuaScript"`
	LuaScriptState      string                           `json:"LuaScriptState"`
	ContainedObjects    []MaximalObjectState             `json:"ContainedObjects"`
	PhysicsMaterial     MaximalPhysicsMaterialState      `json:"PhysicsMaterial"`
	Rigidbody           MaximalRigidbodyState            `json:"Rigidbody"`
	JointFixed          MaximalJointFixedState           `json:"JointFixed"`
	JointHinge          MaximalJointHingeState           `json:"JointHinge"`
	JointSpring         MaximalJointSpringState          `json:"JointSpring"`
	GUID                string                           `json:"GUID"`
	AttachedSnapPoints  []MaximalSnapPoint               `json:"AttachedSnapPoints"`
	AttachedVectorLines []MaximalVectorLine              `json:"AttachedVectorLines"`
	AttachedDecals      []MaximalDecalState              `json:"AttachedDecals"`
	States              map[string]MaximalObjectState    `json:"States"`
	RotationValues      []MaximalRotationValueState      `json:"RotationValues"`
}
