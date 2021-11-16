package twitterscraper

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestGetProfile(t *testing.T) {
	loc := time.FixedZone("UTC", 0)
	joined := time.Date(2010, 01, 18, 8, 49, 30, 0, loc)
	sample := Profile{
		Avatar:    "https://pbs.twimg.com/profile_images/436075027193004032/XlDa2oaz_normal.jpeg",
		Banner:    "https://pbs.twimg.com/profile_banners/106037940/1541084318",
		Biography: "nothing",
		//	Birthday:   "March 21",
		IsPrivate:      false,
		IsVerified:     false,
		Joined:         &joined,
		Location:       "Ukraine",
		Name:           "Nomadic",
		PinnedTweetIDs: []string{},
		URL:            "https://twitter.com/nomadic_ua",
		UserID:         "106037940",
		Username:       "nomadic_ua",
		Website:        "https://nomadic.name",
	}

	profile, err := GetProfile("nomadic_ua")
	if err != nil {
		t.Error(err)
	}

	cmpOptions := cmp.Options{
		cmpopts.IgnoreFields(Profile{}, "FollowersCount"),
		cmpopts.IgnoreFields(Profile{}, "FollowingCount"),
		cmpopts.IgnoreFields(Profile{}, "FriendsCount"),
		cmpopts.IgnoreFields(Profile{}, "LikesCount"),
		cmpopts.IgnoreFields(Profile{}, "ListedCount"),
		cmpopts.IgnoreFields(Profile{}, "TweetsCount"),
	}
	if diff := cmp.Diff(sample, profile, cmpOptions...); diff != "" {
		t.Error("Resulting profile does not match the sample", diff)
	}

	if profile.FollowersCount == 0 {
		t.Error("Expected FollowersCount is greater than zero")
	}
	if profile.FollowingCount == 0 {
		t.Error("Expected FollowingCount is greater than zero")
	}
	if profile.LikesCount == 0 {
		t.Error("Expected LikesCount is greater than zero")
	}
	if profile.TweetsCount == 0 {
		t.Error("Expected TweetsCount is greater than zero")
	}
}

func TestGetProfilePrivate(t *testing.T) {
	loc := time.FixedZone("UTC", 0)
	joined := time.Date(2020, 1, 26, 0, 3, 5, 0, loc)
	sample := Profile{
		Avatar:    "https://pbs.twimg.com/profile_images/1222218816484020224/ik9P1QZt_normal.jpg",
		Banner:    "",
		Biography: "private account",
		//	Birthday:   "March 21",
		IsPrivate:      true,
		IsVerified:     false,
		Joined:         &joined,
		Location:       "",
		Name:           "private account",
		PinnedTweetIDs: []string{},
		URL:            "https://twitter.com/tomdumont",
		UserID:         "1221221876849995777",
		Username:       "tomdumont",
		Website:        "",
	}

	// some random private profile (found via google)
	profile, err := GetProfile("tomdumont")
	if err != nil {
		t.Error(err)
	}

	cmpOptions := cmp.Options{
		cmpopts.IgnoreFields(Profile{}, "FollowersCount"),
		cmpopts.IgnoreFields(Profile{}, "FollowingCount"),
		cmpopts.IgnoreFields(Profile{}, "FriendsCount"),
		cmpopts.IgnoreFields(Profile{}, "LikesCount"),
		cmpopts.IgnoreFields(Profile{}, "ListedCount"),
		cmpopts.IgnoreFields(Profile{}, "TweetsCount"),
	}
	if diff := cmp.Diff(sample, profile, cmpOptions...); diff != "" {
		t.Error("Resulting profile does not match the sample", diff)
	}

	if profile.FollowingCount == 0 {
		t.Error("Expected FollowingCount is greater than zero")
	}
	if profile.LikesCount == 0 {
		t.Error("Expected LikesCount is greater than zero")
	}
	if profile.TweetsCount == 0 {
		t.Error("Expected TweetsCount is greater than zero")
	}
}

func TestGetProfileErrorSuspended(t *testing.T) {
	_, err := GetProfile("123")
	if err == nil {
		t.Error("Expected Error, got success")
	} else {
		if err.Error() != "_Missing: User not found." {
			t.Errorf("Expected error '_Missing: User not found.', got '%s'", err)
		}
	}
}

func TestGetProfileErrorNotFound(t *testing.T) {
	neUser := "sample3123131"
	expectedError := fmt.Sprintf("User '%s' not found", neUser)
	_, err := GetProfile(neUser)
	if err == nil {
		t.Error("Expected Error, got success")
	} else {
		if err.Error() != expectedError {
			t.Errorf("Expected error '%s', got '%s'", expectedError, err)
		}
	}
}

func TestGetUserIDByScreenName(t *testing.T) {
	scraper := New()
	userID, err := scraper.GetUserIDByScreenName("Twitter")
	if err != nil {
		t.Errorf("getUserByScreenName() error = %v", err)
	}
	if userID == "" {
		t.Error("Expected non-empty user ID")
	}
}
