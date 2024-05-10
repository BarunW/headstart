package main

import (
	"log/slog"

	"gopkg.in/ini.v1"
)

type Data struct {
	file *ini.File
}

func NewData(configPath string) (*Data, error) {
	fh := NewFileHandler()
	cfg, err := fh.OpenConfigFile(configPath)
	if err != nil {
		slog.Error("Unable to Open the config file", "Details", err.Error())
		return nil, err
	}
	return &Data{
		file: cfg,
	}, nil
}

func (d Data) GetValueFromSection(section Section, key string) string {
	sec := d.file.Section(string(section))
	return sec.Key(key).Value()
}

func (d Data) GetValFromCommandsSection(key string) string {
	return d.file.Section(string(SECTION_COMMANDS)).Key(key).Value()
}

func (d Data) GetTypeFromLinkTypeSection(key string) string {
	return d.file.Section(string(SECTION_TYPES)).Key(key).Value()
}
