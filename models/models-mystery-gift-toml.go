package models

type MysteryGiftClient struct {
	Name          string            `toml:"name"`
	GiftType      string            `toml:"giftType"`
	Pokemon       PokemonGiftClient `toml:"pokemonGift"`
	BeginningDate string            `toml:"beginningDate"`
	EndDate       string            `toml:"endDate"`
}
type PokemonGiftClient struct {
	OriginalTrainer  string      `toml:"originalTrainer"`
	ID               string      `toml:"id"`
	Species          string      `toml:"species"`
	Types            []string    `toml:"types"`
	Abilities        []Abilities `toml:"abilities"`
	Moves            []Moves     `toml:"moves"`
	Stats            Stats       `toml:"stats"`
	AssetData        AssetData   `toml:"assetData"`
	Nickname         string      `toml:"nickname"`
	LocationOfOrigin string      `toml:"locationOfOrigin"`
	Shiny            bool        `toml:"shiny"`
}

type MysteryGiftsServer struct {
	Mysterygifts []MysteryGiftServer `toml:"mysterygift"`
}

type PokemonGiftServer struct {
	OriginalTrainer  string   `toml:"originalTrainer"`
	ID               int      `toml:"id"`
	Species          string   `toml:"species"`
	Nickname         string   `toml:"nickname"`
	Shiny            bool     `toml:"shiny"`
	LocationOfOrigin string   `toml:"locationOfOrigin"`
	Types            []string `toml:"types"`
	Abilites         []string `toml:"abilites"`
	Stats            []int    `toml:"stats"`
	Moves            []string `toml:"moves"`
}

type MysteryGiftServer struct {
	Name          string            `toml:"name"`
	GiftType      string            `toml:"giftType"`
	BeginningDate string            `toml:"beginningDate"`
	EndDate       string            `toml:"endDate"`
	Pokemongift   PokemonGiftServer `toml:"pokemongift"`
}
