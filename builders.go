package tabletopsimulator

import (
	"errors"
	"strconv"
)

func TTSCard(nickname string, description string, faceURL string, backURL string) *ExhaustiveObjectState {
	newcard := ExhaustiveObjectState{Name: "Card", Transform: DefaultTransform, Nickname: nickname, Description: description, CardID: 0, CustomDeck: make(map[string]ExhaustiveCustomDeck)}
	newcard.CustomDeck["100"] = ExhaustiveCustomDeck{FaceURL: faceURL, BackURL: backURL, NumWidth: 1, NumHeight: 1}
	return &newcard
}

func TTSDeck(nickname string, cards []ExhaustiveObjectState) *ExhaustiveObjectState {
	newdeck := ExhaustiveObjectState{Name: "Deck", Transform: DefaultTransform, Nickname: nickname, ContainedObjects: cards}
	for card := range cards {
		if len(newdeck.ContainedObjects[card].CustomDeck) == 0 {
			continue
		}
		newdeck.ContainedObjects[card].CardID = (card + 1) * 100
		newdeck.DeckIDs = append(newdeck.DeckIDs, (card+1)*100)
		for key, val := range newdeck.ContainedObjects[card].CustomDeck {
			newdeck.ContainedObjects[card].CustomDeck[strconv.Itoa(card+1)] = val
			newdeck.CustomDeck[strconv.Itoa(card+1)] = val
			delete(newdeck.ContainedObjects[card].CustomDeck, key)
		}
	}
	return &newdeck
}

func TTSBag(nickname string, objects []ExhaustiveObjectState) *ExhaustiveObjectState {
	newbag := ExhaustiveObjectState{Name: "Bag", Transform: DefaultTransform, Nickname: nickname, ContainedObjects: objects}
	return &newbag
}

func withName(name string) Option {
	return func(e *ExhaustiveObjectState) {
		e.Name = name
	}
}

func addDecal(name string, imageURL string, size float32, transform ExhaustiveTransform) Option {
	return func(e *ExhaustiveObjectState) {
		e.AttachedDecals = append(e.AttachedDecals, ExhaustiveDecalState{Transform: transform, CustomDecal: ExhaustiveCustomDecal{Name: name, ImageURL: imageURL, Size: size}})
	}
}

func withScript(script string) Option {
	return func(e *ExhaustiveObjectState) {
		e.LuaScript = script
	}
}

func addCustomDeck(faceURL string, backURL string, i int) Option {
	return func(e *ExhaustiveObjectState) {
		e.CustomDeck[strconv.Itoa(i)] = ExhaustiveCustomDeck{FaceURL: faceURL, BackURL: backURL, NumWidth: 1, NumHeight: 1}
	}
}

func addDeckID(i int) Option {
	return func(e *ExhaustiveObjectState) {
		e.DeckIDs = append(e.DeckIDs, i*100)
	}
}

func withNickname(nickname string) Option {
	return func(e *ExhaustiveObjectState) {
		e.Nickname = nickname
	}
}

func withDescription(description string) Option {
	return func(e *ExhaustiveObjectState) {
		e.Description = description
	}
}

func withCardID(cardID int) Option {
	return func(e *ExhaustiveObjectState) {
		e.CardID = cardID
	}
}

func addXMLUI(ui string) Option {
	return func(e *ExhaustiveObjectState) {
		e.XmlUI = ui
	}
}

func addContainedObject(objs ...Option) Option {
	return func(e *ExhaustiveObjectState) {
		newobj, err := NewTTSObject(objs...)
		if err == nil {
			e.ContainedObjects = append(e.ContainedObjects, *newobj)
		}
	}
}

func withAltLookAngle(x float32, y float32, z float32) Option {
	return func(e *ExhaustiveObjectState) {
		e.AltLookAngle = ExhaustiveVector{X: x, Y: y, Z: z}
	}
}

type Option func(*ExhaustiveObjectState)

func NewTTSObject(opts ...Option) (*ExhaustiveObjectState, error) {
	c := &ExhaustiveObjectState{
		Name:      "Bag",
		Transform: DefaultTransform,
	}
	for _, opt := range opts {
		opt(c)
	}
	if c.Name == "Deck" && len(c.ContainedObjects) < 2 {

		return nil, errors.New("Decks need more than one contained object.")
	}
	if c.Name == "Deck" && len(c.DeckIDs) < 2 {
		return nil, errors.New("Decks need more than one value in DeckIDs.")
	}
	return c, nil
}
