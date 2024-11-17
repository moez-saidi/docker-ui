package models

import (
	"errors"
	"strings"
)

type ImageInfo struct {
	Name       string
	Tag        string
	Repository string
}

func (info *ImageInfo) GetFullRepoURL() string {
	err := info.Validate()

	if err != nil {
		return ""
	}

	return strings.TrimSuffix(info.Repository, "/") + "/" + info.Name + ":" + info.Tag
}

func (info *ImageInfo) Validate() error {
	if err := info.validateName(); err != nil {
		return err
	}
	if err := info.validateTag(); err != nil {
		return err
	}
	if err := info.validateRepository(); err != nil {
		return err
	}
	return nil
}

func (info *ImageInfo) validateName() error {
	if strings.TrimSpace(info.Name) == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}

func (info *ImageInfo) validateTag() error {
	if strings.TrimSpace(info.Tag) == "" {
		return errors.New("tag cannot be empty")
	}
	return nil
}

func (info *ImageInfo) validateRepository() error {
	if strings.TrimSpace(info.Repository) == "" {
		return errors.New("repository cannot be empty")
	}
	return nil
}
