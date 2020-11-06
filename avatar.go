package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// ErrNoAvatarURL err is thrown when Avatar instanse cannot return avatar URL
var ErrNoAvatarURL = errors.New("chat: Cannot get avatar URL")

// Avatar represents client's avatar img
type Avatar interface {
	GetAvatarURL(ChatUser) (string, error)
}

// TryAvatar tries to fetch avatar in several ways
type TryAvatar []Avatar

// GetAvatarURL implementation for tryAvatar
func (a TryAvatar) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}

// AuthAvatar is
type AuthAvatar struct{}

// UseAuthAvatar is
var UseAuthAvatar AuthAvatar

// GetAvatarURL is impled
func (AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

// GravatarAvatar is
type GravatarAvatar struct{}

// UseGravatar is
var UseGravatar GravatarAvatar

// GetAvatarURL is impled
func (GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return fmt.Sprintf("//www.gravatar.com/avatar/" + u.UniqueID()), nil
}

// FileSystemAvatar is
type FileSystemAvatar struct{}

// UseFileSystemAvatar is
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL is impled
func (FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := filepath.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}
