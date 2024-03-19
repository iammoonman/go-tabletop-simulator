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

type Decal struct {
	Transform   ExhaustiveTransform `json:"Transform"`
	CustomDecal struct {
		Name     string  `json:"Name"`
		ImageURL string  `json:"ImageURL"`
		Size     float32 `json:"Size"`
	} `json:"CustomDecal"`
}

type Card struct {
	Name        string               `json:"Name"`
	Transform   Transform            `json:"Transform"`
	Nickname    string               `json:"Nickname"`
	Description string               `json:"Description"`
	Memo        string               `json:"Memo"`
	States      map[string]CardState `json:"States"`
	LuaScript   string               `json:"LuaScript"`
	Decals      []Decal              `json:"Decals"`
}

type SingleCard struct {
	Name        string               `json:"Name"`
	Transform   Transform            `json:"Transform"`
	Nickname    string               `json:"Nickname"`
	Description string               `json:"Description"`
	Memo        string               `json:"Memo"`
	States      map[string]CardState `json:"States"`
	LuaScript   string               `json:"LuaScript"`
	CustomDeck  map[string]CardImage `json:"CustomDeck"`
	CardID      uint32               `json:"CardID"`
	Decals      []Decal              `json:"Decals"`
}

type DeckOrSingleCard interface {
	SingleCard() SingleCard
	Deck() Deck
}

type CardBagOptions struct {
	Order int8 `json:"Order"`
}

type CardBag struct {
	Name             string             `json:"Name"`
	Transform        Transform          `json:"Transform"`
	Nickname         string             `json:"Nickname"`
	Description      string             `json:"Description"`
	ContainedObjects []DeckOrSingleCard `json:"ContainedObjects"`
	BagOptions       CardBagOptions     `json:"Bag,omitempty"`
}

func NewSingleCardObject(nickname string, description string, memo string, customDeck CardImage, foil bool) SingleCard {
	var w = SingleCard{Name: "Card", Transform: Transform{ScaleX: 1.0, ScaleY: 1.0, ScaleZ: 1.0}, Nickname: nickname, Description: description, Memo: memo, States: make(map[string]CardState), CardID: 1, LuaScript: "", CustomDeck: make(map[string]CardImage), Decals: []Decal{{CustomDecal: struct {
		Name     string  "json:\"Name\""
		ImageURL string  "json:\"ImageURL\""
		Size     float32 "json:\"Size\""
	}{Name: "StarFoil", ImageURL: "https://i.imgur.com/QnxyMMK.png", Size: 1.0}, Transform: ExhaustiveTransform{
		PosX:   0.0,
		PosY:   0.25,
		PosZ:   0.0,
		RotX:   90.0,
		RotY:   180.0,
		RotZ:   0.0,
		ScaleX: 0.7006438 * 3.1,
		ScaleY: 0.9999966 * 3.1,
		ScaleZ: 15.3846169 * 3.1,
	}}}}
	w.CustomDeck["100"] = customDeck
	if !foil {
		w.Decals = w.Decals[:0]
	}
	return w
}

func NewDeckObject() Deck {
	var w = Deck{Name: "Deck", Transform: Transform{ScaleX: 1.0, ScaleY: 1.0, ScaleZ: 1.0}, DeckIDs: []int{}, ContainedObjects: []Card{}, CustomDeck: make(map[string]CardImage)}
	return w
}

func NewCardEntry(nickname string, description string, memo string, foil bool) Card {
	var c = Card{Name: "Card", Transform: Transform{ScaleX: 1.0, ScaleY: 1.0, ScaleZ: 1.0}, Nickname: nickname, Description: description, Memo: memo, States: map[string]CardState{}, LuaScript: "", Decals: []Decal{{CustomDecal: struct {
		Name     string  "json:\"Name\""
		ImageURL string  "json:\"ImageURL\""
		Size     float32 "json:\"Size\""
	}{Name: "StarFoil", ImageURL: "https://i.imgur.com/QnxyMMK.png", Size: 1.0}, Transform: ExhaustiveTransform{
		PosX:   0.0,
		PosY:   0.25,
		PosZ:   0.0,
		RotX:   90.0,
		RotY:   180.0,
		RotZ:   0.0,
		ScaleX: 0.7006438 * 3.1,
		ScaleY: 0.9999966 * 3.1,
		ScaleZ: 15.3846169 * 3.1,
	}}}}
	if !foil {
		c.Decals = c.Decals[:0]
	}
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

func NewCardBag(nickname string, description string) CardBag {
	return CardBag{Name: "Bag", Transform: NewTransform(), Nickname: nickname, Description: description, ContainedObjects: []DeckOrSingleCard{}, BagOptions: CardBagOptions{Order: 0}}
}
