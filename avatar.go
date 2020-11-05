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
	GetAvatarURL(c *client) (string, error)
}

// AuthAvatar is
type AuthAvatar struct{}

// UseAuthAvatar is
var UseAuthAvatar AuthAvatar

// GetAvatarURL is impled
func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

// GravatarAvatar is
type GravatarAvatar struct{}

// UseGravatar is
var UseGravatar GravatarAvatar

// GetAvatarURL is impled
func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return fmt.Sprintf("//www.gravatar.com/avatar/" + useridStr), nil
		}
	}
	return "", ErrNoAvatarURL
}

// FileSystemAvatar is
type FileSystemAvatar struct{}

// UseFileSystemAvatar is
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL is impled
func (FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			if files, err := ioutil.ReadDir("avatars"); err == nil {
				for _, file := range files {
					if file.IsDir() {
						continue
					}
					if match, _ := filepath.Match(useridStr+"*", file.Name()); match {
						return "/avatars/" + file.Name(), nil
					}
				}
			}
		}
	}
	return "", ErrNoAvatarURL
}
