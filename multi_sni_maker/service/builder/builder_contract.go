package builder

import "sing-box-telegram/entity"

type Builder struct {
	ServerIP         string
	Setting          entity.Setting
	newReality       entity.RealityJson
	privateKey       string
	publicKey        string
	StringConfigZero string
	StringConfigAll  string
	SliceConfigAll   []string
}

func NewBuilder() *Builder {

	return &Builder{}
}
